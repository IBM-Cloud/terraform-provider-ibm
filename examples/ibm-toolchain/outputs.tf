// This allows toolchain_tool_private_worker data to be referenced by other resources and the terraform CLI
// Modify this if only certain data should be exposed
output "ibm_toolchain_tool_private_worker" {
  value       = ibm_toolchain_tool_private_worker.toolchain_tool_private_worker_instance
  description = "toolchain_tool_private_worker resource instance"
}
