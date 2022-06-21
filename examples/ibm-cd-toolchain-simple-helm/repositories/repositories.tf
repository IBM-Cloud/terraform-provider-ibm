resource "ibm_cd_toolchain_tool_hostedgit" "app_repo" {
  toolchain_id = var.toolchain_id
  name         = "app-repo"
  initialization {
    type = "clone"
    source_repo_url = var.app_repo
    private_repo = true
    repo_name = join("-", [split(".", split("/", var.app_repo)[4])[0], formatdate("DDMMYYYYhhmmss", timestamp())])
  }  
  parameters {
    has_issues          = true
    enable_traceability = true
  }
}

resource "ibm_cd_toolchain_tool_hostedgit" "pipeline_repo" {
  toolchain_id = var.toolchain_id
  name         = "pipeline-repo"  
  initialization {
    type = "clone"
    repo_url = var.pipeline_repo
    private_repo = true
    repo_name = join("-", [split(".", split("/", var.pipeline_repo)[4])[0], formatdate("DDMMYYYYhhmmss", timestamp())])
  }
  parameters {
    has_issues          = false
    enable_traceability = false
  }
}

resource "ibm_cd_toolchain_tool_hostedgit" "tekton_tasks_catalog_repo" {
  toolchain_id = var.toolchain_id
  name         = "tasks-repo"
  initialization {
    type = "clone"
    repo_url = var.tekton_tasks_catalog_repo
    private_repo = true
    repo_name = join("-", [split(".", split("/", var.tekton_tasks_catalog_repo)[4])[0], formatdate("DDMMYYYYhhmmss", timestamp())])
  }
  parameters {
    has_issues          = false
    enable_traceability = false
  }
}

output "app_repo_url" {
  value = ibm_cd_toolchain_tool_hostedgit.app_repo.parameters[0].repo_url
}

output "pipeline_repo_url" {
  value = ibm_cd_toolchain_tool_hostedgit.pipeline_repo.parameters[0].repo_url
}

output "tekton_tasks_catalog_repo_url" {
  value = ibm_cd_toolchain_tool_hostedgit.tekton_tasks_catalog_repo.parameters[0].repo_url
}