// This allows collectors data to be referenced by other resources and the terraform CLI
// Modify this if only certain data should be exposed
output "ibm_scc_posture_v2_collectors" {
  value       = ibm_scc_posture_v2_collectors.collectors_instance
  description = "collectors resource instance"
}
// This allows scopes data to be referenced by other resources and the terraform CLI
// Modify this if only certain data should be exposed
output "ibm_scc_posture_v2_scopes" {
  value       = ibm_scc_posture_v2_scopes.scopes_instance
  description = "scopes resource instance"
}
// This allows credentials data to be referenced by other resources and the terraform CLI
// Modify this if only certain data should be exposed
output "ibm_scc_posture_v2_credentials" {
  value       = ibm_scc_posture_v2_credentials.credentials_instance
  description = "credentials resource instance"
}
