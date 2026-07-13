provider "ibm" {
  ibmcloud_api_key = var.ibmcloud_api_key
}

// Create iam_account_settings instance
resource "ibm_iam_account_settings" "iam_account_settings_instance" {
  include_history = true
  restrict_user_domains {
    realm_id                        = "IBMid"
    restrict_invitation             = false
    invitation_email_allow_patterns = ["*@ibm.com"]
  }
}

resource "ibm_iam_account_settings" "iam_account_settings_additional" {
  restrict_create_service_id = "NOT_SET"
  restrict_create_platform_apikey = "NOT_SET"
  restrict_user_list_visibility = "NOT_RESTRICTED"
  if_match = "*"
  mfa = "NONE"
  user_mfa {
    iam_id = var.iam_account_settings_ibmid1
    mfa = "NONE"
  }
  restrict_user_domains {
    realm_id = "IBMid"
    invitation_email_allow_patterns = [
      "*@ibm.com",
      "*@corp.org"
    ]
    restrict_invitation = false
  }
  session_expiration_in_seconds = "NOT_SET"
  session_invalidation_in_seconds = "NOT_SET"
  max_sessions_per_identity = "NOT_SET"
  system_access_token_expiration_in_seconds = "3600"
  system_refresh_token_expiration_in_seconds = "259200"
}

// Create iam_account_settings data source
data "ibm_iam_account_settings" "iam_account_settings_data" {
  include_history = false
  resolve_user_mfa = false
}
