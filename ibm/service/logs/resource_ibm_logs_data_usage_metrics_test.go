// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package logs_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM/logs-go-sdk/logsv0"
)

func TestAccIbmLogsDataUsageMetricsBasic(t *testing.T) {
	var conf logsv0.DataUsageMetricsExportStatus
	enabled := "false"
	enabledUpdate := "true"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheckCloudLogs(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmLogsDataUsageMetricsDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmLogsDataUsageMetricsConfigBasic(enabled),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmLogsDataUsageMetricsExists("ibm_logs_data_usage_metrics.logs_data_usage_metrics_instance", conf),
					resource.TestCheckResourceAttr("ibm_logs_data_usage_metrics.logs_data_usage_metrics_instance", "enabled", enabled),
				),
			},
			resource.TestStep{
				Config: testAccCheckIbmLogsDataUsageMetricsConfigBasic(enabledUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_logs_data_usage_metrics.logs_data_usage_metrics_instance", "enabled", enabledUpdate),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_logs_data_usage_metrics.logs_data_usage_metrics_instance",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIbmLogsDataUsageMetricsConfigBasic(enabled string) string {
	return fmt.Sprintf(`
		resource "ibm_logs_data_usage_metrics" "logs_data_usage_metrics_instance" {
			instance_id = "%s"
			region      = "%s"
			enabled = %s
		}
	`, acc.LogsInstanceId, acc.LogsInstanceRegion, enabled)
}

func testAccCheckIbmLogsDataUsageMetricsExists(n string, obj logsv0.DataUsageMetricsExportStatus) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		_, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		logsClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).LogsV0()
		if err != nil {
			return err
		}
		logsClient = getTestClientWithLogsInstanceEndpoint(logsClient)

		getDataUsageMetricsExportStatusOptions := &logsv0.GetDataUsageMetricsExportStatusOptions{}

		dataUsageMetricsExportStatus, _, err := logsClient.GetDataUsageMetricsExportStatus(getDataUsageMetricsExportStatusOptions)
		if err != nil {
			return err
		}

		obj = *dataUsageMetricsExportStatus
		return nil
	}
}

func testAccCheckIbmLogsDataUsageMetricsDestroy(s *terraform.State) error {
	// As there is no destroy avoid this step
	return nil
}
