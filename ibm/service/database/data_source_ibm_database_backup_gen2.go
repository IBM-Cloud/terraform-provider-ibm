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
	// Gen2 databases use Resource Controller API, not CloudDatabasesV5
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

	if err = d.Set("backup_id", backupID); err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error setting backup_id: %s", err), "(Data) ibm_database_backup", "read")
		return tfErr.GetDiag()
	}
	if err = d.Set("deployment_id", sourceDataServiceCRN); err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error setting deployment_id: %s", err), "(Data) ibm_database_backup", "read")
		return tfErr.GetDiag()
	}
	if err = d.Set("type", backupType); err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error setting type: %s", err), "(Data) ibm_database_backup", "read")
		return tfErr.GetDiag()
	}
	if err = d.Set("status", backupState); err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error setting status: %s", err), "(Data) ibm_database_backup", "read")
		return tfErr.GetDiag()
	}
	if err = d.Set("is_downloadable", false); err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error setting is_downloadable: %s", err), "(Data) ibm_database_backup", "read")
		return tfErr.GetDiag()
	}
	if err = d.Set("is_restorable", backupState == "active"); err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error setting is_restorable: %s", err), "(Data) ibm_database_backup", "read")
		return tfErr.GetDiag()
	}
	if err = d.Set("download_link", ""); err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error setting download_link: %s", err), "(Data) ibm_database_backup", "read")
		return tfErr.GetDiag()
	}
	if err = d.Set("created_at", flex.DateTimeToString(instance.CreatedAt)); err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error setting created_at: %s", err), "(Data) ibm_database_backup", "read")
		return tfErr.GetDiag()
	}

	return nil
}
