// Copyright IBM Corp. 2017, 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/IBM-Cloud/bluemix-go/bmxerror"
	"github.com/IBM-Cloud/bluemix-go/models"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/go-sdk-core/v5/core"
	rc "github.com/IBM/platform-services-go-sdk/resourcecontrollerv2"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var gen2UnsupportedAttrs = []string{
	"point_in_time_recovery_deployment_id",
	"point_in_time_recovery_time",
	"backup_policy",
	"users",
	"allowlist",
	"remote_leader_id",
	"adminpassword",
	"backup_encryption_key_crn",
}

// gen2IgnoredAttrs are attributes that are accepted but have no effect in Gen2
// These generate warnings but don't cause plan failures
var gen2IgnoredAttrs = []string{
	"key_protect_instance",
	"auto_scaling",
	"configuration",
	"logical_replication_slot",
	"offline_restore",
	"async_restore",
	"version_upgrade_skip_backup",
	"skip_initial_backup",
}

const (
	// Parameter keys for Gen2 database configuration
	serviceEndpointsKey = "service-endpoints"
	remoteLeaderIDKey   = "remote_leader_id"

	// Encryption keys
	diskEncryptionKey   = "disk"
	backupEncryptionKey = "backup"
	encryptionKey       = "encryption"
)

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
// Note: Gen2 uses Resource Controller API only, no CloudDatabasesV5 client needed.
type instanceConfigContext struct {
	ctx        context.Context
	d          *schema.ResourceData
	instanceID string
	meta       interface{}
	instance   *rc.ResourceInstance
}

// serviceMetadata encapsulates service-related metadata for database operations.
// Reduces parameter extraction boilerplate and improves testability.
type serviceMetadata struct {
	serviceName string
	catalogCRN  string
}

// updateContext encapsulates the client and location needed for update operations.
// Provides a single point of initialization for update-related operations.
type updateContext struct {
	client   *rc.ResourceControllerV2
	location string
}

type resourceIBMDatabaseGen2Backend struct{}

type gen2AttrBehavior string

const (
	gen2AttrBehaviorUnsupported gen2AttrBehavior = "unsupported"
	gen2AttrBehaviorIgnored     gen2AttrBehavior = "ignored"
)

var gen2AttrGuidance = map[string]string{
	"users": "For user management in Gen2 databases, use the Terraform resource 'ibm_resource_key' instead.\n" +
		"Documentation: https://registry.terraform.io/providers/IBM-Cloud/ibm/latest/docs/resources/resource_key",

	"allowlist": "For IP allowlisting in Gen2 databases, use the Terraform resource 'ibm_cbr_rule' instead.\n" +
		"Documentation: https://registry.terraform.io/providers/IBM-Cloud/ibm/latest/docs/resources/cbr_rule",

	"adminpassword": "Gen2 databases do not create default admin user during provisioning.\n" +
		"Please use the Terraform resource 'ibm_resource_key' to create and manage one.\n" +
		"Documentation: https://registry.terraform.io/providers/IBM-Cloud/ibm/latest/docs/resources/resource_key",

	"point_in_time_recovery_deployment_id": "Gen2 databases do not support restoring from backups using the 'point_in_time_recovery_deployment_id' attribute at this point.\n" +
		"Documentation: https://registry.terraform.io/providers/IBM-Cloud/ibm/latest/docs/resources/database",

	"point_in_time_recovery_time": "Gen2 databases do not support restoring from backups using point_in_time_recovery at this point.\n" +
		"Documentation: https://registry.terraform.io/providers/IBM-Cloud/ibm/latest/docs/resources/database",

	"backup_policy": "Gen2 databases do not support backup_policy at this point.\n" +
		"Documentation: https://registry.terraform.io/providers/IBM-Cloud/ibm/latest/docs/resources/database",

	"backup_encryption_key_crn": "Gen2 databases do not support backup_encryption_key_crn at this point.\n" +
		"Documentation: https://registry.terraform.io/providers/IBM-Cloud/ibm/latest/docs/resources/database",

	"remote_leader_id": "Gen2 databases do not yet support read replica creation and promotion using the 'remote_leader_id' attribute at this point.\n" +
		"Documentation: https://registry.terraform.io/providers/IBM-Cloud/ibm/latest/docs/resources/database",

	"configuration":               "database configuration changes are currently ignored",
	"logical_replication_slot":    "logical replication slot creation is currently ignored",
	"auto_scaling":                "auto-scaling settings are currently ignored",
	"offline_restore":             "offline restore settings are currently ignored",
	"async_restore":               "async restore settings are currently ignored",
	"version_upgrade_skip_backup": "version upgrade backup skipping is currently ignored",
	"skip_initial_backup":         "initial backup skipping is currently ignored",
	"key_protect_instance":        "this Classic key protection setting is currently ignored",
}

func getGen2UnsupportedAttrGuidance(attr string) string {
	return getGen2AttrGuidance(attr, gen2AttrBehaviorUnsupported)
}

func getGen2IgnoredAttrGuidance(attr string) string {
	return getGen2AttrGuidance(attr, gen2AttrBehaviorIgnored)
}

func getGen2AttrGuidance(attr string, behavior gen2AttrBehavior) string {
	if msg, ok := gen2AttrGuidance[attr]; ok {
		return msg
	}

	switch behavior {
	case gen2AttrBehaviorIgnored:
		return "This attribute has no effect in Gen2 databases and is currently ignored."
	default:
		return "This attribute is not supported for Gen2 databases. Please remove it from your configuration."
	}
}

// newResourceIBMDatabaseGen2Backend creates a new Gen2 backend instance
func newResourceIBMDatabaseGen2Backend() resourceIBMDatabaseBackend {
	return &resourceIBMDatabaseGen2Backend{}
}

// getResourceControllerClient initializes and returns the Resource Controller V2 client.
// Centralizes client initialization to reduce duplication and improve testability.
func (g *resourceIBMDatabaseGen2Backend) getResourceControllerClient(meta interface{}) (*rc.ResourceControllerV2, error) {
	rsConClient, err := meta.(conns.ClientSession).ResourceControllerV2API()
	if err != nil {
		return nil, fmt.Errorf("failed to initialize resource controller client: %w", err)
	}
	return rsConClient, nil
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
	warnings := g.WarnIgnoredAttrs(d)

	instance, err := g.createResourceInstance(d, meta)
	if err != nil {
		return appendGen2DiagnosticsErrorsThenWarnings(diag.FromErr(err), warnings)
	}

	d.SetId(*instance.ID)

	_, err = waitForDatabaseInstanceCreate(d, meta, *instance.ID, false)
	if err != nil {
		return appendGen2DiagnosticsErrorsThenWarnings(
			diag.FromErr(fmt.Errorf("error waiting for create database instance (%s) to complete: %w", *instance.ID, err)),
			warnings,
		)
	}

	readDiags := resourceIBMDatabaseInstanceRead(ctx, d, meta)
	return appendGen2DiagnosticsErrorsThenWarnings(readDiags, warnings)
}

// createResourceInstance handles the initial resource instance creation.
// It retrieves service and plan information, builds Gen2 parameters, includes tags, and creates the instance.
// Everything is done in a single API call to avoid unnecessary resource provisioning/deprovisioning.
func (g *resourceIBMDatabaseGen2Backend) createResourceInstance(d *schema.ResourceData, meta interface{}) (*rc.ResourceInstance, error) {
	clientSession := meta.(conns.ClientSession)
	rsConClient, err := g.getResourceControllerClient(meta)
	if err != nil {
		return nil, err
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

	// Build Gen2 parameters (database config + encryption)
	parameters, err := g.buildGen2Parameters(d, serviceName, meta, catalogCRN)
	if err != nil {
		return nil, err
	}
	rsInst.Parameters = parameters

	// Add tags if specified (tags is a TypeSet in the schema)
	if tags, ok := d.GetOk("tags"); ok {
		tagSet := tags.(*schema.Set)
		rsInst.Tags = flex.ExpandStringList(tagSet.List())
	}

	// Create the instance with retry logic
	instance, response, err := g.createInstanceWithRetry(rsConClient, &rsInst)
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

	// Check special case before calling ListDeployments to avoid unnecessary API call
	if serviceName == "databases-for-mongodb" && plan == "enterprise-sharding" {
		return "", "", fmt.Errorf("%s %s is not available yet in this region", serviceName, plan)
	}

	deployments, err := rsCatRepo.ListDeployments(servicePlan)
	if err != nil {
		return "", "", fmt.Errorf("error retrieving deployment for plan %s: %w", plan, err)
	}

	if len(deployments) == 0 {
		return "", "", fmt.Errorf("no deployment found for service plan: %s", plan)
	}

	// Filter and validate deployment location
	catalogCRN, err := g.validateAndGetCatalogCRN(deployments, location, plan)
	if err != nil {
		return "", "", fmt.Errorf("%v", err)
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
// Includes database configuration, encryption settings, and backup_id for restore.
// Note: PITR is not supported in Gen2. backup_id is validated to ensure only Gen2 backups are used.
func (g *resourceIBMDatabaseGen2Backend) buildGen2Parameters(d *schema.ResourceData, serviceName string, meta interface{}, catalogCRN string) (map[string]interface{}, error) {
	// Validate backup_id if provided (only Gen2 coupled and decoupled backups are allowed at this point)
	if backupID, ok := d.GetOk("backup_id"); ok {
		if err := validateGen2BackupCRN(backupID.(string), meta); err != nil {
			return nil, err
		}
	}

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

	// Add restore_backup_id if provided (for restore from backup)
	// Note: Gen2 uses "restore_backup_id" inside dataservices, not "backup_id" at top level
	if backupID, ok := d.GetOk("backup_id"); ok {
		dataservices["restore_backup_id"] = backupID.(string)
	}

	// Build final parameters structure
	parameters := map[string]interface{}{
		"dataservices": dataservices,
	}

	return parameters, nil
}

// buildDBConfig creates database configuration with member group and storage settings.
// Extracts and consolidates member group logic, reducing nested if statements.
// Gen2 supports: members, disk, and host_flavor from groups.
// Note: memory and cpu are NOT supported independently in Gen2 - they are controlled by host_flavor.
func (g *resourceIBMDatabaseGen2Backend) buildDBConfig(d *schema.ResourceData, catalogCRN string, meta interface{}) (map[string]interface{}, error) {
	config := DBConfig{}

	// Version
	if version, ok := d.GetOk("version"); ok {
		config.Version = version.(string)
	}

	// Get member group configuration
	memberGroup := g.getMemberGroup(d)

	// Members count - use from group if specified, otherwise get default from catalog
	members, err := g.getMembersCount(memberGroup, catalogCRN, meta)
	if err != nil {
		return nil, err
	}
	config.Members = members

	// Early return if no member group - simplifies logic below
	if memberGroup == nil {
		return g.dbConfigToMap(config), nil
	}

	// Storage in GB (not MB!) - Gen2 expects per-member allocation
	if memberGroup.Disk != nil {
		storageGB := memberGroup.Disk.Allocation / mbPerGb
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
func (g *resourceIBMDatabaseGen2Backend) getMembersCount(memberGroup *Group, catalogCRN string, meta interface{}) (int, error) {
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
// Includes disk encryption key CRN if configured.
// Note: backup_encryption_key_crn is not supported for Gen2 instances.
func (g *resourceIBMDatabaseGen2Backend) addEncryptionConfig(d *schema.ResourceData, dataservices map[string]interface{}) {
	encryption := make(map[string]interface{}, 1)
	if keyProtect, ok := d.GetOk("key_protect_key"); ok {
		encryption[diskEncryptionKey] = keyProtect.(string)
	}
	if len(encryption) > 0 {
		dataservices[encryptionKey] = encryption
	}
}

// createInstanceWithRetry creates an instance.
// Note: Retry logic can be added in the future if needed.
func (g *resourceIBMDatabaseGen2Backend) createInstanceWithRetry(client *rc.ResourceControllerV2, opts *rc.CreateResourceInstanceOptions) (*rc.ResourceInstance, *core.DetailedResponse, error) {
	instance, response, err := client.CreateResourceInstance(opts)
	return instance, response, err
}

// configureInstance applies post-creation configuration to the instance.
// Note: Group scaling is NOT included here - all group parameters (members, storage, host_flavor)
// are passed during initial creation to avoid unnecessary resource provisioning/deprovisioning.
// Only tags are configured post-creation as they don't affect resource provisioning.
func (g *resourceIBMDatabaseGen2Backend) configureInstance(ctx context.Context, d *schema.ResourceData, meta interface{}, instance *rc.ResourceInstance) error {
	// Initialize configuration context
	configCtx, err := g.initConfigContext(ctx, d, meta, instance)
	if err != nil {
		return err
	}

	// Update tags only - all other configuration is done during initial creation
	if err := g.updateTags(configCtx); err != nil {
		return fmt.Errorf("failed to configure tags: %w", err)
	}

	return nil
}

// initConfigContext initializes the configuration context with validated instance.
// Note: CloudDatabasesV5 client removed as Gen2 uses Resource Controller for all operations.
func (g *resourceIBMDatabaseGen2Backend) initConfigContext(ctx context.Context, d *schema.ResourceData, meta interface{}, instance *rc.ResourceInstance) (*instanceConfigContext, error) {
	if instance == nil || instance.ID == nil {
		return nil, fmt.Errorf("instance or instance ID is nil")
	}

	return &instanceConfigContext{
		ctx:        ctx,
		d:          d,
		instanceID: *instance.ID,
		meta:       meta,
		instance:   instance,
	}, nil
}

// getServiceMetadata retrieves service name and catalog CRN for the database instance.
// Consolidates parameter extraction and catalog lookup into a single operation.
func (g *resourceIBMDatabaseGen2Backend) getServiceMetadata(d *schema.ResourceData, location string, session conns.ClientSession) (*serviceMetadata, error) {
	serviceName := d.Get("service").(string)
	plan := d.Get("plan").(string)

	_, catalogCRN, err := g.getServicePlanAndCatalog(serviceName, plan, location, session)
	if err != nil {
		return nil, fmt.Errorf("failed to get catalog CRN: %w", err)
	}

	return &serviceMetadata{
		serviceName: serviceName,
		catalogCRN:  catalogCRN,
	}, nil
}

// prepareUpdateContext initializes the Resource Controller client and extracts instance location.
// Provides a single point of initialization for update operations, improving testability.
func (g *resourceIBMDatabaseGen2Backend) prepareUpdateContext(configCtx *instanceConfigContext) (*updateContext, error) {
	rsConClient, err := g.getResourceControllerClient(configCtx.meta)
	if err != nil {
		return nil, err
	}

	instanceLocation, err := extractLocationFromCRN(configCtx.instance.CRN)
	if err != nil {
		return nil, err
	}

	return &updateContext{
		client:   rsConClient,
		location: instanceLocation,
	}, nil
}

// updateResourceInstanceParameters updates a resource instance with new parameters.
// Encapsulates the update request construction and execution for reusability.
func (g *resourceIBMDatabaseGen2Backend) updateResourceInstanceParameters(
	rsConClient *rc.ResourceControllerV2,
	instanceID string,
	parameters map[string]interface{},
) error {
	updateReq := rc.UpdateResourceInstanceOptions{
		ID:         &instanceID,
		Parameters: parameters,
	}

	_, response, err := rsConClient.UpdateResourceInstance(&updateReq)
	if err != nil {
		return wrapAPIError("update resource instance", err, response)
	}

	return nil
}

// applyGroupScaling applies scaling configuration to instance groups using Resource Controller.
// Flattens group configuration into parameters and updates the instance via UpdateResourceInstance API.
// This approach is consistent with how groups are handled at CREATE time and removes CloudDatabasesV5 dependency.
func (g *resourceIBMDatabaseGen2Backend) applyGroupScaling(configCtx *instanceConfigContext) error {
	if _, ok := configCtx.d.GetOk("group"); !ok {
		return nil
	}

	// Initialize clients and extract location
	updateCtx, err := g.prepareUpdateContext(configCtx)
	if err != nil {
		return err
	}

	// Get service metadata
	clientSession := configCtx.meta.(conns.ClientSession)
	metadata, err := g.getServiceMetadata(configCtx.d, updateCtx.location, clientSession)
	if err != nil {
		return err
	}

	// Build Gen2 parameters with updated group configuration
	parameters, err := g.buildGen2Parameters(configCtx.d, metadata.serviceName, configCtx.meta, metadata.catalogCRN)
	if err != nil {
		return fmt.Errorf("failed to build parameters: %w", err)
	}

	// Update the instance
	if err := g.updateResourceInstanceParameters(updateCtx.client, configCtx.instanceID, parameters); err != nil {
		return err
	}

	// Wait for update to complete
	_, err = g.waitForGen2InstanceUpdate(configCtx.d, configCtx.meta)
	if err != nil {
		return fmt.Errorf("error waiting for instance update to complete: %w", err)
	}

	return nil
}

// waitForGen2InstanceUpdate waits for a Gen2 database instance update to complete.
// Unlike Classic databases, Gen2 only uses Resource Controller API and doesn't require ICD API checks.
func (g *resourceIBMDatabaseGen2Backend) waitForGen2InstanceUpdate(d *schema.ResourceData, meta interface{}) (interface{}, error) {
	rsConClient, err := meta.(conns.ClientSession).ResourceControllerV2API()
	if err != nil {
		return false, err
	}
	instanceID := d.Id()

	stateConf := &resource.StateChangeConf{
		Pending: []string{databaseInstanceProgressStatus, databaseInstanceInactiveStatus},
		Target:  []string{databaseInstanceSuccessStatus},
		Refresh: func() (interface{}, string, error) {
			rsInst := rc.GetResourceInstanceOptions{
				ID: &instanceID,
			}
			instance, response, err := rsConClient.GetResourceInstance(&rsInst)
			if err != nil {
				if apiErr, ok := err.(bmxerror.RequestFailure); ok && apiErr.StatusCode() == 404 {
					return nil, "", fmt.Errorf("[ERROR] The resource instance %s does not exist anymore: %s %s", d.Id(), err, response)
				}
				return nil, "", fmt.Errorf("[ERROR] GetResourceInstance on %s failed with error %s %s", d.Id(), err, response)
			}
			if *instance.State == databaseInstanceFailStatus {
				return *instance, *instance.State, fmt.Errorf("[ERROR] The resource instance %s failed: %s %s", d.Id(), err, response)
			}
			return *instance, *instance.State, nil
		},
		Timeout:    d.Timeout(schema.TimeoutUpdate),
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	// Gen2 databases don't use ICD API - only Resource Controller
	// No need to call waitForICDReady() like Classic databases do
	return stateConf.WaitForState()
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

// Read retrieves the current state of a database instance.
// Fetches instance details, service info, version, groups, and clears unsupported attributes.
func (g *resourceIBMDatabaseGen2Backend) Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	rsConClient, err := g.getResourceControllerClient(meta)
	if err != nil {
		return diag.FromErr(err)
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
	var diags diag.Diagnostics

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

	// Check for ignored attributes and add warnings
	diags = append(diags, g.WarnIgnoredAttrs(d)...)

	return diags
}

// setBasicAttributes sets basic instance attributes.
// Includes tags, name, status, location, GUID, and resource controller URLs.
func (g *resourceIBMDatabaseGen2Backend) setBasicAttributes(d *schema.ResourceData, instance *rc.ResourceInstance, meta interface{}) error {
	// Use shared Gen2 helper function
	// Resources need service_endpoints and resource_controller_url
	return setGen2BasicAttributes(d, instance, meta, true, true)
}

// setServiceInfo sets service and plan information.
// Retrieves service and plan names from the catalog and clears admin user (not available in Gen2).
func (g *resourceIBMDatabaseGen2Backend) setServiceInfo(d *schema.ResourceData, instance *rc.ResourceInstance, meta interface{}) error {
	// Use shared Gen2 helper function
	return setGen2ServiceInfo(d, instance, meta)
}

// setVersionInfo extracts and sets version information.
// Uses the helper function to extract version from instance extensions.
func (g *resourceIBMDatabaseGen2Backend) setVersionInfo(d *schema.ResourceData, instance *rc.ResourceInstance) {
	// Use shared Gen2 helper function
	// Resources don't include platform_options
	setGen2VersionInfo(d, instance, false)
}

// setGroupsInfo retrieves and sets groups information from catalog.
// Combines instance extensions with catalog metadata to build group configurations.
func (g *resourceIBMDatabaseGen2Backend) setGroupsInfo(d *schema.ResourceData, instance *rc.ResourceInstance, meta interface{}) error {
	// Use shared Gen2 helper function
	return setGen2GroupsInfo(d, instance, meta)
}

// clearUnsupportedAttributes clears attributes not supported in Gen2.
// Sets auto_scaling, allowlist, users, and configuration_schema to nil.
func (g *resourceIBMDatabaseGen2Backend) clearUnsupportedAttributes(d *schema.ResourceData) {
	// Use shared Gen2 helper function
	clearGen2UnsupportedAttributes(d)
}

// diagError creates a diagnostic error with consistent formatting.
func diagError(format string, args ...interface{}) diag.Diagnostics {
	return diag.FromErr(fmt.Errorf("[ERROR] "+format, args...))
}

// updateBasicAttributes updates basic instance attributes.
// Returns true if any updates were made.
func (g *resourceIBMDatabaseGen2Backend) updateBasicAttributes(d *schema.ResourceData, updateReq *rc.UpdateResourceInstanceOptions) bool {
	update := false

	if d.HasChange("name") {
		name := d.Get("name").(string)
		updateReq.Name = &name
		update = true
	}

	if d.HasChange("service_endpoints") {
		update = true
	}

	return update
}

// applyBasicAttributeUpdates updates basic instance attributes and waits for completion.
// Returns diagnostics if any errors occur during the update process.
func (g *resourceIBMDatabaseGen2Backend) applyBasicAttributeUpdates(d *schema.ResourceData, rsConClient *rc.ResourceControllerV2, instanceID string, meta interface{}) diag.Diagnostics {
	updateReq := rc.UpdateResourceInstanceOptions{
		ID: &instanceID,
	}

	if !g.updateBasicAttributes(d, &updateReq) {
		return nil
	}

	_, response, err := rsConClient.UpdateResourceInstance(&updateReq)
	if err != nil {
		return diagError("error updating resource instance: %s %s", err, response)
	}

	_, err = waitForDatabaseInstanceUpdate(d, meta)
	if err != nil {
		return diagError("error waiting for update of resource instance (%s) to complete: %s", d.Id(), err)
	}

	return nil
}

// updateTagsWithDiagnostics updates resource tags and returns diagnostic errors on failure.
// Returns diagnostics if tag update fails.
func (g *resourceIBMDatabaseGen2Backend) updateTagsWithDiagnostics(d *schema.ResourceData, instanceID string, meta interface{}) diag.Diagnostics {
	if !d.HasChange("tags") {
		return nil
	}

	oldList, newList := d.GetChange("tags")
	err := flex.UpdateTagsUsingCRN(oldList, newList, meta, instanceID)
	if err != nil {
		log.Printf("[ERROR] Error on update of Database (%s) tags: %s", d.Id(), err)
		return diagError("error updating tags: %s", err)
	}

	return nil
}

// checkUnsupportedChanges validates that no unsupported Gen2 features are being modified.
// Returns a diagnostic error if any unsupported changes are detected.
// Uses the same guidance messages as plan-time validation for consistency.
func (g *resourceIBMDatabaseGen2Backend) checkUnsupportedChanges(d *schema.ResourceData) diag.Diagnostics {
	// Check all unsupported attributes
	for _, attr := range gen2UnsupportedAttrs {
		if d.HasChange(attr) {
			return diagError("Attribute %q is not supported for Gen2 databases: %s", attr, getGen2UnsupportedAttrGuidance(attr))
		}
	}

	// Special case: version changes are not supported
	if d.HasChange("version") {
		return diagError("Version changes are not supported for Gen2 database instances")
	}

	return nil
}

// applyGroupScalingWithDiagnostics applies group scaling and returns diagnostics.
// Wraps applyGroupScaling to provide consistent diagnostic handling.
func (g *resourceIBMDatabaseGen2Backend) applyGroupScalingWithDiagnostics(ctx context.Context, d *schema.ResourceData, rsConClient *rc.ResourceControllerV2, instanceID string, meta interface{}) diag.Diagnostics {
	if !d.HasChange("group") {
		return nil
	}

	instance, _, err := rsConClient.GetResourceInstance(&rc.GetResourceInstanceOptions{
		ID: &instanceID,
	})
	if err != nil {
		return diagError("error getting resource instance: %s", err)
	}

	configCtx := &instanceConfigContext{
		ctx:        ctx,
		d:          d,
		instanceID: instanceID,
		meta:       meta,
		instance:   instance,
	}

	if err := g.applyGroupScaling(configCtx); err != nil {
		return diagError("error applying group scaling: %s", err)
	}

	return nil
}

// Update modifies an existing IBM Cloud Database Gen2 instance.
// Supports updates to name, tags, and group scaling.
// Many features are not yet supported in Gen2 and will return errors if modified.
func (g *resourceIBMDatabaseGen2Backend) Update(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	warnings := g.WarnIgnoredAttrs(d)

	rsConClient, err := g.getResourceControllerClient(meta)
	if err != nil {
		return appendGen2DiagnosticsErrorsThenWarnings(diag.FromErr(err), warnings)
	}

	instanceID := d.Id()

	if diags := g.checkUnsupportedChanges(d); len(diags) > 0 {
		return appendGen2DiagnosticsErrorsThenWarnings(diags, warnings)
	}

	if diags := g.applyBasicAttributeUpdates(d, rsConClient, instanceID, meta); len(diags) > 0 {
		return appendGen2DiagnosticsErrorsThenWarnings(diags, warnings)
	}

	if diags := g.updateTagsWithDiagnostics(d, instanceID, meta); len(diags) > 0 {
		return appendGen2DiagnosticsErrorsThenWarnings(diags, warnings)
	}

	if diags := g.applyGroupScalingWithDiagnostics(ctx, d, rsConClient, instanceID, meta); len(diags) > 0 {
		return appendGen2DiagnosticsErrorsThenWarnings(diags, warnings)
	}

	readDiags := g.Read(ctx, d, meta)
	return appendGen2DiagnosticsErrorsThenWarnings(readDiags, warnings)
}

// Delete removes a database instance.
// Gen2 and Classic have same behavior for delete operations.
func (g *resourceIBMDatabaseGen2Backend) Delete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return databaseInstanceDelete(ctx, d, meta)
}

// Exists checks if a database instance exists.
// Gen2 and Classic have same behavior for exist check.
func (g *resourceIBMDatabaseGen2Backend) Exists(d *schema.ResourceData, meta interface{}) (bool, error) {
	return databaseInstanceExists(d, meta)
}

// WarnUnsupported returns a single grouped warning for Gen2 attributes that are accepted
// for backward compatibility but ignored.
func (g *resourceIBMDatabaseGen2Backend) WarnUnsupported(ctx context.Context, d *schema.ResourceData) diag.Diagnostics {
	return g.WarnIgnoredAttrs(d)
}

// WarnIgnoredAttrs returns a single grouped warning for Gen2 attributes that are accepted
// for backward compatibility but ignored.
func (g *resourceIBMDatabaseGen2Backend) WarnIgnoredAttrs(d *schema.ResourceData) diag.Diagnostics {
	var ignoredAttrs []string

	for _, attr := range gen2IgnoredAttrs {
		if val, ok := d.GetOk(attr); ok && !isEmptyGen2AttrValue(val) {
			ignoredAttrs = append(ignoredAttrs, attr)
		}
	}

	if len(ignoredAttrs) == 0 {
		return nil
	}

	return diag.Diagnostics{
		{
			Severity: diag.Warning,
			Summary:  "Some configured attributes are ignored for Gen2 databases",
			Detail:   buildGen2IgnoredAttrsWarningDetail(ignoredAttrs),
		},
	}
}

// isZeroValue checks if a value is the zero value for its type
func isZeroValue(val interface{}) bool {
	if val == nil {
		return true
	}

	switch v := val.(type) {
	case string:
		return v == ""
	case int, int8, int16, int32, int64:
		return v == 0
	case uint, uint8, uint16, uint32, uint64:
		return v == 0
	case float32, float64:
		return v == 0
	case bool:
		return !v
	case []interface{}:
		return len(v) == 0
	case map[string]interface{}:
		return len(v) == 0
	default:
		return false
	}
}

// ValidateUnsupportedAttrsDiff validates that unsupported attributes are not configured.
// Returns an error if any Gen2-unsupported attributes are set in the configuration.
// Also includes warnings about ignored attributes in the error message.
func (g *resourceIBMDatabaseGen2Backend) ValidateUnsupportedAttrsDiff(ctx context.Context, d *schema.ResourceDiff, meta interface{}) error {
	var unsupportedAttrs []string

	for _, attr := range gen2UnsupportedAttrs {
		if val, ok := d.GetOk(attr); ok && !isEmptyGen2AttrValue(val) {
			unsupportedAttrs = append(unsupportedAttrs, attr)
		}
	}

	if len(unsupportedAttrs) == 0 {
		return nil
	}

	var msg strings.Builder
	msg.WriteString("The following attributes are not supported for Gen2 databases:\n\n")

	for i, attr := range unsupportedAttrs {
		msg.WriteString(fmt.Sprintf("%d. Attribute: %q\n", i+1, attr))
		if guidance := getGen2UnsupportedAttrGuidance(attr); guidance != "" {
			msg.WriteString(strings.ReplaceAll(guidance, "\n", "\n  "))
			msg.WriteString("\n")
		}
		msg.WriteString("\n")
	}

	return errors.New(msg.String())
}

func (g *resourceIBMDatabaseGen2Backend) ValidateGroupsDiff(ctx context.Context, d *schema.ResourceDiff, meta interface{}) error {
	group, ok := d.GetOk("group")
	if !ok {
		return nil
	}

	groups := expandGroups(group.(*schema.Set).List())
	groupIDs := make([]string, 0, len(groups))
	for _, group := range groups {
		groupIDs = append(groupIDs, group.ID)
	}

	for i, id1 := range groupIDs {
		for j, id2 := range groupIDs {
			if id1 == id2 && i != j {
				return fmt.Errorf("found 2 or more instances of group with group_id %v", id1)
			}
		}
	}

	for _, group := range groups {
		if group == nil {
			continue
		}

		// Gen2 validation: Memory, CPU, and multitenant are not supported
		// Gen2 requires dedicated host flavors
		hasMemory := group.Memory != nil && group.Memory.Allocation > 0
		hasCPU := group.CPU != nil && group.CPU.Allocation > 0
		hasMultitenant := group.HostFlavor != nil && group.HostFlavor.ID == "multitenant"

		if hasMultitenant || hasMemory || hasCPU {
			errMsg := fmt.Sprintf("Configuration error: Gen2 databases do not support the following configuration(s) in group %q:\n\n", group.ID)

			if hasMultitenant {
				errMsg += "   - host_flavor.id: In Gen2 databases, host_flavor cannot be 'multitenant'. Choose a specific dedicated flavor (e.g., \"bx3d.4x20\").\n"
			}
			if hasMemory {
				errMsg += "   - memory: In Gen2 databases, memory allocation is determined by the 'host_flavor' attribute.\n"
			}
			if hasCPU {
				errMsg += "   - cpu: In Gen2 databases, CPU allocation is determined by the 'host_flavor' attribute.\n"
			}

			errMsg += "\n   Example:\n" +
				"     group {\n" +
				fmt.Sprintf("       group_id = %q\n", group.ID) +
				"       host_flavor {\n" +
				"         id = \"bx3d.4x20\"  # Use a dedicated flavor\n" +
				"       }\n" +
				"       disk {\n" +
				"         allocation_mb = 20480\n" +
				"       }\n" +
				"     }\n\n" +
				"   Documentation: https://registry.terraform.io/providers/IBM-Cloud/ibm/latest/docs/resources/database#host_flavor-2\n"

			return errors.New(errMsg)
		}

		if group.HostFlavor != nil && group.HostFlavor.ID != "" && group.HostFlavor.ID != "multitenant" {
			if err := validateGroupHostFlavor(group.ID, "host_flavor", group); err != nil {
				return err
			}
		}
	}

	return nil
}

func (g *resourceIBMDatabaseGen2Backend) ValidateServiceEndpointsDiff(ctx context.Context, d *schema.ResourceDiff, meta interface{}) error {
	serviceEndpoint, serviceEndpointOk := d.GetOk("service_endpoints")
	if serviceEndpointOk && serviceEndpoint.(string) != "" && serviceEndpoint.(string) != "private" {
		return fmt.Errorf("service_endpoints for Gen2 instances is optional, but if set it must be 'private'")
	}
	return nil
}

func (g *resourceIBMDatabaseGen2Backend) ValidateUnsupportedAttrsData(d *schema.ResourceData) error {
	var unsupportedAttrs []string

	for _, attr := range gen2UnsupportedAttrs {
		if val, ok := d.GetOk(attr); ok && !isEmptyGen2AttrValue(val) {
			unsupportedAttrs = append(unsupportedAttrs, attr)
		}
	}

	if len(unsupportedAttrs) == 0 {
		return nil
	}

	var msg strings.Builder
	msg.WriteString("The following attributes are not supported for Gen2 databases:\n\n")

	for i, attr := range unsupportedAttrs {
		msg.WriteString(fmt.Sprintf("%d. Attribute: %q\n", i+1, attr))
		if guidance := getGen2UnsupportedAttrGuidance(attr); guidance != "" {
			msg.WriteString(strings.ReplaceAll(guidance, "\n", "\n  "))
			msg.WriteString("\n")
		}
		msg.WriteString("\n")
	}

	return errors.New(msg.String())
}

func buildGen2IgnoredAttrsWarningDetail(attrs []string) string {
	var b strings.Builder

	b.WriteString("This database uses a Gen2 plan. Some attributes in this configuration are supported for Classic databases but are not implemented for Gen2 yet.\n\n")
	b.WriteString("Terraform will continue, but the following configured attributes will not be applied:\n\n")

	for _, attr := range attrs {
		b.WriteString(fmt.Sprintf("- %s", attr))

		if guidance := getGen2IgnoredAttrGuidance(attr); guidance != "" {
			b.WriteString(fmt.Sprintf(": %s", guidance))
		}

		b.WriteString("\n")
	}

	b.WriteString("\nWhat you can do:\n")
	b.WriteString("- Remove these attributes from the configuration to avoid this warning.\n")
	b.WriteString("- Check the Gen2 documentation for currently supported features.\n")

	return b.String()
}

func appendGen2DiagnosticsErrorsThenWarnings(errors diag.Diagnostics, warnings diag.Diagnostics) diag.Diagnostics {
	var out diag.Diagnostics
	out = append(out, errors...)
	out = append(out, warnings...)
	return out
}

func isEmptyGen2AttrValue(val interface{}) bool {
	if val == nil {
		return true
	}

	switch v := val.(type) {
	case string:
		return v == ""
	case bool:
		return !v
	case int:
		return v == 0
	case int64:
		return v == 0
	case float64:
		return v == 0
	case []interface{}:
		return len(v) == 0
	case map[string]interface{}:
		return len(v) == 0
	default:
		return false
	}
}
