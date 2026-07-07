package database

import (
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

	t.Run("unsupported attr present returns error", func(t *testing.T) {
		d := testGen2DatabaseResourceData(t, map[string]interface{}{
			"adminpassword": "very-secure-password-123",
		})

		err := g.ValidateUnsupportedAttrsData(d)

		requireErrContains(t, err, "adminpassword")
		requireErrContains(t, err, "not supported")
	})

	t.Run("multiple unsupported attrs present are all listed", func(t *testing.T) {
		d := testGen2DatabaseResourceData(t, map[string]interface{}{
			"adminpassword":             "very-secure-password-123",
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
			"adminpassword": "very-secure-password-123",
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

	d := testGen2DatabaseResourceData(t, map[string]interface{}{
		"adminpassword":               "very-secure-password-123",
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
