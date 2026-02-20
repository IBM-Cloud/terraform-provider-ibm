package database

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type resourceIBMDatabaseClassicBackend struct{}

func newResourceIBMDatabaseClassicBackend() resourceIBMDatabaseBackend {
	return &resourceIBMDatabaseClassicBackend{}
}

func (c *resourceIBMDatabaseClassicBackend) WarnUnsupported(context context.Context, d *schema.ResourceData) diag.Diagnostics {
	return nil
}

func (c *resourceIBMDatabaseClassicBackend) Create(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return classicDatabaseInstanceCreate(context, d, meta)
}

func (c *resourceIBMDatabaseClassicBackend) Read(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return classicDatabaseInstanceRead(context, d, meta)
}

func (c *resourceIBMDatabaseClassicBackend) Update(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return classicDatabaseInstanceUpdate(context, d, meta)
}

func (c *resourceIBMDatabaseClassicBackend) Delete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return classicDatabaseInstanceDelete(context, d, meta)
}

func (c *resourceIBMDatabaseClassicBackend) Exists(d *schema.ResourceData, meta interface{}) (bool, error) {
	return classicDatabaseInstanceExists(d, meta)
}

func (c *resourceIBMDatabaseClassicBackend) ValidateUnsupportedAttrsDiff(context context.Context, d *schema.ResourceDiff, meta interface{}) error {
	return nil
}
