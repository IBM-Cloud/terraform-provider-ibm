// Copyright IBM Corp. 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package scc

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM/scc-go-sdk/v4/addonmanagerv1"
)

func DataSourceIBMSccAddonInsights() *schema.Resource {
	return &schema.Resource{
		ReadContext: DataSourceIBMSccAddonInsightsRead,

		Schema: map[string]*schema.Schema{
			"account_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func DataSourceIBMSccAddonInsightsRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	addonManagerClient, err := meta.(conns.ClientSession).AddonManagerV1()
	if err != nil {
		return diag.FromErr(err)
	}

	userDetails, err := meta.(conns.ClientSession).BluemixUserDetails()
	if err != nil {
		return diag.FromErr(err)
	}

	accountID := d.Get("account_id").(string)
	log.Println(fmt.Sprintf("[DEBUG] using specified AccountID %s", accountID))
	if accountID == "" {
		accountID = userDetails.UserAccount
		log.Println(fmt.Sprintf("[DEBUG] AccountID not spedified, using %s", accountID))
	}
	addonManagerClient.AccountID = &accountID

	getSupportedInsightsV2Options := &addonmanagerv1.GetSupportedInsightsV2Options{}

	allInsights, response, err := addonManagerClient.GetSupportedInsightsV2WithContext(context, getSupportedInsightsV2Options)
	if err != nil {
		log.Printf("[DEBUG] GetSupportedInsightsV2WithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("GetSupportedInsightsV2WithContext failed %s\n%s", err, response))
	}

	d.Set("type", allInsights.Type)

	d.SetId(DataSourceIBMSccAddonInsightsID(d))

	return nil
}

// DataSourceIBMSccAddonInsightsID returns a reasonable ID for the list.
func DataSourceIBMSccAddonInsightsID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}
