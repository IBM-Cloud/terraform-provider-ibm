variable "ibmcloud_api_key" {
  description = "IBM Cloud API key"
  type        = string
}

// Resource arguments for account_settings_template_assignment
variable "account_settings_template_assignment_template_id" {
  description = "Template Id."
  type        = string
  default     = null
}
variable "account_settings_template_assignment_template_version" {
  description = "Template version."
  type        = number
  default     = 1
}
variable "account_settings_template_assignment_target_type" {
  description = "Assignment target type."
  type        = string
}
variable "account_settings_template_assignment_target" {
  description = "Assignment target."
  type        = string
}

// Data source arguments for account_settings_template_assignment
variable "account_settings_template_assignment_assignment_id" {
  description = "ID of the Assignment Record."
  type        = string
  default     = null
}
variable "account_settings_template_assignment_include_history" {
  description = "Defines if the entity history is included in the response."
  type        = bool
  default     = false
}
