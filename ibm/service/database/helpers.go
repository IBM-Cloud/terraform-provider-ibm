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

const (
	// Conversion constants
	mbPerGb = 1024

	// HTTP status codes
	httpNotFound = 404

	// Default values
	defaultGroupID     = "member"
	defaultMemberCount = 3

	// Gen2 database operation keys
	deploymentKind     = "deployment"
	dataservicesKey    = "dataservices"
	versionKey         = "version"
	resourcesKey       = "resources"
	platformOptionsKey = "platform_options"
	adminUserKey       = "adminuser"
	autoScalingKey     = "auto_scaling"
	allowlistKey       = "allowlist"
	databaseUserType   = "database"
)

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

// Database service name prefixes mapped to their type keys
var databaseServicePrefixes = map[string]string{
	"databases-for-etcd":          "etcd",
	"databases-for-postgresql":    "postgresql",
	"databases-for-redis":         "redis",
	"databases-for-elasticsearch": "elasticsearch",
	"databases-for-mongodb":       "mongodb",
	"messages-for-rabbitmq":       "rabbitmq",
	"databases-for-mysql":         "mysql",
	"databases-for-enterprisedb":  "enterprisedb",
}

// getDatabaseTypeFromResourceID maps the resource ID or service name to the database type key.
// Used in extensions for Gen2 and in parameters structure.
// Returns an empty string if the resource ID doesn't match any known database service.
func getDatabaseTypeFromResourceID(resourceID string) string {
	for prefix, dbType := range databaseServicePrefixes {
		if strings.HasPrefix(resourceID, prefix) {
			return dbType
		}
	}
	return ""
}

// expandPlatformOptionsFromRCExtension extracts platform options from instance extensions for Gen2.
// Returns a slice containing a single map with disk and backup encryption key CRNs.
// If encryption keys are not found in extensions, empty strings are returned.
func expandPlatformOptionsFromRCExtension(extensions map[string]interface{}) []map[string]interface{} {
	pltOption := map[string]interface{}{
		"disk_encryption_key_crn":   "",
		"backup_encryption_key_crn": "",
	}

	dataservices, ok := extensions["dataservices"].(map[string]interface{})
	if !ok {
		return []map[string]interface{}{pltOption}
	}

	encryption, ok := dataservices["encryption"].(map[string]interface{})
	if !ok {
		return []map[string]interface{}{pltOption}
	}

	if disk, ok := encryption["disk"].(string); ok && disk != "" {
		pltOption["disk_encryption_key_crn"] = disk
	}
	if backup, ok := encryption["backup"].(string); ok && backup != "" {
		pltOption["backup_encryption_key_crn"] = backup
	}

	return []map[string]interface{}{pltOption}
}

// flattenIcdGroupsFromInstanceAndCatalog creates groups data from instance extensions and global catalog metadata for Gen2.
// It combines actual allocation values from the instance with metadata constraints from the catalog.
// Returns a slice of group configurations including memory, CPU, disk, and host flavor information.
func flattenIcdGroupsFromInstanceAndCatalog(instance map[string]interface{}, catalogResources []interface{}, resourceID string) []map[string]interface{} {
	groups := make([]map[string]interface{}, 0, len(catalogResources))

	// Extract allocation values from instance extensions
	allocations := extractDatabaseAllocations(instance, resourceID)

	// Process catalog resources to build group configurations
	for _, resource := range catalogResources {
		resourceMap, ok := resource.(map[string]interface{})
		if !ok {
			continue
		}

		groupID, _ := resourceMap["id"].(string)

		// Use members from instance if available, otherwise use catalog count
		count := allocations.members
		if count == 0 {
			if c, ok := resourceMap["count"].(float64); ok {
				count = int64(c)
			}
		}

		group := map[string]interface{}{
			"group_id":    groupID,
			"count":       count,
			"memory":      buildMemoryConfig(resourceMap, allocations.memoryGB),
			"cpu":         buildCPUConfig(resourceMap, allocations.cpuCount),
			"disk":        buildDiskConfig(resourceMap, allocations.storageGB),
			"host_flavor": buildHostFlavorConfig(allocations.hostFlavorID),
		}
		groups = append(groups, group)
	}

	return groups
}

// databaseAllocations holds resource allocation values extracted from instance extensions.
// Fields are ordered by type for consistency.
type databaseAllocations struct {
	cpuCount     float64
	memoryGB     float64
	storageGB    float64
	members      int64
	hostFlavorID string
}

// extractDatabaseAllocations extracts allocation values from instance extensions for a specific database type
func extractDatabaseAllocations(instance map[string]interface{}, resourceID string) databaseAllocations {
	var alloc databaseAllocations

	dbType := getDatabaseTypeFromResourceID(resourceID)
	if dbType == "" {
		return alloc
	}

	dataservices, ok := instance["dataservices"].(map[string]interface{})
	if !ok {
		return alloc
	}

	dbTypeData, ok := dataservices[dbType].(map[string]interface{})
	if !ok {
		return alloc
	}

	if mem, ok := dbTypeData["memory_gb"].(float64); ok {
		alloc.memoryGB = mem
	}
	if cpu, ok := dbTypeData["cpu_count"].(float64); ok {
		alloc.cpuCount = cpu
	}
	if storage, ok := dbTypeData["storage_gb"].(float64); ok {
		alloc.storageGB = storage
	}
	if m, ok := dbTypeData["members"].(float64); ok {
		alloc.members = int64(m)
	}
	if flavor, ok := dbTypeData["host_flavor"].(string); ok {
		alloc.hostFlavorID = flavor
	}

	return alloc
}

// buildMemoryConfig creates memory configuration from catalog metadata and actual allocation
func buildMemoryConfig(resourceMap map[string]interface{}, memoryGB float64) []map[string]interface{} {
	memory := make(map[string]interface{})

	if memoryData, ok := resourceMap["memory"].(map[string]interface{}); ok {
		memory["units"] = memoryData["units"]
		memory["allocation_mb"] = int64(memoryGB * mbPerGb)
		if minGB, ok := memoryData["minimum_gb"].(float64); ok {
			memory["minimum_mb"] = int64(minGB * mbPerGb)
		}
		if stepGB, ok := memoryData["step_size_gb"].(float64); ok {
			memory["step_size_mb"] = int64(stepGB * mbPerGb)
		}
		memory["is_adjustable"] = memoryData["is_adjustable"]
		memory["can_scale_down"] = memoryData["can_scale_down"]
	}

	return []map[string]interface{}{memory}
}

// buildCPUConfig creates CPU configuration from catalog metadata and actual allocation
func buildCPUConfig(resourceMap map[string]interface{}, cpuCount float64) []map[string]interface{} {
	cpu := make(map[string]interface{})

	if cpuData, ok := resourceMap["cpu"].(map[string]interface{}); ok {
		cpu["units"] = cpuData["units"]
		cpu["allocation_count"] = int64(cpuCount)
		cpu["minimum_count"] = cpuData["minimum_count"]
		cpu["step_size_count"] = cpuData["step_size_count"]
		cpu["is_adjustable"] = cpuData["is_adjustable"]
		cpu["can_scale_down"] = cpuData["can_scale_down"]
	}

	return []map[string]interface{}{cpu}
}

// buildDiskConfig creates disk configuration from catalog metadata and actual allocation
func buildDiskConfig(resourceMap map[string]interface{}, storageGB float64) []map[string]interface{} {
	disk := make(map[string]interface{})

	if diskData, ok := resourceMap["disk"].(map[string]interface{}); ok {
		disk["units"] = diskData["units"]
		disk["allocation_mb"] = int64(storageGB * mbPerGb)
		if minGB, ok := diskData["minimum_gb"].(float64); ok {
			disk["minimum_mb"] = int64(minGB * mbPerGb)
		}
		if stepGB, ok := diskData["step_size_gb"].(float64); ok {
			disk["step_size_mb"] = int64(stepGB * mbPerGb)
		}
		disk["is_adjustable"] = diskData["is_adjustable"]
		disk["can_scale_down"] = diskData["can_scale_down"]
	}

	return []map[string]interface{}{disk}
}

// buildHostFlavorConfig creates host flavor configuration if a flavor ID is provided
func buildHostFlavorConfig(hostFlavorID string) []map[string]interface{} {
	if hostFlavorID == "" {
		return []map[string]interface{}{}
	}

	hostflavor := map[string]interface{}{
		"id":           hostFlavorID,
		"name":         hostFlavorID,
		"hosting_size": "", // Not available in Gen2
	}

	return []map[string]interface{}{hostflavor}
}

// getInitialNodeCountGen2 retrieves the default member count for Gen2 plans from Global Catalog.
// Returns the member count from the catalog metadata, or a default value of 3 if not found.
func getInitialNodeCountGen2(deploymentID string, meta interface{}) (int, error) {
	globalClient, err := meta.(conns.ClientSession).GlobalCatalogV1API()
	if err != nil {
		return 0, fmt.Errorf("failed to get global catalog client: %w", err)
	}

	options := &globalcatalogv1.GetCatalogEntryOptions{
		ID: &deploymentID,
	}

	deployment, _, err := globalClient.GetCatalogEntry(options)
	if err != nil {
		return 0, fmt.Errorf("error retrieving deployment catalog entry: %w", err)
	}

	// Extract member count from deployment metadata
	count := extractMemberCountFromMetadata(deployment)
	if count > 0 {
		return count, nil
	}

	// Return default if not found in metadata
	return defaultMemberCount, nil
}

// extractMemberCountFromMetadata extracts the member count from catalog entry metadata
func extractMemberCountFromMetadata(deployment *globalcatalogv1.CatalogEntry) int {
	if deployment.Metadata == nil || deployment.Metadata.Other == nil {
		return 0
	}

	resources, ok := deployment.Metadata.Other["resources"].([]interface{})
	if !ok {
		return 0
	}

	for _, resource := range resources {
		resourceMap, ok := resource.(map[string]interface{})
		if !ok {
			continue
		}

		groupID, ok := resourceMap["id"].(string)
		if !ok || groupID != "member" {
			continue
		}

		count, ok := resourceMap["count"].(float64)
		if ok && count > 0 {
			return int(count)
		}
	}

	return 0
}

// extractVersionFromExtensions extracts the database version from instance extensions.
// Returns an empty string if the version cannot be found.
func extractVersionFromExtensions(extensions map[string]interface{}, resourceID string) string {
	if extensions == nil {
		return ""
	}

	dbType := getDatabaseTypeFromResourceID(resourceID)
	if dbType == "" {
		return ""
	}

	dataservices, ok := extensions[dataservicesKey].(map[string]interface{})
	if !ok {
		return ""
	}

	dbTypeData, ok := dataservices[dbType].(map[string]interface{})
	if !ok {
		return ""
	}

	version, ok := dbTypeData[versionKey].(string)
	if !ok {
		return ""
	}

	return version
}

// findDeploymentByLocation finds a deployment catalog entry matching the specified location.
// Returns the deployment entry or an error if not found.
func findDeploymentByLocation(globalClient *globalcatalogv1.GlobalCatalogV1, planID string, location string) (*globalcatalogv1.CatalogEntry, error) {
	if globalClient == nil {
		return nil, fmt.Errorf("global catalog client is nil")
	}

	kind := deploymentKind
	childOptions := globalcatalogv1.GetChildObjectsOptions{
		ID:   &planID,
		Kind: &kind,
	}

	children, _, err := globalClient.GetChildObjects(&childOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve plan children: %w", err)
	}

	if children == nil || children.Resources == nil {
		return nil, fmt.Errorf("no deployments found for plan")
	}

	for _, child := range children.Resources {
		if child.Metadata != nil &&
			child.Metadata.Deployment != nil &&
			child.Metadata.Deployment.Location != nil &&
			*child.Metadata.Deployment.Location == location {
			return &child, nil
		}
	}

	return nil, fmt.Errorf("could not find deployment catalog entry for region %s", location)
}
