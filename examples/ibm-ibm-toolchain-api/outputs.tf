// This allows toolchain_tool_insights data to be referenced by other resources and the terraform CLI
// Modify this if only certain data should be exposed
output "ibm_toolchain_tool_insights" {
  value       = ibm_toolchain_tool_insights.toolchain_tool_insights_instance
  description = "toolchain_tool_insights resource instance"
}
