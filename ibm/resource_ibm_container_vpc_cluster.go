package ibm

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	v1 "github.com/IBM-Cloud/bluemix-go/api/container/containerv1"
	v2 "github.com/IBM-Cloud/bluemix-go/api/container/containerv2"
	"github.com/IBM-Cloud/bluemix-go/api/resource/resourcev2/managementv2"
	"github.com/IBM-Cloud/bluemix-go/bmxerror"
	"github.com/hashicorp/terraform-plugin-sdk/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.ibm.com/Bluemix/riaas-go-client/clients/lbaas"
	iserrors "github.ibm.com/Bluemix/riaas-go-client/errors"
	"github.ibm.com/Bluemix/riaas-go-client/session"
)

const (
	deployRequested    = "Deploy requested"
	deployInProgress   = "Deploy in progress"
	ready              = "Ready"
	normal             = "normal"
	masterNodeReady    = "MasterNodeReady"
	oneWorkerNodeReady = "OneWorkerNodeReady"
	ingressReady       = "IngressReady"
)

func resourceIBMContainerVpcCluster() *schema.Resource {
	return &schema.Resource{
		Create:   resourceIBMContainerVpcClusterCreate,
		Read:     resourceIBMContainerVpcClusterRead,
		Update:   resourceIBMContainerVpcClusterUpdate,
		Delete:   resourceIBMContainerVpcClusterDelete,
		Exists:   resourceIBMContainerVpcClusterExists,
		Importer: &schema.ResourceImporter{},

		CustomizeDiff: customdiff.Sequence(
			func(diff *schema.ResourceDiff, v interface{}) error {
				return resourceTagsCustomizeDiff(diff)
			},
		),

		Schema: map[string]*schema.Schema{

			"flavor": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Cluster nodes flavour",
			},

			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The cluster name",
			},

			"vpc_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The vpc id where the cluster is",
			},

			"zones": {
				Type:        schema.TypeList,
				Required:    true,
				ForceNew:    true,
				Description: "Zone info",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Required:    true,
							ForceNew:    true,
							Description: "Zone for the worker pool in a multizone cluster",
						},

						"subnet_id": {
							Type:        schema.TypeString,
							Required:    true,
							ForceNew:    true,
							Description: "The VPC subnet to assign the cluster",
						},
					},
				},
			},
			//Optionals in cluster creation

			"kube_version": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
				DiffSuppressFunc: func(k, o, n string, d *schema.ResourceData) bool {
					if o == "" {
						return false
					}
					new := strings.Split(n, ".")
					old := strings.Split(o, ".")

					if strings.Compare(new[0]+"."+strings.Split(new[1], "_")[0], old[0]+"."+strings.Split(old[1], "_")[0]) == 0 {
						return true
					}
					return false
				},
				Description: "Kubernetes version",
			},

			"service_subnet": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Custom subnet CIDR to provide private IP addresses for services",
				Computed:    true,
			},

			"pod_subnet": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Custom subnet CIDR to provide private IP addresses for pods",
				Computed:    true,
			},

			"worker_count": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     1,
				Description: "Number of worker nodes in the cluster",
			},

			"disable_public_service_endpoint": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Boolean value true if Public service endpoint to be disabled",
			},

			"tags": {
				Type:        schema.TypeSet,
				Optional:    true,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Set:         resourceIBMVPCHash,
				Description: "List of tags for the resources",
			},

			"wait_till": {
				Type:             schema.TypeString,
				Optional:         true,
				Default:          ingressReady,
				DiffSuppressFunc: applyOnce,
				ValidateFunc:     validation.StringInSlice([]string{masterNodeReady, oneWorkerNodeReady, ingressReady}, true),
				Description:      "wait_till can be configured for Master Ready, One worker Ready or Ingress Ready",
			},

			"entitlement": {
				Type:             schema.TypeString,
				Optional:         true,
				DiffSuppressFunc: applyOnce,
				Description:      "Entitlement option reduces additional OCP Licence cost in Openshift Clusters",
			},

			"cos_instance_crn": {
				Type:             schema.TypeString,
				Optional:         true,
				DiffSuppressFunc: applyOnce,
				Description:      "A standard cloud object storage instance CRN to back up the internal registry in your OpenShift on VPC Gen 2 cluster",
			},

			ResourceControllerURL: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The URL of the IBM Cloud dashboard that can be used to explore and view details about this cluster",
			},

			//Get Cluster info Request
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"master_status": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"resource_group_id": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Computed:    true,
				Description: "ID of the resource group.",
			},

			"master_url": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"albs": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"alb_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"enable": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"state": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"load_balancer_hostname": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"resize": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"disable_deployment": {
							Type:     schema.TypeBool,
							Computed: true,
						},
					},
				},
			},
			"public_service_endpoint_url": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"private_service_endpoint_url": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"crn": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "CRN of resource instance",
			},

			"ingress_hostname": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ingress_secret": {
				Type:      schema.TypeString,
				Computed:  true,
				Sensitive: true,
			},

			ResourceName: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The name of the resource",
			},

			ResourceCRN: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The crn of the resource",
			},

			ResourceStatus: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The status of the resource",
			},

			ResourceGroupName: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The resource group name in which resource is provisioned",
			},
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(90 * time.Minute),
			Delete: schema.DefaultTimeout(45 * time.Minute),
		},
	}
}

func resourceIBMContainerVpcClusterCreate(d *schema.ResourceData, meta interface{}) error {

	sess, err := meta.(ClientSession).ISSession()
	if err != nil {
		return err
	}

	var vpcProvider string
	vpcProvider = "vpc-classic"
	if sess.Generation == 2 {
		vpcProvider = "vpc-gen2"
	}

	csClient, err := meta.(ClientSession).VpcContainerAPI()
	if err != nil {
		return err
	}

	disablePublicServiceEndpoint := d.Get("disable_public_service_endpoint").(bool)
	name := d.Get("name").(string)
	var kubeVersion string
	if v, ok := d.GetOk("kube_version"); ok {
		kubeVersion = v.(string)
	}
	podSubnet := d.Get("pod_subnet").(string)
	serviceSubnet := d.Get("service_subnet").(string)
	vpcID := d.Get("vpc_id").(string)
	flavor := d.Get("flavor").(string)
	workerCount := d.Get("worker_count").(int)

	// timeoutStage will define the timeout stage
	var timeoutStage string
	if v, ok := d.GetOk("wait_till"); ok {
		timeoutStage = v.(string)
	}

	var zonesList = make([]v2.Zone, 0)

	if res, ok := d.GetOk("zones"); ok {
		zones := res.([]interface{})
		for _, e := range zones {
			r, _ := e.(map[string]interface{})
			if ID, subnetID := r["name"], r["subnet_id"]; ID != nil && subnetID != nil {
				zoneParam := v2.Zone{}
				zoneParam.ID, zoneParam.SubnetID = ID.(string), subnetID.(string)
				zonesList = append(zonesList, zoneParam)
			}

		}
	}

	workerpool := v2.WorkerPoolConfig{
		VpcID:       vpcID,
		Flavor:      flavor,
		WorkerCount: workerCount,
		Zones:       zonesList,
	}

	params := v2.ClusterCreateRequest{
		DisablePublicServiceEndpoint: disablePublicServiceEndpoint,
		Name:                         name,
		KubeVersion:                  kubeVersion,
		PodSubnet:                    podSubnet,
		ServiceSubnet:                serviceSubnet,
		WorkerPools:                  workerpool,
		Provider:                     vpcProvider,
	}

	// Update params with Entitlement option if provided
	if v, ok := d.GetOk("entitlement"); ok {
		params.DefaultWorkerPoolEntitlement = v.(string)
	}

	// Update params with Cloud Object Store instance CRN id option if provided
	if v, ok := d.GetOk("cos_instance_crn"); ok {
		params.CosInstanceCRN = v.(string)
	}

	targetEnv, err := getVpcClusterTargetHeader(d, meta)
	if err != nil {
		return err
	}

	cls, err := csClient.Clusters().Create(params, targetEnv)

	if err != nil {
		return err
	}
	d.SetId(cls.ID)
	switch strings.ToLower(timeoutStage) {

	case strings.ToLower(masterNodeReady):
		_, err = waitForVpcClusterMasterAvailable(d, meta)
		if err != nil {
			return err
		}

	case strings.ToLower(oneWorkerNodeReady):
		_, err = waitForVpcClusterOneWorkerAvailable(d, meta)
		if err != nil {
			return err
		}

	case strings.ToLower(ingressReady):
		_, err = waitForVpcClusterIngressAvailable(d, meta)
		if err != nil {
			return err
		}

	}
	return resourceIBMContainerVpcClusterUpdate(d, meta)

}

func resourceIBMContainerVpcClusterUpdate(d *schema.ResourceData, meta interface{}) error {

	csClient, err := meta.(ClientSession).VpcContainerAPI()
	if err != nil {
		return err
	}

	targetEnv, err := getVpcClusterTargetHeader(d, meta)
	if err != nil {
		return err
	}

	clusterID := d.Id()

	v := os.Getenv("IC_ENV_TAGS")
	if d.HasChange("tags") || v != "" {
		oldList, newList := d.GetChange("tags")
		cluster, err := csClient.Clusters().GetCluster(clusterID, targetEnv)
		if err != nil {
			return fmt.Errorf("Error retrieving cluster %s: %s", clusterID, err)
		}
		err = UpdateTagsUsingCRN(oldList, newList, meta, cluster.CRN)
		if err != nil {
			log.Printf(
				"An error occured during update of instance (%s) tags: %s", clusterID, err)
		}
	}

	if d.HasChange("kube_version") && !d.IsNewResource() {
		ClusterClient, err := meta.(ClientSession).ContainerAPI()
		if err != nil {
			return err
		}
		var masterVersion string
		if v, ok := d.GetOk("kube_version"); ok {
			masterVersion = v.(string)
		}
		params := v1.ClusterUpdateParam{
			Action:  "update",
			Force:   true,
			Version: masterVersion,
		}

		Env, err := getClusterTargetHeader(d, meta)

		if err != nil {
			return err
		}
		Error := ClusterClient.Clusters().Update(clusterID, params, Env)
		if Error != nil {
			return Error
		}
		_, err = WaitForVpcClusterVersionUpdate(d, meta, targetEnv)
		if err != nil {
			return fmt.Errorf(
				"Error waiting for cluster (%s) version to be updated: %s", d.Id(), err)
		}
		// Update the worker nodes after master node kube-version is updated.

		workers, err := csClient.Workers().ListWorkers(clusterID, false, targetEnv)
		if err != nil {
			return fmt.Errorf("Error retrieving workers for cluster: %s", err)
		}
		for _, worker := range workers {
			_, err := csClient.Workers().ReplaceWokerNode(clusterID, worker.ID, targetEnv)
			// As API returns http response 204 NO CONTENT, error raised will be exempted.
			if err != nil && !strings.Contains(err.Error(), "EmptyResponseBody") {
				return fmt.Errorf("Error replacing the worker node from the cluster: %s", err)
			}
		}
	}

	if d.HasChange("worker_count") && !d.IsNewResource() {
		count := d.Get("worker_count").(int)
		ClusterClient, err := meta.(ClientSession).ContainerAPI()
		if err != nil {
			return err
		}
		Env := v1.ClusterTargetHeader{ResourceGroup: targetEnv.ResourceGroup}

		err = ClusterClient.WorkerPools().ResizeWorkerPool(clusterID, "default", count, Env)
		if err != nil {
			return fmt.Errorf(
				"Error updating the worker_count %d: %s", count, err)
		}
	}

	return resourceIBMContainerVpcClusterRead(d, meta)
}

func resourceIBMContainerVpcClusterRead(d *schema.ResourceData, meta interface{}) error {

	csClient, err := meta.(ClientSession).VpcContainerAPI()
	if err != nil {
		return err
	}
	albsAPI := csClient.Albs()

	targetEnv, err := getVpcClusterTargetHeader(d, meta)
	if err != nil {
		return err
	}

	clusterID := d.Id()
	cls, err := csClient.Clusters().GetCluster(clusterID, targetEnv)
	if err != nil {
		return fmt.Errorf("Error retrieving conatiner vpc cluster: %s", err)
	}

	workerPool, err := csClient.WorkerPools().GetWorkerPool(clusterID, "default", targetEnv)

	var zones = make([]map[string]interface{}, 0)
	for _, zone := range workerPool.Zones {
		for _, subnet := range zone.Subnets {
			if subnet.Primary == true {
				zoneInfo := map[string]interface{}{
					"name":      zone.ID,
					"subnet_id": subnet.ID,
				}
				zones = append(zones, zoneInfo)
			}
		}
	}

	albs, err := albsAPI.ListClusterAlbs(clusterID, targetEnv)
	if err != nil && !strings.Contains(err.Error(), "This operation is not supported for your cluster's version.") {
		return fmt.Errorf("Error retrieving alb's of the cluster %s: %s", clusterID, err)
	}

	d.Set("name", cls.Name)
	d.Set("crn", cls.CRN)
	d.Set("disable_auto_update", cls.DisableAutoUpdate)
	d.Set("master_status", cls.Lifecycle.MasterStatus)
	d.Set("zones", zones)
	if strings.HasSuffix(cls.MasterKubeVersion, "_openshift") {
		d.Set("kube_version", strings.Split(cls.MasterKubeVersion, "_")[0]+"_openshift")
	} else {
		d.Set("kube_version", strings.Split(cls.MasterKubeVersion, "_")[0])

	}
	d.Set("worker_count", workerPool.WorkerCount)
	d.Set("vpc_id", cls.Vpcs[0])
	d.Set("master_url", cls.MasterURL)
	d.Set("flavor", workerPool.Flavor)
	d.Set("service_subnet", cls.ServiceSubnet)
	d.Set("pod_subnet", cls.PodSubnet)
	d.Set("state", cls.State)
	d.Set("region", cls.Region)
	d.Set("ingress_hostname", cls.Ingress.HostName)
	d.Set("ingress_secret", cls.Ingress.SecretName)
	d.Set("albs", flattenVpcAlbs(albs, "all"))
	d.Set("resource_group_id", cls.ResourceGroupID)
	d.Set("public_service_endpoint_url", cls.ServiceEndpoints.PublicServiceEndpointURL)
	d.Set("private_service_endpoint_url", cls.ServiceEndpoints.PrivateServiceEndpointURL)
	if cls.ServiceEndpoints.PublicServiceEndpointURL != "" {
		d.Set("disable_public_service_endpoint", false)
	} else {
		d.Set("disable_public_service_endpoint", true)
	}

	tags, err := GetTagsUsingCRN(meta, cls.CRN)
	if err != nil {
		log.Printf(
			"An error occured during reading of instance (%s) tags : %s", d.Id(), err)
	}
	d.Set("tags", tags)
	controller, err := getBaseController(meta)
	if err != nil {
		return err
	}
	d.Set(ResourceControllerURL, controller+"/kubernetes/clusters")
	d.Set(ResourceName, cls.Name)
	d.Set(ResourceCRN, cls.CRN)
	d.Set(ResourceStatus, cls.State)
	rsMangClient, err := meta.(ClientSession).ResourceManagementAPIv2()
	if err != nil {
		return err
	}
	grp, err := rsMangClient.ResourceGroup().Get(cls.ResourceGroupID)
	if err != nil {
		return err
	}
	d.Set(ResourceGroupName, grp.Name)

	return nil
}

func resourceIBMContainerVpcClusterDelete(d *schema.ResourceData, meta interface{}) error {
	targetEnv, err := getVpcClusterTargetHeader(d, meta)
	if err != nil {
		return err
	}
	csClient, err := meta.(ClientSession).VpcContainerAPI()
	if err != nil {
		return err
	}
	clusterID := d.Id()

	var zonesList = make([]v2.Zone, 0)

	if res, ok := d.GetOk("zones"); ok {
		zones := res.([]interface{})
		for _, e := range zones {
			r, _ := e.(map[string]interface{})
			if ID, subnetID := r["name"], r["subnet_id"]; ID != nil && subnetID != nil {
				zoneParam := v2.Zone{}
				zoneParam.ID, zoneParam.SubnetID = ID.(string), subnetID.(string)
				zonesList = append(zonesList, zoneParam)
			}

		}
	}
	splitZone := strings.Split(zonesList[0].ID, "-")
	region := splitZone[0] + "-" + splitZone[1]

	sess1, err := meta.(ClientSession).ISSession()
	if err != nil {
		log.Println("error creating ISsession", err)
	}

	newSess, err := session.New(sess1.IAMToken, region, int(sess1.Generation), false, 10*time.Minute)
	if err != nil {
		log.Println("error creating new ISsession", err)
	}
	client := lbaas.NewLoadBalancerClient(newSess)

	lbs, err := client.List()
	if err != nil {
		log.Println("cannot retreive load balancers=", err)
	}

	err = csClient.Clusters().Delete(clusterID, targetEnv)
	if err != nil {
		return fmt.Errorf("Error deleting cluster: %s", err)
	}
	_, err = waitForVpcClusterDelete(d, meta)
	if err != nil {
		return err
	}

	if len(lbs) > 0 {
		for _, lb := range lbs {
			if strings.Contains(lb.Name, clusterID) {
				log.Println("Deleting Load Balancer", lb.Name)
				client := lbaas.NewLoadBalancerClient(newSess)
				err = client.Delete(string(lb.ID))
				if err != nil {
					log.Println("error deleting Load Balancer", err)
				}

				_, err = isWaitForDeleted(client, string(lb.ID), d.Timeout(schema.TimeoutDelete))
				if err != nil {
					log.Println("error waiting for Load Balancer to be deleted", err)
				}
			}
		}

	}

	return nil
}

func waitForVpcClusterDelete(d *schema.ResourceData, meta interface{}) (interface{}, error) {
	targetEnv, err := getVpcClusterTargetHeader(d, meta)
	if err != nil {
		return nil, err
	}
	csClient, err := meta.(ClientSession).VpcContainerAPI()
	if err != nil {
		return nil, err
	}
	clusterID := d.Id()
	deleteStateConf := &resource.StateChangeConf{
		Pending: []string{clusterDeletePending},
		Target:  []string{clusterDeleted},
		Refresh: func() (interface{}, string, error) {
			cluster, err := csClient.Clusters().GetCluster(clusterID, targetEnv)
			if err != nil {
				if apiErr, ok := err.(bmxerror.RequestFailure); ok && (apiErr.StatusCode() == 404) {
					return cluster, clusterDeleted, nil
				}
				return nil, "", err
			}
			return cluster, clusterDeletePending, nil
		},
		Timeout:      d.Timeout(schema.TimeoutDelete),
		Delay:        10 * time.Second,
		MinTimeout:   5 * time.Second,
		PollInterval: 5 * time.Second,
	}

	return deleteStateConf.WaitForState()
}

func waitForVpcClusterOneWorkerAvailable(d *schema.ResourceData, meta interface{}) (interface{}, error) {
	targetEnv, err := getVpcClusterTargetHeader(d, meta)
	if err != nil {
		return nil, err
	}
	csClient, err := meta.(ClientSession).VpcContainerAPI()
	if err != nil {
		return nil, err
	}
	clusterID := d.Id()
	createStateConf := &resource.StateChangeConf{
		Pending: []string{deployRequested, deployInProgress},
		Target:  []string{normal},
		Refresh: func() (interface{}, string, error) {
			workers, err := csClient.Workers().ListByWorkerPool(clusterID, "default", false, targetEnv)
			if err != nil {
				return workers, deployInProgress, err
			}
			if len(workers) == 0 {
				return workers, deployInProgress, nil
			}

			for _, worker := range workers {
				log.Println("worker: ", worker.ID)
				log.Println("worker health state:  ", worker.Health.State)

				if worker.Health.State == normal {
					return workers, normal, nil
				}
			}
			return workers, deployInProgress, nil

		},
		Timeout:                   d.Timeout(schema.TimeoutCreate),
		Delay:                     10 * time.Second,
		MinTimeout:                5 * time.Second,
		ContinuousTargetOccurence: 5,
	}
	return createStateConf.WaitForState()
}

func waitForVpcClusterMasterAvailable(d *schema.ResourceData, meta interface{}) (interface{}, error) {
	targetEnv, err := getVpcClusterTargetHeader(d, meta)
	if err != nil {
		return nil, err
	}
	csClient, err := meta.(ClientSession).VpcContainerAPI()
	if err != nil {
		return nil, err
	}
	clusterID := d.Id()
	createStateConf := &resource.StateChangeConf{
		Pending: []string{deployRequested, deployInProgress},
		Target:  []string{ready},
		Refresh: func() (interface{}, string, error) {
			clusterInfo, clusterInfoErr := csClient.Clusters().GetCluster(clusterID, targetEnv)

			if err != nil || clusterInfoErr != nil {
				return clusterInfo, deployInProgress, err
			}

			if clusterInfo.Lifecycle.MasterStatus == ready {
				return clusterInfo, ready, nil
			}
			return clusterInfo, deployInProgress, nil

		},
		Timeout:                   d.Timeout(schema.TimeoutCreate),
		Delay:                     10 * time.Second,
		MinTimeout:                5 * time.Second,
		ContinuousTargetOccurence: 5,
	}
	return createStateConf.WaitForState()
}

func waitForVpcClusterIngressAvailable(d *schema.ResourceData, meta interface{}) (interface{}, error) {
	targetEnv, err := getVpcClusterTargetHeader(d, meta)
	if err != nil {
		return nil, err
	}
	csClient, err := meta.(ClientSession).VpcContainerAPI()
	if err != nil {
		return nil, err
	}
	clusterID := d.Id()
	createStateConf := &resource.StateChangeConf{
		Pending: []string{deployRequested, deployInProgress},
		Target:  []string{ready},
		Refresh: func() (interface{}, string, error) {
			clusterInfo, clusterInfoErr := csClient.Clusters().GetCluster(clusterID, targetEnv)

			if err != nil || clusterInfoErr != nil {
				return clusterInfo, deployInProgress, err
			}

			if clusterInfo.Ingress.HostName != "" {
				return clusterInfo, ready, nil
			}
			return clusterInfo, deployInProgress, nil

		},
		Timeout:                   d.Timeout(schema.TimeoutCreate),
		Delay:                     10 * time.Second,
		MinTimeout:                5 * time.Second,
		ContinuousTargetOccurence: 5,
	}
	return createStateConf.WaitForState()
}

func getVpcClusterTargetHeader(d *schema.ResourceData, meta interface{}) (v2.ClusterTargetHeader, error) {

	resourceGroup := d.Get("resource_group_id").(string)

	sess, err := meta.(ClientSession).BluemixSession()
	if err != nil {
		return v2.ClusterTargetHeader{}, err
	}
	userDetails, err := meta.(ClientSession).BluemixUserDetails()
	if err != nil {
		return v2.ClusterTargetHeader{}, err
	}
	accountID := userDetails.userAccount

	if resourceGroup == "" {
		resourceGroup = sess.Config.ResourceGroup

		if resourceGroup == "" {
			rsMangClient, err := meta.(ClientSession).ResourceManagementAPIv2()
			if err != nil {
				return v2.ClusterTargetHeader{}, err
			}
			resourceGroupQuery := managementv2.ResourceGroupQuery{
				Default:   true,
				AccountID: accountID,
			}
			grpList, err := rsMangClient.ResourceGroup().List(&resourceGroupQuery)
			if err != nil {
				return v2.ClusterTargetHeader{}, err
			}
			if len(grpList) <= 0 {
				return v2.ClusterTargetHeader{}, fmt.Errorf("the targeted resource group could not be found. Make sure you have required permissions to access the resource group")
			}
			resourceGroup = grpList[0].ID
		}
	}

	targetEnv := v2.ClusterTargetHeader{
		ResourceGroup: resourceGroup,
	}
	return targetEnv, nil
}

func resourceIBMContainerVpcClusterExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	csClient, err := meta.(ClientSession).VpcContainerAPI()
	if err != nil {
		return false, err
	}
	targetEnv, err := getVpcClusterTargetHeader(d, meta)
	if err != nil {
		return false, err
	}
	clusterID := d.Id()
	cls, err := csClient.Clusters().GetCluster(clusterID, targetEnv)
	if err != nil {
		if apiErr, ok := err.(bmxerror.RequestFailure); ok {
			if apiErr.StatusCode() == 404 {
				return false, nil
			}
		}
		return false, fmt.Errorf("Error communicating with the API: %s", err)
	}
	return cls.ID == clusterID, nil
}

// WaitForVpcClusterVersionUpdate Waits for cluster creation
func WaitForVpcClusterVersionUpdate(d *schema.ResourceData, meta interface{}, target v2.ClusterTargetHeader) (interface{}, error) {
	csClient, err := meta.(ClientSession).VpcContainerAPI()
	if err != nil {
		return nil, err
	}
	log.Printf("Waiting for cluster (%s) version to be updated.", d.Id())
	id := d.Id()

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"retry", versionUpdating},
		Target:     []string{clusterNormal},
		Refresh:    vpcClusterVersionRefreshFunc(csClient.Clusters(), id, d, target),
		Timeout:    d.Timeout(schema.TimeoutUpdate),
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	return stateConf.WaitForState()
}

func vpcClusterVersionRefreshFunc(client v2.Clusters, instanceID string, d *schema.ResourceData, target v2.ClusterTargetHeader) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		cls, err := client.GetCluster(instanceID, target)
		if err != nil {
			return nil, "retry", fmt.Errorf("Error retrieving conatiner vpc cluster: %s", err)
		}

		// Check active transactions
		log.Println("Checking cluster version", cls.MasterKubeVersion, d.Get("kube_version").(string))
		if strings.Contains(cls.MasterKubeVersion, "pending") {
			return cls, versionUpdating, nil
		}
		return cls, clusterNormal, nil
	}
}

func isWaitForDeleted(lbc *lbaas.LoadBalancerClient, id string, timeout time.Duration) (interface{}, error) {
	log.Printf("Waiting for  (%s) to be deleted.", id)

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"retry", isLBDeleting},
		Target:     []string{},
		Refresh:    isDeleteRefreshFunc(lbc, id),
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	return stateConf.WaitForState()
}

func isDeleteRefreshFunc(lbc *lbaas.LoadBalancerClient, id string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		lb, err := lbc.Get(id)
		if err == nil {
			return lb, isLBDeleting, nil
		}

		iserror, ok := err.(iserrors.RiaasError)
		if ok {
			log.Printf("[DEBUG] %s", iserror.Error())
			if len(iserror.Payload.Errors) == 1 &&
				iserror.Payload.Errors[0].Code == "load_balancer_not_found" {
				return nil, isLBDeleted, nil
			}
		}
		return nil, isLBDeleting, err
	}
}
