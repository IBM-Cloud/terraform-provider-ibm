output "available_locations" {
  value = data.ibm_scc_account_locations.scc_account_locations_instance.locations
}

output "location_details" {
  value = data.ibm_scc_account_location.scc_account_location_instance
}

output "current_location_settings_details" {
  value = data.ibm_scc_account_settings.scc_account_location_settings_instance
}
