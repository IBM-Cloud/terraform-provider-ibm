resource "ibm_cd_tekton_pipeline" "pr_pipeline_instance" {
  pipeline_id = var.pipeline_id
  worker {
    id = "public"
  }
}

resource "ibm_cd_tekton_pipeline_definition" "pr_pipeline_def" {
  pipeline_id   = ibm_cd_tekton_pipeline.pr_pipeline_instance.pipeline_id
  source {
    type = "git"
    properties {
      url         = var.pipeline_repo
      branch      = var.pipeline_repo_branch
      path        = var.pipeline_path
    }
  }
}

resource "ibm_cd_tekton_pipeline_definition" "pr_git_task_def" {
  pipeline_id   = ibm_cd_tekton_pipeline.pr_pipeline_instance.pipeline_id
  source {
    type = "git"
    properties {
      url         = var.tekton_tasks_catalog_repo
      branch      = var.definitions_branch
      path        = "git"
    }
  }
}

resource "ibm_cd_tekton_pipeline_definition" "pr_toolchain_task_def" {
  pipeline_id   = ibm_cd_tekton_pipeline.pr_pipeline_instance.pipeline_id
  source {
    type = "git"
    properties {
      url         = var.tekton_tasks_catalog_repo
      branch      = var.definitions_branch
      path        = "toolchain"
    }
  }
}

resource "ibm_cd_tekton_pipeline_definition" "pr_linter_task_def" {
  pipeline_id   = ibm_cd_tekton_pipeline.pr_pipeline_instance.pipeline_id
  source {
    type = "git"
    properties {
      url         = var.tekton_tasks_catalog_repo
      branch      = var.definitions_branch
      path        = "linter"
    }
  }
}

resource "ibm_cd_tekton_pipeline_definition" "pr_tester_task_def" {
  pipeline_id   = ibm_cd_tekton_pipeline.pr_pipeline_instance.pipeline_id
  source {
    type = "git"
    properties {
      url         = var.tekton_tasks_catalog_repo
      branch      = var.definitions_branch
      path        = "tester"
    }
  }
}

resource "ibm_cd_tekton_pipeline_definition" "pr_utils_task_def" {
  pipeline_id   = ibm_cd_tekton_pipeline.pr_pipeline_instance.pipeline_id
  source {
    type = "git"
    properties {
      url         = var.tekton_tasks_catalog_repo
      branch      = var.definitions_branch
      path        = "utils"
    }
  }
}

resource "ibm_cd_tekton_pipeline_trigger" "pr_pipeline_scm_trigger" {
  pipeline_id     = ibm_cd_tekton_pipeline.pr_pipeline_instance.pipeline_id
  type            = var.pr_pipeline_scm_trigger_type
  name            = var.pr_pipeline_scm_trigger_name
  event_listener  = var.pr_pipeline_scm_trigger_listener_name
  source {
    type = "git"
    properties {
      url       = var.app_repo
      branch    = var.app_repo_branch
    }
  }
  events              = ["pull_request"] 
  max_concurrent_runs = var.pr_pipeline_max_concurrent_runs
}