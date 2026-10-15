// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"context"
	"fmt"

	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const mysqlServiceName = "databases-for-mysql"

// backupsConfigIntent captures what the user explicitly wrote in the backups
// block. The backups schema is Optional+Computed, so the merged diff cannot be
// used to tell configured values from values retained from state.
type backupsConfigIntent struct {
	configured  bool
	preserve    *bool
	copyEnabled *bool
	targetCount int
}

// ctyFirstElem returns the first element of a known, non-null, non-empty list.
func ctyFirstElem(v cty.Value) (cty.Value, bool) {
	if v.IsNull() || !v.IsKnown() || !v.CanIterateElements() || v.LengthInt() == 0 {
		return cty.NilVal, false
	}
	elem := v.Index(cty.NumberIntVal(0))
	if elem.IsNull() || !elem.IsKnown() {
		return cty.NilVal, false
	}
	return elem, true
}

func ctyBoolPtr(v cty.Value) *bool {
	if v.IsNull() || !v.IsKnown() || v.Type() != cty.Bool {
		return nil
	}
	b := v.True()
	return &b
}

// backupsIntentFromRawConfig extracts the user's configured backups values from the raw config.
func backupsIntentFromRawConfig(raw cty.Value) backupsConfigIntent {
	var intent backupsConfigIntent
	if raw.IsNull() || !raw.IsKnown() || !raw.Type().IsObjectType() || !raw.Type().HasAttribute("backups") {
		return intent
	}
	backups := raw.GetAttr("backups")
	if backups.IsNull() || !backups.IsKnown() || !backups.CanIterateElements() || backups.LengthInt() == 0 {
		return intent
	}
	intent.configured = true

	block, ok := ctyFirstElem(backups)
	if !ok {
		return intent
	}
	intent.preserve = ctyBoolPtr(block.GetAttr("preserve"))

	auto, ok := ctyFirstElem(block.GetAttr("automatic_backups"))
	if !ok {
		return intent
	}
	cp, ok := ctyFirstElem(auto.GetAttr("copy"))
	if !ok {
		return intent
	}
	intent.copyEnabled = ctyBoolPtr(cp.GetAttr("enabled"))

	targets := cp.GetAttr("target")
	if !targets.IsNull() && targets.IsKnown() && targets.CanIterateElements() {
		intent.targetCount = targets.LengthInt()
	}
	return intent
}

// backupsRuleInput holds everything validateBackupsRules needs, so the rules can
// be unit tested without a real ResourceDiff.
type backupsRuleInput struct {
	gen2        bool
	service     string
	intent      backupsConfigIntent
	isUpdate    bool
	preserveOld bool
	preserveNew bool
	// preserveChanged is true when the preserve value differs between state and plan.
	preserveChanged bool
}

// validateBackupsRules applies the plan-time rules for the backups block.
func validateBackupsRules(in backupsRuleInput) error {
	if !in.intent.configured {
		return nil
	}

	if !in.gen2 {
		return fmt.Errorf("[ERROR] `backups` is only supported for Gen2 plans (for example standard-gen2). Remove the `backups` block for Classic plans")
	}

	if in.intent.preserve != nil && *in.intent.preserve && in.service == mysqlServiceName {
		return fmt.Errorf("[ERROR] `backups.preserve` cannot be enabled for %s", mysqlServiceName)
	}

	if in.isUpdate && in.preserveChanged && in.preserveOld && !in.preserveNew {
		return fmt.Errorf("[ERROR] `backups.preserve` cannot be disabled once it has been enabled")
	}

	if in.intent.copyEnabled != nil && !*in.intent.copyEnabled && in.intent.targetCount > 0 {
		return fmt.Errorf("[ERROR] `backups.automatic_backups.copy.target` cannot be set when `backups.automatic_backups.copy.enabled` is false")
	}

	return nil
}

// validateBackupsDiff is the CustomizeDiff entry point for the backups block.
func validateBackupsDiff(_ context.Context, d *schema.ResourceDiff, _ interface{}) error {
	intent := backupsIntentFromRawConfig(d.GetRawConfig())
	if !intent.configured {
		return nil
	}

	// Plan not yet known (for example derived from another resource): defer to apply.
	if !d.NewValueKnown("plan") {
		return nil
	}

	in := backupsRuleInput{
		gen2:     isGen2Plan(d.Get("plan").(string)),
		service:  d.Get("service").(string),
		intent:   intent,
		isUpdate: d.Id() != "",
	}

	if in.isUpdate && d.HasChange(backupsPreservePath) && d.NewValueKnown(backupsPreservePath) {
		oldV, newV := d.GetChange(backupsPreservePath)
		in.preserveOld, _ = oldV.(bool)
		in.preserveNew, _ = newV.(bool)
		in.preserveChanged = true
	}

	return validateBackupsRules(in)
}
