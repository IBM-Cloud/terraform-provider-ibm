provider "ibm" {
  ibmcloud_api_key = var.ibmcloud_api_key
}

// Provision iam_account_settings resource instance
resource "ibm_iam_account_settings" "iam_account_settings_instance" {
  mfa = "NONE"
  restrict_create_service_id = "NOT_RESTRICTED"
  user_mfa {
    iam_id = "IBMid-123456789"
    mfa = "NONE"
  }
}