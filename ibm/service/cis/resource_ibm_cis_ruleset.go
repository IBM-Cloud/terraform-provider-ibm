// Copyright IBM Corp. 2024 All Rights Reserved.
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
			Description: "Description of the Rulesets",
		},
		CISRulesetsId: {
			Type:        schema.TypeString,
			Description: "Associated Ruleset ID",
			Optional:    true,
		},
		CISRulesetsKind: {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Kind of the Rulesets",
		},
		CISRulesetsLastUpdatedAt: {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Rulesets Last Updated At",
		},
		CISRulesetsName: {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Name of the Rulesets",
		},
		CISRulesetsPhase: {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Phase of the Rulesets",
		},
		CISRulesetsVersion: {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Version of the Rulesets",
		},
		CISRulesetsRules: {
			Type:        schema.TypeList,
			Optional:    true,
			Description: "Rules of the Rulesets",
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					CISRulesetsRuleId: {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "Id of the Rulesets Rule",
					},
					CISRulesetsRuleVersion: {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "Version of the Rulesets Rule",
					},
					CISRulesetsRuleAction: {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "Action of the Rulesets Rule",
					},
					CISRulesetsRuleActionParameters: {
						Type:        schema.TypeSet,
						Optional:    true,
						Description: "Action parameters of the Rulesets Rule",
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								CISRulesetsRuleId: {
									Type:        schema.TypeString,
									Optional:    true,
									Description: "Id of the Rulesets Rule",
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
												Description: "Enable Disable Rule",
											},
											CISRulesetOverridesSensitivityLevel: {
												Type:        schema.TypeString,
												Optional:    true,
												Description: "Sensitivity Level",
											},
											CISRulesetOverridesRules: {
												Type:        schema.TypeList,
												Optional:    true,
												Description: "Rules",
												Elem: &schema.Resource{
													Schema: map[string]*schema.Schema{
														CISRulesetsId: {
															Type:        schema.TypeString,
															Optional:    true,
															Description: "Id of the Ruleset",
														},
														CISRulesetOverridesEnabled: {
															Type:        schema.TypeBool,
															Optional:    true,
															Description: "Enable Disable Rule",
														},
														CISRulesetOverridesAction: {
															Type:        schema.TypeString,
															Optional:    true,
															Description: "Action to perform",
														},
														CISRulesetOverridesSensitivityLevel: {
															Type:        schema.TypeString,
															Optional:    true,
															Description: "Sensitivity Level",
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
															Description: "Enable Disable Rule",
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
									Description: "Version of the Ruleset",
								},
								CISRuleset: {
									Type:        schema.TypeString,
									Optional:    true,
									Description: "Ruleset ID of the ruleset to apply action to.",
								},
								CISRulesetList: {
									Type:        schema.TypeList,
									Optional:    true,
									Description: "List of Ruleset IDs of the ruleset to apply action to.",
									Elem:        &schema.Schema{Type: schema.TypeString},
								},
								CISRulesetsRuleActionParametersResponse: {
									Type:        schema.TypeSet,
									Optional:    true,
									Description: "Action parameters response of the Rulesets Rule",
									Elem: &schema.Resource{
										Schema: map[string]*schema.Schema{
											CISRulesetsRuleActionParametersResponseContent: {
												Type:        schema.TypeString,
												Optional:    true,
												Description: "Action parameters response content of the Rulesets Rule",
											},
											CISRulesetsRuleActionParametersResponseContentType: {
												Type:        schema.TypeString,
												Optional:    true,
												Description: "Action parameters response type of the Rulesets Rule",
											},
											CISRulesetsRuleActionParametersResponseStatusCode: {
												Type:        schema.TypeInt,
												Optional:    true,
												Description: "Action parameters response status code of the Rulesets Rule",
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
						Description: "Categories of the Rulesets Rule",
						Elem:        &schema.Schema{Type: schema.TypeString},
					},
					CISRulesetsRuleActionEnabled: {
						Type:        schema.TypeBool,
						Optional:    true,
						Description: "Enable/Disable Ruleset Rule",
					},
					CISRulesetsRuleActionDescription: {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "Description of the Rulesets Rule",
					},
					CISRulesetsRuleExpression: {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "Experession of the Rulesets Rule",
					},
					CISRulesetsRuleRef: {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "Reference of the Rulesets Rule",
					},
					CISRulesetsRuleLogging: {
						Type:        schema.TypeMap,
						Optional:    true,
						Description: "Logging of the Rulesets Rule",
						Elem:        &schema.Schema{Type: schema.TypeBool},
					},
					CISRulesetsRulePosition: {
						Type:        schema.TypeSet,
						Optional:    true,
						Description: "Position of Rulesets Rule",
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								CISRulesetsRulePositionAfter: {
									Type:        schema.TypeString,
									Optional:    true,
									Description: "Ruleset before Position of Rulesets Rule",
								},
								CISRulesetsRulePositionBefore: {
									Type:        schema.TypeString,
									Optional:    true,
									Description: "Ruleset after Position of Rulesets Rule",
								},
								CISRulesetsRulePositionIndex: {
									Type:        schema.TypeInt,
									Optional:    true,
									Description: "Index of Rulesets Rule",
								},
							},
						},
					},
					CISRulesetsRuleLastUpdatedAt: {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "Rulesets Rule Last Updated At",
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
				Description: "CIS instance crn",
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
				Description: "Associated Ruleset ID",
				Optional:    true,
			},
			CISRulesetsObjectOutput: {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Container for response information.",
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
		return fmt.Errorf("[ERROR] You can not create a new resource. Please import the resource first. Check documentation for import usage.")
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
		opt.SetKind(rulesetsObject[CISRulesetsKind].(string))
		opt.SetName(rulesetsObject[CISRulesetsName].(string))
		opt.SetPhase(rulesetsObject[CISRulesetsPhase].(string))
		opt.SetRulesetID(rulesetId)

		rulesObj := expandCISRules(rulesetsObject[CISRulesetsRules])
		opt.SetRules(rulesObj)

		result, _, err := sess.UpdateZoneRuleset(opt)

		if err != nil {
			return fmt.Errorf("[ERROR] Error while updating the zone Ruleset %s", err)
		}

		d.SetId(*result.Result.ID)

	} else {
		opt := sess.NewUpdateInstanceRulesetOptions(rulesetId)

		rulesetsObject := d.Get(CISRulesetsObjectOutput).([]interface{})[0].(map[string]interface{})
		opt.SetDescription(rulesetsObject[CISRulesetsDescription].(string))
		opt.SetKind(rulesetsObject[CISRulesetsKind].(string))
		opt.SetName(rulesetsObject[CISRulesetsName].(string))
		opt.SetPhase(rulesetsObject[CISRulesetsPhase].(string))
		opt.SetRulesetID(rulesetId)

		rulesObj := expandCISRules(rulesetsObject[CISRulesetsRules])

		opt.SetRules(rulesObj)

		result, _, err := sess.UpdateInstanceRuleset(opt)

		if err != nil {
			return fmt.Errorf("[ERROR] Error while updating the zone Ruleset %s", err)
		}

		d.SetId(*result.Result.ID)
	}

	return ResourceIBMCISRulesetRead(d, meta)
}

func ResourceIBMCISRulesetRead(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(conns.ClientSession).CisRulesetsSession()
	if err != nil {
		return fmt.Errorf("[ERROR] Error while getting the CisRulesetsSession %s", err)
	}

	rulesetId := d.Get(CISRulesetsId).(string)
	zoneId := d.Get(cisDomainID).(string)
	crn := d.Get(cisID).(string)

	// if reading from state file after importing
	if rulesetId == "" {
		rulesetId, zoneId, crn, _ = flex.ConvertTfToCisThreeVar(d.Id())
	}

	sess.Crn = core.StringPtr(crn)

	if zoneId != "" {
		sess.ZoneIdentifier = core.StringPtr(zoneId)
		opt := sess.NewGetZoneRulesetOptions(rulesetId)
		result, resp, err := sess.GetZoneRuleset(opt)
		if err != nil {
			return fmt.Errorf("[WARN] Resource: Get zone ruleset failed:  %v \n", resp)
		}

		rulesetObj := flattenCISRulesets(*result.Result)

		d.SetId(dataSourceCISRulesetsCheckID(d))
		d.Set(CISRulesetsObjectOutput, rulesetObj)
		d.Set(cisDomainID, zoneId)
		d.Set(CISRulesetsId, rulesetId)
		d.Set(cisID, crn)

	} else {
		opt := sess.NewGetInstanceRulesetOptions(rulesetId)
		result, resp, err := sess.GetInstanceRuleset(opt)
		if err != nil {
			return fmt.Errorf("[WARN] Resource: Get Instance ruleset failed: %v \n", resp)
		}

		rulesetObj := flattenCISRulesets(*result.Result)

		d.SetId(dataSourceCISRulesetsCheckID(d))
		d.Set(CISRulesetsListOutput, rulesetObj)
		d.Set(CISRulesetsId, rulesetId)
		d.Set(cisID, crn)
	}

	return nil
}

func ResourceIBMCISRulesetDelete(d *schema.ResourceData, meta interface{}) error {

	sess, err := meta.(conns.ClientSession).CisRulesetsSession()
	if err != nil {
		return fmt.Errorf("[ERROR] Error while getting the CisRulesetsSession %s", err)
	}

	crn := d.Get(cisID).(string)
	sess.Crn = core.StringPtr(crn)

	zoneId := d.Get(cisDomainID).(string)
	rulesetId := d.Get(CISRulesetsId).(string)

	if zoneId != "" {
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

	ruleObj := obj.([]interface{})[0].(map[string]interface{})
	id := ruleObj[CISRulesetsRuleId].(string)
	expression := ruleObj[CISRulesetsRuleExpression].(string)
	action := ruleObj[CISRulesetsRuleAction].(string)
	description := ruleObj[CISRulesetsRuleActionDescription].(string)
	enabled := ruleObj[CISRulesetsRuleActionEnabled].(bool)
	ref := ruleObj[CISRulesetsRuleRef].(string)
	actionParameterObj := rulesetsv1.ActionParameters{}
	logging := rulesetsv1.Logging{}
	position := rulesetsv1.Position{}

	if reflect.ValueOf(ruleObj[CISRulesetsRuleActionParameters]).IsNil() {
		actionParameterObj = expandCISRulesetsRulesActionParameters(ruleObj[CISRulesetsRuleActionParameters])
	}
	if reflect.ValueOf(ruleObj[CISRulesetsRuleLogging]).IsNil() {
		logging = expandCISRulesetsRulesLogging(ruleObj[CISRulesetsRuleLogging])
	}
	if reflect.ValueOf(ruleObj[CISRulesetsRulePosition]).IsNil() {
		position = expandCISRulesetsRulesPositions(ruleObj[CISRulesetsRulePosition])
	}

	finalResponse := make([]rulesetsv1.RuleCreate, 0)
	ruleRespObj := rulesetsv1.RuleCreate{
		ID:               &id,
		Expression:       &expression,
		Action:           &action,
		Description:      &description,
		Enabled:          &enabled,
		Ref:              &ref,
		ActionParameters: &actionParameterObj,
		Logging:          &logging,
		Position:         &position,
	}
	finalResponse = append(finalResponse, ruleRespObj)

	return finalResponse
}

func expandCISRulesetsRulesLogging(obj interface{}) rulesetsv1.Logging {
	response := obj.(map[string]interface{})
	enabled := response[CISRulesetsRuleActionEnabled].(bool)
	responseObj := rulesetsv1.Logging{
		Enabled: &enabled,
	}
	return responseObj
}

func expandCISRulesetsRulesPositions(obj interface{}) rulesetsv1.Position {
	response := obj.(map[string]interface{})
	before := response[CISRulesetsRulePositionBefore].(string)
	after := response[CISRulesetsRulePositionAfter].(string)
	index := response[CISRulesetsRulePositionIndex].(int64)
	responseObj := rulesetsv1.Position{
		Before: &before,
		After:  &after,
		Index:  &index,
	}
	return responseObj
}

func expandCISRulesetsRulesActionParameters(obj interface{}) rulesetsv1.ActionParameters {

	actionParameterObj := obj.(*schema.Set).List()[0].(map[string]interface{})
	id := actionParameterObj[CISRulesetsRuleId].(string)
	ruleset := actionParameterObj[CISRuleset].(string)
	version := actionParameterObj[CISRulesetsVersion].(string)
	ruleListInterface := actionParameterObj[CISRulesetList].([]interface{})

	ruleList := make([]string, 0)
	for i, v := range ruleListInterface {
		ruleList[i] = fmt.Sprint(v)
	}

	overrideObj := rulesetsv1.Overrides{}
	actionParameterResponse := rulesetsv1.ActionParametersResponse{}

	if reflect.ValueOf(actionParameterObj[CISRulesetOverrides]).IsNil() {
		overrideObj = expandCISRulesetsRulesActionParametersOverrides(actionParameterObj[CISRulesetOverrides])
	}
	if reflect.ValueOf(actionParameterObj[CISRulesetsRuleActionParametersResponse]).IsNil() {
		actionParameterResponse = expandCISRulesetsRulesActionParametersResponse(actionParameterObj[CISRulesetsRuleActionParametersResponse])
	}

	finalResponse := make([]rulesetsv1.ActionParameters, 0)
	actionParameterRespObj := rulesetsv1.ActionParameters{
		ID:        &id,
		Ruleset:   &ruleset,
		Rulesets:  ruleList,
		Version:   &version,
		Overrides: &overrideObj,
		Response:  &actionParameterResponse,
	}
	finalResponse = append(finalResponse, actionParameterRespObj)

	return finalResponse[0]
}

func expandCISRulesetsRulesActionParametersResponse(obj interface{}) rulesetsv1.ActionParametersResponse {
	response := obj.(*schema.Set).List()[0].(map[string]interface{})
	content := response[CISRulesetsRuleActionParametersResponseContent].(string)
	contentType := response[CISRulesetsRuleActionParametersResponseContentType].(string)
	statusCode := response[cisPageRuleActionsValueStatusCode].(int64)

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
	sensitivityOverride := overrideObj[CISRulesetOverridesSensitivityLevel].(string)

	rules := []rulesetsv1.RulesOverride{}
	categories := []rulesetsv1.CategoriesOverride{}
	if reflect.ValueOf(overrideObj[CISRulesetOverridesRules]).IsNil() {
		rules = expandCISRulesetsRulesActionParametersOverridesRules(overrideObj[CISRulesetOverridesRules])
	}
	if reflect.ValueOf(overrideObj[CISRulesetOverridesCategories]).IsNil() {
		categories = expandCISRulesetsRulesActionParametersOverridesCategories(overrideObj[CISRulesetOverridesCategories])
	}

	finalResponse := make([]rulesetsv1.Overrides, 0)
	overrideRespObj := rulesetsv1.Overrides{
		Action:           &actionOverride,
		Enabled:          &enabledOverride,
		SensitivityLevel: &sensitivityOverride,
		Rules:            rules,
		Categories:       categories,
	}
	finalResponse = append(finalResponse, overrideRespObj)

	return finalResponse[0]
}

func expandCISRulesetsRulesActionParametersOverridesCategories(obj interface{}) []rulesetsv1.CategoriesOverride {

	response := obj.([]interface{})[0].(map[string]interface{})

	action := response[CISRulesetOverridesAction].(string)
	enabled := response[CISRulesetOverridesEnabled].(bool)
	category := response[CISRulesetOverridesCategoriesCategory].(string)
	finalResponse := make([]rulesetsv1.CategoriesOverride, 0)
	overrideRespObj := rulesetsv1.CategoriesOverride{
		Action:   &action,
		Enabled:  &enabled,
		Category: &category,
	}
	finalResponse = append(finalResponse, overrideRespObj)

	return finalResponse
}

func expandCISRulesetsRulesActionParametersOverridesRules(obj interface{}) []rulesetsv1.RulesOverride {

	response := obj.([]interface{})[0].(map[string]interface{})
	id := response[CISRulesetsId].(string)
	action := response[CISRulesetOverridesAction].(string)
	enabled := response[CISRulesetOverridesEnabled].(bool)
	sensitivity := response[CISRulesetOverridesSensitivityLevel].(string)

	finalResponse := make([]rulesetsv1.RulesOverride, 0)
	overrideRespObj := rulesetsv1.RulesOverride{
		ID:               &id,
		Action:           &action,
		Enabled:          &enabled,
		SensitivityLevel: &sensitivity,
	}
	finalResponse = append(finalResponse, overrideRespObj)

	return finalResponse
}
