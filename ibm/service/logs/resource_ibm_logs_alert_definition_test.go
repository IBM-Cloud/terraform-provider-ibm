// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package logs_test

import (
	"fmt"
	"testing"

	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/logs-go-sdk/logsv0"
)

func TestAccIbmLogsAlertDefinitionBasic(t *testing.T) {
	var conf logsv0.AlertDefinition
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	typeVar := "logs_immediate_or_unspecified"
	nameUpdate := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	typeVarUpdate := "flow"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmLogsAlertDefinitionDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmLogsAlertDefinitionConfigBasic(name, typeVar),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmLogsAlertDefinitionExists("ibm_logs_alert_definition.logs_alert_definition_instance", conf),
					resource.TestCheckResourceAttr("ibm_logs_alert_definition.logs_alert_definition_instance", "name", name),
					resource.TestCheckResourceAttr("ibm_logs_alert_definition.logs_alert_definition_instance", "type", typeVar),
				),
			},
			resource.TestStep{
				Config: testAccCheckIbmLogsAlertDefinitionConfigBasic(nameUpdate, typeVarUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_logs_alert_definition.logs_alert_definition_instance", "name", nameUpdate),
					resource.TestCheckResourceAttr("ibm_logs_alert_definition.logs_alert_definition_instance", "type", typeVarUpdate),
				),
			},
		},
	})
}

func TestAccIbmLogsAlertDefinitionAllArgs(t *testing.T) {
	var conf logsv0.AlertDefinition
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	description := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	enabled := "false"
	priority := "p5_or_unspecified"
	typeVar := "logs_immediate_or_unspecified"
	phantomMode := "true"
	deleted := "false"
	nameUpdate := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	descriptionUpdate := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	enabledUpdate := "true"
	priorityUpdate := "p1"
	typeVarUpdate := "flow"
	phantomModeUpdate := "false"
	deletedUpdate := "true"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmLogsAlertDefinitionDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmLogsAlertDefinitionConfig(name, description, enabled, priority, typeVar, phantomMode, deleted),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmLogsAlertDefinitionExists("ibm_logs_alert_definition.logs_alert_definition_instance", conf),
					resource.TestCheckResourceAttr("ibm_logs_alert_definition.logs_alert_definition_instance", "name", name),
					resource.TestCheckResourceAttr("ibm_logs_alert_definition.logs_alert_definition_instance", "description", description),
					resource.TestCheckResourceAttr("ibm_logs_alert_definition.logs_alert_definition_instance", "enabled", enabled),
					resource.TestCheckResourceAttr("ibm_logs_alert_definition.logs_alert_definition_instance", "priority", priority),
					resource.TestCheckResourceAttr("ibm_logs_alert_definition.logs_alert_definition_instance", "type", typeVar),
					resource.TestCheckResourceAttr("ibm_logs_alert_definition.logs_alert_definition_instance", "phantom_mode", phantomMode),
					resource.TestCheckResourceAttr("ibm_logs_alert_definition.logs_alert_definition_instance", "deleted", deleted),
				),
			},
			resource.TestStep{
				Config: testAccCheckIbmLogsAlertDefinitionConfig(nameUpdate, descriptionUpdate, enabledUpdate, priorityUpdate, typeVarUpdate, phantomModeUpdate, deletedUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_logs_alert_definition.logs_alert_definition_instance", "name", nameUpdate),
					resource.TestCheckResourceAttr("ibm_logs_alert_definition.logs_alert_definition_instance", "description", descriptionUpdate),
					resource.TestCheckResourceAttr("ibm_logs_alert_definition.logs_alert_definition_instance", "enabled", enabledUpdate),
					resource.TestCheckResourceAttr("ibm_logs_alert_definition.logs_alert_definition_instance", "priority", priorityUpdate),
					resource.TestCheckResourceAttr("ibm_logs_alert_definition.logs_alert_definition_instance", "type", typeVarUpdate),
					resource.TestCheckResourceAttr("ibm_logs_alert_definition.logs_alert_definition_instance", "phantom_mode", phantomModeUpdate),
					resource.TestCheckResourceAttr("ibm_logs_alert_definition.logs_alert_definition_instance", "deleted", deletedUpdate),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_logs_alert_definition.logs_alert_definition",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIbmLogsAlertDefinitionConfigBasic(name string, typeVar string) string {
	return fmt.Sprintf(`
		resource "ibm_logs_alert_definition" "logs_alert_definition_instance" {
			name = "%s"
			type = "%s"
			group_by_keys = "FIXME"
		}
	`, name, typeVar)
}

func testAccCheckIbmLogsAlertDefinitionConfig(name string, description string, enabled string, priority string, typeVar string, phantomMode string, deleted string) string {
	return fmt.Sprintf(`

		resource "ibm_logs_alert_definition" "logs_alert_definition_instance" {
			name = "%s"
			description = "%s"
			enabled = %s
			priority = "%s"
			active_on {
				day_of_week = ["sunday"]
				start_time {
					hours = 14
					minutes = 30
				}
				end_time {
					hours = 14
					minutes = 30
				}
			}
			type = "%s"
			group_by_keys = "FIXME"
			incidents_settings {
				notify_on = "triggered_and_resolved"
				minutes = 30
			}
			notification_group {
				group_by_keys = ["key1","key2"]
				webhooks {
					notify_on = "triggered_and_resolved"
					integration {
						integration_id = 123
					}
					minutes = 15
				}
			}
			entity_labels = "FIXME"
			phantom_mode = %s
			deleted = %s
			logs_immediate {
				logs_filter {
					simple_filter {
						lucene_query = "text:"error""
						label_filters {
							application_name {
								value = "my-app"
								operation = "starts_with"
							}
							subsystem_name {
								value = "my-app"
								operation = "starts_with"
							}
							severities = ["critical"]
						}
					}
				}
				notification_payload_filter = ["obj.field"]
			}
			logs_threshold {
				logs_filter {
					simple_filter {
						lucene_query = "text:"error""
						label_filters {
							application_name {
								value = "my-app"
								operation = "starts_with"
							}
							subsystem_name {
								value = "my-app"
								operation = "starts_with"
							}
							severities = ["critical"]
						}
					}
				}
				undetected_values_management {
					trigger_undetected_values = true
					auto_retire_timeframe = "hours_24"
				}
				rules {
					condition {
						threshold = 100.0
						time_window {
							logs_time_window_specific_value = "hours_36"
						}
					}
					override {
						priority = "p1"
					}
				}
				condition_type = "less_than"
				notification_payload_filter = ["obj.field"]
				evaluation_delay_ms = 60000
			}
			logs_ratio_threshold {
				numerator {
					simple_filter {
						lucene_query = "text:"error""
						label_filters {
							application_name {
								value = "my-app"
								operation = "starts_with"
							}
							subsystem_name {
								value = "my-app"
								operation = "starts_with"
							}
							severities = ["critical"]
						}
					}
				}
				numerator_alias = "numerator_alias"
				denominator {
					simple_filter {
						lucene_query = "text:"error""
						label_filters {
							application_name {
								value = "my-app"
								operation = "starts_with"
							}
							subsystem_name {
								value = "my-app"
								operation = "starts_with"
							}
							severities = ["critical"]
						}
					}
				}
				denominator_alias = "denominator_alias"
				rules {
					condition {
						threshold = 10.0
						time_window {
							logs_ratio_time_window_specific_value = "hours_36"
						}
					}
					override {
						priority = "p1"
					}
				}
				condition_type = "less_than"
				notification_payload_filter = ["obj.field"]
				group_by_for = "denumerator_only"
				undetected_values_management {
					trigger_undetected_values = true
					auto_retire_timeframe = "hours_24"
				}
				ignore_infinity = true
				evaluation_delay_ms = 60000
			}
			logs_time_relative_threshold {
				logs_filter {
					simple_filter {
						lucene_query = "text:"error""
						label_filters {
							application_name {
								value = "my-app"
								operation = "starts_with"
							}
							subsystem_name {
								value = "my-app"
								operation = "starts_with"
							}
							severities = ["critical"]
						}
					}
				}
				rules {
					condition {
						threshold = 100.0
						compared_to = "same_day_last_month"
					}
					override {
						priority = "p1"
					}
				}
				condition_type = "less_than"
				ignore_infinity = true
				notification_payload_filter = ["obj.field"]
				undetected_values_management {
					trigger_undetected_values = true
					auto_retire_timeframe = "hours_24"
				}
				evaluation_delay_ms = 60000
			}
			metric_threshold {
				metric_filter {
					promql = "avg_over_time(metric_name[5m]) > 10"
				}
				rules {
					condition {
						threshold = 100.0
						for_over_pct = 80
						of_the_last {
							metric_time_window_specific_value = "hours_36"
						}
					}
					override {
						priority = "p1"
					}
				}
				condition_type = "less_than_or_equals"
				undetected_values_management {
					trigger_undetected_values = true
					auto_retire_timeframe = "hours_24"
				}
				missing_values {
					replace_with_zero = true
				}
				evaluation_delay_ms = 60000
			}
			flow {
				stages {
					timeframe_ms = "60000"
					timeframe_type = "up_to"
					flow_stages_groups {
						groups {
							alert_defs {
								id = "9fab83da-98cb-4f18-a7ba-b6f0435c9673"
								not = true
							}
							next_op = "or"
							alerts_op = "or"
						}
					}
				}
				enforce_suppression = true
			}
			logs_anomaly {
				logs_filter {
					simple_filter {
						lucene_query = "text:"error""
						label_filters {
							application_name {
								value = "my-app"
								operation = "starts_with"
							}
							subsystem_name {
								value = "my-app"
								operation = "starts_with"
							}
							severities = ["critical"]
						}
					}
				}
				rules {
					condition {
						minimum_threshold = 10.0
						time_window {
							logs_time_window_specific_value = "hours_36"
						}
					}
				}
				condition_type = "more_than_usual_or_unspecified"
				notification_payload_filter = ["obj.field"]
				evaluation_delay_ms = 60000
				anomaly_alert_settings {
					percentage_of_deviation = 10.0
				}
			}
			metric_anomaly {
				metric_filter {
					promql = "avg_over_time(metric_name[5m]) > 10"
				}
				rules {
					condition {
						threshold = 10.0
						for_over_pct = 20
						of_the_last {
							metric_time_window_specific_value = "hours_36"
						}
						min_non_null_values_pct = 10
					}
				}
				condition_type = "less_than_usual"
				evaluation_delay_ms = 60000
				anomaly_alert_settings {
					percentage_of_deviation = 10.0
				}
			}
			logs_new_value {
				logs_filter {
					simple_filter {
						lucene_query = "text:"error""
						label_filters {
							application_name {
								value = "my-app"
								operation = "starts_with"
							}
							subsystem_name {
								value = "my-app"
								operation = "starts_with"
							}
							severities = ["critical"]
						}
					}
				}
				rules {
					condition {
						keypath_to_track = "metadata.field"
						time_window {
							logs_new_value_time_window_specific_value = "months_3"
						}
					}
				}
				notification_payload_filter = ["obj.field"]
			}
			logs_unique_count {
				logs_filter {
					simple_filter {
						lucene_query = "text:"error""
						label_filters {
							application_name {
								value = "my-app"
								operation = "starts_with"
							}
							subsystem_name {
								value = "my-app"
								operation = "starts_with"
							}
							severities = ["critical"]
						}
					}
				}
				rules {
					condition {
						max_unique_count = "100"
						time_window {
							logs_unique_value_time_window_specific_value = "hours_36"
						}
					}
				}
				notification_payload_filter = ["obj.field"]
				max_unique_count_per_group_by_key = "100"
				unique_count_keypath = "obj.field"
			}
		}
	`, name, description, enabled, priority, typeVar, phantomMode, deleted)
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
