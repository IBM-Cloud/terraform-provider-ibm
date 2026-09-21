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

func (c *resourceIBMDatabaseClassicBackend) WarnUnsupported(_ context.Context, d *schema.ResourceData) diag.Diagnostics {
	if group, ok := d.GetOk("group"); ok {
		for _, groupRaw := range group.(*schema.Set).List() {
			if zones, _, ok := memberZonesFromDiff(groupRaw); ok && len(zones) > 0 {
				return diag.Diagnostics{{
					Severity: diag.Warning,
					Summary:  "member_zones is not supported for Classic database plans",
					Detail:   "The member_zones attribute is only applicable to Gen2 database plans. It will be ignored for Classic plans.",
				}}
			}
		}
	}
	return nil
}

func (c *resourceIBMDatabaseClassicBackend) ValidateUnsupportedAttrsDiff(context context.Context, d *schema.ResourceDiff, meta interface{}) error {
	return nil
}

func (c *resourceIBMDatabaseClassicBackend) ValidateGroupsDiff(context context.Context, d *schema.ResourceDiff, meta interface{}) error {
	return validateGroupsDiffClassic(context, d, meta)
}

func (c *resourceIBMDatabaseClassicBackend) ValidateMemberZonesDiff(_ context.Context, _ *schema.ResourceDiff, _ interface{}) error {
	return nil // member_zones is Gen2 only
}

func (c *resourceIBMDatabaseClassicBackend) ValidateServiceEndpointsDiff(context context.Context, d *schema.ResourceDiff, meta interface{}) error {
	return validateServiceEndpointsDiffClassic(context, d, meta)
}
