// This allows toolchain_tool_hashicorpvault data to be referenced by other resources and the terraform CLI
// Modify this if only certain data should be exposed
output "ibm_toolchain_tool_hashicorpvault" {
  value       = ibm_toolchain_tool_hashicorpvault.toolchain_tool_hashicorpvault_instance
  description = "toolchain_tool_hashicorpvault resource instance"
}
