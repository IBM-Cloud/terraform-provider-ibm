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
)

func DataSourceIBMIamIdp() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMIamIdpRead,

		Schema: map[string]*schema.Schema{
			"idp_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Unique identifier of the IDP.",
			},
			"account_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Account where the IdP resides in.",
			},
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Speaking name of the Identity Provider.",
			},
			"type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Type of the IDP.",
			},
			"active": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Defines if the IDP is active.",
			},
			"entity_tag": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Version of the IDP.",
			},
			"share_scope": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of targets which can consume the IdP.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "ID of the account or enterprise.",
						},
						"type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Type of share scope.",
						},
					},
				},
			},
			"created_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Timestamp when the IDP was created.",
			},
			"modified_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Timestamp when the IDP was last modified.",
			},
		},
	}
}

func dataSourceIBMIamIdpRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	iamIdentityClient, err := meta.(conns.ClientSession).IAMIdentityV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_idp", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	idpID := d.Get("idp_id").(string)
	getOpts := iamIdentityClient.NewGetIdpOptions(idpID)
	idp, _, err := iamIdentityClient.GetIdpWithContext(ctx, getOpts)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetIdpWithContext failed: %s", err.Error()), "ibm_iam_idp", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(idpID)

	if idp.AccountID != nil {
		d.Set("account_id", idp.AccountID)
	}
	if idp.Name != nil {
		d.Set("name", idp.Name)
	}
	if idp.Type != nil {
		d.Set("type", idp.Type)
	}
	if idp.Active != nil {
		d.Set("active", idp.Active)
	}
	if idp.EntityTag != nil {
		d.Set("entity_tag", idp.EntityTag)
	}
	if idp.CreatedAt != nil {
		d.Set("created_at", flex.DateTimeToString(idp.CreatedAt))
	}
	if idp.ModifiedAt != nil {
		d.Set("modified_at", flex.DateTimeToString(idp.ModifiedAt))
	}
	if idp.ShareScope != nil {
		d.Set("share_scope", flattenShareScope(idp.ShareScope))
	}
	return nil
}
