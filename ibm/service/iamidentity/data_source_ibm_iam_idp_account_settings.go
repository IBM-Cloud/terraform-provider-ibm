// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package iamidentity

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
)

// DataSourceIBMIamIdpAccountSettings lists IDP settings for an account
// using the ListIDPSettings API (consumable or consumed).
func DataSourceIBMIamIdpAccountSettings() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMIamIdpAccountSettingsRead,

		Schema: map[string]*schema.Schema{
			"account_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Account ID to list IDP settings for.",
			},
			"type": {
				Type:         schema.TypeString,
				Required:     true,
				Description:  "Type of IDP settings to list. Valid values: consumable, consumed.",
				ValidateFunc: validate.ValidateAllowedStringValues([]string{"consumable", "consumed"}),
			},
			"idps": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of IDP account settings.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"idp_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Identity provider ID.",
						},
						"owner_account": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Account that owns the IDP.",
						},
						"owner_account_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Name of the account that owns the IDP.",
						},
						"idp_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Name of the IDP.",
						},
						"idp_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Type of the IDP.",
						},
						"cloud_user_strategy": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Strategy for Cloud User representatives.",
						},
						"active": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether the IDP is active in this account.",
						},
						"ui_default": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether the IDP is the default in this account.",
						},
					},
				},
			},
		},
	}
}

func dataSourceIBMIamIdpAccountSettingsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	iamIdentityClient, err := meta.(conns.ClientSession).IAMIdentityV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_idp_account_settings", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	accountID := d.Get("account_id").(string)
	settingType := d.Get("type").(string)

	listOpts := iamIdentityClient.NewListIDPSettingsOptions(accountID, settingType)
	resp, _, err := iamIdentityClient.ListIDPSettingsWithContext(ctx, listOpts)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("ListIDPSettingsWithContext failed: %s", err.Error()), "ibm_iam_idp_account_settings", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(fmt.Sprintf("%s/%s", accountID, settingType))

	idpsList := make([]map[string]interface{}, 0, len(resp.Idps))
	for _, s := range resp.Idps {
		m := map[string]interface{}{}
		if s.IdpID != nil {
			m["idp_id"] = *s.IdpID
		}
		if s.OwnerAccount != nil {
			m["owner_account"] = *s.OwnerAccount
		}
		if s.OwnerAccountName != nil {
			m["owner_account_name"] = *s.OwnerAccountName
		}
		if s.IdpName != nil {
			m["idp_name"] = *s.IdpName
		}
		if s.IdpType != nil {
			m["idp_type"] = *s.IdpType
		}
		if s.CloudUserStrategy != nil {
			m["cloud_user_strategy"] = *s.CloudUserStrategy
		}
		if s.Active != nil {
			m["active"] = *s.Active
		}
		if s.UIDefault != nil {
			m["ui_default"] = *s.UIDefault
		}
		idpsList = append(idpsList, m)
	}
	d.Set("idps", idpsList)
	return nil
}
