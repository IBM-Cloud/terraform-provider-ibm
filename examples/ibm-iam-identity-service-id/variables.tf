variable "ibmcloud_api_key" {
  description = "IBM Cloud API key"
  type        = string
}

// Resource arguments for iam_service_id
variable "iam_service_id_name" {
  description = "Name of the ServiceID. The name is not checked for uniqueness."
  type        = string
  default     = "name"
}
