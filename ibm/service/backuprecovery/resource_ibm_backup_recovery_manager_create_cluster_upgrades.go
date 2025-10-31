// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.96.1-5136e54a-20241108-203028
 */

package backuprecovery

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/ibm-backup-recovery-sdk-go/backuprecoveryv1"
)

func ResourceIbmBackupRecoveryManagerCreateClusterUpgrades() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIbmBackupRecoveryManagerCreateClusterUpgradesCreate,
		ReadContext:   resourceIbmBackupRecoveryManagerCreateClusterUpgradesRead,
		DeleteContext: resourceIbmBackupRecoveryManagerCreateClusterUpgradesDelete,
		Importer:      &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"upgrade_response_list": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Specifies a list of disks to exclude from being protected. This is only applicable to VM objects.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"error_message": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "Specifies error message if failed to schedule upgrade.",
						},
						"is_upgrade_scheduling_successful": &schema.Schema{
							Type:        schema.TypeBool,
							Optional:    true,
							Computed:    true,
							Description: "Specifies if upgrade scheduling was successsful.",
						},
						"cluster_id": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Computed:    true,
							Description: "Specifies cluster id.",
						},
						"cluster_incarnation_id": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Computed:    true,
							Description: "Specifies cluster incarnation id.",
						},
					},
				},
			},
		},
	}
}

func resourceIbmBackupRecoveryManagerCreateClusterUpgradesCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	managementSreApiClient, err := meta.(conns.ClientSession).BackupRecoveryManagerV1()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_backup_recovery_manager_create_cluster_upgrades", "create", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	createClustersUpgradesOptions := &backuprecoveryv1.CreateClustersUpgradesOptions{}

	if _, ok := d.GetOk("auth_headers"); ok {
		var authHeaders []backuprecoveryv1.AuthHeaderForClusterUpgrade
		for _, v := range d.Get("auth_headers").([]interface{}) {
			value := v.(map[string]interface{})
			authHeadersItem, err := ResourceIbmBackupRecoveryManagerCreateClusterUpgradesMapToAuthHeaderForClusterUpgrade(value)
			if err != nil {
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_backup_recovery_manager_create_cluster_upgrades", "create", "parse-auth_headers").GetDiag()
			}
			authHeaders = append(authHeaders, *authHeadersItem)
		}
		createClustersUpgradesOptions.SetAuthHeaders(authHeaders)
	}
	if _, ok := d.GetOk("clusters"); ok {
		var clusters []backuprecoveryv1.Upgrade
		for _, v := range d.Get("clusters").([]interface{}) {
			value := v.(map[string]interface{})
			clustersItem, err := ResourceIbmBackupRecoveryManagerCreateClusterUpgradesMapToUpgrade(value)
			if err != nil {
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_backup_recovery_manager_create_cluster_upgrades", "create", "parse-clusters").GetDiag()
			}
			clusters = append(clusters, *clustersItem)
		}
		createClustersUpgradesOptions.SetClusters(clusters)
	}
	if _, ok := d.GetOk("interval_for_rolling_upgrade_in_hours"); ok {
		createClustersUpgradesOptions.SetIntervalForRollingUpgradeInHours(int64(d.Get("interval_for_rolling_upgrade_in_hours").(int)))
	}
	if _, ok := d.GetOk("package_url"); ok {
		createClustersUpgradesOptions.SetPackageURL(d.Get("package_url").(string))
	}
	if _, ok := d.GetOk("patch_upgrade_params"); ok {
		patchUpgradeParamsModel, err := ResourceIbmBackupRecoveryManagerCreateClusterUpgradesMapToPatchUpgradeParams(d.Get("patch_upgrade_params.0").(map[string]interface{}))
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_backup_recovery_manager_create_cluster_upgrades", "create", "parse-patch_upgrade_params").GetDiag()
		}
		createClustersUpgradesOptions.SetPatchUpgradeParams(patchUpgradeParamsModel)
	}
	if _, ok := d.GetOk("target_version"); ok {
		createClustersUpgradesOptions.SetTargetVersion(d.Get("target_version").(string))
	}
	if _, ok := d.GetOk("time_stamp_to_upgrade_at_msecs"); ok {
		createClustersUpgradesOptions.SetTimeStampToUpgradeAtMsecs(int64(d.Get("time_stamp_to_upgrade_at_msecs").(int)))
	}
	if _, ok := d.GetOk("type"); ok {
		createClustersUpgradesOptions.SetType(d.Get("type").(string))
	}

	upgradesResponse, _, err := managementSreApiClient.CreateClustersUpgradesWithContext(context, createClustersUpgradesOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("CreateClustersUpgradesWithContext failed: %s", err.Error()), "ibm_backup_recovery_manager_create_cluster_upgrades", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(resourceIbmBackupRecoveryManagerCreateClusterUpgradesID(d))

	if !core.IsNil(upgradesResponse) {
		upgradeResponseListResult := []map[string]interface{}{}
		for _, upgradeResponseListItem := range upgradesResponse {
			upgradeResponseListItemMap, err := ResourceIbmBackupRecoveryManagerCreateClusterUpgradesUpgradeResponseToMap(&upgradeResponseListItem) // #nosec G601
			if err != nil {
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_backup_recovery_manager_create_cluster_upgrades", "read", "upgrade_response_list-to-map").GetDiag()
			}
			upgradeResponseListResult = append(upgradeResponseListResult, upgradeResponseListItemMap)
		}
		if err = d.Set("upgrade_response_list", upgradeResponseListResult); err != nil {
			err = fmt.Errorf("Error setting upgrade_response_list: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_backup_recovery_manager_create_cluster_upgrades", "read", "set-upgrade_response_list").GetDiag()
		}
	}

	return resourceIbmBackupRecoveryManagerCreateClusterUpgradesRead(context, d, meta)
}

func ResourceIbmBackupRecoveryManagerCreateClusterUpgradesUpgradeResponseToMap(model *backuprecoveryv1.UpgradeResponse) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ErrorMessage != nil {
		modelMap["error_message"] = *model.ErrorMessage
	}
	if model.IsUpgradeSchedulingSuccessful != nil {
		modelMap["is_upgrade_scheduling_successful"] = *model.IsUpgradeSchedulingSuccessful
	}
	if model.ClusterID != nil {
		modelMap["cluster_id"] = flex.IntValue(model.ClusterID)
	}
	if model.ClusterIncarnationID != nil {
		modelMap["cluster_incarnation_id"] = flex.IntValue(model.ClusterIncarnationID)
	}
	return modelMap, nil
}

func resourceIbmBackupRecoveryManagerCreateClusterUpgradesID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}

func resourceIbmBackupRecoveryManagerCreateClusterUpgradesRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return nil
}

func resourceIbmBackupRecoveryManagerCreateClusterUpgradesDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// This resource does not support a "delete" operation.
	var diags diag.Diagnostics
	warning := diag.Diagnostic{
		Severity: diag.Warning,
		Summary:  "Delete Not Supported",
		Detail:   "The resource definition will be only be removed from the terraform statefile. This resource cannot be deleted from the backend. ",
	}
	diags = append(diags, warning)
	d.SetId("")
	return diags
}

func ResourceIbmBackupRecoveryManagerCreateClusterUpgradesMapToAuthHeaderForClusterUpgrade(modelMap map[string]interface{}) (*backuprecoveryv1.AuthHeaderForClusterUpgrade, error) {
	model := &backuprecoveryv1.AuthHeaderForClusterUpgrade{}
	model.Key = core.StringPtr(modelMap["key"].(string))
	model.Value = core.StringPtr(modelMap["value"].(string))
	return model, nil
}

func ResourceIbmBackupRecoveryManagerCreateClusterUpgradesMapToUpgrade(modelMap map[string]interface{}) (*backuprecoveryv1.Upgrade, error) {
	model := &backuprecoveryv1.Upgrade{}
	if modelMap["cluster_id"] != nil {
		model.ClusterID = core.Int64Ptr(int64(modelMap["cluster_id"].(int)))
	}
	if modelMap["cluster_incarnation_id"] != nil {
		model.ClusterIncarnationID = core.Int64Ptr(int64(modelMap["cluster_incarnation_id"].(int)))
	}
	if modelMap["current_version"] != nil && modelMap["current_version"].(string) != "" {
		model.CurrentVersion = core.StringPtr(modelMap["current_version"].(string))
	}
	return model, nil
}

func ResourceIbmBackupRecoveryManagerCreateClusterUpgradesMapToPatchUpgradeParams(modelMap map[string]interface{}) (*backuprecoveryv1.PatchUpgradeParams, error) {
	model := &backuprecoveryv1.PatchUpgradeParams{}
	if modelMap["auth_headers"] != nil {
		authHeaders := []backuprecoveryv1.AuthHeaderForClusterUpgrade{}
		for _, authHeadersItem := range modelMap["auth_headers"].([]interface{}) {
			authHeadersItemModel, err := ResourceIbmBackupRecoveryManagerCreateClusterUpgradesMapToAuthHeaderForClusterUpgrade(authHeadersItem.(map[string]interface{}))
			if err != nil {
				return model, err
			}
			authHeaders = append(authHeaders, *authHeadersItemModel)
		}
		model.AuthHeaders = authHeaders
	}
	if modelMap["ignore_pre_checks_failure"] != nil {
		model.IgnorePreChecksFailure = core.BoolPtr(modelMap["ignore_pre_checks_failure"].(bool))
	}
	if modelMap["package_url"] != nil && modelMap["package_url"].(string) != "" {
		model.PackageURL = core.StringPtr(modelMap["package_url"].(string))
	}
	if modelMap["target_version"] != nil && modelMap["target_version"].(string) != "" {
		model.TargetVersion = core.StringPtr(modelMap["target_version"].(string))
	}
	return model, nil
}
