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

// dataSourceIBMDatabaseBackupsGen2Backend holds the ResourceGroupID obtained
// by pickDataSourceBackupsBackend so Read does not need a second
// GetResourceInstance call.
type dataSourceIBMDatabaseBackupsGen2Backend struct {
	resourceGroupID *string
}

func newDataSourceIBMDatabaseBackupsGen2Backend(resourceGroupID *string) dataSourceIBMDatabaseBackupsBackend {
	return &dataSourceIBMDatabaseBackupsGen2Backend{resourceGroupID: resourceGroupID}
}

func (g *dataSourceIBMDatabaseBackupsGen2Backend) Read(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// Gen2 databases use Resource Controller API
	rsConClient, err := meta.(conns.ClientSession).ResourceControllerV2API()
	if err != nil {
		tfErr := flex.TerraformErrorf(err, err.Error(), "(Data) ibm_database_backups", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	// deployment_id is required by pickDataSourceBackupsBackend before this
	// backend is ever selected, so it is always populated here.
	deploymentID := d.Get("deployment_id").(string)
	d.SetId(deploymentID)

	// ResourceGroupID was fetched once by pickDataSourceBackupsBackend —
	// no second GetResourceInstance call needed.
	resourceID := "databases-independent-backups"
	listOptions := &rc.ListResourceInstancesOptions{ResourceID: &resourceID}
	if g.resourceGroupID != nil {
		listOptions.ResourceGroupID = g.resourceGroupID
	}

	backups := []map[string]interface{}{}
	nextURL := ""
	for {
		if nextURL != "" {
			listOptions.Start = &nextURL
		}

		listResponse, response, err := rsConClient.ListResourceInstances(listOptions)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("ListResourceInstances failed: %s\n%s", err.Error(), response), "(Data) ibm_database_backups", "read")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}

		// Filter and map each page as it arrives, appending directly into the
		// accumulator instead of allocating a per-page slice and copying it
		// into backups.
		backups = filterGen2BackupsByDeployment(backups, listResponse.Resources, deploymentID)

		nextURL, err = getInstancesNext(listResponse.NextURL)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("ListResourceInstances failed while parsing NextURL: %s", err.Error()), "(Data) ibm_database_backups", "read")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}

		if nextURL == "" {
			break
		}
	}

	if len(backups) == 0 {
		tfErr := flex.TerraformErrorf(fmt.Errorf("Independent Backup not found for deployment_id: %s", deploymentID), fmt.Sprintf("Independent Backups not found for deployment_id: %s", deploymentID), "(Data) ibm_database_backups", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	if err = d.Set("backups", backups); err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error setting backups: %s", err), "(Data) ibm_database_backups", "read")
		return tfErr.GetDiag()
	}

	return nil
}

// filterGen2BackupsByDeployment filters instances to those whose Gen2 backup
// extensions indicate they belong to deploymentID, maps each match to the
// "backups" schema attribute shape, and appends it to backups. Extracted from
// Read so the deployment scoping logic can be unit tested without a Resource
// Controller client, and so callers can accumulate matches across pages
// without an extra per-page allocation and copy.
func filterGen2BackupsByDeployment(backups []map[string]interface{}, instances []rc.ResourceInstance, deploymentID string) []map[string]interface{} {
	for _, instance := range instances {
		if instance.CRN == nil {
			continue
		}

		sourceDataServiceCRN, backupType := extractGen2BackupExtensions(instance.Extensions)
		if sourceDataServiceCRN != deploymentID {
			continue
		}

		backupState := ""
		if instance.State != nil {
			backupState = *instance.State
		}

		backup := map[string]interface{}{
			"backup_id":       *instance.CRN,
			"deployment_id":   sourceDataServiceCRN,
			"type":            backupType,
			"status":          backupState,
			"is_downloadable": false,
			"is_restorable":   backupState == databaseInstanceSuccessStatus,
			"download_link":   "",
		}
		if instance.CreatedAt != nil {
			backup["created_at"] = flex.DateTimeToString(instance.CreatedAt)
		}
		backups = append(backups, backup)
	}
	return backups
}
