provider "ibm" {
  ibmcloud_api_key = var.ibmcloud_api_key
}

// Provision toolchain_tool_secretsmanager resource instance
resource "ibm_toolchain_tool_secretsmanager" "toolchain_tool_secretsmanager_instance" {
  toolchain_id = var.toolchain_tool_secretsmanager_toolchain_id
  name = var.toolchain_tool_secretsmanager_name
  parameters {
    name = "name"
    region = "region"
    resource_group = "resource_group"
    instance_name = "instance_name"
  }
  parameters_references = var.toolchain_tool_secretsmanager_parameters_references
}
