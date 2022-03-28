provider "ibm" {
  ibmcloud_api_key = var.ibmcloud_api_key
}

// Provision toolchain_tool_sonarqube resource instance
resource "ibm_toolchain_tool_sonarqube" "toolchain_tool_sonarqube_instance" {
  toolchain_id = var.toolchain_tool_sonarqube_toolchain_id
  name = var.toolchain_tool_sonarqube_name
  parameters {
    name = "name"
    dashboard_url = "dashboard_url"
    user_login = "user_login"
    user_password = "user_password"
    blind_connection = true
  }
}
