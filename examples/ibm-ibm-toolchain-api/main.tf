provider "ibm" {
  ibmcloud_api_key = var.ibmcloud_api_key
}

// Provision toolchain_tool_insights resource instance
resource "ibm_toolchain_tool_insights" "toolchain_tool_insights_instance" {
  toolchain_id = var.toolchain_tool_insights_toolchain_id
  parameters_references = var.toolchain_tool_insights_parameters_references
  container {
    guid = "d02d29f1-e7bb-4977-8a6f-26d7b7bb893e"
    type = "organization_guid"
  }
}
