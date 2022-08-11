// This allows schematics_workspace data to be referenced by other resources and the terraform CLI
// Modify this if only certain data should be exposed
output "ibm_schematics_workspace" {
  value       = ibm_schematics_workspace.schematics_workspace_instance
  description = "schematics_workspace resource instance"
}
// This allows schematics_action data to be referenced by other resources and the terraform CLI
// Modify this if only certain data should be exposed
output "ibm_schematics_action" {
  value       = ibm_schematics_action.schematics_action_instance
  description = "schematics_action resource instance"
}
// This allows schematics_job data to be referenced by other resources and the terraform CLI
// Modify this if only certain data should be exposed
output "ibm_schematics_job" {
  value       = ibm_schematics_job.schematics_job_instance
  description = "schematics_job resource instance"
}
