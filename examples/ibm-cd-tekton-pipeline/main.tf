// Base resources
provider "ibm" {
  ibmcloud_api_key = var.ibmcloud_api_key
}
data "ibm_resource_group" "resource_group" {
  name = var.resource_group
}

// Create toolchain instance
resource "ibm_cd_toolchain" "toolchain_instance" {
  name        = var.toolchain_name
  description = var.toolchain_description
  resource_group_id = data.ibm_resource_group.resource_group.id
}

// Create git repo tool instance
resource "ibm_cd_toolchain_tool_hostedgit" "tekton_repo" {
  toolchain_id = ibm_cd_toolchain.toolchain_instance.id
  name         = "tekton-repo"
  initialization {
    type = "clone"
    source_repo_url = var.clone_repo
    private_repo = false
    repo_name = var.repo_name
  }  
  parameters {
    has_issues          = false
    enable_traceability = false
  }
}

// Create tekton pipeline instance
resource "ibm_cd_toolchain_tool_pipeline" "cd_pipeline" {
  toolchain_id = ibm_cd_toolchain.toolchain_instance.id
  parameters {
    name = "tf-pipeline"
    type = "tekton"
  }
}
resource "ibm_cd_tekton_pipeline" "cd_pipeline_instance" {
  pipeline_id = ibm_cd_toolchain_tool_pipeline.cd_pipeline.tool_id
  enable_slack_notifications = false
  enable_partial_cloning = false
  worker {
    id = "public"
  }
}

// Provision cd_tekton_pipeline_definition resource instance
resource "ibm_cd_tekton_pipeline_definition" "cd_tekton_pipeline_definition_instance" {
  pipeline_id = ibm_cd_tekton_pipeline.cd_pipeline_instance.pipeline_id
  scm_source {
    url = ibm_cd_toolchain_tool_hostedgit.tekton_repo.parameters[0].repo_url
    branch = "master"
    path = ".tekton"
  }
}

// Provision cd_tekton_pipeline_property resource instance
resource "ibm_cd_tekton_pipeline_property" "cd_tekton_pipeline_property_instance" {
  pipeline_id = ibm_cd_tekton_pipeline.cd_pipeline_instance.pipeline_id
  name = "env-prop-1"
  value = "Environment text property 1"
  type = "text"
}

// Provision pipeline properties
resource "ibm_cd_tekton_pipeline_property" "apikey" {
  pipeline_id = ibm_cd_tekton_pipeline.cd_pipeline_instance.pipeline_id
  name = "apikey"
  value = var.ibmcloud_api_key
  type = "secure"
}
resource "ibm_cd_tekton_pipeline_property" "cluster_property" {
  pipeline_id = ibm_cd_tekton_pipeline.cd_pipeline_instance.pipeline_id
  name = "cluster"
  value = var.cluster
  type = "text"
}
resource "ibm_cd_tekton_pipeline_property" "cluster_namespace_property" {
  pipeline_id = ibm_cd_tekton_pipeline.cd_pipeline_instance.pipeline_id
  name = "clusterNamespace"
  value = var.cluster_namespace
  type = "text"
}
resource "ibm_cd_tekton_pipeline_property" "cluster_region_property" {
  pipeline_id = ibm_cd_tekton_pipeline.cd_pipeline_instance.pipeline_id
  name = "clusterRegion"
  value = var.region
  type = "text"
}
resource "ibm_cd_tekton_pipeline_property" "registry_region_property" {
  pipeline_id = ibm_cd_tekton_pipeline.cd_pipeline_instance.pipeline_id
  name = "registryRegion"
  value = var.region
  type = "text"
}
resource "ibm_cd_tekton_pipeline_property" "registry_namespace_property" {
  pipeline_id = ibm_cd_tekton_pipeline.cd_pipeline_instance.pipeline_id
  name = "registryNamespace"
  value = var.registry_namespace
  type = "text"
}
resource "ibm_cd_tekton_pipeline_property" "repository_property" {
  pipeline_id = ibm_cd_tekton_pipeline.cd_pipeline_instance.pipeline_id
  name = "repository"
  value = ibm_cd_toolchain_tool_hostedgit.tekton_repo.parameters[0].repo_url
  type = "text"
}

// Provision cd_tekton_pipeline_trigger resource instance
resource "ibm_cd_tekton_pipeline_trigger" "cd_tekton_pipeline_trigger_instance" {
  pipeline_id = ibm_cd_tekton_pipeline.cd_pipeline_instance.pipeline_id
  type = "manual"
  name = "trigger1"
  event_listener = "listener"
  tags = [ "tag1", "tag2" ]
  worker {
    id = "public"
  }
  max_concurrent_runs = 1
}

// Provision cd_tekton_pipeline_trigger_property resource instance
resource "ibm_cd_tekton_pipeline_trigger_property" "cd_tekton_pipeline_trigger_property_instance" {
  pipeline_id = ibm_cd_tekton_pipeline.cd_pipeline_instance.pipeline_id
  trigger_id = ibm_cd_tekton_pipeline_trigger.cd_tekton_pipeline_trigger_instance.trigger_id
  name = "trig-prop-1"
  value = "trigger 1 text property"
  type = "text"
}

// Data sources
// Create cd_tekton_pipeline_definition data source
data "ibm_cd_tekton_pipeline_definition" "cd_tekton_pipeline_definition_instance" {
  pipeline_id = ibm_cd_tekton_pipeline.cd_pipeline_instance.pipeline_id
  definition_id = ibm_cd_tekton_pipeline_definition.cd_tekton_pipeline_definition_instance.definition_id
}
// Create cd_tekton_pipeline_trigger data source
data "ibm_cd_tekton_pipeline_trigger" "cd_tekton_pipeline_trigger_instance" {
  pipeline_id = ibm_cd_tekton_pipeline.cd_pipeline_instance.pipeline_id
  trigger_id = ibm_cd_tekton_pipeline_trigger.cd_tekton_pipeline_trigger_instance.trigger_id
}
// Create cd_tekton_pipeline_trigger_property data source
data "ibm_cd_tekton_pipeline_trigger_property" "cd_tekton_pipeline_trigger_property_instance" {
  pipeline_id = ibm_cd_tekton_pipeline.cd_pipeline_instance.pipeline_id
  trigger_id = ibm_cd_tekton_pipeline_trigger.cd_tekton_pipeline_trigger_instance.trigger_id
  property_name = ibm_cd_tekton_pipeline_trigger_property.cd_tekton_pipeline_trigger_property_instance.name
}
// Create cd_tekton_pipeline_property data source
data "ibm_cd_tekton_pipeline_property" "cd_tekton_pipeline_property_instance" {
  pipeline_id = ibm_cd_tekton_pipeline.cd_pipeline_instance.pipeline_id
  property_name = ibm_cd_tekton_pipeline_property.cd_tekton_pipeline_property_instance.name
}
// Create cd_tekton_pipeline data source
data "ibm_cd_tekton_pipeline" "cd_tekton_pipeline_instance" {
  pipeline_id = ibm_cd_tekton_pipeline.cd_pipeline_instance.pipeline_id
}
