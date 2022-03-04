provider "ibm" {
  ibmcloud_api_key = var.ibmcloud_api_key
}

// Provision toolchain_tool_sonarqube resource instance
resource "ibm_toolchain_tool_sonarqube" "toolchain_tool_sonarqube_instance" {
  toolchain_id = var.toolchain_tool_sonarqube_toolchain_id
  parameters {
    name = "name"
    dashboard_url = "dashboard_url"
    user_login = "user_login"
    user_password = "user_password"
    blind_connection = true
  }
  parameters_references = var.toolchain_tool_sonarqube_parameters_references
  container {
    guid = "d02d29f1-e7bb-4977-8a6f-26d7b7bb893e"
    type = "organization_guid"
  }
}
