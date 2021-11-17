// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"context"
	"fmt"

	"github.com/IBM/networking-go-sdk/firewallrulesv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	ibmCISFirewallrules         = "ibm_cis_firewall_rules"
	cisFirewallrulesID          = "firewall_rule_id"
	cisFilter                   = "filter"
	cisFirewallrulesAction      = "action"
	cisFirewallrulesPaused      = "paused"
	cisFirewallrulesPriority    = "priority"
	cisFirewallrulesDescription = "description"
	cisFirewallrulesList        = "firewall_rules"
)

func resourceIBMCISFirewallrules() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMCISFirewallrulesCreate,
		ReadContext:   resourceIBMCISFirewallrulesRead,
		UpdateContext: resourceIBMCISFirewallrulesUpdate,
		DeleteContext: resourceIBMCISFirewallrulesDelete,
		Importer:      &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			cisID: {
				Type:        schema.TypeString,
				Description: "CIS instance crn",
				Required:    true,
			},
			cisDomainID: {
				Type:             schema.TypeString,
				Description:      "Associated CIS domain",
				Required:         true,
				DiffSuppressFunc: suppressDomainIDDiff,
			},
			cisFilterID: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Firewallrules Existing FilterID",
			},
			cisFirewallrulesAction: {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: InvokeValidator(ibmCISFirewallrules, cisFirewallrulesAction),
				Description:  "Firewallrules Action",
			},
			cisFirewallrulesPriority: {
				Type:         schema.TypeInt,
				Description:  "Firewallrules Action",
				Optional:     true,
				Computed:     true,
				ValidateFunc: InvokeValidator(ibmCISFirewallrules, cisFirewallrulesPriority),
			},
			cisFirewallrulesDescription: {
				Type:         schema.TypeString,
				Optional:     true,
				Description:  "Firewallrules Description",
				ValidateFunc: InvokeValidator(ibmCISFirewallrules, cisFirewallrulesDescription),
			},
			cisFirewallrulesPaused: {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Firewallrules Paused",
			},
		},
	}
}

func resourceIBMCISFirewallrulesCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	sess, err := meta.(ClientSession).BluemixSession()
	if err != nil {
		return diag.FromErr(err)
	}
	xAuthtoken := sess.Config.IAMAccessToken

	cisClient, err := meta.(ClientSession).CisFirewallRulesSession()
	if err != nil {
		return diag.FromErr(err)
	}

	crn := d.Get(cisID).(string)
	zoneID, _, err := convertTftoCisTwoVar(d.Get(cisDomainID).(string))

	var newFirewallRules firewallrulesv1.FirewallRuleInputWithFilterID

	if a, ok := d.GetOk(cisFirewallrulesAction); ok {
		action := a.(string)
		newFirewallRules.Action = &action
	}
	if des, ok := d.GetOk(cisFilterDescription); ok {
		description := des.(string)
		newFirewallRules.Description = &description
	}
	if id, ok := d.GetOk(cisFilterID); ok {
		filterid := id.(string)
		filterModel, _ := cisClient.NewFirewallRuleInputWithFilterIdFilter(filterid)
		newFirewallRules.Filter = filterModel
	}

	opt := cisClient.NewCreateFirewallRulesOptions(xAuthtoken, crn, zoneID)

	opt.SetFirewallRuleInputWithFilterID([]firewallrulesv1.FirewallRuleInputWithFilterID{newFirewallRules})

	result, _, err := cisClient.CreateFirewallRulesWithContext(context, opt)
	if err != nil || result == nil {
		return diag.FromErr(fmt.Errorf("Error reading the  %s", err))
	}
	d.SetId(convertCisToTfThreeVar(*result.Result[0].ID, zoneID, crn))

	return resourceIBMCISFirewallrulesRead(context, d, meta)

}
func resourceIBMCISFirewallrulesRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := meta.(ClientSession).BluemixSession()
	if err != nil {
		return diag.FromErr(err)
	}
	xAuthtoken := sess.Config.IAMAccessToken

	cisClient, err := meta.(ClientSession).CisFirewallRulesSession()
	if err != nil {
		return diag.FromErr(err)
	}
	firwallruleID, zoneID, crn, err := convertTfToCisThreeVar(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	opt := cisClient.NewGetFirewallRuleOptions(xAuthtoken, crn, zoneID, firwallruleID)

	result, response, err := cisClient.GetFirewallRuleWithContext(context, opt)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		return diag.FromErr(fmt.Errorf("Error reading the firewall rules %s:%s", err, response))
	}
	d.Set(cisID, crn)
	d.Set(cisDomainID, zoneID)
	d.Set(cisFilterID, result.Result.Filter.ID)
	d.Set(cisFirewallrulesAction, result.Result.Action)
	d.Set(cisFirewallrulesPaused, result.Result.Paused)
	d.Set(cisFilterDescription, result.Result.Description)

	return nil
}
func resourceIBMCISFirewallrulesUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := meta.(ClientSession).BluemixSession()
	if err != nil {
		return diag.FromErr(err)
	}
	xAuthtoken := sess.Config.IAMAccessToken

	cisClient, err := meta.(ClientSession).CisFirewallRulesSession()
	if err != nil {
		return diag.FromErr(err)
	}

	firewallruleID, zoneID, crn, err := convertTfToCisThreeVar(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	if d.HasChange(cisFilterID) ||
		d.HasChange(cisFirewallrulesAction) ||
		d.HasChange(cisFirewallrulesPaused) ||
		d.HasChange(cisFilterDescription) {

		var updatefirewallrules firewallrulesv1.FirewallRulesUpdateInputItem
		updatefirewallrules.ID = &firewallruleID

		if a, ok := d.GetOk(cisFirewallrulesAction); ok {
			action := a.(string)
			updatefirewallrules.Action = &action
		}
		if p, ok := d.GetOk(cisFirewallrulesPaused); ok {
			paused := p.(bool)
			updatefirewallrules.Paused = &paused
		}
		if des, ok := d.GetOk(cisFilterDescription); ok {
			description := des.(string)
			updatefirewallrules.Description = &description
		}

		if id, ok := d.GetOk(cisFilterID); ok {
			filterid := id.(string)
			filterUpdate, _ := cisClient.NewFirewallRulesUpdateInputItemFilter(filterid)
			updatefirewallrules.Filter = filterUpdate
		}
		opt := cisClient.NewUpdateFirewllRulesOptions(xAuthtoken, crn, zoneID)

		opt.SetFirewallRulesUpdateInputItem([]firewallrulesv1.FirewallRulesUpdateInputItem{updatefirewallrules})

		result, _, err := cisClient.UpdateFirewllRulesWithContext(context, opt)
		if err != nil || result == nil {
			return diag.FromErr(fmt.Errorf("Error updating the firewall rules %s", err))
		}
	}
	return resourceIBMCISFirewallrulesRead(context, d, meta)
}
func resourceIBMCISFirewallrulesDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	sess, err := meta.(ClientSession).BluemixSession()
	if err != nil {
		return diag.FromErr(err)
	}
	xAuthtoken := sess.Config.IAMAccessToken

	cisClient, err := meta.(ClientSession).CisFirewallRulesSession()
	if err != nil {
		return diag.FromErr(err)
	}

	firewallruleid, zoneID, crn, err := convertTfToCisThreeVar(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}
	opt := cisClient.NewDeleteFirewallRulesOptions(xAuthtoken, crn, zoneID, firewallruleid)
	_, response, err := cisClient.DeleteFirewallRulesWithContext(context, opt)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			return nil
		}
		return diag.FromErr(fmt.Errorf("Error deleting the  custom resolver %s:%s", err, response))
	}

	d.SetId("")
	return nil
}
func resourceIBMCISFirewallrulesValidator() *ResourceValidator {
	validateSchema := make([]ValidateSchema, 1)

	validateSchema = append(validateSchema,
		ValidateSchema{
			Identifier:                 cisFirewallrulesAction,
			ValidateFunctionIdentifier: ValidateAllowedStringValue,
			Type:                       TypeString,
			Required:                   true,
			AllowedValues:              "log, allow, challenge, js_challenge, block"})
	validateSchema = append(validateSchema,
		ValidateSchema{
			Identifier:                 cisFirewallrulesDescription,
			ValidateFunctionIdentifier: ValidateAllowedStringValue,
			Type:                       TypeString,
			Required:                   true,
			AllowedValues:              "Firewallrules-creation"})

	validateSchema = append(validateSchema,
		ValidateSchema{
			Identifier:                 cisFirewallrulesPriority,
			ValidateFunctionIdentifier: IntBetween,
			Type:                       TypeInt,
			Optional:                   true,
			MinValue:                   "1",
			MaxValue:                   "2147483647"})
	ibmCISFirewallrulesResourceValidator := ResourceValidator{ResourceName: ibmCISFirewallrules, Schema: validateSchema}
	return &ibmCISFirewallrulesResourceValidator
}
