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
	rsConClient, err := meta.(conns.ClientSession).ResourceControllerV2API()
	if err != nil {
		tfErr := flex.TerraformErrorf(err, err.Error(), "(Data) ibm_database_backup", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	backupID := d.Get("backup_id").(string)
	resourceID := "databases-independent-backups"
	listOptions := &rc.ListResourceInstancesOptions{ResourceID: &resourceID}

	var instances []rc.ResourceInstance
	nextURL := ""
	for {
		if nextURL != "" {
			listOptions.Start = &nextURL
		}

		listResponse, response, err := rsConClient.ListResourceInstances(listOptions)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("ListResourceInstances failed: %s\n%s", err.Error(), response), "(Data) ibm_database_backup", "read")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}

		nextURL, err = getInstancesNext(listResponse.NextURL)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("ListResourceInstances failed while parsing NextURL: %s", err.Error()), "(Data) ibm_database_backup", "read")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}

		instances = append(instances, listResponse.Resources...)
		if nextURL == "" {
			break
		}
	}

	for _, instance := range instances {
		if instance.CRN == nil || *instance.CRN != backupID {
			continue
		}

		d.SetId(backupID)
		if err := d.Set("backup_id", backupID); err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error setting backup_id: %s", err), "(Data) ibm_database_backup", "read")
			return tfErr.GetDiag()
		}
		if err := d.Set("deployment_id", nil); err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error setting deployment_id: %s", err), "(Data) ibm_database_backup", "read")
			return tfErr.GetDiag()
		}
		if err := d.Set("created_at", flex.DateTimeToString(instance.CreatedAt)); err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error setting created_at: %s", err), "(Data) ibm_database_backup", "read")
			return tfErr.GetDiag()
		}
		return nil
	}

	tfErr := flex.TerraformErrorf(fmt.Errorf("independent backup not found: %s", backupID), fmt.Sprintf("independent backup not found: %s", backupID), "(Data) ibm_database_backup", "read")
	log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
	return tfErr.GetDiag()
}
