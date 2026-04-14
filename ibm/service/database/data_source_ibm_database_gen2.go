package database

import (
	"fmt"
	"log"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/platform-services-go-sdk/globalcatalogv1"
	rc "github.com/IBM/platform-services-go-sdk/resourcecontrollerv2"
	rg "github.com/IBM/platform-services-go-sdk/resourcemanagerv2"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
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

	// Set basic attributes
	if err := g.setBasicAttributes(d, instance, meta); err != nil {
		return err
	}

	// Set service and plan information
	if err := g.setServiceInfo(d, instance, meta); err != nil {
		return err
	}

	// Set version information
	g.setVersionInfo(d, instance)

	// Set groups information
	if err := g.setGroupsInfo(d, instance, meta); err != nil {
		return err
	}

	// Clear Gen2 unsupported attributes
	g.clearUnsupportedAttributes(d)

	return nil
}

// setBasicAttributes sets basic instance attributes including tags, name, status, location, and resource controller attributes.
// Uses shared helper functions to reduce duplication with resource file.
func (g *dataSourceIBMDatabaseGen2Backend) setBasicAttributes(d *schema.ResourceData, instance *rc.ResourceInstance, meta interface{}) error {
	// Retrieve and set tags (non-critical operation, log errors but continue)
	if err := setTagsWithLogging(d, *instance.CRN, meta); err != nil {
		return err
	}

	// Set basic instance attributes
	d.Set("name", instance.Name)
	d.Set("status", instance.State)
	d.Set("resource_group_id", instance.ResourceGroupID)
	d.Set("location", instance.RegionID)
	d.Set("guid", instance.GUID)

	// Set resource controller attributes using shared helper
	if err := setResourceControllerAttributes(d, *instance.Name, *instance.CRN, *instance.State, meta); err != nil {
		return err
	}

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

	return nil
}

// setServiceInfo retrieves and sets service and plan information from Global Catalog.
// Clears admin user attribute as it's not available in Gen2.
func (g *dataSourceIBMDatabaseGen2Backend) setServiceInfo(d *schema.ResourceData, instance *rc.ResourceInstance, meta interface{}) error {
	// Retrieve service information from Global Catalog
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

// setVersionInfo extracts and sets version information from instance extensions.
// Also sets platform_options if available.
func (g *dataSourceIBMDatabaseGen2Backend) setVersionInfo(d *schema.ResourceData, instance *rc.ResourceInstance) {
	// Extract version from instance.Extensions based on database type
	version := extractVersionFromExtensions(instance.Extensions, *instance.ResourceID)
	d.Set(versionKey, version)

	// Extract platform_options from instance.Extensions for Gen2
	if instance.Extensions != nil {
		d.Set(platformOptionsKey, expandPlatformOptionsFromRCExtension(instance.Extensions))
	}
}

// setGroupsInfo retrieves and sets groups information from catalog.
// Combines instance extensions with catalog metadata to build group configurations.
func (g *dataSourceIBMDatabaseGen2Backend) setGroupsInfo(d *schema.ResourceData, instance *rc.ResourceInstance, meta interface{}) error {
	globalClient, err := getGlobalCatalogClient(meta)
	if err != nil {
		return err
	}

	// Get groups data from GlobalCatalog for Gen2
	// Find the deployment by getting plan's children and matching by location
	deployment, err := findDeploymentByLocation(globalClient, *instance.ResourcePlanID, *instance.RegionID)
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
	if instance.Extensions != nil && len(catalogResources) > 0 {
		d.Set("groups", flattenIcdGroupsFromInstanceAndCatalog(instance.Extensions, catalogResources, *instance.ResourceID))
	}

	return nil
}

// clearUnsupportedAttributes clears attributes not supported in Gen2.
// Sets auto_scaling and allowlist to nil to prevent stale Classic values.
func (g *dataSourceIBMDatabaseGen2Backend) clearUnsupportedAttributes(d *schema.ResourceData) {
	// Auto scaling is currently not supported in Gen2
	d.Set(autoScalingKey, nil)

	// Allowlist is not supported in Gen2
	d.Set(allowlistKey, nil)
}
