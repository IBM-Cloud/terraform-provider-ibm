---
layout: "ibm"
page_title: "IBM : ibm_logs_alert_definition"
description: |-
  Get information about logs_alert_definition
subcategory: "Cloud Logs"
---

# ibm_logs_alert_definition

Provides a read-only data source to retrieve information about a logs_alert_definition. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_logs_alert_definition" "logs_alert_definition" {
	instance_id 			 = "470e285d-3354-44f8-8119-c91902d23"
	region      			 = "eu-gb"
	logs_alert_definition_id = 3dc02998-0b50-4ea8-b68a-4779d716fa1f
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `instance_id` - (Required, Forces new resource, String)  Cloud Logs Instance GUID.
* `region` - (Optional, Forces new resource, String) Cloud Logs Instance Region.
* `logs_alert_definition_id` - (Required, Forces new resource, String) Alert definition ID.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$/`.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the logs_alert_definition.
* `active_on` - (List) Defining when the alert is active.
Nested schema for **active_on**:
	* `day_of_week` - (List) Days of the week when the alert is active.
	  * Constraints: Allowable list items are: `monday_or_unspecified`, `tuesday`, `wednesday`, `thursday`, `friday`, `saturday`, `sunday`. The maximum length is `7` items. The minimum length is `1` item.
	* `end_time` - (List) Start time of the alert activity.
	Nested schema for **end_time**:
		* `hours` - (Integer) The hour of the day in 24-hour format. Must be an integer between 0 and 23.
		  * Constraints: The maximum value is `23`. The minimum value is `0`.
		* `minutes` - (Integer) Minute of the hour of the day. Must be an integer between 0 and 59.
		  * Constraints: The maximum value is `59`. The minimum value is `0`.
	* `start_time` - (List) Start time of the alert activity.
	Nested schema for **start_time**:
		* `hours` - (Integer) The hour of the day in 24-hour format. Must be an integer between 0 and 23.
		  * Constraints: The maximum value is `23`. The minimum value is `0`.
		* `minutes` - (Integer) Minute of the hour of the day. Must be an integer between 0 and 59.
		  * Constraints: The maximum value is `59`. The minimum value is `0`.
* `alert_version_id` - (String) The previous or old alert ID.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$/`.
* `created_time` - (String) The time when the alert definition was created.
* `deleted` - (Boolean) Whether the alert has been marked as deleted.
* `description` - (String) A detailed description of what the alert monitors and when it triggers.
  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$/`.
* `enabled` - (Boolean) Whether the alert is currently active and monitoring. If true, alert is active.
* `entity_labels` - (Map) Labels used to identify and categorize the alert entity.
* `flow` - (List) Configuration for flow alerts.
Nested schema for **flow**:
	* `enforce_suppression` - (Boolean) Whether to enforce suppression for the flow alert.
	* `stages` - (List) The definition of stages of the flow alert.
	  * Constraints: The maximum length is `4096` items. The minimum length is `1` item.
	Nested schema for **stages**:
		* `flow_stages_groups` - (List) The definition of groups in the flow alert.
		Nested schema for **flow_stages_groups**:
			* `groups` - (List) The definition of an array of groups with alerts and logical operation among those alerts in the flow alert.
			  * Constraints: The maximum length is `4096` items. The minimum length is `1` item.
			Nested schema for **groups**:
				* `alert_defs` - (List) The alert definitions for the flow stage group.
				  * Constraints: The maximum length is `4096` items. The minimum length is `1` item.
				Nested schema for **alert_defs**:
					* `id` - (String) The alert definition ID.
					  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$/`.
					* `not` - (Boolean) Whether or not to negate the alert definition. If true, flow checks for the negate condition of the respective alert.
				* `alerts_op` - (String) The logical operation to apply to the alerts in the group.
				  * Constraints: Allowable values are: `and_or_unspecified`, `or`.
				* `next_op` - (String) The logical operation to apply to the next stage.
				  * Constraints: Allowable values are: `and_or_unspecified`, `or`.
		* `timeframe_ms` - (String) The timeframe for the flow alert in milliseconds.
		  * Constraints: The maximum length is `19` characters. The minimum length is `1` character. The value must match regular expression `/^-?\\d{1,19}$/`.
		* `timeframe_type` - (String) The type of timeframe for the flow alert.
		  * Constraints: Allowable values are: `unspecified`, `up_to`.
* `group_by_keys` - (List) Keys used to group and aggregate alert data.
  * Constraints: The list items must match regular expression `/^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$/`. The maximum length is `2` items. The minimum length is `0` items.
* `incidents_settings` - (List) Incident creation and management settings.
Nested schema for **incidents_settings**:
	* `minutes` - (Integer) The time in minutes before the alert can be triggered again.
	  * Constraints: The maximum value is `4294967295`. The minimum value is `0`.
	* `notify_on` - (String) Indicate if the alert should be triggered or triggered and resolved.
	  * Constraints: Allowable values are: `triggered_only_unspecified`, `triggered_and_resolved`.
* `logs_anomaly` - (List) Configuration for the log-based anomaly detection alerts.
Nested schema for **logs_anomaly**:
	* `anomaly_alert_settings` - (List) The anomaly alert settings configuration.
	Nested schema for **anomaly_alert_settings**:
		* `percentage_of_deviation` - (Float) The percentage of deviation from the baseline when the alert is triggered.
	* `condition_type` - (String) The condition type for the alert.
	  * Constraints: Allowable values are: `more_than_usual_or_unspecified`.
	* `evaluation_delay_ms` - (Integer) The delay in milliseconds before evaluating the alert condition.
	  * Constraints: The maximum value is `2147483647`. The minimum value is `0`.
	* `logs_filter` - (List) The filter to match log entries for immediate alerts.
	Nested schema for **logs_filter**:
		* `simple_filter` - (List) A simple filter that uses a Lucene query and label filters.
		Nested schema for **simple_filter**:
			* `label_filters` - (List) The label filters to filter logs.
			Nested schema for **label_filters**:
				* `application_name` - (List) Filter by application names.
				  * Constraints: The maximum length is `4096` items. The minimum length is `0` items.
				Nested schema for **application_name**:
					* `operation` - (String) The operation to perform on the label value.
					  * Constraints: Allowable values are: `is_or_unspecified`, `includes`, `ends_with`, `starts_with`.
					* `value` - (String) The value used to filter the label.
					  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^.*$/`.
				* `severities` - (List) Filter by log severities.
				  * Constraints: Allowable list items are: `verbose_unspecified`, `debug`, `info`, `warning`, `error`, `critical`. The maximum length is `4096` items. The minimum length is `0` items.
				* `subsystem_name` - (List) Filter by subsystem names.
				  * Constraints: The maximum length is `4096` items. The minimum length is `0` items.
				Nested schema for **subsystem_name**:
					* `operation` - (String) The operation to perform on the label value.
					  * Constraints: Allowable values are: `is_or_unspecified`, `includes`, `ends_with`, `starts_with`.
					* `value` - (String) The value used to filter the label.
					  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^.*$/`.
			* `lucene_query` - (String) The Lucene query to filter logs.
			  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$/`.
	* `notification_payload_filter` - (List) The notification payload filter to specify which fields are included in the notification.
	  * Constraints: The list items must match regular expression `/^[a-zA-Z0-9_.]+$/`. The maximum length is `4096` items. The minimum length is `0` items.
	* `rules` - (List) The condition rules for the log anomaly alert.
	  * Constraints: The maximum length is `1` item. The minimum length is `1` item.
	Nested schema for **rules**:
		* `condition` - (List) The condition for the anomaly alert.
		Nested schema for **condition**:
			* `minimum_threshold` - (Float) The threshold value for the alert condition.
			* `time_window` - (List) The time window for the alert condition.
			Nested schema for **time_window**:
				* `logs_time_window_specific_value` - (String) The time window defined for an alert to be triggered.
				  * Constraints: Allowable values are: `minutes_5_or_unspecified`, `minutes_10`, `minutes_20`, `minutes_15`, `minutes_30`, `hour_1`, `hours_2`, `hours_4`, `hours_6`, `hours_12`, `hours_24`, `hours_36`.
* `logs_immediate` - (List) Configuration for immediate log-based alerts.
Nested schema for **logs_immediate**:
	* `logs_filter` - (List) The filter to match log entries for immediate alerts.
	Nested schema for **logs_filter**:
		* `simple_filter` - (List) A simple filter that uses a Lucene query and label filters.
		Nested schema for **simple_filter**:
			* `label_filters` - (List) The label filters to filter logs.
			Nested schema for **label_filters**:
				* `application_name` - (List) Filter by application names.
				  * Constraints: The maximum length is `4096` items. The minimum length is `0` items.
				Nested schema for **application_name**:
					* `operation` - (String) The operation to perform on the label value.
					  * Constraints: Allowable values are: `is_or_unspecified`, `includes`, `ends_with`, `starts_with`.
					* `value` - (String) The value used to filter the label.
					  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^.*$/`.
				* `severities` - (List) Filter by log severities.
				  * Constraints: Allowable list items are: `verbose_unspecified`, `debug`, `info`, `warning`, `error`, `critical`. The maximum length is `4096` items. The minimum length is `0` items.
				* `subsystem_name` - (List) Filter by subsystem names.
				  * Constraints: The maximum length is `4096` items. The minimum length is `0` items.
				Nested schema for **subsystem_name**:
					* `operation` - (String) The operation to perform on the label value.
					  * Constraints: Allowable values are: `is_or_unspecified`, `includes`, `ends_with`, `starts_with`.
					* `value` - (String) The value used to filter the label.
					  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^.*$/`.
			* `lucene_query` - (String) The Lucene query to filter logs.
			  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$/`.
	* `notification_payload_filter` - (List) The filter to specify which fields are included in the notification payload.
	  * Constraints: The list items must match regular expression `/^[a-zA-Z0-9_.]+$/`. The maximum length is `4096` items. The minimum length is `0` items.
* `logs_new_value` - (List) Configuration for alerts triggered by new log values.
Nested schema for **logs_new_value**:
	* `logs_filter` - (List) The filter to match log entries for immediate alerts.
	Nested schema for **logs_filter**:
		* `simple_filter` - (List) A simple filter that uses a Lucene query and label filters.
		Nested schema for **simple_filter**:
			* `label_filters` - (List) The label filters to filter logs.
			Nested schema for **label_filters**:
				* `application_name` - (List) Filter by application names.
				  * Constraints: The maximum length is `4096` items. The minimum length is `0` items.
				Nested schema for **application_name**:
					* `operation` - (String) The operation to perform on the label value.
					  * Constraints: Allowable values are: `is_or_unspecified`, `includes`, `ends_with`, `starts_with`.
					* `value` - (String) The value used to filter the label.
					  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^.*$/`.
				* `severities` - (List) Filter by log severities.
				  * Constraints: Allowable list items are: `verbose_unspecified`, `debug`, `info`, `warning`, `error`, `critical`. The maximum length is `4096` items. The minimum length is `0` items.
				* `subsystem_name` - (List) Filter by subsystem names.
				  * Constraints: The maximum length is `4096` items. The minimum length is `0` items.
				Nested schema for **subsystem_name**:
					* `operation` - (String) The operation to perform on the label value.
					  * Constraints: Allowable values are: `is_or_unspecified`, `includes`, `ends_with`, `starts_with`.
					* `value` - (String) The value used to filter the label.
					  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^.*$/`.
			* `lucene_query` - (String) The Lucene query to filter logs.
			  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$/`.
	* `notification_payload_filter` - (List) The filter to specify which fields are included in the notification payload.
	  * Constraints: The list items must match regular expression `/^[a-zA-Z0-9_.]+$/`. The maximum length is `4096` items. The minimum length is `0` items.
	* `rules` - (List) The condition rules for the log new value alert.
	  * Constraints: The maximum length is `1` item. The minimum length is `1` item.
	Nested schema for **rules**:
		* `condition` - (List) The condition for detecting new values in logs.
		Nested schema for **condition**:
			* `keypath_to_track` - (String) The keypath to track for new values.
			  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$/`.
			* `time_window` - (List) The time window for detecting new values.
			Nested schema for **time_window**:
				* `logs_new_value_time_window_specific_value` - (String) A time window defined by a specific value.
				  * Constraints: Allowable values are: `hours_12_or_unspecified`, `hours_24`, `hours_48`, `hours_72`, `week_1`, `month_1`, `months_2`, `months_3`.
* `logs_ratio_threshold` - (List) Configuration for the log-based ratio threshold alerts.
Nested schema for **logs_ratio_threshold**:
	* `condition_type` - (String) The condition type for the alert.
	  * Constraints: Allowable values are: `more_than_or_unspecified`, `less_than`.
	* `denominator` - (List) The filter to match log entries for immediate alerts.
	Nested schema for **denominator**:
		* `simple_filter` - (List) A simple filter that uses a Lucene query and label filters.
		Nested schema for **simple_filter**:
			* `label_filters` - (List) The label filters to filter logs.
			Nested schema for **label_filters**:
				* `application_name` - (List) Filter by application names.
				  * Constraints: The maximum length is `4096` items. The minimum length is `0` items.
				Nested schema for **application_name**:
					* `operation` - (String) The operation to perform on the label value.
					  * Constraints: Allowable values are: `is_or_unspecified`, `includes`, `ends_with`, `starts_with`.
					* `value` - (String) The value used to filter the label.
					  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^.*$/`.
				* `severities` - (List) Filter by log severities.
				  * Constraints: Allowable list items are: `verbose_unspecified`, `debug`, `info`, `warning`, `error`, `critical`. The maximum length is `4096` items. The minimum length is `0` items.
				* `subsystem_name` - (List) Filter by subsystem names.
				  * Constraints: The maximum length is `4096` items. The minimum length is `0` items.
				Nested schema for **subsystem_name**:
					* `operation` - (String) The operation to perform on the label value.
					  * Constraints: Allowable values are: `is_or_unspecified`, `includes`, `ends_with`, `starts_with`.
					* `value` - (String) The value used to filter the label.
					  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^.*$/`.
			* `lucene_query` - (String) The Lucene query to filter logs.
			  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$/`.
	* `denominator_alias` - (String) The alias for the denominator filter, used for display purposes.
	  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$/`.
	* `evaluation_delay_ms` - (Integer) The delay in milliseconds before evaluating the alert condition.
	  * Constraints: The maximum value is `2147483647`. The minimum value is `0`.
	* `group_by_for` - (String) The group by settings for the numerator and denominator filters.
	  * Constraints: Allowable values are: `both_or_unspecified`, `numerator_only`, `denumerator_only`.
	* `ignore_infinity` - (Boolean) Determine whether to ignore an infinity result or not. If true, alert is not triggered. When the value of second query is 0, the result of the ratio will be infinity.
	* `notification_payload_filter` - (List) The filter to specify which fields are included in the notification payload.
	  * Constraints: The list items must match regular expression `/^[a-zA-Z0-9_.]+$/`. The maximum length is `4096` items. The minimum length is `0` items.
	* `numerator` - (List) The filter to match log entries for immediate alerts.
	Nested schema for **numerator**:
		* `simple_filter` - (List) A simple filter that uses a Lucene query and label filters.
		Nested schema for **simple_filter**:
			* `label_filters` - (List) The label filters to filter logs.
			Nested schema for **label_filters**:
				* `application_name` - (List) Filter by application names.
				  * Constraints: The maximum length is `4096` items. The minimum length is `0` items.
				Nested schema for **application_name**:
					* `operation` - (String) The operation to perform on the label value.
					  * Constraints: Allowable values are: `is_or_unspecified`, `includes`, `ends_with`, `starts_with`.
					* `value` - (String) The value used to filter the label.
					  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^.*$/`.
				* `severities` - (List) Filter by log severities.
				  * Constraints: Allowable list items are: `verbose_unspecified`, `debug`, `info`, `warning`, `error`, `critical`. The maximum length is `4096` items. The minimum length is `0` items.
				* `subsystem_name` - (List) Filter by subsystem names.
				  * Constraints: The maximum length is `4096` items. The minimum length is `0` items.
				Nested schema for **subsystem_name**:
					* `operation` - (String) The operation to perform on the label value.
					  * Constraints: Allowable values are: `is_or_unspecified`, `includes`, `ends_with`, `starts_with`.
					* `value` - (String) The value used to filter the label.
					  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^.*$/`.
			* `lucene_query` - (String) The Lucene query to filter logs.
			  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$/`.
	* `numerator_alias` - (String) The alias for the numerator filter, used for display purposes.
	  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$/`.
	* `rules` - (List) The condition rules for the ratio alert.
	  * Constraints: The maximum length is `5` items. The minimum length is `1` item.
	Nested schema for **rules**:
		* `condition` - (List) The condition for the ratio alert.
		Nested schema for **condition**:
			* `threshold` - (Float) The threshold value for the alert condition.
			* `time_window` - (List) The time window for the alert condition.
			Nested schema for **time_window**:
				* `logs_ratio_time_window_specific_value` - (String) Specifies the time window for the ratio alert.
				  * Constraints: Allowable values are: `minutes_5_or_unspecified`, `minutes_10`, `minutes_15`, `minutes_30`, `hour_1`, `hours_2`, `hours_4`, `hours_6`, `hours_12`, `hours_24`, `hours_36`.
		* `override` - (List) The override settings for the alert.
		Nested schema for **override**:
			* `priority` - (String) The priority of the alert definition.
			  * Constraints: Allowable values are: `p5_or_unspecified`, `p4`, `p3`, `p2`, `p1`.
	* `undetected_values_management` - (List) Configuration for handling the undetected values in the alert.
	Nested schema for **undetected_values_management**:
		* `auto_retire_timeframe` - (String) The timeframe for auto-retiring the alert when undetected values are detected.
		  * Constraints: Allowable values are: `never_or_unspecified`, `minutes_5`, `minutes_10`, `hour_1`, `hours_2`, `hours_6`, `hours_12`, `hours_24`.
		* `trigger_undetected_values` - (Boolean) Should trigger the alert when undetected values are detected. If true, alert is triggered.
* `logs_threshold` - (List) Configuration for the log-based threshold alerts.
Nested schema for **logs_threshold**:
	* `condition_type` - (String) The condition type for the alert.
	  * Constraints: Allowable values are: `more_than_or_unspecified`, `less_than`.
	* `evaluation_delay_ms` - (Integer) The delay in milliseconds before evaluating the alert condition.
	  * Constraints: The maximum value is `2147483647`. The minimum value is `0`.
	* `logs_filter` - (List) The filter to match log entries for immediate alerts.
	Nested schema for **logs_filter**:
		* `simple_filter` - (List) A simple filter that uses a Lucene query and label filters.
		Nested schema for **simple_filter**:
			* `label_filters` - (List) The label filters to filter logs.
			Nested schema for **label_filters**:
				* `application_name` - (List) Filter by application names.
				  * Constraints: The maximum length is `4096` items. The minimum length is `0` items.
				Nested schema for **application_name**:
					* `operation` - (String) The operation to perform on the label value.
					  * Constraints: Allowable values are: `is_or_unspecified`, `includes`, `ends_with`, `starts_with`.
					* `value` - (String) The value used to filter the label.
					  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^.*$/`.
				* `severities` - (List) Filter by log severities.
				  * Constraints: Allowable list items are: `verbose_unspecified`, `debug`, `info`, `warning`, `error`, `critical`. The maximum length is `4096` items. The minimum length is `0` items.
				* `subsystem_name` - (List) Filter by subsystem names.
				  * Constraints: The maximum length is `4096` items. The minimum length is `0` items.
				Nested schema for **subsystem_name**:
					* `operation` - (String) The operation to perform on the label value.
					  * Constraints: Allowable values are: `is_or_unspecified`, `includes`, `ends_with`, `starts_with`.
					* `value` - (String) The value used to filter the label.
					  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^.*$/`.
			* `lucene_query` - (String) The Lucene query to filter logs.
			  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$/`.
	* `notification_payload_filter` - (List) The filter to specify which fields are included in the notification payload.
	  * Constraints: The list items must match regular expression `/^[a-zA-Z0-9_.]+$/`. The maximum length is `4096` items. The minimum length is `0` items.
	* `rules` - (List) The condition rules for the threshold alert.
	  * Constraints: The maximum length is `5` items. The minimum length is `1` item.
	Nested schema for **rules**:
		* `condition` - (List) The condition for the threshold alert.
		Nested schema for **condition**:
			* `threshold` - (Float) The threshold value for the alert condition.
			* `time_window` - (List) The time window for the alert condition.
			Nested schema for **time_window**:
				* `logs_time_window_specific_value` - (String) The time window defined for an alert to be triggered.
				  * Constraints: Allowable values are: `minutes_5_or_unspecified`, `minutes_10`, `minutes_20`, `minutes_15`, `minutes_30`, `hour_1`, `hours_2`, `hours_4`, `hours_6`, `hours_12`, `hours_24`, `hours_36`.
		* `override` - (List) The override settings for the alert.
		Nested schema for **override**:
			* `priority` - (String) The priority of the alert definition.
			  * Constraints: Allowable values are: `p5_or_unspecified`, `p4`, `p3`, `p2`, `p1`.
	* `undetected_values_management` - (List) Configuration for handling the undetected values in the alert.
	Nested schema for **undetected_values_management**:
		* `auto_retire_timeframe` - (String) The timeframe for auto-retiring the alert when undetected values are detected.
		  * Constraints: Allowable values are: `never_or_unspecified`, `minutes_5`, `minutes_10`, `hour_1`, `hours_2`, `hours_6`, `hours_12`, `hours_24`.
		* `trigger_undetected_values` - (Boolean) Should trigger the alert when undetected values are detected. If true, alert is triggered.
* `logs_time_relative_threshold` - (List) Configuration for time-relative log threshold alerts.
Nested schema for **logs_time_relative_threshold**:
	* `condition_type` - (String) The delay in milliseconds before evaluating the alert condition.
	  * Constraints: Allowable values are: `more_than_or_unspecified`, `less_than`.
	* `evaluation_delay_ms` - (Integer) The delay in milliseconds before evaluating the alert condition.
	  * Constraints: The maximum value is `2147483647`. The minimum value is `0`.
	* `ignore_infinity` - (Boolean) Ignore infinity values in the alert.
	* `logs_filter` - (List) The filter to match log entries for immediate alerts.
	Nested schema for **logs_filter**:
		* `simple_filter` - (List) A simple filter that uses a Lucene query and label filters.
		Nested schema for **simple_filter**:
			* `label_filters` - (List) The label filters to filter logs.
			Nested schema for **label_filters**:
				* `application_name` - (List) Filter by application names.
				  * Constraints: The maximum length is `4096` items. The minimum length is `0` items.
				Nested schema for **application_name**:
					* `operation` - (String) The operation to perform on the label value.
					  * Constraints: Allowable values are: `is_or_unspecified`, `includes`, `ends_with`, `starts_with`.
					* `value` - (String) The value used to filter the label.
					  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^.*$/`.
				* `severities` - (List) Filter by log severities.
				  * Constraints: Allowable list items are: `verbose_unspecified`, `debug`, `info`, `warning`, `error`, `critical`. The maximum length is `4096` items. The minimum length is `0` items.
				* `subsystem_name` - (List) Filter by subsystem names.
				  * Constraints: The maximum length is `4096` items. The minimum length is `0` items.
				Nested schema for **subsystem_name**:
					* `operation` - (String) The operation to perform on the label value.
					  * Constraints: Allowable values are: `is_or_unspecified`, `includes`, `ends_with`, `starts_with`.
					* `value` - (String) The value used to filter the label.
					  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^.*$/`.
			* `lucene_query` - (String) The Lucene query to filter logs.
			  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$/`.
	* `notification_payload_filter` - (List) The filter to specify which fields are included in the notification payload.
	  * Constraints: The list items must match regular expression `/^[a-zA-Z0-9_.]+$/`. The maximum length is `4096` items. The minimum length is `0` items.
	* `rules` - (List) The condition rules for the time-relative alert.
	  * Constraints: The maximum length is `5` items. The minimum length is `1` item.
	Nested schema for **rules**:
		* `condition` - (List) The condition for the time-relative alert.
		Nested schema for **condition**:
			* `compared_to` - (String) The time frame to compare the current value against.
			  * Constraints: Allowable values are: `previous_hour_or_unspecified`, `same_hour_yesterday`, `same_hour_last_week`, `yesterday`, `same_day_last_week`, `same_day_last_month`.
			* `threshold` - (Float) The threshold value for the alert condition.
		* `override` - (List) The override settings for the alert.
		Nested schema for **override**:
			* `priority` - (String) The priority of the alert definition.
			  * Constraints: Allowable values are: `p5_or_unspecified`, `p4`, `p3`, `p2`, `p1`.
	* `undetected_values_management` - (List) Configuration for handling the undetected values in the alert.
	Nested schema for **undetected_values_management**:
		* `auto_retire_timeframe` - (String) The timeframe for auto-retiring the alert when undetected values are detected.
		  * Constraints: Allowable values are: `never_or_unspecified`, `minutes_5`, `minutes_10`, `hour_1`, `hours_2`, `hours_6`, `hours_12`, `hours_24`.
		* `trigger_undetected_values` - (Boolean) Should trigger the alert when undetected values are detected. If true, alert is triggered.
* `logs_unique_count` - (List) Configuration for alerts based on unique log value counts.
Nested schema for **logs_unique_count**:
	* `logs_filter` - (List) The filter to match log entries for immediate alerts.
	Nested schema for **logs_filter**:
		* `simple_filter` - (List) A simple filter that uses a Lucene query and label filters.
		Nested schema for **simple_filter**:
			* `label_filters` - (List) The label filters to filter logs.
			Nested schema for **label_filters**:
				* `application_name` - (List) Filter by application names.
				  * Constraints: The maximum length is `4096` items. The minimum length is `0` items.
				Nested schema for **application_name**:
					* `operation` - (String) The operation to perform on the label value.
					  * Constraints: Allowable values are: `is_or_unspecified`, `includes`, `ends_with`, `starts_with`.
					* `value` - (String) The value used to filter the label.
					  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^.*$/`.
				* `severities` - (List) Filter by log severities.
				  * Constraints: Allowable list items are: `verbose_unspecified`, `debug`, `info`, `warning`, `error`, `critical`. The maximum length is `4096` items. The minimum length is `0` items.
				* `subsystem_name` - (List) Filter by subsystem names.
				  * Constraints: The maximum length is `4096` items. The minimum length is `0` items.
				Nested schema for **subsystem_name**:
					* `operation` - (String) The operation to perform on the label value.
					  * Constraints: Allowable values are: `is_or_unspecified`, `includes`, `ends_with`, `starts_with`.
					* `value` - (String) The value used to filter the label.
					  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^.*$/`.
			* `lucene_query` - (String) The Lucene query to filter logs.
			  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$/`.
	* `max_unique_count_per_group_by_key` - (String) The maximum unique count per group by key.
	  * Constraints: The maximum length is `19` characters. The minimum length is `1` character. The value must match regular expression `/^-?\\d{1,19}$/`.
	* `notification_payload_filter` - (List) The filter to specify which fields are included in the notification payload.
	  * Constraints: The list items must match regular expression `/^[a-zA-Z0-9_.]+$/`. The maximum length is `4096` items. The minimum length is `0` items.
	* `rules` - (List) Rules defining the conditions for the unique count alert.
	  * Constraints: The maximum length is `1` item. The minimum length is `1` item.
	Nested schema for **rules**:
		* `condition` - (List) The condition for detecting unique counts in logs.
		Nested schema for **condition**:
			* `max_unique_count` - (String) The maximum unique count for the alert condition.
			  * Constraints: The maximum length is `19` characters. The minimum length is `1` character. The value must match regular expression `/^-?\\d{1,19}$/`.
			* `time_window` - (List) The time window for the unique count alert.
			Nested schema for **time_window**:
				* `logs_unique_value_time_window_specific_value` - (String) A time window defined by a specific value.
				  * Constraints: Allowable values are: `minute_1_or_unspecified`, `minutes_15`, `minutes_20`, `minutes_30`, `hours_1`, `hours_2`, `hours_4`, `hours_6`, `hours_12`, `hours_24`, `minutes_5`, `minutes_10`, `hours_36`.
	* `unique_count_keypath` - (String) The keypath in the logs to be used for unique count.
	  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$/`.
* `metric_anomaly` - (List) Configuration for metric-based anomaly detection alerts.
Nested schema for **metric_anomaly**:
	* `anomaly_alert_settings` - (List) The anomaly alert settings configuration.
	Nested schema for **anomaly_alert_settings**:
		* `percentage_of_deviation` - (Float) The percentage of deviation from the baseline when the alert is triggered.
	* `condition_type` - (String) The condition type for the alert.
	  * Constraints: Allowable values are: `more_than_usual_or_unspecified`, `less_than_usual`.
	* `evaluation_delay_ms` - (Integer) The delay in milliseconds before evaluating the alert condition.
	  * Constraints: The maximum value is `2147483647`. The minimum value is `0`.
	* `metric_filter` - (List) The filter to match metric entries for the alert.
	Nested schema for **metric_filter**:
		* `promql` - (String) The filter is a PromQL expression.
		  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$/`.
	* `rules` - (List) The condition rules for the metric anomaly alert.
	  * Constraints: The maximum length is `1` item. The minimum length is `1` item.
	Nested schema for **rules**:
		* `condition` - (List) The condition for the metric anomaly alert.
		Nested schema for **condition**:
			* `for_over_pct` - (Integer) The percentage of the metric values that must exceed the threshold to trigger the alert.
			  * Constraints: The maximum value is `4294967295`. The minimum value is `0`.
			* `min_non_null_values_pct` - (Integer) The percentage of non-null values required to trigger the alert.
			  * Constraints: The maximum value is `4294967295`. The minimum value is `0`.
			* `of_the_last` - (List) The time window for the alert condition.
			Nested schema for **of_the_last**:
				* `metric_time_window_dynamic_duration` - (String) The time window as a dynamic value.
				  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$/`.
				* `metric_time_window_specific_value` - (String) The time window as a specific value.
				  * Constraints: Allowable values are: `minutes_1_or_unspecified`, `minutes_5`, `minutes_10`, `minutes_15`, `minutes_30`, `hour_1`, `hours_2`, `hours_4`, `hours_6`, `hours_12`, `hours_24`, `minutes_20`, `hours_36`.
			* `threshold` - (Float) The threshold value for the alert condition.
* `metric_threshold` - (List) Configuration for metric-based threshold alerts.
Nested schema for **metric_threshold**:
	* `condition_type` - (String) The type of the alert condition.
	  * Constraints: Allowable values are: `more_than_or_unspecified`, `less_than`, `more_than_or_equals`, `less_than_or_equals`.
	* `evaluation_delay_ms` - (Integer) The delay in milliseconds before evaluating the alert condition.
	  * Constraints: The maximum value is `2147483647`. The minimum value is `0`.
	* `metric_filter` - (List) The filter to match metric entries for the alert.
	Nested schema for **metric_filter**:
		* `promql` - (String) The filter is a PromQL expression.
		  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$/`.
	* `missing_values` - (List) Configuration for handling missing values in the alert. Only one of `replace_with_zero` or `min_non_null_value_pct` is supported.
	Nested schema for **missing_values**:
		* `min_non_null_values_pct` - (Integer) If set, specifies the minimum percentage of non-null values required for the alert to be triggered.
		  * Constraints: The maximum value is `4294967295`. The minimum value is `0`.
		* `replace_with_zero` - (Boolean) If set to true, missing values will be replaced with zero.
	* `rules` - (List) The condition rules for the metric threshold alert.
	  * Constraints: The maximum length is `5` items. The minimum length is `1` item.
	Nested schema for **rules**:
		* `condition` - (List) The condition for the metric threshold alert.
		Nested schema for **condition**:
			* `for_over_pct` - (Integer) The percentage of values that must exceed the threshold to trigger the alert.
			  * Constraints: The maximum value is `4294967295`. The minimum value is `0`.
			* `of_the_last` - (List) The time window for the alert condition.
			Nested schema for **of_the_last**:
				* `metric_time_window_dynamic_duration` - (String) The time window as a dynamic value.
				  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$/`.
				* `metric_time_window_specific_value` - (String) The time window as a specific value.
				  * Constraints: Allowable values are: `minutes_1_or_unspecified`, `minutes_5`, `minutes_10`, `minutes_15`, `minutes_30`, `hour_1`, `hours_2`, `hours_4`, `hours_6`, `hours_12`, `hours_24`, `minutes_20`, `hours_36`.
			* `threshold` - (Float) The threshold value for the alert condition.
		* `override` - (List) The override settings for the alert.
		Nested schema for **override**:
			* `priority` - (String) The priority of the alert definition.
			  * Constraints: Allowable values are: `p5_or_unspecified`, `p4`, `p3`, `p2`, `p1`.
	* `undetected_values_management` - (List) Configuration for handling the undetected values in the alert.
	Nested schema for **undetected_values_management**:
		* `auto_retire_timeframe` - (String) The timeframe for auto-retiring the alert when undetected values are detected.
		  * Constraints: Allowable values are: `never_or_unspecified`, `minutes_5`, `minutes_10`, `hour_1`, `hours_2`, `hours_6`, `hours_12`, `hours_24`.
		* `trigger_undetected_values` - (Boolean) Should trigger the alert when undetected values are detected. If true, alert is triggered.
* `name` - (String) The name of the alert definition.
  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$/`.
* `notification_group` - (List) Primary notification group for alert events.
Nested schema for **notification_group**:
	* `group_by_keys` - (List) Group the alerts by these keys.
	  * Constraints: The list items must match regular expression `/^[a-zA-Z0-9_.]+$/`. The maximum length is `4096` items. The minimum length is `0` items.
	* `webhooks` - (List) The settings for webhooks associated with the alert definition.
	  * Constraints: The maximum length is `4096` items. The minimum length is `0` items.
	Nested schema for **webhooks**:
		* `integration` - (List) The integration type for webhook notifications.
		Nested schema for **integration**:
			* `integration_id` - (Integer) The integration ID for the notification.
			  * Constraints: The maximum value is `4294967295`. The minimum value is `0`.
		* `minutes` - (Integer) The time in minutes before the notification is sent.
		  * Constraints: The maximum value is `4294967295`. The minimum value is `0`.
		* `notify_on` - (String) Indicate if the alert should be triggered or triggered and resolved.
		  * Constraints: Allowable values are: `triggered_only_unspecified`, `triggered_and_resolved`.
* `phantom_mode` - (Boolean) Whether the alert is in phantom mode (creating incidents or not).
* `priority` - (String) The priority of the alert definition.
  * Constraints: Allowable values are: `p5_or_unspecified`, `p4`, `p3`, `p2`, `p1`.
* `type` - (String) Alert type.
  * Constraints: Allowable values are: `logs_immediate_or_unspecified`, `logs_threshold`, `logs_anomaly`, `logs_ratio_threshold`, `logs_new_value`, `logs_unique_count`, `logs_time_relative_threshold`, `metric_threshold`, `metric_anomaly`, `flow`.
* `updated_time` - (String) The time when the alert definition was last updated.

