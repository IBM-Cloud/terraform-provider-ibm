provider "ibm" {
  ibmcloud_api_key = var.ibmcloud_api_key
}

// Provision cd_toolchain_tool_sonarqube resource instance
resource "ibm_cd_toolchain_tool_sonarqube" "cd_toolchain_tool_sonarqube_instance" {
  toolchain_id = var.cd_toolchain_tool_sonarqube_toolchain_id
  name = var.cd_toolchain_tool_sonarqube_name
  parameters {
    name = "name"
    dashboard_url = "dashboard_url"
    user_login = "user_login"
    user_password = "user_password"
    blind_connection = true
  }
}

// Create cd_toolchain_tool_sonarqube data source
data "ibm_cd_toolchain_tool_sonarqube" "cd_toolchain_tool_sonarqube_instance" {
  toolchain_id = var.cd_toolchain_tool_sonarqube_toolchain_id
  tool_id = var.cd_toolchain_tool_sonarqube_tool_id
}
