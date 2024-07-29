provider "ibm" {
  ibmcloud_api_key = var.ibmcloud_api_key
}

// Provision protection_group_run_request resource instance
resource "ibm_protection_group_run_request" "protection_group_run_request_instance" {
  run_type = var.protection_group_run_request_run_type
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
  uda_params {
    externally_triggered_run_params {
      control_node = "control_node"
      backup_args {
        key = "key"
        value = "value"
      }
    }
  }
}

// Provision recovery_download_files_folders resource instance
resource "ibm_recovery_download_files_folders" "recovery_download_files_folders_instance" {
  name = var.recovery_download_files_folders_name
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
      os_type = "kLinux"
    }
    archival_target_info {
      target_id = 1
      archival_task_id = "archival_task_id"
      target_name = "target_name"
      target_type = "Tape"
      usage_type = "Archival"
      ownership_context = "Local"
      tier_settings {
        cloud_platform = "Oracle"
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
  parent_recovery_id = var.recovery_download_files_folders_parent_recovery_id
  files_and_folders {
    absolute_path = "absolute_path"
    is_directory = true
  }
  glacier_retrieval_type = var.recovery_download_files_folders_glacier_retrieval_type
}

// Provision perform_action_on_protection_group_run_request resource instance
resource "ibm_perform_action_on_protection_group_run_request" "perform_action_on_protection_group_run_request_instance" {
  action = var.perform_action_on_protection_group_run_request_action
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

// Provision protection_group resource instance
resource "ibm_protection_group" "protection_group_instance" {
  name = var.protection_group_name
  policy_id = var.protection_group_policy_id
  priority = var.protection_group_priority
  storage_domain_id = var.protection_group_storage_domain_id
  description = var.protection_group_description
  start_time {
    hour = 0
    minute = 0
    time_zone = "time_zone"
  }
  end_time_usecs = var.protection_group_end_time_usecs
  last_modified_timestamp_usecs = var.protection_group_last_modified_timestamp_usecs
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
  qos_policy = var.protection_group_qos_policy
  abort_in_blackouts = var.protection_group_abort_in_blackouts
  pause_in_blackouts = var.protection_group_pause_in_blackouts
  is_paused = var.protection_group_is_paused
  environment = var.protection_group_environment
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
      objects {
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
  oracle_params {
    objects {
      source_id = 1
      db_params {
        database_id = 1
        db_channels {
          archive_log_retention_days = 1
          archive_log_retention_hours = 1
          credentials {
            username = "username"
            password = "password"
          }
          database_unique_name = "database_unique_name"
          database_uuid = "database_uuid"
          default_channel_count = 1
          database_node_list {
            host_id = "host_id"
            channel_count = 1
            port = 1
            sbt_host_params {
              sbt_library_path = "sbt_library_path"
              view_fs_path = "view_fs_path"
              vip_list = [ "vipList" ]
              vlan_info_list {
                ip_list = [ "ipList" ]
                gateway = "gateway"
                id = 1
                subnet_ip = "subnet_ip"
              }
            }
          }
          max_host_count = 1
          enable_dg_primary_backup = true
          rman_backup_type = "kImageCopy"
        }
      }
    }
    persist_mountpoints = true
    vlan_params {
      vlan_id = 1
      disable_vlan = true
      interface_name = "interface_name"
    }
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
    log_auto_kill_timeout_secs = 1
    incr_auto_kill_timeout_secs = 1
    full_auto_kill_timeout_secs = 1
  }
}

// Provision protection_policy resource instance
resource "ibm_protection_policy" "protection_policy_instance" {
  name = var.protection_policy_name
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
            cloud_platform = "Oracle"
            oracle_tiering {
              tiers {
                move_after_unit = "Days"
                move_after = 1
                tier_type = "kOracleTierStandard"
              }
            }
          }
        }
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
  description = var.protection_policy_description
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
        duration = 1
        data_lock_config {
          mode = "Compliance"
          unit = "Days"
          duration = 1
          enable_worm_on_external_target = true
        }
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
        duration = 1
        data_lock_config {
          mode = "Compliance"
          unit = "Days"
          duration = 1
          enable_worm_on_external_target = true
        }
      }
      target_id = 1
      tier_settings {
        cloud_platform = "Oracle"
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
        duration = 1
        data_lock_config {
          mode = "Compliance"
          unit = "Days"
          duration = 1
          enable_worm_on_external_target = true
        }
      }
      target {
        id = 1
        name = "name"
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
        duration = 1
        data_lock_config {
          mode = "Compliance"
          unit = "Days"
          duration = 1
          enable_worm_on_external_target = true
        }
      }
      params {
        id = 1
        restore_v_mware_params {
          target_vm_folder_id = 1
          target_data_store_id = 1
          enable_copy_recovery = true
          resource_pool_id = 1
          datastore_ids = [ 1 ]
          overwrite_existing_vm = true
          power_off_and_rename_existing_vm = true
          attempt_differential_restore = true
          is_on_prem_deploy = true
        }
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
        duration = 1
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
          duration = 1
          data_lock_config {
            mode = "Compliance"
            unit = "Days"
            duration = 1
            enable_worm_on_external_target = true
          }
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
          duration = 1
          data_lock_config {
            mode = "Compliance"
            unit = "Days"
            duration = 1
            enable_worm_on_external_target = true
          }
        }
        target_id = 1
        tier_settings {
          cloud_platform = "Oracle"
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
          duration = 1
          data_lock_config {
            mode = "Compliance"
            unit = "Days"
            duration = 1
            enable_worm_on_external_target = true
          }
        }
        target {
          id = 1
          name = "name"
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
          duration = 1
          data_lock_config {
            mode = "Compliance"
            unit = "Days"
            duration = 1
            enable_worm_on_external_target = true
          }
        }
        params {
          id = 1
          restore_v_mware_params {
            target_vm_folder_id = 1
            target_data_store_id = 1
            enable_copy_recovery = true
            resource_pool_id = 1
            datastore_ids = [ 1 ]
            overwrite_existing_vm = true
            power_off_and_rename_existing_vm = true
            attempt_differential_restore = true
            is_on_prem_deploy = true
          }
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
          duration = 1
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
  data_lock = var.protection_policy_data_lock
  version = var.protection_policy_version
  is_cbs_enabled = var.protection_policy_is_cbs_enabled
  last_modification_time_usecs = var.protection_policy_last_modification_time_usecs
  template_id = var.protection_policy_template_id
}

// Provision recovery resource instance
resource "ibm_recovery" "recovery_instance" {
  request_initiator_type = var.recovery_request_initiator_type
  name = var.recovery_name
  snapshot_environment = var.recovery_snapshot_environment
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
        os_type = "kLinux"
      }
      archival_target_info {
        target_id = 1
        archival_task_id = "archival_task_id"
        target_name = "target_name"
        target_type = "Tape"
        usage_type = "Archival"
        ownership_context = "Local"
        tier_settings {
          cloud_platform = "Oracle"
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
  oracle_params {
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
        os_type = "kLinux"
      }
      archival_target_info {
        target_id = 1
        archival_task_id = "archival_task_id"
        target_name = "target_name"
        target_type = "Tape"
        usage_type = "Archival"
        ownership_context = "Local"
        tier_settings {
          cloud_platform = "Oracle"
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
      instant_recovery_info {
      }
    }
    recovery_action = "RecoverApps"
    recover_app_params {
      target_environment = "kOracle"
      oracle_target_params {
        recover_to_new_source = true
        new_source_config {
          host {
            id = 1
          }
          recovery_target = "RecoverDatabase"
          recover_database_params {
            restore_time_usecs = 1
            db_channels {
              archive_log_retention_days = 1
              archive_log_retention_hours = 1
              credentials {
                username = "username"
                password = "password"
              }
              database_unique_name = "database_unique_name"
              database_uuid = "database_uuid"
              default_channel_count = 1
              database_node_list {
                host_id = "host_id"
                channel_count = 1
                port = 1
                sbt_host_params {
                  sbt_library_path = "sbt_library_path"
                  view_fs_path = "view_fs_path"
                  vip_list = [ "vipList" ]
                  vlan_info_list {
                    ip_list = [ "ipList" ]
                    gateway = "gateway"
                    id = 1
                    subnet_ip = "subnet_ip"
                  }
                }
              }
              max_host_count = 1
              enable_dg_primary_backup = true
              rman_backup_type = "kImageCopy"
            }
            recovery_mode = true
            shell_evironment_vars {
              key = "key"
              value = "value"
            }
            granular_restore_info {
              granularity_type = "kPDB"
              pdb_restore_params {
                drop_duplicate_pdb = true
                pdb_objects {
                  db_id = "db_id"
                  db_name = "db_name"
                }
                restore_to_existing_cdb = true
                rename_pdb_map {
                  key = "key"
                  value = "value"
                }
                include_in_restore = true
              }
            }
            oracle_archive_log_info {
              range_type = "Time"
              range_info_vec {
                start_of_range = 1
                end_of_range = 1
                protection_group_id = "protection_group_id"
                reset_log_id = 1
                incarnation_id = 1
                thread_id = 1
              }
              archive_log_restore_dest = "archive_log_restore_dest"
            }
            oracle_recovery_validation_info {
              create_dummy_instance = true
            }
            restore_spfile_or_pfile_info {
              should_restore_spfile_or_pfile = true
              file_location = "file_location"
            }
            use_scn_for_restore = true
            database_name = "database_name"
            oracle_base_folder = "oracle_base_folder"
            oracle_home_folder = "oracle_home_folder"
            db_files_destination = "db_files_destination"
            db_config_file_path = "db_config_file_path"
            enable_archive_log_mode = true
            pfile_parameter_map {
              key = "key"
              value = "value"
            }
            bct_file_path = "bct_file_path"
            num_tempfiles = 1
            redo_log_config {
              num_groups = 1
              member_prefix = "member_prefix"
              size_m_bytes = 1
              group_members = [ "groupMembers" ]
            }
            is_multi_stage_restore = true
            oracle_update_restore_options {
              delay_secs = 1
              target_path_vec = [ "targetPathVec" ]
            }
            skip_clone_nid = true
            no_filename_check = true
            new_name_clause = "new_name_clause"
          }
          recover_view_params {
            restore_time_usecs = 1
            db_channels {
              archive_log_retention_days = 1
              archive_log_retention_hours = 1
              credentials {
                username = "username"
                password = "password"
              }
              database_unique_name = "database_unique_name"
              database_uuid = "database_uuid"
              default_channel_count = 1
              database_node_list {
                host_id = "host_id"
                channel_count = 1
                port = 1
                sbt_host_params {
                  sbt_library_path = "sbt_library_path"
                  view_fs_path = "view_fs_path"
                  vip_list = [ "vipList" ]
                  vlan_info_list {
                    ip_list = [ "ipList" ]
                    gateway = "gateway"
                    id = 1
                    subnet_ip = "subnet_ip"
                  }
                }
              }
              max_host_count = 1
              enable_dg_primary_backup = true
              rman_backup_type = "kImageCopy"
            }
            recovery_mode = true
            shell_evironment_vars {
              key = "key"
              value = "value"
            }
            granular_restore_info {
              granularity_type = "kPDB"
              pdb_restore_params {
                drop_duplicate_pdb = true
                pdb_objects {
                  db_id = "db_id"
                  db_name = "db_name"
                }
                restore_to_existing_cdb = true
                rename_pdb_map {
                  key = "key"
                  value = "value"
                }
                include_in_restore = true
              }
            }
            oracle_archive_log_info {
              range_type = "Time"
              range_info_vec {
                start_of_range = 1
                end_of_range = 1
                protection_group_id = "protection_group_id"
                reset_log_id = 1
                incarnation_id = 1
                thread_id = 1
              }
              archive_log_restore_dest = "archive_log_restore_dest"
            }
            oracle_recovery_validation_info {
              create_dummy_instance = true
            }
            restore_spfile_or_pfile_info {
              should_restore_spfile_or_pfile = true
              file_location = "file_location"
            }
            use_scn_for_restore = true
            view_mount_path = "view_mount_path"
          }
        }
        original_source_config {
          restore_time_usecs = 1
          db_channels {
            archive_log_retention_days = 1
            archive_log_retention_hours = 1
            credentials {
              username = "username"
              password = "password"
            }
            database_unique_name = "database_unique_name"
            database_uuid = "database_uuid"
            default_channel_count = 1
            database_node_list {
              host_id = "host_id"
              channel_count = 1
              port = 1
              sbt_host_params {
                sbt_library_path = "sbt_library_path"
                view_fs_path = "view_fs_path"
                vip_list = [ "vipList" ]
                vlan_info_list {
                  ip_list = [ "ipList" ]
                  gateway = "gateway"
                  id = 1
                  subnet_ip = "subnet_ip"
                }
              }
            }
            max_host_count = 1
            enable_dg_primary_backup = true
            rman_backup_type = "kImageCopy"
          }
          recovery_mode = true
          shell_evironment_vars {
            key = "key"
            value = "value"
          }
          granular_restore_info {
            granularity_type = "kPDB"
            pdb_restore_params {
              drop_duplicate_pdb = true
              pdb_objects {
                db_id = "db_id"
                db_name = "db_name"
              }
              restore_to_existing_cdb = true
              rename_pdb_map {
                key = "key"
                value = "value"
              }
              include_in_restore = true
            }
          }
          oracle_archive_log_info {
            range_type = "Time"
            range_info_vec {
              start_of_range = 1
              end_of_range = 1
              protection_group_id = "protection_group_id"
              reset_log_id = 1
              incarnation_id = 1
              thread_id = 1
            }
            archive_log_restore_dest = "archive_log_restore_dest"
          }
          oracle_recovery_validation_info {
            create_dummy_instance = true
          }
          restore_spfile_or_pfile_info {
            should_restore_spfile_or_pfile = true
            file_location = "file_location"
          }
          use_scn_for_restore = true
          roll_forward_log_path_vec = [ "rollForwardLogPathVec" ]
          attempt_complete_recovery = true
          roll_forward_time_msecs = 1
          stop_active_passive = true
        }
      }
      vlan_config {
        id = 1
        disable_vlan = true
      }
    }
  }
}

// Provision search_indexed_object resource instance
resource "ibm_search_indexed_object" "search_indexed_object_instance" {
  protection_group_ids = var.search_indexed_object_protection_group_ids
  storage_domain_ids = var.search_indexed_object_storage_domain_ids
  tenant_id = var.search_indexed_object_tenant_id
  include_tenants = var.search_indexed_object_include_tenants
  tags = var.search_indexed_object_tags
  snapshot_tags = var.search_indexed_object_snapshot_tags
  must_have_tag_ids = var.search_indexed_object_must_have_tag_ids
  might_have_tag_ids = var.search_indexed_object_might_have_tag_ids
  must_have_snapshot_tag_ids = var.search_indexed_object_must_have_snapshot_tag_ids
  might_have_snapshot_tag_ids = var.search_indexed_object_might_have_snapshot_tag_ids
  pagination_cookie = var.search_indexed_object_pagination_cookie
  count = var.search_indexed_object_count
  object_type = var.search_indexed_object_object_type
  use_cached_data = var.search_indexed_object_use_cached_data
  files {
    search_string = "search_string"
    types = [ "File" ]
    source_environments = [ "kSQL" ]
    source_ids = [ 1 ]
    object_ids = [ 1 ]
  }
  public_folders {
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
}

// Provision source_registration resource instance
resource "ibm_source_registration" "source_registration_instance" {
  environment = var.source_registration_environment
  name = var.source_registration_name
  connection_id = var.source_registration_connection_id
  connections {
    connection_id = 1
    entity_id = 1
    connector_group_id = 1
  }
  connector_group_id = var.source_registration_connector_group_id
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
  oracle_params {
    database_entity_info {
      container_database_info {
        database_id = "database_id"
        database_name = "database_name"
      }
      data_guard_info {
        role = "kPrimary"
        standby_type = "kPhysical"
      }
    }
    host_info {
      id = "id"
      name = "name"
      environment = "kPhysical"
    }
  }
}

// Provision update_protection_group_run_request resource instance
resource "ibm_update_protection_group_run_request" "update_protection_group_run_request_instance" {
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

// Provision protection_group_state resource instance
resource "ibm_protection_group_state" "protection_group_state_instance" {
  action = var.protection_group_state_action
  ids = var.protection_group_state_ids
}

// Data source is not linked to a resource instance
// Uncomment if an existing data source instance exists
/*
// Create run_debug_logs data source
data "ibm_run_debug_logs" "run_debug_logs_instance" {
  run_debug_logs_id = var.run_debug_logs_run_debug_logs_id
  run_id = var.run_debug_logs_run_id
  object_id = var.run_debug_logs_object_id
}
*/

// Data source is not linked to a resource instance
// Uncomment if an existing data source instance exists
/*
// Create object_run_debug_logs data source
data "ibm_object_run_debug_logs" "object_run_debug_logs_instance" {
  object_run_debug_logs_id = var.object_run_debug_logs_object_run_debug_logs_id
  run_id = var.object_run_debug_logs_run_id
  object_id = var.object_run_debug_logs_object_id
}
*/

// Data source is not linked to a resource instance
// Uncomment if an existing data source instance exists
/*
// Create run_error_report data source
data "ibm_run_error_report" "run_error_report_instance" {
  run_error_report_id = var.run_error_report_run_error_report_id
  run_id = var.run_error_report_run_id
  object_id = var.run_error_report_object_id
}
*/

// Data source is not linked to a resource instance
// Uncomment if an existing data source instance exists
/*
// Create runs_report data source
data "ibm_runs_report" "runs_report_instance" {
  runs_report_id = var.runs_report_runs_report_id
  run_id = var.runs_report_run_id
  object_id = var.runs_report_object_id
  file_type = var.runs_report_file_type
  name = var.runs_report_name
}
*/

// Data source is not linked to a resource instance
// Uncomment if an existing data source instance exists
/*
// Create recovery_debug_logs data source
data "ibm_recovery_debug_logs" "recovery_debug_logs_instance" {
  recovery_debug_logs_id = var.recovery_debug_logs_recovery_debug_logs_id
}
*/

// Data source is not linked to a resource instance
// Uncomment if an existing data source instance exists
/*
// Create recovery_download_messages data source
data "ibm_recovery_download_messages" "recovery_download_messages_instance" {
  recovery_download_messages_id = var.recovery_download_messages_recovery_download_messages_id
}
*/

// Data source is not linked to a resource instance
// Uncomment if an existing data source instance exists
/*
// Create recovery_download_files data source
data "ibm_recovery_download_files" "recovery_download_files_instance" {
  recovery_download_files_id = var.recovery_download_files_recovery_download_files_id
  start_offset = var.recovery_download_files_start_offset
  length = var.recovery_download_files_length
  file_type = var.recovery_download_files_file_type
  source_name = var.recovery_download_files_source_name
  start_time = var.recovery_download_files_start_time
  include_tenants = var.recovery_download_files_include_tenants
}
*/

// Data source is not linked to a resource instance
// Uncomment if an existing data source instance exists
/*
// Create recovery_fetch_uptier_data data source
data "ibm_recovery_fetch_uptier_data" "recovery_fetch_uptier_data_instance" {
  archive_u_id = var.recovery_fetch_uptier_data_archive_u_id
}
*/

// Data source is not linked to a resource instance
// Uncomment if an existing data source instance exists
/*
// Create protection_run_progress data source
data "ibm_protection_run_progress" "protection_run_progress_instance" {
  run_id = var.protection_run_progress_run_id
  objects = var.protection_run_progress_objects
  tenant_ids = var.protection_run_progress_tenant_ids
  include_tenants = var.protection_run_progress_include_tenants
  include_finished_tasks = var.protection_run_progress_include_finished_tasks
  start_time_usecs = var.protection_run_progress_start_time_usecs
  end_time_usecs = var.protection_run_progress_end_time_usecs
  max_tasks_num = var.protection_run_progress_max_tasks_num
  exclude_object_details = var.protection_run_progress_exclude_object_details
  include_event_logs = var.protection_run_progress_include_event_logs
  max_log_level = var.protection_run_progress_max_log_level
  run_task_path = var.protection_run_progress_run_task_path
  object_task_paths = var.protection_run_progress_object_task_paths
}
*/

// Data source is not linked to a resource instance
// Uncomment if an existing data source instance exists
/*
// Create protection_run_stat data source
data "ibm_protection_run_stat" "protection_run_stat_instance" {
  run_id = var.protection_run_stat_run_id
  objects = var.protection_run_stat_objects
  tenant_ids = var.protection_run_stat_tenant_ids
  include_tenants = var.protection_run_stat_include_tenants
  include_finished_tasks = var.protection_run_stat_include_finished_tasks
  start_time_usecs = var.protection_run_stat_start_time_usecs
  end_time_usecs = var.protection_run_stat_end_time_usecs
  max_tasks_num = var.protection_run_stat_max_tasks_num
  exclude_object_details = var.protection_run_stat_exclude_object_details
  run_task_path = var.protection_run_stat_run_task_path
  object_task_paths = var.protection_run_stat_object_task_paths
}
*/

// Data source is not linked to a resource instance
// Uncomment if an existing data source instance exists
/*
// Create search_objects data source
data "ibm_search_objects" "search_objects_instance" {
  request_initiator_type = var.search_objects_request_initiator_type
  search_string = var.search_objects_search_string
  environments = var.search_objects_environments
  protection_types = var.search_objects_protection_types
  tenant_ids = var.search_objects_tenant_ids
  include_tenants = var.search_objects_include_tenants
  protection_group_ids = var.search_objects_protection_group_ids
  object_ids = var.search_objects_object_ids
  os_types = var.search_objects_os_types
  source_ids = var.search_objects_source_ids
  source_uuids = var.search_objects_source_uuids
  is_protected = var.search_objects_is_protected
  is_deleted = var.search_objects_is_deleted
  last_run_status_list = var.search_objects_last_run_status_list
  region_ids = var.search_objects_region_ids
  cluster_identifiers = var.search_objects_cluster_identifiers
  storage_domain_ids = var.search_objects_storage_domain_ids
  include_deleted_objects = var.search_objects_include_deleted_objects
  pagination_cookie = var.search_objects_pagination_cookie
  count = var.search_objects_count
  must_have_tag_ids = var.search_objects_must_have_tag_ids
  might_have_tag_ids = var.search_objects_might_have_tag_ids
  must_have_snapshot_tag_ids = var.search_objects_must_have_snapshot_tag_ids
  might_have_snapshot_tag_ids = var.search_objects_might_have_snapshot_tag_ids
}
*/

// Data source is not linked to a resource instance
// Uncomment if an existing data source instance exists
/*
// Create search_protected_objects data source
data "ibm_search_protected_objects" "search_protected_objects_instance" {
  request_initiator_type = var.search_protected_objects_request_initiator_type
  search_string = var.search_protected_objects_search_string
  environments = var.search_protected_objects_environments
  snapshot_actions = var.search_protected_objects_snapshot_actions
  object_action_key = var.search_protected_objects_object_action_key
  tenant_ids = var.search_protected_objects_tenant_ids
  include_tenants = var.search_protected_objects_include_tenants
  protection_group_ids = var.search_protected_objects_protection_group_ids
  object_ids = var.search_protected_objects_object_ids
  storage_domain_ids = var.search_protected_objects_storage_domain_ids
  sub_result_size = var.search_protected_objects_sub_result_size
  filter_snapshot_from_usecs = var.search_protected_objects_filter_snapshot_from_usecs
  filter_snapshot_to_usecs = var.search_protected_objects_filter_snapshot_to_usecs
  os_types = var.search_protected_objects_os_types
  source_ids = var.search_protected_objects_source_ids
  run_instance_ids = var.search_protected_objects_run_instance_ids
  cdp_protected_only = var.search_protected_objects_cdp_protected_only
  region_ids = var.search_protected_objects_region_ids
  use_cached_data = var.search_protected_objects_use_cached_data
}
*/

// Data source is not linked to a resource instance
// Uncomment if an existing data source instance exists
/*
// Create protection_group data source
data "ibm_protection_group" "protection_group_instance" {
  protection_group_id = var.data_protection_group_protection_group_id
  request_initiator_type = var.data_protection_group_request_initiator_type
  include_last_run_info = var.data_protection_group_include_last_run_info
  prune_excluded_source_ids = var.data_protection_group_prune_excluded_source_ids
  prune_source_ids = var.data_protection_group_prune_source_ids
}
*/

// Data source is not linked to a resource instance
// Uncomment if an existing data source instance exists
/*
// Create protection_groups data source
data "ibm_protection_groups" "protection_groups_instance" {
  request_initiator_type = var.protection_groups_request_initiator_type
  ids = var.protection_groups_ids
  names = var.protection_groups_names
  policy_ids = var.protection_groups_policy_ids
  storage_domain_id = var.protection_groups_storage_domain_id
  include_groups_with_datalock_only = var.protection_groups_include_groups_with_datalock_only
  environments = var.protection_groups_environments
  is_active = var.protection_groups_is_active
  is_deleted = var.protection_groups_is_deleted
  is_paused = var.protection_groups_is_paused
  last_run_local_backup_status = var.protection_groups_last_run_local_backup_status
  last_run_replication_status = var.protection_groups_last_run_replication_status
  last_run_archival_status = var.protection_groups_last_run_archival_status
  last_run_cloud_spin_status = var.protection_groups_last_run_cloud_spin_status
  last_run_any_status = var.protection_groups_last_run_any_status
  is_last_run_sla_violated = var.protection_groups_is_last_run_sla_violated
  tenant_ids = var.protection_groups_tenant_ids
  include_tenants = var.protection_groups_include_tenants
  include_last_run_info = var.protection_groups_include_last_run_info
  prune_excluded_source_ids = var.protection_groups_prune_excluded_source_ids
  prune_source_ids = var.protection_groups_prune_source_ids
  use_cached_data = var.protection_groups_use_cached_data
}
*/

// Data source is not linked to a resource instance
// Uncomment if an existing data source instance exists
/*
// Create protection_group_run data source
data "ibm_protection_group_run" "protection_group_run_instance" {
  protection_group_run_id = var.protection_group_run_protection_group_run_id
  run_id = var.protection_group_run_run_id
  request_initiator_type = var.protection_group_run_request_initiator_type
  tenant_ids = var.protection_group_run_tenant_ids
  include_tenants = var.protection_group_run_include_tenants
  include_object_details = var.protection_group_run_include_object_details
  use_cached_data = var.protection_group_run_use_cached_data
}
*/

// Data source is not linked to a resource instance
// Uncomment if an existing data source instance exists
/*
// Create protection_group_runs data source
data "ibm_protection_group_runs" "protection_group_runs_instance" {
  protection_group_runs_id = var.protection_group_runs_protection_group_runs_id
  request_initiator_type = var.protection_group_runs_request_initiator_type
  run_id = var.protection_group_runs_run_id
  start_time_usecs = var.protection_group_runs_start_time_usecs
  end_time_usecs = var.protection_group_runs_end_time_usecs
  tenant_ids = var.protection_group_runs_tenant_ids
  include_tenants = var.protection_group_runs_include_tenants
  run_types = var.protection_group_runs_run_types
  include_object_details = var.protection_group_runs_include_object_details
  local_backup_run_status = var.protection_group_runs_local_backup_run_status
  replication_run_status = var.protection_group_runs_replication_run_status
  archival_run_status = var.protection_group_runs_archival_run_status
  cloud_spin_run_status = var.protection_group_runs_cloud_spin_run_status
  num_runs = var.protection_group_runs_num_runs
  exclude_non_restorable_runs = var.protection_group_runs_exclude_non_restorable_runs
  run_tags = var.protection_group_runs_run_tags
  use_cached_data = var.protection_group_runs_use_cached_data
  filter_by_end_time = var.protection_group_runs_filter_by_end_time
  snapshot_target_types = var.protection_group_runs_snapshot_target_types
  only_return_successful_copy_run = var.protection_group_runs_only_return_successful_copy_run
  filter_by_copy_task_end_time = var.protection_group_runs_filter_by_copy_task_end_time
}
*/

// Data source is not linked to a resource instance
// Uncomment if an existing data source instance exists
/*
// Create protection_policies data source
data "ibm_protection_policies" "protection_policies_instance" {
  request_initiator_type = var.protection_policies_request_initiator_type
  ids = var.protection_policies_ids
  policy_names = var.protection_policies_policy_names
  tenant_ids = var.protection_policies_tenant_ids
  include_tenants = var.protection_policies_include_tenants
  types = var.protection_policies_types
  exclude_linked_policies = var.protection_policies_exclude_linked_policies
  include_replicated_policies = var.protection_policies_include_replicated_policies
  include_stats = var.protection_policies_include_stats
}
*/

// Data source is not linked to a resource instance
// Uncomment if an existing data source instance exists
/*
// Create protection_policy data source
data "ibm_protection_policy" "protection_policy_instance" {
  protection_policy_id = var.data_protection_policy_protection_policy_id
  request_initiator_type = var.data_protection_policy_request_initiator_type
}
*/

// Data source is not linked to a resource instance
// Uncomment if an existing data source instance exists
/*
// Create protection_run_summary data source
data "ibm_protection_run_summary" "protection_run_summary_instance" {
  start_time_usecs = var.protection_run_summary_start_time_usecs
  end_time_usecs = var.protection_run_summary_end_time_usecs
  run_status = var.protection_run_summary_run_status
}
*/

// Data source is not linked to a resource instance
// Uncomment if an existing data source instance exists
/*
// Create protection_sources data source
data "ibm_protection_sources" "protection_sources_instance" {
  request_initiator_type = var.protection_sources_request_initiator_type
  tenant_ids = var.protection_sources_tenant_ids
  include_tenants = var.protection_sources_include_tenants
  include_source_credentials = var.protection_sources_include_source_credentials
  encryption_key = var.protection_sources_encryption_key
}
*/

// Data source is not linked to a resource instance
// Uncomment if an existing data source instance exists
/*
// Create recovery data source
data "ibm_recovery" "recovery_instance" {
  recovery_id = var.data_recovery_recovery_id
  include_tenants = var.data_recovery_include_tenants
}
*/

// Data source is not linked to a resource instance
// Uncomment if an existing data source instance exists
/*
// Create recoveries data source
data "ibm_recoveries" "recoveries_instance" {
  ids = var.recoveries_ids
  return_only_child_recoveries = var.recoveries_return_only_child_recoveries
  tenant_ids = var.recoveries_tenant_ids
  include_tenants = var.recoveries_include_tenants
  start_time_usecs = var.recoveries_start_time_usecs
  end_time_usecs = var.recoveries_end_time_usecs
  storage_domain_id = var.recoveries_storage_domain_id
  snapshot_target_type = var.recoveries_snapshot_target_type
  archival_target_type = var.recoveries_archival_target_type
  snapshot_environments = var.recoveries_snapshot_environments
  status = var.recoveries_status
  recovery_actions = var.recoveries_recovery_actions
}
*/

// Data source is not linked to a resource instance
// Uncomment if an existing data source instance exists
/*
// Create source_registrations data source
data "ibm_source_registrations" "source_registrations_instance" {
  ids = var.source_registrations_ids
  tenant_ids = var.source_registrations_tenant_ids
  include_tenants = var.source_registrations_include_tenants
  include_source_credentials = var.source_registrations_include_source_credentials
  encryption_key = var.source_registrations_encryption_key
  use_cached_data = var.source_registrations_use_cached_data
}
*/

// Data source is not linked to a resource instance
// Uncomment if an existing data source instance exists
/*
// Create source_registration data source
data "ibm_source_registration" "source_registration_instance" {
  source_registration_id = var.data_source_registration_source_registration_id
  request_initiator_type = var.data_source_registration_request_initiator_type
}
*/
