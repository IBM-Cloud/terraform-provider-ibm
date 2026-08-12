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

type dataSourceIBMDatabaseBackupsGen2Backend struct{}

func newDataSourceIBMDatabaseBackupsGen2Backend() dataSourceIBMDatabaseBackupsBackend {
	return &dataSourceIBMDatabaseBackupsGen2Backend{}
}

func (g *dataSourceIBMDatabaseBackupsGen2Backend) Read(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// Gen2 databases use Resource Controller API, not CloudDatabasesV5
	// Get the resource controller client to fetch instance details
	rsConClient, err := meta.(conns.ClientSession).ResourceControllerV2API()
	if err != nil {
		tfErr := flex.TerraformErrorf(err, err.Error(), "(Data) ibm_database_backups", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

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
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("ListResourceInstances failed: %s\n%s", err.Error(), response), "(Data) ibm_database_backups", "read")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}

		nextURL, err = getInstancesNext(listResponse.NextURL)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("ListResourceInstances failed while parsing NextURL: %s", err.Error()), "(Data) ibm_database_backups", "read")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}

		instances = append(instances, listResponse.Resources...)
		if nextURL == "" {
			break
		}
	}

	deploymentID := ""
	if v, ok := d.GetOk("deployment_id"); ok {
		deploymentID = v.(string)
	}

	if deploymentID != "" {
		d.SetId(deploymentID)
	} else {
		d.SetId(DataSourceIBMDatabaseBackupsID(d))
	}

	backups := make([]map[string]interface{}, 0, len(instances))
	for _, instance := range instances {
		if instance.CRN == nil {
			continue
		}

		var sourceDataServiceCRN string
		var backupType string
		if instance.Extensions != nil {
			if dataservices, ok := instance.Extensions["dataservices"].(map[string]interface{}); ok {
				if backupData, ok := dataservices["backup"].(map[string]interface{}); ok {
					if crnVal, ok := backupData["source_data_service_crn"].(string); ok {
						sourceDataServiceCRN = crnVal
					}
					if typeVal, ok := backupData["type"].(string); ok {
						backupType = typeVal
					}
				}
			}
		}

		if deploymentID != "" && sourceDataServiceCRN != deploymentID {
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
			"is_restorable":   backupState == "active",
			"download_link":   "",
		}
		if instance.CreatedAt != nil {
			backup["created_at"] = flex.DateTimeToString(instance.CreatedAt)
		}
		backups = append(backups, backup)
	}

	if err = d.Set("backups", backups); err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error setting backups: %s", err), "(Data) ibm_database_backups", "read")
		return tfErr.GetDiag()
	}

	return nil
}
