variable "ibmcloud_api_key" {
  description = "IBM Cloud API key"
  type        = string
}

// Resource arguments for iam_trusted_profile_identities
variable "iam_trusted_profile_identities_profile_id" {
  description = "ID of the trusted profile."
  type        = string
  default     = "profile_id"
}
variable "iam_trusted_profile_identities_if_match" {
  description = "Entity tag of the Identities to be updated. Specify the tag that you retrieved when reading the Profile Identities. This value helps identify parallel usage of this API. Pass * to indicate updating any available version, which may result in stale updates."
  type        = string
  default     = "if_match"
}
variable "iam_trusted_profile_identities" {
  description = "List of identities for the trusted profile."
  type = list(object({
    iam_id      = string
    type        = string
    identifier  = string
    accounts    = list(string)
    description = string
  }))
  default = [
    {
      iam_id      = "IBMid-5500082WK4"
      type        = "user"
      identifier  = "IBMid-5500082WK4"
      accounts    = ["86a1004d3f1848a291de32874cb48120"]
      description = "tf_description_profile identity description"
    }
  ]
}

// Data source arguments for iam_trusted_profile_identities
variable "data_iam_trusted_profile_identities_profile_id" {
  description = "ID of the trusted profile."
  type        = string
  default     = "profile_id"
}

