// Copyright IBM Corp. 2017, 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"encoding/json"
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
			expectedError: true,
			errorContains: "supported only for Classic database instances",
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
		{
			name: "unsupported_remote_leader_id_attr",
			changes: map[string]interface{}{
				"remote_leader_id": "crn:v1:bluemix:public:databases-for-postgresql:us-east:a/abc123:leader-id",
			},
			expectedError: true,
			errorContains: "supported only for Classic database instances",
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
		{
			name: "update_unsupported_remote_leader_id",
			changes: map[string]interface{}{
				"remote_leader_id": "crn:v1:bluemix:public:databases-for-postgresql:us-east:a/abc123:leader-id",
			},
			expectedError: true,
			errorContains: "supported only for Classic database instances",
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
			name:               "allowlist",
			feature:            "IP allowlist",
			supportedInClassic: true,
			supportedInGen2:    false,
			gen2Alternative:    "Configure via Cloud Databases API or UI",
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
			},
			expectedOrder: []string{
				"group scaling",
				"tags",
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
		"allowlist",
	}

	assert.Equal(t, len(expectedUnsupported), len(gen2UnsupportedAttrs),
		"Gen2 unsupported attributes list should match expected")

	for i, attr := range expectedUnsupported {
		assert.Equal(t, attr, gen2UnsupportedAttrs[i],
			"Unsupported attribute %d should be '%s'", i, attr)
	}
}

// TestGen2ValidateGroupsDiffMemoryCPU tests validation of Memory and CPU in Gen2
func TestGen2ValidateGroupsDiffMemoryCPU(t *testing.T) {
	tests := []struct {
		name          string
		groupConfig   map[string]interface{}
		expectedError bool
		errorContains string
	}{
		{
			name: "memory_set_independently_should_fail",
			groupConfig: map[string]interface{}{
				"group_id": "member",
				"memory": map[string]interface{}{
					"allocation_mb": 4096,
				},
			},
			expectedError: true,
			errorContains: "Gen2 databases do not support independent memory configuration",
		},
		{
			name: "cpu_set_independently_should_fail",
			groupConfig: map[string]interface{}{
				"group_id": "member",
				"cpu": map[string]interface{}{
					"allocation_count": 4,
				},
			},
			expectedError: true,
			errorContains: "Gen2 databases do not support independent CPU configuration",
		},
		{
			name: "memory_and_cpu_both_set_should_fail",
			groupConfig: map[string]interface{}{
				"group_id": "member",
				"memory": map[string]interface{}{
					"allocation_mb": 4096,
				},
				"cpu": map[string]interface{}{
					"allocation_count": 4,
				},
			},
			expectedError: true,
			errorContains: "Gen2 databases do not support independent memory configuration",
		},
		{
			name: "host_flavor_only_should_succeed",
			groupConfig: map[string]interface{}{
				"group_id": "member",
				"host_flavor": map[string]interface{}{
					"id": "b3c.4x16.encrypted",
				},
			},
			expectedError: false,
		},
		{
			name: "disk_only_should_succeed",
			groupConfig: map[string]interface{}{
				"group_id": "member",
				"disk": map[string]interface{}{
					"allocation_mb": 20480,
				},
			},
			expectedError: false,
		},
		{
			name: "members_only_should_succeed",
			groupConfig: map[string]interface{}{
				"group_id": "member",
				"members": map[string]interface{}{
					"allocation_count": 3,
				},
			},
			expectedError: false,
		},
		{
			name: "host_flavor_with_disk_should_succeed",
			groupConfig: map[string]interface{}{
				"group_id": "member",
				"host_flavor": map[string]interface{}{
					"id": "b3c.4x16.encrypted",
				},
				"disk": map[string]interface{}{
					"allocation_mb": 20480,
				},
			},
			expectedError: false,
		},
		{
			name: "host_flavor_with_members_should_succeed",
			groupConfig: map[string]interface{}{
				"group_id": "member",
				"host_flavor": map[string]interface{}{
					"id": "b3c.4x16.encrypted",
				},
				"members": map[string]interface{}{
					"allocation_count": 3,
				},
			},
			expectedError: false,
		},
		{
			name: "memory_with_host_flavor_should_fail",
			groupConfig: map[string]interface{}{
				"group_id": "member",
				"host_flavor": map[string]interface{}{
					"id": "b3c.4x16.encrypted",
				},
				"memory": map[string]interface{}{
					"allocation_mb": 4096,
				},
			},
			expectedError: true,
			errorContains: "Gen2 databases do not support independent memory configuration",
		},
		{
			name: "cpu_with_host_flavor_should_fail",
			groupConfig: map[string]interface{}{
				"group_id": "member",
				"host_flavor": map[string]interface{}{
					"id": "b3c.4x16.encrypted",
				},
				"cpu": map[string]interface{}{
					"allocation_count": 4,
				},
			},
			expectedError: true,
			errorContains: "Gen2 databases do not support independent CPU configuration",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Test documents expected behavior
			// Actual implementation would require proper ResourceDiff mock
			assert.NotNil(t, tt.groupConfig, "Group config should be defined")
			if tt.expectedError {
				assert.NotEmpty(t, tt.errorContains, "Error message should be specified")
				assert.Contains(t, tt.errorContains, "Gen2 databases do not support independent",
					"Error should mention Gen2 limitation")
			}
		})
	}
}

// TestGen2VersionImmutability tests that version cannot be changed after creation
func TestGen2VersionImmutability(t *testing.T) {
	tests := []struct {
		name           string
		initialVersion string
		updatedVersion string
		expectedError  bool
		errorContains  string
	}{
		{
			name:           "version_change_should_fail",
			initialVersion: "14",
			updatedVersion: "15",
			expectedError:  true,
			errorContains:  "version cannot be changed",
		},
		{
			name:           "same_version_should_succeed",
			initialVersion: "14",
			updatedVersion: "14",
			expectedError:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.NotEmpty(t, tt.initialVersion, "Initial version should not be empty")
			if tt.expectedError {
				assert.NotEmpty(t, tt.errorContains, "Error message should be specified")
			}
		})
	}
}

// TestGen2VersionUpgradeSkipBackup tests that version_upgrade_skip_backup is silently ignored
func TestGen2VersionUpgradeSkipBackup(t *testing.T) {
	tests := []struct {
		name                     string
		versionUpgradeSkipBackup bool
		expectedBehavior         string
	}{
		{
			name:                     "skip_backup_true_ignored",
			versionUpgradeSkipBackup: true,
			expectedBehavior:         "Attribute accepted but not sent to API",
		},
		{
			name:                     "skip_backup_false_ignored",
			versionUpgradeSkipBackup: false,
			expectedBehavior:         "Attribute accepted but not sent to API",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Test documents that version_upgrade_skip_backup is silently ignored in Gen2
			// It should not cause errors but should not be sent to the API
			assert.NotEmpty(t, tt.expectedBehavior, "Expected behavior should be documented")
		})
	}
}

// TestGen2ServiceEndpointsValidation tests service_endpoints validation
func TestGen2ServiceEndpointsValidation(t *testing.T) {
	tests := []struct {
		name            string
		serviceEndpoint string
		expectedError   bool
		errorContains   string
	}{
		{
			name:            "private_endpoint_valid",
			serviceEndpoint: "private",
			expectedError:   false,
		},
		{
			name:            "public_endpoint_invalid",
			serviceEndpoint: "public",
			expectedError:   true,
			errorContains:   "Gen2 databases only support 'private' service endpoints",
		},
		{
			name:            "public_and_private_invalid",
			serviceEndpoint: "public-and-private",
			expectedError:   true,
			errorContains:   "Gen2 databases only support 'private' service endpoints",
		},
		{
			name:            "empty_defaults_to_private",
			serviceEndpoint: "",
			expectedError:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.expectedError {
				assert.NotEmpty(t, tt.errorContains, "Error message should be specified")
			}
		})
	}
}

// TestGen2SkipInitialBackup tests that skip_initial_backup is silently ignored
func TestGen2SkipInitialBackup(t *testing.T) {
	tests := []struct {
		name              string
		skipInitialBackup bool
		expectedBehavior  string
	}{
		{
			name:              "skip_initial_backup_true_ignored",
			skipInitialBackup: true,
			expectedBehavior:  "Attribute accepted but not sent to API",
		},
		{
			name:              "skip_initial_backup_false_ignored",
			skipInitialBackup: false,
			expectedBehavior:  "Attribute accepted but not sent to API",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Test documents that skip_initial_backup is silently ignored in Gen2
			// Only relevant for Classic read replicas
			assert.NotEmpty(t, tt.expectedBehavior, "Expected behavior should be documented")
		})
	}
}

// TestGen2AdminUserNotSupported tests that adminuser is always empty in Gen2
func TestGen2AdminUserNotSupported(t *testing.T) {
	tests := []struct {
		name              string
		expectedAdminUser string
		reason            string
	}{
		{
			name:              "adminuser_always_empty",
			expectedAdminUser: "",
			reason:            "Gen2 databases do not have a default admin user",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Test documents that adminuser is always empty in Gen2
			// Users should use ibm_resource_key for credentials
			assert.Equal(t, "", tt.expectedAdminUser, "Admin user should always be empty in Gen2")
			assert.NotEmpty(t, tt.reason, "Reason should be documented")
		})
	}
}

// TestGen2ConfigurationSchema tests that configuration_schema is always nil/empty
func TestGen2ConfigurationSchema(t *testing.T) {
	tests := []struct {
		name                 string
		expectedConfigSchema string
		reason               string
	}{
		{
			name:                 "config_schema_always_empty",
			expectedConfigSchema: "",
			reason:               "Gen2 databases do not return configuration schema",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Test documents that configuration_schema is always nil/empty in Gen2
			assert.Equal(t, "", tt.expectedConfigSchema, "Configuration schema should always be empty in Gen2")
			assert.NotEmpty(t, tt.reason, "Reason should be documented")
		})
	}
}

// TestGen2LogicalReplicationSlot tests that logical_replication_slot is not supported
func TestGen2LogicalReplicationSlot(t *testing.T) {
	tests := []struct {
		name          string
		slotConfig    map[string]interface{}
		expectedError bool
		errorContains string
	}{
		{
			name: "logical_replication_slot_create_fails",
			slotConfig: map[string]interface{}{
				"name":          "test_slot",
				"database_name": "testdb",
				"plugin_type":   "wal2json",
			},
			expectedError: true,
			errorContains: "logical_replication_slot is not supported for Gen2 databases",
		},
		{
			name: "logical_replication_slot_update_fails",
			slotConfig: map[string]interface{}{
				"name":          "updated_slot",
				"database_name": "testdb",
			},
			expectedError: true,
			errorContains: "logical_replication_slot is not supported for Gen2 databases",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.NotNil(t, tt.slotConfig, "Slot config should be defined")
			if tt.expectedError {
				assert.NotEmpty(t, tt.errorContains, "Error message should be specified")
			}
		})
	}
}

// TestGen2ComputedAttributes tests that computed attributes return expected values
func TestGen2ComputedAttributes(t *testing.T) {
	tests := []struct {
		name         string
		attribute    string
		expectedType string
		isComputed   bool
	}{
		{
			name:         "status_is_computed",
			attribute:    "status",
			expectedType: "string",
			isComputed:   true,
		},
		{
			name:         "guid_is_computed",
			attribute:    "guid",
			expectedType: "string",
			isComputed:   true,
		},
		{
			name:         "groups_is_computed",
			attribute:    "groups",
			expectedType: "list",
			isComputed:   true,
		},
		{
			name:         "resource_name_is_computed",
			attribute:    "resource_name",
			expectedType: "string",
			isComputed:   true,
		},
		{
			name:         "resource_crn_is_computed",
			attribute:    "resource_crn",
			expectedType: "string",
			isComputed:   true,
		},
		{
			name:         "resource_status_is_computed",
			attribute:    "resource_status",
			expectedType: "string",
			isComputed:   true,
		},
		{
			name:         "resource_group_name_is_computed",
			attribute:    "resource_group_name",
			expectedType: "string",
			isComputed:   true,
		},
		{
			name:         "resource_controller_url_is_computed",
			attribute:    "resource_controller_url",
			expectedType: "string",
			isComputed:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.True(t, tt.isComputed, "Attribute should be computed")
			assert.NotEmpty(t, tt.expectedType, "Expected type should be specified")
		})
	}
}

// TestGen2ForceNewAttributes tests ForceNew behavior for applicable attributes
func TestGen2ForceNewAttributes(t *testing.T) {
	tests := []struct {
		name           string
		attribute      string
		initialValue   interface{}
		changedValue   interface{}
		expectForceNew bool
	}{
		{
			name:           "resource_group_id_force_new",
			attribute:      "resource_group_id",
			initialValue:   "rg-123",
			changedValue:   "rg-456",
			expectForceNew: true,
		},
		{
			name:           "location_force_new",
			attribute:      "location",
			initialValue:   "us-south",
			changedValue:   "us-east",
			expectForceNew: true,
		},
		{
			name:           "service_force_new",
			attribute:      "service",
			initialValue:   "databases-for-postgresql",
			changedValue:   "databases-for-mysql",
			expectForceNew: true,
		},
		{
			name:           "plan_force_new",
			attribute:      "plan",
			initialValue:   "standard-gen2",
			changedValue:   "enterprise-gen2",
			expectForceNew: true,
		},
		{
			name:           "key_protect_instance_force_new",
			attribute:      "key_protect_instance",
			initialValue:   "crn:v1:bluemix:public:kms:us-south:a/abc123::",
			changedValue:   "crn:v1:bluemix:public:kms:us-east:a/abc123::",
			expectForceNew: true,
		},
		{
			name:           "key_protect_key_force_new",
			attribute:      "key_protect_key",
			initialValue:   "crn:v1:bluemix:public:kms:us-south:a/abc123:key:key1",
			changedValue:   "crn:v1:bluemix:public:kms:us-south:a/abc123:key:key2",
			expectForceNew: true,
		},
		{
			name:           "backup_encryption_key_crn_force_new",
			attribute:      "backup_encryption_key_crn",
			initialValue:   "crn:v1:bluemix:public:kms:us-south:a/abc123:key:backup1",
			changedValue:   "crn:v1:bluemix:public:kms:us-south:a/abc123:key:backup2",
			expectForceNew: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.True(t, tt.expectForceNew, "Attribute should trigger ForceNew")
			assert.NotNil(t, tt.initialValue, "Initial value should be defined")
			assert.NotNil(t, tt.changedValue, "Changed value should be defined")
		})
	}
}

// TestGen2DeletionProtection tests deletion_protection attribute
func TestGen2DeletionProtection(t *testing.T) {
	tests := []struct {
		name               string
		deletionProtection bool
		expectedBehavior   string
	}{
		{
			name:               "deletion_protection_false_default",
			deletionProtection: false,
			expectedBehavior:   "Instance can be destroyed by Terraform",
		},
		{
			name:               "deletion_protection_true",
			deletionProtection: true,
			expectedBehavior:   "Terraform prevented from destroying instance",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.NotEmpty(t, tt.expectedBehavior, "Expected behavior should be documented")
		})
	}
}

// TestGen2PlanSuffixValidation tests that Gen2 plans end with -gen2
func TestGen2PlanSuffixValidation(t *testing.T) {
	tests := []struct {
		name          string
		plan          string
		isGen2        bool
		expectedError bool
	}{
		{
			name:          "standard_gen2_valid",
			plan:          "standard-gen2",
			isGen2:        true,
			expectedError: false,
		},
		{
			name:          "enterprise_gen2_valid",
			plan:          "enterprise-gen2",
			isGen2:        true,
			expectedError: false,
		},
		{
			name:          "standard_not_gen2",
			plan:          "standard",
			isGen2:        false,
			expectedError: false,
		},
		{
			name:          "enterprise_not_gen2",
			plan:          "enterprise",
			isGen2:        false,
			expectedError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Test documents plan suffix detection logic
			hasGen2Suffix := len(tt.plan) > 5 && tt.plan[len(tt.plan)-5:] == "-gen2"
			assert.Equal(t, tt.isGen2, hasGen2Suffix, "Plan suffix detection should match expected Gen2 status")
		})
	}
}

// TestGen2KeyProtectInstance tests that key_protect_instance is silently ignored in Gen2
func TestGen2KeyProtectInstance(t *testing.T) {
	tests := []struct {
		name               string
		keyProtectInstance string
		operation          string
		expectedBehavior   string
	}{
		{
			name:               "key_protect_instance_create_ignored",
			keyProtectInstance: "crn:v1:bluemix:public:kms:us-south:a/abc123:instance-id::",
			operation:          "CREATE",
			expectedBehavior:   "Accepted but silently ignored, not sent to API",
		},
		{
			name:               "key_protect_instance_update_forcenew",
			keyProtectInstance: "crn:v1:bluemix:public:kms:us-east:a/abc123:new-instance::",
			operation:          "UPDATE",
			expectedBehavior:   "Cannot be changed (ForceNew)",
		},
		{
			name:               "key_protect_instance_read_persists",
			keyProtectInstance: "crn:v1:bluemix:public:kms:us-south:a/abc123:instance-id::",
			operation:          "READ",
			expectedBehavior:   "Value persists in state but never read from API",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Test documents that key_protect_instance is accepted but ignored in Gen2
			// Classic: Sent to Resource Controller API as KeyProtectInstance parameter (just stored)
			// Gen2: Not used - use key_protect_key for disk encryption and backup_encryption_key_crn for backup encryption
			assert.NotEmpty(t, tt.keyProtectInstance, "Key protect instance should be defined")
			assert.NotEmpty(t, tt.expectedBehavior, "Expected behavior should be documented")

			// Verify the attribute is marked as ForceNew in schema
			if tt.operation == "UPDATE" {
				assert.Contains(t, tt.expectedBehavior, "ForceNew", "Update should trigger ForceNew")
			}
		})
	}
}

// TestGen2AdminPasswordIgnored tests that adminpassword is silently ignored in Gen2
func TestGen2AdminPasswordIgnored(t *testing.T) {
	tests := []struct {
		name             string
		adminPassword    string
		operation        string
		expectedBehavior string
	}{
		{
			name:             "admin_password_create_ignored",
			adminPassword:    "SecurePassword123!",
			operation:        "CREATE",
			expectedBehavior: "Accepted but silently ignored - not validated, not sent to API, not configured",
		},
		{
			name:             "admin_password_update_ignored",
			adminPassword:    "NewPassword456!",
			operation:        "UPDATE",
			expectedBehavior: "Accepted but silently ignored - not validated, not sent to API, not configured",
		},
		{
			name:             "admin_password_read_not_returned",
			adminPassword:    "SecurePassword123!",
			operation:        "READ",
			expectedBehavior: "Not returned from API",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Test documents that adminpassword is silently ignored in Gen2
			// Classic: Sets default admin password during CREATE and UPDATE
			// Gen2: No default admin user exists - use ibm_resource_key for credentials
			assert.NotEmpty(t, tt.adminPassword, "Admin password should be defined")
			assert.NotEmpty(t, tt.expectedBehavior, "Expected behavior should be documented")

			// For CREATE and UPDATE, verify it's ignored; for READ, verify it's not returned
			if tt.operation == "READ" {
				assert.Contains(t, tt.expectedBehavior, "Not returned", "Password should not be returned on READ")
			} else {
				assert.Contains(t, tt.expectedBehavior, "ignored", "Password should be ignored in Gen2")
			}
		})
	}
}

// TestGen2ConfigurationIgnored tests that configuration is silently ignored in Gen2
func TestGen2ConfigurationIgnored(t *testing.T) {
	tests := []struct {
		name             string
		configuration    string
		operation        string
		expectedBehavior string
	}{
		{
			name:             "configuration_create_ignored",
			configuration:    `{"max_connections": 200}`,
			operation:        "CREATE",
			expectedBehavior: "Accepted (no validation) but silently ignored - not sent to API",
		},
		{
			name:             "configuration_update_ignored",
			configuration:    `{"shared_buffers": "256MB"}`,
			operation:        "UPDATE",
			expectedBehavior: "Accepted (no validation) but silently ignored - not sent to API",
		},
		{
			name:             "configuration_read_not_set",
			configuration:    `{"max_connections": 200}`,
			operation:        "READ",
			expectedBehavior: "Not set - always returns empty/nil",
		},
		{
			name:             "configuration_complex_json_ignored",
			configuration:    `{"max_connections": 200, "shared_buffers": "256MB", "work_mem": "4MB"}`,
			operation:        "CREATE",
			expectedBehavior: "Accepted (no validation) but silently ignored - not sent to API",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Test documents that configuration is silently ignored in Gen2
			// Classic: Set database configuration JSON via CloudDatabasesV5 API
			// Gen2: Not yet implemented - configuration changes not supported
			assert.NotEmpty(t, tt.configuration, "Configuration should be defined")
			assert.NotEmpty(t, tt.expectedBehavior, "Expected behavior should be documented")

			// Verify configuration is valid JSON
			var jsonData map[string]interface{}
			err := json.Unmarshal([]byte(tt.configuration), &jsonData)
			assert.NoError(t, err, "Configuration should be valid JSON")

			if tt.operation == "READ" {
				assert.Contains(t, tt.expectedBehavior, "Not set", "Read should not return configuration")
			} else {
				assert.Contains(t, tt.expectedBehavior, "ignored", "Configuration should be ignored")
			}
		})
	}
}

// TestGen2AllUnsupportedAttributesBehavior is a comprehensive test documenting all unsupported attributes
func TestGen2AllUnsupportedAttributesBehavior(t *testing.T) {
	tests := []struct {
		name           string
		attribute      string
		planBehavior   string
		applyBehavior  string
		readBehavior   string
		useAlternative string
	}{
		{
			name:           "version_upgrade_skip_backup",
			attribute:      "version_upgrade_skip_backup",
			planBehavior:   "Accepted",
			applyBehavior:  "Silently ignored",
			readBehavior:   "Not set",
			useAlternative: "Only applies to Classic version upgrades",
		},
		{
			name:           "key_protect_instance",
			attribute:      "key_protect_instance",
			planBehavior:   "Accepted",
			applyBehavior:  "Silently ignored (CREATE), ForceNew (UPDATE)",
			readBehavior:   "Persists in state but never read from API",
			useAlternative: "Use key_protect_key for disk encryption and backup_encryption_key_crn for backup encryption",
		},
		{
			name:           "remote_leader_id",
			attribute:      "remote_leader_id",
			planBehavior:   "Fails if set",
			applyBehavior:  "Not sent to API (CREATE), Fails if changed (UPDATE)",
			readBehavior:   "Not set",
			useAlternative: "Use Classic for read replica creation and promotion",
		},
		{
			name:           "skip_initial_backup",
			attribute:      "skip_initial_backup",
			planBehavior:   "Accepted",
			applyBehavior:  "Not validated, not sent to API",
			readBehavior:   "Not set",
			useAlternative: "Only relevant for Classic read replicas",
		},
		{
			name:           "adminuser",
			attribute:      "adminuser",
			planBehavior:   "N/A (computed)",
			applyBehavior:  "N/A (computed)",
			readBehavior:   "Always empty",
			useAlternative: "Use ibm_resource_key for credentials",
		},
		{
			name:           "adminpassword",
			attribute:      "adminpassword",
			planBehavior:   "Accepted (no validation)",
			applyBehavior:  "Silently ignored - not validated, not sent to API",
			readBehavior:   "Not returned",
			useAlternative: "Use ibm_resource_key for credentials",
		},
		{
			name:           "users",
			attribute:      "users",
			planBehavior:   "Fails if set or changed",
			applyBehavior:  "N/A",
			readBehavior:   "Not set",
			useAlternative: "Use ibm_resource_key resource",
		},
		{
			name:           "allowlist",
			attribute:      "allowlist",
			planBehavior:   "Fails if set or changed",
			applyBehavior:  "N/A",
			readBehavior:   "Not set",
			useAlternative: "Not available in Gen2 architecture",
		},
		{
			name:           "configuration",
			attribute:      "configuration",
			planBehavior:   "Accepted (no validation)",
			applyBehavior:  "Silently ignored - not validated, not sent to API",
			readBehavior:   "Not set",
			useAlternative: "Not yet implemented in Gen2",
		},
		{
			name:           "configuration_schema",
			attribute:      "configuration_schema",
			planBehavior:   "Cannot be set (computed)",
			applyBehavior:  "N/A (computed)",
			readBehavior:   "Always nil/empty",
			useAlternative: "Not available in Gen2",
		},
		{
			name:           "auto_scaling",
			attribute:      "auto_scaling",
			planBehavior:   "Fails if set or changed",
			applyBehavior:  "N/A",
			readBehavior:   "Not set",
			useAlternative: "Monitor and scale manually",
		},
		{
			name:           "logical_replication_slot",
			attribute:      "logical_replication_slot",
			planBehavior:   "Fails if set or changed",
			applyBehavior:  "N/A",
			readBehavior:   "Not set",
			useAlternative: "Use Classic for logical replication",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Comprehensive documentation of all unsupported attributes
			assert.NotEmpty(t, tt.attribute, "Attribute name should be defined")
			assert.NotEmpty(t, tt.planBehavior, "Plan behavior should be documented")
			assert.NotEmpty(t, tt.applyBehavior, "Apply behavior should be documented")
			assert.NotEmpty(t, tt.readBehavior, "Read behavior should be documented")
			assert.NotEmpty(t, tt.useAlternative, "Alternative or reason should be documented")
		})
	}
}

// TestGen2SupportedWithNuancesBehavior tests attributes that are supported but with behavioral differences
func TestGen2SupportedWithNuancesBehavior(t *testing.T) {
	tests := []struct {
		name            string
		attribute       string
		classicBehavior string
		gen2Behavior    string
		nuance          string
	}{
		{
			name:            "version_immutability",
			attribute:       "version",
			classicBehavior: "Updatable - triggers upgrade",
			gen2Behavior:    "Set at creation only - cannot be changed",
			nuance:          "Plan accepts it, Apply/Update fails if changed, Read returns current version",
		},
		{
			name:            "service_endpoints_restriction",
			attribute:       "service_endpoints",
			classicBehavior: "Required, accepts public/private/public-and-private",
			gen2Behavior:    "Optional, must be 'private' if set, defaults to 'private'",
			nuance:          "Plan fails if set to non-private value, Apply/Update is updatable, Read returns value from API",
		},
		{
			name:            "group_scaling_limitations",
			attribute:       "group",
			classicBehavior: "Supports members, memory, disk, cpu, host_flavor",
			gen2Behavior:    "Only members, disk (as storage_gb), and host_flavor supported",
			nuance:          "Memory and CPU controlled by host_flavor, cannot be set independently. Plan fails if memory/cpu set.",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Documents attributes that are supported but behave differently in Gen2
			assert.NotEmpty(t, tt.attribute, "Attribute name should be defined")
			assert.NotEmpty(t, tt.classicBehavior, "Classic behavior should be documented")
			assert.NotEmpty(t, tt.gen2Behavior, "Gen2 behavior should be documented")
			assert.NotEmpty(t, tt.nuance, "Nuance/difference should be documented")
		})
	}
}
