// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceIbmDnsCrForwardingRules() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceIbmDnsCrForwardingRulesRead,

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
			pdnsCRForwardRule: {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						pdnsCRFRRuleID: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Identifier of the forwarding rule.",
						},
						pdnsCRFRDesctiption: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Descriptive text of the forwarding rule.",
						},
						pdnsCRFRType: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Type of the forwarding rule.",
						},
						pdnsCRFRMatch: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The matching zone or hostname.",
						},
						pdnsCRFRForwardTo: {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The upstream DNS servers will be forwarded to.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
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
				},
			},
		},
	}
}

func dataSourceIbmDnsCrForwardingRulesRead(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(ClientSession).PrivateDNSClientSessionScoped()
	if err != nil {
		return err
	}
	instanceID := d.Get(pdnsInstanceID).(string)
	resolverID := d.Get(pdnsCRFRResolverID).(string)

	opt := sess.NewListForwardingRulesOptions(instanceID, resolverID)

	result, response, err := sess.ListForwardingRules(opt)
	if err != nil {
		return fmt.Errorf("error reading list of available forword rules :%s\n%s", err, response)
	}

	forwardRules := make([]interface{}, 0)
	for _, instance := range result.ForwardingRules {
		forwardRule := map[string]interface{}{}
		forwardRule[pdnsCRFRRuleID] = *instance.ID
		forwardRule[pdnsCRFRDesctiption] = *instance.Description
		forwardRule[pdnsCRFRType] = *instance.Type
		forwardRule[pdnsCRFRMatch] = *instance.Match
		forwardRule[pdnsCRFRForwardTo] = instance.ForwardTo

		forwardRules = append(forwardRules, forwardRule)
	}
	d.SetId(dataSourceIBMPrivateDNSForwardrulesID(d))
	d.Set(pdnsInstanceID, instanceID)
	d.Set(pdnsCRFRResolverID, resolverID)
	d.Set(pdnsCRForwardRule, forwardRules)
	return nil
}

func dataSourceIBMPrivateDNSForwardrulesID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}
