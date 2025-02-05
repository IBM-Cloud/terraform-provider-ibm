provider "ibm" {
  ibmcloud_api_key = var.ibmcloud_api_key
}

// Provision backup_recovery_connection_registration_token resource instance
resource "ibm_backup_recovery_connection_registration_token" "backup_recovery_connection_registration_token_instance" {
  connection_id = var.backup_recovery_connection_registration_token_connection_id
  x_ibm_tenant_id = var.backup_recovery_connection_registration_token_x_ibm_tenant_id
}

// Provision backup_recovery_protection_group_run_request resource instance
resource "ibm_backup_recovery_protection_group_run_request" "backup_recovery_protection_group_run_request_instance" {
  x_ibm_tenant_id = var.backup_recovery_protection_group_run_request_x_ibm_tenant_id
  run_type = var.backup_recovery_protection_group_run_request_run_type
  objects {
    id = 1
    app_ids = [ 1 ]
    physical_params {
      metadata_file_path = "metadata_file_path"
    }
  }
  targets_config {
    use_policy_defaults = true
    replications {
      id = 1
      retention {
        unit = "Days"
        duration = 1
        data_lock_config {
          mode = "Compliance"
          unit = "Days"
          duration = 1
          enable_worm_on_external_target = true
        }
      }
    }
    archivals {
      id = 1
      archival_target_type = "Tape"
      retention {
        unit = "Days"
        duration = 1
        data_lock_config {
          mode = "Compliance"
          unit = "Days"
          duration = 1
          enable_worm_on_external_target = true
        }
      }
      copy_only_fully_successful = true
    }
    cloud_replications {
      aws_target {
        region = 1
        source_id = 1
      }
      azure_target {
        resource_group = 1
        source_id = 1
      }
      target_type = "AWS"
      retention {
        unit = "Days"
        duration = 1
        data_lock_config {
          mode = "Compliance"
          unit = "Days"
          duration = 1
          enable_worm_on_external_target = true
        }
      }
    }
  }
}

// Provision backup_recovery_data_source_connection resource instance
resource "ibm_backup_recovery_data_source_connection" "backup_recovery_data_source_connection_instance" {
  x_ibm_tenant_id = var.backup_recovery_data_source_connection_x_ibm_tenant_id
  connection_name = var.backup_recovery_data_source_connection_connection_name
}

// Provision backup_recovery_download_files_folders resource instance
resource "ibm_backup_recovery_download_files_folders" "backup_recovery_download_files_folders_instance" {
  x_ibm_tenant_id = var.backup_recovery_download_files_folders_x_ibm_tenant_id
  documents {
    is_directory = true
    item_id = "item_id"
  }
  name = var.backup_recovery_download_files_folders_name
  object {
    snapshot_id = "snapshot_id"
    point_in_time_usecs = 1
    protection_group_id = "protection_group_id"
    protection_group_name = "protection_group_name"
    object_info {
      id = 1
      name = "name"
      source_id = 1
      source_name = "source_name"
      environment = "kPhysical"
      object_hash = "object_hash"
      object_type = "kCluster"
      logical_size_bytes = 1
      uuid = "uuid"
      global_id = "global_id"
      protection_type = "kAgent"
      sharepoint_site_summary {
        site_web_url = "site_web_url"
      }
      os_type = "kLinux"
      child_objects {
        id = 1
        name = "name"
        source_id = 1
        source_name = "source_name"
        environment = "kPhysical"
        object_hash = "object_hash"
        object_type = "kCluster"
        logical_size_bytes = 1
        uuid = "uuid"
        global_id = "global_id"
        protection_type = "kAgent"
        sharepoint_site_summary {
          site_web_url = "site_web_url"
        }
        os_type = "kLinux"
        v_center_summary {
          is_cloud_env = true
        }
        windows_cluster_summary {
          cluster_source_type = "cluster_source_type"
        }
      }
      v_center_summary {
        is_cloud_env = true
      }
      windows_cluster_summary {
        cluster_source_type = "cluster_source_type"
      }
    }
    archival_target_info {
      target_id = 1
      archival_task_id = "archival_task_id"
      target_name = "target_name"
      target_type = "Tape"
      usage_type = "Archival"
      ownership_context = "Local"
      tier_settings {
        aws_tiering {
          tiers {
            move_after_unit = "Days"
            move_after = 1
            tier_type = "kAmazonS3Standard"
          }
        }
        azure_tiering {
          tiers {
            move_after_unit = "Days"
            move_after = 1
            tier_type = "kAzureTierHot"
          }
        }
        cloud_platform = "AWS"
        google_tiering {
          tiers {
            move_after_unit = "Days"
            move_after = 1
            tier_type = "kGoogleStandard"
          }
        }
        oracle_tiering {
          tiers {
            move_after_unit = "Days"
            move_after = 1
            tier_type = "kOracleTierStandard"
          }
        }
        current_tier_type = "kAmazonS3Standard"
      }
    }
    recover_from_standby = true
  }
  parent_recovery_id = var.backup_recovery_download_files_folders_parent_recovery_id
  files_and_folders {
    absolute_path = "absolute_path"
    is_directory = true
  }
  glacier_retrieval_type = var.backup_recovery_download_files_folders_glacier_retrieval_type
}

// Provision backup_recovery_restore_points resource instance
resource "ibm_backup_recovery_restore_points" "backup_recovery_restore_points_instance" {
  x_ibm_tenant_id = var.backup_recovery_restore_points_x_ibm_tenant_id
  end_time_usecs = var.backup_recovery_restore_points_end_time_usecs
  environment = var.backup_recovery_restore_points_environment
  protection_group_ids = var.backup_recovery_restore_points_protection_group_ids
  source_id = var.backup_recovery_restore_points_source_id
  start_time_usecs = var.backup_recovery_restore_points_start_time_usecs
}

// Provision backup_recovery_perform_action_on_protection_group_run_request resource instance
resource "ibm_backup_recovery_perform_action_on_protection_group_run_request" "backup_recovery_perform_action_on_protection_group_run_request_instance" {
  x_ibm_tenant_id = var.backup_recovery_perform_action_on_protection_group_run_request_x_ibm_tenant_id
  action = var.backup_recovery_perform_action_on_protection_group_run_request_action
  pause_params {
    run_id = "run_id"
  }
  resume_params {
    run_id = "run_id"
  }
  cancel_params {
    run_id = "run_id"
    local_task_id = "local_task_id"
    object_ids = [ 1 ]
    replication_task_id = [ "replicationTaskId" ]
    archival_task_id = [ "archivalTaskId" ]
    cloud_spin_task_id = [ "cloudSpinTaskId" ]
  }
}

// Provision backup_recovery_protection_group resource instance
resource "ibm_backup_recovery_protection_group" "backup_recovery_protection_group_instance" {
  x_ibm_tenant_id = var.backup_recovery_protection_group_x_ibm_tenant_id
  name = var.backup_recovery_protection_group_name
  policy_id = var.backup_recovery_protection_group_policy_id
  priority = var.backup_recovery_protection_group_priority
  description = var.backup_recovery_protection_group_description
  start_time {
    hour = 0
    minute = 0
    time_zone = "time_zone"
  }
  end_time_usecs = var.backup_recovery_protection_group_end_time_usecs
  last_modified_timestamp_usecs = var.backup_recovery_protection_group_last_modified_timestamp_usecs
  alert_policy {
    backup_run_status = [ "kSuccess" ]
    alert_targets {
      email_address = "email_address"
      language = "en-us"
      recipient_type = "kTo"
    }
    raise_object_level_failure_alert = true
    raise_object_level_failure_alert_after_last_attempt = true
    raise_object_level_failure_alert_after_each_attempt = true
  }
  sla {
    backup_run_type = "kIncremental"
    sla_minutes = 1
  }
  qos_policy = var.backup_recovery_protection_group_qos_policy
  abort_in_blackouts = var.backup_recovery_protection_group_abort_in_blackouts
  pause_in_blackouts = var.backup_recovery_protection_group_pause_in_blackouts
  is_paused = var.backup_recovery_protection_group_is_paused
  environment = var.backup_recovery_protection_group_environment
  advanced_configs {
    key = "key"
    value = "value"
  }
  physical_params {
    protection_type = "kFile"
    volume_protection_type_params {
      objects {
        id = 1
        name = "name"
        volume_guids = [ "volumeGuids" ]
        enable_system_backup = true
        excluded_vss_writers = [ "excludedVssWriters" ]
      }
      indexing_policy {
        enable_indexing = true
        include_paths = [ "includePaths" ]
        exclude_paths = [ "excludePaths" ]
      }
      perform_source_side_deduplication = true
      quiesce = true
      continue_on_quiesce_failure = true
      incremental_backup_after_restart = true
      pre_post_script {
        pre_script {
          path = "path"
          params = "params"
          timeout_secs = 1
          is_active = true
          continue_on_error = true
        }
        post_script {
          path = "path"
          params = "params"
          timeout_secs = 1
          is_active = true
        }
      }
      dedup_exclusion_source_ids = [ 1 ]
      excluded_vss_writers = [ "excludedVssWriters" ]
      cobmr_backup = true
    }
    file_protection_type_params {
      excluded_vss_writers = [ "excludedVssWriters" ]
      objects {
        excluded_vss_writers = [ "excludedVssWriters" ]
        id = 1
        file_paths {
          included_path = "included_path"
          excluded_paths = [ "excludedPaths" ]
          skip_nested_volumes = true
        }
        uses_path_level_skip_nested_volume_setting = true
        nested_volume_types_to_skip = [ "nestedVolumeTypesToSkip" ]
        follow_nas_symlink_target = true
        metadata_file_path = "metadata_file_path"
      }
      indexing_policy {
        enable_indexing = true
        include_paths = [ "includePaths" ]
        exclude_paths = [ "excludePaths" ]
      }
      perform_source_side_deduplication = true
      perform_brick_based_deduplication = true
      task_timeouts {
        timeout_mins = 1
        backup_type = "kRegular"
      }
      quiesce = true
      continue_on_quiesce_failure = true
      cobmr_backup = true
      pre_post_script {
        pre_script {
          path = "path"
          params = "params"
          timeout_secs = 1
          is_active = true
          continue_on_error = true
        }
        post_script {
          path = "path"
          params = "params"
          timeout_secs = 1
          is_active = true
        }
      }
      dedup_exclusion_source_ids = [ 1 ]
      global_exclude_paths = [ "globalExcludePaths" ]
      global_exclude_fs = [ "globalExcludeFS" ]
      ignorable_errors = [ "kEOF" ]
      allow_parallel_runs = true
    }
  }
  mssql_params {
    file_protection_type_params {
      aag_backup_preference_type = "kPrimaryReplicaOnly"
      advanced_settings {
        cloned_db_backup_status = "kError"
        db_backup_if_not_online_status = "kError"
        missing_db_backup_status = "kError"
        offline_restoring_db_backup_status = "kError"
        read_only_db_backup_status = "kError"
        report_all_non_autoprotect_db_errors = "kError"
      }
      backup_system_dbs = true
      exclude_filters {
        filter_string = "filter_string"
        is_regular_expression = true
      }
      full_backups_copy_only = true
      log_backup_num_streams = 1
      log_backup_with_clause = "log_backup_with_clause"
      pre_post_script {
        pre_script {
          path = "path"
          params = "params"
          timeout_secs = 1
          is_active = true
          continue_on_error = true
        }
        post_script {
          path = "path"
          params = "params"
          timeout_secs = 1
          is_active = true
        }
      }
      use_aag_preferences_from_server = true
      user_db_backup_preference_type = "kBackupAllDatabases"
      additional_host_params {
        disable_source_side_deduplication = true
        host_id = 1
      }
      objects {
        id = 1
      }
      perform_source_side_deduplication = true
    }
    native_protection_type_params {
      aag_backup_preference_type = "kPrimaryReplicaOnly"
      advanced_settings {
        cloned_db_backup_status = "kError"
        db_backup_if_not_online_status = "kError"
        missing_db_backup_status = "kError"
        offline_restoring_db_backup_status = "kError"
        read_only_db_backup_status = "kError"
        report_all_non_autoprotect_db_errors = "kError"
      }
      backup_system_dbs = true
      exclude_filters {
        filter_string = "filter_string"
        is_regular_expression = true
      }
      full_backups_copy_only = true
      log_backup_num_streams = 1
      log_backup_with_clause = "log_backup_with_clause"
      pre_post_script {
        pre_script {
          path = "path"
          params = "params"
          timeout_secs = 1
          is_active = true
          continue_on_error = true
        }
        post_script {
          path = "path"
          params = "params"
          timeout_secs = 1
          is_active = true
        }
      }
      use_aag_preferences_from_server = true
      user_db_backup_preference_type = "kBackupAllDatabases"
      num_streams = 1
      objects {
        id = 1
      }
      with_clause = "with_clause"
    }
    protection_type = "kFile"
    volume_protection_type_params {
      aag_backup_preference_type = "kPrimaryReplicaOnly"
      advanced_settings {
        cloned_db_backup_status = "kError"
        db_backup_if_not_online_status = "kError"
        missing_db_backup_status = "kError"
        offline_restoring_db_backup_status = "kError"
        read_only_db_backup_status = "kError"
        report_all_non_autoprotect_db_errors = "kError"
      }
      backup_system_dbs = true
      exclude_filters {
        filter_string = "filter_string"
        is_regular_expression = true
      }
      full_backups_copy_only = true
      log_backup_num_streams = 1
      log_backup_with_clause = "log_backup_with_clause"
      pre_post_script {
        pre_script {
          path = "path"
          params = "params"
          timeout_secs = 1
          is_active = true
          continue_on_error = true
        }
        post_script {
          path = "path"
          params = "params"
          timeout_secs = 1
          is_active = true
        }
      }
      use_aag_preferences_from_server = true
      user_db_backup_preference_type = "kBackupAllDatabases"
      additional_host_params {
        enable_system_backup = true
        host_id = 1
        volume_guids = [ "volumeGuids" ]
      }
      backup_db_volumes_only = true
      incremental_backup_after_restart = true
      indexing_policy {
        enable_indexing = true
        include_paths = [ "includePaths" ]
        exclude_paths = [ "excludePaths" ]
      }
      objects {
        id = 1
      }
    }
  }
}

// Provision backup_recovery_protection_policy resource instance
resource "ibm_backup_recovery_protection_policy" "backup_recovery_protection_policy_instance" {
  x_ibm_tenant_id = var.backup_recovery_protection_policy_x_ibm_tenant_id
  name = var.backup_recovery_protection_policy_name
  backup_policy {
    regular {
      incremental {
        schedule {
          unit = "Minutes"
          minute_schedule {
            frequency = 1
          }
          hour_schedule {
            frequency = 1
          }
          day_schedule {
            frequency = 1
          }
          week_schedule {
            day_of_week = [ "Sunday" ]
          }
          month_schedule {
            day_of_week = [ "Sunday" ]
            week_of_month = "First"
            day_of_month = 1
          }
          year_schedule {
            day_of_year = "First"
          }
        }
      }
      full {
        schedule {
          unit = "Days"
          day_schedule {
            frequency = 1
          }
          week_schedule {
            day_of_week = [ "Sunday" ]
          }
          month_schedule {
            day_of_week = [ "Sunday" ]
            week_of_month = "First"
            day_of_month = 1
          }
          year_schedule {
            day_of_year = "First"
          }
        }
      }
      full_backups {
        schedule {
          unit = "Days"
          day_schedule {
            frequency = 1
          }
          week_schedule {
            day_of_week = [ "Sunday" ]
          }
          month_schedule {
            day_of_week = [ "Sunday" ]
            week_of_month = "First"
            day_of_month = 1
          }
          year_schedule {
            day_of_year = "First"
          }
        }
        retention {
          unit = "Days"
          duration = 1
          data_lock_config {
            mode = "Compliance"
            unit = "Days"
            duration = 1
            enable_worm_on_external_target = true
          }
        }
      }
      retention {
        unit = "Days"
        duration = 1
        data_lock_config {
          mode = "Compliance"
          unit = "Days"
          duration = 1
          enable_worm_on_external_target = true
        }
      }
      primary_backup_target {
        target_type = "Local"
        archival_target_settings {
          target_id = 1
          tier_settings {
            aws_tiering {
              tiers {
                move_after_unit = "Days"
                move_after = 1
                tier_type = "kAmazonS3Standard"
              }
            }
            azure_tiering {
              tiers {
                move_after_unit = "Days"
                move_after = 1
                tier_type = "kAzureTierHot"
              }
            }
            cloud_platform = "AWS"
            google_tiering {
              tiers {
                move_after_unit = "Days"
                move_after = 1
                tier_type = "kGoogleStandard"
              }
            }
            oracle_tiering {
              tiers {
                move_after_unit = "Days"
                move_after = 1
                tier_type = "kOracleTierStandard"
              }
            }
          }
        }
        use_default_backup_target = true
      }
    }
    log {
      schedule {
        unit = "Minutes"
        minute_schedule {
          frequency = 1
        }
        hour_schedule {
          frequency = 1
        }
      }
      retention {
        unit = "Days"
        duration = 1
        data_lock_config {
          mode = "Compliance"
          unit = "Days"
          duration = 1
          enable_worm_on_external_target = true
        }
      }
    }
    bmr {
      schedule {
        unit = "Days"
        day_schedule {
          frequency = 1
        }
        week_schedule {
          day_of_week = [ "Sunday" ]
        }
        month_schedule {
          day_of_week = [ "Sunday" ]
          week_of_month = "First"
          day_of_month = 1
        }
        year_schedule {
          day_of_year = "First"
        }
      }
      retention {
        unit = "Days"
        duration = 1
        data_lock_config {
          mode = "Compliance"
          unit = "Days"
          duration = 1
          enable_worm_on_external_target = true
        }
      }
    }
    cdp {
      retention {
        unit = "Minutes"
        duration = 1
        data_lock_config {
          mode = "Compliance"
          unit = "Days"
          duration = 1
          enable_worm_on_external_target = true
        }
      }
    }
    storage_array_snapshot {
      schedule {
        unit = "Minutes"
        minute_schedule {
          frequency = 1
        }
        hour_schedule {
          frequency = 1
        }
        day_schedule {
          frequency = 1
        }
        week_schedule {
          day_of_week = [ "Sunday" ]
        }
        month_schedule {
          day_of_week = [ "Sunday" ]
          week_of_month = "First"
          day_of_month = 1
        }
        year_schedule {
          day_of_year = "First"
        }
      }
      retention {
        unit = "Days"
        duration = 1
        data_lock_config {
          mode = "Compliance"
          unit = "Days"
          duration = 1
          enable_worm_on_external_target = true
        }
      }
    }
    run_timeouts {
      timeout_mins = 1
      backup_type = "kRegular"
    }
  }
  description = var.backup_recovery_protection_policy_description
  blackout_window {
    day = "Sunday"
    start_time {
      hour = 0
      minute = 0
      time_zone = "time_zone"
    }
    end_time {
      hour = 0
      minute = 0
      time_zone = "time_zone"
    }
    config_id = "config_id"
  }
  extended_retention {
    schedule {
      unit = "Runs"
      frequency = 1
    }
    retention {
      unit = "Days"
      duration = 1
      data_lock_config {
        mode = "Compliance"
        unit = "Days"
        duration = 1
        enable_worm_on_external_target = true
      }
    }
    run_type = "Regular"
    config_id = "config_id"
  }
  remote_target_policy {
    replication_targets {
      schedule {
        unit = "Runs"
        frequency = 1
      }
      retention {
        unit = "Days"
        duration = 1
        data_lock_config {
          mode = "Compliance"
          unit = "Days"
          duration = 1
          enable_worm_on_external_target = true
        }
      }
      copy_on_run_success = true
      config_id = "config_id"
      backup_run_type = "Regular"
      run_timeouts {
        timeout_mins = 1
        backup_type = "kRegular"
      }
      log_retention {
        unit = "Days"
        duration = 0
        data_lock_config {
          mode = "Compliance"
          unit = "Days"
          duration = 1
          enable_worm_on_external_target = true
        }
      }
      aws_target_config {
        region = 1
        source_id = 1
      }
      azure_target_config {
        resource_group = 1
        source_id = 1
      }
      target_type = "RemoteCluster"
      remote_target_config {
        cluster_id = 1
      }
    }
    archival_targets {
      schedule {
        unit = "Runs"
        frequency = 1
      }
      retention {
        unit = "Days"
        duration = 1
        data_lock_config {
          mode = "Compliance"
          unit = "Days"
          duration = 1
          enable_worm_on_external_target = true
        }
      }
      copy_on_run_success = true
      config_id = "config_id"
      backup_run_type = "Regular"
      run_timeouts {
        timeout_mins = 1
        backup_type = "kRegular"
      }
      log_retention {
        unit = "Days"
        duration = 0
        data_lock_config {
          mode = "Compliance"
          unit = "Days"
          duration = 1
          enable_worm_on_external_target = true
        }
      }
      target_id = 1
      tier_settings {
        aws_tiering {
          tiers {
            move_after_unit = "Days"
            move_after = 1
            tier_type = "kAmazonS3Standard"
          }
        }
        azure_tiering {
          tiers {
            move_after_unit = "Days"
            move_after = 1
            tier_type = "kAzureTierHot"
          }
        }
        cloud_platform = "AWS"
        google_tiering {
          tiers {
            move_after_unit = "Days"
            move_after = 1
            tier_type = "kGoogleStandard"
          }
        }
        oracle_tiering {
          tiers {
            move_after_unit = "Days"
            move_after = 1
            tier_type = "kOracleTierStandard"
          }
        }
      }
      extended_retention {
        schedule {
          unit = "Runs"
          frequency = 1
        }
        retention {
          unit = "Days"
          duration = 1
          data_lock_config {
            mode = "Compliance"
            unit = "Days"
            duration = 1
            enable_worm_on_external_target = true
          }
        }
        run_type = "Regular"
        config_id = "config_id"
      }
    }
    cloud_spin_targets {
      schedule {
        unit = "Runs"
        frequency = 1
      }
      retention {
        unit = "Days"
        duration = 1
        data_lock_config {
          mode = "Compliance"
          unit = "Days"
          duration = 1
          enable_worm_on_external_target = true
        }
      }
      copy_on_run_success = true
      config_id = "config_id"
      backup_run_type = "Regular"
      run_timeouts {
        timeout_mins = 1
        backup_type = "kRegular"
      }
      log_retention {
        unit = "Days"
        duration = 0
        data_lock_config {
          mode = "Compliance"
          unit = "Days"
          duration = 1
          enable_worm_on_external_target = true
        }
      }
      target {
        aws_params {
          custom_tag_list {
            key = "key"
            value = "value"
          }
          region = 1
          subnet_id = 1
          vpc_id = 1
        }
        azure_params {
          availability_set_id = 1
          network_resource_group_id = 1
          resource_group_id = 1
          storage_account_id = 1
          storage_container_id = 1
          storage_resource_group_id = 1
          temp_vm_resource_group_id = 1
          temp_vm_storage_account_id = 1
          temp_vm_storage_container_id = 1
          temp_vm_subnet_id = 1
          temp_vm_virtual_network_id = 1
        }
        id = 1
      }
    }
    onprem_deploy_targets {
      schedule {
        unit = "Runs"
        frequency = 1
      }
      retention {
        unit = "Days"
        duration = 1
        data_lock_config {
          mode = "Compliance"
          unit = "Days"
          duration = 1
          enable_worm_on_external_target = true
        }
      }
      copy_on_run_success = true
      config_id = "config_id"
      backup_run_type = "Regular"
      run_timeouts {
        timeout_mins = 1
        backup_type = "kRegular"
      }
      log_retention {
        unit = "Days"
        duration = 0
        data_lock_config {
          mode = "Compliance"
          unit = "Days"
          duration = 1
          enable_worm_on_external_target = true
        }
      }
      params {
        id = 1
      }
    }
    rpaas_targets {
      schedule {
        unit = "Runs"
        frequency = 1
      }
      retention {
        unit = "Days"
        duration = 1
        data_lock_config {
          mode = "Compliance"
          unit = "Days"
          duration = 1
          enable_worm_on_external_target = true
        }
      }
      copy_on_run_success = true
      config_id = "config_id"
      backup_run_type = "Regular"
      run_timeouts {
        timeout_mins = 1
        backup_type = "kRegular"
      }
      log_retention {
        unit = "Days"
        duration = 0
        data_lock_config {
          mode = "Compliance"
          unit = "Days"
          duration = 1
          enable_worm_on_external_target = true
        }
      }
      target_id = 1
      target_type = "Tape"
    }
  }
  cascaded_targets_config {
    source_cluster_id = 1
    remote_targets {
      replication_targets {
        schedule {
          unit = "Runs"
          frequency = 1
        }
        retention {
          unit = "Days"
          duration = 1
          data_lock_config {
            mode = "Compliance"
            unit = "Days"
            duration = 1
            enable_worm_on_external_target = true
          }
        }
        copy_on_run_success = true
        config_id = "config_id"
        backup_run_type = "Regular"
        run_timeouts {
          timeout_mins = 1
          backup_type = "kRegular"
        }
        log_retention {
          unit = "Days"
          duration = 0
          data_lock_config {
            mode = "Compliance"
            unit = "Days"
            duration = 1
            enable_worm_on_external_target = true
          }
        }
        aws_target_config {
          region = 1
          source_id = 1
        }
        azure_target_config {
          resource_group = 1
          source_id = 1
        }
        target_type = "RemoteCluster"
        remote_target_config {
          cluster_id = 1
        }
      }
      archival_targets {
        schedule {
          unit = "Runs"
          frequency = 1
        }
        retention {
          unit = "Days"
          duration = 1
          data_lock_config {
            mode = "Compliance"
            unit = "Days"
            duration = 1
            enable_worm_on_external_target = true
          }
        }
        copy_on_run_success = true
        config_id = "config_id"
        backup_run_type = "Regular"
        run_timeouts {
          timeout_mins = 1
          backup_type = "kRegular"
        }
        log_retention {
          unit = "Days"
          duration = 0
          data_lock_config {
            mode = "Compliance"
            unit = "Days"
            duration = 1
            enable_worm_on_external_target = true
          }
        }
        target_id = 1
        tier_settings {
          aws_tiering {
            tiers {
              move_after_unit = "Days"
              move_after = 1
              tier_type = "kAmazonS3Standard"
            }
          }
          azure_tiering {
            tiers {
              move_after_unit = "Days"
              move_after = 1
              tier_type = "kAzureTierHot"
            }
          }
          cloud_platform = "AWS"
          google_tiering {
            tiers {
              move_after_unit = "Days"
              move_after = 1
              tier_type = "kGoogleStandard"
            }
          }
          oracle_tiering {
            tiers {
              move_after_unit = "Days"
              move_after = 1
              tier_type = "kOracleTierStandard"
            }
          }
        }
        extended_retention {
          schedule {
            unit = "Runs"
            frequency = 1
          }
          retention {
            unit = "Days"
            duration = 1
            data_lock_config {
              mode = "Compliance"
              unit = "Days"
              duration = 1
              enable_worm_on_external_target = true
            }
          }
          run_type = "Regular"
          config_id = "config_id"
        }
      }
      cloud_spin_targets {
        schedule {
          unit = "Runs"
          frequency = 1
        }
        retention {
          unit = "Days"
          duration = 1
          data_lock_config {
            mode = "Compliance"
            unit = "Days"
            duration = 1
            enable_worm_on_external_target = true
          }
        }
        copy_on_run_success = true
        config_id = "config_id"
        backup_run_type = "Regular"
        run_timeouts {
          timeout_mins = 1
          backup_type = "kRegular"
        }
        log_retention {
          unit = "Days"
          duration = 0
          data_lock_config {
            mode = "Compliance"
            unit = "Days"
            duration = 1
            enable_worm_on_external_target = true
          }
        }
        target {
          aws_params {
            custom_tag_list {
              key = "key"
              value = "value"
            }
            region = 1
            subnet_id = 1
            vpc_id = 1
          }
          azure_params {
            availability_set_id = 1
            network_resource_group_id = 1
            resource_group_id = 1
            storage_account_id = 1
            storage_container_id = 1
            storage_resource_group_id = 1
            temp_vm_resource_group_id = 1
            temp_vm_storage_account_id = 1
            temp_vm_storage_container_id = 1
            temp_vm_subnet_id = 1
            temp_vm_virtual_network_id = 1
          }
          id = 1
        }
      }
      onprem_deploy_targets {
        schedule {
          unit = "Runs"
          frequency = 1
        }
        retention {
          unit = "Days"
          duration = 1
          data_lock_config {
            mode = "Compliance"
            unit = "Days"
            duration = 1
            enable_worm_on_external_target = true
          }
        }
        copy_on_run_success = true
        config_id = "config_id"
        backup_run_type = "Regular"
        run_timeouts {
          timeout_mins = 1
          backup_type = "kRegular"
        }
        log_retention {
          unit = "Days"
          duration = 0
          data_lock_config {
            mode = "Compliance"
            unit = "Days"
            duration = 1
            enable_worm_on_external_target = true
          }
        }
        params {
          id = 1
        }
      }
      rpaas_targets {
        schedule {
          unit = "Runs"
          frequency = 1
        }
        retention {
          unit = "Days"
          duration = 1
          data_lock_config {
            mode = "Compliance"
            unit = "Days"
            duration = 1
            enable_worm_on_external_target = true
          }
        }
        copy_on_run_success = true
        config_id = "config_id"
        backup_run_type = "Regular"
        run_timeouts {
          timeout_mins = 1
          backup_type = "kRegular"
        }
        log_retention {
          unit = "Days"
          duration = 0
          data_lock_config {
            mode = "Compliance"
            unit = "Days"
            duration = 1
            enable_worm_on_external_target = true
          }
        }
        target_id = 1
        target_type = "Tape"
      }
    }
  }
  retry_options {
    retries = 0
    retry_interval_mins = 1
  }
  data_lock = var.backup_recovery_protection_policy_data_lock
  version = var.backup_recovery_protection_policy_version
  is_cbs_enabled = var.backup_recovery_protection_policy_is_cbs_enabled
  last_modification_time_usecs = var.backup_recovery_protection_policy_last_modification_time_usecs
  template_id = var.backup_recovery_protection_policy_template_id
}

// Provision backup_recovery resource instance
resource "ibm_backup_recovery" "backup_recovery_instance" {
  x_ibm_tenant_id = var.backup_recovery_x_ibm_tenant_id
  request_initiator_type = var.backup_recovery_request_initiator_type
  name = var.backup_recovery_name
  snapshot_environment = var.backup_recovery_snapshot_environment
  physical_params {
    objects {
      snapshot_id = "snapshot_id"
      point_in_time_usecs = 1
      protection_group_id = "protection_group_id"
      protection_group_name = "protection_group_name"
      object_info {
        id = 1
        name = "name"
        source_id = 1
        source_name = "source_name"
        environment = "kPhysical"
        object_hash = "object_hash"
        object_type = "kCluster"
        logical_size_bytes = 1
        uuid = "uuid"
        global_id = "global_id"
        protection_type = "kAgent"
        sharepoint_site_summary {
          site_web_url = "site_web_url"
        }
        os_type = "kLinux"
        child_objects {
          id = 1
          name = "name"
          source_id = 1
          source_name = "source_name"
          environment = "kPhysical"
          object_hash = "object_hash"
          object_type = "kCluster"
          logical_size_bytes = 1
          uuid = "uuid"
          global_id = "global_id"
          protection_type = "kAgent"
          sharepoint_site_summary {
            site_web_url = "site_web_url"
          }
          os_type = "kLinux"
          v_center_summary {
            is_cloud_env = true
          }
          windows_cluster_summary {
            cluster_source_type = "cluster_source_type"
          }
        }
        v_center_summary {
          is_cloud_env = true
        }
        windows_cluster_summary {
          cluster_source_type = "cluster_source_type"
        }
      }
      archival_target_info {
        target_id = 1
        archival_task_id = "archival_task_id"
        target_name = "target_name"
        target_type = "Tape"
        usage_type = "Archival"
        ownership_context = "Local"
        tier_settings {
          aws_tiering {
            tiers {
              move_after_unit = "Days"
              move_after = 1
              tier_type = "kAmazonS3Standard"
            }
          }
          azure_tiering {
            tiers {
              move_after_unit = "Days"
              move_after = 1
              tier_type = "kAzureTierHot"
            }
          }
          cloud_platform = "AWS"
          google_tiering {
            tiers {
              move_after_unit = "Days"
              move_after = 1
              tier_type = "kGoogleStandard"
            }
          }
          oracle_tiering {
            tiers {
              move_after_unit = "Days"
              move_after = 1
              tier_type = "kOracleTierStandard"
            }
          }
          current_tier_type = "kAmazonS3Standard"
        }
      }
      recover_from_standby = true
    }
    recovery_action = "RecoverPhysicalVolumes"
    recover_volume_params {
      target_environment = "kPhysical"
      physical_target_params {
        mount_target {
          id = 1
        }
        volume_mapping {
          source_volume_guid = "source_volume_guid"
          destination_volume_guid = "destination_volume_guid"
        }
        force_unmount_volume = true
        vlan_config {
          id = 1
          disable_vlan = true
        }
      }
    }
    mount_volume_params {
      target_environment = "kPhysical"
      physical_target_params {
        mount_to_original_target = true
        original_target_config {
          server_credentials {
            username = "username"
            password = "password"
          }
        }
        new_target_config {
          mount_target {
            id = 1
          }
          server_credentials {
            username = "username"
            password = "password"
          }
        }
        read_only_mount = true
        volume_names = [ "volumeNames" ]
        vlan_config {
          id = 1
          disable_vlan = true
        }
      }
    }
    recover_file_and_folder_params {
      files_and_folders {
        absolute_path = "absolute_path"
        is_directory = true
        is_view_file_recovery = true
      }
      target_environment = "kPhysical"
      physical_target_params {
        recover_target {
          id = 1
        }
        restore_to_original_paths = true
        overwrite_existing = true
        alternate_restore_directory = "alternate_restore_directory"
        preserve_attributes = true
        preserve_timestamps = true
        preserve_acls = true
        continue_on_error = true
        save_success_files = true
        vlan_config {
          id = 1
          disable_vlan = true
        }
        restore_entity_type = "kRegular"
      }
    }
    download_file_and_folder_params {
      expiry_time_usecs = 1
      files_and_folders {
        absolute_path = "absolute_path"
        is_directory = true
        is_view_file_recovery = true
      }
      download_file_path = "download_file_path"
    }
    system_recovery_params {
      full_nas_path = "full_nas_path"
    }
  }
  mssql_params {
    recover_app_params {
      snapshot_id = "snapshot_id"
      point_in_time_usecs = 1
      protection_group_id = "protection_group_id"
      protection_group_name = "protection_group_name"
      object_info {
        id = 1
        name = "name"
        source_id = 1
        source_name = "source_name"
        environment = "kPhysical"
        object_hash = "object_hash"
        object_type = "kCluster"
        logical_size_bytes = 1
        uuid = "uuid"
        global_id = "global_id"
        protection_type = "kAgent"
        sharepoint_site_summary {
          site_web_url = "site_web_url"
        }
        os_type = "kLinux"
        child_objects {
          id = 1
          name = "name"
          source_id = 1
          source_name = "source_name"
          environment = "kPhysical"
          object_hash = "object_hash"
          object_type = "kCluster"
          logical_size_bytes = 1
          uuid = "uuid"
          global_id = "global_id"
          protection_type = "kAgent"
          sharepoint_site_summary {
            site_web_url = "site_web_url"
          }
          os_type = "kLinux"
          v_center_summary {
            is_cloud_env = true
          }
          windows_cluster_summary {
            cluster_source_type = "cluster_source_type"
          }
        }
        v_center_summary {
          is_cloud_env = true
        }
        windows_cluster_summary {
          cluster_source_type = "cluster_source_type"
        }
      }
      archival_target_info {
        target_id = 1
        archival_task_id = "archival_task_id"
        target_name = "target_name"
        target_type = "Tape"
        usage_type = "Archival"
        ownership_context = "Local"
        tier_settings {
          aws_tiering {
            tiers {
              move_after_unit = "Days"
              move_after = 1
              tier_type = "kAmazonS3Standard"
            }
          }
          azure_tiering {
            tiers {
              move_after_unit = "Days"
              move_after = 1
              tier_type = "kAzureTierHot"
            }
          }
          cloud_platform = "AWS"
          google_tiering {
            tiers {
              move_after_unit = "Days"
              move_after = 1
              tier_type = "kGoogleStandard"
            }
          }
          oracle_tiering {
            tiers {
              move_after_unit = "Days"
              move_after = 1
              tier_type = "kOracleTierStandard"
            }
          }
          current_tier_type = "kAmazonS3Standard"
        }
      }
      recover_from_standby = true
      aag_info {
        name = "name"
        object_id = 1
      }
      host_info {
        id = "id"
        name = "name"
        environment = "kPhysical"
      }
      is_encrypted = true
      sql_target_params {
        new_source_config {
          keep_cdc = true
          multi_stage_restore_options {
            enable_auto_sync = true
            enable_multi_stage_restore = true
          }
          native_log_recovery_with_clause = "native_log_recovery_with_clause"
          native_recovery_with_clause = "native_recovery_with_clause"
          overwriting_policy = "FailIfExists"
          replay_entire_last_log = true
          restore_time_usecs = 1
          secondary_data_files_dir_list {
            directory = "directory"
            filename_pattern = "filename_pattern"
          }
          with_no_recovery = true
          data_file_directory_location = "data_file_directory_location"
          database_name = "database_name"
          host {
            id = 1
          }
          instance_name = "instance_name"
          log_file_directory_location = "log_file_directory_location"
        }
        original_source_config {
          keep_cdc = true
          multi_stage_restore_options {
            enable_auto_sync = true
            enable_multi_stage_restore = true
          }
          native_log_recovery_with_clause = "native_log_recovery_with_clause"
          native_recovery_with_clause = "native_recovery_with_clause"
          overwriting_policy = "FailIfExists"
          replay_entire_last_log = true
          restore_time_usecs = 1
          secondary_data_files_dir_list {
            directory = "directory"
            filename_pattern = "filename_pattern"
          }
          with_no_recovery = true
          capture_tail_logs = true
          data_file_directory_location = "data_file_directory_location"
          log_file_directory_location = "log_file_directory_location"
          new_database_name = "new_database_name"
        }
        recover_to_new_source = true
      }
      target_environment = "kSQL"
    }
    recovery_action = "RecoverApps"
    vlan_config {
      id = 1
      disable_vlan = true
    }
  }
}

// Provision backup_recovery_search_indexed_object resource instance
resource "ibm_backup_recovery_search_indexed_object" "backup_recovery_search_indexed_object_instance" {
  x_ibm_tenant_id = var.backup_recovery_search_indexed_object_x_ibm_tenant_id
  protection_group_ids = var.backup_recovery_search_indexed_object_protection_group_ids
  storage_domain_ids = var.backup_recovery_search_indexed_object_storage_domain_ids
  tenant_id = var.backup_recovery_search_indexed_object_tenant_id
  include_tenants = var.backup_recovery_search_indexed_object_include_tenants
  tags = var.backup_recovery_search_indexed_object_tags
  snapshot_tags = var.backup_recovery_search_indexed_object_snapshot_tags
  must_have_tag_ids = var.backup_recovery_search_indexed_object_must_have_tag_ids
  might_have_tag_ids = var.backup_recovery_search_indexed_object_might_have_tag_ids
  must_have_snapshot_tag_ids = var.backup_recovery_search_indexed_object_must_have_snapshot_tag_ids
  might_have_snapshot_tag_ids = var.backup_recovery_search_indexed_object_might_have_snapshot_tag_ids
  pagination_cookie = var.backup_recovery_search_indexed_object_pagination_cookie
  count = var.backup_recovery_search_indexed_object_count
  object_type = var.backup_recovery_search_indexed_object_object_type
  use_cached_data = var.backup_recovery_search_indexed_object_use_cached_data
  cassandra_params {
    cassandra_object_types = [ "CassandraKeyspaces" ]
    search_string = "search_string"
    source_ids = [ 1 ]
  }
  couchbase_params {
    couchbase_object_types = [ "CouchbaseBuckets" ]
    search_string = "search_string"
    source_ids = [ 1 ]
  }
  email_params {
    attendees_addresses = [ "attendeesAddresses" ]
    bcc_recipient_addresses = [ "bccRecipientAddresses" ]
    cc_recipient_addresses = [ "ccRecipientAddresses" ]
    created_end_time_secs = 1
    created_start_time_secs = 1
    due_date_end_time_secs = 1
    due_date_start_time_secs = 1
    email_address = "email_address"
    email_subject = "email_subject"
    first_name = "first_name"
    folder_names = [ "folderNames" ]
    has_attachment = true
    last_modified_end_time_secs = 1
    last_modified_start_time_secs = 1
    last_name = "last_name"
    middle_name = "middle_name"
    organizer_address = "organizer_address"
    received_end_time_secs = 1
    received_start_time_secs = 1
    recipient_addresses = [ "recipientAddresses" ]
    sender_address = "sender_address"
    source_environment = "kO365"
    task_status_types = [ "NotStarted" ]
    types = [ "Email" ]
    o365_params {
      domain_ids = [ 1 ]
      mailbox_ids = [ 1 ]
    }
  }
  exchange_params {
    search_string = "search_string"
  }
  file_params {
    search_string = "search_string"
    types = [ "File" ]
    source_environments = [ "kVMware" ]
    source_ids = [ 1 ]
    object_ids = [ 1 ]
  }
  hbase_params {
    hbase_object_types = [ "HbaseNamespaces" ]
    search_string = "search_string"
    source_ids = [ 1 ]
  }
  hdfs_params {
    hdfs_types = [ "HDFSFolders" ]
    search_string = "search_string"
    source_ids = [ 1 ]
  }
  hive_params {
    hive_object_types = [ "HiveDatabases" ]
    search_string = "search_string"
    source_ids = [ 1 ]
  }
  mongodb_params {
    mongo_db_object_types = [ "MongoDatabases" ]
    search_string = "search_string"
    source_ids = [ 1 ]
  }
  ms_groups_params {
    mailbox_params {
      attendees_addresses = [ "attendeesAddresses" ]
      bcc_recipient_addresses = [ "bccRecipientAddresses" ]
      cc_recipient_addresses = [ "ccRecipientAddresses" ]
      created_end_time_secs = 1
      created_start_time_secs = 1
      due_date_end_time_secs = 1
      due_date_start_time_secs = 1
      email_address = "email_address"
      email_subject = "email_subject"
      first_name = "first_name"
      folder_names = [ "folderNames" ]
      has_attachment = true
      last_modified_end_time_secs = 1
      last_modified_start_time_secs = 1
      last_name = "last_name"
      middle_name = "middle_name"
      organizer_address = "organizer_address"
      received_end_time_secs = 1
      received_start_time_secs = 1
      recipient_addresses = [ "recipientAddresses" ]
      sender_address = "sender_address"
      source_environment = "kO365"
      task_status_types = [ "NotStarted" ]
      types = [ "Email" ]
    }
    o365_params {
      domain_ids = [ 1 ]
      group_ids = [ 1 ]
      site_ids = [ 1 ]
      teams_ids = [ 1 ]
      user_ids = [ 1 ]
    }
    site_params {
      category_types = [ "Document" ]
      creation_end_time_secs = 1
      creation_start_time_secs = 1
      include_files = true
      include_folders = true
      o365_params {
        domain_ids = [ 1 ]
        group_ids = [ 1 ]
        site_ids = [ 1 ]
        teams_ids = [ 1 ]
        user_ids = [ 1 ]
      }
      owner_names = [ "ownerNames" ]
      search_string = "search_string"
      size_bytes_lower_limit = 1
      size_bytes_upper_limit = 1
    }
  }
  ms_teams_params {
    category_types = [ "Document" ]
    channel_names = [ "channelNames" ]
    channel_params {
      channel_email = "channel_email"
      channel_id = "channel_id"
      channel_name = "channel_name"
      include_private_channels = true
      include_public_channels = true
    }
    creation_end_time_secs = 1
    creation_start_time_secs = 1
    o365_params {
      domain_ids = [ 1 ]
      group_ids = [ 1 ]
      site_ids = [ 1 ]
      teams_ids = [ 1 ]
      user_ids = [ 1 ]
    }
    owner_names = [ "ownerNames" ]
    search_string = "search_string"
    size_bytes_lower_limit = 1
    size_bytes_upper_limit = 1
    types = [ "Channel" ]
  }
  one_drive_params {
    category_types = [ "Document" ]
    creation_end_time_secs = 1
    creation_start_time_secs = 1
    include_files = true
    include_folders = true
    o365_params {
      domain_ids = [ 1 ]
      group_ids = [ 1 ]
      site_ids = [ 1 ]
      teams_ids = [ 1 ]
      user_ids = [ 1 ]
    }
    owner_names = [ "ownerNames" ]
    search_string = "search_string"
    size_bytes_lower_limit = 1
    size_bytes_upper_limit = 1
  }
  public_folder_params {
    search_string = "search_string"
    types = [ "Calendar" ]
    has_attachment = true
    sender_address = "sender_address"
    recipient_addresses = [ "recipientAddresses" ]
    cc_recipient_addresses = [ "ccRecipientAddresses" ]
    bcc_recipient_addresses = [ "bccRecipientAddresses" ]
    received_start_time_secs = 1
    received_end_time_secs = 1
  }
  sfdc_params {
    mutation_types = [ "All" ]
    object_name = "object_name"
    query_string = "query_string"
    snapshot_id = "snapshot_id"
  }
  sharepoint_params {
    category_types = [ "Document" ]
    creation_end_time_secs = 1
    creation_start_time_secs = 1
    include_files = true
    include_folders = true
    o365_params {
      domain_ids = [ 1 ]
      group_ids = [ 1 ]
      site_ids = [ 1 ]
      teams_ids = [ 1 ]
      user_ids = [ 1 ]
    }
    owner_names = [ "ownerNames" ]
    search_string = "search_string"
    size_bytes_lower_limit = 1
    size_bytes_upper_limit = 1
  }
  uda_params {
    search_string = "search_string"
    source_ids = [ 1 ]
  }
}

// Provision backup_recovery_source_registration resource instance
resource "ibm_backup_recovery_source_registration" "backup_recovery_source_registration_instance" {
  x_ibm_tenant_id = var.backup_recovery_source_registration_x_ibm_tenant_id
  environment = var.backup_recovery_source_registration_environment
  name = var.backup_recovery_source_registration_name
  connection_id = var.backup_recovery_source_registration_connection_id
  connections {
    connection_id = 1
    entity_id = 1
    connector_group_id = 1
    data_source_connection_id = "data_source_connection_id"
  }
  connector_group_id = var.backup_recovery_source_registration_connector_group_id
  data_source_connection_id = var.backup_recovery_source_registration_data_source_connection_id
  advanced_configs {
    key = "key"
    value = "value"
  }
  physical_params {
    endpoint = "endpoint"
    force_register = true
    host_type = "kLinux"
    physical_type = "kGroup"
    applications = [ "kSQL" ]
  }
}

// Provision backup_recovery_update_protection_group_run_request resource instance
resource "ibm_backup_recovery_update_protection_group_run_request" "backup_recovery_update_protection_group_run_request_instance" {
  x_ibm_tenant_id = var.backup_recovery_update_protection_group_run_request_x_ibm_tenant_id
  update_protection_group_run_params {
    run_id = "run_id"
    local_snapshot_config {
      enable_legal_hold = true
      delete_snapshot = true
      data_lock = "Compliance"
      days_to_keep = 1
    }
    replication_snapshot_config {
      new_snapshot_config {
        id = 1
        retention {
          unit = "Days"
          duration = 1
          data_lock_config {
            mode = "Compliance"
            unit = "Days"
            duration = 1
            enable_worm_on_external_target = true
          }
        }
      }
      update_existing_snapshot_config {
        id = 1
        name = "name"
        enable_legal_hold = true
        delete_snapshot = true
        resync = true
        data_lock = "Compliance"
        days_to_keep = 1
      }
    }
    archival_snapshot_config {
      new_snapshot_config {
        id = 1
        archival_target_type = "Tape"
        retention {
          unit = "Days"
          duration = 1
          data_lock_config {
            mode = "Compliance"
            unit = "Days"
            duration = 1
            enable_worm_on_external_target = true
          }
        }
        copy_only_fully_successful = true
      }
      update_existing_snapshot_config {
        id = 1
        name = "name"
        archival_target_type = "Tape"
        enable_legal_hold = true
        delete_snapshot = true
        resync = true
        data_lock = "Compliance"
        days_to_keep = 1
      }
    }
  }
}

// Data source is not linked to a resource instance
// Uncomment if an existing data source instance exists

// Create backup_recovery_agent_upgrade_tasks data source
data "ibm_backup_recovery_agent_upgrade_tasks" "backup_recovery_agent_upgrade_tasks_instance" {
  x_ibm_tenant_id = var.backup_recovery_agent_upgrade_tasks_x_ibm_tenant_id
  ids = var.backup_recovery_agent_upgrade_tasks_ids
}


// Data source is not linked to a resource instance
// Uncomment if an existing data source instance exists

// Create backup_recovery_data_source_connections data source
data "ibm_backup_recovery_data_source_connections" "backup_recovery_data_source_connections_instance" {
  x_ibm_tenant_id = var.backup_recovery_data_source_connections_x_ibm_tenant_id
  connection_ids = var.backup_recovery_data_source_connections_connection_ids
  connection_names = var.backup_recovery_data_source_connections_connection_names
}


// Data source is not linked to a resource instance
// Uncomment if an existing data source instance exists

// Create backup_recovery_data_source_connectors data source
data "ibm_backup_recovery_data_source_connectors" "backup_recovery_data_source_connectors_instance" {
  x_ibm_tenant_id = var.backup_recovery_data_source_connectors_x_ibm_tenant_id
  connector_ids = var.backup_recovery_data_source_connectors_connector_ids
  connector_names = var.backup_recovery_data_source_connectors_connector_names
  connection_id = var.backup_recovery_data_source_connectors_connection_id
}


// Data source is not linked to a resource instance
// Uncomment if an existing data source instance exists

// Create backup_recovery_download_agent data source
data "ibm_backup_recovery_download_agent" "backup_recovery_download_agent_instance" {
  x_ibm_tenant_id = var.backup_recovery_download_agent_x_ibm_tenant_id
  platform = var.backup_recovery_download_agent_platform
  linux_params = var.backup_recovery_download_agent_linux_params
  file_path = "./Cohesity_Agent_ibm_rm_20240824_Win_x64_Installer_3.exe"
}


// Data source is not linked to a resource instance
// Uncomment if an existing data source instance exists

// Create backup_recovery_object_snapshots data source
data "ibm_backup_recovery_object_snapshots" "backup_recovery_object_snapshots_instance" {
  object_snapshots_id = var.backup_recovery_object_snapshots_backup_recovery_object_snapshots_id
  x_ibm_tenant_id = var.backup_recovery_object_snapshots_x_ibm_tenant_id
  from_time_usecs = var.backup_recovery_object_snapshots_from_time_usecs
  to_time_usecs = var.backup_recovery_object_snapshots_to_time_usecs
  run_start_from_time_usecs = var.backup_recovery_object_snapshots_run_start_from_time_usecs
  run_start_to_time_usecs = var.backup_recovery_object_snapshots_run_start_to_time_usecs
  snapshot_actions = var.backup_recovery_object_snapshots_snapshot_actions
  run_types = var.backup_recovery_object_snapshots_run_types
  protection_group_ids = var.backup_recovery_object_snapshots_protection_group_ids
  run_instance_ids = var.backup_recovery_object_snapshots_run_instance_ids
  region_ids = var.backup_recovery_object_snapshots_region_ids
  object_action_keys = var.backup_recovery_object_snapshots_object_action_keys
}


// Data source is not linked to a resource instance
// Uncomment if an existing data source instance exists

// Create backup_recovery_search_objects data source
data "ibm_backup_recovery_search_objects" "backup_recovery_search_objects_instance" {
  x_ibm_tenant_id = var.backup_recovery_search_objects_x_ibm_tenant_id
  request_initiator_type = var.backup_recovery_search_objects_request_initiator_type
  search_string = var.backup_recovery_search_objects_search_string
  environments = var.backup_recovery_search_objects_environments
  protection_types = var.backup_recovery_search_objects_protection_types
  protection_group_ids = var.backup_recovery_search_objects_protection_group_ids
  object_ids = var.backup_recovery_search_objects_object_ids
  os_types = var.backup_recovery_search_objects_os_types
  source_ids = var.backup_recovery_search_objects_source_ids
  source_uuids = var.backup_recovery_search_objects_source_uuids
  is_protected = var.backup_recovery_search_objects_is_protected
  is_deleted = var.backup_recovery_search_objects_is_deleted
  last_run_status_list = var.backup_recovery_search_objects_last_run_status_list
  cluster_identifiers = var.backup_recovery_search_objects_cluster_identifiers
  include_deleted_objects = var.backup_recovery_search_objects_include_deleted_objects
  pagination_cookie = var.backup_recovery_search_objects_pagination_cookie
  count = var.backup_recovery_search_objects_count
  must_have_tag_ids = var.backup_recovery_search_objects_must_have_tag_ids
  might_have_tag_ids = var.backup_recovery_search_objects_might_have_tag_ids
  must_have_snapshot_tag_ids = var.backup_recovery_search_objects_must_have_snapshot_tag_ids
  might_have_snapshot_tag_ids = var.backup_recovery_search_objects_might_have_snapshot_tag_ids
  tag_search_name = var.backup_recovery_search_objects_tag_search_name
  tag_names = var.backup_recovery_search_objects_tag_names
  tag_types = var.backup_recovery_search_objects_tag_types
  tag_categories = var.backup_recovery_search_objects_tag_categories
  tag_sub_categories = var.backup_recovery_search_objects_tag_sub_categories
  include_helios_tag_info_for_objects = var.backup_recovery_search_objects_include_helios_tag_info_for_objects
  external_filters = var.backup_recovery_search_objects_external_filters
}


// Data source is not linked to a resource instance
// Uncomment if an existing data source instance exists

// Create backup_recovery_search_protected_objects data source
data "ibm_backup_recovery_search_protected_objects" "backup_recovery_search_protected_objects_instance" {
  x_ibm_tenant_id = var.backup_recovery_search_protected_objects_x_ibm_tenant_id
  request_initiator_type = var.backup_recovery_search_protected_objects_request_initiator_type
  search_string = var.backup_recovery_search_protected_objects_search_string
  environments = var.backup_recovery_search_protected_objects_environments
  snapshot_actions = var.backup_recovery_search_protected_objects_snapshot_actions
  object_action_key = var.backup_recovery_search_protected_objects_object_action_key
  protection_group_ids = var.backup_recovery_search_protected_objects_protection_group_ids
  object_ids = var.backup_recovery_search_protected_objects_object_ids
  sub_result_size = var.backup_recovery_search_protected_objects_sub_result_size
  filter_snapshot_from_usecs = var.backup_recovery_search_protected_objects_filter_snapshot_from_usecs
  filter_snapshot_to_usecs = var.backup_recovery_search_protected_objects_filter_snapshot_to_usecs
  os_types = var.backup_recovery_search_protected_objects_os_types
  source_ids = var.backup_recovery_search_protected_objects_source_ids
  run_instance_ids = var.backup_recovery_search_protected_objects_run_instance_ids
  cdp_protected_only = var.backup_recovery_search_protected_objects_cdp_protected_only
  use_cached_data = var.backup_recovery_search_protected_objects_use_cached_data
}


// Data source is not linked to a resource instance
// Uncomment if an existing data source instance exists

// Create backup_recovery_protection_group data source
data "ibm_backup_recovery_protection_group" "backup_recovery_protection_group_instance" {
  protection_group_id = var.data_backup_recovery_protection_group_backup_recovery_protection_group_id
  x_ibm_tenant_id = var.data_backup_recovery_protection_group_x_ibm_tenant_id
  request_initiator_type = var.data_backup_recovery_protection_group_request_initiator_type
  include_last_run_info = var.data_backup_recovery_protection_group_include_last_run_info
  prune_excluded_source_ids = var.data_backup_recovery_protection_group_prune_excluded_source_ids
  prune_source_ids = var.data_backup_recovery_protection_group_prune_source_ids
}


// Data source is not linked to a resource instance
// Uncomment if an existing data source instance exists

// Create backup_recovery_protection_groups data source
data "ibm_backup_recovery_protection_groups" "backup_recovery_protection_groups_instance" {
  x_ibm_tenant_id = var.backup_recovery_protection_groups_x_ibm_tenant_id
  request_initiator_type = var.backup_recovery_protection_groups_request_initiator_type
  ids = var.backup_recovery_protection_groups_ids
  names = var.backup_recovery_protection_groups_names
  policy_ids = var.backup_recovery_protection_groups_policy_ids
  include_groups_with_datalock_only = var.backup_recovery_protection_groups_include_groups_with_datalock_only
  environments = var.backup_recovery_protection_groups_environments
  is_active = var.backup_recovery_protection_groups_is_active
  is_deleted = var.backup_recovery_protection_groups_is_deleted
  is_paused = var.backup_recovery_protection_groups_is_paused
  last_run_local_backup_status = var.backup_recovery_protection_groups_last_run_local_backup_status
  last_run_replication_status = var.backup_recovery_protection_groups_last_run_replication_status
  last_run_archival_status = var.backup_recovery_protection_groups_last_run_archival_status
  last_run_cloud_spin_status = var.backup_recovery_protection_groups_last_run_cloud_spin_status
  last_run_any_status = var.backup_recovery_protection_groups_last_run_any_status
  is_last_run_sla_violated = var.backup_recovery_protection_groups_is_last_run_sla_violated
  include_last_run_info = var.backup_recovery_protection_groups_include_last_run_info
  prune_excluded_source_ids = var.backup_recovery_protection_groups_prune_excluded_source_ids
  prune_source_ids = var.backup_recovery_protection_groups_prune_source_ids
  use_cached_data = var.backup_recovery_protection_groups_use_cached_data
  source_ids = var.backup_recovery_protection_groups_source_ids
}


// Data source is not linked to a resource instance
// Uncomment if an existing data source instance exists

// Create backup_recovery_protection_group_run data source
data "ibm_backup_recovery_protection_group_run" "backup_recovery_protection_group_run_instance" {
  protection_group_run_id = var.backup_recovery_protection_group_run_backup_recovery_protection_group_run_id
  x_ibm_tenant_id = var.backup_recovery_protection_group_run_x_ibm_tenant_id
  request_initiator_type = var.backup_recovery_protection_group_run_request_initiator_type
  run_id = var.backup_recovery_protection_group_run_run_id
  start_time_usecs = var.backup_recovery_protection_group_run_start_time_usecs
  end_time_usecs = var.backup_recovery_protection_group_run_end_time_usecs
  run_types = var.backup_recovery_protection_group_run_run_types
  include_object_details = var.backup_recovery_protection_group_run_include_object_details
  local_backup_run_status = var.backup_recovery_protection_group_run_local_backup_run_status
  replication_run_status = var.backup_recovery_protection_group_run_replication_run_status
  archival_run_status = var.backup_recovery_protection_group_run_archival_run_status
  cloud_spin_run_status = var.backup_recovery_protection_group_run_cloud_spin_run_status
  num_runs = var.backup_recovery_protection_group_run_num_runs
  exclude_non_restorable_runs = var.backup_recovery_protection_group_run_exclude_non_restorable_runs
  run_tags = var.backup_recovery_protection_group_run_run_tags
  use_cached_data = var.backup_recovery_protection_group_run_use_cached_data
  filter_by_end_time = var.backup_recovery_protection_group_run_filter_by_end_time
  snapshot_target_types = var.backup_recovery_protection_group_run_snapshot_target_types
  only_return_successful_copy_run = var.backup_recovery_protection_group_run_only_return_successful_copy_run
  filter_by_copy_task_end_time = var.backup_recovery_protection_group_run_filter_by_copy_task_end_time
}


// Data source is not linked to a resource instance
// Uncomment if an existing data source instance exists

// Create backup_recovery_protection_group_runs data source
data "ibm_backup_recovery_protection_group_runs" "backup_recovery_protection_group_runs_instance" {
  protection_group_runs_id = var.backup_recovery_protection_group_runs_backup_recovery_protection_group_runs_id
  x_ibm_tenant_id = var.backup_recovery_protection_group_runs_x_ibm_tenant_id
  request_initiator_type = var.backup_recovery_protection_group_runs_request_initiator_type
  run_id = var.backup_recovery_protection_group_runs_run_id
  start_time_usecs = var.backup_recovery_protection_group_runs_start_time_usecs
  end_time_usecs = var.backup_recovery_protection_group_runs_end_time_usecs
  run_types = var.backup_recovery_protection_group_runs_run_types
  include_object_details = var.backup_recovery_protection_group_runs_include_object_details
  local_backup_run_status = var.backup_recovery_protection_group_runs_local_backup_run_status
  replication_run_status = var.backup_recovery_protection_group_runs_replication_run_status
  archival_run_status = var.backup_recovery_protection_group_runs_archival_run_status
  cloud_spin_run_status = var.backup_recovery_protection_group_runs_cloud_spin_run_status
  num_runs = var.backup_recovery_protection_group_runs_num_runs
  exclude_non_restorable_runs = var.backup_recovery_protection_group_runs_exclude_non_restorable_runs
  run_tags = var.backup_recovery_protection_group_runs_run_tags
  use_cached_data = var.backup_recovery_protection_group_runs_use_cached_data
  filter_by_end_time = var.backup_recovery_protection_group_runs_filter_by_end_time
  snapshot_target_types = var.backup_recovery_protection_group_runs_snapshot_target_types
  only_return_successful_copy_run = var.backup_recovery_protection_group_runs_only_return_successful_copy_run
  filter_by_copy_task_end_time = var.backup_recovery_protection_group_runs_filter_by_copy_task_end_time
}


// Data source is not linked to a resource instance
// Uncomment if an existing data source instance exists

// Create backup_recovery_protection_policies data source
data "ibm_backup_recovery_protection_policies" "backup_recovery_protection_policies_instance" {
  x_ibm_tenant_id = var.backup_recovery_protection_policies_x_ibm_tenant_id
  request_initiator_type = var.backup_recovery_protection_policies_request_initiator_type
  ids = var.backup_recovery_protection_policies_ids
  policy_names = var.backup_recovery_protection_policies_policy_names
  types = var.backup_recovery_protection_policies_types
  exclude_linked_policies = var.backup_recovery_protection_policies_exclude_linked_policies
  include_replicated_policies = var.backup_recovery_protection_policies_include_replicated_policies
  include_stats = var.backup_recovery_protection_policies_include_stats
}


// Data source is not linked to a resource instance
// Uncomment if an existing data source instance exists

// Create backup_recovery_protection_policy data source
data "ibm_backup_recovery_protection_policy" "backup_recovery_protection_policy_instance" {
  protection_policy_id = var.data_backup_recovery_protection_policy_backup_recovery_protection_policy_id
  x_ibm_tenant_id = var.data_backup_recovery_protection_policy_x_ibm_tenant_id
  request_initiator_type = var.data_backup_recovery_protection_policy_request_initiator_type
}


// Data source is not linked to a resource instance
// Uncomment if an existing data source instance exists

// Create backup_recovery data source
data "ibm_backup_recovery" "backup_recovery_instance" {
  recovery_id = var.data_backup_recovery_backup_recovery_id
  x_ibm_tenant_id = var.data_backup_recovery_x_ibm_tenant_id
}


// Data source is not linked to a resource instance
// Uncomment if an existing data source instance exists

// Create backup_recoveries data source
data "ibm_backup_recoveries" "backup_recoveries_instance" {
  x_ibm_tenant_id = var.backup_recoveries_x_ibm_tenant_id
  ids = var.backup_recoveries_ids
  return_only_child_recoveries = var.backup_recoveries_return_only_child_recoveries
  start_time_usecs = var.backup_recoveries_start_time_usecs
  end_time_usecs = var.backup_recoveries_end_time_usecs
  snapshot_target_type = var.backup_recoveries_snapshot_target_type
  archival_target_type = var.backup_recoveries_archival_target_type
  snapshot_environments = var.backup_recoveries_snapshot_environments
  status = var.backup_recoveries_status
  recovery_actions = var.backup_recoveries_recovery_actions
}


// Data source is not linked to a resource instance
// Uncomment if an existing data source instance exists

// Create backup_recovery_source_registrations data source
data "ibm_backup_recovery_source_registrations" "backup_recovery_source_registrations_instance" {
  x_ibm_tenant_id = var.backup_recovery_source_registrations_x_ibm_tenant_id
  ids = var.backup_recovery_source_registrations_ids
  include_source_credentials = var.backup_recovery_source_registrations_include_source_credentials
  encryption_key = var.backup_recovery_source_registrations_encryption_key
  use_cached_data = var.backup_recovery_source_registrations_use_cached_data
  include_external_metadata = var.backup_recovery_source_registrations_include_external_metadata
  ignore_tenant_migration_in_progress_check = var.backup_recovery_source_registrations_ignore_tenant_migration_in_progress_check
}


// Data source is not linked to a resource instance
// Uncomment if an existing data source instance exists

// Create backup_recovery_source_registration data source
data "ibm_backup_recovery_source_registration" "backup_recovery_source_registration_instance" {
  source_registration_id = var.data_backup_recovery_source_registration_backup_recovery_source_registration_id
  x_ibm_tenant_id = var.data_backup_recovery_source_registration_x_ibm_tenant_id
  request_initiator_type = var.data_backup_recovery_source_registration_request_initiator_type
}


// Data source is not linked to a resource instance
// Uncomment if an existing data source instance exists

// Create backup_recovery_download_indexed_files data source
data "ibm_backup_recovery_download_indexed_files" "backup_recovery_download_indexed_files_instance" {
  snapshots_id = var.backup_recovery_download_indexed_files_snapshots_id
  x_ibm_tenant_id = var.backup_recovery_download_indexed_files_x_ibm_tenant_id
  file_path = var.backup_recovery_download_indexed_files_file_path
  nvram_file = var.backup_recovery_download_indexed_files_nvram_file
  retry_attempt = var.backup_recovery_download_indexed_files_retry_attempt
  start_offset = var.backup_recovery_download_indexed_files_start_offset
  length = var.backup_recovery_download_indexed_files_length
}

