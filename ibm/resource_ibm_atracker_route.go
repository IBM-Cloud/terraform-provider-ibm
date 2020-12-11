/**
 * (C) Copyright IBM Corp. 2020.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package ibm

import (
	"github.com/IBM/platform-services-go-sdk/atrackerv1"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
	"time"
)

func resourceIBMAtrackerRoute() *schema.Resource {
	return &schema.Resource{
		Create:   resourceIBMAtrackerRouteCreate,
		Read:     resourceIBMAtrackerRouteRead,
		Update:   resourceIBMAtrackerRouteUpdate,
		Delete:   resourceIBMAtrackerRouteDelete,
		Exists:   resourceIBMAtrackerRouteExists,
		Importer: &schema.ResourceImporter{},
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
							Description: "The target ID List. Only one target id is supported. For regional route, the id must be V4 uuid of a target in the same region. For global route, it will be region-code and target-id separated by colon.",
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
					},
				},
			},
			"instance_id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The uuid of ATracker services in this region.",
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

func resourceIBMAtrackerRouteCreate(d *schema.ResourceData, meta interface{}) error {
	atrackerClient, err := meta.(ClientSession).AtrackerV1()
	if err != nil {
		return err
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

	route, response, err := atrackerClient.CreateRoute(createRouteOptions)
	if err != nil {
		log.Printf("[DEBUG] CreateRoute failed %s\n%s", err, response)
		return err
	}

	d.SetId(*route.ID)

	return resourceIBMAtrackerRouteRead(d, meta)
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

func resourceIBMAtrackerRouteRead(d *schema.ResourceData, meta interface{}) error {
	atrackerClient, err := meta.(ClientSession).AtrackerV1()
	if err != nil {
		return err
	}

	getRouteOptions := &atrackerv1.GetRouteOptions{}

	getRouteOptions.SetID(d.Id())

	route, response, err := atrackerClient.GetRoute(getRouteOptions)
	if err != nil {
		log.Printf("[DEBUG] GetRoute failed %s\n%s", err, response)
		return err
	}

	d.Set("name", route.Name)
	d.Set("receive_global_events", route.ReceiveGlobalEvents)
	rules := []map[string]interface{}{}
	for _, rulesItem := range route.Rules {
		rulesItemMap := resourceIBMAtrackerRouteRuleToMap(rulesItem)
		rules = append(rules, rulesItemMap)
	}
	d.Set("rules", rules)
	d.Set("instance_id", route.InstanceID)
	d.Set("crn", route.CRN)
	d.Set("version", intValue(route.Version))

	return nil
}

func resourceIBMAtrackerRouteRuleToMap(rule atrackerv1.Rule) map[string]interface{} {
	ruleMap := map[string]interface{}{}

	ruleMap["target_ids"] = rule.TargetIds

	return ruleMap
}

func resourceIBMAtrackerRouteUpdate(d *schema.ResourceData, meta interface{}) error {
	atrackerClient, err := meta.(ClientSession).AtrackerV1()
	if err != nil {
		return err
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

	_, response, err := atrackerClient.ReplaceRoute(replaceRouteOptions)
	if err != nil {
		log.Printf("[DEBUG] ReplaceRoute failed %s\n%s", err, response)
		return err
	}

	return resourceIBMAtrackerRouteRead(d, meta)
}

func resourceIBMAtrackerRouteDelete(d *schema.ResourceData, meta interface{}) error {
	atrackerClient, err := meta.(ClientSession).AtrackerV1()
	if err != nil {
		return err
	}

	deleteRouteOptions := &atrackerv1.DeleteRouteOptions{}

	deleteRouteOptions.SetID(d.Id())

	response, err := atrackerClient.DeleteRoute(deleteRouteOptions)
	if err != nil {
		log.Printf("[DEBUG] DeleteRoute failed %s\n%s", err, response)
		return err
	}

	d.SetId("")

	return nil
}

func resourceIBMAtrackerRouteExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	atrackerClient, err := meta.(ClientSession).AtrackerV1()
	if err != nil {
		return false, err
	}

	getRouteOptions := &atrackerv1.GetRouteOptions{}

	getRouteOptions.SetID(d.Id())

	route, response, err := atrackerClient.GetRoute(getRouteOptions)
	if err != nil {
		log.Printf("[DEBUG] GetRoute failed %s\n%s", err, response)
		return false, err
	}

	return *route.ID == d.Id(), nil
}
