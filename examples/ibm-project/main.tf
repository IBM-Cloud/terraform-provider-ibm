provider "ibm" {
  ibmcloud_api_key = var.ibmcloud_api_key
}

// Provision project resource instance
resource "ibm_project" "project_instance" {
  resource_group = var.project_resource_group
  location = var.project_location
  name = var.project_name
  description = var.project_description
  destroy_on_delete = var.project_destroy_on_delete
  configs {
    id = "id"
    name = "name"
    labels = [ "labels" ]
    description = "description"
    authorizations {
      method = "API_KEY"
      api_key = "api_key_value"
    }
    locator_id = "locator_id"
    input {
      name = "app_repo_name"
    }
    setting {
      name = "app_repo_name"
      value = "static-website-dev-app-repo"
    }
  }
}

// Provision project_config resource instance
resource "ibm_project_config" "project_config_instance" {
  project_id = ibm_project.project_instance.id
  name = var.project_config_name
  locator_id = var.project_config_locator_id
  labels = var.project_config_labels
  description = var.project_config_description
  authorizations {
    trusted_profile {
      id = "Profile-047ef95f-2f90-4276-8ba3-92512cc4a6ae"
    }
    method = "TRUSTED_PROFILE"
  }
  compliance_profile {
    id = "compliance_profile_id"
    instance_id = "instance_id"
    instance_location = "instance_location"
    attachment_id = "attachment_id"
    profile_name = "profile_name"
  }
  input {
    name = "app_repo_name"
  }
  setting {
    name = "app_repo_name"
    value = "static-website-dev-app-repo"
  }
}

// Create project data source
data "ibm_project" "project_instance" {
  id = ibm_project.project_instance.id
}

// Create project_config data source
data "ibm_project_config" "project_config_instance" {
  project_id = ibm_project.project_instance.id
  id = ibm_project_config.project_config_instance.project_config_id
  version = var.project_config_version
}
