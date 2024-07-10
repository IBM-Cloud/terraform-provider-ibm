// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package scc

import (
	"fmt"
	"strings"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/scc-go-sdk/v5/securityandcompliancecenterapiv3"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	INSTANCE_ID               = "instance_id"
	MAX_REQUIRED_CONFIG_DEPTH = 3
)

// AddSchemaData will add the Schemas 'instance_id' and 'region' to the resource
func AddSchemaData(resource *schema.Resource) *schema.Resource {
	resource.Schema["instance_id"] = &schema.Schema{
		Type:        schema.TypeString,
		Required:    true,
		ForceNew:    true,
		Description: "The ID of the Security and Compliance Center instance.",
	}
	return resource
}

// getRegionData will check if the field region is defined
func getRegionData(client securityandcompliancecenterapiv3.SecurityAndComplianceCenterApiV3, d *schema.ResourceData) string {
	val, ok := d.GetOk("region")
	if ok {
		return val.(string)
	} else {
		url := client.Service.GetServiceURL()
		return strings.Split(url, ".")[1]
	}
}

// setRegionData will set the field "region" field if the field was previously defined
func setRegionData(d *schema.ResourceData, region string) error {
	if val, ok := d.GetOk("region"); ok {
		return d.Set("region", val.(string))
	}
	return nil
}

func getRequiredConfigSchema(currentDepth int) map[string]*schema.Schema {
	baseMap := map[string]*schema.Schema{
		"description": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Description: "The programmatic name of the IBM Cloud service that you want to target with the rule or template.",
		},
		"property": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Description: "The name of the additional attribute that you want to use to further qualify the target.Options differ depending on the service or resource that you are targeting with a rule or template. For more information, refer to the service documentation.",
		},
		"value": &schema.Schema{
			Type:        schema.TypeString,
			Optional:    true,
			Default:     "",
			Description: "The value that you want to apply to `name` field.Options differ depending on the rule or template that you configure. For more information, refer to the service documentation.",
		},
		"operator": &schema.Schema{
			Type:         schema.TypeString,
			Optional:     true,
			Description:  "The way in which the `name` field is compared to its value.There are three types of operators: string, numeric, and boolean.",
			ValidateFunc: validate.InvokeValidator("ibm_scc_rule", "operator"),
		},
	}
	if currentDepth > MAX_REQUIRED_CONFIG_DEPTH {
		return baseMap
	}
	baseMap["and"] = &schema.Schema{
		Type:        schema.TypeList,
		Optional:    true,
		Description: "A condition with the and logical operator.",
		Elem: &schema.Resource{
			Schema: getRequiredConfigSchema(currentDepth + 1),
		},
	}

	baseMap["or"] = &schema.Schema{
		Type:        schema.TypeList,
		Optional:    true,
		Description: "A condition with the or logical operator.",
		Elem: &schema.Resource{
			Schema: getRequiredConfigSchema(currentDepth + 1),
		},
	}

	baseMap["all"] = &schema.Schema{
		Type:        schema.TypeList,
		Optional:    true,
		Description: "A condition with the SubRule all logical operator.",
		Elem: &schema.Resource{
			Schema: getSubRuleSchema(currentDepth + 1),
		},
	}

	baseMap["all_if"] = &schema.Schema{
		Type:        schema.TypeList,
		Optional:    true,
		Description: "A condition with the SubRule all_ifexists logical operator.",
		Elem: &schema.Resource{
			Schema: getSubRuleSchema(currentDepth + 1),
		},
	}

	baseMap["any"] = &schema.Schema{
		Type:        schema.TypeList,
		Optional:    true,
		Description: "A condition with the SubRule any logical operator.",
		Elem: &schema.Resource{
			Schema: getSubRuleSchema(currentDepth + 1),
		},
	}

	baseMap["any_if"] = &schema.Schema{
		Type:        schema.TypeList,
		Optional:    true,
		Description: "A condition with the SubRule any_ifexists logical operator.",
		Elem: &schema.Resource{
			Schema: getSubRuleSchema(currentDepth + 1),
		},
	}
	return baseMap
}

func getTargetSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"service_name": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The target service name.",
		},
		"service_display_name": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "The display name of the target service.",
			// Manual Intervention
			DiffSuppressFunc: func(_, oldVal, newVal string, d *schema.ResourceData) bool {
				if newVal == "" {
					return true
				}
				if strings.ToLower(oldVal) == strings.ToLower(newVal) {
					return true
				}
				return false
			},
			// End Manual Intervention
		},
		"resource_kind": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The target resource kind.",
		},
		"additional_target_attributes": {
			Type:        schema.TypeList,
			Optional:    true,
			Description: "The list of targets supported properties.",
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"name": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "The additional target attribute name.",
					},
					"operator": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "The operator.",
					},
					"value": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "The value.",
					},
				},
			},
		},
	}
}

func getSubRuleSchema(currentDepth int) map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"required_config": {
			Description: "The requirements that must be met to determine the resource's level of compliance in accordance with the rule. Use logical operators (and/or) to define multiple property checks and conditions. To define requirements for a rule, list one or more property check objects in the and array. To add conditions to a property check, use or.",
			Type:        schema.TypeList,
			Required:    true,
			Elem: &schema.Resource{
				Schema: getRequiredConfigSchema(currentDepth + 1),
			},
			MaxItems: 1,
		},
		"target": {
			Description: "The requirements that must be met to determine the resource's level of compliance in accordance with the rule. Use logical operators (and/or) to define multiple property checks and conditions. To define requirements for a rule, list one or more property check objects in the and array. To add conditions to a property check, use or.",
			Type:        schema.TypeList,
			Required:    true,
			Elem: &schema.Resource{
				Schema: getTargetSchema(),
			},
			MaxItems: 1,
		},
	}
}

func ibmSccRuleRequiredConfigToMap(model securityandcompliancecenterapiv3.RequiredConfigIntf) (map[string]interface{}, error) {
	if rc, ok := model.(*securityandcompliancecenterapiv3.RequiredConfig); ok {
		modelMap := make(map[string]interface{})
		fmt.Printf("%#v\n", rc)
		if rc.Description != nil {
			modelMap["description"] = rc.Description
		}
		if rc.And != nil {
			rcItems := []map[string]interface{}{}
			for _, rcItem := range rc.And {
				rcMap, err := ibmSccRuleRequiredConfigToMap(rcItem)
				if err != nil {
					return map[string]interface{}{}, err
				}
				rcItems = append(rcItems, rcMap)
			}
		}
		if rc.Or != nil {
			rcItems := []map[string]interface{}{}
			for _, rcItem := range rc.Or {
				rcMap, err := ibmSccRuleRequiredConfigToMap(rcItem)
				if err != nil {
					return map[string]interface{}{}, err
				}
				rcItems = append(rcItems, rcMap)
			}
		}
		if rc.Property != nil {
			modelMap["property"] = rc.Property
		}
		if rc.Operator != nil {
			modelMap["operator"] = rc.Operator
		}
		if rc.Value != nil {
			modelMap["value"] = rc.Value
		}
		return modelMap, nil
	} else {
		return nil, fmt.Errorf("Unrecognized securityandcompliancecenterapiv3.RequiredConfigIntf subtype encountered %#v", model)
	}
}

func ibmSccRCMapToRequiredConfig(modelMap map[string]interface{}) (securityandcompliancecenterapiv3.RequiredConfigIntf, error) {
	model := &securityandcompliancecenterapiv3.RequiredConfig{}
	if modelMap["description"] != nil && modelMap["description"].(string) != "" {
		model.Description = core.StringPtr(modelMap["description"].(string))
	}
	if modelMap["or"] != nil {
		or := []securityandcompliancecenterapiv3.RequiredConfigIntf{}
		for _, orItem := range modelMap["or"].([]interface{}) {
			orItemModel, err := ibmSccRCMapToRequiredConfig(orItem.(map[string]interface{}))
			if err != nil {
				return model, err
			}
			or = append(or, orItemModel)
		}
		model.Or = or
	}
	if modelMap["and"] != nil {
		and := []securityandcompliancecenterapiv3.RequiredConfigIntf{}
		for _, andItem := range modelMap["and"].([]interface{}) {
			andItemModel, err := ibmSccRCMapToRequiredConfig(andItem.(map[string]interface{}))
			if err != nil {
				return model, err
			}
			and = append(and, andItemModel)
		}
		model.And = and
	}
	if modelMap["any"] != nil {
		anySubRule := securityandcompliancecenterapiv3.RequiredConfigSubRule{}
		anyCondition := modelMap["any"].(map[string]interface{})
		target, err := ibmSccTargetMapToTarget(anyCondition["target"].(map[string]interface{}))
		if err != nil {
			return &anySubRule, err
		}
		anySubRule.Target = target
		rc, err := ibmSccRCMapToRequiredConfig(anyCondition["required_config"].(map[string]interface{}))
		if err != nil {
			return &anySubRule, err
		}
		anySubRule.RequiredConfig = rc.(*securityandcompliancecenterapiv3.RequiredConfig)
		model.Any = &anySubRule
	}
	if modelMap["any_if"] != nil {
		anyIfSubRule := securityandcompliancecenterapiv3.RequiredConfigSubRule{}
		anyIfCondition := modelMap["any_if"].(map[string]interface{})
		target, err := ibmSccTargetMapToTarget(anyIfCondition["target"].(map[string]interface{}))
		if err != nil {
			return &anyIfSubRule, err
		}
		anyIfSubRule.Target = target
		rc, err := ibmSccRCMapToRequiredConfig(anyIfCondition["required_config"].(map[string]interface{}))
		if err != nil {
			return &anyIfSubRule, err
		}
		anyIfSubRule.RequiredConfig = rc.(*securityandcompliancecenterapiv3.RequiredConfig)
		model.AnyIf = &anyIfSubRule
	}
	if modelMap["all"] != nil {
		allSubRule := securityandcompliancecenterapiv3.RequiredConfigSubRule{}
		anyCondition := modelMap["all"].(map[string]interface{})
		target, err := ibmSccTargetMapToTarget(anyCondition["target"].(map[string]interface{}))
		if err != nil {
			return &allSubRule, err
		}
		allSubRule.Target = target
		rc, err := ibmSccRCMapToRequiredConfig(anyCondition["required_config"].(map[string]interface{}))
		if err != nil {
			return &allSubRule, err
		}
		allSubRule.RequiredConfig = rc.(*securityandcompliancecenterapiv3.RequiredConfig)
		model.Any = &allSubRule
	}
	if modelMap["property"] != nil && modelMap["property"].(string) != "" {
		model.Property = core.StringPtr(modelMap["property"].(string))
	}
	if modelMap["operator"] != nil && modelMap["operator"].(string) != "" {
		model.Operator = core.StringPtr(modelMap["operator"].(string))
	}

	if modelMap["value"] != nil && len(modelMap["value"].(string)) > 0 {
		// model.Value = modelMap["value"].(string)
		sLit := strings.Trim(modelMap["value"].(string), "[]")
		sList := strings.Split(sLit, ",")
		if len(sList) == 1 {
			model.Value = modelMap["value"].(string)
		} else {
			model.Value = sList
		}
	}

	return model, nil
}

func ibmSccTargetMapToTarget(modelMap map[string]interface{}) (*securityandcompliancecenterapiv3.Target, error) {
	model := &securityandcompliancecenterapiv3.Target{}
	model.ServiceName = core.StringPtr(modelMap["service_name"].(string))
	if modelMap["service_display_name"] != nil && modelMap["service_display_name"].(string) != "" {
		model.ServiceDisplayName = core.StringPtr(modelMap["service_display_name"].(string))
	}
	model.ResourceKind = core.StringPtr(modelMap["resource_kind"].(string))
	if modelMap["additional_target_attributes"] != nil {
		additionalTargetAttributes := []securityandcompliancecenterapiv3.AdditionalTargetAttribute{}
		for _, additionalTargetAttributesItem := range modelMap["additional_target_attributes"].([]interface{}) {
			additionalTargetAttributesItemModel, err := resourceIbmSccRuleMapToAdditionalTargetAttribute(additionalTargetAttributesItem.(map[string]interface{}))
			if err != nil {
				return model, err
			}
			additionalTargetAttributes = append(additionalTargetAttributes, *additionalTargetAttributesItemModel)
		}
		model.AdditionalTargetAttributes = additionalTargetAttributes
	}
	return model, nil
}

func ibmSccRuleTargetToMap(model *securityandcompliancecenterapiv3.Target) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})

	modelMap["service_name"] = model.ServiceName

	modelMap["resource_kind"] = model.ResourceKind

	if model.ServiceDisplayName != nil {
		modelMap["service_display_name"] = model.ServiceDisplayName
	}

	if model.AdditionalTargetAttributes != nil {
		additionalTargetAttributes := []map[string]interface{}{}
		for _, additionalTargetAttributesItem := range model.AdditionalTargetAttributes {
			additionalTargetAttributesItemMap, err := dataSourceIbmSccRuleAdditionalTargetAttributeToMap(&additionalTargetAttributesItem)
			if err != nil {
				return modelMap, err
			}
			additionalTargetAttributes = append(additionalTargetAttributes, additionalTargetAttributesItemMap)
		}
		modelMap["additional_target_attributes"] = additionalTargetAttributes
	}
	return modelMap, nil
}
