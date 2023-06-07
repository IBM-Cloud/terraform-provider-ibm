// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package iampolicy

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/iampolicymanagementv1"
)

func ResourceIBMIAMPolicyTemplate() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMIAMPolicyTemplateCreate,
		ReadContext:   resourceIBMIAMPolicyTemplateVersionRead,
		UpdateContext: resourceIBMIAMPolicyTemplateUpdate,
		DeleteContext: resourceIBMIAMPolicyTemplateVersionDelete,
		Exists:        resourceIBMIAMPolicyTemplateVersionExists,
		Importer:      &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validate.InvokeValidator("ibm_iam_policy_template", "name"),
				Description:  "name of template.",
			},
			"account_id": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validate.InvokeValidator("ibm_iam_policy_template", "account_id"),
				Description:  "account id where this template will be created.",
			},
			"policy": {
				Type:        schema.TypeList,
				MinItems:    1,
				MaxItems:    1,
				Required:    true,
				Description: "The core set of properties associated with the template's policy objet.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"type": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The policy type; either 'access' or 'authorization'.",
						},
						"description": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Allows the customer to use their own words to record the purpose/context related to a policy.",
						},
						"resource": {
							Type:        schema.TypeList,
							MinItems:    1,
							MaxItems:    1,
							Required:    true,
							Description: "The resource attributes to which the policy grants access.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"attributes": {
										Type:        schema.TypeList,
										Required:    true,
										Description: "List of resource attributes to which the policy grants access.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"key": {
													Type:        schema.TypeString,
													Required:    true,
													Description: "The name of a resource attribute.",
												},
												"operator": {
													Type:        schema.TypeString,
													Required:    true,
													Description: "The operator of an attribute.",
												},
												"value": {
													Type:        schema.TypeString,
													Required:    true,
													Description: "The value of a rule or resource attribute; can be boolean or string for resource attribute. Can be string or an array of strings (e.g., array of days to permit access) for rule attribute.",
												},
											},
										},
									},
									"tags": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "Optional list of resource tags to which the policy grants access.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"key": {
													Type:        schema.TypeString,
													Required:    true,
													Description: "The name of an access management tag.",
												},
												"value": {
													Type:        schema.TypeString,
													Required:    true,
													Description: "The value of an access management tag.",
												},
												"operator": {
													Type:        schema.TypeString,
													Required:    true,
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
							Optional:    true,
							Description: "Indicates pattern of rule, either 'time-based-conditions:once', 'time-based-conditions:weekly:all-day', or 'time-based-conditions:weekly:custom-hours'.",
						},
						"rule": {
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "Additional access conditions associated with the policy.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"key": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The name of an attribute.",
									},
									"operator": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The operator of an attribute.",
									},
									"value": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The value of a rule or resource attribute; can be boolean or string for resource attribute. Can be string or an array of strings (e.g., array of days to permit access) for rule attribute.",
									},
									"conditions": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "List of conditions associated with a policy, e.g., time-based conditions that grant access over a certain time period.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"key": {
													Type:        schema.TypeString,
													Required:    true,
													Description: "The name of an attribute.",
												},
												"operator": {
													Type:        schema.TypeString,
													Required:    true,
													Description: "The operator of an attribute.",
												},
												"value": {
													Type:        schema.TypeString,
													Required:    true,
													Description: "The value of a rule or resource attribute; can be boolean or string for resource attribute. Can be string or an array of strings (e.g., array of days to permit access) for rule attribute.",
												},
											},
										},
									},
								},
							},
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
			"description": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validate.InvokeValidator("ibm_iam_policy_template", "description"),
				Description:  "description of template purpose.",
			},
			"committed": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "committed status for the template.",
			},
			"template_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Template ID.",
			},
			"version": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Template Version.",
			},
		},
	}
}

func ResourceIBMIAMPolicyTemplateValidator() *validate.ResourceValidator {
	validateSchema := make([]validate.ValidateSchema, 0)
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 "name",
			ValidateFunctionIdentifier: validate.StringLenBetween,
			Type:                       validate.TypeString,
			Required:                   true,
			MinValueLength:             1,
			MaxValueLength:             300,
		},
		validate.ValidateSchema{
			Identifier:                 "account_id",
			ValidateFunctionIdentifier: validate.StringLenBetween,
			Type:                       validate.TypeString,
			Required:                   true,
			MinValueLength:             1,
			MaxValueLength:             300,
		},
		validate.ValidateSchema{
			Identifier:                 "description",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Optional:                   true,
			Regexp:                     `^.*$`,
			MinValueLength:             0,
			MaxValueLength:             300,
		},
	)

	resourceValidator := validate.ResourceValidator{ResourceName: "ibm_iam_policy_template", Schema: validateSchema}
	return &resourceValidator
}

func resourceIBMIAMPolicyTemplateCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	iamPolicyManagementClient, err := meta.(conns.ClientSession).IAMPolicyManagementV1API()
	if err != nil {
		return diag.FromErr(err)
	}

	createPolicyTemplateOptions := &iampolicymanagementv1.CreatePolicyTemplateOptions{}

	createPolicyTemplateOptions.SetName(d.Get("name").(string))
	createPolicyTemplateOptions.SetAccountID(d.Get("account_id").(string))

	policyModel, err := resourceIBMPolicyTemplateMapToTemplatePolicy(d.Get("policy.0").(map[string]interface{}), iamPolicyManagementClient)
	if err != nil {
		return diag.FromErr(err)
	}
	createPolicyTemplateOptions.SetPolicy(policyModel)
	if _, ok := d.GetOk("description"); ok {
		createPolicyTemplateOptions.SetDescription(d.Get("description").(string))
	}
	if _, ok := d.GetOk("committed"); ok {
		createPolicyTemplateOptions.SetCommitted(d.Get("committed").(bool))
	}

	policyTemplate, response, err := iamPolicyManagementClient.CreatePolicyTemplateWithContext(context, createPolicyTemplateOptions)
	if err != nil {
		log.Printf("[DEBUG] CreatePolicyTemplateWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("CreatePolicyTemplateWithContext failed %s\n%s", err, response))
	}

	version, _ := strconv.Atoi(*policyTemplate.Version)
	d.SetId(fmt.Sprintf("%s/%d", *policyTemplate.ID, version))
	return resourceIBMIAMPolicyTemplateVersionRead(context, d, meta)
}

func resourceIBMIAMPolicyTemplateUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	_, err := meta.(conns.ClientSession).IAMPolicyManagementV1API()
	if err != nil {
		return diag.FromErr(err)
	}

	if d.HasChange("name") || d.HasChange("account_id") {
		return diag.FromErr(fmt.Errorf("Update failed. Reason: Name and accountId can't be updated"))
	}

	return resourceIBMIAMPolicyTemplateVersionUpdate(context, d, meta)
}

func resourceIBMPolicyTemplateMapToTemplatePolicy(modelMap map[string]interface{}, iamPolicyManagementClient *iampolicymanagementv1.IamPolicyManagementV1) (*iampolicymanagementv1.TemplatePolicy, error) {
	model := &iampolicymanagementv1.TemplatePolicy{}
	model.Type = core.StringPtr(modelMap["type"].(string))
	if modelMap["description"] != nil && modelMap["description"].(string) != "" {
		model.Description = core.StringPtr(modelMap["description"].(string))
	}
	ResourceModel, roleList, err := resourceIBMPolicyTemplateMapToV2PolicyResource(modelMap["resource"].([]interface{})[0].(map[string]interface{}), iamPolicyManagementClient)
	if err != nil {
		return model, err
	}
	model.Resource = ResourceModel
	if modelMap["pattern"] != nil && modelMap["pattern"].(string) != "" {
		model.Pattern = core.StringPtr(modelMap["pattern"].(string))
	}
	if modelMap["rule"] != nil && len(modelMap["rule"].([]interface{})) > 0 {
		RuleModel, err := resourceIBMPolicyTemplateMapToV2PolicyRule(modelMap["rule"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.Rule = RuleModel
	}

	controlModel, err := resourceIBMPolicyTemplateMapToControl(modelMap["roles"].([]interface{}), roleList)
	fmt.Println("=======controlModel========", controlModel)
	fmt.Println("=======roleList========", roleList)
	if err != nil {
		return nil, err
	}

	model.Control = controlModel
	return model, nil
}

func resourceIBMPolicyTemplateMapToV2PolicyResource(modelMap map[string]interface{},
	iamPolicyManagementClient *iampolicymanagementv1.IamPolicyManagementV1) (*iampolicymanagementv1.V2PolicyResource, *iampolicymanagementv1.RoleList, error) {
	model := &iampolicymanagementv1.V2PolicyResource{}
	attributes := []iampolicymanagementv1.V2PolicyResourceAttribute{}
	roleList := &iampolicymanagementv1.RoleList{}
	listRoleOptions := &iampolicymanagementv1.ListRolesOptions{}
	for _, attributesItem := range modelMap["attributes"].([]interface{}) {
		attributesItemModel, err := resourceIBMPolicyTemplateMapToV2PolicyResourceAttribute(attributesItem.(map[string]interface{}))

		if *attributesItemModel.Key == "serviceName" &&
			(*attributesItemModel.Operator == "stringMatch" ||
				*attributesItemModel.Operator == "stringEquals") {
			listRoleOptions.ServiceName = core.StringPtr(attributesItemModel.Value.(string))
		}

		if *attributesItemModel.Key == "service_group_id" && (*attributesItemModel.Operator == "stringMatch" ||
			*attributesItemModel.Operator == "stringEquals") {
			listRoleOptions.ServiceGroupID = core.StringPtr(attributesItemModel.Value.(string))
		}

		if *attributesItemModel.Key == "serviceType" && attributesItemModel.Value.(string) == "service" && (*attributesItemModel.Operator == "stringMatch" ||
			*attributesItemModel.Operator == "stringEquals") {
			listRoleOptions.ServiceName = core.StringPtr("alliamserviceroles")
		}

		roles, _, err := iamPolicyManagementClient.ListRoles(listRoleOptions)
		if err != nil {
			return model, nil, err
		}

		attributes = append(attributes, *attributesItemModel)
		roleList = roles
	}
	model.Attributes = attributes
	if modelMap["tags"] != nil {
		tags := []iampolicymanagementv1.V2PolicyResourceTag{}
		for _, tagsItem := range modelMap["tags"].([]interface{}) {
			tagsItemModel, err := resourceIBMPolicyTemplateMapToV2PolicyResourceTag(tagsItem.(map[string]interface{}))
			if err != nil {
				return model, nil, err
			}
			tags = append(tags, *tagsItemModel)
		}
		model.Tags = tags
	}
	return model, roleList, nil
}

func resourceIBMPolicyTemplateMapToV2PolicyResourceAttribute(modelMap map[string]interface{}) (*iampolicymanagementv1.V2PolicyResourceAttribute, error) {
	model := &iampolicymanagementv1.V2PolicyResourceAttribute{}
	model.Key = core.StringPtr(modelMap["key"].(string))
	model.Operator = core.StringPtr(modelMap["operator"].(string))
	model.Value = modelMap["value"].(string)
	return model, nil
}

func resourceIBMPolicyTemplateMapToV2PolicyResourceTag(modelMap map[string]interface{}) (*iampolicymanagementv1.V2PolicyResourceTag, error) {
	model := &iampolicymanagementv1.V2PolicyResourceTag{}
	model.Key = core.StringPtr(modelMap["key"].(string))
	model.Value = core.StringPtr(modelMap["value"].(string))
	model.Operator = core.StringPtr(modelMap["operator"].(string))
	return model, nil
}

func resourceIBMPolicyTemplateMapToV2PolicyRule(modelMap map[string]interface{}) (iampolicymanagementv1.V2PolicyRuleIntf, error) {
	model := &iampolicymanagementv1.V2PolicyRule{}
	if modelMap["key"] != nil && modelMap["key"].(string) != "" {
		model.Key = core.StringPtr(modelMap["key"].(string))
	}
	if modelMap["operator"] != nil && modelMap["operator"].(string) != "" {
		model.Operator = core.StringPtr(modelMap["operator"].(string))
	}
	if modelMap["value"] != nil {
		model.Value = modelMap["value"].(string)
	}
	if modelMap["conditions"] != nil {
		conditions := []iampolicymanagementv1.RuleAttribute{}
		for _, conditionsItem := range modelMap["conditions"].([]interface{}) {
			conditionsItemModel, err := resourceIBMPolicyTemplateMapToRuleAttribute(conditionsItem.(map[string]interface{}))
			if err != nil {
				return model, err
			}
			conditions = append(conditions, *conditionsItemModel)
		}
		model.Conditions = conditions
	}
	return model, nil
}

func resourceIBMPolicyTemplateMapToRuleAttribute(modelMap map[string]interface{}) (*iampolicymanagementv1.RuleAttribute, error) {
	model := &iampolicymanagementv1.RuleAttribute{}
	model.Key = core.StringPtr(modelMap["key"].(string))
	model.Operator = core.StringPtr(modelMap["operator"].(string))
	model.Value = modelMap["value"].(string)
	return model, nil
}

func resourceIBMPolicyTemplateMapToV2PolicyRuleRuleAttribute(modelMap map[string]interface{}) (*iampolicymanagementv1.V2PolicyRuleRuleAttribute, error) {
	model := &iampolicymanagementv1.V2PolicyRuleRuleAttribute{}
	model.Key = core.StringPtr(modelMap["key"].(string))
	model.Operator = core.StringPtr(modelMap["operator"].(string))
	model.Value = modelMap["value"].(string)
	return model, nil
}

func resourceIBMPolicyTemplateMapToV2PolicyRuleRuleWithConditions(modelMap map[string]interface{}) (*iampolicymanagementv1.V2PolicyRuleRuleWithConditions, error) {
	model := &iampolicymanagementv1.V2PolicyRuleRuleWithConditions{}
	model.Operator = core.StringPtr(modelMap["operator"].(string))
	conditions := []iampolicymanagementv1.RuleAttribute{}
	for _, conditionsItem := range modelMap["conditions"].([]interface{}) {
		conditionsItemModel, err := resourceIBMPolicyTemplateMapToRuleAttribute(conditionsItem.(map[string]interface{}))
		if err != nil {
			return model, err
		}
		conditions = append(conditions, *conditionsItemModel)
	}
	model.Conditions = conditions
	return model, nil
}

func resourceIBMPolicyTemplateMapToControl(roles []interface{}, roleList *iampolicymanagementv1.RoleList) (*iampolicymanagementv1.Control, error) {
	policyRoles := flex.MapRoleListToPolicyRoles(*roleList)

	policyRoles, err := flex.GetRolesFromRoleNames(flex.ExpandStringList(roles), policyRoles)
	if err != nil {
		return &iampolicymanagementv1.Control{}, err
	}
	policyGrant := &iampolicymanagementv1.Grant{
		Roles: flex.MapPolicyRolesToRoles(policyRoles),
	}
	policyControl := &iampolicymanagementv1.Control{
		Grant: policyGrant,
	}
	return policyControl, nil
}

func resourceIBMPolicyTemplateMapToGrant(modelMap map[string]interface{}, roleList *iampolicymanagementv1.RoleList) (*iampolicymanagementv1.Grant, error) {
	model := &iampolicymanagementv1.Grant{}
	roles := []iampolicymanagementv1.Roles{}

	for _, rolesItem := range modelMap["roles"].([]interface{}) {
		rolesItemModel, err := resourceIBMPolicyTemplateMapToRoles(rolesItem.(map[string]interface{}))
		if err != nil {
			return model, err
		}
		roles = append(roles, *rolesItemModel)
	}
	model.Roles = roles
	return model, nil
}

func resourceIBMPolicyTemplateMapToRoles(modelMap map[string]interface{}) (*iampolicymanagementv1.Roles, error) {
	model := &iampolicymanagementv1.Roles{}
	model.RoleID = core.StringPtr(modelMap["role_id"].(string))
	return model, nil
}

func resourceIBMPolicyTemplateTemplatePolicyToMap(model *iampolicymanagementv1.TemplatePolicy, iamPolicyManagementClient *iampolicymanagementv1.IamPolicyManagementV1) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["type"] = model.Type
	if model.Description != nil {
		modelMap["description"] = model.Description
	}
	resourceMap, roleList, err := resourceIBMPolicyTemplateV2PolicyResourceToMap(model.Resource, iamPolicyManagementClient)

	if err != nil {
		return nil, err
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
	modelMap["resource"] = []map[string]interface{}{resourceMap}
	if model.Pattern != nil {
		modelMap["pattern"] = model.Pattern
	}
	if model.Rule != nil {
		ruleMap, err := resourceIBMPolicyTemplateV2PolicyRuleToMap(model.Rule)
		if err != nil {
			return modelMap, err
		}
		modelMap["rule"] = []map[string]interface{}{ruleMap}
	}
	modelMap["roles"] = roleNames
	return modelMap, nil
}

func resourceIBMPolicyTemplateV2PolicyResourceToMap(model *iampolicymanagementv1.V2PolicyResource,
	iamPolicyManagementClient *iampolicymanagementv1.IamPolicyManagementV1) (map[string]interface{}, *iampolicymanagementv1.RoleList, error) {
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

		attributesItemMap, err := resourceIBMPolicyTemplateV2PolicyResourceAttributeToMap(&attributesItem)
		if err != nil {
			return modelMap, nil, err
		}
		attributes = append(attributes, attributesItemMap)
	}
	modelMap["attributes"] = attributes
	if model.Tags != nil {
		tags := []map[string]interface{}{}
		for _, tagsItem := range model.Tags {
			tagsItemMap, err := resourceIBMPolicyTemplateV2PolicyResourceTagToMap(&tagsItem)
			if err != nil {
				return modelMap, nil, err
			}
			tags = append(tags, tagsItemMap)
		}
		modelMap["tags"] = tags
	}
	return modelMap, roles, nil
}

func resourceIBMPolicyTemplateV2PolicyResourceAttributeToMap(model *iampolicymanagementv1.V2PolicyResourceAttribute) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["key"] = model.Key
	modelMap["operator"] = model.Operator
	modelMap["value"] = model.Value
	return modelMap, nil
}

func resourceIBMPolicyTemplateV2PolicyResourceTagToMap(model *iampolicymanagementv1.V2PolicyResourceTag) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["key"] = model.Key
	modelMap["value"] = model.Value
	modelMap["operator"] = model.Operator
	return modelMap, nil
}

func resourceIBMPolicyTemplateV2PolicyRuleToMap(model iampolicymanagementv1.V2PolicyRuleIntf) (map[string]interface{}, error) {
	if _, ok := model.(*iampolicymanagementv1.V2PolicyRuleRuleAttribute); ok {
		return resourceIBMPolicyTemplateV2PolicyRuleRuleAttributeToMap(model.(*iampolicymanagementv1.V2PolicyRuleRuleAttribute))
	} else if _, ok := model.(*iampolicymanagementv1.V2PolicyRuleRuleWithConditions); ok {
		return resourceIBMPolicyTemplateV2PolicyRuleRuleWithConditionsToMap(model.(*iampolicymanagementv1.V2PolicyRuleRuleWithConditions))
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
				conditionsItemMap, err := resourceIBMPolicyTemplateRuleAttributeToMap(&conditionsItem)
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

func resourceIBMPolicyTemplateRuleAttributeToMap(model *iampolicymanagementv1.RuleAttribute) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["key"] = model.Key
	modelMap["operator"] = model.Operator
	modelMap["value"] = model.Value
	return modelMap, nil
}

func resourceIBMPolicyTemplateV2PolicyRuleRuleAttributeToMap(model *iampolicymanagementv1.V2PolicyRuleRuleAttribute) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["key"] = model.Key
	modelMap["operator"] = model.Operator
	modelMap["value"] = model.Value
	return modelMap, nil
}

func resourceIBMPolicyTemplateV2PolicyRuleRuleWithConditionsToMap(model *iampolicymanagementv1.V2PolicyRuleRuleWithConditions) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["operator"] = model.Operator
	conditions := []map[string]interface{}{}
	for _, conditionsItem := range model.Conditions {
		conditionsItemMap, err := resourceIBMPolicyTemplateRuleAttributeToMap(&conditionsItem)
		if err != nil {
			return modelMap, err
		}
		conditions = append(conditions, conditionsItemMap)
	}
	modelMap["conditions"] = conditions
	return modelMap, nil
}

func resourceIBMPolicyTemplateGrantToMap(model *iampolicymanagementv1.Grant) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	roles := []map[string]interface{}{}
	for _, rolesItem := range model.Roles {
		rolesItemMap, err := resourceIBMPolicyTemplateRolesToMap(&rolesItem)
		if err != nil {
			return modelMap, err
		}
		roles = append(roles, rolesItemMap)
	}
	modelMap["roles"] = roles
	return modelMap, nil
}

func resourceIBMPolicyTemplateRolesToMap(model *iampolicymanagementv1.Roles) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["role_id"] = model.RoleID
	return modelMap, nil
}
