provider "ibm" {
  ibmcloud_api_key = var.ibmcloud_api_key
}

// Provision tekton_pipeline_definition resource instance
resource "ibm_tekton_pipeline_definition" "tekton_pipeline_definition_instance" {
  pipeline_id = var.tekton_pipeline_definition_pipeline_id
  scm_source {
    url = "url"
    branch = "branch"
    path = "path"
  }
}

// Provision tekton_pipeline_trigger_property resource instance
resource "ibm_tekton_pipeline_trigger_property" "tekton_pipeline_trigger_property_instance" {
  pipeline_id = var.tekton_pipeline_trigger_property_pipeline_id
  trigger_id = var.tekton_pipeline_trigger_property_trigger_id
  name = var.tekton_pipeline_trigger_property_name
  value = var.tekton_pipeline_trigger_property_value
  options = var.tekton_pipeline_trigger_property_options
  type = var.tekton_pipeline_trigger_property_type
  path = var.tekton_pipeline_trigger_property_path
}

// Provision tekton_pipeline_property resource instance
resource "ibm_tekton_pipeline_property" "tekton_pipeline_property_instance" {
  pipeline_id = var.tekton_pipeline_property_pipeline_id
  name = var.tekton_pipeline_property_name
  value = var.tekton_pipeline_property_value
  options = var.tekton_pipeline_property_options
  type = var.tekton_pipeline_property_type
  path = var.tekton_pipeline_property_path
}

// Provision tekton_pipeline_trigger resource instance
resource "ibm_tekton_pipeline_trigger" "tekton_pipeline_trigger_instance" {
  pipeline_id = var.tekton_pipeline_trigger_pipeline_id
  trigger {
    type = "type"
    name = "start-deploy"
    event_listener = "event_listener"
    id = "id"
    properties {
      name = "name"
      value = "value"
      type = "SECURE"
      path = "path"
      href = "href"
    }
    tags = [ "tags" ]
    worker {
      name = "name"
      type = "private"
      id = "id"
    }
    concurrency {
      max_concurrent_runs = 20
    }
    disabled = true
  }
}

// Create tekton_pipeline_definition data source
data "ibm_tekton_pipeline_definition" "tekton_pipeline_definition_instance" {
  pipeline_id = var.tekton_pipeline_definition_pipeline_id
  definition_id = var.tekton_pipeline_definition_definition_id
}

// Create tekton_pipeline_trigger_property data source
data "ibm_tekton_pipeline_trigger_property" "tekton_pipeline_trigger_property_instance" {
  pipeline_id = var.tekton_pipeline_trigger_property_pipeline_id
  trigger_id = var.tekton_pipeline_trigger_property_trigger_id
  property_name = var.tekton_pipeline_trigger_property_property_name
}

// Create tekton_pipeline_property data source
data "ibm_tekton_pipeline_property" "tekton_pipeline_property_instance" {
  pipeline_id = var.tekton_pipeline_property_pipeline_id
  property_name = var.tekton_pipeline_property_property_name
}

// Create tekton_pipeline_workers data source
data "ibm_tekton_pipeline_workers" "tekton_pipeline_workers_instance" {
  pipeline_id = var.tekton_pipeline_workers_pipeline_id
}

// Create tekton_pipeline_trigger data source
data "ibm_tekton_pipeline_trigger" "tekton_pipeline_trigger_instance" {
  pipeline_id = var.tekton_pipeline_trigger_pipeline_id
  trigger_id = var.tekton_pipeline_trigger_trigger_id
}
