// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package dnsservices

import (
	"context"
	"fmt"
	"strings"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
	"github.com/IBM/go-sdk-core/v5/core"
	dns "github.com/IBM/networking-go-sdk/dnssvcsv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	pdnsCRForwardRule    = "ibm_dns_custom_resolver_forwarding_rule"
	pdnsCRForwardRules   = "rules"
	pdnsCRFRResolverID   = "resolver_id"
	pdnsCRFRDesctiption  = "description"
	pdnsCRFRType         = "type"
	pdnsCRFRMatch        = "match"
	pdnsCRFRForwardTo    = "forward_to"
	pdnsCRFRRuleID       = "rule_id"
	pdnsCRFRCreatedOn    = "created_on"
	pdnsCRFRModifiedOn   = "modified_on"
	pdnsCRFRViews        = "views"
	pdnsCRFRVName        = "view_name"
	pdnsCRFRVDescription = "view_description"
	pdnsCRFRVExpression  = "view_expression"
	pdnsCRFRVForwardTo   = "view_forward_to"
)

func ResourceIBMPrivateDNSForwardingRule() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIbmDnsCrForwardingRuleCreate,
		ReadContext:   resourceIbmDnsCrForwardingRuleRead,
		UpdateContext: resourceIbmDnsCrForwardingRuleUpdate,
		DeleteContext: resourceIbmDnsCrForwardingRuleDelete,
		Importer:      &schema.ResourceImporter{},

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
				ValidateFunc: validate.InvokeValidator(pdnsCRForwardRule, pdnsCRFRType),
				Description:  "Type of the forwarding rule.",
			},
			pdnsCRFRMatch: {
				Type:        schema.TypeString,
				Computed:    true,
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
			pdnsCRFRViews: {
				Type:        schema.TypeSet,
				Description: "An array of views used by forwarding rules",
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						pdnsCRFRVName: {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Unique name of the view.",
						},
						pdnsCRFRVDescription: {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Description of the view.",
						},
						pdnsCRFRVExpression: {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Expression of the view.",
						},
						pdnsCRFRVForwardTo: {
							Type:        schema.TypeList,
							Required:    true,
							Description: "The upstream DNS servers will be forwarded to.",
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
					},
				},
			},
		},
	}
}

func ResourceIBMPrivateDNSForwardingRuleValidator() *validate.ResourceValidator {
	validateSchema := make([]validate.ValidateSchema, 0)
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 "type",
			ValidateFunctionIdentifier: validate.ValidateAllowedStringValue,
			Type:                       validate.TypeString,
			Optional:                   true,
			AllowedValues:              "hostname, zone, Default",
		},
	)

	resourceValidator := validate.ResourceValidator{ResourceName: pdnsCRForwardRule, Schema: validateSchema}
	return &resourceValidator
}

func resourceIbmDnsCrForwardingRuleCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	dnsSvcsClient, err := meta.(conns.ClientSession).PrivateDNSClientSession()
	if err != nil {
		return diag.FromErr(err)
	}
	instanceID := d.Get(pdnsInstanceID).(string)
	resolverID := d.Get(pdnsCRFRResolverID).(string)

	ruleType := d.Get(pdnsCRFRType).(string)
	ruleMatch := d.Get(pdnsCRFRMatch).(string)

	views := d.Get(pdnsCRFRViews).(*schema.Set)

	var forwardingRuleInp dns.ForwardingRuleInputIntf

	if forward, ok := d.GetOk(pdnsCRFRForwardTo); ok {
		if _, ok := d.GetOk(pdnsCRFRViews); ok {
			forwardingRuleInp = new(dns.ForwardingRuleInputForwardingRuleBoth)
			forwardingRuleInp, _ = dnsSvcsClient.NewForwardingRuleInputForwardingRuleBoth(ruleType, ruleMatch, flex.ExpandStringList(forward.([]interface{})), expandPDNSFRViews(views))
		} else {
			forwardingRuleInp, _ = dnsSvcsClient.NewForwardingRuleInputForwardingRuleOnlyForward(ruleType, ruleMatch, flex.ExpandStringList(forward.([]interface{})))
		}

	} else {
		forwardingRuleInp, _ = dnsSvcsClient.NewForwardingRuleInputForwardingRuleOnlyView(ruleType, ruleMatch, expandPDNSFRViews(views))
	}

	opt := dnsSvcsClient.NewCreateForwardingRuleOptions(instanceID, resolverID, forwardingRuleInp)

	result, resp, err := dnsSvcsClient.CreateForwardingRuleWithContext(context, opt)

	if err != nil || result == nil {
		return diag.FromErr(fmt.Errorf("[ERROR] Error creating the forwarding rules %s:%s", err, resp))
	}
	d.SetId(flex.ConvertCisToTfThreeVar(*result.ID, resolverID, instanceID))

	return resourceIbmDnsCrForwardingRuleRead(context, d, meta)
}

func resourceIbmDnsCrForwardingRuleRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	dnsSvcsClient, err := meta.(conns.ClientSession).PrivateDNSClientSession()
	if err != nil {
		return diag.FromErr(err)
	}
	ruleID, resolverID, instanceID, err := flex.ConvertTfToCisThreeVar(d.Id())
	opt := dnsSvcsClient.NewGetForwardingRuleOptions(instanceID, resolverID, ruleID)
	result, resp, err := dnsSvcsClient.GetForwardingRuleWithContext(context, opt)

	if err != nil || result == nil {
		if resp != nil && resp.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		return diag.FromErr(fmt.Errorf("[ERROR] Error reading the forwarding rules %s:%s", err, resp))
	}
	d.Set(pdnsInstanceID, instanceID)
	d.Set(pdnsCRFRResolverID, resolverID)
	d.Set(pdnsCRFRRuleID, ruleID)
	d.Set(pdnsCRFRDesctiption, *result.Description)
	d.Set(pdnsCRFRType, *result.Type)
	d.Set(pdnsCRFRMatch, *result.Match)
	d.Set(pdnsCRFRForwardTo, result.ForwardTo)
	d.Set(pdnsCRFRViews, flattenPDNSFRViews(result.Views))
	return nil

}
func resourceIbmDnsCrForwardingRuleUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	dnsSvcsClient, err := meta.(conns.ClientSession).PrivateDNSClientSession()
	if err != nil {
		return diag.FromErr(err)
	}
	ruleID, resolverID, instanceID, err := flex.ConvertTfToCisThreeVar(d.Id())

	if err != nil {
		return diag.FromErr(err)
	}
	opt := dnsSvcsClient.NewUpdateForwardingRuleOptions(instanceID, resolverID, ruleID)

	if d.HasChange(pdnsCRFRDesctiption) ||
		d.HasChange(pdnsCRFRMatch) ||
		d.HasChange(pdnsCRFRForwardTo) ||
		d.HasChange(pdnsCRFRViews) {
		if des, ok := d.GetOk(pdnsCRFRDesctiption); ok {
			frdesc := des.(string)
			opt.SetDescription(frdesc)
		}
		if _, ok := d.GetOk(pdnsCRFRForwardTo); ok {
			opt.SetForwardTo(flex.ExpandStringList(d.Get(pdnsCRFRForwardTo).([]interface{})))
		}
		if ty, ok := d.GetOk(pdnsCRFRType); ok {
			crtype := ty.(string)
			if strings.ToLower(crtype) == "Default" {
				if match, ok := d.GetOk(pdnsCRFRMatch); ok {
					frmatch := match.(string)
					opt.SetMatch(frmatch)
				}
			}
		}
		if view, ok := d.GetOk(pdnsCRFRViews); ok {
			opt.SetViews(expandPDNSFRViews(view.(*schema.Set)))
		}
		result, resp, err := dnsSvcsClient.UpdateForwardingRuleWithContext(context, opt)
		if err != nil || result == nil {
			return diag.FromErr(fmt.Errorf("[ERROR] Error updating the forwarding rule %s:%s", err, resp))
		}

	}
	return resourceIbmDnsCrForwardingRuleRead(context, d, meta)
}

func resourceIbmDnsCrForwardingRuleDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	dnsSvcsClient, err := meta.(conns.ClientSession).PrivateDNSClientSession()
	if err != nil {
		return diag.FromErr(err)
	}
	ruleID, resolverID, instanceID, err := flex.ConvertTfToCisThreeVar(d.Id())
	opt := dnsSvcsClient.NewDeleteForwardingRuleOptions(instanceID, resolverID, ruleID)
	response, err := dnsSvcsClient.DeleteForwardingRuleWithContext(context, opt)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			return nil
		}
		return diag.FromErr(fmt.Errorf("[ERROR] Error deleting the  Forwarding Rules %s:%s", err, response))
	}
	d.SetId("")
	return nil
}

func expandPDNSFRViews(viewsList *schema.Set) (views []dns.ViewConfig) {
	for _, viewElem := range viewsList.List() {
		viewItem := viewElem.(map[string]interface{})
		view := dns.ViewConfig{
			Name:        core.StringPtr(viewItem[pdnsCRFRVName].(string)),
			Description: core.StringPtr(viewItem[pdnsCRFRVDescription].(string)),
			Expression:  core.StringPtr(viewItem[pdnsCRFRVExpression].(string)),
			ForwardTo:   flex.ExpandStringList(viewItem[pdnsCRFRVForwardTo].([]interface{})),
		}
		views = append(views, view)
	}
	return
}

func flattenPDNSFRViews(list []dns.ViewConfig) []map[string]interface{} {
	views := []map[string]interface{}{}
	for _, view := range list {
		l := map[string]interface{}{
			pdnsCRFRVName:        *view.Name,
			pdnsCRFRVExpression:  *view.Expression,
			pdnsCRFRVDescription: *view.Description,
			pdnsCRFRVForwardTo:   view.ForwardTo,
		}
		views = append(views, l)
	}
	return views
}
