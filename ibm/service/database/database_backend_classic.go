package database

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type classicBackend struct {
	meta interface{}
}

func newClassicBackend(meta interface{}) dbBackend {
	return &classicBackend{meta: meta}
}

func (c *classicBackend) WarnUnsupported(ctx context.Context, d *schema.ResourceData) diag.Diagnostics {
	return nil
}

func (c *classicBackend) Create(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return classicDatabaseInstanceCreate(ctx, d, meta)
}

func (c *classicBackend) Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return classicDatabaseInstanceRead(ctx, d, meta)
}

func (c *classicBackend) Update(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return classicDatabaseInstanceUpdate(ctx, d, meta)
}

func (c *classicBackend) Delete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return classicDatabaseInstanceDelete(ctx, d, meta)
}

func (c *classicBackend) Exists(d *schema.ResourceData, meta interface{}) (bool, error) {
	return classicDatabaseInstanceExists(d, meta)
}

func (c *classicBackend) ValidateUnsupportedAttrsDiff(ctx context.Context, d *schema.ResourceDiff, meta interface{}) error {
	return nil
}
