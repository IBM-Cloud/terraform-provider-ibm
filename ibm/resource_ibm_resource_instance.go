package ibm

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform/flatmap"

	"github.com/IBM-Cloud/bluemix-go/api/resource/resourcev1/controller"
	"github.com/IBM-Cloud/bluemix-go/api/resource/resourcev1/management"
	"github.com/IBM-Cloud/bluemix-go/models"

	"github.com/IBM-Cloud/bluemix-go/bmxerror"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
)

const (
	rsInstanceSuccessStatus  = "active"
	rsInstanceProgressStatus = "in progress"
	rsInstanceInactiveStatus = "inactive"
	rsInstanceFailStatus     = "failed"
	rsInstanceRemovedStatus  = "removed"
)

func resourceIBMResourceInstance() *schema.Resource {
	return &schema.Resource{
		Create:   resourceIBMResourceInstanceCreate,
		Read:     resourceIBMResourceInstanceRead,
		Update:   resourceIBMResourceInstanceUpdate,
		Delete:   resourceIBMResourceInstanceDelete,
		Exists:   resourceIBMResourceInstanceExists,
		Importer: &schema.ResourceImporter{},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "A name for the resource instance",
			},

			"service": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The name of the service offering like cloud-object-storage, kms etc",
			},

			"plan": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The plan type of the service",
			},

			"location": {
				Description: "The location where the instance available",
				Required:    true,
				ForceNew:    true,
				Type:        schema.TypeString,
			},

			"resource_group_id": {
				Description:      "The resource group id",
				Optional:         true,
				ForceNew:         true,
				Type:             schema.TypeString,
				DiffSuppressFunc: applyOnce,
			},

			"parameters": {
				Type:        schema.TypeMap,
				Optional:    true,
				ForceNew:    true,
				Description: "Arbitrary parameters to pass. Must be a JSON object",
			},

			"tags": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set:      schema.HashString,
			},

			"status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Status of resource instance",
			},
		},
	}
}

func resourceIBMResourceInstanceCreate(d *schema.ResourceData, meta interface{}) error {
	rsConClient, err := meta.(ClientSession).ResourceControllerAPI()
	if err != nil {
		return err
	}
	serviceName := d.Get("service").(string)
	plan := d.Get("plan").(string)
	name := d.Get("name").(string)
	location := d.Get("location").(string)

	rsInst := controller.CreateServiceInstanceRequest{
		Name: name,
	}

	rsCatClient, err := meta.(ClientSession).ResourceCatalogAPI()
	if err != nil {
		return err
	}
	rsCatRepo := rsCatClient.ResourceCatalog()

	serviceOff, err := rsCatRepo.FindByName(serviceName, true)
	if err != nil {
		return fmt.Errorf("Error retrieving service offering: %s", err)
	}

	metadata := serviceOff[0].GetMetadata()
	if _, ok := metadata.(*models.ServiceResourceMetadata); !ok {
		return fmt.Errorf("Cannot create instance of resource %s\nUse 'ibm_service_instance' if the resource is a Cloud Foundry service", serviceName)
	}

	servicePlan, err := rsCatRepo.GetServicePlanID(serviceOff[0], plan)
	if err != nil {
		return fmt.Errorf("Error retrieving plan: %s", err)
	}
	rsInst.ServicePlanID = servicePlan

	deployments, err := rsCatRepo.ListDeployments(servicePlan)
	if err != nil {
		return fmt.Errorf("Error retrieving deployment for plan %s : %s", plan, err)
	}
	if len(deployments) == 0 {
		return fmt.Errorf("No deployment found for service plan : %s", plan)
	}
	deployments, supportedLocations := filterDeployments(deployments, location)

	if len(deployments) == 0 {
		locationList := make([]string, 0, len(supportedLocations))
		for l := range supportedLocations {
			locationList = append(locationList, l)
		}
		return fmt.Errorf("No deployment found for service plan %s at location %s.\nValid location(s) are: %q.\nUse 'ibm_service_instance' if the service is a Cloud Foundry service.", plan, location, locationList)
	}

	rsInst.TargetCrn = deployments[0].CatalogCRN

	if rsGrpID, ok := d.GetOk("resource_group_id"); ok {
		rsInst.ResourceGroupID = rsGrpID.(string)
	} else {
		rsMangClient, err := meta.(ClientSession).ResourceManagementAPI()
		if err != nil {
			return err
		}
		resourceGroupQuery := management.ResourceGroupQuery{
			Default: true,
		}
		grpList, err := rsMangClient.ResourceGroup().List(&resourceGroupQuery)
		if err != nil {
			return err
		}
		rsInst.ResourceGroupID = grpList[0].ID
	}

	if parameters, ok := d.GetOk("parameters"); ok {
		rsInst.Parameters = parameters.(map[string]interface{})
	}

	if _, ok := d.GetOk("tags"); ok {
		rsInst.Tags = getServiceTags(d)
	}

	instance, err := rsConClient.ResourceServiceInstance().CreateInstance(rsInst)
	if err != nil {
		return fmt.Errorf("Error creating resource instance: %s", err)
	}

	d.SetId(instance.ID)

	_, err = waitForResourceInstanceCreate(d, meta)
	if err != nil {
		return fmt.Errorf(
			"Error waiting for create resource instance (%s) to be succeeded: %s", d.Id(), err)
	}

	return resourceIBMResourceInstanceRead(d, meta)
}

func resourceIBMResourceInstanceRead(d *schema.ResourceData, meta interface{}) error {
	rsConClient, err := meta.(ClientSession).ResourceControllerAPI()
	if err != nil {
		return err
	}

	instanceID := d.Id()

	instance, err := rsConClient.ResourceServiceInstance().GetInstance(instanceID)
	if err != nil {
		return fmt.Errorf("Error retrieving resource instance: %s", err)
	}

	d.Set("tags", instance.Tags)
	d.Set("name", instance.Name)
	d.Set("status", instance.State)
	d.Set("resource_group_id", instance.ResourceGroupID)
	d.Set("parameters", flatmap.Flatten(instance.Parameters))
	d.Set("location", instance.RegionID)

	rsCatClient, err := meta.(ClientSession).ResourceCatalogAPI()
	if err != nil {
		return err
	}
	rsCatRepo := rsCatClient.ResourceCatalog()

	serviceOff, err := rsCatRepo.GetServiceName(instance.ServiceID)
	if err != nil {
		return fmt.Errorf("Error retrieving service offering: %s", err)
	}

	d.Set("service", serviceOff)

	servicePlan, err := rsCatRepo.GetServicePlanName(instance.ServicePlanID)
	if err != nil {
		return fmt.Errorf("Error retrieving plan: %s", err)
	}
	d.Set("plan", servicePlan)

	return nil
}

func resourceIBMResourceInstanceUpdate(d *schema.ResourceData, meta interface{}) error {
	rsConClient, err := meta.(ClientSession).ResourceControllerAPI()
	if err != nil {
		return err
	}

	instanceID := d.Id()

	updateReq := controller.UpdateServiceInstanceRequest{}
	if d.HasChange("name") {
		updateReq.Name = d.Get("name").(string)
	}

	if d.HasChange("plan") {
		plan := d.Get("plan").(string)
		service := d.Get("service").(string)
		rsCatClient, err := meta.(ClientSession).ResourceCatalogAPI()
		if err != nil {
			return err
		}
		rsCatRepo := rsCatClient.ResourceCatalog()

		serviceOff, err := rsCatRepo.FindByName(service, true)
		if err != nil {
			return fmt.Errorf("Error retrieving service offering: %s", err)
		}

		servicePlan, err := rsCatRepo.GetServicePlanID(serviceOff[0], plan)
		if err != nil {
			return fmt.Errorf("Error retrieving plan: %s", err)
		}

		updateReq.ServicePlanID = servicePlan

	}

	if d.HasChange("tags") {
		tags := getServiceTags(d)
		updateReq.Tags = tags
	}

	_, err = rsConClient.ResourceServiceInstance().UpdateInstance(instanceID, updateReq)
	if err != nil {
		return fmt.Errorf("Error updating resource instance: %s", err)
	}

	_, err = waitForResourceInstanceUpdate(d, meta)
	if err != nil {
		return fmt.Errorf(
			"Error waiting for update resource instance (%s) to be succeeded: %s", d.Id(), err)
	}

	return resourceIBMResourceInstanceRead(d, meta)
}

func resourceIBMResourceInstanceDelete(d *schema.ResourceData, meta interface{}) error {
	rsConClient, err := meta.(ClientSession).ResourceControllerAPI()
	if err != nil {
		return err
	}
	id := d.Id()

	err = rsConClient.ResourceServiceInstance().DeleteInstance(id, true)
	if err != nil {
		return fmt.Errorf("Error deleting resource instance: %s", err)
	}

	_, err = waitForResourceInstanceDelete(d, meta)
	if err != nil {
		return fmt.Errorf(
			"Error waiting for resource instance (%s) to be deleted: %s", d.Id(), err)
	}

	d.SetId("")

	return nil
}
func resourceIBMResourceInstanceExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	rsConClient, err := meta.(ClientSession).ResourceControllerAPI()
	if err != nil {
		return false, err
	}
	instanceID := d.Id()

	instance, err := rsConClient.ResourceServiceInstance().GetInstance(instanceID)
	if err != nil {
		if apiErr, ok := err.(bmxerror.RequestFailure); ok {
			if apiErr.StatusCode() == 404 {
				return false, nil
			}
		}
		return false, fmt.Errorf("Error communicating with the API: %s", err)
	}

	return instance.ID == instanceID, nil
}

func waitForResourceInstanceCreate(d *schema.ResourceData, meta interface{}) (interface{}, error) {
	rsConClient, err := meta.(ClientSession).ResourceControllerAPI()
	if err != nil {
		return false, err
	}
	instanceID := d.Id()

	stateConf := &resource.StateChangeConf{
		Pending: []string{rsInstanceProgressStatus, rsInstanceInactiveStatus},
		Target:  []string{rsInstanceSuccessStatus},
		Refresh: func() (interface{}, string, error) {
			instance, err := rsConClient.ResourceServiceInstance().GetInstance(instanceID)
			if err != nil {
				if apiErr, ok := err.(bmxerror.RequestFailure); ok && apiErr.StatusCode() == 404 {
					return nil, "", fmt.Errorf("The resource instance %s does not exist anymore: %v", d.Id(), err)
				}
				return nil, "", err
			}
			if instance.State == rsInstanceFailStatus {
				return instance, instance.State, fmt.Errorf("The resource instance %s failed: %v", d.Id(), err)
			}
			return instance, instance.State, nil
		},
		Timeout:    d.Timeout(schema.TimeoutCreate),
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	return stateConf.WaitForState()
}

func waitForResourceInstanceUpdate(d *schema.ResourceData, meta interface{}) (interface{}, error) {
	rsConClient, err := meta.(ClientSession).ResourceControllerAPI()
	if err != nil {
		return false, err
	}
	instanceID := d.Id()

	stateConf := &resource.StateChangeConf{
		Pending: []string{rsInstanceProgressStatus, rsInstanceInactiveStatus},
		Target:  []string{rsInstanceSuccessStatus},
		Refresh: func() (interface{}, string, error) {
			instance, err := rsConClient.ResourceServiceInstance().GetInstance(instanceID)
			if err != nil {
				if apiErr, ok := err.(bmxerror.RequestFailure); ok && apiErr.StatusCode() == 404 {
					return nil, "", fmt.Errorf("The resource instance %s does not exist anymore: %v", d.Id(), err)
				}
				return nil, "", err
			}
			if instance.State == rsInstanceFailStatus {
				return instance, instance.State, fmt.Errorf("The resource instance %s failed: %v", d.Id(), err)
			}
			return instance, instance.State, nil
		},
		Timeout:    d.Timeout(schema.TimeoutUpdate),
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	return stateConf.WaitForState()
}

func waitForResourceInstanceDelete(d *schema.ResourceData, meta interface{}) (interface{}, error) {
	rsConClient, err := meta.(ClientSession).ResourceControllerAPI()
	if err != nil {
		return false, err
	}
	instanceID := d.Id()
	stateConf := &resource.StateChangeConf{
		Pending: []string{rsInstanceProgressStatus, rsInstanceInactiveStatus, rsInstanceSuccessStatus},
		Target:  []string{rsInstanceRemovedStatus},
		Refresh: func() (interface{}, string, error) {
			instance, err := rsConClient.ResourceServiceInstance().GetInstance(instanceID)
			if err != nil {
				if apiErr, ok := err.(bmxerror.RequestFailure); ok && apiErr.StatusCode() == 404 {
					return instance, rsInstanceSuccessStatus, nil
				}
				return nil, "", err
			}
			if instance.State == rsInstanceFailStatus {
				return instance, instance.State, fmt.Errorf("The resource instance %s failed to delete: %v", d.Id(), err)
			}
			return instance, instance.State, nil
		},
		Timeout:    d.Timeout(schema.TimeoutDelete),
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	return stateConf.WaitForState()
}

func filterDeployments(deployments []models.ServiceDeployment, location string) ([]models.ServiceDeployment, map[string]bool) {
	supportedDeployments := []models.ServiceDeployment{}
	supportedLocations := make(map[string]bool)
	for _, d := range deployments {
		if d.Metadata.RCCompatible {
			deploymentLocation := d.Metadata.Deployment.Location
			supportedLocations[deploymentLocation] = true
			if deploymentLocation == location {
				supportedDeployments = append(supportedDeployments, d)
			}
		}
	}
	return supportedDeployments, supportedLocations
}
