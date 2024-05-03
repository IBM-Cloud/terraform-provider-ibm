// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package cis

import (
	"fmt"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceIBMCISRulesetsRule() *schema.Resource {
	return &schema.Resource{
		Create:   ResourceIBMCISRulesetsRuleCreate,
		Read:     ResourceIBMCISRulesetsRuleRead,
		Update:   ResourceIBMCISRulesetsRuleUpdate,
		Delete:   ResourceIBMCISRulesetsRuleDelete,
		Importer: &schema.ResourceImporter{},
		Schema: map[string]*schema.Schema{
			cisID: {
				Type:        schema.TypeString,
				Description: "CIS instance crn",
				Required:    true,
				ValidateFunc: validate.InvokeValidator("ibm_cis_rulesets_rule",
					"cis_id"),
			},
		},
	}
}
func ResourceIBMCISRulesetsRuleValidator() *validate.ResourceValidator {
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
		ResourceName: "ibm_cis_rulesets_rule",
		Schema:       validateSchema}
	return &ibmCISRulesetValidator
}

func ResourceIBMCISRulesetsRuleCreate(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(conns.ClientSession).CisRulesetsSession()
	if err != nil {
		return fmt.Errorf("[ERROR] Error while getting the CisRulesetsSession %s", err)
	}
	crn := d.Get(cisID).(string)
	sess.Crn = core.StringPtr(crn)

	return nil
}

func ResourceIBMCISRulesetsRuleRead(d *schema.ResourceData, meta interface{}) error {
	
	return nil
}

func ResourceIBMCISRulesetsRuleUpdate(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(conns.ClientSession).CisRulesetsSession()
	if err != nil {
		return fmt.Errorf("[ERROR] Error while getting the CisRulesetsSession %s", err)
	}

	crn := d.Get(cisID).(string)
	zoneId := d.Get(cisDomainID).(string)
	rulesetId := d.Get(CISRulesetsId).(string)
	ruleId := d.Get(CISRulesetsRuleId).(string)
	sess.Crn = core.StringPtr(crn)

	if zoneId != "" {
		sess.ZoneIdentifier = core.StringPtr(zoneId)

		opt := sess.NewUpdateZoneRulesetRuleOptions(rulesetId, ruleId)

		rulesetsRuleObject := d.Get(CISRulesetsObjectOutput).([]interface{})[0].(map[string]interface{})
		opt.SetDescription(rulesetsRuleObject[CISRulesetsDescription].(string))
		opt.SetAction(rulesetsRuleObject[CISRulesetsRuleAction].(string))
		actionParameters := expandCISRulesetsRulesActionParameters(rulesetsRuleObject[CISRulesetsRuleActionParameters])
		opt.SetActionParameters(&actionParameters)
		opt.SetEnabled(rulesetsRuleObject[CISRulesetsRuleActionEnabled].(bool))
		opt.SetExpression(rulesetsRuleObject[CISRulesetsRuleExpression].(string))
		logging := expandCISRulesetsRulesLogging(rulesetsRuleObject[CISRulesetsRuleLogging])
		opt.SetLogging(&logging)
		opt.SetRef(rulesetsRuleObject[CISRulesetsRuleAction].(string))
		position := expandCISRulesetsRulesPositions(rulesetsRuleObject[CISRulesetsRulePosition])
		opt.SetPosition(&position)
		
		opt.SetRulesetID(ruleId)
		opt.SetRuleID(ruleId)
		opt.SetID(ruleId)

		result, _, err := sess.UpdateZoneRulesetRule(opt)

		if err != nil {
			return fmt.Errorf("[ERROR] Error while updating the zone Ruleset %s", err)
		}

		d.SetId(*result.Result.ID)

	} else {
		opt := sess.NewUpdateInstanceRulesetRuleOptions(rulesetId, ruleId)

		rulesetsRuleObject := d.Get(CISRulesetsObjectOutput).([]interface{})[0].(map[string]interface{})
		opt.SetDescription(rulesetsRuleObject[CISRulesetsDescription].(string))
		opt.SetAction(rulesetsRuleObject[CISRulesetsRuleAction].(string))
		actionParameters := expandCISRulesetsRulesActionParameters(rulesetsRuleObject[CISRulesetsRuleActionParameters])
		opt.SetActionParameters(&actionParameters)
		opt.SetEnabled(rulesetsRuleObject[CISRulesetsRuleActionEnabled].(bool))
		opt.SetExpression(rulesetsRuleObject[CISRulesetsRuleExpression].(string))
		logging := expandCISRulesetsRulesLogging(rulesetsRuleObject[CISRulesetsRuleLogging])
		opt.SetLogging(&logging)
		opt.SetRef(rulesetsRuleObject[CISRulesetsRuleAction].(string))
		position := expandCISRulesetsRulesPositions(rulesetsRuleObject[CISRulesetsRulePosition])
		opt.SetPosition(&position)
		
		opt.SetRulesetID(ruleId)
		opt.SetRuleID(ruleId)
		opt.SetID(ruleId)

		result, _, err := sess.UpdateInstanceRulesetRule(opt)

		if err != nil {
			return fmt.Errorf("[ERROR] Error while updating the zone Ruleset %s", err)
		}

		d.SetId(*result.Result.ID)

	}

	return nil
}

func ResourceIBMCISRulesetsRuleDelete(d *schema.ResourceData, meta interface{}) error {

	sess, err := meta.(conns.ClientSession).CisRulesetsSession()
	if err != nil {
		return fmt.Errorf("[ERROR] Error while getting the CisRulesetsSession %s", err)
	}
	crn := d.Get(cisID).(string)
	sess.Crn = core.StringPtr(crn)

	zoneId := d.Get(cisDomainID).(string)
	rulesetId := d.Get(CISRulesetsId).(string)
	ruleId := d.Get(CISRulesetsRuleId).(string)

	if zoneId != "" {
		opt := sess.NewDeleteZoneRulesetRuleOptions(rulesetId, ruleId)
		_, res, err := sess.DeleteZoneRulesetRule(opt)
		if err != nil {
			return fmt.Errorf("[ERROR] Error deleting the zone ruleset rule %s:%s", err, res)
		}
	} else {
		opt := sess.NewDeleteInstanceRulesetRuleOptions(rulesetId, ruleId)
		_, res, err := sess.DeleteInstanceRulesetRule(opt)
		if err != nil {
			return fmt.Errorf("[ERROR] Error deleting the Instance ruleset rule %s:%s", err, res)
		}
	}

	d.SetId("")
	return nil
}
