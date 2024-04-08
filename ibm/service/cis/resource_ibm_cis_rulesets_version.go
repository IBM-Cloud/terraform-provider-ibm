// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package cis

import (
	"fmt"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceIBMCISRulesetsVersion() *schema.Resource {
	return &schema.Resource{
		Delete:   ResourceIBMCISRulesetsVersionDelete,
		Importer: &schema.ResourceImporter{},
		Schema: map[string]*schema.Schema{
			cisID: {
				Type:        schema.TypeString,
				Description: "CIS instance crn",
				Required:    true,
				ValidateFunc: validate.InvokeValidator("ibm_cis_rulesets_version",
					"cis_id"),
			},
		},
	}
}
func ResourceIBMCISRulesetsVersionValidator() *validate.ResourceValidator {
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
		ResourceName: "ibm_cis_rulesets_version",
		Schema:       validateSchema}
	return &ibmCISRulesetValidator
}

func ResourceIBMCISRulesetsVersionDelete(d *schema.ResourceData, meta interface{}) error {

	_, err := meta.(conns.ClientSession).CisRulesetsSession()
	if err != nil {
		return fmt.Errorf("[ERROR] Error while getting the CisRulesetsSession %s", err)
	}
	// alertID, crn, err := flex.ConvertTftoCisTwoVar(d.Id())
	// if err != nil {
	// 	return err
	// }
	// sess.Crn = core.StringPtr(crn)
	// opt := sess.NewDeleteRulesetOptions(alertID)
	// _, response, err := sess.DeleteRuleset(opt)
	// if err != nil {
	// 	if response != nil && response.StatusCode == 404 {
	// 		return nil
	// 	}
	// 	return fmt.Errorf("[ERROR] Error deleting the alert %s:%s", err, response)
	// }
	return nil
}
