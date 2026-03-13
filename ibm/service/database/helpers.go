// Copyright IBM Corp. 2017, 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"fmt"
	"log"
	"regexp"
	"strings"
	"time"

	"github.com/IBM/cloud-databases-go-sdk/clouddatabasesv5"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
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

func isAttrConfiguredInDiff(d *schema.ResourceDiff, k string) bool {
	v, ok := d.GetOkExists(k)
	if !ok {
		return false
	}
	switch t := v.(type) {
	case string:
		return t != ""
	case []interface{}:
		return len(t) > 0
	case map[string]interface{}:
		return len(t) > 0
	default:
		return true
	}
}

func isGen2Plan(plan string) bool {
	gen2Pattern := regexp.MustCompile(`-gen2($|-.+)`)
	return gen2Pattern.MatchString(strings.ToLower(plan))
}

// getDatabaseTypeFromResourceID maps the resource ID to the database type key used in extensions
func getDatabaseTypeFromResourceID(resourceID string) string {
	if strings.HasPrefix(resourceID, "databases-for-etcd") {
		return "etcd"
	} else if strings.HasPrefix(resourceID, "databases-for-postgresql") {
		return "postgresql"
	} else if strings.HasPrefix(resourceID, "databases-for-redis") {
		return "redis"
	} else if strings.HasPrefix(resourceID, "databases-for-elasticsearch") {
		return "elasticsearch"
	} else if strings.HasPrefix(resourceID, "databases-for-mongodb") {
		return "mongodb"
	} else if strings.HasPrefix(resourceID, "messages-for-rabbitmq") {
		return "rabbitmq"
	} else if strings.HasPrefix(resourceID, "databases-for-mysql") {
		return "mysql"
	} else if strings.HasPrefix(resourceID, "databases-for-enterprisedb") {
		return "enterprisedb"
	}
	return ""
}

// expandPlatformOptionsFromInstance extracts platform options from instance extensions for Gen2
func expandPlatformOptionsFromInstance(instance map[string]interface{}) []map[string]interface{} {
	pltOptions := make([]map[string]interface{}, 0, 1)
	pltOption := make(map[string]interface{})

	if dataservices, ok := instance["dataservices"].(map[string]interface{}); ok {
		if encryption, ok := dataservices["encryption"].(map[string]interface{}); ok {
			if disk, ok := encryption["disk"].(string); ok {
				pltOption["disk_encryption_key_crn"] = disk
			}
			if backup, ok := encryption["backup"].(string); ok {
				pltOption["backup_encryption_key_crn"] = backup
			}
		}
	}

	pltOptions = append(pltOptions, pltOption)
	return pltOptions
}
