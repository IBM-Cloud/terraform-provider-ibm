provider "ibm" {
  ibmcloud_api_key = var.ibmcloud_api_key
}

// Provision project_config resource instance
resource "ibm_project_config" "project_config_instance" {
  project_id = ibm_project.project_instance.id
  definition {
    name = "static-website-dev"
    description = "Website - development"
    authorizations {
      method = "api_key"
      api_key = "<your_apikey_here>"
    }
    locator_id = "1082e7d2-5e2f-0a11-a3bc-f88a8e1931fc.145be7c1-9ec4-4719-b586-584ee52fbed0-global"
    inputs = {
      app_repo_name = "static-website-repo"
    }
  }
}

// Provision project resource instance
resource "ibm_project" "project_instance" {
  location = var.project_location
  resource_group = var.project_resource_group
  definition {
    name = "My static website"
    description = "Sample static website test using the IBM catalog deployable architecture"
    destroy_on_delete = true
    monitoring_enabled = true
    auto_deploy = true
  }
}

// Provision project_environment resource instance
resource "ibm_project_environment" "project_environment_instance" {
  project_id = ibm_project.project_instance.id
  definition {
    name = "environment-stage"
    description = "environment for stage project"
    authorizations {
      method = "api_key"
      api_key = "<your_apikey_here>"
    }
  }
}

// Create project_config data source
data "ibm_project_config" "project_config_instance" {
  project_id = ibm_project.project_instance.id
  project_config_id = ibm_project_config.project_config_instance.project_config_id
}

// Create project data source
data "ibm_project" "project_instance" {
  project_id = ibm_project.project_instance.id
}

// Create project_environment data source
data "ibm_project_environment" "project_environment_instance" {
  project_id = ibm_project.project_instance.id
  project_environment_id = ibm_project_environment.project_environment_instance.project_environment_id
}
