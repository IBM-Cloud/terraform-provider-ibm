// This allows toolchain_tool_pipeline data to be referenced by other resources and the terraform CLI
// Modify this if only certain data should be exposed
output "ibm_toolchain_tool_pipeline" {
  value       = ibm_toolchain_tool_pipeline.toolchain_tool_pipeline_instance
  description = "toolchain_tool_pipeline resource instance"
}
