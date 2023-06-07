// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package iampolicy

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/iampolicymanagementv1"
)

func DataSourceIBMIAMPolicyTemplate() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMPolicyTemplateRead,

		Schema: map[string]*schema.Schema{
			"account_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The account GUID that the policy templates belong to.",
			},
			"policy_templates": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of policy templates.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "name of template.",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "description of template purpose.",
						},
						"account_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "account id where this template will be created.",
						},
						"version": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Template vesrsion.",
						},
						"committed": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Template vesrsion committed status.",
						},
						"policy": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The core set of properties associated with the template's policy objet.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"type": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The policy type; either 'access' or 'authorization'.",
									},
									"description": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Allows the customer to use their own words to record the purpose/context related to a policy.",
									},
									"resource": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "The resource attributes to which the policy grants access.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"attributes": {
													Type:        schema.TypeList,
													Computed:    true,
													Description: "List of resource attributes to which the policy grants access.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"key": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "The name of a resource attribute.",
															},
															"operator": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "The operator of an attribute.",
															},
															"value": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "The value of a rule or resource attribute; can be boolean or string for resource attribute. Can be string or an array of strings (e.g., array of days to permit access) for rule attribute.",
															},
														},
													},
												},
												"tags": {
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Optional list of resource tags to which the policy grants access.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"key": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "The name of an access management tag.",
															},
															"value": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "The value of an access management tag.",
															},
															"operator": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "The operator of an access management tag.",
															},
														},
													},
												},
											},
										},
									},
									"pattern": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Indicates pattern of rule, either 'time-based-conditions:once', 'time-based-conditions:weekly:all-day', or 'time-based-conditions:weekly:custom-hours'.",
									},
									"rule_conditions": {
										Type:        schema.TypeSet,
										Optional:    true,
										Description: "Rule conditions enforced by the policy",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"key": {
													Type:        schema.TypeString,
													Required:    true,
													Description: "Key of the condition",
												},
												"operator": {
													Type:        schema.TypeString,
													Required:    true,
													Description: "Operator of the condition",
												},
												"value": {
													Type:        schema.TypeList,
													Required:    true,
													Elem:        &schema.Schema{Type: schema.TypeString},
													Description: "Value of the condition",
												},
											},
										},
									},

									"rule_operator": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Operator that multiple rule conditions are evaluated over",
									},
									"roles": {
										Type:        schema.TypeList,
										Required:    true,
										Elem:        &schema.Schema{Type: schema.TypeString},
										Description: "Role names of the policy definition",
									},
								},
							},
						},
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The policy template ID.",
						},
					},
				},
			},
		},
	}
}

func dataSourceIBMPolicyTemplateRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	iamPolicyManagementClient, err := meta.(conns.ClientSession).IAMPolicyManagementV1API()
	if err != nil {
		return diag.FromErr(err)
	}

	listPolicyTemplatesOptions := &iampolicymanagementv1.ListPolicyTemplatesOptions{}

	listPolicyTemplatesOptions.SetAccountID(d.Get("account_id").(string))

	policyTemplateCollection, response, err := iamPolicyManagementClient.ListPolicyTemplatesWithContext(context, listPolicyTemplatesOptions)
	if err != nil {
		log.Printf("[DEBUG] ListPolicyTemplatesWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("ListPolicyTemplatesWithContext failed %s\n%s", err, response))
	}

	d.SetId(d.Get("account_id").(string))

	policyTemplates := []map[string]interface{}{}
	if policyTemplateCollection.PolicyTemplates != nil {
		for _, modelItem := range policyTemplateCollection.PolicyTemplates {
			modelMap, err := dataSourceIBMPolicyTemplatePolicyTemplateToMap(&modelItem, iamPolicyManagementClient)
			if err != nil {
				return diag.FromErr(err)
			}
			policyTemplates = append(policyTemplates, modelMap)
		}
	}
	if err = d.Set("policy_templates", policyTemplates); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting policy_templates %s", err))
	}

	return nil
}

func dataSourceIBMPolicyTemplatePolicyTemplateToMap(model *iampolicymanagementv1.PolicyTemplate, iamPolicyManagementClient *iampolicymanagementv1.IamPolicyManagementV1) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["name"] = model.Name
	if model.Description != nil {
		modelMap["description"] = model.Description
	}
	modelMap["account_id"] = model.AccountID
	modelMap["version"] = model.Version
	if model.Committed != nil {
		modelMap["committed"] = model.Committed
	}
	policyMap, err := dataSourceIBMPolicyTemplateTemplatePolicyToMap(model.Policy, iamPolicyManagementClient)
	if err != nil {
		return modelMap, err
	}
	modelMap["policy"] = []map[string]interface{}{policyMap}
	if model.ID != nil {
		modelMap["id"] = model.ID
	}
	return modelMap, nil
}

func dataSourceIBMPolicyTemplateTemplatePolicyToMap(model *iampolicymanagementv1.TemplatePolicy, iamPolicyManagementClient *iampolicymanagementv1.IamPolicyManagementV1) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["type"] = model.Type
	if model.Description != nil {
		modelMap["description"] = model.Description
	}
	resourceMap, roleList, err := dataSourceIBMPolicyTemplateV2PolicyResourceToMap(model.Resource, iamPolicyManagementClient)
	if err != nil {
		return modelMap, err
	}
	modelMap["resource"] = []map[string]interface{}{resourceMap}
	if model.Pattern != nil {
		modelMap["pattern"] = model.Pattern
	}

	if model.Rule != nil {
		modelMap["rule_conditions"] = flex.FlattenRuleConditions(*model.Rule.(*iampolicymanagementv1.V2PolicyRule))
		if len(model.Rule.(*iampolicymanagementv1.V2PolicyRule).Conditions) > 0 {
			modelMap["rule_operator"] = model.Rule.(*iampolicymanagementv1.V2PolicyRule).Operator
		}
	}
	controlResponse := model.Control
	policyRoles := flex.MapRolesToPolicyRoles(controlResponse.Grant.Roles)
	roles := flex.MapRoleListToPolicyRoles(*roleList)

	roleNames := []string{}
	for _, role := range policyRoles {
		role, err := flex.FindRoleByCRN(roles, *role.RoleID)
		if err != nil {
			return nil, err
		}
		roleNames = append(roleNames, *role.DisplayName)
	}
	modelMap["roles"] = roleNames
	return modelMap, nil
}

func dataSourceIBMPolicyTemplateV2PolicyResourceToMap(model *iampolicymanagementv1.V2PolicyResource, iamPolicyManagementClient *iampolicymanagementv1.IamPolicyManagementV1) (map[string]interface{}, *iampolicymanagementv1.RoleList, error) {
	modelMap := make(map[string]interface{})
	attributes := []map[string]interface{}{}
	listRoleOptions := &iampolicymanagementv1.ListRolesOptions{}
	var roles *iampolicymanagementv1.RoleList

	for _, attributesItem := range model.Attributes {
		if *attributesItem.Key == "serviceName" &&
			(*attributesItem.Operator == "stringMatch" ||
				*attributesItem.Operator == "stringEquals") {
			listRoleOptions.ServiceName = core.StringPtr(attributesItem.Value.(string))
		}

		if *attributesItem.Key == "service_group_id" && (*attributesItem.Operator == "stringMatch" ||
			*attributesItem.Operator == "stringEquals") {
			listRoleOptions.ServiceGroupID = core.StringPtr(attributesItem.Value.(string))
		}

		if *attributesItem.Key == "serviceType" && attributesItem.Value.(string) == "service" && (*attributesItem.Operator == "stringMatch" ||
			*attributesItem.Operator == "stringEquals") {
			listRoleOptions.ServiceName = core.StringPtr("alliamserviceroles")
		}

		roleList, _, err := iamPolicyManagementClient.ListRoles(listRoleOptions)
		roles = roleList
		if err != nil {
			return nil, nil, err
		}

		attributesItemMap, err := dataSourceIBMPolicyTemplateV2PolicyResourceAttributeToMap(&attributesItem)
		if err != nil {
			return modelMap, nil, err
		}
		attributes = append(attributes, attributesItemMap)
	}
	modelMap["attributes"] = attributes
	if model.Tags != nil {
		tags := []map[string]interface{}{}
		for _, tagsItem := range model.Tags {
			tagsItemMap, err := dataSourceIBMPolicyTemplateV2PolicyResourceTagToMap(&tagsItem)
			if err != nil {
				return modelMap, nil, err
			}
			tags = append(tags, tagsItemMap)
		}
		modelMap["tags"] = tags
	}
	return modelMap, roles, nil
}

func dataSourceIBMPolicyTemplateV2PolicyResourceAttributeToMap(model *iampolicymanagementv1.V2PolicyResourceAttribute) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["key"] = model.Key
	modelMap["operator"] = model.Operator
	modelMap["value"] = model.Value
	return modelMap, nil
}

func dataSourceIBMPolicyTemplateV2PolicyResourceTagToMap(model *iampolicymanagementv1.V2PolicyResourceTag) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["key"] = model.Key
	modelMap["value"] = model.Value
	modelMap["operator"] = model.Operator
	return modelMap, nil
}

func dataSourceIBMPolicyTemplateV2PolicyRuleToMap(model iampolicymanagementv1.V2PolicyRuleIntf) (map[string]interface{}, error) {
	if _, ok := model.(*iampolicymanagementv1.V2PolicyRuleRuleAttribute); ok {
		return dataSourceIBMPolicyTemplateV2PolicyRuleRuleAttributeToMap(model.(*iampolicymanagementv1.V2PolicyRuleRuleAttribute))
	} else if _, ok := model.(*iampolicymanagementv1.V2PolicyRuleRuleWithConditions); ok {
		return dataSourceIBMPolicyTemplateV2PolicyRuleRuleWithConditionsToMap(model.(*iampolicymanagementv1.V2PolicyRuleRuleWithConditions))
	} else if _, ok := model.(*iampolicymanagementv1.V2PolicyRule); ok {
		modelMap := make(map[string]interface{})
		model := model.(*iampolicymanagementv1.V2PolicyRule)
		if model.Key != nil {
			modelMap["key"] = model.Key
		}
		if model.Operator != nil {
			modelMap["operator"] = model.Operator
		}
		if model.Value != nil {
			modelMap["value"] = model.Value
		}
		if model.Conditions != nil {
			conditions := []map[string]interface{}{}
			for _, conditionsItem := range model.Conditions {
				conditionsItemMap, err := dataSourceIBMPolicyTemplateRuleAttributeToMap(&conditionsItem)
				if err != nil {
					return modelMap, err
				}
				conditions = append(conditions, conditionsItemMap)
			}
			modelMap["conditions"] = conditions
		}
		return modelMap, nil
	} else {
		return nil, fmt.Errorf("Unrecognized iampolicymanagementv1.V2PolicyRuleIntf subtype encountered")
	}
}

func dataSourceIBMPolicyTemplateRuleAttributeToMap(model *iampolicymanagementv1.RuleAttribute) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["key"] = model.Key
	modelMap["operator"] = model.Operator
	modelMap["value"] = model.Value
	return modelMap, nil
}

func dataSourceIBMPolicyTemplateV2PolicyRuleRuleAttributeToMap(model *iampolicymanagementv1.V2PolicyRuleRuleAttribute) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["key"] = model.Key
	modelMap["operator"] = model.Operator
	modelMap["value"] = model.Value
	return modelMap, nil
}

func dataSourceIBMPolicyTemplateV2PolicyRuleRuleWithConditionsToMap(model *iampolicymanagementv1.V2PolicyRuleRuleWithConditions) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["operator"] = model.Operator
	conditions := []map[string]interface{}{}
	for _, conditionsItem := range model.Conditions {
		conditionsItemMap, err := dataSourceIBMPolicyTemplateRuleAttributeToMap(&conditionsItem)
		if err != nil {
			return modelMap, err
		}
		conditions = append(conditions, conditionsItemMap)
	}
	modelMap["conditions"] = conditions
	return modelMap, nil
}

func dataSourceIBMPolicyTemplateControlToMap(model *iampolicymanagementv1.Control) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	grantMap, err := dataSourceIBMPolicyTemplateGrantToMap(model.Grant)
	if err != nil {
		return modelMap, err
	}
	modelMap["grant"] = []map[string]interface{}{grantMap}
	return modelMap, nil
}

func dataSourceIBMPolicyTemplateGrantToMap(model *iampolicymanagementv1.Grant) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	roles := []map[string]interface{}{}
	for _, rolesItem := range model.Roles {
		rolesItemMap, err := dataSourceIBMPolicyTemplateRolesToMap(&rolesItem)
		if err != nil {
			return modelMap, err
		}
		roles = append(roles, rolesItemMap)
	}
	modelMap["roles"] = roles
	return modelMap, nil
}

func dataSourceIBMPolicyTemplateRolesToMap(model *iampolicymanagementv1.Roles) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["role_id"] = model.RoleID
	return modelMap, nil
}
