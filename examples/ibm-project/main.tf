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
    name = "name"
    labels = [ "labels" ]
    description = "description"
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

// Provision project_config resource instance
resource "ibm_project_config" "project_config_instance" {
  project_id = ibm_project.project_instance.id
  name = var.project_config_name
  labels = var.project_config_labels
  description = var.project_config_description
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
  locator_id = var.project_config_locator_id
  input {
    name = "name"
    type = "array"
    value = "anything as a string"
    required = true
  }
  setting {
    name = "name"
    value = "value"
  }
}

// Create project data source
data "ibm_project" "project_instance" {
  id = ibm_project.project_instance.id
}

// Create project_config data source
data "ibm_project_config" "project_config_instance" {
  project_id = ibm_project.project_instance.id
  id = ibm_project_config.project_config_instance.projectConfig_id
}
