// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package appconfiguration

import (
	"fmt"
	"strconv"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM/appconfiguration-go-admin-sdk/appconfigurationv1"
)

func DataSourceIBMAppConfigFeatureRules() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceIbmAppConfigFeatureRulesRead,

		Schema: map[string]*schema.Schema{
			"guid": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "GUID of the App Configuration service. Get it from the service instance credentials section of the dashboard.",
			},
			"environment_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Environment Id.",
			},
			"feature_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Feature Id.",
			},
			"rules": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Array of feature rules.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"rule_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Rule Id.",
						},
						"rule_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Name assigned to the rule.",
						},
						"rules": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "List of targeted segments for this rule.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"segments": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "List of segment ids that are used for targeting.",
										Elem:        &schema.Schema{Type: schema.TypeString},
									},
								},
							},
						},
						"value": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Value to be used for evaluation for this rule.",
						},
						"order": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Order of the rule used during evaluation.",
						},
						"rollout_percentage": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Rollout percentage associated with the rule.",
						},
						"rollout_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The rollout strategy type for the rule. MANUAL is the default. PROGRESSIVE enables automatic phase-based rollout.",
						},
						"rollout_configuration": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Configuration that controls the rollout behaviour for a Progressive rollout type.",
							Elem:        rolloutConfigurationSchema(),
						},
						"href": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Feature flag rule URL.",
						},
					},
				},
			},
			"total_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Total number of rules.",
			},
			"limit": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Number of records returned in the current response.",
			},
			"offset": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Offset of the records returned.",
			},
			"first": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "URL to navigate to the first page of records.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"href": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "URL of the response.",
						},
					},
				},
			},
			"next": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "URL to navigate to the next list of records.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"href": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "URL of the response.",
						},
					},
				},
			},
			"previous": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "URL to navigate to the previous list of records.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"href": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "URL of the response.",
						},
					},
				},
			},
			"last": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "URL to navigate to the last page of records.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"href": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "URL of the response.",
						},
					},
				},
			},
		},
	}
}

func dataSourceIbmAppConfigFeatureRulesRead(d *schema.ResourceData, meta interface{}) error {
	guid := d.Get("guid").(string)

	appconfigClient, err := getAppConfigClient(meta, guid)
	if err != nil {
		return flex.FmtErrorf("%s", err)
	}

	options := &appconfigurationv1.ListFeatureRulesOptions{}
	options.SetEnvironmentID(d.Get("environment_id").(string))
	options.SetFeatureID(d.Get("feature_id").(string))

	result, response, err := appconfigClient.ListFeatureRules(options)
	if err != nil {
		return flex.FmtErrorf("[ERROR] ListFeatureRules failed %s\n%s", err, response)
	}

	d.SetId(fmt.Sprintf("%s/%s/%s", guid, *options.EnvironmentID, *options.FeatureID))

	if result.TotalCount != nil {
		if err = d.Set("total_count", result.TotalCount); err != nil {
			return flex.FmtErrorf("[ERROR] Error setting total_count: %s", err)
		}
	}
	if result.Limit != nil {
		if err = d.Set("limit", result.Limit); err != nil {
			return flex.FmtErrorf("[ERROR] Error setting limit: %s", err)
		}
	}
	if result.Offset != nil {
		if err = d.Set("offset", result.Offset); err != nil {
			return flex.FmtErrorf("[ERROR] Error setting offset: %s", err)
		}
	}
	if result.First != nil {
		if err = d.Set("first", dataSourceFeatureListFlattenPagination(*result.First)); err != nil {
			return flex.FmtErrorf("[ERROR] Error setting first: %s", err)
		}
	}
	if result.Previous != nil {
		if err = d.Set("previous", dataSourceFeatureListFlattenPagination(*result.Previous)); err != nil {
			return flex.FmtErrorf("[ERROR] Error setting previous: %s", err)
		}
	}
	if result.Next != nil {
		if err = d.Set("next", dataSourceFeatureListFlattenPagination(*result.Next)); err != nil {
			return flex.FmtErrorf("[ERROR] Error setting next: %s", err)
		}
	}
	if result.Last != nil {
		if err = d.Set("last", dataSourceFeatureListFlattenPagination(*result.Last)); err != nil {
			return flex.FmtErrorf("[ERROR] Error setting last: %s", err)
		}
	}

	rules := []map[string]interface{}{}
	for _, item := range result.SegmentRules {
		rules = append(rules, dataSourceFeatureRuleWithRuleIDToMap(item))
	}
	if err = d.Set("rules", rules); err != nil {
		return flex.FmtErrorf("[ERROR] Error setting rules: %s", err)
	}

	return nil
}

func dataSourceFeatureRuleWithRuleIDToMap(item appconfigurationv1.FeatureSegmentRuleWithRuleID) map[string]interface{} {
	ruleMap := map[string]interface{}{}

	if item.RuleID != nil {
		ruleMap["rule_id"] = item.RuleID
	}
	if item.RuleName != nil {
		ruleMap["rule_name"] = item.RuleName
	}
	if item.Order != nil {
		ruleMap["order"] = item.Order
	}
	if item.RolloutPercentage != nil {
		ruleMap["rollout_percentage"] = item.RolloutPercentage
	}
	if item.RolloutType != nil {
		ruleMap["rollout_type"] = item.RolloutType
	}
	if item.RolloutConfiguration != nil {
		ruleMap["rollout_configuration"] = dataSourceFeatureFlattenRolloutConfiguration(item.RolloutConfiguration)
	}
	if item.Href != nil {
		ruleMap["href"] = item.Href
	}

	if item.Rules != nil {
		targetSegments := []map[string]interface{}{}
		for _, r := range item.Rules {
			segMap := map[string]interface{}{}
			if r.Segments != nil {
				segMap["segments"] = r.Segments
			}
			targetSegments = append(targetSegments, segMap)
		}
		ruleMap["rules"] = targetSegments
	}

	if item.Value != nil {
		switch v := item.Value.(type) {
		case string:
			ruleMap["value"] = v
		case float64:
			ruleMap["value"] = fmt.Sprintf("%v", v)
		case bool:
			ruleMap["value"] = strconv.FormatBool(v)
		}
	}

	return ruleMap
}
