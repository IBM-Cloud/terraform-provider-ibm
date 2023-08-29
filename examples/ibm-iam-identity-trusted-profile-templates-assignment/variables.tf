variable "ibmcloud_api_key" {
  description = "IBM Cloud API key"
  type        = string
}

// Resource arguments for trusted_profile_template_assignment
variable "trusted_profile_template_assignment_template_id" {
  description = "Template Id."
  type        = string
  default     = "template_id"
}
variable "trusted_profile_template_assignment_template_version" {
  description = "Template version."
  type        = number
  default     = 1
}
variable "trusted_profile_template_assignment_target_type" {
  description = "Assignment target type."
  type        = string
  default     = "Account"
}
variable "trusted_profile_template_assignment_target" {
  description = "Assignment target."
  type        = string
  default     = "target"
}

// Data source arguments for trusted_profile_template_assignment
variable "trusted_profile_template_assignment_assignment_id" {
  description = "ID of the Assignment Record."
  type        = string
  default     = "assignment_id"
}
variable "trusted_profile_template_assignment_include_history" {
  description = "Defines if the entity history is included in the response."
  type        = bool
  default     = false
}
