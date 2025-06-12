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
variable "iam_trusted_profile_identity_description" {
  description = "Description of the identity that can assume the trusted profile. This is optional field for all the types of identities. When this field is not set for the identity type 'serviceid' then the description of the service id is used. Description is recommended for the identity type 'crn' E.g. 'Instance 1234 of IBM Cloud Service project'."
  type        = string
  default     = "description"
}
