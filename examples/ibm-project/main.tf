provider "ibm" {
  ibmcloud_api_key = var.ibmcloud_api_key
}

// Provision project_instance resource instance
resource "ibm_project_instance" "project_instance_instance" {
  name = var.project_instance_name
  description = var.project_instance_description
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
  resource_group = var.project_instance_resource_group
  location = var.project_instance_location
}

// Data source is not linked to a resource instance
// Uncomment if an existing data source instance exists
/*
// Create project_event_notification data source
data "ibm_project_event_notification" "project_event_notification_instance" {
  id = var.project_event_notification_id
  exclude_configs = var.project_event_notification_exclude_configs
  complete = var.project_event_notification_complete
}
*/
