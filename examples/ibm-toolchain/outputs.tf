// This allows toolchain_tool_security_compliance data to be referenced by other resources and the terraform CLI
// Modify this if only certain data should be exposed
output "ibm_toolchain_tool_security_compliance" {
  value       = ibm_toolchain_tool_security_compliance.toolchain_tool_security_compliance_instance
  description = "toolchain_tool_security_compliance resource instance"
}
