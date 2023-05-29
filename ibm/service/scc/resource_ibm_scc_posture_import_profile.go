// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package scc

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/scc-go-sdk/v4/posturemanagementv2"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceIBMSccPostureProfileImport() *schema.Resource {
	return &schema.Resource{
		CreateContext:      resourceIBMSccPostureProfileImport,
		ReadContext:        resourceIBMSccPostureProfileImportRead,
		UpdateContext:      resourceIBMSccPostureProfileImportRead,
		DeleteContext:      resourceIBMSccPostureProfileImportDelete,
		Importer:           &schema.ResourceImporter{},
		DeprecationMessage: "**Removal Notification** Resource Removal: Resource ibm_scc_posture_profile_import is deprecated and being removed.\n This resource will not be available from future release (v1.54.0).",
		Schema: map[string]*schema.Schema{
			"file": {
				Type:        schema.TypeString,
				Description: "File to import",
				Required:    true,
				ForceNew:    true,
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The name of the profile.",
			},
			"description": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "A description of the profile.",
			},
			"version": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The version of the profile.",
			},
			"created_by": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The user who created the profile.",
			},
			"modified_by": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The user who last modified the profile.",
			},
			"base_profile": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The base profile that the controls are pulled from.",
			},
			"type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The type of profile.",
			},
			"no_of_controls": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "no of Controls.",
			},
			"created_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The time that the profile was created in UTC.",
			},
			"updated_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The time that the profile was most recently modified in UTC.",
			},
			"enabled": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "The profile status. If the profile is enabled, the value is true. If the profile is disabled, the value is false.",
			},
		},
	}
}

func resourceIBMSccPostureProfileImport(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	postureManagementClient, err := meta.(conns.ClientSession).PostureManagementV2()
	if err != nil {
		return diag.FromErr(err)
	}

	userDetails, err := meta.(conns.ClientSession).BluemixUserDetails()
	if err != nil {
		return diag.FromErr(fmt.Errorf("[ERROR] Error getting userDetails %s", err))
	}

	importProfilesOptions := &posturemanagementv2.ImportProfilesOptions{}
	accountID := userDetails.UserAccount
	importProfilesOptions.SetAccountID(accountID)

	f, err := os.Open(d.Get("file").(string))
	if err != nil {
		log.Printf("[DEBUG] ImportProfilesWithContext failed to read file %s", err)
		return diag.FromErr(fmt.Errorf("ImportProfilesWithContext failed to read file %s", err))
	}
	importProfilesOptions.SetFile(f)

	profile, response, err := postureManagementClient.ImportProfilesWithContext(context, importProfilesOptions)
	if err != nil {
		log.Printf("[DEBUG] ImportProfilesWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("ImportProfilesWithContext failed %s\n%s", err, response))
	}

	d.SetId(*profile.ProfileID)

	return resourceIBMSccPostureProfileImportRead(context, d, meta)
}

func resourceIBMSccPostureProfileImportRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	postureManagementClient, err := meta.(conns.ClientSession).PostureManagementV2()
	if err != nil {
		return diag.FromErr(err)
	}

	getProfileOptions := &posturemanagementv2.GetProfileOptions{}
	userDetails, err := meta.(conns.ClientSession).BluemixUserDetails()
	if err != nil {
		return diag.FromErr(fmt.Errorf("[ERROR] Error getting userDetails %s", err))
	}

	accountID := userDetails.UserAccount
	getProfileOptions.SetAccountID(accountID)

	getProfileOptions.SetID(d.Id())
	getProfileOptions.SetProfileType("custom")

	profile, response, err := postureManagementClient.GetProfileWithContext(context, getProfileOptions)
	if err != nil {
		log.Printf("[DEBUG] GetProfileWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("GetProfileWithContext failed %s\n%s", err, response))
	}

	d.SetId(*profile.ID)
	if err = d.Set("name", profile.Name); err != nil {
		return diag.FromErr(fmt.Errorf("[ERROR] Error setting name: %s", err))
	}
	if err = d.Set("description", profile.Description); err != nil {
		return diag.FromErr(fmt.Errorf("[ERROR] Error setting description: %s", err))
	}
	if err = d.Set("version", flex.IntValue(profile.Version)); err != nil {
		return diag.FromErr(fmt.Errorf("[ERROR] Error setting version: %s", err))
	}
	if err = d.Set("created_by", profile.CreatedBy); err != nil {
		return diag.FromErr(fmt.Errorf("[ERROR] Error setting created_by: %s", err))
	}
	if err = d.Set("modified_by", profile.ModifiedBy); err != nil {
		return diag.FromErr(fmt.Errorf("[ERROR] Error setting modified_by: %s", err))
	}
	if err = d.Set("base_profile", profile.BaseProfile); err != nil {
		return diag.FromErr(fmt.Errorf("[ERROR] Error setting base_profile: %s", err))
	}
	if err = d.Set("type", profile.Type); err != nil {
		return diag.FromErr(fmt.Errorf("[ERROR] Error setting type: %s", err))
	}
	if err = d.Set("no_of_controls", flex.IntValue(profile.NoOfControls)); err != nil {
		return diag.FromErr(fmt.Errorf("[ERROR] Error setting no_of_controls: %s", err))
	}
	if err = d.Set("created_at", flex.DateTimeToString(profile.CreatedAt)); err != nil {
		return diag.FromErr(fmt.Errorf("[ERROR] Error setting created_at: %s", err))
	}
	if err = d.Set("updated_at", flex.DateTimeToString(profile.UpdatedAt)); err != nil {
		return diag.FromErr(fmt.Errorf("[ERROR] Error setting updated_at: %s", err))
	}
	if err = d.Set("enabled", profile.Enabled); err != nil {
		return diag.FromErr(fmt.Errorf("[ERROR] Error setting enabled: %s", err))
	}

	return nil
}

func resourceIBMSccPostureProfileImportDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	d.SetId("")
	return nil
}
