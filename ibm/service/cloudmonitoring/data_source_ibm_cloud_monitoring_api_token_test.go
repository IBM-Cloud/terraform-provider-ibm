// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package cloudmonitoring_test

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/cloudmonitoring"
	"github.com/IBM/go-sdk-core/v5/core"
)

func TestAccIBMCloudMonitoringAPITokenDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckCloudMonitoring(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCloudMonitoringAPITokenDataSourceConfig(acc.CloudMonitoringInstanceID, ""),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cloud_monitoring_api_token.token", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cloud_monitoring_api_token.token", "value"),
					resource.TestCheckResourceAttrSet("data.ibm_cloud_monitoring_api_token.token", "region"),
				),
			},
			{
				Config: testAccCheckIBMCloudMonitoringAPITokenDataSourceConfig(acc.CloudMonitoringInstanceID, "public"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cloud_monitoring_api_token.token", "value"),
					resource.TestCheckResourceAttr("data.ibm_cloud_monitoring_api_token.token", "endpoint_type", "public"),
				),
			},
		},
	})
}

func TestAccIBMCloudMonitoringAPITokenDataSourceInvalidEndpointType(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config:      testAccCheckIBMCloudMonitoringAPITokenDataSourceConfig("00000000-0000-0000-0000-000000000000", "internal"),
				ExpectError: regexp.MustCompile(`expected endpoint_type to be one of`),
			},
		},
	})
}

func testAccCheckIBMCloudMonitoringAPITokenDataSourceConfig(instanceID, endpointType string) string {
	endpoint := ""
	if endpointType != "" {
		endpoint = fmt.Sprintf("endpoint_type = %q", endpointType)
	}
	return fmt.Sprintf(`
	data "ibm_cloud_monitoring_api_token" "token" {
		instance_id = %q
		%s
	}
	`, instanceID, endpoint)
}

func TestCloudMonitoringInstanceGUID(t *testing.T) {
	cases := []struct {
		in, want string
		wantErr  bool
	}{
		{in: "1234abcd-0000-1111-2222-333344445555", want: "1234abcd-0000-1111-2222-333344445555"},
		{in: "crn:v1:bluemix:public:sysdig-monitor:us-south:a/acct123:1234abcd-0000-1111-2222-333344445555::", want: "1234abcd-0000-1111-2222-333344445555"},
		{in: "  ", wantErr: true},
		{in: "crn:v1:bluemix:public:sysdig-monitor", wantErr: true},
		{in: "crn:v1:bluemix:public:sysdig-monitor:us-south:a/acct123:::", wantErr: true},
	}
	for _, c := range cases {
		got, err := cloudmonitoring.CloudMonitoringInstanceGUID(c.in)
		if c.wantErr != (err != nil) || got != c.want {
			t.Errorf("CloudMonitoringInstanceGUID(%q) = %q, %v", c.in, got, err)
		}
	}
}

func TestCloudMonitoringServiceURL(t *testing.T) {
	if got, _ := cloudmonitoring.CloudMonitoringServiceURL("us-south", "public"); got != "https://us-south.monitoring.cloud.ibm.com" {
		t.Errorf("unexpected public URL %q", got)
	}
	if got, _ := cloudmonitoring.CloudMonitoringServiceURL("eu-de", "private"); got != "https://private.eu-de.monitoring.cloud.ibm.com" {
		t.Errorf("unexpected private URL %q", got)
	}
	if _, err := cloudmonitoring.CloudMonitoringServiceURL("", "public"); err == nil {
		t.Error("expected error for empty region")
	}
}

func TestParseCloudMonitoringAPITokenResponse(t *testing.T) {
	if got, err := cloudmonitoring.ParseCloudMonitoringAPITokenResponse([]byte(`{"token":{"key":"abc"}}`)); err != nil || got != "abc" {
		t.Errorf("got %q, %v", got, err)
	}
	for _, body := range []string{`{}`, `{"token":null}`, `{"token":{"key":""}}`, `not json`} {
		if _, err := cloudmonitoring.ParseCloudMonitoringAPITokenResponse([]byte(body)); err == nil {
			t.Errorf("expected error for body %s", body)
		}
	}
}

func TestFetchCloudMonitoringAPIToken(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet || r.URL.Path != "/api/token" {
			t.Errorf("unexpected request %s %s", r.Method, r.URL.Path)
		}
		if got := r.Header.Get("Authorization"); got != "Bearer iam-token" {
			t.Errorf("unexpected Authorization header %q", got)
		}
		if r.Header.Get("IBMInstanceID") == "bad-guid" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(w, `{"message":"unauthorized"}`)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"token":{"key":"monitoring-key"}}`)
	}))
	defer server.Close()

	service, err := core.NewBaseService(&core.ServiceOptions{
		URL:           server.URL,
		Authenticator: &core.BearerTokenAuthenticator{BearerToken: "iam-token"},
	})
	if err != nil {
		t.Fatal(err)
	}

	got, err := cloudmonitoring.FetchCloudMonitoringAPIToken(context.Background(), service, "good-guid")
	if err != nil || got != "monitoring-key" {
		t.Errorf("got %q, %v", got, err)
	}
	if _, err := cloudmonitoring.FetchCloudMonitoringAPIToken(context.Background(), service, "bad-guid"); err == nil || !regexp.MustCompile(`status 401`).MatchString(err.Error()) {
		t.Errorf("expected status 401 error, got %v", err)
	}
}
