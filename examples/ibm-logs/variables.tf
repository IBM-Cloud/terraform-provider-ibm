variable "ibmcloud_api_key" {
  description = "IBM Cloud API key"
  type        = string
}

// Resource arguments for logs_alert
variable "logs_alert_name" {
  description = "Alert name."
  type        = string
  default     = "Test alert"
}
variable "logs_alert_description" {
  description = "Alert description."
  type        = string
  default     = "Alert if the number of logs reaches a threshold"
}
variable "logs_alert_is_active" {
  description = "Alert is active."
  type        = bool
  default     = true
}
variable "logs_alert_severity" {
  description = "Alert severity."
  type        = string
  default     = "info_or_unspecified"
}
variable "logs_alert_notification_payload_filters" {
  description = "JSON keys to include in the alert notification, if left empty get the full log text in the alert notification."
  type        = list(string)
  default     = [ "notification_payload_filters" ]
}
variable "logs_alert_meta_labels_strings" {
  description = "The Meta labels to add to the alert as string with ':' separator."
  type        = list(string)
  default     = []
}

// Resource arguments for logs_rule_group
variable "logs_rule_group_name" {
  description = "The name of the rule group."
  type        = string
  default     = "mysql-extractrule"
}
variable "logs_rule_group_description" {
  description = "A description for the rule group, should express what is the rule group purpose."
  type        = string
  default     = "mysql audit logs  parser"
}
variable "logs_rule_group_enabled" {
  description = "Whether or not the rule is enabled."
  type        = bool
  default     = true
}
variable "logs_rule_group_order" {
  description = "// The order in which the rule group will be evaluated. The lower the order, the more priority the group will have. Not providing the order will by default create a group with the last order."
  type        = number
  default     = 39
}

// Resource arguments for logs_outgoing_webhook
variable "logs_outgoing_webhook_type" {
  description = "The type of the deployed Outbound Integrations to list."
  type        = string
  default     = "ibm_event_notifications"
}
variable "logs_outgoing_webhook_name" {
  description = "The name of the Outbound Integration."
  type        = string
  default     = "Event Notifications Integration"
}
variable "logs_outgoing_webhook_url" {
  description = "The URL of the Outbound Integration. Null for IBM Event Notifications integration."
  type        = string
  default     = "https://example.com"
}

// Resource arguments for logs_policy
variable "logs_policy_name" {
  description = "Name of policy."
  type        = string
  default     = "My Policy"
}
variable "logs_policy_description" {
  description = "Description of policy."
  type        = string
  default     = "My Policy Description"
}
variable "logs_policy_priority" {
  description = "The data pipeline sources that match the policy rules will go through."
  type        = string
  default     = "type_high"
}
variable "logs_policy_enabled" {
  description = "Flag to enable or disable a policy. This flag is supported only while updating a policy, since the policies are always enabled during creation."
  type        = bool
  default     = true
}

// Resource arguments for logs_dashboard
variable "logs_dashboard_href" {
  description = "Unique identifier for the dashboard."
  type        = string
  default     = "6U1Q8Hpa263Se8PkRKaiE"
}
variable "logs_dashboard_name" {
  description = "Display name of the dashboard."
  type        = string
  default     = "My Dashboard"
}
variable "logs_dashboard_description" {
  description = "Brief description or summary of the dashboard's purpose or content."
  type        = string
  default     = "This dashboard shows the performance of our production environment."
}
variable "logs_dashboard_relative_time_frame" {
  description = "Relative time frame specifying a duration from the current time."
  type        = string
  default     = "1d"
}

// Resource arguments for logs_dashboard_folder
variable "logs_dashboard_folder_name" {
  description = "The dashboard folder name, required."
  type        = string
  default     = "My Folder"
}
variable "logs_dashboard_folder_parent_id" {
  description = "The dashboard folder parent ID, optional. If not set, the folder is a root folder, if set, the folder is a subfolder of the parent folder and needs to be a uuid."
  type        = string
  default     = "9fab83da-98cb-4f18-a7ba-b6f0435c9673"
}

// Resource arguments for logs_e2m
variable "logs_e2m_name" {
  description = "Name of the E2M."
  type        = string
  default     = "Service catalog latency"
}
variable "logs_e2m_description" {
  description = "Description of the E2M."
  type        = string
  default     = "avg and max the latency of catalog service"
}
variable "logs_e2m_type" {
  description = "E2M type."
  type        = string
  default     = "logs2metrics"
}

// Resource arguments for logs_view
variable "logs_view_name" {
  description = "View name."
  type        = string
  default     = "Logs view"
}
variable "logs_view_folder_id" {
  description = "View folder ID."
  type        = string
  default     = "9fab83da-98cb-4f18-a7ba-b6f0435c9673"
}

// Resource arguments for logs_view_folder
variable "logs_view_folder_name" {
  description = "Folder name."
  type        = string
  default     = "My Folder"
}

// Resource arguments for logs_data_access_rule
variable "logs_data_access_rule_display_name" {
  description = "Data Access Rule Display Name."
  type        = string
  default     = "Test Data Access Rule"
}
variable "logs_data_access_rule_description" {
  description = "Optional Data Access Rule Description."
  type        = string
  default     = "Data Access Rule intended for testing"
}
variable "logs_data_access_rule_default_expression" {
  description = "Default expression to use when no filter matches the query."
  type        = string
  default     = "<v1>true"
}

// Resource arguments for logs_enrichment
variable "logs_enrichment_field_name" {
  description = "The enrichment field name."
  type        = string
  default     = "ip"
}

// Resource arguments for logs_data_usage_metrics
variable "logs_data_usage_metrics_enabled" {
  description = "The \"enabled\" parameter for metrics export."
  type        = bool
  default     = true
}

// Resource arguments for logs_stream
variable "logs_stream_name" {
  description = "The name of the Event stream."
  type        = string
  default     = "Live Screen"
}
variable "logs_stream_is_active" {
  description = "Whether the Event stream is active."
  type        = bool
  default     = true
}
variable "logs_stream_dpxl_expression" {
  description = "The DPXL expression of the Event stream."
  type        = string
  default     = "<v1>contains(kubernetes.labels.CX_AZ, 'eu-west-1')"
}
variable "logs_stream_compression_type" {
  description = "The compression type of the stream."
  type        = string
  default     = "gzip"
}

// Resource arguments for logs_alert_definition
variable "logs_alert_definition_name" {
  description = "The name of the alert definition."
  type        = string
  default     = "Unique count alert"
}
variable "logs_alert_definition_description" {
  description = "A detailed description of what the alert monitors and when it triggers."
  type        = string
  default     = "Example of unique count alert from terraform"
}
variable "logs_alert_definition_enabled" {
  description = "Whether the alert is currently active and monitoring."
  type        = bool
  default     = true
}
variable "logs_alert_definition_priority" {
  description = "The priority of the alert definition."
  type        = string
  default     = "p1"
}
variable "logs_alert_definition_type" {
  description = "Alert type."
  type        = string
  default     = "flow"
}
variable "logs_alert_definition_group_by_keys" {
  description = "Keys used to group and aggregate alert data."
  type        = list(string)
  default     = ["key1","key2"]
}
variable "logs_alert_definition_entity_labels" {
  description = "Labels used to identify and categorize the alert entity."
  type        = map(string)
  default     = {"key":"value"}
}
variable "logs_alert_definition_phantom_mode" {
  description = "Whether the alert is in phantom mode (creating incidents or not)."
  type        = bool
  default     = false
}
variable "logs_alert_definition_deleted" {
  description = "Whether the alert has been marked as deleted."
  type        = bool
  default     = false
}

// Data source arguments for logs_alert
variable "data_logs_alert_logs_alert_id" {
  description = "Alert ID."
  type        = 
  default     = 3dc02998-0b50-4ea8-b68a-4779d716fa1f
}


// Data source arguments for logs_rule_group
variable "data_logs_rule_group_group_id" {
  description = "The group ID."
  type        = 
  default     = 3dc02998-0b50-4ea8-b68a-4779d716fa1f
}


// Data source arguments for logs_outgoing_webhooks
variable "logs_outgoing_webhooks_type" {
  description = "The type of the deployed Outbound Integrations to list."
  type        = string
  default     = "ibm_event_notifications"
}

// Data source arguments for logs_outgoing_webhook
variable "data_logs_outgoing_webhook_logs_outgoing_webhook_id" {
  description = "The ID of the Outbound Integration to delete."
  type        = 
  default     = 585bea36-bdd1-4bfb-9a26-51f1f8a12660
}

// Data source arguments for logs_policy
variable "data_logs_policy_logs_policy_id" {
  description = "ID of policy."
  type        = 
  default     = 3dc02998-0b50-4ea8-b68a-4779d716fa1f
}

// Data source arguments for logs_policies
variable "logs_policies_enabled_only" {
  description = "Optionally filter only enabled policies."
  type        = bool
  default     = true
}
variable "logs_policies_source_type" {
  description = "Source type to filter policies by."
  type        = string
  default     = "logs"
}

// Data source arguments for logs_dashboard
variable "data_logs_dashboard_dashboard_id" {
  description = "The ID of the dashboard."
  type        = string
  default     = "dashboard_id"
}




// Data source arguments for logs_e2m
variable "data_logs_e2m_logs_e2m_id" {
  description = "ID of e2m to be deleted."
  type        = string
  default     = "d6a3658e-78d2-47d0-9b81-b2c551f01b09"
}


// Data source arguments for logs_view
variable "data_logs_view_logs_view_id" {
  description = "View ID."
  type        = number
  default     = 52
}


// Data source arguments for logs_view_folder
variable "data_logs_view_folder_logs_view_folder_id" {
  description = "Folder ID."
  type        = 
  default     = 3dc02998-0b50-4ea8-b68a-4779d716fa1f
}


// Data source arguments for logs_data_access_rule
variable "data_logs_data_access_rule_logs_data_access_rule_id" {
  description = "Array of data access rule IDs."
  type        = list()
  default     = ["4f966911-4bda-407e-b069-477394effa59"]
}

// Data source arguments for logs_data_access_rules
variable "logs_data_access_rules_logs_data_access_rules_id" {
  description = "Array of data access rule IDs."
  type        = list()
  default     = ["4f966911-4bda-407e-b069-477394effa59"]
}



// Data source arguments for logs_data_usage_metrics
variable "data_logs_data_usage_metrics_range" {
  description = "Range of days to get the data usage for, by default it will use current month."
  type        = string
  default     = "last_week"
}
variable "data_logs_data_usage_metrics_query" {
  description = "Query to filter daily or detailed the data usage, by default it will use daily one."
  type        = string
  default     = "daily"
}



// Data source arguments for logs_alert_definition
variable "data_logs_alert_definition_logs_alert_definition_id" {
  description = "Alert definition ID."
  type        = 
  default     = 3dc02998-0b50-4ea8-b68a-4779d716fa1f
}

