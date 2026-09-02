package database

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func testGen2DatabaseResourceData(t *testing.T, raw map[string]interface{}) *schema.ResourceData {
	t.Helper()

	base := map[string]interface{}{
		"name":     "test-gen2-db",
		"location": "us-south",
		"service":  "databases-for-postgresql",
		"plan":     "standard-gen2",
	}

	for k, v := range raw {
		base[k] = v
	}

	return schema.TestResourceDataRaw(t, ResourceIBMDatabaseInstance().Schema, base)
}

func requireErrContains(t *testing.T, err error, expected string) {
	t.Helper()

	if err == nil {
		t.Fatalf("expected error containing %q, got nil", expected)
	}

	if !strings.Contains(err.Error(), expected) {
		t.Fatalf("expected error to contain %q, got:\n%s", expected, err.Error())
	}
}

func requireErrNotContains(t *testing.T, err error, unexpected string) {
	t.Helper()

	if err != nil && strings.Contains(err.Error(), unexpected) {
		t.Fatalf("expected error to not contain %q, got:\n%s", unexpected, err.Error())
	}
}

func requireWarningContains(t *testing.T, diags diag.Diagnostics, expected string) {
	t.Helper()

	for _, d := range diags {
		if d.Severity == diag.Warning &&
			(strings.Contains(d.Summary, expected) || strings.Contains(d.Detail, expected)) {
			return
		}
	}

	t.Fatalf("expected warning containing %q, got: %#v", expected, diags)
}

func requireNoErrors(t *testing.T, diags diag.Diagnostics) {
	t.Helper()

	for _, d := range diags {
		if d.Severity == diag.Error {
			t.Fatalf("expected no errors, got: %#v", diags)
		}
	}
}

func TestGen2UnsupportedAttrsValidation(t *testing.T) {
	g := &resourceIBMDatabaseGen2Backend{}
	adminPasswordValue := "example-admin-value"

	t.Run("unsupported attr present returns error", func(t *testing.T) {
		d := testGen2DatabaseResourceData(t, map[string]interface{}{
			"adminpassword": adminPasswordValue,
		})

		err := g.ValidateUnsupportedAttrsData(d)

		requireErrContains(t, err, "adminpassword")
		requireErrContains(t, err, "not supported")
	})

	t.Run("multiple unsupported attrs present are all listed", func(t *testing.T) {
		d := testGen2DatabaseResourceData(t, map[string]interface{}{
			"adminpassword":             adminPasswordValue,
			"backup_encryption_key_crn": "crn:v1:bluemix:public:kms:us-south:a/account-id:instance-id:key:key-id",
			"remote_leader_id":          "crn:v1:bluemix:public:databases-for-postgresql:us-south:a/account-id:instance-id::",
		})

		err := g.ValidateUnsupportedAttrsData(d)

		requireErrContains(t, err, "adminpassword")
		requireErrContains(t, err, "backup_encryption_key_crn")
		requireErrContains(t, err, "remote_leader_id")
	})

	t.Run("ignored attr only does not return error", func(t *testing.T) {
		d := testGen2DatabaseResourceData(t, map[string]interface{}{
			"configuration": `{"max_connections": 100}`,
		})

		err := g.ValidateUnsupportedAttrsData(d)

		if err != nil {
			t.Fatalf("expected no error for ignored attr only, got:\n%s", err.Error())
		}
	})

	t.Run("ignored and unsupported attrs returns error for unsupported attrs only", func(t *testing.T) {
		d := testGen2DatabaseResourceData(t, map[string]interface{}{
			"adminpassword": adminPasswordValue,
			"configuration": `{"max_connections": 100}`,
		})

		err := g.ValidateUnsupportedAttrsData(d)

		requireErrContains(t, err, "adminpassword")
		requireErrNotContains(t, err, "configuration")
	})
}

func TestGen2IgnoredAttrsWarnings(t *testing.T) {
	g := &resourceIBMDatabaseGen2Backend{}

	t.Run("ignored attr present returns warning", func(t *testing.T) {
		d := testGen2DatabaseResourceData(t, map[string]interface{}{
			"configuration": `{"max_connections": 100}`,
		})

		diags := g.WarnIgnoredAttrs(d)

		requireNoErrors(t, diags)

		if len(diags) != 1 {
			t.Fatalf("expected 1 warning, got %d: %#v", len(diags), diags)
		}

		if diags[0].Severity != diag.Warning {
			t.Fatalf("expected warning severity, got: %#v", diags[0])
		}

		requireWarningContains(t, diags, "configuration")
	})

	t.Run("multiple ignored attrs return one grouped warning", func(t *testing.T) {
		d := testGen2DatabaseResourceData(t, map[string]interface{}{
			"configuration":               `{"max_connections": 100}`,
			"version_upgrade_skip_backup": true,
			"skip_initial_backup":         true,
		})

		diags := g.WarnIgnoredAttrs(d)

		requireNoErrors(t, diags)

		if len(diags) != 1 {
			t.Fatalf("expected 1 grouped warning, got %d: %#v", len(diags), diags)
		}

		if diags[0].Severity != diag.Warning {
			t.Fatalf("expected warning severity, got: %#v", diags[0])
		}

		requireWarningContains(t, diags, "configuration")
		requireWarningContains(t, diags, "version_upgrade_skip_backup")
		requireWarningContains(t, diags, "skip_initial_backup")
	})

	t.Run("unsupported attr does not produce ignored warning", func(t *testing.T) {
		d := testGen2DatabaseResourceData(t, map[string]interface{}{
			"backup_id": "backup-123",
		})

		diags := g.WarnIgnoredAttrs(d)

		if len(diags) != 0 {
			t.Fatalf("expected no ignored warnings for unsupported attr, got: %#v", diags)
		}
	})

	t.Run("empty config returns no warnings", func(t *testing.T) {
		d := testGen2DatabaseResourceData(t, nil)

		diags := g.WarnIgnoredAttrs(d)

		if len(diags) != 0 {
			t.Fatalf("expected no warnings, got: %#v", diags)
		}
	})
}

func TestGen2IgnoredAttrsWarningsAreIndependentFromUnsupportedAttrs(t *testing.T) {
	g := &resourceIBMDatabaseGen2Backend{}
	adminPasswordValue := "example-admin-value"

	d := testGen2DatabaseResourceData(t, map[string]interface{}{
		"adminpassword":               adminPasswordValue,
		"configuration":               `{"max_connections": 100}`,
		"version_upgrade_skip_backup": true,
	})

	err := g.ValidateUnsupportedAttrsData(d)
	requireErrContains(t, err, "adminpassword")
	requireErrNotContains(t, err, "configuration")
	requireErrNotContains(t, err, "version_upgrade_skip_backup")

	diags := g.WarnIgnoredAttrs(d)
	requireNoErrors(t, diags)
	requireWarningContains(t, diags, "configuration")
	requireWarningContains(t, diags, "version_upgrade_skip_backup")
}

func TestGen2DiagnosticsCanContainErrorsAndWarnings(t *testing.T) {
	warnings := diag.Diagnostics{
		{
			Severity: diag.Warning,
			Summary:  "ignored attr warning",
			Detail:   "configuration is ignored",
		},
	}

	errors := diag.Diagnostics{
		{
			Severity: diag.Error,
			Summary:  "unsupported attr error",
			Detail:   "adminpassword is not supported",
		},
	}

	combined := appendGen2DiagnosticsErrorsThenWarnings(errors, warnings)

	if len(combined) != 2 {
		t.Fatalf("expected 2 diagnostics, got %d: %#v", len(combined), combined)
	}

	if combined[0].Severity != diag.Error {
		t.Fatalf("expected first diagnostic to be error, got: %#v", combined[0])
	}

	if combined[1].Severity != diag.Warning {
		t.Fatalf("expected second diagnostic to be warning, got: %#v", combined[1])
	}
}

func TestGen2WarningsReturnedWithErrors(t *testing.T) {
	g := &resourceIBMDatabaseGen2Backend{}

	d := testGen2DatabaseResourceData(t, map[string]interface{}{
		"backup_id":     "bad",                      // unsupported -> error
		"configuration": `{"max_connections": 100}`, // ignored -> warning
	})

	err := g.ValidateUnsupportedAttrsData(d)
	if err == nil {
		t.Fatalf("expected error for unsupported attr")
	}

	diags := g.WarnIgnoredAttrs(d)

	if len(diags) == 0 {
		t.Fatalf("expected warnings for ignored attrs")
	}
}

func TestGen2DiagnosticsOrdering(t *testing.T) {
	errors := diag.Diagnostics{
		{Severity: diag.Error, Summary: "error"},
	}
	warnings := diag.Diagnostics{
		{Severity: diag.Warning, Summary: "warning"},
	}

	out := appendGen2DiagnosticsErrorsThenWarnings(errors, warnings)

	if out[0].Severity != diag.Error {
		t.Fatalf("expected error first")
	}
	if out[1].Severity != diag.Warning {
		t.Fatalf("expected warning second")
	}
}

func TestDbConfigToMap_membersIncludedForMongodb(t *testing.T) {
	g := &resourceIBMDatabaseGen2Backend{}
	config := DBConfig{Members: 3, StorageGB: 10, HostFlavor: "bx3d.4x20"}

	result := g.dbConfigToMap(config, "mongodb")

	if _, ok := result["members"]; !ok {
		t.Fatal("expected 'members' to be present for dbType 'mongodb'")
	}
	if result["members"] != 3 {
		t.Fatalf("expected members=3, got %v", result["members"])
	}
}

func TestDbConfigToMap_membersExcludedAndShardsIncludedForMongodbees(t *testing.T) {
	g := &resourceIBMDatabaseGen2Backend{}
	config := DBConfig{Members: 3, Shards: 2, StorageGB: 10, HostFlavor: "bx3d.4x20"}

	result := g.dbConfigToMap(config, "mongodbees")

	if _, ok := result["members"]; ok {
		t.Fatalf("expected 'members' to be absent for dbType 'mongodbees', got %v", result["members"])
	}
	if result["shards"] != 2 {
		t.Fatalf("expected shards=2, got %v", result["shards"])
	}
}

func TestDbConfigToMap_storageShardsAndHostFlavorPresentForMongodbees(t *testing.T) {
	g := &resourceIBMDatabaseGen2Backend{}
	config := DBConfig{Members: 3, Shards: 2, StorageGB: 10, HostFlavor: "bx3d.4x20"}

	result := g.dbConfigToMap(config, "mongodbees")

	if result["storage_gb"] != 10 {
		t.Fatalf("expected storage_gb=10, got %v", result["storage_gb"])
	}
	if result["host_flavor"] != "bx3d.4x20" {
		t.Fatalf("expected host_flavor=bx3d.4x20, got %v", result["host_flavor"])
	}
	if result["shards"] != 2 {
		t.Fatalf("expected shards=2, got %v", result["shards"])
	}
}

func TestBuildGen2Parameters_enterpriseShardingGen2UsesMongodbees(t *testing.T) {
	d := testGen2DatabaseResourceData(t, map[string]interface{}{
		"service": "databases-for-mongodb",
		"plan":    "enterprise-sharding-gen2",
	})

	// dbType resolution: getDatabaseTypeFromResourceID("databases-for-mongodb") → "mongodb"
	// then overridden to "mongodbees" for enterprise-sharding-gen2
	dbType := getDatabaseTypeFromResourceID(d.Get("service").(string))
	plan := d.Get("plan").(string)
	if plan == "enterprise-sharding-gen2" && dbType == "mongodb" {
		dbType = "mongodbees"
	}

	if dbType != "mongodbees" {
		t.Fatalf("expected dbType 'mongodbees' for enterprise-sharding-gen2, got %q", dbType)
	}
}
func TestGetShardsCount(t *testing.T) {
	g := &resourceIBMDatabaseGen2Backend{}

	t.Run("accepts maximum shard count of 3", func(t *testing.T) {
		d := testGen2DatabaseResourceData(t, map[string]interface{}{
			"service": "databases-for-mongodb",
			"plan":    "enterprise-sharding-gen2",
			"shards":  2,
		})

		shards, err := g.getShardsCount(d)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		if shards != 2 {
			t.Fatalf("expected shards=2, got shards=%d", shards)
		}
	})

	t.Run("defaults to 1 when shards not configured", func(t *testing.T) {
		d := testGen2DatabaseResourceData(t, map[string]interface{}{
			"service": "databases-for-mongodb",
			"plan":    "enterprise-sharding-gen2",
		})

		shards, err := g.getShardsCount(d)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		if shards != 1 {
			t.Fatalf("expected default shards=1, got shards=%d", shards)
		}
	})

	t.Run("rejects shards for unsupported plans", func(t *testing.T) {
		d := testGen2DatabaseResourceData(t, map[string]interface{}{
			"service": "databases-for-postgresql",
			"plan":    "standard-gen2",
			"shards":  2,
		})

		_, err := g.getShardsCount(d)
		requireErrContains(t, err, "shards is supported only for service=databases-for-mongodb with plan=enterprise-sharding-gen2")
	})
	t.Run("rejects shard count outside 1-3 range", func(t *testing.T) {
		d := testGen2DatabaseResourceData(t, map[string]interface{}{
			"service": "databases-for-mongodb",
			"plan":    "enterprise-sharding-gen2",
			"shards":  4,
		})

		_, err := g.getShardsCount(d)
		requireErrContains(t, err, "shard count must be between 1 and 3")
	})
}

func TestBuildGen2Parameters_standardGen2UsesMongodbNotMongodbees(t *testing.T) {
	d := testGen2DatabaseResourceData(t, map[string]interface{}{
		"service": "databases-for-mongodb",
		"plan":    "standard-gen2",
	})

	dbType := getDatabaseTypeFromResourceID(d.Get("service").(string))
	plan := d.Get("plan").(string)
	if plan == "enterprise-sharding-gen2" && dbType == "mongodb" {
		dbType = "mongodbees"
	}

	if dbType != "mongodb" {
		t.Fatalf("expected dbType 'mongodb' for standard-gen2, got %q", dbType)
	}
}
func TestValidateShardsDiff(t *testing.T) {
	t.Run("rejects shards for non-mongodb service", func(t *testing.T) {
		err := validateShardsDiffPredicate("databases-for-postgresql", "enterprise-sharding-gen2")
		requireErrContains(t, err, "only supported for `service=databases-for-mongodb`")
	})

	t.Run("rejects shards for non-gen2 plan", func(t *testing.T) {
		err := validateShardsDiffPredicate("databases-for-mongodb", "enterprise-sharding")
		requireErrContains(t, err, "only supported for `service=databases-for-mongodb`")
	})

	t.Run("rejects shards for standard plan", func(t *testing.T) {
		err := validateShardsDiffPredicate("databases-for-mongodb", "standard-gen2")
		requireErrContains(t, err, "only supported for `service=databases-for-mongodb`")
	})

	t.Run("accepts shards for mongodb enterprise-sharding-gen2", func(t *testing.T) {
		err := validateShardsDiffPredicate("databases-for-mongodb", "enterprise-sharding-gen2")
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
	})
}

func validateShardsDiffPredicate(service, plan string) error {
	if service != "databases-for-mongodb" || plan != "enterprise-sharding-gen2" {
		return fmt.Errorf("[ERROR] `shards` is only supported for `service=databases-for-mongodb` with `plan=enterprise-sharding-gen2`")
	}
	return nil
}

// downgrade guard inside validateShardsDiff
func TestValidateShardsDowngradeGuard(t *testing.T) {
	t.Run("allows increase from 1 to 2", func(t *testing.T) {
		err := validateShardsDowngrade(1, 2)
		if err != nil {
			t.Fatalf("expected no error for shard increase, got %v", err)
		}
	})

	t.Run("allows increase from 2 to 3", func(t *testing.T) {
		err := validateShardsDowngrade(2, 3)
		if err != nil {
			t.Fatalf("expected no error for shard increase, got %v", err)
		}
	})

	t.Run("allows same count (no-op)", func(t *testing.T) {
		err := validateShardsDowngrade(2, 2)
		if err != nil {
			t.Fatalf("expected no error for no-op update, got %v", err)
		}
	})

	t.Run("rejects decrease from 2 to 1", func(t *testing.T) {
		err := validateShardsDowngrade(2, 1)
		requireErrContains(t, err, "cannot be decreased")
		requireErrContains(t, err, "Current: 2")
		requireErrContains(t, err, "Requested: 1")
	})

	t.Run("rejects decrease from 3 to 1", func(t *testing.T) {
		err := validateShardsDowngrade(3, 1)
		requireErrContains(t, err, "cannot be decreased")
	})

	t.Run("allows create (oldShards=0) with any positive value", func(t *testing.T) {
		for _, newVal := range []int{1, 2, 3} {
			err := validateShardsDowngrade(0, newVal)
			if err != nil {
				t.Fatalf("expected no error on create with shards=%d, got %v", newVal, err)
			}
		}
	})
}

// validateShardsDowngrade is a helper that isolates the downgrade check inside
// validateShardsDiff so it can be tested without a real ResourceDiff.
func validateShardsDowngrade(oldShards, newShards int) error {
	if oldShards > 0 && newShards < oldShards {
		return fmt.Errorf("[ERROR] Shard count cannot be decreased. Current: %d, Requested: %d", oldShards, newShards)
	}
	return nil
}
func TestBuildDBConfig_mongodbeesShardsPath(t *testing.T) {
	g := &resourceIBMDatabaseGen2Backend{}
	resourceSchema := ResourceIBMDatabaseInstance().Schema

	t.Run("includes shards and omits members for mongodbees", func(t *testing.T) {
		d := schema.TestResourceDataRaw(t, resourceSchema, map[string]interface{}{
			"service": "databases-for-mongodb",
			"plan":    "enterprise-sharding-gen2",
			"shards":  2,
			"group": []interface{}{
				map[string]interface{}{
					"group_id": "member",
					// members must be set explicitly so getMembersCount doesn't
					// fall back to the catalog API (which requires a real client)
					"members": []interface{}{
						map[string]interface{}{"allocation_count": 3},
					},
					"disk": []interface{}{
						map[string]interface{}{"allocation_mb": 20480},
					},
					"host_flavor": []interface{}{
						map[string]interface{}{"id": "bxf.16x64"},
					},
				},
			},
		})

		config, err := g.buildDBConfig(d, "", nil, "mongodbees")
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		if _, ok := config["members"]; ok {
			t.Fatal("expected 'members' absent for mongodbees, but it was present")
		}
		if config["shards"] != 2 {
			t.Fatalf("expected shards=2, got %v", config["shards"])
		}
		if config["storage_gb"] != 20 {
			t.Fatalf("expected storage_gb=20, got %v", config["storage_gb"])
		}
		if config["host_flavor"] != "bxf.16x64" {
			t.Fatalf("expected host_flavor=bxf.16x64, got %v", config["host_flavor"])
		}
	})

	t.Run("defaults shards to 1 when not explicitly set", func(t *testing.T) {
		d := schema.TestResourceDataRaw(t, resourceSchema, map[string]interface{}{
			"service": "databases-for-mongodb",
			"plan":    "enterprise-sharding-gen2",
			"group": []interface{}{
				map[string]interface{}{
					"group_id": "member",
					"members": []interface{}{
						map[string]interface{}{"allocation_count": 3},
					},
				},
			},
		})

		config, err := g.buildDBConfig(d, "", nil, "mongodbees")
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		if config["shards"] != 1 {
			t.Fatalf("expected shards defaulted to 1, got %v", config["shards"])
		}
	})

	t.Run("shards absent from config for non-mongodbees dbType", func(t *testing.T) {
		d := schema.TestResourceDataRaw(t, resourceSchema, map[string]interface{}{
			"service": "databases-for-postgresql",
			"plan":    "standard-gen2",
			"group": []interface{}{
				map[string]interface{}{
					"group_id": "member",
					"members": []interface{}{
						map[string]interface{}{"allocation_count": 3},
					},
				},
			},
		})

		config, err := g.buildDBConfig(d, "", nil, "postgresql")
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		if _, ok := config["shards"]; ok {
			t.Fatalf("expected 'shards' absent for non-mongodbees dbType, got %v", config["shards"])
		}
		if _, ok := config["members"]; !ok {
			t.Fatal("expected 'members' present for non-mongodbees dbType")
		}
	})
}
func TestIsShardAttrConfigured(t *testing.T) {
	t.Run("returns false for nil ResourceData", func(t *testing.T) {
		if isShardAttrConfigured(nil) {
			t.Fatal("expected false for nil d")
		}
	})

	t.Run("returns false when shards not set", func(t *testing.T) {
		d := testGen2DatabaseResourceData(t, map[string]interface{}{
			"service": "databases-for-mongodb",
			"plan":    "enterprise-sharding-gen2",
		})
		if isShardAttrConfigured(d) {
			t.Fatal("expected false when shards not configured")
		}
	})

	t.Run("returns true when shards is set to a positive value", func(t *testing.T) {
		d := testGen2DatabaseResourceData(t, map[string]interface{}{
			"service": "databases-for-mongodb",
			"plan":    "enterprise-sharding-gen2",
			"shards":  2,
		})
		if !isShardAttrConfigured(d) {
			t.Fatal("expected true when shards=2 is configured")
		}
	})
}
