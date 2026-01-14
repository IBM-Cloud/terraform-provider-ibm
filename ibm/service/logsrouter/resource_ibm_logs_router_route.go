// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.108.0-56772134-20251111-102802
*/

package logsrouter

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
	"github.com/IBM/platform-services-go-sdk/logsrouterv3"
)

func ResourceIBMLogsRouterRoute() *schema.Resource {
	return &schema.Resource{
		CreateContext:   resourceIBMLogsRouterRouteCreate,
		ReadContext:     resourceIBMLogsRouterRouteRead,
		UpdateContext:   resourceIBMLogsRouterRouteUpdate,
		DeleteContext:   resourceIBMLogsRouterRouteDelete,
		Importer: &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				ValidateFunc: validate.InvokeValidator("ibm_logs_router_route", "name"),
				Description: "The name of the route.",
			},
			"rules": &schema.Schema{
				Type:        schema.TypeList,
				Required:    true,
				Description: "The routing rules that will be evaluated in their order of the array. Once a rule is matched, the remaining rules in the route definition will be skipped.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The action if the inclusion_filters matches, default is `send` action.",
						},
						"targets": &schema.Schema{
							Type:        schema.TypeList,
							Required:    true,
							Description: "The target ID List. Platform logs will be sent to all targets listed in the rule. You can include targets from other regions.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": &schema.Schema{
										Type:        schema.TypeString,
										Required:    true,
										Description: "The target uuid for a pre-defined platform logs router target.",
									},
									"crn": &schema.Schema{
										Type:        schema.TypeString,
										Required:    true,
										Description: "The CRN of a pre-defined logs-router target.",
									},
									"name": &schema.Schema{
										Type:        schema.TypeString,
										Required:    true,
										Description: "The name of a pre-defined logs-router target.",
									},
									"target_type": &schema.Schema{
										Type:        schema.TypeString,
										Required:    true,
										Description: "The type of the target.",
									},
								},
							},
						},
						"inclusion_filters": &schema.Schema{
							Type:        schema.TypeList,
							Required:    true,
							Description: "A list of conditions to be satisfied for routing platform logs to pre-defined target.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"operand": &schema.Schema{
										Type:        schema.TypeString,
										Required:    true,
										Description: "Part of CRN that can be compared with values. Currently only location is supported.",
									},
									"operator": &schema.Schema{
										Type:        schema.TypeString,
										Required:    true,
										Description: "The operation to be performed between operand and the provided values. 'is' to be used with one value and 'in' can support up to 20 values in the array.",
									},
									"values": &schema.Schema{
										Type:        schema.TypeList,
										Required:    true,
										Description: "The provided string values of the operand to be compared with.",
										Elem:        &schema.Schema{Type: schema.TypeString},
									},
								},
							},
						},
					},
				},
			},
			"managed_by": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				ValidateFunc: validate.InvokeValidator("ibm_logs_router_route", "managed_by"),
				Description: "Present when the route is enterprise-managed (`managed_by: enterprise`).",
			},
			"crn": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The crn of the route resource.",
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
		},
	}
}

func ResourceIBMLogsRouterRouteValidator() *validate.ResourceValidator {
	validateSchema := make([]validate.ValidateSchema, 0)
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 "name",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Required:                   true,
			Regexp:                     `^[a-zA-Z0-9 \-._:]+$`,
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

	resourceValidator := validate.ResourceValidator{ResourceName: "ibm_logs_router_route", Schema: validateSchema}
	return &resourceValidator
}

func resourceIBMLogsRouterRouteCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logsRouterClient, err := meta.(conns.ClientSession).LogsRouterV3()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_logs_router_route", "create", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	createRouteOptions := &logsrouterv3.CreateRouteOptions{}

	createRouteOptions.SetName(d.Get("name").(string))
	var rules []logsrouterv3.RulePrototype
	for _, v := range d.Get("rules").([]interface{}) {
		value := v.(map[string]interface{})
		rulesItem, err := ResourceIBMLogsRouterRouteMapToRulePrototype(value)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_logs_router_route", "create", "parse-rules").GetDiag()
		}
		rules = append(rules, *rulesItem)
	}
	createRouteOptions.SetRules(rules)
	if _, ok := d.GetOk("managed_by"); ok {
		createRouteOptions.SetManagedBy(d.Get("managed_by").(string))
	}

	route, _, err := logsRouterClient.CreateRouteWithContext(context, createRouteOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("CreateRouteWithContext failed: %s", err.Error()), "ibm_logs_router_route", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(*route.ID)

	return resourceIBMLogsRouterRouteRead(context, d, meta)
}

func resourceIBMLogsRouterRouteRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logsRouterClient, err := meta.(conns.ClientSession).LogsRouterV3()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_logs_router_route", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getRouteOptions := &logsrouterv3.GetRouteOptions{}

	getRouteOptions.SetID(d.Id())

	route, response, err := logsRouterClient.GetRouteWithContext(context, getRouteOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetRouteWithContext failed: %s", err.Error()), "ibm_logs_router_route", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	if err = d.Set("name", route.Name); err != nil {
		err = fmt.Errorf("Error setting name: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_logs_router_route", "read", "set-name").GetDiag()
	}
	rules := []map[string]interface{}{}
	for _, rulesItem := range route.Rules {
		rulesItemMap, err := ResourceIBMLogsRouterRouteRuleToMap(&rulesItem) // #nosec G601
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_logs_router_route", "read", "rules-to-map").GetDiag()
		}
		rules = append(rules, rulesItemMap)
	}
	if err = d.Set("rules", rules); err != nil {
		err = fmt.Errorf("Error setting rules: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_logs_router_route", "read", "set-rules").GetDiag()
	}
	if !core.IsNil(route.ManagedBy) {
		if err = d.Set("managed_by", route.ManagedBy); err != nil {
			err = fmt.Errorf("Error setting managed_by: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_logs_router_route", "read", "set-managed_by").GetDiag()
		}
	}
	if err = d.Set("crn", route.CRN); err != nil {
		err = fmt.Errorf("Error setting crn: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_logs_router_route", "read", "set-crn").GetDiag()
	}
	if err = d.Set("created_at", flex.DateTimeToString(route.CreatedAt)); err != nil {
		err = fmt.Errorf("Error setting created_at: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_logs_router_route", "read", "set-created_at").GetDiag()
	}
	if err = d.Set("updated_at", flex.DateTimeToString(route.UpdatedAt)); err != nil {
		err = fmt.Errorf("Error setting updated_at: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_logs_router_route", "read", "set-updated_at").GetDiag()
	}

	return nil
}

func resourceIBMLogsRouterRouteUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logsRouterClient, err := meta.(conns.ClientSession).LogsRouterV3()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_logs_router_route", "update", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	updateRouteOptions := &logsrouterv3.UpdateRouteOptions{}

	updateRouteOptions.SetID(d.Id())

	hasChange := false

	if d.HasChange("name") {
		updateRouteOptions.SetName(d.Get("name").(string))
		hasChange = true
	}
	if d.HasChange("rules") {
		var rules []logsrouterv3.RulePrototype
		for _, v := range d.Get("rules").([]interface{}) {
			value := v.(map[string]interface{})
			rulesItem, err := ResourceIBMLogsRouterRouteMapToRulePrototype(value)
			if err != nil {
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_logs_router_route", "update", "parse-rules").GetDiag()
			}
			rules = append(rules, *rulesItem)
		}
		updateRouteOptions.SetRules(rules)
		hasChange = true
	}

	if hasChange {
		_, _, err = logsRouterClient.UpdateRouteWithContext(context, updateRouteOptions)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("UpdateRouteWithContext failed: %s", err.Error()), "ibm_logs_router_route", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
	}

	return resourceIBMLogsRouterRouteRead(context, d, meta)
}

func resourceIBMLogsRouterRouteDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logsRouterClient, err := meta.(conns.ClientSession).LogsRouterV3()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_logs_router_route", "delete", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	deleteRouteOptions := &logsrouterv3.DeleteRouteOptions{}

	deleteRouteOptions.SetID(d.Id())

	_, err = logsRouterClient.DeleteRouteWithContext(context, deleteRouteOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("DeleteRouteWithContext failed: %s", err.Error()), "ibm_logs_router_route", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId("")

	return nil
}

func ResourceIBMLogsRouterRouteMapToRulePrototype(modelMap map[string]interface{}) (*logsrouterv3.RulePrototype, error) {
	model := &logsrouterv3.RulePrototype{}
	if modelMap["action"] != nil && modelMap["action"].(string) != "" {
		model.Action = core.StringPtr(modelMap["action"].(string))
	}
	targets := []logsrouterv3.TargetIdentity{}
	for _, targetsItem := range modelMap["targets"].([]interface{}) {
		targetsItemModel, err := ResourceIBMLogsRouterRouteMapToTargetIdentity(targetsItem.(map[string]interface{}))
		if err != nil {
			return model, err
		}
		targets = append(targets, *targetsItemModel)
	}
	model.Targets = targets
	inclusionFilters := []logsrouterv3.InclusionFilterPrototype{}
	for _, inclusionFiltersItem := range modelMap["inclusion_filters"].([]interface{}) {
		inclusionFiltersItemModel, err := ResourceIBMLogsRouterRouteMapToInclusionFilterPrototype(inclusionFiltersItem.(map[string]interface{}))
		if err != nil {
			return model, err
		}
		inclusionFilters = append(inclusionFilters, *inclusionFiltersItemModel)
	}
	model.InclusionFilters = inclusionFilters
	return model, nil
}

func ResourceIBMLogsRouterRouteMapToTargetIdentity(modelMap map[string]interface{}) (*logsrouterv3.TargetIdentity, error) {
	model := &logsrouterv3.TargetIdentity{}
	model.ID = core.StringPtr(modelMap["id"].(string))
	return model, nil
}

func ResourceIBMLogsRouterRouteMapToInclusionFilterPrototype(modelMap map[string]interface{}) (*logsrouterv3.InclusionFilterPrototype, error) {
	model := &logsrouterv3.InclusionFilterPrototype{}
	model.Operand = core.StringPtr(modelMap["operand"].(string))
	model.Operator = core.StringPtr(modelMap["operator"].(string))
	values := []string{}
	for _, valuesItem := range modelMap["values"].([]interface{}) {
		values = append(values, valuesItem.(string))
	}
	model.Values = values
	return model, nil
}

func ResourceIBMLogsRouterRouteRuleToMap(model *logsrouterv3.Rule) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Action != nil {
		modelMap["action"] = *model.Action
	}
	targets := []map[string]interface{}{}
	for _, targetsItem := range model.Targets {
		targetsItemMap, err := ResourceIBMLogsRouterRouteTargetReferenceToMap(&targetsItem) // #nosec G601
		if err != nil {
			return modelMap, err
		}
		targets = append(targets, targetsItemMap)
	}
	modelMap["targets"] = targets
	inclusionFilters := []map[string]interface{}{}
	for _, inclusionFiltersItem := range model.InclusionFilters {
		inclusionFiltersItemMap, err := ResourceIBMLogsRouterRouteInclusionFilterToMap(&inclusionFiltersItem) // #nosec G601
		if err != nil {
			return modelMap, err
		}
		inclusionFilters = append(inclusionFilters, inclusionFiltersItemMap)
	}
	modelMap["inclusion_filters"] = inclusionFilters
	return modelMap, nil
}

func ResourceIBMLogsRouterRouteTargetReferenceToMap(model *logsrouterv3.TargetReference) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["id"] = *model.ID
	modelMap["crn"] = *model.CRN
	modelMap["name"] = *model.Name
	modelMap["target_type"] = *model.TargetType
	return modelMap, nil
}

func ResourceIBMLogsRouterRouteInclusionFilterToMap(model *logsrouterv3.InclusionFilter) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["operand"] = *model.Operand
	modelMap["operator"] = *model.Operator
	modelMap["values"] = model.Values
	return modelMap, nil
}
