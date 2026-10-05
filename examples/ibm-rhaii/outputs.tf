output "project_id" {
  description = "ID of the project. Use it as project_id in the Red Hat AI Inference API."
  value       = ibm_rhaii_project.project.project_id
}

output "endpoint" {
  description = "Base URL of the Red Hat AI Inference API for this project"
  value       = ibm_rhaii_project.project.endpoint
}

output "crn" {
  description = "CRN of the project"
  value       = ibm_rhaii_project.project.crn
}
