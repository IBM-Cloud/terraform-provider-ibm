variable "ibmcloud_api_key" {
  description = "IBM Cloud API key"
  type        = string
}

// Data source arguments for iam_user_mfa_enrollments
variable "iam_user_mfa_enrollments_account_id" {
  description = "ID of the account."
  type        = string
  default     = "account_id"
}
variable "iam_user_mfa_enrollments_iam_id" {
  description = "iam_id of the user. This user must be the member of the account."
  type        = string
  default     = "iam_id"
}

// Data source arguments for iam_mfa_report
variable "iam_mfa_report_account_id" {
  description = "ID of the account."
  type        = string
  default     = "account_id"
}
variable "iam_mfa_report_reference" {
  description = "Reference for the report to be generated, You can use 'latest' to get the latest report for the given account."
  type        = string
  default     = "reference"
}
