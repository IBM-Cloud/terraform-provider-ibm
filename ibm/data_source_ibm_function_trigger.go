package ibm

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceIBMFunctionTrigger() *schema.Resource {
	return &schema.Resource{

		Read: dataSourceIBMFunctionTriggerRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of Trigger.",
			},
			"publish": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Trigger Visibility.",
			},
			"version": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Semantic version of the trigger.",
			},

			"annotations": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "All annotations set on trigger by user and those set by the IBM Cloud Function backend/API.",
			},
			"parameters": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "All parameters set on trigger by user and those set by the IBM Cloud Function backend/API.",
			},
		},
	}
}

func dataSourceIBMFunctionTriggerRead(d *schema.ResourceData, meta interface{}) error {
	wskClient, err := meta.(ClientSession).FunctionClient()
	if err != nil {
		return err
	}
	triggerService := wskClient.Triggers
	name := d.Get("name").(string)

	trigger, _, err := triggerService.Get(name)
	if err != nil {
		return fmt.Errorf("Error retrieving IBM Cloud Function Trigger %s : %s", name, err)
	}

	d.SetId(trigger.Name)
	d.Set("name", trigger.Name)
	d.Set("publish", trigger.Publish)
	d.Set("version", trigger.Version)
	annotations, err := flattenAnnotations(trigger.Annotations)
	if err != nil {
		return err
	}
	d.Set("annotations", annotations)
	parameters, err := flattenParameters(trigger.Parameters)
	if err != nil {
		return err
	}
	d.Set("parameters", parameters)

	return nil
}
