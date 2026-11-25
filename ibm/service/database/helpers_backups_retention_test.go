// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/stretchr/testify/require"
)

func retentionTestResourceData(t *testing.T, backups []interface{}) *schema.ResourceData {
	t.Helper()
	raw := map[string]interface{}{
		"name":     "test-gen2-db",
		"location": "us-south",
		"service":  "databases-for-postgresql",
		"plan":     "standard-gen2",
	}
	if backups != nil {
		raw["backups"] = backups
	}
	return schema.TestResourceDataRaw(t, ResourceIBMDatabaseInstance().Schema, raw)
}

func TestParseBackupsRetention(t *testing.T) {
	cases := []struct {
		name  string
		input interface{}
		want  string
		ok    bool
	}{
		{name: "string value", input: "30d", want: "30d", ok: true},
		{name: "float64 value", input: float64(45), want: "45d", ok: true},
		{name: "int value", input: 60, want: "60d", ok: true},
		{name: "empty string", input: "", want: "", ok: false},
		{name: "nil", input: nil, want: "", ok: false},
		{name: "bool is ignored", input: true, want: "", ok: false},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got, ok := parseBackupsRetention(c.input)
			require.Equal(t, c.ok, ok)
			require.Equal(t, c.want, got)
		})
	}
}

func TestFlattenBackupsRetentionFromExtensions(t *testing.T) {
	cases := []struct {
		name       string
		extensions map[string]interface{}
		want       []interface{}
	}{
		{name: "nil extensions", extensions: nil, want: nil},
		{name: "empty extensions", extensions: map[string]interface{}{}, want: nil},
		{
			name:       "dataservices not a map",
			extensions: map[string]interface{}{"dataservices": "x"},
			want:       nil,
		},
		{
			name: "backups missing",
			extensions: map[string]interface{}{
				"dataservices": map[string]interface{}{"other": "x"},
			},
			want: nil,
		},
		{
			name: "backups nil",
			extensions: map[string]interface{}{
				"dataservices": map[string]interface{}{"backups": nil},
			},
			want: nil,
		},
		{
			name: "retention string",
			extensions: map[string]interface{}{
				"dataservices": map[string]interface{}{
					"backups": map[string]interface{}{"retention": "31d"},
				},
			},
			want: []interface{}{map[string]interface{}{"retention": "31d"}},
		},
		{
			name: "numeric retention",
			extensions: map[string]interface{}{
				"dataservices": map[string]interface{}{
					"backups": map[string]interface{}{"retention": float64(45)},
				},
			},
			want: []interface{}{map[string]interface{}{"retention": "45d"}},
		},
		{
			name: "unparseable retention yields nil",
			extensions: map[string]interface{}{
				"dataservices": map[string]interface{}{
					"backups": map[string]interface{}{"retention": true},
				},
			},
			want: nil,
		},
		{
			name: "other backups fields are ignored",
			extensions: map[string]interface{}{
				"dataservices": map[string]interface{}{
					"backups": map[string]interface{}{
						"retention": "30d",
						"preserve":  true,
					},
				},
			},
			want: []interface{}{map[string]interface{}{"retention": "30d"}},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			require.Equal(t, c.want, flattenBackupsRetentionFromExtensions(c.extensions))
		})
	}
}

func TestBuildBackupsRetentionConfig(t *testing.T) {
	t.Run("no backups block returns nil", func(t *testing.T) {
		d := retentionTestResourceData(t, nil)
		require.Nil(t, buildBackupsRetentionConfig(d, false))
	})

	t.Run("empty backups block returns nil", func(t *testing.T) {
		d := retentionTestResourceData(t, []interface{}{map[string]interface{}{}})
		require.Nil(t, buildBackupsRetentionConfig(d, false))
	})

	t.Run("retention set returns payload", func(t *testing.T) {
		d := retentionTestResourceData(t, []interface{}{map[string]interface{}{"retention": "45d"}})
		require.Equal(t, map[string]interface{}{"retention": "45d"}, buildBackupsRetentionConfig(d, false))
	})
}
