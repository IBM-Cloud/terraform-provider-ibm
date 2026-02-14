package database

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type resourceIbmDatabaseClassicBackend struct {
	meta interface{}
}

func newResourceIbmDatabaseClassicBackend(meta interface{}) resourceIbmDatabaseBackend {
	return &resourceIbmDatabaseClassicBackend{meta: meta}
}

func (c *resourceIbmDatabaseClassicBackend) WarnUnsupported(context context.Context, d *schema.ResourceData) diag.Diagnostics {
	return nil
}

func (c *resourceIbmDatabaseClassicBackend) Create(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return classicDatabaseInstanceCreate(context, d, meta)
}

func (c *resourceIbmDatabaseClassicBackend) Read(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return classicDatabaseInstanceRead(context, d, meta)
}

func (c *resourceIbmDatabaseClassicBackend) Update(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return classicDatabaseInstanceUpdate(context, d, meta)
}

func (c *resourceIbmDatabaseClassicBackend) Delete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return classicDatabaseInstanceDelete(context, d, meta)
}

func (c *resourceIbmDatabaseClassicBackend) Exists(d *schema.ResourceData, meta interface{}) (bool, error) {
	return classicDatabaseInstanceExists(d, meta)
}

func (c *resourceIbmDatabaseClassicBackend) ValidateUnsupportedAttrsDiff(context context.Context, d *schema.ResourceDiff, meta interface{}) error {
	return nil
}
