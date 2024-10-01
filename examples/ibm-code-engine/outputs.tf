// This allows code_engine_project data to be referenced by other resources and the terraform CLI
// Modify this if only certain data should be exposed
output "ibm_code_engine_project" {
  value       = ibm_code_engine_project.code_engine_project_instance
  description = "code_engine_project resource instance"
}

// This allows code_engine_app data to be referenced by other resources and the terraform CLI
// Modify this if only certain data should be exposed
output "ibm_code_engine_app" {
  value       = ibm_code_engine_app.code_engine_app_instance
  description = "code_engine_app resource instance"
}
// This allows code_engine_build data to be referenced by other resources and the terraform CLI
// Modify this if only certain data should be exposed
output "ibm_code_engine_build" {
  value       = ibm_code_engine_build.code_engine_build_instance
  description = "code_engine_build resource instance"
}
// This allows code_engine_config_map data to be referenced by other resources and the terraform CLI
// Modify this if only certain data should be exposed
output "ibm_code_engine_config_map" {
  value       = ibm_code_engine_config_map.code_engine_config_map_instance
  description = "code_engine_config_map resource instance"
}
// This allows code_engine_job data to be referenced by other resources and the terraform CLI
// Modify this if only certain data should be exposed
output "ibm_code_engine_job" {
  value       = ibm_code_engine_job.code_engine_job_instance
  description = "code_engine_job resource instance"
}
// This allows code_engine_secret data to be referenced by other resources and the terraform CLI
// Modify this if only certain data should be exposed
output "ibm_code_engine_secret" {
  value       = ibm_code_engine_secret.code_engine_secret_generic
  description = "code_engine_secret resource instance"
  sensitive   = true
}
// This allows code_engine_binding data to be referenced by other resources and the terraform CLI
// Modify this if only certain data should be exposed
output "ibm_code_engine_binding" {
  value       = ibm_code_engine_binding.code_engine_binding_instance
  description = "code_engine_binding resource instance"
}
// This allows code_engine_domain_mapping data to be referenced by other resources and the terraform CLI
// Modify this if only certain data should be exposed
output "ibm_code_engine_domain_mapping" {
  value       = ibm_code_engine_domain_mapping.code_engine_domain_mapping_instance
  description = "code_engine_domain_mapping resource instance"
}
// This allows code_engine_function data to be referenced by other resources and the terraform CLI
// Modify this if only certain data should be exposed
output "ibm_code_engine_function" {
  value       = ibm_code_engine_function.code_engine_function_instance
  description = "code_engine_function resource instance"
}
