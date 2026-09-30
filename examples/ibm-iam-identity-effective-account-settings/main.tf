provider "ibm" {
  ibmcloud_api_key = var.ibmcloud_api_key
}

// Create iam_effective_account_settings data source
data "ibm_iam_effective_account_settings" "iam_effective_account_settings_instance_data" {
  account_id = var.iam_effective_account_settings_account_id
  include_history = false
  resolve_user_mfa = false
}
