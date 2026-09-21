variable "ibmcloud_api_key" {
  description = "IBM Cloud API key"
  type        = string
}

// Resource arguments for iam_serviceid_group
variable "iam_serviceid_group_account_id" {
  description = "ID of the account the service ID group belongs to."
  type        = string
  default     = "account_id"
}
variable "iam_serviceid_group_name" {
  description = "Name of the service ID group. Unique in the account."
  type        = string
  default     = "name"
}

