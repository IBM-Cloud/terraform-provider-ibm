provider "ibm" {
  ibmcloud_api_key = var.ibmcloud_api_key
}

// Data source is not linked to a resource instance
// Uncomment if an existing data source instance exists
/*
// Create iam_effective_account_settings data source
data "ibm_iam_effective_account_settings" "iam_effective_account_settings_instance" {
  account_id = var.iam_effective_account_settings_account_id
  include_history = var.iam_effective_account_settings_include_history
  resolve_user_mfa = var.iam_effective_account_settings_resolve_user_mfa
}
*/
