// This allows atracker_target data to be referenced by other resources and the terraform CLI
// Modify this if only certain data should be exposed
output "ibm_atracker_target" {
  value       = ibm_atracker_target.atracker_target_instance
  description = "atracker_target resource instance"
}
// This allows atracker_route data to be referenced by other resources and the terraform CLI
// Modify this if only certain data should be exposed
output "ibm_atracker_route" {
  value       = ibm_atracker_route.atracker_route_instance
  description = "atracker_route resource instance"
}

// This allows atracker_settings data to be referenced by other resources and the terraform CLI
// Modify this if only certain data should be exposed
output "ibm_atracker_settings" {
  value       = ibm_atracker_settings.atracker_settings_instance
  description = "atracker_settings resource instance"
}
