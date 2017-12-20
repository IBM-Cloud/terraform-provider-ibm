package ibm

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/apache/incubator-openwhisk-client-go/whisk"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceIBMFunctionRule() *schema.Resource {
	return &schema.Resource{
		Create:   resourceIBMFunctionRuleCreate,
		Read:     resourceIBMFunctionRuleRead,
		Update:   resourceIBMFunctionRuleUpdate,
		Delete:   resourceIBMFunctionRuleDelete,
		Exists:   resourceIBMFunctionRuleExists,
		Importer: &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				Description:  "Name of rule.",
				ValidateFunc: validateFunctionName,
			},
			"trigger_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of trigger.",
			},
			"action_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of action.",
				DiffSuppressFunc: func(k, o, n string, d *schema.ResourceData) bool {
					if o == "" {
						return false
					}
					if strings.HasPrefix(n, "/_") {
						temp := strings.Replace(n, "/_", "/"+os.Getenv("FUNCTION_NAMESPACE"), 1)
						if strings.Compare(temp, o) == 0 {
							return true
						}
					}
					if !strings.HasPrefix(n, "/") {
						if strings.HasPrefix(o, "/"+os.Getenv("FUNCTION_NAMESPACE")) {
							return true
						}
					}
					return false
				},
			},
			"status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Status of the rule.",
			},
			"publish": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Rule visbility.",
			},
			"version": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Semantic version of the item.",
			},
		},
	}
}

func resourceIBMFunctionRuleCreate(d *schema.ResourceData, meta interface{}) error {
	wskClient, err := meta.(ClientSession).FunctionClient()
	if err != nil {
		return err
	}
	ruleService := wskClient.Rules

	name := d.Get("name").(string)

	var qualifiedName = new(QualifiedName)

	if qualifiedName, err = NewQualifiedName(name); err != nil {
		return NewQualifiedNameError(name, err)
	}

	trigger := d.Get("trigger_name").(string)
	action := d.Get("action_name").(string)

	triggerName := getQualifiedName(trigger, getNamespaceFromProp())
	actionName := getQualifiedName(action, getNamespaceFromProp())

	payload := whisk.Rule{
		Name:      qualifiedName.GetEntityName(),
		Namespace: qualifiedName.GetNamespace(),
		Trigger:   triggerName,
		Action:    actionName,
	}

	log.Println("[INFO] Creating IBM Cloud Function rule")
	result, _, err := ruleService.Insert(&payload, true)
	if err != nil {
		return fmt.Errorf("Error creating IBM Cloud Function rule: %s", err)
	}

	d.SetId(result.Name)

	return resourceIBMFunctionRuleRead(d, meta)
}

func resourceIBMFunctionRuleRead(d *schema.ResourceData, meta interface{}) error {
	wskClient, err := meta.(ClientSession).FunctionClient()
	if err != nil {
		return err
	}
	ruleService := wskClient.Rules
	id := d.Id()

	rule, _, err := ruleService.Get(id)
	if err != nil {
		return fmt.Errorf("Error retrieving IBM Cloud Function rule %s : %s", id, err)
	}

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

func resourceIBMFunctionRuleUpdate(d *schema.ResourceData, meta interface{}) error {
	wskClient, err := meta.(ClientSession).FunctionClient()
	if err != nil {
		return err
	}
	ruleService := wskClient.Rules

	var qualifiedName = new(QualifiedName)

	if qualifiedName, err = NewQualifiedName(d.Get("name").(string)); err != nil {
		return NewQualifiedNameError(d.Get("name").(string), err)
	}

	payload := whisk.Rule{
		Name:      qualifiedName.GetEntityName(),
		Namespace: qualifiedName.GetNamespace(),
	}
	ischanged := false

	if d.HasChange("trigger_name") {
		trigger := d.Get("trigger_name").(string)
		payload.Trigger = getQualifiedName(trigger, getNamespaceFromProp())
		ischanged = true
	}

	if d.HasChange("action_name") {
		action := d.Get("action_name").(string)
		payload.Action = getQualifiedName(action, getNamespaceFromProp())
		ischanged = true
	}

	if ischanged {
		log.Println("[INFO] Update IBM Cloud Function Rule")
		result, _, err := ruleService.Insert(&payload, true)
		if err != nil {
			return fmt.Errorf("Error updating IBM Cloud Function Rule: %s", err)
		}
		_, _, err = ruleService.SetState(result.Name, "active")
		if err != nil {
			return fmt.Errorf("Error updating IBM Cloud Function Rule: %s", err)
		}
	}

	return resourceIBMFunctionRuleRead(d, meta)
}

func resourceIBMFunctionRuleDelete(d *schema.ResourceData, meta interface{}) error {
	wskClient, err := meta.(ClientSession).FunctionClient()
	if err != nil {
		return err
	}
	ruleService := wskClient.Rules
	id := d.Id()

	_, err = ruleService.Delete(id)
	if err != nil {
		return fmt.Errorf("Error deleting IBM Cloud Function Rule: %s", err)
	}

	d.SetId("")
	return nil
}

func resourceIBMFunctionRuleExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	wskClient, err := meta.(ClientSession).FunctionClient()
	if err != nil {
		return false, err
	}
	ruleService := wskClient.Rules
	id := d.Id()

	rule, resp, err := ruleService.Get(id)
	if err != nil {
		if resp.StatusCode == 404 {
			return false, nil
		}
		return false, fmt.Errorf("Error communicating with IBM Cloud Function Client : %s", err)
	}
	return rule.Name == id, nil
}
