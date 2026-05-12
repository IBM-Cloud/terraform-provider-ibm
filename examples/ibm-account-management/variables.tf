variable "ibmcloud_api_key" {
  description = "IBM Cloud API key"
  type        = string
}

// Data source arguments for account
variable "account_account_id" {
  description = "The unique identifier of the account you want to retrieve."
  type        = string
  default     = "account_id"
}
