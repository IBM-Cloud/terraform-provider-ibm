provider "ibm" {
  ibmcloud_api_key = var.ibmcloud_api_key
}

resource "ibm_resource_instance" "logs_instance" {
  name     = "${var.resource_name_prefix}-instance1"
  service  = "logs"
  plan     = "beta"
  location = "eu-es"
}
resource "ibm_logs_alert" "logs_alert_instance" {
  instance_id = ibm_resource_instance.logs_instance.guid
  region      = ibm_resource_instance.logs_instance.location
  name        = "${var.resource_name_prefix}-alert"
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
resource "ibm_logs_rule_group" "logs_rule_group_instance" {
  instance_id = ibm_resource_instance.logs_instance.guid
  region      = ibm_resource_instance.logs_instance.location
  name        = "${var.resource_name_prefix}-rule-group"
  description = "example rule group decription"
  creator     = "bot@ibm.com"
  enabled     = false
  rule_matchers {
    subsystem_name {
      value = "mysql-cloudwatch"
    }
  }
  rule_subgroups {
    rules {
      name         = "mysql-parse"
      source_field = "text"
      parameters {
        parse_parameters {
          destination_field = "text"
          rule              = "(?P<timestamp>[^,]+),(?P<hostname>[^,]+),(?P<username>[^,]+),(?P<ip>[^,]+),(?P<connectionId>[0-9]+),(?P<queryId>[0-9]+),(?P<operation>[^,]+),(?P<database>[^,]+),'?(?P<object>.*)'?,(?P<returnCode>[0-9]+)"
        }
      }
      enabled = true
      order   = 1
    }

    enabled = true
    order   = 1
  }
  order = 4294967
}
resource "ibm_logs_policy" "logs_policy_instance" {
  instance_id = ibm_resource_instance.logs_instance.guid
  region      = ibm_resource_instance.logs_instance.location
  name        = "${var.resource_name_prefix}-policy"
  description = "example policy decription"
  priority    = "type_medium"
  application_rule {
    name         = "otel-links-test"
    rule_type_id = "start_with"
  }
  log_rules {
    severities = ["info"]
  }
}
resource "ibm_logs_e2m" "logs_e2m_instance" {
  instance_id = ibm_resource_instance.logs_instance.guid
  region      = ibm_resource_instance.logs_instance.location
  name        = "${var.resource_name_prefix}-e2m"
  description = "example E2M decription"
  logs_query {
    applicationname_filters = []
    severity_filters = [
      "debug", "error"
    ]
    subsystemname_filters = []
  }
  type = "logs2metrics"
}
resource "ibm_logs_view" "logs_view_instance" {
  instance_id = ibm_resource_instance.logs_instance.guid
  region      = ibm_resource_instance.logs_instance.location
  name        = "${var.resource_name_prefix}-view"
  filters {
    filters {
      name = "applicationName"
      selected_values = {
        demo = true
      }
    }
    filters {
      name = "subsystemName"
      selected_values = {
        demo = true
      }
    }
    filters {
      name = "operationName"
      selected_values = {
        demo = true
      }
    }
    filters {
      name = "serviceName"
      selected_values = {
        demo = true
      }
    }
    filters {
      name = "severity"
      selected_values = {
        demo = true
      }
    }
  }
  search_query {
    query = "logs"
  }
  time_selection {
    custom_selection {
      from_time = "2024-01-25T11:31:43.152Z"
      to_time   = "2024-01-25T11:37:13.238Z"
    }
  }
}
resource "ibm_logs_view_folder" "logs_view_folder_instance" {
  instance_id = ibm_resource_instance.logs_instance.guid
  region      = ibm_resource_instance.logs_instance.location
  name        = "${var.resource_name_prefix}-view-folder"
}
resource "ibm_logs_outgoing_webhook" "logs_outgoing_webhook_instance" {
  instance_id = ibm_resource_instance.logs_instance.guid
  region      = ibm_resource_instance.logs_instance.location
  name        = "${var.resource_name_prefix}-webhook"
  type        = "ibm_event_notifications"
  ibm_event_notifications {
    event_notifications_instance_id = "6b33da73-28b6-4201-bfea-b2054bb6ae8a"
    region_id                       = "us-south"
  }
}
resource "ibm_logs_dashboard" "logs_dashboard_instance" {
  instance_id = ibm_resource_instance.logs_instance.guid
  region      = ibm_resource_instance.logs_instance.location
  name        = "${var.resource_name_prefix}-dashboard"
  description = "example dashboard description"
  layout {
    sections {
      id {
        value = "b9ca2f71-7d7c-10fb-1a08-c78912705095"
      }
      rows {
        id {
          value = "70b12716-cb18-f933-5a89-3061734eaa2f"
        }
        appearance {
          height = 19
        }
        widgets {
          id {
            value = "6118b86d-860c-c2cb-0cdf-effd62e9f331"
          }
          title       = "test"
          description = "test"
          definition {
            line_chart {
              legend {
                is_visible     = true
                group_by_query = true
              }
              tooltip {
                show_labels = false
                type        = "all"
              }
              query_definitions {
                id           = "13139dad-3d45-16e1-fce2-03517daa71c4"
                color_scheme = "cold"
                name         = "Query 1"
                is_visible   = true
                scale_type   = "linear"
                resolution {
                  buckets_presented = 96
                }
                series_count_limit = 20
                query {
                  logs {
                    group_by = []

                    aggregations {
                      min {
                        observation_field {
                          keypath = [
                            "timestamp",
                          ]
                          scope = "metadata"
                        }
                      }
                    }

                    group_bys {
                      keypath = [
                        "severity",
                      ]
                      scope = "metadata"
                    }
                  }
                }
              }
            }
          }
        }
      }
    }
  }
  filters {
    source {
      logs {
        operator {
          equals {
            selection {
              list {}
            }
          }
        }
        observation_field {
          keypath = ["applicationname"]
          scope   = "label"
        }
      }
    }
    enabled   = true
    collapsed = false
  }
  filters {
    source {
      logs {
        # field = "field"
        operator {
          equals {
            selection {
              all {}
            }
          }
        }
        observation_field {
          keypath = ["subsystemname"]
          scope   = "label"
        }
      }
    }
    enabled   = true
    collapsed = false
  }
  relative_time_frame = "900s"
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
