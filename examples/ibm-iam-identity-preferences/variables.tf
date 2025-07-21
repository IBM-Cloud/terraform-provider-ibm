variable "ibmcloud_api_key" {
  description = "IBM Cloud API key"
  type        = string
}

// Resource arguments for iam_identity_preference
variable "iam_identity_preference_account_id" {
  description = "Account id to update preference for."
  type        = string
  default     = "account_id"
}
variable "iam_identity_preference_iam_id" {
  description = "IAM id to update the preference for."
  type        = string
  default     = "iam_id"
}
variable "iam_identity_preference_service" {
  description = "Service of the preference to be updated."
  type        = string
  default     = "service"
}
variable "iam_identity_preference_preference_id" {
  description = "Identifier of preference to be updated."
  type        = string
  default     = "preference_id"
}
variable "iam_identity_preference_value_string" {
  description = "String value of the preference, only one value property is set, either 'value_string' or 'value_list_of_strings' is present."
  type        = string
  default     = "value_string"
}
variable "iam_identity_preference_value_list_of_strings" {
  description = "List of value of the preference, only one value property is set, either 'value_string' or 'value_list_of_strings' is present."
  type        = list(string)
  default     = [ "value_list_of_strings" ]
}

// Data source arguments for iam_identity_preference
variable "data_iam_identity_preference_account_id" {
  description = "Account id to get preference for."
  type        = string
  default     = "account_id"
}
variable "data_iam_identity_preference_iam_id" {
  description = "IAM id to get the preference for."
  type        = string
  default     = "iam_id"
}
variable "data_iam_identity_preference_service" {
  description = "Service of the preference to be fetched."
  type        = string
  default     = "service"
}
variable "data_iam_identity_preference_preference_id" {
  description = "Identifier of preference to be fetched."
  type        = string
  default     = "preference_id"
}
