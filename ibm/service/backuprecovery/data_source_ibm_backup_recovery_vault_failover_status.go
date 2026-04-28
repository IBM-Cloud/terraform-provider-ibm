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

func DataSourceIbmBackupRecoveryVaultFailoverStatus() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIbmBackupRecoveryVaultFailoverStatusRead,

		Schema: map[string]*schema.Schema{
			"x_ibm_tenant_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Id of the tenant accessing the cluster.",
			},
			"cloud_type": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Specifies the cloud environment type where the Backup and Recovery instance is used. Currently, only 'ibm' is supported for failovers.",
			},
			"vault_ids": &schema.Schema{
				Type:        schema.TypeList,
				Required:    true,
				Description: "Specifies the unique ids of the Backup and Recovery instances i.e. vaults for which the latest failover status is to be fetched.",
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
			},
			"error_message": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Specifies the error message if the batch failover status retrieval failed.",
			},
			"failover_statuses": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Array of failover statuses for the specified Backup and Recovery instances.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"status": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Specifies the status of a vault Failover.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"end_time_usecs": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Specifies the end time of the failover in microseconds since epoch.",
									},
									"error_message": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Specifies the error message if the vault failover failed.",
									},
									"start_time_usecs": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Specifies the start time of the failover in microseconds since epoch.",
									},
									"status": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Specifies the current status of the failover.",
									},
									"uid": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Specifies the unique id of the failover.",
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

func dataSourceIbmBackupRecoveryVaultFailoverStatusRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	backupRecoveryClient, err := meta.(conns.ClientSession).BackupRecoveryV1()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_backup_recovery_vault_failover_status", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getBatchVaultsFailoverStatusOptions := &backuprecoveryv1.GetBatchVaultsFailoverStatusOptions{}

	getBatchVaultsFailoverStatusOptions.SetXIBMTenantID(d.Get("x_ibm_tenant_id").(string))
	getBatchVaultsFailoverStatusOptions.SetCloudType(d.Get("cloud_type").(string))
	var vaultIds []int64
	for _, v := range d.Get("vault_ids").([]interface{}) {
		vaultIdsItem := int64(v.(int))
		vaultIds = append(vaultIds, vaultIdsItem)
	}
	getBatchVaultsFailoverStatusOptions.SetVaultIds(vaultIds)

	getBatchVaultFailoverStatus, _, err := backupRecoveryClient.GetBatchVaultsFailoverStatusWithContext(context, getBatchVaultsFailoverStatusOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetBatchVaultsFailoverStatusWithContext failed: %s", err.Error()), "(Data) ibm_backup_recovery_vault_failover_status", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(dataSourceIbmBackupRecoveryVaultFailoverStatusID(d))

	if !core.IsNil(getBatchVaultFailoverStatus.ErrorMessage) {
		if err = d.Set("error_message", getBatchVaultFailoverStatus.ErrorMessage); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting error_message: %s", err), "(Data) ibm_backup_recovery_vault_failover_status", "read", "set-error_message").GetDiag()
		}
	}

	if !core.IsNil(getBatchVaultFailoverStatus.FailoverStatuses) {
		failoverStatuses := []map[string]interface{}{}
		for _, failoverStatusesItem := range getBatchVaultFailoverStatus.FailoverStatuses {
			failoverStatusesItemMap, err := DataSourceIbmBackupRecoveryVaultFailoverStatusBatchVaultFailoverStatusToMap(&failoverStatusesItem) // #nosec G601
			if err != nil {
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_backup_recovery_vault_failover_status", "read", "failover_statuses-to-map").GetDiag()
			}
			failoverStatuses = append(failoverStatuses, failoverStatusesItemMap)
		}
		if err = d.Set("failover_statuses", failoverStatuses); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting failover_statuses: %s", err), "(Data) ibm_backup_recovery_vault_failover_status", "read", "set-failover_statuses").GetDiag()
		}
	}

	return nil
}

// dataSourceIbmBackupRecoveryVaultFailoverStatusID returns a reasonable ID for the list.
func dataSourceIbmBackupRecoveryVaultFailoverStatusID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}

func DataSourceIbmBackupRecoveryVaultFailoverStatusBatchVaultFailoverStatusToMap(model *backuprecoveryv1.BatchVaultFailoverStatus) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Status != nil {
		statusMap, err := DataSourceIbmBackupRecoveryVaultFailoverStatusVaultFailoverStatusToMap(model.Status)
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

func DataSourceIbmBackupRecoveryVaultFailoverStatusVaultFailoverStatusToMap(model *backuprecoveryv1.VaultFailoverStatus) (map[string]interface{}, error) {
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
