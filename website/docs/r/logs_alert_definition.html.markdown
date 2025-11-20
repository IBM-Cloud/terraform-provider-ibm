---
layout: "ibm"
page_title: "IBM : ibm_logs_alert_definition"
description: |-
  Manages logs_alert_definition.
subcategory: "Cloud Logs"
---

# ibm_logs_alert_definition

Manage ICL Alerts using ibm_logs_alert_definition resource

## Example Usage

### standard immediate alert
```hcl
resource "ibm_logs_alert_definition" "standard_immediate" {

  instance_id = "470e285d-3354-44f8-8119-c91902d23"
  region      = "eu-gb"
  enabled     = true

  name         = "standard-immediate"
  phantom_mode = false
  priority     = "p1"

  type = "logs_immediate_or_unspecified"
  incidents_settings {
    minutes   = 10
    notify_on = "triggered_only_unspecified"
  }
  logs_immediate {
    notification_payload_filter = []

    logs_filter {
      simple_filter {
        lucene_query = "push"
        label_filters {
          application_name {
            operation = "is_or_unspecified"
            value     = "sev1"
          }

          subsystem_name {
            operation = "is_or_unspecified"
            value     = "sev1-logs"
          }
        }
      }
    }
  }
}
```
### standard less than threshold alert
```hcl
resource "ibm_logs_alert_definition" "standard_less_than_threshold" {

  description = "standard-less-than"
  instance_id = "470e285d-3354-44f8-8119-c91902d23"
  region      = "eu-gb"
  enabled     = true

  name         = "standard-less-than-threshold"
  phantom_mode = false

  type = "logs_threshold"
  incidents_settings {
    minutes   = 1
    notify_on = "triggered_only_unspecified"
  }
  logs_threshold {
    condition_type              = "less_than"
    evaluation_delay_ms         = 0
    notification_payload_filter = []
    logs_filter {
      simple_filter {
        lucene_query = "\"push\""
        label_filters {
          application_name {
            operation = "is_or_unspecified"
            value     = "sev1"
          }

          subsystem_name {
            operation = "is_or_unspecified"
            value     = "sev1-logs"
          }
        }
      }
    }
    rules {
      condition {
        threshold = 1

        time_window {
          logs_time_window_specific_value = "minutes_5_or_unspecified"
        }
      }
      override {
        priority = "p2"
      }
    }
    rules {
      condition {
        threshold = 2

        time_window {
          logs_time_window_specific_value = "minutes_10"
        }
      }
      override {
        priority = "p3"
      }
    }
    rules {
      condition {
        threshold = 1

        time_window {
          logs_time_window_specific_value = "minutes_10"
        }
      }
      override {
        priority = "p1"
      }
    }
    undetected_values_management {
      auto_retire_timeframe     = "never_or_unspecified"
      trigger_undetected_values = false
    }
  }
  notification_group {
    webhooks {
      integration {
        integration_id = data.ibm_logs_outgoing_webhook.logs_outgoing_webhook_instance.external_id
      }
    }
  }
}
```
### standard more than alert
```hcl
resource "ibm_logs_alert_definition" "standard_more_than" {

  instance_id = "470e285d-3354-44f8-8119-c91902d23"
  region      = "eu-gb"
  enabled     = true

  name         = "standard-more-than"
  phantom_mode = false
  priority     = "p3"

  type = "logs_threshold"
  incidents_settings {
    minutes   = 1
    notify_on = "triggered_only_unspecified"
  }
  logs_threshold {
    condition_type              = "more_than_or_unspecified"
    evaluation_delay_ms         = 0
    notification_payload_filter = []
    logs_filter {
      simple_filter {
        lucene_query = "\"push\""
        label_filters {
          application_name {
            operation = "is_or_unspecified"
            value     = "sev4"
          }
          subsystem_name {
            operation = "is_or_unspecified"
            value     = "sev4-logs"
          }
        }
      }
    }
    rules {
      condition {
        threshold = 1
        time_window {
          logs_time_window_specific_value = "minutes_10"
        }
      }
      override {
        priority = "p3"
      }
    }
    rules {
      condition {
        threshold = 1
        time_window {
          logs_time_window_specific_value = "minutes_5_or_unspecified"
        }
      }
      override {
        priority = "p2"
      }
    }
  }
  notification_group {
    webhooks {
      integration {
        integration_id = data.ibm_logs_outgoing_webhook.logs_outgoing_webhook_instance.external_id
      }
    }
  }
}
```
### standard more than usual alert
```hcl
resource "ibm_logs_alert_definition" "standard_more_than_usual" {

  instance_id = "470e285d-3354-44f8-8119-c91902d23"
  region      = "eu-gb"
  enabled     = true

  name         = "standard-more-than-usual"
  phantom_mode = false
  priority     = "p5_or_unspecified"

  type = "logs_anomaly"
  incidents_settings {
    minutes   = 1
    notify_on = "triggered_only_unspecified"
  }
  logs_anomaly {
    condition_type              = "more_than_usual_or_unspecified"
    evaluation_delay_ms         = 0
    notification_payload_filter = []
    logs_filter {
      simple_filter {
        lucene_query = "\"push\""
        label_filters {
          severities = []

          application_name {
            operation = "is_or_unspecified"
            value     = "sev5"
          }
          application_name {
            operation = "is_or_unspecified"
            value     = "sev4"
          }
          subsystem_name {
            operation = "is_or_unspecified"
            value     = "sev4-logs"
          }
        }
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
    webhooks {
      integration {
        integration_id = data.ibm_logs_outgoing_webhook.logs_outgoing_webhook_instance.external_id
      }
    }
  }
}
```
### Time relative less than alert
```hcl
resource "ibm_logs_alert_definition" "test_time_relative_less_than" {

  instance_id = "470e285d-3354-44f8-8119-c91902d23"
  region      = "eu-gb"
  enabled     = true

  name         = "test-time-relative-less-than"
  phantom_mode = false
  priority     = "p2"

  type = "logs_time_relative_threshold"
  incidents_settings {
    minutes   = 70
    notify_on = "triggered_only_unspecified"
  }
  logs_time_relative_threshold {
    condition_type              = "less_than"
    evaluation_delay_ms         = 0
    ignore_infinity             = true
    notification_payload_filter = []
    logs_filter {
      simple_filter {
        lucene_query = "\"This is my second log\""
      }
    }
    rules {
      condition {
        compared_to = "previous_hour_or_unspecified"
        threshold   = 4
      }
      override {
        priority = "p2"
      }
    }
    undetected_values_management {
      auto_retire_timeframe     = "never_or_unspecified"
      trigger_undetected_values = false
    }
  }
  notification_group {
    webhooks {
      integration {
        integration_id = data.ibm_logs_outgoing_webhook.logs_outgoing_webhook_instance.external_id
      }
    }
  }
}
```
### Time relative more than alert
```hcl
resource "ibm_logs_alert_definition" "test_time_relative_more_than" {

  instance_id = "470e285d-3354-44f8-8119-c91902d23"
  region      = "eu-gb"
  enabled     = true

  name         = "test-time-relative-more-than"
  phantom_mode = false
  priority     = "p1"

  type = "logs_time_relative_threshold"
  incidents_settings {
    minutes   = 60
    notify_on = "triggered_only_unspecified"
  }
  logs_time_relative_threshold {
    condition_type              = "more_than_or_unspecified"
    evaluation_delay_ms         = 0
    ignore_infinity             = true
    notification_payload_filter = []
    logs_filter {
      simple_filter {
        lucene_query = "\"Push and Query integration test\""
      }
    }
    rules {
      condition {
        compared_to = "previous_hour_or_unspecified"
        threshold   = 1
      }
      override {
        priority = "p1"
      }
    }
  }
  notification_group {
    webhooks {
      integration {
        integration_id = data.ibm_logs_outgoing_webhook.logs_outgoing_webhook_instance.external_id
      }
    }
  }
}
```
### ration less than alert
```hcl
# ibm_logs_alert_definition.ratio_less_than:
resource "ibm_logs_alert_definition" "ratio_less_than" {

  instance_id = "470e285d-3354-44f8-8119-c91902d23"
  region      = "eu-gb"
  enabled     = true

  name         = "ratio less than"
  phantom_mode = false
  priority     = "p2"

  type = "logs_ratio_threshold"
  incidents_settings {
    minutes   = 10
    notify_on = "triggered_only_unspecified"
  }
  logs_ratio_threshold {
    condition_type              = "less_than"
    denominator_alias           = "Query 2"
    evaluation_delay_ms         = 0
    group_by_for                = "both_or_unspecified"
    ignore_infinity             = true
    notification_payload_filter = []
    numerator_alias             = "Query 1"
    denominator {
      simple_filter {
        lucene_query = "\"This is my second log\""
      }
    }
    numerator {
      simple_filter {
        lucene_query = "\"Push and Query integration test\""
      }
    }
    rules {
      condition {
        threshold = 3

        time_window {
          logs_ratio_time_window_specific_value = "minutes_5_or_unspecified"
        }
      }
      override {
        priority = "p2"
      }
    }
    undetected_values_management {
      auto_retire_timeframe     = "never_or_unspecified"
      trigger_undetected_values = false
    }
  }
  notification_group {
    webhooks {
      integration {
        integration_id = data.ibm_logs_outgoing_webhook.logs_outgoing_webhook_instance.external_id
      }
    }
  }
}
```
### ration more than alert with multi conditions
```hcl
resource "ibm_logs_alert_definition" "ratio_more_than_multiple" {

  instance_id = "470e285d-3354-44f8-8119-c91902d23"
  region      = "eu-gb"
  enabled     = true

  name         = "ratio-more-than-multiple"
  phantom_mode = false
  priority     = "p4"

  type = "logs_ratio_threshold"
  incidents_settings {
    minutes   = 10
    notify_on = "triggered_and_resolved"
  }
  logs_ratio_threshold {
    condition_type              = "more_than_or_unspecified"
    denominator_alias           = "Query 2"
    evaluation_delay_ms         = 0
    group_by_for                = "both_or_unspecified"
    ignore_infinity             = true
    notification_payload_filter = []
    numerator_alias             = "Query 1"
    denominator {
      simple_filter {
        lucene_query = "\"This is my second log\""
      }
    }
    numerator {
      simple_filter {
        lucene_query = "\"Push and Query integration test\""
      }
    }
    rules {
      condition {
        threshold = 2

        time_window {
          logs_ratio_time_window_specific_value = "minutes_5_or_unspecified"
        }
      }
      override {
        priority = "p4"
      }
    }
    rules {
      condition {
        threshold = 4

        time_window {
          logs_ratio_time_window_specific_value = "minutes_10"
        }
      }
      override {
        priority = "p1"
      }
    }
  }
  notification_group {
    webhooks {
      integration {
        integration_id = data.ibm_logs_outgoing_webhook.logs_outgoing_webhook_instance.external_id
      }
    }
  }
}

```
### New value alert
```hcl
resource "ibm_logs_alert_definition" "new_value" {

  instance_id = "470e285d-3354-44f8-8119-c91902d23"
  region      = "eu-gb"
  enabled     = true

  name         = "new-value"
  phantom_mode = false
  priority     = "p5_or_unspecified"

  type = "logs_new_value"
  incidents_settings {
    minutes   = 720
    notify_on = "triggered_only_unspecified"
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
  notification_group {
    group_by_keys = [
      "ibm.logId",
    ]
  }
}
```
### Metric less than or equals alert
```hcl
resource "ibm_logs_alert_definition" "metric_alert_less_than_or_equals" {

  instance_id = "470e285d-3354-44f8-8119-c91902d23"
  region      = "eu-gb"
  enabled     = true

  name         = "metric-less-than-or-equals"
  phantom_mode = false
  priority     = "p2"

  type = "metric_threshold"
  incidents_settings {
    minutes   = 10
    notify_on = "triggered_and_resolved"
  }
  metric_threshold {
    condition_type      = "less_than_or_equals"
    evaluation_delay_ms = 0

    metric_filter {
      promql = "duration_cx_sum"
    }

    missing_values {
      replace_with_zero = true
    }

    rules {
      condition {
        for_over_pct = 0
        threshold    = 1

        of_the_last {
          metric_time_window_specific_value = "minutes_10"
        }
      }
      override {
        priority = "p2"
      }
    }

    undetected_values_management {
      auto_retire_timeframe     = "never_or_unspecified"
      trigger_undetected_values = false
    }
  }

  notification_group {
    group_by_keys = []

    webhooks {
      minutes = 0

      integration {
        integration_id = data.ibm_logs_outgoing_webhook.logs_outgoing_webhook_instance.external_id
      }
    }
  }
}
```
### Metric more than or equals alert
```hcl
resource "ibm_logs_alert_definition" "metric_more_than_or_equals_with_usage" {

  instance_id = "470e285d-3354-44f8-8119-c91902d23"
  region      = "eu-gb"
  enabled     = true

  name         = "metric-more-than-or-equals-with-usage"
  phantom_mode = false
  priority     = "p1"

  type = "metric_threshold"

  incidents_settings {
    minutes   = 10
    notify_on = "triggered_only_unspecified"
  }

  metric_threshold {
    condition_type      = "more_than_or_equals"
    evaluation_delay_ms = 0

    metric_filter {
      promql = "cx_data_usage_bytes_total"
    }

    missing_values {
      min_non_null_values_pct = 100
    }

    rules {
      condition {
        for_over_pct = 0
        threshold    = 1

        of_the_last {
          metric_time_window_specific_value = "minutes_10"
        }
      }
      override {
        priority = "p1"
      }
    }
  }

  notification_group {
    group_by_keys = []

    webhooks {
      minutes = 0

      integration {
        integration_id = data.ibm_logs_outgoing_webhook.logs_outgoing_webhook_instance.external_id
      }
    }
  }
}
```
### Flow alert
```hcl
# ibm_logs_alert_definition.flow_alert:
resource "ibm_logs_alert_definition" "flow_alert" {
  instance_id  = "470e285d-3354-44f8-8119-c91902d23"
  region       = "eu-gb"
  enabled      = true
  name         = "flow-alert"
  phantom_mode = false
  priority     = "p1"
  type         = "flow"
  flow {
    enforce_suppression = false
    stages {
      timeframe_ms   = "0"
      timeframe_type = "up_to"
      flow_stages_groups {
        groups {
          alerts_op = "or"
          next_op   = "and_or_unspecified"

          alert_defs {
            id  = ibm_logs_alert_definition.standard_less_than_threshold.alert_def_id
            not = false
          }
          alert_defs {
            id  = ibm_logs_alert_definition.standard_immediate.alert_def_id
            not = false
          }
        }
      }
    }
    stages {
      timeframe_ms   = "3600000"
      timeframe_type = "up_to"
      flow_stages_groups {
        groups {
          alerts_op = "and_or_unspecified"
          next_op   = "and_or_unspecified"

          alert_defs {
            id  = ibm_logs_alert_definition.new_value.alert_def_id
            not = false
          }
        }
      }
    }
  }
  incidents_settings {
    minutes   = 10
    notify_on = "triggered_only_unspecified"
  }
  notification_group {
    webhooks {
      integration {
        integration_id = data.ibm_logs_outgoing_webhook.logs_outgoing_webhook_instance.external_id
      }
    }
  }
}
```
### Unique count alert
```hcl
# ibm_logs_alert_definition.unique_count:
resource "ibm_logs_alert_definition" "unique_count" {

  instance_id = "470e285d-3354-44f8-8119-c91902d23"
  region      = "eu-gb"
  enabled     = true
  group_by_keys = [
    "coralogix.logId",
  ]

  name         = "unique-count"
  phantom_mode = false
  priority     = "p1"

  type = "logs_unique_count"

  incidents_settings {
    minutes   = 5
    notify_on = "triggered_only_unspecified"
  }
  logs_unique_count {
    max_unique_count_per_group_by_key = "10"
    notification_payload_filter       = []
    unique_count_keypath              = "text"
    logs_filter {
      simple_filter {
        lucene_query = "\"push\""
        label_filters {
          severities = []
          application_name {
            operation = "is_or_unspecified"
            value     = "sev1"
          }
          subsystem_name {
            operation = "is_or_unspecified"
            value     = "sev1-logs"
          }
        }
      }
    }
    rules {
      condition {
        max_unique_count = "0"
        time_window {
          logs_unique_value_time_window_specific_value = "minute_1_or_unspecified"
        }
      }
    }
  }
  notification_group {
    group_by_keys = [
      "coralogix.logId",
    ]
  }
}
```

## Argument Reference

You can specify the following arguments for this resource.

* `instance_id` - (Required, Forces new resource, String)  Cloud Logs Instance GUID.
* `region` - (Optional, Forces new resource, String) Cloud Logs Instance Region.
* `active_on` - (Optional, List) Defining when the alert is active.
Nested schema for **active_on**:
	* `day_of_week` - (Required, List) Days of the week when the alert is active.
	  * Constraints: Allowable list items are: `monday_or_unspecified`, `tuesday`, `wednesday`, `thursday`, `friday`, `saturday`, `sunday`. The maximum length is `7` items. The minimum length is `1` item.
	* `end_time` - (Required, List) Start time of the alert activity.
	Nested schema for **end_time**:
		* `hours` - (Optional, Integer) The hour of the day in 24-hour format. Must be an integer between 0 and 23.
		  * Constraints: The maximum value is `23`. The minimum value is `0`.
		* `minutes` - (Optional, Integer) Minute of the hour of the day. Must be an integer between 0 and 59.
		  * Constraints: The maximum value is `59`. The minimum value is `0`.
	* `start_time` - (Required, List) Start time of the alert activity.
	Nested schema for **start_time**:
		* `hours` - (Optional, Integer) The hour of the day in 24-hour format. Must be an integer between 0 and 23.
		  * Constraints: The maximum value is `23`. The minimum value is `0`.
		* `minutes` - (Optional, Integer) Minute of the hour of the day. Must be an integer between 0 and 59.
		  * Constraints: The maximum value is `59`. The minimum value is `0`.
* `deleted` - (Optional, Boolean) Whether the alert has been marked as deleted.
* `description` - (Optional, String) A detailed description of what the alert monitors and when it triggers.
  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$/`.
* `enabled` - (Optional, Boolean) Whether the alert is currently active and monitoring. If true, alert is active.
* `entity_labels` - (Optional, Map) Labels used to identify and categorize the alert entity.
* `flow` - (Optional, List) Configuration for flow alerts.
Nested schema for **flow**:
	* `enforce_suppression` - (Optional, Boolean) Whether to enforce suppression for the flow alert.
	* `stages` - (Required, List) The definition of stages of the flow alert.
	  * Constraints: The maximum length is `4096` items. The minimum length is `1` item.
	Nested schema for **stages**:
		* `flow_stages_groups` - (Required, List) The definition of groups in the flow alert.
		Nested schema for **flow_stages_groups**:
			* `groups` - (Required, List) The definition of an array of groups with alerts and logical operation among those alerts in the flow alert.
			  * Constraints: The maximum length is `4096` items. The minimum length is `1` item.
			Nested schema for **groups**:
				* `alert_defs` - (Required, List) The alert definitions for the flow stage group.
				  * Constraints: The maximum length is `4096` items. The minimum length is `1` item.
				Nested schema for **alert_defs**:
					* `id` - (Required, String) The alert definition ID.
					  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$/`.
					* `not` - (Optional, Boolean) Whether or not to negate the alert definition. If true, flow checks for the negate condition of the respective alert.
				* `alerts_op` - (Required, String) The logical operation to apply to the alerts in the group.
				  * Constraints: Allowable values are: `and_or_unspecified`, `or`.
				* `next_op` - (Required, String) The logical operation to apply to the next stage.
				  * Constraints: Allowable values are: `and_or_unspecified`, `or`.
		* `timeframe_ms` - (Required, String) The timeframe for the flow alert in milliseconds.
		  * Constraints: The maximum length is `19` characters. The minimum length is `1` character. The value must match regular expression `/^-?\\d{1,19}$/`.
		* `timeframe_type` - (Required, String) The type of timeframe for the flow alert.
		  * Constraints: Allowable values are: `unspecified`, `up_to`.
* `group_by_keys` - (Optional, List) Keys used to group and aggregate alert data.
  * Constraints: The list items must match regular expression `/^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$/`. The maximum length is `2` items. The minimum length is `0` items.
* `incidents_settings` - (Optional, List) Incident creation and management settings.
Nested schema for **incidents_settings**:
	* `minutes` - (Optional, Integer) The time in minutes before the alert can be triggered again.
	  * Constraints: The maximum value is `4294967295`. The minimum value is `0`.
	* `notify_on` - (Optional, String) Indicate if the alert should be triggered or triggered and resolved.
	  * Constraints: Allowable values are: `triggered_only_unspecified`, `triggered_and_resolved`.
* `logs_anomaly` - (Optional, List) Configuration for the log-based anomaly detection alerts.
Nested schema for **logs_anomaly**:
	* `anomaly_alert_settings` - (Optional, List) The anomaly alert settings configuration.
	Nested schema for **anomaly_alert_settings**:
		* `percentage_of_deviation` - (Optional, Float) The percentage of deviation from the baseline when the alert is triggered.
	* `condition_type` - (Required, String) The condition type for the alert.
	  * Constraints: Allowable values are: `more_than_usual_or_unspecified`.
	* `evaluation_delay_ms` - (Optional, Integer) The delay in milliseconds before evaluating the alert condition.
	  * Constraints: The maximum value is `2147483647`. The minimum value is `0`.
	* `logs_filter` - (Optional, List) The filter to match log entries for immediate alerts.
	Nested schema for **logs_filter**:
		* `simple_filter` - (Optional, List) A simple filter that uses a Lucene query and label filters.
		Nested schema for **simple_filter**:
			* `label_filters` - (Optional, List) The label filters to filter logs.
			Nested schema for **label_filters**:
				* `application_name` - (Optional, List) Filter by application names.
				  * Constraints: The maximum length is `4096` items. The minimum length is `0` items.
				Nested schema for **application_name**:
					* `operation` - (Required, String) The operation to perform on the label value.
					  * Constraints: Allowable values are: `is_or_unspecified`, `includes`, `ends_with`, `starts_with`.
					* `value` - (Optional, String) The value used to filter the label.
					  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^.*$/`.
				* `severities` - (Optional, List) Filter by log severities.
				  * Constraints: Allowable list items are: `verbose_unspecified`, `debug`, `info`, `warning`, `error`, `critical`. The maximum length is `4096` items. The minimum length is `0` items.
				* `subsystem_name` - (Optional, List) Filter by subsystem names.
				  * Constraints: The maximum length is `4096` items. The minimum length is `0` items.
				Nested schema for **subsystem_name**:
					* `operation` - (Required, String) The operation to perform on the label value.
					  * Constraints: Allowable values are: `is_or_unspecified`, `includes`, `ends_with`, `starts_with`.
					* `value` - (Optional, String) The value used to filter the label.
					  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^.*$/`.
			* `lucene_query` - (Optional, String) The Lucene query to filter logs.
			  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$/`.
	* `notification_payload_filter` - (Optional, List) The notification payload filter to specify which fields are included in the notification.
	  * Constraints: The list items must match regular expression `/^[a-zA-Z0-9_.]+$/`. The maximum length is `4096` items. The minimum length is `0` items.
	* `rules` - (Required, List) The condition rules for the log anomaly alert.
	  * Constraints: The maximum length is `1` item. The minimum length is `1` item.
	Nested schema for **rules**:
		* `condition` - (Required, List) The condition for the anomaly alert.
		Nested schema for **condition**:
			* `minimum_threshold` - (Required, Float) The threshold value for the alert condition.
			* `time_window` - (Required, List) The time window for the alert condition.
			Nested schema for **time_window**:
				* `logs_time_window_specific_value` - (Required, String) The time window defined for an alert to be triggered.
				  * Constraints: Allowable values are: `minutes_5_or_unspecified`, `minutes_10`, `minutes_20`, `minutes_15`, `minutes_30`, `hour_1`, `hours_2`, `hours_4`, `hours_6`, `hours_12`, `hours_24`, `hours_36`.
* `logs_immediate` - (Optional, List) Configuration for immediate log-based alerts.
Nested schema for **logs_immediate**:
	* `logs_filter` - (Optional, List) The filter to match log entries for immediate alerts.
	Nested schema for **logs_filter**:
		* `simple_filter` - (Optional, List) A simple filter that uses a Lucene query and label filters.
		Nested schema for **simple_filter**:
			* `label_filters` - (Optional, List) The label filters to filter logs.
			Nested schema for **label_filters**:
				* `application_name` - (Optional, List) Filter by application names.
				  * Constraints: The maximum length is `4096` items. The minimum length is `0` items.
				Nested schema for **application_name**:
					* `operation` - (Required, String) The operation to perform on the label value.
					  * Constraints: Allowable values are: `is_or_unspecified`, `includes`, `ends_with`, `starts_with`.
					* `value` - (Optional, String) The value used to filter the label.
					  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^.*$/`.
				* `severities` - (Optional, List) Filter by log severities.
				  * Constraints: Allowable list items are: `verbose_unspecified`, `debug`, `info`, `warning`, `error`, `critical`. The maximum length is `4096` items. The minimum length is `0` items.
				* `subsystem_name` - (Optional, List) Filter by subsystem names.
				  * Constraints: The maximum length is `4096` items. The minimum length is `0` items.
				Nested schema for **subsystem_name**:
					* `operation` - (Required, String) The operation to perform on the label value.
					  * Constraints: Allowable values are: `is_or_unspecified`, `includes`, `ends_with`, `starts_with`.
					* `value` - (Optional, String) The value used to filter the label.
					  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^.*$/`.
			* `lucene_query` - (Optional, String) The Lucene query to filter logs.
			  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$/`.
	* `notification_payload_filter` - (Optional, List) The filter to specify which fields are included in the notification payload.
	  * Constraints: The list items must match regular expression `/^[a-zA-Z0-9_.]+$/`. The maximum length is `4096` items. The minimum length is `0` items.
* `logs_new_value` - (Optional, List) Configuration for alerts triggered by new log values.
Nested schema for **logs_new_value**:
	* `logs_filter` - (Optional, List) The filter to match log entries for immediate alerts.
	Nested schema for **logs_filter**:
		* `simple_filter` - (Optional, List) A simple filter that uses a Lucene query and label filters.
		Nested schema for **simple_filter**:
			* `label_filters` - (Optional, List) The label filters to filter logs.
			Nested schema for **label_filters**:
				* `application_name` - (Optional, List) Filter by application names.
				  * Constraints: The maximum length is `4096` items. The minimum length is `0` items.
				Nested schema for **application_name**:
					* `operation` - (Required, String) The operation to perform on the label value.
					  * Constraints: Allowable values are: `is_or_unspecified`, `includes`, `ends_with`, `starts_with`.
					* `value` - (Optional, String) The value used to filter the label.
					  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^.*$/`.
				* `severities` - (Optional, List) Filter by log severities.
				  * Constraints: Allowable list items are: `verbose_unspecified`, `debug`, `info`, `warning`, `error`, `critical`. The maximum length is `4096` items. The minimum length is `0` items.
				* `subsystem_name` - (Optional, List) Filter by subsystem names.
				  * Constraints: The maximum length is `4096` items. The minimum length is `0` items.
				Nested schema for **subsystem_name**:
					* `operation` - (Required, String) The operation to perform on the label value.
					  * Constraints: Allowable values are: `is_or_unspecified`, `includes`, `ends_with`, `starts_with`.
					* `value` - (Optional, String) The value used to filter the label.
					  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^.*$/`.
			* `lucene_query` - (Optional, String) The Lucene query to filter logs.
			  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$/`.
	* `notification_payload_filter` - (Optional, List) The filter to specify which fields are included in the notification payload.
	  * Constraints: The list items must match regular expression `/^[a-zA-Z0-9_.]+$/`. The maximum length is `4096` items. The minimum length is `0` items.
	* `rules` - (Required, List) The condition rules for the log new value alert.
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
* `logs_ratio_threshold` - (Optional, List) Configuration for the log-based ratio threshold alerts.
Nested schema for **logs_ratio_threshold**:
	* `condition_type` - (Required, String) The condition type for the alert.
	  * Constraints: Allowable values are: `more_than_or_unspecified`, `less_than`.
	* `denominator` - (Required, List) The filter to match log entries for immediate alerts.
	Nested schema for **denominator**:
		* `simple_filter` - (Optional, List) A simple filter that uses a Lucene query and label filters.
		Nested schema for **simple_filter**:
			* `label_filters` - (Optional, List) The label filters to filter logs.
			Nested schema for **label_filters**:
				* `application_name` - (Optional, List) Filter by application names.
				  * Constraints: The maximum length is `4096` items. The minimum length is `0` items.
				Nested schema for **application_name**:
					* `operation` - (Required, String) The operation to perform on the label value.
					  * Constraints: Allowable values are: `is_or_unspecified`, `includes`, `ends_with`, `starts_with`.
					* `value` - (Optional, String) The value used to filter the label.
					  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^.*$/`.
				* `severities` - (Optional, List) Filter by log severities.
				  * Constraints: Allowable list items are: `verbose_unspecified`, `debug`, `info`, `warning`, `error`, `critical`. The maximum length is `4096` items. The minimum length is `0` items.
				* `subsystem_name` - (Optional, List) Filter by subsystem names.
				  * Constraints: The maximum length is `4096` items. The minimum length is `0` items.
				Nested schema for **subsystem_name**:
					* `operation` - (Required, String) The operation to perform on the label value.
					  * Constraints: Allowable values are: `is_or_unspecified`, `includes`, `ends_with`, `starts_with`.
					* `value` - (Optional, String) The value used to filter the label.
					  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^.*$/`.
			* `lucene_query` - (Optional, String) The Lucene query to filter logs.
			  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$/`.
	* `denominator_alias` - (Optional, String) The alias for the denominator filter, used for display purposes.
	  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$/`.
	* `evaluation_delay_ms` - (Optional, Integer) The delay in milliseconds before evaluating the alert condition.
	  * Constraints: The maximum value is `2147483647`. The minimum value is `0`.
	* `group_by_for` - (Required, String) The group by settings for the numerator and denominator filters.
	  * Constraints: Allowable values are: `both_or_unspecified`, `numerator_only`, `denumerator_only`.
	* `ignore_infinity` - (Optional, Boolean) Determine whether to ignore an infinity result or not. If true, alert is not triggered. When the value of second query is 0, the result of the ratio will be infinity.
	* `notification_payload_filter` - (Optional, List) The filter to specify which fields are included in the notification payload.
	  * Constraints: The list items must match regular expression `/^[a-zA-Z0-9_.]+$/`. The maximum length is `4096` items. The minimum length is `0` items.
	* `numerator` - (Required, List) The filter to match log entries for immediate alerts.
	Nested schema for **numerator**:
		* `simple_filter` - (Optional, List) A simple filter that uses a Lucene query and label filters.
		Nested schema for **simple_filter**:
			* `label_filters` - (Optional, List) The label filters to filter logs.
			Nested schema for **label_filters**:
				* `application_name` - (Optional, List) Filter by application names.
				  * Constraints: The maximum length is `4096` items. The minimum length is `0` items.
				Nested schema for **application_name**:
					* `operation` - (Required, String) The operation to perform on the label value.
					  * Constraints: Allowable values are: `is_or_unspecified`, `includes`, `ends_with`, `starts_with`.
					* `value` - (Optional, String) The value used to filter the label.
					  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^.*$/`.
				* `severities` - (Optional, List) Filter by log severities.
				  * Constraints: Allowable list items are: `verbose_unspecified`, `debug`, `info`, `warning`, `error`, `critical`. The maximum length is `4096` items. The minimum length is `0` items.
				* `subsystem_name` - (Optional, List) Filter by subsystem names.
				  * Constraints: The maximum length is `4096` items. The minimum length is `0` items.
				Nested schema for **subsystem_name**:
					* `operation` - (Required, String) The operation to perform on the label value.
					  * Constraints: Allowable values are: `is_or_unspecified`, `includes`, `ends_with`, `starts_with`.
					* `value` - (Optional, String) The value used to filter the label.
					  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^.*$/`.
			* `lucene_query` - (Optional, String) The Lucene query to filter logs.
			  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$/`.
	* `numerator_alias` - (Optional, String) The alias for the numerator filter, used for display purposes.
	  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$/`.
	* `rules` - (Required, List) The condition rules for the ratio alert.
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
		* `trigger_undetected_values` - (Required, Boolean) Should trigger the alert when undetected values are detected. If true, alert is triggered.
* `logs_threshold` - (Optional, List) Configuration for the log-based threshold alerts.
Nested schema for **logs_threshold**:
	* `condition_type` - (Required, String) The condition type for the alert.
	  * Constraints: Allowable values are: `more_than_or_unspecified`, `less_than`.
	* `evaluation_delay_ms` - (Optional, Integer) The delay in milliseconds before evaluating the alert condition.
	  * Constraints: The maximum value is `2147483647`. The minimum value is `0`.
	* `logs_filter` - (Optional, List) The filter to match log entries for immediate alerts.
	Nested schema for **logs_filter**:
		* `simple_filter` - (Optional, List) A simple filter that uses a Lucene query and label filters.
		Nested schema for **simple_filter**:
			* `label_filters` - (Optional, List) The label filters to filter logs.
			Nested schema for **label_filters**:
				* `application_name` - (Optional, List) Filter by application names.
				  * Constraints: The maximum length is `4096` items. The minimum length is `0` items.
				Nested schema for **application_name**:
					* `operation` - (Required, String) The operation to perform on the label value.
					  * Constraints: Allowable values are: `is_or_unspecified`, `includes`, `ends_with`, `starts_with`.
					* `value` - (Optional, String) The value used to filter the label.
					  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^.*$/`.
				* `severities` - (Optional, List) Filter by log severities.
				  * Constraints: Allowable list items are: `verbose_unspecified`, `debug`, `info`, `warning`, `error`, `critical`. The maximum length is `4096` items. The minimum length is `0` items.
				* `subsystem_name` - (Optional, List) Filter by subsystem names.
				  * Constraints: The maximum length is `4096` items. The minimum length is `0` items.
				Nested schema for **subsystem_name**:
					* `operation` - (Required, String) The operation to perform on the label value.
					  * Constraints: Allowable values are: `is_or_unspecified`, `includes`, `ends_with`, `starts_with`.
					* `value` - (Optional, String) The value used to filter the label.
					  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^.*$/`.
			* `lucene_query` - (Optional, String) The Lucene query to filter logs.
			  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$/`.
	* `notification_payload_filter` - (Optional, List) The filter to specify which fields are included in the notification payload.
	  * Constraints: The list items must match regular expression `/^[a-zA-Z0-9_.]+$/`. The maximum length is `4096` items. The minimum length is `0` items.
	* `rules` - (Required, List) The condition rules for the threshold alert.
	  * Constraints: The maximum length is `5` items. The minimum length is `1` item.
	Nested schema for **rules**:
		* `condition` - (Required, List) The condition for the threshold alert.
		Nested schema for **condition**:
			* `threshold` - (Required, Float) The threshold value for the alert condition.
			* `time_window` - (Required, List) The time window for the alert condition.
			Nested schema for **time_window**:
				* `logs_time_window_specific_value` - (Required, String) The time window defined for an alert to be triggered.
				  * Constraints: Allowable values are: `minutes_5_or_unspecified`, `minutes_10`, `minutes_20`, `minutes_15`, `minutes_30`, `hour_1`, `hours_2`, `hours_4`, `hours_6`, `hours_12`, `hours_24`, `hours_36`.
		* `override` - (Required, List) The override settings for the alert.
		Nested schema for **override**:
			* `priority` - (Required, String) The priority of the alert definition.
			  * Constraints: Allowable values are: `p5_or_unspecified`, `p4`, `p3`, `p2`, `p1`.
	* `undetected_values_management` - (Optional, List) Configuration for handling the undetected values in the alert.
	Nested schema for **undetected_values_management**:
		* `auto_retire_timeframe` - (Required, String) The timeframe for auto-retiring the alert when undetected values are detected.
		  * Constraints: Allowable values are: `never_or_unspecified`, `minutes_5`, `minutes_10`, `hour_1`, `hours_2`, `hours_6`, `hours_12`, `hours_24`.
		* `trigger_undetected_values` - (Required, Boolean) Should trigger the alert when undetected values are detected. If true, alert is triggered.
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
				* `application_name` - (Optional, List) Filter by application names.
				  * Constraints: The maximum length is `4096` items. The minimum length is `0` items.
				Nested schema for **application_name**:
					* `operation` - (Required, String) The operation to perform on the label value.
					  * Constraints: Allowable values are: `is_or_unspecified`, `includes`, `ends_with`, `starts_with`.
					* `value` - (Optional, String) The value used to filter the label.
					  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^.*$/`.
				* `severities` - (Optional, List) Filter by log severities.
				  * Constraints: Allowable list items are: `verbose_unspecified`, `debug`, `info`, `warning`, `error`, `critical`. The maximum length is `4096` items. The minimum length is `0` items.
				* `subsystem_name` - (Optional, List) Filter by subsystem names.
				  * Constraints: The maximum length is `4096` items. The minimum length is `0` items.
				Nested schema for **subsystem_name**:
					* `operation` - (Required, String) The operation to perform on the label value.
					  * Constraints: Allowable values are: `is_or_unspecified`, `includes`, `ends_with`, `starts_with`.
					* `value` - (Optional, String) The value used to filter the label.
					  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^.*$/`.
			* `lucene_query` - (Optional, String) The Lucene query to filter logs.
			  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$/`.
	* `notification_payload_filter` - (Optional, List) The filter to specify which fields are included in the notification payload.
	  * Constraints: The list items must match regular expression `/^[a-zA-Z0-9_.]+$/`. The maximum length is `4096` items. The minimum length is `0` items.
	* `rules` - (Required, List) The condition rules for the time-relative alert.
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
		* `trigger_undetected_values` - (Required, Boolean) Should trigger the alert when undetected values are detected. If true, alert is triggered.
* `logs_unique_count` - (Optional, List) Configuration for alerts based on unique log value counts.
Nested schema for **logs_unique_count**:
	* `logs_filter` - (Optional, List) The filter to match log entries for immediate alerts.
	Nested schema for **logs_filter**:
		* `simple_filter` - (Optional, List) A simple filter that uses a Lucene query and label filters.
		Nested schema for **simple_filter**:
			* `label_filters` - (Optional, List) The label filters to filter logs.
			Nested schema for **label_filters**:
				* `application_name` - (Optional, List) Filter by application names.
				  * Constraints: The maximum length is `4096` items. The minimum length is `0` items.
				Nested schema for **application_name**:
					* `operation` - (Required, String) The operation to perform on the label value.
					  * Constraints: Allowable values are: `is_or_unspecified`, `includes`, `ends_with`, `starts_with`.
					* `value` - (Optional, String) The value used to filter the label.
					  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^.*$/`.
				* `severities` - (Optional, List) Filter by log severities.
				  * Constraints: Allowable list items are: `verbose_unspecified`, `debug`, `info`, `warning`, `error`, `critical`. The maximum length is `4096` items. The minimum length is `0` items.
				* `subsystem_name` - (Optional, List) Filter by subsystem names.
				  * Constraints: The maximum length is `4096` items. The minimum length is `0` items.
				Nested schema for **subsystem_name**:
					* `operation` - (Required, String) The operation to perform on the label value.
					  * Constraints: Allowable values are: `is_or_unspecified`, `includes`, `ends_with`, `starts_with`.
					* `value` - (Optional, String) The value used to filter the label.
					  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^.*$/`.
			* `lucene_query` - (Optional, String) The Lucene query to filter logs.
			  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$/`.
	* `max_unique_count_per_group_by_key` - (Optional, String) The maximum unique count per group by key.
	  * Constraints: The maximum length is `19` characters. The minimum length is `1` character. The value must match regular expression `/^-?\\d{1,19}$/`.
	* `notification_payload_filter` - (Optional, List) The filter to specify which fields are included in the notification payload.
	  * Constraints: The list items must match regular expression `/^[a-zA-Z0-9_.]+$/`. The maximum length is `4096` items. The minimum length is `0` items.
	* `rules` - (Required, List) Rules defining the conditions for the unique count alert.
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
		* `percentage_of_deviation` - (Optional, Float) The percentage of deviation from the baseline when the alert is triggered.
	* `condition_type` - (Required, String) The condition type for the alert.
	  * Constraints: Allowable values are: `more_than_usual_or_unspecified`, `less_than_usual`.
	* `evaluation_delay_ms` - (Optional, Integer) The delay in milliseconds before evaluating the alert condition.
	  * Constraints: The maximum value is `2147483647`. The minimum value is `0`.
	* `metric_filter` - (Required, List) The filter to match metric entries for the alert.
	Nested schema for **metric_filter**:
		* `promql` - (Required, String) The filter is a PromQL expression.
		  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$/`.
	* `rules` - (Required, List) The condition rules for the metric anomaly alert.
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
	* `missing_values` - (Required, List) Configuration for handling missing values in the alert. Only one of `replace_with_zero` or `min_non_null_value_pct` is supported.
	Nested schema for **missing_values**:
		* `min_non_null_values_pct` - (Optional, Integer) If set, specifies the minimum percentage of non-null values required for the alert to be triggered.
		  * Constraints: The maximum value is `4294967295`. The minimum value is `0`.
		* `replace_with_zero` - (Optional, Boolean) If set to true, missing values will be replaced with zero.
	* `rules` - (Required, List) The condition rules for the metric threshold alert.
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
		* `trigger_undetected_values` - (Required, Boolean) Should trigger the alert when undetected values are detected. If true, alert is triggered.
* `name` - (Required, String) The name of the alert definition.
  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$/`.
* `notification_group` - (Optional, List) Primary notification group for alert events.
Nested schema for **notification_group**:
	* `group_by_keys` - (Optional, List) Group the alerts by these keys.
	  * Constraints: The list items must match regular expression `/^[a-zA-Z0-9_.]+$/`. The maximum length is `4096` items. The minimum length is `0` items.
	* `webhooks` - (Optional, List) The settings for webhooks associated with the alert definition.
	  * Constraints: The maximum length is `4096` items. The minimum length is `0` items.
	Nested schema for **webhooks**:
		* `integration` - (Required, List) The integration type for webhook notifications.
		Nested schema for **integration**:
			* `integration_id` - (Optional, Integer) The integration ID for the notification.
			  * Constraints: The maximum value is `4294967295`. The minimum value is `0`.
		* `minutes` - (Optional, Integer) The time in minutes before the notification is sent.
		  * Constraints: The maximum value is `4294967295`. The minimum value is `0`.
		* `notify_on` - (Optional, String) Indicate if the alert should be triggered or triggered and resolved.
		  * Constraints: Allowable values are: `triggered_only_unspecified`, `triggered_and_resolved`.
* `phantom_mode` - (Optional, Boolean) Whether the alert is in phantom mode (creating incidents or not).
* `priority` - (Optional, String) The priority of the alert definition.
  * Constraints: Allowable values are: `p5_or_unspecified`, `p4`, `p3`, `p2`, `p1`.
* `type` - (Required, String) Alert type.
  * Constraints: Allowable values are: `logs_immediate_or_unspecified`, `logs_threshold`, `logs_anomaly`, `logs_ratio_threshold`, `logs_new_value`, `logs_unique_count`, `logs_time_relative_threshold`, `metric_threshold`, `metric_anomaly`, `flow`.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the logs_alert_definition.
* `alert_def_id` - The unique identifier of the alert definition.
* `alert_version_id` - (String) The previous or old alert ID.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$/`.
* `created_time` - (String) The time when the alert definition was created.
* `updated_time` - (String) The time when the alert definition was last updated.


## Import

You can import the `ibm_logs_alert_definition` resource by using `id`. `id` Alert id is combination of `region`, `instance_id` and `alert_def_id`.

# Syntax
<pre>
$ terraform import ibm_logs_alert_definition.logs_alert_definition < region >/< instance_id >/< alert_id>;
</pre>

# Example
```
$ terraform import ibm_logs_alert_definition.logs_alert_definition eu-gb/3dc02998-0b50-4ea8-b68a-4779d716fa1f/4dc02998-0bc50-0b50-b68a-4779d716fa1f
```
