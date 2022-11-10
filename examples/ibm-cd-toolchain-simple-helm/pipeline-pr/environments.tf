resource "ibm_cd_tekton_pipeline_property" "pr_env_apikey" {
  name           = "apikey"
  type           = "secure"
  value          = format("{vault::%s.ibmcloud-api-key}", var.kp_integration_name)
  pipeline_id   = ibm_cd_tekton_pipeline.pr_pipeline_instance.pipeline_id
}

resource "ibm_cd_tekton_pipeline_property" "pr_env_ibmcloud-api" {
  name           = "ibmcloud-api"
  type           = "text"
  value          = var.ibmcloud_api
  pipeline_id   = ibm_cd_tekton_pipeline.pr_pipeline_instance.pipeline_id
}

resource "ibm_cd_tekton_pipeline_property" "pr_env_pipeline-debug" {
  name           = "pipeline-debug"
  type           = "text"
  value          = "0"
  pipeline_id   = ibm_cd_tekton_pipeline.pr_pipeline_instance.pipeline_id
}