package ibm

import (
	"fmt"

	"github.com/IBM-Bluemix/bluemix-go/api/mccp/mccpv2"
	"github.com/IBM-Bluemix/bluemix-go/bmxerror"
	"github.com/IBM-Bluemix/bluemix-go/helpers"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceIBMServiceInstance() *schema.Resource {
	return &schema.Resource{
		Create:   resourceIBMServiceInstanceCreate,
		Read:     resourceIBMServiceInstanceRead,
		Update:   resourceIBMServiceInstanceUpdate,
		Delete:   resourceIBMServiceInstanceDelete,
		Exists:   resourceIBMServiceInstanceExists,
		Importer: &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "A name for the service instance",
			},

			"space_guid": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The guid of the space in which the instance will be created",
			},

			"service": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The name of the service offering like cleardb, cloudantNOSQLDb etc",
			},

			"credentials": {
				Description: "Credentials asociated with the key",
				Computed:    true,
				Type:        schema.TypeMap,
			},

			"service_plan_guid": {
				Description: "The uniquie identifier of the service offering plan type",
				Computed:    true,
				Type:        schema.TypeString,
			},

			"parameters": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "Arbitrary parameters to pass along to the service broker. Must be a JSON object",
			},

			"plan": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The plan type of the service",
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

func resourceIBMServiceInstanceCreate(d *schema.ResourceData, meta interface{}) error {
	cfClient, err := meta.(ClientSession).MccpAPI()
	if err != nil {
		return err
	}
	serviceName := d.Get("service").(string)
	plan := d.Get("plan").(string)
	name := d.Get("name").(string)
	spaceGUID := d.Get("space_guid").(string)

	svcInst := mccpv2.ServiceInstanceCreateRequest{
		Name:      name,
		SpaceGUID: spaceGUID,
	}

	serviceOff, err := cfClient.ServiceOfferings().FindByLabel(serviceName)
	if err != nil {
		return fmt.Errorf("Error retrieving service offering: %s", err)
	}

	servicePlan, err := cfClient.ServicePlans().FindPlanInServiceOffering(serviceOff.GUID, plan)
	if err != nil {
		return fmt.Errorf("Error retrieving plan: %s", err)
	}
	svcInst.PlanGUID = servicePlan.GUID

	if parameters, ok := d.GetOk("parameters"); ok {
		svcInst.Params = parameters.(map[string]interface{})
	}

	if _, ok := d.GetOk("tags"); ok {
		svcInst.Tags = getServiceTags(d)
	}

	service, err := cfClient.ServiceInstances().Create(svcInst)
	if err != nil {
		return fmt.Errorf("Error creating service: %s", err)
	}

	d.SetId(service.Metadata.GUID)

	return resourceIBMServiceInstanceRead(d, meta)
}

func resourceIBMServiceInstanceRead(d *schema.ResourceData, meta interface{}) error {
	cfClient, err := meta.(ClientSession).MccpAPI()
	if err != nil {
		return err
	}

	serviceGUID := d.Id()

	service, err := cfClient.ServiceInstances().Get(serviceGUID)
	if err != nil {
		return fmt.Errorf("Error retrieving service: %s", err)
	}

	servicePlanGUID := service.Entity.ServicePlanGUID
	d.Set("service_plan_guid", servicePlanGUID)
	d.Set("credentials", service.Entity.Credentials)
	d.Set("tags", service.Entity.Tags)
	d.Set("name", service.Entity.Name)

	p, err := cfClient.ServicePlans().Get(servicePlanGUID)
	if err != nil {
		return err
	}
	d.Set("plan", p.Entity.Name)

	svcOff, err := cfClient.ServiceOfferings().Get(p.Entity.ServiceGUID)
	if err != nil {
		return err
	}
	d.Set("service", svcOff.Entity.Label)

	return nil
}

func resourceIBMServiceInstanceUpdate(d *schema.ResourceData, meta interface{}) error {
	cfClient, err := meta.(ClientSession).MccpAPI()
	if err != nil {
		return err
	}

	serviceGUID := d.Id()

	updateReq := mccpv2.ServiceInstanceUpdateRequest{}
	if d.HasChange("name") {
		updateReq.Name = helpers.String(d.Get("name").(string))
	}

	if d.HasChange("plan") {
		plan := d.Get("plan").(string)
		service := d.Get("service").(string)
		serviceOff, err := cfClient.ServiceOfferings().FindByLabel(service)
		if err != nil {
			return fmt.Errorf("Error retrieving service offering: %s", err)
		}

		servicePlan, err := cfClient.ServicePlans().FindPlanInServiceOffering(serviceOff.GUID, plan)
		if err != nil {
			return fmt.Errorf("Error retrieving plan: %s", err)
		}
		updateReq.PlanGUID = helpers.String(servicePlan.GUID)

	}

	if d.HasChange("parameters") {
		updateReq.Params = d.Get("parameters").(map[string]interface{})
	}

	if d.HasChange("tags") {
		tags := getServiceTags(d)
		updateReq.Tags = &tags
	}

	_, err = cfClient.ServiceInstances().Update(serviceGUID, updateReq)
	if err != nil {
		return fmt.Errorf("Error updating service: %s", err)
	}

	return resourceIBMServiceInstanceRead(d, meta)
}

func resourceIBMServiceInstanceDelete(d *schema.ResourceData, meta interface{}) error {
	cfClient, err := meta.(ClientSession).MccpAPI()
	if err != nil {
		return err
	}
	id := d.Id()

	err = cfClient.ServiceInstances().Delete(id)
	if err != nil {
		return fmt.Errorf("Error deleting service: %s", err)
	}

	d.SetId("")

	return nil
}
func resourceIBMServiceInstanceExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	cfClient, err := meta.(ClientSession).MccpAPI()
	if err != nil {
		return false, err
	}
	serviceGUID := d.Id()

	service, err := cfClient.ServiceInstances().Get(serviceGUID)
	if err != nil {
		if apiErr, ok := err.(bmxerror.RequestFailure); ok {
			if apiErr.StatusCode() == 404 {
				return false, nil
			}
		}
		return false, fmt.Errorf("Error communicating with the API: %s", err)
	}

	return service.Metadata.GUID == serviceGUID, nil
}

func getServiceTags(d *schema.ResourceData) []string {
	tagSet := d.Get("tags").(*schema.Set)

	if tagSet.Len() == 0 {
		empty := []string{}
		return empty
	}

	tags := make([]string, 0, tagSet.Len())
	for _, elem := range tagSet.List() {
		tag := elem.(string)
		tags = append(tags, tag)
	}
	return tags
}
