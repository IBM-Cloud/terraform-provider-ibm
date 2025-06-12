variable "ibmcloud_api_key" {
  description = "IBM Cloud API key"
  type        = string
}

// Resource arguments for billing_report_snapshot
variable "billing_report_snapshot_interval" {
  description = "Frequency of taking the snapshot of the billing reports."
  type        = string
  default     = "daily"
}
variable "billing_report_snapshot_versioning" {
  description = "A new version of report is created or the existing report version is overwritten with every update."
  type        = string
  default     = "new"
}
variable "billing_report_snapshot_report_types" {
  description = "The type of billing reports to take snapshot of. Possible values are [account_summary, enterprise_summary, account_resource_instance_usage]."
  type        = list(string)
  default     = ["account_summary","enterprise_summary","account_resource_instance_usage"]
}
variable "billing_report_snapshot_cos_reports_folder" {
  description = "The billing reports root folder to store the billing reports snapshots. Defaults to \"IBMCloud-Billing-Reports\"."
  type        = string
  default     = "IBMCloud-Billing-Reports"
}
variable "billing_report_snapshot_cos_bucket" {
  description = "The name of the COS bucket to store the snapshot of the billing reports."
  type        = string
  default     = "bucket_name"
}
variable "billing_report_snapshot_cos_location" {
  description = "Region of the COS instance."
  type        = string
  default     = "us-south"
}

// Data source arguments for billing_snapshot_list
variable "billing_snapshot_list_account_id" {
  description = "Account ID for which the billing report snapshot is requested."
  type        = string
  default     = "abc"
}
variable "billing_snapshot_list_month" {
  description = "The month for which billing report snapshot is requested.  Format is yyyy-mm."
  type        = string
  default     = "2023-02"
}
variable "billing_snapshot_list_date_from" {
  description = "Timestamp in milliseconds for which billing report snapshot is requested."
  type        = number
  default     = 1675209600000
}
variable "billing_snapshot_list_date_to" {
  description = "Timestamp in milliseconds for which billing report snapshot is requested."
  type        = number
  default     = 1675987200000
}
variable "billing_snapshot_list_limit" {
  description = "Number of usage records returned. The default value is 30. Maximum value is 200."
  type        = number
  default     = 0
}
