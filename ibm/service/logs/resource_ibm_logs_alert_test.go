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

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/logs"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/observability-c/dragonlog-logs-go-sdk/logsv0"
	"github.com/stretchr/testify/assert"
)

func TestAccIbmLogsAlertBasic(t *testing.T) {
	var conf logsv0.Alert
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
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
		meta_labels_strings = []
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
		  meta_labels_strings = []
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

func TestResourceIbmLogsAlertAlertsV1DateToMap(t *testing.T) {
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

	result, err := logs.ResourceIbmLogsAlertAlertsV1DateToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

// Todo @kavya498: Fix unit testcases
// func TestResourceIbmLogsAlertAlertsV2AlertConditionToMap(t *testing.T) {
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

// 	result, err := logs.ResourceIbmLogsAlertAlertsV2AlertConditionToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

func TestResourceIbmLogsAlertAlertsV2ImmediateConditionToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})

		assert.Equal(t, result, model)
	}

	model := new(logsv0.AlertsV2ImmediateCondition)

	result, err := logs.ResourceIbmLogsAlertAlertsV2ImmediateConditionToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertAlertsV2LessThanConditionToMap(t *testing.T) {
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

	result, err := logs.ResourceIbmLogsAlertAlertsV2LessThanConditionToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertAlertsV2ConditionParametersToMap(t *testing.T) {
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

	result, err := logs.ResourceIbmLogsAlertAlertsV2ConditionParametersToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertAlertsV1MetricAlertConditionParametersToMap(t *testing.T) {
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

	result, err := logs.ResourceIbmLogsAlertAlertsV1MetricAlertConditionParametersToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertAlertsV1MetricAlertPromqlConditionParametersToMap(t *testing.T) {
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

	result, err := logs.ResourceIbmLogsAlertAlertsV1MetricAlertPromqlConditionParametersToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertAlertsV1RelatedExtendedDataToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["cleanup_deadman_duration"] = "cleanup_deadman_duration_never_or_unspecified"
		model["should_trigger_deadman"] = true

		assert.Equal(t, result, model)
	}

	model := new(logsv0.AlertsV1RelatedExtendedData)
	model.CleanupDeadmanDuration = core.StringPtr("cleanup_deadman_duration_never_or_unspecified")
	model.ShouldTriggerDeadman = core.BoolPtr(true)

	result, err := logs.ResourceIbmLogsAlertAlertsV1RelatedExtendedDataToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertAlertsV2MoreThanConditionToMap(t *testing.T) {
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

	result, err := logs.ResourceIbmLogsAlertAlertsV2MoreThanConditionToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertAlertsV2MoreThanUsualConditionToMap(t *testing.T) {
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

	result, err := logs.ResourceIbmLogsAlertAlertsV2MoreThanUsualConditionToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertAlertsV2NewValueConditionToMap(t *testing.T) {
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

	result, err := logs.ResourceIbmLogsAlertAlertsV2NewValueConditionToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertAlertsV2FlowConditionToMap(t *testing.T) {
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

	result, err := logs.ResourceIbmLogsAlertAlertsV2FlowConditionToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertAlertsV1FlowStageToMap(t *testing.T) {
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

	result, err := logs.ResourceIbmLogsAlertAlertsV1FlowStageToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertAlertsV1FlowGroupToMap(t *testing.T) {
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

	result, err := logs.ResourceIbmLogsAlertAlertsV1FlowGroupToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertAlertsV1FlowAlertsToMap(t *testing.T) {
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

	result, err := logs.ResourceIbmLogsAlertAlertsV1FlowAlertsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertAlertsV1FlowAlertToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["id"] = "testString"
		model["not"] = true

		assert.Equal(t, result, model)
	}

	model := new(logsv0.AlertsV1FlowAlert)
	model.ID = core.StringPtr("testString")
	model.Not = core.BoolPtr(true)

	result, err := logs.ResourceIbmLogsAlertAlertsV1FlowAlertToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertAlertsV1FlowTimeframeToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["ms"] = int(0)

		assert.Equal(t, result, model)
	}

	model := new(logsv0.AlertsV1FlowTimeframe)
	model.Ms = core.Int64Ptr(int64(0))

	result, err := logs.ResourceIbmLogsAlertAlertsV1FlowTimeframeToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertAlertsV2UniqueCountConditionToMap(t *testing.T) {
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

	result, err := logs.ResourceIbmLogsAlertAlertsV2UniqueCountConditionToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertAlertsV2LessThanUsualConditionToMap(t *testing.T) {
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

	result, err := logs.ResourceIbmLogsAlertAlertsV2LessThanUsualConditionToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertAlertsV2AlertConditionConditionImmediateToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		alertsV2ImmediateConditionModel := make(map[string]interface{})

		model := make(map[string]interface{})
		model["immediate"] = []map[string]interface{}{alertsV2ImmediateConditionModel}

		assert.Equal(t, result, model)
	}

	alertsV2ImmediateConditionModel := new(logsv0.AlertsV2ImmediateCondition)

	model := new(logsv0.AlertsV2AlertConditionConditionImmediate)
	model.Immediate = alertsV2ImmediateConditionModel

	result, err := logs.ResourceIbmLogsAlertAlertsV2AlertConditionConditionImmediateToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertAlertsV2AlertConditionConditionLessThanToMap(t *testing.T) {
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

	result, err := logs.ResourceIbmLogsAlertAlertsV2AlertConditionConditionLessThanToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertAlertsV2AlertConditionConditionMoreThanToMap(t *testing.T) {
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

	result, err := logs.ResourceIbmLogsAlertAlertsV2AlertConditionConditionMoreThanToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertAlertsV2AlertConditionConditionMoreThanUsualToMap(t *testing.T) {
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

	result, err := logs.ResourceIbmLogsAlertAlertsV2AlertConditionConditionMoreThanUsualToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertAlertsV2AlertConditionConditionNewValueToMap(t *testing.T) {
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

	result, err := logs.ResourceIbmLogsAlertAlertsV2AlertConditionConditionNewValueToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertAlertsV2AlertConditionConditionFlowToMap(t *testing.T) {
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

	result, err := logs.ResourceIbmLogsAlertAlertsV2AlertConditionConditionFlowToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertAlertsV2AlertConditionConditionUniqueCountToMap(t *testing.T) {
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

	result, err := logs.ResourceIbmLogsAlertAlertsV2AlertConditionConditionUniqueCountToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertAlertsV2AlertConditionConditionLessThanUsualToMap(t *testing.T) {
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

	result, err := logs.ResourceIbmLogsAlertAlertsV2AlertConditionConditionLessThanUsualToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertAlertsV2AlertNotificationGroupsToMap(t *testing.T) {
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

	result, err := logs.ResourceIbmLogsAlertAlertsV2AlertNotificationGroupsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

// Todo @kavya498: Fix unit testcases

// func TestResourceIbmLogsAlertAlertsV2AlertNotificationToMap(t *testing.T) {
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

// 	result, err := logs.ResourceIbmLogsAlertAlertsV2AlertNotificationToMap(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

func TestResourceIbmLogsAlertAlertsV2RecipientsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["emails"] = []string{"testString"}

		assert.Equal(t, result, model)
	}

	model := new(logsv0.AlertsV2Recipients)
	model.Emails = []string{"testString"}

	result, err := logs.ResourceIbmLogsAlertAlertsV2RecipientsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertAlertsV2AlertNotificationIntegrationTypeIntegrationIDToMap(t *testing.T) {
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

	result, err := logs.ResourceIbmLogsAlertAlertsV2AlertNotificationIntegrationTypeIntegrationIDToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertAlertsV2AlertNotificationIntegrationTypeRecipientsToMap(t *testing.T) {
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

	result, err := logs.ResourceIbmLogsAlertAlertsV2AlertNotificationIntegrationTypeRecipientsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertAlertsV1AlertFiltersToMap(t *testing.T) {
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

	result, err := logs.ResourceIbmLogsAlertAlertsV1AlertFiltersToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertAlertsV1AlertFiltersMetadataFiltersToMap(t *testing.T) {
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

	result, err := logs.ResourceIbmLogsAlertAlertsV1AlertFiltersMetadataFiltersToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertAlertsV1AlertFiltersRatioAlertToMap(t *testing.T) {
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

	result, err := logs.ResourceIbmLogsAlertAlertsV1AlertFiltersRatioAlertToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertAlertsV1AlertActiveWhenToMap(t *testing.T) {
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

	result, err := logs.ResourceIbmLogsAlertAlertsV1AlertActiveWhenToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertAlertsV1AlertActiveTimeframeToMap(t *testing.T) {
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

	result, err := logs.ResourceIbmLogsAlertAlertsV1AlertActiveTimeframeToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertAlertsV1TimeRangeToMap(t *testing.T) {
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

	result, err := logs.ResourceIbmLogsAlertAlertsV1TimeRangeToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertAlertsV1TimeToMap(t *testing.T) {
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

	result, err := logs.ResourceIbmLogsAlertAlertsV1TimeToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertAlertsV1MetaLabelToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["key"] = "testString"
		model["value"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(logsv0.AlertsV1MetaLabel)
	model.Key = core.StringPtr("testString")
	model.Value = core.StringPtr("testString")

	result, err := logs.ResourceIbmLogsAlertAlertsV1MetaLabelToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertAlertsV1TracingAlertToMap(t *testing.T) {
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

	result, err := logs.ResourceIbmLogsAlertAlertsV1TracingAlertToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertAlertsV1FilterDataToMap(t *testing.T) {
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

	result, err := logs.ResourceIbmLogsAlertAlertsV1FilterDataToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertAlertsV1FiltersToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["values"] = []string{"testString"}
		model["operator"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(logsv0.AlertsV1Filters)
	model.Values = []string{"testString"}
	model.Operator = core.StringPtr("testString")

	result, err := logs.ResourceIbmLogsAlertAlertsV1FiltersToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertAlertsV2AlertIncidentSettingsToMap(t *testing.T) {
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

	result, err := logs.ResourceIbmLogsAlertAlertsV2AlertIncidentSettingsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

// Todo @kavya498: Fix unit testcases
// func TestResourceIbmLogsAlertMapToAlertsV2AlertCondition(t *testing.T) {
// 	checkResult := func(result logsv0.AlertsV2AlertConditionIntf) {
// 		alertsV2ImmediateConditionModel := new(logsv0.AlertsV2ImmediateCondition)

// 		model := new(logsv0.AlertsV2AlertCondition)
// 		model.Immediate = alertsV2ImmediateConditionModel
// 		model.LessThan = alertsV2LessThanConditionModel
// 		model.MoreThan = alertsV2MoreThanConditionModel
// 		model.MoreThanUsual = alertsV2MoreThanUsualConditionModel
// 		model.NewValue = alertsV2NewValueConditionModel
// 		model.Flow = alertsV2FlowConditionModel
// 		model.UniqueCount = alertsV2UniqueCountConditionModel
// 		model.LessThanUsual = alertsV2LessThanUsualConditionModel

// 		assert.Equal(t, result, model)
// 	}

// 	alertsV2ImmediateConditionModel := make(map[string]interface{})

// 	model := make(map[string]interface{})
// 	model["immediate"] = []interface{}{alertsV2ImmediateConditionModel}
// 	model["less_than"] = []interface{}{alertsV2LessThanConditionModel}
// 	model["more_than"] = []interface{}{alertsV2MoreThanConditionModel}
// 	model["more_than_usual"] = []interface{}{alertsV2MoreThanUsualConditionModel}
// 	model["new_value"] = []interface{}{alertsV2NewValueConditionModel}
// 	model["flow"] = []interface{}{alertsV2FlowConditionModel}
// 	model["unique_count"] = []interface{}{alertsV2UniqueCountConditionModel}
// 	model["less_than_usual"] = []interface{}{alertsV2LessThanUsualConditionModel}

// 	result, err := logs.ResourceIbmLogsAlertMapToAlertsV2AlertCondition(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

func TestResourceIbmLogsAlertMapToAlertsV2ImmediateCondition(t *testing.T) {
	checkResult := func(result *logsv0.AlertsV2ImmediateCondition) {
		model := new(logsv0.AlertsV2ImmediateCondition)

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})

	result, err := logs.ResourceIbmLogsAlertMapToAlertsV2ImmediateCondition(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertMapToAlertsV2LessThanCondition(t *testing.T) {
	checkResult := func(result *logsv0.AlertsV2LessThanCondition) {
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

		assert.Equal(t, result, model)
	}

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
	alertsV2ConditionParametersModel["group_by"] = []interface{}{"testString"}
	alertsV2ConditionParametersModel["metric_alert_parameters"] = []interface{}{alertsV1MetricAlertConditionParametersModel}
	alertsV2ConditionParametersModel["metric_alert_promql_parameters"] = []interface{}{alertsV1MetricAlertPromqlConditionParametersModel}
	alertsV2ConditionParametersModel["ignore_infinity"] = true
	alertsV2ConditionParametersModel["relative_timeframe"] = "hour_or_unspecified"
	alertsV2ConditionParametersModel["cardinality_fields"] = []interface{}{"testString"}
	alertsV2ConditionParametersModel["related_extended_data"] = []interface{}{alertsV1RelatedExtendedDataModel}

	model := make(map[string]interface{})
	model["parameters"] = []interface{}{alertsV2ConditionParametersModel}

	result, err := logs.ResourceIbmLogsAlertMapToAlertsV2LessThanCondition(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertMapToAlertsV2ConditionParameters(t *testing.T) {
	checkResult := func(result *logsv0.AlertsV2ConditionParameters) {
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

		assert.Equal(t, result, model)
	}

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
	model["group_by"] = []interface{}{"testString"}
	model["metric_alert_parameters"] = []interface{}{alertsV1MetricAlertConditionParametersModel}
	model["metric_alert_promql_parameters"] = []interface{}{alertsV1MetricAlertPromqlConditionParametersModel}
	model["ignore_infinity"] = true
	model["relative_timeframe"] = "hour_or_unspecified"
	model["cardinality_fields"] = []interface{}{"testString"}
	model["related_extended_data"] = []interface{}{alertsV1RelatedExtendedDataModel}

	result, err := logs.ResourceIbmLogsAlertMapToAlertsV2ConditionParameters(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertMapToAlertsV1MetricAlertConditionParameters(t *testing.T) {
	checkResult := func(result *logsv0.AlertsV1MetricAlertConditionParameters) {
		model := new(logsv0.AlertsV1MetricAlertConditionParameters)
		model.MetricField = core.StringPtr("testString")
		model.MetricSource = core.StringPtr("logs2metrics_or_unspecified")
		model.ArithmeticOperator = core.StringPtr("avg_or_unspecified")
		model.ArithmeticOperatorModifier = core.Int64Ptr(int64(0))
		model.SampleThresholdPercentage = core.Int64Ptr(int64(0))
		model.NonNullPercentage = core.Int64Ptr(int64(0))
		model.SwapNullValues = core.BoolPtr(true)

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["metric_field"] = "testString"
	model["metric_source"] = "logs2metrics_or_unspecified"
	model["arithmetic_operator"] = "avg_or_unspecified"
	model["arithmetic_operator_modifier"] = int(0)
	model["sample_threshold_percentage"] = int(0)
	model["non_null_percentage"] = int(0)
	model["swap_null_values"] = true

	result, err := logs.ResourceIbmLogsAlertMapToAlertsV1MetricAlertConditionParameters(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertMapToAlertsV1MetricAlertPromqlConditionParameters(t *testing.T) {
	checkResult := func(result *logsv0.AlertsV1MetricAlertPromqlConditionParameters) {
		model := new(logsv0.AlertsV1MetricAlertPromqlConditionParameters)
		model.PromqlText = core.StringPtr("testString")
		model.ArithmeticOperatorModifier = core.Int64Ptr(int64(0))
		model.SampleThresholdPercentage = core.Int64Ptr(int64(0))
		model.NonNullPercentage = core.Int64Ptr(int64(0))
		model.SwapNullValues = core.BoolPtr(true)

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["promql_text"] = "testString"
	model["arithmetic_operator_modifier"] = int(0)
	model["sample_threshold_percentage"] = int(0)
	model["non_null_percentage"] = int(0)
	model["swap_null_values"] = true

	result, err := logs.ResourceIbmLogsAlertMapToAlertsV1MetricAlertPromqlConditionParameters(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertMapToAlertsV1RelatedExtendedData(t *testing.T) {
	checkResult := func(result *logsv0.AlertsV1RelatedExtendedData) {
		model := new(logsv0.AlertsV1RelatedExtendedData)
		model.CleanupDeadmanDuration = core.StringPtr("cleanup_deadman_duration_never_or_unspecified")
		model.ShouldTriggerDeadman = core.BoolPtr(true)

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["cleanup_deadman_duration"] = "cleanup_deadman_duration_never_or_unspecified"
	model["should_trigger_deadman"] = true

	result, err := logs.ResourceIbmLogsAlertMapToAlertsV1RelatedExtendedData(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertMapToAlertsV2MoreThanCondition(t *testing.T) {
	checkResult := func(result *logsv0.AlertsV2MoreThanCondition) {
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

		assert.Equal(t, result, model)
	}

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
	alertsV2ConditionParametersModel["group_by"] = []interface{}{"testString"}
	alertsV2ConditionParametersModel["metric_alert_parameters"] = []interface{}{alertsV1MetricAlertConditionParametersModel}
	alertsV2ConditionParametersModel["metric_alert_promql_parameters"] = []interface{}{alertsV1MetricAlertPromqlConditionParametersModel}
	alertsV2ConditionParametersModel["ignore_infinity"] = true
	alertsV2ConditionParametersModel["relative_timeframe"] = "hour_or_unspecified"
	alertsV2ConditionParametersModel["cardinality_fields"] = []interface{}{"testString"}
	alertsV2ConditionParametersModel["related_extended_data"] = []interface{}{alertsV1RelatedExtendedDataModel}

	model := make(map[string]interface{})
	model["parameters"] = []interface{}{alertsV2ConditionParametersModel}
	model["evaluation_window"] = "rolling_or_unspecified"

	result, err := logs.ResourceIbmLogsAlertMapToAlertsV2MoreThanCondition(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertMapToAlertsV2MoreThanUsualCondition(t *testing.T) {
	checkResult := func(result *logsv0.AlertsV2MoreThanUsualCondition) {
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

		assert.Equal(t, result, model)
	}

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
	alertsV2ConditionParametersModel["group_by"] = []interface{}{"testString"}
	alertsV2ConditionParametersModel["metric_alert_parameters"] = []interface{}{alertsV1MetricAlertConditionParametersModel}
	alertsV2ConditionParametersModel["metric_alert_promql_parameters"] = []interface{}{alertsV1MetricAlertPromqlConditionParametersModel}
	alertsV2ConditionParametersModel["ignore_infinity"] = true
	alertsV2ConditionParametersModel["relative_timeframe"] = "hour_or_unspecified"
	alertsV2ConditionParametersModel["cardinality_fields"] = []interface{}{"testString"}
	alertsV2ConditionParametersModel["related_extended_data"] = []interface{}{alertsV1RelatedExtendedDataModel}

	model := make(map[string]interface{})
	model["parameters"] = []interface{}{alertsV2ConditionParametersModel}

	result, err := logs.ResourceIbmLogsAlertMapToAlertsV2MoreThanUsualCondition(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertMapToAlertsV2NewValueCondition(t *testing.T) {
	checkResult := func(result *logsv0.AlertsV2NewValueCondition) {
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

		assert.Equal(t, result, model)
	}

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
	alertsV2ConditionParametersModel["group_by"] = []interface{}{"testString"}
	alertsV2ConditionParametersModel["metric_alert_parameters"] = []interface{}{alertsV1MetricAlertConditionParametersModel}
	alertsV2ConditionParametersModel["metric_alert_promql_parameters"] = []interface{}{alertsV1MetricAlertPromqlConditionParametersModel}
	alertsV2ConditionParametersModel["ignore_infinity"] = true
	alertsV2ConditionParametersModel["relative_timeframe"] = "hour_or_unspecified"
	alertsV2ConditionParametersModel["cardinality_fields"] = []interface{}{"testString"}
	alertsV2ConditionParametersModel["related_extended_data"] = []interface{}{alertsV1RelatedExtendedDataModel}

	model := make(map[string]interface{})
	model["parameters"] = []interface{}{alertsV2ConditionParametersModel}

	result, err := logs.ResourceIbmLogsAlertMapToAlertsV2NewValueCondition(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertMapToAlertsV2FlowCondition(t *testing.T) {
	checkResult := func(result *logsv0.AlertsV2FlowCondition) {
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

		assert.Equal(t, result, model)
	}

	alertsV1FlowAlertModel := make(map[string]interface{})
	alertsV1FlowAlertModel["id"] = "testString"
	alertsV1FlowAlertModel["not"] = true

	alertsV1FlowAlertsModel := make(map[string]interface{})
	alertsV1FlowAlertsModel["op"] = "and"
	alertsV1FlowAlertsModel["values"] = []interface{}{alertsV1FlowAlertModel}

	alertsV1FlowGroupModel := make(map[string]interface{})
	alertsV1FlowGroupModel["alerts"] = []interface{}{alertsV1FlowAlertsModel}
	alertsV1FlowGroupModel["next_op"] = "and"

	alertsV1FlowTimeframeModel := make(map[string]interface{})
	alertsV1FlowTimeframeModel["ms"] = int(0)

	alertsV1FlowStageModel := make(map[string]interface{})
	alertsV1FlowStageModel["groups"] = []interface{}{alertsV1FlowGroupModel}
	alertsV1FlowStageModel["timeframe"] = []interface{}{alertsV1FlowTimeframeModel}

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
	alertsV2ConditionParametersModel["group_by"] = []interface{}{"testString"}
	alertsV2ConditionParametersModel["metric_alert_parameters"] = []interface{}{alertsV1MetricAlertConditionParametersModel}
	alertsV2ConditionParametersModel["metric_alert_promql_parameters"] = []interface{}{alertsV1MetricAlertPromqlConditionParametersModel}
	alertsV2ConditionParametersModel["ignore_infinity"] = true
	alertsV2ConditionParametersModel["relative_timeframe"] = "hour_or_unspecified"
	alertsV2ConditionParametersModel["cardinality_fields"] = []interface{}{"testString"}
	alertsV2ConditionParametersModel["related_extended_data"] = []interface{}{alertsV1RelatedExtendedDataModel}

	model := make(map[string]interface{})
	model["stages"] = []interface{}{alertsV1FlowStageModel}
	model["parameters"] = []interface{}{alertsV2ConditionParametersModel}
	model["enforce_suppression"] = true

	result, err := logs.ResourceIbmLogsAlertMapToAlertsV2FlowCondition(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertMapToAlertsV1FlowStage(t *testing.T) {
	checkResult := func(result *logsv0.AlertsV1FlowStage) {
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

		assert.Equal(t, result, model)
	}

	alertsV1FlowAlertModel := make(map[string]interface{})
	alertsV1FlowAlertModel["id"] = "testString"
	alertsV1FlowAlertModel["not"] = true

	alertsV1FlowAlertsModel := make(map[string]interface{})
	alertsV1FlowAlertsModel["op"] = "and"
	alertsV1FlowAlertsModel["values"] = []interface{}{alertsV1FlowAlertModel}

	alertsV1FlowGroupModel := make(map[string]interface{})
	alertsV1FlowGroupModel["alerts"] = []interface{}{alertsV1FlowAlertsModel}
	alertsV1FlowGroupModel["next_op"] = "and"

	alertsV1FlowTimeframeModel := make(map[string]interface{})
	alertsV1FlowTimeframeModel["ms"] = int(0)

	model := make(map[string]interface{})
	model["groups"] = []interface{}{alertsV1FlowGroupModel}
	model["timeframe"] = []interface{}{alertsV1FlowTimeframeModel}

	result, err := logs.ResourceIbmLogsAlertMapToAlertsV1FlowStage(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertMapToAlertsV1FlowGroup(t *testing.T) {
	checkResult := func(result *logsv0.AlertsV1FlowGroup) {
		alertsV1FlowAlertModel := new(logsv0.AlertsV1FlowAlert)
		alertsV1FlowAlertModel.ID = core.StringPtr("testString")
		alertsV1FlowAlertModel.Not = core.BoolPtr(true)

		alertsV1FlowAlertsModel := new(logsv0.AlertsV1FlowAlerts)
		alertsV1FlowAlertsModel.Op = core.StringPtr("and")
		alertsV1FlowAlertsModel.Values = []logsv0.AlertsV1FlowAlert{*alertsV1FlowAlertModel}

		model := new(logsv0.AlertsV1FlowGroup)
		model.Alerts = alertsV1FlowAlertsModel
		model.NextOp = core.StringPtr("and")

		assert.Equal(t, result, model)
	}

	alertsV1FlowAlertModel := make(map[string]interface{})
	alertsV1FlowAlertModel["id"] = "testString"
	alertsV1FlowAlertModel["not"] = true

	alertsV1FlowAlertsModel := make(map[string]interface{})
	alertsV1FlowAlertsModel["op"] = "and"
	alertsV1FlowAlertsModel["values"] = []interface{}{alertsV1FlowAlertModel}

	model := make(map[string]interface{})
	model["alerts"] = []interface{}{alertsV1FlowAlertsModel}
	model["next_op"] = "and"

	result, err := logs.ResourceIbmLogsAlertMapToAlertsV1FlowGroup(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertMapToAlertsV1FlowAlerts(t *testing.T) {
	checkResult := func(result *logsv0.AlertsV1FlowAlerts) {
		alertsV1FlowAlertModel := new(logsv0.AlertsV1FlowAlert)
		alertsV1FlowAlertModel.ID = core.StringPtr("testString")
		alertsV1FlowAlertModel.Not = core.BoolPtr(true)

		model := new(logsv0.AlertsV1FlowAlerts)
		model.Op = core.StringPtr("and")
		model.Values = []logsv0.AlertsV1FlowAlert{*alertsV1FlowAlertModel}

		assert.Equal(t, result, model)
	}

	alertsV1FlowAlertModel := make(map[string]interface{})
	alertsV1FlowAlertModel["id"] = "testString"
	alertsV1FlowAlertModel["not"] = true

	model := make(map[string]interface{})
	model["op"] = "and"
	model["values"] = []interface{}{alertsV1FlowAlertModel}

	result, err := logs.ResourceIbmLogsAlertMapToAlertsV1FlowAlerts(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertMapToAlertsV1FlowAlert(t *testing.T) {
	checkResult := func(result *logsv0.AlertsV1FlowAlert) {
		model := new(logsv0.AlertsV1FlowAlert)
		model.ID = core.StringPtr("testString")
		model.Not = core.BoolPtr(true)

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["id"] = "testString"
	model["not"] = true

	result, err := logs.ResourceIbmLogsAlertMapToAlertsV1FlowAlert(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertMapToAlertsV1FlowTimeframe(t *testing.T) {
	checkResult := func(result *logsv0.AlertsV1FlowTimeframe) {
		model := new(logsv0.AlertsV1FlowTimeframe)
		model.Ms = core.Int64Ptr(int64(0))

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["ms"] = int(0)

	result, err := logs.ResourceIbmLogsAlertMapToAlertsV1FlowTimeframe(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertMapToAlertsV2UniqueCountCondition(t *testing.T) {
	checkResult := func(result *logsv0.AlertsV2UniqueCountCondition) {
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

		assert.Equal(t, result, model)
	}

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
	alertsV2ConditionParametersModel["group_by"] = []interface{}{"testString"}
	alertsV2ConditionParametersModel["metric_alert_parameters"] = []interface{}{alertsV1MetricAlertConditionParametersModel}
	alertsV2ConditionParametersModel["metric_alert_promql_parameters"] = []interface{}{alertsV1MetricAlertPromqlConditionParametersModel}
	alertsV2ConditionParametersModel["ignore_infinity"] = true
	alertsV2ConditionParametersModel["relative_timeframe"] = "hour_or_unspecified"
	alertsV2ConditionParametersModel["cardinality_fields"] = []interface{}{"testString"}
	alertsV2ConditionParametersModel["related_extended_data"] = []interface{}{alertsV1RelatedExtendedDataModel}

	model := make(map[string]interface{})
	model["parameters"] = []interface{}{alertsV2ConditionParametersModel}

	result, err := logs.ResourceIbmLogsAlertMapToAlertsV2UniqueCountCondition(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertMapToAlertsV2LessThanUsualCondition(t *testing.T) {
	checkResult := func(result *logsv0.AlertsV2LessThanUsualCondition) {
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

		assert.Equal(t, result, model)
	}

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
	alertsV2ConditionParametersModel["group_by"] = []interface{}{"testString"}
	alertsV2ConditionParametersModel["metric_alert_parameters"] = []interface{}{alertsV1MetricAlertConditionParametersModel}
	alertsV2ConditionParametersModel["metric_alert_promql_parameters"] = []interface{}{alertsV1MetricAlertPromqlConditionParametersModel}
	alertsV2ConditionParametersModel["ignore_infinity"] = true
	alertsV2ConditionParametersModel["relative_timeframe"] = "hour_or_unspecified"
	alertsV2ConditionParametersModel["cardinality_fields"] = []interface{}{"testString"}
	alertsV2ConditionParametersModel["related_extended_data"] = []interface{}{alertsV1RelatedExtendedDataModel}

	model := make(map[string]interface{})
	model["parameters"] = []interface{}{alertsV2ConditionParametersModel}

	result, err := logs.ResourceIbmLogsAlertMapToAlertsV2LessThanUsualCondition(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertMapToAlertsV2AlertConditionConditionImmediate(t *testing.T) {
	checkResult := func(result *logsv0.AlertsV2AlertConditionConditionImmediate) {
		alertsV2ImmediateConditionModel := new(logsv0.AlertsV2ImmediateCondition)

		model := new(logsv0.AlertsV2AlertConditionConditionImmediate)
		model.Immediate = alertsV2ImmediateConditionModel

		assert.Equal(t, result, model)
	}

	alertsV2ImmediateConditionModel := make(map[string]interface{})

	model := make(map[string]interface{})
	model["immediate"] = []interface{}{alertsV2ImmediateConditionModel}

	result, err := logs.ResourceIbmLogsAlertMapToAlertsV2AlertConditionConditionImmediate(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertMapToAlertsV2AlertConditionConditionLessThan(t *testing.T) {
	checkResult := func(result *logsv0.AlertsV2AlertConditionConditionLessThan) {
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

		assert.Equal(t, result, model)
	}

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
	alertsV2ConditionParametersModel["group_by"] = []interface{}{"testString"}
	alertsV2ConditionParametersModel["metric_alert_parameters"] = []interface{}{alertsV1MetricAlertConditionParametersModel}
	alertsV2ConditionParametersModel["metric_alert_promql_parameters"] = []interface{}{alertsV1MetricAlertPromqlConditionParametersModel}
	alertsV2ConditionParametersModel["ignore_infinity"] = true
	alertsV2ConditionParametersModel["relative_timeframe"] = "hour_or_unspecified"
	alertsV2ConditionParametersModel["cardinality_fields"] = []interface{}{"testString"}
	alertsV2ConditionParametersModel["related_extended_data"] = []interface{}{alertsV1RelatedExtendedDataModel}

	alertsV2LessThanConditionModel := make(map[string]interface{})
	alertsV2LessThanConditionModel["parameters"] = []interface{}{alertsV2ConditionParametersModel}

	model := make(map[string]interface{})
	model["less_than"] = []interface{}{alertsV2LessThanConditionModel}

	result, err := logs.ResourceIbmLogsAlertMapToAlertsV2AlertConditionConditionLessThan(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertMapToAlertsV2AlertConditionConditionMoreThan(t *testing.T) {
	checkResult := func(result *logsv0.AlertsV2AlertConditionConditionMoreThan) {
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

		assert.Equal(t, result, model)
	}

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
	alertsV2ConditionParametersModel["group_by"] = []interface{}{"testString"}
	alertsV2ConditionParametersModel["metric_alert_parameters"] = []interface{}{alertsV1MetricAlertConditionParametersModel}
	alertsV2ConditionParametersModel["metric_alert_promql_parameters"] = []interface{}{alertsV1MetricAlertPromqlConditionParametersModel}
	alertsV2ConditionParametersModel["ignore_infinity"] = true
	alertsV2ConditionParametersModel["relative_timeframe"] = "hour_or_unspecified"
	alertsV2ConditionParametersModel["cardinality_fields"] = []interface{}{"testString"}
	alertsV2ConditionParametersModel["related_extended_data"] = []interface{}{alertsV1RelatedExtendedDataModel}

	alertsV2MoreThanConditionModel := make(map[string]interface{})
	alertsV2MoreThanConditionModel["parameters"] = []interface{}{alertsV2ConditionParametersModel}
	alertsV2MoreThanConditionModel["evaluation_window"] = "rolling_or_unspecified"

	model := make(map[string]interface{})
	model["more_than"] = []interface{}{alertsV2MoreThanConditionModel}

	result, err := logs.ResourceIbmLogsAlertMapToAlertsV2AlertConditionConditionMoreThan(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertMapToAlertsV2AlertConditionConditionMoreThanUsual(t *testing.T) {
	checkResult := func(result *logsv0.AlertsV2AlertConditionConditionMoreThanUsual) {
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

		assert.Equal(t, result, model)
	}

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
	alertsV2ConditionParametersModel["group_by"] = []interface{}{"testString"}
	alertsV2ConditionParametersModel["metric_alert_parameters"] = []interface{}{alertsV1MetricAlertConditionParametersModel}
	alertsV2ConditionParametersModel["metric_alert_promql_parameters"] = []interface{}{alertsV1MetricAlertPromqlConditionParametersModel}
	alertsV2ConditionParametersModel["ignore_infinity"] = true
	alertsV2ConditionParametersModel["relative_timeframe"] = "hour_or_unspecified"
	alertsV2ConditionParametersModel["cardinality_fields"] = []interface{}{"testString"}
	alertsV2ConditionParametersModel["related_extended_data"] = []interface{}{alertsV1RelatedExtendedDataModel}

	alertsV2MoreThanUsualConditionModel := make(map[string]interface{})
	alertsV2MoreThanUsualConditionModel["parameters"] = []interface{}{alertsV2ConditionParametersModel}

	model := make(map[string]interface{})
	model["more_than_usual"] = []interface{}{alertsV2MoreThanUsualConditionModel}

	result, err := logs.ResourceIbmLogsAlertMapToAlertsV2AlertConditionConditionMoreThanUsual(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertMapToAlertsV2AlertConditionConditionNewValue(t *testing.T) {
	checkResult := func(result *logsv0.AlertsV2AlertConditionConditionNewValue) {
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

		assert.Equal(t, result, model)
	}

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
	alertsV2ConditionParametersModel["group_by"] = []interface{}{"testString"}
	alertsV2ConditionParametersModel["metric_alert_parameters"] = []interface{}{alertsV1MetricAlertConditionParametersModel}
	alertsV2ConditionParametersModel["metric_alert_promql_parameters"] = []interface{}{alertsV1MetricAlertPromqlConditionParametersModel}
	alertsV2ConditionParametersModel["ignore_infinity"] = true
	alertsV2ConditionParametersModel["relative_timeframe"] = "hour_or_unspecified"
	alertsV2ConditionParametersModel["cardinality_fields"] = []interface{}{"testString"}
	alertsV2ConditionParametersModel["related_extended_data"] = []interface{}{alertsV1RelatedExtendedDataModel}

	alertsV2NewValueConditionModel := make(map[string]interface{})
	alertsV2NewValueConditionModel["parameters"] = []interface{}{alertsV2ConditionParametersModel}

	model := make(map[string]interface{})
	model["new_value"] = []interface{}{alertsV2NewValueConditionModel}

	result, err := logs.ResourceIbmLogsAlertMapToAlertsV2AlertConditionConditionNewValue(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertMapToAlertsV2AlertConditionConditionFlow(t *testing.T) {
	checkResult := func(result *logsv0.AlertsV2AlertConditionConditionFlow) {
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

		assert.Equal(t, result, model)
	}

	alertsV1FlowAlertModel := make(map[string]interface{})
	alertsV1FlowAlertModel["id"] = "testString"
	alertsV1FlowAlertModel["not"] = true

	alertsV1FlowAlertsModel := make(map[string]interface{})
	alertsV1FlowAlertsModel["op"] = "and"
	alertsV1FlowAlertsModel["values"] = []interface{}{alertsV1FlowAlertModel}

	alertsV1FlowGroupModel := make(map[string]interface{})
	alertsV1FlowGroupModel["alerts"] = []interface{}{alertsV1FlowAlertsModel}
	alertsV1FlowGroupModel["next_op"] = "and"

	alertsV1FlowTimeframeModel := make(map[string]interface{})
	alertsV1FlowTimeframeModel["ms"] = int(0)

	alertsV1FlowStageModel := make(map[string]interface{})
	alertsV1FlowStageModel["groups"] = []interface{}{alertsV1FlowGroupModel}
	alertsV1FlowStageModel["timeframe"] = []interface{}{alertsV1FlowTimeframeModel}

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
	alertsV2ConditionParametersModel["group_by"] = []interface{}{"testString"}
	alertsV2ConditionParametersModel["metric_alert_parameters"] = []interface{}{alertsV1MetricAlertConditionParametersModel}
	alertsV2ConditionParametersModel["metric_alert_promql_parameters"] = []interface{}{alertsV1MetricAlertPromqlConditionParametersModel}
	alertsV2ConditionParametersModel["ignore_infinity"] = true
	alertsV2ConditionParametersModel["relative_timeframe"] = "hour_or_unspecified"
	alertsV2ConditionParametersModel["cardinality_fields"] = []interface{}{"testString"}
	alertsV2ConditionParametersModel["related_extended_data"] = []interface{}{alertsV1RelatedExtendedDataModel}

	alertsV2FlowConditionModel := make(map[string]interface{})
	alertsV2FlowConditionModel["stages"] = []interface{}{alertsV1FlowStageModel}
	alertsV2FlowConditionModel["parameters"] = []interface{}{alertsV2ConditionParametersModel}
	alertsV2FlowConditionModel["enforce_suppression"] = true

	model := make(map[string]interface{})
	model["flow"] = []interface{}{alertsV2FlowConditionModel}

	result, err := logs.ResourceIbmLogsAlertMapToAlertsV2AlertConditionConditionFlow(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertMapToAlertsV2AlertConditionConditionUniqueCount(t *testing.T) {
	checkResult := func(result *logsv0.AlertsV2AlertConditionConditionUniqueCount) {
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

		assert.Equal(t, result, model)
	}

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
	alertsV2ConditionParametersModel["group_by"] = []interface{}{"testString"}
	alertsV2ConditionParametersModel["metric_alert_parameters"] = []interface{}{alertsV1MetricAlertConditionParametersModel}
	alertsV2ConditionParametersModel["metric_alert_promql_parameters"] = []interface{}{alertsV1MetricAlertPromqlConditionParametersModel}
	alertsV2ConditionParametersModel["ignore_infinity"] = true
	alertsV2ConditionParametersModel["relative_timeframe"] = "hour_or_unspecified"
	alertsV2ConditionParametersModel["cardinality_fields"] = []interface{}{"testString"}
	alertsV2ConditionParametersModel["related_extended_data"] = []interface{}{alertsV1RelatedExtendedDataModel}

	alertsV2UniqueCountConditionModel := make(map[string]interface{})
	alertsV2UniqueCountConditionModel["parameters"] = []interface{}{alertsV2ConditionParametersModel}

	model := make(map[string]interface{})
	model["unique_count"] = []interface{}{alertsV2UniqueCountConditionModel}

	result, err := logs.ResourceIbmLogsAlertMapToAlertsV2AlertConditionConditionUniqueCount(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertMapToAlertsV2AlertConditionConditionLessThanUsual(t *testing.T) {
	checkResult := func(result *logsv0.AlertsV2AlertConditionConditionLessThanUsual) {
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

		assert.Equal(t, result, model)
	}

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
	alertsV2ConditionParametersModel["group_by"] = []interface{}{"testString"}
	alertsV2ConditionParametersModel["metric_alert_parameters"] = []interface{}{alertsV1MetricAlertConditionParametersModel}
	alertsV2ConditionParametersModel["metric_alert_promql_parameters"] = []interface{}{alertsV1MetricAlertPromqlConditionParametersModel}
	alertsV2ConditionParametersModel["ignore_infinity"] = true
	alertsV2ConditionParametersModel["relative_timeframe"] = "hour_or_unspecified"
	alertsV2ConditionParametersModel["cardinality_fields"] = []interface{}{"testString"}
	alertsV2ConditionParametersModel["related_extended_data"] = []interface{}{alertsV1RelatedExtendedDataModel}

	alertsV2LessThanUsualConditionModel := make(map[string]interface{})
	alertsV2LessThanUsualConditionModel["parameters"] = []interface{}{alertsV2ConditionParametersModel}

	model := make(map[string]interface{})
	model["less_than_usual"] = []interface{}{alertsV2LessThanUsualConditionModel}

	result, err := logs.ResourceIbmLogsAlertMapToAlertsV2AlertConditionConditionLessThanUsual(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertMapToAlertsV2AlertNotificationGroups(t *testing.T) {
	checkResult := func(result *logsv0.AlertsV2AlertNotificationGroups) {
		alertsV2AlertNotificationModel := new(logsv0.AlertsV2AlertNotificationIntegrationTypeIntegrationID)
		alertsV2AlertNotificationModel.RetriggeringPeriodSeconds = core.Int64Ptr(int64(0))
		alertsV2AlertNotificationModel.NotifyOn = core.StringPtr("triggered_only")
		alertsV2AlertNotificationModel.IntegrationID = core.Int64Ptr(int64(0))

		model := new(logsv0.AlertsV2AlertNotificationGroups)
		model.GroupByFields = []string{"testString"}
		model.Notifications = []logsv0.AlertsV2AlertNotificationIntf{alertsV2AlertNotificationModel}

		assert.Equal(t, result, model)
	}

	alertsV2AlertNotificationModel := make(map[string]interface{})
	alertsV2AlertNotificationModel["retriggering_period_seconds"] = int(0)
	alertsV2AlertNotificationModel["notify_on"] = "triggered_only"
	alertsV2AlertNotificationModel["integration_id"] = int(0)

	model := make(map[string]interface{})
	model["group_by_fields"] = []interface{}{"testString"}
	model["notifications"] = []interface{}{alertsV2AlertNotificationModel}

	result, err := logs.ResourceIbmLogsAlertMapToAlertsV2AlertNotificationGroups(model)
	assert.Nil(t, err)
	checkResult(result)
}

// Todo @kavya498: Fix unit testcases
// func TestResourceIbmLogsAlertMapToAlertsV2AlertNotification(t *testing.T) {
// 	checkResult := func(result logsv0.AlertsV2AlertNotificationIntf) {
// 		model := new(logsv0.AlertsV2AlertNotification)
// 		model.RetriggeringPeriodSeconds = core.Int64Ptr(int64(0))
// 		model.NotifyOn = core.StringPtr("triggered_only")
// 		model.IntegrationID = core.Int64Ptr(int64(0))
// 		model.Recipients = alertsV2RecipientsModel

// 		assert.Equal(t, result, model)
// 	}

// 	model := make(map[string]interface{})
// 	model["retriggering_period_seconds"] = int(0)
// 	model["notify_on"] = "triggered_only"
// 	model["integration_id"] = int(0)
// 	model["recipients"] = []interface{}{alertsV2RecipientsModel}

// 	result, err := logs.ResourceIbmLogsAlertMapToAlertsV2AlertNotification(model)
// 	assert.Nil(t, err)
// 	checkResult(result)
// }

func TestResourceIbmLogsAlertMapToAlertsV2Recipients(t *testing.T) {
	checkResult := func(result *logsv0.AlertsV2Recipients) {
		model := new(logsv0.AlertsV2Recipients)
		model.Emails = []string{"testString"}

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["emails"] = []interface{}{"testString"}

	result, err := logs.ResourceIbmLogsAlertMapToAlertsV2Recipients(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertMapToAlertsV2AlertNotificationIntegrationTypeIntegrationID(t *testing.T) {
	checkResult := func(result *logsv0.AlertsV2AlertNotificationIntegrationTypeIntegrationID) {
		model := new(logsv0.AlertsV2AlertNotificationIntegrationTypeIntegrationID)
		model.RetriggeringPeriodSeconds = core.Int64Ptr(int64(0))
		model.NotifyOn = core.StringPtr("triggered_only")
		model.IntegrationID = core.Int64Ptr(int64(0))

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["retriggering_period_seconds"] = int(0)
	model["notify_on"] = "triggered_only"
	model["integration_id"] = int(0)

	result, err := logs.ResourceIbmLogsAlertMapToAlertsV2AlertNotificationIntegrationTypeIntegrationID(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertMapToAlertsV2AlertNotificationIntegrationTypeRecipients(t *testing.T) {
	checkResult := func(result *logsv0.AlertsV2AlertNotificationIntegrationTypeRecipients) {
		alertsV2RecipientsModel := new(logsv0.AlertsV2Recipients)
		alertsV2RecipientsModel.Emails = []string{"testString"}

		model := new(logsv0.AlertsV2AlertNotificationIntegrationTypeRecipients)
		model.RetriggeringPeriodSeconds = core.Int64Ptr(int64(0))
		model.NotifyOn = core.StringPtr("triggered_only")
		model.Recipients = alertsV2RecipientsModel

		assert.Equal(t, result, model)
	}

	alertsV2RecipientsModel := make(map[string]interface{})
	alertsV2RecipientsModel["emails"] = []interface{}{"testString"}

	model := make(map[string]interface{})
	model["retriggering_period_seconds"] = int(0)
	model["notify_on"] = "triggered_only"
	model["recipients"] = []interface{}{alertsV2RecipientsModel}

	result, err := logs.ResourceIbmLogsAlertMapToAlertsV2AlertNotificationIntegrationTypeRecipients(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertMapToAlertsV1AlertFilters(t *testing.T) {
	checkResult := func(result *logsv0.AlertsV1AlertFilters) {
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

		assert.Equal(t, result, model)
	}

	alertsV1AlertFiltersMetadataFiltersModel := make(map[string]interface{})
	alertsV1AlertFiltersMetadataFiltersModel["categories"] = []interface{}{"testString"}
	alertsV1AlertFiltersMetadataFiltersModel["applications"] = []interface{}{"testString"}
	alertsV1AlertFiltersMetadataFiltersModel["subsystems"] = []interface{}{"testString"}
	alertsV1AlertFiltersMetadataFiltersModel["computers"] = []interface{}{"testString"}
	alertsV1AlertFiltersMetadataFiltersModel["classes"] = []interface{}{"testString"}
	alertsV1AlertFiltersMetadataFiltersModel["methods"] = []interface{}{"testString"}
	alertsV1AlertFiltersMetadataFiltersModel["ip_addresses"] = []interface{}{"testString"}

	alertsV1AlertFiltersRatioAlertModel := make(map[string]interface{})
	alertsV1AlertFiltersRatioAlertModel["alias"] = "testString"
	alertsV1AlertFiltersRatioAlertModel["text"] = "testString"
	alertsV1AlertFiltersRatioAlertModel["severities"] = []interface{}{"debug_or_unspecified"}
	alertsV1AlertFiltersRatioAlertModel["applications"] = []interface{}{"testString"}
	alertsV1AlertFiltersRatioAlertModel["subsystems"] = []interface{}{"testString"}
	alertsV1AlertFiltersRatioAlertModel["group_by"] = []interface{}{"testString"}

	model := make(map[string]interface{})
	model["severities"] = []interface{}{"debug_or_unspecified"}
	model["metadata"] = []interface{}{alertsV1AlertFiltersMetadataFiltersModel}
	model["alias"] = "testString"
	model["text"] = "testString"
	model["ratio_alerts"] = []interface{}{alertsV1AlertFiltersRatioAlertModel}
	model["filter_type"] = "text_or_unspecified"

	result, err := logs.ResourceIbmLogsAlertMapToAlertsV1AlertFilters(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertMapToAlertsV1AlertFiltersMetadataFilters(t *testing.T) {
	checkResult := func(result *logsv0.AlertsV1AlertFiltersMetadataFilters) {
		model := new(logsv0.AlertsV1AlertFiltersMetadataFilters)
		model.Categories = []string{"testString"}
		model.Applications = []string{"testString"}
		model.Subsystems = []string{"testString"}
		model.Computers = []string{"testString"}
		model.Classes = []string{"testString"}
		model.Methods = []string{"testString"}
		model.IpAddresses = []string{"testString"}

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["categories"] = []interface{}{"testString"}
	model["applications"] = []interface{}{"testString"}
	model["subsystems"] = []interface{}{"testString"}
	model["computers"] = []interface{}{"testString"}
	model["classes"] = []interface{}{"testString"}
	model["methods"] = []interface{}{"testString"}
	model["ip_addresses"] = []interface{}{"testString"}

	result, err := logs.ResourceIbmLogsAlertMapToAlertsV1AlertFiltersMetadataFilters(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertMapToAlertsV1AlertFiltersRatioAlert(t *testing.T) {
	checkResult := func(result *logsv0.AlertsV1AlertFiltersRatioAlert) {
		model := new(logsv0.AlertsV1AlertFiltersRatioAlert)
		model.Alias = core.StringPtr("testString")
		model.Text = core.StringPtr("testString")
		model.Severities = []string{"debug_or_unspecified"}
		model.Applications = []string{"testString"}
		model.Subsystems = []string{"testString"}
		model.GroupBy = []string{"testString"}

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["alias"] = "testString"
	model["text"] = "testString"
	model["severities"] = []interface{}{"debug_or_unspecified"}
	model["applications"] = []interface{}{"testString"}
	model["subsystems"] = []interface{}{"testString"}
	model["group_by"] = []interface{}{"testString"}

	result, err := logs.ResourceIbmLogsAlertMapToAlertsV1AlertFiltersRatioAlert(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertMapToAlertsV1Date(t *testing.T) {
	checkResult := func(result *logsv0.AlertsV1Date) {
		model := new(logsv0.AlertsV1Date)
		model.Year = core.Int64Ptr(int64(38))
		model.Month = core.Int64Ptr(int64(38))
		model.Day = core.Int64Ptr(int64(38))

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["year"] = int(38)
	model["month"] = int(38)
	model["day"] = int(38)

	result, err := logs.ResourceIbmLogsAlertMapToAlertsV1Date(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertMapToAlertsV1AlertActiveWhen(t *testing.T) {
	checkResult := func(result *logsv0.AlertsV1AlertActiveWhen) {
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

		assert.Equal(t, result, model)
	}

	alertsV1TimeModel := make(map[string]interface{})
	alertsV1TimeModel["hours"] = int(38)
	alertsV1TimeModel["minutes"] = int(38)
	alertsV1TimeModel["seconds"] = int(38)

	alertsV1TimeRangeModel := make(map[string]interface{})
	alertsV1TimeRangeModel["start"] = []interface{}{alertsV1TimeModel}
	alertsV1TimeRangeModel["end"] = []interface{}{alertsV1TimeModel}

	alertsV1AlertActiveTimeframeModel := make(map[string]interface{})
	alertsV1AlertActiveTimeframeModel["days_of_week"] = []interface{}{"monday_or_unspecified"}
	alertsV1AlertActiveTimeframeModel["range"] = []interface{}{alertsV1TimeRangeModel}

	model := make(map[string]interface{})
	model["timeframes"] = []interface{}{alertsV1AlertActiveTimeframeModel}

	result, err := logs.ResourceIbmLogsAlertMapToAlertsV1AlertActiveWhen(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertMapToAlertsV1AlertActiveTimeframe(t *testing.T) {
	checkResult := func(result *logsv0.AlertsV1AlertActiveTimeframe) {
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

		assert.Equal(t, result, model)
	}

	alertsV1TimeModel := make(map[string]interface{})
	alertsV1TimeModel["hours"] = int(38)
	alertsV1TimeModel["minutes"] = int(38)
	alertsV1TimeModel["seconds"] = int(38)

	alertsV1TimeRangeModel := make(map[string]interface{})
	alertsV1TimeRangeModel["start"] = []interface{}{alertsV1TimeModel}
	alertsV1TimeRangeModel["end"] = []interface{}{alertsV1TimeModel}

	model := make(map[string]interface{})
	model["days_of_week"] = []interface{}{"monday_or_unspecified"}
	model["range"] = []interface{}{alertsV1TimeRangeModel}

	result, err := logs.ResourceIbmLogsAlertMapToAlertsV1AlertActiveTimeframe(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertMapToAlertsV1TimeRange(t *testing.T) {
	checkResult := func(result *logsv0.AlertsV1TimeRange) {
		alertsV1TimeModel := new(logsv0.AlertsV1Time)
		alertsV1TimeModel.Hours = core.Int64Ptr(int64(38))
		alertsV1TimeModel.Minutes = core.Int64Ptr(int64(38))
		alertsV1TimeModel.Seconds = core.Int64Ptr(int64(38))

		model := new(logsv0.AlertsV1TimeRange)
		model.Start = alertsV1TimeModel
		model.End = alertsV1TimeModel

		assert.Equal(t, result, model)
	}

	alertsV1TimeModel := make(map[string]interface{})
	alertsV1TimeModel["hours"] = int(38)
	alertsV1TimeModel["minutes"] = int(38)
	alertsV1TimeModel["seconds"] = int(38)

	model := make(map[string]interface{})
	model["start"] = []interface{}{alertsV1TimeModel}
	model["end"] = []interface{}{alertsV1TimeModel}

	result, err := logs.ResourceIbmLogsAlertMapToAlertsV1TimeRange(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertMapToAlertsV1Time(t *testing.T) {
	checkResult := func(result *logsv0.AlertsV1Time) {
		model := new(logsv0.AlertsV1Time)
		model.Hours = core.Int64Ptr(int64(38))
		model.Minutes = core.Int64Ptr(int64(38))
		model.Seconds = core.Int64Ptr(int64(38))

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["hours"] = int(38)
	model["minutes"] = int(38)
	model["seconds"] = int(38)

	result, err := logs.ResourceIbmLogsAlertMapToAlertsV1Time(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertMapToAlertsV1MetaLabel(t *testing.T) {
	checkResult := func(result *logsv0.AlertsV1MetaLabel) {
		model := new(logsv0.AlertsV1MetaLabel)
		model.Key = core.StringPtr("testString")
		model.Value = core.StringPtr("testString")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["key"] = "testString"
	model["value"] = "testString"

	result, err := logs.ResourceIbmLogsAlertMapToAlertsV1MetaLabel(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertMapToAlertsV1TracingAlert(t *testing.T) {
	checkResult := func(result *logsv0.AlertsV1TracingAlert) {
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

		assert.Equal(t, result, model)
	}

	alertsV1FiltersModel := make(map[string]interface{})
	alertsV1FiltersModel["values"] = []interface{}{"testString"}
	alertsV1FiltersModel["operator"] = "testString"

	alertsV1FilterDataModel := make(map[string]interface{})
	alertsV1FilterDataModel["field"] = "testString"
	alertsV1FilterDataModel["filters"] = []interface{}{alertsV1FiltersModel}

	model := make(map[string]interface{})
	model["condition_latency"] = int(0)
	model["field_filters"] = []interface{}{alertsV1FilterDataModel}
	model["tag_filters"] = []interface{}{alertsV1FilterDataModel}

	result, err := logs.ResourceIbmLogsAlertMapToAlertsV1TracingAlert(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertMapToAlertsV1FilterData(t *testing.T) {
	checkResult := func(result *logsv0.AlertsV1FilterData) {
		alertsV1FiltersModel := new(logsv0.AlertsV1Filters)
		alertsV1FiltersModel.Values = []string{"testString"}
		alertsV1FiltersModel.Operator = core.StringPtr("testString")

		model := new(logsv0.AlertsV1FilterData)
		model.Field = core.StringPtr("testString")
		model.Filters = []logsv0.AlertsV1Filters{*alertsV1FiltersModel}

		assert.Equal(t, result, model)
	}

	alertsV1FiltersModel := make(map[string]interface{})
	alertsV1FiltersModel["values"] = []interface{}{"testString"}
	alertsV1FiltersModel["operator"] = "testString"

	model := make(map[string]interface{})
	model["field"] = "testString"
	model["filters"] = []interface{}{alertsV1FiltersModel}

	result, err := logs.ResourceIbmLogsAlertMapToAlertsV1FilterData(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertMapToAlertsV1Filters(t *testing.T) {
	checkResult := func(result *logsv0.AlertsV1Filters) {
		model := new(logsv0.AlertsV1Filters)
		model.Values = []string{"testString"}
		model.Operator = core.StringPtr("testString")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["values"] = []interface{}{"testString"}
	model["operator"] = "testString"

	result, err := logs.ResourceIbmLogsAlertMapToAlertsV1Filters(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmLogsAlertMapToAlertsV2AlertIncidentSettings(t *testing.T) {
	checkResult := func(result *logsv0.AlertsV2AlertIncidentSettings) {
		model := new(logsv0.AlertsV2AlertIncidentSettings)
		model.RetriggeringPeriodSeconds = core.Int64Ptr(int64(0))
		model.NotifyOn = core.StringPtr("triggered_only")
		model.UseAsNotificationSettings = core.BoolPtr(true)

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["retriggering_period_seconds"] = int(0)
	model["notify_on"] = "triggered_only"
	model["use_as_notification_settings"] = true

	result, err := logs.ResourceIbmLogsAlertMapToAlertsV2AlertIncidentSettings(model)
	assert.Nil(t, err)
	checkResult(result)
}
