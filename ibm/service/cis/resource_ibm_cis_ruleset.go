// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package cis

import (
	"fmt"
	"log"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/networking-go-sdk/rulesetsv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceIBMCISRuleset() *schema.Resource {
	return &schema.Resource{
		Read:     ResourceIBMCISRulesetRead,
		Update:   ResourceIBMCISRulesetUpdate,
		Delete:   ResourceIBMCISRulesetDelete,
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
				Required:    true,
			},
			CISRulesetsOutput: {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Container for response information.",
				Elem:        CISResponseObject,
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

func ResourceIBMCISRulesetUpdate(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(conns.ClientSession).CisRulesetsSession()
	if err != nil {
		return fmt.Errorf("[ERROR] Error while getting the CisRulesetsSession %s", err)
	}

	zoneId := d.Get(cisDomainID).(string)
	rulesetId := d.Get(CISRulesetsId).(string)

	if zoneId != "" {

		opt := sess.NewUpdateZoneRulesetOptions(rulesetId)

		rulesetObject := d.Get(CISRulesetsOutput).(rulesetsv1.RulesetDetails)
		ruleObject := d.Get(CISRulesetsRules).([]rulesetsv1.RuleCreate)
		opt.SetDescription(*rulesetObject.Description)
		opt.SetRules(ruleObject)

		result, _, err := sess.UpdateZoneRuleset(opt)

		if err != nil {
			return fmt.Errorf("[ERROR] Error while updating the zone Ruleset %s", err)
		}

		d.SetId(*result.Result.ID)

	} else {

		opt := sess.NewUpdateAccountRulesetOptions(rulesetId)

		rulesetObject := d.Get(CISRulesetsOutput).(rulesetsv1.RulesetDetails)
		ruleObject := d.Get(CISRulesetsRules).([]rulesetsv1.RuleCreate)
		opt.SetDescription(*rulesetObject.Description)
		opt.SetRules(ruleObject)

		result, _, err := sess.UpdateAccountRuleset(opt)

		if err != nil {
			return fmt.Errorf("[ERROR] Error while updating the account Ruleset %s", err)
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
			log.Printf("[WARN] Get account ruleset failed: %v\n", resp)
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
		opt := sess.NewDeleteAccountRulesetOptions(rulesetId)
		res, err := sess.DeleteAccountRuleset(opt)
		if err != nil {
			return fmt.Errorf("[ERROR] Error deleting the account ruleset %s:%s", err, res)
		}
	}

	d.SetId("")
	return nil
}
