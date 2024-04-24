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
	CISRulesetsEntryPointOutput = "rulesets"
	CISRulesetPhase             = "ruleset_phase"
	CISRulesetsPhaseListAll     = "list_all"
	CISRulesetVersion           = "ruleset_version"
)

func DataSourceIBMCISRulesetsEntrypointVersions() *schema.Resource {
	return &schema.Resource{
		Read: dataIBMCISRulesetsEntrypointVersionsRead,
		Schema: map[string]*schema.Schema{
			cisID: {
				Type:        schema.TypeString,
				Description: "CIS instance crn",
				Required:    true,
				ValidateFunc: validate.InvokeDataSourceValidator(
					"ibm_cis_rulesets_versions",
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
				Required:    true,
				Description: "Id",
			},
			CISRulesetPhase: {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Ruleset phase",
			},
			CISRulesetsPhaseListAll: {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Ruleset phase",
				Default:     false,
			},
			CISRulesetsEntryPointOutput: {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Container for response information.",
				Elem:        CISResponseObject,
			},
			CISRulesetVersion: {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Ruleset phase",
			},
		},
	}
}
func DataSourceIBMCISRulesetsEntrypointVersionsValidator() *validate.ResourceValidator {

	validateSchema := make([]validate.ValidateSchema, 0)

	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 "cis_id",
			ValidateFunctionIdentifier: validate.ValidateCloudData,
			Type:                       validate.TypeString,
			CloudDataType:              "resource_instance",
			CloudDataRange:             []string{"service:internet-svcs"},
			Required:                   true})

	iBMCISRulesetsEntrypointVersionsValidator := validate.ResourceValidator{
		ResourceName: "ibm_cis_rulesets_entrypoint_versions",
		Schema:       validateSchema}
	return &iBMCISRulesetsEntrypointVersionsValidator
}
func dataIBMCISRulesetsEntrypointVersionsRead(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(conns.ClientSession).CisRulesetsSession()
	if err != nil {
		return err
	}
	crn := d.Get(cisID).(string)
	sess.Crn = core.StringPtr(crn)

	zoneId := d.Get(cisDomainID).(string)
	ruleset_phase := d.Get(CISRulesetPhase).(string)
	ruleset_version := d.Get(CISRulesetVersion).(string)
	list_all := d.Get(CISRulesetsPhaseListAll).(bool)

	if list_all == true {
		opt := sess.NewGetZoneEntryPointRulesetVersionsOptions(ruleset_phase)
		result, resp, err := sess.GetZoneEntryPointRulesetVersions(opt)
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
		}

		d.SetId(dataSourceCISRulesetsCheckID(d))
		d.Set(CISRulesetsOutput, rulesetList)
		d.Set(cisID, crn)

	} else {

		if zoneId != "" {
			sess.ZoneIdentifier = core.StringPtr(zoneId)

			if ruleset_version != "" {
				opt := sess.NewGetZoneEntryPointRulesetVersionOptions(ruleset_phase, ruleset_version)
				result, resp, err := sess.GetZoneEntryPointRulesetVersion(opt)
				if err != nil {
					log.Printf("[WARN] List all Instance rulesets failed: %v\n", resp)
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
				opt := sess.NewGetZoneEntrypointRulesetOptions(ruleset_phase)
				result, resp, err := sess.GetZoneEntrypointRuleset(opt)
				if err != nil {
					log.Printf("[WARN] List all Instance rulesets failed: %v\n", resp)
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
			}

		} else {

			if ruleset_version != "" {
				opt := sess.NewGetInstanceEntryPointRulesetVersionOptions(ruleset_phase, ruleset_version)
				result, resp, err := sess.GetInstanceEntryPointRulesetVersion(opt)
				if err != nil {
					log.Printf("[WARN] List all Instance rulesets failed: %v\n", resp)
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
				opt := sess.NewGetInstanceEntrypointRulesetOptions(ruleset_phase)
				result, resp, err := sess.GetInstanceEntrypointRuleset(opt)
				if err != nil {
					log.Printf("[WARN] List all Instance rulesets failed: %v\n", resp)
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
			}
		}
	}

	return nil
}
func dataSourceCISRulesetsEntrypointVersionsCheckID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}
