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
    type = "terraform_template"
    input {
      name = "name"
      type = "array"
      value = "anything as a string"
      required = true
    }
    output {
      name = "name"
      description = "description"
      value = "anything as a string"
    }
    setting {
      name = "name"
      value = "value"
    }
    id = "id"
    project_id = "project_id"
    version = 1
    is_draft = true
    needs_attention_state = [ "anything as a string" ]
    state = "deleted"
    pipeline_state = "pipeline_failed"
    update_available = true
    created_at = "2021-01-31T09:44:12Z"
    updated_at = "2021-01-31T09:44:12Z"
    last_approved {
      is_forced = true
      comment = "comment"
      timestamp = "2021-01-31T09:44:12Z"
      user_id = "user_id"
    }
    last_save = "2021-01-31T09:44:12Z"
    job_summary {
      plan_summary = { "key" = "anything as a string" }
      apply_summary = { "key" = "anything as a string" }
      destroy_summary = { "key" = "anything as a string" }
      message_summary = { "key" = "anything as a string" }
      plan_messages = { "key" = "anything as a string" }
      apply_messages = { "key" = "anything as a string" }
      destroy_messages = { "key" = "anything as a string" }
    }
    cra_logs {
      cra_version = "cra_version"
      schema_version = "schema_version"
      status = "status"
      summary = { "key" = "anything as a string" }
      timestamp = "2021-01-31T09:44:12Z"
    }
    cost_estimate {
      version = "version"
      currency = "currency"
      total_hourly_cost = "total_hourly_cost"
      total_monthly_cost = "total_monthly_cost"
      past_total_hourly_cost = "past_total_hourly_cost"
      past_total_monthly_cost = "past_total_monthly_cost"
      diff_total_hourly_cost = "diff_total_hourly_cost"
      diff_total_monthly_cost = "diff_total_monthly_cost"
      time_generated = "2021-01-31T09:44:12Z"
      user_id = "user_id"
    }
    last_deployment_job_summary {
      plan_summary = { "key" = "anything as a string" }
      apply_summary = { "key" = "anything as a string" }
      destroy_summary = { "key" = "anything as a string" }
      message_summary = { "key" = "anything as a string" }
      plan_messages = { "key" = "anything as a string" }
      apply_messages = { "key" = "anything as a string" }
      destroy_messages = { "key" = "anything as a string" }
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
