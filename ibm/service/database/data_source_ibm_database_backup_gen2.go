// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	rc "github.com/IBM/platform-services-go-sdk/resourcecontrollerv2"
)

// dataSourceIBMDatabaseBackupGen2Backend holds the backup resource instance and
// its source database instance, both pre-fetched by pickDataSourceBackupBackend,
// so Read can map attributes and check S2S authorization without any extra API calls.
type dataSourceIBMDatabaseBackupGen2Backend struct {
	backupInstance *rc.ResourceInstance
	sourceInstance *rc.ResourceInstance
}

func newDataSourceIBMDatabaseBackupGen2Backend(backupInstance, sourceInstance *rc.ResourceInstance) dataSourceIBMDatabaseBackupBackend {
	return &dataSourceIBMDatabaseBackupGen2Backend{
		backupInstance: backupInstance,
		sourceInstance: sourceInstance,
	}
}

func (g *dataSourceIBMDatabaseBackupGen2Backend) Read(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	backupID := d.Get("backup_id").(string)

	// Use the pre-fetched backup instance when available; fall back to a live
	// API call only when the router could not fetch it (e.g. auth failure).
	instance := g.backupInstance
	if instance == nil {
		rsConClient, err := meta.(conns.ClientSession).ResourceControllerV2API()
		if err != nil {
			tfErr := flex.TerraformErrorf(err, err.Error(), "(Data) ibm_database_backup", "read")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		resp, response, err := rsConClient.GetResourceInstance(&rc.GetResourceInstanceOptions{ID: &backupID})
		if err != nil {
			if response != nil && response.StatusCode == httpNotFound {
				tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Independent Backup not found: %s", backupID), "(Data) ibm_database_backup", "read")
				log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
				return tfErr.GetDiag()
			}
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetResourceInstance failed: %s\n%s", err.Error(), response), "(Data) ibm_database_backup", "read")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		instance = resp
	}

	d.SetId(backupID)

	sourceDataServiceCRN, backupType := extractGen2BackupExtensions(instance.Extensions)

	backupState := ""
	if instance.State != nil {
		backupState = *instance.State
	}

	fields := map[string]interface{}{
		"backup_id":       backupID,
		"deployment_id":   sourceDataServiceCRN,
		"type":            backupType,
		"status":          backupState,
		"is_downloadable": false,
		"is_restorable":   backupState == databaseInstanceSuccessStatus,
		"download_link":   "",
		"created_at":      flex.DateTimeToString(instance.CreatedAt),
	}

	for field, value := range fields {
		if err := d.Set(field, value); err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error setting %s: %s", field, err), "(Data) ibm_database_backup", "read")
			return tfErr.GetDiag()
		}
	}

	// Warn if S2S authorizations are not fully configured on the source database instance.
	// S2S authorization lives on the source instance, not on the backup resource itself.
	// Only applicable to instances using Independent Backups.
	return s2sDiagForInstance(g.sourceInstance)
}

// s2sDiagForInstance returns a non-blocking S2S warning diagnostic when the
// given instance has Independent Backups but the required S2S authorizations
// are missing. Returns nil when no warning is needed.
// Extracted so both the backup and backups Gen2 backends share the same logic
// and unit tests can exercise this function directly.
func s2sDiagForInstance(instance *rc.ResourceInstance) diag.Diagnostics {
	if instance != nil &&
		hasIndependentBackups(instance.Extensions) &&
		!checkS2SAuthorization(instance.Extensions) {
		return diag.Diagnostics{{
			Severity: diag.Warning,
			Summary:  s2sAuthWarningHeader,
			Detail:   s2sAuthWarningDetail,
		}}
	}
	return nil
}
