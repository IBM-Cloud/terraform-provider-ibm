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

// ResourceIBMIamIdpAccountSetting manages the binding of an IDP to an account
// (AddIDPSetting / GetIDPSetting / UpdateIDPSetting / RemoveIDPSetting).
func ResourceIBMIamIdpAccountSetting() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMIamIdpAccountSettingCreate,
		ReadContext:   resourceIBMIamIdpAccountSettingRead,
		UpdateContext: resourceIBMIamIdpAccountSettingUpdate,
		DeleteContext: resourceIBMIamIdpAccountSettingDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"account_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Account which is bound to the IDP.",
			},
			"idp_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Identity provider ID.",
			},
			"cloud_user_strategy": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Strategy for how Cloud User representatives for the IdP users are handled. Valid values: DYNAMIC, STATIC, NEVER.",
			},
			"active": {
				Type:        schema.TypeBool,
				Required:    true,
				Description: "Specifies if the IdP is enabled for usage in the given account context.",
			},
			"ui_default": {
				Type:        schema.TypeBool,
				Required:    true,
				Description: "Specifies if the IdP is used as default in the given account context.",
			},
			// Computed response fields
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
		},
	}
}

func resourceIBMIamIdpAccountSettingCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	iamIdentityClient, err := meta.(conns.ClientSession).IAMIdentityV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_idp_account_setting", "create", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	accountID := d.Get("account_id").(string)
	idpID := d.Get("idp_id").(string)
	cloudUserStrategy := d.Get("cloud_user_strategy").(string)
	active := d.Get("active").(bool)
	uiDefault := d.Get("ui_default").(bool)

	addOpts := iamIdentityClient.NewAddIDPSettingOptions(accountID, idpID, cloudUserStrategy, active, uiDefault)

	setting, _, err := iamIdentityClient.AddIDPSettingWithContext(ctx, addOpts)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("AddIDPSettingWithContext failed: %s", err.Error()), "ibm_iam_idp_account_setting", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(fmt.Sprintf("%s/%s", accountID, *setting.IdpID))
	return resourceIBMIamIdpAccountSettingRead(ctx, d, meta)
}

func resourceIBMIamIdpAccountSettingRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	iamIdentityClient, err := meta.(conns.ClientSession).IAMIdentityV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_idp_account_setting", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	parts, err := flex.SepIdParts(d.Id(), "/")
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_idp_account_setting", "read", "sep-id-parts").GetDiag()
	}
	accountID := parts[0]
	idpID := parts[1]

	getOpts := iamIdentityClient.NewGetIDPSettingOptions(accountID, idpID)
	setting, response, err := iamIdentityClient.GetIDPSettingWithContext(ctx, getOpts)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetIDPSettingWithContext failed: %s", err.Error()), "ibm_iam_idp_account_setting", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.Set("account_id", accountID)
	if setting.IdpID != nil {
		d.Set("idp_id", setting.IdpID)
	}
	if setting.CloudUserStrategy != nil {
		d.Set("cloud_user_strategy", setting.CloudUserStrategy)
	}
	if setting.Active != nil {
		d.Set("active", setting.Active)
	}
	if setting.UIDefault != nil {
		d.Set("ui_default", setting.UIDefault)
	}
	if setting.OwnerAccount != nil {
		d.Set("owner_account", setting.OwnerAccount)
	}
	if setting.OwnerAccountName != nil {
		d.Set("owner_account_name", setting.OwnerAccountName)
	}
	if setting.IdpName != nil {
		d.Set("idp_name", setting.IdpName)
	}
	if setting.IdpType != nil {
		d.Set("idp_type", setting.IdpType)
	}
	return nil
}

func resourceIBMIamIdpAccountSettingUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	iamIdentityClient, err := meta.(conns.ClientSession).IAMIdentityV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_idp_account_setting", "update", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	parts, err := flex.SepIdParts(d.Id(), "/")
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_idp_account_setting", "update", "sep-id-parts").GetDiag()
	}
	accountID := parts[0]
	idpID := parts[1]

	updateOpts := iamIdentityClient.NewUpdateIDPSettingOptions(accountID, idpID)

	if d.HasChange("cloud_user_strategy") {
		updateOpts.SetCloudUserStrategy(d.Get("cloud_user_strategy").(string))
	}
	if d.HasChange("active") {
		updateOpts.SetActive(d.Get("active").(bool))
	}
	if d.HasChange("ui_default") {
		updateOpts.SetUIDefault(d.Get("ui_default").(bool))
	}

	_, _, err = iamIdentityClient.UpdateIDPSettingWithContext(ctx, updateOpts)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("UpdateIDPSettingWithContext failed: %s", err.Error()), "ibm_iam_idp_account_setting", "update")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	return resourceIBMIamIdpAccountSettingRead(ctx, d, meta)
}

func resourceIBMIamIdpAccountSettingDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	iamIdentityClient, err := meta.(conns.ClientSession).IAMIdentityV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_idp_account_setting", "delete", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	parts, err := flex.SepIdParts(d.Id(), "/")
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_idp_account_setting", "delete", "sep-id-parts").GetDiag()
	}
	accountID := parts[0]
	idpID := parts[1]

	removeOpts := iamIdentityClient.NewRemoveIDPSettingOptions(accountID, idpID)
	_, err = iamIdentityClient.RemoveIDPSettingWithContext(ctx, removeOpts)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("RemoveIDPSettingWithContext failed: %s", err.Error()), "ibm_iam_idp_account_setting", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId("")
	return nil
}
