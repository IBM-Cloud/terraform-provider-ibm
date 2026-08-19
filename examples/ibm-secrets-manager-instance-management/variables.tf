variable "ibmcloud_api_key" {
  description = "IBM Cloud API key"
  type        = string
}

// Data source arguments for sm_instance
variable "sm_instance_instance_id" {
  description = "Secrets Manager instance ID."
  type        = string
  default     = "a1b2c3d4-e5f6-7890-abcd-ef1234567890"
}
