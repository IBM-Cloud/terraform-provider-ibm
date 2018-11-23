package ibm

import (
	"fmt"

	v1 "github.com/IBM-Cloud/bluemix-go/api/container/containerv1"
	"github.com/IBM-Cloud/bluemix-go/api/resource/resourcev1/management"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceIBMContainerBindService() *schema.Resource {
	return &schema.Resource{
		Create:   resourceIBMContainerBindServiceCreate,
		Read:     resourceIBMContainerBindServiceRead,
		Update:   resourceIBMContainerBindServiceUpdate,
		Delete:   resourceIBMContainerBindServiceDelete,
		Importer: &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"cluster_name_id": {
				Type:     schema.TypeString,
				ForceNew: true,
				Required: true,
			},
			"service_instance_space_guid": {
				Type:        schema.TypeString,
				Description: "The space guid the service instance belongs to",
				ForceNew:    true,
				Optional:    true,
				Removed:     "This field has been removed",
			},
			"service_instance_name_id": {
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
				Removed:  "This field has been removed. User service_instance_name or service_instance_id instead",
			},
			"service_instance_name": {
				Type:          schema.TypeString,
				ForceNew:      true,
				Optional:      true,
				Computed:      true,
				ConflictsWith: []string{"service_instance_id"},
			},
			"service_instance_id": {
				Type:          schema.TypeString,
				ForceNew:      true,
				Optional:      true,
				Computed:      true,
				ConflictsWith: []string{"service_instance_name"},
			},
			"namespace_id": {
				Type:     schema.TypeString,
				ForceNew: true,
				Required: true,
			},
			"secret_name": {
				Type:      schema.TypeString,
				Computed:  true,
				Sensitive: true,
				Removed:   "This field has been removed",
			},
			"org_guid": {
				Description: "The bluemix organization guid this cluster belongs to",
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
			},
			"space_guid": {
				Description: "The bluemix space guid this cluster belongs to",
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
			},
			"account_guid": {
				Description: "The bluemix account guid this cluster belongs to",
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
			},
			"region": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "The cluster region",
			},
			"resource_group_id": {
				Type:             schema.TypeString,
				Optional:         true,
				Description:      "ID of the resource group.",
				ForceNew:         true,
				DiffSuppressFunc: applyOnce,
			},
			"tags": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set:      schema.HashString,
			},
		},
	}
}

func getClusterTargetHeader(d *schema.ResourceData, meta interface{}) (v1.ClusterTargetHeader, error) {
	orgGUID := d.Get("org_guid").(string)
	spaceGUID := d.Get("space_guid").(string)
	accountGUID := d.Get("account_guid").(string)
	region := d.Get("region").(string)
	resourceGroup := d.Get("resource_group_id").(string)

	sess, err := meta.(ClientSession).BluemixSession()
	if err != nil {
		return v1.ClusterTargetHeader{}, err
	}

	if region == "" {
		region = sess.Config.Region
	}
	if resourceGroup == "" {
		resourceGroup = sess.Config.ResourceGroup

		if resourceGroup == "" {
			rsMangClient, err := meta.(ClientSession).ResourceManagementAPI()
			if err != nil {
				return v1.ClusterTargetHeader{}, err
			}
			resourceGroupQuery := management.ResourceGroupQuery{
				Default: true,
			}
			grpList, err := rsMangClient.ResourceGroup().List(&resourceGroupQuery)
			if err != nil {
				return v1.ClusterTargetHeader{}, err
			}
			resourceGroup = grpList[0].ID
		}
	}

	targetEnv := v1.ClusterTargetHeader{
		OrgID:         orgGUID,
		SpaceID:       spaceGUID,
		AccountID:     accountGUID,
		Region:        region,
		ResourceGroup: resourceGroup,
	}
	return targetEnv, nil
}

func resourceIBMContainerBindServiceCreate(d *schema.ResourceData, meta interface{}) error {
	csClient, err := meta.(ClientSession).ContainerAPI()
	if err != nil {
		return err
	}
	clusterNameID := d.Get("cluster_name_id").(string)
	namespaceID := d.Get("namespace_id").(string)
	var serviceInstanceNameID string
	if serviceInstanceName, ok := d.GetOk("service_instance_name"); ok {
		serviceInstanceNameID = serviceInstanceName.(string)
	} else if serviceInstanceID, ok := d.GetOk("service_instance_id"); ok {
		serviceInstanceNameID = serviceInstanceID.(string)
	} else {
		return fmt.Errorf("Please set either service_instance_name or service_instance_id")
	}
	bindService := v1.ServiceBindRequest{
		ClusterNameOrID:         clusterNameID,
		ServiceInstanceNameOrID: serviceInstanceNameID,
		NamespaceID:             namespaceID,
	}

	targetEnv, err := getClusterTargetHeader(d, meta)
	if err != nil {
		return err
	}
	_, err = csClient.Clusters().BindService(bindService, targetEnv)
	if err != nil {
		return err
	}
	d.SetId(fmt.Sprintf("%s/%s/%s", clusterNameID, serviceInstanceNameID, namespaceID))

	return resourceIBMContainerBindServiceRead(d, meta)
}

func resourceIBMContainerBindServiceUpdate(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceIBMContainerBindServiceRead(d *schema.ResourceData, meta interface{}) error {
	csClient, err := meta.(ClientSession).ContainerAPI()
	if err != nil {
		return err
	}
	parts, err := idParts(d.Id())
	if err != nil {
		return err
	}
	clusterNameID := parts[0]
	serviceInstanceNameID := parts[1]
	namespaceID := parts[2]

	targetEnv, err := getClusterTargetHeader(d, meta)
	if err != nil {
		return err
	}

	boundService, err := csClient.Clusters().FindServiceBoundToCluster(clusterNameID, serviceInstanceNameID, namespaceID, targetEnv)
	if err != nil {
		return err
	}
	d.Set("namespace_id", boundService.Namespace)

	d.Set("service_instance_name", boundService.ServiceName)
	d.Set("service_instance_id", boundService.ServiceID)
	//d.Set(key, boundService.ServiceKeyName)
	//d.Set(key, boundService.ServiceName)
	return nil
}

func resourceIBMContainerBindServiceDelete(d *schema.ResourceData, meta interface{}) error {
	csClient, err := meta.(ClientSession).ContainerAPI()
	if err != nil {
		return err
	}
	parts, err := idParts(d.Id())
	if err != nil {
		return err
	}
	clusterNameID := parts[0]
	serviceInstanceNameID := parts[1]
	namespace := parts[2]
	targetEnv, err := getClusterTargetHeader(d, meta)
	if err != nil {
		return err
	}

	err = csClient.Clusters().UnBindService(clusterNameID, namespace, serviceInstanceNameID, targetEnv)
	if err != nil {
		return fmt.Errorf("Error unbinding service: %s", err)
	}
	return nil
}

//Pure Aramda API not available, we can still find by using k8s api
/*
func resourceIBMContainerBindServiceExists(d *schema.ResourceData, meta interface{}) (bool, error) {

}*/
