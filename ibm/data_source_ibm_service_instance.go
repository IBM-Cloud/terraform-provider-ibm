package ibm

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceIBMServiceInstance() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceIBMServiceInstanceRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Description: "Service instance name for example, cleardb",
				Type:        schema.TypeString,
				Required:    true,
			},

			"credentials": {
				Description: "Credentials asociated with the key",
				Type:        schema.TypeMap,
				Computed:    true,
			},

			"service_plan_guid": {
				Description: "The uniquie identifier of the service offering plan type",
				Type:        schema.TypeString,
				Computed:    true,
			},
		},
	}
}

func dataSourceIBMServiceInstanceRead(d *schema.ResourceData, meta interface{}) error {
	cfClient, err := meta.(ClientSession).MccpAPI()
	if err != nil {
		return err
	}
	siAPI := cfClient.ServiceInstances()
	name := d.Get("name").(string)
	inst, err := siAPI.FindByName(name)
	if err != nil {
		return err
	}

	serviceInstance, err := siAPI.Get(inst.GUID)
	if err != nil {
		return fmt.Errorf("Error retrieving service: %s", err)
	}

	d.SetId(serviceInstance.Metadata.GUID)
	d.Set("credentials", serviceInstance.Entity.Credentials)
	d.Set("service_plan_guid", serviceInstance.Entity.ServicePlanGUID)

	return nil
}
