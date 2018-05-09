package ibm

import (
	"fmt"

	v1 "github.com/IBM-Cloud/bluemix-go/api/container/containerv1"
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
				Required:    true,
			},
			"service_instance_name_id": {
				Type:     schema.TypeString,
				ForceNew: true,
				Required: true,
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
				Required:    true,
				ForceNew:    true,
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

func getClusterTargetHeader(d *schema.ResourceData) v1.ClusterTargetHeader {
	orgGUID := d.Get("org_guid").(string)
	spaceGUID := d.Get("space_guid").(string)
	accountGUID := d.Get("account_guid").(string)

	targetEnv := v1.ClusterTargetHeader{
		OrgID:     orgGUID,
		SpaceID:   spaceGUID,
		AccountID: accountGUID,
	}
	return targetEnv
}

func resourceIBMContainerBindServiceCreate(d *schema.ResourceData, meta interface{}) error {
	csClient, err := meta.(ClientSession).ContainerAPI()
	if err != nil {
		return err
	}
	clusterNameID := d.Get("cluster_name_id").(string)
	serviceInstanceSpaceGUID := d.Get("service_instance_space_guid").(string)
	serviceInstanceNameID := d.Get("service_instance_name_id").(string)
	namespaceID := d.Get("namespace_id").(string)

	bindService := v1.ServiceBindRequest{
		ClusterNameOrID:         clusterNameID,
		SpaceGUID:               serviceInstanceSpaceGUID,
		ServiceInstanceNameOrID: serviceInstanceNameID,
		NamespaceID:             namespaceID,
	}

	targetEnv := getClusterTargetHeader(d)
	bindResp, err := csClient.Clusters().BindService(bindService, targetEnv)
	if err != nil {
		return err
	}
	//Fix me Id would be typically the returned ID from the API, proabably SecretName should be used
	d.SetId(clusterNameID)
	d.Set("service_instance_name_id", serviceInstanceNameID)
	d.Set("namespace_id", namespaceID)
	d.Set("service_instance_space_guid", serviceInstanceSpaceGUID)
	d.Set("secret_name", bindResp.SecretName)

	return resourceIBMContainerBindServiceRead(d, meta)
}

func resourceIBMContainerBindServiceUpdate(d *schema.ResourceData, meta interface{}) error {
	//Only tags are updated and that too locally hence nothing to validate and update in terms of real API at this point
	return nil
}

func resourceIBMContainerBindServiceRead(d *schema.ResourceData, meta interface{}) error {
	//No API to read back the credentials so leave schema as it is
	return nil
}

func resourceIBMContainerBindServiceDelete(d *schema.ResourceData, meta interface{}) error {
	csClient, err := meta.(ClientSession).ContainerAPI()
	if err != nil {
		return err
	}
	clusterID := d.Id()
	namespace := d.Get("namespace_id").(string)
	serviceInstanceNameID := d.Get("service_instance_name_id").(string)
	targetEnv := getClusterTargetHeader(d)

	err = csClient.Clusters().UnBindService(clusterID, namespace, serviceInstanceNameID, targetEnv)
	if err != nil {
		return fmt.Errorf("Error unbinding service: %s", err)
	}
	return nil
}

//Pure Aramda API not available, we can still find by using k8s api
/*
func resourceIBMContainerBindServiceExists(d *schema.ResourceData, meta interface{}) (bool, error) {

}*/
