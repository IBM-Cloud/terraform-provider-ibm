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
				Required:    true,
				Description: "String value of the preference, only one value property is set, either 'value_string' or 'value_list_of_strings' is present.",
			},
			"value_list_of_strings": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "List of value of the preference, only one value property is set, either 'value_string' or 'value_list_of_strings' is present.",
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

func resourceIBMIamIdentityPreferenceRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	iamIdentityClient, err := meta.(conns.ClientSession).IamIdentityV1()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_identity_preference", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getPreferenceOnScopeAccountOptions := &iamidentityv1.GetPreferenceOnScopeAccountOptions{}

	parts, err := flex.SepIdParts(d.Id(), "/")
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_identity_preference", "read", "sep-id-parts").GetDiag()
	}

	getPreferenceOnScopeAccountOptions.SetAccountID(parts[0])
	getPreferenceOnScopeAccountOptions.SetIamID(parts[1])
	getPreferenceOnScopeAccountOptions.SetService(parts[2])
	getPreferenceOnScopeAccountOptions.SetPreferenceID(parts[3])
	getPreferenceOnScopeAccountOptions.SetPreferenceID(parts[4])

	identityPreferenceResponse, response, err := iamIdentityClient.GetPreferenceOnScopeAccountWithContext(context, getPreferenceOnScopeAccountOptions)
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
	iamIdentityClient, err := meta.(conns.ClientSession).IamIdentityV1()
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

	updatePreferenceOnScopeAccountOptions.SetAccountID(parts[0])
	updatePreferenceOnScopeAccountOptions.SetIamID(parts[1])
	updatePreferenceOnScopeAccountOptions.SetService(parts[2])
	updatePreferenceOnScopeAccountOptions.SetPreferenceID(parts[3])
	updatePreferenceOnScopeAccountOptions.SetPreferenceID(parts[4])
	updatePreferenceOnScopeAccountOptions.SetAccountID(d.Get("account_id").(string))
	updatePreferenceOnScopeAccountOptions.SetIamID(d.Get("iam_id").(string))
	updatePreferenceOnScopeAccountOptions.SetService(d.Get("service").(string))
	updatePreferenceOnScopeAccountOptions.SetPreferenceID(d.Get("preference_id").(string))
	updatePreferenceOnScopeAccountOptions.SetValueString(d.Get("value_string").(string))
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
	iamIdentityClient, err := meta.(conns.ClientSession).IamIdentityV1()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_identity_preference", "delete", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	deletePreferenceOnScopeAccountOptions := &iamidentityv1.DeletePreferenceOnScopeAccountOptions{}

	parts, err := flex.SepIdParts(d.Id(), "/")
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_identity_preference", "delete", "sep-id-parts").GetDiag()
	}

	deletePreferenceOnScopeAccountOptions.SetAccountID(parts[0])
	deletePreferenceOnScopeAccountOptions.SetIamID(parts[1])
	deletePreferenceOnScopeAccountOptions.SetService(parts[2])
	deletePreferenceOnScopeAccountOptions.SetPreferenceID(parts[3])
	deletePreferenceOnScopeAccountOptions.SetPreferenceID(parts[4])

	_, _, err = iamIdentityClient.DeletePreferenceOnScopeAccountWithContext(context, deletePreferenceOnScopeAccountOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("DeletePreferenceOnScopeAccountWithContext failed: %s", err.Error()), "ibm_iam_identity_preference", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId("")

	return nil
}
