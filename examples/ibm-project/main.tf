provider "ibm" {
  ibmcloud_api_key = var.ibmcloud_api_key
}

// Provision project_config resource instance
resource "ibm_project_config" "project_config_instance" {
  project_id = ibm_project.project_instance.project_id
  definition {
    name = "name"
    description = "description"
    labels = [ "labels" ]
    authorizations {
      trusted_profile {
        id = "id"
        target_iam_id = "target_iam_id"
      }
      method = "method"
      api_key = "api_key"
    }
    compliance_profile {
      id = "id"
      instance_id = "instance_id"
      instance_location = "instance_location"
      attachment_id = "attachment_id"
      profile_name = "profile_name"
    }
    locator_id = "locator_id"
    input = {  }
    setting = {  }
  }
}

// Provision project resource instance
resource "ibm_project" "project_instance" {
  resource_group = var.project_resource_group
  location = var.project_location
  definition {
    name = "name"
    description = "description"
    destroy_on_delete = true
  }
}

// Create project_config data source
data "ibm_project_config" "project_config_instance" {
  project_id = ibm_project.project_instance.id
  id = ibm_project_config.project_config_instance.project_config_id
}

// Create project data source
data "ibm_project" "project_instance" {
  id = ibm_project.project_instance.id
}
