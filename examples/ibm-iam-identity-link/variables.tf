variable "ibmcloud_api_key" {
  description = "IBM Cloud API key"
  type        = string
}

// Resource arguments for iam_trusted_profile_link
variable "iam_trusted_profile_link_profile_id" {
  description = "ID of the trusted profile."
  type        = string
  default     = "profile_id"
}
variable "iam_trusted_profile_link_cr_type" {
  description = "The compute resource type. Valid values are VSI, PVS, BMS, IKS_SA, ROKS_SA, CE."
  type        = string
  default     = "cr_type"
}
variable "iam_trusted_profile_link_name" {
  description = "Optional name of the Link."
  type        = string
  default     = "placeholder"
}

variable "iam_trusted_profile_link_crn" {
  description = "Link CRN of the resource."
  type        = string
  default     = "link_crn"
}
