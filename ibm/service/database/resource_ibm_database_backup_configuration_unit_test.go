// Copyright IBM Corp. 2017, 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"context"
	"fmt"
	"testing"

	rc "github.com/IBM/platform-services-go-sdk/resourcecontrollerv2"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/stretchr/testify/require"
)

// ---------------------------------------------------------------------------
// Harness: build real *schema.ResourceDiff / *schema.ResourceData values for backup_configuration
// ---------------------------------------------------------------------------

func baseBackupTestAttrs(service, plan string) map[string]interface{} {
	return map[string]interface{}{
		"name":     "test-db",
		"location": "ca-mon",
		"service":  service,
		"plan":     plan,
	}
}

func backupBlock(preserve bool, retention int, copyEnabled bool, regions ...string) []interface{} {
	rs := make([]interface{}, 0, len(regions))
	for _, r := range regions {
		rs = append(rs, r)
	}
	return []interface{}{map[string]interface{}{
		"preserve":       preserve,
		"retention":      retention,
		"copy_enabled":   copyEnabled,
		"target_regions": rs,
	}}
}

// backupInstanceState returns a prior state for an existing instance with the given backup_configuration.
func backupInstanceState(service, plan string, preserve bool, retention int, copyEnabled bool, regions ...string) *terraform.InstanceState {
	attrs := map[string]string{
		"id":                                  "crn:v1:test:instance",
		"name":                                "test-db",
		"location":                            "ca-mon",
		"service":                             service,
		"plan":                                plan,
		"backup_configuration.#":              "1",
		"backup_configuration.0.preserve":     fmt.Sprint(preserve),
		"backup_configuration.0.retention":    fmt.Sprint(retention),
		"backup_configuration.0.copy_enabled": fmt.Sprint(copyEnabled),
		"backup_configuration.0.target_regions.#": fmt.Sprint(len(regions)),
	}
	for _, r := range regions {
		attrs[fmt.Sprintf("backup_configuration.0.target_regions.%d", schema.HashString(r))] = r
	}
	return &terraform.InstanceState{ID: "crn:v1:test:instance", Attributes: attrs}
}

// diffWith plans raw against state using only the given CustomizeDiff, over the real schema.
// A nil customize returns the raw planned diff.
func diffWith(raw map[string]interface{}, state *terraform.InstanceState, customize schema.CustomizeDiffFunc) (*terraform.InstanceDiff, error) {
	r := &schema.Resource{
		Schema:        ResourceIBMDatabaseInstance().Schema,
		CustomizeDiff: customize,
	}
	return r.Diff(context.Background(), state, terraform.NewResourceConfigRaw(raw), nil)
}

// plannedData returns the *schema.ResourceData an apply would see: prior state merged with the planned diff.
func plannedData(t *testing.T, raw map[string]interface{}, state *terraform.InstanceState) *schema.ResourceData {
	t.Helper()
	diff, err := diffWith(raw, state, nil)
	require.NoError(t, err)
	d, err := schema.InternalMap(ResourceIBMDatabaseInstance().Schema).Data(state, diff)
	require.NoError(t, err)
	return d
}

func withBackup(base map[string]interface{}, block []interface{}) map[string]interface{} {
	out := map[string]interface{}{}
	for k, v := range base {
		out[k] = v
	}
	out["backup_configuration"] = block
	return out
}

// ---------------------------------------------------------------------------
// ValidateBackupConfigurationDiff (the real CustomizeDiff, against a real ResourceDiff)
// ---------------------------------------------------------------------------

func TestValidateBackupConfigurationDiff(t *testing.T) {
	pg := baseBackupTestAttrs("databases-for-postgresql", "standard-gen2")

	t.Run("no backup_configuration block is accepted", func(t *testing.T) {
		_, err := diffWith(pg, nil, validateBackupConfigurationDiff)
		require.NoError(t, err)
	})

	t.Run("retention out of range is rejected", func(t *testing.T) {
		for _, r := range []int{1, 29, 61, 90} {
			_, err := diffWith(withBackup(pg, backupBlock(false, r, true)), nil, validateBackupConfigurationDiff)
			requireErrContains(t, err, "retention in backup_configuration must be between 30 and 60 days")
		}
	})

	t.Run("retention within range is accepted", func(t *testing.T) {
		for _, r := range []int{30, 45, 60} {
			_, err := diffWith(withBackup(pg, backupBlock(false, r, true)), nil, validateBackupConfigurationDiff)
			require.NoError(t, err, "retention %d", r)
		}
	})

	t.Run("preserve is rejected for MySQL", func(t *testing.T) {
		my := baseBackupTestAttrs("databases-for-mysql", "standard-gen2")
		_, err := diffWith(withBackup(my, backupBlock(true, 30, true)), nil, validateBackupConfigurationDiff)
		requireErrContains(t, err, "preserve backups is not supported for Databases for MySQL")
	})

	t.Run("preserve=false is accepted for MySQL", func(t *testing.T) {
		my := baseBackupTestAttrs("databases-for-mysql", "standard-gen2")
		_, err := diffWith(withBackup(my, backupBlock(false, 30, true)), nil, validateBackupConfigurationDiff)
		require.NoError(t, err)
	})

	t.Run("preserve is accepted for PostgreSQL and Valkey", func(t *testing.T) {
		for _, svc := range []string{"databases-for-postgresql", "databases-for-valkey"} {
			_, err := diffWith(withBackup(baseBackupTestAttrs(svc, "standard-gen2"), backupBlock(true, 30, true)), nil, validateBackupConfigurationDiff)
			require.NoError(t, err, svc)
		}
	})

	t.Run("preserve true -> false is rejected on update", func(t *testing.T) {
		state := backupInstanceState("databases-for-postgresql", "standard-gen2", true, 30, true)
		_, err := diffWith(withBackup(pg, backupBlock(false, 30, true)), state, validateBackupConfigurationDiff)
		requireErrContains(t, err, "preserve backups cannot be disabled once enabled")
	})

	t.Run("preserve false -> true is accepted on update", func(t *testing.T) {
		state := backupInstanceState("databases-for-postgresql", "standard-gen2", false, 30, true)
		_, err := diffWith(withBackup(pg, backupBlock(true, 30, true)), state, validateBackupConfigurationDiff)
		require.NoError(t, err)
	})

	t.Run("preserve true -> true with another field changed is accepted", func(t *testing.T) {
		state := backupInstanceState("databases-for-postgresql", "standard-gen2", true, 30, true)
		_, err := diffWith(withBackup(pg, backupBlock(true, 50, true)), state, validateBackupConfigurationDiff)
		require.NoError(t, err)
	})

	t.Run("omitting the block on an instance with preserve=true is not a disable", func(t *testing.T) {
		state := backupInstanceState("databases-for-postgresql", "standard-gen2", true, 30, true)
		_, err := diffWith(pg, state, validateBackupConfigurationDiff)
		require.NoError(t, err)
	})

	t.Run("Classic backend accepts anything here (rejection is in the unsupported-attrs validator)", func(t *testing.T) {
		classic := baseBackupTestAttrs("databases-for-postgresql", "standard")
		_, err := diffWith(withBackup(classic, backupBlock(true, 99, true)), nil, validateBackupConfigurationDiff)
		require.NoError(t, err)
	})
}

// TestBackupConfigurationClassicRejected proves Classic plans reject backup_configuration at plan time.
func TestBackupConfigurationClassicRejected(t *testing.T) {
	classic := baseBackupTestAttrs("databases-for-postgresql", "standard")

	t.Run("rejected when set on a Classic plan", func(t *testing.T) {
		_, err := diffWith(withBackup(classic, backupBlock(false, 30, true)), nil, validateUnsupportedAttrsDiff)
		requireErrContains(t, err, "not supported for Classic databases")
		requireErrContains(t, err, `"backup_configuration"`)
	})

	t.Run("not rejected when absent on a Classic plan", func(t *testing.T) {
		_, err := diffWith(classic, nil, validateUnsupportedAttrsDiff)
		require.NoError(t, err)
	})

	t.Run("not treated as unsupported on a Gen2 plan", func(t *testing.T) {
		gen2 := baseBackupTestAttrs("databases-for-postgresql", "standard-gen2")
		_, err := diffWith(withBackup(gen2, backupBlock(false, 30, true)), nil, validateUnsupportedAttrsDiff)
		if err != nil {
			requireErrNotContains(t, err, "backup_configuration")
		}
	})
}

// TestBackupConfigurationValidatorRegistered runs the real resource's full CustomizeDiff chain, so it
// fails if validateBackupConfigurationDiff is ever dropped from ResourceIBMDatabaseInstance().
func TestBackupConfigurationValidatorRegistered(t *testing.T) {
	gen2 := baseBackupTestAttrs("databases-for-postgresql", "standard-gen2")
	gen2["service_endpoints"] = "private"

	r := ResourceIBMDatabaseInstance()
	_, err := r.Diff(context.Background(), nil,
		terraform.NewResourceConfigRaw(withBackup(gen2, backupBlock(false, 90, true))), nil)
	requireErrContains(t, err, "retention in backup_configuration must be between 30 and 60 days")
}

// ---------------------------------------------------------------------------
// Payload shape: create (buildGen2Parameters) and update (backupsParameters)
// ---------------------------------------------------------------------------

func gen2ParamsData(t *testing.T, id string, block []interface{}) *schema.ResourceData {
	t.Helper()
	raw := baseBackupTestAttrs("databases-for-postgresql", "standard-gen2")
	raw["group"] = []interface{}{map[string]interface{}{
		"group_id": "member",
		"members":  []interface{}{map[string]interface{}{"allocation_count": 2}},
	}}
	if block != nil {
		raw["backup_configuration"] = block
	}
	d := schema.TestResourceDataRaw(t, ResourceIBMDatabaseInstance().Schema, raw)
	if id != "" {
		d.SetId(id)
	}
	return d
}

func TestBuildGen2ParametersBackupPayload(t *testing.T) {
	g := &resourceIBMDatabaseGen2Backend{}

	t.Run("create places backups under dataservices, as a sibling of the db-type key", func(t *testing.T) {
		d := gen2ParamsData(t, "", backupBlock(true, 45, true, "us-south", "ca-tor"))

		params, err := g.buildGen2Parameters(d, "databases-for-postgresql", nil, "crn:catalog")
		require.NoError(t, err)

		ds, ok := params["dataservices"].(map[string]interface{})
		require.True(t, ok, "dataservices must be a map")
		require.Contains(t, ds, "postgresql", "db-type key still present")
		require.Contains(t, ds, "backups", "backups must be a direct child of dataservices")

		// ...and must NOT be nested under the db-type key.
		require.NotContains(t, ds["postgresql"], "backups")

		require.Equal(t, map[string]interface{}{
			"preserve":  true,
			"retention": "45d",
			"automatic_backups": map[string]interface{}{
				"copy": map[string]interface{}{
					"enabled": true,
					"target": []map[string]interface{}{
						{"region": "ca-tor"}, {"region": "us-south"}, // sorted: deterministic
					},
				},
			},
		}, ds["backups"])
	})

	t.Run("create without backup_configuration sends no backups key", func(t *testing.T) {
		d := gen2ParamsData(t, "", nil)
		params, err := g.buildGen2Parameters(d, "databases-for-postgresql", nil, "crn:catalog")
		require.NoError(t, err)
		require.NotContains(t, params["dataservices"], "backups")
	})

	t.Run("update never re-sends backups on group/config PATCHes", func(t *testing.T) {
		// An existing instance (ID set) that has backup_configuration in its planned data.
		d := gen2ParamsData(t, "crn:v1:test:instance", backupBlock(true, 45, true, "ca-tor"))
		params, err := g.buildGen2Parameters(d, "databases-for-postgresql", nil, "crn:catalog")
		require.NoError(t, err)
		require.NotContains(t, params["dataservices"], "backups",
			"applyGroupScaling/applyConfigurationUpdate reuse this builder; backups must only be sent by the dedicated call")
	})
}

func TestBackupsParametersEnvelope(t *testing.T) {
	// ADR: {"dataservices": {"backups": {...}}} -- backups is a child of dataservices, not of the db type.
	got := backupsParameters(map[string]interface{}{"retention": "50d"})
	require.Equal(t, map[string]interface{}{
		"dataservices": map[string]interface{}{
			"backups": map[string]interface{}{"retention": "50d"},
		},
	}, got)
}

// ---------------------------------------------------------------------------
// Update payload: only changed fields are sent
// ---------------------------------------------------------------------------

func TestBuildBackupsPayloadUpdate(t *testing.T) {
	pg := baseBackupTestAttrs("databases-for-postgresql", "standard-gen2")
	state := backupInstanceState("databases-for-postgresql", "standard-gen2", false, 30, true, "ca-tor")

	t.Run("only retention changed -> only retention is sent", func(t *testing.T) {
		d := plannedData(t, withBackup(pg, backupBlock(false, 50, true, "ca-tor")), state)
		require.True(t, d.HasChange("backup_configuration"))
		require.Equal(t, map[string]interface{}{"retention": "50d"}, buildBackupsPayload(d, false))
	})

	t.Run("only preserve changed -> only preserve is sent", func(t *testing.T) {
		d := plannedData(t, withBackup(pg, backupBlock(true, 30, true, "ca-tor")), state)
		require.Equal(t, map[string]interface{}{"preserve": true}, buildBackupsPayload(d, false))
	})

	t.Run("copy disabled -> copy.enabled=false and no targets", func(t *testing.T) {
		d := plannedData(t, withBackup(pg, backupBlock(false, 30, false, "ca-tor")), state)
		require.Equal(t, map[string]interface{}{
			"automatic_backups": map[string]interface{}{
				"copy": map[string]interface{}{"enabled": false},
			},
		}, buildBackupsPayload(d, false))
	})

	t.Run("target regions changed -> copy.enabled and targets are sent", func(t *testing.T) {
		d := plannedData(t, withBackup(pg, backupBlock(false, 30, true, "ca-tor", "us-south")), state)
		require.Equal(t, map[string]interface{}{
			"automatic_backups": map[string]interface{}{
				"copy": map[string]interface{}{
					"enabled": true,
					"target":  []map[string]interface{}{{"region": "ca-tor"}, {"region": "us-south"}},
				},
			},
		}, buildBackupsPayload(d, false))
	})

	t.Run("reordering target regions is not a change", func(t *testing.T) {
		two := backupInstanceState("databases-for-postgresql", "standard-gen2", false, 30, true, "ca-tor", "us-south")
		d := plannedData(t, withBackup(pg, backupBlock(false, 30, true, "us-south", "ca-tor")), two)
		require.False(t, d.HasChange("backup_configuration"), "set ordering must not produce a diff")
		require.Nil(t, buildBackupsPayload(d, false))
	})

	t.Run("no change -> nothing to send", func(t *testing.T) {
		d := plannedData(t, withBackup(pg, backupBlock(false, 30, true, "ca-tor")), state)
		require.False(t, d.HasChange("backup_configuration"))
		require.Nil(t, buildBackupsPayload(d, false))
	})

	t.Run("create sends every configured field", func(t *testing.T) {
		d := plannedData(t, withBackup(pg, backupBlock(true, 45, true, "ca-tor")), nil)
		got := buildBackupsPayload(d, true)
		require.Equal(t, true, got["preserve"])
		require.Equal(t, "45d", got["retention"])
		require.Contains(t, got, "automatic_backups")
	})
}

// ---------------------------------------------------------------------------
// Data source: schema parity and populate path
// ---------------------------------------------------------------------------

func nestedSchemaKeys(t *testing.T, s map[string]*schema.Schema, key string) map[string]*schema.Schema {
	t.Helper()
	res, ok := s[key].Elem.(*schema.Resource)
	require.True(t, ok, "%s must be a nested block", key)
	return res.Schema
}

func TestDataSourceBackupConfigurationSchemaParity(t *testing.T) {
	resFields := nestedSchemaKeys(t, ResourceIBMDatabaseInstance().Schema, "backup_configuration")
	dsFields := nestedSchemaKeys(t, DataSourceIBMDatabaseInstance().Schema, "backup_configuration")

	require.True(t, DataSourceIBMDatabaseInstance().Schema["backup_configuration"].Computed)
	require.False(t, DataSourceIBMDatabaseInstance().Schema["backup_configuration"].Optional)

	for name, rf := range resFields {
		df, ok := dsFields[name]
		require.True(t, ok, "data source is missing backup_configuration.%s", name)
		require.Equal(t, rf.Type, df.Type, "type mismatch for backup_configuration.%s", name)
		require.True(t, df.Computed, "data source backup_configuration.%s must be Computed", name)
		require.False(t, df.Optional, "data source backup_configuration.%s must not be Optional", name)
	}
	for name := range dsFields {
		_, ok := resFields[name]
		require.True(t, ok, "data source has unexpected backup_configuration.%s", name)
	}
}

func TestDataSourcePopulatesBackupConfiguration(t *testing.T) {
	d := schema.TestResourceDataRaw(t, DataSourceIBMDatabaseInstance().Schema, map[string]interface{}{"name": "test-db"})

	instance := &rc.ResourceInstance{Extensions: map[string]interface{}{
		"dataservices": map[string]interface{}{
			"backups": map[string]interface{}{
				"preserve":  true,
				"retention": "45d",
				"automatic_backups": map[string]interface{}{
					"copy": map[string]interface{}{
						"enabled": "true",
						"target": []interface{}{
							map[string]interface{}{"region": "us-south"},
							map[string]interface{}{"region": "ca-tor"},
						},
					},
				},
			},
		},
	}}

	require.NoError(t, setGen2BackupConfiguration(d, instance))

	require.Equal(t, 1, d.Get("backup_configuration.#"))
	require.Equal(t, true, d.Get("backup_configuration.0.preserve"))
	require.Equal(t, 45, d.Get("backup_configuration.0.retention"))
	require.Equal(t, true, d.Get("backup_configuration.0.copy_enabled"))
	require.Equal(t, []string{"ca-tor", "us-south"}, targetRegionsOf(d))
}

func TestDataSourceClearsBackupConfigurationWhenAbsent(t *testing.T) {
	d := schema.TestResourceDataRaw(t, DataSourceIBMDatabaseInstance().Schema, map[string]interface{}{"name": "test-db"})
	require.NoError(t, setGen2BackupConfiguration(d, &rc.ResourceInstance{Extensions: map[string]interface{}{
		"dataservices": map[string]interface{}{"backups": map[string]interface{}{"preserve": true}},
	}}))
	require.Equal(t, 1, d.Get("backup_configuration.#"))

	require.NoError(t, setGen2BackupConfiguration(d, &rc.ResourceInstance{Extensions: map[string]interface{}{
		"dataservices": map[string]interface{}{},
	}}))
	require.Equal(t, 0, d.Get("backup_configuration.#"))
}
