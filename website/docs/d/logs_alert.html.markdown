---
layout: "ibm"
page_title: "IBM : ibm_logs_alert"
description: |-
  Get information about logs_alert
subcategory: "Cloud Logs"
---


# ibm_logs_alert

Provides a read-only data source to retrieve information about a logs_alert. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_logs_alert" "logs_alert_instance" {
	instance_id   = ibm_logs_alert.logs_alert_instance.instance_id
	region 		  = ibm_logs_alert.logs_alert_instance.region
	logs_alert_id = ibm_logs_alert.logs_alert_instance.alert_id
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `instance_id` - (Required, String)  Cloud Logs Instance GUID.
* `region` - (Optional, String) Cloud Logs Instance Region.
* `logs_alert_id` - (Required, Forces new resource, String) Alert ID.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$/`.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the logs_alert.
* `active_when` - (List) When should the alert be active.
Nested schema for **active_when**:
	* `timeframes` - (List) Activity timeframes of the alert.
	  * Constraints: The maximum length is `30` items. The minimum length is `1` item.
	Nested schema for **timeframes**:
		* `days_of_week` - (List) Days of the week for activity.
		  * Constraints: Allowable list items are: `monday_or_unspecified`, `tuesday`, `wednesday`, `thursday`, `friday`, `saturday`, `sunday`. The maximum length is `30` items. The minimum length is `1` item.
		* `range` - (List) Time range in the day of the week.
		Nested schema for **range**:
			* `end` - (List) Start time.
			Nested schema for **end**:
				* `hours` - (Integer) Hours of the day.
				* `minutes` - (Integer) Minutes of the hour.
				* `seconds` - (Integer) Seconds of the minute.
			* `start` - (List) Start time.
			Nested schema for **start**:
				* `hours` - (Integer) Hours of the day.
				* `minutes` - (Integer) Minutes of the hour.
				* `seconds` - (Integer) Seconds of the minute.

* `condition` - (List) Alert condition.
Nested schema for **condition**:
	* `flow` - (List) Condition for flow alert.
	Nested schema for **flow**:
		* `enforce_suppression` - (Boolean) Should suppression be enforced on the flow alert.
		* `parameters` - (List) The Less than alert condition parameters.
		Nested schema for **parameters**:
			* `cardinality_fields` - (List) Cardinality fields for unique count alert.
			  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `0` items.
			* `group_by` - (List) The group by fields for the alert condition.
			  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `3` items. The minimum length is `0` items.
			* `ignore_infinity` - (Boolean) Should the evaluation ignore infinity value.
			* `metric_alert_parameters` - (List) The lucene metric alert parameters if it is a lucene metric alert.
			Nested schema for **metric_alert_parameters**:
				* `arithmetic_operator` - (String) The arithmetic operator of the metric promql alert.
				  * Constraints: Allowable values are: `avg_or_unspecified`, `min`, `max`, `sum`, `count`, `percentile`.
				* `arithmetic_operator_modifier` - (Integer) The arithmetic operator modifier of the metric promql alert.
				  * Constraints: The maximum value is `4294967295`. The minimum value is `0`.
				* `metric_field` - (String) The metric field of the metric alert.
				  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
				* `metric_source` - (String) The metric source of the metric alert.
				  * Constraints: Allowable values are: `logs2metrics_or_unspecified`, `prometheus`.
				* `non_null_percentage` - (Integer) Non null percentage of the evaluation.
				  * Constraints: The maximum value is `4294967295`. The minimum value is `0`.
				* `sample_threshold_percentage` - (Integer) The threshold percentage.
				  * Constraints: The maximum value is `4294967295`. The minimum value is `0`.
				* `swap_null_values` - (Boolean) Should we swap null values with zero.
			* `metric_alert_promql_parameters` - (List) The promql metric alert parameters if is is a promql metric alert.
			Nested schema for **metric_alert_promql_parameters**:
				* `arithmetic_operator_modifier` - (Integer) The arithmetic operator of the metric promql alert.
				  * Constraints: The maximum value is `4294967295`. The minimum value is `0`.
				* `non_null_percentage` - (Integer) Non null percentage of the evaluation.
				  * Constraints: The maximum value is `4294967295`. The minimum value is `0`.
				* `promql_text` - (String) The promql text of the metric alert by fields for the alert condition.
				  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
				* `sample_threshold_percentage` - (Integer) The threshold percentage.
				  * Constraints: The maximum value is `4294967295`. The minimum value is `0`.
				* `swap_null_values` - (Boolean) Should we swap null values with zero.
			* `related_extended_data` - (List) Deadman configuration.
			Nested schema for **related_extended_data**:
				* `cleanup_deadman_duration` - (String) Cleanup deadman duration.
				  * Constraints: Allowable values are: `cleanup_deadman_duration_never_or_unspecified`, `cleanup_deadman_duration_5min`, `cleanup_deadman_duration_10min`, `cleanup_deadman_duration_1h`, `cleanup_deadman_duration_2h`, `cleanup_deadman_duration_6h`, `cleanup_deadman_duration_12h`, `cleanup_deadman_duration_24h`.
				* `should_trigger_deadman` - (Boolean) Should we trigger deadman.
			* `relative_timeframe` - (String) The relative timeframe for time relative alerts.
			  * Constraints: Allowable values are: `hour_or_unspecified`, `day`, `week`, `month`.
			* `threshold` - (Float) The threshold for the alert condition.
			* `timeframe` - (String) The timeframe for the alert condition.
			  * Constraints: Allowable values are: `timeframe_5_min_or_unspecified`, `timeframe_10_min`, `timeframe_20_min`, `timeframe_30_min`, `timeframe_1_h`, `timeframe_2_h`, `timeframe_3_h`, `timeframe_4_h`, `timeframe_6_h`, `timeframe_12_h`, `timeframe_24_h`, `timeframe_48_h`, `timeframe_72_h`, `timeframe_1_w`, `timeframe_1_m`, `timeframe_2_m`, `timeframe_3_m`, `timeframe_15_min`, `timeframe_1_min`, `timeframe_2_min`, `timeframe_36_h`.
		* `stages` - (List) The Flow alert condition parameters.
		  * Constraints: The maximum length is `50` items. The minimum length is `0` items.
		Nested schema for **stages**:
			* `groups` - (List) List of groups of alerts.
			  * Constraints: The maximum length is `4096` items. The minimum length is `0` items.
			Nested schema for **groups**:
				* `alerts` - (List) List of alerts.
				Nested schema for **alerts**:
					* `op` - (String) Operator for the alerts.
					  * Constraints: Allowable values are: `and`, `or`.
					* `values` - (List) List of alerts.
					  * Constraints: The maximum length is `4096` items. The minimum length is `0` items.
					Nested schema for **values**:
						* `id` - (String) The alert ID.
						  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
						* `not` - (Boolean) The alert not.
				* `next_op` - (String) Operator for the alerts.
				  * Constraints: Allowable values are: `and`, `or`.
			* `timeframe` - (List) Timeframe for the flow.
			Nested schema for **timeframe**:
				* `ms` - (Integer) Timeframe in milliseconds.
				  * Constraints: The maximum value is `4294967295`. The minimum value is `0`.
	* `immediate` - (List) Condition for immediate standard alert.
	Nested schema for **immediate**:
	* `less_than` - (List) Condition for less than alert.
	Nested schema for **less_than**:
		* `parameters` - (List) The Less than alert condition parameters.
		Nested schema for **parameters**:
			* `cardinality_fields` - (List) Cardinality fields for unique count alert.
			  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `0` items.
			* `group_by` - (List) The group by fields for the alert condition.
			  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `3` items. The minimum length is `0` items.
			* `ignore_infinity` - (Boolean) Should the evaluation ignore infinity value.
			* `metric_alert_parameters` - (List) The lucene metric alert parameters if it is a lucene metric alert.
			Nested schema for **metric_alert_parameters**:
				* `arithmetic_operator` - (String) The arithmetic operator of the metric promql alert.
				  * Constraints: Allowable values are: `avg_or_unspecified`, `min`, `max`, `sum`, `count`, `percentile`.
				* `arithmetic_operator_modifier` - (Integer) The arithmetic operator modifier of the metric promql alert.
				  * Constraints: The maximum value is `4294967295`. The minimum value is `0`.
				* `metric_field` - (String) The metric field of the metric alert.
				  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
				* `metric_source` - (String) The metric source of the metric alert.
				  * Constraints: Allowable values are: `logs2metrics_or_unspecified`, `prometheus`.
				* `non_null_percentage` - (Integer) Non null percentage of the evaluation.
				  * Constraints: The maximum value is `4294967295`. The minimum value is `0`.
				* `sample_threshold_percentage` - (Integer) The threshold percentage.
				  * Constraints: The maximum value is `4294967295`. The minimum value is `0`.
				* `swap_null_values` - (Boolean) Should we swap null values with zero.
			* `metric_alert_promql_parameters` - (List) The promql metric alert parameters if is is a promql metric alert.
			Nested schema for **metric_alert_promql_parameters**:
				* `arithmetic_operator_modifier` - (Integer) The arithmetic operator of the metric promql alert.
				  * Constraints: The maximum value is `4294967295`. The minimum value is `0`.
				* `non_null_percentage` - (Integer) Non null percentage of the evaluation.
				  * Constraints: The maximum value is `4294967295`. The minimum value is `0`.
				* `promql_text` - (String) The promql text of the metric alert by fields for the alert condition.
				  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
				* `sample_threshold_percentage` - (Integer) The threshold percentage.
				  * Constraints: The maximum value is `4294967295`. The minimum value is `0`.
				* `swap_null_values` - (Boolean) Should we swap null values with zero.
			* `related_extended_data` - (List) Deadman configuration.
			Nested schema for **related_extended_data**:
				* `cleanup_deadman_duration` - (String) Cleanup deadman duration.
				  * Constraints: Allowable values are: `cleanup_deadman_duration_never_or_unspecified`, `cleanup_deadman_duration_5min`, `cleanup_deadman_duration_10min`, `cleanup_deadman_duration_1h`, `cleanup_deadman_duration_2h`, `cleanup_deadman_duration_6h`, `cleanup_deadman_duration_12h`, `cleanup_deadman_duration_24h`.
				* `should_trigger_deadman` - (Boolean) Should we trigger deadman.
			* `relative_timeframe` - (String) The relative timeframe for time relative alerts.
			  * Constraints: Allowable values are: `hour_or_unspecified`, `day`, `week`, `month`.
			* `threshold` - (Float) The threshold for the alert condition.
			* `timeframe` - (String) The timeframe for the alert condition.
			  * Constraints: Allowable values are: `timeframe_5_min_or_unspecified`, `timeframe_10_min`, `timeframe_20_min`, `timeframe_30_min`, `timeframe_1_h`, `timeframe_2_h`, `timeframe_3_h`, `timeframe_4_h`, `timeframe_6_h`, `timeframe_12_h`, `timeframe_24_h`, `timeframe_48_h`, `timeframe_72_h`, `timeframe_1_w`, `timeframe_1_m`, `timeframe_2_m`, `timeframe_3_m`, `timeframe_15_min`, `timeframe_1_min`, `timeframe_2_min`, `timeframe_36_h`.
	* `less_than_usual` - (List) Condition for less than usual alert.
	Nested schema for **less_than_usual**:
		* `parameters` - (List) The Less than alert condition parameters.
		Nested schema for **parameters**:
			* `cardinality_fields` - (List) Cardinality fields for unique count alert.
			  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `0` items.
			* `group_by` - (List) The group by fields for the alert condition.
			  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `3` items. The minimum length is `0` items.
			* `ignore_infinity` - (Boolean) Should the evaluation ignore infinity value.
			* `metric_alert_parameters` - (List) The lucene metric alert parameters if it is a lucene metric alert.
			Nested schema for **metric_alert_parameters**:
				* `arithmetic_operator` - (String) The arithmetic operator of the metric promql alert.
				  * Constraints: Allowable values are: `avg_or_unspecified`, `min`, `max`, `sum`, `count`, `percentile`.
				* `arithmetic_operator_modifier` - (Integer) The arithmetic operator modifier of the metric promql alert.
				  * Constraints: The maximum value is `4294967295`. The minimum value is `0`.
				* `metric_field` - (String) The metric field of the metric alert.
				  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
				* `metric_source` - (String) The metric source of the metric alert.
				  * Constraints: Allowable values are: `logs2metrics_or_unspecified`, `prometheus`.
				* `non_null_percentage` - (Integer) Non null percentage of the evaluation.
				  * Constraints: The maximum value is `4294967295`. The minimum value is `0`.
				* `sample_threshold_percentage` - (Integer) The threshold percentage.
				  * Constraints: The maximum value is `4294967295`. The minimum value is `0`.
				* `swap_null_values` - (Boolean) Should we swap null values with zero.
			* `metric_alert_promql_parameters` - (List) The promql metric alert parameters if is is a promql metric alert.
			Nested schema for **metric_alert_promql_parameters**:
				* `arithmetic_operator_modifier` - (Integer) The arithmetic operator of the metric promql alert.
				  * Constraints: The maximum value is `4294967295`. The minimum value is `0`.
				* `non_null_percentage` - (Integer) Non null percentage of the evaluation.
				  * Constraints: The maximum value is `4294967295`. The minimum value is `0`.
				* `promql_text` - (String) The promql text of the metric alert by fields for the alert condition.
				  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
				* `sample_threshold_percentage` - (Integer) The threshold percentage.
				  * Constraints: The maximum value is `4294967295`. The minimum value is `0`.
				* `swap_null_values` - (Boolean) Should we swap null values with zero.
			* `related_extended_data` - (List) Deadman configuration.
			Nested schema for **related_extended_data**:
				* `cleanup_deadman_duration` - (String) Cleanup deadman duration.
				  * Constraints: Allowable values are: `cleanup_deadman_duration_never_or_unspecified`, `cleanup_deadman_duration_5min`, `cleanup_deadman_duration_10min`, `cleanup_deadman_duration_1h`, `cleanup_deadman_duration_2h`, `cleanup_deadman_duration_6h`, `cleanup_deadman_duration_12h`, `cleanup_deadman_duration_24h`.
				* `should_trigger_deadman` - (Boolean) Should we trigger deadman.
			* `relative_timeframe` - (String) The relative timeframe for time relative alerts.
			  * Constraints: Allowable values are: `hour_or_unspecified`, `day`, `week`, `month`.
			* `threshold` - (Float) The threshold for the alert condition.
			* `timeframe` - (String) The timeframe for the alert condition.
			  * Constraints: Allowable values are: `timeframe_5_min_or_unspecified`, `timeframe_10_min`, `timeframe_20_min`, `timeframe_30_min`, `timeframe_1_h`, `timeframe_2_h`, `timeframe_3_h`, `timeframe_4_h`, `timeframe_6_h`, `timeframe_12_h`, `timeframe_24_h`, `timeframe_48_h`, `timeframe_72_h`, `timeframe_1_w`, `timeframe_1_m`, `timeframe_2_m`, `timeframe_3_m`, `timeframe_15_min`, `timeframe_1_min`, `timeframe_2_min`, `timeframe_36_h`.
	* `more_than` - (List) Condition for more than alert.
	Nested schema for **more_than**:
		* `evaluation_window` - (String) The evaluation window for the alert condition.
		  * Constraints: Allowable values are: `rolling_or_unspecified`, `dynamic`.
		* `parameters` - (List) The Less than alert condition parameters.
		Nested schema for **parameters**:
			* `cardinality_fields` - (List) Cardinality fields for unique count alert.
			  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `0` items.
			* `group_by` - (List) The group by fields for the alert condition.
			  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `3` items. The minimum length is `0` items.
			* `ignore_infinity` - (Boolean) Should the evaluation ignore infinity value.
			* `metric_alert_parameters` - (List) The lucene metric alert parameters if it is a lucene metric alert.
			Nested schema for **metric_alert_parameters**:
				* `arithmetic_operator` - (String) The arithmetic operator of the metric promql alert.
				  * Constraints: Allowable values are: `avg_or_unspecified`, `min`, `max`, `sum`, `count`, `percentile`.
				* `arithmetic_operator_modifier` - (Integer) The arithmetic operator modifier of the metric promql alert.
				  * Constraints: The maximum value is `4294967295`. The minimum value is `0`.
				* `metric_field` - (String) The metric field of the metric alert.
				  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
				* `metric_source` - (String) The metric source of the metric alert.
				  * Constraints: Allowable values are: `logs2metrics_or_unspecified`, `prometheus`.
				* `non_null_percentage` - (Integer) Non null percentage of the evaluation.
				  * Constraints: The maximum value is `4294967295`. The minimum value is `0`.
				* `sample_threshold_percentage` - (Integer) The threshold percentage.
				  * Constraints: The maximum value is `4294967295`. The minimum value is `0`.
				* `swap_null_values` - (Boolean) Should we swap null values with zero.
			* `metric_alert_promql_parameters` - (List) The promql metric alert parameters if is is a promql metric alert.
			Nested schema for **metric_alert_promql_parameters**:
				* `arithmetic_operator_modifier` - (Integer) The arithmetic operator of the metric promql alert.
				  * Constraints: The maximum value is `4294967295`. The minimum value is `0`.
				* `non_null_percentage` - (Integer) Non null percentage of the evaluation.
				  * Constraints: The maximum value is `4294967295`. The minimum value is `0`.
				* `promql_text` - (String) The promql text of the metric alert by fields for the alert condition.
				  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
				* `sample_threshold_percentage` - (Integer) The threshold percentage.
				  * Constraints: The maximum value is `4294967295`. The minimum value is `0`.
				* `swap_null_values` - (Boolean) Should we swap null values with zero.
			* `related_extended_data` - (List) Deadman configuration.
			Nested schema for **related_extended_data**:
				* `cleanup_deadman_duration` - (String) Cleanup deadman duration.
				  * Constraints: Allowable values are: `cleanup_deadman_duration_never_or_unspecified`, `cleanup_deadman_duration_5min`, `cleanup_deadman_duration_10min`, `cleanup_deadman_duration_1h`, `cleanup_deadman_duration_2h`, `cleanup_deadman_duration_6h`, `cleanup_deadman_duration_12h`, `cleanup_deadman_duration_24h`.
				* `should_trigger_deadman` - (Boolean) Should we trigger deadman.
			* `relative_timeframe` - (String) The relative timeframe for time relative alerts.
			  * Constraints: Allowable values are: `hour_or_unspecified`, `day`, `week`, `month`.
			* `threshold` - (Float) The threshold for the alert condition.
			* `timeframe` - (String) The timeframe for the alert condition.
			  * Constraints: Allowable values are: `timeframe_5_min_or_unspecified`, `timeframe_10_min`, `timeframe_20_min`, `timeframe_30_min`, `timeframe_1_h`, `timeframe_2_h`, `timeframe_3_h`, `timeframe_4_h`, `timeframe_6_h`, `timeframe_12_h`, `timeframe_24_h`, `timeframe_48_h`, `timeframe_72_h`, `timeframe_1_w`, `timeframe_1_m`, `timeframe_2_m`, `timeframe_3_m`, `timeframe_15_min`, `timeframe_1_min`, `timeframe_2_min`, `timeframe_36_h`.
	* `more_than_usual` - (List) Condition for more than usual alert.
	Nested schema for **more_than_usual**:
		* `parameters` - (List) The Less than alert condition parameters.
		Nested schema for **parameters**:
			* `cardinality_fields` - (List) Cardinality fields for unique count alert.
			  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `0` items.
			* `group_by` - (List) The group by fields for the alert condition.
			  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `3` items. The minimum length is `0` items.
			* `ignore_infinity` - (Boolean) Should the evaluation ignore infinity value.
			* `metric_alert_parameters` - (List) The lucene metric alert parameters if it is a lucene metric alert.
			Nested schema for **metric_alert_parameters**:
				* `arithmetic_operator` - (String) The arithmetic operator of the metric promql alert.
				  * Constraints: Allowable values are: `avg_or_unspecified`, `min`, `max`, `sum`, `count`, `percentile`.
				* `arithmetic_operator_modifier` - (Integer) The arithmetic operator modifier of the metric promql alert.
				  * Constraints: The maximum value is `4294967295`. The minimum value is `0`.
				* `metric_field` - (String) The metric field of the metric alert.
				  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
				* `metric_source` - (String) The metric source of the metric alert.
				  * Constraints: Allowable values are: `logs2metrics_or_unspecified`, `prometheus`.
				* `non_null_percentage` - (Integer) Non null percentage of the evaluation.
				  * Constraints: The maximum value is `4294967295`. The minimum value is `0`.
				* `sample_threshold_percentage` - (Integer) The threshold percentage.
				  * Constraints: The maximum value is `4294967295`. The minimum value is `0`.
				* `swap_null_values` - (Boolean) Should we swap null values with zero.
			* `metric_alert_promql_parameters` - (List) The promql metric alert parameters if is is a promql metric alert.
			Nested schema for **metric_alert_promql_parameters**:
				* `arithmetic_operator_modifier` - (Integer) The arithmetic operator of the metric promql alert.
				  * Constraints: The maximum value is `4294967295`. The minimum value is `0`.
				* `non_null_percentage` - (Integer) Non null percentage of the evaluation.
				  * Constraints: The maximum value is `4294967295`. The minimum value is `0`.
				* `promql_text` - (String) The promql text of the metric alert by fields for the alert condition.
				  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
				* `sample_threshold_percentage` - (Integer) The threshold percentage.
				  * Constraints: The maximum value is `4294967295`. The minimum value is `0`.
				* `swap_null_values` - (Boolean) Should we swap null values with zero.
			* `related_extended_data` - (List) Deadman configuration.
			Nested schema for **related_extended_data**:
				* `cleanup_deadman_duration` - (String) Cleanup deadman duration.
				  * Constraints: Allowable values are: `cleanup_deadman_duration_never_or_unspecified`, `cleanup_deadman_duration_5min`, `cleanup_deadman_duration_10min`, `cleanup_deadman_duration_1h`, `cleanup_deadman_duration_2h`, `cleanup_deadman_duration_6h`, `cleanup_deadman_duration_12h`, `cleanup_deadman_duration_24h`.
				* `should_trigger_deadman` - (Boolean) Should we trigger deadman.
			* `relative_timeframe` - (String) The relative timeframe for time relative alerts.
			  * Constraints: Allowable values are: `hour_or_unspecified`, `day`, `week`, `month`.
			* `threshold` - (Float) The threshold for the alert condition.
			* `timeframe` - (String) The timeframe for the alert condition.
			  * Constraints: Allowable values are: `timeframe_5_min_or_unspecified`, `timeframe_10_min`, `timeframe_20_min`, `timeframe_30_min`, `timeframe_1_h`, `timeframe_2_h`, `timeframe_3_h`, `timeframe_4_h`, `timeframe_6_h`, `timeframe_12_h`, `timeframe_24_h`, `timeframe_48_h`, `timeframe_72_h`, `timeframe_1_w`, `timeframe_1_m`, `timeframe_2_m`, `timeframe_3_m`, `timeframe_15_min`, `timeframe_1_min`, `timeframe_2_min`, `timeframe_36_h`.
	* `new_value` - (List) Condition for new value alert.
	Nested schema for **new_value**:
		* `parameters` - (List) The Less than alert condition parameters.
		Nested schema for **parameters**:
			* `cardinality_fields` - (List) Cardinality fields for unique count alert.
			  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `0` items.
			* `group_by` - (List) The group by fields for the alert condition.
			  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `3` items. The minimum length is `0` items.
			* `ignore_infinity` - (Boolean) Should the evaluation ignore infinity value.
			* `metric_alert_parameters` - (List) The lucene metric alert parameters if it is a lucene metric alert.
			Nested schema for **metric_alert_parameters**:
				* `arithmetic_operator` - (String) The arithmetic operator of the metric promql alert.
				  * Constraints: Allowable values are: `avg_or_unspecified`, `min`, `max`, `sum`, `count`, `percentile`.
				* `arithmetic_operator_modifier` - (Integer) The arithmetic operator modifier of the metric promql alert.
				  * Constraints: The maximum value is `4294967295`. The minimum value is `0`.
				* `metric_field` - (String) The metric field of the metric alert.
				  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
				* `metric_source` - (String) The metric source of the metric alert.
				  * Constraints: Allowable values are: `logs2metrics_or_unspecified`, `prometheus`.
				* `non_null_percentage` - (Integer) Non null percentage of the evaluation.
				  * Constraints: The maximum value is `4294967295`. The minimum value is `0`.
				* `sample_threshold_percentage` - (Integer) The threshold percentage.
				  * Constraints: The maximum value is `4294967295`. The minimum value is `0`.
				* `swap_null_values` - (Boolean) Should we swap null values with zero.
			* `metric_alert_promql_parameters` - (List) The promql metric alert parameters if is is a promql metric alert.
			Nested schema for **metric_alert_promql_parameters**:
				* `arithmetic_operator_modifier` - (Integer) The arithmetic operator of the metric promql alert.
				  * Constraints: The maximum value is `4294967295`. The minimum value is `0`.
				* `non_null_percentage` - (Integer) Non null percentage of the evaluation.
				  * Constraints: The maximum value is `4294967295`. The minimum value is `0`.
				* `promql_text` - (String) The promql text of the metric alert by fields for the alert condition.
				  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
				* `sample_threshold_percentage` - (Integer) The threshold percentage.
				  * Constraints: The maximum value is `4294967295`. The minimum value is `0`.
				* `swap_null_values` - (Boolean) Should we swap null values with zero.
			* `related_extended_data` - (List) Deadman configuration.
			Nested schema for **related_extended_data**:
				* `cleanup_deadman_duration` - (String) Cleanup deadman duration.
				  * Constraints: Allowable values are: `cleanup_deadman_duration_never_or_unspecified`, `cleanup_deadman_duration_5min`, `cleanup_deadman_duration_10min`, `cleanup_deadman_duration_1h`, `cleanup_deadman_duration_2h`, `cleanup_deadman_duration_6h`, `cleanup_deadman_duration_12h`, `cleanup_deadman_duration_24h`.
				* `should_trigger_deadman` - (Boolean) Should we trigger deadman.
			* `relative_timeframe` - (String) The relative timeframe for time relative alerts.
			  * Constraints: Allowable values are: `hour_or_unspecified`, `day`, `week`, `month`.
			* `threshold` - (Float) The threshold for the alert condition.
			* `timeframe` - (String) The timeframe for the alert condition.
			  * Constraints: Allowable values are: `timeframe_5_min_or_unspecified`, `timeframe_10_min`, `timeframe_20_min`, `timeframe_30_min`, `timeframe_1_h`, `timeframe_2_h`, `timeframe_3_h`, `timeframe_4_h`, `timeframe_6_h`, `timeframe_12_h`, `timeframe_24_h`, `timeframe_48_h`, `timeframe_72_h`, `timeframe_1_w`, `timeframe_1_m`, `timeframe_2_m`, `timeframe_3_m`, `timeframe_15_min`, `timeframe_1_min`, `timeframe_2_min`, `timeframe_36_h`.
	* `unique_count` - (List) Condition for unique count alert.
	Nested schema for **unique_count**:
		* `parameters` - (List) The Less than alert condition parameters.
		Nested schema for **parameters**:
			* `cardinality_fields` - (List) Cardinality fields for unique count alert.
			  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `0` items.
			* `group_by` - (List) The group by fields for the alert condition.
			  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `3` items. The minimum length is `0` items.
			* `ignore_infinity` - (Boolean) Should the evaluation ignore infinity value.
			* `metric_alert_parameters` - (List) The lucene metric alert parameters if it is a lucene metric alert.
			Nested schema for **metric_alert_parameters**:
				* `arithmetic_operator` - (String) The arithmetic operator of the metric promql alert.
				  * Constraints: Allowable values are: `avg_or_unspecified`, `min`, `max`, `sum`, `count`, `percentile`.
				* `arithmetic_operator_modifier` - (Integer) The arithmetic operator modifier of the metric promql alert.
				  * Constraints: The maximum value is `4294967295`. The minimum value is `0`.
				* `metric_field` - (String) The metric field of the metric alert.
				  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
				* `metric_source` - (String) The metric source of the metric alert.
				  * Constraints: Allowable values are: `logs2metrics_or_unspecified`, `prometheus`.
				* `non_null_percentage` - (Integer) Non null percentage of the evaluation.
				  * Constraints: The maximum value is `4294967295`. The minimum value is `0`.
				* `sample_threshold_percentage` - (Integer) The threshold percentage.
				  * Constraints: The maximum value is `4294967295`. The minimum value is `0`.
				* `swap_null_values` - (Boolean) Should we swap null values with zero.
			* `metric_alert_promql_parameters` - (List) The promql metric alert parameters if is is a promql metric alert.
			Nested schema for **metric_alert_promql_parameters**:
				* `arithmetic_operator_modifier` - (Integer) The arithmetic operator of the metric promql alert.
				  * Constraints: The maximum value is `4294967295`. The minimum value is `0`.
				* `non_null_percentage` - (Integer) Non null percentage of the evaluation.
				  * Constraints: The maximum value is `4294967295`. The minimum value is `0`.
				* `promql_text` - (String) The promql text of the metric alert by fields for the alert condition.
				  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
				* `sample_threshold_percentage` - (Integer) The threshold percentage.
				  * Constraints: The maximum value is `4294967295`. The minimum value is `0`.
				* `swap_null_values` - (Boolean) Should we swap null values with zero.
			* `related_extended_data` - (List) Deadman configuration.
			Nested schema for **related_extended_data**:
				* `cleanup_deadman_duration` - (String) Cleanup deadman duration.
				  * Constraints: Allowable values are: `cleanup_deadman_duration_never_or_unspecified`, `cleanup_deadman_duration_5min`, `cleanup_deadman_duration_10min`, `cleanup_deadman_duration_1h`, `cleanup_deadman_duration_2h`, `cleanup_deadman_duration_6h`, `cleanup_deadman_duration_12h`, `cleanup_deadman_duration_24h`.
				* `should_trigger_deadman` - (Boolean) Should we trigger deadman.
			* `relative_timeframe` - (String) The relative timeframe for time relative alerts.
			  * Constraints: Allowable values are: `hour_or_unspecified`, `day`, `week`, `month`.
			* `threshold` - (Float) The threshold for the alert condition.
			* `timeframe` - (String) The timeframe for the alert condition.
			  * Constraints: Allowable values are: `timeframe_5_min_or_unspecified`, `timeframe_10_min`, `timeframe_20_min`, `timeframe_30_min`, `timeframe_1_h`, `timeframe_2_h`, `timeframe_3_h`, `timeframe_4_h`, `timeframe_6_h`, `timeframe_12_h`, `timeframe_24_h`, `timeframe_48_h`, `timeframe_72_h`, `timeframe_1_w`, `timeframe_1_m`, `timeframe_2_m`, `timeframe_3_m`, `timeframe_15_min`, `timeframe_1_min`, `timeframe_2_min`, `timeframe_36_h`.

* `description` - (String) Alert description.
  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9_\\-\\s]+$/`.

* `expiration` - (List) Alert expiration date.
Nested schema for **expiration**:
	* `day` - (Integer) Day of the month.
	* `month` - (Integer) Month of the year.
	* `year` - (Integer) Year.

* `filters` - (List) Alert filters.
Nested schema for **filters**:
	* `alias` - (String) The alias of the filter.
	  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
	* `filter_type` - (String) The type of the filter.
	  * Constraints: Allowable values are: `text_or_unspecified`, `template`, `ratio`, `unique_count`, `time_relative`, `metric`, `flow`.
	* `metadata` - (List) The metadata filters.
	Nested schema for **metadata**:
		* `applications` - (List) The applications to filter.
		  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `100` items. The minimum length is `0` items.
		* `subsystems` - (List) The subsystems to filter.
		  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `100` items. The minimum length is `0` items.
	* `ratio_alerts` - (List) The ratio alerts.
	  * Constraints: The maximum length is `4096` items. The minimum length is `0` items.
	Nested schema for **ratio_alerts**:
		* `alias` - (String) The alias of the filter.
		  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
		* `applications` - (List) The applications to filter.
		  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `100` items. The minimum length is `0` items.
		* `group_by` - (List) The group by fields.
		  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `100` items. The minimum length is `0` items.
		* `severities` - (List) The severities to filter.
		  * Constraints: Allowable list items are: `debug_or_unspecified`, `verbose`, `info`, `warning`, `error`, `critical`. The maximum length is `4096` items. The minimum length is `0` items.
		* `subsystems` - (List) The subsystems to filter.
		  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `100` items. The minimum length is `0` items.
		* `text` - (String) The text to filter.
		  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
	* `severities` - (List) The severity of the logs to filter.
	  * Constraints: Allowable list items are: `debug_or_unspecified`, `verbose`, `info`, `warning`, `error`, `critical`. The maximum length is `4096` items. The minimum length is `0` items.
	* `text` - (String) The text to filter.
	  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.

* `incident_settings` - (List) Incident settings, will create the incident based on this configuration.
Nested schema for **incident_settings**:
	* `notify_on` - (String) Notify on setting.
	  * Constraints: Allowable values are: `triggered_only`, `triggered_and_resolved`.
	* `retriggering_period_seconds` - (Integer) The retriggering period of the alert in seconds.
	  * Constraints: The maximum value is `4294967295`. The minimum value is `0`.
	* `use_as_notification_settings` - (Boolean) Use these settings for all notificaion webhook.

* `is_active` - (Boolean) Alert is active.

* `meta_labels` - (List) The Meta labels to add to the alert.
  * Constraints: The maximum length is `200` items. The minimum length is `0` items.
Nested schema for **meta_labels**:
	* `key` - (String) The key of the label.
	  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
	* `value` - (String) The value of the label.
	  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.

* `meta_labels_strings` - (List) The Meta labels to add to the alert as string with ':' separator.
  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `0` items.

* `name` - (String) Alert name.
  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.

* `notification_groups` - (List) Alert notification groups.
  * Constraints: The maximum length is `10` items. The minimum length is `1` item.
Nested schema for **notification_groups**:
	* `group_by_fields` - (List) Group by fields to group the values by.
	  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `20` items. The minimum length is `0` items.
	* `notifications` - (List) Webhook target settings for the the notification.
	  * Constraints: The maximum length is `20` items. The minimum length is `0` items.
	Nested schema for **notifications**:
		* `integration_id` - (Integer) Integration ID.
		  * Constraints: The maximum value is `4294967295`. The minimum value is `0`.
		* `notify_on` - (String) Notify on setting.
		  * Constraints: Allowable values are: `triggered_only`, `triggered_and_resolved`.
		* `recipients` - (List) Recipients.
		Nested schema for **recipients**:
			* `emails` - (List) Email addresses.
			  * Constraints: The list items must match regular expression `/^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\\.[a-zA-Z]{2,}$/`. The maximum length is `20` items. The minimum length is `0` items.
		* `retriggering_period_seconds` - (Integer) Retriggering period of the alert in seconds.
		  * Constraints: The maximum value is `4294967295`. The minimum value is `0`.

* `notification_payload_filters` - (List) JSON keys to include in the alert notification, if left empty get the full log text in the alert notification.
  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `100` items. The minimum length is `0` items.

* `severity` - (String) Alert severity.
  * Constraints: Allowable values are: `info_or_unspecified`, `warning`, `critical`, `error`.

* `unique_identifier` - (String) Alert unique identifier.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$/`.

