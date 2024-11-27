// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package logs_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmLogsDataUsageMetricsDataSourceBasic(t *testing.T) {
	dataUsageMetricsExportStatusEnabled := "false"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckCloudLogs(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmLogsDataUsageMetricsDataSourceConfigBasic(dataUsageMetricsExportStatusEnabled),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_logs_data_usage_metrics.logs_data_usage_metrics_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_data_usage_metrics.logs_data_usage_metrics_instance", "enabled"),
				),
			},
		},
	})
}

func testAccCheckIbmLogsDataUsageMetricsDataSourceConfigBasic(dataUsageMetricsExportStatusEnabled string) string {
	return fmt.Sprintf(`
		resource "ibm_logs_data_usage_metrics" "logs_data_usage_metrics_instance" {
			instance_id = "%s"
			region      = "%s"
			enabled = %s
		}

		data "ibm_logs_data_usage_metrics" "logs_data_usage_metrics_instance" {
			instance_id = ibm_logs_data_usage_metrics.logs_data_usage_metrics_instance.instance_id
			region      = ibm_logs_data_usage_metrics.logs_data_usage_metrics_instance.region
			depends_on = [
				ibm_logs_data_usage_metrics.logs_data_usage_metrics_instance
			]
		}
	`, acc.LogsInstanceId, acc.LogsInstanceRegion, dataUsageMetricsExportStatusEnabled)
}
