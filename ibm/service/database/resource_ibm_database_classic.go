package database

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// classicUnsupportedAttrs are attributes only available on Gen2 plans.
// Setting any of these on a Classic plan produces a plan-time error.
var classicUnsupportedAttrs = []string{
	"maintenance",
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

func (c *resourceIBMDatabaseClassicBackend) WarnUnsupported(context context.Context, d *schema.ResourceData) diag.Diagnostics {
	return nil
}

func (c *resourceIBMDatabaseClassicBackend) ValidateUnsupportedAttrsDiff(context context.Context, d *schema.ResourceDiff, meta interface{}) error {
	if d == nil {
		return nil
	}
	for _, attr := range classicUnsupportedAttrs {
		if val, ok := d.GetOk(attr); ok && !isEmptyGen2AttrValue(val) {
			return fmt.Errorf("attribute %q is only supported for Gen2 database plans and cannot be used with Classic plans", attr)
		}
	}
	return nil
}

func (c *resourceIBMDatabaseClassicBackend) ValidateGroupsDiff(context context.Context, d *schema.ResourceDiff, meta interface{}) error {
	return validateGroupsDiffClassic(context, d, meta)
}

func (c *resourceIBMDatabaseClassicBackend) ValidateServiceEndpointsDiff(context context.Context, d *schema.ResourceDiff, meta interface{}) error {
	return validateServiceEndpointsDiffClassic(context, d, meta)
}

func (c *resourceIBMDatabaseClassicBackend) ValidateMaintenanceWindowDiff(_ context.Context, _ *schema.ResourceDiff, _ interface{}) error {
	return nil
}
