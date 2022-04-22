provider "ibm" {
  ibmcloud_api_key = var.ibmcloud_api_key
}

// Provision toolchain_tool_hashicorpvault resource instance
resource "ibm_toolchain_tool_hashicorpvault" "toolchain_tool_hashicorpvault_instance" {
  toolchain_id = var.toolchain_tool_hashicorpvault_toolchain_id
  name = var.toolchain_tool_hashicorpvault_name
  parameters {
    name = "name"
    server_url = "server_url"
    authentication_method = "token"
    token = "token"
    role_id = "role_id"
    secret_id = "secret_id"
    dashboard_url = "dashboard_url"
    path = "path"
    secret_filter = "secret_filter"
    default_secret = "default_secret"
    username = "username"
    password = "password"
  }
  parameters_references = var.toolchain_tool_hashicorpvault_parameters_references
}
