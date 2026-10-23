// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"context"
	"encoding/json"
	"testing"

	ctyjson "github.com/hashicorp/go-cty/cty/json"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// gen2BackupsParamsData returns resource data with a member group (so buildGen2Parameters
// needs no catalog lookup) and an optional backups block. A non-empty id simulates an
// existing instance.
func gen2BackupsParamsData(t *testing.T, id string, backups []interface{}) *schema.ResourceData {
	t.Helper()
	raw := map[string]interface{}{
		"name":     "test-gen2-db",
		"location": "us-south",
		"service":  "databases-for-postgresql",
		"plan":     "standard-gen2",
		"group": []interface{}{map[string]interface{}{
			"group_id": "member",
			"members":  []interface{}{map[string]interface{}{"allocation_count": 3}},
		}},
	}
	if backups != nil {
		raw["backups"] = backups
	}
	d := schema.TestResourceDataRaw(t, ResourceIBMDatabaseInstance().Schema, raw)
	d.SetId(id)
	return d
}

// TestGen2BuildGen2ParametersBackups verifies buildGen2Parameters sends backups on create
// only, as a sibling of the db-type key, and never on an existing instance.
func TestGen2BuildGen2ParametersBackups(t *testing.T) {
	g := newResourceIBMDatabaseGen2Backend().(*resourceIBMDatabaseGen2Backend)
	backups := []interface{}{map[string]interface{}{
		"preserve":  true,
		"retention": 45,
	}}

	t.Run("create sends backups next to the db-type key", func(t *testing.T) {
		d := gen2BackupsParamsData(t, "", backups)
		params, err := g.buildGen2Parameters(d, "databases-for-postgresql", nil, "")
		require.NoError(t, err)

		dataservices := params["dataservices"].(map[string]interface{})
		assert.Contains(t, dataservices, "postgresql", "db-type key must remain in dataservices")
		assert.Equal(t, map[string]interface{}{"preserve": true, "retention": "45d"}, dataservices["backups"])
	})

	t.Run("update does not send backups", func(t *testing.T) {
		d := gen2BackupsParamsData(t, "crn:v1:test", backups)
		params, err := g.buildGen2Parameters(d, "databases-for-postgresql", nil, "")
		require.NoError(t, err)

		dataservices := params["dataservices"].(map[string]interface{})
		assert.Contains(t, dataservices, "postgresql")
		assert.NotContains(t, dataservices, "backups")
	})

	t.Run("create without a backups block sends no backups", func(t *testing.T) {
		d := gen2BackupsParamsData(t, "", nil)
		params, err := g.buildGen2Parameters(d, "databases-for-postgresql", nil, "")
		require.NoError(t, err)

		assert.NotContains(t, params["dataservices"], "backups")
	})
}

// TestGen2BackupsUpdateGuard verifies applyBackupsWithDiagnostics is a no-op when backups
// have not changed (the guard returns before any client is used).
func TestGen2BackupsUpdateGuard(t *testing.T) {
	g := &resourceIBMDatabaseGen2Backend{}
	d := backupsTestResourceData(t, nil)

	diags := g.applyBackupsWithDiagnostics(d, nil, nil)
	assert.Empty(t, diags)
}

// TestBackupsParameters verifies the update envelope: backups is a direct child of dataservices.
func TestBackupsParameters(t *testing.T) {
	backups := map[string]interface{}{"retention": "45d"}
	assert.Equal(t, map[string]interface{}{
		"dataservices": map[string]interface{}{"backups": backups},
	}, backupsParameters(backups))
}

// backupsUpdateData builds resource data for an existing instance whose prior state holds
// priorBackups, with plannedBackups as the new config, so HasChange reflects a real update.
func backupsUpdateData(t *testing.T, priorBackups, plannedBackups []interface{}) *schema.ResourceData {
	t.Helper()
	res := ResourceIBMDatabaseInstance()
	base := map[string]interface{}{
		"name":     "test-gen2-db",
		"location": "us-south",
		"service":  "databases-for-postgresql",
		"plan":     "standard-gen2",
	}

	prior := schema.TestResourceDataRaw(t, res.Schema, base)
	prior.SetId("crn:v1:test")
	require.NoError(t, prior.Set("backups", priorBackups))
	state := prior.State()

	cfg := map[string]interface{}{"backups": plannedBackups}
	for k, v := range base {
		cfg[k] = v
	}
	diff, err := res.Diff(context.Background(), state, terraform.NewResourceConfigRaw(cfg), nil)
	require.NoError(t, err)
	d, err := schema.InternalMap(res.Schema).Data(state, diff)
	require.NoError(t, err)
	return d
}

// TestBuildBackupsConfigOnlyChanged verifies an update sends only the fields that changed,
// so targets managed outside Terraform are not overwritten by a retention change.
func TestBuildBackupsConfigOnlyChanged(t *testing.T) {
	prior := []interface{}{map[string]interface{}{
		"preserve":  false,
		"retention": 30,
		"automatic_backups": []interface{}{map[string]interface{}{
			"copy": []interface{}{map[string]interface{}{
				"enabled": true,
				"target":  []interface{}{map[string]interface{}{"region": "ca-mon"}},
			}},
		}},
	}}

	t.Run("retention change omits unchanged preserve and targets", func(t *testing.T) {
		d := backupsUpdateData(t, prior, []interface{}{map[string]interface{}{"retention": 45}})
		require.True(t, d.HasChange("backups"))
		assert.Equal(t, map[string]interface{}{"retention": "45d"}, buildBackupsConfig(d, true))
	})

	t.Run("no change sends nothing", func(t *testing.T) {
		d := backupsUpdateData(t, prior, []interface{}{map[string]interface{}{"retention": 30}})
		assert.Nil(t, buildBackupsConfig(d, true))
	})

	t.Run("target order does not matter", func(t *testing.T) {
		two := func(first, second string) []interface{} {
			return []interface{}{map[string]interface{}{
				"automatic_backups": []interface{}{map[string]interface{}{
					"copy": []interface{}{map[string]interface{}{
						"target": []interface{}{
							map[string]interface{}{"region": first},
							map[string]interface{}{"region": second},
						},
					}},
				}},
			}}
		}
		d := backupsUpdateData(t, two("ca-mon", "eu-de"), two("eu-de", "ca-mon"))
		assert.False(t, d.HasChange("backups"))
	})
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

	d := backupsTestResourceData(t, nil)
	require.NoError(t, d.Set("backups", flattenBackupsFromExtensions(extensions)))

	assert.Equal(t, true, d.Get("backups.0.preserve"))
	assert.Equal(t, 45, d.Get("backups.0.retention"))
	assert.Equal(t, true, d.Get("backups.0.automatic_backups.0.copy.0.enabled"))
	var targets []interface{}
	if set, ok := d.Get("backups.0.automatic_backups.0.copy.0.target").(*schema.Set); ok {
		targets = set.List()
	} else if list, ok := d.Get("backups.0.automatic_backups.0.copy.0.target").([]interface{}); ok {
		targets = list
	}
	require.Len(t, targets, 1)
	assert.Equal(t, "ca-mon", targets[0].(map[string]interface{})["region"])
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

	d := backupsTestResourceData(t, nil)
	require.NoError(t, d.Set("backups", flattenBackupsFromExtensions(extensions)))

	assert.Equal(t, false, d.Get("backups.0.preserve"))
	assert.Equal(t, 30, d.Get("backups.0.retention"))
	assert.Equal(t, false, d.Get("backups.0.automatic_backups.0.copy.0.enabled"))
	assert.Equal(t, 0, d.Get("backups.0.automatic_backups.0.copy.0.target.#"))
}

// TestGen2BackupsCopyDisabledDropsRetainedTargets verifies targets retained from state
// (Computed) are not sent when copies are disabled.
func TestGen2BackupsCopyDisabledDropsRetainedTargets(t *testing.T) {
	d := backupsTestResourceData(t, []interface{}{map[string]interface{}{
		"automatic_backups": []interface{}{map[string]interface{}{
			"copy": []interface{}{map[string]interface{}{
				"enabled": false,
				"target":  []interface{}{map[string]interface{}{"region": "ca-mon"}},
			}},
		}},
	}})

	got := buildBackupsConfig(d, false)
	require.NotNil(t, got)
	cp := got["automatic_backups"].(map[string]interface{})["copy"].(map[string]interface{})
	assert.Equal(t, false, cp["enabled"])
	assert.NotContains(t, cp, "target")
}

// TestValidateBackupsDiffThroughResourceDiff runs the real validateBackupsDiff CustomizeDiff
// via Resource.Diff, covering what the unit tests of the inner rules cannot: raw config
// intent, the plan/service lookups and the state-versus-plan comparison.
func TestValidateBackupsDiffThroughResourceDiff(t *testing.T) {
	res := ResourceIBMDatabaseInstance()

	plan := func(service, planName string, backups []interface{}, prior []interface{}) error {
		base := map[string]interface{}{
			"name":     "test-gen2-db",
			"location": "us-south",
			"service":  service,
			"plan":     planName,
		}
		state := &terraform.InstanceState{}
		if prior != nil {
			d := schema.TestResourceDataRaw(t, res.Schema, base)
			d.SetId("crn:v1:test")
			require.NoError(t, d.Set("backups", prior))
			state = d.State()
		}
		cfg := map[string]interface{}{}
		for k, v := range base {
			cfg[k] = v
		}
		if backups != nil {
			cfg["backups"] = backups
		}

		// validateBackupsDiff reads the user's intent from the raw config, which Terraform core
		// supplies via the state. Build it the same way from the resource's implied type.
		cfgJSON, err := json.Marshal(cfg)
		require.NoError(t, err)
		raw, err := ctyjson.Unmarshal(cfgJSON, res.CoreConfigSchema().ImpliedType())
		require.NoError(t, err)
		state.RawConfig = raw

		_, err = res.Diff(context.Background(), state, terraform.NewResourceConfigRaw(cfg), nil)
		return err
	}

	preserve := func(v bool) []interface{} {
		return []interface{}{map[string]interface{}{"preserve": v}}
	}

	t.Run("gen2 postgres with backups is valid", func(t *testing.T) {
		assert.NoError(t, plan("databases-for-postgresql", "standard-gen2", []interface{}{map[string]interface{}{"retention": 45}}, nil))
	})

	t.Run("classic plan with backups is rejected", func(t *testing.T) {
		err := plan("databases-for-postgresql", "standard", []interface{}{map[string]interface{}{"retention": 45}}, nil)
		require.Error(t, err)
		assert.Contains(t, err.Error(), "only supported for Gen2")
	})

	t.Run("preserve is rejected for mysql", func(t *testing.T) {
		err := plan("databases-for-mysql", "standard-gen2", preserve(true), nil)
		require.Error(t, err)
		assert.Contains(t, err.Error(), "cannot be enabled for databases-for-mysql")
	})

	t.Run("preserve cannot go from true to false", func(t *testing.T) {
		err := plan("databases-for-postgresql", "standard-gen2", preserve(false), preserve(true))
		require.Error(t, err)
		assert.Contains(t, err.Error(), "cannot be disabled")
	})

	t.Run("preserve can go from false to true", func(t *testing.T) {
		assert.NoError(t, plan("databases-for-postgresql", "standard-gen2", preserve(true), preserve(false)))
	})

	t.Run("targets with copy disabled are rejected", func(t *testing.T) {
		err := plan("databases-for-postgresql", "standard-gen2", []interface{}{map[string]interface{}{
			"automatic_backups": []interface{}{map[string]interface{}{
				"copy": []interface{}{map[string]interface{}{
					"enabled": false,
					"target":  []interface{}{map[string]interface{}{"region": "ca-mon"}},
				}},
			}},
		}}, nil)
		require.Error(t, err)
		assert.Contains(t, err.Error(), "cannot be set when")
	})
}
