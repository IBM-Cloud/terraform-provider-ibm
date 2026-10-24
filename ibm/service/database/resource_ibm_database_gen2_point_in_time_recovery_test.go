// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"sync/atomic"
	"testing"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	rc "github.com/IBM/platform-services-go-sdk/resourcecontrollerv2"
	"github.com/hashicorp/go-cty/cty"
	sdkretry "github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	testGen2PointInTimeSourceCRN = "crn:v1:bluemix:public:databases-for-postgresql:us-south:a/abc123:source-id::"
	testGen2PointInTime          = "2026-09-27T09:30:00Z"
)

func TestGen2ValidatePointInTimeRestore(t *testing.T) {
	valid := gen2PointInTimeRestore{
		sourceCRN:        testGen2PointInTimeSourceCRN,
		pointInTime:      testGen2PointInTime,
		service:          "databases-for-postgresql",
		location:         "us-south",
		sourceSet:        true,
		sourceKnown:      true,
		pointInTimeSet:   true,
		pointInTimeKnown: true,
	}
	with := func(change func(*gen2PointInTimeRestore)) gen2PointInTimeRestore {
		restore := valid
		change(&restore)
		return restore
	}

	tests := []struct {
		name          string
		restore       gen2PointInTimeRestore
		errorContains string
	}{
		{name: "neither argument set", restore: gen2PointInTimeRestore{service: "databases-for-postgresql", location: "us-south"}},
		{name: "both arguments valid", restore: valid},
		{name: "time with an offset", restore: with(func(r *gen2PointInTimeRestore) { r.pointInTime = "2026-09-27T11:30:00+02:00" })},
		{name: "fractional seconds", restore: with(func(r *gen2PointInTimeRestore) { r.pointInTime = "2026-09-27T09:30:00.250Z" })},
		{
			name: "only the source",
			restore: with(func(r *gen2PointInTimeRestore) {
				r.pointInTime, r.pointInTimeSet, r.pointInTimeKnown = "", false, false
			}),
			errorContains: "must be set together",
		},
		{
			name:          "only the time",
			restore:       with(func(r *gen2PointInTimeRestore) { r.sourceCRN, r.sourceSet, r.sourceKnown = "", false, false }),
			errorContains: "must be set together",
		},
		{
			name:          "empty time",
			restore:       with(func(r *gen2PointInTimeRestore) { r.pointInTime = "" }),
			errorContains: "no restore to the latest time",
		},
		{
			name:          "blank time",
			restore:       with(func(r *gen2PointInTimeRestore) { r.pointInTime = "  " }),
			errorContains: "no restore to the latest time",
		},
		{
			name:          "time without the T separator",
			restore:       with(func(r *gen2PointInTimeRestore) { r.pointInTime = "2026-09-27 09:30:00" }),
			errorContains: "is not an RFC 3339 timestamp",
		},
		{
			name:          "time without an offset",
			restore:       with(func(r *gen2PointInTimeRestore) { r.pointInTime = "2026-09-27T09:30:00" }),
			errorContains: "is not an RFC 3339 timestamp",
		},
		{
			name:          "time with surrounding whitespace",
			restore:       with(func(r *gen2PointInTimeRestore) { r.pointInTime = " " + testGen2PointInTime }),
			errorContains: "is not an RFC 3339 timestamp",
		},
		{
			name:          "combined with backup_id",
			restore:       with(func(r *gen2PointInTimeRestore) { r.backupIDSet = true }),
			errorContains: "backup_id cannot be combined",
		},
		{
			name: "backup CRN as the source",
			restore: with(func(r *gen2PointInTimeRestore) {
				r.sourceCRN = "crn:v1:bluemix:public:databases-for-postgresql:us-south:a/abc123:source-id:backup:backup-id"
			}),
			errorContains: "to restore a backup, use backup_id",
		},
		{
			name: "independent backup CRN as the source",
			restore: with(func(r *gen2PointInTimeRestore) {
				r.sourceCRN = "crn:v1:bluemix:public:databases-independent-backups:us-south:a/abc123:backup-id::"
			}),
			errorContains: "must be an instance of databases-for-postgresql",
		},
		{
			name: "truncated CRN",
			restore: with(func(r *gen2PointInTimeRestore) {
				r.sourceCRN = "crn:v1:bluemix:public:databases-for-postgresql:us-south:a/abc123:source-id"
			}),
			errorContains: "must be the CRN of the source instance",
		},
		{
			name:          "instance ID instead of a CRN",
			restore:       with(func(r *gen2PointInTimeRestore) { r.sourceCRN = "source-id" }),
			errorContains: "must be the CRN of the source instance",
		},
		{
			name:          "source of another service",
			restore:       with(func(r *gen2PointInTimeRestore) { r.service = "databases-for-mysql" }),
			errorContains: "must be an instance of databases-for-mysql",
		},
		{
			name:          "source in another region",
			restore:       with(func(r *gen2PointInTimeRestore) { r.location = "ca-mon" }),
			errorContains: "restores into the source's region only",
		},
		{
			name:    "unknown source skips the CRN checks",
			restore: with(func(r *gen2PointInTimeRestore) { r.sourceCRN, r.sourceKnown = "", false }),
		},
		{
			name:    "unknown time skips the format check",
			restore: with(func(r *gen2PointInTimeRestore) { r.pointInTime, r.pointInTimeKnown = "", false }),
		},
		{
			name:    "unknown service and location skip the comparisons",
			restore: with(func(r *gen2PointInTimeRestore) { r.service, r.location = "", "" }),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateGen2PointInTimeRestore(tt.restore)
			if tt.errorContains == "" {
				assert.NoError(t, err)
				return
			}
			require.Error(t, err)
			assert.Contains(t, err.Error(), tt.errorContains)
		})
	}
}

func testGen2PointInTimeRawConfig(values map[string]cty.Value) cty.Value {
	attrs := map[string]cty.Value{
		pointInTimeRecoveryDeploymentIDKey: cty.NullVal(cty.String),
		pointInTimeRecoveryTimeKey:         cty.NullVal(cty.String),
		"backup_id":                        cty.NullVal(cty.String),
	}
	for k, v := range values {
		attrs[k] = v
	}
	return cty.ObjectVal(attrs)
}

func TestGen2PointInTimeRestoreFromRawConfig(t *testing.T) {
	tests := []struct {
		name     string
		raw      cty.Value
		expected gen2PointInTimeRestore
	}{
		{name: "null configuration", raw: cty.NullVal(cty.EmptyObject)},
		{name: "unknown configuration", raw: cty.UnknownVal(cty.EmptyObject)},
		{name: "arguments absent", raw: cty.EmptyObjectVal},
		{name: "arguments null", raw: testGen2PointInTimeRawConfig(nil)},
		{
			name: "empty time is set",
			raw: testGen2PointInTimeRawConfig(map[string]cty.Value{
				pointInTimeRecoveryTimeKey: cty.StringVal(""),
			}),
			expected: gen2PointInTimeRestore{pointInTimeSet: true, pointInTimeKnown: true},
		},
		{
			name: "unknown values are set but not known",
			raw: testGen2PointInTimeRawConfig(map[string]cty.Value{
				pointInTimeRecoveryDeploymentIDKey: cty.UnknownVal(cty.String),
				pointInTimeRecoveryTimeKey:         cty.UnknownVal(cty.String),
			}),
			expected: gen2PointInTimeRestore{sourceSet: true, pointInTimeSet: true},
		},
		{
			name: "values",
			raw: testGen2PointInTimeRawConfig(map[string]cty.Value{
				pointInTimeRecoveryDeploymentIDKey: cty.StringVal(testGen2PointInTimeSourceCRN),
				pointInTimeRecoveryTimeKey:         cty.StringVal(testGen2PointInTime),
			}),
			expected: gen2PointInTimeRestore{
				sourceCRN:        testGen2PointInTimeSourceCRN,
				pointInTime:      testGen2PointInTime,
				sourceSet:        true,
				sourceKnown:      true,
				pointInTimeSet:   true,
				pointInTimeKnown: true,
			},
		},
		{
			name:     "backup_id set",
			raw:      testGen2PointInTimeRawConfig(map[string]cty.Value{"backup_id": cty.StringVal("crn:v1:backup")}),
			expected: gen2PointInTimeRestore{backupIDSet: true},
		},
		{
			name:     "backup_id unknown",
			raw:      testGen2PointInTimeRawConfig(map[string]cty.Value{"backup_id": cty.UnknownVal(cty.String)}),
			expected: gen2PointInTimeRestore{backupIDSet: true},
		},
		{
			name: "empty backup_id is unset",
			raw:  testGen2PointInTimeRawConfig(map[string]cty.Value{"backup_id": cty.StringVal("")}),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, pointInTimeRestoreFromRawConfig(tt.raw))
		})
	}
}

// planGen2PointInTimeRestore runs only the point-in-time recovery CustomizeDiff, as a plan would.
func planGen2PointInTimeRestore(t *testing.T, instanceID string, config map[string]interface{}) error {
	t.Helper()

	raw := map[string]interface{}{
		"name":     "test-gen2-db",
		"location": "us-south",
		"service":  "databases-for-postgresql",
		"plan":     "standard-gen2",
	}
	for k, v := range config {
		raw[k] = v
	}
	rawConfig := map[string]cty.Value{}
	stateAttributes := map[string]string{}
	for k, v := range raw {
		rawConfig[k] = cty.StringVal(v.(string))
		if instanceID != "" {
			stateAttributes[k] = v.(string)
		}
	}

	res := &schema.Resource{
		Schema:        ResourceIBMDatabaseInstance().Schema,
		CustomizeDiff: validateBackendSpecificPointInTimeRecoveryDiff,
	}
	// An existing instance's state matches its configuration, so no ForceNew diff turns the plan into a create.
	state := &terraform.InstanceState{ID: instanceID, Attributes: stateAttributes, RawConfig: testGen2PointInTimeRawConfig(rawConfig)}
	_, err := res.Diff(context.Background(), state, terraform.NewResourceConfigRaw(raw), nil)
	return err
}

func TestGen2ValidatePointInTimeRecoveryDiff(t *testing.T) {
	tests := []struct {
		name          string
		instanceID    string
		config        map[string]interface{}
		errorContains string
	}{
		{name: "no restore", config: map[string]interface{}{}},
		{
			name: "valid restore",
			config: map[string]interface{}{
				pointInTimeRecoveryDeploymentIDKey: testGen2PointInTimeSourceCRN,
				pointInTimeRecoveryTimeKey:         testGen2PointInTime,
			},
		},
		{
			name: "empty time",
			config: map[string]interface{}{
				pointInTimeRecoveryDeploymentIDKey: testGen2PointInTimeSourceCRN,
				pointInTimeRecoveryTimeKey:         "",
			},
			errorContains: "no restore to the latest time",
		},
		{
			name: "source in another region",
			config: map[string]interface{}{
				"location":                         "ca-mon",
				pointInTimeRecoveryDeploymentIDKey: testGen2PointInTimeSourceCRN,
				pointInTimeRecoveryTimeKey:         testGen2PointInTime,
			},
			errorContains: "restores into the source's region only",
		},
		{
			name: "combined with backup_id",
			config: map[string]interface{}{
				"backup_id":                        "crn:v1:bluemix:public:databases-independent-backups:us-south:a/abc123:backup-id::",
				pointInTimeRecoveryDeploymentIDKey: testGen2PointInTimeSourceCRN,
				pointInTimeRecoveryTimeKey:         testGen2PointInTime,
			},
			errorContains: "backup_id cannot be combined",
		},
		{
			name:       "existing instance is not validated",
			instanceID: "crn:v1:bluemix:public:databases-for-postgresql:us-south:a/abc123:target-id::",
			config: map[string]interface{}{
				pointInTimeRecoveryDeploymentIDKey: testGen2PointInTimeSourceCRN,
				pointInTimeRecoveryTimeKey:         "",
			},
		},
		{
			name: "classic plan keeps restore to latest",
			config: map[string]interface{}{
				"plan":                             "standard",
				pointInTimeRecoveryDeploymentIDKey: testGen2PointInTimeSourceCRN,
				pointInTimeRecoveryTimeKey:         "",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := planGen2PointInTimeRestore(t, tt.instanceID, tt.config)
			if tt.errorContains == "" {
				assert.NoError(t, err)
				return
			}
			require.Error(t, err)
			assert.Contains(t, err.Error(), tt.errorContains)
		})
	}
}

func TestGen2CanonicalPointInTime(t *testing.T) {
	tests := []struct {
		input    string
		expected string
		wantErr  bool
	}{
		{input: "2026-09-27T09:30:00Z", expected: "2026-09-27T09:30:00Z"},
		{input: "2026-09-27T11:30:00+02:00", expected: "2026-09-27T09:30:00Z"},
		{input: "2026-09-27T05:30:00-04:00", expected: "2026-09-27T09:30:00Z"},
		{input: "2026-09-27T09:30:00.250Z", expected: "2026-09-27T09:30:00.25Z"},
		{input: "2026-09-27T09:30:00.000Z", expected: "2026-09-27T09:30:00Z"},
		{input: "", wantErr: true},
		{input: "2026-09-27", wantErr: true},
		{input: "latest", wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got, err := canonicalPointInTime(tt.input)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, tt.expected, got)
		})
	}
}

func testGen2DataservicesParameters(t *testing.T, parameters map[string]interface{}) map[string]interface{} {
	t.Helper()
	dataservices, ok := parameters["dataservices"].(map[string]interface{})
	require.True(t, ok, "parameters must carry a dataservices object")
	return dataservices
}

func testGen2MemberGroup(members int) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"group_id": "member",
			"members": []interface{}{
				map[string]interface{}{"allocation_count": members},
			},
		},
	}
}

func testGen2Retention(retentionDays int) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"point_in_time_recovery": []interface{}{
				map[string]interface{}{"retention_days": retentionDays},
			},
		},
	}
}

func TestGen2BuildCreateParametersPointInTimeRestore(t *testing.T) {
	g := &resourceIBMDatabaseGen2Backend{}

	t.Run("restore inherits the source's members", func(t *testing.T) {
		d := testGen2DatabaseResourceData(t, map[string]interface{}{
			pointInTimeRecoveryDeploymentIDKey: testGen2PointInTimeSourceCRN,
			pointInTimeRecoveryTimeKey:         "2026-09-27T11:30:00+02:00",
		})

		parameters, err := g.buildGen2CreateParameters(d, "databases-for-postgresql", nil, "")
		require.NoError(t, err)

		dataservices := testGen2DataservicesParameters(t, parameters)
		assert.Equal(t, testGen2PointInTimeSourceCRN, dataservices["source_dataservice_crn"])
		assert.Equal(t, testGen2PointInTime, dataservices["point_in_time"])
		assert.NotContains(t, dataservices, "restore_backup_id")
		assert.NotContains(t, dataservices, "backups")
		assert.Equal(t, map[string]interface{}{}, dataservices["postgresql"])
	})

	t.Run("configured members are sent", func(t *testing.T) {
		d := testGen2DatabaseResourceData(t, map[string]interface{}{
			pointInTimeRecoveryDeploymentIDKey: testGen2PointInTimeSourceCRN,
			pointInTimeRecoveryTimeKey:         testGen2PointInTime,
			"group":                            testGen2MemberGroup(3),
		})

		parameters, err := g.buildGen2CreateParameters(d, "databases-for-postgresql", nil, "")
		require.NoError(t, err)

		dbConfig := testGen2DataservicesParameters(t, parameters)["postgresql"].(map[string]interface{})
		assert.Equal(t, 3, dbConfig["members"])
	})

	t.Run("restore with retention", func(t *testing.T) {
		d := testGen2DatabaseResourceData(t, map[string]interface{}{
			pointInTimeRecoveryDeploymentIDKey: testGen2PointInTimeSourceCRN,
			pointInTimeRecoveryTimeKey:         testGen2PointInTime,
			"backups":                          testGen2Retention(14),
		})

		parameters, err := g.buildGen2CreateParameters(d, "databases-for-postgresql", nil, "")
		require.NoError(t, err)

		assert.Equal(t, pointInTimeRecoveryRetention(14), testGen2DataservicesParameters(t, parameters)["backups"])
	})

	t.Run("instance without a restore is unchanged", func(t *testing.T) {
		d := testGen2DatabaseResourceData(t, map[string]interface{}{
			"group": testGen2MemberGroup(2),
		})

		parameters, err := g.buildGen2CreateParameters(d, "databases-for-postgresql", nil, "")
		require.NoError(t, err)

		assert.Equal(t, map[string]interface{}{
			"dataservices": map[string]interface{}{
				"postgresql": map[string]interface{}{"members": 2},
			},
		}, parameters)
	})

	applyTimeErrors := []struct {
		name          string
		config        map[string]interface{}
		errorContains string
	}{
		{
			name: "empty time resolved at apply",
			config: map[string]interface{}{
				pointInTimeRecoveryDeploymentIDKey: testGen2PointInTimeSourceCRN,
				pointInTimeRecoveryTimeKey:         "",
			},
			errorContains: "no restore to the latest time",
		},
		{
			name: "backup CRN resolved at apply",
			config: map[string]interface{}{
				pointInTimeRecoveryDeploymentIDKey: "crn:v1:bluemix:public:databases-for-postgresql:us-south:a/abc123:source-id:backup:backup-id",
				pointInTimeRecoveryTimeKey:         testGen2PointInTime,
			},
			errorContains: "use backup_id",
		},
		{
			name: "combined with backup_id",
			config: map[string]interface{}{
				"backup_id":                        "crn:v1:bluemix:public:databases-independent-backups:us-south:a/abc123:backup-id::",
				pointInTimeRecoveryDeploymentIDKey: testGen2PointInTimeSourceCRN,
				pointInTimeRecoveryTimeKey:         testGen2PointInTime,
			},
			errorContains: "backup_id cannot be combined",
		},
	}
	for _, tt := range applyTimeErrors {
		t.Run(tt.name, func(t *testing.T) {
			d := testGen2DatabaseResourceData(t, tt.config)

			_, err := g.buildGen2CreateParameters(d, "databases-for-postgresql", nil, "")
			require.Error(t, err)
			assert.Contains(t, err.Error(), tt.errorContains)
		})
	}
}

func TestGen2BuildParametersExcludesPointInTimeRestore(t *testing.T) {
	g := &resourceIBMDatabaseGen2Backend{}
	d := testGen2DatabaseResourceData(t, map[string]interface{}{
		pointInTimeRecoveryDeploymentIDKey: testGen2PointInTimeSourceCRN,
		pointInTimeRecoveryTimeKey:         testGen2PointInTime,
	})

	parameters, err := g.buildGen2Parameters(d, "databases-for-postgresql", nil, "")
	require.NoError(t, err)

	dataservices := testGen2DataservicesParameters(t, parameters)
	assert.NotContains(t, dataservices, "source_dataservice_crn")
	assert.NotContains(t, dataservices, "point_in_time")
	assert.NotContains(t, dataservices["postgresql"], "members")
}

func TestGen2AddPointInTimeRecoveryRetention(t *testing.T) {
	t.Run("configured", func(t *testing.T) {
		d := testGen2DatabaseResourceData(t, map[string]interface{}{"backups": testGen2Retention(14)})
		dataservices := map[string]interface{}{}

		addPointInTimeRecoveryRetention(d, dataservices)

		assert.Equal(t, map[string]interface{}{
			"backups": map[string]interface{}{
				"point_in_time_recovery": map[string]interface{}{"retention_days": 14},
			},
		}, dataservices)
	})

	t.Run("absent", func(t *testing.T) {
		d := testGen2DatabaseResourceData(t, nil)
		dataservices := map[string]interface{}{}

		addPointInTimeRecoveryRetention(d, dataservices)

		assert.Empty(t, dataservices)
	})
}

func TestGen2BuildParametersExcludesRetention(t *testing.T) {
	g := &resourceIBMDatabaseGen2Backend{}
	d := testGen2DatabaseResourceData(t, map[string]interface{}{
		"group":   testGen2MemberGroup(2),
		"backups": testGen2Retention(14),
	})

	parameters, err := g.buildGen2Parameters(d, "databases-for-postgresql", nil, "")
	require.NoError(t, err)

	assert.NotContains(t, testGen2DataservicesParameters(t, parameters), "backups")
}

func TestGen2PointInTimeRecoveryRetentionUpdate(t *testing.T) {
	t.Run("changed retention is sent alone", func(t *testing.T) {
		d := testGen2DatabaseResourceData(t, map[string]interface{}{
			"group":   testGen2MemberGroup(2),
			"backups": testGen2Retention(21),
		})

		parameters, ok := pointInTimeRecoveryRetentionUpdate(d)
		require.True(t, ok)
		assert.Equal(t, map[string]interface{}{
			"dataservices": map[string]interface{}{
				"backups": map[string]interface{}{
					"point_in_time_recovery": map[string]interface{}{"retention_days": 21},
				},
			},
		}, parameters)
	})

	t.Run("no retention, no update", func(t *testing.T) {
		d := testGen2DatabaseResourceData(t, map[string]interface{}{"group": testGen2MemberGroup(2)})

		_, ok := pointInTimeRecoveryRetentionUpdate(d)
		assert.False(t, ok)
		assert.Nil(t, (&resourceIBMDatabaseGen2Backend{}).applyPointInTimeRecoveryRetentionWithDiagnostics(context.Background(), d, nil, ""))
	})
}

const testGen2InstanceOperationID = "crn:v1:bluemix:public:databases-for-postgresql:us-south:a/abc123:target-id::"

func testGen2ResourceInstanceJSON(state, lastOperationState string, async bool) string {
	return fmt.Sprintf(`{"id": %q, "state": %q, "last_operation": {"type": "update", "state": %q, "async": %t, "description": "retention_days exceeds the regional maximum", "cancelable": false, "poll": true}}`,
		testGen2InstanceOperationID, state, lastOperationState, async)
}

func TestGen2InstanceOperationStateChangeConf(t *testing.T) {
	inProgress := testGen2ResourceInstanceJSON("active", "in progress", true)

	tests := []struct {
		name          string
		responses     []string
		expectedPolls int
		timesOut      bool
		errorContains string
	}{
		{
			name:      "active instance with an asynchronous update in progress keeps waiting",
			responses: []string{inProgress},
			timesOut:  true,
		},
		{
			name:          "asynchronous update succeeds",
			responses:     []string{inProgress, inProgress, testGen2ResourceInstanceJSON("active", "succeeded", true)},
			expectedPolls: 3,
		},
		{
			name:          "asynchronous update fails",
			responses:     []string{inProgress, testGen2ResourceInstanceJSON("active", "failed", true)},
			expectedPolls: 2,
			errorContains: "retention_days exceeds the regional maximum",
		},
		{
			name:          "synchronous update follows the instance state",
			responses:     []string{testGen2ResourceInstanceJSON("active", "succeeded", false)},
			expectedPolls: 1,
		},
		{
			name:          "failed instance",
			responses:     []string{testGen2ResourceInstanceJSON("failed", "succeeded", false)},
			expectedPolls: 1,
			errorContains: "is in the failed state",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var polls int32
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, "/v2/resource_instances/"+url.PathEscape(testGen2InstanceOperationID), r.URL.EscapedPath())

				poll := int(atomic.AddInt32(&polls, 1))
				w.Header().Set("Content-Type", "application/json")
				_, err := fmt.Fprint(w, tt.responses[min(poll, len(tt.responses))-1])
				assert.NoError(t, err)
			}))
			defer server.Close()

			rsConClient, err := rc.NewResourceControllerV2(&rc.ResourceControllerV2Options{
				URL:           server.URL,
				Authenticator: &core.NoAuthAuthenticator{},
			})
			require.NoError(t, err)

			stateConf := gen2InstanceOperationStateChangeConf(rsConClient, testGen2InstanceOperationID, 500*time.Millisecond)
			stateConf.Delay = 0
			stateConf.PollInterval = 10 * time.Millisecond

			_, err = stateConf.WaitForStateContext(context.Background())

			if tt.timesOut {
				var timeoutErr *sdkretry.TimeoutError
				require.ErrorAs(t, err, &timeoutErr)
				assert.Equal(t, "in progress", timeoutErr.LastState)
				assert.Greater(t, int(atomic.LoadInt32(&polls)), 1, "the wait must keep polling")
				return
			}
			assert.Equal(t, tt.expectedPolls, int(atomic.LoadInt32(&polls)))
			if tt.errorContains == "" {
				assert.NoError(t, err)
				return
			}
			require.Error(t, err)
			assert.Contains(t, err.Error(), tt.errorContains)
		})
	}
}

func TestGen2FlattenPointInTimeRecoveryRetention(t *testing.T) {
	tests := []struct {
		name       string
		extensions map[string]interface{}
		expected   []map[string]interface{}
	}{
		{name: "nil extensions"},
		{name: "no dataservices", extensions: map[string]interface{}{}},
		{name: "no backups", extensions: map[string]interface{}{"dataservices": map[string]interface{}{}}},
		{
			name: "backups without point-in-time recovery",
			extensions: map[string]interface{}{"dataservices": map[string]interface{}{
				"backups": map[string]interface{}{},
			}},
		},
		{
			name: "point-in-time recovery without retention",
			extensions: map[string]interface{}{"dataservices": map[string]interface{}{
				"backups": map[string]interface{}{"point_in_time_recovery": map[string]interface{}{}},
			}},
		},
		{
			name: "retention",
			extensions: map[string]interface{}{"dataservices": map[string]interface{}{
				"backups": map[string]interface{}{"point_in_time_recovery": map[string]interface{}{"retention_days": float64(14)}},
			}},
			expected: []map[string]interface{}{
				{"point_in_time_recovery": []map[string]interface{}{{"retention_days": 14}}},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, flattenPointInTimeRecoveryRetention(tt.extensions))
		})
	}
}

func TestGen2PointInTimeRecoveryRetentionSchema(t *testing.T) {
	backups := ResourceIBMDatabaseInstance().Schema["backups"]
	require.NotNil(t, backups)
	assert.True(t, backups.Optional)
	assert.True(t, backups.Computed)
	assert.Equal(t, 1, backups.MaxItems)

	pointInTimeRecovery := backups.Elem.(*schema.Resource).Schema["point_in_time_recovery"]
	assert.True(t, pointInTimeRecovery.Optional)
	assert.True(t, pointInTimeRecovery.Computed)
	assert.Equal(t, 1, pointInTimeRecovery.MaxItems)

	retentionDays := pointInTimeRecovery.Elem.(*schema.Resource).Schema["retention_days"]
	_, errs := retentionDays.ValidateFunc(6, "retention_days")
	assert.NotEmpty(t, errs)
	_, errs = retentionDays.ValidateFunc(7, "retention_days")
	assert.Empty(t, errs)
	_, errs = retentionDays.ValidateFunc(90, "retention_days")
	assert.Empty(t, errs, "the regional maximum is enforced by the platform")
}

// planWithUnsupportedAttrsValidation runs only the unsupported-attribute CustomizeDiff, as a plan would.
func planWithUnsupportedAttrsValidation(t *testing.T, plan string, config map[string]interface{}) error {
	t.Helper()

	raw := map[string]interface{}{
		"name":     "test-db",
		"location": "us-south",
		"service":  "databases-for-postgresql",
		"plan":     plan,
	}
	for k, v := range config {
		raw[k] = v
	}

	res := &schema.Resource{
		Schema:        ResourceIBMDatabaseInstance().Schema,
		CustomizeDiff: validateUnsupportedAttrsDiff,
	}
	_, err := res.Diff(context.Background(), nil, terraform.NewResourceConfigRaw(raw), nil)
	return err
}

func TestGen2PointInTimeRecoveryUnsupportedAttrs(t *testing.T) {
	restore := map[string]interface{}{
		pointInTimeRecoveryDeploymentIDKey: testGen2PointInTimeSourceCRN,
		pointInTimeRecoveryTimeKey:         testGen2PointInTime,
		"backups":                          testGen2Retention(14),
	}

	t.Run("accepted on a Gen2 plan", func(t *testing.T) {
		assert.NoError(t, planWithUnsupportedAttrsValidation(t, "standard-gen2", restore))
	})

	t.Run("classic plan rejects backups", func(t *testing.T) {
		err := planWithUnsupportedAttrsValidation(t, "standard", map[string]interface{}{"backups": testGen2Retention(14)})
		require.Error(t, err)
		assert.Contains(t, err.Error(), "not supported for Classic databases")
		assert.Contains(t, err.Error(), `"backups"`)
	})

	t.Run("classic plan keeps its restore arguments", func(t *testing.T) {
		assert.NoError(t, planWithUnsupportedAttrsValidation(t, "standard", map[string]interface{}{
			pointInTimeRecoveryDeploymentIDKey: testGen2PointInTimeSourceCRN,
			pointInTimeRecoveryTimeKey:         "",
		}))
	})
}
