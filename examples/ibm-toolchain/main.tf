provider "ibm" {
  ibmcloud_api_key = var.ibmcloud_api_key
}

// Provision toolchain_tool_security_compliance resource instance
resource "ibm_toolchain_tool_security_compliance" "toolchain_tool_security_compliance_instance" {
  toolchain_id = var.toolchain_tool_security_compliance_toolchain_id
  name = var.toolchain_tool_security_compliance_name
  parameters {
    name = "name"
    evidence_repo_name = "evidence_repo_name"
    trigger_scan = "disabled"
    location = "IBM Cloud"
    api-key = "api-key"
    scope = "scope"
    profile = "profile"
  }
  parameters_references = var.toolchain_tool_security_compliance_parameters_references
}
