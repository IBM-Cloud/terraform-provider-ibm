// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package logs_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmLogsAlertsDataSourceBasic(t *testing.T) {
	alertName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	alertIsActive := "false"
	alertSeverity := "info_or_unspecified"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckCloudLogs(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmLogsAlertsDataSourceConfigBasic(alertName, alertIsActive, alertSeverity),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_logs_alerts.logs_alerts_instance", "id"),
				),
			},
		},
	})
}

func TestAccIbmLogsAlertsDataSourceAllArgs(t *testing.T) {
	alertName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	alertDescription := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	alertIsActive := "false"
	alertSeverity := "info_or_unspecified"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckCloudLogs(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmLogsAlertsDataSourceConfig(alertName, alertDescription, alertIsActive, alertSeverity),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_logs_alerts.logs_alerts_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_alerts.logs_alerts_instance", "alerts.#"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_alerts.logs_alerts_instance", "alerts.0.id"),
					resource.TestCheckResourceAttr("data.ibm_logs_alerts.logs_alerts_instance", "alerts.0.name", alertName),
					resource.TestCheckResourceAttr("data.ibm_logs_alerts.logs_alerts_instance", "alerts.0.description", alertDescription),
					resource.TestCheckResourceAttr("data.ibm_logs_alerts.logs_alerts_instance", "alerts.0.is_active", alertIsActive),
					resource.TestCheckResourceAttr("data.ibm_logs_alerts.logs_alerts_instance", "alerts.0.severity", alertSeverity),
					resource.TestCheckResourceAttrSet("data.ibm_logs_alerts.logs_alerts_instance", "alerts.0.unique_identifier"),
				),
			},
		},
	})
}

func testAccCheckIbmLogsAlertsDataSourceConfigBasic(alertName string, alertIsActive string, alertSeverity string) string {
	return fmt.Sprintf(`
		resource "ibm_logs_alert" "logs_alert_instance" {
			instance_id = "%s"
			region      = "%s"
			name        = "%s"
			is_active   = false
			severity    = "info_or_unspecified"
			condition {
			new_value {
				parameters {
				threshold          = 1.0
				timeframe          = "timeframe_12_h"
				group_by           = ["ibm.logId"]
				relative_timeframe = "hour_or_unspecified"
				cardinality_fields = []
				}
			}
			}
			notification_groups {
			group_by_fields = ["ibm.logId"]
			}
			filters {
			text        = "text"
			filter_type = "text_or_unspecified"
			}
			meta_labels_strings = []
			incident_settings {
			retriggering_period_seconds = 43200
			notify_on                   = "triggered_only"
			}
		}

		data "ibm_logs_alerts" "logs_alerts_instance" {
			depends_on = [
				ibm_logs_alert.logs_alert_instance
			]
			instance_id = ibm_logs_alert.logs_alert_instance.instance_id
			region 		= ibm_logs_alert.logs_alert_instance.region
		}
	`, acc.LogsInstanceId, acc.LogsInstanceRegion, alertName)
}

func testAccCheckIbmLogsAlertsDataSourceConfig(alertName string, alertDescription string, alertIsActive string, alertSeverity string) string {
	return fmt.Sprintf(`
		resource "ibm_logs_alert" "logs_alert_instance" {
			instance_id = "%s"
			region      = "%s"
			name        = "%s"
			description = "%s"
			is_active   = %s
			severity    = "%s"
			condition {
			new_value {
				parameters {
				threshold          = 1.0
				timeframe          = "timeframe_12_h"
				group_by           = ["ibm.logId"]
				relative_timeframe = "hour_or_unspecified"
				cardinality_fields = []
				}
			}
			}
			notification_groups {
			group_by_fields = ["ibm.logId"]
			}
			filters {
			text        = "text"
			filter_type = "text_or_unspecified"
			}
			meta_labels_strings = []
			incident_settings {
			retriggering_period_seconds = 43200
			notify_on                   = "triggered_only"
			}
		}


		data "ibm_logs_alerts" "logs_alerts_instance" {
			depends_on = [
				ibm_logs_alert.logs_alert_instance
			]
			instance_id = ibm_logs_alert.logs_alert_instance.instance_id
			region 		= ibm_logs_alert.logs_alert_instance.region
		}
	`, acc.LogsInstanceId, acc.LogsInstanceRegion, alertName, alertDescription, alertIsActive, alertSeverity)
}
