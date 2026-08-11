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

	d.SetId(DataSourceIBMDatabaseBackupsID(d))

	backups := make([]map[string]interface{}, 0, len(instances))
	for _, instance := range instances {
		if instance.CRN == nil {
			continue
		}

		backup := map[string]interface{}{
			"backup_id": *instance.CRN,
		}
		if instance.CreatedAt != nil {
			backup["created_at"] = flex.DateTimeToString(instance.CreatedAt)
		}
		backups = append(backups, backup)
	}

	if err := d.Set("backups", backups); err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error setting backups: %s", err), "(Data) ibm_database_backups", "read")
		return tfErr.GetDiag()
	}

	return nil
}
