// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.104.0-b4a47c49-20250418-184351
 */

package logs_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmLogsAlertDefinitionDataSourceBasic(t *testing.T) {
	alertDefinitionName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	alertDefinitionDescription := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckCloudLogs(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmLogsAlertDefinitionDataSourceConfigBasic(alertDefinitionName, alertDefinitionDescription),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_logs_alert_definition.logs_alert_definition_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_alert_definition.logs_alert_definition_instance", "logs_alert_definition_id"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_alert_definition.logs_alert_definition_instance", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_alert_definition.logs_alert_definition_instance", "type"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_alert_definition.logs_alert_definition_instance", "logs_new_value.#"),
				),
			},
		},
	})
}

func testAccCheckIbmLogsAlertDefinitionDataSourceConfigBasic(alertDefinitionName string, alertDefinitionDescription string) string {
	return fmt.Sprintf(`
		resource "ibm_logs_alert_definition" "logs_alert_definition_instance" {
			instance_id = "%s"
			region      = "%s"
			name        = "%s"
			description = "%s"
			enabled     = true
			group_by_keys    = [
                "ibm.logId",
                ]
			entity_labels = {
				"key" = "value"
			}
			logs_new_value {
				notification_payload_filter = []
				logs_filter {
					simple_filter {
						lucene_query = "text"
					}
				}
				rules {
					condition {
						keypath_to_track = "ibm.logId"
						time_window {
							logs_new_value_time_window_specific_value = "hours_12_or_unspecified"
						}
					}
				}
			}
			priority     = "p5_or_unspecified"
			type         = "logs_new_value"

		}

		data "ibm_logs_alert_definition" "logs_alert_definition_instance" {
			instance_id = "%[1]s"
			region      = "%[2]s"
			logs_alert_definition_id = ibm_logs_alert_definition.logs_alert_definition_instance.alert_def_id
		}
	`, acc.LogsInstanceId, acc.LogsInstanceRegion, alertDefinitionName, alertDefinitionDescription)
}
