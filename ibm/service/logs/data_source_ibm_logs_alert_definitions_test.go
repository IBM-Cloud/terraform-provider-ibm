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

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/logs"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/logs-go-sdk/logsv0"
	"github.com/stretchr/testify/assert"
	. "github.com/IBM-Cloud/terraform-provider-ibm/ibm/unittest"
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

func TestDataSourceIbmLogsAlertDefinitionsAlertDefinitionToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		apisAlertDefinitionTimeOfDayModel := make(map[string]interface{})
		apisAlertDefinitionTimeOfDayModel["hours"] = int(22)
		apisAlertDefinitionTimeOfDayModel["minutes"] = int(30)

		apisAlertDefinitionActivityScheduleModel := make(map[string]interface{})
		apisAlertDefinitionActivityScheduleModel["day_of_week"] = []string{"sunday", "monday_or_unspecified", "tuesday", "wednesday", "thursday", "friday", "saturday"}
		apisAlertDefinitionActivityScheduleModel["start_time"] = []map[string]interface{}{apisAlertDefinitionTimeOfDayModel}
		apisAlertDefinitionActivityScheduleModel["end_time"] = []map[string]interface{}{apisAlertDefinitionTimeOfDayModel}

		apisAlertDefinitionAlertDefIncidentSettingsModel := make(map[string]interface{})
		apisAlertDefinitionAlertDefIncidentSettingsModel["notify_on"] = "triggered_only_unspecified"
		apisAlertDefinitionAlertDefIncidentSettingsModel["minutes"] = int(10)

		apisAlertDefinitionIntegrationTypeModel := make(map[string]interface{})
		apisAlertDefinitionIntegrationTypeModel["integration_id"] = int(123)

		apisAlertDefinitionAlertDefWebhooksSettingsModel := make(map[string]interface{})
		apisAlertDefinitionAlertDefWebhooksSettingsModel["notify_on"] = "triggered_and_resolved"
		apisAlertDefinitionAlertDefWebhooksSettingsModel["integration"] = []map[string]interface{}{apisAlertDefinitionIntegrationTypeModel}
		apisAlertDefinitionAlertDefWebhooksSettingsModel["minutes"] = int(15)

		apisAlertDefinitionAlertDefNotificationGroupModel := make(map[string]interface{})
		apisAlertDefinitionAlertDefNotificationGroupModel["group_by_keys"] = []string{}
		apisAlertDefinitionAlertDefNotificationGroupModel["webhooks"] = []map[string]interface{}{apisAlertDefinitionAlertDefWebhooksSettingsModel}

		apisAlertDefinitionLabelFilterTypeModel := make(map[string]interface{})
		apisAlertDefinitionLabelFilterTypeModel["value"] = "my-app"
		apisAlertDefinitionLabelFilterTypeModel["operation"] = "starts_with"

		apisAlertDefinitionLabelFiltersModel := make(map[string]interface{})
		apisAlertDefinitionLabelFiltersModel["application_name"] = []map[string]interface{}{apisAlertDefinitionLabelFilterTypeModel}
		apisAlertDefinitionLabelFiltersModel["subsystem_name"] = []map[string]interface{}{apisAlertDefinitionLabelFilterTypeModel}
		apisAlertDefinitionLabelFiltersModel["severities"] = []string{}

		apisAlertDefinitionLogsSimpleFilterModel := make(map[string]interface{})
		apisAlertDefinitionLogsSimpleFilterModel["lucene_query"] = "text:"error""
		apisAlertDefinitionLogsSimpleFilterModel["label_filters"] = []map[string]interface{}{apisAlertDefinitionLabelFiltersModel}

		apisAlertDefinitionLogsFilterModel := make(map[string]interface{})
		apisAlertDefinitionLogsFilterModel["simple_filter"] = []map[string]interface{}{apisAlertDefinitionLogsSimpleFilterModel}

		apisAlertDefinitionUndetectedValuesManagementModel := make(map[string]interface{})
		apisAlertDefinitionUndetectedValuesManagementModel["trigger_undetected_values"] = true
		apisAlertDefinitionUndetectedValuesManagementModel["auto_retire_timeframe"] = "hours_24"

		apisAlertDefinitionLogsTimeWindowModel := make(map[string]interface{})
		apisAlertDefinitionLogsTimeWindowModel["logs_time_window_specific_value"] = "minutes_10"

		apisAlertDefinitionLogsThresholdConditionModel := make(map[string]interface{})
		apisAlertDefinitionLogsThresholdConditionModel["threshold"] = float64(1)
		apisAlertDefinitionLogsThresholdConditionModel["time_window"] = []map[string]interface{}{apisAlertDefinitionLogsTimeWindowModel}

		apisAlertDefinitionAlertDefOverrideModel := make(map[string]interface{})
		apisAlertDefinitionAlertDefOverrideModel["priority"] = "p5_or_unspecified"

		apisAlertDefinitionLogsThresholdRuleModel := make(map[string]interface{})
		apisAlertDefinitionLogsThresholdRuleModel["condition"] = []map[string]interface{}{apisAlertDefinitionLogsThresholdConditionModel}
		apisAlertDefinitionLogsThresholdRuleModel["override"] = []map[string]interface{}{apisAlertDefinitionAlertDefOverrideModel}

		apisAlertDefinitionLogsThresholdTypeModel := make(map[string]interface{})
		apisAlertDefinitionLogsThresholdTypeModel["logs_filter"] = []map[string]interface{}{apisAlertDefinitionLogsFilterModel}
		apisAlertDefinitionLogsThresholdTypeModel["undetected_values_management"] = []map[string]interface{}{apisAlertDefinitionUndetectedValuesManagementModel}
		apisAlertDefinitionLogsThresholdTypeModel["rules"] = []map[string]interface{}{apisAlertDefinitionLogsThresholdRuleModel}
		apisAlertDefinitionLogsThresholdTypeModel["condition_type"] = "more_than_or_unspecified"
		apisAlertDefinitionLogsThresholdTypeModel["notification_payload_filter"] = []string{}
		apisAlertDefinitionLogsThresholdTypeModel["evaluation_delay_ms"] = int(60000)

		model := make(map[string]interface{})
		model["id"] = "3dc02998-0b50-4ea8-b68a-4779d716fa1f"
		model["created_time"] = "2021-01-01T00:00:00.000Z"
		model["updated_time"] = "2021-01-01T00:00:00.000Z"
		model["alert_version_id"] = "3dc02998-0b50-4ea8-b68a-4779d716fa1f"
		model["name"] = "Unique count alert"
		model["description"] = "Example of unique count alert from terraform"
		model["enabled"] = true
		model["priority"] = "p1"
		model["active_on"] = []map[string]interface{}{apisAlertDefinitionActivityScheduleModel}
		model["type"] = "flow"
		model["group_by_keys"] = []string{"key1", "key2"}
		model["incidents_settings"] = []map[string]interface{}{apisAlertDefinitionAlertDefIncidentSettingsModel}
		model["notification_group"] = []map[string]interface{}{apisAlertDefinitionAlertDefNotificationGroupModel}
		model["entity_labels"] = map[string]interface{}{"key1": "testString"}
		model["phantom_mode"] = false
		model["deleted"] = false
		model["logs_immediate"] = []map[string]interface{}{apisAlertDefinitionLogsImmediateTypeModel}
		model["logs_threshold"] = []map[string]interface{}{apisAlertDefinitionLogsThresholdTypeModel}
		model["logs_ratio_threshold"] = []map[string]interface{}{apisAlertDefinitionLogsRatioThresholdTypeModel}
		model["logs_time_relative_threshold"] = []map[string]interface{}{apisAlertDefinitionLogsTimeRelativeThresholdTypeModel}
		model["metric_threshold"] = []map[string]interface{}{apisAlertDefinitionMetricThresholdTypeModel}
		model["flow"] = []map[string]interface{}{apisAlertDefinitionFlowTypeModel}
		model["logs_anomaly"] = []map[string]interface{}{apisAlertDefinitionLogsAnomalyTypeModel}
		model["metric_anomaly"] = []map[string]interface{}{apisAlertDefinitionMetricAnomalyTypeModel}
		model["logs_new_value"] = []map[string]interface{}{apisAlertDefinitionLogsNewValueTypeModel}
		model["logs_unique_count"] = []map[string]interface{}{apisAlertDefinitionLogsUniqueCountTypeModel}

		assert.Equal(t, result, model)
	}

	apisAlertDefinitionTimeOfDayModel := new(logsv0.ApisAlertDefinitionTimeOfDay)
	apisAlertDefinitionTimeOfDayModel.Hours = core.Int64Ptr(int64(22))
	apisAlertDefinitionTimeOfDayModel.Minutes = core.Int64Ptr(int64(30))

	apisAlertDefinitionActivityScheduleModel := new(logsv0.ApisAlertDefinitionActivitySchedule)
	apisAlertDefinitionActivityScheduleModel.DayOfWeek = []string{"sunday", "monday_or_unspecified", "tuesday", "wednesday", "thursday", "friday", "saturday"}
	apisAlertDefinitionActivityScheduleModel.StartTime = apisAlertDefinitionTimeOfDayModel
	apisAlertDefinitionActivityScheduleModel.EndTime = apisAlertDefinitionTimeOfDayModel

	apisAlertDefinitionAlertDefIncidentSettingsModel := new(logsv0.ApisAlertDefinitionAlertDefIncidentSettings)
	apisAlertDefinitionAlertDefIncidentSettingsModel.NotifyOn = core.StringPtr("triggered_only_unspecified")
	apisAlertDefinitionAlertDefIncidentSettingsModel.Minutes = core.Int64Ptr(int64(10))

	apisAlertDefinitionIntegrationTypeModel := new(logsv0.ApisAlertDefinitionIntegrationTypeIntegrationTypeIntegrationID)
	apisAlertDefinitionIntegrationTypeModel.IntegrationID = core.Int64Ptr(int64(123))

	apisAlertDefinitionAlertDefWebhooksSettingsModel := new(logsv0.ApisAlertDefinitionAlertDefWebhooksSettings)
	apisAlertDefinitionAlertDefWebhooksSettingsModel.NotifyOn = core.StringPtr("triggered_and_resolved")
	apisAlertDefinitionAlertDefWebhooksSettingsModel.Integration = apisAlertDefinitionIntegrationTypeModel
	apisAlertDefinitionAlertDefWebhooksSettingsModel.Minutes = core.Int64Ptr(int64(15))

	apisAlertDefinitionAlertDefNotificationGroupModel := new(logsv0.ApisAlertDefinitionAlertDefNotificationGroup)
	apisAlertDefinitionAlertDefNotificationGroupModel.GroupByKeys = []string{}
	apisAlertDefinitionAlertDefNotificationGroupModel.Webhooks = []logsv0.ApisAlertDefinitionAlertDefWebhooksSettings{*apisAlertDefinitionAlertDefWebhooksSettingsModel}

	apisAlertDefinitionLabelFilterTypeModel := new(logsv0.ApisAlertDefinitionLabelFilterType)
	apisAlertDefinitionLabelFilterTypeModel.Value = core.StringPtr("my-app")
	apisAlertDefinitionLabelFilterTypeModel.Operation = core.StringPtr("starts_with")

	apisAlertDefinitionLabelFiltersModel := new(logsv0.ApisAlertDefinitionLabelFilters)
	apisAlertDefinitionLabelFiltersModel.ApplicationName = []logsv0.ApisAlertDefinitionLabelFilterType{*apisAlertDefinitionLabelFilterTypeModel}
	apisAlertDefinitionLabelFiltersModel.SubsystemName = []logsv0.ApisAlertDefinitionLabelFilterType{*apisAlertDefinitionLabelFilterTypeModel}
	apisAlertDefinitionLabelFiltersModel.Severities = []string{}

	apisAlertDefinitionLogsSimpleFilterModel := new(logsv0.ApisAlertDefinitionLogsSimpleFilter)
	apisAlertDefinitionLogsSimpleFilterModel.LuceneQuery = core.StringPtr("text:"error"")
	apisAlertDefinitionLogsSimpleFilterModel.LabelFilters = apisAlertDefinitionLabelFiltersModel

	apisAlertDefinitionLogsFilterModel := new(logsv0.ApisAlertDefinitionLogsFilter)
	apisAlertDefinitionLogsFilterModel.SimpleFilter = apisAlertDefinitionLogsSimpleFilterModel

	apisAlertDefinitionUndetectedValuesManagementModel := new(logsv0.ApisAlertDefinitionUndetectedValuesManagement)
	apisAlertDefinitionUndetectedValuesManagementModel.TriggerUndetectedValues = core.BoolPtr(true)
	apisAlertDefinitionUndetectedValuesManagementModel.AutoRetireTimeframe = core.StringPtr("hours_24")

	apisAlertDefinitionLogsTimeWindowModel := new(logsv0.ApisAlertDefinitionLogsTimeWindow)
	apisAlertDefinitionLogsTimeWindowModel.LogsTimeWindowSpecificValue = core.StringPtr("minutes_10")

	apisAlertDefinitionLogsThresholdConditionModel := new(logsv0.ApisAlertDefinitionLogsThresholdCondition)
	apisAlertDefinitionLogsThresholdConditionModel.Threshold = core.Float64Ptr(float64(1))
	apisAlertDefinitionLogsThresholdConditionModel.TimeWindow = apisAlertDefinitionLogsTimeWindowModel

	apisAlertDefinitionAlertDefOverrideModel := new(logsv0.ApisAlertDefinitionAlertDefOverride)
	apisAlertDefinitionAlertDefOverrideModel.Priority = core.StringPtr("p5_or_unspecified")

	apisAlertDefinitionLogsThresholdRuleModel := new(logsv0.ApisAlertDefinitionLogsThresholdRule)
	apisAlertDefinitionLogsThresholdRuleModel.Condition = apisAlertDefinitionLogsThresholdConditionModel
	apisAlertDefinitionLogsThresholdRuleModel.Override = apisAlertDefinitionAlertDefOverrideModel

	apisAlertDefinitionLogsThresholdTypeModel := new(logsv0.ApisAlertDefinitionLogsThresholdType)
	apisAlertDefinitionLogsThresholdTypeModel.LogsFilter = apisAlertDefinitionLogsFilterModel
	apisAlertDefinitionLogsThresholdTypeModel.UndetectedValuesManagement = apisAlertDefinitionUndetectedValuesManagementModel
	apisAlertDefinitionLogsThresholdTypeModel.Rules = []logsv0.ApisAlertDefinitionLogsThresholdRule{*apisAlertDefinitionLogsThresholdRuleModel}
	apisAlertDefinitionLogsThresholdTypeModel.ConditionType = core.StringPtr("more_than_or_unspecified")
	apisAlertDefinitionLogsThresholdTypeModel.NotificationPayloadFilter = []string{}
	apisAlertDefinitionLogsThresholdTypeModel.EvaluationDelayMs = core.Int64Ptr(int64(60000))

	model := new(logsv0.AlertDefinition)
	model.ID = CreateMockUUID("3dc02998-0b50-4ea8-b68a-4779d716fa1f")
	model.CreatedTime = CreateMockDateTime("2021-01-01T00:00:00.000Z")
	model.UpdatedTime = CreateMockDateTime("2021-01-01T00:00:00.000Z")
	model.AlertVersionID = CreateMockUUID("3dc02998-0b50-4ea8-b68a-4779d716fa1f")
	model.Name = core.StringPtr("Unique count alert")
	model.Description = core.StringPtr("Example of unique count alert from terraform")
	model.Enabled = core.BoolPtr(true)
	model.Priority = core.StringPtr("p1")
	model.ActiveOn = apisAlertDefinitionActivityScheduleModel
	model.Type = core.StringPtr("flow")
	model.GroupByKeys = []string{"key1", "key2"}
	model.IncidentsSettings = apisAlertDefinitionAlertDefIncidentSettingsModel
	model.NotificationGroup = apisAlertDefinitionAlertDefNotificationGroupModel
	model.EntityLabels = map[string]string{"key1": "testString"}
	model.PhantomMode = core.BoolPtr(false)
	model.Deleted = core.BoolPtr(false)
	model.LogsImmediate = apisAlertDefinitionLogsImmediateTypeModel
	model.LogsThreshold = apisAlertDefinitionLogsThresholdTypeModel
	model.LogsRatioThreshold = apisAlertDefinitionLogsRatioThresholdTypeModel
	model.LogsTimeRelativeThreshold = apisAlertDefinitionLogsTimeRelativeThresholdTypeModel
	model.MetricThreshold = apisAlertDefinitionMetricThresholdTypeModel
	model.Flow = apisAlertDefinitionFlowTypeModel
	model.LogsAnomaly = apisAlertDefinitionLogsAnomalyTypeModel
	model.MetricAnomaly = apisAlertDefinitionMetricAnomalyTypeModel
	model.LogsNewValue = apisAlertDefinitionLogsNewValueTypeModel
	model.LogsUniqueCount = apisAlertDefinitionLogsUniqueCountTypeModel

	result, err := logs.DataSourceIbmLogsAlertDefinitionsAlertDefinitionToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmLogsAlertDefinitionsApisAlertDefinitionActivityScheduleToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		apisAlertDefinitionTimeOfDayModel := make(map[string]interface{})
		apisAlertDefinitionTimeOfDayModel["hours"] = int(14)
		apisAlertDefinitionTimeOfDayModel["minutes"] = int(30)

		model := make(map[string]interface{})
		model["day_of_week"] = []string{"sunday"}
		model["start_time"] = []map[string]interface{}{apisAlertDefinitionTimeOfDayModel}
		model["end_time"] = []map[string]interface{}{apisAlertDefinitionTimeOfDayModel}

		assert.Equal(t, result, model)
	}

	apisAlertDefinitionTimeOfDayModel := new(logsv0.ApisAlertDefinitionTimeOfDay)
	apisAlertDefinitionTimeOfDayModel.Hours = core.Int64Ptr(int64(14))
	apisAlertDefinitionTimeOfDayModel.Minutes = core.Int64Ptr(int64(30))

	model := new(logsv0.ApisAlertDefinitionActivitySchedule)
	model.DayOfWeek = []string{"sunday"}
	model.StartTime = apisAlertDefinitionTimeOfDayModel
	model.EndTime = apisAlertDefinitionTimeOfDayModel

	result, err := logs.DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionActivityScheduleToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmLogsAlertDefinitionsApisAlertDefinitionTimeOfDayToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["hours"] = int(14)
		model["minutes"] = int(30)

		assert.Equal(t, result, model)
	}

	model := new(logsv0.ApisAlertDefinitionTimeOfDay)
	model.Hours = core.Int64Ptr(int64(14))
	model.Minutes = core.Int64Ptr(int64(30))

	result, err := logs.DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionTimeOfDayToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmLogsAlertDefinitionsApisAlertDefinitionAlertDefIncidentSettingsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["notify_on"] = "triggered_and_resolved"
		model["minutes"] = int(30)

		assert.Equal(t, result, model)
	}

	model := new(logsv0.ApisAlertDefinitionAlertDefIncidentSettings)
	model.NotifyOn = core.StringPtr("triggered_and_resolved")
	model.Minutes = core.Int64Ptr(int64(30))

	result, err := logs.DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionAlertDefIncidentSettingsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmLogsAlertDefinitionsApisAlertDefinitionAlertDefNotificationGroupToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		apisAlertDefinitionIntegrationTypeModel := make(map[string]interface{})
		apisAlertDefinitionIntegrationTypeModel["integration_id"] = int(123)

		apisAlertDefinitionAlertDefWebhooksSettingsModel := make(map[string]interface{})
		apisAlertDefinitionAlertDefWebhooksSettingsModel["notify_on"] = "triggered_and_resolved"
		apisAlertDefinitionAlertDefWebhooksSettingsModel["integration"] = []map[string]interface{}{apisAlertDefinitionIntegrationTypeModel}
		apisAlertDefinitionAlertDefWebhooksSettingsModel["minutes"] = int(15)

		model := make(map[string]interface{})
		model["group_by_keys"] = []string{"key1", "key2"}
		model["webhooks"] = []map[string]interface{}{apisAlertDefinitionAlertDefWebhooksSettingsModel}

		assert.Equal(t, result, model)
	}

	apisAlertDefinitionIntegrationTypeModel := new(logsv0.ApisAlertDefinitionIntegrationTypeIntegrationTypeIntegrationID)
	apisAlertDefinitionIntegrationTypeModel.IntegrationID = core.Int64Ptr(int64(123))

	apisAlertDefinitionAlertDefWebhooksSettingsModel := new(logsv0.ApisAlertDefinitionAlertDefWebhooksSettings)
	apisAlertDefinitionAlertDefWebhooksSettingsModel.NotifyOn = core.StringPtr("triggered_and_resolved")
	apisAlertDefinitionAlertDefWebhooksSettingsModel.Integration = apisAlertDefinitionIntegrationTypeModel
	apisAlertDefinitionAlertDefWebhooksSettingsModel.Minutes = core.Int64Ptr(int64(15))

	model := new(logsv0.ApisAlertDefinitionAlertDefNotificationGroup)
	model.GroupByKeys = []string{"key1", "key2"}
	model.Webhooks = []logsv0.ApisAlertDefinitionAlertDefWebhooksSettings{*apisAlertDefinitionAlertDefWebhooksSettingsModel}

	result, err := logs.DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionAlertDefNotificationGroupToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmLogsAlertDefinitionsApisAlertDefinitionAlertDefWebhooksSettingsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		apisAlertDefinitionIntegrationTypeModel := make(map[string]interface{})
		apisAlertDefinitionIntegrationTypeModel["integration_id"] = int(123)

		model := make(map[string]interface{})
		model["notify_on"] = "triggered_and_resolved"
		model["integration"] = []map[string]interface{}{apisAlertDefinitionIntegrationTypeModel}
		model["minutes"] = int(15)

		assert.Equal(t, result, model)
	}

	apisAlertDefinitionIntegrationTypeModel := new(logsv0.ApisAlertDefinitionIntegrationTypeIntegrationTypeIntegrationID)
	apisAlertDefinitionIntegrationTypeModel.IntegrationID = core.Int64Ptr(int64(123))

	model := new(logsv0.ApisAlertDefinitionAlertDefWebhooksSettings)
	model.NotifyOn = core.StringPtr("triggered_and_resolved")
	model.Integration = apisAlertDefinitionIntegrationTypeModel
	model.Minutes = core.Int64Ptr(int64(15))

	result, err := logs.DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionAlertDefWebhooksSettingsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmLogsAlertDefinitionsApisAlertDefinitionIntegrationTypeToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["integration_id"] = int(123)

		assert.Equal(t, result, model)
	}

	model := new(logsv0.ApisAlertDefinitionIntegrationType)
	model.IntegrationID = core.Int64Ptr(int64(123))

	result, err := logs.DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionIntegrationTypeToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmLogsAlertDefinitionsApisAlertDefinitionIntegrationTypeIntegrationTypeIntegrationIDToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["integration_id"] = int(123)

		assert.Equal(t, result, model)
	}

	model := new(logsv0.ApisAlertDefinitionIntegrationTypeIntegrationTypeIntegrationID)
	model.IntegrationID = core.Int64Ptr(int64(123))

	result, err := logs.DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionIntegrationTypeIntegrationTypeIntegrationIDToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmLogsAlertDefinitionsApisAlertDefinitionLogsImmediateTypeToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		apisAlertDefinitionLabelFilterTypeModel := make(map[string]interface{})
		apisAlertDefinitionLabelFilterTypeModel["value"] = "my-app"
		apisAlertDefinitionLabelFilterTypeModel["operation"] = "starts_with"

		apisAlertDefinitionLabelFiltersModel := make(map[string]interface{})
		apisAlertDefinitionLabelFiltersModel["application_name"] = []map[string]interface{}{apisAlertDefinitionLabelFilterTypeModel}
		apisAlertDefinitionLabelFiltersModel["subsystem_name"] = []map[string]interface{}{apisAlertDefinitionLabelFilterTypeModel}
		apisAlertDefinitionLabelFiltersModel["severities"] = []string{"critical"}

		apisAlertDefinitionLogsSimpleFilterModel := make(map[string]interface{})
		apisAlertDefinitionLogsSimpleFilterModel["lucene_query"] = "text:"error""
		apisAlertDefinitionLogsSimpleFilterModel["label_filters"] = []map[string]interface{}{apisAlertDefinitionLabelFiltersModel}

		apisAlertDefinitionLogsFilterModel := make(map[string]interface{})
		apisAlertDefinitionLogsFilterModel["simple_filter"] = []map[string]interface{}{apisAlertDefinitionLogsSimpleFilterModel}

		model := make(map[string]interface{})
		model["logs_filter"] = []map[string]interface{}{apisAlertDefinitionLogsFilterModel}
		model["notification_payload_filter"] = []string{"obj.field"}

		assert.Equal(t, result, model)
	}

	apisAlertDefinitionLabelFilterTypeModel := new(logsv0.ApisAlertDefinitionLabelFilterType)
	apisAlertDefinitionLabelFilterTypeModel.Value = core.StringPtr("my-app")
	apisAlertDefinitionLabelFilterTypeModel.Operation = core.StringPtr("starts_with")

	apisAlertDefinitionLabelFiltersModel := new(logsv0.ApisAlertDefinitionLabelFilters)
	apisAlertDefinitionLabelFiltersModel.ApplicationName = []logsv0.ApisAlertDefinitionLabelFilterType{*apisAlertDefinitionLabelFilterTypeModel}
	apisAlertDefinitionLabelFiltersModel.SubsystemName = []logsv0.ApisAlertDefinitionLabelFilterType{*apisAlertDefinitionLabelFilterTypeModel}
	apisAlertDefinitionLabelFiltersModel.Severities = []string{"critical"}

	apisAlertDefinitionLogsSimpleFilterModel := new(logsv0.ApisAlertDefinitionLogsSimpleFilter)
	apisAlertDefinitionLogsSimpleFilterModel.LuceneQuery = core.StringPtr("text:"error"")
	apisAlertDefinitionLogsSimpleFilterModel.LabelFilters = apisAlertDefinitionLabelFiltersModel

	apisAlertDefinitionLogsFilterModel := new(logsv0.ApisAlertDefinitionLogsFilter)
	apisAlertDefinitionLogsFilterModel.SimpleFilter = apisAlertDefinitionLogsSimpleFilterModel

	model := new(logsv0.ApisAlertDefinitionLogsImmediateType)
	model.LogsFilter = apisAlertDefinitionLogsFilterModel
	model.NotificationPayloadFilter = []string{"obj.field"}

	result, err := logs.DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionLogsImmediateTypeToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmLogsAlertDefinitionsApisAlertDefinitionLogsFilterToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		apisAlertDefinitionLabelFilterTypeModel := make(map[string]interface{})
		apisAlertDefinitionLabelFilterTypeModel["value"] = "my-app"
		apisAlertDefinitionLabelFilterTypeModel["operation"] = "starts_with"

		apisAlertDefinitionLabelFiltersModel := make(map[string]interface{})
		apisAlertDefinitionLabelFiltersModel["application_name"] = []map[string]interface{}{apisAlertDefinitionLabelFilterTypeModel}
		apisAlertDefinitionLabelFiltersModel["subsystem_name"] = []map[string]interface{}{apisAlertDefinitionLabelFilterTypeModel}
		apisAlertDefinitionLabelFiltersModel["severities"] = []string{"critical"}

		apisAlertDefinitionLogsSimpleFilterModel := make(map[string]interface{})
		apisAlertDefinitionLogsSimpleFilterModel["lucene_query"] = "text:"error""
		apisAlertDefinitionLogsSimpleFilterModel["label_filters"] = []map[string]interface{}{apisAlertDefinitionLabelFiltersModel}

		model := make(map[string]interface{})
		model["simple_filter"] = []map[string]interface{}{apisAlertDefinitionLogsSimpleFilterModel}

		assert.Equal(t, result, model)
	}

	apisAlertDefinitionLabelFilterTypeModel := new(logsv0.ApisAlertDefinitionLabelFilterType)
	apisAlertDefinitionLabelFilterTypeModel.Value = core.StringPtr("my-app")
	apisAlertDefinitionLabelFilterTypeModel.Operation = core.StringPtr("starts_with")

	apisAlertDefinitionLabelFiltersModel := new(logsv0.ApisAlertDefinitionLabelFilters)
	apisAlertDefinitionLabelFiltersModel.ApplicationName = []logsv0.ApisAlertDefinitionLabelFilterType{*apisAlertDefinitionLabelFilterTypeModel}
	apisAlertDefinitionLabelFiltersModel.SubsystemName = []logsv0.ApisAlertDefinitionLabelFilterType{*apisAlertDefinitionLabelFilterTypeModel}
	apisAlertDefinitionLabelFiltersModel.Severities = []string{"critical"}

	apisAlertDefinitionLogsSimpleFilterModel := new(logsv0.ApisAlertDefinitionLogsSimpleFilter)
	apisAlertDefinitionLogsSimpleFilterModel.LuceneQuery = core.StringPtr("text:"error"")
	apisAlertDefinitionLogsSimpleFilterModel.LabelFilters = apisAlertDefinitionLabelFiltersModel

	model := new(logsv0.ApisAlertDefinitionLogsFilter)
	model.SimpleFilter = apisAlertDefinitionLogsSimpleFilterModel

	result, err := logs.DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionLogsFilterToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmLogsAlertDefinitionsApisAlertDefinitionLogsSimpleFilterToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		apisAlertDefinitionLabelFilterTypeModel := make(map[string]interface{})
		apisAlertDefinitionLabelFilterTypeModel["value"] = "my-app"
		apisAlertDefinitionLabelFilterTypeModel["operation"] = "starts_with"

		apisAlertDefinitionLabelFiltersModel := make(map[string]interface{})
		apisAlertDefinitionLabelFiltersModel["application_name"] = []map[string]interface{}{apisAlertDefinitionLabelFilterTypeModel}
		apisAlertDefinitionLabelFiltersModel["subsystem_name"] = []map[string]interface{}{apisAlertDefinitionLabelFilterTypeModel}
		apisAlertDefinitionLabelFiltersModel["severities"] = []string{"critical"}

		model := make(map[string]interface{})
		model["lucene_query"] = "text:"error""
		model["label_filters"] = []map[string]interface{}{apisAlertDefinitionLabelFiltersModel}

		assert.Equal(t, result, model)
	}

	apisAlertDefinitionLabelFilterTypeModel := new(logsv0.ApisAlertDefinitionLabelFilterType)
	apisAlertDefinitionLabelFilterTypeModel.Value = core.StringPtr("my-app")
	apisAlertDefinitionLabelFilterTypeModel.Operation = core.StringPtr("starts_with")

	apisAlertDefinitionLabelFiltersModel := new(logsv0.ApisAlertDefinitionLabelFilters)
	apisAlertDefinitionLabelFiltersModel.ApplicationName = []logsv0.ApisAlertDefinitionLabelFilterType{*apisAlertDefinitionLabelFilterTypeModel}
	apisAlertDefinitionLabelFiltersModel.SubsystemName = []logsv0.ApisAlertDefinitionLabelFilterType{*apisAlertDefinitionLabelFilterTypeModel}
	apisAlertDefinitionLabelFiltersModel.Severities = []string{"critical"}

	model := new(logsv0.ApisAlertDefinitionLogsSimpleFilter)
	model.LuceneQuery = core.StringPtr("text:"error"")
	model.LabelFilters = apisAlertDefinitionLabelFiltersModel

	result, err := logs.DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionLogsSimpleFilterToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmLogsAlertDefinitionsApisAlertDefinitionLabelFiltersToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		apisAlertDefinitionLabelFilterTypeModel := make(map[string]interface{})
		apisAlertDefinitionLabelFilterTypeModel["value"] = "my-app"
		apisAlertDefinitionLabelFilterTypeModel["operation"] = "starts_with"

		model := make(map[string]interface{})
		model["application_name"] = []map[string]interface{}{apisAlertDefinitionLabelFilterTypeModel}
		model["subsystem_name"] = []map[string]interface{}{apisAlertDefinitionLabelFilterTypeModel}
		model["severities"] = []string{"critical"}

		assert.Equal(t, result, model)
	}

	apisAlertDefinitionLabelFilterTypeModel := new(logsv0.ApisAlertDefinitionLabelFilterType)
	apisAlertDefinitionLabelFilterTypeModel.Value = core.StringPtr("my-app")
	apisAlertDefinitionLabelFilterTypeModel.Operation = core.StringPtr("starts_with")

	model := new(logsv0.ApisAlertDefinitionLabelFilters)
	model.ApplicationName = []logsv0.ApisAlertDefinitionLabelFilterType{*apisAlertDefinitionLabelFilterTypeModel}
	model.SubsystemName = []logsv0.ApisAlertDefinitionLabelFilterType{*apisAlertDefinitionLabelFilterTypeModel}
	model.Severities = []string{"critical"}

	result, err := logs.DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionLabelFiltersToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmLogsAlertDefinitionsApisAlertDefinitionLabelFilterTypeToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["value"] = "my-app"
		model["operation"] = "starts_with"

		assert.Equal(t, result, model)
	}

	model := new(logsv0.ApisAlertDefinitionLabelFilterType)
	model.Value = core.StringPtr("my-app")
	model.Operation = core.StringPtr("starts_with")

	result, err := logs.DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionLabelFilterTypeToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmLogsAlertDefinitionsApisAlertDefinitionLogsThresholdTypeToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		apisAlertDefinitionLabelFilterTypeModel := make(map[string]interface{})
		apisAlertDefinitionLabelFilterTypeModel["value"] = "my-app"
		apisAlertDefinitionLabelFilterTypeModel["operation"] = "starts_with"

		apisAlertDefinitionLabelFiltersModel := make(map[string]interface{})
		apisAlertDefinitionLabelFiltersModel["application_name"] = []map[string]interface{}{apisAlertDefinitionLabelFilterTypeModel}
		apisAlertDefinitionLabelFiltersModel["subsystem_name"] = []map[string]interface{}{apisAlertDefinitionLabelFilterTypeModel}
		apisAlertDefinitionLabelFiltersModel["severities"] = []string{"critical"}

		apisAlertDefinitionLogsSimpleFilterModel := make(map[string]interface{})
		apisAlertDefinitionLogsSimpleFilterModel["lucene_query"] = "text:"error""
		apisAlertDefinitionLogsSimpleFilterModel["label_filters"] = []map[string]interface{}{apisAlertDefinitionLabelFiltersModel}

		apisAlertDefinitionLogsFilterModel := make(map[string]interface{})
		apisAlertDefinitionLogsFilterModel["simple_filter"] = []map[string]interface{}{apisAlertDefinitionLogsSimpleFilterModel}

		apisAlertDefinitionUndetectedValuesManagementModel := make(map[string]interface{})
		apisAlertDefinitionUndetectedValuesManagementModel["trigger_undetected_values"] = true
		apisAlertDefinitionUndetectedValuesManagementModel["auto_retire_timeframe"] = "hours_24"

		apisAlertDefinitionLogsTimeWindowModel := make(map[string]interface{})
		apisAlertDefinitionLogsTimeWindowModel["logs_time_window_specific_value"] = "hours_36"

		apisAlertDefinitionLogsThresholdConditionModel := make(map[string]interface{})
		apisAlertDefinitionLogsThresholdConditionModel["threshold"] = float64(100.0)
		apisAlertDefinitionLogsThresholdConditionModel["time_window"] = []map[string]interface{}{apisAlertDefinitionLogsTimeWindowModel}

		apisAlertDefinitionAlertDefOverrideModel := make(map[string]interface{})
		apisAlertDefinitionAlertDefOverrideModel["priority"] = "p1"

		apisAlertDefinitionLogsThresholdRuleModel := make(map[string]interface{})
		apisAlertDefinitionLogsThresholdRuleModel["condition"] = []map[string]interface{}{apisAlertDefinitionLogsThresholdConditionModel}
		apisAlertDefinitionLogsThresholdRuleModel["override"] = []map[string]interface{}{apisAlertDefinitionAlertDefOverrideModel}

		model := make(map[string]interface{})
		model["logs_filter"] = []map[string]interface{}{apisAlertDefinitionLogsFilterModel}
		model["undetected_values_management"] = []map[string]interface{}{apisAlertDefinitionUndetectedValuesManagementModel}
		model["rules"] = []map[string]interface{}{apisAlertDefinitionLogsThresholdRuleModel}
		model["condition_type"] = "less_than"
		model["notification_payload_filter"] = []string{"obj.field"}
		model["evaluation_delay_ms"] = int(60000)

		assert.Equal(t, result, model)
	}

	apisAlertDefinitionLabelFilterTypeModel := new(logsv0.ApisAlertDefinitionLabelFilterType)
	apisAlertDefinitionLabelFilterTypeModel.Value = core.StringPtr("my-app")
	apisAlertDefinitionLabelFilterTypeModel.Operation = core.StringPtr("starts_with")

	apisAlertDefinitionLabelFiltersModel := new(logsv0.ApisAlertDefinitionLabelFilters)
	apisAlertDefinitionLabelFiltersModel.ApplicationName = []logsv0.ApisAlertDefinitionLabelFilterType{*apisAlertDefinitionLabelFilterTypeModel}
	apisAlertDefinitionLabelFiltersModel.SubsystemName = []logsv0.ApisAlertDefinitionLabelFilterType{*apisAlertDefinitionLabelFilterTypeModel}
	apisAlertDefinitionLabelFiltersModel.Severities = []string{"critical"}

	apisAlertDefinitionLogsSimpleFilterModel := new(logsv0.ApisAlertDefinitionLogsSimpleFilter)
	apisAlertDefinitionLogsSimpleFilterModel.LuceneQuery = core.StringPtr("text:"error"")
	apisAlertDefinitionLogsSimpleFilterModel.LabelFilters = apisAlertDefinitionLabelFiltersModel

	apisAlertDefinitionLogsFilterModel := new(logsv0.ApisAlertDefinitionLogsFilter)
	apisAlertDefinitionLogsFilterModel.SimpleFilter = apisAlertDefinitionLogsSimpleFilterModel

	apisAlertDefinitionUndetectedValuesManagementModel := new(logsv0.ApisAlertDefinitionUndetectedValuesManagement)
	apisAlertDefinitionUndetectedValuesManagementModel.TriggerUndetectedValues = core.BoolPtr(true)
	apisAlertDefinitionUndetectedValuesManagementModel.AutoRetireTimeframe = core.StringPtr("hours_24")

	apisAlertDefinitionLogsTimeWindowModel := new(logsv0.ApisAlertDefinitionLogsTimeWindow)
	apisAlertDefinitionLogsTimeWindowModel.LogsTimeWindowSpecificValue = core.StringPtr("hours_36")

	apisAlertDefinitionLogsThresholdConditionModel := new(logsv0.ApisAlertDefinitionLogsThresholdCondition)
	apisAlertDefinitionLogsThresholdConditionModel.Threshold = core.Float64Ptr(float64(100.0))
	apisAlertDefinitionLogsThresholdConditionModel.TimeWindow = apisAlertDefinitionLogsTimeWindowModel

	apisAlertDefinitionAlertDefOverrideModel := new(logsv0.ApisAlertDefinitionAlertDefOverride)
	apisAlertDefinitionAlertDefOverrideModel.Priority = core.StringPtr("p1")

	apisAlertDefinitionLogsThresholdRuleModel := new(logsv0.ApisAlertDefinitionLogsThresholdRule)
	apisAlertDefinitionLogsThresholdRuleModel.Condition = apisAlertDefinitionLogsThresholdConditionModel
	apisAlertDefinitionLogsThresholdRuleModel.Override = apisAlertDefinitionAlertDefOverrideModel

	model := new(logsv0.ApisAlertDefinitionLogsThresholdType)
	model.LogsFilter = apisAlertDefinitionLogsFilterModel
	model.UndetectedValuesManagement = apisAlertDefinitionUndetectedValuesManagementModel
	model.Rules = []logsv0.ApisAlertDefinitionLogsThresholdRule{*apisAlertDefinitionLogsThresholdRuleModel}
	model.ConditionType = core.StringPtr("less_than")
	model.NotificationPayloadFilter = []string{"obj.field"}
	model.EvaluationDelayMs = core.Int64Ptr(int64(60000))

	result, err := logs.DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionLogsThresholdTypeToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmLogsAlertDefinitionsApisAlertDefinitionUndetectedValuesManagementToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["trigger_undetected_values"] = true
		model["auto_retire_timeframe"] = "hours_24"

		assert.Equal(t, result, model)
	}

	model := new(logsv0.ApisAlertDefinitionUndetectedValuesManagement)
	model.TriggerUndetectedValues = core.BoolPtr(true)
	model.AutoRetireTimeframe = core.StringPtr("hours_24")

	result, err := logs.DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionUndetectedValuesManagementToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmLogsAlertDefinitionsApisAlertDefinitionLogsThresholdRuleToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		apisAlertDefinitionLogsTimeWindowModel := make(map[string]interface{})
		apisAlertDefinitionLogsTimeWindowModel["logs_time_window_specific_value"] = "hours_36"

		apisAlertDefinitionLogsThresholdConditionModel := make(map[string]interface{})
		apisAlertDefinitionLogsThresholdConditionModel["threshold"] = float64(100.0)
		apisAlertDefinitionLogsThresholdConditionModel["time_window"] = []map[string]interface{}{apisAlertDefinitionLogsTimeWindowModel}

		apisAlertDefinitionAlertDefOverrideModel := make(map[string]interface{})
		apisAlertDefinitionAlertDefOverrideModel["priority"] = "p1"

		model := make(map[string]interface{})
		model["condition"] = []map[string]interface{}{apisAlertDefinitionLogsThresholdConditionModel}
		model["override"] = []map[string]interface{}{apisAlertDefinitionAlertDefOverrideModel}

		assert.Equal(t, result, model)
	}

	apisAlertDefinitionLogsTimeWindowModel := new(logsv0.ApisAlertDefinitionLogsTimeWindow)
	apisAlertDefinitionLogsTimeWindowModel.LogsTimeWindowSpecificValue = core.StringPtr("hours_36")

	apisAlertDefinitionLogsThresholdConditionModel := new(logsv0.ApisAlertDefinitionLogsThresholdCondition)
	apisAlertDefinitionLogsThresholdConditionModel.Threshold = core.Float64Ptr(float64(100.0))
	apisAlertDefinitionLogsThresholdConditionModel.TimeWindow = apisAlertDefinitionLogsTimeWindowModel

	apisAlertDefinitionAlertDefOverrideModel := new(logsv0.ApisAlertDefinitionAlertDefOverride)
	apisAlertDefinitionAlertDefOverrideModel.Priority = core.StringPtr("p1")

	model := new(logsv0.ApisAlertDefinitionLogsThresholdRule)
	model.Condition = apisAlertDefinitionLogsThresholdConditionModel
	model.Override = apisAlertDefinitionAlertDefOverrideModel

	result, err := logs.DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionLogsThresholdRuleToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmLogsAlertDefinitionsApisAlertDefinitionLogsThresholdConditionToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		apisAlertDefinitionLogsTimeWindowModel := make(map[string]interface{})
		apisAlertDefinitionLogsTimeWindowModel["logs_time_window_specific_value"] = "hours_36"

		model := make(map[string]interface{})
		model["threshold"] = float64(100.0)
		model["time_window"] = []map[string]interface{}{apisAlertDefinitionLogsTimeWindowModel}

		assert.Equal(t, result, model)
	}

	apisAlertDefinitionLogsTimeWindowModel := new(logsv0.ApisAlertDefinitionLogsTimeWindow)
	apisAlertDefinitionLogsTimeWindowModel.LogsTimeWindowSpecificValue = core.StringPtr("hours_36")

	model := new(logsv0.ApisAlertDefinitionLogsThresholdCondition)
	model.Threshold = core.Float64Ptr(float64(100.0))
	model.TimeWindow = apisAlertDefinitionLogsTimeWindowModel

	result, err := logs.DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionLogsThresholdConditionToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmLogsAlertDefinitionsApisAlertDefinitionLogsTimeWindowToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["logs_time_window_specific_value"] = "hours_36"

		assert.Equal(t, result, model)
	}

	model := new(logsv0.ApisAlertDefinitionLogsTimeWindow)
	model.LogsTimeWindowSpecificValue = core.StringPtr("hours_36")

	result, err := logs.DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionLogsTimeWindowToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmLogsAlertDefinitionsApisAlertDefinitionAlertDefOverrideToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["priority"] = "p1"

		assert.Equal(t, result, model)
	}

	model := new(logsv0.ApisAlertDefinitionAlertDefOverride)
	model.Priority = core.StringPtr("p1")

	result, err := logs.DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionAlertDefOverrideToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmLogsAlertDefinitionsApisAlertDefinitionLogsRatioThresholdTypeToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		apisAlertDefinitionLabelFilterTypeModel := make(map[string]interface{})
		apisAlertDefinitionLabelFilterTypeModel["value"] = "my-app"
		apisAlertDefinitionLabelFilterTypeModel["operation"] = "starts_with"

		apisAlertDefinitionLabelFiltersModel := make(map[string]interface{})
		apisAlertDefinitionLabelFiltersModel["application_name"] = []map[string]interface{}{apisAlertDefinitionLabelFilterTypeModel}
		apisAlertDefinitionLabelFiltersModel["subsystem_name"] = []map[string]interface{}{apisAlertDefinitionLabelFilterTypeModel}
		apisAlertDefinitionLabelFiltersModel["severities"] = []string{"critical"}

		apisAlertDefinitionLogsSimpleFilterModel := make(map[string]interface{})
		apisAlertDefinitionLogsSimpleFilterModel["lucene_query"] = "text:"error""
		apisAlertDefinitionLogsSimpleFilterModel["label_filters"] = []map[string]interface{}{apisAlertDefinitionLabelFiltersModel}

		apisAlertDefinitionLogsFilterModel := make(map[string]interface{})
		apisAlertDefinitionLogsFilterModel["simple_filter"] = []map[string]interface{}{apisAlertDefinitionLogsSimpleFilterModel}

		apisAlertDefinitionLogsRatioTimeWindowModel := make(map[string]interface{})
		apisAlertDefinitionLogsRatioTimeWindowModel["logs_ratio_time_window_specific_value"] = "hours_36"

		apisAlertDefinitionLogsRatioConditionModel := make(map[string]interface{})
		apisAlertDefinitionLogsRatioConditionModel["threshold"] = float64(10.0)
		apisAlertDefinitionLogsRatioConditionModel["time_window"] = []map[string]interface{}{apisAlertDefinitionLogsRatioTimeWindowModel}

		apisAlertDefinitionAlertDefOverrideModel := make(map[string]interface{})
		apisAlertDefinitionAlertDefOverrideModel["priority"] = "p1"

		apisAlertDefinitionLogsRatioRulesModel := make(map[string]interface{})
		apisAlertDefinitionLogsRatioRulesModel["condition"] = []map[string]interface{}{apisAlertDefinitionLogsRatioConditionModel}
		apisAlertDefinitionLogsRatioRulesModel["override"] = []map[string]interface{}{apisAlertDefinitionAlertDefOverrideModel}

		apisAlertDefinitionUndetectedValuesManagementModel := make(map[string]interface{})
		apisAlertDefinitionUndetectedValuesManagementModel["trigger_undetected_values"] = true
		apisAlertDefinitionUndetectedValuesManagementModel["auto_retire_timeframe"] = "hours_24"

		model := make(map[string]interface{})
		model["numerator"] = []map[string]interface{}{apisAlertDefinitionLogsFilterModel}
		model["numerator_alias"] = "numerator_alias"
		model["denominator"] = []map[string]interface{}{apisAlertDefinitionLogsFilterModel}
		model["denominator_alias"] = "denominator_alias"
		model["rules"] = []map[string]interface{}{apisAlertDefinitionLogsRatioRulesModel}
		model["condition_type"] = "less_than"
		model["notification_payload_filter"] = []string{"obj.field"}
		model["group_by_for"] = "denumerator_only"
		model["undetected_values_management"] = []map[string]interface{}{apisAlertDefinitionUndetectedValuesManagementModel}
		model["ignore_infinity"] = true
		model["evaluation_delay_ms"] = int(60000)

		assert.Equal(t, result, model)
	}

	apisAlertDefinitionLabelFilterTypeModel := new(logsv0.ApisAlertDefinitionLabelFilterType)
	apisAlertDefinitionLabelFilterTypeModel.Value = core.StringPtr("my-app")
	apisAlertDefinitionLabelFilterTypeModel.Operation = core.StringPtr("starts_with")

	apisAlertDefinitionLabelFiltersModel := new(logsv0.ApisAlertDefinitionLabelFilters)
	apisAlertDefinitionLabelFiltersModel.ApplicationName = []logsv0.ApisAlertDefinitionLabelFilterType{*apisAlertDefinitionLabelFilterTypeModel}
	apisAlertDefinitionLabelFiltersModel.SubsystemName = []logsv0.ApisAlertDefinitionLabelFilterType{*apisAlertDefinitionLabelFilterTypeModel}
	apisAlertDefinitionLabelFiltersModel.Severities = []string{"critical"}

	apisAlertDefinitionLogsSimpleFilterModel := new(logsv0.ApisAlertDefinitionLogsSimpleFilter)
	apisAlertDefinitionLogsSimpleFilterModel.LuceneQuery = core.StringPtr("text:"error"")
	apisAlertDefinitionLogsSimpleFilterModel.LabelFilters = apisAlertDefinitionLabelFiltersModel

	apisAlertDefinitionLogsFilterModel := new(logsv0.ApisAlertDefinitionLogsFilter)
	apisAlertDefinitionLogsFilterModel.SimpleFilter = apisAlertDefinitionLogsSimpleFilterModel

	apisAlertDefinitionLogsRatioTimeWindowModel := new(logsv0.ApisAlertDefinitionLogsRatioTimeWindow)
	apisAlertDefinitionLogsRatioTimeWindowModel.LogsRatioTimeWindowSpecificValue = core.StringPtr("hours_36")

	apisAlertDefinitionLogsRatioConditionModel := new(logsv0.ApisAlertDefinitionLogsRatioCondition)
	apisAlertDefinitionLogsRatioConditionModel.Threshold = core.Float64Ptr(float64(10.0))
	apisAlertDefinitionLogsRatioConditionModel.TimeWindow = apisAlertDefinitionLogsRatioTimeWindowModel

	apisAlertDefinitionAlertDefOverrideModel := new(logsv0.ApisAlertDefinitionAlertDefOverride)
	apisAlertDefinitionAlertDefOverrideModel.Priority = core.StringPtr("p1")

	apisAlertDefinitionLogsRatioRulesModel := new(logsv0.ApisAlertDefinitionLogsRatioRules)
	apisAlertDefinitionLogsRatioRulesModel.Condition = apisAlertDefinitionLogsRatioConditionModel
	apisAlertDefinitionLogsRatioRulesModel.Override = apisAlertDefinitionAlertDefOverrideModel

	apisAlertDefinitionUndetectedValuesManagementModel := new(logsv0.ApisAlertDefinitionUndetectedValuesManagement)
	apisAlertDefinitionUndetectedValuesManagementModel.TriggerUndetectedValues = core.BoolPtr(true)
	apisAlertDefinitionUndetectedValuesManagementModel.AutoRetireTimeframe = core.StringPtr("hours_24")

	model := new(logsv0.ApisAlertDefinitionLogsRatioThresholdType)
	model.Numerator = apisAlertDefinitionLogsFilterModel
	model.NumeratorAlias = core.StringPtr("numerator_alias")
	model.Denominator = apisAlertDefinitionLogsFilterModel
	model.DenominatorAlias = core.StringPtr("denominator_alias")
	model.Rules = []logsv0.ApisAlertDefinitionLogsRatioRules{*apisAlertDefinitionLogsRatioRulesModel}
	model.ConditionType = core.StringPtr("less_than")
	model.NotificationPayloadFilter = []string{"obj.field"}
	model.GroupByFor = core.StringPtr("denumerator_only")
	model.UndetectedValuesManagement = apisAlertDefinitionUndetectedValuesManagementModel
	model.IgnoreInfinity = core.BoolPtr(true)
	model.EvaluationDelayMs = core.Int64Ptr(int64(60000))

	result, err := logs.DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionLogsRatioThresholdTypeToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmLogsAlertDefinitionsApisAlertDefinitionLogsRatioRulesToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		apisAlertDefinitionLogsRatioTimeWindowModel := make(map[string]interface{})
		apisAlertDefinitionLogsRatioTimeWindowModel["logs_ratio_time_window_specific_value"] = "hours_36"

		apisAlertDefinitionLogsRatioConditionModel := make(map[string]interface{})
		apisAlertDefinitionLogsRatioConditionModel["threshold"] = float64(10.0)
		apisAlertDefinitionLogsRatioConditionModel["time_window"] = []map[string]interface{}{apisAlertDefinitionLogsRatioTimeWindowModel}

		apisAlertDefinitionAlertDefOverrideModel := make(map[string]interface{})
		apisAlertDefinitionAlertDefOverrideModel["priority"] = "p1"

		model := make(map[string]interface{})
		model["condition"] = []map[string]interface{}{apisAlertDefinitionLogsRatioConditionModel}
		model["override"] = []map[string]interface{}{apisAlertDefinitionAlertDefOverrideModel}

		assert.Equal(t, result, model)
	}

	apisAlertDefinitionLogsRatioTimeWindowModel := new(logsv0.ApisAlertDefinitionLogsRatioTimeWindow)
	apisAlertDefinitionLogsRatioTimeWindowModel.LogsRatioTimeWindowSpecificValue = core.StringPtr("hours_36")

	apisAlertDefinitionLogsRatioConditionModel := new(logsv0.ApisAlertDefinitionLogsRatioCondition)
	apisAlertDefinitionLogsRatioConditionModel.Threshold = core.Float64Ptr(float64(10.0))
	apisAlertDefinitionLogsRatioConditionModel.TimeWindow = apisAlertDefinitionLogsRatioTimeWindowModel

	apisAlertDefinitionAlertDefOverrideModel := new(logsv0.ApisAlertDefinitionAlertDefOverride)
	apisAlertDefinitionAlertDefOverrideModel.Priority = core.StringPtr("p1")

	model := new(logsv0.ApisAlertDefinitionLogsRatioRules)
	model.Condition = apisAlertDefinitionLogsRatioConditionModel
	model.Override = apisAlertDefinitionAlertDefOverrideModel

	result, err := logs.DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionLogsRatioRulesToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmLogsAlertDefinitionsApisAlertDefinitionLogsRatioConditionToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		apisAlertDefinitionLogsRatioTimeWindowModel := make(map[string]interface{})
		apisAlertDefinitionLogsRatioTimeWindowModel["logs_ratio_time_window_specific_value"] = "hours_36"

		model := make(map[string]interface{})
		model["threshold"] = float64(10.0)
		model["time_window"] = []map[string]interface{}{apisAlertDefinitionLogsRatioTimeWindowModel}

		assert.Equal(t, result, model)
	}

	apisAlertDefinitionLogsRatioTimeWindowModel := new(logsv0.ApisAlertDefinitionLogsRatioTimeWindow)
	apisAlertDefinitionLogsRatioTimeWindowModel.LogsRatioTimeWindowSpecificValue = core.StringPtr("hours_36")

	model := new(logsv0.ApisAlertDefinitionLogsRatioCondition)
	model.Threshold = core.Float64Ptr(float64(10.0))
	model.TimeWindow = apisAlertDefinitionLogsRatioTimeWindowModel

	result, err := logs.DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionLogsRatioConditionToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmLogsAlertDefinitionsApisAlertDefinitionLogsRatioTimeWindowToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["logs_ratio_time_window_specific_value"] = "hours_36"

		assert.Equal(t, result, model)
	}

	model := new(logsv0.ApisAlertDefinitionLogsRatioTimeWindow)
	model.LogsRatioTimeWindowSpecificValue = core.StringPtr("hours_36")

	result, err := logs.DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionLogsRatioTimeWindowToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmLogsAlertDefinitionsApisAlertDefinitionLogsTimeRelativeThresholdTypeToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		apisAlertDefinitionLabelFilterTypeModel := make(map[string]interface{})
		apisAlertDefinitionLabelFilterTypeModel["value"] = "my-app"
		apisAlertDefinitionLabelFilterTypeModel["operation"] = "starts_with"

		apisAlertDefinitionLabelFiltersModel := make(map[string]interface{})
		apisAlertDefinitionLabelFiltersModel["application_name"] = []map[string]interface{}{apisAlertDefinitionLabelFilterTypeModel}
		apisAlertDefinitionLabelFiltersModel["subsystem_name"] = []map[string]interface{}{apisAlertDefinitionLabelFilterTypeModel}
		apisAlertDefinitionLabelFiltersModel["severities"] = []string{"critical"}

		apisAlertDefinitionLogsSimpleFilterModel := make(map[string]interface{})
		apisAlertDefinitionLogsSimpleFilterModel["lucene_query"] = "text:"error""
		apisAlertDefinitionLogsSimpleFilterModel["label_filters"] = []map[string]interface{}{apisAlertDefinitionLabelFiltersModel}

		apisAlertDefinitionLogsFilterModel := make(map[string]interface{})
		apisAlertDefinitionLogsFilterModel["simple_filter"] = []map[string]interface{}{apisAlertDefinitionLogsSimpleFilterModel}

		apisAlertDefinitionLogsTimeRelativeConditionModel := make(map[string]interface{})
		apisAlertDefinitionLogsTimeRelativeConditionModel["threshold"] = float64(100.0)
		apisAlertDefinitionLogsTimeRelativeConditionModel["compared_to"] = "same_day_last_month"

		apisAlertDefinitionAlertDefOverrideModel := make(map[string]interface{})
		apisAlertDefinitionAlertDefOverrideModel["priority"] = "p1"

		apisAlertDefinitionLogsTimeRelativeRuleModel := make(map[string]interface{})
		apisAlertDefinitionLogsTimeRelativeRuleModel["condition"] = []map[string]interface{}{apisAlertDefinitionLogsTimeRelativeConditionModel}
		apisAlertDefinitionLogsTimeRelativeRuleModel["override"] = []map[string]interface{}{apisAlertDefinitionAlertDefOverrideModel}

		apisAlertDefinitionUndetectedValuesManagementModel := make(map[string]interface{})
		apisAlertDefinitionUndetectedValuesManagementModel["trigger_undetected_values"] = true
		apisAlertDefinitionUndetectedValuesManagementModel["auto_retire_timeframe"] = "hours_24"

		model := make(map[string]interface{})
		model["logs_filter"] = []map[string]interface{}{apisAlertDefinitionLogsFilterModel}
		model["rules"] = []map[string]interface{}{apisAlertDefinitionLogsTimeRelativeRuleModel}
		model["condition_type"] = "less_than"
		model["ignore_infinity"] = true
		model["notification_payload_filter"] = []string{"obj.field"}
		model["undetected_values_management"] = []map[string]interface{}{apisAlertDefinitionUndetectedValuesManagementModel}
		model["evaluation_delay_ms"] = int(60000)

		assert.Equal(t, result, model)
	}

	apisAlertDefinitionLabelFilterTypeModel := new(logsv0.ApisAlertDefinitionLabelFilterType)
	apisAlertDefinitionLabelFilterTypeModel.Value = core.StringPtr("my-app")
	apisAlertDefinitionLabelFilterTypeModel.Operation = core.StringPtr("starts_with")

	apisAlertDefinitionLabelFiltersModel := new(logsv0.ApisAlertDefinitionLabelFilters)
	apisAlertDefinitionLabelFiltersModel.ApplicationName = []logsv0.ApisAlertDefinitionLabelFilterType{*apisAlertDefinitionLabelFilterTypeModel}
	apisAlertDefinitionLabelFiltersModel.SubsystemName = []logsv0.ApisAlertDefinitionLabelFilterType{*apisAlertDefinitionLabelFilterTypeModel}
	apisAlertDefinitionLabelFiltersModel.Severities = []string{"critical"}

	apisAlertDefinitionLogsSimpleFilterModel := new(logsv0.ApisAlertDefinitionLogsSimpleFilter)
	apisAlertDefinitionLogsSimpleFilterModel.LuceneQuery = core.StringPtr("text:"error"")
	apisAlertDefinitionLogsSimpleFilterModel.LabelFilters = apisAlertDefinitionLabelFiltersModel

	apisAlertDefinitionLogsFilterModel := new(logsv0.ApisAlertDefinitionLogsFilter)
	apisAlertDefinitionLogsFilterModel.SimpleFilter = apisAlertDefinitionLogsSimpleFilterModel

	apisAlertDefinitionLogsTimeRelativeConditionModel := new(logsv0.ApisAlertDefinitionLogsTimeRelativeCondition)
	apisAlertDefinitionLogsTimeRelativeConditionModel.Threshold = core.Float64Ptr(float64(100.0))
	apisAlertDefinitionLogsTimeRelativeConditionModel.ComparedTo = core.StringPtr("same_day_last_month")

	apisAlertDefinitionAlertDefOverrideModel := new(logsv0.ApisAlertDefinitionAlertDefOverride)
	apisAlertDefinitionAlertDefOverrideModel.Priority = core.StringPtr("p1")

	apisAlertDefinitionLogsTimeRelativeRuleModel := new(logsv0.ApisAlertDefinitionLogsTimeRelativeRule)
	apisAlertDefinitionLogsTimeRelativeRuleModel.Condition = apisAlertDefinitionLogsTimeRelativeConditionModel
	apisAlertDefinitionLogsTimeRelativeRuleModel.Override = apisAlertDefinitionAlertDefOverrideModel

	apisAlertDefinitionUndetectedValuesManagementModel := new(logsv0.ApisAlertDefinitionUndetectedValuesManagement)
	apisAlertDefinitionUndetectedValuesManagementModel.TriggerUndetectedValues = core.BoolPtr(true)
	apisAlertDefinitionUndetectedValuesManagementModel.AutoRetireTimeframe = core.StringPtr("hours_24")

	model := new(logsv0.ApisAlertDefinitionLogsTimeRelativeThresholdType)
	model.LogsFilter = apisAlertDefinitionLogsFilterModel
	model.Rules = []logsv0.ApisAlertDefinitionLogsTimeRelativeRule{*apisAlertDefinitionLogsTimeRelativeRuleModel}
	model.ConditionType = core.StringPtr("less_than")
	model.IgnoreInfinity = core.BoolPtr(true)
	model.NotificationPayloadFilter = []string{"obj.field"}
	model.UndetectedValuesManagement = apisAlertDefinitionUndetectedValuesManagementModel
	model.EvaluationDelayMs = core.Int64Ptr(int64(60000))

	result, err := logs.DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionLogsTimeRelativeThresholdTypeToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmLogsAlertDefinitionsApisAlertDefinitionLogsTimeRelativeRuleToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		apisAlertDefinitionLogsTimeRelativeConditionModel := make(map[string]interface{})
		apisAlertDefinitionLogsTimeRelativeConditionModel["threshold"] = float64(100.0)
		apisAlertDefinitionLogsTimeRelativeConditionModel["compared_to"] = "same_day_last_month"

		apisAlertDefinitionAlertDefOverrideModel := make(map[string]interface{})
		apisAlertDefinitionAlertDefOverrideModel["priority"] = "p1"

		model := make(map[string]interface{})
		model["condition"] = []map[string]interface{}{apisAlertDefinitionLogsTimeRelativeConditionModel}
		model["override"] = []map[string]interface{}{apisAlertDefinitionAlertDefOverrideModel}

		assert.Equal(t, result, model)
	}

	apisAlertDefinitionLogsTimeRelativeConditionModel := new(logsv0.ApisAlertDefinitionLogsTimeRelativeCondition)
	apisAlertDefinitionLogsTimeRelativeConditionModel.Threshold = core.Float64Ptr(float64(100.0))
	apisAlertDefinitionLogsTimeRelativeConditionModel.ComparedTo = core.StringPtr("same_day_last_month")

	apisAlertDefinitionAlertDefOverrideModel := new(logsv0.ApisAlertDefinitionAlertDefOverride)
	apisAlertDefinitionAlertDefOverrideModel.Priority = core.StringPtr("p1")

	model := new(logsv0.ApisAlertDefinitionLogsTimeRelativeRule)
	model.Condition = apisAlertDefinitionLogsTimeRelativeConditionModel
	model.Override = apisAlertDefinitionAlertDefOverrideModel

	result, err := logs.DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionLogsTimeRelativeRuleToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmLogsAlertDefinitionsApisAlertDefinitionLogsTimeRelativeConditionToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["threshold"] = float64(100.0)
		model["compared_to"] = "same_day_last_month"

		assert.Equal(t, result, model)
	}

	model := new(logsv0.ApisAlertDefinitionLogsTimeRelativeCondition)
	model.Threshold = core.Float64Ptr(float64(100.0))
	model.ComparedTo = core.StringPtr("same_day_last_month")

	result, err := logs.DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionLogsTimeRelativeConditionToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmLogsAlertDefinitionsApisAlertDefinitionMetricThresholdTypeToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		apisAlertDefinitionMetricFilterModel := make(map[string]interface{})
		apisAlertDefinitionMetricFilterModel["promql"] = "avg_over_time(metric_name[5m]) > 10"

		apisAlertDefinitionMetricTimeWindowModel := make(map[string]interface{})
		apisAlertDefinitionMetricTimeWindowModel["metric_time_window_specific_value"] = "hours_36"

		apisAlertDefinitionMetricThresholdConditionModel := make(map[string]interface{})
		apisAlertDefinitionMetricThresholdConditionModel["threshold"] = float64(100.0)
		apisAlertDefinitionMetricThresholdConditionModel["for_over_pct"] = int(80)
		apisAlertDefinitionMetricThresholdConditionModel["of_the_last"] = []map[string]interface{}{apisAlertDefinitionMetricTimeWindowModel}

		apisAlertDefinitionAlertDefOverrideModel := make(map[string]interface{})
		apisAlertDefinitionAlertDefOverrideModel["priority"] = "p1"

		apisAlertDefinitionMetricThresholdRuleModel := make(map[string]interface{})
		apisAlertDefinitionMetricThresholdRuleModel["condition"] = []map[string]interface{}{apisAlertDefinitionMetricThresholdConditionModel}
		apisAlertDefinitionMetricThresholdRuleModel["override"] = []map[string]interface{}{apisAlertDefinitionAlertDefOverrideModel}

		apisAlertDefinitionUndetectedValuesManagementModel := make(map[string]interface{})
		apisAlertDefinitionUndetectedValuesManagementModel["trigger_undetected_values"] = true
		apisAlertDefinitionUndetectedValuesManagementModel["auto_retire_timeframe"] = "hours_24"

		apisAlertDefinitionMetricMissingValuesModel := make(map[string]interface{})
		apisAlertDefinitionMetricMissingValuesModel["replace_with_zero"] = true

		model := make(map[string]interface{})
		model["metric_filter"] = []map[string]interface{}{apisAlertDefinitionMetricFilterModel}
		model["rules"] = []map[string]interface{}{apisAlertDefinitionMetricThresholdRuleModel}
		model["condition_type"] = "less_than_or_equals"
		model["undetected_values_management"] = []map[string]interface{}{apisAlertDefinitionUndetectedValuesManagementModel}
		model["missing_values"] = []map[string]interface{}{apisAlertDefinitionMetricMissingValuesModel}
		model["evaluation_delay_ms"] = int(60000)

		assert.Equal(t, result, model)
	}

	apisAlertDefinitionMetricFilterModel := new(logsv0.ApisAlertDefinitionMetricFilter)
	apisAlertDefinitionMetricFilterModel.Promql = core.StringPtr("avg_over_time(metric_name[5m]) > 10")

	apisAlertDefinitionMetricTimeWindowModel := new(logsv0.ApisAlertDefinitionMetricTimeWindowTypeMetricTimeWindowSpecificValue)
	apisAlertDefinitionMetricTimeWindowModel.MetricTimeWindowSpecificValue = core.StringPtr("hours_36")

	apisAlertDefinitionMetricThresholdConditionModel := new(logsv0.ApisAlertDefinitionMetricThresholdCondition)
	apisAlertDefinitionMetricThresholdConditionModel.Threshold = core.Float64Ptr(float64(100.0))
	apisAlertDefinitionMetricThresholdConditionModel.ForOverPct = core.Int64Ptr(int64(80))
	apisAlertDefinitionMetricThresholdConditionModel.OfTheLast = apisAlertDefinitionMetricTimeWindowModel

	apisAlertDefinitionAlertDefOverrideModel := new(logsv0.ApisAlertDefinitionAlertDefOverride)
	apisAlertDefinitionAlertDefOverrideModel.Priority = core.StringPtr("p1")

	apisAlertDefinitionMetricThresholdRuleModel := new(logsv0.ApisAlertDefinitionMetricThresholdRule)
	apisAlertDefinitionMetricThresholdRuleModel.Condition = apisAlertDefinitionMetricThresholdConditionModel
	apisAlertDefinitionMetricThresholdRuleModel.Override = apisAlertDefinitionAlertDefOverrideModel

	apisAlertDefinitionUndetectedValuesManagementModel := new(logsv0.ApisAlertDefinitionUndetectedValuesManagement)
	apisAlertDefinitionUndetectedValuesManagementModel.TriggerUndetectedValues = core.BoolPtr(true)
	apisAlertDefinitionUndetectedValuesManagementModel.AutoRetireTimeframe = core.StringPtr("hours_24")

	apisAlertDefinitionMetricMissingValuesModel := new(logsv0.ApisAlertDefinitionMetricMissingValuesMissingValuesReplaceWithZero)
	apisAlertDefinitionMetricMissingValuesModel.ReplaceWithZero = core.BoolPtr(true)

	model := new(logsv0.ApisAlertDefinitionMetricThresholdType)
	model.MetricFilter = apisAlertDefinitionMetricFilterModel
	model.Rules = []logsv0.ApisAlertDefinitionMetricThresholdRule{*apisAlertDefinitionMetricThresholdRuleModel}
	model.ConditionType = core.StringPtr("less_than_or_equals")
	model.UndetectedValuesManagement = apisAlertDefinitionUndetectedValuesManagementModel
	model.MissingValues = apisAlertDefinitionMetricMissingValuesModel
	model.EvaluationDelayMs = core.Int64Ptr(int64(60000))

	result, err := logs.DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionMetricThresholdTypeToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmLogsAlertDefinitionsApisAlertDefinitionMetricFilterToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["promql"] = "avg_over_time(metric_name[5m]) > 10"

		assert.Equal(t, result, model)
	}

	model := new(logsv0.ApisAlertDefinitionMetricFilter)
	model.Promql = core.StringPtr("avg_over_time(metric_name[5m]) > 10")

	result, err := logs.DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionMetricFilterToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmLogsAlertDefinitionsApisAlertDefinitionMetricThresholdRuleToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		apisAlertDefinitionMetricTimeWindowModel := make(map[string]interface{})
		apisAlertDefinitionMetricTimeWindowModel["metric_time_window_specific_value"] = "hours_36"

		apisAlertDefinitionMetricThresholdConditionModel := make(map[string]interface{})
		apisAlertDefinitionMetricThresholdConditionModel["threshold"] = float64(100.0)
		apisAlertDefinitionMetricThresholdConditionModel["for_over_pct"] = int(80)
		apisAlertDefinitionMetricThresholdConditionModel["of_the_last"] = []map[string]interface{}{apisAlertDefinitionMetricTimeWindowModel}

		apisAlertDefinitionAlertDefOverrideModel := make(map[string]interface{})
		apisAlertDefinitionAlertDefOverrideModel["priority"] = "p1"

		model := make(map[string]interface{})
		model["condition"] = []map[string]interface{}{apisAlertDefinitionMetricThresholdConditionModel}
		model["override"] = []map[string]interface{}{apisAlertDefinitionAlertDefOverrideModel}

		assert.Equal(t, result, model)
	}

	apisAlertDefinitionMetricTimeWindowModel := new(logsv0.ApisAlertDefinitionMetricTimeWindowTypeMetricTimeWindowSpecificValue)
	apisAlertDefinitionMetricTimeWindowModel.MetricTimeWindowSpecificValue = core.StringPtr("hours_36")

	apisAlertDefinitionMetricThresholdConditionModel := new(logsv0.ApisAlertDefinitionMetricThresholdCondition)
	apisAlertDefinitionMetricThresholdConditionModel.Threshold = core.Float64Ptr(float64(100.0))
	apisAlertDefinitionMetricThresholdConditionModel.ForOverPct = core.Int64Ptr(int64(80))
	apisAlertDefinitionMetricThresholdConditionModel.OfTheLast = apisAlertDefinitionMetricTimeWindowModel

	apisAlertDefinitionAlertDefOverrideModel := new(logsv0.ApisAlertDefinitionAlertDefOverride)
	apisAlertDefinitionAlertDefOverrideModel.Priority = core.StringPtr("p1")

	model := new(logsv0.ApisAlertDefinitionMetricThresholdRule)
	model.Condition = apisAlertDefinitionMetricThresholdConditionModel
	model.Override = apisAlertDefinitionAlertDefOverrideModel

	result, err := logs.DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionMetricThresholdRuleToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmLogsAlertDefinitionsApisAlertDefinitionMetricThresholdConditionToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		apisAlertDefinitionMetricTimeWindowModel := make(map[string]interface{})
		apisAlertDefinitionMetricTimeWindowModel["metric_time_window_specific_value"] = "hours_36"

		model := make(map[string]interface{})
		model["threshold"] = float64(100.0)
		model["for_over_pct"] = int(80)
		model["of_the_last"] = []map[string]interface{}{apisAlertDefinitionMetricTimeWindowModel}

		assert.Equal(t, result, model)
	}

	apisAlertDefinitionMetricTimeWindowModel := new(logsv0.ApisAlertDefinitionMetricTimeWindowTypeMetricTimeWindowSpecificValue)
	apisAlertDefinitionMetricTimeWindowModel.MetricTimeWindowSpecificValue = core.StringPtr("hours_36")

	model := new(logsv0.ApisAlertDefinitionMetricThresholdCondition)
	model.Threshold = core.Float64Ptr(float64(100.0))
	model.ForOverPct = core.Int64Ptr(int64(80))
	model.OfTheLast = apisAlertDefinitionMetricTimeWindowModel

	result, err := logs.DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionMetricThresholdConditionToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmLogsAlertDefinitionsApisAlertDefinitionMetricTimeWindowToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["metric_time_window_specific_value"] = "hours_36"
		model["metric_time_window_dynamic_duration"] = "1h30m"

		assert.Equal(t, result, model)
	}

	model := new(logsv0.ApisAlertDefinitionMetricTimeWindow)
	model.MetricTimeWindowSpecificValue = core.StringPtr("hours_36")
	model.MetricTimeWindowDynamicDuration = core.StringPtr("1h30m")

	result, err := logs.DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionMetricTimeWindowToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmLogsAlertDefinitionsApisAlertDefinitionMetricTimeWindowTypeMetricTimeWindowSpecificValueToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["metric_time_window_specific_value"] = "hours_36"

		assert.Equal(t, result, model)
	}

	model := new(logsv0.ApisAlertDefinitionMetricTimeWindowTypeMetricTimeWindowSpecificValue)
	model.MetricTimeWindowSpecificValue = core.StringPtr("hours_36")

	result, err := logs.DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionMetricTimeWindowTypeMetricTimeWindowSpecificValueToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmLogsAlertDefinitionsApisAlertDefinitionMetricTimeWindowTypeMetricTimeWindowDynamicDurationToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["metric_time_window_dynamic_duration"] = "1h30m"

		assert.Equal(t, result, model)
	}

	model := new(logsv0.ApisAlertDefinitionMetricTimeWindowTypeMetricTimeWindowDynamicDuration)
	model.MetricTimeWindowDynamicDuration = core.StringPtr("1h30m")

	result, err := logs.DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionMetricTimeWindowTypeMetricTimeWindowDynamicDurationToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmLogsAlertDefinitionsApisAlertDefinitionMetricMissingValuesToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["replace_with_zero"] = true
		model["min_non_null_values_pct"] = int(80)

		assert.Equal(t, result, model)
	}

	model := new(logsv0.ApisAlertDefinitionMetricMissingValues)
	model.ReplaceWithZero = core.BoolPtr(true)
	model.MinNonNullValuesPct = core.Int64Ptr(int64(80))

	result, err := logs.DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionMetricMissingValuesToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmLogsAlertDefinitionsApisAlertDefinitionMetricMissingValuesMissingValuesReplaceWithZeroToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["replace_with_zero"] = true

		assert.Equal(t, result, model)
	}

	model := new(logsv0.ApisAlertDefinitionMetricMissingValuesMissingValuesReplaceWithZero)
	model.ReplaceWithZero = core.BoolPtr(true)

	result, err := logs.DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionMetricMissingValuesMissingValuesReplaceWithZeroToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmLogsAlertDefinitionsApisAlertDefinitionMetricMissingValuesMissingValuesMinNonNullValuesPctToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["min_non_null_values_pct"] = int(80)

		assert.Equal(t, result, model)
	}

	model := new(logsv0.ApisAlertDefinitionMetricMissingValuesMissingValuesMinNonNullValuesPct)
	model.MinNonNullValuesPct = core.Int64Ptr(int64(80))

	result, err := logs.DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionMetricMissingValuesMissingValuesMinNonNullValuesPctToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmLogsAlertDefinitionsApisAlertDefinitionFlowTypeToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		apisAlertDefinitionFlowStagesGroupsAlertDefsModel := make(map[string]interface{})
		apisAlertDefinitionFlowStagesGroupsAlertDefsModel["id"] = "3dc02998-0b50-4ea8-b68a-4779d716fa1f"
		apisAlertDefinitionFlowStagesGroupsAlertDefsModel["not"] = true

		apisAlertDefinitionFlowStagesGroupModel := make(map[string]interface{})
		apisAlertDefinitionFlowStagesGroupModel["alert_defs"] = []map[string]interface{}{apisAlertDefinitionFlowStagesGroupsAlertDefsModel}
		apisAlertDefinitionFlowStagesGroupModel["next_op"] = "or"
		apisAlertDefinitionFlowStagesGroupModel["alerts_op"] = "or"

		apisAlertDefinitionFlowStagesGroupsModel := make(map[string]interface{})
		apisAlertDefinitionFlowStagesGroupsModel["groups"] = []map[string]interface{}{apisAlertDefinitionFlowStagesGroupModel}

		apisAlertDefinitionFlowStagesModel := make(map[string]interface{})
		apisAlertDefinitionFlowStagesModel["timeframe_ms"] = "60000"
		apisAlertDefinitionFlowStagesModel["timeframe_type"] = "up_to"
		apisAlertDefinitionFlowStagesModel["flow_stages_groups"] = []map[string]interface{}{apisAlertDefinitionFlowStagesGroupsModel}

		model := make(map[string]interface{})
		model["stages"] = []map[string]interface{}{apisAlertDefinitionFlowStagesModel}
		model["enforce_suppression"] = true

		assert.Equal(t, result, model)
	}

	apisAlertDefinitionFlowStagesGroupsAlertDefsModel := new(logsv0.ApisAlertDefinitionFlowStagesGroupsAlertDefs)
	apisAlertDefinitionFlowStagesGroupsAlertDefsModel.ID = CreateMockUUID("3dc02998-0b50-4ea8-b68a-4779d716fa1f")
	apisAlertDefinitionFlowStagesGroupsAlertDefsModel.Not = core.BoolPtr(true)

	apisAlertDefinitionFlowStagesGroupModel := new(logsv0.ApisAlertDefinitionFlowStagesGroup)
	apisAlertDefinitionFlowStagesGroupModel.AlertDefs = []logsv0.ApisAlertDefinitionFlowStagesGroupsAlertDefs{*apisAlertDefinitionFlowStagesGroupsAlertDefsModel}
	apisAlertDefinitionFlowStagesGroupModel.NextOp = core.StringPtr("or")
	apisAlertDefinitionFlowStagesGroupModel.AlertsOp = core.StringPtr("or")

	apisAlertDefinitionFlowStagesGroupsModel := new(logsv0.ApisAlertDefinitionFlowStagesGroups)
	apisAlertDefinitionFlowStagesGroupsModel.Groups = []logsv0.ApisAlertDefinitionFlowStagesGroup{*apisAlertDefinitionFlowStagesGroupModel}

	apisAlertDefinitionFlowStagesModel := new(logsv0.ApisAlertDefinitionFlowStages)
	apisAlertDefinitionFlowStagesModel.TimeframeMs = core.StringPtr("60000")
	apisAlertDefinitionFlowStagesModel.TimeframeType = core.StringPtr("up_to")
	apisAlertDefinitionFlowStagesModel.FlowStagesGroups = apisAlertDefinitionFlowStagesGroupsModel

	model := new(logsv0.ApisAlertDefinitionFlowType)
	model.Stages = []logsv0.ApisAlertDefinitionFlowStages{*apisAlertDefinitionFlowStagesModel}
	model.EnforceSuppression = core.BoolPtr(true)

	result, err := logs.DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionFlowTypeToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmLogsAlertDefinitionsApisAlertDefinitionFlowStagesToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		apisAlertDefinitionFlowStagesGroupsAlertDefsModel := make(map[string]interface{})
		apisAlertDefinitionFlowStagesGroupsAlertDefsModel["id"] = "3dc02998-0b50-4ea8-b68a-4779d716fa1f"
		apisAlertDefinitionFlowStagesGroupsAlertDefsModel["not"] = true

		apisAlertDefinitionFlowStagesGroupModel := make(map[string]interface{})
		apisAlertDefinitionFlowStagesGroupModel["alert_defs"] = []map[string]interface{}{apisAlertDefinitionFlowStagesGroupsAlertDefsModel}
		apisAlertDefinitionFlowStagesGroupModel["next_op"] = "or"
		apisAlertDefinitionFlowStagesGroupModel["alerts_op"] = "or"

		apisAlertDefinitionFlowStagesGroupsModel := make(map[string]interface{})
		apisAlertDefinitionFlowStagesGroupsModel["groups"] = []map[string]interface{}{apisAlertDefinitionFlowStagesGroupModel}

		model := make(map[string]interface{})
		model["timeframe_ms"] = "60000"
		model["timeframe_type"] = "up_to"
		model["flow_stages_groups"] = []map[string]interface{}{apisAlertDefinitionFlowStagesGroupsModel}

		assert.Equal(t, result, model)
	}

	apisAlertDefinitionFlowStagesGroupsAlertDefsModel := new(logsv0.ApisAlertDefinitionFlowStagesGroupsAlertDefs)
	apisAlertDefinitionFlowStagesGroupsAlertDefsModel.ID = CreateMockUUID("3dc02998-0b50-4ea8-b68a-4779d716fa1f")
	apisAlertDefinitionFlowStagesGroupsAlertDefsModel.Not = core.BoolPtr(true)

	apisAlertDefinitionFlowStagesGroupModel := new(logsv0.ApisAlertDefinitionFlowStagesGroup)
	apisAlertDefinitionFlowStagesGroupModel.AlertDefs = []logsv0.ApisAlertDefinitionFlowStagesGroupsAlertDefs{*apisAlertDefinitionFlowStagesGroupsAlertDefsModel}
	apisAlertDefinitionFlowStagesGroupModel.NextOp = core.StringPtr("or")
	apisAlertDefinitionFlowStagesGroupModel.AlertsOp = core.StringPtr("or")

	apisAlertDefinitionFlowStagesGroupsModel := new(logsv0.ApisAlertDefinitionFlowStagesGroups)
	apisAlertDefinitionFlowStagesGroupsModel.Groups = []logsv0.ApisAlertDefinitionFlowStagesGroup{*apisAlertDefinitionFlowStagesGroupModel}

	model := new(logsv0.ApisAlertDefinitionFlowStages)
	model.TimeframeMs = core.StringPtr("60000")
	model.TimeframeType = core.StringPtr("up_to")
	model.FlowStagesGroups = apisAlertDefinitionFlowStagesGroupsModel

	result, err := logs.DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionFlowStagesToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmLogsAlertDefinitionsApisAlertDefinitionFlowStagesGroupsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		apisAlertDefinitionFlowStagesGroupsAlertDefsModel := make(map[string]interface{})
		apisAlertDefinitionFlowStagesGroupsAlertDefsModel["id"] = "3dc02998-0b50-4ea8-b68a-4779d716fa1f"
		apisAlertDefinitionFlowStagesGroupsAlertDefsModel["not"] = true

		apisAlertDefinitionFlowStagesGroupModel := make(map[string]interface{})
		apisAlertDefinitionFlowStagesGroupModel["alert_defs"] = []map[string]interface{}{apisAlertDefinitionFlowStagesGroupsAlertDefsModel}
		apisAlertDefinitionFlowStagesGroupModel["next_op"] = "or"
		apisAlertDefinitionFlowStagesGroupModel["alerts_op"] = "or"

		model := make(map[string]interface{})
		model["groups"] = []map[string]interface{}{apisAlertDefinitionFlowStagesGroupModel}

		assert.Equal(t, result, model)
	}

	apisAlertDefinitionFlowStagesGroupsAlertDefsModel := new(logsv0.ApisAlertDefinitionFlowStagesGroupsAlertDefs)
	apisAlertDefinitionFlowStagesGroupsAlertDefsModel.ID = CreateMockUUID("3dc02998-0b50-4ea8-b68a-4779d716fa1f")
	apisAlertDefinitionFlowStagesGroupsAlertDefsModel.Not = core.BoolPtr(true)

	apisAlertDefinitionFlowStagesGroupModel := new(logsv0.ApisAlertDefinitionFlowStagesGroup)
	apisAlertDefinitionFlowStagesGroupModel.AlertDefs = []logsv0.ApisAlertDefinitionFlowStagesGroupsAlertDefs{*apisAlertDefinitionFlowStagesGroupsAlertDefsModel}
	apisAlertDefinitionFlowStagesGroupModel.NextOp = core.StringPtr("or")
	apisAlertDefinitionFlowStagesGroupModel.AlertsOp = core.StringPtr("or")

	model := new(logsv0.ApisAlertDefinitionFlowStagesGroups)
	model.Groups = []logsv0.ApisAlertDefinitionFlowStagesGroup{*apisAlertDefinitionFlowStagesGroupModel}

	result, err := logs.DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionFlowStagesGroupsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmLogsAlertDefinitionsApisAlertDefinitionFlowStagesGroupToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		apisAlertDefinitionFlowStagesGroupsAlertDefsModel := make(map[string]interface{})
		apisAlertDefinitionFlowStagesGroupsAlertDefsModel["id"] = "3dc02998-0b50-4ea8-b68a-4779d716fa1f"
		apisAlertDefinitionFlowStagesGroupsAlertDefsModel["not"] = true

		model := make(map[string]interface{})
		model["alert_defs"] = []map[string]interface{}{apisAlertDefinitionFlowStagesGroupsAlertDefsModel}
		model["next_op"] = "or"
		model["alerts_op"] = "or"

		assert.Equal(t, result, model)
	}

	apisAlertDefinitionFlowStagesGroupsAlertDefsModel := new(logsv0.ApisAlertDefinitionFlowStagesGroupsAlertDefs)
	apisAlertDefinitionFlowStagesGroupsAlertDefsModel.ID = CreateMockUUID("3dc02998-0b50-4ea8-b68a-4779d716fa1f")
	apisAlertDefinitionFlowStagesGroupsAlertDefsModel.Not = core.BoolPtr(true)

	model := new(logsv0.ApisAlertDefinitionFlowStagesGroup)
	model.AlertDefs = []logsv0.ApisAlertDefinitionFlowStagesGroupsAlertDefs{*apisAlertDefinitionFlowStagesGroupsAlertDefsModel}
	model.NextOp = core.StringPtr("or")
	model.AlertsOp = core.StringPtr("or")

	result, err := logs.DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionFlowStagesGroupToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmLogsAlertDefinitionsApisAlertDefinitionFlowStagesGroupsAlertDefsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["id"] = "3dc02998-0b50-4ea8-b68a-4779d716fa1f"
		model["not"] = true

		assert.Equal(t, result, model)
	}

	model := new(logsv0.ApisAlertDefinitionFlowStagesGroupsAlertDefs)
	model.ID = CreateMockUUID("3dc02998-0b50-4ea8-b68a-4779d716fa1f")
	model.Not = core.BoolPtr(true)

	result, err := logs.DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionFlowStagesGroupsAlertDefsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmLogsAlertDefinitionsApisAlertDefinitionLogsAnomalyTypeToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		apisAlertDefinitionLabelFilterTypeModel := make(map[string]interface{})
		apisAlertDefinitionLabelFilterTypeModel["value"] = "my-app"
		apisAlertDefinitionLabelFilterTypeModel["operation"] = "starts_with"

		apisAlertDefinitionLabelFiltersModel := make(map[string]interface{})
		apisAlertDefinitionLabelFiltersModel["application_name"] = []map[string]interface{}{apisAlertDefinitionLabelFilterTypeModel}
		apisAlertDefinitionLabelFiltersModel["subsystem_name"] = []map[string]interface{}{apisAlertDefinitionLabelFilterTypeModel}
		apisAlertDefinitionLabelFiltersModel["severities"] = []string{"critical"}

		apisAlertDefinitionLogsSimpleFilterModel := make(map[string]interface{})
		apisAlertDefinitionLogsSimpleFilterModel["lucene_query"] = "text:"error""
		apisAlertDefinitionLogsSimpleFilterModel["label_filters"] = []map[string]interface{}{apisAlertDefinitionLabelFiltersModel}

		apisAlertDefinitionLogsFilterModel := make(map[string]interface{})
		apisAlertDefinitionLogsFilterModel["simple_filter"] = []map[string]interface{}{apisAlertDefinitionLogsSimpleFilterModel}

		apisAlertDefinitionLogsTimeWindowModel := make(map[string]interface{})
		apisAlertDefinitionLogsTimeWindowModel["logs_time_window_specific_value"] = "hours_36"

		apisAlertDefinitionLogsAnomalyConditionModel := make(map[string]interface{})
		apisAlertDefinitionLogsAnomalyConditionModel["minimum_threshold"] = float64(10.0)
		apisAlertDefinitionLogsAnomalyConditionModel["time_window"] = []map[string]interface{}{apisAlertDefinitionLogsTimeWindowModel}

		apisAlertDefinitionLogsAnomalyRuleModel := make(map[string]interface{})
		apisAlertDefinitionLogsAnomalyRuleModel["condition"] = []map[string]interface{}{apisAlertDefinitionLogsAnomalyConditionModel}

		apisAlertDefinitionAnomalyAlertSettingsModel := make(map[string]interface{})
		apisAlertDefinitionAnomalyAlertSettingsModel["percentage_of_deviation"] = float64(10.0)

		model := make(map[string]interface{})
		model["logs_filter"] = []map[string]interface{}{apisAlertDefinitionLogsFilterModel}
		model["rules"] = []map[string]interface{}{apisAlertDefinitionLogsAnomalyRuleModel}
		model["condition_type"] = "more_than_usual_or_unspecified"
		model["notification_payload_filter"] = []string{"obj.field"}
		model["evaluation_delay_ms"] = int(60000)
		model["anomaly_alert_settings"] = []map[string]interface{}{apisAlertDefinitionAnomalyAlertSettingsModel}

		assert.Equal(t, result, model)
	}

	apisAlertDefinitionLabelFilterTypeModel := new(logsv0.ApisAlertDefinitionLabelFilterType)
	apisAlertDefinitionLabelFilterTypeModel.Value = core.StringPtr("my-app")
	apisAlertDefinitionLabelFilterTypeModel.Operation = core.StringPtr("starts_with")

	apisAlertDefinitionLabelFiltersModel := new(logsv0.ApisAlertDefinitionLabelFilters)
	apisAlertDefinitionLabelFiltersModel.ApplicationName = []logsv0.ApisAlertDefinitionLabelFilterType{*apisAlertDefinitionLabelFilterTypeModel}
	apisAlertDefinitionLabelFiltersModel.SubsystemName = []logsv0.ApisAlertDefinitionLabelFilterType{*apisAlertDefinitionLabelFilterTypeModel}
	apisAlertDefinitionLabelFiltersModel.Severities = []string{"critical"}

	apisAlertDefinitionLogsSimpleFilterModel := new(logsv0.ApisAlertDefinitionLogsSimpleFilter)
	apisAlertDefinitionLogsSimpleFilterModel.LuceneQuery = core.StringPtr("text:"error"")
	apisAlertDefinitionLogsSimpleFilterModel.LabelFilters = apisAlertDefinitionLabelFiltersModel

	apisAlertDefinitionLogsFilterModel := new(logsv0.ApisAlertDefinitionLogsFilter)
	apisAlertDefinitionLogsFilterModel.SimpleFilter = apisAlertDefinitionLogsSimpleFilterModel

	apisAlertDefinitionLogsTimeWindowModel := new(logsv0.ApisAlertDefinitionLogsTimeWindow)
	apisAlertDefinitionLogsTimeWindowModel.LogsTimeWindowSpecificValue = core.StringPtr("hours_36")

	apisAlertDefinitionLogsAnomalyConditionModel := new(logsv0.ApisAlertDefinitionLogsAnomalyCondition)
	apisAlertDefinitionLogsAnomalyConditionModel.MinimumThreshold = core.Float64Ptr(float64(10.0))
	apisAlertDefinitionLogsAnomalyConditionModel.TimeWindow = apisAlertDefinitionLogsTimeWindowModel

	apisAlertDefinitionLogsAnomalyRuleModel := new(logsv0.ApisAlertDefinitionLogsAnomalyRule)
	apisAlertDefinitionLogsAnomalyRuleModel.Condition = apisAlertDefinitionLogsAnomalyConditionModel

	apisAlertDefinitionAnomalyAlertSettingsModel := new(logsv0.ApisAlertDefinitionAnomalyAlertSettings)
	apisAlertDefinitionAnomalyAlertSettingsModel.PercentageOfDeviation = core.Float32Ptr(float32(10.0))

	model := new(logsv0.ApisAlertDefinitionLogsAnomalyType)
	model.LogsFilter = apisAlertDefinitionLogsFilterModel
	model.Rules = []logsv0.ApisAlertDefinitionLogsAnomalyRule{*apisAlertDefinitionLogsAnomalyRuleModel}
	model.ConditionType = core.StringPtr("more_than_usual_or_unspecified")
	model.NotificationPayloadFilter = []string{"obj.field"}
	model.EvaluationDelayMs = core.Int64Ptr(int64(60000))
	model.AnomalyAlertSettings = apisAlertDefinitionAnomalyAlertSettingsModel

	result, err := logs.DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionLogsAnomalyTypeToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmLogsAlertDefinitionsApisAlertDefinitionLogsAnomalyRuleToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		apisAlertDefinitionLogsTimeWindowModel := make(map[string]interface{})
		apisAlertDefinitionLogsTimeWindowModel["logs_time_window_specific_value"] = "hours_36"

		apisAlertDefinitionLogsAnomalyConditionModel := make(map[string]interface{})
		apisAlertDefinitionLogsAnomalyConditionModel["minimum_threshold"] = float64(10.0)
		apisAlertDefinitionLogsAnomalyConditionModel["time_window"] = []map[string]interface{}{apisAlertDefinitionLogsTimeWindowModel}

		model := make(map[string]interface{})
		model["condition"] = []map[string]interface{}{apisAlertDefinitionLogsAnomalyConditionModel}

		assert.Equal(t, result, model)
	}

	apisAlertDefinitionLogsTimeWindowModel := new(logsv0.ApisAlertDefinitionLogsTimeWindow)
	apisAlertDefinitionLogsTimeWindowModel.LogsTimeWindowSpecificValue = core.StringPtr("hours_36")

	apisAlertDefinitionLogsAnomalyConditionModel := new(logsv0.ApisAlertDefinitionLogsAnomalyCondition)
	apisAlertDefinitionLogsAnomalyConditionModel.MinimumThreshold = core.Float64Ptr(float64(10.0))
	apisAlertDefinitionLogsAnomalyConditionModel.TimeWindow = apisAlertDefinitionLogsTimeWindowModel

	model := new(logsv0.ApisAlertDefinitionLogsAnomalyRule)
	model.Condition = apisAlertDefinitionLogsAnomalyConditionModel

	result, err := logs.DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionLogsAnomalyRuleToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmLogsAlertDefinitionsApisAlertDefinitionLogsAnomalyConditionToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		apisAlertDefinitionLogsTimeWindowModel := make(map[string]interface{})
		apisAlertDefinitionLogsTimeWindowModel["logs_time_window_specific_value"] = "hours_36"

		model := make(map[string]interface{})
		model["minimum_threshold"] = float64(10.0)
		model["time_window"] = []map[string]interface{}{apisAlertDefinitionLogsTimeWindowModel}

		assert.Equal(t, result, model)
	}

	apisAlertDefinitionLogsTimeWindowModel := new(logsv0.ApisAlertDefinitionLogsTimeWindow)
	apisAlertDefinitionLogsTimeWindowModel.LogsTimeWindowSpecificValue = core.StringPtr("hours_36")

	model := new(logsv0.ApisAlertDefinitionLogsAnomalyCondition)
	model.MinimumThreshold = core.Float64Ptr(float64(10.0))
	model.TimeWindow = apisAlertDefinitionLogsTimeWindowModel

	result, err := logs.DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionLogsAnomalyConditionToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmLogsAlertDefinitionsApisAlertDefinitionAnomalyAlertSettingsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["percentage_of_deviation"] = float64(10.0)

		assert.Equal(t, result, model)
	}

	model := new(logsv0.ApisAlertDefinitionAnomalyAlertSettings)
	model.PercentageOfDeviation = core.Float32Ptr(float32(10.0))

	result, err := logs.DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionAnomalyAlertSettingsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmLogsAlertDefinitionsApisAlertDefinitionMetricAnomalyTypeToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		apisAlertDefinitionMetricFilterModel := make(map[string]interface{})
		apisAlertDefinitionMetricFilterModel["promql"] = "avg_over_time(metric_name[5m]) > 10"

		apisAlertDefinitionMetricTimeWindowModel := make(map[string]interface{})
		apisAlertDefinitionMetricTimeWindowModel["metric_time_window_specific_value"] = "hours_36"

		apisAlertDefinitionMetricAnomalyConditionModel := make(map[string]interface{})
		apisAlertDefinitionMetricAnomalyConditionModel["threshold"] = float64(10.0)
		apisAlertDefinitionMetricAnomalyConditionModel["for_over_pct"] = int(20)
		apisAlertDefinitionMetricAnomalyConditionModel["of_the_last"] = []map[string]interface{}{apisAlertDefinitionMetricTimeWindowModel}
		apisAlertDefinitionMetricAnomalyConditionModel["min_non_null_values_pct"] = int(10)

		apisAlertDefinitionMetricAnomalyRuleModel := make(map[string]interface{})
		apisAlertDefinitionMetricAnomalyRuleModel["condition"] = []map[string]interface{}{apisAlertDefinitionMetricAnomalyConditionModel}

		apisAlertDefinitionAnomalyAlertSettingsModel := make(map[string]interface{})
		apisAlertDefinitionAnomalyAlertSettingsModel["percentage_of_deviation"] = float64(10.0)

		model := make(map[string]interface{})
		model["metric_filter"] = []map[string]interface{}{apisAlertDefinitionMetricFilterModel}
		model["rules"] = []map[string]interface{}{apisAlertDefinitionMetricAnomalyRuleModel}
		model["condition_type"] = "less_than_usual"
		model["evaluation_delay_ms"] = int(60000)
		model["anomaly_alert_settings"] = []map[string]interface{}{apisAlertDefinitionAnomalyAlertSettingsModel}

		assert.Equal(t, result, model)
	}

	apisAlertDefinitionMetricFilterModel := new(logsv0.ApisAlertDefinitionMetricFilter)
	apisAlertDefinitionMetricFilterModel.Promql = core.StringPtr("avg_over_time(metric_name[5m]) > 10")

	apisAlertDefinitionMetricTimeWindowModel := new(logsv0.ApisAlertDefinitionMetricTimeWindowTypeMetricTimeWindowSpecificValue)
	apisAlertDefinitionMetricTimeWindowModel.MetricTimeWindowSpecificValue = core.StringPtr("hours_36")

	apisAlertDefinitionMetricAnomalyConditionModel := new(logsv0.ApisAlertDefinitionMetricAnomalyCondition)
	apisAlertDefinitionMetricAnomalyConditionModel.Threshold = core.Float64Ptr(float64(10.0))
	apisAlertDefinitionMetricAnomalyConditionModel.ForOverPct = core.Int64Ptr(int64(20))
	apisAlertDefinitionMetricAnomalyConditionModel.OfTheLast = apisAlertDefinitionMetricTimeWindowModel
	apisAlertDefinitionMetricAnomalyConditionModel.MinNonNullValuesPct = core.Int64Ptr(int64(10))

	apisAlertDefinitionMetricAnomalyRuleModel := new(logsv0.ApisAlertDefinitionMetricAnomalyRule)
	apisAlertDefinitionMetricAnomalyRuleModel.Condition = apisAlertDefinitionMetricAnomalyConditionModel

	apisAlertDefinitionAnomalyAlertSettingsModel := new(logsv0.ApisAlertDefinitionAnomalyAlertSettings)
	apisAlertDefinitionAnomalyAlertSettingsModel.PercentageOfDeviation = core.Float32Ptr(float32(10.0))

	model := new(logsv0.ApisAlertDefinitionMetricAnomalyType)
	model.MetricFilter = apisAlertDefinitionMetricFilterModel
	model.Rules = []logsv0.ApisAlertDefinitionMetricAnomalyRule{*apisAlertDefinitionMetricAnomalyRuleModel}
	model.ConditionType = core.StringPtr("less_than_usual")
	model.EvaluationDelayMs = core.Int64Ptr(int64(60000))
	model.AnomalyAlertSettings = apisAlertDefinitionAnomalyAlertSettingsModel

	result, err := logs.DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionMetricAnomalyTypeToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmLogsAlertDefinitionsApisAlertDefinitionMetricAnomalyRuleToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		apisAlertDefinitionMetricTimeWindowModel := make(map[string]interface{})
		apisAlertDefinitionMetricTimeWindowModel["metric_time_window_specific_value"] = "hours_36"

		apisAlertDefinitionMetricAnomalyConditionModel := make(map[string]interface{})
		apisAlertDefinitionMetricAnomalyConditionModel["threshold"] = float64(10.0)
		apisAlertDefinitionMetricAnomalyConditionModel["for_over_pct"] = int(20)
		apisAlertDefinitionMetricAnomalyConditionModel["of_the_last"] = []map[string]interface{}{apisAlertDefinitionMetricTimeWindowModel}
		apisAlertDefinitionMetricAnomalyConditionModel["min_non_null_values_pct"] = int(10)

		model := make(map[string]interface{})
		model["condition"] = []map[string]interface{}{apisAlertDefinitionMetricAnomalyConditionModel}

		assert.Equal(t, result, model)
	}

	apisAlertDefinitionMetricTimeWindowModel := new(logsv0.ApisAlertDefinitionMetricTimeWindowTypeMetricTimeWindowSpecificValue)
	apisAlertDefinitionMetricTimeWindowModel.MetricTimeWindowSpecificValue = core.StringPtr("hours_36")

	apisAlertDefinitionMetricAnomalyConditionModel := new(logsv0.ApisAlertDefinitionMetricAnomalyCondition)
	apisAlertDefinitionMetricAnomalyConditionModel.Threshold = core.Float64Ptr(float64(10.0))
	apisAlertDefinitionMetricAnomalyConditionModel.ForOverPct = core.Int64Ptr(int64(20))
	apisAlertDefinitionMetricAnomalyConditionModel.OfTheLast = apisAlertDefinitionMetricTimeWindowModel
	apisAlertDefinitionMetricAnomalyConditionModel.MinNonNullValuesPct = core.Int64Ptr(int64(10))

	model := new(logsv0.ApisAlertDefinitionMetricAnomalyRule)
	model.Condition = apisAlertDefinitionMetricAnomalyConditionModel

	result, err := logs.DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionMetricAnomalyRuleToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmLogsAlertDefinitionsApisAlertDefinitionMetricAnomalyConditionToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		apisAlertDefinitionMetricTimeWindowModel := make(map[string]interface{})
		apisAlertDefinitionMetricTimeWindowModel["metric_time_window_specific_value"] = "hours_36"

		model := make(map[string]interface{})
		model["threshold"] = float64(10.0)
		model["for_over_pct"] = int(20)
		model["of_the_last"] = []map[string]interface{}{apisAlertDefinitionMetricTimeWindowModel}
		model["min_non_null_values_pct"] = int(10)

		assert.Equal(t, result, model)
	}

	apisAlertDefinitionMetricTimeWindowModel := new(logsv0.ApisAlertDefinitionMetricTimeWindowTypeMetricTimeWindowSpecificValue)
	apisAlertDefinitionMetricTimeWindowModel.MetricTimeWindowSpecificValue = core.StringPtr("hours_36")

	model := new(logsv0.ApisAlertDefinitionMetricAnomalyCondition)
	model.Threshold = core.Float64Ptr(float64(10.0))
	model.ForOverPct = core.Int64Ptr(int64(20))
	model.OfTheLast = apisAlertDefinitionMetricTimeWindowModel
	model.MinNonNullValuesPct = core.Int64Ptr(int64(10))

	result, err := logs.DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionMetricAnomalyConditionToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmLogsAlertDefinitionsApisAlertDefinitionLogsNewValueTypeToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		apisAlertDefinitionLabelFilterTypeModel := make(map[string]interface{})
		apisAlertDefinitionLabelFilterTypeModel["value"] = "my-app"
		apisAlertDefinitionLabelFilterTypeModel["operation"] = "starts_with"

		apisAlertDefinitionLabelFiltersModel := make(map[string]interface{})
		apisAlertDefinitionLabelFiltersModel["application_name"] = []map[string]interface{}{apisAlertDefinitionLabelFilterTypeModel}
		apisAlertDefinitionLabelFiltersModel["subsystem_name"] = []map[string]interface{}{apisAlertDefinitionLabelFilterTypeModel}
		apisAlertDefinitionLabelFiltersModel["severities"] = []string{"critical"}

		apisAlertDefinitionLogsSimpleFilterModel := make(map[string]interface{})
		apisAlertDefinitionLogsSimpleFilterModel["lucene_query"] = "text:"error""
		apisAlertDefinitionLogsSimpleFilterModel["label_filters"] = []map[string]interface{}{apisAlertDefinitionLabelFiltersModel}

		apisAlertDefinitionLogsFilterModel := make(map[string]interface{})
		apisAlertDefinitionLogsFilterModel["simple_filter"] = []map[string]interface{}{apisAlertDefinitionLogsSimpleFilterModel}

		apisAlertDefinitionLogsNewValueTimeWindowModel := make(map[string]interface{})
		apisAlertDefinitionLogsNewValueTimeWindowModel["logs_new_value_time_window_specific_value"] = "months_3"

		apisAlertDefinitionLogsNewValueConditionModel := make(map[string]interface{})
		apisAlertDefinitionLogsNewValueConditionModel["keypath_to_track"] = "metadata.field"
		apisAlertDefinitionLogsNewValueConditionModel["time_window"] = []map[string]interface{}{apisAlertDefinitionLogsNewValueTimeWindowModel}

		apisAlertDefinitionLogsNewValueRuleModel := make(map[string]interface{})
		apisAlertDefinitionLogsNewValueRuleModel["condition"] = []map[string]interface{}{apisAlertDefinitionLogsNewValueConditionModel}

		model := make(map[string]interface{})
		model["logs_filter"] = []map[string]interface{}{apisAlertDefinitionLogsFilterModel}
		model["rules"] = []map[string]interface{}{apisAlertDefinitionLogsNewValueRuleModel}
		model["notification_payload_filter"] = []string{"obj.field"}

		assert.Equal(t, result, model)
	}

	apisAlertDefinitionLabelFilterTypeModel := new(logsv0.ApisAlertDefinitionLabelFilterType)
	apisAlertDefinitionLabelFilterTypeModel.Value = core.StringPtr("my-app")
	apisAlertDefinitionLabelFilterTypeModel.Operation = core.StringPtr("starts_with")

	apisAlertDefinitionLabelFiltersModel := new(logsv0.ApisAlertDefinitionLabelFilters)
	apisAlertDefinitionLabelFiltersModel.ApplicationName = []logsv0.ApisAlertDefinitionLabelFilterType{*apisAlertDefinitionLabelFilterTypeModel}
	apisAlertDefinitionLabelFiltersModel.SubsystemName = []logsv0.ApisAlertDefinitionLabelFilterType{*apisAlertDefinitionLabelFilterTypeModel}
	apisAlertDefinitionLabelFiltersModel.Severities = []string{"critical"}

	apisAlertDefinitionLogsSimpleFilterModel := new(logsv0.ApisAlertDefinitionLogsSimpleFilter)
	apisAlertDefinitionLogsSimpleFilterModel.LuceneQuery = core.StringPtr("text:"error"")
	apisAlertDefinitionLogsSimpleFilterModel.LabelFilters = apisAlertDefinitionLabelFiltersModel

	apisAlertDefinitionLogsFilterModel := new(logsv0.ApisAlertDefinitionLogsFilter)
	apisAlertDefinitionLogsFilterModel.SimpleFilter = apisAlertDefinitionLogsSimpleFilterModel

	apisAlertDefinitionLogsNewValueTimeWindowModel := new(logsv0.ApisAlertDefinitionLogsNewValueTimeWindow)
	apisAlertDefinitionLogsNewValueTimeWindowModel.LogsNewValueTimeWindowSpecificValue = core.StringPtr("months_3")

	apisAlertDefinitionLogsNewValueConditionModel := new(logsv0.ApisAlertDefinitionLogsNewValueCondition)
	apisAlertDefinitionLogsNewValueConditionModel.KeypathToTrack = core.StringPtr("metadata.field")
	apisAlertDefinitionLogsNewValueConditionModel.TimeWindow = apisAlertDefinitionLogsNewValueTimeWindowModel

	apisAlertDefinitionLogsNewValueRuleModel := new(logsv0.ApisAlertDefinitionLogsNewValueRule)
	apisAlertDefinitionLogsNewValueRuleModel.Condition = apisAlertDefinitionLogsNewValueConditionModel

	model := new(logsv0.ApisAlertDefinitionLogsNewValueType)
	model.LogsFilter = apisAlertDefinitionLogsFilterModel
	model.Rules = []logsv0.ApisAlertDefinitionLogsNewValueRule{*apisAlertDefinitionLogsNewValueRuleModel}
	model.NotificationPayloadFilter = []string{"obj.field"}

	result, err := logs.DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionLogsNewValueTypeToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmLogsAlertDefinitionsApisAlertDefinitionLogsNewValueRuleToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		apisAlertDefinitionLogsNewValueTimeWindowModel := make(map[string]interface{})
		apisAlertDefinitionLogsNewValueTimeWindowModel["logs_new_value_time_window_specific_value"] = "months_3"

		apisAlertDefinitionLogsNewValueConditionModel := make(map[string]interface{})
		apisAlertDefinitionLogsNewValueConditionModel["keypath_to_track"] = "metadata.field"
		apisAlertDefinitionLogsNewValueConditionModel["time_window"] = []map[string]interface{}{apisAlertDefinitionLogsNewValueTimeWindowModel}

		model := make(map[string]interface{})
		model["condition"] = []map[string]interface{}{apisAlertDefinitionLogsNewValueConditionModel}

		assert.Equal(t, result, model)
	}

	apisAlertDefinitionLogsNewValueTimeWindowModel := new(logsv0.ApisAlertDefinitionLogsNewValueTimeWindow)
	apisAlertDefinitionLogsNewValueTimeWindowModel.LogsNewValueTimeWindowSpecificValue = core.StringPtr("months_3")

	apisAlertDefinitionLogsNewValueConditionModel := new(logsv0.ApisAlertDefinitionLogsNewValueCondition)
	apisAlertDefinitionLogsNewValueConditionModel.KeypathToTrack = core.StringPtr("metadata.field")
	apisAlertDefinitionLogsNewValueConditionModel.TimeWindow = apisAlertDefinitionLogsNewValueTimeWindowModel

	model := new(logsv0.ApisAlertDefinitionLogsNewValueRule)
	model.Condition = apisAlertDefinitionLogsNewValueConditionModel

	result, err := logs.DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionLogsNewValueRuleToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmLogsAlertDefinitionsApisAlertDefinitionLogsNewValueConditionToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		apisAlertDefinitionLogsNewValueTimeWindowModel := make(map[string]interface{})
		apisAlertDefinitionLogsNewValueTimeWindowModel["logs_new_value_time_window_specific_value"] = "months_3"

		model := make(map[string]interface{})
		model["keypath_to_track"] = "metadata.field"
		model["time_window"] = []map[string]interface{}{apisAlertDefinitionLogsNewValueTimeWindowModel}

		assert.Equal(t, result, model)
	}

	apisAlertDefinitionLogsNewValueTimeWindowModel := new(logsv0.ApisAlertDefinitionLogsNewValueTimeWindow)
	apisAlertDefinitionLogsNewValueTimeWindowModel.LogsNewValueTimeWindowSpecificValue = core.StringPtr("months_3")

	model := new(logsv0.ApisAlertDefinitionLogsNewValueCondition)
	model.KeypathToTrack = core.StringPtr("metadata.field")
	model.TimeWindow = apisAlertDefinitionLogsNewValueTimeWindowModel

	result, err := logs.DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionLogsNewValueConditionToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmLogsAlertDefinitionsApisAlertDefinitionLogsNewValueTimeWindowToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["logs_new_value_time_window_specific_value"] = "months_3"

		assert.Equal(t, result, model)
	}

	model := new(logsv0.ApisAlertDefinitionLogsNewValueTimeWindow)
	model.LogsNewValueTimeWindowSpecificValue = core.StringPtr("months_3")

	result, err := logs.DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionLogsNewValueTimeWindowToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmLogsAlertDefinitionsApisAlertDefinitionLogsUniqueCountTypeToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		apisAlertDefinitionLabelFilterTypeModel := make(map[string]interface{})
		apisAlertDefinitionLabelFilterTypeModel["value"] = "my-app"
		apisAlertDefinitionLabelFilterTypeModel["operation"] = "starts_with"

		apisAlertDefinitionLabelFiltersModel := make(map[string]interface{})
		apisAlertDefinitionLabelFiltersModel["application_name"] = []map[string]interface{}{apisAlertDefinitionLabelFilterTypeModel}
		apisAlertDefinitionLabelFiltersModel["subsystem_name"] = []map[string]interface{}{apisAlertDefinitionLabelFilterTypeModel}
		apisAlertDefinitionLabelFiltersModel["severities"] = []string{"critical"}

		apisAlertDefinitionLogsSimpleFilterModel := make(map[string]interface{})
		apisAlertDefinitionLogsSimpleFilterModel["lucene_query"] = "text:"error""
		apisAlertDefinitionLogsSimpleFilterModel["label_filters"] = []map[string]interface{}{apisAlertDefinitionLabelFiltersModel}

		apisAlertDefinitionLogsFilterModel := make(map[string]interface{})
		apisAlertDefinitionLogsFilterModel["simple_filter"] = []map[string]interface{}{apisAlertDefinitionLogsSimpleFilterModel}

		apisAlertDefinitionLogsUniqueValueTimeWindowModel := make(map[string]interface{})
		apisAlertDefinitionLogsUniqueValueTimeWindowModel["logs_unique_value_time_window_specific_value"] = "hours_36"

		apisAlertDefinitionLogsUniqueCountConditionModel := make(map[string]interface{})
		apisAlertDefinitionLogsUniqueCountConditionModel["max_unique_count"] = "100"
		apisAlertDefinitionLogsUniqueCountConditionModel["time_window"] = []map[string]interface{}{apisAlertDefinitionLogsUniqueValueTimeWindowModel}

		apisAlertDefinitionLogsUniqueCountRuleModel := make(map[string]interface{})
		apisAlertDefinitionLogsUniqueCountRuleModel["condition"] = []map[string]interface{}{apisAlertDefinitionLogsUniqueCountConditionModel}

		model := make(map[string]interface{})
		model["logs_filter"] = []map[string]interface{}{apisAlertDefinitionLogsFilterModel}
		model["rules"] = []map[string]interface{}{apisAlertDefinitionLogsUniqueCountRuleModel}
		model["notification_payload_filter"] = []string{"obj.field"}
		model["max_unique_count_per_group_by_key"] = "100"
		model["unique_count_keypath"] = "obj.field"

		assert.Equal(t, result, model)
	}

	apisAlertDefinitionLabelFilterTypeModel := new(logsv0.ApisAlertDefinitionLabelFilterType)
	apisAlertDefinitionLabelFilterTypeModel.Value = core.StringPtr("my-app")
	apisAlertDefinitionLabelFilterTypeModel.Operation = core.StringPtr("starts_with")

	apisAlertDefinitionLabelFiltersModel := new(logsv0.ApisAlertDefinitionLabelFilters)
	apisAlertDefinitionLabelFiltersModel.ApplicationName = []logsv0.ApisAlertDefinitionLabelFilterType{*apisAlertDefinitionLabelFilterTypeModel}
	apisAlertDefinitionLabelFiltersModel.SubsystemName = []logsv0.ApisAlertDefinitionLabelFilterType{*apisAlertDefinitionLabelFilterTypeModel}
	apisAlertDefinitionLabelFiltersModel.Severities = []string{"critical"}

	apisAlertDefinitionLogsSimpleFilterModel := new(logsv0.ApisAlertDefinitionLogsSimpleFilter)
	apisAlertDefinitionLogsSimpleFilterModel.LuceneQuery = core.StringPtr("text:"error"")
	apisAlertDefinitionLogsSimpleFilterModel.LabelFilters = apisAlertDefinitionLabelFiltersModel

	apisAlertDefinitionLogsFilterModel := new(logsv0.ApisAlertDefinitionLogsFilter)
	apisAlertDefinitionLogsFilterModel.SimpleFilter = apisAlertDefinitionLogsSimpleFilterModel

	apisAlertDefinitionLogsUniqueValueTimeWindowModel := new(logsv0.ApisAlertDefinitionLogsUniqueValueTimeWindow)
	apisAlertDefinitionLogsUniqueValueTimeWindowModel.LogsUniqueValueTimeWindowSpecificValue = core.StringPtr("hours_36")

	apisAlertDefinitionLogsUniqueCountConditionModel := new(logsv0.ApisAlertDefinitionLogsUniqueCountCondition)
	apisAlertDefinitionLogsUniqueCountConditionModel.MaxUniqueCount = core.StringPtr("100")
	apisAlertDefinitionLogsUniqueCountConditionModel.TimeWindow = apisAlertDefinitionLogsUniqueValueTimeWindowModel

	apisAlertDefinitionLogsUniqueCountRuleModel := new(logsv0.ApisAlertDefinitionLogsUniqueCountRule)
	apisAlertDefinitionLogsUniqueCountRuleModel.Condition = apisAlertDefinitionLogsUniqueCountConditionModel

	model := new(logsv0.ApisAlertDefinitionLogsUniqueCountType)
	model.LogsFilter = apisAlertDefinitionLogsFilterModel
	model.Rules = []logsv0.ApisAlertDefinitionLogsUniqueCountRule{*apisAlertDefinitionLogsUniqueCountRuleModel}
	model.NotificationPayloadFilter = []string{"obj.field"}
	model.MaxUniqueCountPerGroupByKey = core.StringPtr("100")
	model.UniqueCountKeypath = core.StringPtr("obj.field")

	result, err := logs.DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionLogsUniqueCountTypeToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmLogsAlertDefinitionsApisAlertDefinitionLogsUniqueCountRuleToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		apisAlertDefinitionLogsUniqueValueTimeWindowModel := make(map[string]interface{})
		apisAlertDefinitionLogsUniqueValueTimeWindowModel["logs_unique_value_time_window_specific_value"] = "hours_36"

		apisAlertDefinitionLogsUniqueCountConditionModel := make(map[string]interface{})
		apisAlertDefinitionLogsUniqueCountConditionModel["max_unique_count"] = "100"
		apisAlertDefinitionLogsUniqueCountConditionModel["time_window"] = []map[string]interface{}{apisAlertDefinitionLogsUniqueValueTimeWindowModel}

		model := make(map[string]interface{})
		model["condition"] = []map[string]interface{}{apisAlertDefinitionLogsUniqueCountConditionModel}

		assert.Equal(t, result, model)
	}

	apisAlertDefinitionLogsUniqueValueTimeWindowModel := new(logsv0.ApisAlertDefinitionLogsUniqueValueTimeWindow)
	apisAlertDefinitionLogsUniqueValueTimeWindowModel.LogsUniqueValueTimeWindowSpecificValue = core.StringPtr("hours_36")

	apisAlertDefinitionLogsUniqueCountConditionModel := new(logsv0.ApisAlertDefinitionLogsUniqueCountCondition)
	apisAlertDefinitionLogsUniqueCountConditionModel.MaxUniqueCount = core.StringPtr("100")
	apisAlertDefinitionLogsUniqueCountConditionModel.TimeWindow = apisAlertDefinitionLogsUniqueValueTimeWindowModel

	model := new(logsv0.ApisAlertDefinitionLogsUniqueCountRule)
	model.Condition = apisAlertDefinitionLogsUniqueCountConditionModel

	result, err := logs.DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionLogsUniqueCountRuleToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmLogsAlertDefinitionsApisAlertDefinitionLogsUniqueCountConditionToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		apisAlertDefinitionLogsUniqueValueTimeWindowModel := make(map[string]interface{})
		apisAlertDefinitionLogsUniqueValueTimeWindowModel["logs_unique_value_time_window_specific_value"] = "hours_36"

		model := make(map[string]interface{})
		model["max_unique_count"] = "100"
		model["time_window"] = []map[string]interface{}{apisAlertDefinitionLogsUniqueValueTimeWindowModel}

		assert.Equal(t, result, model)
	}

	apisAlertDefinitionLogsUniqueValueTimeWindowModel := new(logsv0.ApisAlertDefinitionLogsUniqueValueTimeWindow)
	apisAlertDefinitionLogsUniqueValueTimeWindowModel.LogsUniqueValueTimeWindowSpecificValue = core.StringPtr("hours_36")

	model := new(logsv0.ApisAlertDefinitionLogsUniqueCountCondition)
	model.MaxUniqueCount = core.StringPtr("100")
	model.TimeWindow = apisAlertDefinitionLogsUniqueValueTimeWindowModel

	result, err := logs.DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionLogsUniqueCountConditionToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmLogsAlertDefinitionsApisAlertDefinitionLogsUniqueValueTimeWindowToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["logs_unique_value_time_window_specific_value"] = "hours_36"

		assert.Equal(t, result, model)
	}

	model := new(logsv0.ApisAlertDefinitionLogsUniqueValueTimeWindow)
	model.LogsUniqueValueTimeWindowSpecificValue = core.StringPtr("hours_36")

	result, err := logs.DataSourceIbmLogsAlertDefinitionsApisAlertDefinitionLogsUniqueValueTimeWindowToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmLogsAlertDefinitionsAlertDefinitionApisAlertDefinitionAlertDefPropertiesTypeDefinitionLogsImmediateToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		apisAlertDefinitionTimeOfDayModel := make(map[string]interface{})
		apisAlertDefinitionTimeOfDayModel["hours"] = int(14)
		apisAlertDefinitionTimeOfDayModel["minutes"] = int(30)

		apisAlertDefinitionActivityScheduleModel := make(map[string]interface{})
		apisAlertDefinitionActivityScheduleModel["day_of_week"] = []string{"sunday"}
		apisAlertDefinitionActivityScheduleModel["start_time"] = []map[string]interface{}{apisAlertDefinitionTimeOfDayModel}
		apisAlertDefinitionActivityScheduleModel["end_time"] = []map[string]interface{}{apisAlertDefinitionTimeOfDayModel}

		apisAlertDefinitionAlertDefIncidentSettingsModel := make(map[string]interface{})
		apisAlertDefinitionAlertDefIncidentSettingsModel["notify_on"] = "triggered_and_resolved"
		apisAlertDefinitionAlertDefIncidentSettingsModel["minutes"] = int(30)

		apisAlertDefinitionIntegrationTypeModel := make(map[string]interface{})
		apisAlertDefinitionIntegrationTypeModel["integration_id"] = int(123)

		apisAlertDefinitionAlertDefWebhooksSettingsModel := make(map[string]interface{})
		apisAlertDefinitionAlertDefWebhooksSettingsModel["notify_on"] = "triggered_and_resolved"
		apisAlertDefinitionAlertDefWebhooksSettingsModel["integration"] = []map[string]interface{}{apisAlertDefinitionIntegrationTypeModel}
		apisAlertDefinitionAlertDefWebhooksSettingsModel["minutes"] = int(15)

		apisAlertDefinitionAlertDefNotificationGroupModel := make(map[string]interface{})
		apisAlertDefinitionAlertDefNotificationGroupModel["group_by_keys"] = []string{"key1", "key2"}
		apisAlertDefinitionAlertDefNotificationGroupModel["webhooks"] = []map[string]interface{}{apisAlertDefinitionAlertDefWebhooksSettingsModel}

		apisAlertDefinitionLabelFilterTypeModel := make(map[string]interface{})
		apisAlertDefinitionLabelFilterTypeModel["value"] = "my-app"
		apisAlertDefinitionLabelFilterTypeModel["operation"] = "starts_with"

		apisAlertDefinitionLabelFiltersModel := make(map[string]interface{})
		apisAlertDefinitionLabelFiltersModel["application_name"] = []map[string]interface{}{apisAlertDefinitionLabelFilterTypeModel}
		apisAlertDefinitionLabelFiltersModel["subsystem_name"] = []map[string]interface{}{apisAlertDefinitionLabelFilterTypeModel}
		apisAlertDefinitionLabelFiltersModel["severities"] = []string{"critical"}

		apisAlertDefinitionLogsSimpleFilterModel := make(map[string]interface{})
		apisAlertDefinitionLogsSimpleFilterModel["lucene_query"] = "text:"error""
		apisAlertDefinitionLogsSimpleFilterModel["label_filters"] = []map[string]interface{}{apisAlertDefinitionLabelFiltersModel}

		apisAlertDefinitionLogsFilterModel := make(map[string]interface{})
		apisAlertDefinitionLogsFilterModel["simple_filter"] = []map[string]interface{}{apisAlertDefinitionLogsSimpleFilterModel}

		apisAlertDefinitionLogsImmediateTypeModel := make(map[string]interface{})
		apisAlertDefinitionLogsImmediateTypeModel["logs_filter"] = []map[string]interface{}{apisAlertDefinitionLogsFilterModel}
		apisAlertDefinitionLogsImmediateTypeModel["notification_payload_filter"] = []string{"obj.field"}

		model := make(map[string]interface{})
		model["id"] = "3dc02998-0b50-4ea8-b68a-4779d716fa1f"
		model["created_time"] = "2021-01-01T00:00:00.000Z"
		model["updated_time"] = "2021-01-01T00:00:00.000Z"
		model["alert_version_id"] = "3dc02998-0b50-4ea8-b68a-4779d716fa1f"
		model["name"] = "Unique count alert"
		model["description"] = "Example of unique count alert from terraform"
		model["enabled"] = true
		model["priority"] = "p1"
		model["active_on"] = []map[string]interface{}{apisAlertDefinitionActivityScheduleModel}
		model["type"] = "flow"
		model["group_by_keys"] = []string{"key1", "key2"}
		model["incidents_settings"] = []map[string]interface{}{apisAlertDefinitionAlertDefIncidentSettingsModel}
		model["notification_group"] = []map[string]interface{}{apisAlertDefinitionAlertDefNotificationGroupModel}
		model["entity_labels"] = map[string]interface{}{"key1": "testString"}
		model["phantom_mode"] = false
		model["deleted"] = false
		model["logs_immediate"] = []map[string]interface{}{apisAlertDefinitionLogsImmediateTypeModel}

		assert.Equal(t, result, model)
	}

	apisAlertDefinitionTimeOfDayModel := new(logsv0.ApisAlertDefinitionTimeOfDay)
	apisAlertDefinitionTimeOfDayModel.Hours = core.Int64Ptr(int64(14))
	apisAlertDefinitionTimeOfDayModel.Minutes = core.Int64Ptr(int64(30))

	apisAlertDefinitionActivityScheduleModel := new(logsv0.ApisAlertDefinitionActivitySchedule)
	apisAlertDefinitionActivityScheduleModel.DayOfWeek = []string{"sunday"}
	apisAlertDefinitionActivityScheduleModel.StartTime = apisAlertDefinitionTimeOfDayModel
	apisAlertDefinitionActivityScheduleModel.EndTime = apisAlertDefinitionTimeOfDayModel

	apisAlertDefinitionAlertDefIncidentSettingsModel := new(logsv0.ApisAlertDefinitionAlertDefIncidentSettings)
	apisAlertDefinitionAlertDefIncidentSettingsModel.NotifyOn = core.StringPtr("triggered_and_resolved")
	apisAlertDefinitionAlertDefIncidentSettingsModel.Minutes = core.Int64Ptr(int64(30))

	apisAlertDefinitionIntegrationTypeModel := new(logsv0.ApisAlertDefinitionIntegrationTypeIntegrationTypeIntegrationID)
	apisAlertDefinitionIntegrationTypeModel.IntegrationID = core.Int64Ptr(int64(123))

	apisAlertDefinitionAlertDefWebhooksSettingsModel := new(logsv0.ApisAlertDefinitionAlertDefWebhooksSettings)
	apisAlertDefinitionAlertDefWebhooksSettingsModel.NotifyOn = core.StringPtr("triggered_and_resolved")
	apisAlertDefinitionAlertDefWebhooksSettingsModel.Integration = apisAlertDefinitionIntegrationTypeModel
	apisAlertDefinitionAlertDefWebhooksSettingsModel.Minutes = core.Int64Ptr(int64(15))

	apisAlertDefinitionAlertDefNotificationGroupModel := new(logsv0.ApisAlertDefinitionAlertDefNotificationGroup)
	apisAlertDefinitionAlertDefNotificationGroupModel.GroupByKeys = []string{"key1", "key2"}
	apisAlertDefinitionAlertDefNotificationGroupModel.Webhooks = []logsv0.ApisAlertDefinitionAlertDefWebhooksSettings{*apisAlertDefinitionAlertDefWebhooksSettingsModel}

	apisAlertDefinitionLabelFilterTypeModel := new(logsv0.ApisAlertDefinitionLabelFilterType)
	apisAlertDefinitionLabelFilterTypeModel.Value = core.StringPtr("my-app")
	apisAlertDefinitionLabelFilterTypeModel.Operation = core.StringPtr("starts_with")

	apisAlertDefinitionLabelFiltersModel := new(logsv0.ApisAlertDefinitionLabelFilters)
	apisAlertDefinitionLabelFiltersModel.ApplicationName = []logsv0.ApisAlertDefinitionLabelFilterType{*apisAlertDefinitionLabelFilterTypeModel}
	apisAlertDefinitionLabelFiltersModel.SubsystemName = []logsv0.ApisAlertDefinitionLabelFilterType{*apisAlertDefinitionLabelFilterTypeModel}
	apisAlertDefinitionLabelFiltersModel.Severities = []string{"critical"}

	apisAlertDefinitionLogsSimpleFilterModel := new(logsv0.ApisAlertDefinitionLogsSimpleFilter)
	apisAlertDefinitionLogsSimpleFilterModel.LuceneQuery = core.StringPtr("text:"error"")
	apisAlertDefinitionLogsSimpleFilterModel.LabelFilters = apisAlertDefinitionLabelFiltersModel

	apisAlertDefinitionLogsFilterModel := new(logsv0.ApisAlertDefinitionLogsFilter)
	apisAlertDefinitionLogsFilterModel.SimpleFilter = apisAlertDefinitionLogsSimpleFilterModel

	apisAlertDefinitionLogsImmediateTypeModel := new(logsv0.ApisAlertDefinitionLogsImmediateType)
	apisAlertDefinitionLogsImmediateTypeModel.LogsFilter = apisAlertDefinitionLogsFilterModel
	apisAlertDefinitionLogsImmediateTypeModel.NotificationPayloadFilter = []string{"obj.field"}

	model := new(logsv0.AlertDefinitionApisAlertDefinitionAlertDefPropertiesTypeDefinitionLogsImmediate)
	model.ID = CreateMockUUID("3dc02998-0b50-4ea8-b68a-4779d716fa1f")
	model.CreatedTime = CreateMockDateTime("2021-01-01T00:00:00.000Z")
	model.UpdatedTime = CreateMockDateTime("2021-01-01T00:00:00.000Z")
	model.AlertVersionID = CreateMockUUID("3dc02998-0b50-4ea8-b68a-4779d716fa1f")
	model.Name = core.StringPtr("Unique count alert")
	model.Description = core.StringPtr("Example of unique count alert from terraform")
	model.Enabled = core.BoolPtr(true)
	model.Priority = core.StringPtr("p1")
	model.ActiveOn = apisAlertDefinitionActivityScheduleModel
	model.Type = core.StringPtr("flow")
	model.GroupByKeys = []string{"key1", "key2"}
	model.IncidentsSettings = apisAlertDefinitionAlertDefIncidentSettingsModel
	model.NotificationGroup = apisAlertDefinitionAlertDefNotificationGroupModel
	model.EntityLabels = map[string]string{"key1": "testString"}
	model.PhantomMode = core.BoolPtr(false)
	model.Deleted = core.BoolPtr(false)
	model.LogsImmediate = apisAlertDefinitionLogsImmediateTypeModel

	result, err := logs.DataSourceIbmLogsAlertDefinitionsAlertDefinitionApisAlertDefinitionAlertDefPropertiesTypeDefinitionLogsImmediateToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmLogsAlertDefinitionsAlertDefinitionApisAlertDefinitionAlertDefPropertiesTypeDefinitionLogsThresholdToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		apisAlertDefinitionTimeOfDayModel := make(map[string]interface{})
		apisAlertDefinitionTimeOfDayModel["hours"] = int(14)
		apisAlertDefinitionTimeOfDayModel["minutes"] = int(30)

		apisAlertDefinitionActivityScheduleModel := make(map[string]interface{})
		apisAlertDefinitionActivityScheduleModel["day_of_week"] = []string{"sunday"}
		apisAlertDefinitionActivityScheduleModel["start_time"] = []map[string]interface{}{apisAlertDefinitionTimeOfDayModel}
		apisAlertDefinitionActivityScheduleModel["end_time"] = []map[string]interface{}{apisAlertDefinitionTimeOfDayModel}

		apisAlertDefinitionAlertDefIncidentSettingsModel := make(map[string]interface{})
		apisAlertDefinitionAlertDefIncidentSettingsModel["notify_on"] = "triggered_and_resolved"
		apisAlertDefinitionAlertDefIncidentSettingsModel["minutes"] = int(30)

		apisAlertDefinitionIntegrationTypeModel := make(map[string]interface{})
		apisAlertDefinitionIntegrationTypeModel["integration_id"] = int(123)

		apisAlertDefinitionAlertDefWebhooksSettingsModel := make(map[string]interface{})
		apisAlertDefinitionAlertDefWebhooksSettingsModel["notify_on"] = "triggered_and_resolved"
		apisAlertDefinitionAlertDefWebhooksSettingsModel["integration"] = []map[string]interface{}{apisAlertDefinitionIntegrationTypeModel}
		apisAlertDefinitionAlertDefWebhooksSettingsModel["minutes"] = int(15)

		apisAlertDefinitionAlertDefNotificationGroupModel := make(map[string]interface{})
		apisAlertDefinitionAlertDefNotificationGroupModel["group_by_keys"] = []string{"key1", "key2"}
		apisAlertDefinitionAlertDefNotificationGroupModel["webhooks"] = []map[string]interface{}{apisAlertDefinitionAlertDefWebhooksSettingsModel}

		apisAlertDefinitionLabelFilterTypeModel := make(map[string]interface{})
		apisAlertDefinitionLabelFilterTypeModel["value"] = "my-app"
		apisAlertDefinitionLabelFilterTypeModel["operation"] = "starts_with"

		apisAlertDefinitionLabelFiltersModel := make(map[string]interface{})
		apisAlertDefinitionLabelFiltersModel["application_name"] = []map[string]interface{}{apisAlertDefinitionLabelFilterTypeModel}
		apisAlertDefinitionLabelFiltersModel["subsystem_name"] = []map[string]interface{}{apisAlertDefinitionLabelFilterTypeModel}
		apisAlertDefinitionLabelFiltersModel["severities"] = []string{"critical"}

		apisAlertDefinitionLogsSimpleFilterModel := make(map[string]interface{})
		apisAlertDefinitionLogsSimpleFilterModel["lucene_query"] = "text:"error""
		apisAlertDefinitionLogsSimpleFilterModel["label_filters"] = []map[string]interface{}{apisAlertDefinitionLabelFiltersModel}

		apisAlertDefinitionLogsFilterModel := make(map[string]interface{})
		apisAlertDefinitionLogsFilterModel["simple_filter"] = []map[string]interface{}{apisAlertDefinitionLogsSimpleFilterModel}

		apisAlertDefinitionUndetectedValuesManagementModel := make(map[string]interface{})
		apisAlertDefinitionUndetectedValuesManagementModel["trigger_undetected_values"] = true
		apisAlertDefinitionUndetectedValuesManagementModel["auto_retire_timeframe"] = "hours_24"

		apisAlertDefinitionLogsTimeWindowModel := make(map[string]interface{})
		apisAlertDefinitionLogsTimeWindowModel["logs_time_window_specific_value"] = "hours_36"

		apisAlertDefinitionLogsThresholdConditionModel := make(map[string]interface{})
		apisAlertDefinitionLogsThresholdConditionModel["threshold"] = float64(100.0)
		apisAlertDefinitionLogsThresholdConditionModel["time_window"] = []map[string]interface{}{apisAlertDefinitionLogsTimeWindowModel}

		apisAlertDefinitionAlertDefOverrideModel := make(map[string]interface{})
		apisAlertDefinitionAlertDefOverrideModel["priority"] = "p1"

		apisAlertDefinitionLogsThresholdRuleModel := make(map[string]interface{})
		apisAlertDefinitionLogsThresholdRuleModel["condition"] = []map[string]interface{}{apisAlertDefinitionLogsThresholdConditionModel}
		apisAlertDefinitionLogsThresholdRuleModel["override"] = []map[string]interface{}{apisAlertDefinitionAlertDefOverrideModel}

		apisAlertDefinitionLogsThresholdTypeModel := make(map[string]interface{})
		apisAlertDefinitionLogsThresholdTypeModel["logs_filter"] = []map[string]interface{}{apisAlertDefinitionLogsFilterModel}
		apisAlertDefinitionLogsThresholdTypeModel["undetected_values_management"] = []map[string]interface{}{apisAlertDefinitionUndetectedValuesManagementModel}
		apisAlertDefinitionLogsThresholdTypeModel["rules"] = []map[string]interface{}{apisAlertDefinitionLogsThresholdRuleModel}
		apisAlertDefinitionLogsThresholdTypeModel["condition_type"] = "less_than"
		apisAlertDefinitionLogsThresholdTypeModel["notification_payload_filter"] = []string{"obj.field"}
		apisAlertDefinitionLogsThresholdTypeModel["evaluation_delay_ms"] = int(60000)

		model := make(map[string]interface{})
		model["id"] = "3dc02998-0b50-4ea8-b68a-4779d716fa1f"
		model["created_time"] = "2021-01-01T00:00:00.000Z"
		model["updated_time"] = "2021-01-01T00:00:00.000Z"
		model["alert_version_id"] = "3dc02998-0b50-4ea8-b68a-4779d716fa1f"
		model["name"] = "Unique count alert"
		model["description"] = "Example of unique count alert from terraform"
		model["enabled"] = true
		model["priority"] = "p1"
		model["active_on"] = []map[string]interface{}{apisAlertDefinitionActivityScheduleModel}
		model["type"] = "flow"
		model["group_by_keys"] = []string{"key1", "key2"}
		model["incidents_settings"] = []map[string]interface{}{apisAlertDefinitionAlertDefIncidentSettingsModel}
		model["notification_group"] = []map[string]interface{}{apisAlertDefinitionAlertDefNotificationGroupModel}
		model["entity_labels"] = map[string]interface{}{"key1": "testString"}
		model["phantom_mode"] = false
		model["deleted"] = false
		model["logs_threshold"] = []map[string]interface{}{apisAlertDefinitionLogsThresholdTypeModel}

		assert.Equal(t, result, model)
	}

	apisAlertDefinitionTimeOfDayModel := new(logsv0.ApisAlertDefinitionTimeOfDay)
	apisAlertDefinitionTimeOfDayModel.Hours = core.Int64Ptr(int64(14))
	apisAlertDefinitionTimeOfDayModel.Minutes = core.Int64Ptr(int64(30))

	apisAlertDefinitionActivityScheduleModel := new(logsv0.ApisAlertDefinitionActivitySchedule)
	apisAlertDefinitionActivityScheduleModel.DayOfWeek = []string{"sunday"}
	apisAlertDefinitionActivityScheduleModel.StartTime = apisAlertDefinitionTimeOfDayModel
	apisAlertDefinitionActivityScheduleModel.EndTime = apisAlertDefinitionTimeOfDayModel

	apisAlertDefinitionAlertDefIncidentSettingsModel := new(logsv0.ApisAlertDefinitionAlertDefIncidentSettings)
	apisAlertDefinitionAlertDefIncidentSettingsModel.NotifyOn = core.StringPtr("triggered_and_resolved")
	apisAlertDefinitionAlertDefIncidentSettingsModel.Minutes = core.Int64Ptr(int64(30))

	apisAlertDefinitionIntegrationTypeModel := new(logsv0.ApisAlertDefinitionIntegrationTypeIntegrationTypeIntegrationID)
	apisAlertDefinitionIntegrationTypeModel.IntegrationID = core.Int64Ptr(int64(123))

	apisAlertDefinitionAlertDefWebhooksSettingsModel := new(logsv0.ApisAlertDefinitionAlertDefWebhooksSettings)
	apisAlertDefinitionAlertDefWebhooksSettingsModel.NotifyOn = core.StringPtr("triggered_and_resolved")
	apisAlertDefinitionAlertDefWebhooksSettingsModel.Integration = apisAlertDefinitionIntegrationTypeModel
	apisAlertDefinitionAlertDefWebhooksSettingsModel.Minutes = core.Int64Ptr(int64(15))

	apisAlertDefinitionAlertDefNotificationGroupModel := new(logsv0.ApisAlertDefinitionAlertDefNotificationGroup)
	apisAlertDefinitionAlertDefNotificationGroupModel.GroupByKeys = []string{"key1", "key2"}
	apisAlertDefinitionAlertDefNotificationGroupModel.Webhooks = []logsv0.ApisAlertDefinitionAlertDefWebhooksSettings{*apisAlertDefinitionAlertDefWebhooksSettingsModel}

	apisAlertDefinitionLabelFilterTypeModel := new(logsv0.ApisAlertDefinitionLabelFilterType)
	apisAlertDefinitionLabelFilterTypeModel.Value = core.StringPtr("my-app")
	apisAlertDefinitionLabelFilterTypeModel.Operation = core.StringPtr("starts_with")

	apisAlertDefinitionLabelFiltersModel := new(logsv0.ApisAlertDefinitionLabelFilters)
	apisAlertDefinitionLabelFiltersModel.ApplicationName = []logsv0.ApisAlertDefinitionLabelFilterType{*apisAlertDefinitionLabelFilterTypeModel}
	apisAlertDefinitionLabelFiltersModel.SubsystemName = []logsv0.ApisAlertDefinitionLabelFilterType{*apisAlertDefinitionLabelFilterTypeModel}
	apisAlertDefinitionLabelFiltersModel.Severities = []string{"critical"}

	apisAlertDefinitionLogsSimpleFilterModel := new(logsv0.ApisAlertDefinitionLogsSimpleFilter)
	apisAlertDefinitionLogsSimpleFilterModel.LuceneQuery = core.StringPtr("text:"error"")
	apisAlertDefinitionLogsSimpleFilterModel.LabelFilters = apisAlertDefinitionLabelFiltersModel

	apisAlertDefinitionLogsFilterModel := new(logsv0.ApisAlertDefinitionLogsFilter)
	apisAlertDefinitionLogsFilterModel.SimpleFilter = apisAlertDefinitionLogsSimpleFilterModel

	apisAlertDefinitionUndetectedValuesManagementModel := new(logsv0.ApisAlertDefinitionUndetectedValuesManagement)
	apisAlertDefinitionUndetectedValuesManagementModel.TriggerUndetectedValues = core.BoolPtr(true)
	apisAlertDefinitionUndetectedValuesManagementModel.AutoRetireTimeframe = core.StringPtr("hours_24")

	apisAlertDefinitionLogsTimeWindowModel := new(logsv0.ApisAlertDefinitionLogsTimeWindow)
	apisAlertDefinitionLogsTimeWindowModel.LogsTimeWindowSpecificValue = core.StringPtr("hours_36")

	apisAlertDefinitionLogsThresholdConditionModel := new(logsv0.ApisAlertDefinitionLogsThresholdCondition)
	apisAlertDefinitionLogsThresholdConditionModel.Threshold = core.Float64Ptr(float64(100.0))
	apisAlertDefinitionLogsThresholdConditionModel.TimeWindow = apisAlertDefinitionLogsTimeWindowModel

	apisAlertDefinitionAlertDefOverrideModel := new(logsv0.ApisAlertDefinitionAlertDefOverride)
	apisAlertDefinitionAlertDefOverrideModel.Priority = core.StringPtr("p1")

	apisAlertDefinitionLogsThresholdRuleModel := new(logsv0.ApisAlertDefinitionLogsThresholdRule)
	apisAlertDefinitionLogsThresholdRuleModel.Condition = apisAlertDefinitionLogsThresholdConditionModel
	apisAlertDefinitionLogsThresholdRuleModel.Override = apisAlertDefinitionAlertDefOverrideModel

	apisAlertDefinitionLogsThresholdTypeModel := new(logsv0.ApisAlertDefinitionLogsThresholdType)
	apisAlertDefinitionLogsThresholdTypeModel.LogsFilter = apisAlertDefinitionLogsFilterModel
	apisAlertDefinitionLogsThresholdTypeModel.UndetectedValuesManagement = apisAlertDefinitionUndetectedValuesManagementModel
	apisAlertDefinitionLogsThresholdTypeModel.Rules = []logsv0.ApisAlertDefinitionLogsThresholdRule{*apisAlertDefinitionLogsThresholdRuleModel}
	apisAlertDefinitionLogsThresholdTypeModel.ConditionType = core.StringPtr("less_than")
	apisAlertDefinitionLogsThresholdTypeModel.NotificationPayloadFilter = []string{"obj.field"}
	apisAlertDefinitionLogsThresholdTypeModel.EvaluationDelayMs = core.Int64Ptr(int64(60000))

	model := new(logsv0.AlertDefinitionApisAlertDefinitionAlertDefPropertiesTypeDefinitionLogsThreshold)
	model.ID = CreateMockUUID("3dc02998-0b50-4ea8-b68a-4779d716fa1f")
	model.CreatedTime = CreateMockDateTime("2021-01-01T00:00:00.000Z")
	model.UpdatedTime = CreateMockDateTime("2021-01-01T00:00:00.000Z")
	model.AlertVersionID = CreateMockUUID("3dc02998-0b50-4ea8-b68a-4779d716fa1f")
	model.Name = core.StringPtr("Unique count alert")
	model.Description = core.StringPtr("Example of unique count alert from terraform")
	model.Enabled = core.BoolPtr(true)
	model.Priority = core.StringPtr("p1")
	model.ActiveOn = apisAlertDefinitionActivityScheduleModel
	model.Type = core.StringPtr("flow")
	model.GroupByKeys = []string{"key1", "key2"}
	model.IncidentsSettings = apisAlertDefinitionAlertDefIncidentSettingsModel
	model.NotificationGroup = apisAlertDefinitionAlertDefNotificationGroupModel
	model.EntityLabels = map[string]string{"key1": "testString"}
	model.PhantomMode = core.BoolPtr(false)
	model.Deleted = core.BoolPtr(false)
	model.LogsThreshold = apisAlertDefinitionLogsThresholdTypeModel

	result, err := logs.DataSourceIbmLogsAlertDefinitionsAlertDefinitionApisAlertDefinitionAlertDefPropertiesTypeDefinitionLogsThresholdToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmLogsAlertDefinitionsAlertDefinitionApisAlertDefinitionAlertDefPropertiesTypeDefinitionLogsRatioThresholdToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		apisAlertDefinitionTimeOfDayModel := make(map[string]interface{})
		apisAlertDefinitionTimeOfDayModel["hours"] = int(14)
		apisAlertDefinitionTimeOfDayModel["minutes"] = int(30)

		apisAlertDefinitionActivityScheduleModel := make(map[string]interface{})
		apisAlertDefinitionActivityScheduleModel["day_of_week"] = []string{"sunday"}
		apisAlertDefinitionActivityScheduleModel["start_time"] = []map[string]interface{}{apisAlertDefinitionTimeOfDayModel}
		apisAlertDefinitionActivityScheduleModel["end_time"] = []map[string]interface{}{apisAlertDefinitionTimeOfDayModel}

		apisAlertDefinitionAlertDefIncidentSettingsModel := make(map[string]interface{})
		apisAlertDefinitionAlertDefIncidentSettingsModel["notify_on"] = "triggered_and_resolved"
		apisAlertDefinitionAlertDefIncidentSettingsModel["minutes"] = int(30)

		apisAlertDefinitionIntegrationTypeModel := make(map[string]interface{})
		apisAlertDefinitionIntegrationTypeModel["integration_id"] = int(123)

		apisAlertDefinitionAlertDefWebhooksSettingsModel := make(map[string]interface{})
		apisAlertDefinitionAlertDefWebhooksSettingsModel["notify_on"] = "triggered_and_resolved"
		apisAlertDefinitionAlertDefWebhooksSettingsModel["integration"] = []map[string]interface{}{apisAlertDefinitionIntegrationTypeModel}
		apisAlertDefinitionAlertDefWebhooksSettingsModel["minutes"] = int(15)

		apisAlertDefinitionAlertDefNotificationGroupModel := make(map[string]interface{})
		apisAlertDefinitionAlertDefNotificationGroupModel["group_by_keys"] = []string{"key1", "key2"}
		apisAlertDefinitionAlertDefNotificationGroupModel["webhooks"] = []map[string]interface{}{apisAlertDefinitionAlertDefWebhooksSettingsModel}

		apisAlertDefinitionLabelFilterTypeModel := make(map[string]interface{})
		apisAlertDefinitionLabelFilterTypeModel["value"] = "my-app"
		apisAlertDefinitionLabelFilterTypeModel["operation"] = "starts_with"

		apisAlertDefinitionLabelFiltersModel := make(map[string]interface{})
		apisAlertDefinitionLabelFiltersModel["application_name"] = []map[string]interface{}{apisAlertDefinitionLabelFilterTypeModel}
		apisAlertDefinitionLabelFiltersModel["subsystem_name"] = []map[string]interface{}{apisAlertDefinitionLabelFilterTypeModel}
		apisAlertDefinitionLabelFiltersModel["severities"] = []string{"critical"}

		apisAlertDefinitionLogsSimpleFilterModel := make(map[string]interface{})
		apisAlertDefinitionLogsSimpleFilterModel["lucene_query"] = "text:"error""
		apisAlertDefinitionLogsSimpleFilterModel["label_filters"] = []map[string]interface{}{apisAlertDefinitionLabelFiltersModel}

		apisAlertDefinitionLogsFilterModel := make(map[string]interface{})
		apisAlertDefinitionLogsFilterModel["simple_filter"] = []map[string]interface{}{apisAlertDefinitionLogsSimpleFilterModel}

		apisAlertDefinitionLogsRatioTimeWindowModel := make(map[string]interface{})
		apisAlertDefinitionLogsRatioTimeWindowModel["logs_ratio_time_window_specific_value"] = "hours_36"

		apisAlertDefinitionLogsRatioConditionModel := make(map[string]interface{})
		apisAlertDefinitionLogsRatioConditionModel["threshold"] = float64(10.0)
		apisAlertDefinitionLogsRatioConditionModel["time_window"] = []map[string]interface{}{apisAlertDefinitionLogsRatioTimeWindowModel}

		apisAlertDefinitionAlertDefOverrideModel := make(map[string]interface{})
		apisAlertDefinitionAlertDefOverrideModel["priority"] = "p1"

		apisAlertDefinitionLogsRatioRulesModel := make(map[string]interface{})
		apisAlertDefinitionLogsRatioRulesModel["condition"] = []map[string]interface{}{apisAlertDefinitionLogsRatioConditionModel}
		apisAlertDefinitionLogsRatioRulesModel["override"] = []map[string]interface{}{apisAlertDefinitionAlertDefOverrideModel}

		apisAlertDefinitionUndetectedValuesManagementModel := make(map[string]interface{})
		apisAlertDefinitionUndetectedValuesManagementModel["trigger_undetected_values"] = true
		apisAlertDefinitionUndetectedValuesManagementModel["auto_retire_timeframe"] = "hours_24"

		apisAlertDefinitionLogsRatioThresholdTypeModel := make(map[string]interface{})
		apisAlertDefinitionLogsRatioThresholdTypeModel["numerator"] = []map[string]interface{}{apisAlertDefinitionLogsFilterModel}
		apisAlertDefinitionLogsRatioThresholdTypeModel["numerator_alias"] = "numerator_alias"
		apisAlertDefinitionLogsRatioThresholdTypeModel["denominator"] = []map[string]interface{}{apisAlertDefinitionLogsFilterModel}
		apisAlertDefinitionLogsRatioThresholdTypeModel["denominator_alias"] = "denominator_alias"
		apisAlertDefinitionLogsRatioThresholdTypeModel["rules"] = []map[string]interface{}{apisAlertDefinitionLogsRatioRulesModel}
		apisAlertDefinitionLogsRatioThresholdTypeModel["condition_type"] = "less_than"
		apisAlertDefinitionLogsRatioThresholdTypeModel["notification_payload_filter"] = []string{"obj.field"}
		apisAlertDefinitionLogsRatioThresholdTypeModel["group_by_for"] = "denumerator_only"
		apisAlertDefinitionLogsRatioThresholdTypeModel["undetected_values_management"] = []map[string]interface{}{apisAlertDefinitionUndetectedValuesManagementModel}
		apisAlertDefinitionLogsRatioThresholdTypeModel["ignore_infinity"] = true
		apisAlertDefinitionLogsRatioThresholdTypeModel["evaluation_delay_ms"] = int(60000)

		model := make(map[string]interface{})
		model["id"] = "3dc02998-0b50-4ea8-b68a-4779d716fa1f"
		model["created_time"] = "2021-01-01T00:00:00.000Z"
		model["updated_time"] = "2021-01-01T00:00:00.000Z"
		model["alert_version_id"] = "3dc02998-0b50-4ea8-b68a-4779d716fa1f"
		model["name"] = "Unique count alert"
		model["description"] = "Example of unique count alert from terraform"
		model["enabled"] = true
		model["priority"] = "p1"
		model["active_on"] = []map[string]interface{}{apisAlertDefinitionActivityScheduleModel}
		model["type"] = "flow"
		model["group_by_keys"] = []string{"key1", "key2"}
		model["incidents_settings"] = []map[string]interface{}{apisAlertDefinitionAlertDefIncidentSettingsModel}
		model["notification_group"] = []map[string]interface{}{apisAlertDefinitionAlertDefNotificationGroupModel}
		model["entity_labels"] = map[string]interface{}{"key1": "testString"}
		model["phantom_mode"] = false
		model["deleted"] = false
		model["logs_ratio_threshold"] = []map[string]interface{}{apisAlertDefinitionLogsRatioThresholdTypeModel}

		assert.Equal(t, result, model)
	}

	apisAlertDefinitionTimeOfDayModel := new(logsv0.ApisAlertDefinitionTimeOfDay)
	apisAlertDefinitionTimeOfDayModel.Hours = core.Int64Ptr(int64(14))
	apisAlertDefinitionTimeOfDayModel.Minutes = core.Int64Ptr(int64(30))

	apisAlertDefinitionActivityScheduleModel := new(logsv0.ApisAlertDefinitionActivitySchedule)
	apisAlertDefinitionActivityScheduleModel.DayOfWeek = []string{"sunday"}
	apisAlertDefinitionActivityScheduleModel.StartTime = apisAlertDefinitionTimeOfDayModel
	apisAlertDefinitionActivityScheduleModel.EndTime = apisAlertDefinitionTimeOfDayModel

	apisAlertDefinitionAlertDefIncidentSettingsModel := new(logsv0.ApisAlertDefinitionAlertDefIncidentSettings)
	apisAlertDefinitionAlertDefIncidentSettingsModel.NotifyOn = core.StringPtr("triggered_and_resolved")
	apisAlertDefinitionAlertDefIncidentSettingsModel.Minutes = core.Int64Ptr(int64(30))

	apisAlertDefinitionIntegrationTypeModel := new(logsv0.ApisAlertDefinitionIntegrationTypeIntegrationTypeIntegrationID)
	apisAlertDefinitionIntegrationTypeModel.IntegrationID = core.Int64Ptr(int64(123))

	apisAlertDefinitionAlertDefWebhooksSettingsModel := new(logsv0.ApisAlertDefinitionAlertDefWebhooksSettings)
	apisAlertDefinitionAlertDefWebhooksSettingsModel.NotifyOn = core.StringPtr("triggered_and_resolved")
	apisAlertDefinitionAlertDefWebhooksSettingsModel.Integration = apisAlertDefinitionIntegrationTypeModel
	apisAlertDefinitionAlertDefWebhooksSettingsModel.Minutes = core.Int64Ptr(int64(15))

	apisAlertDefinitionAlertDefNotificationGroupModel := new(logsv0.ApisAlertDefinitionAlertDefNotificationGroup)
	apisAlertDefinitionAlertDefNotificationGroupModel.GroupByKeys = []string{"key1", "key2"}
	apisAlertDefinitionAlertDefNotificationGroupModel.Webhooks = []logsv0.ApisAlertDefinitionAlertDefWebhooksSettings{*apisAlertDefinitionAlertDefWebhooksSettingsModel}

	apisAlertDefinitionLabelFilterTypeModel := new(logsv0.ApisAlertDefinitionLabelFilterType)
	apisAlertDefinitionLabelFilterTypeModel.Value = core.StringPtr("my-app")
	apisAlertDefinitionLabelFilterTypeModel.Operation = core.StringPtr("starts_with")

	apisAlertDefinitionLabelFiltersModel := new(logsv0.ApisAlertDefinitionLabelFilters)
	apisAlertDefinitionLabelFiltersModel.ApplicationName = []logsv0.ApisAlertDefinitionLabelFilterType{*apisAlertDefinitionLabelFilterTypeModel}
	apisAlertDefinitionLabelFiltersModel.SubsystemName = []logsv0.ApisAlertDefinitionLabelFilterType{*apisAlertDefinitionLabelFilterTypeModel}
	apisAlertDefinitionLabelFiltersModel.Severities = []string{"critical"}

	apisAlertDefinitionLogsSimpleFilterModel := new(logsv0.ApisAlertDefinitionLogsSimpleFilter)
	apisAlertDefinitionLogsSimpleFilterModel.LuceneQuery = core.StringPtr("text:"error"")
	apisAlertDefinitionLogsSimpleFilterModel.LabelFilters = apisAlertDefinitionLabelFiltersModel

	apisAlertDefinitionLogsFilterModel := new(logsv0.ApisAlertDefinitionLogsFilter)
	apisAlertDefinitionLogsFilterModel.SimpleFilter = apisAlertDefinitionLogsSimpleFilterModel

	apisAlertDefinitionLogsRatioTimeWindowModel := new(logsv0.ApisAlertDefinitionLogsRatioTimeWindow)
	apisAlertDefinitionLogsRatioTimeWindowModel.LogsRatioTimeWindowSpecificValue = core.StringPtr("hours_36")

	apisAlertDefinitionLogsRatioConditionModel := new(logsv0.ApisAlertDefinitionLogsRatioCondition)
	apisAlertDefinitionLogsRatioConditionModel.Threshold = core.Float64Ptr(float64(10.0))
	apisAlertDefinitionLogsRatioConditionModel.TimeWindow = apisAlertDefinitionLogsRatioTimeWindowModel

	apisAlertDefinitionAlertDefOverrideModel := new(logsv0.ApisAlertDefinitionAlertDefOverride)
	apisAlertDefinitionAlertDefOverrideModel.Priority = core.StringPtr("p1")

	apisAlertDefinitionLogsRatioRulesModel := new(logsv0.ApisAlertDefinitionLogsRatioRules)
	apisAlertDefinitionLogsRatioRulesModel.Condition = apisAlertDefinitionLogsRatioConditionModel
	apisAlertDefinitionLogsRatioRulesModel.Override = apisAlertDefinitionAlertDefOverrideModel

	apisAlertDefinitionUndetectedValuesManagementModel := new(logsv0.ApisAlertDefinitionUndetectedValuesManagement)
	apisAlertDefinitionUndetectedValuesManagementModel.TriggerUndetectedValues = core.BoolPtr(true)
	apisAlertDefinitionUndetectedValuesManagementModel.AutoRetireTimeframe = core.StringPtr("hours_24")

	apisAlertDefinitionLogsRatioThresholdTypeModel := new(logsv0.ApisAlertDefinitionLogsRatioThresholdType)
	apisAlertDefinitionLogsRatioThresholdTypeModel.Numerator = apisAlertDefinitionLogsFilterModel
	apisAlertDefinitionLogsRatioThresholdTypeModel.NumeratorAlias = core.StringPtr("numerator_alias")
	apisAlertDefinitionLogsRatioThresholdTypeModel.Denominator = apisAlertDefinitionLogsFilterModel
	apisAlertDefinitionLogsRatioThresholdTypeModel.DenominatorAlias = core.StringPtr("denominator_alias")
	apisAlertDefinitionLogsRatioThresholdTypeModel.Rules = []logsv0.ApisAlertDefinitionLogsRatioRules{*apisAlertDefinitionLogsRatioRulesModel}
	apisAlertDefinitionLogsRatioThresholdTypeModel.ConditionType = core.StringPtr("less_than")
	apisAlertDefinitionLogsRatioThresholdTypeModel.NotificationPayloadFilter = []string{"obj.field"}
	apisAlertDefinitionLogsRatioThresholdTypeModel.GroupByFor = core.StringPtr("denumerator_only")
	apisAlertDefinitionLogsRatioThresholdTypeModel.UndetectedValuesManagement = apisAlertDefinitionUndetectedValuesManagementModel
	apisAlertDefinitionLogsRatioThresholdTypeModel.IgnoreInfinity = core.BoolPtr(true)
	apisAlertDefinitionLogsRatioThresholdTypeModel.EvaluationDelayMs = core.Int64Ptr(int64(60000))

	model := new(logsv0.AlertDefinitionApisAlertDefinitionAlertDefPropertiesTypeDefinitionLogsRatioThreshold)
	model.ID = CreateMockUUID("3dc02998-0b50-4ea8-b68a-4779d716fa1f")
	model.CreatedTime = CreateMockDateTime("2021-01-01T00:00:00.000Z")
	model.UpdatedTime = CreateMockDateTime("2021-01-01T00:00:00.000Z")
	model.AlertVersionID = CreateMockUUID("3dc02998-0b50-4ea8-b68a-4779d716fa1f")
	model.Name = core.StringPtr("Unique count alert")
	model.Description = core.StringPtr("Example of unique count alert from terraform")
	model.Enabled = core.BoolPtr(true)
	model.Priority = core.StringPtr("p1")
	model.ActiveOn = apisAlertDefinitionActivityScheduleModel
	model.Type = core.StringPtr("flow")
	model.GroupByKeys = []string{"key1", "key2"}
	model.IncidentsSettings = apisAlertDefinitionAlertDefIncidentSettingsModel
	model.NotificationGroup = apisAlertDefinitionAlertDefNotificationGroupModel
	model.EntityLabels = map[string]string{"key1": "testString"}
	model.PhantomMode = core.BoolPtr(false)
	model.Deleted = core.BoolPtr(false)
	model.LogsRatioThreshold = apisAlertDefinitionLogsRatioThresholdTypeModel

	result, err := logs.DataSourceIbmLogsAlertDefinitionsAlertDefinitionApisAlertDefinitionAlertDefPropertiesTypeDefinitionLogsRatioThresholdToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmLogsAlertDefinitionsAlertDefinitionApisAlertDefinitionAlertDefPropertiesTypeDefinitionLogsTimeRelativeThresholdToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		apisAlertDefinitionTimeOfDayModel := make(map[string]interface{})
		apisAlertDefinitionTimeOfDayModel["hours"] = int(14)
		apisAlertDefinitionTimeOfDayModel["minutes"] = int(30)

		apisAlertDefinitionActivityScheduleModel := make(map[string]interface{})
		apisAlertDefinitionActivityScheduleModel["day_of_week"] = []string{"sunday"}
		apisAlertDefinitionActivityScheduleModel["start_time"] = []map[string]interface{}{apisAlertDefinitionTimeOfDayModel}
		apisAlertDefinitionActivityScheduleModel["end_time"] = []map[string]interface{}{apisAlertDefinitionTimeOfDayModel}

		apisAlertDefinitionAlertDefIncidentSettingsModel := make(map[string]interface{})
		apisAlertDefinitionAlertDefIncidentSettingsModel["notify_on"] = "triggered_and_resolved"
		apisAlertDefinitionAlertDefIncidentSettingsModel["minutes"] = int(30)

		apisAlertDefinitionIntegrationTypeModel := make(map[string]interface{})
		apisAlertDefinitionIntegrationTypeModel["integration_id"] = int(123)

		apisAlertDefinitionAlertDefWebhooksSettingsModel := make(map[string]interface{})
		apisAlertDefinitionAlertDefWebhooksSettingsModel["notify_on"] = "triggered_and_resolved"
		apisAlertDefinitionAlertDefWebhooksSettingsModel["integration"] = []map[string]interface{}{apisAlertDefinitionIntegrationTypeModel}
		apisAlertDefinitionAlertDefWebhooksSettingsModel["minutes"] = int(15)

		apisAlertDefinitionAlertDefNotificationGroupModel := make(map[string]interface{})
		apisAlertDefinitionAlertDefNotificationGroupModel["group_by_keys"] = []string{"key1", "key2"}
		apisAlertDefinitionAlertDefNotificationGroupModel["webhooks"] = []map[string]interface{}{apisAlertDefinitionAlertDefWebhooksSettingsModel}

		apisAlertDefinitionLabelFilterTypeModel := make(map[string]interface{})
		apisAlertDefinitionLabelFilterTypeModel["value"] = "my-app"
		apisAlertDefinitionLabelFilterTypeModel["operation"] = "starts_with"

		apisAlertDefinitionLabelFiltersModel := make(map[string]interface{})
		apisAlertDefinitionLabelFiltersModel["application_name"] = []map[string]interface{}{apisAlertDefinitionLabelFilterTypeModel}
		apisAlertDefinitionLabelFiltersModel["subsystem_name"] = []map[string]interface{}{apisAlertDefinitionLabelFilterTypeModel}
		apisAlertDefinitionLabelFiltersModel["severities"] = []string{"critical"}

		apisAlertDefinitionLogsSimpleFilterModel := make(map[string]interface{})
		apisAlertDefinitionLogsSimpleFilterModel["lucene_query"] = "text:"error""
		apisAlertDefinitionLogsSimpleFilterModel["label_filters"] = []map[string]interface{}{apisAlertDefinitionLabelFiltersModel}

		apisAlertDefinitionLogsFilterModel := make(map[string]interface{})
		apisAlertDefinitionLogsFilterModel["simple_filter"] = []map[string]interface{}{apisAlertDefinitionLogsSimpleFilterModel}

		apisAlertDefinitionLogsTimeRelativeConditionModel := make(map[string]interface{})
		apisAlertDefinitionLogsTimeRelativeConditionModel["threshold"] = float64(100.0)
		apisAlertDefinitionLogsTimeRelativeConditionModel["compared_to"] = "same_day_last_month"

		apisAlertDefinitionAlertDefOverrideModel := make(map[string]interface{})
		apisAlertDefinitionAlertDefOverrideModel["priority"] = "p1"

		apisAlertDefinitionLogsTimeRelativeRuleModel := make(map[string]interface{})
		apisAlertDefinitionLogsTimeRelativeRuleModel["condition"] = []map[string]interface{}{apisAlertDefinitionLogsTimeRelativeConditionModel}
		apisAlertDefinitionLogsTimeRelativeRuleModel["override"] = []map[string]interface{}{apisAlertDefinitionAlertDefOverrideModel}

		apisAlertDefinitionUndetectedValuesManagementModel := make(map[string]interface{})
		apisAlertDefinitionUndetectedValuesManagementModel["trigger_undetected_values"] = true
		apisAlertDefinitionUndetectedValuesManagementModel["auto_retire_timeframe"] = "hours_24"

		apisAlertDefinitionLogsTimeRelativeThresholdTypeModel := make(map[string]interface{})
		apisAlertDefinitionLogsTimeRelativeThresholdTypeModel["logs_filter"] = []map[string]interface{}{apisAlertDefinitionLogsFilterModel}
		apisAlertDefinitionLogsTimeRelativeThresholdTypeModel["rules"] = []map[string]interface{}{apisAlertDefinitionLogsTimeRelativeRuleModel}
		apisAlertDefinitionLogsTimeRelativeThresholdTypeModel["condition_type"] = "less_than"
		apisAlertDefinitionLogsTimeRelativeThresholdTypeModel["ignore_infinity"] = true
		apisAlertDefinitionLogsTimeRelativeThresholdTypeModel["notification_payload_filter"] = []string{"obj.field"}
		apisAlertDefinitionLogsTimeRelativeThresholdTypeModel["undetected_values_management"] = []map[string]interface{}{apisAlertDefinitionUndetectedValuesManagementModel}
		apisAlertDefinitionLogsTimeRelativeThresholdTypeModel["evaluation_delay_ms"] = int(60000)

		model := make(map[string]interface{})
		model["id"] = "3dc02998-0b50-4ea8-b68a-4779d716fa1f"
		model["created_time"] = "2021-01-01T00:00:00.000Z"
		model["updated_time"] = "2021-01-01T00:00:00.000Z"
		model["alert_version_id"] = "3dc02998-0b50-4ea8-b68a-4779d716fa1f"
		model["name"] = "Unique count alert"
		model["description"] = "Example of unique count alert from terraform"
		model["enabled"] = true
		model["priority"] = "p1"
		model["active_on"] = []map[string]interface{}{apisAlertDefinitionActivityScheduleModel}
		model["type"] = "flow"
		model["group_by_keys"] = []string{"key1", "key2"}
		model["incidents_settings"] = []map[string]interface{}{apisAlertDefinitionAlertDefIncidentSettingsModel}
		model["notification_group"] = []map[string]interface{}{apisAlertDefinitionAlertDefNotificationGroupModel}
		model["entity_labels"] = map[string]interface{}{"key1": "testString"}
		model["phantom_mode"] = false
		model["deleted"] = false
		model["logs_time_relative_threshold"] = []map[string]interface{}{apisAlertDefinitionLogsTimeRelativeThresholdTypeModel}

		assert.Equal(t, result, model)
	}

	apisAlertDefinitionTimeOfDayModel := new(logsv0.ApisAlertDefinitionTimeOfDay)
	apisAlertDefinitionTimeOfDayModel.Hours = core.Int64Ptr(int64(14))
	apisAlertDefinitionTimeOfDayModel.Minutes = core.Int64Ptr(int64(30))

	apisAlertDefinitionActivityScheduleModel := new(logsv0.ApisAlertDefinitionActivitySchedule)
	apisAlertDefinitionActivityScheduleModel.DayOfWeek = []string{"sunday"}
	apisAlertDefinitionActivityScheduleModel.StartTime = apisAlertDefinitionTimeOfDayModel
	apisAlertDefinitionActivityScheduleModel.EndTime = apisAlertDefinitionTimeOfDayModel

	apisAlertDefinitionAlertDefIncidentSettingsModel := new(logsv0.ApisAlertDefinitionAlertDefIncidentSettings)
	apisAlertDefinitionAlertDefIncidentSettingsModel.NotifyOn = core.StringPtr("triggered_and_resolved")
	apisAlertDefinitionAlertDefIncidentSettingsModel.Minutes = core.Int64Ptr(int64(30))

	apisAlertDefinitionIntegrationTypeModel := new(logsv0.ApisAlertDefinitionIntegrationTypeIntegrationTypeIntegrationID)
	apisAlertDefinitionIntegrationTypeModel.IntegrationID = core.Int64Ptr(int64(123))

	apisAlertDefinitionAlertDefWebhooksSettingsModel := new(logsv0.ApisAlertDefinitionAlertDefWebhooksSettings)
	apisAlertDefinitionAlertDefWebhooksSettingsModel.NotifyOn = core.StringPtr("triggered_and_resolved")
	apisAlertDefinitionAlertDefWebhooksSettingsModel.Integration = apisAlertDefinitionIntegrationTypeModel
	apisAlertDefinitionAlertDefWebhooksSettingsModel.Minutes = core.Int64Ptr(int64(15))

	apisAlertDefinitionAlertDefNotificationGroupModel := new(logsv0.ApisAlertDefinitionAlertDefNotificationGroup)
	apisAlertDefinitionAlertDefNotificationGroupModel.GroupByKeys = []string{"key1", "key2"}
	apisAlertDefinitionAlertDefNotificationGroupModel.Webhooks = []logsv0.ApisAlertDefinitionAlertDefWebhooksSettings{*apisAlertDefinitionAlertDefWebhooksSettingsModel}

	apisAlertDefinitionLabelFilterTypeModel := new(logsv0.ApisAlertDefinitionLabelFilterType)
	apisAlertDefinitionLabelFilterTypeModel.Value = core.StringPtr("my-app")
	apisAlertDefinitionLabelFilterTypeModel.Operation = core.StringPtr("starts_with")

	apisAlertDefinitionLabelFiltersModel := new(logsv0.ApisAlertDefinitionLabelFilters)
	apisAlertDefinitionLabelFiltersModel.ApplicationName = []logsv0.ApisAlertDefinitionLabelFilterType{*apisAlertDefinitionLabelFilterTypeModel}
	apisAlertDefinitionLabelFiltersModel.SubsystemName = []logsv0.ApisAlertDefinitionLabelFilterType{*apisAlertDefinitionLabelFilterTypeModel}
	apisAlertDefinitionLabelFiltersModel.Severities = []string{"critical"}

	apisAlertDefinitionLogsSimpleFilterModel := new(logsv0.ApisAlertDefinitionLogsSimpleFilter)
	apisAlertDefinitionLogsSimpleFilterModel.LuceneQuery = core.StringPtr("text:"error"")
	apisAlertDefinitionLogsSimpleFilterModel.LabelFilters = apisAlertDefinitionLabelFiltersModel

	apisAlertDefinitionLogsFilterModel := new(logsv0.ApisAlertDefinitionLogsFilter)
	apisAlertDefinitionLogsFilterModel.SimpleFilter = apisAlertDefinitionLogsSimpleFilterModel

	apisAlertDefinitionLogsTimeRelativeConditionModel := new(logsv0.ApisAlertDefinitionLogsTimeRelativeCondition)
	apisAlertDefinitionLogsTimeRelativeConditionModel.Threshold = core.Float64Ptr(float64(100.0))
	apisAlertDefinitionLogsTimeRelativeConditionModel.ComparedTo = core.StringPtr("same_day_last_month")

	apisAlertDefinitionAlertDefOverrideModel := new(logsv0.ApisAlertDefinitionAlertDefOverride)
	apisAlertDefinitionAlertDefOverrideModel.Priority = core.StringPtr("p1")

	apisAlertDefinitionLogsTimeRelativeRuleModel := new(logsv0.ApisAlertDefinitionLogsTimeRelativeRule)
	apisAlertDefinitionLogsTimeRelativeRuleModel.Condition = apisAlertDefinitionLogsTimeRelativeConditionModel
	apisAlertDefinitionLogsTimeRelativeRuleModel.Override = apisAlertDefinitionAlertDefOverrideModel

	apisAlertDefinitionUndetectedValuesManagementModel := new(logsv0.ApisAlertDefinitionUndetectedValuesManagement)
	apisAlertDefinitionUndetectedValuesManagementModel.TriggerUndetectedValues = core.BoolPtr(true)
	apisAlertDefinitionUndetectedValuesManagementModel.AutoRetireTimeframe = core.StringPtr("hours_24")

	apisAlertDefinitionLogsTimeRelativeThresholdTypeModel := new(logsv0.ApisAlertDefinitionLogsTimeRelativeThresholdType)
	apisAlertDefinitionLogsTimeRelativeThresholdTypeModel.LogsFilter = apisAlertDefinitionLogsFilterModel
	apisAlertDefinitionLogsTimeRelativeThresholdTypeModel.Rules = []logsv0.ApisAlertDefinitionLogsTimeRelativeRule{*apisAlertDefinitionLogsTimeRelativeRuleModel}
	apisAlertDefinitionLogsTimeRelativeThresholdTypeModel.ConditionType = core.StringPtr("less_than")
	apisAlertDefinitionLogsTimeRelativeThresholdTypeModel.IgnoreInfinity = core.BoolPtr(true)
	apisAlertDefinitionLogsTimeRelativeThresholdTypeModel.NotificationPayloadFilter = []string{"obj.field"}
	apisAlertDefinitionLogsTimeRelativeThresholdTypeModel.UndetectedValuesManagement = apisAlertDefinitionUndetectedValuesManagementModel
	apisAlertDefinitionLogsTimeRelativeThresholdTypeModel.EvaluationDelayMs = core.Int64Ptr(int64(60000))

	model := new(logsv0.AlertDefinitionApisAlertDefinitionAlertDefPropertiesTypeDefinitionLogsTimeRelativeThreshold)
	model.ID = CreateMockUUID("3dc02998-0b50-4ea8-b68a-4779d716fa1f")
	model.CreatedTime = CreateMockDateTime("2021-01-01T00:00:00.000Z")
	model.UpdatedTime = CreateMockDateTime("2021-01-01T00:00:00.000Z")
	model.AlertVersionID = CreateMockUUID("3dc02998-0b50-4ea8-b68a-4779d716fa1f")
	model.Name = core.StringPtr("Unique count alert")
	model.Description = core.StringPtr("Example of unique count alert from terraform")
	model.Enabled = core.BoolPtr(true)
	model.Priority = core.StringPtr("p1")
	model.ActiveOn = apisAlertDefinitionActivityScheduleModel
	model.Type = core.StringPtr("flow")
	model.GroupByKeys = []string{"key1", "key2"}
	model.IncidentsSettings = apisAlertDefinitionAlertDefIncidentSettingsModel
	model.NotificationGroup = apisAlertDefinitionAlertDefNotificationGroupModel
	model.EntityLabels = map[string]string{"key1": "testString"}
	model.PhantomMode = core.BoolPtr(false)
	model.Deleted = core.BoolPtr(false)
	model.LogsTimeRelativeThreshold = apisAlertDefinitionLogsTimeRelativeThresholdTypeModel

	result, err := logs.DataSourceIbmLogsAlertDefinitionsAlertDefinitionApisAlertDefinitionAlertDefPropertiesTypeDefinitionLogsTimeRelativeThresholdToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmLogsAlertDefinitionsAlertDefinitionApisAlertDefinitionAlertDefPropertiesTypeDefinitionMetricThresholdToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		apisAlertDefinitionTimeOfDayModel := make(map[string]interface{})
		apisAlertDefinitionTimeOfDayModel["hours"] = int(14)
		apisAlertDefinitionTimeOfDayModel["minutes"] = int(30)

		apisAlertDefinitionActivityScheduleModel := make(map[string]interface{})
		apisAlertDefinitionActivityScheduleModel["day_of_week"] = []string{"sunday"}
		apisAlertDefinitionActivityScheduleModel["start_time"] = []map[string]interface{}{apisAlertDefinitionTimeOfDayModel}
		apisAlertDefinitionActivityScheduleModel["end_time"] = []map[string]interface{}{apisAlertDefinitionTimeOfDayModel}

		apisAlertDefinitionAlertDefIncidentSettingsModel := make(map[string]interface{})
		apisAlertDefinitionAlertDefIncidentSettingsModel["notify_on"] = "triggered_and_resolved"
		apisAlertDefinitionAlertDefIncidentSettingsModel["minutes"] = int(30)

		apisAlertDefinitionIntegrationTypeModel := make(map[string]interface{})
		apisAlertDefinitionIntegrationTypeModel["integration_id"] = int(123)

		apisAlertDefinitionAlertDefWebhooksSettingsModel := make(map[string]interface{})
		apisAlertDefinitionAlertDefWebhooksSettingsModel["notify_on"] = "triggered_and_resolved"
		apisAlertDefinitionAlertDefWebhooksSettingsModel["integration"] = []map[string]interface{}{apisAlertDefinitionIntegrationTypeModel}
		apisAlertDefinitionAlertDefWebhooksSettingsModel["minutes"] = int(15)

		apisAlertDefinitionAlertDefNotificationGroupModel := make(map[string]interface{})
		apisAlertDefinitionAlertDefNotificationGroupModel["group_by_keys"] = []string{"key1", "key2"}
		apisAlertDefinitionAlertDefNotificationGroupModel["webhooks"] = []map[string]interface{}{apisAlertDefinitionAlertDefWebhooksSettingsModel}

		apisAlertDefinitionMetricFilterModel := make(map[string]interface{})
		apisAlertDefinitionMetricFilterModel["promql"] = "avg_over_time(metric_name[5m]) > 10"

		apisAlertDefinitionMetricTimeWindowModel := make(map[string]interface{})
		apisAlertDefinitionMetricTimeWindowModel["metric_time_window_specific_value"] = "hours_36"

		apisAlertDefinitionMetricThresholdConditionModel := make(map[string]interface{})
		apisAlertDefinitionMetricThresholdConditionModel["threshold"] = float64(100.0)
		apisAlertDefinitionMetricThresholdConditionModel["for_over_pct"] = int(80)
		apisAlertDefinitionMetricThresholdConditionModel["of_the_last"] = []map[string]interface{}{apisAlertDefinitionMetricTimeWindowModel}

		apisAlertDefinitionAlertDefOverrideModel := make(map[string]interface{})
		apisAlertDefinitionAlertDefOverrideModel["priority"] = "p1"

		apisAlertDefinitionMetricThresholdRuleModel := make(map[string]interface{})
		apisAlertDefinitionMetricThresholdRuleModel["condition"] = []map[string]interface{}{apisAlertDefinitionMetricThresholdConditionModel}
		apisAlertDefinitionMetricThresholdRuleModel["override"] = []map[string]interface{}{apisAlertDefinitionAlertDefOverrideModel}

		apisAlertDefinitionUndetectedValuesManagementModel := make(map[string]interface{})
		apisAlertDefinitionUndetectedValuesManagementModel["trigger_undetected_values"] = true
		apisAlertDefinitionUndetectedValuesManagementModel["auto_retire_timeframe"] = "hours_24"

		apisAlertDefinitionMetricMissingValuesModel := make(map[string]interface{})
		apisAlertDefinitionMetricMissingValuesModel["replace_with_zero"] = true

		apisAlertDefinitionMetricThresholdTypeModel := make(map[string]interface{})
		apisAlertDefinitionMetricThresholdTypeModel["metric_filter"] = []map[string]interface{}{apisAlertDefinitionMetricFilterModel}
		apisAlertDefinitionMetricThresholdTypeModel["rules"] = []map[string]interface{}{apisAlertDefinitionMetricThresholdRuleModel}
		apisAlertDefinitionMetricThresholdTypeModel["condition_type"] = "less_than_or_equals"
		apisAlertDefinitionMetricThresholdTypeModel["undetected_values_management"] = []map[string]interface{}{apisAlertDefinitionUndetectedValuesManagementModel}
		apisAlertDefinitionMetricThresholdTypeModel["missing_values"] = []map[string]interface{}{apisAlertDefinitionMetricMissingValuesModel}
		apisAlertDefinitionMetricThresholdTypeModel["evaluation_delay_ms"] = int(60000)

		model := make(map[string]interface{})
		model["id"] = "3dc02998-0b50-4ea8-b68a-4779d716fa1f"
		model["created_time"] = "2021-01-01T00:00:00.000Z"
		model["updated_time"] = "2021-01-01T00:00:00.000Z"
		model["alert_version_id"] = "3dc02998-0b50-4ea8-b68a-4779d716fa1f"
		model["name"] = "Unique count alert"
		model["description"] = "Example of unique count alert from terraform"
		model["enabled"] = true
		model["priority"] = "p1"
		model["active_on"] = []map[string]interface{}{apisAlertDefinitionActivityScheduleModel}
		model["type"] = "flow"
		model["group_by_keys"] = []string{"key1", "key2"}
		model["incidents_settings"] = []map[string]interface{}{apisAlertDefinitionAlertDefIncidentSettingsModel}
		model["notification_group"] = []map[string]interface{}{apisAlertDefinitionAlertDefNotificationGroupModel}
		model["entity_labels"] = map[string]interface{}{"key1": "testString"}
		model["phantom_mode"] = false
		model["deleted"] = false
		model["metric_threshold"] = []map[string]interface{}{apisAlertDefinitionMetricThresholdTypeModel}

		assert.Equal(t, result, model)
	}

	apisAlertDefinitionTimeOfDayModel := new(logsv0.ApisAlertDefinitionTimeOfDay)
	apisAlertDefinitionTimeOfDayModel.Hours = core.Int64Ptr(int64(14))
	apisAlertDefinitionTimeOfDayModel.Minutes = core.Int64Ptr(int64(30))

	apisAlertDefinitionActivityScheduleModel := new(logsv0.ApisAlertDefinitionActivitySchedule)
	apisAlertDefinitionActivityScheduleModel.DayOfWeek = []string{"sunday"}
	apisAlertDefinitionActivityScheduleModel.StartTime = apisAlertDefinitionTimeOfDayModel
	apisAlertDefinitionActivityScheduleModel.EndTime = apisAlertDefinitionTimeOfDayModel

	apisAlertDefinitionAlertDefIncidentSettingsModel := new(logsv0.ApisAlertDefinitionAlertDefIncidentSettings)
	apisAlertDefinitionAlertDefIncidentSettingsModel.NotifyOn = core.StringPtr("triggered_and_resolved")
	apisAlertDefinitionAlertDefIncidentSettingsModel.Minutes = core.Int64Ptr(int64(30))

	apisAlertDefinitionIntegrationTypeModel := new(logsv0.ApisAlertDefinitionIntegrationTypeIntegrationTypeIntegrationID)
	apisAlertDefinitionIntegrationTypeModel.IntegrationID = core.Int64Ptr(int64(123))

	apisAlertDefinitionAlertDefWebhooksSettingsModel := new(logsv0.ApisAlertDefinitionAlertDefWebhooksSettings)
	apisAlertDefinitionAlertDefWebhooksSettingsModel.NotifyOn = core.StringPtr("triggered_and_resolved")
	apisAlertDefinitionAlertDefWebhooksSettingsModel.Integration = apisAlertDefinitionIntegrationTypeModel
	apisAlertDefinitionAlertDefWebhooksSettingsModel.Minutes = core.Int64Ptr(int64(15))

	apisAlertDefinitionAlertDefNotificationGroupModel := new(logsv0.ApisAlertDefinitionAlertDefNotificationGroup)
	apisAlertDefinitionAlertDefNotificationGroupModel.GroupByKeys = []string{"key1", "key2"}
	apisAlertDefinitionAlertDefNotificationGroupModel.Webhooks = []logsv0.ApisAlertDefinitionAlertDefWebhooksSettings{*apisAlertDefinitionAlertDefWebhooksSettingsModel}

	apisAlertDefinitionMetricFilterModel := new(logsv0.ApisAlertDefinitionMetricFilter)
	apisAlertDefinitionMetricFilterModel.Promql = core.StringPtr("avg_over_time(metric_name[5m]) > 10")

	apisAlertDefinitionMetricTimeWindowModel := new(logsv0.ApisAlertDefinitionMetricTimeWindowTypeMetricTimeWindowSpecificValue)
	apisAlertDefinitionMetricTimeWindowModel.MetricTimeWindowSpecificValue = core.StringPtr("hours_36")

	apisAlertDefinitionMetricThresholdConditionModel := new(logsv0.ApisAlertDefinitionMetricThresholdCondition)
	apisAlertDefinitionMetricThresholdConditionModel.Threshold = core.Float64Ptr(float64(100.0))
	apisAlertDefinitionMetricThresholdConditionModel.ForOverPct = core.Int64Ptr(int64(80))
	apisAlertDefinitionMetricThresholdConditionModel.OfTheLast = apisAlertDefinitionMetricTimeWindowModel

	apisAlertDefinitionAlertDefOverrideModel := new(logsv0.ApisAlertDefinitionAlertDefOverride)
	apisAlertDefinitionAlertDefOverrideModel.Priority = core.StringPtr("p1")

	apisAlertDefinitionMetricThresholdRuleModel := new(logsv0.ApisAlertDefinitionMetricThresholdRule)
	apisAlertDefinitionMetricThresholdRuleModel.Condition = apisAlertDefinitionMetricThresholdConditionModel
	apisAlertDefinitionMetricThresholdRuleModel.Override = apisAlertDefinitionAlertDefOverrideModel

	apisAlertDefinitionUndetectedValuesManagementModel := new(logsv0.ApisAlertDefinitionUndetectedValuesManagement)
	apisAlertDefinitionUndetectedValuesManagementModel.TriggerUndetectedValues = core.BoolPtr(true)
	apisAlertDefinitionUndetectedValuesManagementModel.AutoRetireTimeframe = core.StringPtr("hours_24")

	apisAlertDefinitionMetricMissingValuesModel := new(logsv0.ApisAlertDefinitionMetricMissingValuesMissingValuesReplaceWithZero)
	apisAlertDefinitionMetricMissingValuesModel.ReplaceWithZero = core.BoolPtr(true)

	apisAlertDefinitionMetricThresholdTypeModel := new(logsv0.ApisAlertDefinitionMetricThresholdType)
	apisAlertDefinitionMetricThresholdTypeModel.MetricFilter = apisAlertDefinitionMetricFilterModel
	apisAlertDefinitionMetricThresholdTypeModel.Rules = []logsv0.ApisAlertDefinitionMetricThresholdRule{*apisAlertDefinitionMetricThresholdRuleModel}
	apisAlertDefinitionMetricThresholdTypeModel.ConditionType = core.StringPtr("less_than_or_equals")
	apisAlertDefinitionMetricThresholdTypeModel.UndetectedValuesManagement = apisAlertDefinitionUndetectedValuesManagementModel
	apisAlertDefinitionMetricThresholdTypeModel.MissingValues = apisAlertDefinitionMetricMissingValuesModel
	apisAlertDefinitionMetricThresholdTypeModel.EvaluationDelayMs = core.Int64Ptr(int64(60000))

	model := new(logsv0.AlertDefinitionApisAlertDefinitionAlertDefPropertiesTypeDefinitionMetricThreshold)
	model.ID = CreateMockUUID("3dc02998-0b50-4ea8-b68a-4779d716fa1f")
	model.CreatedTime = CreateMockDateTime("2021-01-01T00:00:00.000Z")
	model.UpdatedTime = CreateMockDateTime("2021-01-01T00:00:00.000Z")
	model.AlertVersionID = CreateMockUUID("3dc02998-0b50-4ea8-b68a-4779d716fa1f")
	model.Name = core.StringPtr("Unique count alert")
	model.Description = core.StringPtr("Example of unique count alert from terraform")
	model.Enabled = core.BoolPtr(true)
	model.Priority = core.StringPtr("p1")
	model.ActiveOn = apisAlertDefinitionActivityScheduleModel
	model.Type = core.StringPtr("flow")
	model.GroupByKeys = []string{"key1", "key2"}
	model.IncidentsSettings = apisAlertDefinitionAlertDefIncidentSettingsModel
	model.NotificationGroup = apisAlertDefinitionAlertDefNotificationGroupModel
	model.EntityLabels = map[string]string{"key1": "testString"}
	model.PhantomMode = core.BoolPtr(false)
	model.Deleted = core.BoolPtr(false)
	model.MetricThreshold = apisAlertDefinitionMetricThresholdTypeModel

	result, err := logs.DataSourceIbmLogsAlertDefinitionsAlertDefinitionApisAlertDefinitionAlertDefPropertiesTypeDefinitionMetricThresholdToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmLogsAlertDefinitionsAlertDefinitionApisAlertDefinitionAlertDefPropertiesTypeDefinitionFlowToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		apisAlertDefinitionTimeOfDayModel := make(map[string]interface{})
		apisAlertDefinitionTimeOfDayModel["hours"] = int(14)
		apisAlertDefinitionTimeOfDayModel["minutes"] = int(30)

		apisAlertDefinitionActivityScheduleModel := make(map[string]interface{})
		apisAlertDefinitionActivityScheduleModel["day_of_week"] = []string{"sunday"}
		apisAlertDefinitionActivityScheduleModel["start_time"] = []map[string]interface{}{apisAlertDefinitionTimeOfDayModel}
		apisAlertDefinitionActivityScheduleModel["end_time"] = []map[string]interface{}{apisAlertDefinitionTimeOfDayModel}

		apisAlertDefinitionAlertDefIncidentSettingsModel := make(map[string]interface{})
		apisAlertDefinitionAlertDefIncidentSettingsModel["notify_on"] = "triggered_and_resolved"
		apisAlertDefinitionAlertDefIncidentSettingsModel["minutes"] = int(30)

		apisAlertDefinitionIntegrationTypeModel := make(map[string]interface{})
		apisAlertDefinitionIntegrationTypeModel["integration_id"] = int(123)

		apisAlertDefinitionAlertDefWebhooksSettingsModel := make(map[string]interface{})
		apisAlertDefinitionAlertDefWebhooksSettingsModel["notify_on"] = "triggered_and_resolved"
		apisAlertDefinitionAlertDefWebhooksSettingsModel["integration"] = []map[string]interface{}{apisAlertDefinitionIntegrationTypeModel}
		apisAlertDefinitionAlertDefWebhooksSettingsModel["minutes"] = int(15)

		apisAlertDefinitionAlertDefNotificationGroupModel := make(map[string]interface{})
		apisAlertDefinitionAlertDefNotificationGroupModel["group_by_keys"] = []string{"key1", "key2"}
		apisAlertDefinitionAlertDefNotificationGroupModel["webhooks"] = []map[string]interface{}{apisAlertDefinitionAlertDefWebhooksSettingsModel}

		apisAlertDefinitionFlowStagesGroupsAlertDefsModel := make(map[string]interface{})
		apisAlertDefinitionFlowStagesGroupsAlertDefsModel["id"] = "3dc02998-0b50-4ea8-b68a-4779d716fa1f"
		apisAlertDefinitionFlowStagesGroupsAlertDefsModel["not"] = true

		apisAlertDefinitionFlowStagesGroupModel := make(map[string]interface{})
		apisAlertDefinitionFlowStagesGroupModel["alert_defs"] = []map[string]interface{}{apisAlertDefinitionFlowStagesGroupsAlertDefsModel}
		apisAlertDefinitionFlowStagesGroupModel["next_op"] = "or"
		apisAlertDefinitionFlowStagesGroupModel["alerts_op"] = "or"

		apisAlertDefinitionFlowStagesGroupsModel := make(map[string]interface{})
		apisAlertDefinitionFlowStagesGroupsModel["groups"] = []map[string]interface{}{apisAlertDefinitionFlowStagesGroupModel}

		apisAlertDefinitionFlowStagesModel := make(map[string]interface{})
		apisAlertDefinitionFlowStagesModel["timeframe_ms"] = "60000"
		apisAlertDefinitionFlowStagesModel["timeframe_type"] = "up_to"
		apisAlertDefinitionFlowStagesModel["flow_stages_groups"] = []map[string]interface{}{apisAlertDefinitionFlowStagesGroupsModel}

		apisAlertDefinitionFlowTypeModel := make(map[string]interface{})
		apisAlertDefinitionFlowTypeModel["stages"] = []map[string]interface{}{apisAlertDefinitionFlowStagesModel}
		apisAlertDefinitionFlowTypeModel["enforce_suppression"] = true

		model := make(map[string]interface{})
		model["id"] = "3dc02998-0b50-4ea8-b68a-4779d716fa1f"
		model["created_time"] = "2021-01-01T00:00:00.000Z"
		model["updated_time"] = "2021-01-01T00:00:00.000Z"
		model["alert_version_id"] = "3dc02998-0b50-4ea8-b68a-4779d716fa1f"
		model["name"] = "Unique count alert"
		model["description"] = "Example of unique count alert from terraform"
		model["enabled"] = true
		model["priority"] = "p1"
		model["active_on"] = []map[string]interface{}{apisAlertDefinitionActivityScheduleModel}
		model["type"] = "flow"
		model["group_by_keys"] = []string{"key1", "key2"}
		model["incidents_settings"] = []map[string]interface{}{apisAlertDefinitionAlertDefIncidentSettingsModel}
		model["notification_group"] = []map[string]interface{}{apisAlertDefinitionAlertDefNotificationGroupModel}
		model["entity_labels"] = map[string]interface{}{"key1": "testString"}
		model["phantom_mode"] = false
		model["deleted"] = false
		model["flow"] = []map[string]interface{}{apisAlertDefinitionFlowTypeModel}

		assert.Equal(t, result, model)
	}

	apisAlertDefinitionTimeOfDayModel := new(logsv0.ApisAlertDefinitionTimeOfDay)
	apisAlertDefinitionTimeOfDayModel.Hours = core.Int64Ptr(int64(14))
	apisAlertDefinitionTimeOfDayModel.Minutes = core.Int64Ptr(int64(30))

	apisAlertDefinitionActivityScheduleModel := new(logsv0.ApisAlertDefinitionActivitySchedule)
	apisAlertDefinitionActivityScheduleModel.DayOfWeek = []string{"sunday"}
	apisAlertDefinitionActivityScheduleModel.StartTime = apisAlertDefinitionTimeOfDayModel
	apisAlertDefinitionActivityScheduleModel.EndTime = apisAlertDefinitionTimeOfDayModel

	apisAlertDefinitionAlertDefIncidentSettingsModel := new(logsv0.ApisAlertDefinitionAlertDefIncidentSettings)
	apisAlertDefinitionAlertDefIncidentSettingsModel.NotifyOn = core.StringPtr("triggered_and_resolved")
	apisAlertDefinitionAlertDefIncidentSettingsModel.Minutes = core.Int64Ptr(int64(30))

	apisAlertDefinitionIntegrationTypeModel := new(logsv0.ApisAlertDefinitionIntegrationTypeIntegrationTypeIntegrationID)
	apisAlertDefinitionIntegrationTypeModel.IntegrationID = core.Int64Ptr(int64(123))

	apisAlertDefinitionAlertDefWebhooksSettingsModel := new(logsv0.ApisAlertDefinitionAlertDefWebhooksSettings)
	apisAlertDefinitionAlertDefWebhooksSettingsModel.NotifyOn = core.StringPtr("triggered_and_resolved")
	apisAlertDefinitionAlertDefWebhooksSettingsModel.Integration = apisAlertDefinitionIntegrationTypeModel
	apisAlertDefinitionAlertDefWebhooksSettingsModel.Minutes = core.Int64Ptr(int64(15))

	apisAlertDefinitionAlertDefNotificationGroupModel := new(logsv0.ApisAlertDefinitionAlertDefNotificationGroup)
	apisAlertDefinitionAlertDefNotificationGroupModel.GroupByKeys = []string{"key1", "key2"}
	apisAlertDefinitionAlertDefNotificationGroupModel.Webhooks = []logsv0.ApisAlertDefinitionAlertDefWebhooksSettings{*apisAlertDefinitionAlertDefWebhooksSettingsModel}

	apisAlertDefinitionFlowStagesGroupsAlertDefsModel := new(logsv0.ApisAlertDefinitionFlowStagesGroupsAlertDefs)
	apisAlertDefinitionFlowStagesGroupsAlertDefsModel.ID = CreateMockUUID("3dc02998-0b50-4ea8-b68a-4779d716fa1f")
	apisAlertDefinitionFlowStagesGroupsAlertDefsModel.Not = core.BoolPtr(true)

	apisAlertDefinitionFlowStagesGroupModel := new(logsv0.ApisAlertDefinitionFlowStagesGroup)
	apisAlertDefinitionFlowStagesGroupModel.AlertDefs = []logsv0.ApisAlertDefinitionFlowStagesGroupsAlertDefs{*apisAlertDefinitionFlowStagesGroupsAlertDefsModel}
	apisAlertDefinitionFlowStagesGroupModel.NextOp = core.StringPtr("or")
	apisAlertDefinitionFlowStagesGroupModel.AlertsOp = core.StringPtr("or")

	apisAlertDefinitionFlowStagesGroupsModel := new(logsv0.ApisAlertDefinitionFlowStagesGroups)
	apisAlertDefinitionFlowStagesGroupsModel.Groups = []logsv0.ApisAlertDefinitionFlowStagesGroup{*apisAlertDefinitionFlowStagesGroupModel}

	apisAlertDefinitionFlowStagesModel := new(logsv0.ApisAlertDefinitionFlowStages)
	apisAlertDefinitionFlowStagesModel.TimeframeMs = core.StringPtr("60000")
	apisAlertDefinitionFlowStagesModel.TimeframeType = core.StringPtr("up_to")
	apisAlertDefinitionFlowStagesModel.FlowStagesGroups = apisAlertDefinitionFlowStagesGroupsModel

	apisAlertDefinitionFlowTypeModel := new(logsv0.ApisAlertDefinitionFlowType)
	apisAlertDefinitionFlowTypeModel.Stages = []logsv0.ApisAlertDefinitionFlowStages{*apisAlertDefinitionFlowStagesModel}
	apisAlertDefinitionFlowTypeModel.EnforceSuppression = core.BoolPtr(true)

	model := new(logsv0.AlertDefinitionApisAlertDefinitionAlertDefPropertiesTypeDefinitionFlow)
	model.ID = CreateMockUUID("3dc02998-0b50-4ea8-b68a-4779d716fa1f")
	model.CreatedTime = CreateMockDateTime("2021-01-01T00:00:00.000Z")
	model.UpdatedTime = CreateMockDateTime("2021-01-01T00:00:00.000Z")
	model.AlertVersionID = CreateMockUUID("3dc02998-0b50-4ea8-b68a-4779d716fa1f")
	model.Name = core.StringPtr("Unique count alert")
	model.Description = core.StringPtr("Example of unique count alert from terraform")
	model.Enabled = core.BoolPtr(true)
	model.Priority = core.StringPtr("p1")
	model.ActiveOn = apisAlertDefinitionActivityScheduleModel
	model.Type = core.StringPtr("flow")
	model.GroupByKeys = []string{"key1", "key2"}
	model.IncidentsSettings = apisAlertDefinitionAlertDefIncidentSettingsModel
	model.NotificationGroup = apisAlertDefinitionAlertDefNotificationGroupModel
	model.EntityLabels = map[string]string{"key1": "testString"}
	model.PhantomMode = core.BoolPtr(false)
	model.Deleted = core.BoolPtr(false)
	model.Flow = apisAlertDefinitionFlowTypeModel

	result, err := logs.DataSourceIbmLogsAlertDefinitionsAlertDefinitionApisAlertDefinitionAlertDefPropertiesTypeDefinitionFlowToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmLogsAlertDefinitionsAlertDefinitionApisAlertDefinitionAlertDefPropertiesTypeDefinitionLogsAnomalyToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		apisAlertDefinitionTimeOfDayModel := make(map[string]interface{})
		apisAlertDefinitionTimeOfDayModel["hours"] = int(14)
		apisAlertDefinitionTimeOfDayModel["minutes"] = int(30)

		apisAlertDefinitionActivityScheduleModel := make(map[string]interface{})
		apisAlertDefinitionActivityScheduleModel["day_of_week"] = []string{"sunday"}
		apisAlertDefinitionActivityScheduleModel["start_time"] = []map[string]interface{}{apisAlertDefinitionTimeOfDayModel}
		apisAlertDefinitionActivityScheduleModel["end_time"] = []map[string]interface{}{apisAlertDefinitionTimeOfDayModel}

		apisAlertDefinitionAlertDefIncidentSettingsModel := make(map[string]interface{})
		apisAlertDefinitionAlertDefIncidentSettingsModel["notify_on"] = "triggered_and_resolved"
		apisAlertDefinitionAlertDefIncidentSettingsModel["minutes"] = int(30)

		apisAlertDefinitionIntegrationTypeModel := make(map[string]interface{})
		apisAlertDefinitionIntegrationTypeModel["integration_id"] = int(123)

		apisAlertDefinitionAlertDefWebhooksSettingsModel := make(map[string]interface{})
		apisAlertDefinitionAlertDefWebhooksSettingsModel["notify_on"] = "triggered_and_resolved"
		apisAlertDefinitionAlertDefWebhooksSettingsModel["integration"] = []map[string]interface{}{apisAlertDefinitionIntegrationTypeModel}
		apisAlertDefinitionAlertDefWebhooksSettingsModel["minutes"] = int(15)

		apisAlertDefinitionAlertDefNotificationGroupModel := make(map[string]interface{})
		apisAlertDefinitionAlertDefNotificationGroupModel["group_by_keys"] = []string{"key1", "key2"}
		apisAlertDefinitionAlertDefNotificationGroupModel["webhooks"] = []map[string]interface{}{apisAlertDefinitionAlertDefWebhooksSettingsModel}

		apisAlertDefinitionLabelFilterTypeModel := make(map[string]interface{})
		apisAlertDefinitionLabelFilterTypeModel["value"] = "my-app"
		apisAlertDefinitionLabelFilterTypeModel["operation"] = "starts_with"

		apisAlertDefinitionLabelFiltersModel := make(map[string]interface{})
		apisAlertDefinitionLabelFiltersModel["application_name"] = []map[string]interface{}{apisAlertDefinitionLabelFilterTypeModel}
		apisAlertDefinitionLabelFiltersModel["subsystem_name"] = []map[string]interface{}{apisAlertDefinitionLabelFilterTypeModel}
		apisAlertDefinitionLabelFiltersModel["severities"] = []string{"critical"}

		apisAlertDefinitionLogsSimpleFilterModel := make(map[string]interface{})
		apisAlertDefinitionLogsSimpleFilterModel["lucene_query"] = "text:"error""
		apisAlertDefinitionLogsSimpleFilterModel["label_filters"] = []map[string]interface{}{apisAlertDefinitionLabelFiltersModel}

		apisAlertDefinitionLogsFilterModel := make(map[string]interface{})
		apisAlertDefinitionLogsFilterModel["simple_filter"] = []map[string]interface{}{apisAlertDefinitionLogsSimpleFilterModel}

		apisAlertDefinitionLogsTimeWindowModel := make(map[string]interface{})
		apisAlertDefinitionLogsTimeWindowModel["logs_time_window_specific_value"] = "hours_36"

		apisAlertDefinitionLogsAnomalyConditionModel := make(map[string]interface{})
		apisAlertDefinitionLogsAnomalyConditionModel["minimum_threshold"] = float64(10.0)
		apisAlertDefinitionLogsAnomalyConditionModel["time_window"] = []map[string]interface{}{apisAlertDefinitionLogsTimeWindowModel}

		apisAlertDefinitionLogsAnomalyRuleModel := make(map[string]interface{})
		apisAlertDefinitionLogsAnomalyRuleModel["condition"] = []map[string]interface{}{apisAlertDefinitionLogsAnomalyConditionModel}

		apisAlertDefinitionAnomalyAlertSettingsModel := make(map[string]interface{})
		apisAlertDefinitionAnomalyAlertSettingsModel["percentage_of_deviation"] = float64(10.0)

		apisAlertDefinitionLogsAnomalyTypeModel := make(map[string]interface{})
		apisAlertDefinitionLogsAnomalyTypeModel["logs_filter"] = []map[string]interface{}{apisAlertDefinitionLogsFilterModel}
		apisAlertDefinitionLogsAnomalyTypeModel["rules"] = []map[string]interface{}{apisAlertDefinitionLogsAnomalyRuleModel}
		apisAlertDefinitionLogsAnomalyTypeModel["condition_type"] = "more_than_usual_or_unspecified"
		apisAlertDefinitionLogsAnomalyTypeModel["notification_payload_filter"] = []string{"obj.field"}
		apisAlertDefinitionLogsAnomalyTypeModel["evaluation_delay_ms"] = int(60000)
		apisAlertDefinitionLogsAnomalyTypeModel["anomaly_alert_settings"] = []map[string]interface{}{apisAlertDefinitionAnomalyAlertSettingsModel}

		model := make(map[string]interface{})
		model["id"] = "3dc02998-0b50-4ea8-b68a-4779d716fa1f"
		model["created_time"] = "2021-01-01T00:00:00.000Z"
		model["updated_time"] = "2021-01-01T00:00:00.000Z"
		model["alert_version_id"] = "3dc02998-0b50-4ea8-b68a-4779d716fa1f"
		model["name"] = "Unique count alert"
		model["description"] = "Example of unique count alert from terraform"
		model["enabled"] = true
		model["priority"] = "p1"
		model["active_on"] = []map[string]interface{}{apisAlertDefinitionActivityScheduleModel}
		model["type"] = "flow"
		model["group_by_keys"] = []string{"key1", "key2"}
		model["incidents_settings"] = []map[string]interface{}{apisAlertDefinitionAlertDefIncidentSettingsModel}
		model["notification_group"] = []map[string]interface{}{apisAlertDefinitionAlertDefNotificationGroupModel}
		model["entity_labels"] = map[string]interface{}{"key1": "testString"}
		model["phantom_mode"] = false
		model["deleted"] = false
		model["logs_anomaly"] = []map[string]interface{}{apisAlertDefinitionLogsAnomalyTypeModel}

		assert.Equal(t, result, model)
	}

	apisAlertDefinitionTimeOfDayModel := new(logsv0.ApisAlertDefinitionTimeOfDay)
	apisAlertDefinitionTimeOfDayModel.Hours = core.Int64Ptr(int64(14))
	apisAlertDefinitionTimeOfDayModel.Minutes = core.Int64Ptr(int64(30))

	apisAlertDefinitionActivityScheduleModel := new(logsv0.ApisAlertDefinitionActivitySchedule)
	apisAlertDefinitionActivityScheduleModel.DayOfWeek = []string{"sunday"}
	apisAlertDefinitionActivityScheduleModel.StartTime = apisAlertDefinitionTimeOfDayModel
	apisAlertDefinitionActivityScheduleModel.EndTime = apisAlertDefinitionTimeOfDayModel

	apisAlertDefinitionAlertDefIncidentSettingsModel := new(logsv0.ApisAlertDefinitionAlertDefIncidentSettings)
	apisAlertDefinitionAlertDefIncidentSettingsModel.NotifyOn = core.StringPtr("triggered_and_resolved")
	apisAlertDefinitionAlertDefIncidentSettingsModel.Minutes = core.Int64Ptr(int64(30))

	apisAlertDefinitionIntegrationTypeModel := new(logsv0.ApisAlertDefinitionIntegrationTypeIntegrationTypeIntegrationID)
	apisAlertDefinitionIntegrationTypeModel.IntegrationID = core.Int64Ptr(int64(123))

	apisAlertDefinitionAlertDefWebhooksSettingsModel := new(logsv0.ApisAlertDefinitionAlertDefWebhooksSettings)
	apisAlertDefinitionAlertDefWebhooksSettingsModel.NotifyOn = core.StringPtr("triggered_and_resolved")
	apisAlertDefinitionAlertDefWebhooksSettingsModel.Integration = apisAlertDefinitionIntegrationTypeModel
	apisAlertDefinitionAlertDefWebhooksSettingsModel.Minutes = core.Int64Ptr(int64(15))

	apisAlertDefinitionAlertDefNotificationGroupModel := new(logsv0.ApisAlertDefinitionAlertDefNotificationGroup)
	apisAlertDefinitionAlertDefNotificationGroupModel.GroupByKeys = []string{"key1", "key2"}
	apisAlertDefinitionAlertDefNotificationGroupModel.Webhooks = []logsv0.ApisAlertDefinitionAlertDefWebhooksSettings{*apisAlertDefinitionAlertDefWebhooksSettingsModel}

	apisAlertDefinitionLabelFilterTypeModel := new(logsv0.ApisAlertDefinitionLabelFilterType)
	apisAlertDefinitionLabelFilterTypeModel.Value = core.StringPtr("my-app")
	apisAlertDefinitionLabelFilterTypeModel.Operation = core.StringPtr("starts_with")

	apisAlertDefinitionLabelFiltersModel := new(logsv0.ApisAlertDefinitionLabelFilters)
	apisAlertDefinitionLabelFiltersModel.ApplicationName = []logsv0.ApisAlertDefinitionLabelFilterType{*apisAlertDefinitionLabelFilterTypeModel}
	apisAlertDefinitionLabelFiltersModel.SubsystemName = []logsv0.ApisAlertDefinitionLabelFilterType{*apisAlertDefinitionLabelFilterTypeModel}
	apisAlertDefinitionLabelFiltersModel.Severities = []string{"critical"}

	apisAlertDefinitionLogsSimpleFilterModel := new(logsv0.ApisAlertDefinitionLogsSimpleFilter)
	apisAlertDefinitionLogsSimpleFilterModel.LuceneQuery = core.StringPtr("text:"error"")
	apisAlertDefinitionLogsSimpleFilterModel.LabelFilters = apisAlertDefinitionLabelFiltersModel

	apisAlertDefinitionLogsFilterModel := new(logsv0.ApisAlertDefinitionLogsFilter)
	apisAlertDefinitionLogsFilterModel.SimpleFilter = apisAlertDefinitionLogsSimpleFilterModel

	apisAlertDefinitionLogsTimeWindowModel := new(logsv0.ApisAlertDefinitionLogsTimeWindow)
	apisAlertDefinitionLogsTimeWindowModel.LogsTimeWindowSpecificValue = core.StringPtr("hours_36")

	apisAlertDefinitionLogsAnomalyConditionModel := new(logsv0.ApisAlertDefinitionLogsAnomalyCondition)
	apisAlertDefinitionLogsAnomalyConditionModel.MinimumThreshold = core.Float64Ptr(float64(10.0))
	apisAlertDefinitionLogsAnomalyConditionModel.TimeWindow = apisAlertDefinitionLogsTimeWindowModel

	apisAlertDefinitionLogsAnomalyRuleModel := new(logsv0.ApisAlertDefinitionLogsAnomalyRule)
	apisAlertDefinitionLogsAnomalyRuleModel.Condition = apisAlertDefinitionLogsAnomalyConditionModel

	apisAlertDefinitionAnomalyAlertSettingsModel := new(logsv0.ApisAlertDefinitionAnomalyAlertSettings)
	apisAlertDefinitionAnomalyAlertSettingsModel.PercentageOfDeviation = core.Float32Ptr(float32(10.0))

	apisAlertDefinitionLogsAnomalyTypeModel := new(logsv0.ApisAlertDefinitionLogsAnomalyType)
	apisAlertDefinitionLogsAnomalyTypeModel.LogsFilter = apisAlertDefinitionLogsFilterModel
	apisAlertDefinitionLogsAnomalyTypeModel.Rules = []logsv0.ApisAlertDefinitionLogsAnomalyRule{*apisAlertDefinitionLogsAnomalyRuleModel}
	apisAlertDefinitionLogsAnomalyTypeModel.ConditionType = core.StringPtr("more_than_usual_or_unspecified")
	apisAlertDefinitionLogsAnomalyTypeModel.NotificationPayloadFilter = []string{"obj.field"}
	apisAlertDefinitionLogsAnomalyTypeModel.EvaluationDelayMs = core.Int64Ptr(int64(60000))
	apisAlertDefinitionLogsAnomalyTypeModel.AnomalyAlertSettings = apisAlertDefinitionAnomalyAlertSettingsModel

	model := new(logsv0.AlertDefinitionApisAlertDefinitionAlertDefPropertiesTypeDefinitionLogsAnomaly)
	model.ID = CreateMockUUID("3dc02998-0b50-4ea8-b68a-4779d716fa1f")
	model.CreatedTime = CreateMockDateTime("2021-01-01T00:00:00.000Z")
	model.UpdatedTime = CreateMockDateTime("2021-01-01T00:00:00.000Z")
	model.AlertVersionID = CreateMockUUID("3dc02998-0b50-4ea8-b68a-4779d716fa1f")
	model.Name = core.StringPtr("Unique count alert")
	model.Description = core.StringPtr("Example of unique count alert from terraform")
	model.Enabled = core.BoolPtr(true)
	model.Priority = core.StringPtr("p1")
	model.ActiveOn = apisAlertDefinitionActivityScheduleModel
	model.Type = core.StringPtr("flow")
	model.GroupByKeys = []string{"key1", "key2"}
	model.IncidentsSettings = apisAlertDefinitionAlertDefIncidentSettingsModel
	model.NotificationGroup = apisAlertDefinitionAlertDefNotificationGroupModel
	model.EntityLabels = map[string]string{"key1": "testString"}
	model.PhantomMode = core.BoolPtr(false)
	model.Deleted = core.BoolPtr(false)
	model.LogsAnomaly = apisAlertDefinitionLogsAnomalyTypeModel

	result, err := logs.DataSourceIbmLogsAlertDefinitionsAlertDefinitionApisAlertDefinitionAlertDefPropertiesTypeDefinitionLogsAnomalyToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmLogsAlertDefinitionsAlertDefinitionApisAlertDefinitionAlertDefPropertiesTypeDefinitionMetricAnomalyToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		apisAlertDefinitionTimeOfDayModel := make(map[string]interface{})
		apisAlertDefinitionTimeOfDayModel["hours"] = int(14)
		apisAlertDefinitionTimeOfDayModel["minutes"] = int(30)

		apisAlertDefinitionActivityScheduleModel := make(map[string]interface{})
		apisAlertDefinitionActivityScheduleModel["day_of_week"] = []string{"sunday"}
		apisAlertDefinitionActivityScheduleModel["start_time"] = []map[string]interface{}{apisAlertDefinitionTimeOfDayModel}
		apisAlertDefinitionActivityScheduleModel["end_time"] = []map[string]interface{}{apisAlertDefinitionTimeOfDayModel}

		apisAlertDefinitionAlertDefIncidentSettingsModel := make(map[string]interface{})
		apisAlertDefinitionAlertDefIncidentSettingsModel["notify_on"] = "triggered_and_resolved"
		apisAlertDefinitionAlertDefIncidentSettingsModel["minutes"] = int(30)

		apisAlertDefinitionIntegrationTypeModel := make(map[string]interface{})
		apisAlertDefinitionIntegrationTypeModel["integration_id"] = int(123)

		apisAlertDefinitionAlertDefWebhooksSettingsModel := make(map[string]interface{})
		apisAlertDefinitionAlertDefWebhooksSettingsModel["notify_on"] = "triggered_and_resolved"
		apisAlertDefinitionAlertDefWebhooksSettingsModel["integration"] = []map[string]interface{}{apisAlertDefinitionIntegrationTypeModel}
		apisAlertDefinitionAlertDefWebhooksSettingsModel["minutes"] = int(15)

		apisAlertDefinitionAlertDefNotificationGroupModel := make(map[string]interface{})
		apisAlertDefinitionAlertDefNotificationGroupModel["group_by_keys"] = []string{"key1", "key2"}
		apisAlertDefinitionAlertDefNotificationGroupModel["webhooks"] = []map[string]interface{}{apisAlertDefinitionAlertDefWebhooksSettingsModel}

		apisAlertDefinitionMetricFilterModel := make(map[string]interface{})
		apisAlertDefinitionMetricFilterModel["promql"] = "avg_over_time(metric_name[5m]) > 10"

		apisAlertDefinitionMetricTimeWindowModel := make(map[string]interface{})
		apisAlertDefinitionMetricTimeWindowModel["metric_time_window_specific_value"] = "hours_36"

		apisAlertDefinitionMetricAnomalyConditionModel := make(map[string]interface{})
		apisAlertDefinitionMetricAnomalyConditionModel["threshold"] = float64(10.0)
		apisAlertDefinitionMetricAnomalyConditionModel["for_over_pct"] = int(20)
		apisAlertDefinitionMetricAnomalyConditionModel["of_the_last"] = []map[string]interface{}{apisAlertDefinitionMetricTimeWindowModel}
		apisAlertDefinitionMetricAnomalyConditionModel["min_non_null_values_pct"] = int(10)

		apisAlertDefinitionMetricAnomalyRuleModel := make(map[string]interface{})
		apisAlertDefinitionMetricAnomalyRuleModel["condition"] = []map[string]interface{}{apisAlertDefinitionMetricAnomalyConditionModel}

		apisAlertDefinitionAnomalyAlertSettingsModel := make(map[string]interface{})
		apisAlertDefinitionAnomalyAlertSettingsModel["percentage_of_deviation"] = float64(10.0)

		apisAlertDefinitionMetricAnomalyTypeModel := make(map[string]interface{})
		apisAlertDefinitionMetricAnomalyTypeModel["metric_filter"] = []map[string]interface{}{apisAlertDefinitionMetricFilterModel}
		apisAlertDefinitionMetricAnomalyTypeModel["rules"] = []map[string]interface{}{apisAlertDefinitionMetricAnomalyRuleModel}
		apisAlertDefinitionMetricAnomalyTypeModel["condition_type"] = "less_than_usual"
		apisAlertDefinitionMetricAnomalyTypeModel["evaluation_delay_ms"] = int(60000)
		apisAlertDefinitionMetricAnomalyTypeModel["anomaly_alert_settings"] = []map[string]interface{}{apisAlertDefinitionAnomalyAlertSettingsModel}

		model := make(map[string]interface{})
		model["id"] = "3dc02998-0b50-4ea8-b68a-4779d716fa1f"
		model["created_time"] = "2021-01-01T00:00:00.000Z"
		model["updated_time"] = "2021-01-01T00:00:00.000Z"
		model["alert_version_id"] = "3dc02998-0b50-4ea8-b68a-4779d716fa1f"
		model["name"] = "Unique count alert"
		model["description"] = "Example of unique count alert from terraform"
		model["enabled"] = true
		model["priority"] = "p1"
		model["active_on"] = []map[string]interface{}{apisAlertDefinitionActivityScheduleModel}
		model["type"] = "flow"
		model["group_by_keys"] = []string{"key1", "key2"}
		model["incidents_settings"] = []map[string]interface{}{apisAlertDefinitionAlertDefIncidentSettingsModel}
		model["notification_group"] = []map[string]interface{}{apisAlertDefinitionAlertDefNotificationGroupModel}
		model["entity_labels"] = map[string]interface{}{"key1": "testString"}
		model["phantom_mode"] = false
		model["deleted"] = false
		model["metric_anomaly"] = []map[string]interface{}{apisAlertDefinitionMetricAnomalyTypeModel}

		assert.Equal(t, result, model)
	}

	apisAlertDefinitionTimeOfDayModel := new(logsv0.ApisAlertDefinitionTimeOfDay)
	apisAlertDefinitionTimeOfDayModel.Hours = core.Int64Ptr(int64(14))
	apisAlertDefinitionTimeOfDayModel.Minutes = core.Int64Ptr(int64(30))

	apisAlertDefinitionActivityScheduleModel := new(logsv0.ApisAlertDefinitionActivitySchedule)
	apisAlertDefinitionActivityScheduleModel.DayOfWeek = []string{"sunday"}
	apisAlertDefinitionActivityScheduleModel.StartTime = apisAlertDefinitionTimeOfDayModel
	apisAlertDefinitionActivityScheduleModel.EndTime = apisAlertDefinitionTimeOfDayModel

	apisAlertDefinitionAlertDefIncidentSettingsModel := new(logsv0.ApisAlertDefinitionAlertDefIncidentSettings)
	apisAlertDefinitionAlertDefIncidentSettingsModel.NotifyOn = core.StringPtr("triggered_and_resolved")
	apisAlertDefinitionAlertDefIncidentSettingsModel.Minutes = core.Int64Ptr(int64(30))

	apisAlertDefinitionIntegrationTypeModel := new(logsv0.ApisAlertDefinitionIntegrationTypeIntegrationTypeIntegrationID)
	apisAlertDefinitionIntegrationTypeModel.IntegrationID = core.Int64Ptr(int64(123))

	apisAlertDefinitionAlertDefWebhooksSettingsModel := new(logsv0.ApisAlertDefinitionAlertDefWebhooksSettings)
	apisAlertDefinitionAlertDefWebhooksSettingsModel.NotifyOn = core.StringPtr("triggered_and_resolved")
	apisAlertDefinitionAlertDefWebhooksSettingsModel.Integration = apisAlertDefinitionIntegrationTypeModel
	apisAlertDefinitionAlertDefWebhooksSettingsModel.Minutes = core.Int64Ptr(int64(15))

	apisAlertDefinitionAlertDefNotificationGroupModel := new(logsv0.ApisAlertDefinitionAlertDefNotificationGroup)
	apisAlertDefinitionAlertDefNotificationGroupModel.GroupByKeys = []string{"key1", "key2"}
	apisAlertDefinitionAlertDefNotificationGroupModel.Webhooks = []logsv0.ApisAlertDefinitionAlertDefWebhooksSettings{*apisAlertDefinitionAlertDefWebhooksSettingsModel}

	apisAlertDefinitionMetricFilterModel := new(logsv0.ApisAlertDefinitionMetricFilter)
	apisAlertDefinitionMetricFilterModel.Promql = core.StringPtr("avg_over_time(metric_name[5m]) > 10")

	apisAlertDefinitionMetricTimeWindowModel := new(logsv0.ApisAlertDefinitionMetricTimeWindowTypeMetricTimeWindowSpecificValue)
	apisAlertDefinitionMetricTimeWindowModel.MetricTimeWindowSpecificValue = core.StringPtr("hours_36")

	apisAlertDefinitionMetricAnomalyConditionModel := new(logsv0.ApisAlertDefinitionMetricAnomalyCondition)
	apisAlertDefinitionMetricAnomalyConditionModel.Threshold = core.Float64Ptr(float64(10.0))
	apisAlertDefinitionMetricAnomalyConditionModel.ForOverPct = core.Int64Ptr(int64(20))
	apisAlertDefinitionMetricAnomalyConditionModel.OfTheLast = apisAlertDefinitionMetricTimeWindowModel
	apisAlertDefinitionMetricAnomalyConditionModel.MinNonNullValuesPct = core.Int64Ptr(int64(10))

	apisAlertDefinitionMetricAnomalyRuleModel := new(logsv0.ApisAlertDefinitionMetricAnomalyRule)
	apisAlertDefinitionMetricAnomalyRuleModel.Condition = apisAlertDefinitionMetricAnomalyConditionModel

	apisAlertDefinitionAnomalyAlertSettingsModel := new(logsv0.ApisAlertDefinitionAnomalyAlertSettings)
	apisAlertDefinitionAnomalyAlertSettingsModel.PercentageOfDeviation = core.Float32Ptr(float32(10.0))

	apisAlertDefinitionMetricAnomalyTypeModel := new(logsv0.ApisAlertDefinitionMetricAnomalyType)
	apisAlertDefinitionMetricAnomalyTypeModel.MetricFilter = apisAlertDefinitionMetricFilterModel
	apisAlertDefinitionMetricAnomalyTypeModel.Rules = []logsv0.ApisAlertDefinitionMetricAnomalyRule{*apisAlertDefinitionMetricAnomalyRuleModel}
	apisAlertDefinitionMetricAnomalyTypeModel.ConditionType = core.StringPtr("less_than_usual")
	apisAlertDefinitionMetricAnomalyTypeModel.EvaluationDelayMs = core.Int64Ptr(int64(60000))
	apisAlertDefinitionMetricAnomalyTypeModel.AnomalyAlertSettings = apisAlertDefinitionAnomalyAlertSettingsModel

	model := new(logsv0.AlertDefinitionApisAlertDefinitionAlertDefPropertiesTypeDefinitionMetricAnomaly)
	model.ID = CreateMockUUID("3dc02998-0b50-4ea8-b68a-4779d716fa1f")
	model.CreatedTime = CreateMockDateTime("2021-01-01T00:00:00.000Z")
	model.UpdatedTime = CreateMockDateTime("2021-01-01T00:00:00.000Z")
	model.AlertVersionID = CreateMockUUID("3dc02998-0b50-4ea8-b68a-4779d716fa1f")
	model.Name = core.StringPtr("Unique count alert")
	model.Description = core.StringPtr("Example of unique count alert from terraform")
	model.Enabled = core.BoolPtr(true)
	model.Priority = core.StringPtr("p1")
	model.ActiveOn = apisAlertDefinitionActivityScheduleModel
	model.Type = core.StringPtr("flow")
	model.GroupByKeys = []string{"key1", "key2"}
	model.IncidentsSettings = apisAlertDefinitionAlertDefIncidentSettingsModel
	model.NotificationGroup = apisAlertDefinitionAlertDefNotificationGroupModel
	model.EntityLabels = map[string]string{"key1": "testString"}
	model.PhantomMode = core.BoolPtr(false)
	model.Deleted = core.BoolPtr(false)
	model.MetricAnomaly = apisAlertDefinitionMetricAnomalyTypeModel

	result, err := logs.DataSourceIbmLogsAlertDefinitionsAlertDefinitionApisAlertDefinitionAlertDefPropertiesTypeDefinitionMetricAnomalyToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmLogsAlertDefinitionsAlertDefinitionApisAlertDefinitionAlertDefPropertiesTypeDefinitionLogsNewValueToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		apisAlertDefinitionTimeOfDayModel := make(map[string]interface{})
		apisAlertDefinitionTimeOfDayModel["hours"] = int(14)
		apisAlertDefinitionTimeOfDayModel["minutes"] = int(30)

		apisAlertDefinitionActivityScheduleModel := make(map[string]interface{})
		apisAlertDefinitionActivityScheduleModel["day_of_week"] = []string{"sunday"}
		apisAlertDefinitionActivityScheduleModel["start_time"] = []map[string]interface{}{apisAlertDefinitionTimeOfDayModel}
		apisAlertDefinitionActivityScheduleModel["end_time"] = []map[string]interface{}{apisAlertDefinitionTimeOfDayModel}

		apisAlertDefinitionAlertDefIncidentSettingsModel := make(map[string]interface{})
		apisAlertDefinitionAlertDefIncidentSettingsModel["notify_on"] = "triggered_and_resolved"
		apisAlertDefinitionAlertDefIncidentSettingsModel["minutes"] = int(30)

		apisAlertDefinitionIntegrationTypeModel := make(map[string]interface{})
		apisAlertDefinitionIntegrationTypeModel["integration_id"] = int(123)

		apisAlertDefinitionAlertDefWebhooksSettingsModel := make(map[string]interface{})
		apisAlertDefinitionAlertDefWebhooksSettingsModel["notify_on"] = "triggered_and_resolved"
		apisAlertDefinitionAlertDefWebhooksSettingsModel["integration"] = []map[string]interface{}{apisAlertDefinitionIntegrationTypeModel}
		apisAlertDefinitionAlertDefWebhooksSettingsModel["minutes"] = int(15)

		apisAlertDefinitionAlertDefNotificationGroupModel := make(map[string]interface{})
		apisAlertDefinitionAlertDefNotificationGroupModel["group_by_keys"] = []string{"key1", "key2"}
		apisAlertDefinitionAlertDefNotificationGroupModel["webhooks"] = []map[string]interface{}{apisAlertDefinitionAlertDefWebhooksSettingsModel}

		apisAlertDefinitionLabelFilterTypeModel := make(map[string]interface{})
		apisAlertDefinitionLabelFilterTypeModel["value"] = "my-app"
		apisAlertDefinitionLabelFilterTypeModel["operation"] = "starts_with"

		apisAlertDefinitionLabelFiltersModel := make(map[string]interface{})
		apisAlertDefinitionLabelFiltersModel["application_name"] = []map[string]interface{}{apisAlertDefinitionLabelFilterTypeModel}
		apisAlertDefinitionLabelFiltersModel["subsystem_name"] = []map[string]interface{}{apisAlertDefinitionLabelFilterTypeModel}
		apisAlertDefinitionLabelFiltersModel["severities"] = []string{"critical"}

		apisAlertDefinitionLogsSimpleFilterModel := make(map[string]interface{})
		apisAlertDefinitionLogsSimpleFilterModel["lucene_query"] = "text:"error""
		apisAlertDefinitionLogsSimpleFilterModel["label_filters"] = []map[string]interface{}{apisAlertDefinitionLabelFiltersModel}

		apisAlertDefinitionLogsFilterModel := make(map[string]interface{})
		apisAlertDefinitionLogsFilterModel["simple_filter"] = []map[string]interface{}{apisAlertDefinitionLogsSimpleFilterModel}

		apisAlertDefinitionLogsNewValueTimeWindowModel := make(map[string]interface{})
		apisAlertDefinitionLogsNewValueTimeWindowModel["logs_new_value_time_window_specific_value"] = "months_3"

		apisAlertDefinitionLogsNewValueConditionModel := make(map[string]interface{})
		apisAlertDefinitionLogsNewValueConditionModel["keypath_to_track"] = "metadata.field"
		apisAlertDefinitionLogsNewValueConditionModel["time_window"] = []map[string]interface{}{apisAlertDefinitionLogsNewValueTimeWindowModel}

		apisAlertDefinitionLogsNewValueRuleModel := make(map[string]interface{})
		apisAlertDefinitionLogsNewValueRuleModel["condition"] = []map[string]interface{}{apisAlertDefinitionLogsNewValueConditionModel}

		apisAlertDefinitionLogsNewValueTypeModel := make(map[string]interface{})
		apisAlertDefinitionLogsNewValueTypeModel["logs_filter"] = []map[string]interface{}{apisAlertDefinitionLogsFilterModel}
		apisAlertDefinitionLogsNewValueTypeModel["rules"] = []map[string]interface{}{apisAlertDefinitionLogsNewValueRuleModel}
		apisAlertDefinitionLogsNewValueTypeModel["notification_payload_filter"] = []string{"obj.field"}

		model := make(map[string]interface{})
		model["id"] = "3dc02998-0b50-4ea8-b68a-4779d716fa1f"
		model["created_time"] = "2021-01-01T00:00:00.000Z"
		model["updated_time"] = "2021-01-01T00:00:00.000Z"
		model["alert_version_id"] = "3dc02998-0b50-4ea8-b68a-4779d716fa1f"
		model["name"] = "Unique count alert"
		model["description"] = "Example of unique count alert from terraform"
		model["enabled"] = true
		model["priority"] = "p1"
		model["active_on"] = []map[string]interface{}{apisAlertDefinitionActivityScheduleModel}
		model["type"] = "flow"
		model["group_by_keys"] = []string{"key1", "key2"}
		model["incidents_settings"] = []map[string]interface{}{apisAlertDefinitionAlertDefIncidentSettingsModel}
		model["notification_group"] = []map[string]interface{}{apisAlertDefinitionAlertDefNotificationGroupModel}
		model["entity_labels"] = map[string]interface{}{"key1": "testString"}
		model["phantom_mode"] = false
		model["deleted"] = false
		model["logs_new_value"] = []map[string]interface{}{apisAlertDefinitionLogsNewValueTypeModel}

		assert.Equal(t, result, model)
	}

	apisAlertDefinitionTimeOfDayModel := new(logsv0.ApisAlertDefinitionTimeOfDay)
	apisAlertDefinitionTimeOfDayModel.Hours = core.Int64Ptr(int64(14))
	apisAlertDefinitionTimeOfDayModel.Minutes = core.Int64Ptr(int64(30))

	apisAlertDefinitionActivityScheduleModel := new(logsv0.ApisAlertDefinitionActivitySchedule)
	apisAlertDefinitionActivityScheduleModel.DayOfWeek = []string{"sunday"}
	apisAlertDefinitionActivityScheduleModel.StartTime = apisAlertDefinitionTimeOfDayModel
	apisAlertDefinitionActivityScheduleModel.EndTime = apisAlertDefinitionTimeOfDayModel

	apisAlertDefinitionAlertDefIncidentSettingsModel := new(logsv0.ApisAlertDefinitionAlertDefIncidentSettings)
	apisAlertDefinitionAlertDefIncidentSettingsModel.NotifyOn = core.StringPtr("triggered_and_resolved")
	apisAlertDefinitionAlertDefIncidentSettingsModel.Minutes = core.Int64Ptr(int64(30))

	apisAlertDefinitionIntegrationTypeModel := new(logsv0.ApisAlertDefinitionIntegrationTypeIntegrationTypeIntegrationID)
	apisAlertDefinitionIntegrationTypeModel.IntegrationID = core.Int64Ptr(int64(123))

	apisAlertDefinitionAlertDefWebhooksSettingsModel := new(logsv0.ApisAlertDefinitionAlertDefWebhooksSettings)
	apisAlertDefinitionAlertDefWebhooksSettingsModel.NotifyOn = core.StringPtr("triggered_and_resolved")
	apisAlertDefinitionAlertDefWebhooksSettingsModel.Integration = apisAlertDefinitionIntegrationTypeModel
	apisAlertDefinitionAlertDefWebhooksSettingsModel.Minutes = core.Int64Ptr(int64(15))

	apisAlertDefinitionAlertDefNotificationGroupModel := new(logsv0.ApisAlertDefinitionAlertDefNotificationGroup)
	apisAlertDefinitionAlertDefNotificationGroupModel.GroupByKeys = []string{"key1", "key2"}
	apisAlertDefinitionAlertDefNotificationGroupModel.Webhooks = []logsv0.ApisAlertDefinitionAlertDefWebhooksSettings{*apisAlertDefinitionAlertDefWebhooksSettingsModel}

	apisAlertDefinitionLabelFilterTypeModel := new(logsv0.ApisAlertDefinitionLabelFilterType)
	apisAlertDefinitionLabelFilterTypeModel.Value = core.StringPtr("my-app")
	apisAlertDefinitionLabelFilterTypeModel.Operation = core.StringPtr("starts_with")

	apisAlertDefinitionLabelFiltersModel := new(logsv0.ApisAlertDefinitionLabelFilters)
	apisAlertDefinitionLabelFiltersModel.ApplicationName = []logsv0.ApisAlertDefinitionLabelFilterType{*apisAlertDefinitionLabelFilterTypeModel}
	apisAlertDefinitionLabelFiltersModel.SubsystemName = []logsv0.ApisAlertDefinitionLabelFilterType{*apisAlertDefinitionLabelFilterTypeModel}
	apisAlertDefinitionLabelFiltersModel.Severities = []string{"critical"}

	apisAlertDefinitionLogsSimpleFilterModel := new(logsv0.ApisAlertDefinitionLogsSimpleFilter)
	apisAlertDefinitionLogsSimpleFilterModel.LuceneQuery = core.StringPtr("text:"error"")
	apisAlertDefinitionLogsSimpleFilterModel.LabelFilters = apisAlertDefinitionLabelFiltersModel

	apisAlertDefinitionLogsFilterModel := new(logsv0.ApisAlertDefinitionLogsFilter)
	apisAlertDefinitionLogsFilterModel.SimpleFilter = apisAlertDefinitionLogsSimpleFilterModel

	apisAlertDefinitionLogsNewValueTimeWindowModel := new(logsv0.ApisAlertDefinitionLogsNewValueTimeWindow)
	apisAlertDefinitionLogsNewValueTimeWindowModel.LogsNewValueTimeWindowSpecificValue = core.StringPtr("months_3")

	apisAlertDefinitionLogsNewValueConditionModel := new(logsv0.ApisAlertDefinitionLogsNewValueCondition)
	apisAlertDefinitionLogsNewValueConditionModel.KeypathToTrack = core.StringPtr("metadata.field")
	apisAlertDefinitionLogsNewValueConditionModel.TimeWindow = apisAlertDefinitionLogsNewValueTimeWindowModel

	apisAlertDefinitionLogsNewValueRuleModel := new(logsv0.ApisAlertDefinitionLogsNewValueRule)
	apisAlertDefinitionLogsNewValueRuleModel.Condition = apisAlertDefinitionLogsNewValueConditionModel

	apisAlertDefinitionLogsNewValueTypeModel := new(logsv0.ApisAlertDefinitionLogsNewValueType)
	apisAlertDefinitionLogsNewValueTypeModel.LogsFilter = apisAlertDefinitionLogsFilterModel
	apisAlertDefinitionLogsNewValueTypeModel.Rules = []logsv0.ApisAlertDefinitionLogsNewValueRule{*apisAlertDefinitionLogsNewValueRuleModel}
	apisAlertDefinitionLogsNewValueTypeModel.NotificationPayloadFilter = []string{"obj.field"}

	model := new(logsv0.AlertDefinitionApisAlertDefinitionAlertDefPropertiesTypeDefinitionLogsNewValue)
	model.ID = CreateMockUUID("3dc02998-0b50-4ea8-b68a-4779d716fa1f")
	model.CreatedTime = CreateMockDateTime("2021-01-01T00:00:00.000Z")
	model.UpdatedTime = CreateMockDateTime("2021-01-01T00:00:00.000Z")
	model.AlertVersionID = CreateMockUUID("3dc02998-0b50-4ea8-b68a-4779d716fa1f")
	model.Name = core.StringPtr("Unique count alert")
	model.Description = core.StringPtr("Example of unique count alert from terraform")
	model.Enabled = core.BoolPtr(true)
	model.Priority = core.StringPtr("p1")
	model.ActiveOn = apisAlertDefinitionActivityScheduleModel
	model.Type = core.StringPtr("flow")
	model.GroupByKeys = []string{"key1", "key2"}
	model.IncidentsSettings = apisAlertDefinitionAlertDefIncidentSettingsModel
	model.NotificationGroup = apisAlertDefinitionAlertDefNotificationGroupModel
	model.EntityLabels = map[string]string{"key1": "testString"}
	model.PhantomMode = core.BoolPtr(false)
	model.Deleted = core.BoolPtr(false)
	model.LogsNewValue = apisAlertDefinitionLogsNewValueTypeModel

	result, err := logs.DataSourceIbmLogsAlertDefinitionsAlertDefinitionApisAlertDefinitionAlertDefPropertiesTypeDefinitionLogsNewValueToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmLogsAlertDefinitionsAlertDefinitionApisAlertDefinitionAlertDefPropertiesTypeDefinitionLogsUniqueCountToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		apisAlertDefinitionTimeOfDayModel := make(map[string]interface{})
		apisAlertDefinitionTimeOfDayModel["hours"] = int(14)
		apisAlertDefinitionTimeOfDayModel["minutes"] = int(30)

		apisAlertDefinitionActivityScheduleModel := make(map[string]interface{})
		apisAlertDefinitionActivityScheduleModel["day_of_week"] = []string{"sunday"}
		apisAlertDefinitionActivityScheduleModel["start_time"] = []map[string]interface{}{apisAlertDefinitionTimeOfDayModel}
		apisAlertDefinitionActivityScheduleModel["end_time"] = []map[string]interface{}{apisAlertDefinitionTimeOfDayModel}

		apisAlertDefinitionAlertDefIncidentSettingsModel := make(map[string]interface{})
		apisAlertDefinitionAlertDefIncidentSettingsModel["notify_on"] = "triggered_and_resolved"
		apisAlertDefinitionAlertDefIncidentSettingsModel["minutes"] = int(30)

		apisAlertDefinitionIntegrationTypeModel := make(map[string]interface{})
		apisAlertDefinitionIntegrationTypeModel["integration_id"] = int(123)

		apisAlertDefinitionAlertDefWebhooksSettingsModel := make(map[string]interface{})
		apisAlertDefinitionAlertDefWebhooksSettingsModel["notify_on"] = "triggered_and_resolved"
		apisAlertDefinitionAlertDefWebhooksSettingsModel["integration"] = []map[string]interface{}{apisAlertDefinitionIntegrationTypeModel}
		apisAlertDefinitionAlertDefWebhooksSettingsModel["minutes"] = int(15)

		apisAlertDefinitionAlertDefNotificationGroupModel := make(map[string]interface{})
		apisAlertDefinitionAlertDefNotificationGroupModel["group_by_keys"] = []string{"key1", "key2"}
		apisAlertDefinitionAlertDefNotificationGroupModel["webhooks"] = []map[string]interface{}{apisAlertDefinitionAlertDefWebhooksSettingsModel}

		apisAlertDefinitionLabelFilterTypeModel := make(map[string]interface{})
		apisAlertDefinitionLabelFilterTypeModel["value"] = "my-app"
		apisAlertDefinitionLabelFilterTypeModel["operation"] = "starts_with"

		apisAlertDefinitionLabelFiltersModel := make(map[string]interface{})
		apisAlertDefinitionLabelFiltersModel["application_name"] = []map[string]interface{}{apisAlertDefinitionLabelFilterTypeModel}
		apisAlertDefinitionLabelFiltersModel["subsystem_name"] = []map[string]interface{}{apisAlertDefinitionLabelFilterTypeModel}
		apisAlertDefinitionLabelFiltersModel["severities"] = []string{"critical"}

		apisAlertDefinitionLogsSimpleFilterModel := make(map[string]interface{})
		apisAlertDefinitionLogsSimpleFilterModel["lucene_query"] = "text:"error""
		apisAlertDefinitionLogsSimpleFilterModel["label_filters"] = []map[string]interface{}{apisAlertDefinitionLabelFiltersModel}

		apisAlertDefinitionLogsFilterModel := make(map[string]interface{})
		apisAlertDefinitionLogsFilterModel["simple_filter"] = []map[string]interface{}{apisAlertDefinitionLogsSimpleFilterModel}

		apisAlertDefinitionLogsUniqueValueTimeWindowModel := make(map[string]interface{})
		apisAlertDefinitionLogsUniqueValueTimeWindowModel["logs_unique_value_time_window_specific_value"] = "hours_36"

		apisAlertDefinitionLogsUniqueCountConditionModel := make(map[string]interface{})
		apisAlertDefinitionLogsUniqueCountConditionModel["max_unique_count"] = "100"
		apisAlertDefinitionLogsUniqueCountConditionModel["time_window"] = []map[string]interface{}{apisAlertDefinitionLogsUniqueValueTimeWindowModel}

		apisAlertDefinitionLogsUniqueCountRuleModel := make(map[string]interface{})
		apisAlertDefinitionLogsUniqueCountRuleModel["condition"] = []map[string]interface{}{apisAlertDefinitionLogsUniqueCountConditionModel}

		apisAlertDefinitionLogsUniqueCountTypeModel := make(map[string]interface{})
		apisAlertDefinitionLogsUniqueCountTypeModel["logs_filter"] = []map[string]interface{}{apisAlertDefinitionLogsFilterModel}
		apisAlertDefinitionLogsUniqueCountTypeModel["rules"] = []map[string]interface{}{apisAlertDefinitionLogsUniqueCountRuleModel}
		apisAlertDefinitionLogsUniqueCountTypeModel["notification_payload_filter"] = []string{"obj.field"}
		apisAlertDefinitionLogsUniqueCountTypeModel["max_unique_count_per_group_by_key"] = "100"
		apisAlertDefinitionLogsUniqueCountTypeModel["unique_count_keypath"] = "obj.field"

		model := make(map[string]interface{})
		model["id"] = "3dc02998-0b50-4ea8-b68a-4779d716fa1f"
		model["created_time"] = "2021-01-01T00:00:00.000Z"
		model["updated_time"] = "2021-01-01T00:00:00.000Z"
		model["alert_version_id"] = "3dc02998-0b50-4ea8-b68a-4779d716fa1f"
		model["name"] = "Unique count alert"
		model["description"] = "Example of unique count alert from terraform"
		model["enabled"] = true
		model["priority"] = "p1"
		model["active_on"] = []map[string]interface{}{apisAlertDefinitionActivityScheduleModel}
		model["type"] = "flow"
		model["group_by_keys"] = []string{"key1", "key2"}
		model["incidents_settings"] = []map[string]interface{}{apisAlertDefinitionAlertDefIncidentSettingsModel}
		model["notification_group"] = []map[string]interface{}{apisAlertDefinitionAlertDefNotificationGroupModel}
		model["entity_labels"] = map[string]interface{}{"key1": "testString"}
		model["phantom_mode"] = false
		model["deleted"] = false
		model["logs_unique_count"] = []map[string]interface{}{apisAlertDefinitionLogsUniqueCountTypeModel}

		assert.Equal(t, result, model)
	}

	apisAlertDefinitionTimeOfDayModel := new(logsv0.ApisAlertDefinitionTimeOfDay)
	apisAlertDefinitionTimeOfDayModel.Hours = core.Int64Ptr(int64(14))
	apisAlertDefinitionTimeOfDayModel.Minutes = core.Int64Ptr(int64(30))

	apisAlertDefinitionActivityScheduleModel := new(logsv0.ApisAlertDefinitionActivitySchedule)
	apisAlertDefinitionActivityScheduleModel.DayOfWeek = []string{"sunday"}
	apisAlertDefinitionActivityScheduleModel.StartTime = apisAlertDefinitionTimeOfDayModel
	apisAlertDefinitionActivityScheduleModel.EndTime = apisAlertDefinitionTimeOfDayModel

	apisAlertDefinitionAlertDefIncidentSettingsModel := new(logsv0.ApisAlertDefinitionAlertDefIncidentSettings)
	apisAlertDefinitionAlertDefIncidentSettingsModel.NotifyOn = core.StringPtr("triggered_and_resolved")
	apisAlertDefinitionAlertDefIncidentSettingsModel.Minutes = core.Int64Ptr(int64(30))

	apisAlertDefinitionIntegrationTypeModel := new(logsv0.ApisAlertDefinitionIntegrationTypeIntegrationTypeIntegrationID)
	apisAlertDefinitionIntegrationTypeModel.IntegrationID = core.Int64Ptr(int64(123))

	apisAlertDefinitionAlertDefWebhooksSettingsModel := new(logsv0.ApisAlertDefinitionAlertDefWebhooksSettings)
	apisAlertDefinitionAlertDefWebhooksSettingsModel.NotifyOn = core.StringPtr("triggered_and_resolved")
	apisAlertDefinitionAlertDefWebhooksSettingsModel.Integration = apisAlertDefinitionIntegrationTypeModel
	apisAlertDefinitionAlertDefWebhooksSettingsModel.Minutes = core.Int64Ptr(int64(15))

	apisAlertDefinitionAlertDefNotificationGroupModel := new(logsv0.ApisAlertDefinitionAlertDefNotificationGroup)
	apisAlertDefinitionAlertDefNotificationGroupModel.GroupByKeys = []string{"key1", "key2"}
	apisAlertDefinitionAlertDefNotificationGroupModel.Webhooks = []logsv0.ApisAlertDefinitionAlertDefWebhooksSettings{*apisAlertDefinitionAlertDefWebhooksSettingsModel}

	apisAlertDefinitionLabelFilterTypeModel := new(logsv0.ApisAlertDefinitionLabelFilterType)
	apisAlertDefinitionLabelFilterTypeModel.Value = core.StringPtr("my-app")
	apisAlertDefinitionLabelFilterTypeModel.Operation = core.StringPtr("starts_with")

	apisAlertDefinitionLabelFiltersModel := new(logsv0.ApisAlertDefinitionLabelFilters)
	apisAlertDefinitionLabelFiltersModel.ApplicationName = []logsv0.ApisAlertDefinitionLabelFilterType{*apisAlertDefinitionLabelFilterTypeModel}
	apisAlertDefinitionLabelFiltersModel.SubsystemName = []logsv0.ApisAlertDefinitionLabelFilterType{*apisAlertDefinitionLabelFilterTypeModel}
	apisAlertDefinitionLabelFiltersModel.Severities = []string{"critical"}

	apisAlertDefinitionLogsSimpleFilterModel := new(logsv0.ApisAlertDefinitionLogsSimpleFilter)
	apisAlertDefinitionLogsSimpleFilterModel.LuceneQuery = core.StringPtr("text:"error"")
	apisAlertDefinitionLogsSimpleFilterModel.LabelFilters = apisAlertDefinitionLabelFiltersModel

	apisAlertDefinitionLogsFilterModel := new(logsv0.ApisAlertDefinitionLogsFilter)
	apisAlertDefinitionLogsFilterModel.SimpleFilter = apisAlertDefinitionLogsSimpleFilterModel

	apisAlertDefinitionLogsUniqueValueTimeWindowModel := new(logsv0.ApisAlertDefinitionLogsUniqueValueTimeWindow)
	apisAlertDefinitionLogsUniqueValueTimeWindowModel.LogsUniqueValueTimeWindowSpecificValue = core.StringPtr("hours_36")

	apisAlertDefinitionLogsUniqueCountConditionModel := new(logsv0.ApisAlertDefinitionLogsUniqueCountCondition)
	apisAlertDefinitionLogsUniqueCountConditionModel.MaxUniqueCount = core.StringPtr("100")
	apisAlertDefinitionLogsUniqueCountConditionModel.TimeWindow = apisAlertDefinitionLogsUniqueValueTimeWindowModel

	apisAlertDefinitionLogsUniqueCountRuleModel := new(logsv0.ApisAlertDefinitionLogsUniqueCountRule)
	apisAlertDefinitionLogsUniqueCountRuleModel.Condition = apisAlertDefinitionLogsUniqueCountConditionModel

	apisAlertDefinitionLogsUniqueCountTypeModel := new(logsv0.ApisAlertDefinitionLogsUniqueCountType)
	apisAlertDefinitionLogsUniqueCountTypeModel.LogsFilter = apisAlertDefinitionLogsFilterModel
	apisAlertDefinitionLogsUniqueCountTypeModel.Rules = []logsv0.ApisAlertDefinitionLogsUniqueCountRule{*apisAlertDefinitionLogsUniqueCountRuleModel}
	apisAlertDefinitionLogsUniqueCountTypeModel.NotificationPayloadFilter = []string{"obj.field"}
	apisAlertDefinitionLogsUniqueCountTypeModel.MaxUniqueCountPerGroupByKey = core.StringPtr("100")
	apisAlertDefinitionLogsUniqueCountTypeModel.UniqueCountKeypath = core.StringPtr("obj.field")

	model := new(logsv0.AlertDefinitionApisAlertDefinitionAlertDefPropertiesTypeDefinitionLogsUniqueCount)
	model.ID = CreateMockUUID("3dc02998-0b50-4ea8-b68a-4779d716fa1f")
	model.CreatedTime = CreateMockDateTime("2021-01-01T00:00:00.000Z")
	model.UpdatedTime = CreateMockDateTime("2021-01-01T00:00:00.000Z")
	model.AlertVersionID = CreateMockUUID("3dc02998-0b50-4ea8-b68a-4779d716fa1f")
	model.Name = core.StringPtr("Unique count alert")
	model.Description = core.StringPtr("Example of unique count alert from terraform")
	model.Enabled = core.BoolPtr(true)
	model.Priority = core.StringPtr("p1")
	model.ActiveOn = apisAlertDefinitionActivityScheduleModel
	model.Type = core.StringPtr("flow")
	model.GroupByKeys = []string{"key1", "key2"}
	model.IncidentsSettings = apisAlertDefinitionAlertDefIncidentSettingsModel
	model.NotificationGroup = apisAlertDefinitionAlertDefNotificationGroupModel
	model.EntityLabels = map[string]string{"key1": "testString"}
	model.PhantomMode = core.BoolPtr(false)
	model.Deleted = core.BoolPtr(false)
	model.LogsUniqueCount = apisAlertDefinitionLogsUniqueCountTypeModel

	result, err := logs.DataSourceIbmLogsAlertDefinitionsAlertDefinitionApisAlertDefinitionAlertDefPropertiesTypeDefinitionLogsUniqueCountToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}
