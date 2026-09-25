// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package appconfiguration

import (
	"fmt"
	"strconv"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
	"github.com/IBM/appconfiguration-go-admin-sdk/appconfigurationv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceIBMAppConfigFeatureRule() *schema.Resource {
	return &schema.Resource{
		Create:   resourceIbmAppConfigFeatureRuleCreate,
		Read:     resourceIbmAppConfigFeatureRuleRead,
		Update:   resourceIbmAppConfigFeatureRuleUpdate,
		Delete:   resourceIbmAppConfigFeatureRuleDelete,
		Importer: &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"guid": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "GUID of the App Configuration service. Get it from the service instance credentials section of the dashboard.",
			},
			"environment_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Environment Id.",
			},
			"feature_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Feature Id.",
			},
			"rule_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Rule Id. Allowed special characters are dot ( . ), hyphen( - ), underscore ( _ ) only.",
			},
			"rule_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Name assigned to the rule.",
			},
			"rules": {
				Type:        schema.TypeList,
				Required:    true,
				Description: "List of targeted segments for this rule.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"segments": {
							Type:        schema.TypeList,
							Required:    true,
							Description: "List of segment ids that are used for targeting.",
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
					},
				},
			},
			"value": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Value to be used for evaluation for this rule. The value can be BOOLEAN, STRING or a NUMERIC value as per the feature type.",
			},
			"order": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Order of the rule used during evaluation.",
			},
			"rollout_percentage": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Rollout percentage associated with the rule.",
			},
			"rollout_type": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validate.InvokeValidator("ibm_app_config_feature_rule", "rollout_type"),
				Description:  "The rollout strategy type for the rule. MANUAL is the default. PROGRESSIVE enables automatic phase-based rollout.",
			},
			"rollout_configuration": {
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Description: "Configuration that controls the rollout behaviour for a Progressive rollout type.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"duration_preset": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "A predefined duration preset that sets the overall pace of the rollout. Use CUSTOM to define phases manually.",
						},
						"start_at": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The UTC date and time at which the rollout should start, in ISO 8601 format.",
						},
						"status": {
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "The current execution status of a rollout (QUEUED, RUNNING, STOPPED).",
						},
						"phases": {
							Type:        schema.TypeList,
							Required:    true,
							Description: "The ordered list of rollout phases.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"percentage": {
										Type:        schema.TypeInt,
										Required:    true,
										Description: "The rollout percentage target for this phase.",
									},
									"duration": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "The length of time to wait before advancing to the next phase.",
									},
									"duration_type": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The unit of time for the duration field (minutes, hours, days).",
									},
								},
							},
						},
					},
				},
			},
			"href": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Feature flag rule URL.",
			},
		},
	}
}

func ResourceIBMAppConfigFeatureRuleValidator() *validate.ResourceValidator {
	validateSchema := make([]validate.ValidateSchema, 0)
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 "rollout_type",
			ValidateFunctionIdentifier: validate.ValidateAllowedStringValue,
			Type:                       validate.TypeString,
			Optional:                   true,
			AllowedValues:              "MANUAL, PROGRESSIVE",
		},
	)
	resourceValidator := validate.ResourceValidator{
		ResourceName: "ibm_app_config_feature_rule",
		Schema:       validateSchema,
	}
	return &resourceValidator
}

func resourceIbmAppConfigFeatureRuleCreate(d *schema.ResourceData, meta interface{}) error {
	guid := d.Get("guid").(string)
	appconfigClient, err := getAppConfigClient(meta, guid)
	if err != nil {
		return flex.FmtErrorf("%s", err)
	}

	environmentID := d.Get("environment_id").(string)
	featureID := d.Get("feature_id").(string)

	// Build segments rules list
	targetSegments, err := resourceIbmAppConfigFeatureRuleMapToTargetSegments(d.Get("rules").([]interface{}))
	if err != nil {
		return err
	}

	// Parse value — we need the feature type to format correctly; read it from the feature
	featureType, featureFormat, err := getFeatureTypeAndFormat(appconfigClient, environmentID, featureID)
	if err != nil {
		return err
	}
	value, err := formatValue(featureType, featureFormat, d.Get("value").(string))
	if err != nil {
		return err
	}

	options := &appconfigurationv1.CreateFeatureRuleOptions{}
	options.SetEnvironmentID(environmentID)
	options.SetFeatureID(featureID)
	options.SetRuleID(d.Get("rule_id").(string))
	options.SetRules(targetSegments)
	options.SetValue(value)

	if _, ok := GetFieldExists(d, "rule_name"); ok {
		options.SetRuleName(d.Get("rule_name").(string))
	}
	if _, ok := GetFieldExists(d, "rollout_percentage"); ok {
		options.SetRolloutPercentage(int64(d.Get("rollout_percentage").(int)))
	}
	if _, ok := GetFieldExists(d, "rollout_type"); ok {
		options.SetRolloutType(d.Get("rollout_type").(string))
	}
	if _, ok := GetFieldExists(d, "rollout_configuration"); ok {
		rc, err := resourceIbmAppConfigFeatureMapToRolloutConfiguration(d.Get("rollout_configuration").([]interface{}))
		if err != nil {
			return err
		}
		if rc != nil {
			options.SetRolloutConfiguration(rc)
		}
	}

	result, response, err := appconfigClient.CreateFeatureRule(options)
	if err != nil {
		return flex.FmtErrorf("[ERROR] CreateFeatureRule failed %s\n%s", err, response)
	}

	d.SetId(fmt.Sprintf("%s/%s/%s/%s", guid, environmentID, featureID, *result.RuleID))
	return resourceIbmAppConfigFeatureRuleRead(d, meta)
}

func resourceIbmAppConfigFeatureRuleRead(d *schema.ResourceData, meta interface{}) error {
	parts, err := flex.IdParts(d.Id())
	if err != nil {
		return nil
	}
	if len(parts) < 4 {
		return flex.FmtErrorf("[ERROR] Invalid resource ID, expected guid/environment_id/feature_id/rule_id")
	}
	guid, environmentID, featureID, ruleID := parts[0], parts[1], parts[2], parts[3]

	appconfigClient, err := getAppConfigClient(meta, guid)
	if err != nil {
		return flex.FmtErrorf("%s", err)
	}

	options := &appconfigurationv1.GetFeatureRuleOptions{}
	options.SetEnvironmentID(environmentID)
	options.SetFeatureID(featureID)
	options.SetRuleID(ruleID)

	result, response, err := appconfigClient.GetFeatureRule(options)
	if err != nil {
		return flex.FmtErrorf("[ERROR] GetFeatureRule failed %s\n%s", err, response)
	}

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
		if err = d.Set("rollout_configuration", resourceIbmAppConfigFeatureFlattenRolloutConfiguration(result.RolloutConfiguration)); err != nil {
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

func resourceIbmAppConfigFeatureRuleUpdate(d *schema.ResourceData, meta interface{}) error {
	parts, err := flex.IdParts(d.Id())
	if err != nil {
		return nil
	}
	if len(parts) < 4 {
		return flex.FmtErrorf("[ERROR] Invalid resource ID, expected guid/environment_id/feature_id/rule_id")
	}
	guid, environmentID, featureID, ruleID := parts[0], parts[1], parts[2], parts[3]

	appconfigClient, err := getAppConfigClient(meta, guid)
	if err != nil {
		return flex.FmtErrorf("%s", err)
	}

	if !d.HasChanges("rules", "value", "rule_name", "rollout_percentage", "rollout_type", "rollout_configuration") {
		return nil
	}

	options := &appconfigurationv1.UpdateFeatureRuleOptions{}
	options.SetEnvironmentID(environmentID)
	options.SetFeatureID(featureID)
	options.SetRuleID(ruleID)

	if d.HasChange("rules") {
		targetSegments, err := resourceIbmAppConfigFeatureRuleMapToTargetSegments(d.Get("rules").([]interface{}))
		if err != nil {
			return err
		}
		options.SetRules(targetSegments)
	}

	if d.HasChange("value") {
		featureType, featureFormat, err := getFeatureTypeAndFormat(appconfigClient, environmentID, featureID)
		if err != nil {
			return err
		}
		value, err := formatValue(featureType, featureFormat, d.Get("value").(string))
		if err != nil {
			return err
		}
		options.SetValue(value)
	}

	if _, ok := GetFieldExists(d, "rule_name"); ok {
		options.SetRuleName(d.Get("rule_name").(string))
	}
	if _, ok := GetFieldExists(d, "rollout_percentage"); ok {
		options.SetRolloutPercentage(int64(d.Get("rollout_percentage").(int)))
	}
	if _, ok := GetFieldExists(d, "rollout_type"); ok {
		options.SetRolloutType(d.Get("rollout_type").(string))
	}
	if _, ok := GetFieldExists(d, "rollout_configuration"); ok {
		rc, err := resourceIbmAppConfigFeatureMapToRolloutConfiguration(d.Get("rollout_configuration").([]interface{}))
		if err != nil {
			return err
		}
		if rc != nil {
			options.SetRolloutConfiguration(rc)
		}
	}

	_, response, err := appconfigClient.UpdateFeatureRule(options)
	if err != nil {
		return flex.FmtErrorf("[ERROR] UpdateFeatureRule failed %s\n%s", err, response)
	}

	return resourceIbmAppConfigFeatureRuleRead(d, meta)
}

func resourceIbmAppConfigFeatureRuleDelete(d *schema.ResourceData, meta interface{}) error {
	parts, err := flex.IdParts(d.Id())
	if err != nil {
		return nil
	}
	if len(parts) < 4 {
		return flex.FmtErrorf("[ERROR] Invalid resource ID, expected guid/environment_id/feature_id/rule_id")
	}
	guid, environmentID, featureID, ruleID := parts[0], parts[1], parts[2], parts[3]

	appconfigClient, err := getAppConfigClient(meta, guid)
	if err != nil {
		return flex.FmtErrorf("%s", err)
	}

	options := &appconfigurationv1.DeleteFeatureRuleOptions{}
	options.SetEnvironmentID(environmentID)
	options.SetFeatureID(featureID)
	options.SetRuleID(ruleID)

	_, response, err := appconfigClient.DeleteFeatureRule(options)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		return flex.FmtErrorf("[ERROR] DeleteFeatureRule failed %s\n%s", err, response)
	}

	d.SetId("")
	return nil
}

// getFeatureTypeAndFormat fetches the type and format of a feature, used to correctly format rule values.
func getFeatureTypeAndFormat(client *appconfigurationv1.AppConfigurationV1, environmentID, featureID string) (string, interface{}, error) {
	opts := &appconfigurationv1.GetFeatureOptions{}
	opts.SetEnvironmentID(environmentID)
	opts.SetFeatureID(featureID)

	feature, response, err := client.GetFeature(opts)
	if err != nil {
		return "", nil, flex.FmtErrorf("[ERROR] GetFeature (for rule value formatting) failed %s\n%s", err, response)
	}

	var featureFormat interface{} = nil
	if feature.Format != nil {
		featureFormat = *feature.Format
	}
	return *feature.Type, featureFormat, nil
}

func resourceIbmAppConfigFeatureRuleMapToTargetSegments(rulesList []interface{}) ([]appconfigurationv1.TargetSegments, error) {
	result := []appconfigurationv1.TargetSegments{}
	for _, r := range rulesList {
		ruleMap := r.(map[string]interface{})
		target := appconfigurationv1.TargetSegments{}
		segments := []string{}
		for _, s := range ruleMap["segments"].([]interface{}) {
			segments = append(segments, s.(string))
		}
		target.Segments = segments
		result = append(result, target)
	}
	return result, nil
}

func resourceIbmAppConfigFeatureRuleWithIDToMap(rule appconfigurationv1.FeatureSegmentRuleWithRuleID) map[string]interface{} {
	ruleMap := map[string]interface{}{}

	if rule.RuleID != nil {
		ruleMap["rule_id"] = rule.RuleID
	}
	if rule.RuleName != nil {
		ruleMap["rule_name"] = rule.RuleName
	}
	if rule.Order != nil {
		ruleMap["order"] = flex.IntValue(rule.Order)
	}
	if rule.RolloutPercentage != nil {
		ruleMap["rollout_percentage"] = flex.IntValue(rule.RolloutPercentage)
	}
	if rule.RolloutType != nil {
		ruleMap["rollout_type"] = rule.RolloutType
	}
	if rule.RolloutConfiguration != nil {
		ruleMap["rollout_configuration"] = resourceIbmAppConfigFeatureFlattenRolloutConfiguration(rule.RolloutConfiguration)
	}
	if rule.Href != nil {
		ruleMap["href"] = rule.Href
	}

	if rule.Rules != nil {
		segments := []map[string]interface{}{}
		for _, r := range rule.Rules {
			segMap := map[string]interface{}{}
			if r.Segments != nil {
				segMap["segments"] = r.Segments
			}
			segments = append(segments, segMap)
		}
		ruleMap["rules"] = segments
	}

	if rule.Value != nil {
		switch v := rule.Value.(type) {
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
