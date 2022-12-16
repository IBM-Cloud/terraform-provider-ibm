// Copyright IBM Corp. 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package metricsrouter

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
	"github.com/IBM/platform-services-go-sdk/metricsrouterv3"
)

func ResourceIBMMetricsRouterRoute() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMMetricsRouterRouteCreate,
		ReadContext:   resourceIBMMetricsRouterRouteRead,
		UpdateContext: resourceIBMMetricsRouterRouteUpdate,
		DeleteContext: resourceIBMMetricsRouterRouteDelete,
		Importer:      &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validate.InvokeValidator("ibm_metrics_router_route", "name"),
				Description:  "The name of the route. The name must be 1000 characters or less and cannot include any special characters other than `(space) - . _ :`. Do not include any personal identifying information (PII) in any resource names.",
			},
			"rules": &schema.Schema{
				Type:        schema.TypeList,
				Required:    true,
				Description: "Routing rules that will be evaluated in their order of the array.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"target_ids": &schema.Schema{
							Type:        schema.TypeList,
							Required:    true,
							Description: "The target ID List. All the metrics will be sent to all targets listed in the rule. You can include targets from other regions.",
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
						"inclusion_filters": &schema.Schema{
							Type:        schema.TypeList,
							Required:    true,
							Description: "A list of conditions to be satisfied for routing metrics to pre-defined target.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"operand": &schema.Schema{
										Type:        schema.TypeString,
										Required:    true,
										Description: "Part of CRN that can be compared with values.",
									},
									"operator": &schema.Schema{
										Type:        schema.TypeString,
										Required:    true,
										Description: "The operation to be performed between operand and the provided values. 'is' to be used with one value and 'in' can support upto 20 values in the array.",
									},
									"value": &schema.Schema{
										Type:        schema.TypeList,
										Required:    true,
										Description: "The provided values of the in operand to be compared with.",
										Elem:        &schema.Schema{Type: schema.TypeString},
									},
								},
							},
						},
					},
				},
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

func ResourceIBMMetricsRouterRouteValidator() *validate.ResourceValidator {
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
	)

	resourceValidator := validate.ResourceValidator{ResourceName: "ibm_metrics_router_route", Schema: validateSchema}
	return &resourceValidator
}

func resourceIBMMetricsRouterRouteCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	metricsRouterClient, err := meta.(conns.ClientSession).MetricsRouterV3()
	if err != nil {
		return diag.FromErr(err)
	}

	createRouteOptions := &metricsrouterv3.CreateRouteOptions{}

	createRouteOptions.SetName(d.Get("name").(string))
	var rules []metricsrouterv3.RulePrototype
	for _, e := range d.Get("rules").([]interface{}) {
		value := e.(map[string]interface{})
		rulesItem, err := resourceIBMMetricsRouterRouteMapToRulePrototype(value)
		if err != nil {
			return diag.FromErr(err)
		}
		rules = append(rules, *rulesItem)
	}
	createRouteOptions.SetRules(rules)

	route, response, err := metricsRouterClient.CreateRouteWithContext(context, createRouteOptions)
	if err != nil {
		log.Printf("[DEBUG] CreateRouteWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("CreateRouteWithContext failed %s\n%s", err, response))
	}

	d.SetId(*route.ID)

	return resourceIBMMetricsRouterRouteRead(context, d, meta)
}

func resourceIBMMetricsRouterRouteRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	metricsRouterClient, err := meta.(conns.ClientSession).MetricsRouterV3()
	if err != nil {
		return diag.FromErr(err)
	}

	getRouteOptions := &metricsrouterv3.GetRouteOptions{}

	getRouteOptions.SetID(d.Id())

	route, response, err := metricsRouterClient.GetRouteWithContext(context, getRouteOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		log.Printf("[DEBUG] GetRouteWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("GetRouteWithContext failed %s\n%s", err, response))
	}

	if err = d.Set("name", route.Name); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting name: %s", err))
	}
	rules := []map[string]interface{}{}
	for _, rulesItem := range route.Rules {
		rulesItemMap, err := resourceIBMMetricsRouterRouteRulePrototypeToMap(&rulesItem)
		if err != nil {
			return diag.FromErr(err)
		}
		rules = append(rules, rulesItemMap)
	}
	if err = d.Set("rules", rules); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting rules: %s", err))
	}
	if err = d.Set("crn", route.CRN); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting crn: %s", err))
	}
	if err = d.Set("version", flex.IntValue(route.Version)); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting version: %s", err))
	}
	if err = d.Set("created_at", flex.DateTimeToString(route.CreatedAt)); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting created_at: %s", err))
	}
	if err = d.Set("updated_at", flex.DateTimeToString(route.UpdatedAt)); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting updated_at: %s", err))
	}
	if err = d.Set("api_version", flex.IntValue(route.APIVersion)); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting api_version: %s", err))
	}
	if err = d.Set("message", route.Message); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting message: %s", err))
	}

	return nil
}

func resourceIBMMetricsRouterRouteUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	metricsRouterClient, err := meta.(conns.ClientSession).MetricsRouterV3()
	if err != nil {
		return diag.FromErr(err)
	}

	replaceRouteOptions := &metricsrouterv3.ReplaceRouteOptions{}

	replaceRouteOptions.SetID(d.Id())
	replaceRouteOptions.SetName(d.Get("name").(string))
	var rules []metricsrouterv3.RulePrototype
	for _, e := range d.Get("rules").([]interface{}) {
		value := e.(map[string]interface{})
		rulesItem, err := resourceIBMMetricsRouterRouteMapToRulePrototype(value)
		if err != nil {
			return diag.FromErr(err)
		}
		rules = append(rules, *rulesItem)
	}
	replaceRouteOptions.SetRules(rules)

	_, response, err := metricsRouterClient.ReplaceRouteWithContext(context, replaceRouteOptions)
	if err != nil {
		log.Printf("[DEBUG] ReplaceRouteWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("ReplaceRouteWithContext failed %s\n%s", err, response))
	}

	return resourceIBMMetricsRouterRouteRead(context, d, meta)
}

func resourceIBMMetricsRouterRouteDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	metricsRouterClient, err := meta.(conns.ClientSession).MetricsRouterV3()
	if err != nil {
		return diag.FromErr(err)
	}

	deleteRouteOptions := &metricsrouterv3.DeleteRouteOptions{}

	deleteRouteOptions.SetID(d.Id())

	response, err := metricsRouterClient.DeleteRouteWithContext(context, deleteRouteOptions)
	if err != nil {
		log.Printf("[DEBUG] DeleteRouteWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("DeleteRouteWithContext failed %s\n%s", err, response))
	}

	d.SetId("")

	return nil
}

func resourceIBMMetricsRouterRouteMapToRulePrototype(modelMap map[string]interface{}) (*metricsrouterv3.RulePrototype, error) {
	model := &metricsrouterv3.RulePrototype{}
	targetIds := []string{}
	for _, targetIdsItem := range modelMap["target_ids"].([]interface{}) {
		targetIds = append(targetIds, targetIdsItem.(string))
	}
	model.TargetIds = targetIds
	inclusionFilters := []metricsrouterv3.InclusionFilter{}
	for _, inclusionFiltersItem := range modelMap["inclusion_filters"].([]interface{}) {
		inclusionFiltersItemModel, err := resourceIBMMetricsRouterRouteMapToInclusionFilter(inclusionFiltersItem.(map[string]interface{}))
		if err != nil {
			return model, err
		}
		inclusionFilters = append(inclusionFilters, *inclusionFiltersItemModel)
	}
	model.InclusionFilters = inclusionFilters
	return model, nil
}

func resourceIBMMetricsRouterRouteMapToInclusionFilter(modelMap map[string]interface{}) (*metricsrouterv3.InclusionFilter, error) {
	model := &metricsrouterv3.InclusionFilter{}
	model.Operand = core.StringPtr(modelMap["operand"].(string))
	model.Operator = core.StringPtr(modelMap["operator"].(string))
	value := []string{}
	for _, valueItem := range modelMap["value"].([]interface{}) {
		value = append(value, valueItem.(string))
	}
	model.Value = value
	return model, nil
}

func resourceIBMMetricsRouterRouteRulePrototypeToMap(model *metricsrouterv3.Rule) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["target_ids"] = model.TargetIds
	inclusionFilters := []map[string]interface{}{}
	for _, inclusionFiltersItem := range model.InclusionFilters {
		inclusionFiltersItemMap, err := resourceIBMMetricsRouterRouteInclusionFilterToMap(&inclusionFiltersItem)
		if err != nil {
			return modelMap, err
		}
		inclusionFilters = append(inclusionFilters, inclusionFiltersItemMap)
	}
	modelMap["inclusion_filters"] = inclusionFilters
	return modelMap, nil
}

func resourceIBMMetricsRouterRouteInclusionFilterToMap(model *metricsrouterv3.InclusionFilter) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["operand"] = model.Operand
	modelMap["operator"] = model.Operator
	modelMap["value"] = model.Value
	return modelMap, nil
}
