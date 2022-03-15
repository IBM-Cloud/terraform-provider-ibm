// This allows toolchain_tool_git data to be referenced by other resources and the terraform CLI
// Modify this if only certain data should be exposed
output "ibm_toolchain_tool_git" {
  value       = ibm_toolchain_tool_git.toolchain_tool_git_instance
  description = "toolchain_tool_git resource instance"
}
