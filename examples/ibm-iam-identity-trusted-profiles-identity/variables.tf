variable "ibmcloud_api_key" {
  description = "IBM Cloud API key"
  type        = string
}

// Resource arguments for iam_trusted_profile_identity
variable "iam_trusted_profile_identity_profile_id" {
  description = "ID of the trusted profile."
  type        = string
  default     = "profile_id"
}
variable "iam_trusted_profile_identity_identity_type" {
  description = "Type of the identity."
  type        = string
  default     = "user"
}
variable "iam_trusted_profile_identity_identifier" {
  description = "Identifier of the identity that can assume the trusted profiles. This can be a user identifier (IAM id), serviceid or crn. Internally it uses account id of the service id for the identifier 'serviceid' and for the identifier 'crn' it uses account id contained in the CRN."
  type        = string
  default     = "identifier"
}
variable "iam_trusted_profile_identity_type" {
  description = "Type of the identity."
  type        = string
  default     = "user"
}
variable "iam_trusted_profile_identity_accounts" {
  description = "Only valid for the type user. Accounts from which a user can assume the trusted profile."
  type        = list(string)
  default     = [ "accounts" ]
}
