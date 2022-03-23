// This allows toolchain_tool_secretsmanager data to be referenced by other resources and the terraform CLI
// Modify this if only certain data should be exposed
output "ibm_toolchain_tool_secretsmanager" {
  value       = ibm_toolchain_tool_secretsmanager.toolchain_tool_secretsmanager_instance
  description = "toolchain_tool_secretsmanager resource instance"
}
