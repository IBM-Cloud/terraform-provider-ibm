resource "ibm_cd_tekton_pipeline_property" "pr_env_apikey" {
  name           = "apikey"
  type           = "SECURE"
  value          = format("{vault::%s.ibmcloud-api-key}", var.kp_integration_name)
  pipeline_id    = var.pipeline_id           
}

resource "ibm_cd_tekton_pipeline_property" "pr_env_ibmcloud-api" {
  name           = "ibmcloud-api"
  type           = "TEXT"
  value          = "https://cloud.ibm.com"
  pipeline_id    = var.pipeline_id         
}

resource "ibm_cd_tekton_pipeline_property" "pr_env_pipeline-debug" {
  name           = "pipeline-debug"
  type           = "TEXT"
  value          = "0"
  pipeline_id    = var.pipeline_id         
}