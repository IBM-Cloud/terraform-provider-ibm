// This allows tekton_pipeline_property data to be referenced by other resources and the terraform CLI
// Modify this if only certain data should be exposed
output "ibm_tekton_pipeline_property" {
  value       = ibm_tekton_pipeline_property.tekton_pipeline_property_instance
  description = "tekton_pipeline_property resource instance"
}
