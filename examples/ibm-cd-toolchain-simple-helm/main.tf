data "ibm_resource_group" "resource_group" {
  name = var.resource_group
}

resource "ibm_cd_toolchain" "toolchain_instance" {
  name        = var.toolchain_name
  description = var.toolchain_description
  resource_group_id = data.ibm_resource_group.resource_group.id
}

module "repositories" {
  source                          = "./repositories"
  toolchain_id                    = ibm_cd_toolchain.toolchain_instance.id
  resource_group                  = data.ibm_resource_group.resource_group.id  
  ibmcloud_api_key                = var.ibmcloud_api_key
  region                          = var.region  
  app_repo                        = var.app_repo
  pipeline_repo                   = var.pipeline_repo
  tekton_tasks_catalog_repo       = var.tekton_tasks_catalog_repo
  repositories_prefix             = var.app_name
}

resource "ibm_cd_toolchain_tool_pipeline" "ci_pipeline" {
  toolchain_id = ibm_cd_toolchain.toolchain_instance.id
  parameters {
    name = "ci-pipeline"
  }
}

module "pipeline-ci" {
  source                    = "./pipeline-ci"
  depends_on                = [ module.repositories ]
  ibmcloud_api_key          = var.ibmcloud_api_key
  ibmcloud_api              = var.ibmcloud_api
  region                    = var.region
  pipeline_id               = split("/", ibm_cd_toolchain_tool_pipeline.ci_pipeline.id)[1]
  resource_group            = var.resource_group
  app_name                  = var.app_name
  app_image_name            = var.app_image_name  
  cluster_name              = var.cluster_name
  cluster_namespace         = var.cluster_namespace
  cluster_region            = var.cluster_region
  registry_namespace        = var.registry_namespace
  registry_region           = var.registry_region
  commons_hosted_region     = var.commons_hosted_region
  app_repo                  = module.repositories.app_repo_url 
  app_repo_branch           = var.app_repo_branch
  pipeline_repo             = module.repositories.pipeline_repo_url
  pipeline_repo_branch      = var.pipeline_repo_branch
  tekton_tasks_catalog_repo = module.repositories.tekton_tasks_catalog_repo_url
  definitions_branch        = var.definitions_branch
  kp_integration_name       = module.integrations.keyprotect_integration_name
}

resource "ibm_cd_toolchain_tool_pipeline" "pr_pipeline" {
  toolchain_id = ibm_cd_toolchain.toolchain_instance.id
  parameters {
    name = "pr-pipeline"
  }
}

module "pipeline-pr" {
  source                    = "./pipeline-pr"
  depends_on                = [ module.repositories ]
  ibmcloud_api_key          = var.ibmcloud_api_key
  ibmcloud_api              = var.ibmcloud_api
  region                    = var.region  
  pipeline_id               = split("/", ibm_cd_toolchain_tool_pipeline.pr_pipeline.id)[1]
  resource_group            = var.resource_group
  app_name                  = var.app_name
  app_repo                  = module.repositories.app_repo_url 
  app_repo_branch           = var.app_repo_branch
  pipeline_repo             = module.repositories.pipeline_repo_url
  pipeline_repo_branch      = var.pipeline_repo_branch
  tekton_tasks_catalog_repo = module.repositories.tekton_tasks_catalog_repo_url  
  definitions_branch        = var.definitions_branch
  kp_integration_name       = module.integrations.keyprotect_integration_name
}

module "services" {
  source                      = "./services"
  kp_name                     = var.kp_name
  kp_region                   = var.kp_region
  region                      = var.region
  ibmcloud_api                = var.ibmcloud_api
  cluster_name                = var.cluster_name
  cluster_namespace           = var.cluster_namespace
  cluster_region              = var.cluster_region
  registry_namespace          = var.registry_namespace
  registry_region             = var.registry_region   
}

module "integrations" {
  source                      = "./integrations"
  depends_on                  = [ module.repositories, module.services ]
  toolchain_id                = ibm_cd_toolchain.toolchain_instance.id
  region                      = var.region  
  resource_group              = var.resource_group
  key_protect_instance_region = var.kp_region
  key_protect_instance_name   = module.services.key_protect_instance_name
  key_protect_instance_guid   = module.services.key_protect_instance_guid
}

output "toolchain_id" {
  value = ibm_cd_toolchain.toolchain_instance.id
}