variable "ibmcloud_api_key" {
  description = "IBM Cloud API key"
  type        = string
}

// Resource arguments for logs_alert
variable "logs_alert_name" {
  description = "Alert name."
  type        = string
  default     = "Unique count alert"
}
variable "logs_alert_description" {
  description = "Alert description."
  type        = string
  default     = "Example of unique count alert from terraform"
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
  default     = ["notification_payload_filters"]
}
variable "logs_alert_meta_labels_strings" {
  description = "The Meta labels to add to the alert as string with ':' separator."
  type        = list(string)
  default     = ["meta_labels_strings"]
}

// Resource arguments for logs_rule_group
variable "logs_rule_group_name" {
  description = "The name of the rule group."
  type        = string
  default     = "rule group"
}
variable "logs_rule_group_description" {
  description = "A description for the rule group, should express what is the rule group purpose."
  type        = string
  default     = "Rule group to extract severity from logs"
}
variable "logs_rule_group_creator" {
  description = "The creator of the rule group."
  type        = string
  default     = "terraform-rules-creator"
}
variable "logs_rule_group_enabled" {
  description = "Whether or not the rule is enabled."
  type        = bool
  default     = true
}
variable "logs_rule_group_order" {
  description = "// The order in which the rule group will be evaluated. The lower the order, the more priority the group will have. Not providing the order will by default create a group with the last order."
  type        = number
  default     = 0
}

// Resource arguments for logs_outgoing_webhook
variable "logs_outgoing_webhook_type" {
  description = "Outbound webhook type."
  type        = string
  default     = "ibm_event_notifications"
}
variable "logs_outgoing_webhook_name" {
  description = "The name of the outbound webhook."
  type        = string
  default     = "name"
}
variable "logs_outgoing_webhook_url" {
  description = "The URL of the outbound webhook."
  type        = string
  default     = "url"
}

// Resource arguments for logs_policy
variable "logs_policy_name" {
  description = "name of policy."
  type        = string
  default     = "name"
}
variable "logs_policy_description" {
  description = "description of policy."
  type        = string
  default     = "description"
}
variable "logs_policy_priority" {
  description = "the data pipeline sources that match the policy rules will go through."
  type        = string
  default     = "type_unspecified"
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
  default     = "relative_time_frame"
}
variable "logs_dashboard_false" {
  description = ""
  type        = bool
  default     = true
}
variable "logs_dashboard_two_minutes" {
  description = ""
  type        = bool
  default     = true
}
variable "logs_dashboard_five_minutes" {
  description = ""
  type        = bool
  default     = true
}

// Resource arguments for logs_e2m
variable "logs_e2m_name" {
  description = "E2M name."
  type        = string
  default     = "Service catalog latency"
}
variable "logs_e2m_description" {
  description = "E2m description."
  type        = string
  default     = "avg and max the latency of catalog service"
}
variable "logs_e2m_type" {
  description = "e2m type."
  type        = string
  default     = "unspecified"
}

// Resource arguments for logs_view
variable "logs_view_name" {
  description = "View name."
  type        = string
  default     = "name"
}
variable "logs_view_folder_id" {
  description = "View folder id."
  type        = string
  default     = "9fab83da-98cb-4f18-a7ba-b6f0435c9673"
}

// Resource arguments for logs_view_folder
variable "logs_view_folder_name" {
  description = "Folder name."
  type        = string
  default     = "My folder"
}

// Data source arguments for logs_alert
variable "data_logs_alert_logs_alert_id" {
  description = "Alert id."
  type        = string
  default     = "3dc02998-0b50-4ea8-b68a-4779d716fa1f"
}


// Data source arguments for logs_rule_group
variable "data_logs_rule_group_group_id" {
  description = "The group id."
  type        = string
  default     = "3dc02998-0b50-4ea8-b68a-4779d716fa1f"
}


// Data source arguments for logs_outgoing_webhooks
variable "logs_outgoing_webhooks_type" {
  description = "Outbound webhook type."
  type        = string
  default     = "placeholder"
}

// Data source arguments for logs_outgoing_webhook
variable "data_logs_outgoing_webhook_logs_outgoing_webhook_id" {
  description = "Outbound webhook ID."
  type        = string
  default     = "logs_outgoing_webhook_id"
}

// Data source arguments for logs_policy
variable "data_logs_policy_logs_policy_id" {
  description = "id of policy."
  type        = string
  default     = "logs_policy_id"
}

// Data source arguments for logs_policies
variable "logs_policies_enabled_only" {
  description = "optionally filter only enabled policies."
  type        = bool
  default     = false
}
variable "logs_policies_source_type" {
  description = "Source type to filter policies by."
  type        = string
  default     = "placeholder"
}

// Data source arguments for logs_dashboard
variable "data_logs_dashboard_dashboard_id" {
  description = "The ID of the dashboard."
  type        = string
  default     = "dashboard_id"
}

// Data source arguments for logs_e2m
variable "data_logs_e2m_logs_e2m_id" {
  description = "id of e2m to be deleted."
  type        = string
  default     = "logs_e2m_id"
}


// Data source arguments for logs_view
variable "data_logs_view_logs_view_id" {
  description = "View id."
  type        = number
  default     = 2
}


// Data source arguments for logs_view_folder
variable "data_logs_view_folder_logs_view_folder_id" {
  description = "Folder id."
  type        = string
  default     = "3dc02998-0b50-4ea8-b68a-4779d716fa1f"
}

