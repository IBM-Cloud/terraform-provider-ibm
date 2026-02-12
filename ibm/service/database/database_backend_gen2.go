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

type gen2ResourceBackend struct {
	meta interface{}
}

func newGen2ResourceBackend(meta interface{}) dbResourceBackend {
	return &gen2ResourceBackend{meta: meta}
}

func (g *gen2ResourceBackend) WarnUnsupported(ctx context.Context, d *schema.ResourceData) diag.Diagnostics {
	return nil
}

func (g *gen2ResourceBackend) Create(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return diag.Errorf("gen2 backend not implemented yet (plan=%q)", d.Get("plan").(string))
}

func (g *gen2ResourceBackend) Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return diag.Errorf("gen2 backend not implemented yet")
}

func (g *gen2ResourceBackend) Update(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return diag.Errorf("gen2 backend not implemented yet")
}

func (g *gen2ResourceBackend) Delete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return diag.Errorf("gen2 backend not implemented yet")
}

func (g *gen2ResourceBackend) Exists(d *schema.ResourceData, meta interface{}) (bool, error) {
	return false, fmt.Errorf("gen2 backend not implemented yet")
}

func (g *gen2ResourceBackend) ValidateUnsupportedAttrsDiff(ctx context.Context, d *schema.ResourceDiff, meta interface{}) error {
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
