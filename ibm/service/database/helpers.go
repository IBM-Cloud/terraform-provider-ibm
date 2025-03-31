// Copyright IBM Corp. 2017, 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"fmt"
	"log"
	"time"

	"github.com/IBM/cloud-databases-go-sdk/clouddatabasesv5"
	"github.com/IBM/go-sdk-core/v5/core"
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

func (t *TimeoutHelper) futureTimeToISO(duration time.Duration) string {
	return t.Now.Add(duration).UTC().Format(time.RFC3339)
}

func (t *TimeoutHelper) calculateExpirationDatetime(timeoutDuration time.Duration) string {
	if t.isMoreThan24Hours(timeoutDuration) {
		return t.futureTimeToISO(24 * time.Hour)
	}

	return t.futureTimeToISO(timeoutDuration)
}

func (tm *TaskManager) IsMatchingTaskInProgress(description string) (bool, *clouddatabasesv5.Task, error) {
	opts := &clouddatabasesv5.ListDeploymentTasksOptions{
		ID: core.StringPtr(tm.InstanceID),
	}

	resp, _, err := tm.Client.ListDeploymentTasks(opts)
	if err != nil {
		return false, nil, fmt.Errorf("failed to list tasks for instance: %w", err)
	}

	for _, task := range resp.Tasks {
		if task.Status == nil || task.Description == nil {
			continue
		}
		status := *task.Status
		desc := *task.Description

		if (status == databaseTaskRunningStatus || status == databaseTaskQueuedStatus) && desc == description {
			log.Printf("[INFO] Found matching task in progress: %s (status: %s)", desc, status)
			return true, &task, nil
		}
	}

	return false, nil, nil
}
