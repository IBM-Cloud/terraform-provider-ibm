package database

import (
	"strings"
	"testing"
)

func TestValidateConfigurationForService_gen2SkipsSDKValidation(t *testing.T) {
	// Gen2-only fields that do not exist in the Classic SDK struct must be accepted
	// without error for any plan matching the gen2 naming pattern.
	gen2OnlyFields := []struct {
		name   string
		config string
	}{
		{"max_worker_processes", `{"max_worker_processes": 8}`},
		{"max_logical_replication_workers", `{"max_logical_replication_workers": 4}`},
		{"pgaudit.log", `{"pgaudit.log": "all"}`},
		{"pgaudit.role", `{"pgaudit.role": "auditor"}`},
		{"boolean log_connections", `{"log_connections": true}`},
		{"boolean log_disconnections", `{"log_disconnections": false}`},
		{"multiple gen2-only fields", `{"max_worker_processes": 8, "max_logical_replication_workers": 4, "pgaudit.log": "all", "pgaudit.role": "auditor"}`},
	}

	for _, tc := range gen2OnlyFields {
		t.Run(tc.name, func(t *testing.T) {
			err := validateConfigurationForService("databases-for-postgresql", "standard-gen2", tc.config)
			if err != nil {
				t.Errorf("expected no error for Gen2 field %q, got: %s", tc.name, err)
			}
		})
	}
}

func TestValidateConfigurationForService_gen2AcceptsClassicFields(t *testing.T) {
	// Fields that exist in Classic must also pass for Gen2.
	err := validateConfigurationForService("databases-for-postgresql", "standard-gen2", `{"max_connections": 100}`)
	if err != nil {
		t.Errorf("expected no error for classic field on Gen2 plan, got: %s", err)
	}
}

func TestValidateConfigurationForService_classicRejectsUnknownFields(t *testing.T) {
	// Unknown fields must still be rejected for Classic plans.
	err := validateConfigurationForService("databases-for-postgresql", "standard", `{"max_worker_processes": 8}`)
	if err == nil {
		t.Fatal("expected error for unknown field on Classic plan, got nil")
	}
	if !strings.Contains(err.Error(), "invalid field") {
		t.Errorf("expected 'invalid field' in error, got: %s", err)
	}
}

func TestValidateConfigurationForService_classicAcceptsKnownFields(t *testing.T) {
	// Known Classic fields must be accepted on a Classic plan.
	err := validateConfigurationForService("databases-for-postgresql", "standard", `{"max_connections": 100}`)
	if err != nil {
		t.Errorf("expected no error for known Classic field, got: %s", err)
	}
}

func TestValidateConfigurationForService_invalidJSON(t *testing.T) {
	// Malformed JSON must be rejected for both Classic and Gen2.
	for _, plan := range []string{"standard", "standard-gen2"} {
		err := validateConfigurationForService("databases-for-postgresql", plan, `{not valid json}`)
		if err == nil {
			t.Errorf("expected error for invalid JSON on plan %q, got nil", plan)
		}
		if !strings.Contains(err.Error(), "configuration JSON invalid") {
			t.Errorf("expected 'configuration JSON invalid' in error for plan %q, got: %s", plan, err)
		}
	}
}

func TestValidateConfigurationForService_unsupportedServiceClassic(t *testing.T) {
	// A service with no configuration schema must be rejected on Classic plans.
	err := validateConfigurationForService("databases-for-elasticsearch", "standard", `{"some_key": 1}`)
	if err == nil {
		t.Fatal("expected error for unsupported service on Classic plan, got nil")
	}
	if !strings.Contains(err.Error(), "configuration is not supported") {
		t.Errorf("expected 'configuration is not supported' in error, got: %s", err)
	}
}

func TestValidateConfigurationForService_unsupportedServiceGen2Passes(t *testing.T) {
	// On Gen2 plans, any service is accepted — the API decides what's valid.
	err := validateConfigurationForService("databases-for-elasticsearch", "standard-gen2", `{"some_key": 1}`)
	if err != nil {
		t.Errorf("expected no error for unsupported service on Gen2 plan, got: %s", err)
	}
}

func TestValidateConfigurationForService_enterpriseShardingGen2(t *testing.T) {
	// enterprise-sharding-gen2 is also a Gen2 plan and must skip SDK validation.
	err := validateConfigurationForService("databases-for-mongodb", "enterprise-sharding-gen2", `{"max_worker_processes": 8}`)
	if err != nil {
		t.Errorf("expected no error for Gen2 field on enterprise-sharding-gen2 plan, got: %s", err)
	}
}
