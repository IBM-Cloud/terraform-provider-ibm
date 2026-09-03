// Copyright IBM Corp. 2017, 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"fmt"
	"testing"
	"time"

	"github.com/IBM/cloud-databases-go-sdk/clouddatabasesv5"
	"github.com/IBM/go-sdk-core/v5/core"
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
