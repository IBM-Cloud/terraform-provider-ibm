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

// dataSourceIBMDatabaseBackupGen2Backend holds the source database instance
// fetched by pickDataSourceBackupBackend so Read can check S2S authorization
// without a second GetResourceInstance call.
type dataSourceIBMDatabaseBackupGen2Backend struct {
	sourceInstance *rc.ResourceInstance
}

func newDataSourceIBMDatabaseBackupGen2Backend(sourceInstance *rc.ResourceInstance) dataSourceIBMDatabaseBackupBackend {
	return &dataSourceIBMDatabaseBackupGen2Backend{sourceInstance: sourceInstance}
}

func (g *dataSourceIBMDatabaseBackupGen2Backend) Read(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// Gen2 databases use Resource Controller API
	// Get the resource controller client to fetch instance details
	rsConClient, err := meta.(conns.ClientSession).ResourceControllerV2API()
	if err != nil {
		tfErr := flex.TerraformErrorf(err, err.Error(), "(Data) ibm_database_backup", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	backupID := d.Get("backup_id").(string)

	// Get the backup instance to verify it exists and is accessible.
	instance, response, err := rsConClient.GetResourceInstance(&rc.GetResourceInstanceOptions{
		ID: &backupID,
	})
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
		if err = d.Set(field, value); err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error setting %s: %s", field, err), "(Data) ibm_database_backup", "read")
			return tfErr.GetDiag()
		}
	}

	// Warn if S2S authorizations are not fully configured on the source database instance.
	// S2S authorization lives on the source instance, not on the backup resource itself.
	// Only applicable to instances using Independent Backups.
	if g.sourceInstance != nil &&
		hasIndependentBackups(g.sourceInstance.Extensions) &&
		!checkS2SAuthorization(g.sourceInstance.Extensions) {
		return diag.Diagnostics{{
			Severity: diag.Warning,
			Summary:  s2sAuthWarningHeader,
			Detail:   s2sAuthWarningDetail,
		}}
	}

	return nil
}
