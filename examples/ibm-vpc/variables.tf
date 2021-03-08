variable "ibmcloud_api_key" {
  description = "IBM Cloud API key"
  type        = string
}

// Data source arguments for is_dedicated_host
variable "is_dedicated_host_id" {
  description = "The unique identifier for this virtual server instance."
  type        = string
  default     = "1e09281b-f177-46fb-baf1-bc152b2e391a"
}

// Data source arguments for is_dedicated_hosts
variable "is_dedicated_hosts_id" {
  description = "The unique identifier for this dedicated host."
  type        = string
  default     = "1e09281b-f177-46fb-baf1-bc152b2e391a"
}
