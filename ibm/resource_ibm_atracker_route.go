// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM/platform-services-go-sdk/atrackerv1"
)

func resourceIBMAtrackerRoute() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMAtrackerRouteCreate,
		ReadContext:   resourceIBMAtrackerRouteRead,
		UpdateContext: resourceIBMAtrackerRouteUpdate,
		DeleteContext: resourceIBMAtrackerRouteDelete,
		Importer:      &schema.ResourceImporter{},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(60 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the route. Must be 180 characters or less and cannot include any special characters other than `(space) - . _ :`.",
			},
			"receive_global_events": &schema.Schema{
				Type:        schema.TypeBool,
				Required:    true,
				Description: "Whether or not all global events should be forwarded to this region.",
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
							Description: "The target ID List. Only one target id is supported.",
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
					},
				},
			},
			"crn": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The crn of this route type resource.",
			},
			"version": &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The version of this route.",
			},
		},
	}
}

func resourceIBMAtrackerRouteCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	atrackerClient, err := meta.(ClientSession).AtrackerV1()
	if err != nil {
		return diag.FromErr(err)
	}

	createRouteOptions := &atrackerv1.CreateRouteOptions{}

	createRouteOptions.SetName(d.Get("name").(string))
	createRouteOptions.SetReceiveGlobalEvents(d.Get("receive_global_events").(bool))
	var rules []atrackerv1.Rule
	for _, e := range d.Get("rules").([]interface{}) {
		value := e.(map[string]interface{})
		rulesItem := resourceIBMAtrackerRouteMapToRule(value)
		rules = append(rules, rulesItem)
	}
	createRouteOptions.SetRules(rules)

	route, response, err := atrackerClient.CreateRouteWithContext(context, createRouteOptions)
	if err != nil {
		log.Printf("[DEBUG] CreateRouteWithContext failed %s\n%s", err, response)
		return diag.FromErr(err)
	}

	d.SetId(*route.ID)

	return resourceIBMAtrackerRouteRead(context, d, meta)
}

func resourceIBMAtrackerRouteMapToRule(ruleMap map[string]interface{}) atrackerv1.Rule {
	rule := atrackerv1.Rule{}

	targetIds := []string{}
	for _, targetIdsItem := range ruleMap["target_ids"].([]interface{}) {
		targetIds = append(targetIds, targetIdsItem.(string))
	}
	rule.TargetIds = targetIds

	return rule
}

func resourceIBMAtrackerRouteRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	atrackerClient, err := meta.(ClientSession).AtrackerV1()
	if err != nil {
		return diag.FromErr(err)
	}

	getRouteOptions := &atrackerv1.GetRouteOptions{}

	getRouteOptions.SetID(d.Id())

	route, response, err := atrackerClient.GetRouteWithContext(context, getRouteOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		log.Printf("[DEBUG] GetRouteWithContext failed %s\n%s", err, response)
		return diag.FromErr(err)
	}

	if err = d.Set("name", route.Name); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting name: %s", err))
	}
	if err = d.Set("receive_global_events", route.ReceiveGlobalEvents); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting receive_global_events: %s", err))
	}
	rules := []map[string]interface{}{}
	for _, rulesItem := range route.Rules {
		rulesItemMap := resourceIBMAtrackerRouteRuleToMap(rulesItem)
		rules = append(rules, rulesItemMap)
	}
	if err = d.Set("rules", rules); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting rules: %s", err))
	}
	if err = d.Set("crn", route.CRN); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting crn: %s", err))
	}
	if err = d.Set("version", intValue(route.Version)); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting version: %s", err))
	}

	return nil
}

func resourceIBMAtrackerRouteRuleToMap(rule atrackerv1.Rule) map[string]interface{} {
	ruleMap := map[string]interface{}{}

	ruleMap["target_ids"] = rule.TargetIds

	return ruleMap
}

func resourceIBMAtrackerRouteUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	atrackerClient, err := meta.(ClientSession).AtrackerV1()
	if err != nil {
		return diag.FromErr(err)
	}

	replaceRouteOptions := &atrackerv1.ReplaceRouteOptions{}

	replaceRouteOptions.SetID(d.Id())
	replaceRouteOptions.SetName(d.Get("name").(string))
	replaceRouteOptions.SetReceiveGlobalEvents(d.Get("receive_global_events").(bool))
	var rules []atrackerv1.Rule
	for _, e := range d.Get("rules").([]interface{}) {
		value := e.(map[string]interface{})
		rulesItem := resourceIBMAtrackerRouteMapToRule(value)
		rules = append(rules, rulesItem)
	}
	replaceRouteOptions.SetRules(rules)

	_, response, err := atrackerClient.ReplaceRouteWithContext(context, replaceRouteOptions)
	if err != nil {
		log.Printf("[DEBUG] ReplaceRouteWithContext failed %s\n%s", err, response)
		return diag.FromErr(err)
	}

	return resourceIBMAtrackerRouteRead(context, d, meta)
}

func resourceIBMAtrackerRouteDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	atrackerClient, err := meta.(ClientSession).AtrackerV1()
	if err != nil {
		return diag.FromErr(err)
	}

	deleteRouteOptions := &atrackerv1.DeleteRouteOptions{}

	deleteRouteOptions.SetID(d.Id())

	response, err := atrackerClient.DeleteRouteWithContext(context, deleteRouteOptions)
	if err != nil {
		log.Printf("[DEBUG] DeleteRouteWithContext failed %s\n%s", err, response)
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
