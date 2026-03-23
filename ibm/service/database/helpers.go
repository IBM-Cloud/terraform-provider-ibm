// Copyright IBM Corp. 2017, 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"fmt"
	"log"
	"regexp"
	"strings"
	"time"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM/cloud-databases-go-sdk/clouddatabasesv5"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/globalcatalogv1"
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

// getDatabaseTypeFromResourceID maps the resource ID or service name to the database type key
// Used in extensions for Gen2 and in parameters structure
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

// expandPlatformOptionsFromRCExtension extracts platform options from instance extensions for Gen2
func expandPlatformOptionsFromRCExtension(extensions map[string]interface{}) []map[string]interface{} {
	pltOption := make(map[string]interface{})

	// Initialize with empty strings
	pltOption["disk_encryption_key_crn"] = ""
	pltOption["backup_encryption_key_crn"] = ""

	if dataservices, ok := extensions["dataservices"].(map[string]interface{}); ok {
		if encryption, ok := dataservices["encryption"].(map[string]interface{}); ok {
			if disk, ok := encryption["disk"].(string); ok {
				pltOption["disk_encryption_key_crn"] = disk
			}
			if backup, ok := encryption["backup"].(string); ok {
				pltOption["backup_encryption_key_crn"] = backup
			}
		}
	}

	// Always return platform options with empty strings if keys are not found
	return []map[string]interface{}{pltOption}
}

// flattenIcdGroupsFromInstanceAndCatalog creates groups data from instance extensions and global catalog metadata for Gen2
func flattenIcdGroupsFromInstanceAndCatalog(instance map[string]interface{}, catalogResources []interface{}, resourceID string) []map[string]interface{} {
	groups := make([]map[string]interface{}, 0)

	// Get allocation values from instance extensions
	var memoryGB, cpuCount, storageGB float64
	var members int64
	var hostFlavorID string

	// Get the database type from resource ID
	dbType := getDatabaseTypeFromResourceID(resourceID)

	if dataservices, ok := instance["dataservices"].(map[string]interface{}); ok {
		// Access the specific database type data (e.g., "mongodb", "postgresql", etc.)
		if dbType != "" {
			if dbTypeData, ok := dataservices[dbType].(map[string]interface{}); ok {
				if mem, ok := dbTypeData["memory_gb"].(float64); ok {
					memoryGB = mem
				}
				if cpu, ok := dbTypeData["cpu_count"].(float64); ok {
					cpuCount = cpu
				}
				if storage, ok := dbTypeData["storage_gb"].(float64); ok {
					storageGB = storage
				}
				if m, ok := dbTypeData["members"].(float64); ok {
					members = int64(m)
				}
				if flavor, ok := dbTypeData["host_flavor"].(string); ok {
					hostFlavorID = flavor
				}
			}
		}
	}

	// Process catalog resources to get min/max/step values
	for _, resource := range catalogResources {
		resourceMap, ok := resource.(map[string]interface{})
		if !ok {
			continue
		}

		groupID, _ := resourceMap["id"].(string)

		// Use members from instance extensions if available, otherwise use count from catalog
		count := members
		if count == 0 {
			if c, ok := resourceMap["count"].(float64); ok {
				count = int64(c)
			}
		}

		// Memory
		memorys := make([]map[string]interface{}, 1)
		memory := make(map[string]interface{})
		if memoryData, ok := resourceMap["memory"].(map[string]interface{}); ok {
			memory["units"] = memoryData["units"]
			memory["allocation_mb"] = int64(memoryGB * 1024) // Convert GB to MB
			if minGB, ok := memoryData["minimum_gb"].(float64); ok {
				memory["minimum_mb"] = int64(minGB * 1024)
			}
			if stepGB, ok := memoryData["step_size_gb"].(float64); ok {
				memory["step_size_mb"] = int64(stepGB * 1024)
			}
			memory["is_adjustable"] = memoryData["is_adjustable"]
			memory["can_scale_down"] = memoryData["can_scale_down"]
		}
		memorys[0] = memory

		// CPU
		cpus := make([]map[string]interface{}, 1)
		cpu := make(map[string]interface{})
		if cpuData, ok := resourceMap["cpu"].(map[string]interface{}); ok {
			cpu["units"] = cpuData["units"]
			cpu["allocation_count"] = int64(cpuCount)
			cpu["minimum_count"] = cpuData["minimum_count"]
			cpu["step_size_count"] = cpuData["step_size_count"]
			cpu["is_adjustable"] = cpuData["is_adjustable"]
			cpu["can_scale_down"] = cpuData["can_scale_down"]
		}
		cpus[0] = cpu

		// Disk
		disks := make([]map[string]interface{}, 1)
		disk := make(map[string]interface{})
		if diskData, ok := resourceMap["disk"].(map[string]interface{}); ok {
			disk["units"] = diskData["units"]
			disk["allocation_mb"] = int64(storageGB * 1024) // Convert GB to MB
			if minGB, ok := diskData["minimum_gb"].(float64); ok {
				disk["minimum_mb"] = int64(minGB * 1024)
			}
			if stepGB, ok := diskData["step_size_gb"].(float64); ok {
				disk["step_size_mb"] = int64(stepGB * 1024)
			}
			disk["is_adjustable"] = diskData["is_adjustable"]
			disk["can_scale_down"] = diskData["can_scale_down"]
		}
		disks[0] = disk

		// Host Flavor
		hostflavors := make([]map[string]interface{}, 0)
		if hostFlavorID != "" {
			hostflavors = make([]map[string]interface{}, 1)
			hostflavor := make(map[string]interface{})
			hostflavor["id"] = hostFlavorID
			hostflavor["name"] = hostFlavorID
			hostflavor["hosting_size"] = "" // Not available in Gen2
			hostflavors[0] = hostflavor
		}

		group := map[string]interface{}{
			"group_id":    groupID,
			"count":       int64(count),
			"memory":      memorys,
			"cpu":         cpus,
			"disk":        disks,
			"host_flavor": hostflavors,
		}
		groups = append(groups, group)
	}

	return groups
}

// getInitialNodeCountGen2 gets the default member count for Gen2 plans from Global Catalog
func getInitialNodeCountGen2(deploymentID string, meta interface{}) (int, error) {
	globalClient, err := meta.(conns.ClientSession).GlobalCatalogV1API()
	if err != nil {
		return 0, err
	}

	// Get the deployment catalog entry
	options := globalcatalogv1.GetCatalogEntryOptions{
		ID: &deploymentID,
	}
	deployment, _, err := globalClient.GetCatalogEntry(&options)
	if err != nil {
		return 0, fmt.Errorf("[ERROR] Error retrieving deployment catalog entry: %s", err)
	}

	// Extract member count from deployment metadata resources
	if deployment.Metadata != nil && deployment.Metadata.Other != nil {
		if resources, ok := deployment.Metadata.Other["resources"].([]interface{}); ok {
			for _, resource := range resources {
				if resourceMap, ok := resource.(map[string]interface{}); ok {
					if groupID, ok := resourceMap["id"].(string); ok && groupID == "member" {
						if count, ok := resourceMap["count"].(float64); ok {
							return int(count), nil
						}
					}
				}
			}
		}
	}

	// Default to 3 if not found in metadata
	return 3, nil
}
