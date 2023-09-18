// This output allows scc_provider_type_instance data to be referenced by other resources and the terraform CLI
// Modify this output if only certain data should be exposed
output "ibm_scc_provider_type_instance" {
  value       = ibm_scc_provider_type_instance.scc_provider_type_instance_instance
  description = "scc_provider_type_instance resource instance"
}
