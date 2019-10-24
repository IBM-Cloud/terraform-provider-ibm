package ibm

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceIBMServicePlan() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceIBMServicePlanRead,

		Schema: map[string]*schema.Schema{
			"service": {
				Description: "Service name for example, cloudantNoSQLDB",
				Type:        schema.TypeString,
				Required:    true,
			},

			"plan": {
				Description: "The plan type ex- shared ",
				Type:        schema.TypeString,
				Required:    true,
			},
		},
	}
}

func dataSourceIBMServicePlanRead(d *schema.ResourceData, meta interface{}) error {
	cfClient, err := meta.(ClientSession).MccpAPI()
	if err != nil {
		return err
	}
	soffAPI := cfClient.ServiceOfferings()
	spAPI := cfClient.ServicePlans()

	service := d.Get("service").(string)
	plan := d.Get("plan").(string)
	serviceOff, err := soffAPI.FindByLabel(service)
	if err != nil {
		return fmt.Errorf("Error retrieving service offering: %s", err)
	}
	servicePlan, err := spAPI.FindPlanInServiceOffering(serviceOff.GUID, plan)
	if err != nil {
		return fmt.Errorf("Error retrieving plan: %s", err)
	}

	d.SetId(servicePlan.GUID)
	return nil
}
