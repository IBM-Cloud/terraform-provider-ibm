// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM/appconfiguration-go-admin-sdk/appconfigurationv1"
	"github.com/IBM/go-sdk-core/v5/core"
)

func resourceIbmAppConfigSegment() *schema.Resource {
	return &schema.Resource{
		Create:   resourceIbmAppConfigSegmentCreate,
		Read:     resourceIbmAppConfigSegmentRead,
		Update:   resourceIbmAppConfigSegmentUpdate,
		Delete:   resourceIbmAppConfigSegmentDelete,
		Importer: &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"guid": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "GUID of the App Configuration service. Get it from the service instance credentials section of the dashboard.",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Segment name.",
			},
			"segment_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Segment id.",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Segment description.",
			},
			"tags": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Tags associated with the segments.",
			},
			"rules": {
				Type:        schema.TypeList,
				Required:    true,
				Description: "List of rules that determine if the entity is part of the segment.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"attribute_name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Attribute name.",
						},
						"operator": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Operator to be used for the evaluation if the entity is part of the segment.",
						},
						"values": {
							Type:        schema.TypeList,
							Required:    true,
							Description: "List of values. Entities matching any of the given values will be considered to be part of the segment.",
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
					},
				},
			},
			"created_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Creation time of the segment.",
			},
			"updated_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Last modified time of the segment data.",
			},
			"href": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Segment URL.",
			},
		},
	}
}

func resourceIbmAppConfigSegmentCreate(d *schema.ResourceData, meta interface{}) error {
	guid := d.Get("guid").(string)
	appconfigClient, err := getAppConfigClient(meta, guid)
	if err != nil {
		return err
	}
	options := &appconfigurationv1.CreateSegmentOptions{}

	options.SetName(d.Get("name").(string))
	options.SetSegmentID(d.Get("segment_id").(string))

	if _, ok := d.GetOk("description"); ok {
		options.SetDescription(d.Get("description").(string))
	}
	if _, ok := d.GetOk("tags"); ok {
		options.SetTags(d.Get("tags").(string))
	}

	var rules []appconfigurationv1.Rule
	for _, e := range d.Get("rules").([]interface{}) {
		value := e.(map[string]interface{})
		rulesItem := resourceIbmAppConfigSegmentMapToRuleObject(value)
		rules = append(rules, rulesItem)
	}
	options.SetRules(rules)

	result, response, err := appconfigClient.CreateSegment(options)
	if err != nil {
		log.Printf("[DEBUG] CreateSegment failed %s\n%s", err, response)
		return err
	}

	d.SetId(fmt.Sprintf("%s/%s", guid, *result.SegmentID))

	return resourceIbmAppConfigSegmentRead(d, meta)
}

func resourceIbmAppConfigSegmentMapToRuleObject(ruleObjectMap map[string]interface{}) appconfigurationv1.Rule {
	ruleObject := appconfigurationv1.Rule{}

	ruleObject.AttributeName = core.StringPtr(ruleObjectMap["attribute_name"].(string))
	ruleObject.Operator = core.StringPtr(ruleObjectMap["operator"].(string))
	values := []string{}
	for _, valuesItem := range ruleObjectMap["values"].([]interface{}) {
		values = append(values, valuesItem.(string))
	}
	ruleObject.Values = values

	return ruleObject
}

func resourceIbmAppConfigSegmentRead(d *schema.ResourceData, meta interface{}) error {
	parts, err := idParts(d.Id())
	if err != nil {
		return nil
	}
	appconfigClient, err := getAppConfigClient(meta, parts[0])
	if err != nil {
		return err
	}
	options := &appconfigurationv1.GetSegmentOptions{}

	options.SetSegmentID(parts[1])

	result, response, err := appconfigClient.GetSegment(options)
	if err != nil {
		log.Printf("[DEBUG] GetSegment failed %s\n%s", err, response)
		return err
	}

	if result.Name != nil {
		if err = d.Set("name", result.Name); err != nil {
			return fmt.Errorf("error setting name: %s", err)
		}
	}
	if result.SegmentID != nil {
		if err = d.Set("segment_id", result.SegmentID); err != nil {
			return fmt.Errorf("error setting segment_id: %s", err)
		}
	}
	if result.Description != nil {
		if err = d.Set("description", result.Description); err != nil {
			return fmt.Errorf("error setting description: %s", err)
		}
	}
	if result.Tags != nil {
		if err = d.Set("tags", result.Tags); err != nil {
			return fmt.Errorf("error setting tags: %s", err)
		}
	}

	if result.Rules != nil {
		rules := []map[string]interface{}{}
		for _, rulesItem := range result.Rules {
			rulesItemMap := resourceIbmAppConfigSegmentRuleObjectToMap(rulesItem)
			rules = append(rules, rulesItemMap)
		}
		if err = d.Set("rules", rules); err != nil {
			return fmt.Errorf("error setting rules: %s", err)
		}
	}
	if result.CreatedTime != nil {
		if err = d.Set("created_time", result.CreatedTime.String()); err != nil {
			return fmt.Errorf("error setting created_time: %s", err)
		}
	}
	if result.UpdatedTime != nil {
		if err = d.Set("updated_time", result.UpdatedTime.String()); err != nil {
			return fmt.Errorf("error setting updated_time: %s", err)
		}
	}
	if result.Href != nil {
		if err = d.Set("href", result.Href); err != nil {
			return fmt.Errorf("error setting href: %s", err)
		}
	}

	return nil
}

func resourceIbmAppConfigSegmentRuleObjectToMap(ruleObject appconfigurationv1.Rule) map[string]interface{} {
	ruleObjectMap := map[string]interface{}{}

	ruleObjectMap["attribute_name"] = ruleObject.AttributeName
	ruleObjectMap["operator"] = ruleObject.Operator
	ruleObjectMap["values"] = ruleObject.Values

	return ruleObjectMap
}

func resourceIbmAppConfigSegmentUpdate(d *schema.ResourceData, meta interface{}) error {
	if ok := d.HasChanges("name", "tags", "color_code", "description"); ok {
		parts, err := idParts(d.Id())
		if err != nil {
			return nil
		}
		appconfigClient, err := getAppConfigClient(meta, parts[0])
		if err != nil {
			return err
		}

		options := &appconfigurationv1.UpdateSegmentOptions{}

		options.SetSegmentID(parts[1])
		options.SetName(d.Get("name").(string))

		if _, ok := d.GetOk("description"); ok {
			options.SetDescription(d.Get("description").(string))
		}
		if _, ok := d.GetOk("tags"); ok {
			options.SetTags(d.Get("tags").(string))
		}
		var rules []appconfigurationv1.Rule
		for _, e := range d.Get("rules").([]interface{}) {
			value := e.(map[string]interface{})
			rulesItem := resourceIbmAppConfigSegmentMapToRuleObject(value)
			rules = append(rules, rulesItem)
		}
		options.SetRules(rules)

		_, response, err := appconfigClient.UpdateSegment(options)
		if err != nil {
			log.Printf("[DEBUG] UpdateSegment failed %s\n%s", err, response)
			return err
		}

		return resourceIbmAppConfigSegmentRead(d, meta)
	}
	return nil
}

func resourceIbmAppConfigSegmentDelete(d *schema.ResourceData, meta interface{}) error {
	parts, err := idParts(d.Id())
	if err != nil {
		return nil
	}

	appconfigClient, err := getAppConfigClient(meta, parts[0])
	if err != nil {
		return err
	}

	options := &appconfigurationv1.DeleteSegmentOptions{}

	options.SetSegmentID(parts[1])

	response, err := appconfigClient.DeleteSegment(options)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		log.Printf("[DEBUG] DeleteSegment failed %s\n%s", err, response)
		return err
	}

	d.SetId("")

	return nil
}
