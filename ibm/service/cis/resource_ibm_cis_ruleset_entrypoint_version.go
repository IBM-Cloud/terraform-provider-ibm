// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package cis

import (
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceIBMCISRulesetEntryPointVersion() *schema.Resource {
	return &schema.Resource{
		Read:     ResourceIBMCISRulesetEntryPointVersionRead,
		Create:   ResourceIBMCISRulesetEntryPointVersionUpdate,
		Update:   ResourceIBMCISRulesetEntryPointVersionUpdate,
		Delete:   ResourceIBMCISRulesetEntryPointVersionDelete,
		Importer: &schema.ResourceImporter{},
		Schema: map[string]*schema.Schema{
			cisID: {
				Type:        schema.TypeString,
				Description: "CIS instance crn",
				Required:    true,
			},
			cisDomainID: {
				Type:             schema.TypeString,
				Description:      "Associated CIS domain",
				Optional:         true,
				DiffSuppressFunc: suppressDomainIDDiff,
			},
			CISRulesetPhase: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Ruleset phase",
			},
			CISRulesetsEntryPointOutput: {
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Description: "Input block: ruleset description and rules to deploy.",
				Elem:        CISResourceResponseObject,
			},
			CISRulesetsEntryPointResponseOutput: {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Output block: full ruleset response as returned by the CIS API after apply.",
				Elem:        CISResponseObject,
			},
		},
	}
}
func ResourceIBMCISRulesetEntryPointVersionValidator() *validate.ResourceValidator {
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
		ResourceName: "ibm_cis_ruleset_entrypoint_version",
		Schema:       validateSchema}
	return &ibmCISRulesetValidator
}

func ResourceIBMCISRulesetEntryPointVersionRead(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(conns.ClientSession).CisRulesetsSession()
	if err != nil {
		return flex.FmtErrorf("[ERROR] Error while getting the CisRulesetsSession %s", err)
	}

	ruleset_phase, zoneId, crn, err := flex.ConvertTfToCisThreeVar(d.Id())
	if err != nil {
		return flex.FmtErrorf("[ERROR] Error while ConvertTftoCisThreeVar %s", err)
	}
	sess.Crn = core.StringPtr(crn)

	if zoneId != "" {
		sess.ZoneIdentifier = core.StringPtr(zoneId)
		opt := sess.Clone().NewGetZoneEntrypointRulesetOptions(ruleset_phase)
		result, resp, err := sess.GetZoneEntrypointRuleset(opt)
		if err != nil {
			return flex.FmtErrorf("[ERROR] Get zone ruleset failed: %s %v", err, resp)
		}
		// rulesets (input) is intentionally not overwritten — it holds the
		// user-configured values and must not be overwritten with API values
		// or the SDK will generate a perpetual diff on every plan.
		// rulesets_response is Computed-only: it receives the full API response
		// and is never compared against user config, so no diff is produced.
		d.Set(CISRulesetsEntryPointResponseOutput, flattenCISRulesets(*result.Result))
		d.Set(cisDomainID, zoneId)
		d.Set(cisID, crn)
		d.Set(CISRulesetPhase, ruleset_phase)

	} else {
		opt := sess.NewGetInstanceEntrypointRulesetOptions(ruleset_phase)
		result, resp, err := sess.GetInstanceEntrypointRuleset(opt)
		if err != nil {
			return flex.FmtErrorf("[ERROR] Get instance ruleset failed: %s %v", err, resp)
		}
		// Same rationale as zone branch above.
		d.Set(CISRulesetsEntryPointResponseOutput, flattenCISRulesets(*result.Result))
		d.Set(cisDomainID, zoneId)
		d.Set(cisID, crn)
		d.Set(CISRulesetPhase, ruleset_phase)

	}

	return nil
}

func ResourceIBMCISRulesetEntryPointVersionUpdate(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(conns.ClientSession).CisRulesetsSession()
	if err != nil {
		return flex.FmtErrorf("[ERROR] Error while getting the CisRulesetsSession %s", err)
	}

	crn := d.Get(cisID).(string)
	sess.Crn = core.StringPtr(crn)

	zoneId := d.Get(cisDomainID).(string)
	ruleset_phase := d.Get(CISRulesetPhase).(string)

	if zoneId != "" {
		sess.ZoneIdentifier = &zoneId

		opt := sess.NewUpdateZoneEntrypointRulesetOptions(ruleset_phase)

		rulesetsObject := d.Get(CISRulesetsEntryPointOutput).([]interface{})[0].(map[string]interface{})
		opt.SetDescription(rulesetsObject[CISRulesetsDescription].(string))

		rulesObj := expandCISRules(rulesetsObject[CISRulesetsRules])
		opt.SetRules(rulesObj)

		result, resp, err := sess.UpdateZoneEntrypointRuleset(opt)
		if err != nil || result == nil {
			return flex.FmtErrorf("[ERROR] Error while Update Zone Entrypoint Rulesets %s %s", err, resp)
		}

	} else {
		opt := sess.NewUpdateInstanceEntrypointRulesetOptions(ruleset_phase)

		rulesetsObject := d.Get(CISRulesetsEntryPointOutput).([]interface{})[0].(map[string]interface{})
		opt.SetDescription(rulesetsObject[CISRulesetsDescription].(string))

		rulesObj := expandCISRules(rulesetsObject[CISRulesetsRules])
		opt.SetRules(rulesObj)

		result, resp, err := sess.UpdateInstanceEntrypointRuleset(opt)
		if err != nil || result == nil {
			return flex.FmtErrorf("[ERROR] Error while Update Entrypoint Rulesets %s %s", err, resp)
		}

	}
	d.SetId(dataSourceCISRulesetsEPCheckID(d))
	return ResourceIBMCISRulesetEntryPointVersionRead(d, meta)
}

func ResourceIBMCISRulesetEntryPointVersionDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}
