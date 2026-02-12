package database

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type dbResourceBackend interface {
	Create(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics
	Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics
	Update(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics
	Delete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics
	Exists(d *schema.ResourceData, meta interface{}) (bool, error)

	// WarnUnsupported should emit warnings (diag.Warning) for fields you will ignore.
	// For classic backend this should return nil.
	WarnUnsupported(ctx context.Context, d *schema.ResourceData) diag.Diagnostics

	ValidateUnsupportedAttrsDiff(ctx context.Context, d *schema.ResourceDiff, meta interface{}) error
}

func isGen2Plan(plan string) bool {
	return strings.Contains(strings.ToLower(plan), "gen2")
}

func pickResourceBackend(d *schema.ResourceData, meta interface{}) dbResourceBackend {
	plan := d.Get("plan").(string)
	if isGen2Plan(plan) {
		return newGen2ResourceBackend(meta)
	}
	return newClassicResourceBackend(meta)
}

func pickResourceBackendFromDiff(d *schema.ResourceDiff, meta interface{}) dbResourceBackend {
	planRaw, ok := d.GetOk("plan")
	if !ok {
		// No plan yet; default to classic to avoid blocking planning unexpectedly.
		return newClassicResourceBackend(meta)
	}

	plan, ok := planRaw.(string)
	if !ok {
		return newClassicResourceBackend(meta)
	}

	if isGen2Plan(plan) {
		return newGen2ResourceBackend(meta)
	}
	return newClassicResourceBackend(meta)
}
