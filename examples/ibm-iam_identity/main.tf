provider "ibm" {
  ibmcloud_api_key = var.ibmcloud_api_key
}

// Provision iam_account_settings resource instance
resource "ibm_iam_account_settings" "iam_account_settings_instance" {
  include_history = var.iam_account_settings_include_history
}