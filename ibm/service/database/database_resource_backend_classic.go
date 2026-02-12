package database

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type classicResourceBackend struct {
	meta interface{}
}

func newClassicResourceBackend(meta interface{}) dbResourceBackend {
	return &classicResourceBackend{meta: meta}
}

func (c *classicResourceBackend) WarnUnsupported(ctx context.Context, d *schema.ResourceData) diag.Diagnostics {
	return nil
}

func (c *classicResourceBackend) Create(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return classicDatabaseInstanceCreate(ctx, d, meta)
}

func (c *classicResourceBackend) Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return classicDatabaseInstanceRead(ctx, d, meta)
}

func (c *classicResourceBackend) Update(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return classicDatabaseInstanceUpdate(ctx, d, meta)
}

func (c *classicResourceBackend) Delete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return classicDatabaseInstanceDelete(ctx, d, meta)
}

func (c *classicResourceBackend) Exists(d *schema.ResourceData, meta interface{}) (bool, error) {
	return classicDatabaseInstanceExists(d, meta)
}

func (c *classicResourceBackend) ValidateUnsupportedAttrsDiff(ctx context.Context, d *schema.ResourceDiff, meta interface{}) error {
	return nil
}
