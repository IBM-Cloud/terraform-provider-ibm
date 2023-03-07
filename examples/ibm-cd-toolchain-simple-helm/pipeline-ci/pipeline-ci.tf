resource "ibm_cd_tekton_pipeline" "ci_pipeline_instance" {
  pipeline_id = var.pipeline_id
  worker {
    id = "public"
  }
}

resource "ibm_cd_tekton_pipeline_definition" "ci_pipeline_def" {
  pipeline_id  = ibm_cd_tekton_pipeline.ci_pipeline_instance.pipeline_id
  source {
    type = "git"
    properties {
      url = var.pipeline_repo
      branch = var.pipeline_repo_branch
      path = ".pipeline"
    }
  }
}

resource "ibm_cd_tekton_pipeline_definition" "ci_git_task_def" {
  pipeline_id  = ibm_cd_tekton_pipeline.ci_pipeline_instance.pipeline_id
  source {
    type = "git"
    properties {
      url = var.tekton_tasks_catalog_repo
      branch = var.definitions_branch
      path = "git"
    }
  }
}

resource "ibm_cd_tekton_pipeline_definition" "ci_cr_task_def" {
  pipeline_id  = ibm_cd_tekton_pipeline.ci_pipeline_instance.pipeline_id
  source {
    type = "git"
    properties {
      url = var.tekton_tasks_catalog_repo
      branch = var.definitions_branch
      path = "container-registry"
    }
  }
}

resource "ibm_cd_tekton_pipeline_definition" "ci_kube_task_def" {
  pipeline_id  = ibm_cd_tekton_pipeline.ci_pipeline_instance.pipeline_id
  source {
    type = "git"
    properties {
      url = var.tekton_tasks_catalog_repo
      branch = var.definitions_branch
      path = "kubernetes-service"
    }
  }
}

resource "ibm_cd_tekton_pipeline_definition" "ci_toolchain_task_def" {
  pipeline_id  = ibm_cd_tekton_pipeline.ci_pipeline_instance.pipeline_id
  source {
    type = "git"
    properties {
      url = var.tekton_tasks_catalog_repo
      branch = var.definitions_branch
      path = "toolchain"
    }
  }
}

resource "ibm_cd_tekton_pipeline_definition" "ci_insights_task_def" {
  pipeline_id  = ibm_cd_tekton_pipeline.ci_pipeline_instance.pipeline_id
  source {
    type = "git"
    properties {
      url = var.tekton_tasks_catalog_repo
      branch = var.definitions_branch
      path = "devops-insights"
    }
  }
}

resource "ibm_cd_tekton_pipeline_definition" "ci_linter_task_def" {
  pipeline_id  = ibm_cd_tekton_pipeline.ci_pipeline_instance.pipeline_id
  source {
    type = "git"
    properties {
      url = var.tekton_tasks_catalog_repo
      branch = var.definitions_branch
      path = "linter"
    }
  }
}

resource "ibm_cd_tekton_pipeline_definition" "ci_tester_task_def" {
  pipeline_id   = ibm_cd_tekton_pipeline.ci_pipeline_instance.pipeline_id
  source {
    type = "git"
    properties {
      url = var.tekton_tasks_catalog_repo
      branch = var.definitions_branch
      path = "tester"
    }
  }
}

resource "ibm_cd_tekton_pipeline_definition" "ci_utils_task_def" {
  pipeline_id  = ibm_cd_tekton_pipeline.ci_pipeline_instance.pipeline_id
  source {
    type = "git"
    properties {
      url = var.tekton_tasks_catalog_repo
      branch = var.definitions_branch
      path = "utils"
    }
  }
}

resource "ibm_cd_tekton_pipeline_trigger" "ci_pipeline_scm_trigger" {
  pipeline_id     = ibm_cd_tekton_pipeline.ci_pipeline_instance.pipeline_id
  type            = var.ci_pipeline_scm_trigger_type
  name            = var.ci_pipeline_scm_trigger_name
  event_listener  = var.ci_pipeline_scm_trigger_listener_name
  source {
    type = "git"
    properties {
      url     = var.app_repo
      branch  = var.app_repo_branch
    }
  }
  events              = ["push"]
  enabled             = var.ci_pipeline_scm_trigger_enabled
  max_concurrent_runs = var.ci_pipeline_max_concurrent_runs
}

resource "ibm_cd_tekton_pipeline_trigger" "ci_pipeline_manual_trigger" {
  pipeline_id     = ibm_cd_tekton_pipeline.ci_pipeline_instance.pipeline_id
  type            = "manual"
  name            = "manual-run"
  event_listener  = "manual-run"
}

resource "ibm_cd_tekton_pipeline_trigger_property" "ci_pipeline_manual_trigger_branch" {
  pipeline_id    = ibm_cd_tekton_pipeline.ci_pipeline_instance.pipeline_id
  trigger_id     = ibm_cd_tekton_pipeline_trigger.ci_pipeline_manual_trigger.trigger_id
  name           = "branch"
  type           = "text"
  value          = var.app_repo_branch
}
