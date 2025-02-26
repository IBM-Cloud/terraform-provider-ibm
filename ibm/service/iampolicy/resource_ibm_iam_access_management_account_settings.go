// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package iampolicy

import (
	"context"
	"fmt"
	"encoding/json"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	// "github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/iampolicymanagementv1"
)

func ResourceIBMIAMAccessManagementAccountSettings() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMAccessManagementAccountSettingsUpdate,
		ReadContext:   resourceIBMAccessManagementAccountSettingsRead,
		UpdateContext: resourceIBMAccessManagementAccountSettingsUpdate,
		DeleteContext: resourceIBMAccessManagementAccountSettingsReset,
		Importer:      &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"external_account_identity_interaction": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "How external accounts can interact in relation to the requested account.",
				DiffSuppressFunc: flex.SuppressEquivalentJSON,
				StateFunc: func(v interface{}) string {
					json, _ := flex.NormalizeJSONString(v)
					return json
				},
				// Elem: &schema.Resource{
				// 	Schema: map[string]*schema.Schema{
				// 		"identity_types": {
				// 			Type:        schema.TypeSet,
				// 			Required:    true,
				// 			Description: "The settings for each identity type.",
				// 			Elem: &schema.Resource{
				// 				Schema: map[string]*schema.Schema{
				// 					"user": {
				// 						Type:        schema.TypeSet,
				// 						Optional:    true,
				// 						Description: "The core set of properties associated with a user identity type.",
				// 						Elem: &schema.Resource{
				// 							Schema: map[string]*schema.Schema{
				// 								"state": {
				// 									Type:        schema.TypeString,
				// 									Required:    true,
				// 									Description: "The state of the user identity type.",
				// 								},
				// 								"external_allowed_accounts": {
				// 									Type:        schema.TypeList,
				// 									Required:    true,
				// 									Description: "List of accounts that the state applies to for the user identity type.",
				// 									Elem:        &schema.Schema{Type: schema.TypeString},
				// 								},
				// 							},
				// 						},
				// 					},
				// 					"service_id": {
				// 						Type:        schema.TypeSet,
				// 						Optional:    true,
				// 						Description: "The core set of properties associated with a serviceID identity type.",
				// 						Elem: &schema.Resource{
				// 							Schema: map[string]*schema.Schema{
				// 								"state": {
				// 									Type:        schema.TypeString,
				// 									Required:    true,
				// 									Description: "The state of the serviceId identity type.",
				// 								},
				// 								"external_allowed_accounts": {
				// 									Type:        schema.TypeList,
				// 									Required:    true,
				// 									Description: "List of accounts that the state applies to for the serviceId identity type.",
				// 									Elem:        &schema.Schema{Type: schema.TypeString},
				// 								},
				// 							},
				// 						},
				// 					},
				// 					"service": {
				// 						Type:        schema.TypeSet,
				// 						Optional:    true,
				// 						Description: "The core set of properties associated with a service identity type.",
				// 						Elem: &schema.Resource{
				// 							Schema: map[string]*schema.Schema{
				// 								"state": {
				// 									Type:        schema.TypeString,
				// 									Required:    true,
				// 									Description: "The state of the service identity type.",
				// 								},
				// 								"external_allowed_accounts": {
				// 									Type:        schema.TypeList,
				// 									Required:    true,
				// 									Description: "List of accounts that the state applies to for the service identity type.",
				// 									Elem:        &schema.Schema{Type: schema.TypeString},
				// 								},
				// 							},
				// 						},
				// 					},
				// 				},
				// 			},
				// 		},
				// 	},
				// },
			},
			"accept_language": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "default",
				Description: "Language code for translations* `default` - English* `de` -  German (Standard)* `en` - English* `es` - Spanish (Spain)* `fr` - French (Standard)* `it` - Italian (Standard)* `ja` - Japanese* `ko` - Korean* `pt-br` - Portuguese (Brazil)* `zh-cn` - Chinese (Simplified, PRC)* `zh-tw` - (Chinese, Taiwan).",
			},
			"account_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The account GUID that the Access Management Account Settings belong to.",
			},
		},
	}
}

func resourceIBMAccessManagementAccountSettingsCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// There is no create,
	return resourceIBMAccessManagementAccountSettingsRead(context, d, meta)
}

func resourceIBMAccessManagementAccountSettingsRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	iamPolicyManagementClient, err := meta.(conns.ClientSession).IAMPolicyManagementV1API()
	if err != nil {
		tfErr := flex.TerraformErrorf(err, err.Error(), "ibm_iam_access_management_account_settings", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	var accountID string

	if _, ok := d.GetOk("account_id"); ok {
		accountID = d.Get("account_id").(string)
	}
	
	getSettingsOptions := &iampolicymanagementv1.GetSettingsOptions{
		AccountID:      &accountID,
	}


	if _, ok := d.GetOk("accept_language"); ok {
		getSettingsOptions.SetAcceptLanguage(d.Get("accept_language").(string))
	}

	accountSettingsAccessManagement, _, err := iamPolicyManagementClient.GetSettingsWithContext(context, getSettingsOptions)
	
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetSettingsWithContext failed: %s", err.Error()), "ibm_iam_access_management_account_settings", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	if _, ok := d.GetOk("external_account_identity_interaction"); ok {
		b, err := json.Marshal(accountSettingsAccessManagement.ExternalAccountIdentityInteraction)
		if err != nil {
			return diag.FromErr(fmt.Errorf("failed to Marshal accountSettingsAccessManagement.ExternalAccountIdentityInteraction %s", err))
		}
		jsonInt, _ := flex.NormalizeJSONString(string(b))
		d.Set("external_account_identity_interaction", jsonInt)
	}
	
	d.SetId(fmt.Sprintf("amSettings-%s", accountID))
	return nil
}

func resourceIBMAccessManagementAccountSettingsUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	iamPolicyManagementClient, err := meta.(conns.ClientSession).IAMPolicyManagementV1API()
	if err != nil {
		tfErr := flex.TerraformErrorf(err, err.Error(), "ibm_iam_access_management_account_settings", "update")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	updateSettingsOptions := &iampolicymanagementv1.UpdateSettingsOptions{}

	hasChange := false

	if d.HasChange("external_account_identity_interaction") {
		if interaction, ok := d.GetOk("external_account_identity_interaction"); ok {
			externalAccountIdentityInteractionPatch := &iampolicymanagementv1.ExternalAccountIdentityInteractionPatch{}
			json.Unmarshal([]byte(interaction.(string)), &externalAccountIdentityInteractionPatch)
			updateSettingsOptions.ExternalAccountIdentityInteraction = externalAccountIdentityInteractionPatch
		}
		hasChange = true
	}

	if hasChange {
		var accountID string
		if _, ok := d.GetOk("account_id"); ok {
			accountID = d.Get("account_id").(string)
		}
		
		getSettingsOptions := &iampolicymanagementv1.GetSettingsOptions{
			AccountID:      &accountID,
		}

		if _, ok := d.GetOk("accept_language"); ok {
			getSettingsOptions.SetAcceptLanguage(d.Get("accept_language").(string))
		}
		_, response, err := iamPolicyManagementClient.GetSettingsWithContext(context, getSettingsOptions)
		if err != nil {
			if response != nil && response.StatusCode == 404 {
				d.SetId("")
				return nil
			}
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetPolicyAssignmentWithContext failed: %s", err.Error()), "ibm_iam_access_management_account_settings", "read")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}

		updateSettingsOptions.SetIfMatch(response.Headers.Get("ETag"))
		updateSettingsOptions.SetAccountID(accountID)
		_, _, err = iamPolicyManagementClient.UpdateSettingsWithContext(context, updateSettingsOptions)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("UpdateSettingsWithContext failed: %s", err.Error()), "ibm_iam_access_management_account_settings", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
	}

	return resourceIBMAccessManagementAccountSettingsRead(context, d, meta)
}

func resourceIBMAccessManagementAccountSettingsReset(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
// 	iamPolicyManagementClient, err := meta.(conns.ClientSession).IAMPolicyManagementV1API()

// Update Settings to enabled and empty array for each category since there is no real delete functionality for this API

// 	d.SetId("")

	return nil
}