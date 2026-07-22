// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"context"
	"fmt"
	"log"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	rc "github.com/IBM/platform-services-go-sdk/resourcecontrollerv2"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// dataSourceIBMDatabaseTaskGen2Backend implements task retrieval for Gen2 databases using RC API.
type dataSourceIBMDatabaseTaskGen2Backend struct {
	utils gen2TaskUtils
}

func newDataSourceIBMDatabaseTaskGen2Backend() dataSourceIBMDatabaseTaskBackend {
	return &dataSourceIBMDatabaseTaskGen2Backend{}
}

// Read retrieves task details for a Gen2 database instance using Resource Controller API.
func (g *dataSourceIBMDatabaseTaskGen2Backend) Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	rsConClient, err := meta.(conns.ClientSession).ResourceControllerV2API()
	if err != nil {
		tfErr := flex.TerraformErrorf(err, "failed to get Resource Controller client", "(Data) ibm_database_task", "read")
		return tfErr.GetDiag()
	}

	deploymentID := d.Get("task_id").(string)

	getInstanceOptions := &rc.GetResourceInstanceOptions{
		ID: &deploymentID,
	}

	instance, response, err := rsConClient.GetResourceInstanceWithContext(ctx, getInstanceOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("failed to get instance details: %s\n%s", err.Error(), response), "(Data) ibm_database_task", "read")
		return tfErr.GetDiag()
	}

	d.SetId(*instance.ID)
	log.Printf("[DEBUG] Setting Gen2 task ID: %s", *instance.ID)

	if err = d.Set("task_id", ""); err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error setting task_id: %s", err), "(Data) ibm_database_task", "read")
		return tfErr.GetDiag()
	}

	if err = d.Set("deployment_id", instance.ID); err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error setting deployment_id: %s", err), "(Data) ibm_database_task", "read")
		return tfErr.GetDiag()
	}

	description := g.utils.getOperationDescription(instance)
	if err = d.Set("description", description); err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error setting description: %s", err), "(Data) ibm_database_task", "read")
		return tfErr.GetDiag()
	}

	status := g.utils.mapStateToStatus(instance)
	log.Printf("[DEBUG] Gen2 task status: %s", status)
	if err = d.Set("status", status); err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error setting status: %s", err), "(Data) ibm_database_task", "read")
		return tfErr.GetDiag()
	}

	progressPercent := g.utils.calculateProgress(instance)
	log.Printf("[DEBUG] Gen2 task progress: %d%%", progressPercent)
	if err = d.Set("progress_percent", progressPercent); err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error setting progress_percent: %s", err), "(Data) ibm_database_task", "read")
		return tfErr.GetDiag()
	}

	createdAt := g.utils.getOperationTime(instance)
	if err = d.Set("created_at", createdAt); err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error setting created_at: %s", err), "(Data) ibm_database_task", "read")
		return tfErr.GetDiag()
	}

	return nil
}
