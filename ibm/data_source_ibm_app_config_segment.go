// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM/appconfiguration-go-admin-sdk/appconfigurationv1"
)

func dataSourceIbmAppConfigSegment() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceIbmAppConfigSegmentRead,

		Schema: map[string]*schema.Schema{
			"guid": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "GUID of the App Configuration service. Get it from the service instance credentials section of the dashboard.",
			},
			"segment_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Segment Id.",
			},
			"includes": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Include feature and property details in the response.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Segment name.",
			},
			"description": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Segment description.",
			},
			"tags": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Tags associated with the segments.",
			},
			"rules": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of rules that determine if the entity is part of the segment.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"attribute_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Attribute name.",
						},
						"operator": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Operator to be used for the evaluation if the entity is part of the segment.",
						},
						"values": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "List of values. Entities matching any of the given values will be considered to be part of the segment.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
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

func dataSourceIbmAppConfigSegmentRead(d *schema.ResourceData, meta interface{}) error {
	guid := d.Get("guid").(string)

	appconfigClient, err := getAppConfigClient(meta, guid)
	if err != nil {
		return err
	}

	options := &appconfigurationv1.GetSegmentOptions{}

	options.SetSegmentID(d.Get("segment_id").(string))

	if _, ok := d.GetOk("includes"); ok {
		includes := []string{}
		for _, segmentsItem := range d.Get("includes").([]interface{}) {
			includes = append(includes, segmentsItem.(string))
		}
		options.SetInclude(includes)
	}

	result, response, err := appconfigClient.GetSegment(options)
	if err != nil {
		log.Printf("[DEBUG] GetSegment failed %s\n%s", err, response)
		return err
	}

	d.SetId(fmt.Sprintf("%s/%s", guid, *result.SegmentID))

	if result.Name != nil {
		if err = d.Set("name", result.Name); err != nil {
			return fmt.Errorf("error setting name: %s", err)
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
		err = d.Set("rules", dataSourceSegmentFlattenRules(result.Rules))
		if err != nil {
			return fmt.Errorf("error setting rules %s", err)
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

func dataSourceSegmentFlattenRules(result []appconfigurationv1.Rule) (rules []map[string]interface{}) {
	for _, rulesItem := range result {
		rules = append(rules, dataSourceSegmentRulesToMap(rulesItem))
	}

	return rules
}

func dataSourceSegmentRulesToMap(rulesItem appconfigurationv1.Rule) (rulesMap map[string]interface{}) {
	rulesMap = map[string]interface{}{}

	if rulesItem.AttributeName != nil {
		rulesMap["attribute_name"] = rulesItem.AttributeName
	}
	if rulesItem.Operator != nil {
		rulesMap["operator"] = rulesItem.Operator
	}
	if rulesItem.Values != nil {
		rulesMap["values"] = rulesItem.Values
	}

	return rulesMap
}
