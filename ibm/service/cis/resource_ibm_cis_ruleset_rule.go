// Copyright IBM Corp. 2024, 2025. All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package cis

import (
	"reflect"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/networking-go-sdk/rulesetsv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var CISRulesetsRulesObject = &schema.Resource{
	Schema: map[string]*schema.Schema{
		CISRulesetsRuleId: {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "ID of the rulesets rule",
			Computed:    true,
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
						Computed:    true,
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
												Description: "Score threshold for the override rule",
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
			Description: "Position of the rulesets rule",
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					CISRulesetsRulePositionAfter: {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "Ruleset before position of the rulesets rule",
					},
					CISRulesetsRulePositionBefore: {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "Ruleset after position of the rulesets rule",
					},
					CISRulesetsRulePositionIndex: {
						Type:        schema.TypeInt,
						Optional:    true,
						Description: "Index of the rulesets rule",
					},
				},
			},
		},
		CISRulesetsRuleRateLimit: {
			Type:        schema.TypeList,
			Optional:    true,
			MaxItems:    1,
			Description: "Ratelimit of the Rulesets Rule",
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					CISRulesetsRuleRateLimitCharacteristics: {
						Type:        schema.TypeList,
						Optional:    true,
						Description: "List of Characteristics of the ratelimit on rulesets rule.",
						Elem:        &schema.Schema{Type: schema.TypeString},
					},
					CISRulesetsRuleRateLimitCountingExpression: {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "Counting expression of the ratelimit on rulesets rule.",
					},
					CISRulesetsRuleRateLimitMitigationTimeout: {
						Type:        schema.TypeInt,
						Optional:    true,
						Description: "Mitigation timeout of the ratelimit on rulesets rule.",
					},
					CISRulesetsRuleRateLimitPeriod: {
						Type:        schema.TypeInt,
						Optional:    true,
						Description: "Period of the ratelimit on rulesets rule.",
					},
					CISRulesetsRuleRateLimitRequestsPerPeriod: {
						Type:        schema.TypeInt,
						Optional:    true,
						Description: "Requests per period of the ratelimit on rulesets rule.",
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
}

func ResourceIBMCISRulesetRule() *schema.Resource {
	return &schema.Resource{
		Create:   ResourceIBMCISRulesetRuleCreate,
		Read:     ResourceIBMCISRulesetRuleRead,
		Update:   ResourceIBMCISRulesetRuleUpdate,
		Delete:   ResourceIBMCISRulesetRuleDelete,
		Importer: &schema.ResourceImporter{},
		Schema: map[string]*schema.Schema{
			cisID: {
				Type:        schema.TypeString,
				Description: "CIS instance CRN",
				Required:    true,
				ValidateFunc: validate.InvokeValidator("ibm_cis_ruleset_rule",
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
				Required:    true,
			},
			CISRulesetsRule: {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Rules of the rulesets",
				MaxItems:    1,
				Elem:        CISRulesetsRulesObject,
			},
		},
	}
}
func ResourceIBMCISRulesetRuleValidator() *validate.ResourceValidator {
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
		ResourceName: "ibm_cis_ruleset_rule",
		Schema:       validateSchema}
	return &ibmCISRulesetValidator
}

func ResourceIBMCISRulesetRuleCreate(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(conns.ClientSession).CisRulesetsSession()
	if err != nil {
		return flex.FmtErrorf("[ERROR] Error while getting the CisRulesetsSession %s", err)
	}
	crn := d.Get(cisID).(string)
	zoneId := d.Get(cisDomainID).(string)
	rulesetId := d.Get(CISRulesetsId).(string)

	sess.Crn = core.StringPtr(crn)

	if zoneId != "" {
		sess.ZoneIdentifier = core.StringPtr(zoneId)
		opt := sess.NewCreateZoneRulesetRuleOptions(rulesetId)

		rulesObject := d.Get(CISRulesetsRule).([]interface{})[0].(map[string]interface{})

		opt.SetRulesetID(rulesetId)
		opt.SetExpression(rulesObject[CISRulesetsRuleExpression].(string))
		opt.SetAction(rulesObject[CISRulesetsRuleAction].(string))
		opt.SetDescription(rulesObject[CISRulesetsRuleActionDescription].(string))
		opt.SetEnabled(rulesObject[CISRulesetsRuleActionEnabled].(bool))
		opt.SetRef(rulesObject[CISRulesetsRuleRef].(string))

		position := rulesetsv1.Position{}
		if !reflect.ValueOf(rulesObject[CISRulesetsRulePosition]).IsNil() {
			position, err = expandCISRulesetsRulesPositions(rulesObject[CISRulesetsRulePosition])
			if err != nil {
				return flex.FmtErrorf("[ERROR] Error while creating the zone Rule %s", err)
			}
		}
		opt.SetPosition(&position)

		ratelimit := rulesetsv1.Ratelimit{}
		if v, ok := rulesObject[CISRulesetsRuleRateLimit]; ok && v != nil {
			ratelimit, err = expandCISRulesetsRulesRateLimits(v)
			if err != nil {
				return flex.FmtErrorf("[ERROR] Error while creating the zone Rule: %s", err)
			}

			if !DataSourceCISRulesetsRuleIsEmptyRateLimit(ratelimit) {
				opt.SetRatelimit(&ratelimit)
			}
		}

		actionParameterObj := rulesetsv1.ActionParameters{}
		if len(rulesObject[CISRulesetsRuleActionParameters].(*schema.Set).List()) != 0 {
			actionParameterObj = expandCISRulesetsRulesActionParameters(rulesObject[CISRulesetsRuleActionParameters])
		}
		opt.SetActionParameters(&actionParameterObj)

		result, resp, err := sess.CreateZoneRulesetRule(opt)

		if err != nil {
			return flex.FmtErrorf("[ERROR] Error while creating the zone Rule %s", resp)
		}
		len_rules := len(result.Result.Rules)

		// When creating a rule response is resulted as list of rules.
		// To get the index we have to check if index,after or before is provided by the user.
		// If not provided then we will take the last rule from the list as the new rule is added at the end.

		rule_id := ""
		if len(rulesObject[CISRulesetsRulePosition].(*schema.Set).List()) != 0 {
			response := rulesObject[CISRulesetsRulePosition].(*schema.Set).List()[0].(map[string]interface{})
			before := response[CISRulesetsRulePositionBefore].(string)
			after := response[CISRulesetsRulePositionAfter].(string)
			index := int64(response[CISRulesetsRulePositionIndex].(int))

			if after != "" {
				for i, rule := range result.Result.Rules {
					if *rule.ID == after {
						opt.SetID(*result.Result.Rules[i+1].ID)
						rule_id = *result.Result.Rules[i+1].ID
						break
					}
				}
			} else if before != "" {
				for i, rule := range result.Result.Rules {
					if *rule.ID == before {
						opt.SetID(*result.Result.Rules[i-1].ID)
						rule_id = *result.Result.Rules[i-1].ID
						break
					}
				}
			} else if index != 0 {
				opt.SetID(*result.Result.Rules[index-1].ID)
				rule_id = *result.Result.Rules[index-1].ID
			}

		} else {
			opt.SetID(*result.Result.Rules[len_rules-1].ID)
			rule_id = *result.Result.Rules[len_rules-1].ID
		}

		d.SetId(dataSourceCISRulesetsRuleCheckID(d, rule_id))

	} else {
		opt := sess.NewCreateInstanceRulesetRuleOptions(rulesetId)

		rulesObject := d.Get(CISRulesetsRule).([]interface{})[0].(map[string]interface{})

		opt.SetRulesetID(rulesetId)
		opt.SetExpression(rulesObject[CISRulesetsRuleExpression].(string))
		opt.SetAction(rulesObject[CISRulesetsRuleAction].(string))
		opt.SetDescription(rulesObject[CISRulesetsRuleActionDescription].(string))
		opt.SetEnabled(rulesObject[CISRulesetsRuleActionEnabled].(bool))
		opt.SetRef(rulesObject[CISRulesetsRuleRef].(string))

		position := rulesetsv1.Position{}
		if reflect.ValueOf(rulesObject[CISRulesetsRulePosition]).IsNil() {
			position, err = expandCISRulesetsRulesPositions(rulesObject[CISRulesetsRulePosition])
			if err != nil {
				return flex.FmtErrorf("[ERROR] Error while creating the instance Rule %s", err)
			}
		}
		opt.SetPosition(&position)

		actionParameterObj := rulesetsv1.ActionParameters{}
		if len(rulesObject[CISRulesetsRuleActionParameters].(*schema.Set).List()) != 0 {
			actionParameterObj = expandCISRulesetsRulesActionParameters(rulesObject[CISRulesetsRuleActionParameters])
		}
		opt.SetActionParameters(&actionParameterObj)

		result, resp, err := sess.CreateInstanceRulesetRule(opt)

		if err != nil {
			return flex.FmtErrorf("[ERROR] Error while creating the instance Rule %s", resp)
		}

		len_rules := len(result.Result.Rules)
		opt.SetID(*result.Result.Rules[len_rules-1].ID)

		d.SetId(dataSourceCISRulesetsRuleCheckID(d, *result.Result.Rules[len_rules-1].ID))
	}
	return nil
}

func ResourceIBMCISRulesetRuleRead(d *schema.ResourceData, meta interface{}) error {

	return nil
}

func ResourceIBMCISRulesetRuleUpdate(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(conns.ClientSession).CisRulesetsSession()
	if err != nil {
		return flex.FmtErrorf("[ERROR] Error while getting the CisRulesetsSession %s", err)
	}

	ruleId, rulesetId, zoneId, crn, _ := flex.ConvertTfToCisFourVar(d.Id())
	sess.Crn = core.StringPtr(crn)

	if zoneId != "" {
		sess.ZoneIdentifier = core.StringPtr(zoneId)

		opt := sess.NewUpdateZoneRulesetRuleOptions(rulesetId, ruleId)

		rulesetsRuleObject := d.Get(CISRulesetsRule).([]interface{})[0].(map[string]interface{})
		opt.SetDescription(rulesetsRuleObject[CISRulesetsDescription].(string))
		opt.SetAction(rulesetsRuleObject[CISRulesetsRuleAction].(string))
		if rulesetsRuleObject[CISRulesetsRuleActionParameters] != nil {
			actionParameters := expandCISRulesetsRulesActionParameters(rulesetsRuleObject[CISRulesetsRuleActionParameters])
			opt.SetActionParameters(&actionParameters)
		}
		opt.SetEnabled(rulesetsRuleObject[CISRulesetsRuleActionEnabled].(bool))
		opt.SetExpression(rulesetsRuleObject[CISRulesetsRuleExpression].(string))
		opt.SetRef(rulesetsRuleObject[CISRulesetsRuleRef].(string))
		if d.HasChange(CISRulesetsRule + ".0." + CISRulesetsRulePosition) {
			position, positionError := expandCISRulesetsRulesPositions(rulesetsRuleObject[CISRulesetsRulePosition])
			if positionError != nil {
				return flex.FmtErrorf("[ERROR] Error while updating the zone Ruleset %s", positionError)
			}
			opt.SetPosition(&position)
		}

		if v, ok := rulesetsRuleObject[CISRulesetsRuleRateLimit]; ok && v != nil {
			ratelimit, ratelimitErr := expandCISRulesetsRulesRateLimits(v)
			if ratelimitErr != nil {
				return flex.FmtErrorf("[ERROR] Error while updating the zone Ruleset: %s", ratelimitErr)
			}

			if !DataSourceCISRulesetsRuleIsEmptyRateLimit(ratelimit) {
				opt.SetRatelimit(&ratelimit)
			}
		}

		opt.SetRulesetID(rulesetId)
		opt.SetRuleID(ruleId)
		opt.SetID(ruleId)

		_, _, err := sess.UpdateZoneRulesetRule(opt)

		if err != nil {
			return flex.FmtErrorf("[ERROR] Error while updating the zone Ruleset %s", err)
		}

		d.SetId(dataSourceCISRulesetsRuleCheckID(d, ruleId))

	} else {
		opt := sess.NewUpdateInstanceRulesetRuleOptions(rulesetId, ruleId)

		rulesetsRuleObject := d.Get(CISRulesetsRule).([]interface{})[0].(map[string]interface{})
		opt.SetDescription(rulesetsRuleObject[CISRulesetsDescription].(string))
		opt.SetAction(rulesetsRuleObject[CISRulesetsRuleAction].(string))
		actionParameters := expandCISRulesetsRulesActionParameters(rulesetsRuleObject[CISRulesetsRuleActionParameters])
		opt.SetActionParameters(&actionParameters)
		opt.SetEnabled(rulesetsRuleObject[CISRulesetsRuleActionEnabled].(bool))
		opt.SetExpression(rulesetsRuleObject[CISRulesetsRuleExpression].(string))
		opt.SetRef(rulesetsRuleObject[CISRulesetsRuleAction].(string))
		position, err := expandCISRulesetsRulesPositions(rulesetsRuleObject[CISRulesetsRulePosition])
		if err != nil {
			return flex.FmtErrorf("[ERROR] Error while updating the instance Ruleset %s", err)
		}
		opt.SetPosition(&position)

		opt.SetRulesetID(rulesetId)
		opt.SetRuleID(ruleId)
		opt.SetID(ruleId)

		_, _, err = sess.UpdateInstanceRulesetRule(opt)

		if err != nil {
			return flex.FmtErrorf("[ERROR] Error while updating the instance Ruleset %s", err)
		}

		d.SetId(dataSourceCISRulesetsRuleCheckID(d, ruleId))
	}
	return nil
}

func ResourceIBMCISRulesetRuleDelete(d *schema.ResourceData, meta interface{}) error {

	sess, err := meta.(conns.ClientSession).CisRulesetsSession()
	if err != nil {
		return flex.FmtErrorf("[ERROR] Error while getting the CisRulesetsSession %s", err)
	}

	ruleId, rulesetId, zoneId, crn, _ := flex.ConvertTfToCisFourVar(d.Id())
	sess.Crn = core.StringPtr(crn)

	if zoneId != "" {
		sess.ZoneIdentifier = core.StringPtr(zoneId)
		opt := sess.NewDeleteZoneRulesetRuleOptions(rulesetId, ruleId)
		_, res, err := sess.DeleteZoneRulesetRule(opt)
		if err != nil {
			return flex.FmtErrorf("[ERROR] Error deleting the zone ruleset rule %s:%s", err, res)
		}
	} else {
		opt := sess.NewDeleteInstanceRulesetRuleOptions(rulesetId, ruleId)
		_, res, err := sess.DeleteInstanceRulesetRule(opt)
		if err != nil {
			return flex.FmtErrorf("[ERROR] Error deleting the Instance ruleset rule %s:%s", err, res)
		}
	}

	d.SetId("")
	return nil
}

func dataSourceCISRulesetsRuleCheckID(d *schema.ResourceData, ruleId string) string {
	return ruleId + ":" + d.Get(CISRulesetsId).(string) + ":" + d.Get(cisDomainID).(string) + ":" + d.Get(cisID).(string)
}

func DataSourceCISRulesetsRuleIsEmptyRateLimit(r rulesetsv1.Ratelimit) bool {
	return len(r.Characteristics) == 0 &&
		r.CountingExpression == nil &&
		r.MitigationTimeout == nil &&
		r.Period == nil &&
		r.RequestsPerPeriod == nil
}
