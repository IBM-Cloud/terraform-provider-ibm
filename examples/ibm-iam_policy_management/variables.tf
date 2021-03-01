variable "ibmcloud_api_key" {
  description = "IBM Cloud API key"
  type        = string
}

// Resource arguments for iam_policy
variable "iam_policy_type" {
  description = "The policy type; either 'access' or 'authorization'."
  type        = string
  default     = "type"
}
variable "iam_policy_subjects" {
  description = "The subjects associated with a policy."
  type        = list(object({ example=string }))
  default     = [ { example: "object" } ]
}
variable "iam_policy_roles" {
  description = "A set of role cloud resource names (CRNs) granted by the policy."
  type        = list(object({ example=string }))
  default     = [ { example: "object" } ]
}
variable "iam_policy_resources" {
  description = "The resources associated with a policy."
  type        = list(object({ example=string }))
  default     = [ { example: "object" } ]
}
variable "iam_policy_description" {
  description = "Customer-defined description."
  type        = string
  default     = "placeholder"
}
variable "iam_policy_accept_language" {
  description = "Translation language code."
  type        = string
  default     = "placeholder"
}

// Resource arguments for iam_custom_role
variable "iam_custom_role_display_name" {
  description = "The display name of the role that is shown in the console."
  type        = string
  default     = "display_name"
}
variable "iam_custom_role_actions" {
  description = "The actions of the role."
  type        = list(string)
  default     = [ "actions" ]
}
variable "iam_custom_role_name" {
  description = "The name of the role that is used in the CRN. Can only be alphanumeric and has to be capitalized."
  type        = string
  default     = "name"
}
variable "iam_custom_role_account_id" {
  description = "The account GUID."
  type        = string
  default     = "account_id"
}
variable "iam_custom_role_service_name" {
  description = "The service name."
  type        = string
  default     = "service_name"
}
variable "iam_custom_role_description" {
  description = "The description of the role."
  type        = string
  default     = "placeholder"
}
variable "iam_custom_role_accept_language" {
  description = "Translation language code."
  type        = string
  default     = "placeholder"
}

// Data source arguments for iam_policy
variable "iam_policy_policy_id" {
  description = "The policy ID."
  type        = string
  default     = "policy_id"
}

// Data source arguments for iam_custom_role
variable "iam_custom_role_role_id" {
  description = "The role ID."
  type        = string
  default     = "role_id"
}
