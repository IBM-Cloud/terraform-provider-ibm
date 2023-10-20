variable "ibmcloud_api_key" {
  description = "IBM Cloud API key"
  type        = string
}

// Resource arguments for account_settings_template
variable "account_settings_template_name" {
  description = "The name of the trusted profile template. This is visible only in the enterprise account."
  type        = string
}
variable "account_settings_template_description" {
  description = "The description of the trusted profile template. Describe the template for enterprise account users."
  type        = string
  default     = null
}

// Data source arguments for account_settings_template
variable "account_settings_template_template_id" {
  description = "ID of the account settings template."
  type        = string
  default     = null
}
variable "account_settings_template_version" {
  description = "Version of the account settings template."
  type        = string
  default     = null
}
variable "account_settings_template_include_history" {
  description = "Defines if the entity history is included in the response."
  type        = bool
  default     = false
}
