package database

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var classicUnsupportedAttrs = []string{
	"shards",
}

type resourceIBMDatabaseClassicBackend struct{}

func newResourceIBMDatabaseClassicBackend() resourceIBMDatabaseBackend {
	return &resourceIBMDatabaseClassicBackend{}
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
	return databaseInstanceDelete(context, d, meta)
}

func (c *resourceIBMDatabaseClassicBackend) Exists(d *schema.ResourceData, meta interface{}) (bool, error) {
	return databaseInstanceExists(d, meta)
}

func (c *resourceIBMDatabaseClassicBackend) WarnUnsupported(_ context.Context, _ *schema.ResourceData) diag.Diagnostics {
	// Unsupported attribute validation for Classic is handled at plan time by
	// validateUnsupportedAttrsDiffClassic (via ValidateUnsupportedAttrsDiff /
	// CustomizeDiff), so there is nothing to warn about at apply time.
	return nil
}

func (c *resourceIBMDatabaseClassicBackend) ValidateUnsupportedAttrsDiff(context context.Context, d *schema.ResourceDiff, meta interface{}) error {
	return validateUnsupportedAttrsDiffClassic(context, d, meta)
}

func (c *resourceIBMDatabaseClassicBackend) ValidateGroupsDiff(context context.Context, d *schema.ResourceDiff, meta interface{}) error {
	return validateGroupsDiffClassic(context, d, meta)
}

func (c *resourceIBMDatabaseClassicBackend) ValidateMemberZonesDiff(_ context.Context, d *schema.ResourceDiff, _ interface{}) error {
	// member_zones is Gen2-only. Block it at plan time using the shared traversal helper.
	return memberZonesInRawConfig(d, func(zones []string, _ int) error {
		return fmt.Errorf("'member_zones' is not supported for Classic database plans. " +
			"This attribute is only applicable to Gen2 plans.")
	})
}

func (c *resourceIBMDatabaseClassicBackend) ValidateServiceEndpointsDiff(context context.Context, d *schema.ResourceDiff, meta interface{}) error {
	return validateServiceEndpointsDiffClassic(context, d, meta)
}

func (c *resourceIBMDatabaseClassicBackend) ValidateShardsDiff(context context.Context, d *schema.ResourceDiff, meta interface{}) error {
	return nil
}
