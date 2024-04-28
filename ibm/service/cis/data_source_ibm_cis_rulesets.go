// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package cis

import (
	"encoding/json"
	"log"
	"reflect"
	"time"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/networking-go-sdk/rulesetsv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	CISRulesetsOutput                                  = "rulesets"
	CISRulesetsDescription                             = "description"
	CISRulesetsKind                                    = "kind"
	CISRulesetsName                                    = "name"
	CISRulesetsPhase                                   = "phase"
	CISRulesetsLastUpdatedAt                           = "last_updated"
	CISRulesetsVersion                                 = "version"
	CISRulesetsRules                                   = "rules"
	CISRulesetsRuleId                                  = "id"
	CISRulesetsRuleVersion                             = "rule_version"
	CISRulesetsRuleAction                              = "rule_action"
	CISRulesetsRuleActionParameters                    = "action_parameters"
	CISRulesetsRuleActionParametersResponse            = "response"
	CISRulesetsRuleActionParametersResponseContent     = "content"
	CISRulesetsRuleActionParametersResponseContentType = "content_type"
	CISRulesetsRuleActionParametersResponseStatusCode  = "status_code"
	CISRulesetsRuleExpression                          = "rule_expression"
	CISRulesetsRuleRef                                 = "rule_ref"
	CISRulesetsRuleLogging                             = "rule_logging"
	CISRulesetsRuleLoggingEnabled                      = "enabled"
	CISRulesetsRuleLastUpdatedAt                       = "rule_last_updated_at"
	CISRulesetsId                                      = "ruleset_id"
	CISRuleset                                         = "ruleset"
	CISRulesetList                                     = "rulesets"
	CISRulesetOverrides                                = "overrides"
	CISRulesetOverridesAction                          = "action"
	CISRulesetOverridesEnabled                         = "enabled"
	CISRulesetOverridesSensitivityLevel                = "sensitivity_level"
	CISRulesetOverridesCategories                      = "categories"
	CISRulesetOverridesCategoriesCategory              = "category"
	CISRulesetOverridesRules                           = "rules"
)

var CISResponseObject = &schema.Resource{
	Schema: map[string]*schema.Schema{
		CISRulesetsDescription: {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Description of the Rulesets",
		},
		CISRulesetsKind: {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Kind of the Rulesets",
		},
		CISRulesetsName: {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Name of the Rulesets",
		},
		CISRulesetsPhase: {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Phase of the Rulesets",
		},
		CISRulesetsLastUpdatedAt: {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Rulesets Last Updated At",
		},
		CISRulesetsVersion: {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Version of the Rulesets",
		},
		CISRulesetsRules: {
			Type:        schema.TypeList,
			Computed:    true,
			Optional:    true,
			Description: "Rules of the Rulesets",
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					CISRulesetsRuleId: {
						Type:        schema.TypeString,
						Computed:    true,
						Description: "Id of the Rulesets Rule",
					},
					CISRulesetsRuleVersion: {
						Type:        schema.TypeString,
						Computed:    true,
						Description: "Version of the Rulesets Rule",
					},
					CISRulesetsRuleAction: {
						Type:        schema.TypeString,
						Computed:    true,
						Description: "Action of the Rulesets Rule",
					},
					CISRulesetsRuleActionParameters: {
						Type:        schema.TypeSet,
						Computed:    true,
						Description: "Action parameters of the Rulesets Rule",
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								CISRulesetsRuleActionParametersResponse: {
									Type:        schema.TypeMap,
									Computed:    true,
									Description: "Action parameters response of the Rulesets Rule",
									// Elem: &schema.Resource{
									// 	Schema: map[string]*schema.Schema{
									// 		CISRulesetsRuleActionParametersResponseContent: {
									// 			Type:        schema.TypeString,
									// 			Computed:    true,
									// 			Description: "Action parameters response content of the Rulesets Rule",
									// 		},
									// 		CISRulesetsRuleActionParametersResponseContentType: {
									// 			Type:        schema.TypeString,
									// 			Computed:    true,
									// 			Description: "Action parameters response type of the Rulesets Rule",
									// 		},
									// 		CISRulesetsRuleActionParametersResponseStatusCode: {
									// 			Type:        schema.TypeString,
									// 			Computed:    true,
									// 			Description: "Action parameters response status code of the Rulesets Rule",
									// 		},
									// 	},
									// },
								},
								CISRulesetsRuleId: {
									Type:        schema.TypeString,
									Computed:    true,
									Description: "Id of the Rulesets Rule",
								},
								CISRuleset: {
									Type:        schema.TypeString,
									Computed:    true,
									Description: "Ruleset ID of the ruleset to apply action to.",
								},
								CISRulesetsVersion: {
									Type:        schema.TypeString,
									Computed:    true,
									Description: "Version of the Ruleset",
								},
								CISRulesetList: {
									Type:        schema.TypeList,
									Computed:    true,
									Description: "List of Ruleset IDs of the ruleset to apply action to.",
									Elem:        schema.TypeString,
								},
								CISRulesetOverrides: {
									Type:        schema.TypeSet,
									Computed:    true,
									Description: "Override options",
									Elem: &schema.Resource{
										Schema: map[string]*schema.Schema{
											CISRulesetOverridesAction: {
												Type:        schema.TypeString,
												Computed:    true,
												Description: "Action to perform",
											},
											CISRulesetOverridesEnabled: {
												Type:        schema.TypeBool,
												Computed:    true,
												Description: "Enable Disable Rule",
											},
											CISRulesetOverridesSensitivityLevel: {
												Type:        schema.TypeString,
												Computed:    true,
												Description: "Sensitivity Level",
											},
											CISRulesetOverridesCategories: {
												Type:        schema.TypeList,
												Computed:    true,
												Description: "Categories",
												Elem: &schema.Resource{
													Schema: map[string]*schema.Schema{
														CISRulesetOverridesCategoriesCategory: {
															Type:        schema.TypeString,
															Computed:    true,
															Description: "Category",
														},
														CISRulesetOverridesEnabled: {
															Type:        schema.TypeBool,
															Computed:    true,
															Description: "Enable Disable Rule",
														},
														CISRulesetOverridesAction: {
															Type:        schema.TypeString,
															Computed:    true,
															Description: "Action to perform",
														},
													},
												},
											},
											CISRulesetOverridesRules: {
												Type:        schema.TypeList,
												Computed:    true,
												Description: "Rules",
												Elem: &schema.Resource{
													Schema: map[string]*schema.Schema{
														CISRulesetsId: {
															Type:        schema.TypeString,
															Computed:    true,
															Description: "Id",
														},
														CISRulesetOverridesEnabled: {
															Type:        schema.TypeBool,
															Computed:    true,
															Description: "Enable Disable Rule",
														},
														CISRulesetOverridesAction: {
															Type:        schema.TypeString,
															Computed:    true,
															Description: "Action to perform",
														},
														CISRulesetOverridesSensitivityLevel: {
															Type:        schema.TypeString,
															Computed:    true,
															Description: "Sensitivity Level",
														},
													},
												},
											},
										},
									},
								},
							},
						},
					},
					CISRulesetsRuleExpression: {
						Type:        schema.TypeString,
						Computed:    true,
						Description: "Experession of the Rulesets Rule",
					},
					CISRulesetsRuleRef: {
						Type:        schema.TypeString,
						Computed:    true,
						Description: "Reference of the Rulesets Rule",
					},
					CISRulesetsRuleLogging: {
						Type:        schema.TypeMap,
						Computed:    true,
						Description: "Logging of the Rulesets Rule",
						Elem:        &schema.Schema{Type: schema.TypeBool},
						// Elem: &schema.Resource{
						// 	Schema: map[string]*schema.Schema{
						// 		CISRulesetsRuleLoggingEnabled: {
						// 			Type:        schema.TypeBool,
						// 			Computed:    true,
						// 			Description: "Logging Enabled",
						// 		},
						// 	},
						// },
					},
					CISRulesetsRuleLastUpdatedAt: {
						Type:        schema.TypeString,
						Computed:    true,
						Description: "Rulesets Rule Last Updated At",
					},
				},
			},
		},
		CISRulesetsId: {
			Type:        schema.TypeString,
			Description: "Associated Ruleset ID",
			Computed:    true,
		},
	},
}

func DataSourceIBMCISRulesets() *schema.Resource {
	return &schema.Resource{
		Read: dataIBMCISRulesetsRead,
		Schema: map[string]*schema.Schema{
			cisID: {
				Type:        schema.TypeString,
				Description: "CIS instance crn",
				Required:    true,
				ValidateFunc: validate.InvokeDataSourceValidator(
					"ibm_cis_rulesets",
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
			CISRulesetsOutput: {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Container for response information.",
				Elem:        CISResponseObject,
			},
			"rulesets_obj": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "Container for response information.",
				Elem:        CISResponseObject,
			},
		},
	}
}

func DataSourceIBMCISRulesetsValidator() *validate.ResourceValidator {

	validateSchema := make([]validate.ValidateSchema, 0)

	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 "cis_id",
			ValidateFunctionIdentifier: validate.ValidateCloudData,
			Type:                       validate.TypeString,
			CloudDataType:              "resource_instance",
			CloudDataRange:             []string{"service:internet-svcs"},
			Required:                   true})

	iBMCISRulesetsValidator := validate.ResourceValidator{
		ResourceName: "ibm_cis_rulesets",
		Schema:       validateSchema}
	return &iBMCISRulesetsValidator
}

func dataIBMCISRulesetsRead(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(conns.ClientSession).CisRulesetsSession()
	if err != nil {
		return err
	}
	crn := d.Get(cisID).(string)
	sess.Crn = core.StringPtr(crn)

	zoneId := d.Get(cisDomainID).(string)
	rulesetId := d.Get(CISRulesetsId).(string)

	if zoneId != "" {
		sess.ZoneIdentifier = core.StringPtr(zoneId)

		if rulesetId != "" {
			opt := sess.NewGetZoneRulesetOptions(rulesetId)
			result, resp, err := sess.GetZoneRuleset(opt)
			if err != nil {
				log.Printf("[WARN] List all Instance rulesets failed: %v\n", resp)
				return err
			}
			rulesetObj := flattenCISRulesets(*result.Result)

			d.SetId(dataSourceCISRulesetsCheckID(d))
			d.Set("rulesets_obj", rulesetObj)
			d.Set(cisID, crn)

		} else {
			opt := sess.NewGetZoneRulesetsOptions()
			result, resp, err := sess.GetZoneRulesets(opt)
			if err != nil {
				log.Printf("[WARN] List all Instance rulesets failed: %v\n", resp)
				return err
			}

			rulesetList := make([]map[string]interface{}, 0)
			for _, rulesetObj := range result.Result {
				rulesetOutput := map[string]interface{}{}
				rulesetOutput[CISRulesetsDescription] = *rulesetObj.Description
				rulesetOutput[CISRulesetsKind] = *rulesetObj.Kind
				rulesetOutput[CISRulesetsName] = *rulesetObj.Name
				rulesetOutput[CISRulesetsPhase] = *rulesetObj.Phase
				rulesetOutput[CISRulesetsLastUpdatedAt] = *rulesetObj.LastUpdated
				rulesetOutput[CISRulesetsVersion] = *rulesetObj.Version
				rulesetOutput[CISRulesetsId] = *&rulesetObj.ID

				if rulesetOutput[CISRulesetsPhase] == "http_request_firewall_managed" {
					rulesetList = append(rulesetList, rulesetOutput)
				}
			}

			d.SetId(dataSourceCISRulesetsCheckID(d))
			d.Set(CISRulesetsOutput, rulesetList)
			d.Set(cisID, crn)
		}

	} else {

		if rulesetId != "" {
			opt := sess.NewGetInstanceRulesetOptions(rulesetId)
			result, resp, err := sess.GetInstanceRuleset(opt)
			if err != nil {
				log.Printf("[WARN] List all Instance rulesets failed: %v\n", resp)
				return err
			}

			rulesetObj := flattenCISRulesets(*result.Result)

			d.SetId(dataSourceCISRulesetsCheckID(d))
			d.Set(CISRulesetsOutput, rulesetObj)
			d.Set(cisID, crn)

		} else {
			opt := sess.NewGetInstanceRulesetsOptions()
			result, resp, err := sess.GetInstanceRulesets(opt)
			if err != nil {
				log.Printf("[WARN] List all Instance rulesets failed: %v\n", resp)
				return err
			}

			rulesetList := make([]map[string]interface{}, 0)
			for _, rulesetObj := range result.Result {
				rulesetOutput := map[string]interface{}{}
				rulesetOutput[CISRulesetsDescription] = *rulesetObj.Description
				rulesetOutput[CISRulesetsKind] = *rulesetObj.Kind
				rulesetOutput[CISRulesetsName] = *rulesetObj.Name
				rulesetOutput[CISRulesetsPhase] = *rulesetObj.Phase
				rulesetOutput[CISRulesetsLastUpdatedAt] = *rulesetObj.LastUpdated
				rulesetOutput[CISRulesetsVersion] = *rulesetObj.Version
				rulesetOutput[CISRulesetsId] = *&rulesetObj.ID

				if rulesetOutput[CISRulesetsKind] == "http_request_firewall_managed" {
					rulesetList = append(rulesetList, rulesetOutput)
				}
			}

			d.SetId(dataSourceCISRulesetsCheckID(d))
			d.Set(CISRulesetsOutput, rulesetList)
			d.Set(cisID, crn)
		}
	}

	return nil
}

func dataSourceCISRulesetsCheckID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}

func flattenCISRulesets(rulesetObj rulesetsv1.RulesetDetails) interface{} {

	finalrulesetObj := make([]interface{}, 0)

	rulesetOutput := map[string]interface{}{}
	rulesetOutput[CISRulesetsDescription] = *rulesetObj.Description
	rulesetOutput[CISRulesetsKind] = *rulesetObj.Kind
	rulesetOutput[CISRulesetsName] = *rulesetObj.Name
	rulesetOutput[CISRulesetsPhase] = *rulesetObj.Phase
	rulesetOutput[CISRulesetsLastUpdatedAt] = *rulesetObj.LastUpdated
	rulesetOutput[CISRulesetsVersion] = *rulesetObj.Version

	ruleDetailsList := make([]map[string]interface{}, 0)
	for _, ruleDetailsObj := range rulesetObj.Rules {
		ruleDetails := map[string]interface{}{}
		ruleDetails[CISRulesetsRuleId] = ruleDetailsObj.ID
		ruleDetails[CISRulesetsRuleVersion] = ruleDetailsObj.Version
		ruleDetails[CISRulesetsRuleAction] = ruleDetailsObj.Action
		ruleDetails[CISRulesetsRuleExpression] = ruleDetailsObj.Expression
		ruleDetails[CISRulesetsRuleRef] = ruleDetailsObj.Ref
		ruleDetails[CISRulesetsRuleLastUpdatedAt] = ruleDetailsObj.LastUpdated
		ruleDetails[CISRulesetsRuleLogging] = ruleDetailsObj.Logging

		flattenedActionParameter := flattenCISRulesetsRuleActionParameters(ruleDetailsObj.ActionParameters)
		// Check if returned interface value is nil
		if flattenedActionParameter == nil || reflect.ValueOf(flattenedActionParameter).IsNil() {
			ruleDetails[CISRulesetsRuleActionParameters] = flattenedActionParameter
		}

		ruleDetailsList = append(ruleDetailsList, ruleDetails)
	}

	rulesetOutput[CISRulesetsRules] = ruleDetailsList

	finalrulesetObj = append(finalrulesetObj, rulesetOutput)

	return finalrulesetObj
}

func flattenCISRulesetsRuleActionParameters(rulesetsRuleActionParameterObj *rulesetsv1.ActionParameters) interface{} {
	resultObj := make([]interface{}, 0)
	actionParametersOutput := map[string]interface{}{}
	resultOutput := map[string]interface{}{}

	res, _ := json.Marshal(rulesetsRuleActionParameterObj)
	json.Unmarshal(res, &actionParametersOutput)

	if val, ok := actionParametersOutput["id"]; ok {
		resultOutput[CISRulesetsRuleId] = val.(string)
	}
	if val, ok := actionParametersOutput["ruleset"]; ok {
		resultOutput[CISRuleset] = val.(string)
	}
	if val, ok := actionParametersOutput["version"]; ok {
		resultOutput[CISRulesetsVersion] = val.(string)
	}

	if val, ok := actionParametersOutput["rulesets"]; ok {
		resultOutput[CISRulesetList] = val
	}
	if val, ok := actionParametersOutput["response"]; ok {
		resultOutput[CISRulesetsRuleActionParametersResponse] = val
	}

	if val, ok := actionParametersOutput["overrides"]; ok {
		resultOutput[CISRulesetOverrides] = val
	}

	resultObj = append(resultObj, resultOutput)

	return resultObj
}
