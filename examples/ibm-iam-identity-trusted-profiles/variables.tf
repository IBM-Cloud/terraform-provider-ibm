variable "ibmcloud_api_key" {
  description = "IBM Cloud API key"
  type        = string
}

// Resource arguments for iam_trusted_profiles
variable "iam_trusted_profiles_name" {
  description = "Name of the trusted profile. The name is checked for uniqueness. Therefore trusted profiles with the same names can not exist in the same account."
  type        = string
  default     = "name"
}
variable "iam_trusted_profiles_account_id" {
  description = "The account ID of the trusted profile."
  type        = string
  default     = "account_id"
}
variable "iam_trusted_profiles_description" {
  description = "The optional description of the trusted profile. The 'description' property is only available if a description was provided during creation of trusted profile."
  type        = string
  default     = "placeholder"
}

// Data source arguments for iam_trusted_profiles
variable "iam_trusted_profiles_profile_id" {
  description = "ID of the trusted profile to get."
  type        = string
  default     = "profile_id"
}
