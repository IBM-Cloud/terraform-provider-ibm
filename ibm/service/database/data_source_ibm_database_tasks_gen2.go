// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"context"
	"fmt"
	"time"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	rc "github.com/IBM/platform-services-go-sdk/resourcecontrollerv2"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// dataSourceIBMDatabaseTasksGen2Backend implements tasks retrieval for Gen2 databases using RC API.
type dataSourceIBMDatabaseTasksGen2Backend struct{}

func newDataSourceIBMDatabaseTasksGen2Backend() dataSourceIBMDatabaseTasksBackend {
	return &dataSourceIBMDatabaseTasksGen2Backend{}
}

// Read retrieves task list for a Gen2 database instance using Resource Controller API.
// Note: Gen2 databases don't have individual task tracking like classic databases.
// This implementation returns a single "task" representing the current instance state.
func (g *dataSourceIBMDatabaseTasksGen2Backend) Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	rsConClient, err := meta.(conns.ClientSession).ResourceControllerV2API()
	if err != nil {
		tfErr := flex.TerraformErrorf(err, "failed to get Resource Controller client", "(Data) ibm_database_tasks", "read")
		return tfErr.GetDiag()
	}

	deploymentID := d.Get("deployment_id").(string)

	getInstanceOptions := &rc.GetResourceInstanceOptions{
		ID: &deploymentID,
	}

	instance, response, err := rsConClient.GetResourceInstanceWithContext(ctx, getInstanceOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("failed to get instance details: %s\n%s", err.Error(), response), "(Data) ibm_database_tasks", "read")
		return tfErr.GetDiag()
	}

	// Validate Gen2 instance
	if instance.ResourcePlanID == nil {
		tfErr := flex.TerraformErrorf(fmt.Errorf("instance resource plan ID is nil"), "cannot determine database generation", "(Data) ibm_database_tasks", "read")
		return tfErr.GetDiag()
	}

	if !isGen2Plan(*instance.ResourcePlanID) {
		tfErr := flex.TerraformErrorf(
			fmt.Errorf("instance is not a Gen2 database"),
			"this data source is for Gen2 databases only",
			"(Data) ibm_database_tasks",
			"read")
		return tfErr.GetDiag()
	}

	d.SetId(deploymentID)

	// Gen2 doesn't have multiple tasks like classic databases
	// Create a single task entry representing the current instance state
	tasks := []map[string]interface{}{
		g.instanceToTaskMap(instance),
	}

	if err = d.Set("tasks", tasks); err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error setting tasks: %s", err), "(Data) ibm_database_tasks", "read")
		return tfErr.GetDiag()
	}

	return nil
}

// instanceToTaskMap converts a Gen2 instance to a task map structure
func (g *dataSourceIBMDatabaseTasksGen2Backend) instanceToTaskMap(instance *rc.ResourceInstance) map[string]interface{} {
	taskMap := make(map[string]interface{})

	// Gen2 doesn't have task_id from last_operation, so set it to empty
	taskMap["task_id"] = ""

	if instance.ID != nil {
		taskMap["deployment_id"] = *instance.ID
	}

	taskMap["description"] = g.getOperationDescription(instance)
	taskMap["status"] = g.mapStateToStatus(instance)
	taskMap["progress_percent"] = g.calculateProgress(instance)
	taskMap["created_at"] = g.getOperationTime(instance)

	return taskMap
}

func (g *dataSourceIBMDatabaseTasksGen2Backend) getOperationDescription(instance *rc.ResourceInstance) string {
	if instance.LastOperation != nil {
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

func (g *dataSourceIBMDatabaseTasksGen2Backend) mapStateToStatus(instance *rc.ResourceInstance) string {
	if instance.State == nil {
		return "unknown"
	}

	state := *instance.State
	switch state {
	case "active":
		return "completed"
	case "provisioning", "in progress":
		return "running"
	case "failed":
		return "failed"
	case "inactive":
		return "queued"
	case "removed":
		return "completed"
	default:
		return state
	}
}

func (g *dataSourceIBMDatabaseTasksGen2Backend) calculateProgress(instance *rc.ResourceInstance) int {
	if instance.State == nil {
		return 0
	}

	state := *instance.State
	switch state {
	case "active":
		return 100
	case "provisioning":
		return 50
	case "in progress":
		return 75
	case "failed", "removed":
		return 100
	case "inactive":
		return 0
	default:
		return 0
	}
}

func (g *dataSourceIBMDatabaseTasksGen2Backend) getOperationTime(instance *rc.ResourceInstance) string {
	if instance.UpdatedAt != nil {
		return flex.DateTimeToString(instance.UpdatedAt)
	}
	if instance.CreatedAt != nil {
		return flex.DateTimeToString(instance.CreatedAt)
	}

	return time.Now().UTC().Format(time.RFC3339)
}

// Made with Bob
