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
variable "iam_serviceid_group_description" {
  description = "Description of the service ID group."
  type        = string
  default     = "description"
}

// Data source arguments for iam_serviceid_group
variable "data_iam_serviceid_group_iam_serviceid_group_id" {
  description = "Unique ID of the service ID group."
  type        = string
  default     = "iam_serviceid_group_id"
}
