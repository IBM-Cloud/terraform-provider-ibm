// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package logs_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmLogsAlertDataSourceBasic(t *testing.T) {
	alertName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	alertIsActive := "false"
	alertSeverity := "info_or_unspecified"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckCloudLogs(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmLogsAlertDataSourceConfigBasic(alertName, alertIsActive, alertSeverity),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_logs_alert.logs_alert_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_alert.logs_alert_instance", "logs_alert_id"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_alert.logs_alert_instance", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_alert.logs_alert_instance", "is_active"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_alert.logs_alert_instance", "severity"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_alert.logs_alert_instance", "condition.#"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_alert.logs_alert_instance", "notification_groups.#"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_alert.logs_alert_instance", "filters.#"),
				),
			},
		},
	})
}

func TestAccIbmLogsAlertDataSourceAllArgs(t *testing.T) {
	alertName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	alertDescription := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	alertIsActive := "false"
	alertSeverity := "info_or_unspecified"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckCloudLogs(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmLogsAlertDataSourceConfig(alertName, alertDescription, alertIsActive, alertSeverity),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_logs_alert.logs_alert_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_alert.logs_alert_instance", "logs_alert_id"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_alert.logs_alert_instance", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_alert.logs_alert_instance", "description"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_alert.logs_alert_instance", "is_active"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_alert.logs_alert_instance", "severity"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_alert.logs_alert_instance", "expiration.#"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_alert.logs_alert_instance", "condition.#"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_alert.logs_alert_instance", "notification_groups.#"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_alert.logs_alert_instance", "filters.#"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_alert.logs_alert_instance", "active_when.#"),
					// resource.TestCheckResourceAttrSet("data.ibm_logs_alert.logs_alert_instance", "notification_payload_filters.#"),
					// resource.TestCheckResourceAttrSet("data.ibm_logs_alert.logs_alert_instance", "meta_labels.#"),
					// resource.TestCheckResourceAttrSet("data.ibm_logs_alert.logs_alert_instance", "meta_labels.0.key"),
					// resource.TestCheckResourceAttrSet("data.ibm_logs_alert.logs_alert_instance", "meta_labels.0.value"),
					// resource.TestCheckResourceAttrSet("data.ibm_logs_alert.logs_alert_instance", "meta_labels_strings.#"),
					// resource.TestCheckResourceAttrSet("data.ibm_logs_alert.logs_alert_instance", "tracing_alert.#"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_alert.logs_alert_instance", "unique_identifier"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_alert.logs_alert_instance", "incident_settings.#"),
				),
			},
		},
	})
}

func testAccCheckIbmLogsAlertDataSourceConfigBasic(alertName string, alertIsActive string, alertSeverity string) string {
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

		data "ibm_logs_alert" "logs_alert_instance" {
			instance_id   = ibm_logs_alert.logs_alert_instance.instance_id
			region 		  = ibm_logs_alert.logs_alert_instance.region
			logs_alert_id = ibm_logs_alert.logs_alert_instance.alert_id
		}
	`, acc.LogsInstanceId, acc.LogsInstanceRegion, alertName)
}

func testAccCheckIbmLogsAlertDataSourceConfig(alertName string, alertDescription string, alertIsActive string, alertSeverity string) string {
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

		data "ibm_logs_alert" "logs_alert_instance" {
			instance_id   = ibm_logs_alert.logs_alert_instance.instance_id
			region 		  = ibm_logs_alert.logs_alert_instance.region
			logs_alert_id = ibm_logs_alert.logs_alert_instance.alert_id
		}
	`, acc.LogsInstanceId, acc.LogsInstanceRegion, alertName, alertDescription, alertIsActive, alertSeverity)
}
