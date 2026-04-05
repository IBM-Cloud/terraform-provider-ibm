// Copyright IBM Corp. 2017, 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"

	"github.com/IBM-Cloud/bluemix-go/models"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/cloud-databases-go-sdk/clouddatabasesv5"
	"github.com/IBM/go-sdk-core/v5/core"
	rc "github.com/IBM/platform-services-go-sdk/resourcecontrollerv2"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var gen2UnsupportedAttrs = []string{
	"backup_policy",
	"users",
	"auto_scaling",
	"allowlist",
	"configuration_schema",
	"logical_replication_slot",
}

// DBConfig represents database-specific configuration for Gen2 parameters.
// Replaces map[string]interface{} for type safety and compile-time validation.
type DBConfig struct {
	Version    string `json:"version,omitempty"`
	Members    int    `json:"members"`
	StorageGB  int    `json:"storage_gb,omitempty"`
	HostFlavor string `json:"host_flavor,omitempty"`
}

// instanceConfigContext encapsulates shared context for instance configuration steps.
// This reduces parameter passing and makes the configuration flow more maintainable.
type instanceConfigContext struct {
	ctx                  context.Context
	d                    *schema.ResourceData
	instanceID           string
	cloudDatabasesClient *clouddatabasesv5.CloudDatabasesV5
	meta                 interface{}
	instance             *rc.ResourceInstance
}

type resourceIBMDatabaseGen2Backend struct{}

// newResourceIBMDatabaseGen2Backend creates a new Gen2 backend instance
func newResourceIBMDatabaseGen2Backend() resourceIBMDatabaseBackend {
	return &resourceIBMDatabaseGen2Backend{}
}

// wrapErr wraps an error with a context string for better error messages.
// Reduces duplication of fmt.Errorf calls throughout the codebase.
func wrapErr(ctx string, err error) error {
	return fmt.Errorf("%s: %w", ctx, err)
}

// emptyStringsWithErr returns two empty strings and an error.
// Helper to reduce duplication of `return "", "", fmt.Errorf(...)` statements.
func emptyStringsWithErr(format string, args ...interface{}) (string, string, error) {
	return "", "", fmt.Errorf(format, args...)
}

// Create provisions a new IBM Cloud Database Gen2 instance.
// It handles resource creation, scaling configuration, encryption setup,
// and post-provisioning tasks like password updates and allowlist configuration.
//
// Parameters:
//   - ctx: Context for cancellation and timeouts
//   - d: Terraform resource data containing configuration
//   - meta: Provider metadata with API clients
//
// Returns:
//   - diag.Diagnostics: Any errors or warnings encountered
func (g *resourceIBMDatabaseGen2Backend) Create(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// Create the resource instance
	instance, err := g.createResourceInstance(ctx, d, meta)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(*instance.ID)

	// Wait for instance creation to complete
	_, err = waitForDatabaseInstanceCreate(d, meta, *instance.ID, false)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error waiting for create database instance (%s) to complete: %w", *instance.ID, err))
	}

	// Configure the instance with additional settings
	if err := g.configureInstance(ctx, d, meta, instance); err != nil {
		return diag.FromErr(err)
	}

	return resourceIBMDatabaseInstanceRead(ctx, d, meta)
}

// createResourceInstance handles the initial resource instance creation.
// It retrieves service and plan information, builds Gen2 parameters, and creates the instance.
func (g *resourceIBMDatabaseGen2Backend) createResourceInstance(ctx context.Context, d *schema.ResourceData, meta interface{}) (*rc.ResourceInstance, error) {
	clientSession := meta.(conns.ClientSession)
	rsConClient, err := clientSession.ResourceControllerV2API()
	if err != nil {
		return nil, fmt.Errorf("failed to initialize resource controller client: %w", err)
	}

	serviceName := d.Get("service").(string)
	plan := d.Get("plan").(string)
	name := d.Get("name").(string)
	location := d.Get("location").(string)

	rsInst := rc.CreateResourceInstanceOptions{
		Name: &name,
	}

	// Get service offering and plan
	servicePlan, catalogCRN, err := g.getServicePlanAndCatalog(serviceName, plan, location, clientSession)
	if err != nil {
		return nil, err
	}

	rsInst.ResourcePlanID = &servicePlan
	rsInst.Target = &catalogCRN

	// Set resource group
	if err := g.setResourceGroup(d, meta, &rsInst); err != nil {
		return nil, err
	}

	// Build Gen2 parameters
	parameters, err := g.buildGen2Parameters(d, serviceName, meta, catalogCRN)
	if err != nil {
		return nil, err
	}
	rsInst.Parameters = parameters

	// Create the instance with retry logic
	instance, response, err := g.createInstanceWithRetry(ctx, rsConClient, &rsInst)
	if err != nil {
		return nil, fmt.Errorf("error creating database instance: %w (response: %v)", err, response)
	}

	return instance, nil
}

// getServicePlanAndCatalog retrieves the service plan ID and catalog CRN.
// It validates that the plan is available in the specified location.
func (g *resourceIBMDatabaseGen2Backend) getServicePlanAndCatalog(serviceName, plan, location string, meta conns.ClientSession) (string, string, error) {
	rsCatClient, err := meta.ResourceCatalogAPI()
	if err != nil {
		return emptyStringsWithErr("failed to initialize resource catalog client: %w", err)
	}
	rsCatRepo := rsCatClient.ResourceCatalog()

	serviceOff, err := rsCatRepo.FindByName(serviceName, true)
	if err != nil {
		return emptyStringsWithErr("error retrieving database service offering: %w", err)
	}

	servicePlan, err := rsCatRepo.GetServicePlanID(serviceOff[0], plan)
	if err != nil {
		return emptyStringsWithErr("error retrieving plan: %w", err)
	}

	// Check special case before calling ListDeployments to avoid unnecessary API call
	if serviceName == "databases-for-mongodb" && plan == "enterprise-sharding" {
		return emptyStringsWithErr("%s %s is not available yet in this region", serviceName, plan)
	}

	deployments, err := rsCatRepo.ListDeployments(servicePlan)
	if err != nil {
		return emptyStringsWithErr("error retrieving deployment for plan %s: %w", plan, err)
	}

	if len(deployments) == 0 {
		return emptyStringsWithErr("no deployment found for service plan: %s", plan)
	}

	// Filter and validate deployment location
	catalogCRN, err := g.validateAndGetCatalogCRN(deployments, location, plan)
	if err != nil {
		return emptyStringsWithErr("%v", err)
	}

	return servicePlan, catalogCRN, nil
}

// validateAndGetCatalogCRN filters deployments by location and returns the catalog CRN.
// Extracted to reduce nesting and improve readability of getServicePlanAndCatalog.
func (g *resourceIBMDatabaseGen2Backend) validateAndGetCatalogCRN(deployments []models.ServiceDeployment, location, plan string) (string, error) {
	filtered, supportedLocations := filterDatabaseDeployments(deployments, location)

	if len(filtered) == 0 {
		// Convert map keys to slice for strings.Join
		locations := make([]string, 0, len(supportedLocations))
		for l := range supportedLocations {
			locations = append(locations, l)
		}
		locationList := strings.Join(locations, ", ")
		return "", fmt.Errorf("no deployment found for service plan %s at location %s. Valid location(s) are: %s",
			plan, location, locationList)
	}

	return filtered[0].CatalogCRN, nil
}

// setResourceGroup sets the resource group for the instance.
// Uses the configured resource group or defaults to the account's default resource group.
func (g *resourceIBMDatabaseGen2Backend) setResourceGroup(d *schema.ResourceData, meta interface{}, rsInst *rc.CreateResourceInstanceOptions) error {
	if rsGrpID, ok := d.GetOk("resource_group_id"); ok {
		rgID := rsGrpID.(string)
		rsInst.ResourceGroup = &rgID
	} else {
		defaultRg, err := flex.DefaultResourceGroup(meta)
		if err != nil {
			return fmt.Errorf("failed to get default resource group: %w", err)
		}
		rsInst.ResourceGroup = &defaultRg
	}
	return nil
}

// buildGen2Parameters constructs the Gen2-specific parameters structure.
// Includes database configuration, encryption, restore, and PITR settings.
func (g *resourceIBMDatabaseGen2Backend) buildGen2Parameters(d *schema.ResourceData, serviceName string, meta interface{}, catalogCRN string) (map[string]interface{}, error) {
	// Get the database type for the dataservices key
	dbType := getDatabaseTypeFromResourceID(serviceName)
	if dbType == "" {
		return nil, fmt.Errorf("unable to determine database type from service name: %s", serviceName)
	}

	// Build database configuration using typed struct
	dbConfig, err := g.buildDBConfig(d, catalogCRN, meta)
	if err != nil {
		return nil, err
	}

	// Build dataservices structure
	dataservices := map[string]interface{}{
		dbType: dbConfig,
	}

	// Handle encryption
	g.addEncryptionConfig(d, dataservices)

	// Handle restore from backup
	g.addRestoreConfig(d, dataservices)

	// Handle point-in-time recovery
	g.addPITRConfig(d, dataservices)

	// Handle read replica
	if remoteLeader, ok := d.GetOk("remote_leader_id"); ok {
		dataservices["remote_leader_id"] = remoteLeader.(string)
	}

	// Build final parameters structure
	parameters := map[string]interface{}{
		"dataservices": dataservices,
	}

	return parameters, nil
}

// buildDBConfig creates database configuration with member group and storage settings.
// Extracts and consolidates member group logic, reducing nested if statements.
func (g *resourceIBMDatabaseGen2Backend) buildDBConfig(d *schema.ResourceData, catalogCRN string, meta interface{}) (map[string]interface{}, error) {
	config := DBConfig{}

	// Version
	if version, ok := d.GetOk("version"); ok {
		config.Version = version.(string)
	}

	// Get member group configuration
	memberGroup := g.getMemberGroup(d)

	// Members count
	members, err := g.getMembersCount(d, memberGroup, catalogCRN, meta)
	if err != nil {
		return nil, err
	}
	config.Members = members

	// Early return if no member group - simplifies logic below
	if memberGroup == nil {
		return g.dbConfigToMap(config), nil
	}

	// Storage in GB (not MB!) - guard clause eliminates nested if
	if memberGroup.Disk != nil {
		// Disk allocation is per member in MB, convert to GB for total
		storageGB := (memberGroup.Disk.Allocation * members) / mbPerGb
		config.StorageGB = storageGB
	}

	// Host flavor - guard clause eliminates nested if
	if memberGroup.HostFlavor != nil {
		config.HostFlavor = memberGroup.HostFlavor.ID
	}

	return g.dbConfigToMap(config), nil
}

// dbConfigToMap converts DBConfig struct to map[string]interface{} for API compatibility.
// Only includes non-zero values to avoid sending unnecessary fields.
func (g *resourceIBMDatabaseGen2Backend) dbConfigToMap(config DBConfig) map[string]interface{} {
	result := make(map[string]interface{})

	if config.Version != "" {
		result["version"] = config.Version
	}
	result["members"] = config.Members
	if config.StorageGB > 0 {
		result["storage_gb"] = config.StorageGB
	}
	if config.HostFlavor != "" {
		result["host_flavor"] = config.HostFlavor
	}

	return result
}

// getMemberGroup extracts the member group configuration from schema.
// Returns the group with ID "member" or nil if not found.
func (g *resourceIBMDatabaseGen2Backend) getMemberGroup(d *schema.ResourceData) *Group {
	if group, ok := d.GetOk("group"); ok {
		groups := expandGroups(group.(*schema.Set).List())
		for _, grp := range groups {
			if grp.ID == defaultGroupID {
				return grp
			}
		}
	}
	return nil
}

// getMembersCount determines the number of members for the instance.
// Uses the configured member count or retrieves the default from the catalog.
func (g *resourceIBMDatabaseGen2Backend) getMembersCount(d *schema.ResourceData, memberGroup *Group, catalogCRN string, meta interface{}) (int, error) {
	if memberGroup != nil && memberGroup.Members != nil {
		return memberGroup.Members.Allocation, nil
	}

	// Get initial node count from catalog
	members, err := getInitialNodeCountGen2(catalogCRN, meta)
	if err != nil {
		return 0, fmt.Errorf("failed to get initial node count: %w", err)
	}
	return members, nil
}

// addEncryptionConfig adds encryption configuration to dataservices.
// Includes disk and backup encryption key CRNs if configured.
func (g *resourceIBMDatabaseGen2Backend) addEncryptionConfig(d *schema.ResourceData, dataservices map[string]interface{}) {
	encryption := make(map[string]interface{}, 2)
	if keyProtect, ok := d.GetOk("key_protect_key"); ok {
		encryption["disk"] = keyProtect.(string)
	}
	if backUpEncryptionKey, ok := d.GetOk("backup_encryption_key_crn"); ok {
		encryption["backup"] = backUpEncryptionKey.(string)
	}
	if len(encryption) > 0 {
		dataservices["encryption"] = encryption
	}
}

// addRestoreConfig adds restore configuration to dataservices.
// Includes backup ID and restore mode settings if configured.
func (g *resourceIBMDatabaseGen2Backend) addRestoreConfig(d *schema.ResourceData, dataservices map[string]interface{}) {
	if backupID, ok := d.GetOk("backup_id"); ok {
		dataservices["restore_backup_id"] = backupID.(string)
	}

	if offlineRestore, ok := d.GetOk("offline_restore"); ok {
		dataservices["offline_restore"] = offlineRestore.(bool)
	}

	if asyncRestore, ok := d.GetOk("async_restore"); ok {
		dataservices["async_restore"] = asyncRestore.(bool)
	}
}

// addPITRConfig adds point-in-time recovery configuration to dataservices.
// Includes deployment ID and recovery time if configured.
func (g *resourceIBMDatabaseGen2Backend) addPITRConfig(d *schema.ResourceData, dataservices map[string]interface{}) {
	if pitrID, ok := d.GetOk("point_in_time_recovery_deployment_id"); ok {
		dataservices["point_in_time_recovery_deployment_id"] = pitrID.(string)
	}

	pitrOk := !d.GetRawConfig().AsValueMap()["point_in_time_recovery_time"].IsNull()
	if pitrTime, ok := d.GetOk("point_in_time_recovery_time"); pitrOk {
		if !ok {
			pitrTime = ""
		}
		pitrTimeTrimmed := strings.TrimSpace(pitrTime.(string))
		dataservices["point_in_time_recovery_time"] = pitrTimeTrimmed
	}
}

// createInstanceWithRetry creates an instance.
// Note: Retry logic can be added in the future if needed.
func (g *resourceIBMDatabaseGen2Backend) createInstanceWithRetry(ctx context.Context, client *rc.ResourceControllerV2, opts *rc.CreateResourceInstanceOptions) (*rc.ResourceInstance, *core.DetailedResponse, error) {
	instance, response, err := client.CreateResourceInstance(opts)
	return instance, response, err
}

// configureInstance applies post-creation configuration to the instance.
// Includes scaling, tags, passwords, allowlist, auto-scaling, users, and database settings.
func (g *resourceIBMDatabaseGen2Backend) configureInstance(ctx context.Context, d *schema.ResourceData, meta interface{}, instance *rc.ResourceInstance) error {
	// Initialize configuration context
	configCtx, err := g.initConfigContext(ctx, d, meta, instance)
	if err != nil {
		return err
	}

	// Define configuration steps in order of execution
	type configStep struct {
		name string
		fn   func(*instanceConfigContext) error
	}

	configSteps := []configStep{
		{name: "group scaling", fn: g.applyGroupScaling},
		{name: "tags", fn: g.updateTags},
		{name: "admin password", fn: g.updateAdminPassword},
		{name: "allowlist", fn: g.configureAllowlist},
		{name: "auto-scaling", fn: g.configureAutoScaling},
		{name: "users", fn: g.configureUsers},
		{name: "database settings", fn: g.configureDatabaseSettings},
		{name: "logical replication", fn: g.configureLogicalReplication},
	}

	// Execute configuration steps sequentially
	for _, step := range configSteps {
		if err := step.fn(configCtx); err != nil {
			return fmt.Errorf("failed to configure %s: %w", step.name, err)
		}
	}

	return nil
}

// initConfigContext initializes the configuration context with validated instance and client.
func (g *resourceIBMDatabaseGen2Backend) initConfigContext(ctx context.Context, d *schema.ResourceData, meta interface{}, instance *rc.ResourceInstance) (*instanceConfigContext, error) {
	if instance == nil || instance.ID == nil {
		return nil, fmt.Errorf("instance or instance ID is nil")
	}

	clientSession := meta.(conns.ClientSession)
	cloudDatabasesClient, err := clientSession.CloudDatabasesV5()
	if err != nil {
		return nil, fmt.Errorf("failed to initialize cloud databases client: %w", err)
	}

	return &instanceConfigContext{
		ctx:                  ctx,
		d:                    d,
		instanceID:           *instance.ID,
		cloudDatabasesClient: cloudDatabasesClient,
		meta:                 meta,
		instance:             instance,
	}, nil
}

// applyGroupScaling applies scaling configuration to instance groups.
// Compares desired configuration with current state and applies changes as needed.
func (g *resourceIBMDatabaseGen2Backend) applyGroupScaling(configCtx *instanceConfigContext) error {
	group, ok := configCtx.d.GetOk("group")
	if !ok {
		return nil
	}

	groups := expandGroups(group.(*schema.Set).List())
	groupsResponse, err := getGroups(configCtx.instanceID, configCtx.meta)
	if err != nil {
		return fmt.Errorf("failed to get groups: %w", err)
	}
	currentGroups := normalizeGroups(groupsResponse)

	for _, grp := range groups {
		if err := g.scaleGroup(configCtx, grp, currentGroups); err != nil {
			return err
		}
	}

	return nil
}

// scaleGroup scales a specific group if needed.
// buildGroupScaling constructs scaling configuration and returns whether scaling is needed.
// It compares the desired group configuration with the current state and builds the appropriate
// scaling request. The nodeCount parameter is updated if member allocation changes.
func (g *resourceIBMDatabaseGen2Backend) buildGroupScaling(grp, currentGroup *Group, nodeCount int) (*clouddatabasesv5.GroupScaling, bool) {
	groupScaling := &clouddatabasesv5.GroupScaling{}
	needsScaling := false

	if grp.Members != nil && grp.Members.Allocation != currentGroup.Members.Allocation {
		groupScaling.Members = &clouddatabasesv5.GroupScalingMembers{
			AllocationCount: core.Int64Ptr(int64(grp.Members.Allocation)),
		}
		nodeCount = grp.Members.Allocation
		needsScaling = true
	}

	if grp.Memory != nil && grp.Memory.Allocation*nodeCount != currentGroup.Memory.Allocation {
		groupScaling.Memory = &clouddatabasesv5.GroupScalingMemory{
			AllocationMb: core.Int64Ptr(int64(grp.Memory.Allocation * nodeCount)),
		}
		needsScaling = true
	}

	if grp.Disk != nil && grp.Disk.Allocation*nodeCount != currentGroup.Disk.Allocation {
		groupScaling.Disk = &clouddatabasesv5.GroupScalingDisk{
			AllocationMb: core.Int64Ptr(int64(grp.Disk.Allocation * nodeCount)),
		}
		needsScaling = true
	}

	if grp.CPU != nil && grp.CPU.Allocation*nodeCount != currentGroup.CPU.Allocation {
		groupScaling.CPU = &clouddatabasesv5.GroupScalingCPU{
			AllocationCount: core.Int64Ptr(int64(grp.CPU.Allocation * nodeCount)),
		}
		needsScaling = true
	}

	if grp.HostFlavor != nil {
		groupScaling.HostFlavor = &clouddatabasesv5.GroupScalingHostFlavor{
			ID: core.StringPtr(grp.HostFlavor.ID),
		}
		needsScaling = true
	}

	return groupScaling, needsScaling
}

// findGroupByID locates a group in the slice by ID using a map for O(1) lookup.
// Returns the group and true if found, nil and false otherwise.
func findGroupByID(groups []Group, id string) (*Group, bool) {
	groupMap := make(map[string]*Group, len(groups))
	for i := range groups {
		groupMap[groups[i].ID] = &groups[i]
	}
	group, exists := groupMap[id]
	return group, exists
}

// executeGroupScaling sends the scaling request to the API and returns the task ID.
func (g *resourceIBMDatabaseGen2Backend) executeGroupScaling(configCtx *instanceConfigContext, groupID string, groupScaling *clouddatabasesv5.GroupScaling) (string, error) {
	opts := &clouddatabasesv5.SetDeploymentScalingGroupOptions{
		ID:      &configCtx.instanceID,
		GroupID: &groupID,
		Group:   groupScaling,
	}

	resp, _, err := configCtx.cloudDatabasesClient.SetDeploymentScalingGroup(opts)
	if err != nil {
		return "", fmt.Errorf("failed to set deployment scaling group: %w", err)
	}

	return *resp.Task.ID, nil
}

// waitForScalingTaskComplete waits for the scaling task to complete.
func (g *resourceIBMDatabaseGen2Backend) waitForScalingTaskComplete(configCtx *instanceConfigContext, taskID string) error {
	_, err := waitForDatabaseTaskComplete(taskID, configCtx.d, configCtx.meta, configCtx.d.Timeout(schema.TimeoutUpdate))
	if err != nil {
		return fmt.Errorf("error waiting for scaling task to complete: %w", err)
	}
	return nil
}

// scaleGroup handles horizontal and vertical scaling for members, memory, disk, CPU, and host flavor.
// It compares the desired group configuration with the current state and applies changes as needed.
func (g *resourceIBMDatabaseGen2Backend) scaleGroup(configCtx *instanceConfigContext, grp *Group, currentGroups []Group) error {
	currentGroup, exists := findGroupByID(currentGroups, grp.ID)
	if !exists {
		return fmt.Errorf("current group %s not found", grp.ID)
	}

	nodeCount := currentGroup.Members.Allocation

	// Skip scaling if this is the default group with no member changes
	if grp.ID == defaultGroupID && (grp.Members == nil || grp.Members.Allocation == nodeCount) {
		return nil
	}

	// Build scaling configuration and check if any changes are needed
	groupScaling, needsScaling := g.buildGroupScaling(grp, currentGroup, nodeCount)
	if !needsScaling {
		return nil
	}

	// Apply scaling changes
	taskID, err := g.executeGroupScaling(configCtx, grp.ID, groupScaling)
	if err != nil {
		return err
	}

	return g.waitForScalingTaskComplete(configCtx, taskID)
}

// updateTags updates resource tags.
// Compares old and new tags and applies changes using the CRN.
func (g *resourceIBMDatabaseGen2Backend) updateTags(configCtx *instanceConfigContext) error {
	v := os.Getenv("IC_ENV_TAGS")
	if _, ok := configCtx.d.GetOk("tags"); ok || v != "" {
		oldList, newList := configCtx.d.GetChange("tags")
		err := flex.UpdateTagsUsingCRN(oldList, newList, configCtx.meta, *configCtx.instance.CRN)
		if err != nil {
			return fmt.Errorf("failed to update tags: %w", err)
		}
	}
	return nil
}

// updateAdminPassword updates the admin password if provided.
// getAdminUsername retrieves the admin username from deployment info.
// Extracts and returns the database admin username for the specified instance.
func (g *resourceIBMDatabaseGen2Backend) getAdminUsername(configCtx *instanceConfigContext) (string, error) {
	getDeploymentInfoOptions := &clouddatabasesv5.GetDeploymentInfoOptions{
		ID: core.StringPtr(configCtx.instanceID),
	}

	response, httpResp, err := configCtx.cloudDatabasesClient.GetDeploymentInfoWithContext(configCtx.ctx, getDeploymentInfoOptions)
	if err != nil {
		if httpResp != nil && httpResp.StatusCode == httpNotFound {
			return "", fmt.Errorf("database instance not found in configured region: %w", err)
		}
		return "", fmt.Errorf("failed to get deployment info: %w", err)
	}

	if response == nil || response.Deployment == nil {
		return "", fmt.Errorf("deployment info response is nil")
	}

	return response.Deployment.AdminUsernames[databaseUserType], nil
}

// waitForTaskCompletion waits for a database task to complete.
// Provides consistent task waiting with proper timeout handling.
func (g *resourceIBMDatabaseGen2Backend) waitForTaskCompletion(taskID string, configCtx *instanceConfigContext, timeoutType string) error {
	_, err := waitForDatabaseTaskComplete(taskID, configCtx.d, configCtx.meta, configCtx.d.Timeout(timeoutType))
	if err != nil {
		return fmt.Errorf("task %s failed: %w", taskID, err)
	}
	return nil
}

// updateAdminPassword updates the admin password if provided.
// Retrieves the admin username from deployment info and updates the password.
func (g *resourceIBMDatabaseGen2Backend) updateAdminPassword(configCtx *instanceConfigContext) error {
	pw, ok := configCtx.d.GetOk("adminpassword")
	if !ok {
		return nil
	}

	adminPassword, ok := pw.(string)
	if !ok {
		return fmt.Errorf("adminpassword must be a string, got %T", pw)
	}

	adminUser, err := g.getAdminUsername(configCtx)
	if err != nil {
		return err
	}

	updateUserOptions := &clouddatabasesv5.UpdateUserOptions{
		ID:       core.StringPtr(configCtx.instanceID),
		UserType: core.StringPtr(databaseUserType),
		Username: core.StringPtr(adminUser),
		User: &clouddatabasesv5.UserUpdatePasswordSetting{
			Password: &adminPassword,
		},
	}

	updateUserResponse, response, err := configCtx.cloudDatabasesClient.UpdateUserWithContext(configCtx.ctx, updateUserOptions)
	if err != nil {
		return fmt.Errorf("UpdateUser (%s) failed: %w (response: %v)", *updateUserOptions.Username, err, response)
	}

	return g.waitForTaskCompletion(*updateUserResponse.Task.ID, configCtx, schema.TimeoutCreate)
}

// configureAllowlist configures the IP allowlist.
// Sets the list of allowed IP addresses for database access.
func (g *resourceIBMDatabaseGen2Backend) configureAllowlist(configCtx *instanceConfigContext) error {
	_, hasAllowlist := configCtx.d.GetOk("allowlist")
	if !hasAllowlist {
		return nil
	}

	ipAddresses := configCtx.d.Get("allowlist").(*schema.Set)
	entries := flex.ExpandAllowlist(ipAddresses)

	setAllowlistOptions := &clouddatabasesv5.SetAllowlistOptions{
		ID:          &configCtx.instanceID,
		IPAddresses: entries,
	}

	setAllowlistResponse, _, err := configCtx.cloudDatabasesClient.SetAllowlist(setAllowlistOptions)
	if err != nil {
		return fmt.Errorf("error updating database allowlists: %w", err)
	}

	taskId := *setAllowlistResponse.Task.ID

	_, err = waitForDatabaseTaskComplete(taskId, configCtx.d, configCtx.meta, configCtx.d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("error waiting for update of database (%s) allowlist task to complete: %w", configCtx.instanceID, err)
	}

	return nil
}

// configureAutoScaling configures auto-scaling settings.
// Sets disk and memory auto-scaling conditions if configured.
func (g *resourceIBMDatabaseGen2Backend) configureAutoScaling(configCtx *instanceConfigContext) error {
	autoScaling := configCtx.d.Get("auto_scaling").([]interface{})
	if len(autoScaling) == 0 {
		return nil
	}

	autoscalingSetGroupAutoscaling := &clouddatabasesv5.AutoscalingSetGroupAutoscaling{}

	if diskRecord, ok := configCtx.d.GetOk("auto_scaling.0.disk"); ok {
		diskGroup, err := expandAutoscalingDiskGroup(configCtx.d, diskRecord)
		if err != nil {
			return fmt.Errorf("error expanding disk autoscaling: %w", err)
		}
		autoscalingSetGroupAutoscaling.Disk = diskGroup
	}

	if memoryRecord, ok := configCtx.d.GetOk("auto_scaling.0.memory"); ok {
		memoryGroup, err := expandAutoscalingMemoryGroup(configCtx.d, memoryRecord)
		if err != nil {
			return fmt.Errorf("error expanding memory autoscaling: %w", err)
		}
		autoscalingSetGroupAutoscaling.Memory = memoryGroup
	}

	if autoscalingSetGroupAutoscaling.Disk != nil || autoscalingSetGroupAutoscaling.Memory != nil {
		setAutoscalingConditionsOptions := &clouddatabasesv5.SetAutoscalingConditionsOptions{
			ID:          &configCtx.instanceID,
			GroupID:     core.StringPtr(defaultGroupID),
			Autoscaling: autoscalingSetGroupAutoscaling,
		}

		setAutoscalingConditionsResponse, _, err := configCtx.cloudDatabasesClient.SetAutoscalingConditionsWithContext(configCtx.ctx, setAutoscalingConditionsOptions)
		if err != nil {
			return fmt.Errorf("error updating database auto_scaling: %w", err)
		}

		if err := g.waitForTaskCompletion(*setAutoscalingConditionsResponse.Task.ID, configCtx, schema.TimeoutCreate); err != nil {
			return fmt.Errorf("error waiting for database (%s) auto_scaling update: %w", configCtx.instanceID, err)
		}
	}

	return nil
}

// configureUsers configures database users.
// Attempts to update existing users or creates new ones if they don't exist.
func (g *resourceIBMDatabaseGen2Backend) configureUsers(configCtx *instanceConfigContext) error {
	userList, ok := configCtx.d.GetOk("users")
	if !ok {
		return nil
	}

	users := expandUsers(userList.(*schema.Set).List())
	for _, user := range users {
		// Note: Some db users exist after provisioning (i.e. admin, repl)
		// so we must attempt both methods
		err := user.Update(configCtx.instanceID, configCtx.d, configCtx.meta)

		if err != nil {
			err = user.Create(configCtx.instanceID, configCtx.d, configCtx.meta)
		}

		if err != nil {
			return fmt.Errorf("error configuring user %s: %w", user.Username, err)
		}
	}

	return nil
}

// configureDatabaseSettings configures database-specific settings.
// Applies custom configuration JSON to the database instance.
func (g *resourceIBMDatabaseGen2Backend) configureDatabaseSettings(configCtx *instanceConfigContext) error {
	config, ok := configCtx.d.GetOk("configuration")
	if !ok {
		return nil
	}

	var rawConfig map[string]json.RawMessage
	err := json.Unmarshal([]byte(config.(string)), &rawConfig)
	if err != nil {
		return fmt.Errorf("configuration JSON invalid: %w", err)
	}

	var configuration clouddatabasesv5.ConfigurationIntf = new(clouddatabasesv5.Configuration)
	err = core.UnmarshalModel(rawConfig, "", &configuration, clouddatabasesv5.UnmarshalConfiguration)
	if err != nil {
		return fmt.Errorf("database configuration is invalid: %w", err)
	}

	updateDatabaseConfigurationOptions := &clouddatabasesv5.UpdateDatabaseConfigurationOptions{
		ID:            &configCtx.instanceID,
		Configuration: configuration,
	}

	updateDatabaseConfigurationResponse, response, err := configCtx.cloudDatabasesClient.UpdateDatabaseConfiguration(updateDatabaseConfigurationOptions)

	if err != nil {
		return fmt.Errorf("error updating database configuration failed: %w (response: %v)", err, response)
	}

	taskID := *updateDatabaseConfigurationResponse.Task.ID

	icdId := flex.EscapeUrlParm(configCtx.instanceID)
	_, err = waitForDatabaseTaskComplete(taskID, configCtx.d, configCtx.meta, configCtx.d.Timeout(schema.TimeoutUpdate))
	if err != nil {
		return fmt.Errorf("error waiting for database (%s) configuration update task to complete: %w", icdId, err)
	}

	return nil
}

// configureLogicalReplication configures logical replication slots for PostgreSQL.
// Only applicable to PostgreSQL databases; creates replication slots as configured.
func (g *resourceIBMDatabaseGen2Backend) configureLogicalReplication(configCtx *instanceConfigContext) error {
	if _, ok := configCtx.d.GetOk("logical_replication_slot"); !ok {
		return nil
	}

	service := configCtx.d.Get("service").(string)
	if service != "databases-for-postgresql" {
		return fmt.Errorf("logical Replication can only be set for databases-for-postgresql instances")
	}

	_, logicalReplicationList := configCtx.d.GetChange("logical_replication_slot")

	add := logicalReplicationList.(*schema.Set).List()

	for _, entry := range add {
		newEntry := entry.(map[string]interface{})
		logicalReplicationSlot := &clouddatabasesv5.LogicalReplicationSlot{
			Name:         core.StringPtr(newEntry["name"].(string)),
			DatabaseName: core.StringPtr(newEntry["database_name"].(string)),
			PluginType:   core.StringPtr(newEntry["plugin_type"].(string)),
		}

		createLogicalReplicationOptions := &clouddatabasesv5.CreateLogicalReplicationSlotOptions{
			ID:                     &configCtx.instanceID,
			LogicalReplicationSlot: logicalReplicationSlot,
		}

		createLogicalRepSlotResponse, response, err := configCtx.cloudDatabasesClient.CreateLogicalReplicationSlot(createLogicalReplicationOptions)
		if err != nil {
			return fmt.Errorf("CreateLogicalReplicationSlot (%s) failed: %w (response: %v)", *createLogicalReplicationOptions.LogicalReplicationSlot.Name, err, response)
		}

		taskID := *createLogicalRepSlotResponse.Task.ID
		_, err = waitForDatabaseTaskComplete(taskID, configCtx.d, configCtx.meta, configCtx.d.Timeout(schema.TimeoutUpdate))
		if err != nil {
			return fmt.Errorf("error waiting for database (%s) logical replication slot (%s) create task to complete: %w", configCtx.instanceID, *createLogicalReplicationOptions.LogicalReplicationSlot.Name, err)
		}
	}

	return nil
}

// Read retrieves the current state of a database instance.
// Fetches instance details, service info, version, groups, and clears unsupported attributes.
func (g *resourceIBMDatabaseGen2Backend) Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	rsConClient, err := meta.(conns.ClientSession).ResourceControllerV2API()
	if err != nil {
		return diag.FromErr(fmt.Errorf("failed to initialize resource controller client: %w", err))
	}

	instanceID := d.Id()
	rsInst := rc.GetResourceInstanceOptions{
		ID: &instanceID,
	}
	instance, response, err := rsConClient.GetResourceInstance(&rsInst)

	// Check if resource is unavailable (not found or removed)
	if unavailable, diags := g.isResourceUnavailable(instance, response, err, d); unavailable {
		return diags
	}

	if err != nil {
		return diag.FromErr(fmt.Errorf("error retrieving resource instance: %w (response: %v)", err, response))
	}

	// Populate all resource attributes
	return g.populateResourceData(d, instance, meta)
}

// isResourceUnavailable checks if the resource is not found or in a removed state.
// Implements recommendations #1, #2, and #4:
// - Extracts duplicate error handling logic
// - Uses HTTP status code instead of string matching
// - Consolidates state validation logic
// Returns true if the resource should be removed from state, along with any diagnostics.
func (g *resourceIBMDatabaseGen2Backend) isResourceUnavailable(instance *rc.ResourceInstance, response *core.DetailedResponse, err error, d *schema.ResourceData) (bool, diag.Diagnostics) {
	// Check for 404 errors using status code (more robust than string matching)
	if err != nil && response != nil && response.StatusCode == httpNotFound {
		log.Printf("[WARN] Removing record from state because it's not found via the API")
		d.SetId("")
		return true, nil
	}

	// Check for removed state using constant
	if instance != nil && instance.State != nil && strings.Contains(*instance.State, instanceStateRemoved) {
		log.Printf("[WARN] Removing instance from TF state because it's now in removed state")
		d.SetId("")
		return true, nil
	}

	return false, nil
}

// populateResourceData orchestrates setting all resource attributes.
// Implements recommendation #5: Extract attribute setting logic.
// Calls individual setter methods in sequence and returns any errors encountered.
func (g *resourceIBMDatabaseGen2Backend) populateResourceData(d *schema.ResourceData, instance *rc.ResourceInstance, meta interface{}) diag.Diagnostics {
	// Set basic attributes
	if err := g.setBasicAttributes(d, instance, meta); err != nil {
		return diag.FromErr(err)
	}

	// Set service and plan information
	if err := g.setServiceInfo(d, instance, meta); err != nil {
		return diag.FromErr(err)
	}

	// Set version information
	g.setVersionInfo(d, instance)

	// Set groups information
	if err := g.setGroupsInfo(d, instance, meta); err != nil {
		return diag.FromErr(err)
	}

	// Clear Gen2 unsupported attributes
	g.clearUnsupportedAttributes(d)

	return nil
}

// setBasicAttributes sets basic instance attributes.
// Includes tags, name, status, location, GUID, and resource controller URLs.
func (g *resourceIBMDatabaseGen2Backend) setBasicAttributes(d *schema.ResourceData, instance *rc.ResourceInstance, meta interface{}) error {
	tags, err := flex.GetTagsUsingCRN(meta, *instance.CRN)
	if err != nil {
		log.Printf("Error on get of ibm Database tags (%s) tags: %s", d.Id(), err)
	}
	d.Set("tags", tags)
	d.Set("name", *instance.Name)
	d.Set("status", *instance.State)
	d.Set("resource_group_id", *instance.ResourceGroupID)

	var instanceLocation string
	if instance.CRN != nil {
		location := strings.Split(*instance.CRN, ":")
		if len(location) > 5 {
			instanceLocation = location[5]
			d.Set("location", instanceLocation)
		}
	}
	d.Set("guid", *instance.GUID)

	if instance.Parameters != nil {
		if endpoint, ok := instance.Parameters["service-endpoints"]; ok {
			d.Set("service_endpoints", endpoint)
		}
	}

	d.Set(flex.ResourceName, *instance.Name)
	d.Set(flex.ResourceCRN, *instance.CRN)
	d.Set(flex.ResourceStatus, *instance.State)
	d.Set(flex.ResourceGroupName, *instance.ResourceGroupCRN)

	rcontroller, err := flex.GetBaseController(meta)
	if err != nil {
		return fmt.Errorf("failed to get base controller: %w", err)
	}
	d.Set(flex.ResourceControllerURL, rcontroller+"/services/"+url.QueryEscape(*instance.CRN))

	return nil
}

// setServiceInfo sets service and plan information.
// Retrieves service and plan names from the catalog and clears admin user (not available in Gen2).
func (g *resourceIBMDatabaseGen2Backend) setServiceInfo(d *schema.ResourceData, instance *rc.ResourceInstance, meta interface{}) error {
	rsCatClient, err := meta.(conns.ClientSession).ResourceCatalogAPI()
	if err != nil {
		return fmt.Errorf("failed to initialize resource catalog client: %w", err)
	}
	rsCatRepo := rsCatClient.ResourceCatalog()

	serviceOff, err := rsCatRepo.GetServiceName(*instance.ResourceID)
	if err != nil {
		return fmt.Errorf("error retrieving service offering: %w", err)
	}
	d.Set("service", serviceOff)

	servicePlan, err := rsCatRepo.GetServicePlanName(*instance.ResourcePlanID)
	if err != nil {
		return fmt.Errorf("error retrieving plan: %w", err)
	}
	d.Set("plan", servicePlan)

	// Admin user is not available in Gen2
	d.Set(adminUserKey, nil)

	return nil
}

// setVersionInfo extracts and sets version information.
// Uses the helper function to extract version from instance extensions.
func (g *resourceIBMDatabaseGen2Backend) setVersionInfo(d *schema.ResourceData, instance *rc.ResourceInstance) {
	version := ""
	if instance.Extensions != nil && instance.ResourceID != nil {
		version = extractVersionFromExtensions(instance.Extensions, *instance.ResourceID)
	}
	d.Set(versionKey, version)
}

// setGroupsInfo retrieves and sets groups information from catalog.
// Combines instance extensions with catalog metadata to build group configurations.
func (g *resourceIBMDatabaseGen2Backend) setGroupsInfo(d *schema.ResourceData, instance *rc.ResourceInstance, meta interface{}) error {
	if instance.CRN == nil {
		return fmt.Errorf("instance CRN is nil")
	}

	location := strings.Split(*instance.CRN, ":")
	if len(location) <= 5 {
		return fmt.Errorf("invalid CRN format")
	}
	instanceLocation := location[5]

	globalClient, err := meta.(conns.ClientSession).GlobalCatalogV1API()
	if err != nil {
		return fmt.Errorf("failed to initialize global catalog client: %w", err)
	}

	catalogDeployment, err := findDeploymentByLocation(globalClient, *instance.ResourcePlanID, instanceLocation)
	if err != nil {
		return err
	}

	// Extract resources from deployment metadata
	var catalogResources []interface{}
	if catalogDeployment.Metadata != nil && catalogDeployment.Metadata.Other != nil {
		if resources, ok := catalogDeployment.Metadata.Other[resourcesKey].([]interface{}); ok {
			catalogResources = resources
		}
	}

	// Flatten groups using instance extensions and catalog metadata
	if instance.Extensions != nil && len(catalogResources) > 0 && instance.ResourceID != nil {
		d.Set("groups", flattenIcdGroupsFromInstanceAndCatalog(instance.Extensions, catalogResources, *instance.ResourceID))
	}

	return nil
}

// clearUnsupportedAttributes clears attributes not supported in Gen2.
// Sets auto_scaling, allowlist, users, and configuration_schema to nil.
func (g *resourceIBMDatabaseGen2Backend) clearUnsupportedAttributes(d *schema.ResourceData) {
	d.Set(autoScalingKey, nil)
	d.Set(allowlistKey, nil)
	d.Set("users", nil)
	d.Set("configuration_schema", nil)
}

// Update updates an existing database instance.
// TODO: Gen2 update logic is not yet implemented. This is a known limitation.
// Users should use the Classic backend for update operations until this is implemented.
func (g *resourceIBMDatabaseGen2Backend) Update(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	rsConClient, err := meta.(conns.ClientSession).ResourceControllerV2API()
	if err != nil {
		return diag.FromErr(err)
	}

	instanceID := d.Id()
	updateReq := rc.UpdateResourceInstanceOptions{
		ID: &instanceID,
	}
	update := false
	if d.HasChange("name") {
		name := d.Get("name").(string)
		updateReq.Name = &name
		update = true
	}
	if d.HasChange("service_endpoints") {
		params := Params{}
		params.ServiceEndpoints = d.Get("service_endpoints").(string)
		parameters, _ := json.Marshal(params)
		var raw map[string]interface{}
		json.Unmarshal(parameters, &raw)
		updateReq.Parameters = raw
		update = true
	}

	if update {
		_, response, err := rsConClient.UpdateResourceInstance(&updateReq)
		if err != nil {
			return diag.FromErr(fmt.Errorf("[ERROR] Error updating resource instance: %s %s", err, response))
		}

		_, err = waitForDatabaseInstanceUpdate(d, meta)
		if err != nil {
			return diag.FromErr(fmt.Errorf(
				"[ERROR] Error waiting for update of resource instance (%s) to complete: %s", d.Id(), err))
		}
	}

	if d.HasChange("tags") {
		oldList, newList := d.GetChange("tags")
		err = flex.UpdateTagsUsingCRN(oldList, newList, meta, instanceID)
		if err != nil {
			log.Printf(
				"[ERROR] Error on update of Database (%s) tags: %s", d.Id(), err)
		}
	}

	cloudDatabasesClient, err := meta.(conns.ClientSession).CloudDatabasesV5()
	if err != nil {
		return diag.FromErr(fmt.Errorf("[ERROR] Error getting database client settings: %s", err))
	}

	if err != nil {
		return diag.FromErr(fmt.Errorf("[ERROR] Error getting database client settings: %s", err))
	}
	icdId := flex.EscapeUrlParm(instanceID)

	if d.HasChange("configuration") {
		if config, ok := d.GetOk("configuration"); ok {
			var rawConfig map[string]json.RawMessage
			err = json.Unmarshal([]byte(config.(string)), &rawConfig)
			if err != nil {
				return diag.FromErr(fmt.Errorf("[ERROR] configuration JSON invalid\n%s", err))
			}

			var configuration clouddatabasesv5.ConfigurationIntf = new(clouddatabasesv5.Configuration)
			err = core.UnmarshalModel(rawConfig, "", &configuration, clouddatabasesv5.UnmarshalConfiguration)
			if err != nil {
				return diag.FromErr(err)
			}

			updateDatabaseConfigurationOptions := &clouddatabasesv5.UpdateDatabaseConfigurationOptions{
				ID:            &instanceID,
				Configuration: configuration,
			}

			updateDatabaseConfigurationResponse, response, err := cloudDatabasesClient.UpdateDatabaseConfiguration(updateDatabaseConfigurationOptions)

			if err != nil {
				return diag.FromErr(fmt.Errorf(
					"[ERROR] Error updating database configuration failed %s\n%s", err, response))
			}

			taskID := *updateDatabaseConfigurationResponse.Task.ID

			_, err = waitForDatabaseTaskComplete(taskID, d, meta, d.Timeout(schema.TimeoutUpdate))
			if err != nil {
				return diag.FromErr(fmt.Errorf(
					"[ERROR] Error waiting for database (%s) configuration update task to complete: %s", icdId, err))
			}
		}
	}

	if d.HasChange("group") {
		oldGroup, newGroup := d.GetChange("group")
		if oldGroup == nil {
			oldGroup = new(schema.Set)
		}
		if newGroup == nil {
			newGroup = new(schema.Set)
		}

		os := oldGroup.(*schema.Set)
		ns := newGroup.(*schema.Set)

		groupChanges := expandGroups(ns.Difference(os).List())

		groupsResponse, err := getGroups(instanceID, meta)
		if err != nil {
			return diag.FromErr(fmt.Errorf(
				"[ERROR] Error geting group (%s) scaling group update task to complete: %s", icdId, err))
		}

		currentGroups := normalizeGroups(groupsResponse)

		for _, group := range groupChanges {
			groupScaling := &clouddatabasesv5.GroupScaling{}
			var currentGroup *Group
			for _, g := range currentGroups {
				if g.ID == group.ID {
					currentGroup = &g
					break
				}
			}

			if currentGroup == nil {
				return diag.FromErr(fmt.Errorf(
					"[ERROR]  (%s) group does not exist: %s", icdId, err))
			}
			nodeCount := currentGroup.Members.Allocation

			if group.Members != nil && group.Members.Allocation != currentGroup.Members.Allocation {
				groupScaling.Members = &clouddatabasesv5.GroupScalingMembers{AllocationCount: core.Int64Ptr(int64(group.Members.Allocation))}
				nodeCount = group.Members.Allocation
			}
			if group.Memory != nil && group.Memory.Allocation*nodeCount != currentGroup.Memory.Allocation {
				groupScaling.Memory = &clouddatabasesv5.GroupScalingMemory{AllocationMb: core.Int64Ptr(int64(group.Memory.Allocation * nodeCount))}
			}
			if group.Disk != nil && group.Disk.Allocation*nodeCount != currentGroup.Disk.Allocation {
				groupScaling.Disk = &clouddatabasesv5.GroupScalingDisk{AllocationMb: core.Int64Ptr(int64(group.Disk.Allocation * nodeCount))}
			}
			if group.CPU != nil && group.CPU.Allocation*nodeCount != currentGroup.CPU.Allocation {
				groupScaling.CPU = &clouddatabasesv5.GroupScalingCPU{AllocationCount: core.Int64Ptr(int64(group.CPU.Allocation * nodeCount))}
			}
			if group.HostFlavor != nil {
				groupScaling.HostFlavor = &clouddatabasesv5.GroupScalingHostFlavor{ID: core.StringPtr(group.HostFlavor.ID)}
			}

			if groupScaling.Members != nil || groupScaling.Memory != nil || groupScaling.Disk != nil || groupScaling.CPU != nil || groupScaling.HostFlavor != nil {
				setDeploymentScalingGroupOptions := &clouddatabasesv5.SetDeploymentScalingGroupOptions{
					ID:      &instanceID,
					GroupID: &group.ID,
					Group:   groupScaling,
				}

				setDeploymentScalingGroupResponse, response, err := cloudDatabasesClient.SetDeploymentScalingGroup(setDeploymentScalingGroupOptions)

				if err != nil {
					return diag.FromErr(fmt.Errorf("[ERROR] SetDeploymentScalingGroup (%s) failed %s\n%s", group.ID, err, response))
				}

				// API may return HTTP 204 No Content if no change made
				if response.StatusCode == 202 {
					taskIDLink := *setDeploymentScalingGroupResponse.Task.ID

					_, err = waitForDatabaseTaskComplete(taskIDLink, d, meta, d.Timeout(schema.TimeoutCreate))

					if err != nil {
						return diag.FromErr(err)
					}
				}
			}
		}
	}

	if d.HasChange("auto_scaling.0") {
		return diag.FromErr(fmt.Errorf("[ERROR] Auto scaling is not supported for Gen2 database instances"))
	}

	if d.HasChange("adminpassword") {
		return diag.FromErr(fmt.Errorf("[ERROR] Admin password management is not supported for Gen2 database instances. In Gen2, there is no default admin user. Users should manage credentials using the ibm_resource_key resource (https://registry.terraform.io/providers/IBM-Cloud/ibm/latest/docs/resources/resource_key)"))
	}

	if d.HasChange("allowlist") {
		return diag.FromErr(fmt.Errorf("[ERROR] Allowlist is not supported for Gen2 database instances"))
	}

	if d.HasChange("users") {
		return diag.FromErr(fmt.Errorf("[ERROR] User management is not supported for Gen2 database instances. Users should manage credentials using the ibm_resource_key resource (https://registry.terraform.io/providers/IBM-Cloud/ibm/latest/docs/resources/resource_key)"))
	}

	if d.HasChange("logical_replication_slot") {
		return diag.FromErr(fmt.Errorf("[ERROR] Logical replication slot management is not supported for Gen2 database instances. Please use the Classic backend for logical replication slot operations"))
	}

	if d.HasChange("remote_leader_id") {
		remoteLeaderId := d.Get("remote_leader_id").(string)

		if remoteLeaderId == "" {
			skipInitialBackup := false
			if skip, ok := d.GetOk("skip_initial_backup"); ok {
				skipInitialBackup = skip.(bool)
			}

			promoteReadOnlyReplicaOptions := &clouddatabasesv5.PromoteReadOnlyReplicaOptions{
				ID: &instanceID,
				Promotion: map[string]interface{}{
					"skip_initial_backup": skipInitialBackup,
				},
			}

			promoteReadReplicaResponse, response, err := cloudDatabasesClient.PromoteReadOnlyReplica(promoteReadOnlyReplicaOptions)

			if err != nil {
				return diag.FromErr(fmt.Errorf("[ERROR] Error promoting read replica: %s\n%s", err, response))
			}

			taskID := *promoteReadReplicaResponse.Task.ID
			_, err = waitForDatabaseTaskComplete(taskID, d, meta, d.Timeout(schema.TimeoutUpdate))

			if err != nil {
				return diag.FromErr(fmt.Errorf("[ERROR] Error promoting read replica: %s", err))
			}
		}
	}

	if d.HasChange("version") {
		return diag.FromErr(fmt.Errorf("[ERROR] Version changes are not supported for Gen2 database instances."))
	}

	return g.Read(ctx, d, meta)

}

// Delete removes a database instance.
// TODO: Gen2 delete logic is not yet implemented. This is a known limitation.
// Users should use the Classic backend for delete operations until this is implemented.
func (g *resourceIBMDatabaseGen2Backend) Delete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return databaseInstanceDelete(ctx, d, meta)
}

// Exists checks if a database instance exists.
// TODO: Gen2 exists check is not yet implemented. This is a known limitation.
// Users should use the Classic backend until this is implemented.
func (g *resourceIBMDatabaseGen2Backend) Exists(d *schema.ResourceData, meta interface{}) (bool, error) {
	return databaseInstanceExists(d, meta)
}

// WarnUnsupported returns warnings for unsupported features.
// Currently returns no warnings; reserved for future use.
func (g *resourceIBMDatabaseGen2Backend) WarnUnsupported(ctx context.Context, d *schema.ResourceData) diag.Diagnostics {
	return nil
}

// ValidateUnsupportedAttrsDiff validates that unsupported attributes are not configured.
// Returns an error if any Gen2-unsupported attributes are set in the configuration.
func (g *resourceIBMDatabaseGen2Backend) ValidateUnsupportedAttrsDiff(ctx context.Context, d *schema.ResourceDiff, meta interface{}) error {
	var bad []string
	for _, k := range gen2UnsupportedAttrs {
		if isAttrConfiguredInDiff(d, k) {
			bad = append(bad, k)
		}
	}
	if len(bad) == 0 {
		return nil
	}

	planRaw, _ := d.GetOk("plan")
	plan, _ := planRaw.(string)

	return fmt.Errorf(
		"plan %q indicates Gen2. The following attributes are not supported for Gen2 and must be removed: %s",
		strings.TrimSpace(plan),
		strings.Join(bad, ", "),
	)
}
