package database

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type dataSourceIbmDatabaseClassicBackend struct{}

func newDataSourceIbmDatabaseClassicBackend() dataSourceIbmDatabaseBackend {
	return &dataSourceIbmDatabaseClassicBackend{}
}

func (c *dataSourceIbmDatabaseClassicBackend) Read(d *schema.ResourceData, meta interface{}) error {
	return classicDataSourceIBMDatabaseInstanceRead(d, meta)
}
