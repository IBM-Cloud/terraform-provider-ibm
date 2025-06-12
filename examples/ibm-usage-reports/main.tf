provider "ibm" {
  ibmcloud_api_key = var.ibmcloud_api_key
}

// Provision billing_report_snapshot resource instance
resource "ibm_billing_report_snapshot" "billing_report_snapshot_instance" {
  interval = var.billing_report_snapshot_interval
  versioning = var.billing_report_snapshot_versioning
  report_types = var.billing_report_snapshot_report_types
  cos_reports_folder = var.billing_report_snapshot_cos_reports_folder
  cos_bucket = var.billing_report_snapshot_cos_bucket
  cos_location = var.billing_report_snapshot_cos_location
}

// Create billing_snapshot_list data source
data "ibm_billing_snapshot_list" "billing_snapshot_list_instance" {
  month = var.billing_snapshot_list_month
  date_from = var.billing_snapshot_list_date_from
  date_to = var.billing_snapshot_list_date_to
}
