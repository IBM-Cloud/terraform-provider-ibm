variable "ibmcloud_api_key" {
  description = "IBM Cloud API key"
  type        = string
}

// Resource arguments for trusted_profile_template
variable "trusted_profile_template_name" {
  description = "The name of the trusted profile template. This is visible only in the enterprise account."
  type        = string
  default     = "name"
}
variable "trusted_profile_template_description" {
  description = "The description of the trusted profile template. Describe the template for enterprise account users."
  type        = string
  default     = null
}

// Data source arguments for trusted_profile_template
variable "trusted_profile_template_template_id" {
  description = "ID of the trusted profile template."
  type        = string
  default     = null
}
variable "trusted_profile_template_version" {
  description = "Version of the Profile Template."
  type        = string
  default     = null
}
variable "trusted_profile_template_include_history" {
  description = "Defines if the entity history is included in the response."
  type        = bool
  default     = false
}
