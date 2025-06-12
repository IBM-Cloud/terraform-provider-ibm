provider "ibm" {
  ibmcloud_api_key = var.ibmcloud_api_key
}

// Provision schematics_workspace resource instance
resource "ibm_schematics_workspace" "schematics_workspace_instance" {
  applied_shareddata_ids = var.schematics_workspace_applied_shareddata_ids
  catalog_ref = var.schematics_workspace_catalog_ref
  description = var.schematics_workspace_description
  location = var.schematics_workspace_location
  name = var.schematics_workspace_name
  resource_group = var.schematics_workspace_resource_group
  shared_data = var.schematics_workspace_shared_data
  tags = var.schematics_workspace_tags
  template_data = var.schematics_workspace_template_data
  template_ref = var.schematics_workspace_template_ref
  template_repo = var.schematics_workspace_template_repo
  type = var.schematics_workspace_type
  workspace_status = var.schematics_workspace_workspace_status
  x_github_token = var.schematics_workspace_x_github_token
}

// Provision schematics_action resource instance
resource "ibm_schematics_action" "schematics_action_instance" {
  name = var.schematics_action_name
  description = var.schematics_action_description
  location = var.schematics_action_location
  resource_group = var.schematics_action_resource_group
  tags = var.schematics_action_tags
  user_state = var.schematics_action_user_state
  source_readme_url = var.schematics_action_source_readme_url
  source = var.schematics_action_source
  source_type = var.schematics_action_source_type
  command_parameter = var.schematics_action_command_parameter
  bastion = var.schematics_action_bastion
  targets_ini = var.schematics_action_targets_ini
  credentials = var.schematics_action_credentials
  inputs = var.schematics_action_inputs
  outputs = var.schematics_action_outputs
  settings = var.schematics_action_settings
  trigger_record_id = var.schematics_action_trigger_record_id
  state = var.schematics_action_state
  sys_lock = var.schematics_action_sys_lock
  x_github_token = var.schematics_action_x_github_token
}

// Provision schematics_job resource instance
resource "ibm_schematics_job" "schematics_job_instance" {
  command_object = var.schematics_job_command_object
  command_object_id = var.schematics_job_command_object_id
  command_name = var.schematics_job_command_name
  command_parameter = var.schematics_job_command_parameter
  command_options = var.schematics_job_command_options
  inputs = var.schematics_job_inputs
  settings = var.schematics_job_settings
  tags = var.schematics_job_tags
  location = var.schematics_job_location
  status = var.schematics_job_status
  data = var.schematics_job_data
  bastion = var.schematics_job_bastion
  log_summary = var.schematics_job_log_summary
  x_github_token = var.schematics_job_x_github_token
}

// Create schematics_output data source
data "ibm_schematics_output" "schematics_output_instance" {
  workspace_id = var.schematics_output_workspace_id
}

// Create schematics_state data source
data "ibm_schematics_state" "schematics_state_instance" {
  workspace_id = var.schematics_state_workspace_id
  template_id = var.schematics_state_template_id
}

// Create schematics_workspace data source
data "ibm_schematics_workspace" "schematics_workspace_instance" {
  workspace_id = var.schematics_workspace_workspace_id
}

// Create schematics_action data source
data "ibm_schematics_action" "schematics_action_instance" {
  action_id = var.schematics_action_action_id
}

// Create schematics_job data source
data "ibm_schematics_job" "schematics_job_instance" {
  job_id = var.schematics_job_job_id
}

// Provision schematics_policy resource instance
resource "ibm_schematics_policy" "schematics_policy_instance" {
  name = var.schematics_policy_name
  description = var.schematics_policy_description
  resource_group = var.schematics_policy_resource_group
  tags = var.schematics_policy_tags
  location = var.schematics_policy_location
  kind = var.schematics_policy_kind
  target {
    selector_kind = "ids"
    selector_ids = [ "selector_ids" ]
    selector_scope {
      kind = "workspace"
      tags = [ "tags" ]
      resource_groups = [ "resource_groups" ]
      locations = [ "us-south" ]
    }
  }
  parameter {
    agent_assignment_policy_parameter {
      selector_kind = "ids"
      selector_ids = [ "selector_ids" ]
      selector_scope {
        kind = "workspace"
        tags = [ "tags" ]
        resource_groups = [ "resource_groups" ]
        locations = [ "us-south" ]
      }
    }
  }
  scoped_resources {
    kind = "workspace"
    id = "id"
  }
}

// Provision schematics_agent resource instance
resource "ibm_schematics_agent" "schematics_agent_instance" {
  name = var.schematics_agent_name
  resource_group = var.schematics_agent_resource_group
  version = var.schematics_agent_version
  schematics_location = var.schematics_agent_schematics_location
  agent_location = var.schematics_agent_agent_location
  agent_infrastructure {
    infra_type = "ibm_kubernetes"
    cluster_id = "cluster_id"
    cluster_resource_group = "cluster_resource_group"
    cos_instance_name = "cos_instance_name"
    cos_bucket_name = "cos_bucket_name"
    cos_bucket_region = "cos_bucket_region"
  }
  description = var.schematics_agent_description
  tags = var.schematics_agent_tags
  agent_metadata {
    name = "purpose"
    value = ["git", "terraform", "ansible"]
  }
  agent_inputs {
    name = "name"
    value = "value"
    use_default = true
    metadata {
      type = "boolean"
      aliases = [ "aliases" ]
      description = "description"
      cloud_data_type = "cloud_data_type"
      default_value = "default_value"
      link_status = "normal"
      secure = true
      immutable = true
      hidden = true
      required = true
      options = [ "options" ]
      min_value = 1
      max_value = 1
      min_length = 1
      max_length = 1
      matches = "matches"
      position = 1
      group_by = "group_by"
      source = "source"
    }
  }
  user_state {
    state = "enable"
  }
  agent_kpi {
    availability_indicator = "available"
    lifecycle_indicator = "consistent"
    percent_usage_indicator = "percent_usage_indicator"
    application_indicators = [ null ]
    infra_indicators = [ null ]
  }
}

// Provision schematics_agent_prs resource instance
resource "ibm_schematics_agent_prs" "schematics_agent_prs_instance" {
  agent_id = var.schematics_agent_prs_agent_id
  force = var.schematics_agent_prs_force
}

// Provision schematics_agent_deploy resource instance
resource "ibm_schematics_agent_deploy" "schematics_agent_deploy_instance" {
  agent_id = var.schematics_agent_deploy_agent_id
  force = var.schematics_agent_deploy_force
}

// Provision schematics_agent_health resource instance
resource "ibm_schematics_agent_health" "schematics_agent_health_instance" {
  agent_id = var.schematics_agent_health_agent_id
  force = var.schematics_agent_health_force
}

// Create schematics_policies data source
data "ibm_schematics_policies" "schematics_policies_instance" {
  policy_kind = var.schematics_policies_policy_kind
}


// Create schematics_policy data source
data "ibm_schematics_policy" "schematics_policy_instance" {
  policy_id = ibm_schematics_policy.schematics_policy_instance.id
}

// Create schematics_agents data source
data "ibm_schematics_agents" "schematics_agents_instance" {
}
// Create schematics_agent data source
data "ibm_schematics_agent" "schematics_agent_instance" {
  agent_id = ibm_schematics_agent.schematics_agent_instance.id
}
// Create schematics_agent_prs data source
data "ibm_schematics_agent_prs" "schematics_agent_prs_instance" {
  agent_id = ibm_schematics_agent.schematics_agent_instance.id
}
// Create schematics_agent_deploy data source
data "ibm_schematics_agent_deploy" "schematics_agent_deploy_instance" {
  agent_id = ibm_schematics_agent.schematics_agent_instance.id
}
// Create schematics_agent_health data source
data "ibm_schematics_agent_health" "schematics_agent_health_instance" {
  agent_id = ibm_schematics_agent.schematics_agent_instance.id
}