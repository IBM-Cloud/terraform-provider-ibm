// Copyright IBM Corp. 2017, 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"fmt"
	"log"
	"time"

	"github.com/IBM/cloud-databases-go-sdk/clouddatabasesv5"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/go-openapi/strfmt"
)

/*  TODO Move other helper functions here */
type TimeoutHelper struct {
	Now time.Time
}

// Allows mocking
type DeploymentTaskFetcher interface {
	ListDeploymentTasks(opts *clouddatabasesv5.ListDeploymentTasksOptions) (*clouddatabasesv5.Tasks, *core.DetailedResponse, error)
}
type TaskManager struct {
	Client     DeploymentTaskFetcher
	InstanceID string
}

func (t *TimeoutHelper) isMoreThan24Hours(duration time.Duration) bool {
	return duration > 24*time.Hour
}

func (t *TimeoutHelper) futureTimeToISO(duration time.Duration) strfmt.DateTime {
	utcTime := t.Now.Add(duration).UTC()
	return strfmt.DateTime(utcTime)
}

func (t *TimeoutHelper) calculateExpirationDatetime(timeoutDuration time.Duration) strfmt.DateTime {
	if t.isMoreThan24Hours(timeoutDuration) {
		return t.futureTimeToISO(24 * time.Hour)
	}

	return t.futureTimeToISO(timeoutDuration)
}

func (tm *TaskManager) matchingTaskInProgress(taskType string) (bool, *clouddatabasesv5.Task, error) {
	opts := &clouddatabasesv5.ListDeploymentTasksOptions{
		ID: core.StringPtr(tm.InstanceID),
	}

	resp, _, err := tm.Client.ListDeploymentTasks(opts)
	if err != nil {
		return false, nil, fmt.Errorf("failed to list tasks for instance: %w", err)
	}

	for _, task := range resp.Tasks {
		if task.Status == nil || task.ResourceType == nil {
			continue
		}
		id := *task.ID
		createdAt := *task.CreatedAt
		status := *task.Status
		progress := *task.ProgressPercent
		description := *task.Description
		resourceType := *task.ResourceType

		if (status == databaseTaskRunningStatus || status == databaseTaskQueuedStatus) && resourceType == taskType {
			log.Printf("[INFO] Found matching task in progress:\n"+
				"  Type: %s\n"+
				"  Created at: %s\n"+
				"  Status: %s\n"+
				"  Current progress percent: %d\n"+
				"  Description: %s\n"+
				"  ID: %s\n",
				resourceType, createdAt, status, progress, description, id)
			return true, &task, nil
		}
	}

	return false, nil, nil
}
