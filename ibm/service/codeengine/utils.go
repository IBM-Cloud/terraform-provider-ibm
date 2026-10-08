// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package codeengine

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// codeEngineMemorySuffixBytes maps the unit suffix accepted by Code Engine to
// its byte multiplier.  Code Engine accepts M (Megabyte) and G (Gigabyte) as
// the only suffixes for memory / ephemeral-storage limits, where M = 10^6 and
// G = 10^9 (SI, not binary).
var codeEngineMemorySuffixBytes = map[string]float64{
	"M": 1e6,
	"G": 1e9,
}

// quantityRe matches a Code Engine size string such as "500M", "0.5G", "1G".
var quantityRe = regexp.MustCompile(`^([0-9]+(?:\.[0-9]*)?)([MG])$`)

// parseCodeEngineQuantity converts a Code Engine size string (e.g. "500M",
// "0.5G") to the equivalent number of bytes.  It returns -1 when the string
// does not match the expected format.
func parseCodeEngineQuantity(s string) float64 {
	s = strings.TrimSpace(s)
	m := quantityRe.FindStringSubmatch(s)
	if m == nil {
		return -1
	}
	value, err := strconv.ParseFloat(m[1], 64)
	if err != nil {
		return -1
	}
	multiplier, ok := codeEngineMemorySuffixBytes[m[2]]
	if !ok {
		return -1
	}
	return value * multiplier
}

// SuppressCodeEngineQuantityDiff is a Terraform DiffSuppressFunc that treats
// two Code Engine size strings as equal when they represent the same number of
// bytes (e.g. "500M" == "0.5G").  This prevents spurious in-place updates that
// would otherwise be triggered because the IBM Cloud API normalises the stored
// value to a different but equivalent representation (e.g. "0.5G" → "500M").
func SuppressCodeEngineQuantityDiff(_, oldVal, newVal string, _ *schema.ResourceData) bool {
	oldBytes := parseCodeEngineQuantity(oldVal)
	newBytes := parseCodeEngineQuantity(newVal)
	// If either value cannot be parsed fall back to a literal comparison so
	// that an unexpected format is always treated as a real change.
	if oldBytes < 0 || newBytes < 0 {
		return oldVal == newVal
	}
	return oldBytes == newBytes
}
