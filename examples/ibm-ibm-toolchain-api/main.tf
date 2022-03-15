provider "ibm" {
  ibmcloud_api_key = var.ibmcloud_api_key
}

// Provision toolchain_tool_git resource instance
resource "ibm_toolchain_tool_git" "toolchain_tool_git_instance" {
  git_provider = var.toolchain_tool_git_git_provider
  toolchain_id = var.toolchain_tool_git_toolchain_id
  initialization {
    repo_name = "repo_name"
    repo_url = "repo_url"
    source_repo_url = "source_repo_url"
    type = "new"
    private_repo = true
  }
  parameters {
    enable_traceability = true
    has_issues = true
    repo_name = "repo_name"
    repo_url = "repo_url"
    source_repo_url = "source_repo_url"
    type = "new"
    private_repo = true
  }
  container {
    guid = "d02d29f1-e7bb-4977-8a6f-26d7b7bb893e"
    type = "organization_guid"
  }
}
