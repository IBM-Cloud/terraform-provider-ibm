package database

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type dataSourceIBMDatabaseClassicBackend struct{}

func newDataSourceIBMDatabaseClassicBackend() dataSourceIBMDatabaseBackend {
	return &dataSourceIBMDatabaseClassicBackend{}
}

func (c *dataSourceIBMDatabaseClassicBackend) Read(d *schema.ResourceData, meta interface{}) error {
	return classicDataSourceIBMDatabaseInstanceRead(d, meta)
}
