output "project_id" {
  description = "ID of the project. Use it as project_id in the Red Hat AI Inference API."
  value       = ibm_rhaii_project.project.project_id
}

output "endpoint" {
  description = "Base URL of the Red Hat AI Inference API for this project"
  value       = ibm_rhaii_project.project.endpoint
}

output "private_endpoint" {
  description = "Private base URL of the Red Hat AI Inference API for this project, reachable from the IBM Cloud private network"
  value       = ibm_rhaii_project.project.private_endpoint
}

output "plan_id" {
  description = "Global catalog ID of the pricing plan of the project"
  value       = ibm_rhaii_project.project.plan_id
}

output "crn" {
  description = "CRN of the project"
  value       = ibm_rhaii_project.project.crn
}

output "model_ids" {
  description = "IDs of the inference models available in the project"
  value       = data.ibm_rhaii_inference_models.models.models[*].id
}

output "model_status" {
  description = "Status of the selected inference model"
  value       = data.ibm_rhaii_inference_model.model.status
}
