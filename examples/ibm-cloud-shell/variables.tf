variable "ibmcloud_api_key" {
  description = "IBM Cloud API key"
  type        = string
}

// Resource and Data source arguments for cloud_shell_account_settings
variable "cloud_shell_account_settings_account_id" {
  description = "The account ID in which the account settings belong to."
  type        = string
  default     = "account_id"
}
variable "cloud_shell_account_settings_default_enable_new_features" {
  description = "You can choose which Cloud Shell features are available in the account and whether any new features are enabled as they become available. The feature settings apply only to the enabled Cloud Shell locations."
  type        = bool
  default     = false
}
variable "cloud_shell_account_settings_default_enable_new_regions" {
  description = "Set whether Cloud Shell is enabled in a specific location for the account. The location determines where user and session data are stored. By default, users are routed to the nearest available location."
  type        = bool
  default     = false
}
variable "cloud_shell_account_settings_enabled" {
  description = "When enabled, Cloud Shell is available to all users in the account."
  type        = bool
  default     = false
}
variable "cloud_shell_account_settings_features" {
  description = "List of Cloud Shell features."
  type        = list(object({ example=string }))
  default     = [ { example: "object" } ]
}
variable "cloud_shell_account_settings_regions" {
  description = "List of Cloud Shell region settings."
  type        = list(object({ example=string }))
  default     = [ { example: "object" } ]
}
variable "cloud_shell_account_settings_tags" {
  description = "Local tags associated with cloud_shell_account_settings"
  type        = set(string)
  default     = []
}
