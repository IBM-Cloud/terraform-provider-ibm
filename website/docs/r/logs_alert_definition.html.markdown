---
layout: "ibm"
page_title: "IBM : ibm_logs_alert_definition"
description: |-
  Manages logs_alert_definition.
subcategory: "Cloud Logs"
---

# ibm_logs_alert_definition

Create, update, and delete logs_alert_definitions with this resource.

## Example Usage

```hcl
resource "ibm_logs_alert_definition" "logs_alert_definition_instance" {
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
  deleted = false
  description = "Example of unique count alert from terraform"
  enabled = true
  entity_labels = {"key":"value"}
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
  incidents_settings {
		notify_on = "triggered_and_resolved"
		minutes = 30
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
  name = "Unique count alert"
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
  phantom_mode = false
  priority = "p1"
  type = "flow"
}
```

## Argument Reference

You can specify the following arguments for this resource.

* `active_on` - (Optional, List) Defining when the alert is active.
Nested schema for **active_on**:
	* `day_of_week` - (Required, List) Days of the week when the alert is active.
	  * Constraints: Allowable list items are: `monday_or_unspecified`, `tuesday`, `wednesday`, `thursday`, `friday`, `saturday`, `sunday`. The maximum length is `7` items. The minimum length is `1` item.
	* `end_time` - (Required, List) Start time of the alert activity.
	Nested schema for **end_time**:
		* `hours` - (Optional, Integer) Hours of day in 24 hour format. Should be from 0 to 23.
		  * Constraints: The maximum value is `23`. The minimum value is `0`.
		* `minutes` - (Optional, Integer) Minutes of hour of day. Must be from 0 to 59.
		  * Constraints: The maximum value is `59`. The minimum value is `0`.
	* `start_time` - (Required, List) Start time of the alert activity.
	Nested schema for **start_time**:
		* `hours` - (Optional, Integer) Hours of day in 24 hour format. Should be from 0 to 23.
		  * Constraints: The maximum value is `23`. The minimum value is `0`.
		* `minutes` - (Optional, Integer) Minutes of hour of day. Must be from 0 to 59.
		  * Constraints: The maximum value is `59`. The minimum value is `0`.
* `deleted` - (Optional, Boolean) Whether the alert has been marked as deleted.
* `description` - (Optional, String) A detailed description of what the alert monitors and when it triggers.
  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$/`.
* `enabled` - (Optional, Boolean) Whether the alert is currently active and monitoring.
* `entity_labels` - (Optional, Map) Labels used to identify and categorize the alert entity.
* `flow` - (Optional, List) Configuration for flow-based alerts.
Nested schema for **flow**:
	* `enforce_suppression` - (Optional, Boolean) Whether to enforce suppression for the flow alert.
	* `stages` - (Required, List) The stages of the flow alert.
	  * Constraints: The maximum length is `4096` items. The minimum length is `1` item.
	Nested schema for **stages**:
		* `flow_stages_groups` - (Required, List) Flow stages groups.
		Nested schema for **flow_stages_groups**:
			* `groups` - (Required, List) The groups of stages in the flow alert.
			  * Constraints: The maximum length is `4096` items. The minimum length is `1` item.
			Nested schema for **groups**:
				* `alert_defs` - (Required, List) The alert definitions for the flow stage group.
				  * Constraints: The maximum length is `4096` items. The minimum length is `1` item.
				Nested schema for **alert_defs**:
					* `id` - (Required, String) The alert definition ID.
					  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$/`.
					* `not` - (Optional, Boolean) Whether to negate the alert definition or not.
				* `alerts_op` - (Required, String) The logical operation to apply to the alerts in the group.
				  * Constraints: Allowable values are: `and_or_unspecified`, `or`.
				* `next_op` - (Required, String) The logical operation to apply to the next stage.
				  * Constraints: Allowable values are: `and_or_unspecified`, `or`.
		* `timeframe_ms` - (Required, String) The timeframe for the flow alert in milliseconds.
		  * Constraints: The maximum length is `19` characters. The minimum length is `1` character. The value must match regular expression `/^-?\\d{1,19}$/`.
		* `timeframe_type` - (Required, String) The type of timeframe for the flow alert.
		  * Constraints: Allowable values are: `unspecified`, `up_to`.
* `group_by_keys` - (Required, List) Keys used to group and aggregate alert data.
  * Constraints: The list items must match regular expression `/^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$/`. The maximum length is `2` items. The minimum length is `0` items.
* `incidents_settings` - (Optional, List) Incident creation and management settings.
Nested schema for **incidents_settings**:
	* `minutes` - (Optional, Integer) The time in minutes before the alert can be retriggered.
	  * Constraints: The maximum value is `4294967295`. The minimum value is `0`.
	* `notify_on` - (Optional, String) The condition to notify about the alert.
	  * Constraints: Allowable values are: `triggered_only_unspecified`, `triggered_and_resolved`.
* `logs_anomaly` - (Optional, List) Configuration for log-based anomaly detection alerts.
Nested schema for **logs_anomaly**:
	* `anomaly_alert_settings` - (Optional, List) The anomaly alert settings configuration.
	Nested schema for **anomaly_alert_settings**:
		* `percentage_of_deviation` - (Optional, Float) The percentage of deviation from the baseline for triggering the alert.
	* `condition_type` - (Required, String) The type of condition for the alert.
	  * Constraints: Allowable values are: `more_than_usual_or_unspecified`.
	* `evaluation_delay_ms` - (Optional, Integer) The delay in milliseconds before evaluating the alert condition.
	  * Constraints: The maximum value is `2147483647`. The minimum value is `0`.
	* `logs_filter` - (Optional, List) The filter to match log entries for immediate alerts.
	Nested schema for **logs_filter**:
		* `simple_filter` - (Optional, List) A simple filter that uses a Lucene query and label filters.
		Nested schema for **simple_filter**:
			* `label_filters` - (Optional, List) The label filters to filter logs.
			Nested schema for **label_filters**:
				* `application_name` - (Required, List) Filter by application names.
				  * Constraints: The maximum length is `4096` items. The minimum length is `0` items.
				Nested schema for **application_name**:
					* `operation` - (Required, String) The operation to perform on the label value.
					  * Constraints: Allowable values are: `is_or_unspecified`, `includes`, `ends_with`, `starts_with`.
					* `value` - (Optional, String) / The value of the label to filter by.
					  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^.*$/`.
				* `severities` - (Required, List) Filter by log severities.
				  * Constraints: Allowable list items are: `verbose_unspecified`, `debug`, `info`, `warning`, `error`, `critical`. The maximum length is `4096` items. The minimum length is `0` items.
				* `subsystem_name` - (Required, List) Filter by subsystem names.
				  * Constraints: The maximum length is `4096` items. The minimum length is `0` items.
				Nested schema for **subsystem_name**:
					* `operation` - (Required, String) The operation to perform on the label value.
					  * Constraints: Allowable values are: `is_or_unspecified`, `includes`, `ends_with`, `starts_with`.
					* `value` - (Optional, String) / The value of the label to filter by.
					  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^.*$/`.
			* `lucene_query` - (Optional, String) The Lucene query to filter logs.
			  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$/`.
	* `notification_payload_filter` - (Optional, List) The notification payload filter to specify which fields to include in the notification.
	  * Constraints: The list items must match regular expression `/^[a-zA-Z0-9_.]+$/`. The maximum length is `4096` items. The minimum length is `0` items.
	* `rules` - (Required, List) The rules for the log anomaly alert.
	  * Constraints: The maximum length is `1` item. The minimum length is `1` item.
	Nested schema for **rules**:
		* `condition` - (Required, List) The condition for the anomaly alert.
		Nested schema for **condition**:
			* `minimum_threshold` - (Required, Float) The threshold value for the alert condition.
			* `time_window` - (Required, List) The time window for the alert condition.
			Nested schema for **time_window**:
				* `logs_time_window_specific_value` - (Required, String) A time window defined by a specific value.
				  * Constraints: Allowable values are: `minutes_5_or_unspecified`, `minutes_10`, `minutes_20`, `minutes_15`, `minutes_30`, `hour_1`, `hours_2`, `hours_4`, `hours_6`, `hours_12`, `hours_24`, `hours_36`.
* `logs_immediate` - (Optional, List) Configuration for immediate log-based alerts.
Nested schema for **logs_immediate**:
	* `logs_filter` - (Optional, List) The filter to match log entries for immediate alerts.
	Nested schema for **logs_filter**:
		* `simple_filter` - (Optional, List) A simple filter that uses a Lucene query and label filters.
		Nested schema for **simple_filter**:
			* `label_filters` - (Optional, List) The label filters to filter logs.
			Nested schema for **label_filters**:
				* `application_name` - (Required, List) Filter by application names.
				  * Constraints: The maximum length is `4096` items. The minimum length is `0` items.
				Nested schema for **application_name**:
					* `operation` - (Required, String) The operation to perform on the label value.
					  * Constraints: Allowable values are: `is_or_unspecified`, `includes`, `ends_with`, `starts_with`.
					* `value` - (Optional, String) / The value of the label to filter by.
					  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^.*$/`.
				* `severities` - (Required, List) Filter by log severities.
				  * Constraints: Allowable list items are: `verbose_unspecified`, `debug`, `info`, `warning`, `error`, `critical`. The maximum length is `4096` items. The minimum length is `0` items.
				* `subsystem_name` - (Required, List) Filter by subsystem names.
				  * Constraints: The maximum length is `4096` items. The minimum length is `0` items.
				Nested schema for **subsystem_name**:
					* `operation` - (Required, String) The operation to perform on the label value.
					  * Constraints: Allowable values are: `is_or_unspecified`, `includes`, `ends_with`, `starts_with`.
					* `value` - (Optional, String) / The value of the label to filter by.
					  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^.*$/`.
			* `lucene_query` - (Optional, String) The Lucene query to filter logs.
			  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$/`.
	* `notification_payload_filter` - (Optional, List) The filter to specify which fields to include in the notification payload.
	  * Constraints: The list items must match regular expression `/^[a-zA-Z0-9_.]+$/`. The maximum length is `4096` items. The minimum length is `0` items.
* `logs_new_value` - (Optional, List) Configuration for alerts triggered by new log values.
Nested schema for **logs_new_value**:
	* `logs_filter` - (Optional, List) The filter to match log entries for immediate alerts.
	Nested schema for **logs_filter**:
		* `simple_filter` - (Optional, List) A simple filter that uses a Lucene query and label filters.
		Nested schema for **simple_filter**:
			* `label_filters` - (Optional, List) The label filters to filter logs.
			Nested schema for **label_filters**:
				* `application_name` - (Required, List) Filter by application names.
				  * Constraints: The maximum length is `4096` items. The minimum length is `0` items.
				Nested schema for **application_name**:
					* `operation` - (Required, String) The operation to perform on the label value.
					  * Constraints: Allowable values are: `is_or_unspecified`, `includes`, `ends_with`, `starts_with`.
					* `value` - (Optional, String) / The value of the label to filter by.
					  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^.*$/`.
				* `severities` - (Required, List) Filter by log severities.
				  * Constraints: Allowable list items are: `verbose_unspecified`, `debug`, `info`, `warning`, `error`, `critical`. The maximum length is `4096` items. The minimum length is `0` items.
				* `subsystem_name` - (Required, List) Filter by subsystem names.
				  * Constraints: The maximum length is `4096` items. The minimum length is `0` items.
				Nested schema for **subsystem_name**:
					* `operation` - (Required, String) The operation to perform on the label value.
					  * Constraints: Allowable values are: `is_or_unspecified`, `includes`, `ends_with`, `starts_with`.
					* `value` - (Optional, String) / The value of the label to filter by.
					  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^.*$/`.
			* `lucene_query` - (Optional, String) The Lucene query to filter logs.
			  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$/`.
	* `notification_payload_filter` - (Optional, List) The filter to specify which fields to include in the notification payload.
	  * Constraints: The list items must match regular expression `/^[a-zA-Z0-9_.]+$/`. The maximum length is `4096` items. The minimum length is `0` items.
	* `rules` - (Required, List) The rules for the log new value alert.
	  * Constraints: The maximum length is `1` item. The minimum length is `1` item.
	Nested schema for **rules**:
		* `condition` - (Required, List) The condition for detecting new values in logs.
		Nested schema for **condition**:
			* `keypath_to_track` - (Required, String) The keypath to track for new values.
			  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$/`.
			* `time_window` - (Required, List) The time window for detecting new values.
			Nested schema for **time_window**:
				* `logs_new_value_time_window_specific_value` - (Required, String) A time window defined by a specific value.
				  * Constraints: Allowable values are: `hours_12_or_unspecified`, `hours_24`, `hours_48`, `hours_72`, `week_1`, `month_1`, `months_2`, `months_3`.
* `logs_ratio_threshold` - (Optional, List) Configuration for log-based ratio threshold alerts.
Nested schema for **logs_ratio_threshold**:
	* `condition_type` - (Required, String) The type of condition for the alert.
	  * Constraints: Allowable values are: `more_than_or_unspecified`, `less_than`.
	* `denominator` - (Required, List) The filter to match log entries for immediate alerts.
	Nested schema for **denominator**:
		* `simple_filter` - (Optional, List) A simple filter that uses a Lucene query and label filters.
		Nested schema for **simple_filter**:
			* `label_filters` - (Optional, List) The label filters to filter logs.
			Nested schema for **label_filters**:
				* `application_name` - (Required, List) Filter by application names.
				  * Constraints: The maximum length is `4096` items. The minimum length is `0` items.
				Nested schema for **application_name**:
					* `operation` - (Required, String) The operation to perform on the label value.
					  * Constraints: Allowable values are: `is_or_unspecified`, `includes`, `ends_with`, `starts_with`.
					* `value` - (Optional, String) / The value of the label to filter by.
					  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^.*$/`.
				* `severities` - (Required, List) Filter by log severities.
				  * Constraints: Allowable list items are: `verbose_unspecified`, `debug`, `info`, `warning`, `error`, `critical`. The maximum length is `4096` items. The minimum length is `0` items.
				* `subsystem_name` - (Required, List) Filter by subsystem names.
				  * Constraints: The maximum length is `4096` items. The minimum length is `0` items.
				Nested schema for **subsystem_name**:
					* `operation` - (Required, String) The operation to perform on the label value.
					  * Constraints: Allowable values are: `is_or_unspecified`, `includes`, `ends_with`, `starts_with`.
					* `value` - (Optional, String) / The value of the label to filter by.
					  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^.*$/`.
			* `lucene_query` - (Optional, String) The Lucene query to filter logs.
			  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$/`.
	* `denominator_alias` - (Optional, String) The alias for the denominator filter, used for display purposes.
	  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$/`.
	* `evaluation_delay_ms` - (Optional, Integer) The delay in milliseconds before evaluating the alert condition.
	  * Constraints: The maximum value is `2147483647`. The minimum value is `0`.
	* `group_by_for` - (Required, String) The group by settings for the numerator and denominator filters.
	  * Constraints: Allowable values are: `both_or_unspecified`, `numerator_only`, `denumerator_only`.
	* `ignore_infinity` - (Optional, Boolean) The configuration for ignoring infinity values in the ratio.
	* `notification_payload_filter` - (Optional, List) The filter to specify which fields to include in the notification payload.
	  * Constraints: The list items must match regular expression `/^[a-zA-Z0-9_.]+$/`. The maximum length is `4096` items. The minimum length is `0` items.
	* `numerator` - (Required, List) The filter to match log entries for immediate alerts.
	Nested schema for **numerator**:
		* `simple_filter` - (Optional, List) A simple filter that uses a Lucene query and label filters.
		Nested schema for **simple_filter**:
			* `label_filters` - (Optional, List) The label filters to filter logs.
			Nested schema for **label_filters**:
				* `application_name` - (Required, List) Filter by application names.
				  * Constraints: The maximum length is `4096` items. The minimum length is `0` items.
				Nested schema for **application_name**:
					* `operation` - (Required, String) The operation to perform on the label value.
					  * Constraints: Allowable values are: `is_or_unspecified`, `includes`, `ends_with`, `starts_with`.
					* `value` - (Optional, String) / The value of the label to filter by.
					  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^.*$/`.
				* `severities` - (Required, List) Filter by log severities.
				  * Constraints: Allowable list items are: `verbose_unspecified`, `debug`, `info`, `warning`, `error`, `critical`. The maximum length is `4096` items. The minimum length is `0` items.
				* `subsystem_name` - (Required, List) Filter by subsystem names.
				  * Constraints: The maximum length is `4096` items. The minimum length is `0` items.
				Nested schema for **subsystem_name**:
					* `operation` - (Required, String) The operation to perform on the label value.
					  * Constraints: Allowable values are: `is_or_unspecified`, `includes`, `ends_with`, `starts_with`.
					* `value` - (Optional, String) / The value of the label to filter by.
					  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^.*$/`.
			* `lucene_query` - (Optional, String) The Lucene query to filter logs.
			  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$/`.
	* `numerator_alias` - (Optional, String) The alias for the numerator filter, used for display purposes.
	  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$/`.
	* `rules` - (Required, List) The rules for the ratio alert.
	  * Constraints: The maximum length is `5` items. The minimum length is `1` item.
	Nested schema for **rules**:
		* `condition` - (Required, List) The condition for the ratio alert.
		Nested schema for **condition**:
			* `threshold` - (Required, Float) The threshold value for the alert condition.
			* `time_window` - (Required, List) The time window for the alert condition.
			Nested schema for **time_window**:
				* `logs_ratio_time_window_specific_value` - (Required, String) Specifies the time window for the ratio alert.
				  * Constraints: Allowable values are: `minutes_5_or_unspecified`, `minutes_10`, `minutes_15`, `minutes_30`, `hour_1`, `hours_2`, `hours_4`, `hours_6`, `hours_12`, `hours_24`, `hours_36`.
		* `override` - (Required, List) The override settings for the alert.
		Nested schema for **override**:
			* `priority` - (Required, String) The priority of the alert definition.
			  * Constraints: Allowable values are: `p5_or_unspecified`, `p4`, `p3`, `p2`, `p1`.
	* `undetected_values_management` - (Optional, List) Configuration for handling the undetected values in the alert.
	Nested schema for **undetected_values_management**:
		* `auto_retire_timeframe` - (Required, String) The timeframe for auto-retiring the alert when undetected values are detected.
		  * Constraints: Allowable values are: `never_or_unspecified`, `minutes_5`, `minutes_10`, `hour_1`, `hours_2`, `hours_6`, `hours_12`, `hours_24`.
		* `trigger_undetected_values` - (Required, Boolean) Should trigger the alert when undetected values are detected.
* `logs_threshold` - (Optional, List) Configuration for log-based threshold alerts.
Nested schema for **logs_threshold**:
	* `condition_type` - (Required, String) The type of condition for the alert.
	  * Constraints: Allowable values are: `more_than_or_unspecified`, `less_than`.
	* `evaluation_delay_ms` - (Optional, Integer) The delay in milliseconds before evaluating the alert condition.
	  * Constraints: The maximum value is `2147483647`. The minimum value is `0`.
	* `logs_filter` - (Optional, List) The filter to match log entries for immediate alerts.
	Nested schema for **logs_filter**:
		* `simple_filter` - (Optional, List) A simple filter that uses a Lucene query and label filters.
		Nested schema for **simple_filter**:
			* `label_filters` - (Optional, List) The label filters to filter logs.
			Nested schema for **label_filters**:
				* `application_name` - (Required, List) Filter by application names.
				  * Constraints: The maximum length is `4096` items. The minimum length is `0` items.
				Nested schema for **application_name**:
					* `operation` - (Required, String) The operation to perform on the label value.
					  * Constraints: Allowable values are: `is_or_unspecified`, `includes`, `ends_with`, `starts_with`.
					* `value` - (Optional, String) / The value of the label to filter by.
					  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^.*$/`.
				* `severities` - (Required, List) Filter by log severities.
				  * Constraints: Allowable list items are: `verbose_unspecified`, `debug`, `info`, `warning`, `error`, `critical`. The maximum length is `4096` items. The minimum length is `0` items.
				* `subsystem_name` - (Required, List) Filter by subsystem names.
				  * Constraints: The maximum length is `4096` items. The minimum length is `0` items.
				Nested schema for **subsystem_name**:
					* `operation` - (Required, String) The operation to perform on the label value.
					  * Constraints: Allowable values are: `is_or_unspecified`, `includes`, `ends_with`, `starts_with`.
					* `value` - (Optional, String) / The value of the label to filter by.
					  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^.*$/`.
			* `lucene_query` - (Optional, String) The Lucene query to filter logs.
			  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$/`.
	* `notification_payload_filter` - (Optional, List) The filter to specify which fields to include in the notification payload.
	  * Constraints: The list items must match regular expression `/^[a-zA-Z0-9_.]+$/`. The maximum length is `4096` items. The minimum length is `0` items.
	* `rules` - (Required, List) The rules for the threshold alert.
	  * Constraints: The maximum length is `5` items. The minimum length is `1` item.
	Nested schema for **rules**:
		* `condition` - (Required, List) The condition for the threshold alert.
		Nested schema for **condition**:
			* `threshold` - (Required, Float) The threshold value for the alert condition.
			* `time_window` - (Required, List) The time window for the alert condition.
			Nested schema for **time_window**:
				* `logs_time_window_specific_value` - (Required, String) A time window defined by a specific value.
				  * Constraints: Allowable values are: `minutes_5_or_unspecified`, `minutes_10`, `minutes_20`, `minutes_15`, `minutes_30`, `hour_1`, `hours_2`, `hours_4`, `hours_6`, `hours_12`, `hours_24`, `hours_36`.
		* `override` - (Required, List) The override settings for the alert.
		Nested schema for **override**:
			* `priority` - (Required, String) The priority of the alert definition.
			  * Constraints: Allowable values are: `p5_or_unspecified`, `p4`, `p3`, `p2`, `p1`.
	* `undetected_values_management` - (Optional, List) Configuration for handling the undetected values in the alert.
	Nested schema for **undetected_values_management**:
		* `auto_retire_timeframe` - (Required, String) The timeframe for auto-retiring the alert when undetected values are detected.
		  * Constraints: Allowable values are: `never_or_unspecified`, `minutes_5`, `minutes_10`, `hour_1`, `hours_2`, `hours_6`, `hours_12`, `hours_24`.
		* `trigger_undetected_values` - (Required, Boolean) Should trigger the alert when undetected values are detected.
* `logs_time_relative_threshold` - (Optional, List) Configuration for time-relative log threshold alerts.
Nested schema for **logs_time_relative_threshold**:
	* `condition_type` - (Required, String) The delay in milliseconds before evaluating the alert condition.
	  * Constraints: Allowable values are: `more_than_or_unspecified`, `less_than`.
	* `evaluation_delay_ms` - (Optional, Integer) The delay in milliseconds before evaluating the alert condition.
	  * Constraints: The maximum value is `2147483647`. The minimum value is `0`.
	* `ignore_infinity` - (Optional, Boolean) Ignore infinity values in the alert.
	* `logs_filter` - (Optional, List) The filter to match log entries for immediate alerts.
	Nested schema for **logs_filter**:
		* `simple_filter` - (Optional, List) A simple filter that uses a Lucene query and label filters.
		Nested schema for **simple_filter**:
			* `label_filters` - (Optional, List) The label filters to filter logs.
			Nested schema for **label_filters**:
				* `application_name` - (Required, List) Filter by application names.
				  * Constraints: The maximum length is `4096` items. The minimum length is `0` items.
				Nested schema for **application_name**:
					* `operation` - (Required, String) The operation to perform on the label value.
					  * Constraints: Allowable values are: `is_or_unspecified`, `includes`, `ends_with`, `starts_with`.
					* `value` - (Optional, String) / The value of the label to filter by.
					  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^.*$/`.
				* `severities` - (Required, List) Filter by log severities.
				  * Constraints: Allowable list items are: `verbose_unspecified`, `debug`, `info`, `warning`, `error`, `critical`. The maximum length is `4096` items. The minimum length is `0` items.
				* `subsystem_name` - (Required, List) Filter by subsystem names.
				  * Constraints: The maximum length is `4096` items. The minimum length is `0` items.
				Nested schema for **subsystem_name**:
					* `operation` - (Required, String) The operation to perform on the label value.
					  * Constraints: Allowable values are: `is_or_unspecified`, `includes`, `ends_with`, `starts_with`.
					* `value` - (Optional, String) / The value of the label to filter by.
					  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^.*$/`.
			* `lucene_query` - (Optional, String) The Lucene query to filter logs.
			  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$/`.
	* `notification_payload_filter` - (Optional, List) The filter to specify which fields to include in the notification payload.
	  * Constraints: The list items must match regular expression `/^[a-zA-Z0-9_.]+$/`. The maximum length is `4096` items. The minimum length is `0` items.
	* `rules` - (Required, List) The rules for the time-relative alert.
	  * Constraints: The maximum length is `5` items. The minimum length is `1` item.
	Nested schema for **rules**:
		* `condition` - (Required, List) The condition for the time-relative alert.
		Nested schema for **condition**:
			* `compared_to` - (Required, String) The time frame to compare the current value against.
			  * Constraints: Allowable values are: `previous_hour_or_unspecified`, `same_hour_yesterday`, `same_hour_last_week`, `yesterday`, `same_day_last_week`, `same_day_last_month`.
			* `threshold` - (Required, Float) The threshold value for the alert condition.
		* `override` - (Required, List) The override settings for the alert.
		Nested schema for **override**:
			* `priority` - (Required, String) The priority of the alert definition.
			  * Constraints: Allowable values are: `p5_or_unspecified`, `p4`, `p3`, `p2`, `p1`.
	* `undetected_values_management` - (Optional, List) Configuration for handling the undetected values in the alert.
	Nested schema for **undetected_values_management**:
		* `auto_retire_timeframe` - (Required, String) The timeframe for auto-retiring the alert when undetected values are detected.
		  * Constraints: Allowable values are: `never_or_unspecified`, `minutes_5`, `minutes_10`, `hour_1`, `hours_2`, `hours_6`, `hours_12`, `hours_24`.
		* `trigger_undetected_values` - (Required, Boolean) Should trigger the alert when undetected values are detected.
* `logs_unique_count` - (Optional, List) Configuration for alerts based on unique log value counts.
Nested schema for **logs_unique_count**:
	* `logs_filter` - (Optional, List) The filter to match log entries for immediate alerts.
	Nested schema for **logs_filter**:
		* `simple_filter` - (Optional, List) A simple filter that uses a Lucene query and label filters.
		Nested schema for **simple_filter**:
			* `label_filters` - (Optional, List) The label filters to filter logs.
			Nested schema for **label_filters**:
				* `application_name` - (Required, List) Filter by application names.
				  * Constraints: The maximum length is `4096` items. The minimum length is `0` items.
				Nested schema for **application_name**:
					* `operation` - (Required, String) The operation to perform on the label value.
					  * Constraints: Allowable values are: `is_or_unspecified`, `includes`, `ends_with`, `starts_with`.
					* `value` - (Optional, String) / The value of the label to filter by.
					  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^.*$/`.
				* `severities` - (Required, List) Filter by log severities.
				  * Constraints: Allowable list items are: `verbose_unspecified`, `debug`, `info`, `warning`, `error`, `critical`. The maximum length is `4096` items. The minimum length is `0` items.
				* `subsystem_name` - (Required, List) Filter by subsystem names.
				  * Constraints: The maximum length is `4096` items. The minimum length is `0` items.
				Nested schema for **subsystem_name**:
					* `operation` - (Required, String) The operation to perform on the label value.
					  * Constraints: Allowable values are: `is_or_unspecified`, `includes`, `ends_with`, `starts_with`.
					* `value` - (Optional, String) / The value of the label to filter by.
					  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^.*$/`.
			* `lucene_query` - (Optional, String) The Lucene query to filter logs.
			  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$/`.
	* `max_unique_count_per_group_by_key` - (Optional, String) The maximum unique count per group by key.
	  * Constraints: The maximum length is `19` characters. The minimum length is `1` character. The value must match regular expression `/^-?\\d{1,19}$/`.
	* `notification_payload_filter` - (Optional, List) The filter to specify which fields to include in the notification payload.
	  * Constraints: The list items must match regular expression `/^[a-zA-Z0-9_.]+$/`. The maximum length is `4096` items. The minimum length is `0` items.
	* `rules` - (Required, List) The rules for the log unique count alert.
	  * Constraints: The maximum length is `1` item. The minimum length is `1` item.
	Nested schema for **rules**:
		* `condition` - (Required, List) The condition for detecting unique counts in logs.
		Nested schema for **condition**:
			* `max_unique_count` - (Required, String) The maximum unique count for the alert condition.
			  * Constraints: The maximum length is `19` characters. The minimum length is `1` character. The value must match regular expression `/^-?\\d{1,19}$/`.
			* `time_window` - (Required, List) The time window for the unique count alert.
			Nested schema for **time_window**:
				* `logs_unique_value_time_window_specific_value` - (Required, String) A time window defined by a specific value.
				  * Constraints: Allowable values are: `minute_1_or_unspecified`, `minutes_15`, `minutes_20`, `minutes_30`, `hours_1`, `hours_2`, `hours_4`, `hours_6`, `hours_12`, `hours_24`, `minutes_5`, `minutes_10`, `hours_36`.
	* `unique_count_keypath` - (Required, String) The keypath in the logs to be used for unique count.
	  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$/`.
* `metric_anomaly` - (Optional, List) Configuration for metric-based anomaly detection alerts.
Nested schema for **metric_anomaly**:
	* `anomaly_alert_settings` - (Optional, List) The anomaly alert settings configuration.
	Nested schema for **anomaly_alert_settings**:
		* `percentage_of_deviation` - (Optional, Float) The percentage of deviation from the baseline for triggering the alert.
	* `condition_type` - (Required, String) The type of condition for the alert.
	  * Constraints: Allowable values are: `more_than_usual_or_unspecified`, `less_than_usual`.
	* `evaluation_delay_ms` - (Optional, Integer) The delay in milliseconds before evaluating the alert condition.
	  * Constraints: The maximum value is `2147483647`. The minimum value is `0`.
	* `metric_filter` - (Required, List) The filter to match metric entries for the alert.
	Nested schema for **metric_filter**:
		* `promql` - (Required, String) The filter is a PromQL expression.
		  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$/`.
	* `rules` - (Required, List) The rules for the metric anomaly alert.
	  * Constraints: The maximum length is `1` item. The minimum length is `1` item.
	Nested schema for **rules**:
		* `condition` - (Required, List) The condition for the metric anomaly alert.
		Nested schema for **condition**:
			* `for_over_pct` - (Optional, Integer) The percentage of the metric values that must exceed the threshold to trigger the alert.
			  * Constraints: The maximum value is `4294967295`. The minimum value is `0`.
			* `min_non_null_values_pct` - (Required, Integer) The percentage of non-null values required to trigger the alert.
			  * Constraints: The maximum value is `4294967295`. The minimum value is `0`.
			* `of_the_last` - (Required, List) The time window for the alert condition.
			Nested schema for **of_the_last**:
				* `metric_time_window_dynamic_duration` - (Optional, String) The time window as a dynamic value.
				  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$/`.
				* `metric_time_window_specific_value` - (Optional, String) The time window as a specific value.
				  * Constraints: Allowable values are: `minutes_1_or_unspecified`, `minutes_5`, `minutes_10`, `minutes_15`, `minutes_30`, `hour_1`, `hours_2`, `hours_4`, `hours_6`, `hours_12`, `hours_24`, `minutes_20`, `hours_36`.
			* `threshold` - (Required, Float) The threshold value for the alert condition.
* `metric_threshold` - (Optional, List) Configuration for metric-based threshold alerts.
Nested schema for **metric_threshold**:
	* `condition_type` - (Required, String) The type of the alert condition.
	  * Constraints: Allowable values are: `more_than_or_unspecified`, `less_than`, `more_than_or_equals`, `less_than_or_equals`.
	* `evaluation_delay_ms` - (Optional, Integer) The delay in milliseconds before evaluating the alert condition.
	  * Constraints: The maximum value is `2147483647`. The minimum value is `0`.
	* `metric_filter` - (Required, List) The filter to match metric entries for the alert.
	Nested schema for **metric_filter**:
		* `promql` - (Required, String) The filter is a PromQL expression.
		  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$/`.
	* `missing_values` - (Required, List) Configuration for handling missing values in the alert.
	Nested schema for **missing_values**:
		* `min_non_null_values_pct` - (Optional, Integer) If set, specifies the minimum percentage of non-null values required for the alert to be triggered.
		  * Constraints: The maximum value is `4294967295`. The minimum value is `0`.
		* `replace_with_zero` - (Optional, Boolean) If set to true, missing values will be replaced with zero.
	* `rules` - (Required, List) The rules for the metric threshold alert.
	  * Constraints: The maximum length is `5` items. The minimum length is `1` item.
	Nested schema for **rules**:
		* `condition` - (Required, List) The condition for the metric threshold alert.
		Nested schema for **condition**:
			* `for_over_pct` - (Required, Integer) The percentage of values that must exceed the threshold to trigger the alert.
			  * Constraints: The maximum value is `4294967295`. The minimum value is `0`.
			* `of_the_last` - (Required, List) The time window for the alert condition.
			Nested schema for **of_the_last**:
				* `metric_time_window_dynamic_duration` - (Optional, String) The time window as a dynamic value.
				  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$/`.
				* `metric_time_window_specific_value` - (Optional, String) The time window as a specific value.
				  * Constraints: Allowable values are: `minutes_1_or_unspecified`, `minutes_5`, `minutes_10`, `minutes_15`, `minutes_30`, `hour_1`, `hours_2`, `hours_4`, `hours_6`, `hours_12`, `hours_24`, `minutes_20`, `hours_36`.
			* `threshold` - (Required, Float) The threshold value for the alert condition.
		* `override` - (Required, List) The override settings for the alert.
		Nested schema for **override**:
			* `priority` - (Required, String) The priority of the alert definition.
			  * Constraints: Allowable values are: `p5_or_unspecified`, `p4`, `p3`, `p2`, `p1`.
	* `undetected_values_management` - (Optional, List) Configuration for handling the undetected values in the alert.
	Nested schema for **undetected_values_management**:
		* `auto_retire_timeframe` - (Required, String) The timeframe for auto-retiring the alert when undetected values are detected.
		  * Constraints: Allowable values are: `never_or_unspecified`, `minutes_5`, `minutes_10`, `hour_1`, `hours_2`, `hours_6`, `hours_12`, `hours_24`.
		* `trigger_undetected_values` - (Required, Boolean) Should trigger the alert when undetected values are detected.
* `name` - (Required, String) The name of the alert definition.
  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$/`.
* `notification_group` - (Optional, List) Primary notification group for alert events.
Nested schema for **notification_group**:
	* `group_by_keys` - (Required, List) The keys to group the alerts by.
	  * Constraints: The list items must match regular expression `/^[a-zA-Z0-9_.]+$/`. The maximum length is `4096` items. The minimum length is `0` items.
	* `webhooks` - (Required, List) The settings for webhooks associated with the alert definition.
	  * Constraints: The maximum length is `4096` items. The minimum length is `0` items.
	Nested schema for **webhooks**:
		* `integration` - (Required, List) The integration type for webhook notifications.
		Nested schema for **integration**:
			* `integration_id` - (Optional, Integer) The integration ID for the notification.
			  * Constraints: The maximum value is `4294967295`. The minimum value is `0`.
		* `minutes` - (Optional, Integer) The time in minutes before the notification is sent.
		  * Constraints: The maximum value is `4294967295`. The minimum value is `0`.
		* `notify_on` - (Optional, String) The condition to notify about the alert.
		  * Constraints: Allowable values are: `triggered_only_unspecified`, `triggered_and_resolved`.
* `phantom_mode` - (Optional, Boolean) Whether the alert is in phantom mode (creating incidents or not).
* `priority` - (Optional, String) The priority of the alert definition.
  * Constraints: Allowable values are: `p5_or_unspecified`, `p4`, `p3`, `p2`, `p1`.
* `type` - (Required, String) Alert type.
  * Constraints: Allowable values are: `logs_immediate_or_unspecified`, `logs_threshold`, `logs_anomaly`, `logs_ratio_threshold`, `logs_new_value`, `logs_unique_count`, `logs_time_relative_threshold`, `metric_threshold`, `metric_anomaly`, `flow`.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the logs_alert_definition.
* `alert_version_id` - (String) The old alert ID.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$/`.
* `created_time` - (String) The time when the alert definition was created.
* `updated_time` - (String) The time when the alert definition was last updated.


## Import

You can import the `ibm_logs_alert_definition` resource by using `id`. This is the alert definition's persistent ID (does not change on replace), AKA UniqueIdentifier.

# Syntax
<pre>
$ terraform import ibm_logs_alert_definition.logs_alert_definition &lt;id&gt;
</pre>

# Example
```
$ terraform import ibm_logs_alert_definition.logs_alert_definition 3dc02998-0b50-4ea8-b68a-4779d716fa1f
```
