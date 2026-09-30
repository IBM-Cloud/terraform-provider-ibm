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

func DataSourceIBMIamIdps() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMIamIdpsRead,

		Schema: map[string]*schema.Schema{
			"account_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Account ID to query IDPs for.",
			},
			"idps": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of Identity Providers.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"idp_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Unique identifier of the IDP.",
						},
						"account_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Account where the IdP resides.",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Name of the Identity Provider.",
						},
						"type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Type of the IDP.",
						},
						"active": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether the IDP is active.",
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
				},
			},
		},
	}
}

func dataSourceIBMIamIdpsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	iamIdentityClient, err := meta.(conns.ClientSession).IAMIdentityV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_idps", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	accountID := d.Get("account_id").(string)
	listOpts := iamIdentityClient.NewListIdpsOptions(accountID)
	idpsResponse, _, err := iamIdentityClient.ListIdpsWithContext(ctx, listOpts)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("ListIdpsWithContext failed: %s", err.Error()), "ibm_iam_idps", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(accountID)

	idpsList := make([]map[string]interface{}, 0, len(idpsResponse.Idps))
	for _, idp := range idpsResponse.Idps {
		m := map[string]interface{}{}
		if idp.IdpID != nil {
			m["idp_id"] = *idp.IdpID
		}
		if idp.AccountID != nil {
			m["account_id"] = *idp.AccountID
		}
		if idp.Name != nil {
			m["name"] = *idp.Name
		}
		if idp.Type != nil {
			m["type"] = *idp.Type
		}
		if idp.Active != nil {
			m["active"] = *idp.Active
		}
		if idp.EntityTag != nil {
			m["entity_tag"] = *idp.EntityTag
		}
		if idp.ShareScope != nil {
			m["share_scope"] = flattenShareScope(idp.ShareScope)
		}
		if idp.CreatedAt != nil {
			m["created_at"] = flex.DateTimeToString(idp.CreatedAt)
		}
		if idp.ModifiedAt != nil {
			m["modified_at"] = flex.DateTimeToString(idp.ModifiedAt)
		}
		idpsList = append(idpsList, m)
	}
	d.Set("idps", idpsList)
	return nil
}
