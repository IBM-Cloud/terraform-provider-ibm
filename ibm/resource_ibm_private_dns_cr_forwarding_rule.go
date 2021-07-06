// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	pdnsCRForwardRule   = "ibm_dns_cr_forwarding_rule"
	pdnsCRFRResolverID  = "resolver_id"
	pdnsCRFRDesctiption = "description"
	pdnsCRFRType        = "type"
	pdnsCRFRMatch       = "match"
	pdnsCRFRForwardTo   = "forward_to"
	pdnsCRFRRuleID      = "rule_id"
	pdnsCRFRCreatedOn   = "created_on"
	pdnsCRFRModifiedOn  = "modified_on"
)

func resourceIbmDnsCrForwardingRule() *schema.Resource {
	return &schema.Resource{
		Create:   resourceIbmDnsCrForwardingRuleCreate,
		Read:     resourceIbmDnsCrForwardingRuleRead,
		Update:   resourceIbmDnsCrForwardingRuleUpdate,
		Delete:   resourceIbmDnsCrForwardingRuleDelete,
		Importer: &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			pdnsInstanceID: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The unique identifier of a service instance.",
			},
			pdnsCRFRResolverID: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The unique identifier of a custom resolver.",
			},
			pdnsCRFRDesctiption: {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Descriptive text of the forwarding rule.",
			},
			pdnsCRFRType: {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: InvokeValidator(pdnsCRForwardRule, "type"),
				Description:  "Type of the forwarding rule.",
			},
			pdnsCRFRMatch: {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The matching zone or hostname.",
			},
			pdnsCRFRForwardTo: {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "The upstream DNS servers will be forwarded to.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			pdnsCRFRRuleID: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "the time when a forwarding rule ID is created, RFC3339 format.",
			},
			pdnsCRFRCreatedOn: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "the time when a forwarding rule is created, RFC3339 format.",
			},
			pdnsCRFRModifiedOn: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "the recent time when a forwarding rule is modified, RFC3339 format.",
			},
		},
	}
}

func resourceIbmDnsCrForwardingRuleValidator() *ResourceValidator {
	validateSchema := make([]ValidateSchema, 1)
	validateSchema = append(validateSchema,
		ValidateSchema{
			Identifier:                 "type",
			ValidateFunctionIdentifier: ValidateAllowedStringValue,
			Type:                       TypeString,
			Optional:                   true,
			AllowedValues:              "hostname, zone",
		},
	)

	resourceValidator := ResourceValidator{ResourceName: pdnsCRForwardRule, Schema: validateSchema}
	return &resourceValidator
}

func resourceIbmDnsCrForwardingRuleCreate(d *schema.ResourceData, meta interface{}) error {
	dnsSvcsClient, err := meta.(ClientSession).PrivateDNSClientSessionScoped()
	if err != nil {
		return err
	}
	instanceID := d.Get(pdnsInstanceID).(string)
	resolverID := d.Get(pdnsCRFRResolverID).(string)
	opt := dnsSvcsClient.NewCreateForwardingRuleOptions(instanceID, resolverID)

	if des, ok := d.GetOk(pdnsCRFRDesctiption); ok {
		opt.SetDescription(des.(string))
	}
	if t, ok := d.GetOk(pdnsCRFRType); ok {
		opt.SetType(t.(string))
	}
	if m, ok := d.GetOk(pdnsCRFRMatch); ok {
		opt.SetMatch(m.(string))
	}
	if _, ok := d.GetOk(pdnsCRFRForwardTo); ok {
		opt.SetForwardTo(expandStringList(d.Get(pdnsCRFRForwardTo).([]interface{})))
	}
	result, resp, err := dnsSvcsClient.CreateForwardingRule(opt)

	if err != nil || result == nil {
		return fmt.Errorf("Error while adding forword rule :%s %s", err, resp)
	}
	d.SetId(convertCisToTfThreeVar(*result.ID, resolverID, instanceID))

	return resourceIbmDnsCrForwardingRuleRead(d, meta)
}

func resourceIbmDnsCrForwardingRuleRead(d *schema.ResourceData, meta interface{}) error {
	dnsSvcsClient, err := meta.(ClientSession).PrivateDNSClientSessionScoped()
	if err != nil {
		return err
	}
	ruleID, resolverID, instanceID, err := convertTfToCisThreeVar(d.Id())
	opt := dnsSvcsClient.NewGetForwardingRuleOptions(instanceID, resolverID, ruleID)
	result, resp, err := dnsSvcsClient.GetForwardingRule(opt)
	if err != nil {
		return fmt.Errorf("GetForwardingRule  failed %s\n%s", err, resp)
	}
	d.Set(pdnsInstanceID, instanceID)
	d.Set(pdnsCRFRResolverID, resolverID)
	d.Set(pdnsCRFRRuleID, ruleID)
	d.Set(pdnsCRFRDesctiption, *result.Description)
	d.Set(pdnsCRFRType, *result.Type)
	d.Set(pdnsCRFRMatch, *result.Match)
	d.Set(pdnsCRFRForwardTo, result.ForwardTo)
	return nil

}
func resourceIbmDnsCrForwardingRuleUpdate(d *schema.ResourceData, meta interface{}) error {
	dnsSvcsClient, err := meta.(ClientSession).PrivateDNSClientSessionScoped()
	if err != nil {
		return err
	}
	ruleID, resolverID, instanceID, err := convertTfToCisThreeVar(d.Id())

	if err != nil {
		return err
	}

	opt := dnsSvcsClient.NewUpdateForwardingRuleOptions(instanceID, resolverID, ruleID)
	if d.HasChange(pdnsCRFRDesctiption) ||
		d.HasChange(pdnsCRFRMatch) ||
		d.HasChange(pdnsCRFRForwardTo) {

		if des, ok := d.GetOk(pdnsCRFRDesctiption); ok {
			frdes := des.(string)
			opt.SetDescription(frdes)
		}
		if ma, ok := d.GetOk(pdnsCRFRMatch); ok {
			frmatch := ma.(string)
			opt.SetMatch(frmatch)
		}
		if _, ok := d.GetOk(pdnsCRFRForwardTo); ok {
			opt.SetForwardTo(expandStringList(d.Get(pdnsCRFRForwardTo).([]interface{})))
		}

		result, _, err := dnsSvcsClient.UpdateForwardingRule(opt)
		if err != nil {
			return fmt.Errorf("Error updating pdns ForwardingRule: %s", err)
		}
		if *result.ID == "" {
			return fmt.Errorf("Error failed to find id in Update response; resource was empty")
		}
	}
	return resourceIbmDnsCrForwardingRuleRead(d, meta)
}

func resourceIbmDnsCrForwardingRuleDelete(d *schema.ResourceData, meta interface{}) error {
	dnsSvcsClient, err := meta.(ClientSession).PrivateDNSClientSessionScoped()
	if err != nil {
		return err
	}
	ruleID, resolverID, instanceID, err := convertTfToCisThreeVar(d.Id())
	opt := dnsSvcsClient.NewDeleteForwardingRuleOptions(instanceID, resolverID, ruleID)
	response, err := dnsSvcsClient.DeleteForwardingRule(opt)
	if err != nil {
		log.Printf("[DEBUG] DeleteForwardingRule failed %s\n%s", err, response)
		return fmt.Errorf("DeleteForwardingRule failed %s\n%s", err, response)
	}
	d.SetId("")
	return nil
}
