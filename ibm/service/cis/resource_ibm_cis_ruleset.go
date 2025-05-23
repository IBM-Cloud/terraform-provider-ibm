// Copyright IBM Corp. 2024, 2025. All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package cis

import (
	"fmt"
	"reflect"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/networking-go-sdk/rulesetsv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var CISResourceResponseObject = &schema.Resource{
	Schema: map[string]*schema.Schema{
		CISRulesetsDescription: {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Description of the rulesets",
		},
		CISRulesetsId: {
			Type:        schema.TypeString,
			Description: "Associated ruleset ID",
			Optional:    true,
		},
		CISRulesetsKind: {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Kind of the rulesets",
		},
		CISRulesetsLastUpdatedAt: {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Rulesets last updated at",
		},
		CISRulesetsName: {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Name of the rulesets",
		},
		CISRulesetsPhase: {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Phase of the rulesets",
		},
		CISRulesetsVersion: {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Version of the rulesets",
		},
		CISRulesetsRules: {
			Type:        schema.TypeList,
			Optional:    true,
			Description: "Rules of the rulesets",
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					CISRulesetsRuleId: {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "ID of the rulesets rule",
					},
					CISRulesetsRuleVersion: {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "Version of the rulesets rule",
					},
					CISRulesetsRuleAction: {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "Action of the rulesets rule",
					},
					CISRulesetsRuleActionParameters: {
						Type:        schema.TypeSet,
						Optional:    true,
						Description: "Action parameters of the rulesets rule",
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								CISRulesetsRuleId: {
									Type:        schema.TypeString,
									Optional:    true,
									Description: "ID of the rulesets rule",
								},
								CISRulesetOverrides: {
									Type:        schema.TypeSet,
									Optional:    true,
									Description: "Override options",
									Elem: &schema.Resource{
										Schema: map[string]*schema.Schema{
											CISRulesetOverridesAction: {
												Type:        schema.TypeString,
												Optional:    true,
												Description: "Action to perform",
											},
											CISRulesetOverridesEnabled: {
												Type:        schema.TypeBool,
												Optional:    true,
												Description: "Enable/Disable rule",
											},
											// CISRulesetOverridesSensitivityLevel: {
											// 	Type:        schema.TypeString,
											// 	Optional:    true,
											// 	Description: "Sensitivity level",
											// },
											CISRulesetOverridesRules: {
												Type:        schema.TypeList,
												Optional:    true,
												Description: "Rules",
												Elem: &schema.Resource{
													Schema: map[string]*schema.Schema{
														CISRulesetRuleId: {
															Type:        schema.TypeString,
															Optional:    true,
															Description: "ID of the ruleset",
														},
														CISRulesetOverridesEnabled: {
															Type:        schema.TypeBool,
															Optional:    true,
															Description: "Enable/Disable rule",
														},
														CISRulesetOverridesAction: {
															Type:        schema.TypeString,
															Optional:    true,
															Description: "Action to perform",
														},
														CISRulesetOverridesSensitivityLevel: {
															Type:        schema.TypeString,
															Optional:    true,
															Description: "Sensitivity level",
														},
														CISRulesetOverridesScoreThreshold: {
															Type:        schema.TypeInt,
															Optional:    true,
															Description: "Score Threshold",
														},
													},
												},
											},
											CISRulesetOverridesCategories: {
												Type:        schema.TypeList,
												Optional:    true,
												Description: "Categories",
												Elem: &schema.Resource{
													Schema: map[string]*schema.Schema{
														CISRulesetOverridesCategoriesCategory: {
															Type:        schema.TypeString,
															Optional:    true,
															Description: "Category",
														},
														CISRulesetOverridesEnabled: {
															Type:        schema.TypeBool,
															Optional:    true,
															Description: "Enable/Disable rule",
														},
														CISRulesetOverridesAction: {
															Type:        schema.TypeString,
															Optional:    true,
															Description: "Action to perform",
														},
													},
												},
											},
										},
									},
								},
								CISRulesetsVersion: {
									Type:        schema.TypeString,
									Optional:    true,
									Description: "Version of the ruleset",
								},
								CISRuleset: {
									Type:        schema.TypeString,
									Optional:    true,
									Description: "Ruleset of the rule",
								},
								CISRulesetsRulePhases: {
									Type:        schema.TypeList,
									Optional:    true,
									Description: "Phases of the rule",
									Elem:        &schema.Schema{Type: schema.TypeString},
								},
								CISRulesetsRuleProducts: {
									Type:        schema.TypeList,
									Optional:    true,
									Description: "Products of the rule",
									Elem:        &schema.Schema{Type: schema.TypeString},
								},
								CISRulesetList: {
									Type:        schema.TypeList,
									Optional:    true,
									Description: "List of ruleset IDs of the ruleset to apply action to",
									Elem:        &schema.Schema{Type: schema.TypeString},
								},
								CISRulesetsRuleActionParametersResponse: {
									Type:        schema.TypeSet,
									Optional:    true,
									Description: "Action parameters response of the rulesets rule",
									Elem: &schema.Resource{
										Schema: map[string]*schema.Schema{
											CISRulesetsRuleActionParametersResponseContent: {
												Type:        schema.TypeString,
												Optional:    true,
												Description: "Action parameters response content of the rulesets rule",
											},
											CISRulesetsRuleActionParametersResponseContentType: {
												Type:        schema.TypeString,
												Optional:    true,
												Description: "Action parameters response type of the rulesets rule",
											},
											CISRulesetsRuleActionParametersResponseStatusCode: {
												Type:        schema.TypeInt,
												Optional:    true,
												Description: "Action parameters response status code of the rulesets rule",
											},
										},
									},
								},
							},
						},
					},
					CISRulesetsRuleActionCategories: {
						Type:        schema.TypeList,
						Optional:    true,
						Description: "Categories of the rulesets rule",
						Elem:        &schema.Schema{Type: schema.TypeString},
					},
					CISRulesetsRuleActionEnabled: {
						Type:        schema.TypeBool,
						Optional:    true,
						Description: "Enable/Disable ruleset rule",
					},
					CISRulesetsRuleActionDescription: {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "Description of the rulesets rule",
					},
					CISRulesetsRuleExpression: {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "Expression of the rulesets rule",
					},
					CISRulesetsRuleRef: {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "Reference of the rulesets rule",
					},
					CISRulesetsRuleLogging: {
						Type:        schema.TypeMap,
						Optional:    true,
						Description: "Logging of the rulesets rule",
						Elem:        &schema.Schema{Type: schema.TypeBool},
					},
					CISRulesetsRulePosition: {
						Type:        schema.TypeSet,
						Optional:    true,
						Description: "Position of rulesets rule",
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								CISRulesetsRulePositionAfter: {
									Type:        schema.TypeString,
									Optional:    true,
									Description: "Ruleset before position of rulesets rule",
								},
								CISRulesetsRulePositionBefore: {
									Type:        schema.TypeString,
									Optional:    true,
									Description: "Ruleset after position of rulesets rule",
								},
								CISRulesetsRulePositionIndex: {
									Type:        schema.TypeInt,
									Optional:    true,
									Description: "Index of rulesets rule",
								},
							},
						},
					},
					CISRulesetsRuleLastUpdatedAt: {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "Rulesets rule last updated at",
					},
				},
			},
		},
	},
}

func ResourceIBMCISRuleset() *schema.Resource {
	return &schema.Resource{
		Read:     ResourceIBMCISRulesetRead,
		Update:   ResourceIBMCISRulesetUpdate,
		Delete:   ResourceIBMCISRulesetDelete,
		Create:   ResourceIBMCISRulesetRead,
		Importer: &schema.ResourceImporter{},
		Schema: map[string]*schema.Schema{
			cisID: {
				Type:        schema.TypeString,
				Description: "CIS instance CRN",
				Required:    true,
				ValidateFunc: validate.InvokeValidator("ibm_cis_ruleset",
					"cis_id"),
			},
			cisDomainID: {
				Type:             schema.TypeString,
				Description:      "Associated CIS domain",
				Optional:         true,
				DiffSuppressFunc: suppressDomainIDDiff,
			},
			CISRulesetsId: {
				Type:        schema.TypeString,
				Description: "Associated ruleset ID",
				Optional:    true,
			},
			CISRulesetsObjectOutput: {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Container for response information",
				Elem:        CISResourceResponseObject,
			},
		},
	}
}
func ResourceIBMCISRulesetValidator() *validate.ResourceValidator {
	validateSchema := make([]validate.ValidateSchema, 0)
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 "cis_id",
			ValidateFunctionIdentifier: validate.ValidateCloudData,
			Type:                       validate.TypeString,
			CloudDataType:              "resource_instance",
			CloudDataRange:             []string{"service:internet-svcs"},
			Required:                   true})
	ibmCISRulesetValidator := validate.ResourceValidator{
		ResourceName: "ibm_cis_ruleset",
		Schema:       validateSchema}
	return &ibmCISRulesetValidator
}

func ResourceIBMCISRulesetCreate(d *schema.ResourceData, meta interface{}) error {
	// check if it is a new resource, if true then return error that user need to import it first
	if d.IsNewResource() {
		return fmt.Errorf("[ERROR] You can not create a new resource. Please import the resource first. Check documentation for import usage")
	}
	return nil
}

func ResourceIBMCISRulesetUpdate(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(conns.ClientSession).CisRulesetsSession()
	if err != nil {
		return fmt.Errorf("[ERROR] Error while getting the CisRulesetsSession %s", err)
	}

	crn := d.Get(cisID).(string)
	zoneId := d.Get(cisDomainID).(string)
	rulesetId := d.Get(CISRulesetsId).(string)
	sess.Crn = core.StringPtr(crn)

	if zoneId != "" {
		sess.ZoneIdentifier = core.StringPtr(zoneId)

		opt := sess.NewUpdateZoneRulesetOptions(rulesetId)

		rulesetsObject := d.Get(CISRulesetsObjectOutput).([]interface{})[0].(map[string]interface{})
		opt.SetDescription(rulesetsObject[CISRulesetsDescription].(string))
		opt.SetName(rulesetsObject[CISRulesetsName].(string))
		opt.SetRulesetID(rulesetId)

		rulesObj := expandCISRules(rulesetsObject[CISRulesetsRules])
		opt.SetRules(rulesObj)

		_, resp, err := sess.UpdateZoneRuleset(opt)

		if err != nil {
			return fmt.Errorf("[ERROR] Error while updating the zone Ruleset %s", resp)
		}

		d.SetId(dataSourceCISRulesetsCheckID(d))

	} else {
		opt := sess.NewUpdateInstanceRulesetOptions(rulesetId)

		rulesetsObject := d.Get(CISRulesetsObjectOutput).([]interface{})[0].(map[string]interface{})
		opt.SetDescription(rulesetsObject[CISRulesetsDescription].(string))
		opt.SetName(rulesetsObject[CISRulesetsName].(string))
		opt.SetRulesetID(rulesetId)

		rulesObj := expandCISRules(rulesetsObject[CISRulesetsRules])
		opt.SetRules(rulesObj)

		_, _, err := sess.UpdateInstanceRuleset(opt)

		if err != nil {
			return fmt.Errorf("[ERROR] Error while updating the zone Ruleset %s", err)
		}

		d.SetId(dataSourceCISRulesetsCheckID(d))
	}

	return ResourceIBMCISRulesetRead(d, meta)
}

func ResourceIBMCISRulesetRead(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(conns.ClientSession).CisRulesetsSession()
	if err != nil {
		return fmt.Errorf("[ERROR] Error while getting the CisRulesetsSession %s", err)
	}

	rulesetId, zoneId, crn, _ := flex.ConvertTfToCisThreeVar(d.Id())
	d.Set(CISRulesetsId, rulesetId)
	d.Set(cisDomainID, zoneId)
	d.Set(cisID, crn)

	sess.Crn = core.StringPtr(crn)

	if zoneId != "" {
		sess.ZoneIdentifier = core.StringPtr(zoneId)
		opt := sess.NewGetZoneRulesetOptions(rulesetId)
		result, resp, err := sess.GetZoneRuleset(opt)
		if err != nil {
			return fmt.Errorf("[WARN] Resource: Get zone ruleset failed:  %v \n", resp)
		}

		rulesetObj := flattenCISRulesets(*result.Result)

		d.Set(CISRulesetsObjectOutput, rulesetObj)
		d.Set(cisDomainID, zoneId)
		d.Set(cisID, crn)
		d.SetId(dataSourceCISRulesetsCheckID(d))

	} else {
		opt := sess.NewGetInstanceRulesetOptions(rulesetId)
		result, resp, err := sess.GetInstanceRuleset(opt)
		if err != nil {
			return fmt.Errorf("[WARN] Resource: Get Instance ruleset failed: %v \n", resp)
		}

		rulesetObj := flattenCISRulesets(*result.Result)

		d.Set(CISRulesetsListOutput, rulesetObj)
		d.Set(CISRulesetsId, rulesetId)
		d.Set(cisID, crn)
		d.SetId(dataSourceCISRulesetsCheckID(d))
	}

	return nil
}

func ResourceIBMCISRulesetDelete(d *schema.ResourceData, meta interface{}) error {

	sess, err := meta.(conns.ClientSession).CisRulesetsSession()
	if err != nil {
		return fmt.Errorf("[ERROR] Error while getting the CisRulesetsSession %s", err)
	}

	rulesetId, zoneId, crn, _ := flex.ConvertTfToCisThreeVar(d.Id())
	sess.Crn = core.StringPtr(crn)

	if zoneId != "" {
		sess.ZoneIdentifier = core.StringPtr(zoneId)
		opt := sess.NewDeleteZoneRulesetOptions(rulesetId)
		res, err := sess.DeleteZoneRuleset(opt)
		if err != nil {
			return fmt.Errorf("[ERROR] Error deleting the zone ruleset %s:%s", err, res)
		}
	} else {
		opt := sess.NewDeleteInstanceRulesetOptions(rulesetId)
		res, err := sess.DeleteInstanceRuleset(opt)
		if err != nil {
			return fmt.Errorf("[ERROR] Error deleting the Instance ruleset %s:%s", err, res)
		}
	}

	d.SetId("")
	return nil
}

func expandCISRules(obj interface{}) []rulesetsv1.RuleCreate {

	finalResponse := make([]rulesetsv1.RuleCreate, 0)

	listResponse := obj.([]interface{})
	for _, val := range listResponse {

		ruleObj := val.(map[string]interface{})

		id := ruleObj[CISRulesetsRuleId].(string)
		expression := ruleObj[CISRulesetsRuleExpression].(string)
		action := ruleObj[CISRulesetsRuleAction].(string)
		description := ruleObj[CISRulesetsRuleActionDescription].(string)
		enabled := ruleObj[CISRulesetsRuleActionEnabled].(bool)
		ref := ruleObj[CISRulesetsRuleRef].(string)

		actionParameterObj := rulesetsv1.ActionParameters{}

		if len(ruleObj[CISRulesetsRuleActionParameters].(*schema.Set).List()) != 0 {
			actionParameterObj = expandCISRulesetsRulesActionParameters(ruleObj[CISRulesetsRuleActionParameters])
		}

		position := rulesetsv1.Position{}
		if len(ruleObj[CISRulesetsRulePosition].(*schema.Set).List()) != 0 {
			var err error
			position, err = expandCISRulesetsRulesPositions(ruleObj[CISRulesetsRulePosition])
			if err != nil {
				fmt.Printf("[ERROR] Error while expanding the CIS Rulesets Rule Position %s", err)
				return []rulesetsv1.RuleCreate{}
			}
		}

		ruleRespObj := rulesetsv1.RuleCreate{
			Expression:       &expression,
			Action:           &action,
			Description:      &description,
			Enabled:          &enabled,
			Ref:              &ref,
			ActionParameters: &actionParameterObj,
			Position:         &position,
		}

		if id != "" {
			ruleRespObj.ID = &id
		}

		finalResponse = append(finalResponse, ruleRespObj)
	}

	return finalResponse
}

func expandCISRulesetsRulesPositions(obj interface{}) (rulesetsv1.Position, error) {
	responseObj := rulesetsv1.Position{}
	if len(obj.(*schema.Set).List()) != 0 {
		response := obj.(*schema.Set).List()[0].(map[string]interface{})

		before := response[CISRulesetsRulePositionBefore].(string)
		after := response[CISRulesetsRulePositionAfter].(string)
		index := int64(response[CISRulesetsRulePositionIndex].(int))

		if before != "" && after == "" && index == 0 {
			responseObj = rulesetsv1.Position{
				Before: &before,
			}
		} else if after != "" && before == "" && index == 0 {
			responseObj = rulesetsv1.Position{
				After: &after,
			}
		} else if index != 0 && before == "" && after == "" {
			responseObj = rulesetsv1.Position{
				Index: &index,
			}
		} else {
			return rulesetsv1.Position{}, fmt.Errorf("only one of 'before', 'after', or 'index' can be set")
		}
	}
	return responseObj, nil
}

func expandCISRulesetsRulesActionParameters(obj interface{}) rulesetsv1.ActionParameters {

	actionParameterRespObj := rulesetsv1.ActionParameters{}
	// return empty object if action parameter is not provided.
	if len(obj.(*schema.Set).List()) == 0 {
		return actionParameterRespObj
	}

	actionParameterObj := obj.(*schema.Set).List()[0].(map[string]interface{})

	id := actionParameterObj[CISRulesetsRuleId].(string)
	if id != "" {
		actionParameterRespObj.ID = &id
	}
	version := actionParameterObj[CISRulesetsVersion].(string)
	if version != "" {
		actionParameterRespObj.Version = &version
	}
	ruleListInterface := actionParameterObj[CISRulesetList].([]interface{})

	ruleList := make([]string, 0)
	for i, v := range ruleListInterface {
		ruleList[i] = fmt.Sprint(v)
	}
	actionParameterRespObj.Rulesets = ruleList

	ruleset := actionParameterObj[CISRuleset].(string)
	if ruleset != "" {
		actionParameterRespObj.Ruleset = &ruleset
	}

	phases := actionParameterObj[CISRulesetsRulePhases].([]interface{})
	phasesList := flex.ExpandStringList(phases)
	actionParameterRespObj.Phases = phasesList

	products := actionParameterObj[CISRulesetsRuleProducts].([]interface{})
	productsList := flex.ExpandStringList(products)
	actionParameterRespObj.Products = productsList

	finalResponse := make([]rulesetsv1.ActionParameters, 0)

	overrideObj := rulesetsv1.Overrides{}
	if len(actionParameterObj[CISRulesetOverrides].(*schema.Set).List()) != 0 {
		overrideObj = expandCISRulesetsRulesActionParametersOverrides(actionParameterObj[CISRulesetOverrides])
		actionParameterRespObj.Overrides = &overrideObj
	}

	resObj := rulesetsv1.ActionParametersResponse{}
	if len(actionParameterObj[CISRulesetsRuleActionParametersResponse].(*schema.Set).List()) != 0 {
		resObj = expandCISRulesetsRulesActionParametersResponse(actionParameterObj[CISRulesetsRuleActionParametersResponse])
		actionParameterRespObj.Response = &resObj
	}

	finalResponse = append(finalResponse, actionParameterRespObj)

	return finalResponse[0]
}

func expandCISRulesetsRulesActionParametersResponse(obj interface{}) rulesetsv1.ActionParametersResponse {
	response := obj.(*schema.Set).List()[0].(map[string]interface{})
	content := response[CISRulesetsRuleActionParametersResponseContent].(string)
	contentType := response[CISRulesetsRuleActionParametersResponseContentType].(string)
	statusCode := int64(response[cisPageRuleActionsValueStatusCode].(int))

	responseObj := rulesetsv1.ActionParametersResponse{
		Content:     &content,
		ContentType: &contentType,
		StatusCode:  &statusCode,
	}

	return responseObj
}

func expandCISRulesetsRulesActionParametersOverrides(obj interface{}) rulesetsv1.Overrides {

	overrideObj := obj.(*schema.Set).List()[0].(map[string]interface{})
	actionOverride := overrideObj[CISRulesetOverridesAction].(string)
	enabledOverride := overrideObj[CISRulesetOverridesEnabled].(bool)

	rules := []rulesetsv1.RulesOverride{}
	categories := []rulesetsv1.CategoriesOverride{}
	if !reflect.ValueOf(overrideObj[CISRulesetOverridesRules]).IsNil() {
		rules = expandCISRulesetsRulesActionParametersOverridesRules(overrideObj[CISRulesetOverridesRules])
	}
	if !reflect.ValueOf(overrideObj[CISRulesetOverridesCategories]).IsNil() {
		categories = expandCISRulesetsRulesActionParametersOverridesCategories(overrideObj[CISRulesetOverridesCategories])
	}

	finalResponse := make([]rulesetsv1.Overrides, 0)
	overrideRespObj := rulesetsv1.Overrides{
		Rules:      rules,
		Categories: categories,
	}
	if actionOverride != "" {
		overrideRespObj.Action = &actionOverride
	}
	if enabledOverride {
		overrideRespObj.Enabled = &enabledOverride
	}
	finalResponse = append(finalResponse, overrideRespObj)

	return finalResponse[0]
}

func expandCISRulesetsRulesActionParametersOverridesCategories(obj interface{}) []rulesetsv1.CategoriesOverride {
	finalResponse := make([]rulesetsv1.CategoriesOverride, 0)

	listResponse := obj.([]interface{})

	for _, val := range listResponse {
		response := val.(map[string]interface{})
		action := response[CISRulesetOverridesAction].(string)
		enabled := response[CISRulesetOverridesEnabled].(bool)
		category := response[CISRulesetOverridesCategoriesCategory].(string)
		overrideRespObj := rulesetsv1.CategoriesOverride{
			Category: &category,
			Enabled:  &enabled,
		}
		if action != "" {
			overrideRespObj.Action = &action
		}
		finalResponse = append(finalResponse, overrideRespObj)

	}

	return finalResponse
}

func expandCISRulesetsRulesActionParametersOverridesRules(obj interface{}) []rulesetsv1.RulesOverride {
	finalResponse := make([]rulesetsv1.RulesOverride, 0)

	listResponse := obj.([]interface{})
	for _, val := range listResponse {
		response := val.(map[string]interface{})
		id := response[CISRulesetRuleId].(string)
		action := response[CISRulesetOverridesAction].(string)
		enabled := response[CISRulesetOverridesEnabled].(bool)
		score := int64(response[CISRulesetOverridesScoreThreshold].(int))

		overrideRespObj := rulesetsv1.RulesOverride{
			ID:      &id,
			Enabled: &enabled,
		}
		if action != "" {
			overrideRespObj.Action = &action
		}
		if score != 0 {
			overrideRespObj.ScoreThreshold = &score
		}
		finalResponse = append(finalResponse, overrideRespObj)
	}

	return finalResponse
}
