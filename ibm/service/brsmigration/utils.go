// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package brsmigration

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func getBrsInstance(d *schema.ResourceData, parts ...string) string {
	if v, ok := d.GetOk("brs_instance"); ok && v.(string) != "" {
		return v.(string)
	}
	if v, ok := d.GetOk("brs_crn"); ok && v.(string) != "" {
		return v.(string)
	}
	for _, part := range parts {
		if strings.HasPrefix(part, "crn:") {
			return part
		}
	}
	if len(parts) >= 3 {
		return parts[0]
	}
	return ""
}
