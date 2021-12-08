provider "ibm" {
  ibmcloud_api_key = var.ibmcloud_api_key
}

// Provision collectors resource instance
resource "ibm_scc_posture_v2_collectors" "collectors_instance" {
  name = var.collectors_name
  is_public = var.collectors_is_public
  managed_by = var.collectors_managed_by
  description = var.collectors_description
  passphrase = var.collectors_passphrase
  is_ubi_image = var.collectors_is_ubi_image
}

// Provision scopes resource instance
resource "ibm_scc_posture_v2_scopes" "scopes_instance" {
  name = var.scopes_name
  description = var.scopes_description
  collector_ids = var.scopes_collector_ids
  credential_id = var.scopes_credential_id
  credential_type = var.scopes_credential_type
}

// Provision credentials resource instance
resource "ibm_scc_posture_v2_credentials" "credentials_instance" {
  enabled = var.credentials_enabled
  type = var.credentials_type
  name = var.credentials_name
  description = var.credentials_description
  display_fields {
    ibm_api_key = "sample_api_key"
  }
  group {
    id = "1"
    passphrase = "passphrase"
  }
  purpose = var.credentials_purpose
}

// Create list_scopes data source
//data "ibm_scc_posture_v2_list_scopes" "list_scopes_instance" {
//}

// Create profileDetails data source
data "ibm_scc_posture_v2_profileDetails" "profileDetails_instance" {
  id = var.profileDetails_id
  profile_type = "4"
}

// Create list_profiles data source
data "ibm_scc_posture_v2_list_profiles" "list_profiles_instance" {
}

// Create list_latest_scans data source
data "ibm_scc_posture_v2_list_latest_scans" "list_latest_scans_instance" {
}

// Create scans_summary data source
data "ibm_scc_posture_v2_scans_summary" "scans_summary_instance" {
  scan_id = var.scans_summary_scan_id
  profile_id = var.scans_summary_profile_id
}

// Create scan_summaries data source
data "ibm_scc_posture_v2_scan_summaries" "scan_summaries_instance" {
  report_setting_id = var.scan_summaries_report_setting_id
}

// Create group_profile_details data source
data "ibm_scc_posture_v2_group_profile_details" "group_profile_details_instance" {
  profile_id = var.group_profile_details_profile_id
}

// Create scope_correlation data source
data "ibm_scc_posture_v2_scope_correlation" "scope_correlation_instance" {
  correlation_id = var.scope_correlation_correlation_id
}
