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
}

type resourceIBMDatabaseGen2Backend struct{}

// newResourceIBMDatabaseGen2Backend creates a new Gen2 backend instance
func newResourceIBMDatabaseGen2Backend() resourceIBMDatabaseBackend {
	return &resourceIBMDatabaseGen2Backend{}
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
	rsConClient, err := meta.(conns.ClientSession).ResourceControllerV2API()
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
	servicePlan, catalogCRN, err := g.getServicePlanAndCatalog(serviceName, plan, location, meta)
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
func (g *resourceIBMDatabaseGen2Backend) getServicePlanAndCatalog(serviceName, plan, location string, meta interface{}) (string, string, error) {
	rsCatClient, err := meta.(conns.ClientSession).ResourceCatalogAPI()
	if err != nil {
		return "", "", fmt.Errorf("failed to initialize resource catalog client: %w", err)
	}
	rsCatRepo := rsCatClient.ResourceCatalog()

	serviceOff, err := rsCatRepo.FindByName(serviceName, true)
	if err != nil {
		return "", "", fmt.Errorf("error retrieving database service offering: %w", err)
	}

	servicePlan, err := rsCatRepo.GetServicePlanID(serviceOff[0], plan)
	if err != nil {
		return "", "", fmt.Errorf("error retrieving plan: %w", err)
	}

	deployments, err := rsCatRepo.ListDeployments(servicePlan)
	if err != nil {
		if serviceName == "databases-for-mongodb" && plan == "enterprise-sharding" {
			return "", "", fmt.Errorf("%s %s is not available yet in this region", serviceName, plan)
		}
		return "", "", fmt.Errorf("error retrieving deployment for plan %s: %w", plan, err)
	}

	if len(deployments) == 0 {
		return "", "", fmt.Errorf("no deployment found for service plan: %s", plan)
	}

	deployments, supportedLocations := filterDatabaseDeployments(deployments, location)

	if len(deployments) == 0 {
		var locationList strings.Builder
		first := true
		for l := range supportedLocations {
			if !first {
				locationList.WriteString(", ")
			}
			locationList.WriteString(l)
			first = false
		}
		return "", "", fmt.Errorf("no deployment found for service plan %s at location %s. Valid location(s) are: %s", plan, location, locationList.String())
	}

	return servicePlan, deployments[0].CatalogCRN, nil
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

	// Build the database-specific configuration
	dbConfig := make(map[string]interface{}, 5)

	// Version
	if version, ok := d.GetOk("version"); ok {
		dbConfig["version"] = version.(string)
	}

	// Get member group configuration
	memberGroup := g.getMemberGroup(d)

	// Members count
	members, err := g.getMembersCount(d, memberGroup, catalogCRN, meta)
	if err != nil {
		return nil, err
	}
	dbConfig["members"] = members

	// Storage in GB (not MB!)
	if memberGroup != nil && memberGroup.Disk != nil {
		// Disk allocation is per member in MB, convert to GB for total
		storageGB := (memberGroup.Disk.Allocation * members) / mbPerGb
		dbConfig["storage_gb"] = storageGB
	}

	// Host flavor
	if memberGroup != nil && memberGroup.HostFlavor != nil {
		dbConfig["host_flavor"] = memberGroup.HostFlavor.ID
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
	if instance == nil || instance.ID == nil {
		return fmt.Errorf("instance or instance ID is nil")
	}

	instanceID := *instance.ID

	cloudDatabasesClient, err := meta.(conns.ClientSession).CloudDatabasesV5()
	if err != nil {
		return fmt.Errorf("failed to initialize cloud databases client: %w", err)
	}

	// Apply group scaling if configured
	if err := g.applyGroupScaling(ctx, d, instanceID, cloudDatabasesClient, meta); err != nil {
		return err
	}

	// Update tags
	if err := g.updateTags(d, instance, meta); err != nil {
		log.Printf("Error on create of ibm database (%s) tags: %s", d.Id(), err)
	}

	// Update admin password if provided
	if err := g.updateAdminPassword(ctx, d, instanceID, cloudDatabasesClient, meta); err != nil {
		return err
	}

	// Configure allowlist if provided
	if err := g.configureAllowlist(ctx, d, instanceID, cloudDatabasesClient, meta); err != nil {
		return err
	}

	// Configure auto-scaling if provided
	if err := g.configureAutoScaling(ctx, d, instanceID, cloudDatabasesClient, meta); err != nil {
		return err
	}

	// Configure users if provided
	if err := g.configureUsers(ctx, d, instanceID, meta); err != nil {
		return err
	}

	// Configure database settings
	if err := g.configureDatabaseSettings(ctx, d, instanceID, cloudDatabasesClient, meta); err != nil {
		return err
	}

	// Configure logical replication slots if provided
	if err := g.configureLogicalReplication(ctx, d, instanceID, cloudDatabasesClient, meta); err != nil {
		return err
	}

	return nil
}

// applyGroupScaling applies scaling configuration to instance groups.
// Compares desired configuration with current state and applies changes as needed.
func (g *resourceIBMDatabaseGen2Backend) applyGroupScaling(ctx context.Context, d *schema.ResourceData, instanceID string, client *clouddatabasesv5.CloudDatabasesV5, meta interface{}) error {
	group, ok := d.GetOk("group")
	if !ok {
		return nil
	}

	groups := expandGroups(group.(*schema.Set).List())
	groupsResponse, err := getGroups(instanceID, meta)
	if err != nil {
		return fmt.Errorf("failed to get groups: %w", err)
	}
	currentGroups := normalizeGroups(groupsResponse)

	for _, grp := range groups {
		if err := g.scaleGroup(ctx, d, instanceID, grp, currentGroups, client, meta); err != nil {
			return err
		}
	}

	return nil
}

// scaleGroup scales a specific group if needed.
// Handles horizontal and vertical scaling for members, memory, disk, CPU, and host flavor.
func (g *resourceIBMDatabaseGen2Backend) scaleGroup(ctx context.Context, d *schema.ResourceData, instanceID string, grp *Group, currentGroups []Group, client *clouddatabasesv5.CloudDatabasesV5, meta interface{}) error {
	groupScaling := &clouddatabasesv5.GroupScaling{}
	var currentGroup *Group
	var nodeCount int

	for i := range currentGroups {
		if currentGroups[i].ID == grp.ID {
			currentGroup = &currentGroups[i]
			nodeCount = currentGroup.Members.Allocation
			break
		}
	}

	if currentGroup == nil {
		return fmt.Errorf("current group %s not found", grp.ID)
	}

	if grp.ID == defaultGroupID && (grp.Members == nil || grp.Members.Allocation == nodeCount) {
		// No Horizontal Scaling needed
		return nil
	}

	if grp.Members != nil && grp.Members.Allocation != currentGroup.Members.Allocation {
		groupScaling.Members = &clouddatabasesv5.GroupScalingMembers{AllocationCount: core.Int64Ptr(int64(grp.Members.Allocation))}
		nodeCount = grp.Members.Allocation
	}
	if grp.Memory != nil && grp.Memory.Allocation*nodeCount != currentGroup.Memory.Allocation {
		groupScaling.Memory = &clouddatabasesv5.GroupScalingMemory{AllocationMb: core.Int64Ptr(int64(grp.Memory.Allocation * nodeCount))}
	}
	if grp.Disk != nil && grp.Disk.Allocation*nodeCount != currentGroup.Disk.Allocation {
		groupScaling.Disk = &clouddatabasesv5.GroupScalingDisk{AllocationMb: core.Int64Ptr(int64(grp.Disk.Allocation * nodeCount))}
	}
	if grp.CPU != nil && grp.CPU.Allocation*nodeCount != currentGroup.CPU.Allocation {
		groupScaling.CPU = &clouddatabasesv5.GroupScalingCPU{AllocationCount: core.Int64Ptr(int64(grp.CPU.Allocation * nodeCount))}
	}
	if grp.HostFlavor != nil {
		groupScaling.HostFlavor = &clouddatabasesv5.GroupScalingHostFlavor{ID: core.StringPtr(grp.HostFlavor.ID)}
	}

	if groupScaling.Members != nil || groupScaling.Memory != nil || groupScaling.Disk != nil || groupScaling.CPU != nil || groupScaling.HostFlavor != nil {
		setDeploymentScalingGroupOptions := &clouddatabasesv5.SetDeploymentScalingGroupOptions{
			ID:      &instanceID,
			GroupID: &grp.ID,
			Group:   groupScaling,
		}

		setDeploymentScalingGroupResponse, _, err := client.SetDeploymentScalingGroup(setDeploymentScalingGroupOptions)
		if err != nil {
			return fmt.Errorf("failed to set deployment scaling group: %w", err)
		}

		taskIDLink := *setDeploymentScalingGroupResponse.Task.ID

		_, err = waitForDatabaseTaskComplete(taskIDLink, d, meta, d.Timeout(schema.TimeoutCreate))
		if err != nil {
			return fmt.Errorf("error waiting for scaling task to complete: %w", err)
		}
	}

	return nil
}

// updateTags updates resource tags.
// Compares old and new tags and applies changes using the CRN.
func (g *resourceIBMDatabaseGen2Backend) updateTags(d *schema.ResourceData, instance *rc.ResourceInstance, meta interface{}) error {
	v := os.Getenv("IC_ENV_TAGS")
	if _, ok := d.GetOk("tags"); ok || v != "" {
		oldList, newList := d.GetChange("tags")
		err := flex.UpdateTagsUsingCRN(oldList, newList, meta, *instance.CRN)
		if err != nil {
			return fmt.Errorf("failed to update tags: %w", err)
		}
	}
	return nil
}

// updateAdminPassword updates the admin password if provided.
// Retrieves the admin username from deployment info and updates the password.
func (g *resourceIBMDatabaseGen2Backend) updateAdminPassword(ctx context.Context, d *schema.ResourceData, instanceID string, client *clouddatabasesv5.CloudDatabasesV5, meta interface{}) error {
	pw, ok := d.GetOk("adminpassword")
	if !ok {
		return nil
	}

	adminPassword := pw.(string)

	getDeploymentInfoOptions := &clouddatabasesv5.GetDeploymentInfoOptions{
		ID: core.StringPtr(instanceID),
	}
	getDeploymentInfoResponse, response, err := client.GetDeploymentInfo(getDeploymentInfoOptions)

	if err != nil {
		if response != nil && response.StatusCode == httpNotFound {
			return fmt.Errorf("the database instance was not found in the region set for the Provider, or the default of us-south. Specify the correct region in the provider definition, or create a provider alias for the correct region: %w", err)
		}
		return fmt.Errorf("error getting database config while updating adminpassword for: %s: %w", instanceID, err)
	}

	if getDeploymentInfoResponse == nil || getDeploymentInfoResponse.Deployment == nil {
		return fmt.Errorf("deployment info response is nil")
	}

	deployment := getDeploymentInfoResponse.Deployment
	adminUser := deployment.AdminUsernames[databaseUserType]

	user := &clouddatabasesv5.UserUpdatePasswordSetting{
		Password: &adminPassword,
	}

	updateUserOptions := &clouddatabasesv5.UpdateUserOptions{
		ID:       core.StringPtr(instanceID),
		UserType: core.StringPtr(databaseUserType),
		Username: core.StringPtr(adminUser),
		User:     user,
	}

	updateUserResponse, response, err := client.UpdateUser(updateUserOptions)
	if err != nil {
		return fmt.Errorf("UpdateUser (%s) failed: %w (response: %v)", *updateUserOptions.Username, err, response)
	}

	taskID := *updateUserResponse.Task.ID
	_, err = waitForDatabaseTaskComplete(taskID, d, meta, d.Timeout(schema.TimeoutCreate))

	if err != nil {
		return fmt.Errorf("error updating database admin password: %w", err)
	}

	return nil
}

// configureAllowlist configures the IP allowlist.
// Sets the list of allowed IP addresses for database access.
func (g *resourceIBMDatabaseGen2Backend) configureAllowlist(ctx context.Context, d *schema.ResourceData, instanceID string, client *clouddatabasesv5.CloudDatabasesV5, meta interface{}) error {
	_, hasAllowlist := d.GetOk("allowlist")
	if !hasAllowlist {
		return nil
	}

	ipAddresses := d.Get("allowlist").(*schema.Set)
	entries := flex.ExpandAllowlist(ipAddresses)

	setAllowlistOptions := &clouddatabasesv5.SetAllowlistOptions{
		ID:          &instanceID,
		IPAddresses: entries,
	}

	setAllowlistResponse, _, err := client.SetAllowlist(setAllowlistOptions)
	if err != nil {
		return fmt.Errorf("error updating database allowlists: %w", err)
	}

	taskId := *setAllowlistResponse.Task.ID

	_, err = waitForDatabaseTaskComplete(taskId, d, meta, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("error waiting for update of database (%s) allowlist task to complete: %w", instanceID, err)
	}

	return nil
}

// configureAutoScaling configures auto-scaling settings.
// Sets disk and memory auto-scaling conditions if configured.
func (g *resourceIBMDatabaseGen2Backend) configureAutoScaling(ctx context.Context, d *schema.ResourceData, instanceID string, client *clouddatabasesv5.CloudDatabasesV5, meta interface{}) error {
	if _, ok := d.GetOk("auto_scaling.0"); !ok {
		return nil
	}

	autoscalingSetGroupAutoscaling := &clouddatabasesv5.AutoscalingSetGroupAutoscaling{}

	if diskRecord, ok := d.GetOk("auto_scaling.0.disk"); ok {
		diskGroup, err := expandAutoscalingDiskGroup(d, diskRecord)
		if err != nil {
			return fmt.Errorf("error in getting diskGroup from expandAutoscalingDiskGroup: %w", err)
		}
		autoscalingSetGroupAutoscaling.Disk = diskGroup
	}

	if memoryRecord, ok := d.GetOk("auto_scaling.0.memory"); ok {
		memoryGroup, err := expandAutoscalingMemoryGroup(d, memoryRecord)
		if err != nil {
			return fmt.Errorf("error in getting memoryBody from expandAutoscalingMemoryGroup: %w", err)
		}
		autoscalingSetGroupAutoscaling.Memory = memoryGroup
	}

	if autoscalingSetGroupAutoscaling.Disk != nil || autoscalingSetGroupAutoscaling.Memory != nil {
		setAutoscalingConditionsOptions := &clouddatabasesv5.SetAutoscalingConditionsOptions{
			ID:          &instanceID,
			GroupID:     core.StringPtr(defaultGroupID),
			Autoscaling: autoscalingSetGroupAutoscaling,
		}

		setAutoscalingConditionsResponse, _, err := client.SetAutoscalingConditions(setAutoscalingConditionsOptions)
		if err != nil {
			return fmt.Errorf("error updating database auto_scaling: %w", err)
		}

		taskId := *setAutoscalingConditionsResponse.Task.ID

		_, err = waitForDatabaseTaskComplete(taskId, d, meta, d.Timeout(schema.TimeoutCreate))
		if err != nil {
			return fmt.Errorf("error waiting for database (%s) memory auto_scaling group update task to complete: %w", instanceID, err)
		}
	}

	return nil
}

// configureUsers configures database users.
// Attempts to update existing users or creates new ones if they don't exist.
func (g *resourceIBMDatabaseGen2Backend) configureUsers(ctx context.Context, d *schema.ResourceData, instanceID string, meta interface{}) error {
	userList, ok := d.GetOk("users")
	if !ok {
		return nil
	}

	users := expandUsers(userList.(*schema.Set).List())
	for _, user := range users {
		// Note: Some db users exist after provisioning (i.e. admin, repl)
		// so we must attempt both methods
		err := user.Update(instanceID, d, meta)

		if err != nil {
			err = user.Create(instanceID, d, meta)
		}

		if err != nil {
			return fmt.Errorf("error configuring user %s: %w", user.Username, err)
		}
	}

	return nil
}

// configureDatabaseSettings configures database-specific settings.
// Applies custom configuration JSON to the database instance.
func (g *resourceIBMDatabaseGen2Backend) configureDatabaseSettings(ctx context.Context, d *schema.ResourceData, instanceID string, client *clouddatabasesv5.CloudDatabasesV5, meta interface{}) error {
	config, ok := d.GetOk("configuration")
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
		ID:            &instanceID,
		Configuration: configuration,
	}

	updateDatabaseConfigurationResponse, response, err := client.UpdateDatabaseConfiguration(updateDatabaseConfigurationOptions)

	if err != nil {
		return fmt.Errorf("error updating database configuration failed: %w (response: %v)", err, response)
	}

	taskID := *updateDatabaseConfigurationResponse.Task.ID

	icdId := flex.EscapeUrlParm(instanceID)
	_, err = waitForDatabaseTaskComplete(taskID, d, meta, d.Timeout(schema.TimeoutUpdate))
	if err != nil {
		return fmt.Errorf("error waiting for database (%s) configuration update task to complete: %w", icdId, err)
	}

	return nil
}

// configureLogicalReplication configures logical replication slots for PostgreSQL.
// Only applicable to PostgreSQL databases; creates replication slots as configured.
func (g *resourceIBMDatabaseGen2Backend) configureLogicalReplication(ctx context.Context, d *schema.ResourceData, instanceID string, client *clouddatabasesv5.CloudDatabasesV5, meta interface{}) error {
	if _, ok := d.GetOk("logical_replication_slot"); !ok {
		return nil
	}

	service := d.Get("service").(string)
	if service != "databases-for-postgresql" {
		return fmt.Errorf("logical Replication can only be set for databases-for-postgresql instances")
	}

	_, logicalReplicationList := d.GetChange("logical_replication_slot")

	add := logicalReplicationList.(*schema.Set).List()

	for _, entry := range add {
		newEntry := entry.(map[string]interface{})
		logicalReplicationSlot := &clouddatabasesv5.LogicalReplicationSlot{
			Name:         core.StringPtr(newEntry["name"].(string)),
			DatabaseName: core.StringPtr(newEntry["database_name"].(string)),
			PluginType:   core.StringPtr(newEntry["plugin_type"].(string)),
		}

		createLogicalReplicationOptions := &clouddatabasesv5.CreateLogicalReplicationSlotOptions{
			ID:                     &instanceID,
			LogicalReplicationSlot: logicalReplicationSlot,
		}

		createLogicalRepSlotResponse, response, err := client.CreateLogicalReplicationSlot(createLogicalReplicationOptions)
		if err != nil {
			return fmt.Errorf("CreateLogicalReplicationSlot (%s) failed: %w (response: %v)", *createLogicalReplicationOptions.LogicalReplicationSlot.Name, err, response)
		}

		taskID := *createLogicalRepSlotResponse.Task.ID
		_, err = waitForDatabaseTaskComplete(taskID, d, meta, d.Timeout(schema.TimeoutUpdate))
		if err != nil {
			return fmt.Errorf("error waiting for database (%s) logical replication slot (%s) create task to complete: %w", instanceID, *createLogicalReplicationOptions.LogicalReplicationSlot.Name, err)
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
	if err != nil {
		if strings.Contains(err.Error(), "Object not found") ||
			strings.Contains(err.Error(), "status code: 404") {
			log.Printf("[WARN] Removing record from state because it's not found via the API")
			d.SetId("")
			return nil
		}
		return diag.FromErr(fmt.Errorf("error retrieving resource instance: %w (response: %v)", err, response))
	}

	if instance.State != nil && strings.Contains(*instance.State, "removed") {
		log.Printf("[WARN] Removing instance from TF state because it's now in removed state")
		d.SetId("")
		return nil
	}

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
	return diag.Errorf("Update operation for Gen2 backend is not yet implemented. Please contact support or use the Classic backend for update operations.")
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
