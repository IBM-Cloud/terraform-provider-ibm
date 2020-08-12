variable "ibmcloud_api_key" {
  description = "holds the user api key"
}

variable "resource_group" {
  description = "holds the resource group"
}

variable "network_id" {
description = "holds vpc crn id "
}
variable "network_account_id" {
description = "holds the ID of the account which owns the network that is being connected. Generally only used if the network is in a different account than the gateway"
}