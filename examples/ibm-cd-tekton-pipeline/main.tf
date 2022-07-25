provider "ibm" {
  ibmcloud_api_key = var.ibmcloud_api_key
}

// Provision cd_tekton_pipeline_definition resource instance
resource "ibm_cd_tekton_pipeline_definition" "cd_tekton_pipeline_definition_instance" {
  pipeline_id = var.cd_tekton_pipeline_definition_pipeline_id
  scm_source {
    url = "url"
    branch = "branch"
    tag = "tag"
    path = "path"
  }
}

// Provision cd_tekton_pipeline_trigger_property resource instance
resource "ibm_cd_tekton_pipeline_trigger_property" "cd_tekton_pipeline_trigger_property_instance" {
  pipeline_id = var.cd_tekton_pipeline_trigger_property_pipeline_id
  trigger_id = var.cd_tekton_pipeline_trigger_property_trigger_id
  name = var.cd_tekton_pipeline_trigger_property_name
  value = var.cd_tekton_pipeline_trigger_property_value
  enum = var.cd_tekton_pipeline_trigger_property_enum
  default = var.cd_tekton_pipeline_trigger_property_default
  type = var.cd_tekton_pipeline_trigger_property_type
  path = var.cd_tekton_pipeline_trigger_property_path
}

// Provision cd_tekton_pipeline_property resource instance
resource "ibm_cd_tekton_pipeline_property" "cd_tekton_pipeline_property_instance" {
  pipeline_id = var.cd_tekton_pipeline_property_pipeline_id
  name = var.cd_tekton_pipeline_property_name
  value = var.cd_tekton_pipeline_property_value
  enum = var.cd_tekton_pipeline_property_enum
  default = var.cd_tekton_pipeline_property_default
  type = var.cd_tekton_pipeline_property_type
  path = var.cd_tekton_pipeline_property_path
}

// Provision cd_tekton_pipeline_trigger resource instance
resource "ibm_cd_tekton_pipeline_trigger" "cd_tekton_pipeline_trigger_instance" {
  pipeline_id = var.cd_tekton_pipeline_trigger_pipeline_id
  trigger {
    source_trigger_id = "source_trigger_id"
    name = "start-deploy"
  }
}

// Provision cd_tekton_pipeline resource instance
resource "ibm_cd_tekton_pipeline" "cd_tekton_pipeline_instance" {
  worker {
    id = "id"
  }
}

// Create cd_tekton_pipeline_definition data source
data "ibm_cd_tekton_pipeline_definition" "cd_tekton_pipeline_definition_instance" {
  pipeline_id = var.cd_tekton_pipeline_definition_pipeline_id
  definition_id = var.cd_tekton_pipeline_definition_definition_id
}

// Create cd_tekton_pipeline_trigger_property data source
data "ibm_cd_tekton_pipeline_trigger_property" "cd_tekton_pipeline_trigger_property_instance" {
  pipeline_id = var.cd_tekton_pipeline_trigger_property_pipeline_id
  trigger_id = var.cd_tekton_pipeline_trigger_property_trigger_id
  property_name = var.cd_tekton_pipeline_trigger_property_property_name
}

// Create cd_tekton_pipeline_property data source
data "ibm_cd_tekton_pipeline_property" "cd_tekton_pipeline_property_instance" {
  pipeline_id = var.cd_tekton_pipeline_property_pipeline_id
  property_name = var.cd_tekton_pipeline_property_property_name
}

// Create cd_tekton_pipeline_trigger data source
data "ibm_cd_tekton_pipeline_trigger" "cd_tekton_pipeline_trigger_instance" {
  pipeline_id = var.cd_tekton_pipeline_trigger_pipeline_id
  trigger_id = var.cd_tekton_pipeline_trigger_trigger_id
}

// Create cd_tekton_pipeline data source
data "ibm_cd_tekton_pipeline" "cd_tekton_pipeline_instance" {
  id = var.cd_tekton_pipeline_id
}
