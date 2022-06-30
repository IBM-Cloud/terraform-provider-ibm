// This allows cd_toolchain_tool_sonarqube data to be referenced by other resources and the terraform CLI
// Modify this if only certain data should be exposed
output "ibm_cd_toolchain_tool_sonarqube" {
  value       = ibm_cd_toolchain_tool_sonarqube.cd_toolchain_tool_sonarqube_instance
  description = "cd_toolchain_tool_sonarqube resource instance"
}
