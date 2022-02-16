variable "ibmcloud_api_key" {
  description = "IBM Cloud API key"
  type        = string
}

// Data source arguments for scc_account_location
variable "scc_account_location_location_id" {
  description = "The programatic ID of the location that you want to work in."
  type        = string
  default     = "us"
}

// Resource arguments for ibm_scc_account_settings
variable "ibm_scc_account_settings_location_id" {
  description = "The programatic ID of the location that you want to work in."
  type        = string
}
