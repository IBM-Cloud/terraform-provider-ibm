// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package cis

import (
	"log"
	"time"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
	"github.com/IBM/go-sdk-core/v5/core"
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
	CISRulesetsId                                      = "id"
)

var CISRulesetsRuleElement = &schema.Resource{
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
						Type:        schema.TypeSet,
						Computed:    true,
						Description: "Action parameters response of the Rulesets Rule",
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								CISRulesetsRuleActionParametersResponseContent: {
									Type:        schema.TypeString,
									Computed:    true,
									Description: "Action parameters response content of the Rulesets Rule",
								},
								CISRulesetsRuleActionParametersResponseContentType: {
									Type:        schema.TypeString,
									Computed:    true,
									Description: "Action parameters response type of the Rulesets Rule",
								},
								CISRulesetsRuleActionParametersResponseStatusCode: {
									Type:        schema.TypeString,
									Computed:    true,
									Description: "Action parameters response status code of the Rulesets Rule",
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
			Type:        schema.TypeSet,
			Computed:    true,
			Description: "Logging of the Rulesets Rule",
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					CISRulesetsRuleLoggingEnabled: {
						Type:        schema.TypeBool,
						Computed:    true,
						Description: "Logging Enabled",
					},
				},
			},
		},
		CISRulesetsRuleLastUpdatedAt: {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Rulesets Rule Last Updated At",
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
				Elem: &schema.Resource{
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
							Description: "Rules of the Rulesets",
							Elem:        CISRulesetsRuleElement,
						},
					},
				},
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
				log.Printf("[WARN] List all account rulesets failed: %v\n", resp)
				return err
			}
			rulesetObj := result.Result

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
				ruleDetails[CISRulesetsRuleId] = *&ruleDetailsObj.ID
				ruleDetails[CISRulesetsRuleVersion] = *&ruleDetailsObj.Version
				ruleDetails[CISRulesetsRuleAction] = *&ruleDetailsObj.Action
				ruleDetails[CISRulesetsRuleExpression] = *&ruleDetailsObj.Expression
				ruleDetails[CISRulesetsRuleRef] = *&ruleDetailsObj.Ref
				ruleDetails[CISRulesetsRuleLastUpdatedAt] = *&ruleDetailsObj.LastUpdated

				ruleDetailsLoggingObj := map[string]interface{}{}
				ruleDetailsLogging := *&ruleDetailsObj.Logging
				ruleDetailsLoggingObj[CISRulesetsRuleLoggingEnabled] = *ruleDetailsLogging.Enabled
				ruleDetails[CISRulesetsRuleLogging] = ruleDetailsLoggingObj

				ruleDetailsActionParametersObj := map[string]interface{}{}
				ruleDetailsActionParameters := *&ruleDetailsObj.ActionParameters
				ruleDetailsActionParametersResponseObj := map[string]interface{}{}
				ruleDetailsActionParametersResponse := *&ruleDetailsActionParameters.Response
				ruleDetailsActionParametersResponseObj[CISRulesetsRuleActionParametersResponseContent] = *ruleDetailsActionParametersResponse.Content
				ruleDetailsActionParametersResponseObj[CISRulesetsRuleActionParametersResponseContentType] = *ruleDetailsActionParametersResponse.ContentType
				ruleDetailsActionParametersResponseObj[CISRulesetsRuleActionParametersResponseStatusCode] = *ruleDetailsActionParametersResponse.StatusCode
				ruleDetails[CISRulesetsRules] = ruleDetailsActionParametersObj

				ruleDetailsList = append(ruleDetailsList, ruleDetails)
			}

			rulesetOutput[CISRulesetsRules] = ruleDetailsList

			d.SetId(dataSourceCISRulesetsCheckID(d))
			d.Set(CISRulesetsOutput, rulesetOutput)
			d.Set(cisID, crn)

		} else {
			opt := sess.NewGetZoneRulesetsOptions()
			result, resp, err := sess.GetZoneRulesets(opt)
			if err != nil {
				log.Printf("[WARN] List all account rulesets failed: %v\n", resp)
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
			}

			d.SetId(dataSourceCISRulesetsCheckID(d))
			d.Set(CISRulesetsOutput, rulesetList)
			d.Set(cisID, crn)
		}

	} else {

		if rulesetId != "" {
			opt := sess.NewGetAccountRulesetOptions(rulesetId)
			result, resp, err := sess.GetAccountRuleset(opt)
			if err != nil {
				log.Printf("[WARN] List all account rulesets failed: %v\n", resp)
				return err
			}

			rulesetObj := result.Result

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
				ruleDetails[CISRulesetsRuleId] = *&ruleDetailsObj.ID
				ruleDetails[CISRulesetsRuleVersion] = *&ruleDetailsObj.Version
				ruleDetails[CISRulesetsRuleAction] = *&ruleDetailsObj.Action
				ruleDetails[CISRulesetsRuleExpression] = *&ruleDetailsObj.Expression
				ruleDetails[CISRulesetsRuleRef] = *&ruleDetailsObj.Ref
				ruleDetails[CISRulesetsRuleLastUpdatedAt] = *&ruleDetailsObj.LastUpdated

				ruleDetailsLoggingObj := map[string]interface{}{}
				ruleDetailsLogging := *&ruleDetailsObj.Logging
				ruleDetailsLoggingObj[CISRulesetsRuleLoggingEnabled] = *ruleDetailsLogging.Enabled
				ruleDetails[CISRulesetsRuleLogging] = ruleDetailsLoggingObj

				ruleDetailsActionParametersObj := map[string]interface{}{}
				ruleDetailsActionParameters := *&ruleDetailsObj.ActionParameters
				ruleDetailsActionParametersResponseObj := map[string]interface{}{}
				ruleDetailsActionParametersResponse := *&ruleDetailsActionParameters.Response
				ruleDetailsActionParametersResponseObj[CISRulesetsRuleActionParametersResponseContent] = *ruleDetailsActionParametersResponse.Content
				ruleDetailsActionParametersResponseObj[CISRulesetsRuleActionParametersResponseContentType] = *ruleDetailsActionParametersResponse.ContentType
				ruleDetailsActionParametersResponseObj[CISRulesetsRuleActionParametersResponseStatusCode] = *ruleDetailsActionParametersResponse.StatusCode
				ruleDetails[CISRulesetsRules] = ruleDetailsActionParametersObj

				ruleDetailsList = append(ruleDetailsList, ruleDetails)
			}

			rulesetOutput[CISRulesetsRules] = ruleDetailsList

			d.SetId(dataSourceCISRulesetsCheckID(d))
			d.Set(CISRulesetsOutput, rulesetOutput)
			d.Set(cisID, crn)

		} else {
			opt := sess.NewGetAccountRulesetsOptions()
			result, resp, err := sess.GetAccountRulesets(opt)
			if err != nil {
				log.Printf("[WARN] List all account rulesets failed: %v\n", resp)
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
