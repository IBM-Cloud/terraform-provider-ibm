variable "ibmcloud_api_key" {
  description = "IBM Cloud API key"
  type        = string
}

// Data source arguments for is_dedicated_host_profile
variable "is_dedicated_host_profile_name" {
  description = "The globally unique name for this virtual server instance profile."
  type        = string
  default     = "bc1-4x16"
}

// Data source arguments for is_dedicated_host_profiles
variable "is_dedicated_host_profiles_name" {
  description = "The globally unique name for this dedicated host profile."
  type        = string
  default     = "mx2-host-152x1216"
}
