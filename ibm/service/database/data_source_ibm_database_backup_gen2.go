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

type dataSourceIBMDatabaseBackupGen2Backend struct{}

func newDataSourceIBMDatabaseBackupGen2Backend() dataSourceIBMDatabaseBackupBackend {
	return &dataSourceIBMDatabaseBackupGen2Backend{}
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

	// Get the instance to verify it exists and is accessible
	instance, response, err := rsConClient.GetResourceInstance(&rc.GetResourceInstanceOptions{
		ID: &backupID,
	})
	if err != nil {
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

	return nil
}
