provider "ibm" {
  ibmcloud_api_key = var.ibmcloud_api_key
}

// Provision logs_alert resource instance
resource "ibm_logs_alert" "logs_alert_instance" {
  name        = var.logs_alert_name
  description = var.logs_alert_description
  is_active   = var.logs_alert_is_active
  severity    = var.logs_alert_severity
  expiration {
    year  = 1
    month = 1
    day   = 1
  }
  condition {
    immediate = {}
  }
  notification_groups {
    group_by_fields = ["group_by_fields"]
    notifications {
      retriggering_period_seconds = 0
      notify_on                   = "triggered_only"
      integration_id              = 0
    }
  }
  filters {
    severities = ["debug_or_unspecified"]
    metadata {
      categories   = ["categories"]
      applications = ["applications"]
      subsystems   = ["subsystems"]
      computers    = ["computers"]
      classes      = ["classes"]
      methods      = ["methods"]
      ip_addresses = ["ip_addresses"]
    }
    alias = "alias"
    text  = "text"
    ratio_alerts {
      alias        = "alias"
      text         = "text"
      severities   = ["debug_or_unspecified"]
      applications = ["applications"]
      subsystems   = ["subsystems"]
      group_by     = ["group_by"]
    }
    filter_type = "text_or_unspecified"
  }
  active_when {
    timeframes {
      days_of_week = ["monday_or_unspecified"]
      range {
        start {
          hours   = 1
          minutes = 1
          seconds = 1
        }
        end {
          hours   = 1
          minutes = 1
          seconds = 1
        }
      }
    }
  }
  notification_payload_filters = var.logs_alert_notification_payload_filters
  meta_labels {
    key   = "key"
    value = "value"
  }
  meta_labels_strings = var.logs_alert_meta_labels_strings
  tracing_alert {
    condition_latency = 0
    field_filters {
      field = "field"
      filters {
        values   = ["values"]
        operator = "operator"
      }
    }
    tag_filters {
      field = "field"
      filters {
        values   = ["values"]
        operator = "operator"
      }
    }
  }
  incident_settings {
    retriggering_period_seconds  = 0
    notify_on                    = "triggered_only"
    use_as_notification_settings = true
  }
}

// Provision logs_rule_group resource instance
resource "ibm_logs_rule_group" "logs_rule_group_instance" {
  name        = var.logs_rule_group_name
  description = var.logs_rule_group_description
  creator     = var.logs_rule_group_creator
  enabled     = var.logs_rule_group_enabled
  rule_matchers {
    application_name {
      value = "value"
    }
  }
  rule_subgroups {
    id = "9fab83da-98cb-4f18-a7ba-b6f0435c9673"
    rules {
      id           = "9fab83da-98cb-4f18-a7ba-b6f0435c9673"
      name         = "name"
      description  = "description"
      source_field = "logObj.source"
      parameters {
        extract_parameters {
          rule = "rule"
        }
      }
      enabled = true
      order   = 0
    }
    enabled = true
    order   = 0
  }
  order = var.logs_rule_group_order
}

// Provision logs_outgoing_webhook resource instance
resource "ibm_logs_outgoing_webhook" "logs_outgoing_webhook_instance" {
  type = var.logs_outgoing_webhook_type
  name = var.logs_outgoing_webhook_name
  url  = var.logs_outgoing_webhook_url
  ibm_event_notifications {
    event_notifications_instance_id = "9fab83da-98cb-4f18-a7ba-b6f0435c9673"
    region_id                       = "region_id"
  }
}

// Provision logs_policy resource instance
resource "ibm_logs_policy" "logs_policy_instance" {
  name        = var.logs_policy_name
  description = var.logs_policy_description
  priority    = var.logs_policy_priority
  application_rule {
    rule_type_id = "unspecified"
    name         = "name"
  }
  subsystem_rule {
    rule_type_id = "unspecified"
    name         = "name"
  }
  archive_retention {
    id = "id"
  }
  log_rules {
    severities = ["unspecified"]
  }
}

// Provision logs_dashboard resource instance
resource "ibm_logs_dashboard" "logs_dashboard_instance" {
  href        = var.logs_dashboard_href
  name        = var.logs_dashboard_name
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
          title       = "Response time"
          description = "The average response time of the system"
          definition {
            line_chart {
              legend {
                is_visible     = true
                columns        = ["unspecified"]
                group_by_query = true
              }
              tooltip {
                show_labels = true
                type        = "unspecified"
              }
              query_definitions {
                id = "9fab83da-98cb-4f18-a7ba-b6f0435c9673"
                query {
                  logs {
                    lucene_query {
                      value = "coralogix.metadata.applicationName:'production'"
                    }
                    group_by = ["group_by"]
                    aggregations {
                      count = {}
                    }
                    filters {
                      operator {
                        equals {
                          selection {
                            all = {}
                          }
                        }
                      }
                      observation_field {
                        keypath = ["keypath"]
                        scope   = "unspecified"
                      }
                    }
                    group_bys {
                      keypath = ["keypath"]
                      scope   = "unspecified"
                    }
                  }
                }
                series_name_template = "{{severity}}"
                series_count_limit   = "10"
                unit                 = "unspecified"
                scale_type           = "unspecified"
                name                 = "CPU usage"
                is_visible           = true
                color_scheme         = "classic"
                resolution {
                  interval          = "1m"
                  buckets_presented = 100
                }
                data_mode_type = "high_unspecified"
              }
            }
          }
        }
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
              keypath = ["keypath"]
              scope   = "unspecified"
            }
          }
        }
        selection {
          all = {}
        }
        values_order_direction = "unspecified"
      }
    }
    display_name = "Service Name"
  }
  filters {
    source {
      logs {
        operator {
          equals {
            selection {
              all = {}
            }
          }
        }
        observation_field {
          keypath = ["keypath"]
          scope   = "unspecified"
        }
      }
    }
    enabled   = true
    collapsed = true
  }
  annotations {
    href    = "9fab83da-98cb-4f18-a7ba-b6f0435c9673"
    id      = "9fab83da-98cb-4f18-a7ba-b6f0435c9673"
    name    = "Deployments"
    enabled = true
    source {
      metrics {
        promql_query {
          value = "sum(up)"
        }
        strategy {
          start_time_metric = {}
        }
        message_template = "message_template"
        labels           = ["labels"]
      }
    }
  }
  absolute_time_frame {
    from = "2021-01-31T09:44:12Z"
    to   = "2021-01-31T09:44:12Z"
  }
  relative_time_frame = var.logs_dashboard_relative_time_frame
  folder_id {
    value = "9fab83da-98cb-4f18-a7ba-b6f0435c9673"
  }
  folder_path {
    segments = ["segments"]
  }
  false        = var.logs_dashboard_false
  two_minutes  = var.logs_dashboard_two_minutes
  five_minutes = var.logs_dashboard_five_minutes
}

// Provision logs_e2m resource instance
resource "ibm_logs_e2m" "logs_e2m_instance" {
  name        = var.logs_e2m_name
  description = var.logs_e2m_description
  metric_labels {
    target_label = "target_label"
    source_field = "source_field"
  }
  metric_fields {
    target_base_metric_name = "target_base_metric_name"
    source_field            = "source_field"
    aggregations {
      enabled            = true
      agg_type           = "unspecified"
      target_metric_name = "target_metric_name"
      samples {
        sample_type = "unspecified"
      }
    }
  }
  type = var.logs_e2m_type
  spans_query {
    lucene                  = "lucene"
    applicationname_filters = ["applicationname_filters"]
    subsystemname_filters   = ["subsystemname_filters"]
    action_filters          = ["action_filters"]
    service_filters         = ["service_filters"]
  }
  logs_query {
    lucene                  = "lucene"
    alias                   = "alias"
    applicationname_filters = ["applicationname_filters"]
    subsystemname_filters   = ["subsystemname_filters"]
    severity_filters        = ["unspecified"]
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
      name            = "applicationName"
      selected_values = { "cs-rest-test1" : true, "demo" : true }
    }
  }
  folder_id = var.logs_view_folder_id
}

// Provision logs_view_folder resource instance
resource "ibm_logs_view_folder" "logs_view_folder_instance" {
  name = var.logs_view_folder_name
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
