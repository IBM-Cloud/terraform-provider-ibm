// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package logs_test

import (
	"fmt"
	"testing"

	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/logs-go-sdk/logsv0"
	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/conns"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/flex"
)

func TestAccIbmLogsAlertBasic(t *testing.T) {
	var conf logsv0.Alert
	name := fmt.Sprintf("tf_name_応答時間モニター!_%d", acctest.RandIntRange(10, 100))
	isActive := "false"
	severity := "info_or_unspecified"
	nameUpdate := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	isActiveUpdate := "true"
	severityUpdate := "error"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheckCloudLogs(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmLogsAlertDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmLogsAlertConfigBasic(name, isActive, severity),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmLogsAlertExists("ibm_logs_alert.logs_alert_instance", conf),
					resource.TestCheckResourceAttr("ibm_logs_alert.logs_alert_instance", "name", name),
					resource.TestCheckResourceAttr("ibm_logs_alert.logs_alert_instance", "is_active", isActive),
					resource.TestCheckResourceAttr("ibm_logs_alert.logs_alert_instance", "severity", severity),
				),
			},
			resource.TestStep{
				Config: testAccCheckIbmLogsAlertConfigBasic(nameUpdate, isActiveUpdate, severityUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_logs_alert.logs_alert_instance", "name", nameUpdate),
					resource.TestCheckResourceAttr("ibm_logs_alert.logs_alert_instance", "is_active", isActiveUpdate),
					resource.TestCheckResourceAttr("ibm_logs_alert.logs_alert_instance", "severity", severityUpdate),
				),
			},
		},
	})
}

func TestAccIbmLogsAlertAllArgs(t *testing.T) {
	var conf logsv0.Alert
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	description := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	isActive := "false"
	severity := "info_or_unspecified"
	nameUpdate := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	descriptionUpdate := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	isActiveUpdate := "true"
	severityUpdate := "error"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheckCloudLogs(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmLogsAlertDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmLogsAlertConfig(name, description, isActive, severity),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmLogsAlertExists("ibm_logs_alert.logs_alert_instance", conf),
					resource.TestCheckResourceAttr("ibm_logs_alert.logs_alert_instance", "name", name),
					resource.TestCheckResourceAttr("ibm_logs_alert.logs_alert_instance", "description", description),
					resource.TestCheckResourceAttr("ibm_logs_alert.logs_alert_instance", "is_active", isActive),
					resource.TestCheckResourceAttr("ibm_logs_alert.logs_alert_instance", "severity", severity),
				),
			},
			resource.TestStep{
				Config: testAccCheckIbmLogsAlertConfig(nameUpdate, descriptionUpdate, isActiveUpdate, severityUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_logs_alert.logs_alert_instance", "name", nameUpdate),
					resource.TestCheckResourceAttr("ibm_logs_alert.logs_alert_instance", "description", descriptionUpdate),
					resource.TestCheckResourceAttr("ibm_logs_alert.logs_alert_instance", "is_active", isActiveUpdate),
					resource.TestCheckResourceAttr("ibm_logs_alert.logs_alert_instance", "severity", severityUpdate),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_logs_alert.logs_alert_instance",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIbmLogsAlertConfigBasic(name string, isActive string, severity string) string {
	return fmt.Sprintf(`
	resource "ibm_logs_alert" "logs_alert_instance" {
		instance_id = "%s"
		region      = "%s"
		name        = "%s"
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
		meta_labels {
			key   = "label1"
			value = "value"
		}
		meta_labels {
			key   = "label2"
			value = "true"
		}
		incident_settings {
			retriggering_period_seconds = 43200
			notify_on                   = "triggered_only"
		}
	}
`, acc.LogsInstanceId, acc.LogsInstanceRegion, name, isActive, severity)
}

func testAccCheckIbmLogsAlertConfig(name string, description string, isActive string, severity string) string {
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
		meta_labels {
			key   = "label1"
			value = "value"
		}
		meta_labels {
			key   = "label2"
			value = "value"
		}
		meta_labels {
			key   = "label3"
			value = "value"
		}
		incident_settings {
			retriggering_period_seconds = 43200
			notify_on                   = "triggered_only"
		}
	}
	`, acc.LogsInstanceId, acc.LogsInstanceRegion, name, description, isActive, severity)
}

func testAccCheckIbmLogsAlertExists(n string, obj logsv0.Alert) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		logsClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).LogsV0()
		if err != nil {
			return err
		}
		logsClient = getTestClientWithLogsInstanceEndpoint(logsClient)

		resourceID, err := flex.IdParts(rs.Primary.ID)
		if err != nil {
			return err
		}

		getAlertOptions := &logsv0.GetAlertOptions{}

		getAlertOptions.SetID(core.UUIDPtr(strfmt.UUID(resourceID[2])))

		alert, _, err := logsClient.GetAlert(getAlertOptions)
		if err != nil {
			return err
		}

		obj = *alert
		return nil
	}
}

func testAccCheckIbmLogsAlertDestroy(s *terraform.State) error {
	logsClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).LogsV0()
	if err != nil {
		return err
	}
	logsClient = getTestClientWithLogsInstanceEndpoint(logsClient)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_logs_alert" {
			continue
		}

		resourceID, err := flex.IdParts(rs.Primary.ID)
		if err != nil {
			return err
		}
		getAlertOptions := &logsv0.GetAlertOptions{}

		getAlertOptions.SetID(core.UUIDPtr(strfmt.UUID(resourceID[2])))

		// Try to find the key
		_, response, err := logsClient.GetAlert(getAlertOptions)

		if err == nil {
			return fmt.Errorf("logs_alert still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for logs_alert (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
