// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.108.0-56772134-20251111-102802
 */

package atracker

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/atrackerv2"
)

func ResourceIBMAtrackerRoute() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMAtrackerRouteCreate,
		ReadContext:   resourceIBMAtrackerRouteRead,
		UpdateContext: resourceIBMAtrackerRouteUpdate,
		DeleteContext: resourceIBMAtrackerRouteDelete,
		Importer:      &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validate.InvokeValidator("ibm_atracker_route", "name"),
				Description:  "The name of the route.",
			},
			"rules": &schema.Schema{
				Type:        schema.TypeList,
				Required:    true,
				Description: "The routing rules that will be evaluated in their order of the array. Once a rule is matched, the remaining rules in the route definition will be skipped.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"target_ids": &schema.Schema{
							Type:        schema.TypeList,
							Required:    true,
							Description: "The target ID List. All the events will be send to all targets listed in the rule. You can include targets from other regions.",
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
						"locations": &schema.Schema{
							Type:        schema.TypeList,
							Required:    true,
							Description: "Logs from these locations will be sent to the targets specified. Locations is a superset of regions including global and *.",
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
					},
				},
			},
			"managed_by": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Default:  "account",
				// Suppress the diff state transition from nil, account as they are equivalent.
				DiffSuppressFunc: func(k, old, newv string, d *schema.ResourceData) bool {
					if old == "" && newv == "account" {
						return true
					} else if old == "account" && newv == "" {
						return true
					} else {
						return false
					}
				},
				ValidateFunc: validate.InvokeValidator("ibm_atracker_route", "managed_by"),
				Description:  "Present when the route is enterprise-managed (`managed_by: enterprise`).",
			},
			"crn": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The crn of the route resource.",
			},
			"version": &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The version of the route.",
			},
			"created_at": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The timestamp of the route creation time.",
			},
			"updated_at": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The timestamp of the route last updated time.",
			},
			"api_version": &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The API version of the route.",
			},
			"message": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "An optional message containing information about the route.",
			},
		},
	}
}

func ResourceIBMAtrackerRouteValidator() *validate.ResourceValidator {
	validateSchema := make([]validate.ValidateSchema, 0)
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 "name",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Required:                   true,
			Regexp:                     `^[a-zA-Z0-9 -._:]+$`,
			MinValueLength:             1,
			MaxValueLength:             1000,
		},
		validate.ValidateSchema{
			Identifier:                 "managed_by",
			ValidateFunctionIdentifier: validate.ValidateAllowedStringValue,
			Type:                       validate.TypeString,
			Optional:                   true,
			AllowedValues:              "account, enterprise",
		},
	)

	resourceValidator := validate.ResourceValidator{ResourceName: "ibm_atracker_route", Schema: validateSchema}
	return &resourceValidator
}

func resourceIBMAtrackerRouteCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	atrackerClient, err := meta.(conns.ClientSession).AtrackerV2()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_atracker_route", "create", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	createRouteOptions := &atrackerv2.CreateRouteOptions{}

	createRouteOptions.SetName(d.Get("name").(string))
	var rules []atrackerv2.RulePrototype
	for _, v := range d.Get("rules").([]interface{}) {
		value := v.(map[string]interface{})
		rulesItem, err := ResourceIBMAtrackerRouteMapToRulePrototype(value)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_atracker_route", "create", "parse-rules").GetDiag()
		}
		rules = append(rules, *rulesItem)
	}
	createRouteOptions.SetRules(rules)
	if _, ok := d.GetOk("managed_by"); ok {
		createRouteOptions.SetManagedBy(d.Get("managed_by").(string))
	}

	route, _, err := atrackerClient.CreateRouteWithContext(context, createRouteOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("CreateRouteWithContext failed: %s", err.Error()), "ibm_atracker_route", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(*route.ID)
	d.Set("api_version", 2)

	return resourceIBMAtrackerRouteRead(context, d, meta)
}

func resourceIBMAtrackerRouteRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	atrackerClient, err := meta.(conns.ClientSession).AtrackerV2()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_atracker_route", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getRouteOptions := &atrackerv2.GetRouteOptions{}

	getRouteOptions.SetID(d.Id())

	route, response, err := atrackerClient.GetRouteWithContext(context, getRouteOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetRouteWithContext failed: %s", err.Error()), "ibm_atracker_route", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	if err = d.Set("name", route.Name); err != nil {
		err = fmt.Errorf("Error setting name: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_atracker_route", "read", "set-name").GetDiag()
	}
	rules := []map[string]interface{}{}
	for _, rulesItem := range route.Rules {
		rulesItemMap, err := ResourceIBMAtrackerRouteRuleToMap(&rulesItem) // #nosec G601
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_atracker_route", "read", "rules-to-map").GetDiag()
		}
		rules = append(rules, rulesItemMap)
	}
	if err = d.Set("rules", rules); err != nil {
		err = fmt.Errorf("Error setting rules: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_atracker_route", "read", "set-rules").GetDiag()
	}
	if !core.IsNil(route.ManagedBy) {
		if err = d.Set("managed_by", route.ManagedBy); err != nil {
			err = fmt.Errorf("Error setting managed_by: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_atracker_route", "read", "set-managed_by").GetDiag()
		}
	}
	if err = d.Set("crn", route.CRN); err != nil {
		err = fmt.Errorf("Error setting crn: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_atracker_route", "read", "set-crn").GetDiag()
	}
	if !core.IsNil(route.Version) {
		if err = d.Set("version", flex.IntValue(route.Version)); err != nil {
			err = fmt.Errorf("Error setting version: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_atracker_route", "read", "set-version").GetDiag()
		}
	}
	if err = d.Set("created_at", flex.DateTimeToString(route.CreatedAt)); err != nil {
		err = fmt.Errorf("Error setting created_at: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_atracker_route", "read", "set-created_at").GetDiag()
	}
	if err = d.Set("updated_at", flex.DateTimeToString(route.UpdatedAt)); err != nil {
		err = fmt.Errorf("Error setting updated_at: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_atracker_route", "read", "set-updated_at").GetDiag()
	}
	if err = d.Set("api_version", flex.IntValue(route.APIVersion)); err != nil {
		err = fmt.Errorf("Error setting api_version: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_atracker_route", "read", "set-api_version").GetDiag()
	}
	if !core.IsNil(route.Message) {
		if err = d.Set("message", route.Message); err != nil {
			err = fmt.Errorf("Error setting message: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_atracker_route", "read", "set-message").GetDiag()
		}
	}

	return nil
}

func resourceIBMAtrackerRouteUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	atrackerClient, err := meta.(conns.ClientSession).AtrackerV2()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_atracker_route", "update", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	replaceRouteOptions := &atrackerv2.ReplaceRouteOptions{}

	replaceRouteOptions.SetID(d.Id())
	replaceRouteOptions.SetName(d.Get("name").(string))
	var rules []atrackerv2.RulePrototype
	for _, v := range d.Get("rules").([]interface{}) {
		value := v.(map[string]interface{})
		rulesItem, err := ResourceIBMAtrackerRouteMapToRulePrototype(value)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_atracker_route", "update", "parse-rules").GetDiag()
		}
		rules = append(rules, *rulesItem)
	}
	replaceRouteOptions.SetRules(rules)
	if _, ok := d.GetOk("managed_by"); ok {
		replaceRouteOptions.SetManagedBy(d.Get("managed_by").(string))
	}

	_, _, err = atrackerClient.ReplaceRouteWithContext(context, replaceRouteOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("ReplaceRouteWithContext failed: %s", err.Error()), "ibm_atracker_route", "update")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	return resourceIBMAtrackerRouteRead(context, d, meta)
}

func resourceIBMAtrackerRouteDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	atrackerClient, err := meta.(conns.ClientSession).AtrackerV2()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_atracker_route", "delete", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	deleteRouteOptions := &atrackerv2.DeleteRouteOptions{}

	deleteRouteOptions.SetID(d.Id())

	_, err = atrackerClient.DeleteRouteWithContext(context, deleteRouteOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("DeleteRouteWithContext failed: %s", err.Error()), "ibm_atracker_route", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId("")

	return nil
}

func ResourceIBMAtrackerRouteMapToRulePrototype(modelMap map[string]interface{}) (*atrackerv2.RulePrototype, error) {
	model := &atrackerv2.RulePrototype{}
	targetIds := []string{}
	for _, targetIdsItem := range modelMap["target_ids"].([]interface{}) {
		targetIds = append(targetIds, targetIdsItem.(string))
	}
	model.TargetIds = targetIds
	if modelMap["locations"] != nil {
		locations := []string{}
		for _, locationsItem := range modelMap["locations"].([]interface{}) {
			locations = append(locations, locationsItem.(string))
		}
		model.Locations = locations
	}
	return model, nil
}

func ResourceIBMAtrackerRouteRuleToMap(model *atrackerv2.Rule) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["target_ids"] = model.TargetIds
	modelMap["locations"] = model.Locations
	return modelMap, nil
}
