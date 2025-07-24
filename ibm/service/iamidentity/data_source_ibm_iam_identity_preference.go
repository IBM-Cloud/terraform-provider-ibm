// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.105.1-067d600b-20250616-154447
 */

package iamidentity

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/iamidentityv1"
)

func DataSourceIBMIamIdentityPreference() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMIamIdentityPreferenceRead,

		Schema: map[string]*schema.Schema{
			"account_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Account id to get preference for.",
			},
			"iam_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "IAM id to get the preference for.",
			},
			"service": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Service of the preference to be fetched.",
			},
			"preference_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Identifier of preference to be fetched.",
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
	}
}

func dataSourceIBMIamIdentityPreferenceRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	iamIdentityClient, err := meta.(conns.ClientSession).IamIdentityV1()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_iam_identity_preference", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getPreferenceOnScopeAccountOptions := &iamidentityv1.GetPreferenceOnScopeAccountOptions{}

	getPreferenceOnScopeAccountOptions.SetAccountID(d.Get("account_id").(string))
	getPreferenceOnScopeAccountOptions.SetIamID(d.Get("iam_id").(string))
	getPreferenceOnScopeAccountOptions.SetService(d.Get("service").(string))
	getPreferenceOnScopeAccountOptions.SetPreferenceID(d.Get("preference_id").(string))

	identityPreferenceResponse, _, err := iamIdentityClient.GetPreferenceOnScopeAccountWithContext(context, getPreferenceOnScopeAccountOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetPreferenceOnScopeAccountWithContext failed: %s", err.Error()), "(Data) ibm_iam_identity_preference", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(fmt.Sprintf("%s/%s/%s/%s", *getPreferenceOnScopeAccountOptions.AccountID, *getPreferenceOnScopeAccountOptions.IamID, *getPreferenceOnScopeAccountOptions.Service, *getPreferenceOnScopeAccountOptions.PreferenceID))

	if !core.IsNil(identityPreferenceResponse.Scope) {
		if err = d.Set("scope", identityPreferenceResponse.Scope); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting scope: %s", err), "(Data) ibm_iam_identity_preference", "read", "set-scope").GetDiag()
		}
	}

	if !core.IsNil(identityPreferenceResponse.ValueString) {
		if err = d.Set("value_string", identityPreferenceResponse.ValueString); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting value_string: %s", err), "(Data) ibm_iam_identity_preference", "read", "set-value_string").GetDiag()
		}
	}

	if !core.IsNil(identityPreferenceResponse.ValueListOfStrings) {
		valueListOfStrings := []interface{}{}
		for _, valueListOfStringsItem := range identityPreferenceResponse.ValueListOfStrings {
			valueListOfStrings = append(valueListOfStrings, valueListOfStringsItem)
		}
		if err = d.Set("value_list_of_strings", valueListOfStrings); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting value_list_of_strings: %s", err), "(Data) ibm_iam_identity_preference", "read", "set-value_list_of_strings").GetDiag()
		}
	}

	return nil
}
