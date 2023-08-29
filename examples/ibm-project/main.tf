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
    project_id = "project_id"
    version = 1
    is_draft = true
    needs_attention_state = [ "anything as a string" ]
    state = "approved"
    approved_version {
      needs_attention_state = [ "anything as a string" ]
      state = "approved"
      version = 1
      href = "href"
    }
    installed_version {
      needs_attention_state = [ "anything as a string" ]
      state = "approved"
      version = 1
      href = "href"
    }
    definition {
      name = "name"
      description = "description"
    }
    check_job {
      id = "id"
      href = "href"
    }
    install_job {
      id = "id"
      href = "href"
    }
    uninstall_job {
      id = "id"
      href = "href"
    }
    href = "href"
  }
}

// Provision project_config resource instance
resource "ibm_project_config" "project_config_instance" {
  project_id = ibm_project.project_instance.id
  name = var.project_config_name
  description = var.project_config_description
  labels = var.project_config_labels
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
  }
  setting {
  }
}

// Create project data source
data "ibm_project" "project_instance" {
  id = ibm_project.project_instance.id
}

// Create project_config data source
data "ibm_project_config" "project_config_instance" {
  project_id = ibm_project.project_instance.id
  id = ibm_project_config.project_config_instance.projectConfigCanonical_id
}
