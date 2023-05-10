provider "ibm" {
  ibmcloud_api_key = var.ibmcloud_api_key
}

// Provision project_instance resource instance
resource "ibm_project_instance" "project_instance" {
  name = "My static website"
  description = "Sample static website test using the IBM catalog deployable architecture"
  configs {
    name = "static-website-dev"
    labels = [ "env:dev", "billing:internal" ]
    description = "Website - development"
    locator_id = "1082e7d2-5e2f-0a11-a3bc-f88a8e1931fc.145be7c1-9ec4-4719-b586-584ee52fbed0-global"
    input {
      name = "app_repo_name"
    }
    setting {
      name = "app_repo_name"
      value = "static-website-dev-app-repo"
    }
  }
  resource_group = "Default"
  location = "us-south"
}

// Data source is not linked to a resource instance
// Uncomment if an existing data source instance exists
/*
// Create project_event_notification data source
data "ibm_project_event_notification" "project_event_notification_instance" {
  project_id = var.project_event_notification_id
}
*/
