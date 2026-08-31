// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"fmt"
	"time"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	rc "github.com/IBM/platform-services-go-sdk/resourcecontrollerv2"
)

// gen2GetOperationDescription extracts a human-readable description from the instance's last operation or state.
func gen2GetOperationDescription(instance *rc.ResourceInstance) string {
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

// gen2MapStateToStatus converts Resource Controller instance state to task status.
// Maps Gen2 instance states to standardized task statuses for consistency with classic databases.
func gen2MapStateToStatus(instance *rc.ResourceInstance) string {
	if instance.State == nil {
		// Instance state is not available
		return "unknown"
	}

	state := *instance.State
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

// gen2CalculateProgress estimates task completion percentage based on instance state.
// Note: RC API doesn't provide granular progress data, so these are approximations.
func gen2CalculateProgress(instance *rc.ResourceInstance) int {
	if instance.State == nil {
		// No state information available
		return 0
	}

	state := *instance.State
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

// gen2GetOperationTime returns the most recent timestamp for the instance operation.
// Prefers UpdatedAt over CreatedAt, falls back to current time if neither is available.
func gen2GetOperationTime(instance *rc.ResourceInstance) string {
	if instance.UpdatedAt != nil {
		return flex.DateTimeToString(instance.UpdatedAt)
	}
	if instance.CreatedAt != nil {
		return flex.DateTimeToString(instance.CreatedAt)
	}

	return time.Now().UTC().Format(time.RFC3339)
}
