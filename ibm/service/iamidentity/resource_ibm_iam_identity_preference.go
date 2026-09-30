// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.106.0-09823488-20250707-071701
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

func ResourceIBMIamIdentityPreference() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMIamIdentityPreferenceCreate,
		ReadContext:   resourceIBMIamIdentityPreferenceRead,
		UpdateContext: resourceIBMIamIdentityPreferenceUpdate,
		DeleteContext: resourceIBMIamIdentityPreferenceDelete,
		Importer:      &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"account_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Account id to update preference for.",
			},
			"iam_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "IAM id to update the preference for.",
			},
			"service": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Service of the preference to be updated.",
			},
			"preference_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Identifier of preference to be updated.",
			},
			"value_string": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "String value of the preference, only one value property is set, either 'value_string' or 'value_list_of_strings' is present.",
			},
			"value_list_of_strings": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "List of values of the preference, only one value property is set, either 'value_string' or 'value_list_of_strings' is present.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"scope": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Scope of the preference, 'global' or 'account'.",
			},
		},
	}
}

func resourceIBMIamIdentityPreferenceCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// Identity preferences always exist with a default state and cannot be created.
	// We mask this from Terraform users by constructing the ID and calling the update function.
	d.SetId(fmt.Sprintf("%s/%s/%s/%s",
		d.Get("account_id").(string),
		d.Get("iam_id").(string),
		d.Get("service").(string),
		d.Get("preference_id").(string),
	))

	return resourceIBMIamIdentityPreferenceUpdate(context, d, meta)
}

func resourceIBMIamIdentityPreferenceRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	iamIdentityClient, err := meta.(conns.ClientSession).IAMIdentityV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_identity_preference", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getPreferencesOnScopeAccountOptions := &iamidentityv1.GetPreferencesOnScopeAccountOptions{}

	parts, err := flex.SepIdParts(d.Id(), "/")
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_identity_preference", "read", "sep-id-parts").GetDiag()
	}

	getPreferencesOnScopeAccountOptions.SetAccountID(parts[0])
	getPreferencesOnScopeAccountOptions.SetIamID(parts[1])
	getPreferencesOnScopeAccountOptions.SetService(parts[2])
	getPreferencesOnScopeAccountOptions.SetPreferenceID(parts[3])

	identityPreferenceResponse, response, err := iamIdentityClient.GetPreferencesOnScopeAccountWithContext(context, getPreferencesOnScopeAccountOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetPreferenceOnScopeAccountWithContext failed: %s", err.Error()), "ibm_iam_identity_preference", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	if err = d.Set("account_id", identityPreferenceResponse.AccountID); err != nil {
		err = fmt.Errorf("Error setting account_id: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_identity_preference", "read", "set-account_id").GetDiag()
	}
	// iam_id is not part of the response, so extract the value from the ID parts
	if err = d.Set("iam_id", parts[1]); err != nil {
		err = fmt.Errorf("Error setting iam_id: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_identity_preference", "read", "set-iam_id").GetDiag()
	}
	if err = d.Set("service", identityPreferenceResponse.Service); err != nil {
		err = fmt.Errorf("Error setting service: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_identity_preference", "read", "set-service").GetDiag()
	}
	if err = d.Set("value_string", identityPreferenceResponse.ValueString); err != nil {
		err = fmt.Errorf("Error setting value_string: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_identity_preference", "read", "set-value_string").GetDiag()
	}
	if !core.IsNil(identityPreferenceResponse.ValueListOfStrings) {
		if err = d.Set("value_list_of_strings", identityPreferenceResponse.ValueListOfStrings); err != nil {
			err = fmt.Errorf("Error setting value_list_of_strings: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_identity_preference", "read", "set-value_list_of_strings").GetDiag()
		}
	}
	if !core.IsNil(identityPreferenceResponse.Scope) {
		if err = d.Set("scope", identityPreferenceResponse.Scope); err != nil {
			err = fmt.Errorf("Error setting scope: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_identity_preference", "read", "set-scope").GetDiag()
		}
	}
	if !core.IsNil(identityPreferenceResponse.ID) {
		if err = d.Set("preference_id", identityPreferenceResponse.ID); err != nil {
			err = fmt.Errorf("Error setting preference_id: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_identity_preference", "read", "set-preference_id").GetDiag()
		}
	}

	return nil
}

func resourceIBMIamIdentityPreferenceUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	iamIdentityClient, err := meta.(conns.ClientSession).IAMIdentityV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_identity_preference", "update", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	updatePreferenceOnScopeAccountOptions := &iamidentityv1.UpdatePreferenceOnScopeAccountOptions{}

	parts, err := flex.SepIdParts(d.Id(), "/")
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_identity_preference", "update", "sep-id-parts").GetDiag()
	}

	// Lock to prevent concurrent updates to the same account/iam_id/service combination
	mutexKey := fmt.Sprintf("iam_identity_preference_%s_%s_%s", parts[0], parts[1], parts[2])
	conns.IbmMutexKV.Lock(mutexKey)
	defer conns.IbmMutexKV.Unlock(mutexKey)

	updatePreferenceOnScopeAccountOptions.SetAccountID(parts[0])
	updatePreferenceOnScopeAccountOptions.SetIamID(parts[1])
	updatePreferenceOnScopeAccountOptions.SetService(parts[2])
	updatePreferenceOnScopeAccountOptions.SetPreferenceID(parts[3])
	if _, ok := d.GetOk("value_string"); ok {
		updatePreferenceOnScopeAccountOptions.SetValueString(d.Get("value_string").(string))
	}
	if _, ok := d.GetOk("value_list_of_strings"); ok {
		var valueListOfStrings []string
		for _, v := range d.Get("value_list_of_strings").([]interface{}) {
			valueListOfStringsItem := v.(string)
			valueListOfStrings = append(valueListOfStrings, valueListOfStringsItem)
		}
		updatePreferenceOnScopeAccountOptions.SetValueListOfStrings(valueListOfStrings)
	}

	_, _, err = iamIdentityClient.UpdatePreferenceOnScopeAccountWithContext(context, updatePreferenceOnScopeAccountOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("UpdatePreferenceOnScopeAccountWithContext failed: %s", err.Error()), "ibm_iam_identity_preference", "update")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	return resourceIBMIamIdentityPreferenceRead(context, d, meta)
}

func resourceIBMIamIdentityPreferenceDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	iamIdentityClient, err := meta.(conns.ClientSession).IAMIdentityV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_identity_preference", "delete", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	deletePreferenceOnScopeAccountOptions := &iamidentityv1.DeletePreferencesOnScopeAccountOptions{}

	parts, err := flex.SepIdParts(d.Id(), "/")
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_identity_preference", "delete", "sep-id-parts").GetDiag()
	}

	// Lock to prevent concurrent deletes to the same account/iam_id/service combination
	mutexKey := fmt.Sprintf("iam_identity_preference_%s_%s_%s", parts[0], parts[1], parts[2])
	conns.IbmMutexKV.Lock(mutexKey)
	defer conns.IbmMutexKV.Unlock(mutexKey)

	deletePreferenceOnScopeAccountOptions.SetAccountID(parts[0])
	deletePreferenceOnScopeAccountOptions.SetIamID(parts[1])
	deletePreferenceOnScopeAccountOptions.SetService(parts[2])
	deletePreferenceOnScopeAccountOptions.SetPreferenceID(parts[3])

	_, err = iamIdentityClient.DeletePreferencesOnScopeAccount(deletePreferenceOnScopeAccountOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("DeletePreferenceOnScopeAccountWithContext failed: %s", err.Error()), "ibm_iam_identity_preference", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId("")

	return nil
}
