variable "ibmcloud_api_key" {
  description = "IBM Cloud API key"
  type        = string
}

// Resource arguments for iam_trusted_profiles_claim_rule
variable "iam_trusted_profile_claim_rule_profile_id" {
  description = "ID of the trusted profile to create a claim rule."
  type        = string
  default     = "profile_id"
}
variable "iam_trusted_profile_claim_rule_type" {
  description = "Type of the calim rule, either 'Profile-SAML' or 'Profile-CR'."
  type        = string
  default     = "type"
}
variable "iam_trusted_profile_claim_rule_conditions" {
  description = "Conditions of this claim rule."
  type        = list(object({ claim=string }))
  default     = [ { "claim" : "claim", "operator" : "operator", "value" : "value" } ]
}
variable "iam_trusted_profile_claim_rule_name" {
  description = "Name of the claim rule to be created or updated."
  type        = string
  default     = "placeholder"
}
variable "iam_trusted_profile_claim_rule_realm_name" {
  description = "The realm name of the Idp this claim rule applies to. This field is required only if the type is specified as 'Profile-SAML'."
  type        = string
  default     = "placeholder"
}
variable "iam_trusted_profile_claim_rule_cr_type" {
  description = "The compute resource type the rule applies to, required only if type is specified as 'Profile-CR'. Valid values are VSI, IKS_SA, ROKS_SA."
  type        = string
  default     = "placeholder"
}
variable "iam_trusted_profile_claim_rule_expiration" {
  description = "Session expiration in seconds, only required if type is 'Profile-SAML'."
  type        = number
  default     = 0
}

// Data source arguments for iam_trusted_profiles_claim_rule
variable "data_source_iam_trusted_profile_claim_rule_profile_id" {
  description = "ID of the trusted profile."
  type        = string
  default     = "profile_id"
}
variable "iam_trusted_profile_claim_rule_rule_id" {
  description = "ID of the claim rule to get."
  type        = string
  default     = "rule_id"
}

// Data source arguments for iam_trusted_profile_claim_rules
variable "iam_trusted_profile_claim_rule_profile_id" {
  description = "ID of the trusted profile."
  type        = string
  default     = "profile_id"
}