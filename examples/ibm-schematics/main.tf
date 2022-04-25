provider "ibm" {
  ibmcloud_api_key = var.ibmcloud_api_key
}

// Provision schematics_workspace resource instance
resource "ibm_schematics_workspace" "schematics_workspace_instance" {
  applied_shareddata_ids = var.schematics_workspace_applied_shareddata_ids
  catalog_ref {
    dry_run = true
    owning_account = "owning_account"
    item_icon_url = "item_icon_url"
    item_id = "item_id"
    item_name = "item_name"
    item_readme_url = "item_readme_url"
    item_url = "item_url"
    launch_url = "launch_url"
    offering_version = "offering_version"
  }
  dependencies {
    parents = [ "parents" ]
    children = [ "children" ]
  }
  description = var.schematics_workspace_description
  location = var.schematics_workspace_location
  name = var.schematics_workspace_name
  resource_group = var.schematics_workspace_resource_group
  shared_data {
    cluster_created_on = "cluster_created_on"
    cluster_id = "cluster_id"
    cluster_name = "cluster_name"
    cluster_type = "cluster_type"
    entitlement_keys = [ null ]
    namespace = "namespace"
    region = "region"
    resource_group_id = "resource_group_id"
    worker_count = 1
    worker_machine_type = "worker_machine_type"
  }
  tags = var.schematics_workspace_tags
  template_data {
    env_values = [ null ]
    env_values_metadata {
      hidden = true
      name = "name"
      secure = true
    }
    folder = "folder"
    compact = true
    init_state_file = "init_state_file"
    injectors {
      tft_git_url = "tft_git_url"
      tft_git_token = "tft_git_token"
      tft_prefix = "tft_prefix"
      injection_type = "injection_type"
      tft_name = "tft_name"
      tft_parameters {
        name = "name"
        value = "value"
      }
    }
    type = "type"
    uninstall_script_name = "uninstall_script_name"
    values = "values"
    values_metadata = [ null ]
    variablestore {
      description = "description"
      name = "name"
      secure = true
      type = "type"
      use_default = true
      value = "value"
    }
  }
  template_ref = var.schematics_workspace_template_ref
  template_repo {
    branch = "branch"
    release = "release"
    repo_sha_value = "repo_sha_value"
    repo_url = "repo_url"
    url = "url"
  }
  type = var.schematics_workspace_type
  workspace_status {
    frozen = true
    frozen_at = "2021-01-31T09:44:12Z"
    frozen_by = "frozen_by"
    locked = true
    locked_by = "locked_by"
    locked_time = "2021-01-31T09:44:12Z"
  }
  x_github_token = var.schematics_workspace_x_github_token
}

// Provision schematics_action resource instance
resource "ibm_schematics_action" "schematics_action_instance" {
  name = var.schematics_action_name
  description = var.schematics_action_description
  location = var.schematics_action_location
  resource_group = var.schematics_action_resource_group
  bastion_connection_type = var.schematics_action_bastion_connection_type
  inventory_connection_type = var.schematics_action_inventory_connection_type
  tags = var.schematics_action_tags
  user_state {
    state = "draft"
    set_by = "set_by"
    set_at = "2021-01-31T09:44:12Z"
  }
  source_readme_url = var.schematics_action_source_readme_url
  source {
    source_type = "local"
    git {
      computed_git_repo_url = "computed_git_repo_url"
      git_repo_url = "git_repo_url"
      git_token = "git_token"
      git_repo_folder = "git_repo_folder"
      git_release = "git_release"
      git_branch = "git_branch"
    }
    catalog {
      catalog_name = "catalog_name"
      offering_name = "offering_name"
      offering_version = "offering_version"
      offering_kind = "offering_kind"
      offering_id = "offering_id"
      offering_version_id = "offering_version_id"
      offering_repo_url = "offering_repo_url"
    }
  }
  source_type = var.schematics_action_source_type
  command_parameter = var.schematics_action_command_parameter
  inventory = var.schematics_action_inventory
  credentials {
    name = "name"
    value = "value"
    use_default = true
    metadata {
      type = "string"
      aliases = [ "aliases" ]
      description = "description"
      cloud_data_type = "cloud_data_type"
      default_value = "default_value"
      link_status = "normal"
      immutable = true
      hidden = true
      required = true
      position = 1
      group_by = "group_by"
      source = "source"
    }
    link = "link"
  }
  bastion {
    name = "name"
    host = "host"
  }
  bastion_credential {
    name = "name"
    value = "value"
    use_default = true
    metadata {
      type = "string"
      aliases = [ "aliases" ]
      description = "description"
      cloud_data_type = "cloud_data_type"
      default_value = "default_value"
      link_status = "normal"
      immutable = true
      hidden = true
      required = true
      position = 1
      group_by = "group_by"
      source = "source"
    }
    link = "link"
  }
  targets_ini = var.schematics_action_targets_ini
  action_inputs {
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
    link = "link"
  }
  action_outputs {
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
    link = "link"
  }
  settings {
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
    link = "link"
  }
  state {
    status_code = "normal"
    status_job_id = "status_job_id"
    status_message = "status_message"
  }
  sys_lock {
    sys_locked = true
    sys_locked_by = "sys_locked_by"
    sys_locked_at = "2021-01-31T09:44:12Z"
  }
  x_github_token = var.schematics_action_x_github_token
}

// Provision schematics_job resource instance
resource "ibm_schematics_job" "schematics_job_instance" {
  refresh_token = var.schematics_job_refresh_token
  command_object = var.schematics_job_command_object
  command_object_id = var.schematics_job_command_object_id
  command_name = var.schematics_job_command_name
  command_parameter = var.schematics_job_command_parameter
  command_options = var.schematics_job_command_options
  job_inputs {
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
    link = "link"
  }
  job_env_settings {
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
    link = "link"
  }
  tags = var.schematics_job_tags
  location = var.schematics_job_location
  status {
    position_in_queue = 1.0
    total_in_queue = 1.0
    workspace_job_status {
      workspace_name = "workspace_name"
      status_code = "job_pending"
      status_message = "status_message"
      flow_status {
        flow_id = "flow_id"
        flow_name = "flow_name"
        status_code = "job_pending"
        status_message = "status_message"
        workitems {
          workspace_id = "workspace_id"
          workspace_name = "workspace_name"
          job_id = "job_id"
          status_code = "job_pending"
          status_message = "status_message"
          updated_at = "2021-01-31T09:44:12Z"
        }
        updated_at = "2021-01-31T09:44:12Z"
      }
      template_status {
        template_id = "template_id"
        template_name = "template_name"
        flow_index = 1
        status_code = "job_pending"
        status_message = "status_message"
        updated_at = "2021-01-31T09:44:12Z"
      }
      updated_at = "2021-01-31T09:44:12Z"
      commands {
        name = "name"
        outcome = "outcome"
      }
    }
    action_job_status {
      action_name = "action_name"
      status_code = "job_pending"
      status_message = "status_message"
      bastion_status_code = "none"
      bastion_status_message = "bastion_status_message"
      targets_status_code = "none"
      targets_status_message = "targets_status_message"
      updated_at = "2021-01-31T09:44:12Z"
    }
    system_job_status {
      system_status_message = "system_status_message"
      system_status_code = "job_pending"
      schematics_resource_status {
        status_code = "job_pending"
        status_message = "status_message"
        schematics_resource_id = "schematics_resource_id"
        updated_at = "2021-01-31T09:44:12Z"
      }
      updated_at = "2021-01-31T09:44:12Z"
    }
    flow_job_status {
      flow_id = "flow_id"
      flow_name = "flow_name"
      status_code = "job_pending"
      status_message = "status_message"
      workitems {
        workspace_id = "workspace_id"
        workspace_name = "workspace_name"
        job_id = "job_id"
        status_code = "job_pending"
        status_message = "status_message"
        updated_at = "2021-01-31T09:44:12Z"
      }
      updated_at = "2021-01-31T09:44:12Z"
    }
  }
  data {
    job_type = "repo_download_job"
    workspace_job_data {
      workspace_name = "workspace_name"
      flow_id = "flow_id"
      flow_name = "flow_name"
      inputs {
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
        link = "link"
      }
      outputs {
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
        link = "link"
      }
      settings {
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
        link = "link"
      }
      template_data {
        template_id = "template_id"
        template_name = "template_name"
        flow_index = 1
        inputs {
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
          link = "link"
        }
        outputs {
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
          link = "link"
        }
        settings {
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
          link = "link"
        }
        updated_at = "2021-01-31T09:44:12Z"
      }
      updated_at = "2021-01-31T09:44:12Z"
    }
    action_job_data {
      action_name = "action_name"
      inputs {
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
        link = "link"
      }
      outputs {
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
        link = "link"
      }
      settings {
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
        link = "link"
      }
      updated_at = "2021-01-31T09:44:12Z"
      inventory_record {
        name = "name"
        id = "id"
        description = "description"
        location = "us-south"
        resource_group = "resource_group"
        created_at = "2021-01-31T09:44:12Z"
        created_by = "created_by"
        updated_at = "2021-01-31T09:44:12Z"
        updated_by = "updated_by"
        inventories_ini = "inventories_ini"
        resource_queries = [ "resource_queries" ]
      }
      materialized_inventory = "materialized_inventory"
    }
    system_job_data {
      key_id = "key_id"
      schematics_resource_id = [ "schematics_resource_id" ]
      updated_at = "2021-01-31T09:44:12Z"
    }
    flow_job_data {
      flow_id = "flow_id"
      flow_name = "flow_name"
      workitems {
        command_object_id = "command_object_id"
        command_object_name = "command_object_name"
        layers = "layers"
        source_type = "local"
        source {
          source_type = "local"
          git {
            computed_git_repo_url = "computed_git_repo_url"
            git_repo_url = "git_repo_url"
            git_token = "git_token"
            git_repo_folder = "git_repo_folder"
            git_release = "git_release"
            git_branch = "git_branch"
          }
          catalog {
            catalog_name = "catalog_name"
            offering_name = "offering_name"
            offering_version = "offering_version"
            offering_kind = "offering_kind"
            offering_id = "offering_id"
            offering_version_id = "offering_version_id"
            offering_repo_url = "offering_repo_url"
          }
        }
        inputs {
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
          link = "link"
        }
        outputs {
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
          link = "link"
        }
        settings {
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
          link = "link"
        }
        last_job {
          command_object = "workspace"
          command_object_name = "command_object_name"
          command_object_id = "command_object_id"
          command_name = "workspace_plan"
          job_id = "job_id"
          job_status = "job_pending"
        }
        updated_at = "2021-01-31T09:44:12Z"
      }
      updated_at = "2021-01-31T09:44:12Z"
    }
  }
  bastion {
    name = "name"
    host = "host"
  }
  log_summary {
    job_id = "job_id"
    job_type = "repo_download_job"
    log_start_at = "2021-01-31T09:44:12Z"
    log_analyzed_till = "2021-01-31T09:44:12Z"
    elapsed_time = 1.0
    log_errors {
      error_code = "error_code"
      error_msg = "error_msg"
      error_count = 1.0
    }
    repo_download_job {
      scanned_file_count = 1.0
      quarantined_file_count = 1.0
      detected_filetype = "detected_filetype"
      inputs_count = "inputs_count"
      outputs_count = "outputs_count"
    }
    workspace_job {
      resources_add = 1.0
      resources_modify = 1.0
      resources_destroy = 1.0
    }
    flow_job {
      workitems_completed = 1.0
      workitems_pending = 1.0
      workitems_failed = 1.0
      workitems {
        workspace_id = "workspace_id"
        job_id = "job_id"
        resources_add = 1.0
        resources_modify = 1.0
        resources_destroy = 1.0
        log_url = "log_url"
      }
    }
    action_job {
      target_count = 1.0
      task_count = 1.0
      play_count = 1.0
      recap {
        target = [ "target" ]
        ok = 1.0
        changed = 1.0
        failed = 1.0
        skipped = 1.0
        unreachable = 1.0
      }
    }
    system_job {
      target_count = 1.0
      success = 1.0
      failed = 1.0
    }
  }
}

// Provision schematics_inventory resource instance
resource "ibm_schematics_inventory" "schematics_inventory_instance" {
  name = var.schematics_inventory_name
  description = var.schematics_inventory_description
  location = var.schematics_inventory_location
  resource_group = var.schematics_inventory_resource_group
  inventories_ini = var.schematics_inventory_inventories_ini
  resource_queries = var.schematics_inventory_resource_queries
}

// Provision schematics_resource_query resource instance
resource "ibm_schematics_resource_query" "schematics_resource_query_instance" {
  type = var.schematics_resource_query_type
  name = var.schematics_resource_query_name
  queries {
    query_type = "workspaces"
    query_condition {
      name = "name"
      value = "value"
      description = "description"
    }
    query_select = [ "query_select" ]
  }
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

// Create schematics_inventory data source
data "ibm_schematics_inventory" "schematics_inventory_instance" {
  inventory_id = var.schematics_inventory_inventory_id
}

// Create schematics_resource_query data source
data "ibm_schematics_resource_query" "schematics_resource_query_instance" {
  query_id = var.schematics_resource_query_query_id
}
