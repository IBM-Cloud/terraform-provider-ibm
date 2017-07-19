package ibm

import (
	"fmt"
	"strings"

	v1 "github.com/IBM-Bluemix/bluemix-go/api/container/containerv1"
	"github.com/IBM-Bluemix/bluemix-go/bmxerror"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceIBMContainerBindService() *schema.Resource {
	return &schema.Resource{
		Create:   resourceIBMContainerBindServiceCreate,
		Read:     resourceIBMContainerBindServiceRead,
		Delete:   resourceIBMContainerBindServiceDelete,
		Exists:   resourceIBMContainerBindServiceExists,
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
				Type:     schema.TypeString,
				Computed: true,
			},
			"org_guid": {
				Description: "The bluemix organization guid this cluster belongs to",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
			"space_guid": {
				Description: "The bluemix space guid this cluster belongs to",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
			"account_guid": {
				Description: "The bluemix account guid this cluster belongs to",
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
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

	d.SetId(bindResp.SecretName)
	d.Set("service_instance_name_id", serviceInstanceNameID)
	d.Set("cluster_name_id", clusterNameID)
	d.Set("namespace_id", namespaceID)
	d.Set("space_guid", serviceInstanceSpaceGUID)
	d.Set("secret_name", bindResp.SecretName)

	return resourceIBMContainerBindServiceRead(d, meta)
}

func resourceIBMContainerBindServiceRead(d *schema.ResourceData, meta interface{}) error {
	csClient, err := meta.(ClientSession).ContainerAPI()
	if err != nil {
		return err
	}
	clusterID := d.Get("cluster_name_id").(string)
	namespace := d.Get("namespace_id").(string)
	serviceInstanceNameID := d.Get("service_instance_name_id").(string)
	targetEnv := getClusterTargetHeader(d)
	_, err = csClient.Clusters().FindServiceBoundToCluster(clusterID, serviceInstanceNameID, namespace, targetEnv)
	if err != nil {
		return fmt.Errorf("Error finding service: %s", err)
	}
	return nil
}

func resourceIBMContainerBindServiceDelete(d *schema.ResourceData, meta interface{}) error {
	csClient, err := meta.(ClientSession).ContainerAPI()
	if err != nil {
		return err
	}
	clusterID := d.Get("cluster_name_id").(string)
	namespace := d.Get("namespace_id").(string)
	serviceInstanceNameID := d.Get("service_instance_name_id").(string)
	targetEnv := getClusterTargetHeader(d)

	err = csClient.Clusters().UnBindService(clusterID, namespace, serviceInstanceNameID, targetEnv)
	if err != nil {
		return fmt.Errorf("Error unbinding service: %s", err)
	}
	return nil
}

func resourceIBMContainerBindServiceExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	csClient, err := meta.(ClientSession).ContainerAPI()
	if err != nil {
		return false, err
	}
	clusterID := d.Get("cluster_name_id").(string)
	namespace := d.Get("namespace_id").(string)
	serviceInstanceNameID := d.Get("service_instance_name_id").(string)
	targetEnv := getClusterTargetHeader(d)

	service, err := csClient.Clusters().FindServiceBoundToCluster(clusterID, serviceInstanceNameID, namespace, targetEnv)
	if err != nil {
		if apiErr, ok := err.(bmxerror.RequestFailure); ok {
			if apiErr.StatusCode() == 404 {
				return false, nil
			}
		}
		return false, fmt.Errorf("Error communicating with the API: %s", err)
	}
	return (strings.Compare(serviceInstanceNameID, service.ServiceName) == 0 || strings.Compare(serviceInstanceNameID, service.ServiceID) == 0), nil
}
