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

func DataSourceIBMAppConfigFeatureRule() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceIbmAppConfigFeatureRuleRead,

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
			"rule_id": {
				Type:        schema.TypeString,
				Required:    true,
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
	}
}

func dataSourceIbmAppConfigFeatureRuleRead(d *schema.ResourceData, meta interface{}) error {
	guid := d.Get("guid").(string)

	appconfigClient, err := getAppConfigClient(meta, guid)
	if err != nil {
		return flex.FmtErrorf("%s", err)
	}

	options := &appconfigurationv1.GetFeatureRuleOptions{}
	options.SetEnvironmentID(d.Get("environment_id").(string))
	options.SetFeatureID(d.Get("feature_id").(string))
	options.SetRuleID(d.Get("rule_id").(string))

	result, response, err := appconfigClient.GetFeatureRule(options)
	if err != nil {
		return flex.FmtErrorf("[ERROR] GetFeatureRule failed %s\n%s", err, response)
	}

	d.SetId(fmt.Sprintf("%s/%s/%s/%s", guid, *options.EnvironmentID, *options.FeatureID, *result.RuleID))

	if result.RuleID != nil {
		if err = d.Set("rule_id", result.RuleID); err != nil {
			return flex.FmtErrorf("[ERROR] Error setting rule_id: %s", err)
		}
	}
	if result.RuleName != nil {
		if err = d.Set("rule_name", result.RuleName); err != nil {
			return flex.FmtErrorf("[ERROR] Error setting rule_name: %s", err)
		}
	}
	if result.Order != nil {
		if err = d.Set("order", result.Order); err != nil {
			return flex.FmtErrorf("[ERROR] Error setting order: %s", err)
		}
	}
	if result.RolloutPercentage != nil {
		if err = d.Set("rollout_percentage", result.RolloutPercentage); err != nil {
			return flex.FmtErrorf("[ERROR] Error setting rollout_percentage: %s", err)
		}
	}
	if result.RolloutType != nil {
		if err = d.Set("rollout_type", result.RolloutType); err != nil {
			return flex.FmtErrorf("[ERROR] Error setting rollout_type: %s", err)
		}
	}
	if result.RolloutConfiguration != nil {
		if err = d.Set("rollout_configuration", dataSourceFeatureFlattenRolloutConfiguration(result.RolloutConfiguration)); err != nil {
			return flex.FmtErrorf("[ERROR] Error setting rollout_configuration: %s", err)
		}
	}
	if result.Href != nil {
		if err = d.Set("href", result.Href); err != nil {
			return flex.FmtErrorf("[ERROR] Error setting href: %s", err)
		}
	}

	if result.Rules != nil {
		rules := []map[string]interface{}{}
		for _, r := range result.Rules {
			rulesMap := map[string]interface{}{}
			if r.Segments != nil {
				rulesMap["segments"] = r.Segments
			}
			rules = append(rules, rulesMap)
		}
		if err = d.Set("rules", rules); err != nil {
			return flex.FmtErrorf("[ERROR] Error setting rules: %s", err)
		}
	}

	if result.Value != nil {
		switch v := result.Value.(type) {
		case string:
			d.Set("value", v)
		case float64:
			d.Set("value", fmt.Sprintf("%v", v))
		case bool:
			d.Set("value", strconv.FormatBool(v))
		}
	}

	return nil
}
