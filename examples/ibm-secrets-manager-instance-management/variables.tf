variable "ibmcloud_api_key" {
  description = "IBM Cloud API key"
  type        = string
}

// Data source arguments for sm_instance
variable "sm_instance_instance_id" {
  description = "The service instance ID."
  type        = string
  default     = "60b40daa-1fd3-4f35-a994-2409cc0f270c"
}
