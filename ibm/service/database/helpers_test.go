// Copyright IBM Corp. 2017, 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"fmt"
	"sort"
	"strings"
	"testing"
	"time"

	"github.com/IBM/cloud-databases-go-sdk/clouddatabasesv5"
	"github.com/IBM/go-sdk-core/v5/core"
	rc "github.com/IBM/platform-services-go-sdk/resourcecontrollerv2"
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/stretchr/testify/require"
)

func TestIsMoreThan24Hours(t *testing.T) {
	mockNow := time.Date(2025, 1, 1, 15, 0, 0, 0, time.UTC)
	helper := TimeoutHelper{Now: mockNow}

	testcases := []struct {
		description string
		duration    time.Duration
		expected    bool
	}{
		{
			description: "When duration is EXACTLY 24 hours, Expect false",
			duration:    24 * time.Hour,
			expected:    false,
		},
		{
			description: "When duration is MORE than 24 hours, Expect true",
			duration:    25 * time.Hour,
			expected:    true,
		},
		{
			description: "When duration is LESS than 24 hours, Expect false",
			duration:    23 * time.Hour,
			expected:    false,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.description, func(t *testing.T) {
			result := helper.isMoreThan24Hours(tc.duration)
			require.Equal(t, tc.expected, result)
		})
	}
}

func TestFutureTimeToISO(t *testing.T) {
	mockNow := time.Date(2025, 1, 1, 10, 0, 0, 0, time.UTC)
	helper := TimeoutHelper{Now: mockNow}

	result := helper.futureTimeToISO(30 * time.Minute)
	expected := strfmt.DateTime(result)

	require.Equal(t, expected, result)
}

func TestCalculateExpirationDatetime(t *testing.T) {
	mockNow := time.Date(2025, 1, 1, 15, 0, 0, 0, time.UTC)
	helper := TimeoutHelper{Now: mockNow}

	expected24Hours := strfmt.DateTime(helper.futureTimeToISO(24 * time.Hour))
	expected20minutes := strfmt.DateTime(helper.futureTimeToISO(20 * time.Minute))

	testcases := []struct {
		description string
		duration    time.Duration
		expected    strfmt.DateTime
	}{
		{
			description: "When duration is EXACTLY 24 hours, Expect an ISO 24 hrs from now",
			duration:    24 * time.Hour,
			expected:    expected24Hours,
		},
		{
			description: "When duration is MORE than 24 hours, Expect an ISO 24 hrs from now as that is the maximum",
			duration:    25 * time.Hour,
			expected:    expected24Hours,
		},
		{
			description: "When duration is LESS than 24 hours, Expect an ISO of now + duration",
			duration:    20 * time.Minute,
			expected:    expected20minutes,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.description, func(t *testing.T) {
			result := helper.calculateExpirationDatetime(tc.duration)
			require.Equal(t, tc.expected, result)
		})
	}
}
func TestExtractShardsFromExtensions(t *testing.T) {
	t.Run("returns shard count from mongodbees extensions", func(t *testing.T) {
		extensions := map[string]interface{}{
			dataservicesKey: map[string]interface{}{
				"mongodbees": map[string]interface{}{
					"shards": float64(2),
				},
			},
		}

		shards := extractShardsFromExtensions(extensions)
		require.Equal(t, 2, shards)
	})

	t.Run("returns zero when shards key is missing", func(t *testing.T) {
		extensions := map[string]interface{}{
			dataservicesKey: map[string]interface{}{
				"mongodbees": map[string]interface{}{},
			},
		}

		shards := extractShardsFromExtensions(extensions)
		require.Equal(t, 0, shards)
	})

	t.Run("returns zero when mongodbees key is absent", func(t *testing.T) {
		extensions := map[string]interface{}{
			dataservicesKey: map[string]interface{}{
				"otherdb": map[string]interface{}{
					"shards": float64(2),
				},
			},
		}

		shards := extractShardsFromExtensions(extensions)
		require.Equal(t, 0, shards)
	})
}

type MockTaskClient struct {
	Tasks []clouddatabasesv5.Task
	Err   error
}

func (m *MockTaskClient) ListDeploymentTasks(opts *clouddatabasesv5.ListDeploymentTasksOptions) (*clouddatabasesv5.Tasks, *core.DetailedResponse, error) {
	if m.Err != nil {
		return nil, nil, m.Err
	}
	return &clouddatabasesv5.Tasks{
		Tasks: m.Tasks,
	}, &core.DetailedResponse{}, nil
}

func TestMatchingTaskInProgress(t *testing.T) {
	str := "2025-05-12T10:00:00Z"
	parsedTime, _ := time.Parse(time.RFC3339, str)
	mockCreatedAt := strfmt.DateTime(parsedTime)

	testcases := []struct {
		description        string
		mockTasks          []clouddatabasesv5.Task
		mockError          error
		instanceID         string
		matchResourceType  string
		expectedInProgress bool
		expectedTask       clouddatabasesv5.Task
		expectError        bool
	}{
		{
			description: "When matching task is running, Expect true and matching task",
			mockTasks: []clouddatabasesv5.Task{
				{
					ID:              core.StringPtr("123"),
					Status:          core.StringPtr(databaseTaskRunningStatus),
					ResourceType:    core.StringPtr(taskRestore),
					CreatedAt:       &mockCreatedAt,
					ProgressPercent: core.Int64Ptr(74),
					Description:     core.StringPtr("Restore running"),
				},
				{
					ID:              core.StringPtr("1234"),
					Status:          core.StringPtr(databaseTaskRunningStatus),
					ResourceType:    core.StringPtr(taskUpgrade),
					CreatedAt:       &mockCreatedAt,
					ProgressPercent: core.Int64Ptr(74),
					Description:     core.StringPtr("Upgrade running"),
				},
			},
			instanceID:         "inst-1",
			matchResourceType:  taskUpgrade,
			expectedInProgress: true,
			expectedTask: clouddatabasesv5.Task{
				ID:              core.StringPtr("1234"),
				Status:          core.StringPtr(databaseTaskRunningStatus),
				ResourceType:    core.StringPtr(taskUpgrade),
				CreatedAt:       &mockCreatedAt,
				ProgressPercent: core.Int64Ptr(74),
				Description:     core.StringPtr("Upgrade running"),
			},
		},
		{
			description: "When matching task is queued, Expect true and matching task",
			mockTasks: []clouddatabasesv5.Task{
				{
					ID:              core.StringPtr("123"),
					Status:          core.StringPtr(databaseTaskQueuedStatus),
					ResourceType:    core.StringPtr(taskRestore),
					CreatedAt:       &mockCreatedAt,
					ProgressPercent: core.Int64Ptr(74),
					Description:     core.StringPtr("Restore running"),
				},
				{
					ID:              core.StringPtr("234"),
					Status:          core.StringPtr(databaseTaskQueuedStatus),
					ResourceType:    core.StringPtr(taskUpgrade),
					CreatedAt:       &mockCreatedAt,
					ProgressPercent: core.Int64Ptr(74),
					Description:     core.StringPtr("Upgrade running"),
				},
			},
			instanceID:         "inst-2",
			matchResourceType:  taskUpgrade,
			expectedInProgress: true,
			expectedTask: clouddatabasesv5.Task{
				ID:              core.StringPtr("234"),
				Status:          core.StringPtr(databaseTaskQueuedStatus),
				ResourceType:    core.StringPtr(taskUpgrade),
				CreatedAt:       &mockCreatedAt,
				ProgressPercent: core.Int64Ptr(74),
				Description:     core.StringPtr("Upgrade running"),
			},
		},
		{
			description: "When matching task is completed, Expect false",
			mockTasks: []clouddatabasesv5.Task{
				{
					ID:              core.StringPtr("101"),
					Status:          core.StringPtr(databaseTaskCompletedStatus),
					ResourceType:    core.StringPtr(taskUpgrade),
					CreatedAt:       &mockCreatedAt,
					ProgressPercent: core.Int64Ptr(74),
					Description:     core.StringPtr("Upgrade running"),
				},
				{
					ID:              core.StringPtr("102"),
					Status:          core.StringPtr(databaseTaskQueuedStatus),
					ResourceType:    core.StringPtr("backup"),
					CreatedAt:       &mockCreatedAt,
					ProgressPercent: core.Int64Ptr(74),
					Description:     core.StringPtr("backup running"),
				},
			},
			instanceID:         "inst-4",
			matchResourceType:  taskUpgrade,
			expectedInProgress: false,
		},
		{
			description: "When matching task is NOT the running task, Expect false",
			mockTasks: []clouddatabasesv5.Task{
				{
					ID:              core.StringPtr("789"),
					Status:          core.StringPtr(databaseTaskRunningStatus),
					ResourceType:    core.StringPtr(taskRestore),
					CreatedAt:       &mockCreatedAt,
					ProgressPercent: core.Int64Ptr(74),
					Description:     core.StringPtr("Restore running"),
				},
			},
			instanceID:         "inst-3",
			matchResourceType:  taskUpgrade,
			expectedInProgress: false,
		},
		{
			description:        "When there is an error getting tasks, Expect error",
			mockError:          fmt.Errorf("API error"),
			instanceID:         "inst-5",
			matchResourceType:  taskUpgrade,
			expectError:        true,
			expectedInProgress: false,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.description, func(t *testing.T) {
			mockClient := &MockTaskClient{
				Tasks: tc.mockTasks,
				Err:   tc.mockError,
			}

			tm := &TaskManager{
				Client:     mockClient,
				InstanceID: tc.instanceID,
			}

			inProgress, task, err := tm.matchingTaskInProgress(tc.matchResourceType)

			if tc.expectError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.expectedInProgress, inProgress)

				if tc.expectedInProgress {
					require.NotNil(t, task)
					require.Equal(t, tc.expectedTask, *task)
				} else {
					require.Nil(t, task)
				}
			}
		})
	}
}

func TestIsGen2Plan(t *testing.T) {
	cases := []struct {
		plan string
		want bool
	}{
		{"databases-for-postgresql-standard", false},
		{"databases-for-postgresql-gen2", true},
		{"databases-for-postgresql-gen2-dev", true},
		{"standard-gen2", true},
		{"standard", false},
		{"", false},
	}
	for _, c := range cases {
		if got := isGen2Plan(c.plan); got != c.want {
			t.Errorf("isGen2Plan(%q) = %v, want %v", c.plan, got, c.want)
		}
	}
}

// TestClearGen2UnsupportedAttributes tests the clearGen2UnsupportedAttributes function
func TestClearGen2UnsupportedAttributes(t *testing.T) {
	adminPasswordValue := "example-admin-value"

	d := schema.TestResourceDataRaw(t, map[string]*schema.Schema{
		"adminuser": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"adminpassword": {
			Type:      schema.TypeString,
			Optional:  true,
			Sensitive: true,
		},
		"auto_scaling": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"enabled": {
						Type:     schema.TypeBool,
						Optional: true,
					},
				},
			},
		},
		"allowlist": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"address": {
						Type:     schema.TypeString,
						Optional: true,
					},
				},
			},
		},
		"users": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"name": {
						Type:     schema.TypeString,
						Optional: true,
					},
				},
			},
		},
		"configuration_schema": {
			Type:     schema.TypeString,
			Optional: true,
		},
	}, map[string]interface{}{
		"adminuser":            "admin",
		"adminpassword":        adminPasswordValue,
		"auto_scaling":         []interface{}{map[string]interface{}{"enabled": true}},
		"allowlist":            []interface{}{map[string]interface{}{"address": "1.2.3.4"}},
		"users":                []interface{}{map[string]interface{}{"name": "user1"}},
		"configuration_schema": "some_schema",
	})

	clearGen2UnsupportedAttributes(d)

	// Verify all Gen2 unsupported attributes are cleared (d.Set(key, nil) results in empty values, not nil)

	adminuser := d.Get("adminuser")
	require.Equal(t, "", adminuser, "adminuser should be empty string after clearing")

	adminpassword := d.Get("adminpassword")
	require.Equal(t, "", adminpassword, "adminpassword should be empty string after clearing")

	autoScaling := d.Get("auto_scaling")
	require.NotNil(t, autoScaling, "auto_scaling should be set to empty value")
	require.Empty(t, autoScaling, "auto_scaling should be empty after clearing")

	allowlist := d.Get("allowlist")
	require.NotNil(t, allowlist, "allowlist should be set to empty value")
	require.Empty(t, allowlist, "allowlist should be empty after clearing")

	users := d.Get("users")
	require.NotNil(t, users, "users should be set to empty value")
	require.Empty(t, users, "users should be empty after clearing")

	configSchema := d.Get("configuration_schema")
	require.Equal(t, "", configSchema, "configuration_schema should be empty string after clearing")

	// Note: platform_options.backup_encryption_key_crn is also not supported in Gen2,
	// but it's handled by the data source implementation which only sets disk_encryption_key_crn
}

func TestExtractDeploymentIDFromCRN(t *testing.T) {
	testcases := []struct {
		description   string
		catalogCRN    string
		expectedID    string
		expectError   bool
		errorContains string
	}{
		{
			description: "Valid CRN with deployment ID",
			catalogCRN:  "crn:v1:bluemix:public:globalcatalog::::deployment:standard-gen2-deployment-ca-mon-11b01c58",
			expectedID:  "standard-gen2-deployment-ca-mon-11b01c58",
			expectError: false,
		},
		{
			description: "Valid CRN with different deployment ID",
			catalogCRN:  "crn:v1:bluemix:public:globalcatalog::::deployment:databases-for-postgresql-standard-us-south",
			expectedID:  "databases-for-postgresql-standard-us-south",
			expectError: false,
		},
		{
			description:   "Invalid CRN - missing deployment prefix",
			catalogCRN:    "crn:v1:bluemix:public:globalcatalog::::standard-gen2-deployment-ca-mon-11b01c58",
			expectError:   true,
			errorContains: "invalid catalog CRN format",
		},
		{
			description:   "Invalid CRN - empty deployment ID",
			catalogCRN:    "crn:v1:bluemix:public:globalcatalog::::deployment:",
			expectError:   true,
			errorContains: "empty deployment ID",
		},
		{
			description:   "Invalid CRN - multiple deployment prefixes",
			catalogCRN:    "crn:v1:bluemix:public:globalcatalog::::deployment:test:deployment:another",
			expectError:   true,
			errorContains: "invalid catalog CRN format",
		},
		{
			description:   "Empty CRN",
			catalogCRN:    "",
			expectError:   true,
			errorContains: "invalid catalog CRN format",
		},
	}

	for _, tc := range testcases {
		t.Run(tc.description, func(t *testing.T) {
			deploymentID, err := extractDeploymentIDFromCRN(tc.catalogCRN)

			if tc.expectError {
				require.Error(t, err)
				require.Contains(t, err.Error(), tc.errorContains)
				require.Empty(t, deploymentID)
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.expectedID, deploymentID)
			}
		})
	}
}

// ---------------------------------------------------------------------------
// isTruthy
// ---------------------------------------------------------------------------

func TestIsTruthy(t *testing.T) {
	cases := []struct {
		name string
		in   interface{}
		want bool
	}{
		{"bool true", true, true},
		{"bool false", false, false},
		{"string true", "true", true},
		{"string false", "false", false},
		{"string TRUE (case sensitive)", "TRUE", false},
		{"nil", nil, false},
		{"int 1", 1, false},
		{"empty string", "", false},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			require.Equal(t, c.want, isTruthy(c.in))
		})
	}
}

// ---------------------------------------------------------------------------
// s2sAuthWarning sentinel error
// ---------------------------------------------------------------------------

func TestS2sAuthWarningError(t *testing.T) {
	w := &s2sAuthWarning{}
	require.Equal(t, s2sAuthWarningHeader, w.Error())
}

// ---------------------------------------------------------------------------
// checkS2SAuthorization
// ---------------------------------------------------------------------------

func TestCheckS2SAuthorization(t *testing.T) {
	cases := []struct {
		name       string
		extensions map[string]interface{}
		want       bool
	}{
		{
			name:       "nil extensions",
			extensions: nil,
			want:       false,
		},
		{
			name:       "empty extensions map",
			extensions: map[string]interface{}{},
			want:       false,
		},
		{
			name: "dataservices key missing",
			extensions: map[string]interface{}{
				"other_key": "value",
			},
			want: false,
		},
		{
			name: "dataservices is not a map",
			extensions: map[string]interface{}{
				"dataservices": "not-a-map",
			},
			want: false,
		},
		{
			name: "authorizations key missing inside dataservices",
			extensions: map[string]interface{}{
				"dataservices": map[string]interface{}{
					"other_key": "value",
				},
			},
			want: false,
		},
		{
			name: "authorizations is nil",
			extensions: map[string]interface{}{
				"dataservices": map[string]interface{}{
					"authorizations": nil,
				},
			},
			want: false,
		},
		{
			name: "authorizations is empty map",
			extensions: map[string]interface{}{
				"dataservices": map[string]interface{}{
					"authorizations": map[string]interface{}{},
				},
			},
			want: false,
		},
		{
			name: "both flags true (bool)",
			extensions: map[string]interface{}{
				"dataservices": map[string]interface{}{
					"authorizations": map[string]interface{}{
						"independent_backups": true,
						"resource_group":      true,
					},
				},
			},
			want: true,
		},
		{
			name: "both flags true (string)",
			extensions: map[string]interface{}{
				"dataservices": map[string]interface{}{
					"authorizations": map[string]interface{}{
						"independent_backups": "true",
						"resource_group":      "true",
					},
				},
			},
			want: true,
		},
		{
			name: "independent_backups false (bool)",
			extensions: map[string]interface{}{
				"dataservices": map[string]interface{}{
					"authorizations": map[string]interface{}{
						"independent_backups": false,
						"resource_group":      true,
					},
				},
			},
			want: false,
		},
		{
			name: "resource_group false (bool)",
			extensions: map[string]interface{}{
				"dataservices": map[string]interface{}{
					"authorizations": map[string]interface{}{
						"independent_backups": true,
						"resource_group":      false,
					},
				},
			},
			want: false,
		},
		{
			name: "both flags false (bool)",
			extensions: map[string]interface{}{
				"dataservices": map[string]interface{}{
					"authorizations": map[string]interface{}{
						"independent_backups": false,
						"resource_group":      false,
					},
				},
			},
			want: false,
		},
		{
			name: "independent_backups missing",
			extensions: map[string]interface{}{
				"dataservices": map[string]interface{}{
					"authorizations": map[string]interface{}{
						"resource_group": true,
					},
				},
			},
			want: false,
		},
		{
			name: "resource_group missing",
			extensions: map[string]interface{}{
				"dataservices": map[string]interface{}{
					"authorizations": map[string]interface{}{
						"independent_backups": true,
					},
				},
			},
			want: false,
		},
		{
			name: "mixed: bool true + string true",
			extensions: map[string]interface{}{
				"dataservices": map[string]interface{}{
					"authorizations": map[string]interface{}{
						"independent_backups": true,
						"resource_group":      "true",
					},
				},
			},
			want: true,
		},
		{
			name: "resource_group string false",
			extensions: map[string]interface{}{
				"dataservices": map[string]interface{}{
					"authorizations": map[string]interface{}{
						"independent_backups": true,
						"resource_group":      "false",
					},
				},
			},
			want: false,
		},
		{
			name: "authorizations is not a map (wrong type)",
			extensions: map[string]interface{}{
				"dataservices": map[string]interface{}{
					"authorizations": "not-a-map",
				},
			},
			want: false,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			require.Equal(t, c.want, checkS2SAuthorization(c.extensions))
		})
	}
}

// ---------------------------------------------------------------------------
// hasIndependentBackups
// ---------------------------------------------------------------------------

func TestHasIndependentBackups(t *testing.T) {
	cases := []struct {
		name       string
		extensions map[string]interface{}
		want       bool
	}{
		{
			name:       "nil extensions",
			extensions: nil,
			want:       false,
		},
		{
			name:       "empty extensions",
			extensions: map[string]interface{}{},
			want:       false,
		},
		{
			name: "dataservices key missing",
			extensions: map[string]interface{}{
				"other_key": "value",
			},
			want: false,
		},
		{
			name: "dataservices is not a map",
			extensions: map[string]interface{}{
				"dataservices": "not-a-map",
			},
			want: false,
		},
		{
			name: "backups key missing inside dataservices",
			extensions: map[string]interface{}{
				"dataservices": map[string]interface{}{
					"other_key": "value",
				},
			},
			want: false,
		},
		{
			name: "backups key is nil",
			extensions: map[string]interface{}{
				"dataservices": map[string]interface{}{
					"backups": nil,
				},
			},
			want: false,
		},
		{
			name: "backups key is present and non-nil",
			extensions: map[string]interface{}{
				"dataservices": map[string]interface{}{
					"backups": map[string]interface{}{
						"automatic_backups": map[string]interface{}{
							"enabled": true,
							"window": map[string]interface{}{
								"start_time": "09:04Z",
							},
						},
						"preserve":  false,
						"retention": "30d",
					},
				},
			},
			want: true,
		},
		{
			name: "backups key is an empty map (still non-nil)",
			extensions: map[string]interface{}{
				"dataservices": map[string]interface{}{
					"backups": map[string]interface{}{},
				},
			},
			want: true,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			require.Equal(t, c.want, hasIndependentBackups(c.extensions))
		})
	}
}

// ---------------------------------------------------------------------------
// s2sAuthWarningHeader / s2sAuthWarningDetail constants are non-empty
// ---------------------------------------------------------------------------

func TestS2SWarningConstants(t *testing.T) {
	require.NotEmpty(t, s2sAuthWarningHeader, "s2sAuthWarningHeader must not be empty")
	require.NotEmpty(t, s2sAuthWarningDetail, "s2sAuthWarningDetail must not be empty")
}
func TestExtractGen2BackupExtensions(t *testing.T) {
	testcases := []struct {
		description        string
		extensions         map[string]interface{}
		expectedSourceCRN  string
		expectedBackupType string
	}{
		{
			description: "Valid extensions with source_data_service_crn and type",
			extensions: map[string]interface{}{
				"dataservices": map[string]interface{}{
					"backup": map[string]interface{}{
						"source_data_service_crn": "crn:v1:bluemix:public:databases-for-postgresql:us-south:a/abc123:deployment-id::",
						"type":                    "on_demand",
					},
				},
			},
			expectedSourceCRN:  "crn:v1:bluemix:public:databases-for-postgresql:us-south:a/abc123:deployment-id::",
			expectedBackupType: "on_demand",
		},
		{
			description: "Valid extensions with scheduled backup type",
			extensions: map[string]interface{}{
				"dataservices": map[string]interface{}{
					"backup": map[string]interface{}{
						"source_data_service_crn": "crn:v1:bluemix:public:databases-for-mysql:us-east:a/abc123:deployment-id::",
						"type":                    "scheduled",
					},
				},
			},
			expectedSourceCRN:  "crn:v1:bluemix:public:databases-for-mysql:us-east:a/abc123:deployment-id::",
			expectedBackupType: "scheduled",
		},
		{
			description:        "Nil extensions",
			extensions:         nil,
			expectedSourceCRN:  "",
			expectedBackupType: "",
		},
		{
			description:        "Empty extensions map",
			extensions:         map[string]interface{}{},
			expectedSourceCRN:  "",
			expectedBackupType: "",
		},
		{
			description: "Missing dataservices key",
			extensions: map[string]interface{}{
				"other_key": "some_value",
			},
			expectedSourceCRN:  "",
			expectedBackupType: "",
		},
		{
			description: "dataservices is not a map",
			extensions: map[string]interface{}{
				"dataservices": "not-a-map",
			},
			expectedSourceCRN:  "",
			expectedBackupType: "",
		},
		{
			description: "Missing backup key within dataservices",
			extensions: map[string]interface{}{
				"dataservices": map[string]interface{}{
					"other_key": "some_value",
				},
			},
			expectedSourceCRN:  "",
			expectedBackupType: "",
		},
		{
			description: "backup is not a map",
			extensions: map[string]interface{}{
				"dataservices": map[string]interface{}{
					"backup": "not-a-map",
				},
			},
			expectedSourceCRN:  "",
			expectedBackupType: "",
		},
		{
			description: "Missing source_data_service_crn field",
			extensions: map[string]interface{}{
				"dataservices": map[string]interface{}{
					"backup": map[string]interface{}{
						"type": "on_demand",
					},
				},
			},
			expectedSourceCRN:  "",
			expectedBackupType: "on_demand",
		},
		{
			description: "Missing type field",
			extensions: map[string]interface{}{
				"dataservices": map[string]interface{}{
					"backup": map[string]interface{}{
						"source_data_service_crn": "crn:v1:bluemix:public:databases-for-postgresql:us-south:a/abc123:deployment-id::",
					},
				},
			},
			expectedSourceCRN:  "crn:v1:bluemix:public:databases-for-postgresql:us-south:a/abc123:deployment-id::",
			expectedBackupType: "",
		},
		{
			description: "source_data_service_crn is not a string",
			extensions: map[string]interface{}{
				"dataservices": map[string]interface{}{
					"backup": map[string]interface{}{
						"source_data_service_crn": 12345,
						"type":                    "on_demand",
					},
				},
			},
			expectedSourceCRN:  "",
			expectedBackupType: "on_demand",
		},
		{
			description: "type is not a string",
			extensions: map[string]interface{}{
				"dataservices": map[string]interface{}{
					"backup": map[string]interface{}{
						"source_data_service_crn": "crn:v1:bluemix:public:databases-for-postgresql:us-south:a/abc123:deployment-id::",
						"type":                    42,
					},
				},
			},
			expectedSourceCRN:  "crn:v1:bluemix:public:databases-for-postgresql:us-south:a/abc123:deployment-id::",
			expectedBackupType: "",
		},
	}

	for _, tc := range testcases {
		t.Run(tc.description, func(t *testing.T) {
			sourceDataServiceCRN, backupType := extractGen2BackupExtensions(tc.extensions)

			require.Equal(t, tc.expectedSourceCRN, sourceDataServiceCRN)
			require.Equal(t, tc.expectedBackupType, backupType)
		})
	}
}

func TestGetInstancesNext(t *testing.T) {
	testcases := []struct {
		description string
		next        *string
		expected    string
		expectError bool
	}{
		{
			description: "Nil next returns empty string and no error",
			next:        nil,
			expected:    "",
		},
		{
			description: "URL with next_url query parameter",
			next:        core.StringPtr("https://api.example.com/v2/resource_instances?next_url=abc123"),
			expected:    "abc123",
		},
		{
			description: "URL without next_url query parameter",
			next:        core.StringPtr("https://api.example.com/v2/resource_instances?start=abc123"),
			expected:    "",
		},
		{
			description: "Empty string URL",
			next:        core.StringPtr(""),
			expected:    "",
		},
		{
			description: "Malformed URL returns error",
			next:        core.StringPtr("https://api.example.com/v2/resource_instances/%zz"),
			expected:    "",
			expectError: true,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.description, func(t *testing.T) {
			result, err := getInstancesNext(tc.next)

			if tc.expectError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
			require.Equal(t, tc.expected, result)
		})
	}
}

// ---------------------------------------------------------------------------
// gen2GetOperationDescription
// ---------------------------------------------------------------------------

func TestGen2GetOperationDescription_LastOperationDescription(t *testing.T) {
	desc := "Provisioning complete"
	instance := &rc.ResourceInstance{
		LastOperation: &rc.ResourceInstanceLastOperation{
			Description: core.StringPtr(desc),
			Type:        core.StringPtr("provision"),
		},
	}
	got := gen2GetOperationDescription(instance)
	if got != desc {
		t.Errorf("expected %q, got %q", desc, got)
	}
}

func TestGen2GetOperationDescription_LastOperationTypeFallback(t *testing.T) {
	instance := &rc.ResourceInstance{
		LastOperation: &rc.ResourceInstanceLastOperation{
			Description: core.StringPtr(""),
			Type:        core.StringPtr("update"),
		},
	}
	got := gen2GetOperationDescription(instance)
	want := "Operation: update"
	if got != want {
		t.Errorf("expected %q, got %q", want, got)
	}
}

func TestGen2GetOperationDescription_StateFallback(t *testing.T) {
	instance := &rc.ResourceInstance{
		State: core.StringPtr("active"),
	}
	got := gen2GetOperationDescription(instance)
	want := "Instance state: active"
	if got != want {
		t.Errorf("expected %q, got %q", want, got)
	}
}

func TestGen2GetOperationDescription_DefaultFallback(t *testing.T) {
	instance := &rc.ResourceInstance{}
	got := gen2GetOperationDescription(instance)
	want := "Gen2 database instance operation"
	if got != want {
		t.Errorf("expected %q, got %q", want, got)
	}
}

// ---------------------------------------------------------------------------
// gen2MapStateToStatus
// ---------------------------------------------------------------------------

func TestGen2MapStateToStatus(t *testing.T) {
	tests := []struct {
		state *string
		want  string
	}{
		{nil, "unknown"},
		{core.StringPtr("active"), "completed"},
		{core.StringPtr("provisioning"), "running"},
		{core.StringPtr("in progress"), "running"},
		{core.StringPtr("removed"), "completed"},
		{core.StringPtr("failed"), "failed"},
		{core.StringPtr("inactive"), "inactive"},
		{core.StringPtr("unknown-state"), "unknown-state"},
	}

	for _, tc := range tests {
		label := "nil"
		if tc.state != nil {
			label = *tc.state
		}
		t.Run(label, func(t *testing.T) {
			instance := &rc.ResourceInstance{State: tc.state}
			got := gen2MapStateToStatus(instance)
			if got != tc.want {
				t.Errorf("state=%q: expected %q, got %q", label, tc.want, got)
			}
		})
	}
}

// ---------------------------------------------------------------------------
// gen2CalculateProgress
// ---------------------------------------------------------------------------

func TestGen2CalculateProgress(t *testing.T) {
	tests := []struct {
		state *string
		want  int
	}{
		{nil, 0},
		{core.StringPtr("active"), 100},
		{core.StringPtr("provisioning"), 50},
		{core.StringPtr("in progress"), 75},
		{core.StringPtr("failed"), 100},
		{core.StringPtr("removed"), 100},
		{core.StringPtr("inactive"), 0},
		{core.StringPtr("unknown-state"), 0},
	}

	for _, tc := range tests {
		label := "nil"
		if tc.state != nil {
			label = *tc.state
		}
		t.Run(label, func(t *testing.T) {
			instance := &rc.ResourceInstance{State: tc.state}
			got := gen2CalculateProgress(instance)
			if got != tc.want {
				t.Errorf("state=%q: expected %d, got %d", label, tc.want, got)
			}
		})
	}
}

// ---------------------------------------------------------------------------
// gen2GetOperationTime
// ---------------------------------------------------------------------------

func TestGen2GetOperationTime_UpdatedAt(t *testing.T) {
	updatedAt := strfmt.DateTime(time.Date(2024, 6, 1, 12, 0, 0, 0, time.UTC))
	createdAt := strfmt.DateTime(time.Date(2024, 6, 1, 11, 0, 0, 0, time.UTC))
	instance := &rc.ResourceInstance{
		UpdatedAt: &updatedAt,
		CreatedAt: &createdAt,
	}
	got := gen2GetOperationTime(instance)
	// UpdatedAt should be preferred over CreatedAt
	if !strings.Contains(got, "2024") {
		t.Errorf("unexpected time string: %q", got)
	}
}

func TestGen2GetOperationTime_CreatedAtFallback(t *testing.T) {
	createdAt := strfmt.DateTime(time.Date(2024, 6, 1, 10, 0, 0, 0, time.UTC))
	instance := &rc.ResourceInstance{
		CreatedAt: &createdAt,
	}
	got := gen2GetOperationTime(instance)
	if !strings.Contains(got, "2024") {
		t.Errorf("unexpected time string: %q", got)
	}
}

func TestGen2GetOperationTime_NowFallback(t *testing.T) {
	before := time.Now().UTC().Add(-2 * time.Second)
	instance := &rc.ResourceInstance{}
	got := gen2GetOperationTime(instance)
	after := time.Now().UTC().Add(2 * time.Second)

	parsed, err := time.Parse(time.RFC3339, got)
	if err != nil {
		t.Fatalf("returned string is not RFC3339: %q", got)
	}
	if parsed.Before(before) || parsed.After(after) {
		t.Errorf("expected time near now, got %q", got)
	}
}

func TestExtractDatabaseAllocations(t *testing.T) {
	t.Run("extracts allocations from mongodb key for standard-gen2", func(t *testing.T) {
		instance := map[string]interface{}{
			"dataservices": map[string]interface{}{
				"mongodb": map[string]interface{}{
					"storage_gb":  float64(30),
					"host_flavor": "bx3d.4x20",
					"members":     float64(3),
				},
			},
		}
		alloc := extractDatabaseAllocations(instance, "databases-for-mongodb", "standard-gen2")
		require.Equal(t, float64(30), alloc.storageGB)
		require.Equal(t, "bx3d.4x20", alloc.hostFlavorID)
		require.Equal(t, int64(3), alloc.members)
	})

	t.Run("falls back to mongodbees key for enterprise-sharding-gen2", func(t *testing.T) {
		instance := map[string]interface{}{
			"dataservices": map[string]interface{}{
				"mongodbees": map[string]interface{}{
					"storage_gb":  float64(60),
					"host_flavor": "bx3d.8x40",
					"shards":      float64(2),
				},
			},
		}
		alloc := extractDatabaseAllocations(instance, "databases-for-mongodb", "enterprise-sharding-gen2")
		require.Equal(t, float64(60), alloc.storageGB)
		require.Equal(t, "bx3d.8x40", alloc.hostFlavorID)
		require.Equal(t, int64(2), alloc.shards)
	})

	t.Run("returns zero allocations when dataservices key is absent", func(t *testing.T) {
		instance := map[string]interface{}{}
		alloc := extractDatabaseAllocations(instance, "databases-for-mongodb", "standard-gen2")
		require.Equal(t, float64(0), alloc.storageGB)
		require.Equal(t, "", alloc.hostFlavorID)
	})
}
func TestSharedSetShardsInfo(t *testing.T) {
	t.Run("does not set shards when instance is nil", func(t *testing.T) {
		d := schema.TestResourceDataRaw(t, ResourceIBMDatabaseInstance().Schema, map[string]interface{}{
			"service": "databases-for-mongodb",
			"plan":    "enterprise-sharding-gen2",
		})

		err := sharedSetShardsInfo(d, nil)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		if v := d.Get("shards").(int); v != 0 {
			t.Fatalf("expected shards unset (0) for nil instance, got %d", v)
		}
	})

	t.Run("does not set shards when instance.Extensions is nil", func(t *testing.T) {
		d := schema.TestResourceDataRaw(t, ResourceIBMDatabaseInstance().Schema, map[string]interface{}{
			"service": "databases-for-mongodb",
			"plan":    "enterprise-sharding-gen2",
		})

		instance := &rc.ResourceInstance{Extensions: nil}
		err := sharedSetShardsInfo(d, instance)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		if v := d.Get("shards").(int); v != 0 {
			t.Fatalf("expected shards unset (0) for nil extensions, got %d", v)
		}
	})

	t.Run("reads shard count from extensions for enterprise-sharding-gen2", func(t *testing.T) {
		d := schema.TestResourceDataRaw(t, ResourceIBMDatabaseInstance().Schema, map[string]interface{}{
			"service": "databases-for-mongodb",
			"plan":    "enterprise-sharding-gen2",
		})

		instance := &rc.ResourceInstance{
			Extensions: map[string]interface{}{
				dataservicesKey: map[string]interface{}{
					"mongodbees": map[string]interface{}{
						"shards": float64(3),
					},
				},
			},
		}
		err := sharedSetShardsInfo(d, instance)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		if v := d.Get("shards").(int); v != 3 {
			t.Fatalf("expected shards=3 from extensions, got %d", v)
		}
	})

	t.Run("does not set shards when extensions exist but shards key is absent", func(t *testing.T) {
		d := schema.TestResourceDataRaw(t, ResourceIBMDatabaseInstance().Schema, map[string]interface{}{
			"service": "databases-for-mongodb",
			"plan":    "enterprise-sharding-gen2",
		})

		instance := &rc.ResourceInstance{
			Extensions: map[string]interface{}{
				dataservicesKey: map[string]interface{}{
					"mongodbees": map[string]interface{}{
						// no "shards" key
					},
				},
			},
		}
		err := sharedSetShardsInfo(d, instance)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		if v := d.Get("shards").(int); v != 0 {
			t.Fatalf("expected shards unset (0) when key absent, got %d", v)
		}
	})
}

func TestFlattenBackupConfiguration(t *testing.T) {
	t.Run("returns nil for nil input", func(t *testing.T) {
		require.Nil(t, flattenBackupConfiguration(nil))
	})

	t.Run("returns nil for empty map", func(t *testing.T) {
		require.Nil(t, flattenBackupConfiguration(map[string]interface{}{}))
	})

	t.Run("flattens full backup configuration", func(t *testing.T) {
		backups := map[string]interface{}{
			"preserve":  true,
			"retention": "45d",
			"automatic_backups": map[string]interface{}{
				"copy": map[string]interface{}{
					"enabled": true,
					"target": []interface{}{
						map[string]interface{}{"region": "ca-mon"},
						map[string]interface{}{"region": "us-south"},
					},
				},
			},
		}

		result := flattenBackupConfiguration(backups)
		require.Len(t, result, 1)
		require.Equal(t, true, result[0]["preserve"])
		require.Equal(t, 45, result[0]["retention"])
		require.Equal(t, []string{"ca-mon", "us-south"}, result[0]["target_regions"])
	})

	t.Run("flattens values returned by the API", func(t *testing.T) {
		backups := map[string]interface{}{
			"preserve":  false,
			"retention": "30d",
		}

		result := flattenBackupConfiguration(backups)
		require.Len(t, result, 1)
		require.Equal(t, false, result[0]["preserve"])
		require.Equal(t, 30, result[0]["retention"])
	})

	t.Run("does not invent defaults for fields the API did not return", func(t *testing.T) {
		result := flattenBackupConfiguration(map[string]interface{}{"preserve": true})
		require.Len(t, result, 1)
		require.Equal(t, map[string]interface{}{"preserve": true}, result[0])
	})

	t.Run("omits retention that cannot be parsed instead of defaulting", func(t *testing.T) {
		result := flattenBackupConfiguration(map[string]interface{}{
			"preserve":  true,
			"retention": "bogus",
		})
		require.Len(t, result, 1)
		_, hasRetention := result[0]["retention"]
		require.False(t, hasRetention)
		require.Equal(t, true, result[0]["preserve"])
	})

	t.Run("omits preserve that cannot be parsed", func(t *testing.T) {
		result := flattenBackupConfiguration(map[string]interface{}{"preserve": "maybe"})
		require.Len(t, result, 1)
		_, hasPreserve := result[0]["preserve"]
		require.False(t, hasPreserve)
	})

	t.Run("flattens preserve returned as string", func(t *testing.T) {
		// The ADR example provisioning payload uses string booleans.
		result := flattenBackupConfiguration(map[string]interface{}{"preserve": "true"})
		require.Len(t, result, 1)
		require.Equal(t, true, result[0]["preserve"])
	})

	t.Run("flattens target regions and ignores other copy fields", func(t *testing.T) {
		// The ADR metadata also carries copy.enabled and per-target account_id/resource_group_id;
		// only the region is surfaced.
		backups := map[string]interface{}{
			"automatic_backups": map[string]interface{}{
				"copy": map[string]interface{}{
					"enabled": "true",
					"target": []interface{}{map[string]interface{}{
						"region":            "ca-tor",
						"account_id":        "23b09aee04da4545b6e32805fa93249d",
						"resource_group_id": "c21a4e8564c14d1aab2a9a8b441904eb",
					}},
				},
			},
		}

		result := flattenBackupConfiguration(backups)
		require.Len(t, result, 1)
		require.Equal(t, map[string]interface{}{"target_regions": []string{"ca-tor"}}, result[0])
	})
}

func TestParseBoolValue(t *testing.T) {
	v, ok := parseBoolValue(true)
	require.True(t, ok)
	require.True(t, v)

	v, ok = parseBoolValue("false")
	require.True(t, ok)
	require.False(t, v)

	_, ok = parseBoolValue("maybe")
	require.False(t, ok)

	_, ok = parseBoolValue(nil)
	require.False(t, ok)
}

func TestParseRetentionDays(t *testing.T) {
	valid := []struct {
		in   interface{}
		want int
	}{
		{"30d", 30},
		{"45d", 45},
		{" 60 ", 60},
		{"60", 60},
		{40, 40},
		{int64(55), 55},
		{float64(50), 50},
	}
	for _, c := range valid {
		got, err := parseRetentionDays(c.in)
		require.NoError(t, err, "input %v", c.in)
		require.Equal(t, c.want, got, "input %v", c.in)
	}

	for _, in := range []interface{}{"invalid", "", "30h", nil, true, []string{"30d"}} {
		_, err := parseRetentionDays(in)
		require.Error(t, err, "input %v", in)
	}
}

// targetRegionsOf returns the sorted target_regions of the first backup_configuration block.
func targetRegionsOf(d *schema.ResourceData) []string {
	set, _ := d.Get("backup_configuration.0.target_regions").(*schema.Set)
	if set == nil {
		return nil
	}
	var out []string
	for _, r := range set.List() {
		out = append(out, r.(string))
	}
	sort.Strings(out)
	return out
}

func TestSetGen2BackupConfiguration(t *testing.T) {
	newData := func() *schema.ResourceData {
		return schema.TestResourceDataRaw(t, ResourceIBMDatabaseInstance().Schema, map[string]interface{}{
			"backup_configuration": []interface{}{
				map[string]interface{}{
					"preserve":       true,
					"retention":      45,
					"target_regions": []interface{}{"ca-mon"},
				},
			},
		})
	}
	instanceWith := func(backups interface{}) *rc.ResourceInstance {
		return &rc.ResourceInstance{
			Extensions: map[string]interface{}{
				"dataservices": map[string]interface{}{"backups": backups},
			},
		}
	}

	t.Run("API values overwrite state", func(t *testing.T) {
		d := newData()
		require.NoError(t, setGen2BackupConfiguration(d, instanceWith(map[string]interface{}{
			"preserve":  true,
			"retention": "60d",
			"automatic_backups": map[string]interface{}{
				"copy": map[string]interface{}{
					"target": []interface{}{map[string]interface{}{"region": "us-south"}},
				},
			},
		})))

		require.Equal(t, 60, d.Get("backup_configuration.0.retention"))
		require.Equal(t, []string{"us-south"}, targetRegionsOf(d))
	})

	t.Run("fields missing or unparseable in API response keep current value", func(t *testing.T) {
		d := newData()
		require.NoError(t, setGen2BackupConfiguration(d, instanceWith(map[string]interface{}{
			"retention": "bogus",
			"preserve":  true,
		})))

		require.Equal(t, 45, d.Get("backup_configuration.0.retention"))
		require.Equal(t, true, d.Get("backup_configuration.0.preserve"))
		require.Equal(t, []string{"ca-mon"}, targetRegionsOf(d))
	})

	t.Run("target regions are order-insensitive", func(t *testing.T) {
		d := newData()
		require.NoError(t, setGen2BackupConfiguration(d, instanceWith(map[string]interface{}{
			"automatic_backups": map[string]interface{}{
				"copy": map[string]interface{}{
					"target": []interface{}{
						map[string]interface{}{"region": "us-south"},
						map[string]interface{}{"region": "ca-tor"},
					},
				},
			},
		})))
		first := targetRegionsOf(d)

		d2 := newData()
		require.NoError(t, setGen2BackupConfiguration(d2, instanceWith(map[string]interface{}{
			"automatic_backups": map[string]interface{}{
				"copy": map[string]interface{}{
					"target": []interface{}{
						map[string]interface{}{"region": "ca-tor"},
						map[string]interface{}{"region": "us-south"},
					},
				},
			},
		})))

		require.Equal(t, []string{"ca-tor", "us-south"}, first)
		require.Equal(t, first, targetRegionsOf(d2))
		require.Equal(t,
			d.Get("backup_configuration.0.target_regions").(*schema.Set).Len(),
			d2.Get("backup_configuration.0.target_regions").(*schema.Set).Len())
	})

	t.Run("API with no backups block clears stale state", func(t *testing.T) {
		for name, ext := range map[string]map[string]interface{}{
			"no dataservices":  {},
			"no backups key":   {"dataservices": map[string]interface{}{}},
			"nil backups":      {"dataservices": map[string]interface{}{"backups": nil}},
			"empty backups":    {"dataservices": map[string]interface{}{"backups": map[string]interface{}{}}},
			"nil extensions":   nil,
			"wrong type value": {"dataservices": map[string]interface{}{"backups": "oops"}},
		} {
			t.Run(name, func(t *testing.T) {
				d := newData()
				require.NoError(t, setGen2BackupConfiguration(d, &rc.ResourceInstance{Extensions: ext}))
				require.Equal(t, 0, d.Get("backup_configuration.#"))
			})
		}
	})

	t.Run("nil instance is a no-op", func(t *testing.T) {
		d := newData()
		require.NoError(t, setGen2BackupConfiguration(d, nil))
		require.Equal(t, 45, d.Get("backup_configuration.0.retention"))
	})
}

func TestExpandBackupConfiguration(t *testing.T) {
	t.Run("returns nil for empty input", func(t *testing.T) {
		require.Nil(t, expandBackupConfiguration(nil))
		require.Nil(t, expandBackupConfiguration([]interface{}{}))
	})

	t.Run("accepts a schema.Set and sorts targets deterministically", func(t *testing.T) {
		cfg := []interface{}{
			map[string]interface{}{
				"target_regions": schema.NewSet(schema.HashString, []interface{}{"us-south", "ca-mon", "eu-de"}),
			},
		}

		result := expandBackupConfiguration(cfg)
		copyObj := result["automatic_backups"].(map[string]interface{})["copy"].(map[string]interface{})
		targets := copyObj["target"].([]map[string]interface{})
		require.Equal(t, []map[string]interface{}{
			{"region": "ca-mon"}, {"region": "eu-de"}, {"region": "us-south"},
		}, targets)
	})

	t.Run("expands full backup configuration", func(t *testing.T) {
		cfg := []interface{}{
			map[string]interface{}{
				"preserve":       true,
				"retention":      45,
				"target_regions": []interface{}{"ca-mon", "us-south"},
			},
		}

		result := expandBackupConfiguration(cfg)
		require.NotNil(t, result)
		require.Equal(t, true, result["preserve"])
		require.Equal(t, "45d", result["retention"])

		autoBackups, ok := result["automatic_backups"].(map[string]interface{})
		require.True(t, ok)
		copyObj, ok := autoBackups["copy"].(map[string]interface{})
		require.True(t, ok)
		targets, ok := copyObj["target"].([]map[string]interface{})
		require.True(t, ok)
		require.Len(t, targets, 2)
		require.Equal(t, "ca-mon", targets[0]["region"])
		require.Equal(t, "us-south", targets[1]["region"])
	})

	t.Run("no target regions sends no automatic_backups", func(t *testing.T) {
		cfg := []interface{}{
			map[string]interface{}{
				"preserve":       false,
				"retention":      30,
				"target_regions": []interface{}{},
			},
		}

		result := expandBackupConfiguration(cfg)
		require.Equal(t, map[string]interface{}{"preserve": false, "retention": "30d"}, result)
	})
}
