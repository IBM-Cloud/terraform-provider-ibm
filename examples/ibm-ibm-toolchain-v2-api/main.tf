provider "ibm" {
  ibmcloud_api_key = var.ibmcloud_api_key
}

// Provision toolchain resource instance
resource "ibm_toolchain" "toolchain_instance" {
  post_toolchain_request {
    name = "TestToolchainV2"
    description = "A sample toolchain to test the API"
    resource_group_id = "6a9a01f2cff54a7f966f803d92877123"
    generator = "API"
    template = { "key": null }
  }
}

// Provision toolchain_integration resource instance
resource "ibm_toolchain_integration" "toolchain_integration_instance" {
  toolchain_id = var.toolchain_integration_toolchain_id
  service_id = var.toolchain_integration_service_id
  name = var.toolchain_integration_name
  parameters = var.toolchain_integration_parameters
}
