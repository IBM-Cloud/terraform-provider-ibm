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
	CISRulesetsRuleTag = "rulesets_rule_tag"
)

func DataSourceIBMCISRulesetsRulesByTag() *schema.Resource {
	return &schema.Resource{
		Read: dataIBMCISRulesetsRead,
		Schema: map[string]*schema.Schema{
			cisID: {
				Type:        schema.TypeString,
				Description: "CIS instance crn",
				Required:    true,
				ValidateFunc: validate.InvokeDataSourceValidator(
					"ibm_cis_rulesets_rules_by_tag",
					"cis_id"),
			},
			CISRulesetsId: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Id",
			},
			CISRulesetVersion: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Ruleset version",
			},
			CISRulesetsRuleTag: {
				Type:        schema.TypeBool,
				Required:    true,
				Description: "Rulesets rule tag",
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
func DataSourceIBMCISRulesetsRulesByTagValidator() *validate.ResourceValidator {

	validateSchema := make([]validate.ValidateSchema, 0)

	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 "cis_id",
			ValidateFunctionIdentifier: validate.ValidateCloudData,
			Type:                       validate.TypeString,
			CloudDataType:              "resource_instance",
			CloudDataRange:             []string{"service:internet-svcs"},
			Required:                   true})

	iBMCISRulesetsRulesValidator := validate.ResourceValidator{
		ResourceName: "ibm_cis_rulesets_rules_by_tag",
		Schema:       validateSchema}
	return &iBMCISRulesetsRulesValidator
}
func dataIBMCISRulesetsRulesRead(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(conns.ClientSession).CisRulesetsSession()
	if err != nil {
		return err
	}
	crn := d.Get(cisID).(string)
	sess.Crn = core.StringPtr(crn)

	rulesetId := d.Get(CISRulesetsId).(string)
	rulesetVersion := d.Get(CISRulesetVersion).(string)
	rulesetRuleTag := d.Get(CISRulesetsRuleTag).(string)

	opt := sess.NewGetInstanceRulesetVersionByTagOptions(rulesetId, rulesetVersion, rulesetRuleTag)
	result, resp, err := sess.GetInstanceRulesetVersionByTag(opt)
	if err != nil {
		log.Printf("[WARN] List all rulesets version rules by tag failed: %v\n", resp)
		return err
	}

	rulesetObj := flattenCISRulesets(d, *result.Result)

	d.SetId(dataSourceCISRulesetsCheckID(d))
	d.Set(CISRulesetsOutput, rulesetObj)
	d.Set(cisID, crn)

	return nil
}
func dataSourceCISRulesetsRulesCheckID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}
