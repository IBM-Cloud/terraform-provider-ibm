provider "ibm" {
  ibmcloud_api_key = var.ibmcloud_api_key
}

// Create list_scopes data source
data "ibm_scc_posture_scopes" "list_scopes_instance" {
  scope_id = var.list_scopes_scope_id
}

// Create list_profiles data source
data "ibm_scc_posture_profiles" "list_profiles_instance" {
  profile_id = var.list_profiles_profile_id
}

// Create list_latest_scans data source
data "ibm_scc_posture_latest_scans" "list_latest_scans_instance" {
  scan_id = var.list_latest_scans_scan_id
}

// Create scans_summary data source
data "ibm_scc_posture_scan_summary" "scans_summary_instance" {
  scan_id = var.scans_summary_scan_id
  profile_id = var.scans_summary_profile_id
}

// Create scan_summaries data source
data "ibm_scc_posture_scan_summaries" "scan_summaries_instance" {
  profile_id = var.scan_summaries_profile_id
  scope_id = var.scan_summaries_scope_id
  scan_id = var.scan_summaries_scan_id
}
