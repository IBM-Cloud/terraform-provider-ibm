provider "ibm" {
  ibmcloud_api_key = var.ibmcloud_api_key
}

// Provision logs_alert resource instance
resource "ibm_logs_alert" "logs_alert_instance" {
  name = var.logs_alert_name
  description = var.logs_alert_description
  is_active = var.logs_alert_is_active
  severity = var.logs_alert_severity
  expiration {
    year = 2012
    month = 12
    day = 24
  }
  condition {
    immediate = {  }
  }
  notification_groups {
    group_by_fields = ["cpu"]
    notifications {
      retriggering_period_seconds = 60
      notify_on = "triggered_and_resolved"
      integration_id = 123
    }
  }
  filters {
    severities = ["critical"]
    metadata {
      applications = ["CpuMonitoring","WebApi"]
      subsystems = ["SnapshotGenerator","PermissionControl"]
    }
    alias = "monitorQuery"
    text = "_exists_:"container_name""
    ratio_alerts {
      alias = "TopLevelAlert"
      text = "_exists_:"container_name""
      severities = [ "critical" ]
      applications = ["CpuMonitoring","WebApi"]
      subsystems = ["SnapshotGenerator","PermissionControl"]
      group_by = ["Host","Thread"]
    }
    filter_type = "flow"
  }
  active_when {
    timeframes {
      days_of_week = ["sunday"]
      range {
        start {
          hours = 22
          minutes = 22
          seconds = 22
        }
        end {
          hours = 22
          minutes = 22
          seconds = 22
        }
      }
    }
  }
  notification_payload_filters = var.logs_alert_notification_payload_filters
  meta_labels {
    key = "ColorLabel"
    value = "Red"
  }
  meta_labels_strings = var.logs_alert_meta_labels_strings
  incident_settings {
    retriggering_period_seconds = 60
    notify_on = "triggered_and_resolved"
    use_as_notification_settings = true
  }
}

// Provision logs_rule_group resource instance
resource "ibm_logs_rule_group" "logs_rule_group_instance" {
  name = var.logs_rule_group_name
  description = var.logs_rule_group_description
  enabled = var.logs_rule_group_enabled
  rule_matchers {
    application_name {
      value = "my-application"
    }
  }
  rule_subgroups {
    id = "9fab83da-98cb-4f18-a7ba-b6f0435c9673"
    rules {
      id = "9fab83da-98cb-4f18-a7ba-b6f0435c9673"
      name = "Extract service and region"
      description = "Extracts the service and region from the source field"
      source_field = "logObj.source"
      parameters {
        extract_parameters {
          rule = "^http:\\/\\/my\\.service\\.com\\/#(?P<service>\\w+)\\-(?P<region>[^_]+)_"
        }
      }
      enabled = true
      order = 1
    }
    enabled = true
    order = 1
  }
  order = var.logs_rule_group_order
}

// Provision logs_outgoing_webhook resource instance
resource "ibm_logs_outgoing_webhook" "logs_outgoing_webhook_instance" {
  type = var.logs_outgoing_webhook_type
  name = var.logs_outgoing_webhook_name
  url = var.logs_outgoing_webhook_url
  ibm_event_notifications {
    event_notifications_instance_id = "9fab83da-98cb-4f18-a7ba-b6f0435c9673"
    region_id = "eu-es"
    source_id = "crn:v1:staging:public:logs:eu-gb:a/223af6f4260f42ebe23e95fcddd33cb7:63a3e4be-cb73-4f52-898e-8e93484a70a5::"
    source_name = "IBM Cloud Event Notifications"
    endpoint_type = "private"
  }
}

// Provision logs_policy resource instance
resource "ibm_logs_policy" "logs_policy_instance" {
  before {
    id = "9fab83da-98cb-4f18-a7ba-b6f0435c9673"
  }
  name = var.logs_policy_name
  description = var.logs_policy_description
  priority = var.logs_policy_priority
  enabled = var.logs_policy_enabled
  application_rule {
    rule_type_id = "includes"
    name = "Rule Name"
  }
  subsystem_rule {
    rule_type_id = "includes"
    name = "Rule Name"
  }
  archive_retention {
    id = "9fab83da-98cb-4f18-a7ba-b6f0435c9673"
  }
  log_rules {
    severities = ["critical"]
  }
}

// Provision logs_dashboard resource instance
resource "ibm_logs_dashboard" "logs_dashboard_instance" {
  href = var.logs_dashboard_href
  name = var.logs_dashboard_name
  description = var.logs_dashboard_description
  layout {
    sections {
      href = "href"
      id {
        value = "9fab83da-98cb-4f18-a7ba-b6f0435c9673"
      }
      rows {
        href = "href"
        id {
          value = "9fab83da-98cb-4f18-a7ba-b6f0435c9673"
        }
        appearance {
          height = 5
        }
        widgets {
          href = "href"
          id {
            value = "9fab83da-98cb-4f18-a7ba-b6f0435c9673"
          }
          title = "Response time"
          description = "The average response time of the system"
          definition {
            line_chart {
              legend {
                is_visible = true
                columns = ["name"]
                group_by_query = true
              }
              tooltip {
                show_labels = true
                type = "single"
              }
              query_definitions {
                id = "9fab83da-98cb-4f18-a7ba-b6f0435c9673"
                query {
                  logs {
                    lucene_query {
                      value = "coralogix.metadata.applicationName:"production""
                    }
                    group_by = [ "group_by" ]
                    aggregations {
                      count = {  }
                    }
                    filters {
                      operator {
                        equals {
                          selection {
                            all = {  }
                          }
                        }
                      }
                      observation_field {
                        keypath = ["applicationname"]
                        scope = "metadata"
                      }
                    }
                    group_bys {
                      keypath = ["applicationname"]
                      scope = "metadata"
                    }
                  }
                }
                series_name_template = "{{severity}}"
                series_count_limit = "10"
                unit = "usd"
                scale_type = "logarithmic"
                name = "CPU usage"
                is_visible = true
                color_scheme = "classic"
                resolution {
                  interval = "1m"
                  buckets_presented = 100
                }
                data_mode_type = "archive"
              }
              stacked_line = "relative"
            }
          }
          created_at = "2021-01-01T00:00:00.000Z"
          updated_at = "2021-01-01T00:00:00.000Z"
        }
      }
      options {
        internal = {  }
      }
    }
  }
  variables {
    name = "service_name"
    definition {
      multi_select {
        source {
          logs_path {
            observation_field {
              keypath = ["applicationname"]
              scope = "metadata"
            }
          }
        }
        selection {
          all = {  }
        }
        values_order_direction = "desc"
        selection_options {
          selection_type = "single"
        }
      }
    }
    display_name = "Service Name"
    description = "description"
    display_type = "nothing"
  }
  filters {
    source {
      logs {
        operator {
          equals {
            selection {
              all = {  }
            }
          }
        }
        observation_field {
          keypath = ["applicationname"]
          scope = "metadata"
        }
      }
    }
    enabled = true
    collapsed = true
  }
  annotations {
    href = "9fab83da-98cb-4f18-a7ba-b6f0435c9673"
    id = "9fab83da-98cb-4f18-a7ba-b6f0435c9673"
    name = "Deployments"
    enabled = true
    source {
      metrics {
        promql_query {
          value = "sum(up)"
        }
        strategy {
          start_time_metric = {  }
        }
        message_template = "message_template"
        labels = [ "labels" ]
      }
    }
  }
  absolute_time_frame {
    from = "2021-01-01T00:00:00.000Z"
    to = "2021-01-01T00:00:00.000Z"
  }
  relative_time_frame = var.logs_dashboard_relative_time_frame
  folder_id {
    value = "9fab83da-98cb-4f18-a7ba-b6f0435c9673"
  }
  folder_path {
    segments = ["production","payments"]
  }
  false {
  }
  two_minutes {
  }
  five_minutes {
  }
}

// Provision logs_dashboard_folder resource instance
resource "ibm_logs_dashboard_folder" "logs_dashboard_folder_instance" {
  name = var.logs_dashboard_folder_name
  parent_id = var.logs_dashboard_folder_parent_id
}

// Provision logs_e2m resource instance
resource "ibm_logs_e2m" "logs_e2m_instance" {
  name = var.logs_e2m_name
  description = var.logs_e2m_description
  metric_labels {
    target_label = "alias_label_name"
    source_field = "log_obj.string_value"
  }
  metric_fields {
    target_base_metric_name = "alias_field_name"
    source_field = "log_obj.numeric_field"
    aggregations {
      enabled = true
      agg_type = "samples"
      target_metric_name = "alias_field_name_agg_func"
      samples {
        sample_type = "max"
      }
    }
  }
  type = var.logs_e2m_type
  logs_query {
    lucene = "log_obj.numeric_field: [50 TO 100]"
    alias = "new_query"
    applicationname_filters = ["app_name"]
    subsystemname_filters = ["sub_name"]
    severity_filters = ["critical"]
  }
}

// Provision logs_view resource instance
resource "ibm_logs_view" "logs_view_instance" {
  name = var.logs_view_name
  search_query {
    query = "error"
  }
  time_selection {
    quick_selection {
      caption = "Last hour"
      seconds = 3600
    }
  }
  filters {
    filters {
      name = "applicationName"
      selected_values = {"cs-rest-test1":true,"demo":true}
    }
  }
  folder_id = var.logs_view_folder_id
}

// Provision logs_view_folder resource instance
resource "ibm_logs_view_folder" "logs_view_folder_instance" {
  name = var.logs_view_folder_name
}

// Provision logs_data_access_rule resource instance
resource "ibm_logs_data_access_rule" "logs_data_access_rule_instance" {
  display_name = var.logs_data_access_rule_display_name
  description = var.logs_data_access_rule_description
  filters {
    entity_type = "logs"
    expression = "true"
  }
  default_expression = var.logs_data_access_rule_default_expression
}

// Provision logs_enrichment resource instance
resource "ibm_logs_enrichment" "logs_enrichment_instance" {
  field_name = var.logs_enrichment_field_name
  enrichment_type {
    geo_ip = {  }
  }
}

// Provision logs_data_usage_metrics resource instance
resource "ibm_logs_data_usage_metrics" "logs_data_usage_metrics_instance" {
  enabled = var.logs_data_usage_metrics_enabled
}

// Provision logs_stream resource instance
resource "ibm_logs_stream" "logs_stream_instance" {
  name = var.logs_stream_name
  is_active = var.logs_stream_is_active
  dpxl_expression = var.logs_stream_dpxl_expression
  compression_type = var.logs_stream_compression_type
  ibm_event_streams {
    brokers = "kafka01.example.com:9093"
    topic = "live.screen"
  }
}

// Provision logs_alert_definition resource instance
resource "ibm_logs_alert_definition" "logs_alert_definition_instance" {
  name = var.logs_alert_definition_name
  description = var.logs_alert_definition_description
  enabled = var.logs_alert_definition_enabled
  priority = var.logs_alert_definition_priority
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
  type = var.logs_alert_definition_type
  group_by_keys = var.logs_alert_definition_group_by_keys
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
  entity_labels = var.logs_alert_definition_entity_labels
  phantom_mode = var.logs_alert_definition_phantom_mode
  deleted = var.logs_alert_definition_deleted
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

// Data source is not linked to a resource instance
// Uncomment if an existing data source instance exists
/*
// Create logs_alert data source
data "ibm_logs_alert" "logs_alert_instance" {
  logs_alert_id = var.data_logs_alert_logs_alert_id
}
*/

// Data source is not linked to a resource instance
// Uncomment if an existing data source instance exists
/*
// Create logs_alerts data source
data "ibm_logs_alerts" "logs_alerts_instance" {
}
*/

// Data source is not linked to a resource instance
// Uncomment if an existing data source instance exists
/*
// Create logs_rule_group data source
data "ibm_logs_rule_group" "logs_rule_group_instance" {
  group_id = var.data_logs_rule_group_group_id
}
*/

// Data source is not linked to a resource instance
// Uncomment if an existing data source instance exists
/*
// Create logs_rule_groups data source
data "ibm_logs_rule_groups" "logs_rule_groups_instance" {
}
*/

// Data source is not linked to a resource instance
// Uncomment if an existing data source instance exists
/*
// Create logs_outgoing_webhooks data source
data "ibm_logs_outgoing_webhooks" "logs_outgoing_webhooks_instance" {
  type = var.logs_outgoing_webhooks_type
}
*/

// Data source is not linked to a resource instance
// Uncomment if an existing data source instance exists
/*
// Create logs_outgoing_webhook data source
data "ibm_logs_outgoing_webhook" "logs_outgoing_webhook_instance" {
  logs_outgoing_webhook_id = var.data_logs_outgoing_webhook_logs_outgoing_webhook_id
}
*/

// Data source is not linked to a resource instance
// Uncomment if an existing data source instance exists
/*
// Create logs_policy data source
data "ibm_logs_policy" "logs_policy_instance" {
  logs_policy_id = var.data_logs_policy_logs_policy_id
}
*/

// Data source is not linked to a resource instance
// Uncomment if an existing data source instance exists
/*
// Create logs_policies data source
data "ibm_logs_policies" "logs_policies_instance" {
  enabled_only = var.logs_policies_enabled_only
  source_type = var.logs_policies_source_type
}
*/

// Data source is not linked to a resource instance
// Uncomment if an existing data source instance exists
/*
// Create logs_dashboard data source
data "ibm_logs_dashboard" "logs_dashboard_instance" {
  dashboard_id = var.data_logs_dashboard_dashboard_id
}
*/

// Data source is not linked to a resource instance
// Uncomment if an existing data source instance exists
/*
// Create logs_dashboard_folder data source
data "ibm_logs_dashboard_folder" "logs_dashboard_folder_instance" {
}
*/

// Data source is not linked to a resource instance
// Uncomment if an existing data source instance exists
/*
// Create logs_dashboard_folders data source
data "ibm_logs_dashboard_folders" "logs_dashboard_folders_instance" {
}
*/

// Data source is not linked to a resource instance
// Uncomment if an existing data source instance exists
/*
// Create logs_dashboards data source
data "ibm_logs_dashboards" "logs_dashboards_instance" {
}
*/

// Data source is not linked to a resource instance
// Uncomment if an existing data source instance exists
/*
// Create logs_e2m data source
data "ibm_logs_e2m" "logs_e2m_instance" {
  logs_e2m_id = var.data_logs_e2m_logs_e2m_id
}
*/

// Data source is not linked to a resource instance
// Uncomment if an existing data source instance exists
/*
// Create logs_e2ms data source
data "ibm_logs_e2ms" "logs_e2ms_instance" {
}
*/

// Data source is not linked to a resource instance
// Uncomment if an existing data source instance exists
/*
// Create logs_view data source
data "ibm_logs_view" "logs_view_instance" {
  logs_view_id = var.data_logs_view_logs_view_id
}
*/

// Data source is not linked to a resource instance
// Uncomment if an existing data source instance exists
/*
// Create logs_views data source
data "ibm_logs_views" "logs_views_instance" {
}
*/

// Data source is not linked to a resource instance
// Uncomment if an existing data source instance exists
/*
// Create logs_view_folder data source
data "ibm_logs_view_folder" "logs_view_folder_instance" {
  logs_view_folder_id = var.data_logs_view_folder_logs_view_folder_id
}
*/

// Data source is not linked to a resource instance
// Uncomment if an existing data source instance exists
/*
// Create logs_view_folders data source
data "ibm_logs_view_folders" "logs_view_folders_instance" {
}
*/

// Data source is not linked to a resource instance
// Uncomment if an existing data source instance exists
/*
// Create logs_data_access_rule data source
data "ibm_logs_data_access_rule" "logs_data_access_rule_instance" {
  logs_data_access_rule_id = var.data_logs_data_access_rule_logs_data_access_rule_id
}
*/

// Data source is not linked to a resource instance
// Uncomment if an existing data source instance exists
/*
// Create logs_data_access_rules data source
data "ibm_logs_data_access_rules" "logs_data_access_rules_instance" {
  logs_data_access_rules_id = var.logs_data_access_rules_logs_data_access_rules_id
}
*/

// Data source is not linked to a resource instance
// Uncomment if an existing data source instance exists
/*
// Create logs_enrichment data source
data "ibm_logs_enrichment" "logs_enrichment_instance" {
}
*/

// Data source is not linked to a resource instance
// Uncomment if an existing data source instance exists
/*
// Create logs_enrichments data source
data "ibm_logs_enrichments" "logs_enrichments_instance" {
}
*/

// Data source is not linked to a resource instance
// Uncomment if an existing data source instance exists
/*
// Create logs_data_usage_metrics data source
data "ibm_logs_data_usage_metrics" "logs_data_usage_metrics_instance" {
  range = var.data_logs_data_usage_metrics_range
  query = var.data_logs_data_usage_metrics_query
}
*/

// Data source is not linked to a resource instance
// Uncomment if an existing data source instance exists
/*
// Create logs_stream data source
data "ibm_logs_stream" "logs_stream_instance" {
}
*/

// Data source is not linked to a resource instance
// Uncomment if an existing data source instance exists
/*
// Create logs_streams data source
data "ibm_logs_streams" "logs_streams_instance" {
}
*/

// Data source is not linked to a resource instance
// Uncomment if an existing data source instance exists
/*
// Create logs_alert_definition data source
data "ibm_logs_alert_definition" "logs_alert_definition_instance" {
  logs_alert_definition_id = var.data_logs_alert_definition_logs_alert_definition_id
}
*/

// Data source is not linked to a resource instance
// Uncomment if an existing data source instance exists
/*
// Create logs_alert_definitions data source
data "ibm_logs_alert_definitions" "logs_alert_definitions_instance" {
}
*/
