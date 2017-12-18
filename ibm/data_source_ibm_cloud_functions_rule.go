package ibm

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceIBMCloudFunctionsRule() *schema.Resource {
	return &schema.Resource{

		Read: dataSourceIBMCloudFunctionsRuleRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the rule.",
			},
			"trigger_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Name of the trigger.",
			},
			"action_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Name of an action.",
			},
			"status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Status of the rule.",
			},
			"publish": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Rule Visibility.",
			},
			"version": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Semantic version of the rule",
			},
		},
	}
}

func dataSourceIBMCloudFunctionsRuleRead(d *schema.ResourceData, meta interface{}) error {
	wskClient, err := meta.(ClientSession).CloudFunctionsClient()
	if err != nil {
		return err
	}
	ruleService := wskClient.Rules
	name := d.Get("name").(string)

	rule, _, err := ruleService.Get(name)
	if err != nil {
		return fmt.Errorf("Error retrieving IBM Cloud Functions Rule %s : %s", name, err)
	}

	d.SetId(rule.Name)
	d.Set("name", rule.Name)
	d.Set("publish", rule.Publish)
	d.Set("version", rule.Version)
	d.Set("status", rule.Status)
	d.Set("trigger_name", rule.Trigger.(map[string]interface{})["name"])
	path := rule.Action.(map[string]interface{})["path"]
	actionName := rule.Action.(map[string]interface{})["name"]
	d.Set("action_name", fmt.Sprintf("/%s/%s", path, actionName))
	return nil
}
