// This allows toolchain data to be referenced by other resources and the terraform CLI
// Modify this if only certain data should be exposed
output "ibm_toolchain" {
  value       = ibm_toolchain.toolchain_instance
  description = "toolchain resource instance"
}
// This allows toolchain_integration data to be referenced by other resources and the terraform CLI
// Modify this if only certain data should be exposed
output "ibm_toolchain_integration" {
  value       = ibm_toolchain_integration.toolchain_integration_instance
  description = "toolchain_integration resource instance"
}
