// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func gen2BackupsTestData(t *testing.T, backups []interface{}) *schema.ResourceData {
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

// TestGen2BackupsInjectedIntoDataservices verifies the backups payload sits next to
// the db-type key inside dataservices, mirroring buildGen2Parameters.
func TestGen2BackupsInjectedIntoDataservices(t *testing.T) {
	d := gen2BackupsTestData(t, []interface{}{map[string]interface{}{
		"preserve":  true,
		"retention": 45,
	}})

	dataservices := map[string]interface{}{"postgresql": map[string]interface{}{"members": 3}}
	if backups := buildBackupsConfig(d); backups != nil {
		dataservices["backups"] = backups
	}

	backups, ok := dataservices["backups"].(map[string]interface{})
	require.True(t, ok, "dataservices must contain a backups map")
	assert.Equal(t, true, backups["preserve"])
	assert.Equal(t, "45d", backups["retention"])
	assert.Contains(t, dataservices, "postgresql", "db-type key must remain in dataservices")
}

// TestGen2BackupsNotSentWhenAbsent verifies no backups key is produced without a backups block.
func TestGen2BackupsNotSentWhenAbsent(t *testing.T) {
	d := gen2BackupsTestData(t, nil)

	dataservices := map[string]interface{}{"postgresql": map[string]interface{}{"members": 3}}
	if backups := buildBackupsConfig(d); backups != nil {
		dataservices["backups"] = backups
	}

	assert.NotContains(t, dataservices, "backups")
}

// TestGen2BackupsUpdateGuard verifies applyBackupsWithDiagnostics is a no-op when backups
// have not changed (the guard returns before any client is used).
func TestGen2BackupsUpdateGuard(t *testing.T) {
	g := &resourceIBMDatabaseGen2Backend{}
	d := gen2BackupsTestData(t, nil)

	diags := g.applyBackupsWithDiagnostics(context.Background(), d, nil, "test-id", nil)
	assert.Empty(t, diags)
}

// TestGen2BackupsFlattenAcceptedBySchema verifies that the flatten output is accepted by
// the resource schema and reads back with the expected values, which is what
// populateResourceData relies on.
func TestGen2BackupsFlattenAcceptedBySchema(t *testing.T) {
	extensions := map[string]interface{}{
		"dataservices": map[string]interface{}{
			"backups": map[string]interface{}{
				"preserve":  true,
				"retention": "45d",
				"automatic_backups": map[string]interface{}{
					"copy": map[string]interface{}{
						"enabled": true,
						"target": []interface{}{
							map[string]interface{}{"region": "ca-mon", "account_id": "a", "resource_group_id": "rg"},
						},
					},
				},
			},
		},
	}

	d := gen2BackupsTestData(t, nil)
	require.NoError(t, d.Set("backups", flattenBackupsFromExtensions(extensions)))

	assert.Equal(t, true, d.Get("backups.0.preserve"))
	assert.Equal(t, 45, d.Get("backups.0.retention"))
	assert.Equal(t, true, d.Get("backups.0.automatic_backups.0.copy.0.enabled"))
	assert.Equal(t, "ca-mon", d.Get("backups.0.automatic_backups.0.copy.0.target.0.region"))
	assert.Equal(t, 1, d.Get("backups.0.automatic_backups.0.copy.0.target.#"))
}

// TestGen2BackupsFlattenCopyDisabledAcceptedBySchema covers the copy-disabled shape (no targets).
func TestGen2BackupsFlattenCopyDisabledAcceptedBySchema(t *testing.T) {
	extensions := map[string]interface{}{
		"dataservices": map[string]interface{}{
			"backups": map[string]interface{}{
				"preserve":  false,
				"retention": "30d",
				"automatic_backups": map[string]interface{}{
					"copy": map[string]interface{}{"enabled": false},
				},
			},
		},
	}

	d := gen2BackupsTestData(t, nil)
	require.NoError(t, d.Set("backups", flattenBackupsFromExtensions(extensions)))

	assert.Equal(t, false, d.Get("backups.0.preserve"))
	assert.Equal(t, 30, d.Get("backups.0.retention"))
	assert.Equal(t, false, d.Get("backups.0.automatic_backups.0.copy.0.enabled"))
	assert.Equal(t, 0, d.Get("backups.0.automatic_backups.0.copy.0.target.#"))
}

// TestGen2BackupsCopyDisabledDropsRetainedTargets verifies targets retained from state
// (Computed) are not sent when copies are disabled.
func TestGen2BackupsCopyDisabledDropsRetainedTargets(t *testing.T) {
	d := gen2BackupsTestData(t, []interface{}{map[string]interface{}{
		"automatic_backups": []interface{}{map[string]interface{}{
			"copy": []interface{}{map[string]interface{}{
				"enabled": false,
				"target":  []interface{}{map[string]interface{}{"region": "ca-mon"}},
			}},
		}},
	}})

	got := buildBackupsConfig(d)
	require.NotNil(t, got)
	cp := got["automatic_backups"].(map[string]interface{})["copy"].(map[string]interface{})
	assert.Equal(t, false, cp["enabled"])
	assert.NotContains(t, cp, "target")
}
