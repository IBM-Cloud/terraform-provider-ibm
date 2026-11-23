// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	testGen2WindowInstanceCRN = "crn:v1:bluemix:public:databases-for-postgresql:ca-mon:a/23b09aee04da4545b6e32805fa93249d:00000000-0000-0000-0000-000000000000::"
	testGen2WindowPath        = "/v1/backups/instances/crn:v1:bluemix:public:databases-for-postgresql:ca-mon:a%2F23b09aee04da4545b6e32805fa93249d:00000000-0000-0000-0000-000000000000::/point_in_time_recovery"
	testGen2WindowURL         = "https://api.postgresql.ca-mon.dataservices.cloud.ibm.com" + testGen2WindowPath
)

func windowURLExtensions(windowURL interface{}) map[string]interface{} {
	return map[string]interface{}{
		"dataservices": map[string]interface{}{
			"point_in_time_recovery_status": map[string]interface{}{
				"restorable_window_url": windowURL,
			},
		},
	}
}

func TestGen2RestorableWindowURL(t *testing.T) {
	tests := []struct {
		name          string
		extensions    map[string]interface{}
		expected      string
		errorContains string
	}{
		{name: "published URL is used verbatim", extensions: windowURLExtensions(testGen2WindowURL), expected: testGen2WindowURL},
		{name: "nil extensions", errorContains: "publish no dataservices.point_in_time_recovery_status.restorable_window_url"},
		{name: "no status", extensions: map[string]interface{}{"dataservices": map[string]interface{}{}}, errorContains: "refreshes its extensions"},
		{name: "empty URL", extensions: windowURLExtensions(""), errorContains: "ask IBM Cloud support to sync it"},
		{name: "URL of another type", extensions: windowURLExtensions(42), errorContains: "publish no"},
		{name: "plain http", extensions: windowURLExtensions("http://api.postgresql.ca-mon.dataservices.cloud.ibm.com" + testGen2WindowPath), errorContains: "is not an https URL under cloud.ibm.com"},
		{name: "host outside IBM Cloud", extensions: windowURLExtensions("https://example.com" + testGen2WindowPath), errorContains: "is not an https URL under cloud.ibm.com"},
		{name: "suffix without a label boundary", extensions: windowURLExtensions("https://evilcloud.ibm.com" + testGen2WindowPath), errorContains: "is not an https URL under cloud.ibm.com"},
		{name: "unparsable URL", extensions: windowURLExtensions("https://[::1"), errorContains: "is not an https URL under cloud.ibm.com"},
		{name: "IPv6 zone that ends in cloud.ibm.com", extensions: windowURLExtensions("https://[::1%25.cloud.ibm.com]" + testGen2WindowPath), errorContains: "is not an https URL under cloud.ibm.com"},
		{name: "IPv4-mapped address with a cloud.ibm.com zone", extensions: windowURLExtensions("https://[::ffff:203.0.113.7%25x.cloud.ibm.com]" + testGen2WindowPath), errorContains: "is not an https URL under cloud.ibm.com"},
		{name: "IPv4 address", extensions: windowURLExtensions("https://203.0.113.7" + testGen2WindowPath), errorContains: "is not an https URL under cloud.ibm.com"},
		{name: "explicit port is kept", extensions: windowURLExtensions("https://api.postgresql.ca-mon.dataservices.cloud.ibm.com:8443" + testGen2WindowPath), expected: "https://api.postgresql.ca-mon.dataservices.cloud.ibm.com:8443" + testGen2WindowPath},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := restorableWindowURL(testGen2WindowInstanceCRN, tt.extensions)
			if tt.errorContains != "" {
				require.Error(t, err)
				assert.Contains(t, err.Error(), tt.errorContains)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, tt.expected, got)
		})
	}
}

func newTestWindowService(t *testing.T, serverURL string) *core.BaseService {
	t.Helper()
	service, err := core.NewBaseService(&core.ServiceOptions{
		URL:           serverURL,
		Authenticator: &core.BearerTokenAuthenticator{BearerToken: "test-token"},
	})
	require.NoError(t, err)
	return service
}

func TestGen2FetchPointInTimeRecoveryWindow(t *testing.T) {
	t.Run("authenticated GET keeps the escaped instance CRN", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodGet, r.Method)
			assert.Equal(t, testGen2WindowPath, r.URL.EscapedPath())
			assert.Equal(t, "Bearer test-token", r.Header.Get("Authorization"))
			assert.Equal(t, "application/json", r.Header.Get("Accept"))

			crnSegment, err := url.QueryUnescape(strings.Split(r.URL.EscapedPath(), "/")[4])
			assert.NoError(t, err)
			assert.Equal(t, testGen2WindowInstanceCRN, crnSegment)

			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(`{
				"earliest_restorable_at": "2026-09-20T09:30:01.000Z",
				"latest_restorable_at": "2026-09-27T09:30:00.000Z",
				"retention_days": 8,
				"archiving": {"status": "healthy", "updated_at": "2026-09-27T09:29:40.000Z"},
				"unavailable_periods": [
					{"from": "2026-09-22T01:00:00.000Z", "until": "2026-09-22T03:00:01.000Z"}
				],
				"not_restorable_reason": null,
				"evaluated_at": "2026-09-27T09:30:00.000Z"
			}`))
		}))
		defer server.Close()

		window, response, err := fetchPointInTimeRecoveryWindow(context.Background(), newTestWindowService(t, server.URL), server.URL+testGen2WindowPath)
		require.NoError(t, err)
		require.NotNil(t, response)
		assert.Equal(t, http.StatusOK, response.StatusCode)

		assert.Equal(t, "2026-09-20T09:30:01.000Z", *window.EarliestRestorableAt)
		assert.Equal(t, "2026-09-27T09:30:00.000Z", *window.LatestRestorableAt)
		assert.Equal(t, int64(8), *window.RetentionDays)
		assert.Equal(t, "healthy", *window.Archiving.Status)
		require.Len(t, window.UnavailablePeriods, 1)
		assert.Equal(t, "2026-09-22T01:00:00.000Z", *window.UnavailablePeriods[0].From)
		assert.Equal(t, "2026-09-22T03:00:01.000Z", *window.UnavailablePeriods[0].Until)
		assert.Nil(t, window.NotRestorableReason)
	})

	t.Run("an error response is returned with its status", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusForbidden)
			_, _ = w.Write([]byte(`{"errors":[{"code":"forbidden","message":"You are not authorized to read this instance's point-in-time recovery window."}],"trace":"t-1","status_code":403}`))
		}))
		defer server.Close()

		window, response, err := fetchPointInTimeRecoveryWindow(context.Background(), newTestWindowService(t, server.URL), server.URL+testGen2WindowPath)
		require.Error(t, err)
		assert.Nil(t, window)
		require.NotNil(t, response)
		assert.Equal(t, http.StatusForbidden, response.StatusCode)
		assert.Contains(t, err.Error(), "You are not authorized")
	})
}

func TestGen2FlattenPointInTimeRecoveryWindow(t *testing.T) {
	earliest := "2026-09-20T09:30:01.000Z"
	latest := "2026-09-27T09:30:00.000Z"
	beforeGap := "2026-09-22T00:59:59.999Z"
	from := "2026-09-22T01:00:00.000Z"
	until := "2026-09-22T03:00:01.000Z"
	healthy := "healthy"
	delayed := "delayed"
	archiveDelayed := "archive_delayed"
	archiveGap := "archive_gap"
	noEligibleBackup := "no_eligible_backup"
	retentionDays := int64(14)

	tests := []struct {
		name     string
		window   pointInTimeRecoveryWindow
		expected map[string]interface{}
	}{
		{
			name: "restorable with a closed period",
			window: pointInTimeRecoveryWindow{
				EarliestRestorableAt: &earliest,
				LatestRestorableAt:   &latest,
				RetentionDays:        &retentionDays,
				Archiving:            &pointInTimeRecoveryArchiving{Status: &healthy},
				UnavailablePeriods:   []pointInTimeRecoveryUnavailablePeriod{{From: &from, Until: &until}},
			},
			expected: map[string]interface{}{
				"earliest_point_in_time_recovery_time": earliest,
				"latest_point_in_time_recovery_time":   latest,
				"retention_days":                       14,
				"archiving_status":                     healthy,
				"unavailable_periods":                  []map[string]interface{}{{"from": from, "until": until}},
				"not_restorable_reason":                "",
			},
		},
		{
			name: "delayed archiving: no times and archive_delayed",
			window: pointInTimeRecoveryWindow{
				RetentionDays:       &retentionDays,
				Archiving:           &pointInTimeRecoveryArchiving{Status: &delayed},
				NotRestorableReason: &archiveDelayed,
			},
			expected: map[string]interface{}{
				"earliest_point_in_time_recovery_time": "",
				"latest_point_in_time_recovery_time":   "",
				"retention_days":                       14,
				"archiving_status":                     delayed,
				"unavailable_periods":                  []map[string]interface{}{},
				"not_restorable_reason":                archiveDelayed,
			},
		},
		{
			name: "open period after a gap holds latest before it",
			window: pointInTimeRecoveryWindow{
				EarliestRestorableAt: &earliest,
				LatestRestorableAt:   &beforeGap,
				RetentionDays:        &retentionDays,
				Archiving:            &pointInTimeRecoveryArchiving{Status: &healthy},
				UnavailablePeriods:   []pointInTimeRecoveryUnavailablePeriod{{From: &from}},
			},
			expected: map[string]interface{}{
				"earliest_point_in_time_recovery_time": earliest,
				"latest_point_in_time_recovery_time":   beforeGap,
				"retention_days":                       14,
				"archiving_status":                     healthy,
				"unavailable_periods":                  []map[string]interface{}{{"from": from, "until": ""}},
				"not_restorable_reason":                "",
			},
		},
		{
			name: "no backup to start from yet",
			window: pointInTimeRecoveryWindow{
				RetentionDays:       &retentionDays,
				Archiving:           &pointInTimeRecoveryArchiving{Status: &healthy},
				NotRestorableReason: &noEligibleBackup,
			},
			expected: map[string]interface{}{
				"earliest_point_in_time_recovery_time": "",
				"latest_point_in_time_recovery_time":   "",
				"retention_days":                       14,
				"archiving_status":                     healthy,
				"unavailable_periods":                  []map[string]interface{}{},
				"not_restorable_reason":                noEligibleBackup,
			},
		},
		{
			name: "an open gap that covers every time",
			window: pointInTimeRecoveryWindow{
				RetentionDays:       &retentionDays,
				Archiving:           &pointInTimeRecoveryArchiving{Status: &healthy},
				UnavailablePeriods:  []pointInTimeRecoveryUnavailablePeriod{{From: &from}},
				NotRestorableReason: &archiveGap,
			},
			expected: map[string]interface{}{
				"earliest_point_in_time_recovery_time": "",
				"latest_point_in_time_recovery_time":   "",
				"retention_days":                       14,
				"archiving_status":                     healthy,
				"unavailable_periods":                  []map[string]interface{}{{"from": from, "until": ""}},
				"not_restorable_reason":                archiveGap,
			},
		},
		{
			name:   "empty response",
			window: pointInTimeRecoveryWindow{},
			expected: map[string]interface{}{
				"earliest_point_in_time_recovery_time": "",
				"latest_point_in_time_recovery_time":   "",
				"retention_days":                       0,
				"archiving_status":                     "",
				"unavailable_periods":                  []map[string]interface{}{},
				"not_restorable_reason":                "",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, flattenPointInTimeRecoveryWindow(&tt.window))
		})
	}
}

func TestGen2PointInTimeRecoveryWindowSchema(t *testing.T) {
	dataSourceSchema := DataSourceIBMDatabasePointInTimeRecovery().Schema

	for _, key := range []string{
		"earliest_point_in_time_recovery_time",
		"latest_point_in_time_recovery_time",
		"retention_days",
		"archiving_status",
		"unavailable_periods",
		"not_restorable_reason",
	} {
		require.Contains(t, dataSourceSchema, key)
		assert.True(t, dataSourceSchema[key].Computed, "%s must be computed", key)
	}

	earliest := "2026-09-20T09:30:01.000Z"
	beforeGap := "2026-09-22T00:59:59.999Z"
	from := "2026-09-22T01:00:00.000Z"
	retentionDays := int64(8)
	window := &pointInTimeRecoveryWindow{
		EarliestRestorableAt: &earliest,
		LatestRestorableAt:   &beforeGap,
		RetentionDays:        &retentionDays,
		UnavailablePeriods:   []pointInTimeRecoveryUnavailablePeriod{{From: &from}},
	}
	d := schema.TestResourceDataRaw(t, dataSourceSchema, map[string]interface{}{"deployment_id": testGen2WindowInstanceCRN})
	for key, value := range flattenPointInTimeRecoveryWindow(window) {
		require.NoError(t, d.Set(key, value), "setting %s", key)
	}
	assert.Equal(t, earliest, d.Get("earliest_point_in_time_recovery_time"))
	assert.Equal(t, beforeGap, d.Get("latest_point_in_time_recovery_time"))
	assert.Equal(t, 8, d.Get("retention_days"))
	assert.Equal(t, from, d.Get("unavailable_periods.0.from"))
	assert.Equal(t, "", d.Get("unavailable_periods.0.until"))
}
