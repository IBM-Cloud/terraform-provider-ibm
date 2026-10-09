// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM/cloud-databases-go-sdk/clouddatabasesv5"
	"github.com/IBM/go-sdk-core/v5/core"
	rc "github.com/IBM/platform-services-go-sdk/resourcecontrollerv2"
)

const (
	testPointInTimeRecoveryDataSourceResource = "(Data) ibm_database_point_in_time_recovery"
	testWindowServerErrorBody                 = `{"errors":[{"code":"internal_server_error","message":"The restorable window could not be read. Retry the request later. If the problem persists, open a support case and include the trace."}],"trace":"trace-123","status_code":500}`
	testWindowServerErrorSummary              = "HTTP 500: The restorable window could not be read. Retry the request later. If the problem persists, open a support case and include the trace. (trace: trace-123)"
)

// pointInTimeRecoveryTestSession panics on any client the data source is not expected to use.
type pointInTimeRecoveryTestSession struct {
	conns.ClientSession
	resourceController    *rc.ResourceControllerV2
	resourceControllerErr error
	cloudDatabases        *clouddatabasesv5.CloudDatabasesV5
	cloudDatabasesErr     error
}

func (s pointInTimeRecoveryTestSession) ResourceControllerV2API() (*rc.ResourceControllerV2, error) {
	return s.resourceController, s.resourceControllerErr
}

func (s pointInTimeRecoveryTestSession) CloudDatabasesV5() (*clouddatabasesv5.CloudDatabasesV5, error) {
	return s.cloudDatabases, s.cloudDatabasesErr
}

// newCloudTestSession routes every request to server whatever host its URL names,
// so the window URL can stay under cloud.ibm.com as restorableWindowURL requires.
func newCloudTestSession(t *testing.T, server *httptest.Server) pointInTimeRecoveryTestSession {
	t.Helper()
	transport := server.Client().Transport.(*http.Transport).Clone()
	transport.DialContext = func(ctx context.Context, network, _ string) (net.Conn, error) {
		var dialer net.Dialer
		return dialer.DialContext(ctx, network, server.Listener.Addr().String())
	}
	transport.TLSClientConfig.ServerName = "example.com"
	httpClient := &http.Client{Transport: transport}

	resourceController, err := rc.NewResourceControllerV2(&rc.ResourceControllerV2Options{
		URL:           "https://resource-controller.cloud.ibm.com",
		Authenticator: &core.BearerTokenAuthenticator{BearerToken: "test-token"},
	})
	require.NoError(t, err)
	resourceController.Service.SetHTTPClient(httpClient)

	cloudDatabases, err := clouddatabasesv5.NewCloudDatabasesV5(&clouddatabasesv5.CloudDatabasesV5Options{
		URL:           "https://api.ca-mon.databases.cloud.ibm.com/v5/ibm",
		Authenticator: &core.NoAuthAuthenticator{},
	})
	require.NoError(t, err)
	cloudDatabases.Service.SetHTTPClient(httpClient)

	return pointInTimeRecoveryTestSession{resourceController: resourceController, cloudDatabases: cloudDatabases}
}

func respondJSON(status int, body string) http.HandlerFunc {
	return func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		_, _ = w.Write([]byte(body))
	}
}

func newPointInTimeRecoveryTestData(t *testing.T) *schema.ResourceData {
	t.Helper()
	return schema.TestResourceDataRaw(t, DataSourceIBMDatabasePointInTimeRecovery().Schema, map[string]interface{}{"deployment_id": testGen2WindowInstanceCRN})
}

func requirePointInTimeRecoveryProblem(t *testing.T, diags diag.Diagnostics) string {
	t.Helper()
	require.Len(t, diags, 1)
	require.Equal(t, diag.Error, diags[0].Severity)

	var problem struct {
		Summary   string `yaml:"summary"`
		Resource  string `yaml:"resource"`
		Operation string `yaml:"operation"`
	}
	require.NoError(t, yaml.Unmarshal([]byte(diags[0].Summary), &problem), diags[0].Summary)
	assert.Equal(t, testPointInTimeRecoveryDataSourceResource, problem.Resource)
	assert.Equal(t, "read", problem.Operation)
	return problem.Summary
}

func TestDescribeRequestFailure(t *testing.T) {
	tests := []struct {
		name     string
		respond  http.HandlerFunc
		expected string
	}{
		{name: "body that is not JSON is reported decoded", respond: http.NotFound, expected: "HTTP 404: 404 page not found"},
		{
			name:     "JSON error body is reported by its API message and trace",
			respond:  respondJSON(http.StatusForbidden, `{"errors":[{"code":"forbidden","message":"Viewer access to the instance is required."}],"trace":"t-1","status_code":403}`),
			expected: "HTTP 403: Viewer access to the instance is required. (trace: t-1)",
		},
		{name: "server error keeps the trace that its message asks for", respond: respondJSON(http.StatusInternalServerError, testWindowServerErrorBody), expected: testWindowServerErrorSummary},
		{
			name:     "JSON error body without a trace is reported by its API message alone",
			respond:  respondJSON(http.StatusForbidden, `{"errors":[{"code":"forbidden","message":"Viewer access to the instance is required."}],"status_code":403}`),
			expected: "HTTP 403: Viewer access to the instance is required.",
		},
		{
			name:     "trace that is not a string is left out",
			respond:  respondJSON(http.StatusForbidden, `{"errors":[{"code":"forbidden","message":"Viewer access to the instance is required."}],"trace":42,"status_code":403}`),
			expected: "HTTP 403: Viewer access to the instance is required.",
		},
		{
			name:     "empty trace is left out",
			respond:  respondJSON(http.StatusForbidden, `{"errors":[{"code":"forbidden","message":"Viewer access to the instance is required."}],"trace":"","status_code":403}`),
			expected: "HTTP 403: Viewer access to the instance is required.",
		},
		{name: "empty body is reported by its status text", respond: func(w http.ResponseWriter, _ *http.Request) { w.WriteHeader(http.StatusBadGateway) }, expected: "HTTP 502: Bad Gateway"},
		{
			name:     "JSON error body whose message the SDK cannot find is reported in full",
			respond:  respondJSON(http.StatusNotFound, `{"errors":"Deployment <none> not found & not restorable"}`),
			expected: `HTTP 404: {"errors":"Deployment <none> not found & not restorable"}`,
		},
		{
			name:     "JSON error body with an empty message is reported in full",
			respond:  respondJSON(http.StatusInternalServerError, `{"message":"","trace":"t-2"}`),
			expected: `HTTP 500: {"message":"","trace":"t-2"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(tt.respond)
			defer server.Close()

			_, response, err := fetchPointInTimeRecoveryWindow(context.Background(), newTestWindowService(t, server.URL), server.URL+testGen2WindowPath)
			require.Error(t, err)
			assert.Equal(t, tt.expected, describeRequestFailure(err, response))
		})
	}

	t.Run("a successful status keeps the SDK error", func(t *testing.T) {
		server := httptest.NewServer(respondJSON(http.StatusOK, `{"earliest_restorable_at":`))
		defer server.Close()

		_, response, err := fetchPointInTimeRecoveryWindow(context.Background(), newTestWindowService(t, server.URL), server.URL+testGen2WindowPath)
		require.Error(t, err)
		require.NotNil(t, response)
		assert.Equal(t, err.Error(), describeRequestFailure(err, response))
	})

	t.Run("no response keeps the error", func(t *testing.T) {
		err := errors.New("dial tcp: connect: connection refused")
		assert.Equal(t, err.Error(), describeRequestFailure(err, nil))
	})

	iamFailures := []struct {
		name        string
		iamResponse http.HandlerFunc
		expected    string
	}{
		{
			name: "an IAM token failure with a body that is not JSON keeps the authenticate step",
			iamResponse: func(w http.ResponseWriter, _ *http.Request) {
				w.Header().Set("Content-Type", "text/html")
				w.WriteHeader(http.StatusServiceUnavailable)
				_, _ = w.Write([]byte("<html><body>upstream unavailable</body></html>"))
			},
			expected: "HTTP 503: " + fmt.Sprintf(core.ERRORMSG_AUTHENTICATE_ERROR, "<html><body>upstream unavailable</body></html>"),
		},
		{
			name:        "an IAM token failure with a JSON body keeps the authenticate step",
			iamResponse: respondJSON(http.StatusBadRequest, `{"errorCode":"BXNIM0415E","errorMessage":"Provided API key could not be found."}`),
			expected:    "HTTP 400: " + fmt.Sprintf(core.ERRORMSG_AUTHENTICATE_ERROR, "Provided API key could not be found."),
		},
	}
	for _, tt := range iamFailures {
		t.Run(tt.name, func(t *testing.T) {
			iam := httptest.NewServer(tt.iamResponse)
			defer iam.Close()
			api := httptest.NewServer(http.HandlerFunc(func(_ http.ResponseWriter, r *http.Request) {
				t.Errorf("unexpected request to %s", r.URL.Path)
			}))
			defer api.Close()

			service, err := core.NewBaseService(&core.ServiceOptions{
				URL:           api.URL,
				Authenticator: &core.IamAuthenticator{ApiKey: "test-api-key", URL: iam.URL}, // pragma: allowlist secret
			})
			require.NoError(t, err)

			_, response, err := fetchPointInTimeRecoveryWindow(context.Background(), service, api.URL+testGen2WindowPath)
			require.Error(t, err)
			assert.Equal(t, tt.expected, describeRequestFailure(err, response))
		})
	}
}

func TestPointInTimeRecoveryReadReportsHTTPFailures(t *testing.T) {
	gen2Instance := respondJSON(http.StatusOK, `{
		"resource_plan_id": "databases-for-postgresql-standard-gen2",
		"extensions": {"dataservices": {"point_in_time_recovery_status": {"restorable_window_url": "`+testGen2WindowURL+`"}}}
	}`)
	classicInstance := respondJSON(http.StatusOK, `{"resource_plan_id": "databases-for-postgresql-standard"}`)

	tests := []struct {
		name            string
		instance        http.HandlerFunc
		pointInTimeData http.HandlerFunc
		expectedSummary string
	}{
		{
			name:            "Gen2 window route that does not exist",
			instance:        gen2Instance,
			pointInTimeData: http.NotFound,
			expectedSummary: "GET point-in-time recovery window failed: HTTP 404: 404 page not found",
		},
		{
			name:     "Gen2 window refused by the API",
			instance: gen2Instance,
			pointInTimeData: respondJSON(http.StatusForbidden,
				`{"errors":[{"code":"forbidden","message":"You are not authorized to read this instance's point-in-time recovery window."}],"trace":"t-1","status_code":403}`),
			expectedSummary: "GET point-in-time recovery window failed: HTTP 403: You are not authorized to read this instance's point-in-time recovery window. (trace: t-1)",
		},
		{
			name:            "Gen2 window that the API cannot read",
			instance:        gen2Instance,
			pointInTimeData: respondJSON(http.StatusInternalServerError, testWindowServerErrorBody),
			expectedSummary: "GET point-in-time recovery window failed: " + testWindowServerErrorSummary,
		},
		{
			name:            "Classic data route that does not exist",
			instance:        classicInstance,
			pointInTimeData: http.NotFound,
			expectedSummary: "GetPitrDataWithContext failed: HTTP 404: 404 page not found",
		},
		{
			name:            "Classic error body that the SDK reads no message from",
			instance:        classicInstance,
			pointInTimeData: respondJSON(http.StatusNotFound, `{"errors":{"id":["is not a valid deployment"]}}`),
			expectedSummary: `GetPitrDataWithContext failed: HTTP 404: {"errors":{"id":["is not a valid deployment"]}}`,
		},
		{
			name:            "instance lookup refused",
			instance:        respondJSON(http.StatusNotFound, `{"message":"Instance not found","status_code":404}`),
			expectedSummary: "failed to get resource instance: HTTP 404: Instance not found",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if strings.HasPrefix(r.URL.Path, "/v2/resource_instances/") {
					tt.instance(w, r)
					return
				}
				if tt.pointInTimeData == nil {
					t.Errorf("unexpected request to %s", r.URL.Path)
					http.Error(w, "unexpected request", http.StatusInternalServerError)
					return
				}
				tt.pointInTimeData(w, r)
			}))
			defer server.Close()

			diags := DataSourceIBMDatabasePointInTimeRecoveryRead(context.Background(), newPointInTimeRecoveryTestData(t), newCloudTestSession(t, server))

			assert.Equal(t, tt.expectedSummary, requirePointInTimeRecoveryProblem(t, diags))
			assert.NotContains(t, diags[0].Summary, "RawResult")
		})
	}
}

func TestPointInTimeRecoveryReadLabelsClientFailures(t *testing.T) {
	clientErr := errors.New("no IAM access token")

	tests := []struct {
		name           string
		read           func(context.Context, *schema.ResourceData, interface{}) diag.Diagnostics
		meta           interface{}
		summaryContain string
	}{
		{
			name:           "Resource Controller client for the plan lookup",
			read:           DataSourceIBMDatabasePointInTimeRecoveryRead,
			meta:           pointInTimeRecoveryTestSession{resourceControllerErr: clientErr},
			summaryContain: clientErr.Error(),
		},
		{
			name:           "Classic client",
			read:           newDataSourceIBMDatabasePointInTimeRecoveryClassicBackend().Read,
			meta:           pointInTimeRecoveryTestSession{cloudDatabasesErr: clientErr},
			summaryContain: clientErr.Error(),
		},
		{
			name:           "Gen2 window URL not published",
			read:           newDataSourceIBMDatabasePointInTimeRecoveryGen2Backend(&rc.ResourceInstance{}).Read,
			summaryContain: "publish no dataservices.point_in_time_recovery_status.restorable_window_url",
		},
		{
			name:           "Gen2 Resource Controller client for the window",
			read:           newDataSourceIBMDatabasePointInTimeRecoveryGen2Backend(&rc.ResourceInstance{Extensions: windowURLExtensions(testGen2WindowURL)}).Read,
			meta:           pointInTimeRecoveryTestSession{resourceControllerErr: clientErr},
			summaryContain: clientErr.Error(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			diags := tt.read(context.Background(), newPointInTimeRecoveryTestData(t), tt.meta)

			assert.Contains(t, requirePointInTimeRecoveryProblem(t, diags), tt.summaryContain)
		})
	}
}

func TestPointInTimeRecoveryReadSucceeds(t *testing.T) {
	gen2Instance := respondJSON(http.StatusOK, `{
		"resource_plan_id": "databases-for-postgresql-standard-gen2",
		"extensions": {"dataservices": {"point_in_time_recovery_status": {"restorable_window_url": "`+testGen2WindowURL+`"}}}
	}`)
	classicInstance := respondJSON(http.StatusOK, `{"resource_plan_id": "databases-for-postgresql-standard"}`)

	tests := []struct {
		name                  string
		instance              http.HandlerFunc
		pointInTimeData       string
		expected              map[string]string
		warningContains       string
		expectedAuthorization string
	}{
		{
			name:     "Gen2 window with a closed and an open period",
			instance: gen2Instance,
			pointInTimeData: `{
				"earliest_restorable_at": "2026-09-20T09:30:01.000Z",
				"latest_restorable_at": "2026-09-23T10:59:59.999Z",
				"retention_days": 7,
				"archiving": {"status": "healthy", "updated_at": "2026-09-27T09:29:40.000Z"},
				"unavailable_periods": [
					{"from": "2026-09-22T01:00:00.000Z", "until": "2026-09-22T03:00:01.000Z"},
					{"from": "2026-09-23T11:00:00.000Z", "until": null}
				],
				"not_restorable_reason": null,
				"evaluated_at": "2026-09-27T09:30:00.000Z"
			}`,
			expected: map[string]string{
				"earliest_point_in_time_recovery_time": "2026-09-20T09:30:01.000Z",
				"latest_point_in_time_recovery_time":   "2026-09-23T10:59:59.999Z",
				"retention_days":                       "7",
				"archiving_status":                     "healthy",
				"unavailable_periods.#":                "2",
				"unavailable_periods.0.until":          "2026-09-22T03:00:01.000Z",
				"unavailable_periods.1.from":           "2026-09-23T11:00:00.000Z",
				"unavailable_periods.1.until":          "",
				"not_restorable_reason":                "",
			},
			expectedAuthorization: "Bearer test-token",
		},
		{
			name:     "Gen2 window while archiving is delayed",
			instance: gen2Instance,
			pointInTimeData: `{
				"earliest_restorable_at": null,
				"latest_restorable_at": null,
				"retention_days": 7,
				"archiving": {"status": "delayed", "updated_at": "2026-10-09T20:00:00.000Z"},
				"unavailable_periods": [],
				"not_restorable_reason": "archive_delayed",
				"evaluated_at": "2026-10-09T20:04:41.956Z"
			}`,
			expected: map[string]string{
				"earliest_point_in_time_recovery_time": "",
				"latest_point_in_time_recovery_time":   "",
				"archiving_status":                     "delayed",
				"not_restorable_reason":                "archive_delayed",
			},
			warningContains:       "Retry once archiving_status is healthy again.",
			expectedAuthorization: "Bearer test-token",
		},
		{
			name:            "Classic data",
			instance:        classicInstance,
			pointInTimeData: `{"point_in_time_recovery_data": {"earliest_point_in_time_recovery_time": "2026-09-20T09:30:01Z"}}`,
			expected: map[string]string{
				"earliest_point_in_time_recovery_time": "2026-09-20T09:30:01Z",
				"latest_point_in_time_recovery_time":   "",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var dataAuthorization string
			server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if strings.HasPrefix(r.URL.Path, "/v2/resource_instances/") {
					tt.instance(w, r)
					return
				}
				dataAuthorization = r.Header.Get("Authorization")
				respondJSON(http.StatusOK, tt.pointInTimeData)(w, r)
			}))
			defer server.Close()

			d := newPointInTimeRecoveryTestData(t)
			diags := DataSourceIBMDatabasePointInTimeRecoveryRead(context.Background(), d, newCloudTestSession(t, server))

			if tt.warningContains == "" {
				assert.Empty(t, diags)
			} else {
				require.Len(t, diags, 1)
				assert.Equal(t, diag.Warning, diags[0].Severity)
				assert.Contains(t, diags[0].Detail, testGen2WindowInstanceCRN)
				assert.Contains(t, diags[0].Detail, tt.warningContains)
			}
			assert.Equal(t, testGen2WindowInstanceCRN, d.Id())
			attributes := d.State().Attributes
			for key, value := range tt.expected {
				assert.Equal(t, value, attributes[key], key)
			}
			assert.Equal(t, tt.expectedAuthorization, dataAuthorization)
		})
	}
}
