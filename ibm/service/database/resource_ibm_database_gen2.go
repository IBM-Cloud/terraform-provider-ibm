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
	"github.com/IBM/platform-services-go-sdk/globalcatalogv1"
	rc "github.com/IBM/platform-services-go-sdk/resourcecontrollerv2"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var gen2UnsupportedAttrs = []string{
	// TODO: update the list
	"backup_policy",
	"users",
}

type resourceIBMDatabaseGen2Backend struct{}

func newResourceIBMDatabaseGen2Backend() resourceIBMDatabaseBackend {
	return &resourceIBMDatabaseGen2Backend{}
}

func (g *resourceIBMDatabaseGen2Backend) Create(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	rsConClient, err := meta.(conns.ClientSession).ResourceControllerV2API()
	if err != nil {
		return diag.FromErr(err)
	}

	serviceName := d.Get("service").(string)
	plan := d.Get("plan").(string)
	name := d.Get("name").(string)
	location := d.Get("location").(string)

	rsInst := rc.CreateResourceInstanceOptions{
		Name: &name,
	}

	rsCatClient, err := meta.(conns.ClientSession).ResourceCatalogAPI()
	if err != nil {
		return diag.FromErr(err)
	}
	rsCatRepo := rsCatClient.ResourceCatalog()

	serviceOff, err := rsCatRepo.FindByName(serviceName, true)
	if err != nil {
		return diag.FromErr(fmt.Errorf("[ERROR] Error retrieving database service offering: %s", err))
	}

	servicePlan, err := rsCatRepo.GetServicePlanID(serviceOff[0], plan)
	if err != nil {
		return diag.FromErr(fmt.Errorf("[ERROR] Error retrieving plan: %s", err))
	}
	rsInst.ResourcePlanID = &servicePlan

	deployments, err := rsCatRepo.ListDeployments(servicePlan)
	if err != nil {
		if serviceName == "databases-for-mongodb" && plan == "enterprise-sharding" {
			return diag.FromErr(fmt.Errorf("%s %s is not available yet in this region", serviceName, plan))
		} else {
			return diag.FromErr(fmt.Errorf("[ERROR] Error retrieving deployment for plan %s : %s", plan, err))
		}
	}
	if len(deployments) == 0 {
		return diag.FromErr(fmt.Errorf("[ERROR] No deployment found for service plan : %s", plan))
	}
	deployments, supportedLocations := filterDatabaseDeployments(deployments, location)

	if len(deployments) == 0 {
		locationList := make([]string, 0, len(supportedLocations))
		for l := range supportedLocations {
			locationList = append(locationList, l)
		}
		return diag.FromErr(fmt.Errorf("[ERROR] No deployment found for service plan %s at location %s.\nValid location(s) are: %q", plan, location, locationList))
	}
	catalogCRN := deployments[0].CatalogCRN
	rsInst.Target = &catalogCRN

	if rsGrpID, ok := d.GetOk("resource_group_id"); ok {
		rgID := rsGrpID.(string)
		rsInst.ResourceGroup = &rgID
	} else {
		defaultRg, err := flex.DefaultResourceGroup(meta)
		if err != nil {
			return diag.FromErr(err)
		}
		rsInst.ResourceGroup = &defaultRg
	}

	params := Params{}

	if version, ok := d.GetOk("version"); ok {
		params.Version = version.(string)
	}
	if keyProtect, ok := d.GetOk("key_protect_key"); ok {
		params.KeyProtectKey = keyProtect.(string)
	}
	if keyProtectInstance, ok := d.GetOk("key_protect_instance"); ok {
		params.KeyProtectInstance = keyProtectInstance.(string)
	}
	if backupID, ok := d.GetOk("backup_id"); ok {
		params.BackupID = backupID.(string)
	}
	if backUpEncryptionKey, ok := d.GetOk("backup_encryption_key_crn"); ok {
		params.BackUpEncryptionCRN = backUpEncryptionKey.(string)
	}
	if remoteLeader, ok := d.GetOk("remote_leader_id"); ok {
		params.RemoteLeaderID = remoteLeader.(string)
	}

	if pitrID, ok := d.GetOk("point_in_time_recovery_deployment_id"); ok {
		params.PITRDeploymentID = pitrID.(string)
	}

	pitrOk := !d.GetRawConfig().AsValueMap()["point_in_time_recovery_time"].IsNull()
	if pitrTime, ok := d.GetOk("point_in_time_recovery_time"); pitrOk {
		if !ok {
			pitrTime = ""
		}

		pitrTimeTrimmed := strings.TrimSpace(pitrTime.(string))
		params.PITRTimeStamp = &pitrTimeTrimmed
	}

	if offlineRestore, ok := d.GetOk("offline_restore"); ok {
		params.OfflineRestore = offlineRestore.(bool)
	}

	if asyncRestore, ok := d.GetOk("async_restore"); ok {
		params.AsyncRestore = asyncRestore.(bool)
	}

	var initialNodeCount int
	var sourceCRN string

	if params.PITRDeploymentID != "" {
		sourceCRN = params.PITRDeploymentID
	}

	if params.RemoteLeaderID != "" {
		sourceCRN = params.RemoteLeaderID
	}

	if sourceCRN != "" {
		group, err := getMemberGroup(sourceCRN, meta)

		if err != nil {
			return diag.FromErr(
				fmt.Errorf("[ERROR] Error fetching source formation group: %s", err)) // raise error
		}

		if group != nil {
			initialNodeCount = group.Members.Allocation
		}
	} else {
		initialNodeCount, err = getInitialNodeCount(serviceName, plan, meta)
		if err != nil {
			return diag.FromErr(err)
		}
	}

	var memberGroup *Group

	if group, ok := d.GetOk("group"); ok {
		groups := expandGroups(group.(*schema.Set).List())

		for _, g := range groups {
			if g.ID == "member" {
				memberGroup = g
				break
			}
		}
	}

	if memberGroup != nil && memberGroup.Memory != nil {
		params.Memory = memberGroup.Memory.Allocation * initialNodeCount
	}

	if memberGroup != nil && memberGroup.Disk != nil {
		params.Disk = memberGroup.Disk.Allocation * initialNodeCount
	}

	if memberGroup != nil && memberGroup.CPU != nil {
		params.CPU = memberGroup.CPU.Allocation * initialNodeCount
	}

	if memberGroup != nil && memberGroup.HostFlavor != nil {
		params.HostFlavor = memberGroup.HostFlavor.ID
	}

	serviceEndpoint := d.Get("service_endpoints").(string)
	params.ServiceEndpoints = serviceEndpoint
	parameters, _ := json.Marshal(params)
	var raw map[string]interface{}
	json.Unmarshal(parameters, &raw)
	//paramString := string(parameters[:])
	rsInst.Parameters = raw

	instance, response, err := rsConClient.CreateResourceInstance(&rsInst)
	if err != nil {
		return diag.FromErr(
			fmt.Errorf("[ERROR] Error creating database instance: %s %s", err, response))
	}
	d.SetId(*instance.ID)

	_, err = waitForDatabaseInstanceCreate(d, meta, *instance.ID)
	if err != nil {
		return diag.FromErr(
			fmt.Errorf(
				"[ERROR] Error waiting for create database instance (%s) to complete: %s", *instance.ID, err))
	}

	cloudDatabasesClient, err := meta.(conns.ClientSession).CloudDatabasesV5()
	if err != nil {
		return diag.FromErr(err)
	}

	if group, ok := d.GetOk("group"); ok {
		groups := expandGroups(group.(*schema.Set).List())
		groupsResponse, err := getGroups(*instance.ID, meta)
		if err != nil {
			return diag.FromErr(err)
		}
		currentGroups := normalizeGroups(groupsResponse)

		for _, g := range groups {
			groupScaling := &clouddatabasesv5.GroupScaling{}
			var currentGroup *Group
			var nodeCount int

			for _, cg := range currentGroups {
				if cg.ID == g.ID {
					currentGroup = &cg
					nodeCount = currentGroup.Members.Allocation
				}
			}

			if g.ID == "member" && (g.Members == nil || g.Members.Allocation == nodeCount) {
				// No Horizontal Scaling needed
				continue
			}

			if g.Members != nil && g.Members.Allocation != currentGroup.Members.Allocation {
				groupScaling.Members = &clouddatabasesv5.GroupScalingMembers{AllocationCount: core.Int64Ptr(int64(g.Members.Allocation))}
				nodeCount = g.Members.Allocation
			}
			if g.Memory != nil && g.Memory.Allocation*nodeCount != currentGroup.Memory.Allocation {
				groupScaling.Memory = &clouddatabasesv5.GroupScalingMemory{AllocationMb: core.Int64Ptr(int64(g.Memory.Allocation * nodeCount))}
			}
			if g.Disk != nil && g.Disk.Allocation*nodeCount != currentGroup.Disk.Allocation {
				groupScaling.Disk = &clouddatabasesv5.GroupScalingDisk{AllocationMb: core.Int64Ptr(int64(g.Disk.Allocation * nodeCount))}
			}
			if g.CPU != nil && g.CPU.Allocation*nodeCount != currentGroup.CPU.Allocation {
				groupScaling.CPU = &clouddatabasesv5.GroupScalingCPU{AllocationCount: core.Int64Ptr(int64(g.CPU.Allocation * nodeCount))}
			}
			if g.HostFlavor != nil {
				groupScaling.HostFlavor = &clouddatabasesv5.GroupScalingHostFlavor{ID: core.StringPtr(g.HostFlavor.ID)}
			}

			if groupScaling.Members != nil || groupScaling.Memory != nil || groupScaling.Disk != nil || groupScaling.CPU != nil || groupScaling.HostFlavor != nil {
				setDeploymentScalingGroupOptions := &clouddatabasesv5.SetDeploymentScalingGroupOptions{
					ID:      instance.ID,
					GroupID: &g.ID,
					Group:   groupScaling,
				}

				setDeploymentScalingGroupResponse, _, err := cloudDatabasesClient.SetDeploymentScalingGroup(setDeploymentScalingGroupOptions)

				taskIDLink := *setDeploymentScalingGroupResponse.Task.ID

				_, err = waitForDatabaseTaskComplete(taskIDLink, d, meta, d.Timeout(schema.TimeoutCreate))

				if err != nil {
					return diag.FromErr(err)
				}
			}
		}
	}

	v := os.Getenv("IC_ENV_TAGS")
	if _, ok := d.GetOk("tags"); ok || v != "" {
		oldList, newList := d.GetChange("tags")
		err = flex.UpdateTagsUsingCRN(oldList, newList, meta, *instance.CRN)
		if err != nil {
			log.Printf(
				"Error on create of ibm database (%s) tags: %s", d.Id(), err)
		}
	}

	instanceID := *instance.ID
	icdId := flex.EscapeUrlParm(instanceID)

	if pw, ok := d.GetOk("adminpassword"); ok {
		adminPassword := pw.(string)

		getDeploymentInfoOptions := &clouddatabasesv5.GetDeploymentInfoOptions{
			ID: core.StringPtr(instanceID),
		}
		getDeploymentInfoResponse, response, err := cloudDatabasesClient.GetDeploymentInfo(getDeploymentInfoOptions)

		if err != nil {
			if response.StatusCode == 404 {
				return diag.FromErr(fmt.Errorf("[ERROR] The database instance was not found in the region set for the Provider, or the default of us-south. Specify the correct region in the provider definition, or create a provider alias for the correct region. %v", err))
			}
			return diag.FromErr(fmt.Errorf("[ERROR] Error getting database config while updating adminpassword for: %s with error %s", instanceID, err))
		}
		deployment := getDeploymentInfoResponse.Deployment

		adminUser := deployment.AdminUsernames["database"]

		user := &clouddatabasesv5.UserUpdatePasswordSetting{
			Password: &adminPassword,
		}

		updateUserOptions := &clouddatabasesv5.UpdateUserOptions{
			ID:       core.StringPtr(instanceID),
			UserType: core.StringPtr("database"),
			Username: core.StringPtr(adminUser),
			User:     user,
		}

		updateUserResponse, response, err := cloudDatabasesClient.UpdateUser(updateUserOptions)
		if err != nil {
			return diag.FromErr(fmt.Errorf("[ERROR] UpdateUser (%s) failed %s\n%s", *updateUserOptions.Username, err, response))
		}

		taskID := *updateUserResponse.Task.ID
		_, err = waitForDatabaseTaskComplete(taskID, d, meta, d.Timeout(schema.TimeoutCreate))

		if err != nil {
			return diag.FromErr(fmt.Errorf("[ERROR] Error updating database admin password: %s", err))
		}
	}

	_, hasAllowlist := d.GetOk("allowlist")

	if hasAllowlist {
		var ipAddresses *schema.Set

		ipAddresses = d.Get("allowlist").(*schema.Set)

		entries := flex.ExpandAllowlist(ipAddresses)

		setAllowlistOptions := &clouddatabasesv5.SetAllowlistOptions{
			ID:          &instanceID,
			IPAddresses: entries,
		}

		setAllowlistResponse, _, err := cloudDatabasesClient.SetAllowlist(setAllowlistOptions)
		if err != nil {
			return diag.FromErr(fmt.Errorf("[ERROR] Error updating database allowlists: %s", err))
		}

		taskId := *setAllowlistResponse.Task.ID

		_, err = waitForDatabaseTaskComplete(taskId, d, meta, d.Timeout(schema.TimeoutCreate))
		if err != nil {
			return diag.FromErr(fmt.Errorf(
				"[ERROR] Error waiting for update of database (%s) allowlist task to complete: %s", instanceID, err))
		}
	}

	if _, ok := d.GetOk("auto_scaling.0"); ok {
		autoscalingSetGroupAutoscaling := &clouddatabasesv5.AutoscalingSetGroupAutoscaling{}

		if diskRecord, ok := d.GetOk("auto_scaling.0.disk"); ok {
			diskGroup, err := expandAutoscalingDiskGroup(d, diskRecord)
			if err != nil {
				return diag.FromErr(fmt.Errorf("[ERROR] Error in getting diskGroup from expandAutoscalingDiskGroup %s", err))
			}
			autoscalingSetGroupAutoscaling.Disk = diskGroup
		}

		if memoryRecord, ok := d.GetOk("auto_scaling.0.memory"); ok {
			memoryGroup, err := expandAutoscalingMemoryGroup(d, memoryRecord)
			if err != nil {
				return diag.FromErr(fmt.Errorf("[ERROR] Error in getting memoryBody from expandAutoscalingMemoryGroup %s", err))
			}

			autoscalingSetGroupAutoscaling.Memory = memoryGroup
		}

		if autoscalingSetGroupAutoscaling.Disk != nil || autoscalingSetGroupAutoscaling.Memory != nil {
			setAutoscalingConditionsOptions := &clouddatabasesv5.SetAutoscalingConditionsOptions{
				ID:          &instanceID,
				GroupID:     core.StringPtr("member"),
				Autoscaling: autoscalingSetGroupAutoscaling,
			}

			setAutoscalingConditionsResponse, _, err := cloudDatabasesClient.SetAutoscalingConditions(setAutoscalingConditionsOptions)
			if err != nil {
				return diag.FromErr(fmt.Errorf("[ERROR] Error updating database auto_scaling: %s", err))
			}

			taskId := *setAutoscalingConditionsResponse.Task.ID

			_, err = waitForDatabaseTaskComplete(taskId, d, meta, d.Timeout(schema.TimeoutCreate))
			if err != nil {
				return diag.FromErr(fmt.Errorf("[ERROR] Error waiting for database (%s) memory auto_scaling group update task to complete: %s", instanceID, err))
			}
		}
	}

	if userList, ok := d.GetOk("users"); ok {
		if err != nil {
			return diag.FromErr(fmt.Errorf("[ERROR] Error getting database client settings: %s", err))
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
				return diag.FromErr(err)
			}
		}
	}

	if config, ok := d.GetOk("configuration"); ok {
		var rawConfig map[string]json.RawMessage
		err = json.Unmarshal([]byte(config.(string)), &rawConfig)
		if err != nil {
			return diag.FromErr(fmt.Errorf("[ERROR] configuration JSON invalid\n%s", err))
		}

		var configuration clouddatabasesv5.ConfigurationIntf = new(clouddatabasesv5.Configuration)
		err = core.UnmarshalModel(rawConfig, "", &configuration, clouddatabasesv5.UnmarshalConfiguration)
		if err != nil {
			return diag.FromErr(fmt.Errorf("[ERROR] database configuration is invalid"))
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

	if _, ok := d.GetOk("logical_replication_slot"); ok {
		service := d.Get("service").(string)
		if service != "databases-for-postgresql" {
			return diag.FromErr(fmt.Errorf("[ERROR] Error Logical Replication can only be set for databases-for-postgresql instances"))
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

			createLogicalRepSlotResponse, response, err := cloudDatabasesClient.CreateLogicalReplicationSlot(createLogicalReplicationOptions)
			if err != nil {
				return diag.FromErr(fmt.Errorf("[ERROR] CreateLogicalReplicationSlot (%s) failed %s\n%s", *createLogicalReplicationOptions.LogicalReplicationSlot.Name, err, response))
			}

			taskID := *createLogicalRepSlotResponse.Task.ID
			_, err = waitForDatabaseTaskComplete(taskID, d, meta, d.Timeout(schema.TimeoutUpdate))
			if err != nil {
				return diag.FromErr(fmt.Errorf(
					"[ERROR] Error waiting for database (%s) logical replication slot (%s) create task to complete: %s", instanceID, *createLogicalReplicationOptions.LogicalReplicationSlot.Name, err))
			}
		}
	}

	return resourceIBMDatabaseInstanceRead(context, d, meta)
}

func (g *resourceIBMDatabaseGen2Backend) Read(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	rsConClient, err := meta.(conns.ClientSession).ResourceControllerV2API()
	if err != nil {
		return diag.FromErr(err)
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
		return diag.FromErr(fmt.Errorf("[ERROR] Error retrieving resource instance: %s %s", err, response))
	}
	if strings.Contains(*instance.State, "removed") {
		log.Printf("[WARN] Removing instance from TF state because it's now in removed state")
		d.SetId("")
		return nil
	}

	tags, err := flex.GetTagsUsingCRN(meta, *instance.CRN)
	if err != nil {
		log.Printf(
			"Error on get of ibm Database tags (%s) tags: %s", d.Id(), err)
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
		return diag.FromErr(err)
	}
	d.Set(flex.ResourceControllerURL, rcontroller+"/services/"+url.QueryEscape(*instance.CRN))

	rsCatClient, err := meta.(conns.ClientSession).ResourceCatalogAPI()
	if err != nil {
		return diag.FromErr(err)
	}
	rsCatRepo := rsCatClient.ResourceCatalog()

	serviceOff, err := rsCatRepo.GetServiceName(*instance.ResourceID)
	if err != nil {
		return diag.FromErr(fmt.Errorf("[ERROR] Error retrieving service offering: %s", err))
	}

	d.Set("service", serviceOff)

	servicePlan, err := rsCatRepo.GetServicePlanName(*instance.ResourcePlanID)
	if err != nil {
		return diag.FromErr(fmt.Errorf("[ERROR] Error retrieving plan: %s", err))
	}
	d.Set("plan", servicePlan)

	// Admin user is not available in Gen2. Users should manage credentials using ibm_resource_key.
	// Clear it from state if it was previously set (e.g., if the state was carried forward from a Classic instance).
	d.Set("adminuser", nil)

	// Extract version from instance.Extensions based on database type
	var version string
	if instance.Extensions != nil {
		dbType := getDatabaseTypeFromResourceID(*instance.ResourceID)
		if dbType != "" {
			if dataservices, ok := instance.Extensions["dataservices"].(map[string]interface{}); ok {
				if dbTypeData, ok := dataservices[dbType].(map[string]interface{}); ok {
					if v, ok := dbTypeData["version"].(string); ok {
						version = v
					}
				}
			}
		}
	}
	d.Set("version", version)

	// Get groups data from GlobalCatalog for Gen2
	// Find the deployment by getting plan's children and matching by location
	globalClient, err := meta.(conns.ClientSession).GlobalCatalogV1API()
	if err != nil {
		return diag.FromErr(err)
	}

	var catalogDeployment *globalcatalogv1.CatalogEntry
	kind := "deployment"
	childOptions := globalcatalogv1.GetChildObjectsOptions{
		ID:   instance.ResourcePlanID,
		Kind: &kind,
	}
	children, _, err := globalClient.GetChildObjects(&childOptions)
	if err != nil {
		return diag.FromErr(fmt.Errorf("[ERROR] Error retrieving plan children: %s", err))
	}

	if children != nil && children.Resources != nil {
		for _, child := range children.Resources {
			// Check if this deployment's location matches the instance region
			if child.Metadata != nil &&
				child.Metadata.Deployment != nil &&
				child.Metadata.Deployment.Location != nil &&
				*child.Metadata.Deployment.Location == instanceLocation {
				catalogDeployment = &child
				break
			}
		}
	}

	if catalogDeployment == nil {
		return diag.FromErr(fmt.Errorf("[ERROR] Could not find deployment catalog entry for region %s", instanceLocation))
	}

	// Extract resources from deployment metadata
	var catalogResources []interface{}
	if catalogDeployment.Metadata != nil && catalogDeployment.Metadata.Other != nil {
		if resources, ok := catalogDeployment.Metadata.Other["resources"].([]interface{}); ok {
			catalogResources = resources
		}
	}

	// Flatten groups using instance extensions and catalog metadata
	if instance.Extensions != nil && len(catalogResources) > 0 {
		d.Set("groups", flattenIcdGroupsFromInstanceAndCatalog(instance.Extensions, catalogResources, *instance.ResourceID))
	}

	// Auto scaling is currently not supported in Gen2. Clear it from state if it was previously set
	// (e.g., if the state was carried forward from a Classic instance).
	d.Set("auto_scaling", nil)

	// Allowlist is not supported in Gen2. Clear it from state if it was previously set
	// (e.g., if the state was carried forward from a Classic instance).
	d.Set("allowlist", nil)

	// Users are not managed in Gen2 via this resource. Users should manage credentials using ibm_resource_key.
	// Clear it from state if it was previously set (e.g., if the state was carried forward from a Classic instance).
	d.Set("users", nil)

	// Configuration schema is currently not supported in Gen2. Clear it from state if it was previously set
	// (e.g., if the state was carried forward from a Classic instance).
	d.Set("configuration_schema", nil)

	return nil

}

func (g *resourceIBMDatabaseGen2Backend) Update(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return diag.Errorf("gen2 backend not implemented yet")
}

func (g *resourceIBMDatabaseGen2Backend) Delete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return diag.Errorf("gen2 backend not implemented yet")
}

func (g *resourceIBMDatabaseGen2Backend) Exists(d *schema.ResourceData, meta interface{}) (bool, error) {
	return false, fmt.Errorf("gen2 backend not implemented yet")
}

func (g *resourceIBMDatabaseGen2Backend) WarnUnsupported(context context.Context, d *schema.ResourceData) diag.Diagnostics {
	return nil
}

func (g *resourceIBMDatabaseGen2Backend) ValidateUnsupportedAttrsDiff(context context.Context, d *schema.ResourceDiff, meta interface{}) error {
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
