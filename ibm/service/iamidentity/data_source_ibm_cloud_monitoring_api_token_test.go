// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package iamidentity_test

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/iamidentity"
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
		got, err := iamidentity.CloudMonitoringInstanceGUID(c.in)
		if c.wantErr != (err != nil) || got != c.want {
			t.Errorf("CloudMonitoringInstanceGUID(%q) = %q, %v", c.in, got, err)
		}
	}
}

func TestCloudMonitoringAPITokenURL(t *testing.T) {
	t.Setenv(iamidentity.CloudMonitoringEndpointEnv, "")
	if got, _ := iamidentity.CloudMonitoringAPITokenURL("us-south", "public"); got != "https://us-south.monitoring.cloud.ibm.com/api/token" {
		t.Errorf("unexpected public URL %q", got)
	}
	if got, _ := iamidentity.CloudMonitoringAPITokenURL("eu-de", "private"); got != "https://private.eu-de.monitoring.cloud.ibm.com/api/token" {
		t.Errorf("unexpected private URL %q", got)
	}
	if _, err := iamidentity.CloudMonitoringAPITokenURL("", "public"); err == nil {
		t.Error("expected error for empty region")
	}
	t.Setenv(iamidentity.CloudMonitoringEndpointEnv, "https://example.test/")
	if got, _ := iamidentity.CloudMonitoringAPITokenURL("", "public"); got != "https://example.test/api/token" {
		t.Errorf("unexpected override URL %q", got)
	}
}

func TestParseCloudMonitoringAPITokenResponse(t *testing.T) {
	if got, err := iamidentity.ParseCloudMonitoringAPITokenResponse([]byte(`{"token":{"key":"abc"}}`)); err != nil || got != "abc" {
		t.Errorf("got %q, %v", got, err)
	}
	for _, body := range []string{`{}`, `{"token":null}`, `{"token":{"key":""}}`, `not json`} {
		if _, err := iamidentity.ParseCloudMonitoringAPITokenResponse([]byte(body)); err == nil {
			t.Errorf("expected error for body %s", body)
		}
	}
}
