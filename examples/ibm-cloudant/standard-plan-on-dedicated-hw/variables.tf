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

// Resource arguments for cloudant_dedicated_hardware_name
variable "cloudant_dedicated_hardware_name" {
  description = "The service instance name of the IBM Cloudant dedicated hardware plan."
  type = string
}
