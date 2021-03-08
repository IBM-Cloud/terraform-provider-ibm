variable "ibmcloud_api_key" {
  description = "IBM Cloud API key"
  type        = string
}

// Data source arguments for is_dedicated_hosts
variable "is_dedicated_hosts_name" {
  description = "The unique user-defined name for this dedicated host. If unspecified, the name will be a hyphenated list of randomly-selected words."
  type        = string
  default     = "my-host"
}
