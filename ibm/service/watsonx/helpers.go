// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package watsonx

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/path"
)

func pathOf(name string) path.Path {
	return path.Root(name)
}

// splitImportID splits "<region>/<id>" import IDs.
func splitImportID(id string) (string, string, bool) {
	parts := strings.SplitN(id, "/", 2)
	if len(parts) != 2 || parts[0] == "" || parts[1] == "" {
		return "", "", false
	}
	return parts[0], parts[1], true
}
