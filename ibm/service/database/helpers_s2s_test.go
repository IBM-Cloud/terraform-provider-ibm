// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// ---------------------------------------------------------------------------
// isTruthy
// ---------------------------------------------------------------------------

func TestIsTruthy(t *testing.T) {
	cases := []struct {
		name string
		in   interface{}
		want bool
	}{
		{"bool true", true, true},
		{"bool false", false, false},
		{"string true", "true", true},
		{"string false", "false", false},
		{"string TRUE (case sensitive)", "TRUE", false},
		{"nil", nil, false},
		{"int 1", 1, false},
		{"empty string", "", false},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			require.Equal(t, c.want, isTruthy(c.in))
		})
	}
}

// ---------------------------------------------------------------------------
// s2sAuthWarning sentinel error
// ---------------------------------------------------------------------------

func TestS2sAuthWarningError(t *testing.T) {
	w := &s2sAuthWarning{}
	require.Equal(t, s2sAuthWarningHeader, w.Error())
}

// ---------------------------------------------------------------------------
// checkS2SAuthorization
// ---------------------------------------------------------------------------

func TestCheckS2SAuthorization(t *testing.T) {
	cases := []struct {
		name       string
		extensions map[string]interface{}
		want       bool
	}{
		{
			name:       "nil extensions",
			extensions: nil,
			want:       false,
		},
		{
			name:       "empty extensions map",
			extensions: map[string]interface{}{},
			want:       false,
		},
		{
			name: "authorizations key missing",
			extensions: map[string]interface{}{
				"other_key": "value",
			},
			want: false,
		},
		{
			name: "authorizations is nil",
			extensions: map[string]interface{}{
				"authorizations": nil,
			},
			want: false,
		},
		{
			name: "authorizations is empty map",
			extensions: map[string]interface{}{
				"authorizations": map[string]interface{}{},
			},
			want: false,
		},
		{
			name: "both flags true (bool)",
			extensions: map[string]interface{}{
				"authorizations": map[string]interface{}{
					"independent_backups": true,
					"resource_group":      true,
				},
			},
			want: true,
		},
		{
			name: "both flags true (string)",
			extensions: map[string]interface{}{
				"authorizations": map[string]interface{}{
					"independent_backups": "true",
					"resource_group":      "true",
				},
			},
			want: true,
		},
		{
			name: "independent_backups false (bool)",
			extensions: map[string]interface{}{
				"authorizations": map[string]interface{}{
					"independent_backups": false,
					"resource_group":      true,
				},
			},
			want: false,
		},
		{
			name: "resource_group false (bool)",
			extensions: map[string]interface{}{
				"authorizations": map[string]interface{}{
					"independent_backups": true,
					"resource_group":      false,
				},
			},
			want: false,
		},
		{
			name: "both flags false (bool)",
			extensions: map[string]interface{}{
				"authorizations": map[string]interface{}{
					"independent_backups": false,
					"resource_group":      false,
				},
			},
			want: false,
		},
		{
			name: "independent_backups missing",
			extensions: map[string]interface{}{
				"authorizations": map[string]interface{}{
					"resource_group": true,
				},
			},
			want: false,
		},
		{
			name: "resource_group missing",
			extensions: map[string]interface{}{
				"authorizations": map[string]interface{}{
					"independent_backups": true,
				},
			},
			want: false,
		},
		{
			name: "mixed: bool true + string true",
			extensions: map[string]interface{}{
				"authorizations": map[string]interface{}{
					"independent_backups": true,
					"resource_group":      "true",
				},
			},
			want: true,
		},
		{
			name: "resource_group string false",
			extensions: map[string]interface{}{
				"authorizations": map[string]interface{}{
					"independent_backups": true,
					"resource_group":      "false",
				},
			},
			want: false,
		},
		{
			name: "authorizations is not a map (wrong type)",
			extensions: map[string]interface{}{
				"authorizations": "not-a-map",
			},
			want: false,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			require.Equal(t, c.want, checkS2SAuthorization(c.extensions))
		})
	}
}

// ---------------------------------------------------------------------------
// s2sAuthWarningHeader / s2sAuthWarningDetail constants are non-empty
// ---------------------------------------------------------------------------

func TestS2SWarningConstants(t *testing.T) {
	require.NotEmpty(t, s2sAuthWarningHeader, "s2sAuthWarningHeader must not be empty")
	require.NotEmpty(t, s2sAuthWarningDetail, "s2sAuthWarningDetail must not be empty")
}
