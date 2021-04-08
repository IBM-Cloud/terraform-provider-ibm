provider "ibm" {
  ibmcloud_api_key = var.ibmcloud_api_key
}

// Read iam_account_settings data source
data "ibm_iam_account_settings" "iam_account_settings_source" {
}

// Provision iam_account_settings resource instance
resource "ibm_iam_account_settings" "iam_account_settings_instance" {
  mfa = "LEVEL3"
  restrict_create_service_id = "NOT_RESTRICTED"
}