variable "ibmcloud_api_key" {
  description = "IBM Cloud API key"
  type        = string
}

// Resource arguments

variable "account_id" {
  description = "The ID of the account to target found in: https://cloud.ibm.com/account/settings"
  type        = string
}

variable "resource_group_id" {
  description = "The ID of the account's resource group to target found in: https://cloud.ibm.com/account/resource-groups"
  type        = string
}
