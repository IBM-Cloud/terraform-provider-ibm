// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"testing"

	"github.com/hashicorp/go-cty/cty"
	"github.com/stretchr/testify/require"
)

func boolPtr(b bool) *bool { return &b }

func TestValidateBackupsRules(t *testing.T) {
	configured := func(preserve, copyEnabled *bool, targets int) backupsConfigIntent {
		return backupsConfigIntent{configured: true, preserve: preserve, copyEnabled: copyEnabled, targetCount: targets}
	}

	cases := []struct {
		name    string
		in      backupsRuleInput
		wantErr string
	}{
		{
			name: "no backups block is always valid, even on Classic",
			in:   backupsRuleInput{gen2: false, service: "databases-for-postgresql"},
		},
		{
			name:    "classic plan with backups block is rejected",
			in:      backupsRuleInput{gen2: false, service: "databases-for-postgresql", intent: configured(nil, nil, 0)},
			wantErr: "only supported for Gen2 plans",
		},
		{
			name: "gen2 with empty backups block is valid",
			in:   backupsRuleInput{gen2: true, service: "databases-for-postgresql", intent: configured(nil, nil, 0)},
		},
		{
			name: "preserve true on postgres is valid",
			in:   backupsRuleInput{gen2: true, service: "databases-for-postgresql", intent: configured(boolPtr(true), nil, 0)},
		},
		{
			name:    "preserve true on mysql is rejected",
			in:      backupsRuleInput{gen2: true, service: "databases-for-mysql", intent: configured(boolPtr(true), nil, 0)},
			wantErr: "cannot be enabled for databases-for-mysql",
		},
		{
			name: "preserve false on mysql is valid",
			in:   backupsRuleInput{gen2: true, service: "databases-for-mysql", intent: configured(boolPtr(false), nil, 0)},
		},
		{
			name: "preserve omitted on mysql is valid",
			in:   backupsRuleInput{gen2: true, service: "databases-for-mysql", intent: configured(nil, nil, 0)},
		},
		{
			name: "preserve false to true on update is valid",
			in: backupsRuleInput{gen2: true, service: "databases-for-postgresql", intent: configured(boolPtr(true), nil, 0),
				isUpdate: true, preserveChanged: true, preserveOld: false, preserveNew: true},
		},
		{
			name: "preserve true to false on update is rejected",
			in: backupsRuleInput{gen2: true, service: "databases-for-postgresql", intent: configured(boolPtr(false), nil, 0),
				isUpdate: true, preserveChanged: true, preserveOld: true, preserveNew: false},
			wantErr: "cannot be disabled once it has been enabled",
		},
		{
			name: "preserve unchanged true on update is valid",
			in: backupsRuleInput{gen2: true, service: "databases-for-postgresql", intent: configured(boolPtr(true), nil, 0),
				isUpdate: true, preserveChanged: false, preserveOld: true, preserveNew: true},
		},
		{
			name: "copy enabled with target is valid",
			in:   backupsRuleInput{gen2: true, service: "databases-for-postgresql", intent: configured(nil, boolPtr(true), 1)},
		},
		{
			name: "copy disabled without target is valid",
			in:   backupsRuleInput{gen2: true, service: "databases-for-postgresql", intent: configured(nil, boolPtr(false), 0)},
		},
		{
			name:    "copy disabled with target is rejected",
			in:      backupsRuleInput{gen2: true, service: "databases-for-postgresql", intent: configured(nil, boolPtr(false), 1)},
			wantErr: "cannot be set when",
		},
		{
			name: "copy enabled omitted with target is valid",
			in:   backupsRuleInput{gen2: true, service: "databases-for-postgresql", intent: configured(nil, nil, 2)},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			err := validateBackupsRules(c.in)
			if c.wantErr == "" {
				require.NoError(t, err)
				return
			}
			require.Error(t, err)
			require.Contains(t, err.Error(), c.wantErr)
		})
	}
}

func TestBackupsIntentFromRawConfig(t *testing.T) {
	copyType := cty.Object(map[string]cty.Type{
		"enabled": cty.Bool,
		"target":  cty.List(cty.Object(map[string]cty.Type{"region": cty.String})),
	})
	autoType := cty.Object(map[string]cty.Type{"copy": cty.List(copyType)})
	backupsType := cty.Object(map[string]cty.Type{
		"preserve":          cty.Bool,
		"retention":         cty.Number,
		"automatic_backups": cty.List(autoType),
	})
	t.Run("no backups attribute in config", func(t *testing.T) {
		got := backupsIntentFromRawConfig(cty.ObjectVal(map[string]cty.Value{"name": cty.StringVal("x")}))
		require.False(t, got.configured)
	})

	t.Run("backups null", func(t *testing.T) {
		got := backupsIntentFromRawConfig(cty.ObjectVal(map[string]cty.Value{"backups": cty.NullVal(cty.List(backupsType))}))
		require.False(t, got.configured)
	})

	t.Run("empty backups list", func(t *testing.T) {
		got := backupsIntentFromRawConfig(cty.ObjectVal(map[string]cty.Value{"backups": cty.ListValEmpty(backupsType)}))
		require.False(t, got.configured)
	})

	t.Run("backups block with only retention", func(t *testing.T) {
		block := cty.ObjectVal(map[string]cty.Value{
			"preserve":          cty.NullVal(cty.Bool),
			"retention":         cty.NumberIntVal(45),
			"automatic_backups": cty.NullVal(cty.List(autoType)),
		})
		raw := cty.ObjectVal(map[string]cty.Value{"backups": cty.ListVal([]cty.Value{block})})
		got := backupsIntentFromRawConfig(raw)
		require.True(t, got.configured)
		require.Nil(t, got.preserve)
		require.Nil(t, got.copyEnabled)
		require.Equal(t, 0, got.targetCount)
	})

	t.Run("full block", func(t *testing.T) {
		target := cty.ObjectVal(map[string]cty.Value{"region": cty.StringVal("ca-mon")})
		cp := cty.ObjectVal(map[string]cty.Value{
			"enabled": cty.False,
			"target":  cty.ListVal([]cty.Value{target, target}),
		})
		auto := cty.ObjectVal(map[string]cty.Value{"copy": cty.ListVal([]cty.Value{cp})})
		block := cty.ObjectVal(map[string]cty.Value{
			"preserve":          cty.True,
			"retention":         cty.NumberIntVal(60),
			"automatic_backups": cty.ListVal([]cty.Value{auto}),
		})
		raw := cty.ObjectVal(map[string]cty.Value{"backups": cty.ListVal([]cty.Value{block})})
		got := backupsIntentFromRawConfig(raw)
		require.True(t, got.configured)
		require.NotNil(t, got.preserve)
		require.True(t, *got.preserve)
		require.NotNil(t, got.copyEnabled)
		require.False(t, *got.copyEnabled)
		require.Equal(t, 2, got.targetCount)
	})

	t.Run("unknown preserve is treated as not set", func(t *testing.T) {
		block := cty.ObjectVal(map[string]cty.Value{
			"preserve":          cty.UnknownVal(cty.Bool),
			"retention":         cty.NullVal(cty.Number),
			"automatic_backups": cty.NullVal(cty.List(autoType)),
		})
		raw := cty.ObjectVal(map[string]cty.Value{"backups": cty.ListVal([]cty.Value{block})})
		got := backupsIntentFromRawConfig(raw)
		require.True(t, got.configured)
		require.Nil(t, got.preserve)
	})
}
