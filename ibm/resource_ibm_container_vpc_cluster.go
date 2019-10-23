package ibm

import (
	"fmt"
	"log"
	"strings"
	"time"

	v2 "github.com/IBM-Cloud/bluemix-go/api/container/containerv2"
	"github.com/IBM-Cloud/bluemix-go/api/resource/resourcev1/management"
	"github.com/IBM-Cloud/bluemix-go/bmxerror"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
)

const (
	deployRequested  = "Deploy requested"
	deployInProgress = "Deploy in progress"
	ready            = "Ready"
)

func resourceIBMContainerVpcCluster() *schema.Resource {
	return &schema.Resource{
		Create:   resourceIBMContainerVpcClusterCreate,
		Read:     resourceIBMContainerVpcClusterRead,
		Delete:   resourceIBMContainerVpcClusterDelete,
		Exists:   resourceIBMContainerVpcClusterExists,
		Importer: &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{

			"flavor": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
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
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
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
			},

			"service_subnet": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "172.21.0.0/16",
				ForceNew:    true,
				Description: "Custom subnet CIDR to provide private IP addresses for services",
			},

			"pod_subnet": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "172.30.0.0/16",
				ForceNew:    true,
				Description: "Custom subnet CIDR to provide private IP addresses for pods",
			},

			"worker_count": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1,
				ForceNew: true,
			},

			"disable_public_service_endpoint": {
				Type:     schema.TypeBool,
				Optional: true,
				ForceNew: true,
				Default:  true,
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

			"created_date": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"vpcs": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "list of VPCs",
			},

			"worker_zones": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "list of zones",
			},

			"ingress": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"hostname": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"secret": {
							Type:      schema.TypeString,
							Computed:  true,
							Sensitive: true,
						},
					},
				},
			},

			"master_status": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"master_status_modified_date": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"master_health": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"master_state": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"modified_date": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"owner_email": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"resource_group_id": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Computed:    true,
				Description: "ID of the resource group.",
			},

			"account_guid": {
				Description: "The bluemix account guid this cluster belongs to",
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Deprecated:  "This field is deprecated",
			},

			"crn": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "CRN of resource instance",
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

	var zonesList = make([]v2.Zone, 0)

	if res, ok := d.GetOk("zones"); ok {
		zones := res.([]interface{})
		for _, e := range zones {
			r, _ := e.(map[string]interface{})
			if ID, subnetID := r["id"], r["subnet_id"]; ID != nil && subnetID != nil {
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
	}

	targetEnv, err := getVpcClusterTargetHeader(d, meta)
	if err != nil {
		return err
	}

	cls, err := csClient.Clusters().Create(params, targetEnv)

	log.Println("Cluster creation response: ", cls)
	if err != nil {
		return err
	}
	d.SetId(cls.ID)
	_, err = waitForVpcClusterCreate(d, meta)
	if err != nil {
		return err
	}
	return resourceIBMContainerVpcClusterRead(d, meta)
}

func resourceIBMContainerVpcClusterRead(d *schema.ResourceData, meta interface{}) error {

	csClient, err := meta.(ClientSession).VpcContainerAPI()
	if err != nil {
		return err
	}

	targetEnv, err := getVpcClusterTargetHeader(d, meta)
	if err != nil {
		return err
	}

	clusterID := d.Id()
	cls, err := csClient.Clusters().GetCluster(clusterID, targetEnv)
	if err != nil {
		return fmt.Errorf("Error retrieving conatiner vpc cluster: %s", err)
	}

	d.Set("id", cls.ID)
	//mandatory fields[zone]
	d.Set("name", cls.Name)
	d.Set("created_date", cls.CreatedDate)
	d.Set("crn", cls.CRN)
	d.Set("datacenter", cls.DataCenter)
	d.Set("disable_auto_update", cls.DisableAutoUpdate)
	d.Set("master_kube_version", cls.MasterKubeVersion)
	d.Set("master_status", cls.Lifecycle.MasterStatus)
	d.Set("master_status_modified_date", cls.Lifecycle.MasterStatusModifiedDate)
	d.Set("master_health", cls.Lifecycle.MasterHealth)
	d.Set("master_state", cls.Lifecycle.MasterState)
	d.Set("modified_date", cls.Lifecycle.ModifiedDate)
	d.Set("service_subnet", cls.ServiceSubnet)
	d.Set("pod_subnet", cls.PodSubnet)
	d.Set("state", cls.State)
	d.Set("worker_zones", flattenStringList(cls.WorkerZones))
	d.Set("region", cls.Region)
	d.Set("resource_group_id", cls.ResourceGroupID)
	d.Set("state", cls.State)
	d.Set("vpcs", flattenStringList(cls.Vpcs))

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
	rsMangClient, err := meta.(ClientSession).ResourceManagementAPI()
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
	err = csClient.Clusters().Delete(clusterID, targetEnv)
	if err != nil {
		return fmt.Errorf("Error deleting cluster: %s", err)
	}
	_, err = waitForVpcClusterDelete(d, meta)
	if err != nil {
		return err
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

func waitForVpcClusterCreate(d *schema.ResourceData, meta interface{}) (interface{}, error) {
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
			clusterInfo, err := csClient.Clusters().GetCluster(clusterID, targetEnv)
			if err != nil {
				return clusterInfo, deployInProgress, nil
			}
			if (clusterInfo.Lifecycle == v2.LifeCycleInfo{}) {
				return clusterInfo, deployInProgress, nil
			}
			log.Println("Master Node Status:", clusterInfo.Lifecycle.MasterStatus)
			log.Println("Checking cluster state", strings.Compare(clusterInfo.Lifecycle.MasterStatus, ready))
			if strings.Compare(clusterInfo.Lifecycle.MasterStatus, ready) != 0 {
				return clusterInfo, deployInProgress, nil
			}
			return clusterInfo, ready, nil
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
	accountGUID := d.Get("account_guid").(string)

	sess, err := meta.(ClientSession).BluemixSession()
	if err != nil {
		return v2.ClusterTargetHeader{}, err
	}

	if resourceGroup == "" {
		resourceGroup = sess.Config.ResourceGroup

		if resourceGroup == "" {
			rsMangClient, err := meta.(ClientSession).ResourceManagementAPI()
			if err != nil {
				return v2.ClusterTargetHeader{}, err
			}
			resourceGroupQuery := management.ResourceGroupQuery{
				Default: true,
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
		AccountID:     accountGUID,
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
