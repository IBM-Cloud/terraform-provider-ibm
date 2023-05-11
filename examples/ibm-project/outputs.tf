// This allows project_instance data to be referenced by other resources and the terraform CLI
// Modify this if only certain data should be exposed
output "ibm_project_instance" {
  value       = ibm_project_instance.project_instance
  description = "project_instance resource instance"
}
