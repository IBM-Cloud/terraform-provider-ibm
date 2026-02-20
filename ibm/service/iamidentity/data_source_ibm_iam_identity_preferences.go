// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.107.1-41b0fbd0-20250825-080732
 */

package iamidentity

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/platform-services-go-sdk/iamidentityv1"
)

func DataSourceIBMIamIdentityPreferences() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMIamIdentityPreferencesRead,

		Schema: map[string]*schema.Schema{
			"account_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Account id to get preferences for.",
			},
			"iam_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "IAM id to get the preferences for.",
			},
			"preferences": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of Identity Preferences.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"service": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Service of the preference.",
						},
						"id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Unique ID of the preference.",
						},
						"account_id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Account ID of the preference, only present for scope 'account'.",
						},
						"scope": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Scope of the preference, 'global' or 'account'.",
						},
						"value_string": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "String value of the preference, only one value property is set, either 'value_string' or 'value_list_of_strings' is present.",
						},
						"value_list_of_strings": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "List of value of the preference, only one value property is set, either 'value_string' or 'value_list_of_strings' is present.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
		},
	}
}

func dataSourceIBMIamIdentityPreferencesRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	iamIdentityClient, err := meta.(conns.ClientSession).IAMIdentityV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_iam_identity_preferences", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getAllPreferencesOnScopeAccountOptions := &iamidentityv1.GetAllPreferencesOnScopeAccountOptions{}

	getAllPreferencesOnScopeAccountOptions.SetAccountID(d.Get("account_id").(string))
	getAllPreferencesOnScopeAccountOptions.SetIamID(d.Get("iam_id").(string))

	identityPreferencesResponse, _, err := iamIdentityClient.GetAllPreferencesOnScopeAccountWithContext(context, getAllPreferencesOnScopeAccountOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetAllPreferencesOnScopeAccountWithContext failed: %s", err.Error()), "(Data) ibm_iam_identity_preferences", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(dataSourceIBMIamIdentityPreferencesID(d))

	preferences := []map[string]interface{}{}
	for _, preferencesItem := range identityPreferencesResponse.Preferences {
		preferencesItemMap, err := DataSourceIBMIamIdentityPreferencesIdentityPreferenceResponseToMap(&preferencesItem) // #nosec G601
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_iam_identity_preferences", "read", "preferences-to-map").GetDiag()
		}
		preferences = append(preferences, preferencesItemMap)
	}
	if err = d.Set("preferences", preferences); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting preferences: %s", err), "(Data) ibm_iam_identity_preferences", "read", "set-preferences").GetDiag()
	}

	return nil
}

// dataSourceIBMIamIdentityPreferencesID returns a reasonable ID for the list.
func dataSourceIBMIamIdentityPreferencesID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}

func DataSourceIBMIamIdentityPreferencesIdentityPreferenceResponseToMap(model *iamidentityv1.IdentityPreferenceResponse) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Service != nil {
		modelMap["service"] = *model.Service
	}
	if model.ID != nil {
		modelMap["id"] = *model.ID
	}
	if model.AccountID != nil {
		modelMap["account_id"] = *model.AccountID
	}
	if model.Scope != nil {
		modelMap["scope"] = *model.Scope
	}
	if model.ValueString != nil {
		modelMap["value_string"] = *model.ValueString
	}
	if model.ValueListOfStrings != nil {
		modelMap["value_list_of_strings"] = model.ValueListOfStrings
	}
	return modelMap, nil
}
