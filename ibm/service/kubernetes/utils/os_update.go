package kubernetesutils

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
)

var (
	versionedOSRegexp = regexp.MustCompile(`^(UBUNTU|RHEL|REDHAT)_(\d+)_([a-zA-Z0-9]+)$`)
	InplaceOSUpdate   = customdiff.ForceNewIfChange("operating_system", func(ctx context.Context, oldValue, newValue, meta interface{}) bool {
		oldMatch := versionedOSRegexp.FindStringSubmatch(oldValue.(string))
		newMatch := versionedOSRegexp.FindStringSubmatch(newValue.(string))
		// Not a versioned OS, needs replacement
		if oldMatch == nil || newMatch == nil {
			return true
		}

		oldOS := oldMatch[1]
		newOS := newMatch[1]
		oldVersion := oldMatch[2]
		newVersion := newMatch[2]

		// Ubuntu can be upgraded in-place.
		if oldOS == "UBUNTU" && newOS == "UBUNTU" && oldVersion < newVersion {
			return false
		}

		// RHEL can be upgraded in-place.
		// NOTE: REDHAT might be used in older configs
		/*if (oldOS == "REDHAT" || oldOS == "RHEL") && (newOS == "REDHAT" || newOS == "RHEL") && oldVersion < newVersion {
			return false
		}*/

		return true
	})
)
