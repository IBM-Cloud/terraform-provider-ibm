variable "ibmcloud_api_key" {
  description = "IBM Cloud API key"
  type        = string
}

// Variables for iam_account_settings
variable "iam_account_settings_account_id" {
  description = "Unique ID of the account."
  type        = string
  default     = "account_id"
}

variable "iam_account_settings_ibmid1" {
  description = "Defines the IBM id to target for user_mfa setting."
  type        = string
}
