// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.113.1-d76630af-20260320-135953
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

func DataSourceIbmBackupRecoveryVaultRecoveryScanStatus() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIbmBackupRecoveryVaultRecoveryScanStatusRead,

		Schema: map[string]*schema.Schema{
			"x_ibm_tenant_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Id of the tenant accessing the cluster.",
			},
			"cloud_type": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Specifies the cloud environment type where the Backup and Recovery instance is used. Currently, only 'ibm' is supported for recover scans.",
			},
			"vault_ids": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Specifies the unique ids of the Backup and Recovery instances for which the latest recovery scan status is to be fetched.",
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
			},
			"error_message": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Specifies the error message if the batch recovery scan status retrieval failed.",
			},
			"recovery_scan_statuses": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Array of recovery scan statuses for the specified Backup and Recovery instances.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"status": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Specifies the status of a Recovery Scan.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"end_time_usecs": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Specifies the end time of the recovery scan in microseconds since epoch.",
									},
									"error_message": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Specifies the error message if the recovery scan failed.",
									},
									"start_time_usecs": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Specifies the start time of the recovery scan in microseconds since epoch.",
									},
									"status": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Specifies the current status of the recovery scan.",
									},
									"uid": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Specifies the unique id of the recovery scan.",
									},
								},
							},
						},
						"vault_id": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Specifies the unique id of the Backup and Recovery instance.",
						},
					},
				},
			},
		},
	}
}

func dataSourceIbmBackupRecoveryVaultRecoveryScanStatusRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	backupRecoveryClient, err := meta.(conns.ClientSession).BackupRecoveryV1()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_backup_recovery_vault_recovery_scan_status", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getBatchVaultRecoveryScanStatusOptions := &backuprecoveryv1.GetBatchVaultRecoveryScanStatusOptions{}

	getBatchVaultRecoveryScanStatusOptions.SetXIBMTenantID(d.Get("x_ibm_tenant_id").(string))
	getBatchVaultRecoveryScanStatusOptions.SetCloudType(d.Get("cloud_type").(string))
	if _, ok := d.GetOk("vault_ids"); ok {
		var vaultIds []int64
		for _, v := range d.Get("vault_ids").([]interface{}) {
			vaultIdsItem := int64(v.(int))
			vaultIds = append(vaultIds, vaultIdsItem)
		}
		getBatchVaultRecoveryScanStatusOptions.SetVaultIds(vaultIds)
	}

	getBatchVaultRecoveryScanStatus, _, err := backupRecoveryClient.GetBatchVaultRecoveryScanStatusWithContext(context, getBatchVaultRecoveryScanStatusOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetBatchVaultRecoveryScanStatusWithContext failed: %s", err.Error()), "(Data) ibm_backup_recovery_vault_recovery_scan_status", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(dataSourceIbmBackupRecoveryVaultRecoveryScanStatusID(d))

	if !core.IsNil(getBatchVaultRecoveryScanStatus.ErrorMessage) {
		if err = d.Set("error_message", getBatchVaultRecoveryScanStatus.ErrorMessage); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting error_message: %s", err), "(Data) ibm_backup_recovery_vault_recovery_scan_status", "read", "set-error_message").GetDiag()
		}
	}

	if !core.IsNil(getBatchVaultRecoveryScanStatus.RecoveryScanStatuses) {
		recoveryScanStatuses := []map[string]interface{}{}
		for _, recoveryScanStatusesItem := range getBatchVaultRecoveryScanStatus.RecoveryScanStatuses {
			recoveryScanStatusesItemMap, err := DataSourceIbmBackupRecoveryVaultRecoveryScanStatusBatchVaultRecoveryScanStatusToMap(&recoveryScanStatusesItem) // #nosec G601
			if err != nil {
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_backup_recovery_vault_recovery_scan_status", "read", "recovery_scan_statuses-to-map").GetDiag()
			}
			recoveryScanStatuses = append(recoveryScanStatuses, recoveryScanStatusesItemMap)
		}
		if err = d.Set("recovery_scan_statuses", recoveryScanStatuses); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting recovery_scan_statuses: %s", err), "(Data) ibm_backup_recovery_vault_recovery_scan_status", "read", "set-recovery_scan_statuses").GetDiag()
		}
	}

	return nil
}

// dataSourceIbmBackupRecoveryVaultRecoveryScanStatusID returns a reasonable ID for the list.
func dataSourceIbmBackupRecoveryVaultRecoveryScanStatusID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}

func DataSourceIbmBackupRecoveryVaultRecoveryScanStatusBatchVaultRecoveryScanStatusToMap(model *backuprecoveryv1.BatchVaultRecoveryScanStatus) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Status != nil {
		statusMap, err := DataSourceIbmBackupRecoveryVaultRecoveryScanStatusRecoveryScanStatusToMap(model.Status)
		if err != nil {
			return modelMap, err
		}
		modelMap["status"] = []map[string]interface{}{statusMap}
	}
	if model.VaultID != nil {
		modelMap["vault_id"] = flex.IntValue(model.VaultID)
	}
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryVaultRecoveryScanStatusRecoveryScanStatusToMap(model *backuprecoveryv1.RecoveryScanStatus) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.EndTimeUsecs != nil {
		modelMap["end_time_usecs"] = flex.IntValue(model.EndTimeUsecs)
	}
	if model.ErrorMessage != nil {
		modelMap["error_message"] = *model.ErrorMessage
	}
	if model.StartTimeUsecs != nil {
		modelMap["start_time_usecs"] = flex.IntValue(model.StartTimeUsecs)
	}
	if model.Status != nil {
		modelMap["status"] = *model.Status
	}
	if model.Uid != nil {
		modelMap["uid"] = *model.Uid
	}
	return modelMap, nil
}
