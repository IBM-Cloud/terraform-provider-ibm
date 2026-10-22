// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package codeengine

import (
	"testing"
)

func TestParseCodeEngineQuantity(t *testing.T) {
	cases := []struct {
		input string
		want  float64
	}{
		{"500M", 500e6},
		{"0.5G", 0.5e9},
		{"1G", 1e9},
		{"400M", 400e6},
		{"4G", 4e9},
		{"2G", 2e9},
		{"250M", 250e6},
		{"0.25G", 0.25e9},
		// Invalid inputs should return -1
		{"", -1},
		{"500", -1},
		{"G", -1},
		{"500K", -1},
		{"abc", -1},
	}

	for _, tc := range cases {
		got := parseCodeEngineQuantity(tc.input)
		if got != tc.want {
			t.Errorf("parseCodeEngineQuantity(%q) = %v, want %v", tc.input, got, tc.want)
		}
	}
}

func TestSuppressCodeEngineQuantityDiff(t *testing.T) {
	cases := []struct {
		old      string
		new      string
		suppress bool
	}{
		// Equivalent values in different units — must suppress
		{"500M", "0.5G", true},
		{"0.5G", "500M", true},
		{"1G", "1000M", true},
		{"400M", "0.4G", true},
		{"4G", "4000M", true},
		{"250M", "0.25G", true},
		// Identical strings — must suppress
		{"500M", "500M", true},
		{"1G", "1G", true},
		// Different values — must NOT suppress
		{"500M", "1G", false},
		{"1G", "2G", false},
		{"400M", "500M", false},
		// Unparseable values fall back to literal comparison
		{"500M", "500", false},
		{"500", "500", true},
		{"", "", true},
		{"500M", "", false},
	}

	for _, tc := range cases {
		got := SuppressCodeEngineQuantityDiff("", tc.old, tc.new, nil)
		if got != tc.suppress {
			t.Errorf("SuppressCodeEngineQuantityDiff(%q, %q) = %v, want %v", tc.old, tc.new, got, tc.suppress)
		}
	}
}
