// This allows project data to be referenced by other resources and the terraform CLI
// Modify this if only certain data should be exposed
output "ibm_project" {
  value       = ibm_project.project_instance
  description = "project resource instance"
}
