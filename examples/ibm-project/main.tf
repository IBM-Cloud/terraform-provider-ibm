provider "ibm" {
  ibmcloud_api_key = var.ibmcloud_api_key
}

// Provision project_instance resource instance
resource "ibm_project_instance" "project_instance_instance" {
  resource_group = var.project_instance_resource_group
  location = var.project_instance_location
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
      value = "anything as a string"
    }
    setting {
      name = "name"
      value = "value"
    }
  }
}

// Create project_event_notification data source
data "ibm_project_event_notification" "project_event_notification_instance" {
  id = ibm_project_instance.project_instance_instance.id
}
