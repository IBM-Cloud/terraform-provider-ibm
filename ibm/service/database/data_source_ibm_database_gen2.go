package database

import (
	"fmt"

	rc "github.com/IBM/platform-services-go-sdk/resourcecontrollerv2"

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
	// Use shared Gen2 helper function
	// Data sources don't need service_endpoints or resource_controller_url
	return setGen2BasicAttributes(d, instance, meta, false, false)
}

// setServiceInfo retrieves and sets service and plan information from Global Catalog.
// Clears admin user attribute as it's not available in Gen2.
func (g *dataSourceIBMDatabaseGen2Backend) setServiceInfo(d *schema.ResourceData, instance *rc.ResourceInstance, meta interface{}) error {
	// Use shared Gen2 helper function
	return setGen2ServiceInfo(d, instance, meta)
}

// setVersionInfo extracts and sets version information from instance extensions.
// Also sets platform_options if available.
func (g *dataSourceIBMDatabaseGen2Backend) setVersionInfo(d *schema.ResourceData, instance *rc.ResourceInstance) {
	// Use shared Gen2 helper function
	// Data sources include platform_options
	setGen2VersionInfo(d, instance, true)
}

// setGroupsInfo retrieves and sets groups information from catalog.
// Combines instance extensions with catalog metadata to build group configurations.
func (g *dataSourceIBMDatabaseGen2Backend) setGroupsInfo(d *schema.ResourceData, instance *rc.ResourceInstance, meta interface{}) error {
	// Use shared Gen2 helper function
	return setGen2GroupsInfo(d, instance, meta)
}

// clearUnsupportedAttributes clears attributes not supported in Gen2.
// Sets auto_scaling and allowlist to nil to prevent stale Classic values.
func (g *dataSourceIBMDatabaseGen2Backend) clearUnsupportedAttributes(d *schema.ResourceData) {
	// Use shared Gen2 helper function
	clearGen2UnsupportedAttributes(d)
}
