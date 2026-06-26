// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	rc "github.com/IBM/platform-services-go-sdk/resourcecontrollerv2"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// dataSourceIBMDatabaseTaskGen2Backend implements task retrieval for Gen2 databases using RC API.
type dataSourceIBMDatabaseTaskGen2Backend struct{}

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

	description := g.getOperationDescription(instance)
	if err = d.Set("description", description); err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error setting description: %s", err), "(Data) ibm_database_task", "read")
		return tfErr.GetDiag()
	}

	status := g.mapStateToStatus(instance)
	if err = d.Set("status", status); err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error setting status: %s", err), "(Data) ibm_database_task", "read")
		return tfErr.GetDiag()
	}

	progressPercent := g.calculateProgress(instance)
	if err = d.Set("progress_percent", progressPercent); err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error setting progress_percent: %s", err), "(Data) ibm_database_task", "read")
		return tfErr.GetDiag()
	}

	createdAt := g.getOperationTime(instance)
	if err = d.Set("created_at", createdAt); err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error setting created_at: %s", err), "(Data) ibm_database_task", "read")
		return tfErr.GetDiag()
	}

	return nil
}

func (g *dataSourceIBMDatabaseTaskGen2Backend) getOperationDescription(instance *rc.ResourceInstance) string {
	if instance.LastOperation != nil {
		log.Printf("[DEBUG] Gen2 task LastOperation: %+v", instance.LastOperation)
		if instance.LastOperation.Description != nil && *instance.LastOperation.Description != "" {
			return *instance.LastOperation.Description
		}
		if instance.LastOperation.Type != nil && *instance.LastOperation.Type != "" {
			return fmt.Sprintf("Operation: %s", *instance.LastOperation.Type)
		}
	}

	if instance.State != nil {
		return fmt.Sprintf("Instance state: %s", *instance.State)
	}

	return "Gen2 database instance operation"
}

func (g *dataSourceIBMDatabaseTaskGen2Backend) mapStateToStatus(instance *rc.ResourceInstance) string {
	if instance.State == nil {
		// Instance state is not available
		return "unknown"
	}

	state := *instance.State
	log.Printf("[DEBUG] Gen2 task mapping state '%s' to status", state)
	switch state {
	case "active":
		// Instance is fully provisioned and operational
		return "completed"
	case "provisioning", "in progress":
		// Instance is being created or an operation is in progress
		return "running"
	case "removed":
		// Instance has been deleted, operation is complete
		return "completed"
	default:
		// Return the original state for any unmapped states
		return state
	}
}

func (g *dataSourceIBMDatabaseTaskGen2Backend) calculateProgress(instance *rc.ResourceInstance) int {
	if instance.State == nil {
		// No state information available
		return 0
	}

	state := *instance.State
	log.Printf("[DEBUG] Gen2 task calculating progress for state '%s'", state)
	switch state {
	case "active":
		// Instance is fully provisioned and operational - 100% complete
		return 100
	case "provisioning":
		// Instance is being created - estimated at 50% (midpoint of provisioning process)
		// Note: Actual progress may vary; RC API doesn't provide granular progress data
		return 50
	case "in progress":
		// Operation is in progress - estimated at 75% (nearing completion)
		// Note: This is an approximation as RC API doesn't provide actual progress percentage
		return 75
	case "failed", "removed":
		// Operation has completed (either failed or instance removed) - 100% done
		return 100
	case "inactive":
		// Instance is stopped/suspended - no progress (0%)
		return 0
	default:
		// Unknown state - assume no progress
		return 0
	}
}

func (g *dataSourceIBMDatabaseTaskGen2Backend) getOperationTime(instance *rc.ResourceInstance) string {
	if instance.UpdatedAt != nil {
		return flex.DateTimeToString(instance.UpdatedAt)
	}
	if instance.CreatedAt != nil {
		return flex.DateTimeToString(instance.CreatedAt)
	}

	return time.Now().UTC().Format(time.RFC3339)
}
