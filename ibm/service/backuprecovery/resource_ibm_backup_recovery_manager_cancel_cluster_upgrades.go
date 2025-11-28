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

func ResourceIbmBackupRecoveryManagerCancelClusterUpgrades() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIbmBackupRecoveryManagerCancelClusterUpgradesCreate,
		ReadContext:   resourceIbmBackupRecoveryManagerCancelClusterUpgradesRead,
		DeleteContext: resourceIbmBackupRecoveryManagerCancelClusterUpgradesDelete,
		UpdateContext: resourceIbmBackupRecoveryManagerCancelClusterUpgradesUpdate,
		Importer:      &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"cluster_identifiers": &schema.Schema{
				Type:        schema.TypeList,
				Required:    true,
				ForceNew:    true,
				Description: "Specifies the list of cluster identifiers. The format is clusterId:clusterIncarnationId.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"cancelled_upgrade_response_list": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Specifies list of cluster scheduled uprgade cancel response.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"error_message": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "Specifies an error message if failed to cancel a scheduled upgrade.",
						},
						"is_upgrade_cancel_successful": &schema.Schema{
							Type:        schema.TypeBool,
							Optional:    true,
							Computed:    true,
							Description: "Specifies if scheduled upgrade cancelling was successful.",
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

func resourceIbmBackupRecoveryManagerCancelClusterUpgradesCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	managementApiClient, err := meta.(conns.ClientSession).BackupRecoveryManagerV1()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_backup_recovery_manager_cancel_cluster_upgrades", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	bmxsession, err := meta.(conns.ClientSession).BluemixSession()
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("unable to get clientSession"), "ibm_backup_recovery_manager_cancel_cluster_upgrades", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	endpointType := d.Get("endpoint_type").(string)
	instanceId, region := getInstanceIdAndRegion(d)
	managementApiClient, err = setManagerClientAuth(managementApiClient, bmxsession, region, endpointType)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("unable to set authenticator for clientSession: %s", err), "ibm_backup_recovery_manager_cancel_cluster_upgrades", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	if instanceId != "" {
		managementApiClient = getManagerClientWithInstanceEndpoint(managementApiClient, bmxsession, instanceId, region, endpointType)
	}

	deleteClustersUpgradesOptions := &backuprecoveryv1.DeleteClustersUpgradesOptions{}

	if _, ok := d.GetOk("cluster_identifiers"); ok {
		var clusterIdentifiers []string
		for _, v := range d.Get("cluster_identifiers").([]interface{}) {
			clusterIdentifiersItem := v.(string)
			clusterIdentifiers = append(clusterIdentifiers, clusterIdentifiersItem)
		}
		deleteClustersUpgradesOptions.SetClusterIdentifiers(clusterIdentifiers)
	}

	upgradesCancelResponse, _, err := managementApiClient.DeleteClustersUpgradesWithContext(context, deleteClustersUpgradesOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("DeleteClustersUpgradesWithContext failed: %s", err.Error()), "ibm_backup_recovery_manager_cancel_cluster_upgrades", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(resourceIbmBackupRecoveryManagerCancelClusterUpgradesID(d))

	if !core.IsNil(upgradesCancelResponse) {
		cancelledUpgradeResponseListResult := []map[string]interface{}{}
		for _, cancelledUpgradeResponseListItem := range upgradesCancelResponse {
			cancelledUpgradeResponseListItemMap, err := ResourceIbmBackupRecoveryManagerCancelClusterUpgradesUpgradeCancelResponseToMap(&cancelledUpgradeResponseListItem) // #nosec G601
			if err != nil {
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_backup_recovery_manager_cancel_cluster_upgrades", "read", "cancelled_upgrade_response_list-to-map").GetDiag()
			}
			cancelledUpgradeResponseListResult = append(cancelledUpgradeResponseListResult, cancelledUpgradeResponseListItemMap)
		}
		if err = d.Set("cancelled_upgrade_response_list", cancelledUpgradeResponseListResult); err != nil {
			err = fmt.Errorf("Error setting cancelled_upgrade_response_list: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_backup_recovery_manager_cancel_cluster_upgrades", "read", "set-cancelled_upgrade_response_list").GetDiag()
		}
	}

	return resourceIbmBackupRecoveryManagerCancelClusterUpgradesRead(context, d, meta)
}

func resourceIbmBackupRecoveryManagerCancelClusterUpgradesID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}

func resourceIbmBackupRecoveryManagerCancelClusterUpgradesRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return nil
}

func resourceIbmBackupRecoveryManagerCancelClusterUpgradesUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// This resource does not support a "delete" operation.
	var diags diag.Diagnostics
	warning := diag.Diagnostic{
		Severity: diag.Warning,
		Summary:  "Update Not Supported",
		Detail:   "The resource definition will be only be removed from the terraform statefile. This resource cannot be deleted from the backend. ",
	}
	diags = append(diags, warning)
	d.SetId("")
	return diags
}

func resourceIbmBackupRecoveryManagerCancelClusterUpgradesDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

func ResourceIbmBackupRecoveryManagerCancelClusterUpgradesUpgradeCancelResponseToMap(model *backuprecoveryv1.UpgradeCancelResponse) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ErrorMessage != nil {
		modelMap["error_message"] = *model.ErrorMessage
	}
	if model.IsUpgradeCancelSuccessful != nil {
		modelMap["is_upgrade_cancel_successful"] = *model.IsUpgradeCancelSuccessful
	}
	if model.ClusterID != nil {
		modelMap["cluster_id"] = flex.IntValue(model.ClusterID)
	}
	if model.ClusterIncarnationID != nil {
		modelMap["cluster_incarnation_id"] = flex.IntValue(model.ClusterIncarnationID)
	}
	return modelMap, nil
}
