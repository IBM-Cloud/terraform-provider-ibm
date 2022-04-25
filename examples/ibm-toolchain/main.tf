provider "ibm" {
  ibmcloud_api_key = var.ibmcloud_api_key
}

// Provision toolchain_tool_private_worker resource instance
resource "ibm_toolchain_tool_private_worker" "toolchain_tool_private_worker_instance" {
  toolchain_id = var.toolchain_tool_private_worker_toolchain_id
  name = var.toolchain_tool_private_worker_name
  parameters {
    name = "name"
    workerQueueCredentials = "workerQueueCredentials"
    workerQueueIdentifier = "workerQueueIdentifier"
  }
  parameters_references = var.toolchain_tool_private_worker_parameters_references
}
