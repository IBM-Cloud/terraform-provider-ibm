variable "ibmcloud_api_key" {
  description = "IBM Cloud API key"
  type        = string
}

// Data source arguments for sm_instance
variable "sm_instance_instance_id" {
  description = "The service instance ID."
  type        = string
  default     = "bfc50c2e-d66d-4f37-9ccf-9713f8325b39"
}
