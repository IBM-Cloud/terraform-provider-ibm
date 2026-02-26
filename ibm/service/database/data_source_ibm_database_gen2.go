package database

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type dataSourceIBMDatabaseGen2Backend struct{}

func newDataSourceIBMDatabaseGen2Backend() dataSourceIBMDatabaseBackend {
	return &dataSourceIBMDatabaseGen2Backend{}
}

func (g *dataSourceIBMDatabaseGen2Backend) Read(d *schema.ResourceData, meta interface{}) error {
	// NOTE - Edge case: potential stale values for unsupported Gen2 attributes.
	// If this data source was previously resolved to a Classic instance, all
	// attributes (including ones not supported by Gen2) would have been set.
	// If the same filters later resolve to a Gen2 instance (e.g., name/location/service),
	// Terraform will not automatically clear attributes that are no longer set,
	// unlike a resource which would ForceNew on such changes.
	// As a result, the Gen2 read path may only set supported attributes while
	// previously populated Classic-only attributes remain stale in state.
	// There is no clean mechanism to fully reset datasource state, and doing so
	// is generally considered an anti-pattern.
	// If this becomes an issue, unsupported attributes could be explicitly set
	// to null via d.Set() to ensure stale values are cleared.
	return fmt.Errorf("gen2 backend not implemented yet")
}
