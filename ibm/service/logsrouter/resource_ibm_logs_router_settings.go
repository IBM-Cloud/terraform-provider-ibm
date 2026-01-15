// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.108.0-56772134-20251111-102802
 */

package logsrouter

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
	"github.com/IBM/platform-services-go-sdk/logsrouterv3"
)

func ResourceIBMLogsRouterSettings() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMLogsRouterSettingsCreate,
		ReadContext:   resourceIBMLogsRouterSettingsRead,
		UpdateContext: resourceIBMLogsRouterSettingsUpdate,
		DeleteContext: resourceIBMLogsRouterSettingsDelete,
		Importer:      &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"default_targets": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "A list of default target references. Enterprise-managed targets are not supported.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:        schema.TypeString,
							Required:    true,
							Description: "The target uuid for a pre-defined platform logs router target.",
						},
					},
				},
			},
			"permitted_target_regions": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "If present then only these regions may be used to define a target.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"primary_metadata_region": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validate.InvokeValidator("ibm_logs_router_v3_settings", "primary_metadata_region"),
				Description:  "To store all your meta data in a single region.",
			},
			"backup_metadata_region": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validate.InvokeValidator("ibm_logs_router_v3_settings", "backup_metadata_region"),
				Description:  "To backup all your meta data in a different region.",
			},
			"private_api_endpoint_only": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "If you set this true then you cannot access api through public network.",
			},
			"api_version": &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "API version used for IBM Cloud Logs Routing service under the account.",
			},
		},
	}
}

func ResourceIBMLogsRouterSettingsValidator() *validate.ResourceValidator {
	validateSchema := make([]validate.ValidateSchema, 0)
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 "primary_metadata_region",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Optional:                   true,
			Regexp:                     `^[a-zA-Z0-9 \-_]+$`,
			MinValueLength:             3,
			MaxValueLength:             256,
		},
		validate.ValidateSchema{
			Identifier:                 "backup_metadata_region",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Optional:                   true,
			Regexp:                     `^[a-zA-Z0-9 \-_]+$`,
			MinValueLength:             3,
			MaxValueLength:             256,
		},
	)

	resourceValidator := validate.ResourceValidator{ResourceName: "ibm_logs_router_v3_settings", Schema: validateSchema}
	return &resourceValidator
}

func resourceIBMLogsRouterSettingsCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logsRouterClient, err := meta.(conns.ClientSession).LogsRouterV3()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_logs_router_v3_settings", "delete", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	updateSettingsOptions := &logsrouterv3.UpdateSettingsOptions{}

	if _, ok := d.GetOk("default_targets"); ok {
		var defaultTargets []logsrouterv3.TargetIdentity
		for _, v := range d.Get("default_targets").([]interface{}) {
			value := v.(map[string]interface{})
			defaultTargetsItem, err := ResourceIBMLogsRouterSettingsMapToTargetIdentity(value)
			if err != nil {
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_logs_router_v3_settings", "delete", "parse-default_targets").GetDiag()
			}
			defaultTargets = append(defaultTargets, *defaultTargetsItem)
		}
		updateSettingsOptions.SetDefaultTargets(defaultTargets)
	}
	if _, ok := d.GetOk("permitted_target_regions"); ok {
		var permittedTargetRegions []string
		for _, v := range d.Get("permitted_target_regions").([]interface{}) {
			permittedTargetRegionsItem := v.(string)
			permittedTargetRegions = append(permittedTargetRegions, permittedTargetRegionsItem)
		}
		updateSettingsOptions.SetPermittedTargetRegions(permittedTargetRegions)
	}
	if _, ok := d.GetOk("primary_metadata_region"); ok {
		updateSettingsOptions.SetPrimaryMetadataRegion(d.Get("primary_metadata_region").(string))
	}
	if _, ok := d.GetOk("backup_metadata_region"); ok {
		updateSettingsOptions.SetBackupMetadataRegion(d.Get("backup_metadata_region").(string))
	}
	if _, ok := d.GetOk("private_api_endpoint_only"); ok {
		updateSettingsOptions.SetPrivateAPIEndpointOnly(d.Get("private_api_endpoint_only").(bool))
	}

	setting, _, err := logsRouterClient.UpdateSettingsWithContext(context, updateSettingsOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("UpdateSettingsWithContext failed: %s", err.Error()), "ibm_logs_router_v3_settings", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(*setting.PrimaryMetadataRegion)

	return resourceIBMLogsRouterSettingsRead(context, d, meta)
}

func resourceIBMLogsRouterSettingsRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logsRouterClient, err := meta.(conns.ClientSession).LogsRouterV3()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_logs_router_v3_settings", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getSettingsOptions := &logsrouterv3.GetSettingsOptions{}

	setting, response, err := logsRouterClient.GetSettingsWithContext(context, getSettingsOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetSettingsWithContext failed: %s", err.Error()), "ibm_logs_router_v3_settings", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	if !core.IsNil(setting.DefaultTargets) {
		defaultTargets := []map[string]interface{}{}
		for _, defaultTargetsItem := range setting.DefaultTargets {
			defaultTargetsItemMap, err := ResourceIBMLogsRouterSettingsTargetReferenceToMap(&defaultTargetsItem) // #nosec G601
			if err != nil {
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_logs_router_v3_settings", "read", "default_targets-to-map").GetDiag()
			}
			defaultTargets = append(defaultTargets, defaultTargetsItemMap)
		}
		if err = d.Set("default_targets", defaultTargets); err != nil {
			err = fmt.Errorf("Error setting default_targets: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_logs_router_v3_settings", "read", "set-default_targets").GetDiag()
		}
	}
	if !core.IsNil(setting.PermittedTargetRegions) {
		if err = d.Set("permitted_target_regions", setting.PermittedTargetRegions); err != nil {
			err = fmt.Errorf("Error setting permitted_target_regions: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_logs_router_v3_settings", "read", "set-permitted_target_regions").GetDiag()
		}
	}
	if !core.IsNil(setting.PrimaryMetadataRegion) {
		if err = d.Set("primary_metadata_region", setting.PrimaryMetadataRegion); err != nil {
			err = fmt.Errorf("Error setting primary_metadata_region: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_logs_router_v3_settings", "read", "set-primary_metadata_region").GetDiag()
		}
	}
	if !core.IsNil(setting.BackupMetadataRegion) {
		if err = d.Set("backup_metadata_region", setting.BackupMetadataRegion); err != nil {
			err = fmt.Errorf("Error setting backup_metadata_region: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_logs_router_v3_settings", "read", "set-backup_metadata_region").GetDiag()
		}
	}
	if !core.IsNil(setting.PrivateAPIEndpointOnly) {
		if err = d.Set("private_api_endpoint_only", setting.PrivateAPIEndpointOnly); err != nil {
			err = fmt.Errorf("Error setting private_api_endpoint_only: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_logs_router_v3_settings", "read", "set-private_api_endpoint_only").GetDiag()
		}
	}
	if err = d.Set("api_version", flex.IntValue(setting.APIVersion)); err != nil {
		err = fmt.Errorf("Error setting api_version: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_logs_router_v3_settings", "read", "set-api_version").GetDiag()
	}

	return nil
}

func resourceIBMLogsRouterSettingsUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logsRouterClient, err := meta.(conns.ClientSession).LogsRouterV3()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_logs_router_v3_settings", "delete", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	updateSettingsOptions := &logsrouterv3.UpdateSettingsOptions{}

	updateSettingsOptions.SetPrimaryMetadataRegion(d.Id())

	hasChange := false

	if d.HasChange("default_targets") {
		if _, ok := d.GetOk("default_targets"); ok {
			var defaultTargets []logsrouterv3.TargetIdentity
			for _, e := range d.Get("default_targets").([]interface{}) {
				value := e.(map[string]interface{})
				defaultTargetsItem, err := ResourceIBMLogsRouterSettingsMapToTargetIdentity(value)
				if err != nil {
					return diag.FromErr(err)
				}
				defaultTargets = append(defaultTargets, *defaultTargetsItem)
			}
			updateSettingsOptions.SetDefaultTargets(defaultTargets)
		} else {
			// In this case, need to remove all the default_targets
			updateSettingsOptions.SetDefaultTargets([]logsrouterv3.TargetIdentity{})
		}
		hasChange = true
	}
	if d.HasChange("permitted_target_regions") {
		if _, ok := d.GetOk("permitted_target_regions"); ok {
			var permittedTargetRegions []string
			for _, v := range d.Get("permitted_target_regions").([]interface{}) {
				permittedTargetRegionsItem := v.(string)
				permittedTargetRegions = append(permittedTargetRegions, permittedTargetRegionsItem)
			}
			updateSettingsOptions.SetPermittedTargetRegions(permittedTargetRegions)
		} else {
			// In this case, need to remove all the permitted_target_regions
			updateSettingsOptions.SetPermittedTargetRegions([]string{})
		}
		hasChange = true
	}
	if d.HasChange("primary_metadata_region") {
		updateSettingsOptions.SetPrimaryMetadataRegion(d.Get("primary_metadata_region").(string))
		hasChange = true
	}
	if d.HasChange("backup_metadata_region") {
		updateSettingsOptions.SetBackupMetadataRegion(d.Get("backup_metadata_region").(string))
		hasChange = true
	}
	if d.HasChange("private_api_endpoint_only") {
		updateSettingsOptions.SetPrivateAPIEndpointOnly(d.Get("private_api_endpoint_only").(bool))
		hasChange = true
	}

	if hasChange {
		_, _, err = logsRouterClient.UpdateSettingsWithContext(context, updateSettingsOptions)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("UpdateSettingsWithContext failed: %s", err.Error()), "ibm_logs_router_v3_settings", "delete")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
	}

	return resourceIBMLogsRouterSettingsRead(context, d, meta)
}

func resourceIBMLogsRouterSettingsDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logsRouterClient, err := meta.(conns.ClientSession).LogsRouterV3()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_logs_router_v3_settings", "delete", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	// Retrieve old settings and put them for required fields.  Remove all other fields
	settings, response, err := logsRouterClient.GetSettingsWithContext(context, &logsrouterv3.GetSettingsOptions{})
	if err != nil {
		log.Printf("[DEBUG] UpdateSettingsWithContext with GetSettingsWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("with GetSettingsWithContext failed %s\n%s", err, response))
	}

	updateSettingsOptions := &logsrouterv3.UpdateSettingsOptions{}

	updateSettingsOptions.PrimaryMetadataRegion = settings.PrimaryMetadataRegion
	updateSettingsOptions.BackupMetadataRegion = settings.BackupMetadataRegion
	updateSettingsOptions.PrivateAPIEndpointOnly = settings.PrivateAPIEndpointOnly
	updateSettingsOptions.PermittedTargetRegions = []string{}
	updateSettingsOptions.DefaultTargets = []logsrouterv3.TargetIdentity{}

	_, res, err := logsRouterClient.UpdateSettingsWithContext(context, updateSettingsOptions)
	if err != nil {
		log.Printf("[DEBUG] UpdateSettingsWithContext failed %s\n%s", err, res)
		return diag.FromErr(fmt.Errorf("UpdateSettingsWithContext failed %s\n%s", err, res))
	}

	d.SetId("")

	return nil
}

func ResourceIBMLogsRouterSettingsMapToTargetIdentity(modelMap map[string]interface{}) (*logsrouterv3.TargetIdentity, error) {
	model := &logsrouterv3.TargetIdentity{}
	model.ID = core.StringPtr(modelMap["id"].(string))
	return model, nil
}

func ResourceIBMLogsRouterSettingsTargetReferenceToMap(model *logsrouterv3.TargetReference) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["id"] = *model.ID
	return modelMap, nil
}
