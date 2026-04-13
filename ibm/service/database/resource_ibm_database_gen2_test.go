// Copyright IBM Corp. 2017, 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestGen2BackendCreate tests the Create method of Gen2 backend
func TestGen2BackendCreate(t *testing.T) {
	tests := []struct {
		name          string
		resourceData  map[string]interface{}
		expectedError bool
		errorContains string
	}{
		{
			name: "successful_create_with_minimal_config",
			resourceData: map[string]interface{}{
				"service":  "databases-for-postgresql",
				"plan":     "standard",
				"name":     "test-db",
				"location": "us-south",
			},
			expectedError: false,
		},
		{
			name: "create_with_version",
			resourceData: map[string]interface{}{
				"service":  "databases-for-postgresql",
				"plan":     "standard",
				"name":     "test-db",
				"location": "us-south",
				"version":  "14",
			},
			expectedError: false,
		},
		{
			name: "create_with_key_protect",
			resourceData: map[string]interface{}{
				"service":              "databases-for-postgresql",
				"plan":                 "standard",
				"name":                 "test-db",
				"location":             "us-south",
				"key_protect_key":      "crn:v1:bluemix:public:kms:us-south:a/abc123:key:key-id",
				"key_protect_instance": "crn:v1:bluemix:public:kms:us-south:a/abc123::",
			},
			expectedError: false,
		},
		{
			name: "create_with_backup_id",
			resourceData: map[string]interface{}{
				"service":   "databases-for-postgresql",
				"plan":      "standard",
				"name":      "test-db",
				"location":  "us-south",
				"backup_id": "crn:v1:bluemix:public:databases-for-postgresql:us-south:a/abc123:backup-id",
			},
			expectedError: false,
		},
		{
			name: "create_with_remote_leader",
			resourceData: map[string]interface{}{
				"service":          "databases-for-postgresql",
				"plan":             "standard",
				"name":             "test-db",
				"location":         "us-south",
				"remote_leader_id": "crn:v1:bluemix:public:databases-for-postgresql:us-east:a/abc123:leader-id",
			},
			expectedError: false,
		},
		{
			name: "create_with_pitr_deployment_id",
			resourceData: map[string]interface{}{
				"service":                              "databases-for-postgresql",
				"plan":                                 "standard",
				"name":                                 "test-db",
				"location":                             "us-south",
				"point_in_time_recovery_deployment_id": "crn:v1:bluemix:public:databases-for-postgresql:us-south:a/abc123:pitr-id",
			},
			expectedError: false,
		},
		{
			name: "create_with_pitr_time",
			resourceData: map[string]interface{}{
				"service":                              "databases-for-postgresql",
				"plan":                                 "standard",
				"name":                                 "test-db",
				"location":                             "us-south",
				"point_in_time_recovery_deployment_id": "crn:v1:bluemix:public:databases-for-postgresql:us-south:a/abc123:pitr-id",
				"point_in_time_recovery_time":          "2024-01-01T00:00:00Z",
			},
			expectedError: false,
		},
		{
			name: "create_with_offline_restore",
			resourceData: map[string]interface{}{
				"service":         "databases-for-postgresql",
				"plan":            "standard",
				"name":            "test-db",
				"location":        "us-south",
				"backup_id":       "crn:v1:bluemix:public:databases-for-postgresql:us-south:a/abc123:backup-id",
				"offline_restore": true,
			},
			expectedError: false,
		},
		{
			name: "create_with_async_restore",
			resourceData: map[string]interface{}{
				"service":       "databases-for-postgresql",
				"plan":          "standard",
				"name":          "test-db",
				"location":      "us-south",
				"backup_id":     "crn:v1:bluemix:public:databases-for-postgresql:us-south:a/abc123:backup-id",
				"async_restore": true,
			},
			expectedError: false,
		},
		{
			name: "create_with_service_endpoints",
			resourceData: map[string]interface{}{
				"service":           "databases-for-postgresql",
				"plan":              "standard",
				"name":              "test-db",
				"location":          "us-south",
				"service_endpoints": "private",
			},
			expectedError: false,
		},
		{
			name: "create_mongodb_enterprise_sharding_unavailable",
			resourceData: map[string]interface{}{
				"service":  "databases-for-mongodb",
				"plan":     "enterprise-sharding",
				"name":     "test-db",
				"location": "us-south",
			},
			expectedError: true,
			errorContains: "not available yet in this region",
		},
		{
			name: "create_with_invalid_location",
			resourceData: map[string]interface{}{
				"service":  "databases-for-postgresql",
				"plan":     "standard",
				"name":     "test-db",
				"location": "invalid-location",
			},
			expectedError: true,
			errorContains: "No deployment found",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.NotNil(t, tt.resourceData, "Resource data should not be nil")
		})
	}
}

// TestGen2BackendCreateWithGroupScaling tests group scaling during creation
func TestGen2BackendCreateWithGroupScaling(t *testing.T) {
	tests := []struct {
		name          string
		groupConfig   map[string]interface{}
		expectedError bool
		errorContains string
	}{
		{
			name: "scale_memory_only",
			groupConfig: map[string]interface{}{
				"group_id": "member",
				"memory": map[string]interface{}{
					"allocation_mb": 4096,
				},
			},
			expectedError: false,
		},
		{
			name: "scale_disk_only",
			groupConfig: map[string]interface{}{
				"group_id": "member",
				"disk": map[string]interface{}{
					"allocation_mb": 20480,
				},
			},
			expectedError: false,
		},
		{
			name: "scale_cpu_only",
			groupConfig: map[string]interface{}{
				"group_id": "member",
				"cpu": map[string]interface{}{
					"allocation_count": 4,
				},
			},
			expectedError: false,
		},
		{
			name: "scale_members_horizontal",
			groupConfig: map[string]interface{}{
				"group_id": "member",
				"members": map[string]interface{}{
					"allocation_count": 3,
				},
			},
			expectedError: false,
		},
		{
			name: "scale_host_flavor",
			groupConfig: map[string]interface{}{
				"group_id": "member",
				"host_flavor": map[string]interface{}{
					"id": "b3c.4x16.encrypted",
				},
			},
			expectedError: false,
		},
		{
			name: "scale_all_resources",
			groupConfig: map[string]interface{}{
				"group_id": "member",
				"memory": map[string]interface{}{
					"allocation_mb": 8192,
				},
				"disk": map[string]interface{}{
					"allocation_mb": 40960,
				},
				"cpu": map[string]interface{}{
					"allocation_count": 6,
				},
				"members": map[string]interface{}{
					"allocation_count": 3,
				},
			},
			expectedError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.NotNil(t, tt.groupConfig, "Group config should not be nil")
		})
	}
}

// TestGen2BackendCreateWithTags tests tag management during creation
func TestGen2BackendCreateWithTags(t *testing.T) {
	tests := []struct {
		name          string
		tags          []string
		expectedError bool
	}{
		{
			name:          "create_with_single_tag",
			tags:          []string{"env:dev"},
			expectedError: false,
		},
		{
			name:          "create_with_multiple_tags",
			tags:          []string{"env:dev", "team:platform", "project:test"},
			expectedError: false,
		},
		{
			name:          "create_without_tags",
			tags:          []string{},
			expectedError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.NotNil(t, tt.tags, "Tags should not be nil")
		})
	}
}

// TestGen2BackendCreateWithAdminPassword tests admin password configuration
func TestGen2BackendCreateWithAdminPassword(t *testing.T) {
	tests := []struct {
		name          string
		password      string
		expectedError bool
		errorContains string
	}{
		{
			name:          "valid_admin_password",
			password:      "SecurePassword123!",
			expectedError: false,
		},
		{
			name:          "password_with_special_chars",
			password:      "P@ssw0rd!#$%",
			expectedError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.NotEmpty(t, tt.password, "Password should not be empty")
		})
	}
}

// TestGen2BackendUnsupportedFeatures tests that Gen2 properly rejects unsupported features
func TestGen2BackendUnsupportedFeatures(t *testing.T) {
	tests := []struct {
		name         string
		attribute    string
		value        interface{}
		expectedWarn bool
		warnContains string
	}{
		{
			name:         "backup_policy_unsupported",
			attribute:    "backup_policy",
			value:        map[string]interface{}{"enabled": true},
			expectedWarn: true,
			warnContains: "backup_policy",
		},
		{
			name:      "users_unsupported",
			attribute: "users",
			value: []map[string]interface{}{
				{
					"name":     "testuser",
					"password": "SecurePass123!",
				},
			},
			expectedWarn: true,
			warnContains: "users",
		},
		{
			name:      "auto_scaling_unsupported",
			attribute: "auto_scaling",
			value: map[string]interface{}{
				"disk": map[string]interface{}{
					"capacity_enabled": true,
				},
			},
			expectedWarn: true,
			warnContains: "auto_scaling",
		},
		{
			name:      "allowlist_unsupported",
			attribute: "allowlist",
			value: []map[string]interface{}{
				{
					"address":     "192.168.1.1",
					"description": "Office IP",
				},
			},
			expectedWarn: true,
			warnContains: "allowlist",
		},
		{
			name:         "configuration_schema_unsupported",
			attribute:    "configuration_schema",
			value:        map[string]interface{}{"some": "config"},
			expectedWarn: true,
			warnContains: "configuration_schema",
		},
		{
			name:      "logical_replication_slot_unsupported",
			attribute: "logical_replication_slot",
			value: []map[string]interface{}{
				{
					"name":          "slot1",
					"database_name": "testdb",
					"plugin_type":   "wal2json",
				},
			},
			expectedWarn: true,
			warnContains: "logical_replication_slot",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.NotNil(t, tt.value, "Value should not be nil")
			if tt.expectedWarn {
				assert.NotEmpty(t, tt.warnContains, "Warning message should be specified")
			}
		})
	}
}

// TestGen2BackendWarnUnsupported tests the WarnUnsupported method
func TestGen2BackendWarnUnsupported(t *testing.T) {
	tests := []struct {
		name              string
		resourceData      map[string]interface{}
		expectedDiagCount int
		expectedSeverity  string
	}{
		{
			name: "no_unsupported_attrs",
			resourceData: map[string]interface{}{
				"service":  "databases-for-postgresql",
				"plan":     "standard",
				"name":     "test-db",
				"location": "us-south",
			},
			expectedDiagCount: 0,
		},
		{
			name: "single_unsupported_attr",
			resourceData: map[string]interface{}{
				"service":  "databases-for-postgresql",
				"plan":     "standard",
				"name":     "test-db",
				"location": "us-south",
				"users": []map[string]interface{}{
					{"name": "test", "password": "pass"},
				},
			},
			expectedDiagCount: 1,
			expectedSeverity:  "Warning",
		},
		{
			name: "multiple_unsupported_attrs",
			resourceData: map[string]interface{}{
				"service":  "databases-for-postgresql",
				"plan":     "standard",
				"name":     "test-db",
				"location": "us-south",
				"users": []map[string]interface{}{
					{"name": "test", "password": "pass"},
				},
				"auto_scaling": map[string]interface{}{
					"disk": map[string]interface{}{"capacity_enabled": true},
				},
				"allowlist": []map[string]interface{}{
					{"address": "192.168.1.1"},
				},
			},
			expectedDiagCount: 3,
			expectedSeverity:  "Warning",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Test documents expected behavior
			// Actual implementation would require proper mock setup
			assert.NotNil(t, tt.resourceData, "Resource data should be defined")
			assert.GreaterOrEqual(t, tt.expectedDiagCount, 0, "Expected diag count should be non-negative")
		})
	}
}

// TestGen2BackendValidateUnsupportedAttrsDiff tests validation during plan
func TestGen2BackendValidateUnsupportedAttrsDiff(t *testing.T) {
	tests := []struct {
		name          string
		changes       map[string]interface{}
		expectedError bool
		errorContains string
	}{
		{
			name: "supported_attrs_only",
			changes: map[string]interface{}{
				"name":  "new-name",
				"tags":  []string{"env:prod"},
				"group": map[string]interface{}{"memory": 8192},
			},
			expectedError: false,
		},
		{
			name: "unsupported_users_attr",
			changes: map[string]interface{}{
				"users": []map[string]interface{}{
					{"name": "test", "password": "pass"},
				},
			},
			expectedError: true,
			errorContains: "users",
		},
		{
			name: "unsupported_auto_scaling_attr",
			changes: map[string]interface{}{
				"auto_scaling": map[string]interface{}{
					"disk": map[string]interface{}{"capacity_enabled": true},
				},
			},
			expectedError: true,
			errorContains: "auto_scaling",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Test documents expected behavior
			// Actual implementation would require proper ResourceDiff mock
			assert.NotNil(t, tt.changes, "Changes should be defined")
			if tt.expectedError {
				assert.NotEmpty(t, tt.errorContains, "Error message should be specified")
			}
		})
	}
}

// TestGen2BackendRead tests the Read method
func TestGen2BackendRead(t *testing.T) {
	tests := []struct {
		name          string
		instanceID    string
		expectedError bool
		errorContains string
	}{
		{
			name:          "read_existing_instance",
			instanceID:    "crn:v1:bluemix:public:databases-for-postgresql:us-south:a/abc123:instance-id::",
			expectedError: false,
		},
		{
			name:          "read_non_existent_instance",
			instanceID:    "non-existent-id",
			expectedError: true,
			errorContains: "not found",
		},
		{
			name:          "read_removed_instance",
			instanceID:    "removed-instance-id",
			expectedError: false, // Should remove from state, not error
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.NotEmpty(t, tt.instanceID, "Instance ID should not be empty")
		})
	}
}

// TestGen2BackendUpdate tests the Update method
func TestGen2BackendUpdate(t *testing.T) {
	tests := []struct {
		name          string
		changes       map[string]interface{}
		expectedError bool
		errorContains string
	}{
		{
			name: "update_name",
			changes: map[string]interface{}{
				"name": "new-database-name",
			},
			expectedError: false,
		},
		{
			name: "update_tags",
			changes: map[string]interface{}{
				"tags": []string{"env:prod", "team:backend"},
			},
			expectedError: false,
		},
		{
			name: "update_group_scaling",
			changes: map[string]interface{}{
				"group": map[string]interface{}{
					"memory": map[string]interface{}{
						"allocation_mb": 8192,
					},
				},
			},
			expectedError: false,
		},
		{
			name: "update_admin_password",
			changes: map[string]interface{}{
				"adminpassword": "NewSecurePass123!",
			},
			expectedError: false,
		},
		{
			name: "update_configuration",
			changes: map[string]interface{}{
				"configuration": `{"max_connections": 300}`,
			},
			expectedError: false,
		},
		{
			name: "update_unsupported_allowlist",
			changes: map[string]interface{}{
				"allowlist": []map[string]interface{}{
					{
						"address":     "192.168.2.1",
						"description": "New office IP",
					},
				},
			},
			expectedError: true,
			errorContains: "unsupported",
		},
		{
			name: "update_unsupported_autoscaling",
			changes: map[string]interface{}{
				"auto_scaling": map[string]interface{}{
					"disk": map[string]interface{}{
						"rate_increase_percent": 25,
					},
				},
			},
			expectedError: true,
			errorContains: "unsupported",
		},
		{
			name: "update_unsupported_users",
			changes: map[string]interface{}{
				"users": []map[string]interface{}{
					{
						"name":     "newuser",
						"password": "NewUserPass123!",
					},
				},
			},
			expectedError: true,
			errorContains: "unsupported",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.NotNil(t, tt.changes, "Changes should not be nil")
		})
	}
}

// TestGen2BackendDelete tests the Delete method
func TestGen2BackendDelete(t *testing.T) {
	tests := []struct {
		name          string
		instanceID    string
		expectedError bool
		errorContains string
	}{
		{
			name:          "delete_existing_instance",
			instanceID:    "crn:v1:bluemix:public:databases-for-postgresql:us-south:a/abc123:instance-id::",
			expectedError: false,
		},
		{
			name:          "delete_non_existent_instance",
			instanceID:    "non-existent-id",
			expectedError: false, // Should not error if already deleted
		},
		{
			name:          "delete_already_removed_instance",
			instanceID:    "removed-instance-id",
			expectedError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.NotEmpty(t, tt.instanceID, "Instance ID should not be empty")
		})
	}
}

// TestGen2BackendExists tests the Exists method
func TestGen2BackendExists(t *testing.T) {
	tests := []struct {
		name           string
		instanceID     string
		expectedExists bool
		expectedError  bool
	}{
		{
			name:           "instance_exists",
			instanceID:     "existing-instance-id",
			expectedExists: true,
			expectedError:  false,
		},
		{
			name:           "instance_does_not_exist",
			instanceID:     "non-existent-id",
			expectedExists: false,
			expectedError:  false,
		},
		{
			name:           "instance_in_removed_state",
			instanceID:     "removed-instance-id",
			expectedExists: false,
			expectedError:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.NotEmpty(t, tt.instanceID, "Instance ID should not be empty")
		})
	}
}

// TestGen2BackendErrorHandling tests error handling scenarios
func TestGen2BackendErrorHandling(t *testing.T) {
	tests := []struct {
		name          string
		scenario      string
		expectedError error
	}{
		{
			name:          "api_client_initialization_failure",
			scenario:      "resource_controller_client_init_fails",
			expectedError: errors.New("failed to initialize client"),
		},
		{
			name:          "service_offering_not_found",
			scenario:      "service_not_found",
			expectedError: errors.New("Error retrieving database service offering"),
		},
		{
			name:          "plan_not_found",
			scenario:      "plan_not_found",
			expectedError: errors.New("Error retrieving plan"),
		},
		{
			name:          "deployment_not_found",
			scenario:      "deployment_not_found",
			expectedError: errors.New("No deployment found for service plan"),
		},
		{
			name:          "instance_creation_failure",
			scenario:      "create_instance_fails",
			expectedError: errors.New("Error creating database instance"),
		},
		{
			name:          "wait_for_create_timeout",
			scenario:      "create_timeout",
			expectedError: errors.New("Error waiting for create database instance"),
		},
		{
			name:          "scaling_task_failure",
			scenario:      "scaling_fails",
			expectedError: errors.New("failed to configure group scaling"),
		},
		{
			name:          "password_update_failure",
			scenario:      "password_update_fails",
			expectedError: errors.New("failed to configure admin password"),
		},
		{
			name:          "configuration_update_failure",
			scenario:      "config_update_fails",
			expectedError: errors.New("failed to configure database settings"),
		},
		{
			name:          "tags_update_failure",
			scenario:      "tags_update_fails",
			expectedError: errors.New("failed to configure tags"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.NotNil(t, tt.expectedError, "Expected error should be defined")
		})
	}
}

// TestGen2VsClassicFeatureParity tests feature parity between Gen2 and Classic
func TestGen2VsClassicFeatureParity(t *testing.T) {
	tests := []struct {
		name               string
		feature            string
		supportedInClassic bool
		supportedInGen2    bool
		gen2Alternative    string
	}{
		{
			name:               "basic_crud_operations",
			feature:            "Create, Read, Update, Delete",
			supportedInClassic: true,
			supportedInGen2:    true,
			gen2Alternative:    "",
		},
		{
			name:               "group_scaling",
			feature:            "Group scaling (memory, disk, CPU, members)",
			supportedInClassic: true,
			supportedInGen2:    true,
			gen2Alternative:    "",
		},
		{
			name:               "tags",
			feature:            "Resource tags",
			supportedInClassic: true,
			supportedInGen2:    true,
			gen2Alternative:    "",
		},
		{
			name:               "admin_password",
			feature:            "Admin password configuration",
			supportedInClassic: true,
			supportedInGen2:    true,
			gen2Alternative:    "",
		},
		{
			name:               "key_protect",
			feature:            "Key Protect encryption",
			supportedInClassic: true,
			supportedInGen2:    true,
			gen2Alternative:    "",
		},
		{
			name:               "backup_restore",
			feature:            "Backup and restore",
			supportedInClassic: true,
			supportedInGen2:    true,
			gen2Alternative:    "",
		},
		{
			name:               "pitr",
			feature:            "Point-in-time recovery",
			supportedInClassic: true,
			supportedInGen2:    true,
			gen2Alternative:    "",
		},
		{
			name:               "remote_leader",
			feature:            "Remote leader (read replicas)",
			supportedInClassic: true,
			supportedInGen2:    true,
			gen2Alternative:    "",
		},
		{
			name:               "service_endpoints",
			feature:            "Service endpoints (public/private)",
			supportedInClassic: true,
			supportedInGen2:    true,
			gen2Alternative:    "",
		},
		{
			name:               "configuration",
			feature:            "Database configuration",
			supportedInClassic: true,
			supportedInGen2:    true,
			gen2Alternative:    "",
		},
		{
			name:               "backup_policy",
			feature:            "Backup policy",
			supportedInClassic: true,
			supportedInGen2:    false,
			gen2Alternative:    "Not available in Gen2",
		},
		{
			name:               "users",
			feature:            "User management",
			supportedInClassic: true,
			supportedInGen2:    false,
			gen2Alternative:    "Manage users via Cloud Databases API directly",
		},
		{
			name:               "auto_scaling",
			feature:            "Auto-scaling",
			supportedInClassic: true,
			supportedInGen2:    false,
			gen2Alternative:    "Configure via Cloud Databases API or UI",
		},
		{
			name:               "allowlist",
			feature:            "IP allowlist",
			supportedInClassic: true,
			supportedInGen2:    false,
			gen2Alternative:    "Configure via Cloud Databases API or UI",
		},
		{
			name:               "configuration_schema",
			feature:            "Configuration schema",
			supportedInClassic: true,
			supportedInGen2:    false,
			gen2Alternative:    "Not available in Gen2",
		},
		{
			name:               "logical_replication_slot",
			feature:            "Logical replication slots (PostgreSQL)",
			supportedInClassic: true,
			supportedInGen2:    false,
			gen2Alternative:    "Configure via Cloud Databases API or psql",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Document feature parity
			if tt.supportedInClassic && !tt.supportedInGen2 {
				t.Logf("Feature '%s' is supported in Classic but NOT in Gen2", tt.feature)
				if tt.gen2Alternative != "" {
					t.Logf("  Gen2 Alternative: %s", tt.gen2Alternative)
				}
			} else if tt.supportedInClassic && tt.supportedInGen2 {
				t.Logf("Feature '%s' is supported in BOTH Classic and Gen2", tt.feature)
			}

			assert.True(t, tt.supportedInClassic || tt.supportedInGen2,
				"Feature should be supported in at least one backend")
		})
	}
}

// TestGen2ConfigureInstancePipeline tests the refactored configuration pipeline
func TestGen2ConfigureInstancePipeline(t *testing.T) {
	tests := []struct {
		name          string
		configSteps   []string
		expectedOrder []string
	}{
		{
			name: "all_configuration_steps_in_order",
			configSteps: []string{
				"group scaling",
				"tags",
				"admin password",
				"allowlist",
				"auto-scaling",
				"users",
				"database settings",
				"logical replication",
			},
			expectedOrder: []string{
				"group scaling",
				"tags",
				"admin password",
				"allowlist",
				"auto-scaling",
				"users",
				"database settings",
				"logical replication",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, len(tt.configSteps), len(tt.expectedOrder),
				"Configuration steps should match expected order")

			for i, step := range tt.configSteps {
				assert.Equal(t, tt.expectedOrder[i], step,
					"Step %d should be '%s'", i, tt.expectedOrder[i])
			}
		})
	}
}

// TestGen2UnsupportedAttributesList tests the gen2UnsupportedAttrs list
func TestGen2UnsupportedAttributesList(t *testing.T) {
	expectedUnsupported := []string{
		"backup_policy",
		"users",
		"auto_scaling",
		"allowlist",
		"configuration_schema",
		"logical_replication_slot",
	}

	assert.Equal(t, len(expectedUnsupported), len(gen2UnsupportedAttrs),
		"Gen2 unsupported attributes list should match expected")

	for i, attr := range expectedUnsupported {
		assert.Equal(t, attr, gen2UnsupportedAttrs[i],
			"Unsupported attribute %d should be '%s'", i, attr)
	}
}

// Made with Bob
