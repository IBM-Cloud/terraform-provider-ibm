// This allows tekton_pipeline_definition data to be referenced by other resources and the terraform CLI
// Modify this if only certain data should be exposed
output "ibm_cd_tekton_pipeline_definition" {
  value       = ibm_cd_tekton_pipeline_definition.tekton_pipeline_definition_instance
  description = "tekton_pipeline_definition resource instance"
}
// This allows tekton_pipeline_trigger_property data to be referenced by other resources and the terraform CLI
// Modify this if only certain data should be exposed
output "ibm_cd_tekton_pipeline_trigger_property" {
  value       = ibm_cd_tekton_pipeline_trigger_property.tekton_pipeline_trigger_property_instance
  description = "tekton_pipeline_trigger_property resource instance"
}
// This allows tekton_pipeline_property data to be referenced by other resources and the terraform CLI
// Modify this if only certain data should be exposed
output "ibm_cd_tekton_pipeline_property" {
  value       = ibm_cd_tekton_pipeline_property.tekton_pipeline_property_instance
  description = "tekton_pipeline_property resource instance"
}
// This allows tekton_pipeline_trigger data to be referenced by other resources and the terraform CLI
// Modify this if only certain data should be exposed
output "ibm_cd_tekton_pipeline_trigger" {
  value       = ibm_cd_tekton_pipeline_trigger.tekton_pipeline_trigger_instance
  description = "tekton_pipeline_trigger resource instance"
}
// This allows tekton_pipeline data to be referenced by other resources and the terraform CLI
// Modify this if only certain data should be exposed
output "ibm_cd_tekton_pipeline" {
  value       = ibm_cd_tekton_pipeline.tekton_pipeline_instance
  description = "tekton_pipeline resource instance"
}
