provider "ibm" {
  ibmcloud_api_key = var.ibmcloud_api_key
}

resource "ibm_iam_account_settings_template" "account_settings_template_instance" {
  name = var.account_settings_template_name
  account_settings {
    restrict_create_service_id = "RESTRICTED"
    restrict_create_platform_apikey = "RESTRICTED"
    allowed_ip_addresses = "allowed_ip_addresses"
    mfa = "NONE"
    user_mfa {
      iam_id = "iam_id"
      mfa = "NONE"
    }
    session_expiration_in_seconds = "session_expiration_in_seconds"
    session_invalidation_in_seconds = "session_invalidation_in_seconds"
    max_sessions_per_identity = "max_sessions_per_identity"
    system_access_token_expiration_in_seconds = "system_access_token_expiration_in_seconds"
    system_refresh_token_expiration_in_seconds = "system_refresh_token_expiration_in_seconds"
    restrict_user_list_visibility = "RESTRICTED"
    restrict_user_domains {
      account_sufficient = true
      restrictions {
        realm_id = "IBMid"
        invitation_email_allow_patterns = *.*@company.com
        restrict_invitation = true
      }
    }
  }
}

resource "ibm_iam_account_settings_template" "account_settings_template_new_version" {
  template_id = ibm_iam_account_settings_template.account_settings_template_instance.id
  name = var.account_settings_template_name
  description = "Description for version 2"
#  committed = true
  account_settings {
    mfa = "LEVEL3"
  }
}

// data source for a pre-existing template
#data "ibm_iam_account_settings_template" "account_settings_template_instance" {
#  template_id = var.account_settings_template_template_id
#  version = var.account_settings_template_version
#}