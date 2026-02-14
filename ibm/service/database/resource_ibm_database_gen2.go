package database

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var gen2UnsupportedAttrs = []string{
	// TODO: update the list
	"backup_policy",
	"service_endpoints",
	"users",
	"groups",
}

type resourceIbmDatabaseGen2Backend struct {
	meta interface{}
}

func newResourceIbmDatabaseGen2Backend(meta interface{}) resourceIbmDatabaseBackend {
	return &resourceIbmDatabaseGen2Backend{meta: meta}
}

func (g *resourceIbmDatabaseGen2Backend) WarnUnsupported(context context.Context, d *schema.ResourceData) diag.Diagnostics {
	return nil
}

func (g *resourceIbmDatabaseGen2Backend) Create(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return diag.Errorf("gen2 backend not implemented yet (plan=%q)", d.Get("plan").(string))
}

func (g *resourceIbmDatabaseGen2Backend) Read(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return diag.Errorf("gen2 backend not implemented yet")
}

func (g *resourceIbmDatabaseGen2Backend) Update(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return diag.Errorf("gen2 backend not implemented yet")
}

func (g *resourceIbmDatabaseGen2Backend) Delete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return diag.Errorf("gen2 backend not implemented yet")
}

func (g *resourceIbmDatabaseGen2Backend) Exists(d *schema.ResourceData, meta interface{}) (bool, error) {
	return false, fmt.Errorf("gen2 backend not implemented yet")
}

func (g *resourceIbmDatabaseGen2Backend) ValidateUnsupportedAttrsDiff(context context.Context, d *schema.ResourceDiff, meta interface{}) error {
	var bad []string
	for _, k := range gen2UnsupportedAttrs {
		if isAttrConfiguredInDiff(d, k) {
			bad = append(bad, k)
		}
	}
	if len(bad) == 0 {
		return nil
	}

	planRaw, _ := d.GetOk("plan")
	plan, _ := planRaw.(string)

	return fmt.Errorf(
		"plan %q indicates Gen2. The following attributes are not supported for Gen2 and must be removed: %s",
		strings.TrimSpace(plan),
		strings.Join(bad, ", "),
	)
}
