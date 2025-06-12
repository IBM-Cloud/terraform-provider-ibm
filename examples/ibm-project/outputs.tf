// This output allows project_config data to be referenced by other resources and the terraform CLI
// Modify this output if only certain data should be exposed
output "ibm_project_config" {
  value       = ibm_project_config.project_config_instance
  description = "project_config resource instance"
  sensitive   = true
}
// This output allows project data to be referenced by other resources and the terraform CLI
// Modify this output if only certain data should be exposed
output "ibm_project" {
  value       = ibm_project.project_instance
  description = "project resource instance"
}
// This output allows project_environment data to be referenced by other resources and the terraform CLI
// Modify this output if only certain data should be exposed
output "ibm_project_environment" {
  value       = ibm_project_environment.project_environment_instance
  description = "project_environment resource instance"
  sensitive   = true
}
