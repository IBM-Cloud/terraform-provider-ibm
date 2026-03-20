// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.111.0-1bfb72c2-20260206-185521
 */

package accountmanagement

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/platform-services-go-sdk/accountmanagementv4"
)

func DataSourceIbmAccount() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIbmAccountRead,

		Schema: map[string]*schema.Schema{
			"account_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The unique identifier of the account you want to retrieve.",
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"owner": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"owner_userid": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"owner_iamid": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"linked_softlayer_account": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"team_directory_enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
			"traits": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"eu_supported": &schema.Schema{
							Type:     schema.TypeBool,
							Computed: true,
						},
						"poc": &schema.Schema{
							Type:     schema.TypeBool,
							Computed: true,
						},
						"hippa": &schema.Schema{
							Type:     schema.TypeBool,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceIbmAccountRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	accountManagementClient, err := meta.(conns.ClientSession).AccountManagementV4()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_account_info", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getAccountOptions := &accountmanagementv4.GetAccountOptions{}

	getAccountOptions.SetAccountID(d.Get("account_id").(string))

	accountResponse, _, err := accountManagementClient.GetAccountWithContext(context, getAccountOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetAccountWithContext failed: %s", err.Error()), "(Data) ibm_account_info", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(*accountResponse.ID)

	if err = d.Set("name", accountResponse.Name); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting name: %s", err), "(Data) ibm_account_info", "read", "set-name").GetDiag()
	}

	if err = d.Set("owner", accountResponse.Owner); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting owner: %s", err), "(Data) ibm_account_info", "read", "set-owner").GetDiag()
	}

	if err = d.Set("owner_userid", accountResponse.OwnerUserid); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting owner_userid: %s", err), "(Data) ibm_account_info", "read", "set-owner_userid").GetDiag()
	}

	if err = d.Set("owner_iamid", accountResponse.OwnerIamid); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting owner_iamid: %s", err), "(Data) ibm_account_info", "read", "set-owner_iamid").GetDiag()
	}

	if err = d.Set("type", accountResponse.Type); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting type: %s", err), "(Data) ibm_account_info", "read", "set-type").GetDiag()
	}

	if err = d.Set("status", accountResponse.Status); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting status: %s", err), "(Data) ibm_account_info", "read", "set-status").GetDiag()
	}

	if err = d.Set("linked_softlayer_account", accountResponse.LinkedSoftlayerAccount); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting linked_softlayer_account: %s", err), "(Data) ibm_account_info", "read", "set-linked_softlayer_account").GetDiag()
	}

	if err = d.Set("team_directory_enabled", accountResponse.TeamDirectoryEnabled); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting team_directory_enabled: %s", err), "(Data) ibm_account_info", "read", "set-team_directory_enabled").GetDiag()
	}

	traits := []map[string]interface{}{}
	traitsMap, err := DataSourceIbmAccountAccountResponseTraitsToMap(accountResponse.Traits)
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_account_info", "read", "traits-to-map").GetDiag()
	}
	traits = append(traits, traitsMap)
	if err = d.Set("traits", traits); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting traits: %s", err), "(Data) ibm_account_info", "read", "set-traits").GetDiag()
	}

	return nil
}

func DataSourceIbmAccountAccountResponseTraitsToMap(model *accountmanagementv4.AccountResponseTraits) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["eu_supported"] = *model.EuSupported
	modelMap["poc"] = *model.Poc
	modelMap["hippa"] = *model.Hippa
	return modelMap, nil
}
