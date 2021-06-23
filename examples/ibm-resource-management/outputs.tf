// This allows resource_alias data to be referenced by other resources and the terraform CLI
// Modify this if only certain data should be exposed
output "ibm_resource_alias" {
  value       = ibm_resource_alias.resource_alias_instance
  description = "resource_alias resource instance"
}
// This allows resource_binding data to be referenced by other resources and the terraform CLI
// Modify this if only certain data should be exposed
output "ibm_resource_binding" {
  value       = ibm_resource_binding.resource_binding_instance
  description = "resource_binding resource instance"
}
