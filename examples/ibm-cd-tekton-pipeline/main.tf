provider "ibm" {
  ibmcloud_api_key = var.ibmcloud_api_key
}

// Provision tekton_pipeline_definition resource instance
resource "ibm_cd_tekton_pipeline_definition" "tekton_pipeline_definition_instance" {
  pipeline_id = var.tekton_pipeline_definition_pipeline_id
  scm_source {
    url = "url"
    branch = "branch"
    tag = "tag"
    path = "path"
  }
}

// Provision tekton_pipeline_trigger_property resource instance
resource "ibm_cd_tekton_pipeline_trigger_property" "tekton_pipeline_trigger_property_instance" {
  pipeline_id = var.tekton_pipeline_trigger_property_pipeline_id
  trigger_id = var.tekton_pipeline_trigger_property_trigger_id
  name = var.tekton_pipeline_trigger_property_name
  value = var.tekton_pipeline_trigger_property_value
  enum = var.tekton_pipeline_trigger_property_enum
  default = var.tekton_pipeline_trigger_property_default
  type = var.tekton_pipeline_trigger_property_type
  path = var.tekton_pipeline_trigger_property_path
}

// Provision tekton_pipeline_property resource instance
resource "ibm_cd_tekton_pipeline_property" "tekton_pipeline_property_instance" {
  pipeline_id = var.tekton_pipeline_property_pipeline_id
  name = var.tekton_pipeline_property_name
  value = var.tekton_pipeline_property_value
  enum = var.tekton_pipeline_property_enum
  default = var.tekton_pipeline_property_default
  type = var.tekton_pipeline_property_type
  path = var.tekton_pipeline_property_path
}

// Provision tekton_pipeline_trigger resource instance
resource "ibm_cd_tekton_pipeline_trigger" "tekton_pipeline_trigger_instance" {
  pipeline_id = var.tekton_pipeline_trigger_pipeline_id
  trigger {
    source_trigger_id = "source_trigger_id"
    name = "start-deploy"
  }
}

// Provision tekton_pipeline resource instance
resource "ibm_cd_tekton_pipeline" "tekton_pipeline_instance" {
  worker {
    id = "id"
  }
}

// Create tekton_pipeline_definition data source
data "ibm_cd_tekton_pipeline_definition" "tekton_pipeline_definition_instance" {
  pipeline_id = var.tekton_pipeline_definition_pipeline_id
  definition_id = var.tekton_pipeline_definition_definition_id
}

// Create tekton_pipeline_trigger_property data source
data "ibm_cd_tekton_pipeline_trigger_property" "tekton_pipeline_trigger_property_instance" {
  pipeline_id = var.tekton_pipeline_trigger_property_pipeline_id
  trigger_id = var.tekton_pipeline_trigger_property_trigger_id
  property_name = var.tekton_pipeline_trigger_property_property_name
}

// Create tekton_pipeline_property data source
data "ibm_cd_tekton_pipeline_property" "tekton_pipeline_property_instance" {
  pipeline_id = var.tekton_pipeline_property_pipeline_id
  property_name = var.tekton_pipeline_property_property_name
}

// Create tekton_pipeline_trigger data source
data "ibm_cd_tekton_pipeline_trigger" "tekton_pipeline_trigger_instance" {
  pipeline_id = var.tekton_pipeline_trigger_pipeline_id
  trigger_id = var.tekton_pipeline_trigger_trigger_id
}

// Create tekton_pipeline data source
data "ibm_cd_tekton_pipeline" "tekton_pipeline_instance" {
  id = var.tekton_pipeline_id
}
