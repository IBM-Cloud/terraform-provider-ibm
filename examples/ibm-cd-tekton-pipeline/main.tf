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
    service_instance_id = "service_instance_id"
  }
}

// Provision cd_tekton_pipeline_trigger_property resource instance
resource "ibm_cd_tekton_pipeline_trigger_property" "cd_tekton_pipeline_trigger_property_instance" {
  pipeline_id = var.cd_tekton_pipeline_trigger_property_pipeline_id
  trigger_id = var.cd_tekton_pipeline_trigger_property_trigger_id
  name = var.cd_tekton_pipeline_trigger_property_name
  value = var.cd_tekton_pipeline_trigger_property_value
  enum = var.cd_tekton_pipeline_trigger_property_enum
  type = var.cd_tekton_pipeline_trigger_property_type
  path = var.cd_tekton_pipeline_trigger_property_path
}

// Provision cd_tekton_pipeline_property resource instance
resource "ibm_cd_tekton_pipeline_property" "cd_tekton_pipeline_property_instance" {
  pipeline_id = var.cd_tekton_pipeline_property_pipeline_id
  name = var.cd_tekton_pipeline_property_name
  value = var.cd_tekton_pipeline_property_value
  enum = var.cd_tekton_pipeline_property_enum
  type = var.cd_tekton_pipeline_property_type
  path = var.cd_tekton_pipeline_property_path
}

// Provision cd_tekton_pipeline_trigger resource instance
resource "ibm_cd_tekton_pipeline_trigger" "cd_tekton_pipeline_trigger_instance" {
  pipeline_id = var.cd_tekton_pipeline_trigger_pipeline_id
  type = var.cd_tekton_pipeline_trigger_type
  name = var.cd_tekton_pipeline_trigger_name
  event_listener = var.cd_tekton_pipeline_trigger_event_listener
  tags = var.cd_tekton_pipeline_trigger_tags
  worker {
    name = "name"
    type = "private"
    id = "id"
  }
  max_concurrent_runs = var.cd_tekton_pipeline_trigger_max_concurrent_runs
  disabled = var.cd_tekton_pipeline_trigger_disabled
  secret {
    type = "token_matches"
    value = "value"
    source = "header"
    key_name = "key_name"
    algorithm = "md4"
  }
  cron = var.cd_tekton_pipeline_trigger_cron
  timezone = var.cd_tekton_pipeline_trigger_timezone
  scm_source {
    url = "url"
    branch = "branch"
    pattern = "pattern"
    blind_connection = true
    hook_id = "hook_id"
    service_instance_id = "service_instance_id"
  }
  events {
    push = true
    pull_request_closed = true
    pull_request = true
  }
}

// Provision cd_tekton_pipeline resource instance
resource "ibm_cd_tekton_pipeline" "cd_tekton_pipeline_instance" {
  enable_slack_notifications = var.cd_tekton_pipeline_enable_slack_notifications
  enable_partial_cloning = var.cd_tekton_pipeline_enable_partial_cloning
  worker {
    id = "id"
  }
}

// Data source is not linked to a resource instance
// Uncomment if an existing data source instance exists
/*
// Create cd_tekton_pipeline_definition data source
data "ibm_cd_tekton_pipeline_definition" "cd_tekton_pipeline_definition_instance" {
  pipeline_id = var.cd_tekton_pipeline_definition_pipeline_id
  definition_id = var.cd_tekton_pipeline_definition_definition_id
}
*/

// Data source is not linked to a resource instance
// Uncomment if an existing data source instance exists
/*
// Create cd_tekton_pipeline_trigger_property data source
data "ibm_cd_tekton_pipeline_trigger_property" "cd_tekton_pipeline_trigger_property_instance" {
  pipeline_id = var.cd_tekton_pipeline_trigger_property_pipeline_id
  trigger_id = var.cd_tekton_pipeline_trigger_property_trigger_id
  property_name = var.cd_tekton_pipeline_trigger_property_property_name
}
*/

// Data source is not linked to a resource instance
// Uncomment if an existing data source instance exists
/*
// Create cd_tekton_pipeline_property data source
data "ibm_cd_tekton_pipeline_property" "cd_tekton_pipeline_property_instance" {
  pipeline_id = var.cd_tekton_pipeline_property_pipeline_id
  property_name = var.cd_tekton_pipeline_property_property_name
}
*/

// Data source is not linked to a resource instance
// Uncomment if an existing data source instance exists
/*
// Create cd_tekton_pipeline_trigger data source
data "ibm_cd_tekton_pipeline_trigger" "cd_tekton_pipeline_trigger_instance" {
  pipeline_id = var.cd_tekton_pipeline_trigger_pipeline_id
  trigger_id = var.cd_tekton_pipeline_trigger_trigger_id
}
*/

// Data source is not linked to a resource instance
// Uncomment if an existing data source instance exists
/*
// Create cd_tekton_pipeline data source
data "ibm_cd_tekton_pipeline" "cd_tekton_pipeline_instance" {
  id = var.cd_tekton_pipeline_id
}
*/
