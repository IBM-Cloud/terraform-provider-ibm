// Copyright IBM Corp. 2017, 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"fmt"
	"log"
	"net/url"
	"regexp"
	"strings"
	"time"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/cloud-databases-go-sdk/clouddatabasesv5"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/globalcatalogv1"
	rc "github.com/IBM/platform-services-go-sdk/resourcecontrollerv2"
	rg "github.com/IBM/platform-services-go-sdk/resourcemanagerv2"
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

	// Instance states - shared across Classic and Gen2
	instanceStateRemoved               = "removed"
	databaseInstanceSuccessStatus      = "active"
	databaseInstanceProvisioningStatus = "provisioning"
	databaseInstanceProgressStatus     = "in progress"
	databaseInstanceInactiveStatus     = "inactive"
	databaseInstanceFailStatus         = "failed"
	databaseInstanceRemovedStatus      = "removed"

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

// extractLocationFromCRN extracts the location (region) from an IBM Cloud CRN.
// CRN format: crn:version:cname:ctype:service-name:location:scope:service-instance:resource-type:resource
// Returns the location field (index 5) or an error if the CRN is invalid.
func extractLocationFromCRN(crn *string) (string, error) {
	if crn == nil {
		return "", fmt.Errorf("CRN is nil")
	}
	parts := strings.Split(*crn, ":")
	if len(parts) <= 5 {
		return "", fmt.Errorf("invalid CRN format: expected at least 6 parts, got %d", len(parts))
	}
	return parts[5], nil
}

// wrapAPIError wraps an API error with operation context and response details.
// Provides consistent error formatting across API calls.
func wrapAPIError(operation string, err error, response interface{}) error {
	return fmt.Errorf("failed to %s: %w (response: %v)", operation, err, response)
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

// getGlobalCatalogClient initializes and returns the Global Catalog V1 client.
// Centralizes client initialization to reduce duplication and improve testability.
func getGlobalCatalogClient(meta interface{}) (*globalcatalogv1.GlobalCatalogV1, error) {
	client, err := meta.(conns.ClientSession).GlobalCatalogV1API()
	if err != nil {
		return nil, fmt.Errorf("failed to get global catalog client: %w", err)
	}
	return client, nil
}

// getResourceManagerClient initializes and returns the Resource Manager V2 client.
// Centralizes client initialization to reduce duplication and improve testability.
func getResourceManagerClient(meta interface{}) (interface{}, error) {
	client, err := meta.(conns.ClientSession).ResourceManagerV2API()
	if err != nil {
		return nil, fmt.Errorf("failed to get resource manager client: %w", err)
	}
	return client, nil
}

// setTagsWithLogging retrieves and sets tags for a resource, logging errors instead of failing.
// Returns error only for critical failures, logs warnings for non-critical issues.
func setTagsWithLogging(d *schema.ResourceData, crn string, meta interface{}) error {
	tags, err := flex.GetTagsUsingCRN(meta, crn)
	if err != nil {
		log.Printf("[WARN] Failed to retrieve tags for resource %s: %v", crn, err)
	}
	return d.Set("tags", tags)
}

// buildResourceControllerURL constructs the resource controller URL for a given CRN.
// Standardizes URL building across resources and data sources.
func buildResourceControllerURL(meta interface{}, crn string) (string, error) {
	rcontroller, err := flex.GetBaseController(meta)
	if err != nil {
		return "", fmt.Errorf("failed to get base controller: %w", err)
	}
	return rcontroller + "/services/" + url.QueryEscape(crn), nil
}

// setResourceControllerAttributes sets common flex resource controller attributes.
// Reduces duplication of setting name, CRN, status, and controller URL.
func setResourceControllerAttributes(d *schema.ResourceData, name, crn, state string, meta interface{}) error {
	d.Set(flex.ResourceName, name)
	d.Set(flex.ResourceCRN, crn)
	d.Set(flex.ResourceStatus, state)

	controllerURL, err := buildResourceControllerURL(meta, crn)
	if err != nil {
		return err
	}
	d.Set(flex.ResourceControllerURL, controllerURL)

	return nil
}

// setGen2BasicAttributes sets basic instance attributes including tags, name, status, location, and resource controller attributes.
// This function is shared between data source and resource implementations.
// Parameters:
//   - includeServiceEndpoints: if true, sets service_endpoints from instance parameters (resource only)
//   - includeResourceControllerURL: if true, sets resource_controller_url (resource only)
func setGen2BasicAttributes(d *schema.ResourceData, instance *rc.ResourceInstance, meta interface{}, includeServiceEndpoints, includeResourceControllerURL bool) error {
	// Retrieve and set tags (non-critical operation, log errors but continue)
	tags, err := flex.GetTagsUsingCRN(meta, *instance.CRN)
	if err != nil {
		log.Printf("[WARN] Error on get of ibm Database tags (%s) tags: %s", d.Id(), err)
	}
	d.Set("tags", tags)

	// Set basic instance attributes
	d.Set("name", instance.Name)
	d.Set("status", instance.State)
	d.Set("resource_group_id", instance.ResourceGroupID)
	d.Set("guid", instance.GUID)

	// Set location - try to extract from CRN first, fallback to RegionID
	var instanceLocation string
	if instance.CRN != nil {
		var err error
		instanceLocation, err = extractLocationFromCRN(instance.CRN)
		if err == nil {
			d.Set("location", instanceLocation)
		}
	}
	if instanceLocation == "" && instance.RegionID != nil {
		d.Set("location", instance.RegionID)
	}

	// Set service endpoints if requested (resource only)
	if includeServiceEndpoints && instance.Parameters != nil {
		if endpoint, ok := instance.Parameters["service_endpoints"]; ok {
			d.Set("service_endpoints", endpoint)
		}
	}

	// Set resource controller attributes
	d.Set(flex.ResourceName, instance.Name)
	d.Set(flex.ResourceCRN, instance.CRN)
	d.Set(flex.ResourceStatus, instance.State)

	// Retrieve and set resource group name
	rMgtClient, err := getResourceManagerClient(meta)
	if err != nil {
		return err
	}
	getResourceGroupOptions := rg.GetResourceGroupOptions{
		ID: instance.ResourceGroupID,
	}
	resourceGroup, resp, err := rMgtClient.(*rg.ResourceManagerV2).GetResourceGroup(&getResourceGroupOptions)
	if err != nil || resourceGroup == nil {
		log.Printf("[WARN] Failed to retrieve resource group: %v %v", err, resp)
	}
	if resourceGroup != nil && resourceGroup.Name != nil {
		d.Set(flex.ResourceGroupName, resourceGroup.Name)
	}

	// Set resource controller URL if requested (resource only)
	if includeResourceControllerURL {
		rcontroller, err := flex.GetBaseController(meta)
		if err != nil {
			return fmt.Errorf("failed to get base controller: %w", err)
		}
		d.Set(flex.ResourceControllerURL, rcontroller+"/services/"+url.QueryEscape(*instance.CRN))
	}

	return nil
}

// setGen2ServiceInfo retrieves and sets service and plan information from Global Catalog.
// Clears admin user attribute as it's not available in Gen2.
// This function is shared between data source and resource implementations.
func setGen2ServiceInfo(d *schema.ResourceData, instance *rc.ResourceInstance, meta interface{}) error {
	// Get global catalog client
	globalClient, err := getGlobalCatalogClient(meta)
	if err != nil {
		return err
	}

	// Get service offering details
	serviceOptions := globalcatalogv1.GetCatalogEntryOptions{
		ID: instance.ResourceID,
	}
	service, _, err := globalClient.GetCatalogEntry(&serviceOptions)
	if err != nil {
		return fmt.Errorf("failed to retrieve service offering: %w", err)
	}
	d.Set("service", service.Name)

	// Get plan details
	planOptions := globalcatalogv1.GetCatalogEntryOptions{
		ID: instance.ResourcePlanID,
	}
	plan, _, err := globalClient.GetCatalogEntry(&planOptions)
	if err != nil {
		return fmt.Errorf("failed to retrieve plan: %w", err)
	}
	d.Set("plan", plan.Name)

	// Clear Gen2-unsupported attributes to prevent stale Classic values
	// Admin user is not available in Gen2. Users should manage credentials using ibm_resource_key.
	d.Set(adminUserKey, nil)

	return nil
}

// setGen2VersionInfo extracts and sets version information from instance extensions.
// Also sets platform_options if includePlatformOptions is true (data source only).
// This function is shared between data source and resource implementations.
func setGen2VersionInfo(d *schema.ResourceData, instance *rc.ResourceInstance, includePlatformOptions bool) {
	// Extract version from instance.Extensions based on database type
	version := ""
	if instance.Extensions != nil && instance.ResourceID != nil {
		version = extractVersionFromExtensions(instance.Extensions, *instance.ResourceID)
	}
	d.Set(versionKey, version)

	// Extract platform_options from instance.Extensions for Gen2 (data source only)
	if includePlatformOptions && instance.Extensions != nil {
		d.Set(platformOptionsKey, expandPlatformOptionsFromRCExtension(instance.Extensions))
	}
}

// setGen2GroupsInfo retrieves and sets groups information from catalog.
// Combines instance extensions with catalog metadata to build group configurations.
// This function is shared between data source and resource implementations.
func setGen2GroupsInfo(d *schema.ResourceData, instance *rc.ResourceInstance, meta interface{}) error {
	// Extract location - try CRN first, fallback to RegionID
	var instanceLocation string
	if instance.CRN != nil {
		var err error
		instanceLocation, err = extractLocationFromCRN(instance.CRN)
		if err != nil && instance.RegionID != nil {
			instanceLocation = *instance.RegionID
		}
	} else if instance.RegionID != nil {
		instanceLocation = *instance.RegionID
	}

	if instanceLocation == "" {
		return fmt.Errorf("unable to determine instance location")
	}

	// Get global catalog client
	globalClient, err := getGlobalCatalogClient(meta)
	if err != nil {
		return err
	}

	// Get groups data from GlobalCatalog for Gen2
	// Find the deployment by getting plan's children and matching by location
	deployment, err := findDeploymentByLocation(globalClient, *instance.ResourcePlanID, instanceLocation)
	if err != nil {
		return err
	}

	// Extract resources from deployment metadata
	var catalogResources []interface{}
	if deployment.Metadata != nil && deployment.Metadata.Other != nil {
		if resources, ok := deployment.Metadata.Other[resourcesKey].([]interface{}); ok {
			catalogResources = resources
		}
	}

	// Flatten groups using instance extensions and catalog metadata
	if instance.Extensions != nil && len(catalogResources) > 0 && instance.ResourceID != nil {
		d.Set("groups", flattenIcdGroupsFromInstanceAndCatalog(instance.Extensions, catalogResources, *instance.ResourceID))
	}

	return nil
}

// clearGen2UnsupportedAttributes clears attributes not supported in Gen2.
// Sets allowlist, users, and configuration_schema to nil to prevent stale Classic values.
// Note: auto_scaling and logical_replication_slot are silently ignored but NOT cleared
// to avoid drift detection when users have these in their configuration.
// This function is shared between data source and resource implementations.
func clearGen2UnsupportedAttributes(d *schema.ResourceData) {
	// Allowlist is not supported in Gen2
	d.Set(allowlistKey, nil)

	// Users management is not supported in Gen2 (use ibm_resource_key instead)
	d.Set("users", nil)

	// Configuration schema is not supported in Gen2
	d.Set("configuration_schema", nil)
}
