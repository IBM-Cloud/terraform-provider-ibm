// This allows toolchain_tool_sonarqube data to be referenced by other resources and the terraform CLI
// Modify this if only certain data should be exposed
output "ibm_toolchain_tool_sonarqube" {
  value       = ibm_toolchain_tool_sonarqube.toolchain_tool_sonarqube_instance
  description = "toolchain_tool_sonarqube resource instance"
}
