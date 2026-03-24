package database

import (
	"fmt"
	"log"
	"net/url"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/platform-services-go-sdk/globalcatalogv1"
	rg "github.com/IBM/platform-services-go-sdk/resourcemanagerv2"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	// Constants for Gen2 database operations
	deploymentKind     = "deployment"
	dataservicesKey    = "dataservices"
	versionKey         = "version"
	resourcesKey       = "resources"
	platformOptionsKey = "platform_options"
	adminUserKey       = "adminuser"
	autoScalingKey     = "auto_scaling"
	allowlistKey       = "allowlist"
)

type dataSourceIBMDatabaseGen2Backend struct{}

func newDataSourceIBMDatabaseGen2Backend() dataSourceIBMDatabaseBackend {
	return &dataSourceIBMDatabaseGen2Backend{}
}

// Read retrieves and populates the state for a Gen2 database instance data source.
// It handles the migration scenario where a data source may have previously resolved
// to a Classic instance and now resolves to a Gen2 instance, ensuring stale Classic-only
// attributes are properly cleared from state.
//
// IMPORTANT: Gen2 Migration Edge Case
// =====================================
// When a data source previously resolved to a Classic instance and later resolves
// to a Gen2 instance (e.g., due to filter changes), Terraform does not automatically
// clear attributes that are no longer set. This differs from resources which would
// trigger ForceNew on such changes.
//
// Result: Gen2-unsupported attributes (adminuser, auto_scaling, allowlist) may
// persist with stale Classic values in state.
//
// Mitigation: These attributes are explicitly set to nil below to ensure state
// consistency. This is the recommended approach as fully resetting data source
// state is considered an anti-pattern in Terraform.
func (g *dataSourceIBMDatabaseGen2Backend) Read(d *schema.ResourceData, meta interface{}) error {
	// Find the database instance
	instance, err := findInstance(d, meta)
	if err != nil {
		return fmt.Errorf("failed to find database instance: %w", err)
	}
	if instance == nil || instance.ID == nil {
		return fmt.Errorf("database instance not found or missing ID")
	}
	d.SetId(*instance.ID)

	// Retrieve and set tags (non-critical operation, log errors but continue)
	tags, err := flex.GetTagsUsingCRN(meta, d.Id())
	if err != nil {
		log.Printf("[WARN] Failed to retrieve tags for database instance %s: %v", d.Id(), err)
	}
	d.Set("tags", tags)

	// Set basic instance attributes
	d.Set("name", instance.Name)
	d.Set("status", instance.State)
	d.Set("resource_group_id", instance.ResourceGroupID)
	d.Set("location", instance.RegionID)
	d.Set("guid", instance.GUID)

	// Retrieve service information from Global Catalog
	globalClient, err := meta.(conns.ClientSession).GlobalCatalogV1API()
	if err != nil {
		return fmt.Errorf("failed to get global catalog client: %w", err)
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

	// Set resource controller attributes
	d.Set(flex.ResourceName, instance.Name)
	d.Set(flex.ResourceCRN, instance.CRN)
	d.Set(flex.ResourceStatus, instance.State)

	// Retrieve resource group information
	rMgtClient, err := meta.(conns.ClientSession).ResourceManagerV2API()
	if err != nil {
		return fmt.Errorf("failed to get resource manager client: %w", err)
	}
	getResourceGroupOptions := rg.GetResourceGroupOptions{
		ID: instance.ResourceGroupID,
	}
	resourceGroup, resp, err := rMgtClient.GetResourceGroup(&getResourceGroupOptions)
	if err != nil || resourceGroup == nil {
		log.Printf("[WARN] Failed to retrieve resource group: %v %v", err, resp)
	}
	if resourceGroup != nil && resourceGroup.Name != nil {
		d.Set(flex.ResourceGroupName, resourceGroup.Name)
	}

	// Set resource controller URL
	rcontroller, err := flex.GetBaseController(meta)
	if err != nil {
		return fmt.Errorf("failed to get base controller: %w", err)
	}
	d.Set(flex.ResourceControllerURL, rcontroller+"/services/"+url.QueryEscape(*instance.CRN))

	// Clear Gen2-unsupported attributes to prevent stale Classic values
	// Admin user is not available in Gen2. Users should manage credentials using ibm_resource_key.
	d.Set(adminUserKey, nil)

	// Extract version from instance.Extensions based on database type
	var version string
	if instance.Extensions != nil {
		dbType := getDatabaseTypeFromResourceID(*instance.ResourceID)
		if dbType != "" {
			if dataservices, ok := instance.Extensions[dataservicesKey].(map[string]interface{}); ok {
				if dbTypeData, ok := dataservices[dbType].(map[string]interface{}); ok {
					if v, ok := dbTypeData[versionKey].(string); ok {
						version = v
					}
				}
			}
		}
	}
	d.Set(versionKey, version)

	// Extract platform_options from instance.Extensions for Gen2
	if instance.Extensions != nil {
		d.Set(platformOptionsKey, expandPlatformOptionsFromRCExtension(instance.Extensions))
	}

	// Get groups data from GlobalCatalog for Gen2
	// Find the deployment by getting plan's children and matching by location
	var deployment *globalcatalogv1.CatalogEntry
	kind := deploymentKind
	childOptions := globalcatalogv1.GetChildObjectsOptions{
		ID:   instance.ResourcePlanID,
		Kind: &kind,
	}
	children, _, err := globalClient.GetChildObjects(&childOptions)
	if err != nil {
		return fmt.Errorf("failed to retrieve plan children: %w", err)
	}

	if children != nil && children.Resources != nil {
		for _, child := range children.Resources {
			// Check if this deployment's location matches the instance region
			if child.Metadata != nil &&
				child.Metadata.Deployment != nil &&
				child.Metadata.Deployment.Location != nil &&
				*child.Metadata.Deployment.Location == *instance.RegionID {
				deployment = &child
				break
			}
		}
	}

	if deployment == nil {
		return fmt.Errorf("could not find deployment catalog entry for region %s", *instance.RegionID)
	}

	// Extract resources from deployment metadata
	var catalogResources []interface{}
	if deployment.Metadata != nil && deployment.Metadata.Other != nil {
		if resources, ok := deployment.Metadata.Other[resourcesKey].([]interface{}); ok {
			catalogResources = resources
		}
	}

	// Flatten groups using instance extensions and catalog metadata
	if instance.Extensions != nil && len(catalogResources) > 0 {
		d.Set("groups", flattenIcdGroupsFromInstanceAndCatalog(instance.Extensions, catalogResources, *instance.ResourceID))
	}

	// Clear additional Gen2-unsupported attributes
	// Auto scaling is currently not supported in Gen2
	d.Set(autoScalingKey, nil)

	// Allowlist is not supported in Gen2
	d.Set(allowlistKey, nil)

	return nil
}
