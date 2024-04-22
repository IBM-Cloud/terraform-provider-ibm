// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package cis

import (
	"fmt"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/networking-go-sdk/rulesetsv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceIBMCISRulesetsEntryPointVersion() *schema.Resource {
	return &schema.Resource{
		Read:     ResourceIBMCISRulesetsEntryPointVersionRead,
		Update:   ResourceIBMCISRulesetsEntryPointVersionUpdate,
		Delete:   ResourceIBMCISRulesetsEntryPointVersionDelete,
		Importer: &schema.ResourceImporter{},
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
			CISRulesetPhase: {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Ruleset phase",
			},
			CISRulesetsEntryPointOutput: {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Container for response information.",
				Elem:        CISResponseObject,
			},
		},
	}
}
func ResourceIBMCISRulesetsEntryPointVersionValidator() *validate.ResourceValidator {
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
		ResourceName: "ibm_cis_rulesets_entrypoint_version",
		Schema:       validateSchema}
	return &ibmCISRulesetValidator
}

func ResourceIBMCISRulesetsEntryPointVersionRead(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(conns.ClientSession).CisRulesetsSession()
	if err != nil {
		return fmt.Errorf("[ERROR] Error while getting the CisRulesetsSession %s", err)
	}

	crn := d.Get(cisID).(string)
	sess.Crn = core.StringPtr(crn)

	zoneId := d.Get(cisDomainID).(string)
	rulesetId := d.Id()

	if zoneId != "" {
		sess.ZoneIdentifier = core.StringPtr(zoneId)
		opt := sess.NewGetZoneRulesetOptions(rulesetId)
		result, resp, err := sess.GetZoneRuleset(opt)
		if err != nil {
			return fmt.Errorf("[WARN] Get zone ruleset failed: %v\n", resp)
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
		opt := sess.NewGetAccountRulesetOptions(rulesetId)
		result, resp, err := sess.GetAccountRuleset(opt)
		if err != nil {
			return fmt.Errorf("[WARN] Get account ruleset failed: %v\n", resp)
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

	return nil
}

func ResourceIBMCISRulesetsEntryPointVersionUpdate(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(conns.ClientSession).CisRulesetsSession()
	if err != nil {
		return fmt.Errorf("[ERROR] Error while getting the CisRulesetsSession %s", err)
	}

	ruleset_phase, zoneId, crn, err := flex.ConvertTfToCisThreeVar(d.Id())
	if err != nil {
		return fmt.Errorf("[ERROR] Error while ConvertTftoCisThreeVar %s", err)
	}
	sess.Crn = core.StringPtr(crn)

	if zoneId != "" {
		sess.ZoneIdentifier = &zoneId

		opt := sess.NewUpdateZoneEntrypointRulesetOptions(ruleset_phase)

		if d.HasChange(CISRulesetsDescription) {
			if val, ok := d.GetOk(CISRulesetsDescription); ok {
				opt.SetDescription(val.(string))
			}
		}
		if d.HasChange(CISRulesetsKind) {
			if val, ok := d.GetOk(CISRulesetsKind); ok {
				opt.SetKind(val.(string))
			}
		}
		if d.HasChange(CISRulesetsName) {
			if val, ok := d.GetOk(CISRulesetsName); ok {
				opt.SetName(val.(string))
			}
		}
		if d.HasChange(CISRulesetsRules) {
			if val, ok := d.GetOk(CISRulesetsRules); ok {
				opt.SetRules(val.([]rulesetsv1.RuleCreate))
			}
		}

		result, resp, err := sess.UpdateZoneEntrypointRuleset(opt)
		if err != nil || result == nil {
			return fmt.Errorf("[ERROR] Error while Update Zone Entrypoint Rulesets %s %s", err, resp)
		}

	} else {
		opt := sess.NewUpdateAccountEntrypointRulesetOptions(ruleset_phase)

		if d.HasChange(CISRulesetsDescription) {
			if val, ok := d.GetOk(CISRulesetsDescription); ok {
				opt.SetDescription(val.(string))
			}
		}
		if d.HasChange(CISRulesetsKind) {
			if val, ok := d.GetOk(CISRulesetsKind); ok {
				opt.SetKind(val.(string))
			}
		}
		if d.HasChange(CISRulesetsName) {
			if val, ok := d.GetOk(CISRulesetsName); ok {
				opt.SetName(val.(string))
			}
		}
		if d.HasChange(CISRulesetsRules) {
			if val, ok := d.GetOk(CISRulesetsRules); ok {
				opt.SetRules(val.([]rulesetsv1.RuleCreate))
			}
		}

		result, resp, err := sess.UpdateAccountEntrypointRuleset(opt)
		if err != nil || result == nil {
			return fmt.Errorf("[ERROR] Error while Update Entrypoint Rulesets %s %s", err, resp)
		}

	}
	return ResourceIBMCISRulesetsEntryPointVersionRead(d, meta)
}

func ResourceIBMCISRulesetsEntryPointVersionDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}
