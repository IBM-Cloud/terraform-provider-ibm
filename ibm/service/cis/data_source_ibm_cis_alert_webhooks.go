// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package cis

import (
	"fmt"
	"time"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const cisWebhooksList = "cis_webhooks_list"

func DataSourceIBMCISWebhooks() *schema.Resource {
	return &schema.Resource{
		Read: dataIBMCISWebhooksRead,
		Schema: map[string]*schema.Schema{
			cisID: {
				Type:        schema.TypeString,
				Description: "CIS instance crn",
				Required:    true,
			},
			cisWebhooksList: {
				Type:        schema.TypeList,
				Description: "Collection of Webhook details",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						cisWebhooksID: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Webhook ID",
						},
						cisWebhooksName: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Webhook Name",
						},
						cisWebhooksURL: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Webhook URL",
						},
						cisWebhooksType: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Webhook Type",
						},
					},
				},
			},
		},
	}
}
func dataIBMCISWebhooksRead(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(conns.ClientSession).CisWebhooksSession()
	if err != nil {
		return fmt.Errorf("[ERROR] Error while getting the CisWebhooksSession %s", err)
	}
	crn := d.Get(cisID).(string)
	sess.Crn = core.StringPtr(crn)
	opt := sess.NewListWebhooksOptions()

	result, resp, err := sess.ListWebhooks(opt)
	if err != nil || result == nil {
		return fmt.Errorf("[ERROR] Error Listing all Webhooks %q: %s %s", d.Id(), err, resp)
	}

	webhooks := make([]map[string]interface{}, 0)

	for _, instance := range result.Result {
		webhook := map[string]interface{}{}
		webhook[cisWebhooksID] = *instance.ID
		webhook[cisWebhooksName] = *instance.Name
		webhook[cisWebhooksURL] = *instance.URL
		webhook[cisWebhooksType] = *instance.Type
		webhooks = append(webhooks, webhook)
	}
	d.SetId(dataSourceCISFiltersCheckID(d))
	d.Set(cisID, crn)
	d.Set(cisWebhooksList, webhooks)
	return nil
}

func dataSourceCISWebhooksCheckID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}
