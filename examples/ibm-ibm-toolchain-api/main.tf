provider "ibm" {
  ibmcloud_api_key = var.ibmcloud_api_key
}

// Provision toolchain_tool_git resource instance
resource "ibm_toolchain_tool_git" "toolchain_tool_git_instance" {
  toolchain_id = var.toolchain_tool_git_toolchain_id
  provider = var.toolchain_tool_git_provider
  parameters {
    repo_url = "repo_url"
    action = "clone"
    legal = true
    enable_traceability = true
    has_issues = true
  }
  parameters_references = var.toolchain_tool_git_parameters_references
  container {
    guid = "d02d29f1-e7bb-4977-8a6f-26d7b7bb893e"
    type = "organization_guid"
  }
}
