// Resource arguments for ibmcloud_api_key
variable "ibmcloud_api_key" {
  description = "IBM Cloud API key."
  type        = string
  sensitive   = true
}

// Resource arguments for service_region
variable "service_region" {
  description = "Region in which service has to be provisioned."
  type        = string
  default     = "us-south"
}
