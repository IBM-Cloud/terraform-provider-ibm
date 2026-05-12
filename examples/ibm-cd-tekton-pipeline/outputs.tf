// This allows cd_tekton_pipeline data to be referenced by other resources and the terraform CLI
// Modify this if only certain data should be exposed"
output "ibm_cd_tekton_pipeline" {
  value       = ibm_cd_tekton_pipeline.cd_pipeline_instance
  description = "cd_pipeline_instance resource instance"
}
// This allows cd_tekton_pipeline_definition data to be referenced by other resources and the terraform CLI
// Modify this if only certain data should be exposed
output "ibm_cd_tekton_pipeline_definition" {
  value       = ibm_cd_tekton_pipeline_definition.cd_tekton_pipeline_definition_instance
  description = "cd_tekton_pipeline_definition resource instance"
}
// This allows cd_tekton_pipeline_trigger_property data to be referenced by other resources and the terraform CLI
// Modify this if only certain data should be exposed
output "ibm_cd_tekton_pipeline_trigger_property" {
  value       = ibm_cd_tekton_pipeline_trigger_property.cd_tekton_pipeline_trigger_property_instance
  description = "cd_tekton_pipeline_trigger_property resource instance"
}
// This allows cd_tekton_pipeline_property data to be referenced by other resources and the terraform CLI
// Modify this if only certain data should be exposed
output "ibm_cd_tekton_pipeline_property" {
  value       = ibm_cd_tekton_pipeline_property.cd_tekton_pipeline_property_instance
  description = "cd_tekton_pipeline_property resource instance"
}
// This allows cd_tekton_pipeline_trigger data to be referenced by other resources and the terraform CLI
// Modify this if only certain data should be exposed
output "ibm_cd_tekton_pipeline_trigger" {
  value       = ibm_cd_tekton_pipeline_trigger.cd_tekton_pipeline_trigger_instance
  description = "cd_tekton_pipeline_trigger resource instance"
}