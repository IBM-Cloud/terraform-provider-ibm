// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package logs_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/logs"
	. "github.com/IBM-Cloud/terraform-provider-ibm/ibm/unittest"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/observability-c/dragonlog-logs-go-sdk/logsv0"
	"github.com/stretchr/testify/assert"
)

func TestAccIbmLogsAlertsDataSourceBasic(t *testing.T) {
	alertName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	alertIsActive := "false"
	alertSeverity := "info_or_unspecified"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
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
		PreCheck:  func() { acc.TestAccPreCheck(t) },
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
			name = "%s"
			is_active = %s
			severity = "%s"
			condition {
				immediate = {  }
			}
			notification_groups {
				group_by_fields = [ "group_by_fields" ]
				notifications {
					retriggering_period_seconds = 0
					notify_on = "triggered_only"
					integration_id = 0
				}
			}
			filters {
				severities = [ "debug_or_unspecified" ]
				metadata {
					categories = [ "categories" ]
					applications = [ "applications" ]
					subsystems = [ "subsystems" ]
					computers = [ "computers" ]
					classes = [ "classes" ]
					methods = [ "methods" ]
					ip_addresses = [ "ip_addresses" ]
				}
				alias = "alias"
				text = "text"
				ratio_alerts {
					alias = "alias"
					text = "text"
					severities = [ "debug_or_unspecified" ]
					applications = [ "applications" ]
					subsystems = [ "subsystems" ]
					group_by = [ "group_by" ]
				}
				filter_type = "text_or_unspecified"
			}
		}

		data "ibm_logs_alerts" "logs_alerts_instance" {
			depends_on = [
				ibm_logs_alert.logs_alert_instance
			]
		}
	`, alertName, alertIsActive, alertSeverity)
}

func testAccCheckIbmLogsAlertsDataSourceConfig(alertName string, alertDescription string, alertIsActive string, alertSeverity string) string {
	return fmt.Sprintf(`
		resource "ibm_logs_alert" "logs_alert_instance" {
			name = "%s"
			description = "%s"
			is_active = %s
			severity = "%s"
			expiration {
				year = 1
				month = 1
				day = 1
			}
			condition {
				immediate = {  }
			}
			notification_groups {
				group_by_fields = [ "group_by_fields" ]
				notifications {
					retriggering_period_seconds = 0
					notify_on = "triggered_only"
					integration_id = 0
				}
			}
			filters {
				severities = [ "debug_or_unspecified" ]
				metadata {
					categories = [ "categories" ]
					applications = [ "applications" ]
					subsystems = [ "subsystems" ]
					computers = [ "computers" ]
					classes = [ "classes" ]
					methods = [ "methods" ]
					ip_addresses = [ "ip_addresses" ]
				}
				alias = "alias"
				text = "text"
				ratio_alerts {
					alias = "alias"
					text = "text"
					severities = [ "debug_or_unspecified" ]
					applications = [ "applications" ]
					subsystems = [ "subsystems" ]
					group_by = [ "group_by" ]
				}
				filter_type = "text_or_unspecified"
			}
			active_when {
				timeframes {
					days_of_week = [ "monday_or_unspecified" ]
					range {
						start {
							hours = 1
							minutes = 1
							seconds = 1
						}
						end {
							hours = 1
							minutes = 1
							seconds = 1
						}
					}
				}
			}
			notification_payload_filters = "FIXME"
			meta_labels {
				key = "key"
				value = "value"
			}
			meta_labels_strings = "FIXME"
			tracing_alert {
				condition_latency = 0
				field_filters {
					field = "field"
					filters {
						values = [ "values" ]
						operator = "operator"
					}
				}
				tag_filters {
					field = "field"
					filters {
						values = [ "values" ]
						operator = "operator"
					}
				}
			}
			incident_settings {
				retriggering_period_seconds = 0
				notify_on = "triggered_only"
				use_as_notification_settings = true
			}
		}

		data "ibm_logs_alerts" "logs_alerts_instance" {
			depends_on = [
				ibm_logs_alert.logs_alert_instance
			]
		}
	`, alertName, alertDescription, alertIsActive, alertSeverity)
}

func TestDataSourceIbmLogsAlertsAlertToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		alertsV1DateModel := make(map[string]interface{})
		alertsV1DateModel["year"] = int(38)
		alertsV1DateModel["month"] = int(38)
		alertsV1DateModel["day"] = int(38)

		alertsV1MetricAlertConditionParametersModel := make(map[string]interface{})
		alertsV1MetricAlertConditionParametersModel["metric_field"] = "testString"
		alertsV1MetricAlertConditionParametersModel["metric_source"] = "logs2metrics_or_unspecified"
		alertsV1MetricAlertConditionParametersModel["arithmetic_operator"] = "avg_or_unspecified"
		alertsV1MetricAlertConditionParametersModel["arithmetic_operator_modifier"] = int(0)
		alertsV1MetricAlertConditionParametersModel["sample_threshold_percentage"] = int(0)
		alertsV1MetricAlertConditionParametersModel["non_null_percentage"] = int(0)
		alertsV1MetricAlertConditionParametersModel["swap_null_values"] = true

		alertsV1MetricAlertPromqlConditionParametersModel := make(map[string]interface{})
		alertsV1MetricAlertPromqlConditionParametersModel["promql_text"] = "testString"
		alertsV1MetricAlertPromqlConditionParametersModel["arithmetic_operator_modifier"] = int(0)
		alertsV1MetricAlertPromqlConditionParametersModel["sample_threshold_percentage"] = int(0)
		alertsV1MetricAlertPromqlConditionParametersModel["non_null_percentage"] = int(0)
		alertsV1MetricAlertPromqlConditionParametersModel["swap_null_values"] = true

		alertsV1RelatedExtendedDataModel := make(map[string]interface{})
		alertsV1RelatedExtendedDataModel["cleanup_deadman_duration"] = "cleanup_deadman_duration_never_or_unspecified"
		alertsV1RelatedExtendedDataModel["should_trigger_deadman"] = true

		alertsV2ConditionParametersModel := make(map[string]interface{})
		alertsV2ConditionParametersModel["threshold"] = float64(1)
		alertsV2ConditionParametersModel["timeframe"] = "timeframe_10_min"
		alertsV2ConditionParametersModel["group_by"] = []string{"coralogix.metadata.applicationName"}
		alertsV2ConditionParametersModel["metric_alert_parameters"] = []map[string]interface{}{alertsV1MetricAlertConditionParametersModel}
		alertsV2ConditionParametersModel["metric_alert_promql_parameters"] = []map[string]interface{}{alertsV1MetricAlertPromqlConditionParametersModel}
		alertsV2ConditionParametersModel["ignore_infinity"] = true
		alertsV2ConditionParametersModel["relative_timeframe"] = "hour_or_unspecified"
		alertsV2ConditionParametersModel["cardinality_fields"] = []string{"testString"}
		alertsV2ConditionParametersModel["related_extended_data"] = []map[string]interface{}{alertsV1RelatedExtendedDataModel}

		alertsV2MoreThanConditionModel := make(map[string]interface{})
		alertsV2MoreThanConditionModel["parameters"] = []map[string]interface{}{alertsV2ConditionParametersModel}
		alertsV2MoreThanConditionModel["evaluation_window"] = "rolling_or_unspecified"

		alertsV2AlertConditionModel := make(map[string]interface{})
		alertsV2AlertConditionModel["more_than"] = []map[string]interface{}{alertsV2MoreThanConditionModel}

		alertsV2AlertNotificationModel := make(map[string]interface{})
		alertsV2AlertNotificationModel["retriggering_period_seconds"] = int(0)
		alertsV2AlertNotificationModel["notify_on"] = "triggered_only"
		alertsV2AlertNotificationModel["integration_id"] = int(0)

		alertsV2AlertNotificationGroupsModel := make(map[string]interface{})
		alertsV2AlertNotificationGroupsModel["group_by_fields"] = []string{"coralogix.metadata.applicationName"}
		alertsV2AlertNotificationGroupsModel["notifications"] = []map[string]interface{}{alertsV2AlertNotificationModel}

		alertsV1AlertFiltersMetadataFiltersModel := make(map[string]interface{})
		alertsV1AlertFiltersMetadataFiltersModel["categories"] = []string{"testString"}
		alertsV1AlertFiltersMetadataFiltersModel["applications"] = []string{"testString"}
		alertsV1AlertFiltersMetadataFiltersModel["subsystems"] = []string{"testString"}
		alertsV1AlertFiltersMetadataFiltersModel["computers"] = []string{"testString"}
		alertsV1AlertFiltersMetadataFiltersModel["classes"] = []string{"testString"}
		alertsV1AlertFiltersMetadataFiltersModel["methods"] = []string{"testString"}
		alertsV1AlertFiltersMetadataFiltersModel["ip_addresses"] = []string{"testString"}

		alertsV1AlertFiltersRatioAlertModel := make(map[string]interface{})
		alertsV1AlertFiltersRatioAlertModel["alias"] = "testString"
		alertsV1AlertFiltersRatioAlertModel["text"] = "testString"
		alertsV1AlertFiltersRatioAlertModel["severities"] = []string{"debug_or_unspecified"}
		alertsV1AlertFiltersRatioAlertModel["applications"] = []string{"testString"}
		alertsV1AlertFiltersRatioAlertModel["subsystems"] = []string{"testString"}
		alertsV1AlertFiltersRatioAlertModel["group_by"] = []string{"testString"}

		alertsV1AlertFiltersModel := make(map[string]interface{})
		alertsV1AlertFiltersModel["severities"] = []string{"info"}
		alertsV1AlertFiltersModel["metadata"] = []map[string]interface{}{alertsV1AlertFiltersMetadataFiltersModel}
		alertsV1AlertFiltersModel["alias"] = "testString"
		alertsV1AlertFiltersModel["text"] = "initiator.id.keyword:iam-ServiceId-10820fd6-c3fe-414e-8fd5-44ce95f7d34d AND action.keyword:cloud-object-storage.object.create"
		alertsV1AlertFiltersModel["ratio_alerts"] = []map[string]interface{}{alertsV1AlertFiltersRatioAlertModel}
		alertsV1AlertFiltersModel["filter_type"] = "text_or_unspecified"

		alertsV1TimeModel := make(map[string]interface{})
		alertsV1TimeModel["hours"] = int(18)
		alertsV1TimeModel["minutes"] = int(30)
		alertsV1TimeModel["seconds"] = int(38)

		alertsV1TimeRangeModel := make(map[string]interface{})
		alertsV1TimeRangeModel["start"] = []map[string]interface{}{alertsV1TimeModel}
		alertsV1TimeRangeModel["end"] = []map[string]interface{}{alertsV1TimeModel}

		alertsV1AlertActiveTimeframeModel := make(map[string]interface{})
		alertsV1AlertActiveTimeframeModel["days_of_week"] = []string{"sunday", "monday_or_unspecified", "tuesday", "wednesday", "thursday", "friday", "saturday"}
		alertsV1AlertActiveTimeframeModel["range"] = []map[string]interface{}{alertsV1TimeRangeModel}

		alertsV1AlertActiveWhenModel := make(map[string]interface{})
		alertsV1AlertActiveWhenModel["timeframes"] = []map[string]interface{}{alertsV1AlertActiveTimeframeModel}

		alertsV1MetaLabelModel := make(map[string]interface{})
		alertsV1MetaLabelModel["key"] = "env"
		alertsV1MetaLabelModel["value"] = "dev"

		alertsV1FiltersModel := make(map[string]interface{})
		alertsV1FiltersModel["values"] = []string{"testString"}
		alertsV1FiltersModel["operator"] = "testString"

		alertsV1FilterDataModel := make(map[string]interface{})
		alertsV1FilterDataModel["field"] = "testString"
		alertsV1FilterDataModel["filters"] = []map[string]interface{}{alertsV1FiltersModel}

		alertsV1TracingAlertModel := make(map[string]interface{})
		alertsV1TracingAlertModel["condition_latency"] = int(0)
		alertsV1TracingAlertModel["field_filters"] = []map[string]interface{}{alertsV1FilterDataModel}
		alertsV1TracingAlertModel["tag_filters"] = []map[string]interface{}{alertsV1FilterDataModel}

		alertsV2AlertIncidentSettingsModel := make(map[string]interface{})
		alertsV2AlertIncidentSettingsModel["retriggering_period_seconds"] = int(300)
		alertsV2AlertIncidentSettingsModel["notify_on"] = "triggered_only"
		alertsV2AlertIncidentSettingsModel["use_as_notification_settings"] = true

		model := make(map[string]interface{})
		model["id"] = "3dc02998-0b50-4ea8-b68a-4779d716fa1f"
		model["name"] = "Unique count alert"
		model["description"] = "Example of unique count alert from terraform"
		model["is_active"] = true
		model["severity"] = "info_or_unspecified"
		model["expiration"] = []map[string]interface{}{alertsV1DateModel}
		model["condition"] = []map[string]interface{}{alertsV2AlertConditionModel}
		model["notification_groups"] = []map[string]interface{}{alertsV2AlertNotificationGroupsModel}
		model["filters"] = []map[string]interface{}{alertsV1AlertFiltersModel}
		model["active_when"] = []map[string]interface{}{alertsV1AlertActiveWhenModel}
		model["notification_payload_filters"] = []string{"testString"}
		model["meta_labels"] = []map[string]interface{}{alertsV1MetaLabelModel}
		model["meta_labels_strings"] = []string{"testString"}
		model["tracing_alert"] = []map[string]interface{}{alertsV1TracingAlertModel}
		model["unique_identifier"] = "3dc02998-0b50-4ea8-b68a-4779d716fa1f"
		model["incident_settings"] = []map[string]interface{}{alertsV2AlertIncidentSettingsModel}

		assert.Equal(t, result, model)
	}

	alertsV1DateModel := new(logsv0.AlertsV1Date)
	alertsV1DateModel.Year = core.Int64Ptr(int64(38))
	alertsV1DateModel.Month = core.Int64Ptr(int64(38))
	alertsV1DateModel.Day = core.Int64Ptr(int64(38))

	alertsV1MetricAlertConditionParametersModel := new(logsv0.AlertsV1MetricAlertConditionParameters)
	alertsV1MetricAlertConditionParametersModel.MetricField = core.StringPtr("testString")
	alertsV1MetricAlertConditionParametersModel.MetricSource = core.StringPtr("logs2metrics_or_unspecified")
	alertsV1MetricAlertConditionParametersModel.ArithmeticOperator = core.StringPtr("avg_or_unspecified")
	alertsV1MetricAlertConditionParametersModel.ArithmeticOperatorModifier = core.Int64Ptr(int64(0))
	alertsV1MetricAlertConditionParametersModel.SampleThresholdPercentage = core.Int64Ptr(int64(0))
	alertsV1MetricAlertConditionParametersModel.NonNullPercentage = core.Int64Ptr(int64(0))
	alertsV1MetricAlertConditionParametersModel.SwapNullValues = core.BoolPtr(true)

	alertsV1MetricAlertPromqlConditionParametersModel := new(logsv0.AlertsV1MetricAlertPromqlConditionParameters)
	alertsV1MetricAlertPromqlConditionParametersModel.PromqlText = core.StringPtr("testString")
	alertsV1MetricAlertPromqlConditionParametersModel.ArithmeticOperatorModifier = core.Int64Ptr(int64(0))
	alertsV1MetricAlertPromqlConditionParametersModel.SampleThresholdPercentage = core.Int64Ptr(int64(0))
	alertsV1MetricAlertPromqlConditionParametersModel.NonNullPercentage = core.Int64Ptr(int64(0))
	alertsV1MetricAlertPromqlConditionParametersModel.SwapNullValues = core.BoolPtr(true)

	alertsV1RelatedExtendedDataModel := new(logsv0.AlertsV1RelatedExtendedData)
	alertsV1RelatedExtendedDataModel.CleanupDeadmanDuration = core.StringPtr("cleanup_deadman_duration_never_or_unspecified")
	alertsV1RelatedExtendedDataModel.ShouldTriggerDeadman = core.BoolPtr(true)

	alertsV2ConditionParametersModel := new(logsv0.AlertsV2ConditionParameters)
	alertsV2ConditionParametersModel.Threshold = core.Float64Ptr(float64(1))
	alertsV2ConditionParametersModel.Timeframe = core.StringPtr("timeframe_10_min")
	alertsV2ConditionParametersModel.GroupBy = []string{"coralogix.metadata.applicationName"}
	alertsV2ConditionParametersModel.MetricAlertParameters = alertsV1MetricAlertConditionParametersModel
	alertsV2ConditionParametersModel.MetricAlertPromqlParameters = alertsV1MetricAlertPromqlConditionParametersModel
	alertsV2ConditionParametersModel.IgnoreInfinity = core.BoolPtr(true)
	alertsV2ConditionParametersModel.RelativeTimeframe = core.StringPtr("hour_or_unspecified")
	alertsV2ConditionParametersModel.CardinalityFields = []string{"testString"}
	alertsV2ConditionParametersModel.RelatedExtendedData = alertsV1RelatedExtendedDataModel

	alertsV2MoreThanConditionModel := new(logsv0.AlertsV2MoreThanCondition)
	alertsV2MoreThanConditionModel.Parameters = alertsV2ConditionParametersModel
	alertsV2MoreThanConditionModel.EvaluationWindow = core.StringPtr("rolling_or_unspecified")

	alertsV2AlertConditionModel := new(logsv0.AlertsV2AlertConditionConditionMoreThan)
	alertsV2AlertConditionModel.MoreThan = alertsV2MoreThanConditionModel

	alertsV2AlertNotificationModel := new(logsv0.AlertsV2AlertNotificationIntegrationTypeIntegrationID)
	alertsV2AlertNotificationModel.RetriggeringPeriodSeconds = core.Int64Ptr(int64(0))
	alertsV2AlertNotificationModel.NotifyOn = core.StringPtr("triggered_only")
	alertsV2AlertNotificationModel.IntegrationID = core.Int64Ptr(int64(0))

	alertsV2AlertNotificationGroupsModel := new(logsv0.AlertsV2AlertNotificationGroups)
	alertsV2AlertNotificationGroupsModel.GroupByFields = []string{"coralogix.metadata.applicationName"}
	alertsV2AlertNotificationGroupsModel.Notifications = []logsv0.AlertsV2AlertNotificationIntf{alertsV2AlertNotificationModel}

	alertsV1AlertFiltersMetadataFiltersModel := new(logsv0.AlertsV1AlertFiltersMetadataFilters)
	alertsV1AlertFiltersMetadataFiltersModel.Categories = []string{"testString"}
	alertsV1AlertFiltersMetadataFiltersModel.Applications = []string{"testString"}
	alertsV1AlertFiltersMetadataFiltersModel.Subsystems = []string{"testString"}
	alertsV1AlertFiltersMetadataFiltersModel.Computers = []string{"testString"}
	alertsV1AlertFiltersMetadataFiltersModel.Classes = []string{"testString"}
	alertsV1AlertFiltersMetadataFiltersModel.Methods = []string{"testString"}
	alertsV1AlertFiltersMetadataFiltersModel.IpAddresses = []string{"testString"}

	alertsV1AlertFiltersRatioAlertModel := new(logsv0.AlertsV1AlertFiltersRatioAlert)
	alertsV1AlertFiltersRatioAlertModel.Alias = core.StringPtr("testString")
	alertsV1AlertFiltersRatioAlertModel.Text = core.StringPtr("testString")
	alertsV1AlertFiltersRatioAlertModel.Severities = []string{"debug_or_unspecified"}
	alertsV1AlertFiltersRatioAlertModel.Applications = []string{"testString"}
	alertsV1AlertFiltersRatioAlertModel.Subsystems = []string{"testString"}
	alertsV1AlertFiltersRatioAlertModel.GroupBy = []string{"testString"}

	alertsV1AlertFiltersModel := new(logsv0.AlertsV1AlertFilters)
	alertsV1AlertFiltersModel.Severities = []string{"info"}
	alertsV1AlertFiltersModel.Metadata = alertsV1AlertFiltersMetadataFiltersModel
	alertsV1AlertFiltersModel.Alias = core.StringPtr("testString")
	alertsV1AlertFiltersModel.Text = core.StringPtr("initiator.id.keyword:iam-ServiceId-10820fd6-c3fe-414e-8fd5-44ce95f7d34d AND action.keyword:cloud-object-storage.object.create")
	alertsV1AlertFiltersModel.RatioAlerts = []logsv0.AlertsV1AlertFiltersRatioAlert{*alertsV1AlertFiltersRatioAlertModel}
	alertsV1AlertFiltersModel.FilterType = core.StringPtr("text_or_unspecified")

	alertsV1TimeModel := new(logsv0.AlertsV1Time)
	alertsV1TimeModel.Hours = core.Int64Ptr(int64(18))
	alertsV1TimeModel.Minutes = core.Int64Ptr(int64(30))
	alertsV1TimeModel.Seconds = core.Int64Ptr(int64(38))

	alertsV1TimeRangeModel := new(logsv0.AlertsV1TimeRange)
	alertsV1TimeRangeModel.Start = alertsV1TimeModel
	alertsV1TimeRangeModel.End = alertsV1TimeModel

	alertsV1AlertActiveTimeframeModel := new(logsv0.AlertsV1AlertActiveTimeframe)
	alertsV1AlertActiveTimeframeModel.DaysOfWeek = []string{"sunday", "monday_or_unspecified", "tuesday", "wednesday", "thursday", "friday", "saturday"}
	alertsV1AlertActiveTimeframeModel.Range = alertsV1TimeRangeModel

	alertsV1AlertActiveWhenModel := new(logsv0.AlertsV1AlertActiveWhen)
	alertsV1AlertActiveWhenModel.Timeframes = []logsv0.AlertsV1AlertActiveTimeframe{*alertsV1AlertActiveTimeframeModel}

	alertsV1MetaLabelModel := new(logsv0.AlertsV1MetaLabel)
	alertsV1MetaLabelModel.Key = core.StringPtr("env")
	alertsV1MetaLabelModel.Value = core.StringPtr("dev")

	alertsV1FiltersModel := new(logsv0.AlertsV1Filters)
	alertsV1FiltersModel.Values = []string{"testString"}
	alertsV1FiltersModel.Operator = core.StringPtr("testString")

	alertsV1FilterDataModel := new(logsv0.AlertsV1FilterData)
	alertsV1FilterDataModel.Field = core.StringPtr("testString")
	alertsV1FilterDataModel.Filters = []logsv0.AlertsV1Filters{*alertsV1FiltersModel}

	alertsV1TracingAlertModel := new(logsv0.AlertsV1TracingAlert)
	alertsV1TracingAlertModel.ConditionLatency = core.Int64Ptr(int64(0))
	alertsV1TracingAlertModel.FieldFilters = []logsv0.AlertsV1FilterData{*alertsV1FilterDataModel}
	alertsV1TracingAlertModel.TagFilters = []logsv0.AlertsV1FilterData{*alertsV1FilterDataModel}

	alertsV2AlertIncidentSettingsModel := new(logsv0.AlertsV2AlertIncidentSettings)
	alertsV2AlertIncidentSettingsModel.RetriggeringPeriodSeconds = core.Int64Ptr(int64(300))
	alertsV2AlertIncidentSettingsModel.NotifyOn = core.StringPtr("triggered_only")
	alertsV2AlertIncidentSettingsModel.UseAsNotificationSettings = core.BoolPtr(true)

	model := new(logsv0.Alert)
	model.ID = CreateMockUUID("3dc02998-0b50-4ea8-b68a-4779d716fa1f")
	model.Name = core.StringPtr("Unique count alert")
	model.Description = core.StringPtr("Example of unique count alert from terraform")
	model.IsActive = core.BoolPtr(true)
	model.Severity = core.StringPtr("info_or_unspecified")
	model.Expiration = alertsV1DateModel
	model.Condition = alertsV2AlertConditionModel
	model.NotificationGroups = []logsv0.AlertsV2AlertNotificationGroups{*alertsV2AlertNotificationGroupsModel}
	model.Filters = alertsV1AlertFiltersModel
	model.ActiveWhen = alertsV1AlertActiveWhenModel
	model.NotificationPayloadFilters = []string{"testString"}
	model.MetaLabels = []logsv0.AlertsV1MetaLabel{*alertsV1MetaLabelModel}
	model.MetaLabelsStrings = []string{"testString"}
	model.TracingAlert = alertsV1TracingAlertModel
	model.UniqueIdentifier = CreateMockUUID("3dc02998-0b50-4ea8-b68a-4779d716fa1f")
	model.IncidentSettings = alertsV2AlertIncidentSettingsModel

	result, err := logs.DataSourceIbmLogsAlertsAlertToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmLogsAlertsAlertsV1DateToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["year"] = int(38)
		model["month"] = int(38)
		model["day"] = int(38)

		assert.Equal(t, result, model)
	}

	model := new(logsv0.AlertsV1Date)
	model.Year = core.Int64Ptr(int64(38))
	model.Month = core.Int64Ptr(int64(38))
	model.Day = core.Int64Ptr(int64(38))

	result, err := logs.DataSourceIbmLogsAlertsAlertsV1DateToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

// Todo @kavya498: Fix unit testcases
// func TestDataSourceIbmLogsAlertsAlertsV2AlertConditionToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		alertsV2ImmediateConditionModel := make(map[string]interface{})

// 		model := make(map[string]interface{})
// 		model["immediate"] = []map[string]interface{}{alertsV2ImmediateConditionModel}
// 		model["less_than"] = []map[string]interface{}{alertsV2LessThanConditionModel}
// 		model["more_than"] = []map[string]interface{}{alertsV2MoreThanConditionModel}
// 		model["more_than_usual"] = []map[string]interface{}{alertsV2MoreThanUsualConditionModel}
// 		model["new_value"] = []map[string]interface{}{alertsV2NewValueConditionModel}
// 		model["flow"] = []map[string]interface{}{alertsV2FlowConditionModel}
// 		model["unique_count"] = []map[string]interface{}{alertsV2UniqueCountConditionModel}
// 		model["less_than_usual"] = []map[string]interface{}{alertsV2LessThanUsualConditionModel}

// 		assert.Equal(t, result, model)
// 	}

// 	alertsV2ImmediateConditionModel := new(logsv0.AlertsV2ImmediateCondition)

// 	model := new(logsv0.AlertsV2AlertCondition)
// 	model.Immediate = alertsV2ImmediateConditionModel
// 	model.LessThan = alertsV2LessThanConditionModel
// 	model.MoreThan = alertsV2MoreThanConditionModel
// 	model.MoreThanUsual = alertsV2MoreThanUsualConditionModel
// 	model.NewValue = alertsV2NewValueConditionModel
// 	model.Flow = alertsV2FlowConditionModel
// 	model.UniqueCount = alertsV2UniqueCountConditionModel
// 	model.LessThanUsual = alertsV2LessThanUsualConditionModel

// 	result, err := logs.DataSourceIbmLogsAlertsAlertsV2AlertConditionToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

func TestDataSourceIbmLogsAlertsAlertsV2ImmediateConditionToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})

		assert.Equal(t, result, model)
	}

	model := new(logsv0.AlertsV2ImmediateCondition)

	result, err := logs.DataSourceIbmLogsAlertsAlertsV2ImmediateConditionToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmLogsAlertsAlertsV2LessThanConditionToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		alertsV1MetricAlertConditionParametersModel := make(map[string]interface{})
		alertsV1MetricAlertConditionParametersModel["metric_field"] = "testString"
		alertsV1MetricAlertConditionParametersModel["metric_source"] = "logs2metrics_or_unspecified"
		alertsV1MetricAlertConditionParametersModel["arithmetic_operator"] = "avg_or_unspecified"
		alertsV1MetricAlertConditionParametersModel["arithmetic_operator_modifier"] = int(0)
		alertsV1MetricAlertConditionParametersModel["sample_threshold_percentage"] = int(0)
		alertsV1MetricAlertConditionParametersModel["non_null_percentage"] = int(0)
		alertsV1MetricAlertConditionParametersModel["swap_null_values"] = true

		alertsV1MetricAlertPromqlConditionParametersModel := make(map[string]interface{})
		alertsV1MetricAlertPromqlConditionParametersModel["promql_text"] = "testString"
		alertsV1MetricAlertPromqlConditionParametersModel["arithmetic_operator_modifier"] = int(0)
		alertsV1MetricAlertPromqlConditionParametersModel["sample_threshold_percentage"] = int(0)
		alertsV1MetricAlertPromqlConditionParametersModel["non_null_percentage"] = int(0)
		alertsV1MetricAlertPromqlConditionParametersModel["swap_null_values"] = true

		alertsV1RelatedExtendedDataModel := make(map[string]interface{})
		alertsV1RelatedExtendedDataModel["cleanup_deadman_duration"] = "cleanup_deadman_duration_never_or_unspecified"
		alertsV1RelatedExtendedDataModel["should_trigger_deadman"] = true

		alertsV2ConditionParametersModel := make(map[string]interface{})
		alertsV2ConditionParametersModel["threshold"] = float64(72.5)
		alertsV2ConditionParametersModel["timeframe"] = "timeframe_5_min_or_unspecified"
		alertsV2ConditionParametersModel["group_by"] = []string{"testString"}
		alertsV2ConditionParametersModel["metric_alert_parameters"] = []map[string]interface{}{alertsV1MetricAlertConditionParametersModel}
		alertsV2ConditionParametersModel["metric_alert_promql_parameters"] = []map[string]interface{}{alertsV1MetricAlertPromqlConditionParametersModel}
		alertsV2ConditionParametersModel["ignore_infinity"] = true
		alertsV2ConditionParametersModel["relative_timeframe"] = "hour_or_unspecified"
		alertsV2ConditionParametersModel["cardinality_fields"] = []string{"testString"}
		alertsV2ConditionParametersModel["related_extended_data"] = []map[string]interface{}{alertsV1RelatedExtendedDataModel}

		model := make(map[string]interface{})
		model["parameters"] = []map[string]interface{}{alertsV2ConditionParametersModel}

		assert.Equal(t, result, model)
	}

	alertsV1MetricAlertConditionParametersModel := new(logsv0.AlertsV1MetricAlertConditionParameters)
	alertsV1MetricAlertConditionParametersModel.MetricField = core.StringPtr("testString")
	alertsV1MetricAlertConditionParametersModel.MetricSource = core.StringPtr("logs2metrics_or_unspecified")
	alertsV1MetricAlertConditionParametersModel.ArithmeticOperator = core.StringPtr("avg_or_unspecified")
	alertsV1MetricAlertConditionParametersModel.ArithmeticOperatorModifier = core.Int64Ptr(int64(0))
	alertsV1MetricAlertConditionParametersModel.SampleThresholdPercentage = core.Int64Ptr(int64(0))
	alertsV1MetricAlertConditionParametersModel.NonNullPercentage = core.Int64Ptr(int64(0))
	alertsV1MetricAlertConditionParametersModel.SwapNullValues = core.BoolPtr(true)

	alertsV1MetricAlertPromqlConditionParametersModel := new(logsv0.AlertsV1MetricAlertPromqlConditionParameters)
	alertsV1MetricAlertPromqlConditionParametersModel.PromqlText = core.StringPtr("testString")
	alertsV1MetricAlertPromqlConditionParametersModel.ArithmeticOperatorModifier = core.Int64Ptr(int64(0))
	alertsV1MetricAlertPromqlConditionParametersModel.SampleThresholdPercentage = core.Int64Ptr(int64(0))
	alertsV1MetricAlertPromqlConditionParametersModel.NonNullPercentage = core.Int64Ptr(int64(0))
	alertsV1MetricAlertPromqlConditionParametersModel.SwapNullValues = core.BoolPtr(true)

	alertsV1RelatedExtendedDataModel := new(logsv0.AlertsV1RelatedExtendedData)
	alertsV1RelatedExtendedDataModel.CleanupDeadmanDuration = core.StringPtr("cleanup_deadman_duration_never_or_unspecified")
	alertsV1RelatedExtendedDataModel.ShouldTriggerDeadman = core.BoolPtr(true)

	alertsV2ConditionParametersModel := new(logsv0.AlertsV2ConditionParameters)
	alertsV2ConditionParametersModel.Threshold = core.Float64Ptr(float64(72.5))
	alertsV2ConditionParametersModel.Timeframe = core.StringPtr("timeframe_5_min_or_unspecified")
	alertsV2ConditionParametersModel.GroupBy = []string{"testString"}
	alertsV2ConditionParametersModel.MetricAlertParameters = alertsV1MetricAlertConditionParametersModel
	alertsV2ConditionParametersModel.MetricAlertPromqlParameters = alertsV1MetricAlertPromqlConditionParametersModel
	alertsV2ConditionParametersModel.IgnoreInfinity = core.BoolPtr(true)
	alertsV2ConditionParametersModel.RelativeTimeframe = core.StringPtr("hour_or_unspecified")
	alertsV2ConditionParametersModel.CardinalityFields = []string{"testString"}
	alertsV2ConditionParametersModel.RelatedExtendedData = alertsV1RelatedExtendedDataModel

	model := new(logsv0.AlertsV2LessThanCondition)
	model.Parameters = alertsV2ConditionParametersModel

	result, err := logs.DataSourceIbmLogsAlertsAlertsV2LessThanConditionToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmLogsAlertsAlertsV2ConditionParametersToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		alertsV1MetricAlertConditionParametersModel := make(map[string]interface{})
		alertsV1MetricAlertConditionParametersModel["metric_field"] = "testString"
		alertsV1MetricAlertConditionParametersModel["metric_source"] = "logs2metrics_or_unspecified"
		alertsV1MetricAlertConditionParametersModel["arithmetic_operator"] = "avg_or_unspecified"
		alertsV1MetricAlertConditionParametersModel["arithmetic_operator_modifier"] = int(0)
		alertsV1MetricAlertConditionParametersModel["sample_threshold_percentage"] = int(0)
		alertsV1MetricAlertConditionParametersModel["non_null_percentage"] = int(0)
		alertsV1MetricAlertConditionParametersModel["swap_null_values"] = true

		alertsV1MetricAlertPromqlConditionParametersModel := make(map[string]interface{})
		alertsV1MetricAlertPromqlConditionParametersModel["promql_text"] = "testString"
		alertsV1MetricAlertPromqlConditionParametersModel["arithmetic_operator_modifier"] = int(0)
		alertsV1MetricAlertPromqlConditionParametersModel["sample_threshold_percentage"] = int(0)
		alertsV1MetricAlertPromqlConditionParametersModel["non_null_percentage"] = int(0)
		alertsV1MetricAlertPromqlConditionParametersModel["swap_null_values"] = true

		alertsV1RelatedExtendedDataModel := make(map[string]interface{})
		alertsV1RelatedExtendedDataModel["cleanup_deadman_duration"] = "cleanup_deadman_duration_never_or_unspecified"
		alertsV1RelatedExtendedDataModel["should_trigger_deadman"] = true

		model := make(map[string]interface{})
		model["threshold"] = float64(72.5)
		model["timeframe"] = "timeframe_5_min_or_unspecified"
		model["group_by"] = []string{"testString"}
		model["metric_alert_parameters"] = []map[string]interface{}{alertsV1MetricAlertConditionParametersModel}
		model["metric_alert_promql_parameters"] = []map[string]interface{}{alertsV1MetricAlertPromqlConditionParametersModel}
		model["ignore_infinity"] = true
		model["relative_timeframe"] = "hour_or_unspecified"
		model["cardinality_fields"] = []string{"testString"}
		model["related_extended_data"] = []map[string]interface{}{alertsV1RelatedExtendedDataModel}

		assert.Equal(t, result, model)
	}

	alertsV1MetricAlertConditionParametersModel := new(logsv0.AlertsV1MetricAlertConditionParameters)
	alertsV1MetricAlertConditionParametersModel.MetricField = core.StringPtr("testString")
	alertsV1MetricAlertConditionParametersModel.MetricSource = core.StringPtr("logs2metrics_or_unspecified")
	alertsV1MetricAlertConditionParametersModel.ArithmeticOperator = core.StringPtr("avg_or_unspecified")
	alertsV1MetricAlertConditionParametersModel.ArithmeticOperatorModifier = core.Int64Ptr(int64(0))
	alertsV1MetricAlertConditionParametersModel.SampleThresholdPercentage = core.Int64Ptr(int64(0))
	alertsV1MetricAlertConditionParametersModel.NonNullPercentage = core.Int64Ptr(int64(0))
	alertsV1MetricAlertConditionParametersModel.SwapNullValues = core.BoolPtr(true)

	alertsV1MetricAlertPromqlConditionParametersModel := new(logsv0.AlertsV1MetricAlertPromqlConditionParameters)
	alertsV1MetricAlertPromqlConditionParametersModel.PromqlText = core.StringPtr("testString")
	alertsV1MetricAlertPromqlConditionParametersModel.ArithmeticOperatorModifier = core.Int64Ptr(int64(0))
	alertsV1MetricAlertPromqlConditionParametersModel.SampleThresholdPercentage = core.Int64Ptr(int64(0))
	alertsV1MetricAlertPromqlConditionParametersModel.NonNullPercentage = core.Int64Ptr(int64(0))
	alertsV1MetricAlertPromqlConditionParametersModel.SwapNullValues = core.BoolPtr(true)

	alertsV1RelatedExtendedDataModel := new(logsv0.AlertsV1RelatedExtendedData)
	alertsV1RelatedExtendedDataModel.CleanupDeadmanDuration = core.StringPtr("cleanup_deadman_duration_never_or_unspecified")
	alertsV1RelatedExtendedDataModel.ShouldTriggerDeadman = core.BoolPtr(true)

	model := new(logsv0.AlertsV2ConditionParameters)
	model.Threshold = core.Float64Ptr(float64(72.5))
	model.Timeframe = core.StringPtr("timeframe_5_min_or_unspecified")
	model.GroupBy = []string{"testString"}
	model.MetricAlertParameters = alertsV1MetricAlertConditionParametersModel
	model.MetricAlertPromqlParameters = alertsV1MetricAlertPromqlConditionParametersModel
	model.IgnoreInfinity = core.BoolPtr(true)
	model.RelativeTimeframe = core.StringPtr("hour_or_unspecified")
	model.CardinalityFields = []string{"testString"}
	model.RelatedExtendedData = alertsV1RelatedExtendedDataModel

	result, err := logs.DataSourceIbmLogsAlertsAlertsV2ConditionParametersToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmLogsAlertsAlertsV1MetricAlertConditionParametersToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["metric_field"] = "testString"
		model["metric_source"] = "logs2metrics_or_unspecified"
		model["arithmetic_operator"] = "avg_or_unspecified"
		model["arithmetic_operator_modifier"] = int(0)
		model["sample_threshold_percentage"] = int(0)
		model["non_null_percentage"] = int(0)
		model["swap_null_values"] = true

		assert.Equal(t, result, model)
	}

	model := new(logsv0.AlertsV1MetricAlertConditionParameters)
	model.MetricField = core.StringPtr("testString")
	model.MetricSource = core.StringPtr("logs2metrics_or_unspecified")
	model.ArithmeticOperator = core.StringPtr("avg_or_unspecified")
	model.ArithmeticOperatorModifier = core.Int64Ptr(int64(0))
	model.SampleThresholdPercentage = core.Int64Ptr(int64(0))
	model.NonNullPercentage = core.Int64Ptr(int64(0))
	model.SwapNullValues = core.BoolPtr(true)

	result, err := logs.DataSourceIbmLogsAlertsAlertsV1MetricAlertConditionParametersToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmLogsAlertsAlertsV1MetricAlertPromqlConditionParametersToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["promql_text"] = "testString"
		model["arithmetic_operator_modifier"] = int(0)
		model["sample_threshold_percentage"] = int(0)
		model["non_null_percentage"] = int(0)
		model["swap_null_values"] = true

		assert.Equal(t, result, model)
	}

	model := new(logsv0.AlertsV1MetricAlertPromqlConditionParameters)
	model.PromqlText = core.StringPtr("testString")
	model.ArithmeticOperatorModifier = core.Int64Ptr(int64(0))
	model.SampleThresholdPercentage = core.Int64Ptr(int64(0))
	model.NonNullPercentage = core.Int64Ptr(int64(0))
	model.SwapNullValues = core.BoolPtr(true)

	result, err := logs.DataSourceIbmLogsAlertsAlertsV1MetricAlertPromqlConditionParametersToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmLogsAlertsAlertsV1RelatedExtendedDataToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["cleanup_deadman_duration"] = "cleanup_deadman_duration_never_or_unspecified"
		model["should_trigger_deadman"] = true

		assert.Equal(t, result, model)
	}

	model := new(logsv0.AlertsV1RelatedExtendedData)
	model.CleanupDeadmanDuration = core.StringPtr("cleanup_deadman_duration_never_or_unspecified")
	model.ShouldTriggerDeadman = core.BoolPtr(true)

	result, err := logs.DataSourceIbmLogsAlertsAlertsV1RelatedExtendedDataToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmLogsAlertsAlertsV2MoreThanConditionToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		alertsV1MetricAlertConditionParametersModel := make(map[string]interface{})
		alertsV1MetricAlertConditionParametersModel["metric_field"] = "testString"
		alertsV1MetricAlertConditionParametersModel["metric_source"] = "logs2metrics_or_unspecified"
		alertsV1MetricAlertConditionParametersModel["arithmetic_operator"] = "avg_or_unspecified"
		alertsV1MetricAlertConditionParametersModel["arithmetic_operator_modifier"] = int(0)
		alertsV1MetricAlertConditionParametersModel["sample_threshold_percentage"] = int(0)
		alertsV1MetricAlertConditionParametersModel["non_null_percentage"] = int(0)
		alertsV1MetricAlertConditionParametersModel["swap_null_values"] = true

		alertsV1MetricAlertPromqlConditionParametersModel := make(map[string]interface{})
		alertsV1MetricAlertPromqlConditionParametersModel["promql_text"] = "testString"
		alertsV1MetricAlertPromqlConditionParametersModel["arithmetic_operator_modifier"] = int(0)
		alertsV1MetricAlertPromqlConditionParametersModel["sample_threshold_percentage"] = int(0)
		alertsV1MetricAlertPromqlConditionParametersModel["non_null_percentage"] = int(0)
		alertsV1MetricAlertPromqlConditionParametersModel["swap_null_values"] = true

		alertsV1RelatedExtendedDataModel := make(map[string]interface{})
		alertsV1RelatedExtendedDataModel["cleanup_deadman_duration"] = "cleanup_deadman_duration_never_or_unspecified"
		alertsV1RelatedExtendedDataModel["should_trigger_deadman"] = true

		alertsV2ConditionParametersModel := make(map[string]interface{})
		alertsV2ConditionParametersModel["threshold"] = float64(72.5)
		alertsV2ConditionParametersModel["timeframe"] = "timeframe_5_min_or_unspecified"
		alertsV2ConditionParametersModel["group_by"] = []string{"testString"}
		alertsV2ConditionParametersModel["metric_alert_parameters"] = []map[string]interface{}{alertsV1MetricAlertConditionParametersModel}
		alertsV2ConditionParametersModel["metric_alert_promql_parameters"] = []map[string]interface{}{alertsV1MetricAlertPromqlConditionParametersModel}
		alertsV2ConditionParametersModel["ignore_infinity"] = true
		alertsV2ConditionParametersModel["relative_timeframe"] = "hour_or_unspecified"
		alertsV2ConditionParametersModel["cardinality_fields"] = []string{"testString"}
		alertsV2ConditionParametersModel["related_extended_data"] = []map[string]interface{}{alertsV1RelatedExtendedDataModel}

		model := make(map[string]interface{})
		model["parameters"] = []map[string]interface{}{alertsV2ConditionParametersModel}
		model["evaluation_window"] = "rolling_or_unspecified"

		assert.Equal(t, result, model)
	}

	alertsV1MetricAlertConditionParametersModel := new(logsv0.AlertsV1MetricAlertConditionParameters)
	alertsV1MetricAlertConditionParametersModel.MetricField = core.StringPtr("testString")
	alertsV1MetricAlertConditionParametersModel.MetricSource = core.StringPtr("logs2metrics_or_unspecified")
	alertsV1MetricAlertConditionParametersModel.ArithmeticOperator = core.StringPtr("avg_or_unspecified")
	alertsV1MetricAlertConditionParametersModel.ArithmeticOperatorModifier = core.Int64Ptr(int64(0))
	alertsV1MetricAlertConditionParametersModel.SampleThresholdPercentage = core.Int64Ptr(int64(0))
	alertsV1MetricAlertConditionParametersModel.NonNullPercentage = core.Int64Ptr(int64(0))
	alertsV1MetricAlertConditionParametersModel.SwapNullValues = core.BoolPtr(true)

	alertsV1MetricAlertPromqlConditionParametersModel := new(logsv0.AlertsV1MetricAlertPromqlConditionParameters)
	alertsV1MetricAlertPromqlConditionParametersModel.PromqlText = core.StringPtr("testString")
	alertsV1MetricAlertPromqlConditionParametersModel.ArithmeticOperatorModifier = core.Int64Ptr(int64(0))
	alertsV1MetricAlertPromqlConditionParametersModel.SampleThresholdPercentage = core.Int64Ptr(int64(0))
	alertsV1MetricAlertPromqlConditionParametersModel.NonNullPercentage = core.Int64Ptr(int64(0))
	alertsV1MetricAlertPromqlConditionParametersModel.SwapNullValues = core.BoolPtr(true)

	alertsV1RelatedExtendedDataModel := new(logsv0.AlertsV1RelatedExtendedData)
	alertsV1RelatedExtendedDataModel.CleanupDeadmanDuration = core.StringPtr("cleanup_deadman_duration_never_or_unspecified")
	alertsV1RelatedExtendedDataModel.ShouldTriggerDeadman = core.BoolPtr(true)

	alertsV2ConditionParametersModel := new(logsv0.AlertsV2ConditionParameters)
	alertsV2ConditionParametersModel.Threshold = core.Float64Ptr(float64(72.5))
	alertsV2ConditionParametersModel.Timeframe = core.StringPtr("timeframe_5_min_or_unspecified")
	alertsV2ConditionParametersModel.GroupBy = []string{"testString"}
	alertsV2ConditionParametersModel.MetricAlertParameters = alertsV1MetricAlertConditionParametersModel
	alertsV2ConditionParametersModel.MetricAlertPromqlParameters = alertsV1MetricAlertPromqlConditionParametersModel
	alertsV2ConditionParametersModel.IgnoreInfinity = core.BoolPtr(true)
	alertsV2ConditionParametersModel.RelativeTimeframe = core.StringPtr("hour_or_unspecified")
	alertsV2ConditionParametersModel.CardinalityFields = []string{"testString"}
	alertsV2ConditionParametersModel.RelatedExtendedData = alertsV1RelatedExtendedDataModel

	model := new(logsv0.AlertsV2MoreThanCondition)
	model.Parameters = alertsV2ConditionParametersModel
	model.EvaluationWindow = core.StringPtr("rolling_or_unspecified")

	result, err := logs.DataSourceIbmLogsAlertsAlertsV2MoreThanConditionToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmLogsAlertsAlertsV2MoreThanUsualConditionToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		alertsV1MetricAlertConditionParametersModel := make(map[string]interface{})
		alertsV1MetricAlertConditionParametersModel["metric_field"] = "testString"
		alertsV1MetricAlertConditionParametersModel["metric_source"] = "logs2metrics_or_unspecified"
		alertsV1MetricAlertConditionParametersModel["arithmetic_operator"] = "avg_or_unspecified"
		alertsV1MetricAlertConditionParametersModel["arithmetic_operator_modifier"] = int(0)
		alertsV1MetricAlertConditionParametersModel["sample_threshold_percentage"] = int(0)
		alertsV1MetricAlertConditionParametersModel["non_null_percentage"] = int(0)
		alertsV1MetricAlertConditionParametersModel["swap_null_values"] = true

		alertsV1MetricAlertPromqlConditionParametersModel := make(map[string]interface{})
		alertsV1MetricAlertPromqlConditionParametersModel["promql_text"] = "testString"
		alertsV1MetricAlertPromqlConditionParametersModel["arithmetic_operator_modifier"] = int(0)
		alertsV1MetricAlertPromqlConditionParametersModel["sample_threshold_percentage"] = int(0)
		alertsV1MetricAlertPromqlConditionParametersModel["non_null_percentage"] = int(0)
		alertsV1MetricAlertPromqlConditionParametersModel["swap_null_values"] = true

		alertsV1RelatedExtendedDataModel := make(map[string]interface{})
		alertsV1RelatedExtendedDataModel["cleanup_deadman_duration"] = "cleanup_deadman_duration_never_or_unspecified"
		alertsV1RelatedExtendedDataModel["should_trigger_deadman"] = true

		alertsV2ConditionParametersModel := make(map[string]interface{})
		alertsV2ConditionParametersModel["threshold"] = float64(72.5)
		alertsV2ConditionParametersModel["timeframe"] = "timeframe_5_min_or_unspecified"
		alertsV2ConditionParametersModel["group_by"] = []string{"testString"}
		alertsV2ConditionParametersModel["metric_alert_parameters"] = []map[string]interface{}{alertsV1MetricAlertConditionParametersModel}
		alertsV2ConditionParametersModel["metric_alert_promql_parameters"] = []map[string]interface{}{alertsV1MetricAlertPromqlConditionParametersModel}
		alertsV2ConditionParametersModel["ignore_infinity"] = true
		alertsV2ConditionParametersModel["relative_timeframe"] = "hour_or_unspecified"
		alertsV2ConditionParametersModel["cardinality_fields"] = []string{"testString"}
		alertsV2ConditionParametersModel["related_extended_data"] = []map[string]interface{}{alertsV1RelatedExtendedDataModel}

		model := make(map[string]interface{})
		model["parameters"] = []map[string]interface{}{alertsV2ConditionParametersModel}

		assert.Equal(t, result, model)
	}

	alertsV1MetricAlertConditionParametersModel := new(logsv0.AlertsV1MetricAlertConditionParameters)
	alertsV1MetricAlertConditionParametersModel.MetricField = core.StringPtr("testString")
	alertsV1MetricAlertConditionParametersModel.MetricSource = core.StringPtr("logs2metrics_or_unspecified")
	alertsV1MetricAlertConditionParametersModel.ArithmeticOperator = core.StringPtr("avg_or_unspecified")
	alertsV1MetricAlertConditionParametersModel.ArithmeticOperatorModifier = core.Int64Ptr(int64(0))
	alertsV1MetricAlertConditionParametersModel.SampleThresholdPercentage = core.Int64Ptr(int64(0))
	alertsV1MetricAlertConditionParametersModel.NonNullPercentage = core.Int64Ptr(int64(0))
	alertsV1MetricAlertConditionParametersModel.SwapNullValues = core.BoolPtr(true)

	alertsV1MetricAlertPromqlConditionParametersModel := new(logsv0.AlertsV1MetricAlertPromqlConditionParameters)
	alertsV1MetricAlertPromqlConditionParametersModel.PromqlText = core.StringPtr("testString")
	alertsV1MetricAlertPromqlConditionParametersModel.ArithmeticOperatorModifier = core.Int64Ptr(int64(0))
	alertsV1MetricAlertPromqlConditionParametersModel.SampleThresholdPercentage = core.Int64Ptr(int64(0))
	alertsV1MetricAlertPromqlConditionParametersModel.NonNullPercentage = core.Int64Ptr(int64(0))
	alertsV1MetricAlertPromqlConditionParametersModel.SwapNullValues = core.BoolPtr(true)

	alertsV1RelatedExtendedDataModel := new(logsv0.AlertsV1RelatedExtendedData)
	alertsV1RelatedExtendedDataModel.CleanupDeadmanDuration = core.StringPtr("cleanup_deadman_duration_never_or_unspecified")
	alertsV1RelatedExtendedDataModel.ShouldTriggerDeadman = core.BoolPtr(true)

	alertsV2ConditionParametersModel := new(logsv0.AlertsV2ConditionParameters)
	alertsV2ConditionParametersModel.Threshold = core.Float64Ptr(float64(72.5))
	alertsV2ConditionParametersModel.Timeframe = core.StringPtr("timeframe_5_min_or_unspecified")
	alertsV2ConditionParametersModel.GroupBy = []string{"testString"}
	alertsV2ConditionParametersModel.MetricAlertParameters = alertsV1MetricAlertConditionParametersModel
	alertsV2ConditionParametersModel.MetricAlertPromqlParameters = alertsV1MetricAlertPromqlConditionParametersModel
	alertsV2ConditionParametersModel.IgnoreInfinity = core.BoolPtr(true)
	alertsV2ConditionParametersModel.RelativeTimeframe = core.StringPtr("hour_or_unspecified")
	alertsV2ConditionParametersModel.CardinalityFields = []string{"testString"}
	alertsV2ConditionParametersModel.RelatedExtendedData = alertsV1RelatedExtendedDataModel

	model := new(logsv0.AlertsV2MoreThanUsualCondition)
	model.Parameters = alertsV2ConditionParametersModel

	result, err := logs.DataSourceIbmLogsAlertsAlertsV2MoreThanUsualConditionToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmLogsAlertsAlertsV2NewValueConditionToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		alertsV1MetricAlertConditionParametersModel := make(map[string]interface{})
		alertsV1MetricAlertConditionParametersModel["metric_field"] = "testString"
		alertsV1MetricAlertConditionParametersModel["metric_source"] = "logs2metrics_or_unspecified"
		alertsV1MetricAlertConditionParametersModel["arithmetic_operator"] = "avg_or_unspecified"
		alertsV1MetricAlertConditionParametersModel["arithmetic_operator_modifier"] = int(0)
		alertsV1MetricAlertConditionParametersModel["sample_threshold_percentage"] = int(0)
		alertsV1MetricAlertConditionParametersModel["non_null_percentage"] = int(0)
		alertsV1MetricAlertConditionParametersModel["swap_null_values"] = true

		alertsV1MetricAlertPromqlConditionParametersModel := make(map[string]interface{})
		alertsV1MetricAlertPromqlConditionParametersModel["promql_text"] = "testString"
		alertsV1MetricAlertPromqlConditionParametersModel["arithmetic_operator_modifier"] = int(0)
		alertsV1MetricAlertPromqlConditionParametersModel["sample_threshold_percentage"] = int(0)
		alertsV1MetricAlertPromqlConditionParametersModel["non_null_percentage"] = int(0)
		alertsV1MetricAlertPromqlConditionParametersModel["swap_null_values"] = true

		alertsV1RelatedExtendedDataModel := make(map[string]interface{})
		alertsV1RelatedExtendedDataModel["cleanup_deadman_duration"] = "cleanup_deadman_duration_never_or_unspecified"
		alertsV1RelatedExtendedDataModel["should_trigger_deadman"] = true

		alertsV2ConditionParametersModel := make(map[string]interface{})
		alertsV2ConditionParametersModel["threshold"] = float64(72.5)
		alertsV2ConditionParametersModel["timeframe"] = "timeframe_5_min_or_unspecified"
		alertsV2ConditionParametersModel["group_by"] = []string{"testString"}
		alertsV2ConditionParametersModel["metric_alert_parameters"] = []map[string]interface{}{alertsV1MetricAlertConditionParametersModel}
		alertsV2ConditionParametersModel["metric_alert_promql_parameters"] = []map[string]interface{}{alertsV1MetricAlertPromqlConditionParametersModel}
		alertsV2ConditionParametersModel["ignore_infinity"] = true
		alertsV2ConditionParametersModel["relative_timeframe"] = "hour_or_unspecified"
		alertsV2ConditionParametersModel["cardinality_fields"] = []string{"testString"}
		alertsV2ConditionParametersModel["related_extended_data"] = []map[string]interface{}{alertsV1RelatedExtendedDataModel}

		model := make(map[string]interface{})
		model["parameters"] = []map[string]interface{}{alertsV2ConditionParametersModel}

		assert.Equal(t, result, model)
	}

	alertsV1MetricAlertConditionParametersModel := new(logsv0.AlertsV1MetricAlertConditionParameters)
	alertsV1MetricAlertConditionParametersModel.MetricField = core.StringPtr("testString")
	alertsV1MetricAlertConditionParametersModel.MetricSource = core.StringPtr("logs2metrics_or_unspecified")
	alertsV1MetricAlertConditionParametersModel.ArithmeticOperator = core.StringPtr("avg_or_unspecified")
	alertsV1MetricAlertConditionParametersModel.ArithmeticOperatorModifier = core.Int64Ptr(int64(0))
	alertsV1MetricAlertConditionParametersModel.SampleThresholdPercentage = core.Int64Ptr(int64(0))
	alertsV1MetricAlertConditionParametersModel.NonNullPercentage = core.Int64Ptr(int64(0))
	alertsV1MetricAlertConditionParametersModel.SwapNullValues = core.BoolPtr(true)

	alertsV1MetricAlertPromqlConditionParametersModel := new(logsv0.AlertsV1MetricAlertPromqlConditionParameters)
	alertsV1MetricAlertPromqlConditionParametersModel.PromqlText = core.StringPtr("testString")
	alertsV1MetricAlertPromqlConditionParametersModel.ArithmeticOperatorModifier = core.Int64Ptr(int64(0))
	alertsV1MetricAlertPromqlConditionParametersModel.SampleThresholdPercentage = core.Int64Ptr(int64(0))
	alertsV1MetricAlertPromqlConditionParametersModel.NonNullPercentage = core.Int64Ptr(int64(0))
	alertsV1MetricAlertPromqlConditionParametersModel.SwapNullValues = core.BoolPtr(true)

	alertsV1RelatedExtendedDataModel := new(logsv0.AlertsV1RelatedExtendedData)
	alertsV1RelatedExtendedDataModel.CleanupDeadmanDuration = core.StringPtr("cleanup_deadman_duration_never_or_unspecified")
	alertsV1RelatedExtendedDataModel.ShouldTriggerDeadman = core.BoolPtr(true)

	alertsV2ConditionParametersModel := new(logsv0.AlertsV2ConditionParameters)
	alertsV2ConditionParametersModel.Threshold = core.Float64Ptr(float64(72.5))
	alertsV2ConditionParametersModel.Timeframe = core.StringPtr("timeframe_5_min_or_unspecified")
	alertsV2ConditionParametersModel.GroupBy = []string{"testString"}
	alertsV2ConditionParametersModel.MetricAlertParameters = alertsV1MetricAlertConditionParametersModel
	alertsV2ConditionParametersModel.MetricAlertPromqlParameters = alertsV1MetricAlertPromqlConditionParametersModel
	alertsV2ConditionParametersModel.IgnoreInfinity = core.BoolPtr(true)
	alertsV2ConditionParametersModel.RelativeTimeframe = core.StringPtr("hour_or_unspecified")
	alertsV2ConditionParametersModel.CardinalityFields = []string{"testString"}
	alertsV2ConditionParametersModel.RelatedExtendedData = alertsV1RelatedExtendedDataModel

	model := new(logsv0.AlertsV2NewValueCondition)
	model.Parameters = alertsV2ConditionParametersModel

	result, err := logs.DataSourceIbmLogsAlertsAlertsV2NewValueConditionToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmLogsAlertsAlertsV2FlowConditionToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		alertsV1FlowAlertModel := make(map[string]interface{})
		alertsV1FlowAlertModel["id"] = "testString"
		alertsV1FlowAlertModel["not"] = true

		alertsV1FlowAlertsModel := make(map[string]interface{})
		alertsV1FlowAlertsModel["op"] = "and"
		alertsV1FlowAlertsModel["values"] = []map[string]interface{}{alertsV1FlowAlertModel}

		alertsV1FlowGroupModel := make(map[string]interface{})
		alertsV1FlowGroupModel["alerts"] = []map[string]interface{}{alertsV1FlowAlertsModel}
		alertsV1FlowGroupModel["next_op"] = "and"

		alertsV1FlowTimeframeModel := make(map[string]interface{})
		alertsV1FlowTimeframeModel["ms"] = int(0)

		alertsV1FlowStageModel := make(map[string]interface{})
		alertsV1FlowStageModel["groups"] = []map[string]interface{}{alertsV1FlowGroupModel}
		alertsV1FlowStageModel["timeframe"] = []map[string]interface{}{alertsV1FlowTimeframeModel}

		alertsV1MetricAlertConditionParametersModel := make(map[string]interface{})
		alertsV1MetricAlertConditionParametersModel["metric_field"] = "testString"
		alertsV1MetricAlertConditionParametersModel["metric_source"] = "logs2metrics_or_unspecified"
		alertsV1MetricAlertConditionParametersModel["arithmetic_operator"] = "avg_or_unspecified"
		alertsV1MetricAlertConditionParametersModel["arithmetic_operator_modifier"] = int(0)
		alertsV1MetricAlertConditionParametersModel["sample_threshold_percentage"] = int(0)
		alertsV1MetricAlertConditionParametersModel["non_null_percentage"] = int(0)
		alertsV1MetricAlertConditionParametersModel["swap_null_values"] = true

		alertsV1MetricAlertPromqlConditionParametersModel := make(map[string]interface{})
		alertsV1MetricAlertPromqlConditionParametersModel["promql_text"] = "testString"
		alertsV1MetricAlertPromqlConditionParametersModel["arithmetic_operator_modifier"] = int(0)
		alertsV1MetricAlertPromqlConditionParametersModel["sample_threshold_percentage"] = int(0)
		alertsV1MetricAlertPromqlConditionParametersModel["non_null_percentage"] = int(0)
		alertsV1MetricAlertPromqlConditionParametersModel["swap_null_values"] = true

		alertsV1RelatedExtendedDataModel := make(map[string]interface{})
		alertsV1RelatedExtendedDataModel["cleanup_deadman_duration"] = "cleanup_deadman_duration_never_or_unspecified"
		alertsV1RelatedExtendedDataModel["should_trigger_deadman"] = true

		alertsV2ConditionParametersModel := make(map[string]interface{})
		alertsV2ConditionParametersModel["threshold"] = float64(72.5)
		alertsV2ConditionParametersModel["timeframe"] = "timeframe_5_min_or_unspecified"
		alertsV2ConditionParametersModel["group_by"] = []string{"testString"}
		alertsV2ConditionParametersModel["metric_alert_parameters"] = []map[string]interface{}{alertsV1MetricAlertConditionParametersModel}
		alertsV2ConditionParametersModel["metric_alert_promql_parameters"] = []map[string]interface{}{alertsV1MetricAlertPromqlConditionParametersModel}
		alertsV2ConditionParametersModel["ignore_infinity"] = true
		alertsV2ConditionParametersModel["relative_timeframe"] = "hour_or_unspecified"
		alertsV2ConditionParametersModel["cardinality_fields"] = []string{"testString"}
		alertsV2ConditionParametersModel["related_extended_data"] = []map[string]interface{}{alertsV1RelatedExtendedDataModel}

		model := make(map[string]interface{})
		model["stages"] = []map[string]interface{}{alertsV1FlowStageModel}
		model["parameters"] = []map[string]interface{}{alertsV2ConditionParametersModel}
		model["enforce_suppression"] = true

		assert.Equal(t, result, model)
	}

	alertsV1FlowAlertModel := new(logsv0.AlertsV1FlowAlert)
	alertsV1FlowAlertModel.ID = core.StringPtr("testString")
	alertsV1FlowAlertModel.Not = core.BoolPtr(true)

	alertsV1FlowAlertsModel := new(logsv0.AlertsV1FlowAlerts)
	alertsV1FlowAlertsModel.Op = core.StringPtr("and")
	alertsV1FlowAlertsModel.Values = []logsv0.AlertsV1FlowAlert{*alertsV1FlowAlertModel}

	alertsV1FlowGroupModel := new(logsv0.AlertsV1FlowGroup)
	alertsV1FlowGroupModel.Alerts = alertsV1FlowAlertsModel
	alertsV1FlowGroupModel.NextOp = core.StringPtr("and")

	alertsV1FlowTimeframeModel := new(logsv0.AlertsV1FlowTimeframe)
	alertsV1FlowTimeframeModel.Ms = core.Int64Ptr(int64(0))

	alertsV1FlowStageModel := new(logsv0.AlertsV1FlowStage)
	alertsV1FlowStageModel.Groups = []logsv0.AlertsV1FlowGroup{*alertsV1FlowGroupModel}
	alertsV1FlowStageModel.Timeframe = alertsV1FlowTimeframeModel

	alertsV1MetricAlertConditionParametersModel := new(logsv0.AlertsV1MetricAlertConditionParameters)
	alertsV1MetricAlertConditionParametersModel.MetricField = core.StringPtr("testString")
	alertsV1MetricAlertConditionParametersModel.MetricSource = core.StringPtr("logs2metrics_or_unspecified")
	alertsV1MetricAlertConditionParametersModel.ArithmeticOperator = core.StringPtr("avg_or_unspecified")
	alertsV1MetricAlertConditionParametersModel.ArithmeticOperatorModifier = core.Int64Ptr(int64(0))
	alertsV1MetricAlertConditionParametersModel.SampleThresholdPercentage = core.Int64Ptr(int64(0))
	alertsV1MetricAlertConditionParametersModel.NonNullPercentage = core.Int64Ptr(int64(0))
	alertsV1MetricAlertConditionParametersModel.SwapNullValues = core.BoolPtr(true)

	alertsV1MetricAlertPromqlConditionParametersModel := new(logsv0.AlertsV1MetricAlertPromqlConditionParameters)
	alertsV1MetricAlertPromqlConditionParametersModel.PromqlText = core.StringPtr("testString")
	alertsV1MetricAlertPromqlConditionParametersModel.ArithmeticOperatorModifier = core.Int64Ptr(int64(0))
	alertsV1MetricAlertPromqlConditionParametersModel.SampleThresholdPercentage = core.Int64Ptr(int64(0))
	alertsV1MetricAlertPromqlConditionParametersModel.NonNullPercentage = core.Int64Ptr(int64(0))
	alertsV1MetricAlertPromqlConditionParametersModel.SwapNullValues = core.BoolPtr(true)

	alertsV1RelatedExtendedDataModel := new(logsv0.AlertsV1RelatedExtendedData)
	alertsV1RelatedExtendedDataModel.CleanupDeadmanDuration = core.StringPtr("cleanup_deadman_duration_never_or_unspecified")
	alertsV1RelatedExtendedDataModel.ShouldTriggerDeadman = core.BoolPtr(true)

	alertsV2ConditionParametersModel := new(logsv0.AlertsV2ConditionParameters)
	alertsV2ConditionParametersModel.Threshold = core.Float64Ptr(float64(72.5))
	alertsV2ConditionParametersModel.Timeframe = core.StringPtr("timeframe_5_min_or_unspecified")
	alertsV2ConditionParametersModel.GroupBy = []string{"testString"}
	alertsV2ConditionParametersModel.MetricAlertParameters = alertsV1MetricAlertConditionParametersModel
	alertsV2ConditionParametersModel.MetricAlertPromqlParameters = alertsV1MetricAlertPromqlConditionParametersModel
	alertsV2ConditionParametersModel.IgnoreInfinity = core.BoolPtr(true)
	alertsV2ConditionParametersModel.RelativeTimeframe = core.StringPtr("hour_or_unspecified")
	alertsV2ConditionParametersModel.CardinalityFields = []string{"testString"}
	alertsV2ConditionParametersModel.RelatedExtendedData = alertsV1RelatedExtendedDataModel

	model := new(logsv0.AlertsV2FlowCondition)
	model.Stages = []logsv0.AlertsV1FlowStage{*alertsV1FlowStageModel}
	model.Parameters = alertsV2ConditionParametersModel
	model.EnforceSuppression = core.BoolPtr(true)

	result, err := logs.DataSourceIbmLogsAlertsAlertsV2FlowConditionToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmLogsAlertsAlertsV1FlowStageToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		alertsV1FlowAlertModel := make(map[string]interface{})
		alertsV1FlowAlertModel["id"] = "testString"
		alertsV1FlowAlertModel["not"] = true

		alertsV1FlowAlertsModel := make(map[string]interface{})
		alertsV1FlowAlertsModel["op"] = "and"
		alertsV1FlowAlertsModel["values"] = []map[string]interface{}{alertsV1FlowAlertModel}

		alertsV1FlowGroupModel := make(map[string]interface{})
		alertsV1FlowGroupModel["alerts"] = []map[string]interface{}{alertsV1FlowAlertsModel}
		alertsV1FlowGroupModel["next_op"] = "and"

		alertsV1FlowTimeframeModel := make(map[string]interface{})
		alertsV1FlowTimeframeModel["ms"] = int(0)

		model := make(map[string]interface{})
		model["groups"] = []map[string]interface{}{alertsV1FlowGroupModel}
		model["timeframe"] = []map[string]interface{}{alertsV1FlowTimeframeModel}

		assert.Equal(t, result, model)
	}

	alertsV1FlowAlertModel := new(logsv0.AlertsV1FlowAlert)
	alertsV1FlowAlertModel.ID = core.StringPtr("testString")
	alertsV1FlowAlertModel.Not = core.BoolPtr(true)

	alertsV1FlowAlertsModel := new(logsv0.AlertsV1FlowAlerts)
	alertsV1FlowAlertsModel.Op = core.StringPtr("and")
	alertsV1FlowAlertsModel.Values = []logsv0.AlertsV1FlowAlert{*alertsV1FlowAlertModel}

	alertsV1FlowGroupModel := new(logsv0.AlertsV1FlowGroup)
	alertsV1FlowGroupModel.Alerts = alertsV1FlowAlertsModel
	alertsV1FlowGroupModel.NextOp = core.StringPtr("and")

	alertsV1FlowTimeframeModel := new(logsv0.AlertsV1FlowTimeframe)
	alertsV1FlowTimeframeModel.Ms = core.Int64Ptr(int64(0))

	model := new(logsv0.AlertsV1FlowStage)
	model.Groups = []logsv0.AlertsV1FlowGroup{*alertsV1FlowGroupModel}
	model.Timeframe = alertsV1FlowTimeframeModel

	result, err := logs.DataSourceIbmLogsAlertsAlertsV1FlowStageToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmLogsAlertsAlertsV1FlowGroupToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		alertsV1FlowAlertModel := make(map[string]interface{})
		alertsV1FlowAlertModel["id"] = "testString"
		alertsV1FlowAlertModel["not"] = true

		alertsV1FlowAlertsModel := make(map[string]interface{})
		alertsV1FlowAlertsModel["op"] = "and"
		alertsV1FlowAlertsModel["values"] = []map[string]interface{}{alertsV1FlowAlertModel}

		model := make(map[string]interface{})
		model["alerts"] = []map[string]interface{}{alertsV1FlowAlertsModel}
		model["next_op"] = "and"

		assert.Equal(t, result, model)
	}

	alertsV1FlowAlertModel := new(logsv0.AlertsV1FlowAlert)
	alertsV1FlowAlertModel.ID = core.StringPtr("testString")
	alertsV1FlowAlertModel.Not = core.BoolPtr(true)

	alertsV1FlowAlertsModel := new(logsv0.AlertsV1FlowAlerts)
	alertsV1FlowAlertsModel.Op = core.StringPtr("and")
	alertsV1FlowAlertsModel.Values = []logsv0.AlertsV1FlowAlert{*alertsV1FlowAlertModel}

	model := new(logsv0.AlertsV1FlowGroup)
	model.Alerts = alertsV1FlowAlertsModel
	model.NextOp = core.StringPtr("and")

	result, err := logs.DataSourceIbmLogsAlertsAlertsV1FlowGroupToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmLogsAlertsAlertsV1FlowAlertsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		alertsV1FlowAlertModel := make(map[string]interface{})
		alertsV1FlowAlertModel["id"] = "testString"
		alertsV1FlowAlertModel["not"] = true

		model := make(map[string]interface{})
		model["op"] = "and"
		model["values"] = []map[string]interface{}{alertsV1FlowAlertModel}

		assert.Equal(t, result, model)
	}

	alertsV1FlowAlertModel := new(logsv0.AlertsV1FlowAlert)
	alertsV1FlowAlertModel.ID = core.StringPtr("testString")
	alertsV1FlowAlertModel.Not = core.BoolPtr(true)

	model := new(logsv0.AlertsV1FlowAlerts)
	model.Op = core.StringPtr("and")
	model.Values = []logsv0.AlertsV1FlowAlert{*alertsV1FlowAlertModel}

	result, err := logs.DataSourceIbmLogsAlertsAlertsV1FlowAlertsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmLogsAlertsAlertsV1FlowAlertToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["id"] = "testString"
		model["not"] = true

		assert.Equal(t, result, model)
	}

	model := new(logsv0.AlertsV1FlowAlert)
	model.ID = core.StringPtr("testString")
	model.Not = core.BoolPtr(true)

	result, err := logs.DataSourceIbmLogsAlertsAlertsV1FlowAlertToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmLogsAlertsAlertsV1FlowTimeframeToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["ms"] = int(0)

		assert.Equal(t, result, model)
	}

	model := new(logsv0.AlertsV1FlowTimeframe)
	model.Ms = core.Int64Ptr(int64(0))

	result, err := logs.DataSourceIbmLogsAlertsAlertsV1FlowTimeframeToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmLogsAlertsAlertsV2UniqueCountConditionToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		alertsV1MetricAlertConditionParametersModel := make(map[string]interface{})
		alertsV1MetricAlertConditionParametersModel["metric_field"] = "testString"
		alertsV1MetricAlertConditionParametersModel["metric_source"] = "logs2metrics_or_unspecified"
		alertsV1MetricAlertConditionParametersModel["arithmetic_operator"] = "avg_or_unspecified"
		alertsV1MetricAlertConditionParametersModel["arithmetic_operator_modifier"] = int(0)
		alertsV1MetricAlertConditionParametersModel["sample_threshold_percentage"] = int(0)
		alertsV1MetricAlertConditionParametersModel["non_null_percentage"] = int(0)
		alertsV1MetricAlertConditionParametersModel["swap_null_values"] = true

		alertsV1MetricAlertPromqlConditionParametersModel := make(map[string]interface{})
		alertsV1MetricAlertPromqlConditionParametersModel["promql_text"] = "testString"
		alertsV1MetricAlertPromqlConditionParametersModel["arithmetic_operator_modifier"] = int(0)
		alertsV1MetricAlertPromqlConditionParametersModel["sample_threshold_percentage"] = int(0)
		alertsV1MetricAlertPromqlConditionParametersModel["non_null_percentage"] = int(0)
		alertsV1MetricAlertPromqlConditionParametersModel["swap_null_values"] = true

		alertsV1RelatedExtendedDataModel := make(map[string]interface{})
		alertsV1RelatedExtendedDataModel["cleanup_deadman_duration"] = "cleanup_deadman_duration_never_or_unspecified"
		alertsV1RelatedExtendedDataModel["should_trigger_deadman"] = true

		alertsV2ConditionParametersModel := make(map[string]interface{})
		alertsV2ConditionParametersModel["threshold"] = float64(72.5)
		alertsV2ConditionParametersModel["timeframe"] = "timeframe_5_min_or_unspecified"
		alertsV2ConditionParametersModel["group_by"] = []string{"testString"}
		alertsV2ConditionParametersModel["metric_alert_parameters"] = []map[string]interface{}{alertsV1MetricAlertConditionParametersModel}
		alertsV2ConditionParametersModel["metric_alert_promql_parameters"] = []map[string]interface{}{alertsV1MetricAlertPromqlConditionParametersModel}
		alertsV2ConditionParametersModel["ignore_infinity"] = true
		alertsV2ConditionParametersModel["relative_timeframe"] = "hour_or_unspecified"
		alertsV2ConditionParametersModel["cardinality_fields"] = []string{"testString"}
		alertsV2ConditionParametersModel["related_extended_data"] = []map[string]interface{}{alertsV1RelatedExtendedDataModel}

		model := make(map[string]interface{})
		model["parameters"] = []map[string]interface{}{alertsV2ConditionParametersModel}

		assert.Equal(t, result, model)
	}

	alertsV1MetricAlertConditionParametersModel := new(logsv0.AlertsV1MetricAlertConditionParameters)
	alertsV1MetricAlertConditionParametersModel.MetricField = core.StringPtr("testString")
	alertsV1MetricAlertConditionParametersModel.MetricSource = core.StringPtr("logs2metrics_or_unspecified")
	alertsV1MetricAlertConditionParametersModel.ArithmeticOperator = core.StringPtr("avg_or_unspecified")
	alertsV1MetricAlertConditionParametersModel.ArithmeticOperatorModifier = core.Int64Ptr(int64(0))
	alertsV1MetricAlertConditionParametersModel.SampleThresholdPercentage = core.Int64Ptr(int64(0))
	alertsV1MetricAlertConditionParametersModel.NonNullPercentage = core.Int64Ptr(int64(0))
	alertsV1MetricAlertConditionParametersModel.SwapNullValues = core.BoolPtr(true)

	alertsV1MetricAlertPromqlConditionParametersModel := new(logsv0.AlertsV1MetricAlertPromqlConditionParameters)
	alertsV1MetricAlertPromqlConditionParametersModel.PromqlText = core.StringPtr("testString")
	alertsV1MetricAlertPromqlConditionParametersModel.ArithmeticOperatorModifier = core.Int64Ptr(int64(0))
	alertsV1MetricAlertPromqlConditionParametersModel.SampleThresholdPercentage = core.Int64Ptr(int64(0))
	alertsV1MetricAlertPromqlConditionParametersModel.NonNullPercentage = core.Int64Ptr(int64(0))
	alertsV1MetricAlertPromqlConditionParametersModel.SwapNullValues = core.BoolPtr(true)

	alertsV1RelatedExtendedDataModel := new(logsv0.AlertsV1RelatedExtendedData)
	alertsV1RelatedExtendedDataModel.CleanupDeadmanDuration = core.StringPtr("cleanup_deadman_duration_never_or_unspecified")
	alertsV1RelatedExtendedDataModel.ShouldTriggerDeadman = core.BoolPtr(true)

	alertsV2ConditionParametersModel := new(logsv0.AlertsV2ConditionParameters)
	alertsV2ConditionParametersModel.Threshold = core.Float64Ptr(float64(72.5))
	alertsV2ConditionParametersModel.Timeframe = core.StringPtr("timeframe_5_min_or_unspecified")
	alertsV2ConditionParametersModel.GroupBy = []string{"testString"}
	alertsV2ConditionParametersModel.MetricAlertParameters = alertsV1MetricAlertConditionParametersModel
	alertsV2ConditionParametersModel.MetricAlertPromqlParameters = alertsV1MetricAlertPromqlConditionParametersModel
	alertsV2ConditionParametersModel.IgnoreInfinity = core.BoolPtr(true)
	alertsV2ConditionParametersModel.RelativeTimeframe = core.StringPtr("hour_or_unspecified")
	alertsV2ConditionParametersModel.CardinalityFields = []string{"testString"}
	alertsV2ConditionParametersModel.RelatedExtendedData = alertsV1RelatedExtendedDataModel

	model := new(logsv0.AlertsV2UniqueCountCondition)
	model.Parameters = alertsV2ConditionParametersModel

	result, err := logs.DataSourceIbmLogsAlertsAlertsV2UniqueCountConditionToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmLogsAlertsAlertsV2LessThanUsualConditionToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		alertsV1MetricAlertConditionParametersModel := make(map[string]interface{})
		alertsV1MetricAlertConditionParametersModel["metric_field"] = "testString"
		alertsV1MetricAlertConditionParametersModel["metric_source"] = "logs2metrics_or_unspecified"
		alertsV1MetricAlertConditionParametersModel["arithmetic_operator"] = "avg_or_unspecified"
		alertsV1MetricAlertConditionParametersModel["arithmetic_operator_modifier"] = int(0)
		alertsV1MetricAlertConditionParametersModel["sample_threshold_percentage"] = int(0)
		alertsV1MetricAlertConditionParametersModel["non_null_percentage"] = int(0)
		alertsV1MetricAlertConditionParametersModel["swap_null_values"] = true

		alertsV1MetricAlertPromqlConditionParametersModel := make(map[string]interface{})
		alertsV1MetricAlertPromqlConditionParametersModel["promql_text"] = "testString"
		alertsV1MetricAlertPromqlConditionParametersModel["arithmetic_operator_modifier"] = int(0)
		alertsV1MetricAlertPromqlConditionParametersModel["sample_threshold_percentage"] = int(0)
		alertsV1MetricAlertPromqlConditionParametersModel["non_null_percentage"] = int(0)
		alertsV1MetricAlertPromqlConditionParametersModel["swap_null_values"] = true

		alertsV1RelatedExtendedDataModel := make(map[string]interface{})
		alertsV1RelatedExtendedDataModel["cleanup_deadman_duration"] = "cleanup_deadman_duration_never_or_unspecified"
		alertsV1RelatedExtendedDataModel["should_trigger_deadman"] = true

		alertsV2ConditionParametersModel := make(map[string]interface{})
		alertsV2ConditionParametersModel["threshold"] = float64(72.5)
		alertsV2ConditionParametersModel["timeframe"] = "timeframe_5_min_or_unspecified"
		alertsV2ConditionParametersModel["group_by"] = []string{"testString"}
		alertsV2ConditionParametersModel["metric_alert_parameters"] = []map[string]interface{}{alertsV1MetricAlertConditionParametersModel}
		alertsV2ConditionParametersModel["metric_alert_promql_parameters"] = []map[string]interface{}{alertsV1MetricAlertPromqlConditionParametersModel}
		alertsV2ConditionParametersModel["ignore_infinity"] = true
		alertsV2ConditionParametersModel["relative_timeframe"] = "hour_or_unspecified"
		alertsV2ConditionParametersModel["cardinality_fields"] = []string{"testString"}
		alertsV2ConditionParametersModel["related_extended_data"] = []map[string]interface{}{alertsV1RelatedExtendedDataModel}

		model := make(map[string]interface{})
		model["parameters"] = []map[string]interface{}{alertsV2ConditionParametersModel}

		assert.Equal(t, result, model)
	}

	alertsV1MetricAlertConditionParametersModel := new(logsv0.AlertsV1MetricAlertConditionParameters)
	alertsV1MetricAlertConditionParametersModel.MetricField = core.StringPtr("testString")
	alertsV1MetricAlertConditionParametersModel.MetricSource = core.StringPtr("logs2metrics_or_unspecified")
	alertsV1MetricAlertConditionParametersModel.ArithmeticOperator = core.StringPtr("avg_or_unspecified")
	alertsV1MetricAlertConditionParametersModel.ArithmeticOperatorModifier = core.Int64Ptr(int64(0))
	alertsV1MetricAlertConditionParametersModel.SampleThresholdPercentage = core.Int64Ptr(int64(0))
	alertsV1MetricAlertConditionParametersModel.NonNullPercentage = core.Int64Ptr(int64(0))
	alertsV1MetricAlertConditionParametersModel.SwapNullValues = core.BoolPtr(true)

	alertsV1MetricAlertPromqlConditionParametersModel := new(logsv0.AlertsV1MetricAlertPromqlConditionParameters)
	alertsV1MetricAlertPromqlConditionParametersModel.PromqlText = core.StringPtr("testString")
	alertsV1MetricAlertPromqlConditionParametersModel.ArithmeticOperatorModifier = core.Int64Ptr(int64(0))
	alertsV1MetricAlertPromqlConditionParametersModel.SampleThresholdPercentage = core.Int64Ptr(int64(0))
	alertsV1MetricAlertPromqlConditionParametersModel.NonNullPercentage = core.Int64Ptr(int64(0))
	alertsV1MetricAlertPromqlConditionParametersModel.SwapNullValues = core.BoolPtr(true)

	alertsV1RelatedExtendedDataModel := new(logsv0.AlertsV1RelatedExtendedData)
	alertsV1RelatedExtendedDataModel.CleanupDeadmanDuration = core.StringPtr("cleanup_deadman_duration_never_or_unspecified")
	alertsV1RelatedExtendedDataModel.ShouldTriggerDeadman = core.BoolPtr(true)

	alertsV2ConditionParametersModel := new(logsv0.AlertsV2ConditionParameters)
	alertsV2ConditionParametersModel.Threshold = core.Float64Ptr(float64(72.5))
	alertsV2ConditionParametersModel.Timeframe = core.StringPtr("timeframe_5_min_or_unspecified")
	alertsV2ConditionParametersModel.GroupBy = []string{"testString"}
	alertsV2ConditionParametersModel.MetricAlertParameters = alertsV1MetricAlertConditionParametersModel
	alertsV2ConditionParametersModel.MetricAlertPromqlParameters = alertsV1MetricAlertPromqlConditionParametersModel
	alertsV2ConditionParametersModel.IgnoreInfinity = core.BoolPtr(true)
	alertsV2ConditionParametersModel.RelativeTimeframe = core.StringPtr("hour_or_unspecified")
	alertsV2ConditionParametersModel.CardinalityFields = []string{"testString"}
	alertsV2ConditionParametersModel.RelatedExtendedData = alertsV1RelatedExtendedDataModel

	model := new(logsv0.AlertsV2LessThanUsualCondition)
	model.Parameters = alertsV2ConditionParametersModel

	result, err := logs.DataSourceIbmLogsAlertsAlertsV2LessThanUsualConditionToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmLogsAlertsAlertsV2AlertConditionConditionImmediateToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		alertsV2ImmediateConditionModel := make(map[string]interface{})

		model := make(map[string]interface{})
		model["immediate"] = []map[string]interface{}{alertsV2ImmediateConditionModel}

		assert.Equal(t, result, model)
	}

	alertsV2ImmediateConditionModel := new(logsv0.AlertsV2ImmediateCondition)

	model := new(logsv0.AlertsV2AlertConditionConditionImmediate)
	model.Immediate = alertsV2ImmediateConditionModel

	result, err := logs.DataSourceIbmLogsAlertsAlertsV2AlertConditionConditionImmediateToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmLogsAlertsAlertsV2AlertConditionConditionLessThanToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		alertsV1MetricAlertConditionParametersModel := make(map[string]interface{})
		alertsV1MetricAlertConditionParametersModel["metric_field"] = "testString"
		alertsV1MetricAlertConditionParametersModel["metric_source"] = "logs2metrics_or_unspecified"
		alertsV1MetricAlertConditionParametersModel["arithmetic_operator"] = "avg_or_unspecified"
		alertsV1MetricAlertConditionParametersModel["arithmetic_operator_modifier"] = int(0)
		alertsV1MetricAlertConditionParametersModel["sample_threshold_percentage"] = int(0)
		alertsV1MetricAlertConditionParametersModel["non_null_percentage"] = int(0)
		alertsV1MetricAlertConditionParametersModel["swap_null_values"] = true

		alertsV1MetricAlertPromqlConditionParametersModel := make(map[string]interface{})
		alertsV1MetricAlertPromqlConditionParametersModel["promql_text"] = "testString"
		alertsV1MetricAlertPromqlConditionParametersModel["arithmetic_operator_modifier"] = int(0)
		alertsV1MetricAlertPromqlConditionParametersModel["sample_threshold_percentage"] = int(0)
		alertsV1MetricAlertPromqlConditionParametersModel["non_null_percentage"] = int(0)
		alertsV1MetricAlertPromqlConditionParametersModel["swap_null_values"] = true

		alertsV1RelatedExtendedDataModel := make(map[string]interface{})
		alertsV1RelatedExtendedDataModel["cleanup_deadman_duration"] = "cleanup_deadman_duration_never_or_unspecified"
		alertsV1RelatedExtendedDataModel["should_trigger_deadman"] = true

		alertsV2ConditionParametersModel := make(map[string]interface{})
		alertsV2ConditionParametersModel["threshold"] = float64(72.5)
		alertsV2ConditionParametersModel["timeframe"] = "timeframe_5_min_or_unspecified"
		alertsV2ConditionParametersModel["group_by"] = []string{"testString"}
		alertsV2ConditionParametersModel["metric_alert_parameters"] = []map[string]interface{}{alertsV1MetricAlertConditionParametersModel}
		alertsV2ConditionParametersModel["metric_alert_promql_parameters"] = []map[string]interface{}{alertsV1MetricAlertPromqlConditionParametersModel}
		alertsV2ConditionParametersModel["ignore_infinity"] = true
		alertsV2ConditionParametersModel["relative_timeframe"] = "hour_or_unspecified"
		alertsV2ConditionParametersModel["cardinality_fields"] = []string{"testString"}
		alertsV2ConditionParametersModel["related_extended_data"] = []map[string]interface{}{alertsV1RelatedExtendedDataModel}

		alertsV2LessThanConditionModel := make(map[string]interface{})
		alertsV2LessThanConditionModel["parameters"] = []map[string]interface{}{alertsV2ConditionParametersModel}

		model := make(map[string]interface{})
		model["less_than"] = []map[string]interface{}{alertsV2LessThanConditionModel}

		assert.Equal(t, result, model)
	}

	alertsV1MetricAlertConditionParametersModel := new(logsv0.AlertsV1MetricAlertConditionParameters)
	alertsV1MetricAlertConditionParametersModel.MetricField = core.StringPtr("testString")
	alertsV1MetricAlertConditionParametersModel.MetricSource = core.StringPtr("logs2metrics_or_unspecified")
	alertsV1MetricAlertConditionParametersModel.ArithmeticOperator = core.StringPtr("avg_or_unspecified")
	alertsV1MetricAlertConditionParametersModel.ArithmeticOperatorModifier = core.Int64Ptr(int64(0))
	alertsV1MetricAlertConditionParametersModel.SampleThresholdPercentage = core.Int64Ptr(int64(0))
	alertsV1MetricAlertConditionParametersModel.NonNullPercentage = core.Int64Ptr(int64(0))
	alertsV1MetricAlertConditionParametersModel.SwapNullValues = core.BoolPtr(true)

	alertsV1MetricAlertPromqlConditionParametersModel := new(logsv0.AlertsV1MetricAlertPromqlConditionParameters)
	alertsV1MetricAlertPromqlConditionParametersModel.PromqlText = core.StringPtr("testString")
	alertsV1MetricAlertPromqlConditionParametersModel.ArithmeticOperatorModifier = core.Int64Ptr(int64(0))
	alertsV1MetricAlertPromqlConditionParametersModel.SampleThresholdPercentage = core.Int64Ptr(int64(0))
	alertsV1MetricAlertPromqlConditionParametersModel.NonNullPercentage = core.Int64Ptr(int64(0))
	alertsV1MetricAlertPromqlConditionParametersModel.SwapNullValues = core.BoolPtr(true)

	alertsV1RelatedExtendedDataModel := new(logsv0.AlertsV1RelatedExtendedData)
	alertsV1RelatedExtendedDataModel.CleanupDeadmanDuration = core.StringPtr("cleanup_deadman_duration_never_or_unspecified")
	alertsV1RelatedExtendedDataModel.ShouldTriggerDeadman = core.BoolPtr(true)

	alertsV2ConditionParametersModel := new(logsv0.AlertsV2ConditionParameters)
	alertsV2ConditionParametersModel.Threshold = core.Float64Ptr(float64(72.5))
	alertsV2ConditionParametersModel.Timeframe = core.StringPtr("timeframe_5_min_or_unspecified")
	alertsV2ConditionParametersModel.GroupBy = []string{"testString"}
	alertsV2ConditionParametersModel.MetricAlertParameters = alertsV1MetricAlertConditionParametersModel
	alertsV2ConditionParametersModel.MetricAlertPromqlParameters = alertsV1MetricAlertPromqlConditionParametersModel
	alertsV2ConditionParametersModel.IgnoreInfinity = core.BoolPtr(true)
	alertsV2ConditionParametersModel.RelativeTimeframe = core.StringPtr("hour_or_unspecified")
	alertsV2ConditionParametersModel.CardinalityFields = []string{"testString"}
	alertsV2ConditionParametersModel.RelatedExtendedData = alertsV1RelatedExtendedDataModel

	alertsV2LessThanConditionModel := new(logsv0.AlertsV2LessThanCondition)
	alertsV2LessThanConditionModel.Parameters = alertsV2ConditionParametersModel

	model := new(logsv0.AlertsV2AlertConditionConditionLessThan)
	model.LessThan = alertsV2LessThanConditionModel

	result, err := logs.DataSourceIbmLogsAlertsAlertsV2AlertConditionConditionLessThanToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmLogsAlertsAlertsV2AlertConditionConditionMoreThanToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		alertsV1MetricAlertConditionParametersModel := make(map[string]interface{})
		alertsV1MetricAlertConditionParametersModel["metric_field"] = "testString"
		alertsV1MetricAlertConditionParametersModel["metric_source"] = "logs2metrics_or_unspecified"
		alertsV1MetricAlertConditionParametersModel["arithmetic_operator"] = "avg_or_unspecified"
		alertsV1MetricAlertConditionParametersModel["arithmetic_operator_modifier"] = int(0)
		alertsV1MetricAlertConditionParametersModel["sample_threshold_percentage"] = int(0)
		alertsV1MetricAlertConditionParametersModel["non_null_percentage"] = int(0)
		alertsV1MetricAlertConditionParametersModel["swap_null_values"] = true

		alertsV1MetricAlertPromqlConditionParametersModel := make(map[string]interface{})
		alertsV1MetricAlertPromqlConditionParametersModel["promql_text"] = "testString"
		alertsV1MetricAlertPromqlConditionParametersModel["arithmetic_operator_modifier"] = int(0)
		alertsV1MetricAlertPromqlConditionParametersModel["sample_threshold_percentage"] = int(0)
		alertsV1MetricAlertPromqlConditionParametersModel["non_null_percentage"] = int(0)
		alertsV1MetricAlertPromqlConditionParametersModel["swap_null_values"] = true

		alertsV1RelatedExtendedDataModel := make(map[string]interface{})
		alertsV1RelatedExtendedDataModel["cleanup_deadman_duration"] = "cleanup_deadman_duration_never_or_unspecified"
		alertsV1RelatedExtendedDataModel["should_trigger_deadman"] = true

		alertsV2ConditionParametersModel := make(map[string]interface{})
		alertsV2ConditionParametersModel["threshold"] = float64(72.5)
		alertsV2ConditionParametersModel["timeframe"] = "timeframe_5_min_or_unspecified"
		alertsV2ConditionParametersModel["group_by"] = []string{"testString"}
		alertsV2ConditionParametersModel["metric_alert_parameters"] = []map[string]interface{}{alertsV1MetricAlertConditionParametersModel}
		alertsV2ConditionParametersModel["metric_alert_promql_parameters"] = []map[string]interface{}{alertsV1MetricAlertPromqlConditionParametersModel}
		alertsV2ConditionParametersModel["ignore_infinity"] = true
		alertsV2ConditionParametersModel["relative_timeframe"] = "hour_or_unspecified"
		alertsV2ConditionParametersModel["cardinality_fields"] = []string{"testString"}
		alertsV2ConditionParametersModel["related_extended_data"] = []map[string]interface{}{alertsV1RelatedExtendedDataModel}

		alertsV2MoreThanConditionModel := make(map[string]interface{})
		alertsV2MoreThanConditionModel["parameters"] = []map[string]interface{}{alertsV2ConditionParametersModel}
		alertsV2MoreThanConditionModel["evaluation_window"] = "rolling_or_unspecified"

		model := make(map[string]interface{})
		model["more_than"] = []map[string]interface{}{alertsV2MoreThanConditionModel}

		assert.Equal(t, result, model)
	}

	alertsV1MetricAlertConditionParametersModel := new(logsv0.AlertsV1MetricAlertConditionParameters)
	alertsV1MetricAlertConditionParametersModel.MetricField = core.StringPtr("testString")
	alertsV1MetricAlertConditionParametersModel.MetricSource = core.StringPtr("logs2metrics_or_unspecified")
	alertsV1MetricAlertConditionParametersModel.ArithmeticOperator = core.StringPtr("avg_or_unspecified")
	alertsV1MetricAlertConditionParametersModel.ArithmeticOperatorModifier = core.Int64Ptr(int64(0))
	alertsV1MetricAlertConditionParametersModel.SampleThresholdPercentage = core.Int64Ptr(int64(0))
	alertsV1MetricAlertConditionParametersModel.NonNullPercentage = core.Int64Ptr(int64(0))
	alertsV1MetricAlertConditionParametersModel.SwapNullValues = core.BoolPtr(true)

	alertsV1MetricAlertPromqlConditionParametersModel := new(logsv0.AlertsV1MetricAlertPromqlConditionParameters)
	alertsV1MetricAlertPromqlConditionParametersModel.PromqlText = core.StringPtr("testString")
	alertsV1MetricAlertPromqlConditionParametersModel.ArithmeticOperatorModifier = core.Int64Ptr(int64(0))
	alertsV1MetricAlertPromqlConditionParametersModel.SampleThresholdPercentage = core.Int64Ptr(int64(0))
	alertsV1MetricAlertPromqlConditionParametersModel.NonNullPercentage = core.Int64Ptr(int64(0))
	alertsV1MetricAlertPromqlConditionParametersModel.SwapNullValues = core.BoolPtr(true)

	alertsV1RelatedExtendedDataModel := new(logsv0.AlertsV1RelatedExtendedData)
	alertsV1RelatedExtendedDataModel.CleanupDeadmanDuration = core.StringPtr("cleanup_deadman_duration_never_or_unspecified")
	alertsV1RelatedExtendedDataModel.ShouldTriggerDeadman = core.BoolPtr(true)

	alertsV2ConditionParametersModel := new(logsv0.AlertsV2ConditionParameters)
	alertsV2ConditionParametersModel.Threshold = core.Float64Ptr(float64(72.5))
	alertsV2ConditionParametersModel.Timeframe = core.StringPtr("timeframe_5_min_or_unspecified")
	alertsV2ConditionParametersModel.GroupBy = []string{"testString"}
	alertsV2ConditionParametersModel.MetricAlertParameters = alertsV1MetricAlertConditionParametersModel
	alertsV2ConditionParametersModel.MetricAlertPromqlParameters = alertsV1MetricAlertPromqlConditionParametersModel
	alertsV2ConditionParametersModel.IgnoreInfinity = core.BoolPtr(true)
	alertsV2ConditionParametersModel.RelativeTimeframe = core.StringPtr("hour_or_unspecified")
	alertsV2ConditionParametersModel.CardinalityFields = []string{"testString"}
	alertsV2ConditionParametersModel.RelatedExtendedData = alertsV1RelatedExtendedDataModel

	alertsV2MoreThanConditionModel := new(logsv0.AlertsV2MoreThanCondition)
	alertsV2MoreThanConditionModel.Parameters = alertsV2ConditionParametersModel
	alertsV2MoreThanConditionModel.EvaluationWindow = core.StringPtr("rolling_or_unspecified")

	model := new(logsv0.AlertsV2AlertConditionConditionMoreThan)
	model.MoreThan = alertsV2MoreThanConditionModel

	result, err := logs.DataSourceIbmLogsAlertsAlertsV2AlertConditionConditionMoreThanToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmLogsAlertsAlertsV2AlertConditionConditionMoreThanUsualToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		alertsV1MetricAlertConditionParametersModel := make(map[string]interface{})
		alertsV1MetricAlertConditionParametersModel["metric_field"] = "testString"
		alertsV1MetricAlertConditionParametersModel["metric_source"] = "logs2metrics_or_unspecified"
		alertsV1MetricAlertConditionParametersModel["arithmetic_operator"] = "avg_or_unspecified"
		alertsV1MetricAlertConditionParametersModel["arithmetic_operator_modifier"] = int(0)
		alertsV1MetricAlertConditionParametersModel["sample_threshold_percentage"] = int(0)
		alertsV1MetricAlertConditionParametersModel["non_null_percentage"] = int(0)
		alertsV1MetricAlertConditionParametersModel["swap_null_values"] = true

		alertsV1MetricAlertPromqlConditionParametersModel := make(map[string]interface{})
		alertsV1MetricAlertPromqlConditionParametersModel["promql_text"] = "testString"
		alertsV1MetricAlertPromqlConditionParametersModel["arithmetic_operator_modifier"] = int(0)
		alertsV1MetricAlertPromqlConditionParametersModel["sample_threshold_percentage"] = int(0)
		alertsV1MetricAlertPromqlConditionParametersModel["non_null_percentage"] = int(0)
		alertsV1MetricAlertPromqlConditionParametersModel["swap_null_values"] = true

		alertsV1RelatedExtendedDataModel := make(map[string]interface{})
		alertsV1RelatedExtendedDataModel["cleanup_deadman_duration"] = "cleanup_deadman_duration_never_or_unspecified"
		alertsV1RelatedExtendedDataModel["should_trigger_deadman"] = true

		alertsV2ConditionParametersModel := make(map[string]interface{})
		alertsV2ConditionParametersModel["threshold"] = float64(72.5)
		alertsV2ConditionParametersModel["timeframe"] = "timeframe_5_min_or_unspecified"
		alertsV2ConditionParametersModel["group_by"] = []string{"testString"}
		alertsV2ConditionParametersModel["metric_alert_parameters"] = []map[string]interface{}{alertsV1MetricAlertConditionParametersModel}
		alertsV2ConditionParametersModel["metric_alert_promql_parameters"] = []map[string]interface{}{alertsV1MetricAlertPromqlConditionParametersModel}
		alertsV2ConditionParametersModel["ignore_infinity"] = true
		alertsV2ConditionParametersModel["relative_timeframe"] = "hour_or_unspecified"
		alertsV2ConditionParametersModel["cardinality_fields"] = []string{"testString"}
		alertsV2ConditionParametersModel["related_extended_data"] = []map[string]interface{}{alertsV1RelatedExtendedDataModel}

		alertsV2MoreThanUsualConditionModel := make(map[string]interface{})
		alertsV2MoreThanUsualConditionModel["parameters"] = []map[string]interface{}{alertsV2ConditionParametersModel}

		model := make(map[string]interface{})
		model["more_than_usual"] = []map[string]interface{}{alertsV2MoreThanUsualConditionModel}

		assert.Equal(t, result, model)
	}

	alertsV1MetricAlertConditionParametersModel := new(logsv0.AlertsV1MetricAlertConditionParameters)
	alertsV1MetricAlertConditionParametersModel.MetricField = core.StringPtr("testString")
	alertsV1MetricAlertConditionParametersModel.MetricSource = core.StringPtr("logs2metrics_or_unspecified")
	alertsV1MetricAlertConditionParametersModel.ArithmeticOperator = core.StringPtr("avg_or_unspecified")
	alertsV1MetricAlertConditionParametersModel.ArithmeticOperatorModifier = core.Int64Ptr(int64(0))
	alertsV1MetricAlertConditionParametersModel.SampleThresholdPercentage = core.Int64Ptr(int64(0))
	alertsV1MetricAlertConditionParametersModel.NonNullPercentage = core.Int64Ptr(int64(0))
	alertsV1MetricAlertConditionParametersModel.SwapNullValues = core.BoolPtr(true)

	alertsV1MetricAlertPromqlConditionParametersModel := new(logsv0.AlertsV1MetricAlertPromqlConditionParameters)
	alertsV1MetricAlertPromqlConditionParametersModel.PromqlText = core.StringPtr("testString")
	alertsV1MetricAlertPromqlConditionParametersModel.ArithmeticOperatorModifier = core.Int64Ptr(int64(0))
	alertsV1MetricAlertPromqlConditionParametersModel.SampleThresholdPercentage = core.Int64Ptr(int64(0))
	alertsV1MetricAlertPromqlConditionParametersModel.NonNullPercentage = core.Int64Ptr(int64(0))
	alertsV1MetricAlertPromqlConditionParametersModel.SwapNullValues = core.BoolPtr(true)

	alertsV1RelatedExtendedDataModel := new(logsv0.AlertsV1RelatedExtendedData)
	alertsV1RelatedExtendedDataModel.CleanupDeadmanDuration = core.StringPtr("cleanup_deadman_duration_never_or_unspecified")
	alertsV1RelatedExtendedDataModel.ShouldTriggerDeadman = core.BoolPtr(true)

	alertsV2ConditionParametersModel := new(logsv0.AlertsV2ConditionParameters)
	alertsV2ConditionParametersModel.Threshold = core.Float64Ptr(float64(72.5))
	alertsV2ConditionParametersModel.Timeframe = core.StringPtr("timeframe_5_min_or_unspecified")
	alertsV2ConditionParametersModel.GroupBy = []string{"testString"}
	alertsV2ConditionParametersModel.MetricAlertParameters = alertsV1MetricAlertConditionParametersModel
	alertsV2ConditionParametersModel.MetricAlertPromqlParameters = alertsV1MetricAlertPromqlConditionParametersModel
	alertsV2ConditionParametersModel.IgnoreInfinity = core.BoolPtr(true)
	alertsV2ConditionParametersModel.RelativeTimeframe = core.StringPtr("hour_or_unspecified")
	alertsV2ConditionParametersModel.CardinalityFields = []string{"testString"}
	alertsV2ConditionParametersModel.RelatedExtendedData = alertsV1RelatedExtendedDataModel

	alertsV2MoreThanUsualConditionModel := new(logsv0.AlertsV2MoreThanUsualCondition)
	alertsV2MoreThanUsualConditionModel.Parameters = alertsV2ConditionParametersModel

	model := new(logsv0.AlertsV2AlertConditionConditionMoreThanUsual)
	model.MoreThanUsual = alertsV2MoreThanUsualConditionModel

	result, err := logs.DataSourceIbmLogsAlertsAlertsV2AlertConditionConditionMoreThanUsualToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmLogsAlertsAlertsV2AlertConditionConditionNewValueToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		alertsV1MetricAlertConditionParametersModel := make(map[string]interface{})
		alertsV1MetricAlertConditionParametersModel["metric_field"] = "testString"
		alertsV1MetricAlertConditionParametersModel["metric_source"] = "logs2metrics_or_unspecified"
		alertsV1MetricAlertConditionParametersModel["arithmetic_operator"] = "avg_or_unspecified"
		alertsV1MetricAlertConditionParametersModel["arithmetic_operator_modifier"] = int(0)
		alertsV1MetricAlertConditionParametersModel["sample_threshold_percentage"] = int(0)
		alertsV1MetricAlertConditionParametersModel["non_null_percentage"] = int(0)
		alertsV1MetricAlertConditionParametersModel["swap_null_values"] = true

		alertsV1MetricAlertPromqlConditionParametersModel := make(map[string]interface{})
		alertsV1MetricAlertPromqlConditionParametersModel["promql_text"] = "testString"
		alertsV1MetricAlertPromqlConditionParametersModel["arithmetic_operator_modifier"] = int(0)
		alertsV1MetricAlertPromqlConditionParametersModel["sample_threshold_percentage"] = int(0)
		alertsV1MetricAlertPromqlConditionParametersModel["non_null_percentage"] = int(0)
		alertsV1MetricAlertPromqlConditionParametersModel["swap_null_values"] = true

		alertsV1RelatedExtendedDataModel := make(map[string]interface{})
		alertsV1RelatedExtendedDataModel["cleanup_deadman_duration"] = "cleanup_deadman_duration_never_or_unspecified"
		alertsV1RelatedExtendedDataModel["should_trigger_deadman"] = true

		alertsV2ConditionParametersModel := make(map[string]interface{})
		alertsV2ConditionParametersModel["threshold"] = float64(72.5)
		alertsV2ConditionParametersModel["timeframe"] = "timeframe_5_min_or_unspecified"
		alertsV2ConditionParametersModel["group_by"] = []string{"testString"}
		alertsV2ConditionParametersModel["metric_alert_parameters"] = []map[string]interface{}{alertsV1MetricAlertConditionParametersModel}
		alertsV2ConditionParametersModel["metric_alert_promql_parameters"] = []map[string]interface{}{alertsV1MetricAlertPromqlConditionParametersModel}
		alertsV2ConditionParametersModel["ignore_infinity"] = true
		alertsV2ConditionParametersModel["relative_timeframe"] = "hour_or_unspecified"
		alertsV2ConditionParametersModel["cardinality_fields"] = []string{"testString"}
		alertsV2ConditionParametersModel["related_extended_data"] = []map[string]interface{}{alertsV1RelatedExtendedDataModel}

		alertsV2NewValueConditionModel := make(map[string]interface{})
		alertsV2NewValueConditionModel["parameters"] = []map[string]interface{}{alertsV2ConditionParametersModel}

		model := make(map[string]interface{})
		model["new_value"] = []map[string]interface{}{alertsV2NewValueConditionModel}

		assert.Equal(t, result, model)
	}

	alertsV1MetricAlertConditionParametersModel := new(logsv0.AlertsV1MetricAlertConditionParameters)
	alertsV1MetricAlertConditionParametersModel.MetricField = core.StringPtr("testString")
	alertsV1MetricAlertConditionParametersModel.MetricSource = core.StringPtr("logs2metrics_or_unspecified")
	alertsV1MetricAlertConditionParametersModel.ArithmeticOperator = core.StringPtr("avg_or_unspecified")
	alertsV1MetricAlertConditionParametersModel.ArithmeticOperatorModifier = core.Int64Ptr(int64(0))
	alertsV1MetricAlertConditionParametersModel.SampleThresholdPercentage = core.Int64Ptr(int64(0))
	alertsV1MetricAlertConditionParametersModel.NonNullPercentage = core.Int64Ptr(int64(0))
	alertsV1MetricAlertConditionParametersModel.SwapNullValues = core.BoolPtr(true)

	alertsV1MetricAlertPromqlConditionParametersModel := new(logsv0.AlertsV1MetricAlertPromqlConditionParameters)
	alertsV1MetricAlertPromqlConditionParametersModel.PromqlText = core.StringPtr("testString")
	alertsV1MetricAlertPromqlConditionParametersModel.ArithmeticOperatorModifier = core.Int64Ptr(int64(0))
	alertsV1MetricAlertPromqlConditionParametersModel.SampleThresholdPercentage = core.Int64Ptr(int64(0))
	alertsV1MetricAlertPromqlConditionParametersModel.NonNullPercentage = core.Int64Ptr(int64(0))
	alertsV1MetricAlertPromqlConditionParametersModel.SwapNullValues = core.BoolPtr(true)

	alertsV1RelatedExtendedDataModel := new(logsv0.AlertsV1RelatedExtendedData)
	alertsV1RelatedExtendedDataModel.CleanupDeadmanDuration = core.StringPtr("cleanup_deadman_duration_never_or_unspecified")
	alertsV1RelatedExtendedDataModel.ShouldTriggerDeadman = core.BoolPtr(true)

	alertsV2ConditionParametersModel := new(logsv0.AlertsV2ConditionParameters)
	alertsV2ConditionParametersModel.Threshold = core.Float64Ptr(float64(72.5))
	alertsV2ConditionParametersModel.Timeframe = core.StringPtr("timeframe_5_min_or_unspecified")
	alertsV2ConditionParametersModel.GroupBy = []string{"testString"}
	alertsV2ConditionParametersModel.MetricAlertParameters = alertsV1MetricAlertConditionParametersModel
	alertsV2ConditionParametersModel.MetricAlertPromqlParameters = alertsV1MetricAlertPromqlConditionParametersModel
	alertsV2ConditionParametersModel.IgnoreInfinity = core.BoolPtr(true)
	alertsV2ConditionParametersModel.RelativeTimeframe = core.StringPtr("hour_or_unspecified")
	alertsV2ConditionParametersModel.CardinalityFields = []string{"testString"}
	alertsV2ConditionParametersModel.RelatedExtendedData = alertsV1RelatedExtendedDataModel

	alertsV2NewValueConditionModel := new(logsv0.AlertsV2NewValueCondition)
	alertsV2NewValueConditionModel.Parameters = alertsV2ConditionParametersModel

	model := new(logsv0.AlertsV2AlertConditionConditionNewValue)
	model.NewValue = alertsV2NewValueConditionModel

	result, err := logs.DataSourceIbmLogsAlertsAlertsV2AlertConditionConditionNewValueToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmLogsAlertsAlertsV2AlertConditionConditionFlowToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		alertsV1FlowAlertModel := make(map[string]interface{})
		alertsV1FlowAlertModel["id"] = "testString"
		alertsV1FlowAlertModel["not"] = true

		alertsV1FlowAlertsModel := make(map[string]interface{})
		alertsV1FlowAlertsModel["op"] = "and"
		alertsV1FlowAlertsModel["values"] = []map[string]interface{}{alertsV1FlowAlertModel}

		alertsV1FlowGroupModel := make(map[string]interface{})
		alertsV1FlowGroupModel["alerts"] = []map[string]interface{}{alertsV1FlowAlertsModel}
		alertsV1FlowGroupModel["next_op"] = "and"

		alertsV1FlowTimeframeModel := make(map[string]interface{})
		alertsV1FlowTimeframeModel["ms"] = int(0)

		alertsV1FlowStageModel := make(map[string]interface{})
		alertsV1FlowStageModel["groups"] = []map[string]interface{}{alertsV1FlowGroupModel}
		alertsV1FlowStageModel["timeframe"] = []map[string]interface{}{alertsV1FlowTimeframeModel}

		alertsV1MetricAlertConditionParametersModel := make(map[string]interface{})
		alertsV1MetricAlertConditionParametersModel["metric_field"] = "testString"
		alertsV1MetricAlertConditionParametersModel["metric_source"] = "logs2metrics_or_unspecified"
		alertsV1MetricAlertConditionParametersModel["arithmetic_operator"] = "avg_or_unspecified"
		alertsV1MetricAlertConditionParametersModel["arithmetic_operator_modifier"] = int(0)
		alertsV1MetricAlertConditionParametersModel["sample_threshold_percentage"] = int(0)
		alertsV1MetricAlertConditionParametersModel["non_null_percentage"] = int(0)
		alertsV1MetricAlertConditionParametersModel["swap_null_values"] = true

		alertsV1MetricAlertPromqlConditionParametersModel := make(map[string]interface{})
		alertsV1MetricAlertPromqlConditionParametersModel["promql_text"] = "testString"
		alertsV1MetricAlertPromqlConditionParametersModel["arithmetic_operator_modifier"] = int(0)
		alertsV1MetricAlertPromqlConditionParametersModel["sample_threshold_percentage"] = int(0)
		alertsV1MetricAlertPromqlConditionParametersModel["non_null_percentage"] = int(0)
		alertsV1MetricAlertPromqlConditionParametersModel["swap_null_values"] = true

		alertsV1RelatedExtendedDataModel := make(map[string]interface{})
		alertsV1RelatedExtendedDataModel["cleanup_deadman_duration"] = "cleanup_deadman_duration_never_or_unspecified"
		alertsV1RelatedExtendedDataModel["should_trigger_deadman"] = true

		alertsV2ConditionParametersModel := make(map[string]interface{})
		alertsV2ConditionParametersModel["threshold"] = float64(72.5)
		alertsV2ConditionParametersModel["timeframe"] = "timeframe_5_min_or_unspecified"
		alertsV2ConditionParametersModel["group_by"] = []string{"testString"}
		alertsV2ConditionParametersModel["metric_alert_parameters"] = []map[string]interface{}{alertsV1MetricAlertConditionParametersModel}
		alertsV2ConditionParametersModel["metric_alert_promql_parameters"] = []map[string]interface{}{alertsV1MetricAlertPromqlConditionParametersModel}
		alertsV2ConditionParametersModel["ignore_infinity"] = true
		alertsV2ConditionParametersModel["relative_timeframe"] = "hour_or_unspecified"
		alertsV2ConditionParametersModel["cardinality_fields"] = []string{"testString"}
		alertsV2ConditionParametersModel["related_extended_data"] = []map[string]interface{}{alertsV1RelatedExtendedDataModel}

		alertsV2FlowConditionModel := make(map[string]interface{})
		alertsV2FlowConditionModel["stages"] = []map[string]interface{}{alertsV1FlowStageModel}
		alertsV2FlowConditionModel["parameters"] = []map[string]interface{}{alertsV2ConditionParametersModel}
		alertsV2FlowConditionModel["enforce_suppression"] = true

		model := make(map[string]interface{})
		model["flow"] = []map[string]interface{}{alertsV2FlowConditionModel}

		assert.Equal(t, result, model)
	}

	alertsV1FlowAlertModel := new(logsv0.AlertsV1FlowAlert)
	alertsV1FlowAlertModel.ID = core.StringPtr("testString")
	alertsV1FlowAlertModel.Not = core.BoolPtr(true)

	alertsV1FlowAlertsModel := new(logsv0.AlertsV1FlowAlerts)
	alertsV1FlowAlertsModel.Op = core.StringPtr("and")
	alertsV1FlowAlertsModel.Values = []logsv0.AlertsV1FlowAlert{*alertsV1FlowAlertModel}

	alertsV1FlowGroupModel := new(logsv0.AlertsV1FlowGroup)
	alertsV1FlowGroupModel.Alerts = alertsV1FlowAlertsModel
	alertsV1FlowGroupModel.NextOp = core.StringPtr("and")

	alertsV1FlowTimeframeModel := new(logsv0.AlertsV1FlowTimeframe)
	alertsV1FlowTimeframeModel.Ms = core.Int64Ptr(int64(0))

	alertsV1FlowStageModel := new(logsv0.AlertsV1FlowStage)
	alertsV1FlowStageModel.Groups = []logsv0.AlertsV1FlowGroup{*alertsV1FlowGroupModel}
	alertsV1FlowStageModel.Timeframe = alertsV1FlowTimeframeModel

	alertsV1MetricAlertConditionParametersModel := new(logsv0.AlertsV1MetricAlertConditionParameters)
	alertsV1MetricAlertConditionParametersModel.MetricField = core.StringPtr("testString")
	alertsV1MetricAlertConditionParametersModel.MetricSource = core.StringPtr("logs2metrics_or_unspecified")
	alertsV1MetricAlertConditionParametersModel.ArithmeticOperator = core.StringPtr("avg_or_unspecified")
	alertsV1MetricAlertConditionParametersModel.ArithmeticOperatorModifier = core.Int64Ptr(int64(0))
	alertsV1MetricAlertConditionParametersModel.SampleThresholdPercentage = core.Int64Ptr(int64(0))
	alertsV1MetricAlertConditionParametersModel.NonNullPercentage = core.Int64Ptr(int64(0))
	alertsV1MetricAlertConditionParametersModel.SwapNullValues = core.BoolPtr(true)

	alertsV1MetricAlertPromqlConditionParametersModel := new(logsv0.AlertsV1MetricAlertPromqlConditionParameters)
	alertsV1MetricAlertPromqlConditionParametersModel.PromqlText = core.StringPtr("testString")
	alertsV1MetricAlertPromqlConditionParametersModel.ArithmeticOperatorModifier = core.Int64Ptr(int64(0))
	alertsV1MetricAlertPromqlConditionParametersModel.SampleThresholdPercentage = core.Int64Ptr(int64(0))
	alertsV1MetricAlertPromqlConditionParametersModel.NonNullPercentage = core.Int64Ptr(int64(0))
	alertsV1MetricAlertPromqlConditionParametersModel.SwapNullValues = core.BoolPtr(true)

	alertsV1RelatedExtendedDataModel := new(logsv0.AlertsV1RelatedExtendedData)
	alertsV1RelatedExtendedDataModel.CleanupDeadmanDuration = core.StringPtr("cleanup_deadman_duration_never_or_unspecified")
	alertsV1RelatedExtendedDataModel.ShouldTriggerDeadman = core.BoolPtr(true)

	alertsV2ConditionParametersModel := new(logsv0.AlertsV2ConditionParameters)
	alertsV2ConditionParametersModel.Threshold = core.Float64Ptr(float64(72.5))
	alertsV2ConditionParametersModel.Timeframe = core.StringPtr("timeframe_5_min_or_unspecified")
	alertsV2ConditionParametersModel.GroupBy = []string{"testString"}
	alertsV2ConditionParametersModel.MetricAlertParameters = alertsV1MetricAlertConditionParametersModel
	alertsV2ConditionParametersModel.MetricAlertPromqlParameters = alertsV1MetricAlertPromqlConditionParametersModel
	alertsV2ConditionParametersModel.IgnoreInfinity = core.BoolPtr(true)
	alertsV2ConditionParametersModel.RelativeTimeframe = core.StringPtr("hour_or_unspecified")
	alertsV2ConditionParametersModel.CardinalityFields = []string{"testString"}
	alertsV2ConditionParametersModel.RelatedExtendedData = alertsV1RelatedExtendedDataModel

	alertsV2FlowConditionModel := new(logsv0.AlertsV2FlowCondition)
	alertsV2FlowConditionModel.Stages = []logsv0.AlertsV1FlowStage{*alertsV1FlowStageModel}
	alertsV2FlowConditionModel.Parameters = alertsV2ConditionParametersModel
	alertsV2FlowConditionModel.EnforceSuppression = core.BoolPtr(true)

	model := new(logsv0.AlertsV2AlertConditionConditionFlow)
	model.Flow = alertsV2FlowConditionModel

	result, err := logs.DataSourceIbmLogsAlertsAlertsV2AlertConditionConditionFlowToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmLogsAlertsAlertsV2AlertConditionConditionUniqueCountToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		alertsV1MetricAlertConditionParametersModel := make(map[string]interface{})
		alertsV1MetricAlertConditionParametersModel["metric_field"] = "testString"
		alertsV1MetricAlertConditionParametersModel["metric_source"] = "logs2metrics_or_unspecified"
		alertsV1MetricAlertConditionParametersModel["arithmetic_operator"] = "avg_or_unspecified"
		alertsV1MetricAlertConditionParametersModel["arithmetic_operator_modifier"] = int(0)
		alertsV1MetricAlertConditionParametersModel["sample_threshold_percentage"] = int(0)
		alertsV1MetricAlertConditionParametersModel["non_null_percentage"] = int(0)
		alertsV1MetricAlertConditionParametersModel["swap_null_values"] = true

		alertsV1MetricAlertPromqlConditionParametersModel := make(map[string]interface{})
		alertsV1MetricAlertPromqlConditionParametersModel["promql_text"] = "testString"
		alertsV1MetricAlertPromqlConditionParametersModel["arithmetic_operator_modifier"] = int(0)
		alertsV1MetricAlertPromqlConditionParametersModel["sample_threshold_percentage"] = int(0)
		alertsV1MetricAlertPromqlConditionParametersModel["non_null_percentage"] = int(0)
		alertsV1MetricAlertPromqlConditionParametersModel["swap_null_values"] = true

		alertsV1RelatedExtendedDataModel := make(map[string]interface{})
		alertsV1RelatedExtendedDataModel["cleanup_deadman_duration"] = "cleanup_deadman_duration_never_or_unspecified"
		alertsV1RelatedExtendedDataModel["should_trigger_deadman"] = true

		alertsV2ConditionParametersModel := make(map[string]interface{})
		alertsV2ConditionParametersModel["threshold"] = float64(72.5)
		alertsV2ConditionParametersModel["timeframe"] = "timeframe_5_min_or_unspecified"
		alertsV2ConditionParametersModel["group_by"] = []string{"testString"}
		alertsV2ConditionParametersModel["metric_alert_parameters"] = []map[string]interface{}{alertsV1MetricAlertConditionParametersModel}
		alertsV2ConditionParametersModel["metric_alert_promql_parameters"] = []map[string]interface{}{alertsV1MetricAlertPromqlConditionParametersModel}
		alertsV2ConditionParametersModel["ignore_infinity"] = true
		alertsV2ConditionParametersModel["relative_timeframe"] = "hour_or_unspecified"
		alertsV2ConditionParametersModel["cardinality_fields"] = []string{"testString"}
		alertsV2ConditionParametersModel["related_extended_data"] = []map[string]interface{}{alertsV1RelatedExtendedDataModel}

		alertsV2UniqueCountConditionModel := make(map[string]interface{})
		alertsV2UniqueCountConditionModel["parameters"] = []map[string]interface{}{alertsV2ConditionParametersModel}

		model := make(map[string]interface{})
		model["unique_count"] = []map[string]interface{}{alertsV2UniqueCountConditionModel}

		assert.Equal(t, result, model)
	}

	alertsV1MetricAlertConditionParametersModel := new(logsv0.AlertsV1MetricAlertConditionParameters)
	alertsV1MetricAlertConditionParametersModel.MetricField = core.StringPtr("testString")
	alertsV1MetricAlertConditionParametersModel.MetricSource = core.StringPtr("logs2metrics_or_unspecified")
	alertsV1MetricAlertConditionParametersModel.ArithmeticOperator = core.StringPtr("avg_or_unspecified")
	alertsV1MetricAlertConditionParametersModel.ArithmeticOperatorModifier = core.Int64Ptr(int64(0))
	alertsV1MetricAlertConditionParametersModel.SampleThresholdPercentage = core.Int64Ptr(int64(0))
	alertsV1MetricAlertConditionParametersModel.NonNullPercentage = core.Int64Ptr(int64(0))
	alertsV1MetricAlertConditionParametersModel.SwapNullValues = core.BoolPtr(true)

	alertsV1MetricAlertPromqlConditionParametersModel := new(logsv0.AlertsV1MetricAlertPromqlConditionParameters)
	alertsV1MetricAlertPromqlConditionParametersModel.PromqlText = core.StringPtr("testString")
	alertsV1MetricAlertPromqlConditionParametersModel.ArithmeticOperatorModifier = core.Int64Ptr(int64(0))
	alertsV1MetricAlertPromqlConditionParametersModel.SampleThresholdPercentage = core.Int64Ptr(int64(0))
	alertsV1MetricAlertPromqlConditionParametersModel.NonNullPercentage = core.Int64Ptr(int64(0))
	alertsV1MetricAlertPromqlConditionParametersModel.SwapNullValues = core.BoolPtr(true)

	alertsV1RelatedExtendedDataModel := new(logsv0.AlertsV1RelatedExtendedData)
	alertsV1RelatedExtendedDataModel.CleanupDeadmanDuration = core.StringPtr("cleanup_deadman_duration_never_or_unspecified")
	alertsV1RelatedExtendedDataModel.ShouldTriggerDeadman = core.BoolPtr(true)

	alertsV2ConditionParametersModel := new(logsv0.AlertsV2ConditionParameters)
	alertsV2ConditionParametersModel.Threshold = core.Float64Ptr(float64(72.5))
	alertsV2ConditionParametersModel.Timeframe = core.StringPtr("timeframe_5_min_or_unspecified")
	alertsV2ConditionParametersModel.GroupBy = []string{"testString"}
	alertsV2ConditionParametersModel.MetricAlertParameters = alertsV1MetricAlertConditionParametersModel
	alertsV2ConditionParametersModel.MetricAlertPromqlParameters = alertsV1MetricAlertPromqlConditionParametersModel
	alertsV2ConditionParametersModel.IgnoreInfinity = core.BoolPtr(true)
	alertsV2ConditionParametersModel.RelativeTimeframe = core.StringPtr("hour_or_unspecified")
	alertsV2ConditionParametersModel.CardinalityFields = []string{"testString"}
	alertsV2ConditionParametersModel.RelatedExtendedData = alertsV1RelatedExtendedDataModel

	alertsV2UniqueCountConditionModel := new(logsv0.AlertsV2UniqueCountCondition)
	alertsV2UniqueCountConditionModel.Parameters = alertsV2ConditionParametersModel

	model := new(logsv0.AlertsV2AlertConditionConditionUniqueCount)
	model.UniqueCount = alertsV2UniqueCountConditionModel

	result, err := logs.DataSourceIbmLogsAlertsAlertsV2AlertConditionConditionUniqueCountToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmLogsAlertsAlertsV2AlertConditionConditionLessThanUsualToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		alertsV1MetricAlertConditionParametersModel := make(map[string]interface{})
		alertsV1MetricAlertConditionParametersModel["metric_field"] = "testString"
		alertsV1MetricAlertConditionParametersModel["metric_source"] = "logs2metrics_or_unspecified"
		alertsV1MetricAlertConditionParametersModel["arithmetic_operator"] = "avg_or_unspecified"
		alertsV1MetricAlertConditionParametersModel["arithmetic_operator_modifier"] = int(0)
		alertsV1MetricAlertConditionParametersModel["sample_threshold_percentage"] = int(0)
		alertsV1MetricAlertConditionParametersModel["non_null_percentage"] = int(0)
		alertsV1MetricAlertConditionParametersModel["swap_null_values"] = true

		alertsV1MetricAlertPromqlConditionParametersModel := make(map[string]interface{})
		alertsV1MetricAlertPromqlConditionParametersModel["promql_text"] = "testString"
		alertsV1MetricAlertPromqlConditionParametersModel["arithmetic_operator_modifier"] = int(0)
		alertsV1MetricAlertPromqlConditionParametersModel["sample_threshold_percentage"] = int(0)
		alertsV1MetricAlertPromqlConditionParametersModel["non_null_percentage"] = int(0)
		alertsV1MetricAlertPromqlConditionParametersModel["swap_null_values"] = true

		alertsV1RelatedExtendedDataModel := make(map[string]interface{})
		alertsV1RelatedExtendedDataModel["cleanup_deadman_duration"] = "cleanup_deadman_duration_never_or_unspecified"
		alertsV1RelatedExtendedDataModel["should_trigger_deadman"] = true

		alertsV2ConditionParametersModel := make(map[string]interface{})
		alertsV2ConditionParametersModel["threshold"] = float64(72.5)
		alertsV2ConditionParametersModel["timeframe"] = "timeframe_5_min_or_unspecified"
		alertsV2ConditionParametersModel["group_by"] = []string{"testString"}
		alertsV2ConditionParametersModel["metric_alert_parameters"] = []map[string]interface{}{alertsV1MetricAlertConditionParametersModel}
		alertsV2ConditionParametersModel["metric_alert_promql_parameters"] = []map[string]interface{}{alertsV1MetricAlertPromqlConditionParametersModel}
		alertsV2ConditionParametersModel["ignore_infinity"] = true
		alertsV2ConditionParametersModel["relative_timeframe"] = "hour_or_unspecified"
		alertsV2ConditionParametersModel["cardinality_fields"] = []string{"testString"}
		alertsV2ConditionParametersModel["related_extended_data"] = []map[string]interface{}{alertsV1RelatedExtendedDataModel}

		alertsV2LessThanUsualConditionModel := make(map[string]interface{})
		alertsV2LessThanUsualConditionModel["parameters"] = []map[string]interface{}{alertsV2ConditionParametersModel}

		model := make(map[string]interface{})
		model["less_than_usual"] = []map[string]interface{}{alertsV2LessThanUsualConditionModel}

		assert.Equal(t, result, model)
	}

	alertsV1MetricAlertConditionParametersModel := new(logsv0.AlertsV1MetricAlertConditionParameters)
	alertsV1MetricAlertConditionParametersModel.MetricField = core.StringPtr("testString")
	alertsV1MetricAlertConditionParametersModel.MetricSource = core.StringPtr("logs2metrics_or_unspecified")
	alertsV1MetricAlertConditionParametersModel.ArithmeticOperator = core.StringPtr("avg_or_unspecified")
	alertsV1MetricAlertConditionParametersModel.ArithmeticOperatorModifier = core.Int64Ptr(int64(0))
	alertsV1MetricAlertConditionParametersModel.SampleThresholdPercentage = core.Int64Ptr(int64(0))
	alertsV1MetricAlertConditionParametersModel.NonNullPercentage = core.Int64Ptr(int64(0))
	alertsV1MetricAlertConditionParametersModel.SwapNullValues = core.BoolPtr(true)

	alertsV1MetricAlertPromqlConditionParametersModel := new(logsv0.AlertsV1MetricAlertPromqlConditionParameters)
	alertsV1MetricAlertPromqlConditionParametersModel.PromqlText = core.StringPtr("testString")
	alertsV1MetricAlertPromqlConditionParametersModel.ArithmeticOperatorModifier = core.Int64Ptr(int64(0))
	alertsV1MetricAlertPromqlConditionParametersModel.SampleThresholdPercentage = core.Int64Ptr(int64(0))
	alertsV1MetricAlertPromqlConditionParametersModel.NonNullPercentage = core.Int64Ptr(int64(0))
	alertsV1MetricAlertPromqlConditionParametersModel.SwapNullValues = core.BoolPtr(true)

	alertsV1RelatedExtendedDataModel := new(logsv0.AlertsV1RelatedExtendedData)
	alertsV1RelatedExtendedDataModel.CleanupDeadmanDuration = core.StringPtr("cleanup_deadman_duration_never_or_unspecified")
	alertsV1RelatedExtendedDataModel.ShouldTriggerDeadman = core.BoolPtr(true)

	alertsV2ConditionParametersModel := new(logsv0.AlertsV2ConditionParameters)
	alertsV2ConditionParametersModel.Threshold = core.Float64Ptr(float64(72.5))
	alertsV2ConditionParametersModel.Timeframe = core.StringPtr("timeframe_5_min_or_unspecified")
	alertsV2ConditionParametersModel.GroupBy = []string{"testString"}
	alertsV2ConditionParametersModel.MetricAlertParameters = alertsV1MetricAlertConditionParametersModel
	alertsV2ConditionParametersModel.MetricAlertPromqlParameters = alertsV1MetricAlertPromqlConditionParametersModel
	alertsV2ConditionParametersModel.IgnoreInfinity = core.BoolPtr(true)
	alertsV2ConditionParametersModel.RelativeTimeframe = core.StringPtr("hour_or_unspecified")
	alertsV2ConditionParametersModel.CardinalityFields = []string{"testString"}
	alertsV2ConditionParametersModel.RelatedExtendedData = alertsV1RelatedExtendedDataModel

	alertsV2LessThanUsualConditionModel := new(logsv0.AlertsV2LessThanUsualCondition)
	alertsV2LessThanUsualConditionModel.Parameters = alertsV2ConditionParametersModel

	model := new(logsv0.AlertsV2AlertConditionConditionLessThanUsual)
	model.LessThanUsual = alertsV2LessThanUsualConditionModel

	result, err := logs.DataSourceIbmLogsAlertsAlertsV2AlertConditionConditionLessThanUsualToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmLogsAlertsAlertsV2AlertNotificationGroupsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		alertsV2AlertNotificationModel := make(map[string]interface{})
		alertsV2AlertNotificationModel["retriggering_period_seconds"] = int(0)
		alertsV2AlertNotificationModel["notify_on"] = "triggered_only"
		alertsV2AlertNotificationModel["integration_id"] = int(0)

		model := make(map[string]interface{})
		model["group_by_fields"] = []string{"testString"}
		model["notifications"] = []map[string]interface{}{alertsV2AlertNotificationModel}

		assert.Equal(t, result, model)
	}

	alertsV2AlertNotificationModel := new(logsv0.AlertsV2AlertNotificationIntegrationTypeIntegrationID)
	alertsV2AlertNotificationModel.RetriggeringPeriodSeconds = core.Int64Ptr(int64(0))
	alertsV2AlertNotificationModel.NotifyOn = core.StringPtr("triggered_only")
	alertsV2AlertNotificationModel.IntegrationID = core.Int64Ptr(int64(0))

	model := new(logsv0.AlertsV2AlertNotificationGroups)
	model.GroupByFields = []string{"testString"}
	model.Notifications = []logsv0.AlertsV2AlertNotificationIntf{alertsV2AlertNotificationModel}

	result, err := logs.DataSourceIbmLogsAlertsAlertsV2AlertNotificationGroupsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

// Todo @kavya498: Fix unit testcases
// func TestDataSourceIbmLogsAlertsAlertsV2AlertNotificationToMap(t *testing.T) {
// 	checkResult := func(result map[string]interface{}) {
// 		model := make(map[string]interface{})
// 		model["retriggering_period_seconds"] = int(0)
// 		model["notify_on"] = "triggered_only"
// 		model["integration_id"] = int(0)
// 		model["recipients"] = []map[string]interface{}{alertsV2RecipientsModel}

// 		assert.Equal(t, result, model)
// 	}

// 	model := new(logsv0.AlertsV2AlertNotification)
// 	model.RetriggeringPeriodSeconds = core.Int64Ptr(int64(0))
// 	model.NotifyOn = core.StringPtr("triggered_only")
// 	model.IntegrationID = core.Int64Ptr(int64(0))
// 	model.Recipients = alertsV2RecipientsModel

// 	result, err := logs.DataSourceIbmLogsAlertsAlertsV2AlertNotificationToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

func TestDataSourceIbmLogsAlertsAlertsV2RecipientsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["emails"] = []string{"testString"}

		assert.Equal(t, result, model)
	}

	model := new(logsv0.AlertsV2Recipients)
	model.Emails = []string{"testString"}

	result, err := logs.DataSourceIbmLogsAlertsAlertsV2RecipientsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmLogsAlertsAlertsV2AlertNotificationIntegrationTypeIntegrationIDToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["retriggering_period_seconds"] = int(0)
		model["notify_on"] = "triggered_only"
		model["integration_id"] = int(0)

		assert.Equal(t, result, model)
	}

	model := new(logsv0.AlertsV2AlertNotificationIntegrationTypeIntegrationID)
	model.RetriggeringPeriodSeconds = core.Int64Ptr(int64(0))
	model.NotifyOn = core.StringPtr("triggered_only")
	model.IntegrationID = core.Int64Ptr(int64(0))

	result, err := logs.DataSourceIbmLogsAlertsAlertsV2AlertNotificationIntegrationTypeIntegrationIDToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmLogsAlertsAlertsV2AlertNotificationIntegrationTypeRecipientsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		alertsV2RecipientsModel := make(map[string]interface{})
		alertsV2RecipientsModel["emails"] = []string{"testString"}

		model := make(map[string]interface{})
		model["retriggering_period_seconds"] = int(0)
		model["notify_on"] = "triggered_only"
		model["recipients"] = []map[string]interface{}{alertsV2RecipientsModel}

		assert.Equal(t, result, model)
	}

	alertsV2RecipientsModel := new(logsv0.AlertsV2Recipients)
	alertsV2RecipientsModel.Emails = []string{"testString"}

	model := new(logsv0.AlertsV2AlertNotificationIntegrationTypeRecipients)
	model.RetriggeringPeriodSeconds = core.Int64Ptr(int64(0))
	model.NotifyOn = core.StringPtr("triggered_only")
	model.Recipients = alertsV2RecipientsModel

	result, err := logs.DataSourceIbmLogsAlertsAlertsV2AlertNotificationIntegrationTypeRecipientsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmLogsAlertsAlertsV1AlertFiltersToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		alertsV1AlertFiltersMetadataFiltersModel := make(map[string]interface{})
		alertsV1AlertFiltersMetadataFiltersModel["categories"] = []string{"testString"}
		alertsV1AlertFiltersMetadataFiltersModel["applications"] = []string{"testString"}
		alertsV1AlertFiltersMetadataFiltersModel["subsystems"] = []string{"testString"}
		alertsV1AlertFiltersMetadataFiltersModel["computers"] = []string{"testString"}
		alertsV1AlertFiltersMetadataFiltersModel["classes"] = []string{"testString"}
		alertsV1AlertFiltersMetadataFiltersModel["methods"] = []string{"testString"}
		alertsV1AlertFiltersMetadataFiltersModel["ip_addresses"] = []string{"testString"}

		alertsV1AlertFiltersRatioAlertModel := make(map[string]interface{})
		alertsV1AlertFiltersRatioAlertModel["alias"] = "testString"
		alertsV1AlertFiltersRatioAlertModel["text"] = "testString"
		alertsV1AlertFiltersRatioAlertModel["severities"] = []string{"debug_or_unspecified"}
		alertsV1AlertFiltersRatioAlertModel["applications"] = []string{"testString"}
		alertsV1AlertFiltersRatioAlertModel["subsystems"] = []string{"testString"}
		alertsV1AlertFiltersRatioAlertModel["group_by"] = []string{"testString"}

		model := make(map[string]interface{})
		model["severities"] = []string{"debug_or_unspecified"}
		model["metadata"] = []map[string]interface{}{alertsV1AlertFiltersMetadataFiltersModel}
		model["alias"] = "testString"
		model["text"] = "testString"
		model["ratio_alerts"] = []map[string]interface{}{alertsV1AlertFiltersRatioAlertModel}
		model["filter_type"] = "text_or_unspecified"

		assert.Equal(t, result, model)
	}

	alertsV1AlertFiltersMetadataFiltersModel := new(logsv0.AlertsV1AlertFiltersMetadataFilters)
	alertsV1AlertFiltersMetadataFiltersModel.Categories = []string{"testString"}
	alertsV1AlertFiltersMetadataFiltersModel.Applications = []string{"testString"}
	alertsV1AlertFiltersMetadataFiltersModel.Subsystems = []string{"testString"}
	alertsV1AlertFiltersMetadataFiltersModel.Computers = []string{"testString"}
	alertsV1AlertFiltersMetadataFiltersModel.Classes = []string{"testString"}
	alertsV1AlertFiltersMetadataFiltersModel.Methods = []string{"testString"}
	alertsV1AlertFiltersMetadataFiltersModel.IpAddresses = []string{"testString"}

	alertsV1AlertFiltersRatioAlertModel := new(logsv0.AlertsV1AlertFiltersRatioAlert)
	alertsV1AlertFiltersRatioAlertModel.Alias = core.StringPtr("testString")
	alertsV1AlertFiltersRatioAlertModel.Text = core.StringPtr("testString")
	alertsV1AlertFiltersRatioAlertModel.Severities = []string{"debug_or_unspecified"}
	alertsV1AlertFiltersRatioAlertModel.Applications = []string{"testString"}
	alertsV1AlertFiltersRatioAlertModel.Subsystems = []string{"testString"}
	alertsV1AlertFiltersRatioAlertModel.GroupBy = []string{"testString"}

	model := new(logsv0.AlertsV1AlertFilters)
	model.Severities = []string{"debug_or_unspecified"}
	model.Metadata = alertsV1AlertFiltersMetadataFiltersModel
	model.Alias = core.StringPtr("testString")
	model.Text = core.StringPtr("testString")
	model.RatioAlerts = []logsv0.AlertsV1AlertFiltersRatioAlert{*alertsV1AlertFiltersRatioAlertModel}
	model.FilterType = core.StringPtr("text_or_unspecified")

	result, err := logs.DataSourceIbmLogsAlertsAlertsV1AlertFiltersToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmLogsAlertsAlertsV1AlertFiltersMetadataFiltersToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["categories"] = []string{"testString"}
		model["applications"] = []string{"testString"}
		model["subsystems"] = []string{"testString"}
		model["computers"] = []string{"testString"}
		model["classes"] = []string{"testString"}
		model["methods"] = []string{"testString"}
		model["ip_addresses"] = []string{"testString"}

		assert.Equal(t, result, model)
	}

	model := new(logsv0.AlertsV1AlertFiltersMetadataFilters)
	model.Categories = []string{"testString"}
	model.Applications = []string{"testString"}
	model.Subsystems = []string{"testString"}
	model.Computers = []string{"testString"}
	model.Classes = []string{"testString"}
	model.Methods = []string{"testString"}
	model.IpAddresses = []string{"testString"}

	result, err := logs.DataSourceIbmLogsAlertsAlertsV1AlertFiltersMetadataFiltersToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmLogsAlertsAlertsV1AlertFiltersRatioAlertToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["alias"] = "testString"
		model["text"] = "testString"
		model["severities"] = []string{"debug_or_unspecified"}
		model["applications"] = []string{"testString"}
		model["subsystems"] = []string{"testString"}
		model["group_by"] = []string{"testString"}

		assert.Equal(t, result, model)
	}

	model := new(logsv0.AlertsV1AlertFiltersRatioAlert)
	model.Alias = core.StringPtr("testString")
	model.Text = core.StringPtr("testString")
	model.Severities = []string{"debug_or_unspecified"}
	model.Applications = []string{"testString"}
	model.Subsystems = []string{"testString"}
	model.GroupBy = []string{"testString"}

	result, err := logs.DataSourceIbmLogsAlertsAlertsV1AlertFiltersRatioAlertToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmLogsAlertsAlertsV1AlertActiveWhenToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		alertsV1TimeModel := make(map[string]interface{})
		alertsV1TimeModel["hours"] = int(38)
		alertsV1TimeModel["minutes"] = int(38)
		alertsV1TimeModel["seconds"] = int(38)

		alertsV1TimeRangeModel := make(map[string]interface{})
		alertsV1TimeRangeModel["start"] = []map[string]interface{}{alertsV1TimeModel}
		alertsV1TimeRangeModel["end"] = []map[string]interface{}{alertsV1TimeModel}

		alertsV1AlertActiveTimeframeModel := make(map[string]interface{})
		alertsV1AlertActiveTimeframeModel["days_of_week"] = []string{"monday_or_unspecified"}
		alertsV1AlertActiveTimeframeModel["range"] = []map[string]interface{}{alertsV1TimeRangeModel}

		model := make(map[string]interface{})
		model["timeframes"] = []map[string]interface{}{alertsV1AlertActiveTimeframeModel}

		assert.Equal(t, result, model)
	}

	alertsV1TimeModel := new(logsv0.AlertsV1Time)
	alertsV1TimeModel.Hours = core.Int64Ptr(int64(38))
	alertsV1TimeModel.Minutes = core.Int64Ptr(int64(38))
	alertsV1TimeModel.Seconds = core.Int64Ptr(int64(38))

	alertsV1TimeRangeModel := new(logsv0.AlertsV1TimeRange)
	alertsV1TimeRangeModel.Start = alertsV1TimeModel
	alertsV1TimeRangeModel.End = alertsV1TimeModel

	alertsV1AlertActiveTimeframeModel := new(logsv0.AlertsV1AlertActiveTimeframe)
	alertsV1AlertActiveTimeframeModel.DaysOfWeek = []string{"monday_or_unspecified"}
	alertsV1AlertActiveTimeframeModel.Range = alertsV1TimeRangeModel

	model := new(logsv0.AlertsV1AlertActiveWhen)
	model.Timeframes = []logsv0.AlertsV1AlertActiveTimeframe{*alertsV1AlertActiveTimeframeModel}

	result, err := logs.DataSourceIbmLogsAlertsAlertsV1AlertActiveWhenToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmLogsAlertsAlertsV1AlertActiveTimeframeToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		alertsV1TimeModel := make(map[string]interface{})
		alertsV1TimeModel["hours"] = int(38)
		alertsV1TimeModel["minutes"] = int(38)
		alertsV1TimeModel["seconds"] = int(38)

		alertsV1TimeRangeModel := make(map[string]interface{})
		alertsV1TimeRangeModel["start"] = []map[string]interface{}{alertsV1TimeModel}
		alertsV1TimeRangeModel["end"] = []map[string]interface{}{alertsV1TimeModel}

		model := make(map[string]interface{})
		model["days_of_week"] = []string{"monday_or_unspecified"}
		model["range"] = []map[string]interface{}{alertsV1TimeRangeModel}

		assert.Equal(t, result, model)
	}

	alertsV1TimeModel := new(logsv0.AlertsV1Time)
	alertsV1TimeModel.Hours = core.Int64Ptr(int64(38))
	alertsV1TimeModel.Minutes = core.Int64Ptr(int64(38))
	alertsV1TimeModel.Seconds = core.Int64Ptr(int64(38))

	alertsV1TimeRangeModel := new(logsv0.AlertsV1TimeRange)
	alertsV1TimeRangeModel.Start = alertsV1TimeModel
	alertsV1TimeRangeModel.End = alertsV1TimeModel

	model := new(logsv0.AlertsV1AlertActiveTimeframe)
	model.DaysOfWeek = []string{"monday_or_unspecified"}
	model.Range = alertsV1TimeRangeModel

	result, err := logs.DataSourceIbmLogsAlertsAlertsV1AlertActiveTimeframeToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmLogsAlertsAlertsV1TimeRangeToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		alertsV1TimeModel := make(map[string]interface{})
		alertsV1TimeModel["hours"] = int(38)
		alertsV1TimeModel["minutes"] = int(38)
		alertsV1TimeModel["seconds"] = int(38)

		model := make(map[string]interface{})
		model["start"] = []map[string]interface{}{alertsV1TimeModel}
		model["end"] = []map[string]interface{}{alertsV1TimeModel}

		assert.Equal(t, result, model)
	}

	alertsV1TimeModel := new(logsv0.AlertsV1Time)
	alertsV1TimeModel.Hours = core.Int64Ptr(int64(38))
	alertsV1TimeModel.Minutes = core.Int64Ptr(int64(38))
	alertsV1TimeModel.Seconds = core.Int64Ptr(int64(38))

	model := new(logsv0.AlertsV1TimeRange)
	model.Start = alertsV1TimeModel
	model.End = alertsV1TimeModel

	result, err := logs.DataSourceIbmLogsAlertsAlertsV1TimeRangeToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmLogsAlertsAlertsV1TimeToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["hours"] = int(38)
		model["minutes"] = int(38)
		model["seconds"] = int(38)

		assert.Equal(t, result, model)
	}

	model := new(logsv0.AlertsV1Time)
	model.Hours = core.Int64Ptr(int64(38))
	model.Minutes = core.Int64Ptr(int64(38))
	model.Seconds = core.Int64Ptr(int64(38))

	result, err := logs.DataSourceIbmLogsAlertsAlertsV1TimeToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmLogsAlertsAlertsV1MetaLabelToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["key"] = "testString"
		model["value"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(logsv0.AlertsV1MetaLabel)
	model.Key = core.StringPtr("testString")
	model.Value = core.StringPtr("testString")

	result, err := logs.DataSourceIbmLogsAlertsAlertsV1MetaLabelToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmLogsAlertsAlertsV1TracingAlertToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		alertsV1FiltersModel := make(map[string]interface{})
		alertsV1FiltersModel["values"] = []string{"testString"}
		alertsV1FiltersModel["operator"] = "testString"

		alertsV1FilterDataModel := make(map[string]interface{})
		alertsV1FilterDataModel["field"] = "testString"
		alertsV1FilterDataModel["filters"] = []map[string]interface{}{alertsV1FiltersModel}

		model := make(map[string]interface{})
		model["condition_latency"] = int(0)
		model["field_filters"] = []map[string]interface{}{alertsV1FilterDataModel}
		model["tag_filters"] = []map[string]interface{}{alertsV1FilterDataModel}

		assert.Equal(t, result, model)
	}

	alertsV1FiltersModel := new(logsv0.AlertsV1Filters)
	alertsV1FiltersModel.Values = []string{"testString"}
	alertsV1FiltersModel.Operator = core.StringPtr("testString")

	alertsV1FilterDataModel := new(logsv0.AlertsV1FilterData)
	alertsV1FilterDataModel.Field = core.StringPtr("testString")
	alertsV1FilterDataModel.Filters = []logsv0.AlertsV1Filters{*alertsV1FiltersModel}

	model := new(logsv0.AlertsV1TracingAlert)
	model.ConditionLatency = core.Int64Ptr(int64(0))
	model.FieldFilters = []logsv0.AlertsV1FilterData{*alertsV1FilterDataModel}
	model.TagFilters = []logsv0.AlertsV1FilterData{*alertsV1FilterDataModel}

	result, err := logs.DataSourceIbmLogsAlertsAlertsV1TracingAlertToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmLogsAlertsAlertsV1FilterDataToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		alertsV1FiltersModel := make(map[string]interface{})
		alertsV1FiltersModel["values"] = []string{"testString"}
		alertsV1FiltersModel["operator"] = "testString"

		model := make(map[string]interface{})
		model["field"] = "testString"
		model["filters"] = []map[string]interface{}{alertsV1FiltersModel}

		assert.Equal(t, result, model)
	}

	alertsV1FiltersModel := new(logsv0.AlertsV1Filters)
	alertsV1FiltersModel.Values = []string{"testString"}
	alertsV1FiltersModel.Operator = core.StringPtr("testString")

	model := new(logsv0.AlertsV1FilterData)
	model.Field = core.StringPtr("testString")
	model.Filters = []logsv0.AlertsV1Filters{*alertsV1FiltersModel}

	result, err := logs.DataSourceIbmLogsAlertsAlertsV1FilterDataToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmLogsAlertsAlertsV1FiltersToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["values"] = []string{"testString"}
		model["operator"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(logsv0.AlertsV1Filters)
	model.Values = []string{"testString"}
	model.Operator = core.StringPtr("testString")

	result, err := logs.DataSourceIbmLogsAlertsAlertsV1FiltersToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmLogsAlertsAlertsV2AlertIncidentSettingsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["retriggering_period_seconds"] = int(0)
		model["notify_on"] = "triggered_only"
		model["use_as_notification_settings"] = true

		assert.Equal(t, result, model)
	}

	model := new(logsv0.AlertsV2AlertIncidentSettings)
	model.RetriggeringPeriodSeconds = core.Int64Ptr(int64(0))
	model.NotifyOn = core.StringPtr("triggered_only")
	model.UseAsNotificationSettings = core.BoolPtr(true)

	result, err := logs.DataSourceIbmLogsAlertsAlertsV2AlertIncidentSettingsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}
