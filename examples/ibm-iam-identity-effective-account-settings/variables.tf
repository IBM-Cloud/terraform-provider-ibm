variable "ibmcloud_api_key" {
  description = "IBM Cloud API key"
  type        = string
}

// Data source arguments for iam_effective_account_settings
variable "iam_effective_account_settings_account_id" {
  description = "Unique ID of the account."
  type        = string
  default     = "account_id"
}
variable "iam_effective_account_settings_include_history" {
  description = "Defines if the entity history is included in the response."
  type        = bool
  default     = false
}
variable "iam_effective_account_settings_resolve_user_mfa" {
  description = "Enrich MFA exemptions with user information."
  type        = bool
  default     = false
}
