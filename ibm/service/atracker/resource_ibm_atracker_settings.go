// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.101.0-62624c1e-20250225-192301
 */

package atracker

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/atrackerv2"
)

func ResourceIBMAtrackerSettings() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMAtrackerSettingsCreate,
		ReadContext:   resourceIBMAtrackerSettingsRead,
		UpdateContext: resourceIBMAtrackerSettingsUpdate,
		DeleteContext: resourceIBMAtrackerSettingsDelete,
		Importer:      &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"default_targets": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "The target ID List. In the event that no routing rule causes the event to be sent to a target, these targets will receive the event.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"permitted_target_regions": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "If present then only these regions may be used to define a target.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"metadata_region_primary": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validate.InvokeValidator("ibm_atracker_settings", "metadata_region_primary"),
				Description:  "To store all your meta data in a single region.",
			},
			"metadata_region_backup": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validate.InvokeValidator("ibm_atracker_settings", "metadata_region_backup"),
				Description:  "To store all your meta data in a backup region.",
			},
			"private_api_endpoint_only": &schema.Schema{
				Type:        schema.TypeBool,
				Required:    true,
				Description: "If you set this true then you cannot access api through public network.",
			},
			"api_version": &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "API version used for configuring IBM Cloud Activity Tracker Event Routing resources in the account.",
			},
			"message": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "An optional message containing information about the audit log locations.",
			},
		},
	}
}

func ResourceIBMAtrackerSettingsValidator() *validate.ResourceValidator {
	validateSchema := make([]validate.ValidateSchema, 0)
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 "metadata_region_primary",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Required:                   true,
			Regexp:                     `^[a-zA-Z0-9 -_]`,
			MinValueLength:             3,
			MaxValueLength:             256,
		},
		validate.ValidateSchema{
			Identifier:                 "metadata_region_backup",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Optional:                   true,
			Regexp:                     `^[a-zA-Z0-9 -_]`,
			MinValueLength:             3,
			MaxValueLength:             256,
		},
	)

	resourceValidator := validate.ResourceValidator{ResourceName: "ibm_atracker_settings", Schema: validateSchema}
	return &resourceValidator
}

func resourceIBMAtrackerSettingsCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	atrackerClient, err := meta.(conns.ClientSession).AtrackerV2()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_atracker_settings", "delete", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	putSettingsOptions := &atrackerv2.PutSettingsOptions{}

	putSettingsOptions.SetMetadataRegionPrimary(d.Get("metadata_region_primary").(string))
	putSettingsOptions.SetPrivateAPIEndpointOnly(d.Get("private_api_endpoint_only").(bool))
	if _, ok := d.GetOk("default_targets"); ok {
		var defaultTargets []string
		for _, v := range d.Get("default_targets").([]interface{}) {
			defaultTargetsItem := v.(string)
			defaultTargets = append(defaultTargets, defaultTargetsItem)
		}
		putSettingsOptions.SetDefaultTargets(defaultTargets)
	}
	if _, ok := d.GetOk("permitted_target_regions"); ok {
		var permittedTargetRegions []string
		for _, v := range d.Get("permitted_target_regions").([]interface{}) {
			permittedTargetRegionsItem := v.(string)
			permittedTargetRegions = append(permittedTargetRegions, permittedTargetRegionsItem)
		}
		putSettingsOptions.SetPermittedTargetRegions(permittedTargetRegions)
	}
	if _, ok := d.GetOk("metadata_region_backup"); ok {
		putSettingsOptions.SetMetadataRegionBackup(d.Get("metadata_region_backup").(string))
	}

	settings, _, err := atrackerClient.PutSettingsWithContext(context, putSettingsOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("PutSettingsWithContext failed: %s", err.Error()), "ibm_atracker_settings", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(*settings.MetadataRegionPrimary)

	return resourceIBMAtrackerSettingsRead(context, d, meta)
}

func resourceIBMAtrackerSettingsRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	atrackerClient, err := meta.(conns.ClientSession).AtrackerV2()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_atracker_settings", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getSettingsOptions := &atrackerv2.GetSettingsOptions{}

	settings, response, err := atrackerClient.GetSettingsWithContext(context, getSettingsOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetSettingsWithContext failed: %s", err.Error()), "ibm_atracker_settings", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	if !core.IsNil(settings.DefaultTargets) {
		if err = d.Set("default_targets", settings.DefaultTargets); err != nil {
			err = fmt.Errorf("Error setting default_targets: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_atracker_settings", "read", "set-default_targets").GetDiag()
		}
	}
	if !core.IsNil(settings.PermittedTargetRegions) {
		if err = d.Set("permitted_target_regions", settings.PermittedTargetRegions); err != nil {
			err = fmt.Errorf("Error setting permitted_target_regions: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_atracker_settings", "read", "set-permitted_target_regions").GetDiag()
		}
	}
	if err = d.Set("metadata_region_primary", settings.MetadataRegionPrimary); err != nil {
		err = fmt.Errorf("Error setting metadata_region_primary: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_atracker_settings", "read", "set-metadata_region_primary").GetDiag()
	}
	if !core.IsNil(settings.MetadataRegionBackup) {
		if err = d.Set("metadata_region_backup", settings.MetadataRegionBackup); err != nil {
			err = fmt.Errorf("Error setting metadata_region_backup: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_atracker_settings", "read", "set-metadata_region_backup").GetDiag()
		}
	}
	if err = d.Set("private_api_endpoint_only", settings.PrivateAPIEndpointOnly); err != nil {
		err = fmt.Errorf("Error setting private_api_endpoint_only: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_atracker_settings", "read", "set-private_api_endpoint_only").GetDiag()
	}
	if err = d.Set("api_version", flex.IntValue(settings.APIVersion)); err != nil {
		err = fmt.Errorf("Error setting api_version: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_atracker_settings", "read", "set-api_version").GetDiag()
	}
	if !core.IsNil(settings.Message) {
		if err = d.Set("message", settings.Message); err != nil {
			err = fmt.Errorf("Error setting message: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_atracker_settings", "read", "set-message").GetDiag()
		}
	}

	return nil
}

func resourceIBMAtrackerSettingsUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	atrackerClient, err := meta.(conns.ClientSession).AtrackerV2()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_atracker_settings", "delete", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	putSettingsOptions := &atrackerv2.PutSettingsOptions{}

	putSettingsOptions.SetMetadataRegionPrimary(d.Id())
	putSettingsOptions.SetMetadataRegionPrimary(d.Get("metadata_region_primary").(string))
	putSettingsOptions.SetPrivateAPIEndpointOnly(d.Get("private_api_endpoint_only").(bool))
	if _, ok := d.GetOk("default_targets"); ok {
		var defaultTargets []string
		for _, v := range d.Get("default_targets").([]interface{}) {
			defaultTargetsItem := v.(string)
			defaultTargets = append(defaultTargets, defaultTargetsItem)
		}
		putSettingsOptions.SetDefaultTargets(defaultTargets)
	}
	if _, ok := d.GetOk("permitted_target_regions"); ok {
		var permittedTargetRegions []string
		for _, v := range d.Get("permitted_target_regions").([]interface{}) {
			permittedTargetRegionsItem := v.(string)
			permittedTargetRegions = append(permittedTargetRegions, permittedTargetRegionsItem)
		}
		putSettingsOptions.SetPermittedTargetRegions(permittedTargetRegions)
	}
	if _, ok := d.GetOk("metadata_region_backup"); ok {
		putSettingsOptions.SetMetadataRegionBackup(d.Get("metadata_region_backup").(string))
	}

	_, _, err = atrackerClient.PutSettingsWithContext(context, putSettingsOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("PutSettingsWithContext failed: %s", err.Error()), "ibm_atracker_settings", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	return resourceIBMAtrackerSettingsRead(context, d, meta)
}

func resourceIBMAtrackerSettingsDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	atrackerClient, err := meta.(conns.ClientSession).AtrackerV2()
	if err != nil {
		return diag.FromErr(err)
	}

	// Retrieve old settings and put them for required fields.  Remove all other fields
	settings, getResponse, err := atrackerClient.GetSettingsWithContext(context, &atrackerv2.GetSettingsOptions{})
	if err != nil {
		log.Printf("[DEBUG] PutSettingsWithContext with GetSettingsWithContext failed %s\n%s", err, getResponse)
		return diag.FromErr(fmt.Errorf("GetSettingsWithContext failed %s\n%s", err, getResponse))
	}
	putSettingsOptions := &atrackerv2.PutSettingsOptions{}

	putSettingsOptions.MetadataRegionPrimary = settings.MetadataRegionPrimary
	putSettingsOptions.PrivateAPIEndpointOnly = settings.PrivateAPIEndpointOnly
	putSettingsOptions.PermittedTargetRegions = []string{}
	putSettingsOptions.DefaultTargets = []string{}

	_, response, err := atrackerClient.PutSettingsWithContext(context, putSettingsOptions)
	if err != nil {
		log.Printf("[DEBUG] PutSettingsWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("PutSettingsWithContext failed %s\n%s", err, response))
	}

	d.SetId("")

	return nil
}
