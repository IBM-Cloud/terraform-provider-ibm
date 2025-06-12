---
layout: "ibm"
page_title: "IBM : ibm_logs_alert"
description: |-
  Manages logs_alert.
subcategory: "Cloud Logs"
---

# ibm_logs_alert

Create, update, and delete logs_alerts with this resource.

## Example Usage

```hcl
resource "ibm_resource_instance" "logs_instance" {
  name     = "logs-instance"
  service  = "logs"
  plan     = "standard"
  location = "eu-gb"
}
resource "ibm_logs_alert" "logs_alert_instance" {
  instance_id = ibm_resource_instance.logs_instance.guid
  region      = ibm_resource_instance.logs_instance.location
  name        = "example-alert-decription"
  is_active   = true
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
```

## Argument Reference

You can specify the following arguments for this resource.
* `instance_id` - (Required, Forces new resource, String)  Cloud Logs Instance GUID.
* `region` - (Optional, Forces new resource, String) Cloud Logs Instance Region.
* `endpoint_type` - (Optional, String) Cloud Logs Instance Endpoint type. Allowed values `public` and `private`.
* `active_when` - (Optional, List) When should the alert be active.
Nested schema for **active_when**:
	* `timeframes` - (Required, List) Activity timeframes of the alert.
	  * Constraints: The maximum length is `30` items. The minimum length is `1` item.
	Nested schema for **timeframes**:
		* `days_of_week` - (Required, List) Days of the week for activity.
		  * Constraints: Allowable list items are: `monday_or_unspecified`, `tuesday`, `wednesday`, `thursday`, `friday`, `saturday`, `sunday`. The maximum length is `30` items. The minimum length is `1` item.
		* `range` - (Required, List) Time range in the day of the week.
		Nested schema for **range**:
			* `end` - (Required, List) Start time.
			Nested schema for **end**:
				* `hours` - (Optional, Integer) Hours of the day.
				  * Constraints: The maximum value is `24`. 
				* `minutes` - (Optional, Integer) Minutes of the hour.
				  * Constraints: The maximum value is `59`. 
				* `seconds` - (Optional, Integer) Seconds of the minute.
				  * Constraints: The maximum value is `59`. 
			* `start` - (Required, List) Start time.
			Nested schema for **start**:
				* `hours` - (Optional, Integer) Hours of the day.
				  * Constraints: The maximum value is `24`. 
				* `minutes` - (Optional, Integer) Minutes of the hour.
				  * Constraints: The maximum value is `59`. 
				* `seconds` - (Optional, Integer) Seconds of the minute.
				  * Constraints: The maximum value is `59`. 
* `condition` - (Required, List) Alert condition.
Nested schema for **condition**:
	* `flow` - (Optional, List) Condition for flow alert.
	Nested schema for **flow**:
		* `enforce_suppression` - (Optional, Boolean) Should suppression be enforced on the flow alert.
		* `parameters` - (Optional, List) The Less than alert condition parameters.
		Nested schema for **parameters**:
			* `cardinality_fields` - (Optional, List) Cardinality fields for unique count alert.
			  * Constraints: The list items must match regular expression `/^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$/`. The maximum length is `4096` items. The minimum length is `0` items.
			* `group_by` - (Optional, List) The group by fields for the alert condition.
			  * Constraints: The list items must match regular expression `/^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$/`. The maximum length is `3` items. The minimum length is `0` items.
			* `ignore_infinity` - (Optional, Boolean) Should the evaluation ignore infinity value.
			* `metric_alert_parameters` - (Optional, List) The lucene metric alert parameters if it is a lucene metric alert.
			Nested schema for **metric_alert_parameters**:
				* `arithmetic_operator` - (Required, String) The arithmetic operator of the metric promql alert.
				  * Constraints: Allowable values are: `avg_or_unspecified`, `min`, `max`, `sum`, `count`, `percentile`.
				* `arithmetic_operator_modifier` - (Optional, Integer) The arithmetic operator modifier of the metric promql alert.
				  * Constraints: The maximum value is `4294967295`. 
				* `metric_field` - (Required, String) The metric field of the metric alert.
				  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$/`.
				* `metric_source` - (Required, String) The metric source of the metric alert.
				  * Constraints: Allowable values are: `logs2metrics_or_unspecified`, `prometheus`.
				* `non_null_percentage` - (Optional, Integer) Non null percentage of the evaluation.
				  * Constraints: The maximum value is `4294967295`. 
				* `sample_threshold_percentage` - (Optional, Integer) The threshold percentage.
				  * Constraints: The maximum value is `4294967295`. 
				* `swap_null_values` - (Optional, Boolean) Should we swap null values with zero.
			* `metric_alert_promql_parameters` - (Optional, List) The promql metric alert parameters if is is a promql metric alert.
			Nested schema for **metric_alert_promql_parameters**:
				* `arithmetic_operator_modifier` - (Optional, Integer) The arithmetic operator of the metric promql alert.
				  * Constraints: The maximum value is `4294967295`. 
				* `non_null_percentage` - (Optional, Integer) Non null percentage of the evaluation.
				  * Constraints: The maximum value is `4294967295`. 
				* `promql_text` - (Required, String) The promql text of the metric alert by fields for the alert condition.
				  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$/`.
				* `sample_threshold_percentage` - (Required, Integer) The threshold percentage.
				  * Constraints: The maximum value is `4294967295`. 
				* `swap_null_values` - (Optional, Boolean) Should we swap null values with zero.
			* `related_extended_data` - (Optional, List) Deadman configuration.
			Nested schema for **related_extended_data**:
				* `cleanup_deadman_duration` - (Optional, String) Cleanup deadman duration.
				  * Constraints: Allowable values are: `cleanup_deadman_duration_never_or_unspecified`, `cleanup_deadman_duration_5min`, `cleanup_deadman_duration_10min`, `cleanup_deadman_duration_1h`, `cleanup_deadman_duration_2h`, `cleanup_deadman_duration_6h`, `cleanup_deadman_duration_12h`, `cleanup_deadman_duration_24h`.
				* `should_trigger_deadman` - (Optional, Boolean) Should we trigger deadman.
			* `relative_timeframe` - (Optional, String) The relative timeframe for time relative alerts.
			  * Constraints: Allowable values are: `hour_or_unspecified`, `day`, `week`, `month`.
			* `threshold` - (Optional, Float) The threshold for the alert condition.
			* `timeframe` - (Required, String) The timeframe for the alert condition.
			  * Constraints: Allowable values are: `timeframe_5_min_or_unspecified`, `timeframe_10_min`, `timeframe_20_min`, `timeframe_30_min`, `timeframe_1_h`, `timeframe_2_h`, `timeframe_3_h`, `timeframe_4_h`, `timeframe_6_h`, `timeframe_12_h`, `timeframe_24_h`, `timeframe_48_h`, `timeframe_72_h`, `timeframe_1_w`, `timeframe_1_m`, `timeframe_2_m`, `timeframe_3_m`, `timeframe_15_min`, `timeframe_1_min`, `timeframe_2_min`, `timeframe_36_h`.
		* `stages` - (Optional, List) The Flow alert condition parameters.
		  * Constraints: The maximum length is `50` items. The minimum length is `0` items.
		Nested schema for **stages**:
			* `groups` - (Optional, List) List of groups of alerts.
			  * Constraints: The maximum length is `4096` items. The minimum length is `0` items.
			Nested schema for **groups**:
				* `alerts` - (Optional, List) List of alerts.
				Nested schema for **alerts**:
					* `op` - (Optional, String) Operator for the alerts.
					  * Constraints: Allowable values are: `and`, `or`.
					* `values` - (Optional, List) List of alerts.
					  * Constraints: The maximum length is `4096` items. The minimum length is `0` items.
					Nested schema for **values**:
						* `id` - (Optional, String) The alert ID.
						  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$/`.
						* `not` - (Optional, Boolean) The alert not.
				* `next_op` - (Optional, String) Operator for the alerts.
				  * Constraints: Allowable values are: `and`, `or`.
			* `timeframe` - (Optional, List) Timeframe for the flow.
			Nested schema for **timeframe**:
				* `ms` - (Optional, Integer) Timeframe in milliseconds.
				  * Constraints: The maximum value is `4294967295`. 
	* `immediate` - (Optional, List) Condition for immediate standard alert.
	Nested schema for **immediate**:
	* `less_than` - (Optional, List) Condition for less than alert.
	Nested schema for **less_than**:
		* `parameters` - (Required, List) The Less than alert condition parameters.
		Nested schema for **parameters**:
			* `cardinality_fields` - (Optional, List) Cardinality fields for unique count alert.
			  * Constraints: The list items must match regular expression `/^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$/`. The maximum length is `4096` items. The minimum length is `0` items.
			* `group_by` - (Optional, List) The group by fields for the alert condition.
			  * Constraints: The list items must match regular expression `/^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$/`. The maximum length is `3` items. The minimum length is `0` items.
			* `ignore_infinity` - (Optional, Boolean) Should the evaluation ignore infinity value.
			* `metric_alert_parameters` - (Optional, List) The lucene metric alert parameters if it is a lucene metric alert.
			Nested schema for **metric_alert_parameters**:
				* `arithmetic_operator` - (Required, String) The arithmetic operator of the metric promql alert.
				  * Constraints: Allowable values are: `avg_or_unspecified`, `min`, `max`, `sum`, `count`, `percentile`.
				* `arithmetic_operator_modifier` - (Optional, Integer) The arithmetic operator modifier of the metric promql alert.
				  * Constraints: The maximum value is `4294967295`. 
				* `metric_field` - (Required, String) The metric field of the metric alert.
				  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$/`.
				* `metric_source` - (Required, String) The metric source of the metric alert.
				  * Constraints: Allowable values are: `logs2metrics_or_unspecified`, `prometheus`.
				* `non_null_percentage` - (Optional, Integer) Non null percentage of the evaluation.
				  * Constraints: The maximum value is `4294967295`. 
				* `sample_threshold_percentage` - (Optional, Integer) The threshold percentage.
				  * Constraints: The maximum value is `4294967295`. 
				* `swap_null_values` - (Optional, Boolean) Should we swap null values with zero.
			* `metric_alert_promql_parameters` - (Optional, List) The promql metric alert parameters if is is a promql metric alert.
			Nested schema for **metric_alert_promql_parameters**:
				* `arithmetic_operator_modifier` - (Optional, Integer) The arithmetic operator of the metric promql alert.
				  * Constraints: The maximum value is `4294967295`. 
				* `non_null_percentage` - (Optional, Integer) Non null percentage of the evaluation.
				  * Constraints: The maximum value is `4294967295`. 
				* `promql_text` - (Required, String) The promql text of the metric alert by fields for the alert condition.
				  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$/`.
				* `sample_threshold_percentage` - (Required, Integer) The threshold percentage.
				  * Constraints: The maximum value is `4294967295`. 
				* `swap_null_values` - (Optional, Boolean) Should we swap null values with zero.
			* `related_extended_data` - (Optional, List) Deadman configuration.
			Nested schema for **related_extended_data**:
				* `cleanup_deadman_duration` - (Optional, String) Cleanup deadman duration.
				  * Constraints: Allowable values are: `cleanup_deadman_duration_never_or_unspecified`, `cleanup_deadman_duration_5min`, `cleanup_deadman_duration_10min`, `cleanup_deadman_duration_1h`, `cleanup_deadman_duration_2h`, `cleanup_deadman_duration_6h`, `cleanup_deadman_duration_12h`, `cleanup_deadman_duration_24h`.
				* `should_trigger_deadman` - (Optional, Boolean) Should we trigger deadman.
			* `relative_timeframe` - (Optional, String) The relative timeframe for time relative alerts.
			  * Constraints: Allowable values are: `hour_or_unspecified`, `day`, `week`, `month`.
			* `threshold` - (Optional, Float) The threshold for the alert condition.
			* `timeframe` - (Required, String) The timeframe for the alert condition.
			  * Constraints: Allowable values are: `timeframe_5_min_or_unspecified`, `timeframe_10_min`, `timeframe_20_min`, `timeframe_30_min`, `timeframe_1_h`, `timeframe_2_h`, `timeframe_3_h`, `timeframe_4_h`, `timeframe_6_h`, `timeframe_12_h`, `timeframe_24_h`, `timeframe_48_h`, `timeframe_72_h`, `timeframe_1_w`, `timeframe_1_m`, `timeframe_2_m`, `timeframe_3_m`, `timeframe_15_min`, `timeframe_1_min`, `timeframe_2_min`, `timeframe_36_h`.
	* `less_than_usual` - (Optional, List) Condition for less than usual alert.
	Nested schema for **less_than_usual**:
		* `parameters` - (Required, List) The Less than alert condition parameters.
		Nested schema for **parameters**:
			* `cardinality_fields` - (Optional, List) Cardinality fields for unique count alert.
			  * Constraints: The list items must match regular expression `/^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$/`. The maximum length is `4096` items. The minimum length is `0` items.
			* `group_by` - (Optional, List) The group by fields for the alert condition.
			  * Constraints: The list items must match regular expression `/^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$/`. The maximum length is `3` items. The minimum length is `0` items.
			* `ignore_infinity` - (Optional, Boolean) Should the evaluation ignore infinity value.
			* `metric_alert_parameters` - (Optional, List) The lucene metric alert parameters if it is a lucene metric alert.
			Nested schema for **metric_alert_parameters**:
				* `arithmetic_operator` - (Required, String) The arithmetic operator of the metric promql alert.
				  * Constraints: Allowable values are: `avg_or_unspecified`, `min`, `max`, `sum`, `count`, `percentile`.
				* `arithmetic_operator_modifier` - (Optional, Integer) The arithmetic operator modifier of the metric promql alert.
				  * Constraints: The maximum value is `4294967295`. 
				* `metric_field` - (Required, String) The metric field of the metric alert.
				  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$/`.
				* `metric_source` - (Required, String) The metric source of the metric alert.
				  * Constraints: Allowable values are: `logs2metrics_or_unspecified`, `prometheus`.
				* `non_null_percentage` - (Optional, Integer) Non null percentage of the evaluation.
				  * Constraints: The maximum value is `4294967295`. 
				* `sample_threshold_percentage` - (Optional, Integer) The threshold percentage.
				  * Constraints: The maximum value is `4294967295`. 
				* `swap_null_values` - (Optional, Boolean) Should we swap null values with zero.
			* `metric_alert_promql_parameters` - (Optional, List) The promql metric alert parameters if is is a promql metric alert.
			Nested schema for **metric_alert_promql_parameters**:
				* `arithmetic_operator_modifier` - (Optional, Integer) The arithmetic operator of the metric promql alert.
				  * Constraints: The maximum value is `4294967295`. 
				* `non_null_percentage` - (Optional, Integer) Non null percentage of the evaluation.
				  * Constraints: The maximum value is `4294967295`. 
				* `promql_text` - (Required, String) The promql text of the metric alert by fields for the alert condition.
				  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$/`.
				* `sample_threshold_percentage` - (Required, Integer) The threshold percentage.
				  * Constraints: The maximum value is `4294967295`. 
				* `swap_null_values` - (Optional, Boolean) Should we swap null values with zero.
			* `related_extended_data` - (Optional, List) Deadman configuration.
			Nested schema for **related_extended_data**:
				* `cleanup_deadman_duration` - (Optional, String) Cleanup deadman duration.
				  * Constraints: Allowable values are: `cleanup_deadman_duration_never_or_unspecified`, `cleanup_deadman_duration_5min`, `cleanup_deadman_duration_10min`, `cleanup_deadman_duration_1h`, `cleanup_deadman_duration_2h`, `cleanup_deadman_duration_6h`, `cleanup_deadman_duration_12h`, `cleanup_deadman_duration_24h`.
				* `should_trigger_deadman` - (Optional, Boolean) Should we trigger deadman.
			* `relative_timeframe` - (Optional, String) The relative timeframe for time relative alerts.
			  * Constraints: Allowable values are: `hour_or_unspecified`, `day`, `week`, `month`.
			* `threshold` - (Optional, Float) The threshold for the alert condition.
			* `timeframe` - (Required, String) The timeframe for the alert condition.
			  * Constraints: Allowable values are: `timeframe_5_min_or_unspecified`, `timeframe_10_min`, `timeframe_20_min`, `timeframe_30_min`, `timeframe_1_h`, `timeframe_2_h`, `timeframe_3_h`, `timeframe_4_h`, `timeframe_6_h`, `timeframe_12_h`, `timeframe_24_h`, `timeframe_48_h`, `timeframe_72_h`, `timeframe_1_w`, `timeframe_1_m`, `timeframe_2_m`, `timeframe_3_m`, `timeframe_15_min`, `timeframe_1_min`, `timeframe_2_min`, `timeframe_36_h`.
	* `more_than` - (Optional, List) Condition for more than alert.
	Nested schema for **more_than**:
		* `evaluation_window` - (Optional, String) The evaluation window for the alert condition.
		  * Constraints: Allowable values are: `rolling_or_unspecified`, `dynamic`.
		* `parameters` - (Required, List) The Less than alert condition parameters.
		Nested schema for **parameters**:
			* `cardinality_fields` - (Optional, List) Cardinality fields for unique count alert.
			  * Constraints: The list items must match regular expression `/^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$/`. The maximum length is `4096` items. The minimum length is `0` items.
			* `group_by` - (Optional, List) The group by fields for the alert condition.
			  * Constraints: The list items must match regular expression `/^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$/`. The maximum length is `3` items. The minimum length is `0` items.
			* `ignore_infinity` - (Optional, Boolean) Should the evaluation ignore infinity value.
			* `metric_alert_parameters` - (Optional, List) The lucene metric alert parameters if it is a lucene metric alert.
			Nested schema for **metric_alert_parameters**:
				* `arithmetic_operator` - (Required, String) The arithmetic operator of the metric promql alert.
				  * Constraints: Allowable values are: `avg_or_unspecified`, `min`, `max`, `sum`, `count`, `percentile`.
				* `arithmetic_operator_modifier` - (Optional, Integer) The arithmetic operator modifier of the metric promql alert.
				  * Constraints: The maximum value is `4294967295`. 
				* `metric_field` - (Required, String) The metric field of the metric alert.
				  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$/`.
				* `metric_source` - (Required, String) The metric source of the metric alert.
				  * Constraints: Allowable values are: `logs2metrics_or_unspecified`, `prometheus`.
				* `non_null_percentage` - (Optional, Integer) Non null percentage of the evaluation.
				  * Constraints: The maximum value is `4294967295`. 
				* `sample_threshold_percentage` - (Optional, Integer) The threshold percentage.
				  * Constraints: The maximum value is `4294967295`. 
				* `swap_null_values` - (Optional, Boolean) Should we swap null values with zero.
			* `metric_alert_promql_parameters` - (Optional, List) The promql metric alert parameters if is is a promql metric alert.
			Nested schema for **metric_alert_promql_parameters**:
				* `arithmetic_operator_modifier` - (Optional, Integer) The arithmetic operator of the metric promql alert.
				  * Constraints: The maximum value is `4294967295`. 
				* `non_null_percentage` - (Optional, Integer) Non null percentage of the evaluation.
				  * Constraints: The maximum value is `4294967295`. 
				* `promql_text` - (Required, String) The promql text of the metric alert by fields for the alert condition.
				  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$/`.
				* `sample_threshold_percentage` - (Required, Integer) The threshold percentage.
				  * Constraints: The maximum value is `4294967295`. 
				* `swap_null_values` - (Optional, Boolean) Should we swap null values with zero.
			* `related_extended_data` - (Optional, List) Deadman configuration.
			Nested schema for **related_extended_data**:
				* `cleanup_deadman_duration` - (Optional, String) Cleanup deadman duration.
				  * Constraints: Allowable values are: `cleanup_deadman_duration_never_or_unspecified`, `cleanup_deadman_duration_5min`, `cleanup_deadman_duration_10min`, `cleanup_deadman_duration_1h`, `cleanup_deadman_duration_2h`, `cleanup_deadman_duration_6h`, `cleanup_deadman_duration_12h`, `cleanup_deadman_duration_24h`.
				* `should_trigger_deadman` - (Optional, Boolean) Should we trigger deadman.
			* `relative_timeframe` - (Optional, String) The relative timeframe for time relative alerts.
			  * Constraints: Allowable values are: `hour_or_unspecified`, `day`, `week`, `month`.
			* `threshold` - (Optional, Float) The threshold for the alert condition.
			* `timeframe` - (Required, String) The timeframe for the alert condition.
			  * Constraints: Allowable values are: `timeframe_5_min_or_unspecified`, `timeframe_10_min`, `timeframe_20_min`, `timeframe_30_min`, `timeframe_1_h`, `timeframe_2_h`, `timeframe_3_h`, `timeframe_4_h`, `timeframe_6_h`, `timeframe_12_h`, `timeframe_24_h`, `timeframe_48_h`, `timeframe_72_h`, `timeframe_1_w`, `timeframe_1_m`, `timeframe_2_m`, `timeframe_3_m`, `timeframe_15_min`, `timeframe_1_min`, `timeframe_2_min`, `timeframe_36_h`.
	* `more_than_usual` - (Optional, List) Condition for more than usual alert.
	Nested schema for **more_than_usual**:
		* `parameters` - (Required, List) The Less than alert condition parameters.
		Nested schema for **parameters**:
			* `cardinality_fields` - (Optional, List) Cardinality fields for unique count alert.
			  * Constraints: The list items must match regular expression `/^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$/`. The maximum length is `4096` items. The minimum length is `0` items.
			* `group_by` - (Optional, List) The group by fields for the alert condition.
			  * Constraints: The list items must match regular expression `/^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$/`. The maximum length is `3` items. The minimum length is `0` items.
			* `ignore_infinity` - (Optional, Boolean) Should the evaluation ignore infinity value.
			* `metric_alert_parameters` - (Optional, List) The lucene metric alert parameters if it is a lucene metric alert.
			Nested schema for **metric_alert_parameters**:
				* `arithmetic_operator` - (Required, String) The arithmetic operator of the metric promql alert.
				  * Constraints: Allowable values are: `avg_or_unspecified`, `min`, `max`, `sum`, `count`, `percentile`.
				* `arithmetic_operator_modifier` - (Optional, Integer) The arithmetic operator modifier of the metric promql alert.
				  * Constraints: The maximum value is `4294967295`. 
				* `metric_field` - (Required, String) The metric field of the metric alert.
				  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$/`.
				* `metric_source` - (Required, String) The metric source of the metric alert.
				  * Constraints: Allowable values are: `logs2metrics_or_unspecified`, `prometheus`.
				* `non_null_percentage` - (Optional, Integer) Non null percentage of the evaluation.
				  * Constraints: The maximum value is `4294967295`. 
				* `sample_threshold_percentage` - (Optional, Integer) The threshold percentage.
				  * Constraints: The maximum value is `4294967295`. 
				* `swap_null_values` - (Optional, Boolean) Should we swap null values with zero.
			* `metric_alert_promql_parameters` - (Optional, List) The promql metric alert parameters if is is a promql metric alert.
			Nested schema for **metric_alert_promql_parameters**:
				* `arithmetic_operator_modifier` - (Optional, Integer) The arithmetic operator of the metric promql alert.
				  * Constraints: The maximum value is `4294967295`. 
				* `non_null_percentage` - (Optional, Integer) Non null percentage of the evaluation.
				  * Constraints: The maximum value is `4294967295`. 
				* `promql_text` - (Required, String) The promql text of the metric alert by fields for the alert condition.
				  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$/`.
				* `sample_threshold_percentage` - (Required, Integer) The threshold percentage.
				  * Constraints: The maximum value is `4294967295`. 
				* `swap_null_values` - (Optional, Boolean) Should we swap null values with zero.
			* `related_extended_data` - (Optional, List) Deadman configuration.
			Nested schema for **related_extended_data**:
				* `cleanup_deadman_duration` - (Optional, String) Cleanup deadman duration.
				  * Constraints: Allowable values are: `cleanup_deadman_duration_never_or_unspecified`, `cleanup_deadman_duration_5min`, `cleanup_deadman_duration_10min`, `cleanup_deadman_duration_1h`, `cleanup_deadman_duration_2h`, `cleanup_deadman_duration_6h`, `cleanup_deadman_duration_12h`, `cleanup_deadman_duration_24h`.
				* `should_trigger_deadman` - (Optional, Boolean) Should we trigger deadman.
			* `relative_timeframe` - (Optional, String) The relative timeframe for time relative alerts.
			  * Constraints: Allowable values are: `hour_or_unspecified`, `day`, `week`, `month`.
			* `threshold` - (Optional, Float) The threshold for the alert condition.
			* `timeframe` - (Required, String) The timeframe for the alert condition.
			  * Constraints: Allowable values are: `timeframe_5_min_or_unspecified`, `timeframe_10_min`, `timeframe_20_min`, `timeframe_30_min`, `timeframe_1_h`, `timeframe_2_h`, `timeframe_3_h`, `timeframe_4_h`, `timeframe_6_h`, `timeframe_12_h`, `timeframe_24_h`, `timeframe_48_h`, `timeframe_72_h`, `timeframe_1_w`, `timeframe_1_m`, `timeframe_2_m`, `timeframe_3_m`, `timeframe_15_min`, `timeframe_1_min`, `timeframe_2_min`, `timeframe_36_h`.
	* `new_value` - (Optional, List) Condition for new value alert.
	Nested schema for **new_value**:
		* `parameters` - (Required, List) The Less than alert condition parameters.
		Nested schema for **parameters**:
			* `cardinality_fields` - (Optional, List) Cardinality fields for unique count alert.
			  * Constraints: The list items must match regular expression `/^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$/`. The maximum length is `4096` items. The minimum length is `0` items.
			* `group_by` - (Optional, List) The group by fields for the alert condition.
			  * Constraints: The list items must match regular expression `/^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$/`. The maximum length is `3` items. The minimum length is `0` items.
			* `ignore_infinity` - (Optional, Boolean) Should the evaluation ignore infinity value.
			* `metric_alert_parameters` - (Optional, List) The lucene metric alert parameters if it is a lucene metric alert.
			Nested schema for **metric_alert_parameters**:
				* `arithmetic_operator` - (Required, String) The arithmetic operator of the metric promql alert.
				  * Constraints: Allowable values are: `avg_or_unspecified`, `min`, `max`, `sum`, `count`, `percentile`.
				* `arithmetic_operator_modifier` - (Optional, Integer) The arithmetic operator modifier of the metric promql alert.
				  * Constraints: The maximum value is `4294967295`. 
				* `metric_field` - (Required, String) The metric field of the metric alert.
				  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$/`.
				* `metric_source` - (Required, String) The metric source of the metric alert.
				  * Constraints: Allowable values are: `logs2metrics_or_unspecified`, `prometheus`.
				* `non_null_percentage` - (Optional, Integer) Non null percentage of the evaluation.
				  * Constraints: The maximum value is `4294967295`. 
				* `sample_threshold_percentage` - (Optional, Integer) The threshold percentage.
				  * Constraints: The maximum value is `4294967295`. 
				* `swap_null_values` - (Optional, Boolean) Should we swap null values with zero.
			* `metric_alert_promql_parameters` - (Optional, List) The promql metric alert parameters if is is a promql metric alert.
			Nested schema for **metric_alert_promql_parameters**:
				* `arithmetic_operator_modifier` - (Optional, Integer) The arithmetic operator of the metric promql alert.
				  * Constraints: The maximum value is `4294967295`. 
				* `non_null_percentage` - (Optional, Integer) Non null percentage of the evaluation.
				  * Constraints: The maximum value is `4294967295`. 
				* `promql_text` - (Required, String) The promql text of the metric alert by fields for the alert condition.
				  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$/`.
				* `sample_threshold_percentage` - (Required, Integer) The threshold percentage.
				  * Constraints: The maximum value is `4294967295`. 
				* `swap_null_values` - (Optional, Boolean) Should we swap null values with zero.
			* `related_extended_data` - (Optional, List) Deadman configuration.
			Nested schema for **related_extended_data**:
				* `cleanup_deadman_duration` - (Optional, String) Cleanup deadman duration.
				  * Constraints: Allowable values are: `cleanup_deadman_duration_never_or_unspecified`, `cleanup_deadman_duration_5min`, `cleanup_deadman_duration_10min`, `cleanup_deadman_duration_1h`, `cleanup_deadman_duration_2h`, `cleanup_deadman_duration_6h`, `cleanup_deadman_duration_12h`, `cleanup_deadman_duration_24h`.
				* `should_trigger_deadman` - (Optional, Boolean) Should we trigger deadman.
			* `relative_timeframe` - (Optional, String) The relative timeframe for time relative alerts.
			  * Constraints: Allowable values are: `hour_or_unspecified`, `day`, `week`, `month`.
			* `threshold` - (Optional, Float) The threshold for the alert condition.
			* `timeframe` - (Required, String) The timeframe for the alert condition.
			  * Constraints: Allowable values are: `timeframe_5_min_or_unspecified`, `timeframe_10_min`, `timeframe_20_min`, `timeframe_30_min`, `timeframe_1_h`, `timeframe_2_h`, `timeframe_3_h`, `timeframe_4_h`, `timeframe_6_h`, `timeframe_12_h`, `timeframe_24_h`, `timeframe_48_h`, `timeframe_72_h`, `timeframe_1_w`, `timeframe_1_m`, `timeframe_2_m`, `timeframe_3_m`, `timeframe_15_min`, `timeframe_1_min`, `timeframe_2_min`, `timeframe_36_h`.
	* `unique_count` - (Optional, List) Condition for unique count alert.
	Nested schema for **unique_count**:
		* `parameters` - (Required, List) The Less than alert condition parameters.
		Nested schema for **parameters**:
			* `cardinality_fields` - (Optional, List) Cardinality fields for unique count alert.
			  * Constraints: The list items must match regular expression `/^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$/`. The maximum length is `4096` items. The minimum length is `0` items.
			* `group_by` - (Optional, List) The group by fields for the alert condition.
			  * Constraints: The list items must match regular expression `/^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$/`. The maximum length is `3` items. The minimum length is `0` items.
			* `ignore_infinity` - (Optional, Boolean) Should the evaluation ignore infinity value.
			* `metric_alert_parameters` - (Optional, List) The lucene metric alert parameters if it is a lucene metric alert.
			Nested schema for **metric_alert_parameters**:
				* `arithmetic_operator` - (Required, String) The arithmetic operator of the metric promql alert.
				  * Constraints: Allowable values are: `avg_or_unspecified`, `min`, `max`, `sum`, `count`, `percentile`.
				* `arithmetic_operator_modifier` - (Optional, Integer) The arithmetic operator modifier of the metric promql alert.
				  * Constraints: The maximum value is `4294967295`. 
				* `metric_field` - (Required, String) The metric field of the metric alert.
				  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$/`.
				* `metric_source` - (Required, String) The metric source of the metric alert.
				  * Constraints: Allowable values are: `logs2metrics_or_unspecified`, `prometheus`.
				* `non_null_percentage` - (Optional, Integer) Non null percentage of the evaluation.
				  * Constraints: The maximum value is `4294967295`. 
				* `sample_threshold_percentage` - (Optional, Integer) The threshold percentage.
				  * Constraints: The maximum value is `4294967295`. 
				* `swap_null_values` - (Optional, Boolean) Should we swap null values with zero.
			* `metric_alert_promql_parameters` - (Optional, List) The promql metric alert parameters if is is a promql metric alert.
			Nested schema for **metric_alert_promql_parameters**:
				* `arithmetic_operator_modifier` - (Optional, Integer) The arithmetic operator of the metric promql alert.
				  * Constraints: The maximum value is `4294967295`. 
				* `non_null_percentage` - (Optional, Integer) Non null percentage of the evaluation.
				  * Constraints: The maximum value is `4294967295`. 
				* `promql_text` - (Required, String) The promql text of the metric alert by fields for the alert condition.
				  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$/`.
				* `sample_threshold_percentage` - (Required, Integer) The threshold percentage.
				  * Constraints: The maximum value is `4294967295`. 
				* `swap_null_values` - (Optional, Boolean) Should we swap null values with zero.
			* `related_extended_data` - (Optional, List) Deadman configuration.
			Nested schema for **related_extended_data**:
				* `cleanup_deadman_duration` - (Optional, String) Cleanup deadman duration.
				  * Constraints: Allowable values are: `cleanup_deadman_duration_never_or_unspecified`, `cleanup_deadman_duration_5min`, `cleanup_deadman_duration_10min`, `cleanup_deadman_duration_1h`, `cleanup_deadman_duration_2h`, `cleanup_deadman_duration_6h`, `cleanup_deadman_duration_12h`, `cleanup_deadman_duration_24h`.
				* `should_trigger_deadman` - (Optional, Boolean) Should we trigger deadman.
			* `relative_timeframe` - (Optional, String) The relative timeframe for time relative alerts.
			  * Constraints: Allowable values are: `hour_or_unspecified`, `day`, `week`, `month`.
			* `threshold` - (Optional, Float) The threshold for the alert condition.
			* `timeframe` - (Required, String) The timeframe for the alert condition.
			  * Constraints: Allowable values are: `timeframe_5_min_or_unspecified`, `timeframe_10_min`, `timeframe_20_min`, `timeframe_30_min`, `timeframe_1_h`, `timeframe_2_h`, `timeframe_3_h`, `timeframe_4_h`, `timeframe_6_h`, `timeframe_12_h`, `timeframe_24_h`, `timeframe_48_h`, `timeframe_72_h`, `timeframe_1_w`, `timeframe_1_m`, `timeframe_2_m`, `timeframe_3_m`, `timeframe_15_min`, `timeframe_1_min`, `timeframe_2_min`, `timeframe_36_h`.
* `description` - (Optional, String) Alert description.
  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^[\\p{L}0-9_\\-\\s]+$/`.
* `expiration` - (Optional, List) Alert expiration date.
Nested schema for **expiration**:
	* `day` - (Optional, Integer) Day of the month.
	  * Constraints: The maximum value is `31`. The minimum value is `1`.
	* `month` - (Optional, Integer) Month of the year.
	  * Constraints: The maximum value is `12`. The minimum value is `1`.
	* `year` - (Optional, Integer) Year.
	  * Constraints: The maximum value is `2147483647`. 
* `filters` - (Optional, List) Alert filters.
Nested schema for **filters**:
	* `alias` - (Optional, String) The alias of the filter.
	  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$/`.
	* `filter_type` - (Optional, String) The type of the filter.
	  * Constraints: Allowable values are: `text_or_unspecified`, `template`, `ratio`, `unique_count`, `time_relative`, `metric`, `flow`.
	* `metadata` - (Optional, List) The metadata filters.
	Nested schema for **metadata**:
		* `applications` - (Optional, List) The applications to filter.
		  * Constraints: The list items must match regular expression `/^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$/`. The maximum length is `100` items. The minimum length is `0` items.
		* `subsystems` - (Optional, List) The subsystems to filter.
		  * Constraints: The list items must match regular expression `/^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$/`. The maximum length is `100` items. The minimum length is `0` items.
	* `ratio_alerts` - (Optional, List) The ratio alerts.
	  * Constraints: The maximum length is `4096` items. The minimum length is `0` items.
	Nested schema for **ratio_alerts**:
		* `alias` - (Required, String) The alias of the filter.
		  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$/`.
		* `applications` - (Optional, List) The applications to filter.
		  * Constraints: The list items must match regular expression `/^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$/`. The maximum length is `100` items. The minimum length is `0` items.
		* `group_by` - (Optional, List) The group by fields.
		  * Constraints: The list items must match regular expression `/^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$/`. The maximum length is `100` items. The minimum length is `0` items.
		* `severities` - (Optional, List) The severities to filter.
		  * Constraints: Allowable list items are: `debug_or_unspecified`, `verbose`, `info`, `warning`, `error`, `critical`. The maximum length is `4096` items. The minimum length is `0` items.
		* `subsystems` - (Optional, List) The subsystems to filter.
		  * Constraints: The list items must match regular expression `/^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$/`. The maximum length is `100` items. The minimum length is `0` items.
		* `text` - (Optional, String) The text to filter.
		  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$/`.
	* `severities` - (Optional, List) The severity of the logs to filter.
	  * Constraints: Allowable list items are: `debug_or_unspecified`, `verbose`, `info`, `warning`, `error`, `critical`. The maximum length is `4096` items. The minimum length is `0` items.
	* `text` - (Optional, String) The text to filter.
	  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$/`.
* `incident_settings` - (Optional, List) Incident settings, will create the incident based on this configuration.
Nested schema for **incident_settings**:
	* `notify_on` - (Optional, String) Notify on setting.
	  * Constraints: Allowable values are: `triggered_only`, `triggered_and_resolved`.
	* `retriggering_period_seconds` - (Optional, Integer) The retriggering period of the alert in seconds.
	  * Constraints: The maximum value is `4294967295`. 
	* `use_as_notification_settings` - (Optional, Boolean) Use these settings for all notificaion webhook.
* `is_active` - (Required, Boolean) Alert is active.
* `meta_labels` - (Optional, List) The Meta labels to add to the alert.
  * Constraints: The maximum length is `200` items. The minimum length is `0` items.
Nested schema for **meta_labels**:
	* `key` - (Optional, String) The key of the label.
	  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$/`.
	* `value` - (Optional, String) The value of the label.
	  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$/`.
* `meta_labels_strings` - (Optional, List) The Meta labels to add to the alert as string with ':' separator.
  * Constraints: The list items must match regular expression `/^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$/`. The maximum length is `4096` items. The minimum length is `0` items.
* `name` - (Required, String) Alert name.
  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$/`.
* `notification_groups` - (Optional, List) Alert notification groups.
  * Constraints: The maximum length is `10` items. The minimum length is `1` item.
Nested schema for **notification_groups**:
	* `group_by_fields` - (Optional, List) Group by fields to group the values by.
	  * Constraints: The list items must match regular expression `/^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$/`. The maximum length is `20` items. The minimum length is `0` items.
	* `notifications` - (Optional, List) Webhook target settings for the the notification.
	  * Constraints: The maximum length is `20` items. The minimum length is `0` items.
	Nested schema for **notifications**:
		* `integration_id` - (Optional, Integer) Integration ID.
		  * Constraints: The maximum value is `4294967295`. 
		* `notify_on` - (Optional, String) Notify on setting.
		  * Constraints: Allowable values are: `triggered_only`, `triggered_and_resolved`.
		* `recipients` - (Optional, List) Recipients.
		Nested schema for **recipients**:
			* `emails` - (Optional, List) Email addresses.
			  * Constraints: The list items must match regular expression `/^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\\.[a-zA-Z]{2,}$/`. The maximum length is `20` items. The minimum length is `0` items.
		* `retriggering_period_seconds` - (Optional, Integer) Retriggering period of the alert in seconds.
		  * Constraints: The maximum value is `4294967295`. 
* `notification_payload_filters` - (Optional, List) JSON keys to include in the alert notification, if left empty get the full log text in the alert notification.
  * Constraints: The list items must match regular expression `/^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$/`. The maximum length is `100` items. The minimum length is `0` items.
* `severity` - (Required, String) Alert severity.
  * Constraints: Allowable values are: `info_or_unspecified`, `warning`, `critical`, `error`.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the logs_alert resource.
* `alert_id` - The unique identifier of the logs alert.
* `unique_identifier` - (String) Alert unique identifier.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$/`.


## Import

You can import the `ibm_logs_alert` resource by using `id`. `id` Alert id is combination of `region`, `instance_id` and `alert_id`.

# Syntax
<pre>
$ terraform import ibm_logs_alert.logs_alert < region >/< instance_id >/< alert_id >;
</pre>

# Example
```
$ terraform import ibm_logs_alert.logs_alert eu-gb/3dc02998-0b50-4ea8-b68a-4779d716fa1f/4dc02998-0bc50-0b50-b68a-4779d716fa1f
```
