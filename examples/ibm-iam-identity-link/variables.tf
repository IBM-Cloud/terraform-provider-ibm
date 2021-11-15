variable "ibmcloud_api_key" {
  description = "IBM Cloud API key"
  type        = string
}

// Resource arguments for iam_trusted_profiles_link
variable "iam_trusted_profiles_link_profile_id" {
  description = "ID of the trusted profile."
  type        = string
  default     = "profile_id"
}
variable "iam_trusted_profiles_link_cr_type" {
  description = "The compute resource type. Valid values are VSI, IKS_SA, ROKS_SA."
  type        = string
  default     = "cr_type"
}
variable "iam_trusted_profiles_link_name" {
  description = "Optional name of the Link."
  type        = string
  default     = "placeholder"
}

// Data source arguments for iam_trusted_profiles_link
variable "data_source_iam_trusted_profiles_link_profile_id" {
  description = "ID of the trusted profile."
  type        = string
  default     = "profile_id"
}
variable "iam_trusted_profiles_link_link_id" {
  description = "ID of the link."
  type        = string
  default     = "link_id"
}

// Data source arguments for iam_trusted_profile_link_list
variable "iam_trusted_profile_link_list_profile_id" {
  description = "ID of the trusted profile."
  type        = string
  default     = "profile_id"
}
