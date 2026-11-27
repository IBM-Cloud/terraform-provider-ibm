// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"context"
	"fmt"

	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// isBackupsBlockConfigured reports whether the user wrote a backups block.
// The backups schema is Optional+Computed, so the raw config is used to tell a
// configured block from values retained in state.
func isBackupsBlockConfigured(raw cty.Value) bool {
	if raw.IsNull() || !raw.IsKnown() || !raw.Type().IsObjectType() || !raw.Type().HasAttribute("backups") {
		return false
	}
	backups := raw.GetAttr("backups")
	return !backups.IsNull() && backups.IsKnown() && backups.CanIterateElements() && backups.LengthInt() > 0
}

// validateBackupsDiff is the CustomizeDiff entry point for the backups block.
// backups is Gen2 only, so the plan fails when it is configured on a Classic plan.
func validateBackupsDiff(_ context.Context, d *schema.ResourceDiff, _ interface{}) error {
	if !isBackupsBlockConfigured(d.GetRawConfig()) {
		return nil
	}

	// Plan not yet known (for example derived from another resource): defer to apply.
	if !d.NewValueKnown("plan") {
		return nil
	}

	if !isGen2Plan(d.Get("plan").(string)) {
		return fmt.Errorf("[ERROR] `backups` is only supported for Gen2 plans (for example standard-gen2). Remove the `backups` block for Classic plans")
	}

	return nil
}
