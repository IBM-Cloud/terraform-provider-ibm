// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.104.0-b4a47c49-20250418-184351
 */

package logs_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmLogsAlertDefinitionsDataSourceBasic(t *testing.T) {
	alertDefinitionName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	alertDefinitionType := "logs_immediate_or_unspecified"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmLogsAlertDefinitionsDataSourceConfigBasic(alertDefinitionName, alertDefinitionType),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_logs_alert_definitions.logs_alert_definitions_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_alert_definitions.logs_alert_definitions_instance", "alert_definitions.#"),
					resource.TestCheckResourceAttr("data.ibm_logs_alert_definitions.logs_alert_definitions_instance", "alert_definitions.0.name", alertDefinitionName),
					resource.TestCheckResourceAttr("data.ibm_logs_alert_definitions.logs_alert_definitions_instance", "alert_definitions.0.type", alertDefinitionType),
				),
			},
		},
	})
}

func TestAccIbmLogsAlertDefinitionsDataSourceAllArgs(t *testing.T) {
	alertDefinitionName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	alertDefinitionDescription := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	alertDefinitionEnabled := "false"
	alertDefinitionPriority := "p5_or_unspecified"
	alertDefinitionType := "logs_immediate_or_unspecified"
	alertDefinitionPhantomMode := "true"
	alertDefinitionDeleted := "false"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmLogsAlertDefinitionsDataSourceConfig(alertDefinitionName, alertDefinitionDescription, alertDefinitionEnabled, alertDefinitionPriority, alertDefinitionType, alertDefinitionPhantomMode, alertDefinitionDeleted),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_logs_alert_definitions.logs_alert_definitions_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_alert_definitions.logs_alert_definitions_instance", "alert_definitions.#"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_alert_definitions.logs_alert_definitions_instance", "alert_definitions.0.id"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_alert_definitions.logs_alert_definitions_instance", "alert_definitions.0.created_time"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_alert_definitions.logs_alert_definitions_instance", "alert_definitions.0.updated_time"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_alert_definitions.logs_alert_definitions_instance", "alert_definitions.0.alert_version_id"),
					resource.TestCheckResourceAttr("data.ibm_logs_alert_definitions.logs_alert_definitions_instance", "alert_definitions.0.name", alertDefinitionName),
					resource.TestCheckResourceAttr("data.ibm_logs_alert_definitions.logs_alert_definitions_instance", "alert_definitions.0.description", alertDefinitionDescription),
					resource.TestCheckResourceAttr("data.ibm_logs_alert_definitions.logs_alert_definitions_instance", "alert_definitions.0.enabled", alertDefinitionEnabled),
					resource.TestCheckResourceAttr("data.ibm_logs_alert_definitions.logs_alert_definitions_instance", "alert_definitions.0.priority", alertDefinitionPriority),
					resource.TestCheckResourceAttr("data.ibm_logs_alert_definitions.logs_alert_definitions_instance", "alert_definitions.0.type", alertDefinitionType),
					resource.TestCheckResourceAttr("data.ibm_logs_alert_definitions.logs_alert_definitions_instance", "alert_definitions.0.phantom_mode", alertDefinitionPhantomMode),
					resource.TestCheckResourceAttr("data.ibm_logs_alert_definitions.logs_alert_definitions_instance", "alert_definitions.0.deleted", alertDefinitionDeleted),
				),
			},
		},
	})
}

func testAccCheckIbmLogsAlertDefinitionsDataSourceConfigBasic(alertDefinitionName string, alertDefinitionType string) string {
	return fmt.Sprintf(`
		resource "ibm_logs_alert_definition" "logs_alert_definition_instance" {
			name = "%s"
			type = "%s"
			group_by_keys = "FIXME"
		}

		data "ibm_logs_alert_definitions" "logs_alert_definitions_instance" {
			depends_on = [
				ibm_logs_alert_definition.logs_alert_definition_instance
			]
		}
	`, alertDefinitionName, alertDefinitionType)
}

func testAccCheckIbmLogsAlertDefinitionsDataSourceConfig(alertDefinitionName string, alertDefinitionDescription string, alertDefinitionEnabled string, alertDefinitionPriority string, alertDefinitionType string, alertDefinitionPhantomMode string, alertDefinitionDeleted string) string {
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

		data "ibm_logs_alert_definitions" "logs_alert_definitions_instance" {
			depends_on = [
				ibm_logs_alert_definition.logs_alert_definition_instance
			]
		}
	`, alertDefinitionName, alertDefinitionDescription, alertDefinitionEnabled, alertDefinitionPriority, alertDefinitionType, alertDefinitionPhantomMode, alertDefinitionDeleted)
}
