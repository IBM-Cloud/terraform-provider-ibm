// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package logs_test

import (
	"fmt"
	"testing"

	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/logs-go-sdk/logsv0"
)

func TestAccIbmLogsAlertDefinitionBasic(t *testing.T) {
	var conf logsv0.AlertDefinition
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	description := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	nameUpdate := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	descriptionUpdate := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheckCloudLogs(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmLogsAlertDefinitionDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmLogsAlertDefinitionConfigBasic(name, description),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmLogsAlertDefinitionExists("ibm_logs_alert_definition.logs_alert_definition_instance", conf),
					resource.TestCheckResourceAttr("ibm_logs_alert_definition.logs_alert_definition_instance", "name", name),
				),
			},
			resource.TestStep{
				Config: testAccCheckIbmLogsAlertDefinitionConfigBasic(nameUpdate, descriptionUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_logs_alert_definition.logs_alert_definition_instance", "name", nameUpdate),
				),
			},
		},
	})
}

func TestAccIbmLogsAlertDefinitionAllArgs(t *testing.T) {
	var conf logsv0.AlertDefinition
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	nameUpdate := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	description := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	descriptionUpdate := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	enabled := "false"
	enabledUpdate := "true"
	priority := "p2"
	priorityUpdate := "p3"
	typeVar := "logs_anomaly"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheckCloudLogs(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmLogsAlertDefinitionDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmLogsAlertDefinitionConfig(name, description, enabled, priority, typeVar),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmLogsAlertDefinitionExists("ibm_logs_alert_definition.logs_alert_definition_instance", conf),
					resource.TestCheckResourceAttr("ibm_logs_alert_definition.logs_alert_definition_instance", "name", name),
					resource.TestCheckResourceAttr("ibm_logs_alert_definition.logs_alert_definition_instance", "description", description),
					resource.TestCheckResourceAttr("ibm_logs_alert_definition.logs_alert_definition_instance", "enabled", enabled),
					resource.TestCheckResourceAttr("ibm_logs_alert_definition.logs_alert_definition_instance", "priority", priority),
					resource.TestCheckResourceAttr("ibm_logs_alert_definition.logs_alert_definition_instance", "type", typeVar),
				),
			},
			resource.TestStep{
				Config: testAccCheckIbmLogsAlertDefinitionConfig(nameUpdate, descriptionUpdate, enabledUpdate, priorityUpdate, typeVar),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_logs_alert_definition.logs_alert_definition_instance", "name", nameUpdate),
					resource.TestCheckResourceAttr("ibm_logs_alert_definition.logs_alert_definition_instance", "description", descriptionUpdate),
					resource.TestCheckResourceAttr("ibm_logs_alert_definition.logs_alert_definition_instance", "enabled", enabledUpdate),
					resource.TestCheckResourceAttr("ibm_logs_alert_definition.logs_alert_definition_instance", "priority", priorityUpdate),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_logs_alert_definition.logs_alert_definition_instance",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIbmLogsAlertDefinitionConfigBasic(name string, description string) string {
	return fmt.Sprintf(`
	resource "ibm_logs_alert_definition" "logs_alert_definition_instance" {
		instance_id = "%s"
		region      = "%s"
		name        = "%s"
		description = "%s"
		enabled     = true
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
	`, acc.LogsInstanceId, acc.LogsInstanceRegion, name, description)
}

func testAccCheckIbmLogsAlertDefinitionConfig(name string, description string, enabled string, priority string, typeVar string) string {
	return fmt.Sprintf(`
	resource "ibm_logs_alert_definition" "logs_alert_definition_instance" {
		instance_id   = "%s"
		region        = "%s"
		name          = "%s"
		description   = "%s"
		enabled       = %s
		group_by_keys = []
		incidents_settings {
			minutes   = 1
			notify_on = "triggered_only_unspecified"
		}
		logs_anomaly {
			condition_type      = "more_than_usual_or_unspecified"
			evaluation_delay_ms = 0
			logs_filter {
				simple_filter {
					label_filters {
						application_name {
							operation = "is_or_unspecified"
							value     = "sev5"
						}
						application_name {
							operation = "is_or_unspecified"
							value     = "sev4"
						}
						severities = []
						subsystem_name {
							operation = "is_or_unspecified"
							value     = "sev4-logs"
						}
					}
					lucene_query = "\"push1\""
				}
			}
			rules {
				condition {
						minimum_threshold = 1
						time_window {
						logs_time_window_specific_value = "minutes_5_or_unspecified"
					}
				}
			}
		}
		notification_group {
			group_by_keys = []
			webhooks {
			integration {
				integration_id = 27
			}
			minutes   = 0
			notify_on = ""
			}
		}
		priority     = "%s"
		type         = "%s"

	}
	`, acc.LogsInstanceId, acc.LogsInstanceRegion, name, description, enabled, priority, typeVar)
}

func testAccCheckIbmLogsAlertDefinitionExists(n string, obj logsv0.AlertDefinition) resource.TestCheckFunc {

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

		getAlertDefOptions := &logsv0.GetAlertDefOptions{}

		getAlertDefOptions.SetID(core.UUIDPtr(strfmt.UUID(resourceID[2])))

		alertDefinitionIntf, _, err := logsClient.GetAlertDef(getAlertDefOptions)
		if err != nil {
			return err
		}

		alertDefinition := alertDefinitionIntf.(*logsv0.AlertDefinition)
		obj = *alertDefinition
		return nil
	}
}

func testAccCheckIbmLogsAlertDefinitionDestroy(s *terraform.State) error {
	logsClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).LogsV0()
	if err != nil {
		return err
	}
	logsClient = getTestClientWithLogsInstanceEndpoint(logsClient)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_logs_alert_definition" {
			continue
		}
		resourceID, err := flex.IdParts(rs.Primary.ID)
		if err != nil {
			return err
		}

		getAlertDefOptions := &logsv0.GetAlertDefOptions{}

		getAlertDefOptions.SetID(core.UUIDPtr(strfmt.UUID(resourceID[2])))

		// Try to find the key
		_, response, err := logsClient.GetAlertDef(getAlertDefOptions)

		if err == nil {
			return fmt.Errorf("logs_alert_definition still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for logs_alert_definition (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
