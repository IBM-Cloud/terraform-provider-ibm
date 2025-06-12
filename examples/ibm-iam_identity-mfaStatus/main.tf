provider "ibm" {
  ibmcloud_api_key = var.ibmcloud_api_key
}

// Data source is not linked to a resource instance
// Uncomment if an existing data source instance exists

// Create iam_user_mfa_enrollments data source
data "ibm_iam_user_mfa_enrollments" "iam_user_mfa_enrollments_instance" {
  account_id = var.iam_user_mfa_enrollments_account_id
  iam_id = var.iam_user_mfa_enrollments_iam_id
}


// Data source is not linked to a resource instance
// Uncomment if an existing data source instance exists
/*
// Create iam_mfa_report data source
data "ibm_iam_mfa_report" "iam_mfa_report_instance" {
  account_id = var.iam_mfa_report_account_id
  reference = var.iam_mfa_report_reference
}
*/

