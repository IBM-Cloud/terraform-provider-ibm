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
	// TODO
	// Document clearly in the datasource docs which attributes are Classic-only.
	// In Gen2 Read:
	// do not set Classic-only string fields (adminuser, adminpassword) at all, or set them to null if your SDK path supports it safely.
	// set collection fields to empty lists/sets for stability if they’re not supported.

	// NOTE - EDGE CASE: Possibility of having stale values in unsupported Gen2 attributes
	// If the data source state was previously initialized with Classic, all attrib values would have set, including those not supported by Gen2
	// Now if the instance becomes gen2 (because, the filtering name/location/service etc matched to gen2)
	// unlike a TF resource which would ForceNew when plan changes; a datasource will not clear stale attribute values.
	// Now, Gen2 will call d.Set only for Gen2 valid attribs - which will result in data-source having stale values for those unsupported attributes
	// There is no direct way to "clean" everything in the state - it is also anti-patern.
	// If this becomes a real concern, the option we have is to set all those unsupported attributes explicity to `null` using d.Set()
	return fmt.Errorf("gen2 backend not implemented yet")
}
