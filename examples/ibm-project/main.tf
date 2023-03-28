provider "ibm" {
  ibmcloud_api_key = var.ibmcloud_api_key
}

// Provision project resource instance
resource "ibm_project" "project_instance" {
  name = var.project_name
  description = var.project_description
  configs {
    id = "id"
    name = "name"
    labels = [ "labels" ]
    description = "description"
    locator_id = "locator_id"
    input {
      name = "name"
    }
    setting {
      name = "name"
      value = "value"
    }
  }
  resource_group = var.project_resource_group
  location = var.project_location
}

// Data source is not linked to a resource instance
// Uncomment if an existing data source instance exists
/*
// Create project data source
data "ibm_project" "project_instance" {
  id = var.project_id
  exclude_configs = var.project_exclude_configs
  complete = var.project_complete
}
*/

// Data source is not linked to a resource instance
// Uncomment if an existing data source instance exists
/*
// Create event_notification data source
data "ibm_event_notification" "event_notification_instance" {
  id = var.event_notification_id
}
*/
