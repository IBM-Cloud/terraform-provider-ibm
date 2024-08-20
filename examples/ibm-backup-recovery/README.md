# Examples for IBM Backup recovery API

These examples illustrate how to use the resources and data sources associated with IBM Backup recovery API.

The following resources are supported:
* ibm_protection_group_run_request
* ibm_recovery_download_files_folders
* ibm_perform_action_on_protection_group_run_request
* ibm_protection_group
* ibm_protection_policy
* ibm_recovery
* ibm_search_indexed_object
* ibm_source_registration
* ibm_update_protection_group_run_request
* ibm_protection_group_state

The following data sources are supported:
* ibm_run_debug_logs
* ibm_object_run_debug_logs
* ibm_run_error_report
* ibm_runs_report
* ibm_recovery_debug_logs
* ibm_recovery_download_messages
* ibm_recovery_download_files
* ibm_recovery_fetch_uptier_data
* ibm_protection_run_progress
* ibm_protection_run_stat
* ibm_search_objects
* ibm_search_protected_objects
* ibm_protection_group
* ibm_protection_groups
* ibm_protection_group_run
* ibm_protection_group_runs
* ibm_protection_policies
* ibm_protection_policy
* ibm_protection_run_summary
* ibm_protection_sources
* ibm_recovery
* ibm_recoveries
* ibm_source_registrations
* ibm_source_registration

## Usage

To run this example, execute the following commands:

```bash
$ terraform init
$ terraform plan
$ terraform apply
```

Run `terraform destroy` when you don't need these resources.

## IBM Backup recovery API resources

### Resource: ibm_protection_group_run_request

```hcl
resource "ibm_protection_group_run_request" "protection_group_run_request_instance" {
  run_type = var.protection_group_run_request_run_type
  objects = var.protection_group_run_request_objects
  targets_config = var.protection_group_run_request_targets_config
  uda_params = var.protection_group_run_request_uda_params
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| ibmcloud\_api\_key | IBM Cloud API key | `string` | true |
| run_type | Type of protection run. 'kRegular' indicates an incremental (CBT) backup. Incremental backups utilizing CBT (if supported) are captured of the target protection objects. The first run of a kRegular schedule captures all the blocks. 'kFull' indicates a full (no CBT) backup. A complete backup (all blocks) of the target protection objects are always captured and Change Block Tracking (CBT) is not utilized. 'kLog' indicates a Database Log backup. Capture the database transaction logs to allow rolling back to a specific point in time. 'kSystem' indicates system volume backup. It produces an image for bare metal recovery. | `string` | true |
| objects | Specifies the list of objects to be protected by this Protection Group run. These can be leaf objects or non-leaf objects in the protection hierarchy. This must be specified only if a subset of objects from the Protection Groups needs to be protected. | `list()` | false |
| targets_config | Specifies the replication and archival targets. | `` | false |
| uda_params | Specifies the parameters for Universal Data Adapter protection run. | `` | false |

#### Outputs

| Name | Description |
|------|-------------|
| run_id | The unique ID. |

### Resource: ibm_recovery_download_files_folders

```hcl
resource "ibm_recovery_download_files_folders" "recovery_download_files_folders_instance" {
  name = var.recovery_download_files_folders_name
  object = var.recovery_download_files_folders_object
  parent_recovery_id = var.recovery_download_files_folders_parent_recovery_id
  files_and_folders = var.recovery_download_files_folders_files_and_folders
  glacier_retrieval_type = var.recovery_download_files_folders_glacier_retrieval_type
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| ibmcloud\_api\_key | IBM Cloud API key | `string` | true |
| name | Specifies the name of the recovery task. This field must be set and must be a unique name. | `string` | true |
| object | Specifies the common snapshot parameters for a protected object. | `` | true |
| parent_recovery_id | If current recovery is child task triggered through another parent recovery operation, then this field will specify the id of the parent recovery. | `string` | false |
| files_and_folders | Specifies the list of files and folders to download. | `list()` | true |
| glacier_retrieval_type | Specifies the glacier retrieval type when restoring or downloding files or folders from a Glacier-based cloud snapshot. | `string` | false |

### Resource: ibm_perform_action_on_protection_group_run_request

```hcl
resource "ibm_perform_action_on_protection_group_run_request" "perform_action_on_protection_group_run_request_instance" {
  action = var.perform_action_on_protection_group_run_request_action
  pause_params = var.perform_action_on_protection_group_run_request_pause_params
  resume_params = var.perform_action_on_protection_group_run_request_resume_params
  cancel_params = var.perform_action_on_protection_group_run_request_cancel_params
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| ibmcloud\_api\_key | IBM Cloud API key | `string` | true |
| action | Specifies the type of the action which will be performed on protection runs. | `string` | true |
| pause_params | Specifies the pause action params for a protection run. | `list()` | false |
| resume_params | Specifies the resume action params for a protection run. | `list()` | false |
| cancel_params | Specifies the cancel action params for a protection run. | `list()` | false |

#### Outputs

| Name | Description |
|------|-------------|
| run_id | The unique ID. |

### Resource: ibm_protection_group

```hcl
resource "ibm_protection_group" "protection_group_instance" {
  name = var.protection_group_name
  policy_id = var.protection_group_policy_id
  priority = var.protection_group_priority
  storage_domain_id = var.protection_group_storage_domain_id
  description = var.protection_group_description
  start_time = var.protection_group_start_time
  end_time_usecs = var.protection_group_end_time_usecs
  last_modified_timestamp_usecs = var.protection_group_last_modified_timestamp_usecs
  alert_policy = var.protection_group_alert_policy
  sla = var.protection_group_sla
  qos_policy = var.protection_group_qos_policy
  abort_in_blackouts = var.protection_group_abort_in_blackouts
  pause_in_blackouts = var.protection_group_pause_in_blackouts
  is_paused = var.protection_group_is_paused
  environment = var.protection_group_environment
  advanced_configs = var.protection_group_advanced_configs
  physical_params = var.protection_group_physical_params
  oracle_params = var.protection_group_oracle_params
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| ibmcloud\_api\_key | IBM Cloud API key | `string` | true |
| name | Specifies the name of the Protection Group. | `string` | true |
| policy_id | Specifies the unique id of the Protection Policy associated with the Protection Group. The Policy provides retry settings Protection Schedules, Priority, SLA, etc. | `string` | true |
| priority | Specifies the priority of the Protection Group. | `string` | false |
| storage_domain_id | Specifies the Storage Domain (View Box) ID where this Protection Group writes data. | `number` | false |
| description | Specifies a description of the Protection Group. | `string` | false |
| start_time | Specifies the time of day. Used for scheduling purposes. | `` | false |
| end_time_usecs | Specifies the end time in micro seconds for this Protection Group. If this is not specified, the Protection Group won't be ended. | `number` | false |
| last_modified_timestamp_usecs | Specifies the last time this protection group was updated. If this is passed into a PUT request, then the backend will validate that the timestamp passed in matches the time that the protection group was actually last modified. If the two timestamps do not match, then the request will be rejected with a stale error. | `number` | false |
| alert_policy | Specifies a policy for alerting users of the status of a Protection Group. | `` | false |
| sla | Specifies the SLA parameters for this Protection Group. | `list()` | false |
| qos_policy | Specifies whether the Protection Group will be written to HDD or SSD. | `string` | false |
| abort_in_blackouts | Specifies whether currently executing jobs should abort if a blackout period specified by a policy starts. Available only if the selected policy has at least one blackout period. Default value is false. | `bool` | false |
| pause_in_blackouts | Specifies whether currently executing jobs should be paused if a blackout period specified by a policy starts. Available only if the selected policy has at least one blackout period. Default value is false. This field should not be set to true if 'abortInBlackouts' is sent as true. | `bool` | false |
| is_paused | Specifies if the the Protection Group is paused. New runs are not scheduled for the paused Protection Groups. Active run if any is not impacted. | `bool` | false |
| environment | Specifies the environment of the Protection Group. | `string` | true |
| advanced_configs | Specifies the advanced configuration for a protection job. | `list()` | false |
| physical_params |  | `` | false |
| oracle_params | Specifies the parameters to create Oracle Protection Group. | `` | false |

#### Outputs

| Name | Description |
|------|-------------|
| cluster_id | Specifies the cluster ID. |
| region_id | Specifies the region ID. |
| is_active | Specifies if the Protection Group is active or not. |
| is_deleted | Specifies if the Protection Group has been deleted. |
| last_run | Specifies the parameters which are common between Protection Group runs of all Protection Groups. |
| permissions | Specifies the list of tenants that have permissions for this protection group. |
| is_protect_once | Specifies if the the Protection Group is using a protect once type of policy. This field is helpful to identify run happen for this group. |
| missing_entities | Specifies the Information about missing entities. |
| invalid_entities | Specifies the Information about invalid entities. An entity will be considered invalid if it is part of an active protection group but has lost compatibility for the given backup type. |
| num_protected_objects | Specifies the number of protected objects of the Protection Group. |

### Resource: ibm_protection_policy

```hcl
resource "ibm_protection_policy" "protection_policy_instance" {
  name = var.protection_policy_name
  backup_policy = var.protection_policy_backup_policy
  description = var.protection_policy_description
  blackout_window = var.protection_policy_blackout_window
  extended_retention = var.protection_policy_extended_retention
  remote_target_policy = var.protection_policy_remote_target_policy
  cascaded_targets_config = var.protection_policy_cascaded_targets_config
  retry_options = var.protection_policy_retry_options
  data_lock = var.protection_policy_data_lock
  version = var.protection_policy_version
  is_cbs_enabled = var.protection_policy_is_cbs_enabled
  last_modification_time_usecs = var.protection_policy_last_modification_time_usecs
  template_id = var.protection_policy_template_id
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| ibmcloud\_api\_key | IBM Cloud API key | `string` | true |
| name | Specifies the name of the Protection Policy. | `string` | true |
| backup_policy | Specifies the backup schedule and retentions of a Protection Policy. | `` | true |
| description | Specifies the description of the Protection Policy. | `string` | false |
| blackout_window | List of Blackout Windows. If specified, this field defines blackout periods when new Group Runs are not started. If a Group Run has been scheduled but not yet executed and the blackout period starts, the behavior depends on the policy field AbortInBlackoutPeriod. | `list()` | false |
| extended_retention | Specifies additional retention policies that should be applied to the backup snapshots. A backup snapshot will be retained up to a time that is the maximum of all retention policies that are applicable to it. | `list()` | false |
| remote_target_policy | Specifies the replication, archival and cloud spin targets of Protection Policy. | `` | false |
| cascaded_targets_config | Specifies the configuration for cascaded replications. Using cascaded replication, replication cluster(Rx) can further replicate and archive the snapshot copies to further targets. Its recommended to create cascaded configuration where protection group will be created. | `list()` | false |
| retry_options | Retry Options of a Protection Policy when a Protection Group run fails. | `` | false |
| data_lock | This field is now deprecated. Please use the DataLockConfig in the backup retention. | `string` | false |
| version | Specifies the current policy verison. Policy version is incremented for optionally supporting new features and differentialting across releases. | `number` | false |
| is_cbs_enabled | Specifies true if Calender Based Schedule is supported by client. Default value is assumed as false for this feature. | `bool` | false |
| last_modification_time_usecs | Specifies the last time this Policy was updated. If this is passed into a PUT request, then the backend will validate that the timestamp passed in matches the time that the policy was actually last modified. If the two timestamps do not match, then the request will be rejected with a stale error. | `number` | false |
| template_id | Specifies the parent policy template id to which the policy is linked to. This field is set only when policy is created from template. | `string` | false |

#### Outputs

| Name | Description |
|------|-------------|
| is_usable | This field is set to true if the linked policy which is internally created from a policy templates qualifies as usable to create more policies on the cluster. If the linked policy is partially filled and can not create a working policy then this field will be set to false. In case of normal policy created on the cluster, this field wont be populated. |
| is_replicated | This field is set to true when policy is the replicated policy. |
| num_protection_groups | Specifies the number of protection groups using the protection policy. |
| num_protected_objects | Specifies the number of protected objects using the protection policy. |

### Resource: ibm_recovery

```hcl
resource "ibm_recovery" "recovery_instance" {
  request_initiator_type = var.recovery_request_initiator_type
  name = var.recovery_name
  snapshot_environment = var.recovery_snapshot_environment
  physical_params = var.recovery_physical_params
  oracle_params = var.recovery_oracle_params
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| ibmcloud\_api\_key | IBM Cloud API key | `string` | true |
| request_initiator_type | Specifies the type of request from UI, which is used for services like magneto to determine the priority of requests. | `string` | false |
| name | Specifies the name of the Recovery. | `string` | true |
| snapshot_environment | Specifies the type of snapshot environment for which the Recovery was performed. | `string` | true |
| physical_params | Specifies the recovery options specific to Physical environment. | `` | false |
| oracle_params | Specifies the recovery options specific to oracle environment. | `` | false |

#### Outputs

| Name | Description |
|------|-------------|
| start_time_usecs | Specifies the start time of the Recovery in Unix timestamp epoch in microseconds. |
| end_time_usecs | Specifies the end time of the Recovery in Unix timestamp epoch in microseconds. This field will be populated only after Recovery is finished. |
| status | Status of the Recovery. 'Running' indicates that the Recovery is still running. 'Canceled' indicates that the Recovery has been cancelled. 'Canceling' indicates that the Recovery is in the process of being cancelled. 'Failed' indicates that the Recovery has failed. 'Succeeded' indicates that the Recovery has finished successfully. 'SucceededWithWarning' indicates that the Recovery finished successfully, but there were some warning messages. 'Skipped' indicates that the Recovery task was skipped. |
| progress_task_id | Progress monitor task id for Recovery. |
| recovery_action | Specifies the type of recover action. |
| permissions | Specifies the list of tenants that have permissions for this recovery. |
| creation_info | Specifies the information about the creation of the protection group or recovery. |
| can_tear_down | Specifies whether it's possible to tear down the objects created by the recovery. |
| tear_down_status | Specifies the status of the tear down operation. This is only set when the canTearDown is set to true. 'DestroyScheduled' indicates that the tear down is ready to schedule. 'Destroying' indicates that the tear down is still running. 'Destroyed' indicates that the tear down succeeded. 'DestroyError' indicates that the tear down failed. |
| tear_down_message | Specifies the error message about the tear down operation if it fails. |
| messages | Specifies messages about the recovery. |
| is_parent_recovery | Specifies whether the current recovery operation has created child recoveries. This is currently used in SQL recovery where multiple child recoveries can be tracked under a common/parent recovery. |
| parent_recovery_id | If current recovery is child recovery triggered by another parent recovery operation, then this field willt specify the id of the parent recovery. |
| retrieve_archive_tasks | Specifies the list of persistent state of a retrieve of an archive task. |
| is_multi_stage_restore | Specifies whether the current recovery operation is a multi-stage restore operation. This is currently used by VMware recoveres for the migration/hot-standby use case. |

### Resource: ibm_search_indexed_object

```hcl
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
  files = var.search_indexed_object_files
  public_folders = var.search_indexed_object_public_folders
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| ibmcloud\_api\_key | IBM Cloud API key | `string` | true |
| protection_group_ids | Specifies a list of Protection Group ids to filter the indexed objects. If specified, the objects indexed by specified Protection Group ids will be returned. | `list(string)` | false |
| storage_domain_ids | Specifies the Storage Domain ids to filter indexed objects for which Protection Groups are writing data to Cohesity Views on the specified Storage Domains. | `list(number)` | false |
| tenant_id | TenantId contains id of the tenant for which objects are to be returned. | `string` | false |
| include_tenants | If true, the response will include objects which belongs to all tenants which the current user has permission to see. Default value is false. | `bool` | false |
| tags | "This field is deprecated. Please use mightHaveTagIds.". | `list(string)` | false |
| snapshot_tags | "This field is deprecated. Please use mightHaveSnapshotTagIds.". | `list(string)` | false |
| must_have_tag_ids | Specifies tags which must be all present in the document. | `list(string)` | false |
| might_have_tag_ids | Specifies list of tags, one or more of which might be present in the document. These are OR'ed together and the resulting criteria AND'ed with the rest of the query. | `list(string)` | false |
| must_have_snapshot_tag_ids | Specifies snapshot tags which must be all present in the document. | `list(string)` | false |
| might_have_snapshot_tag_ids | Specifies list of snapshot tags, one or more of which might be present in the document. These are OR'ed together and the resulting criteria AND'ed with the rest of the query. | `list(string)` | false |
| pagination_cookie | Specifies the pagination cookie with which subsequent parts of the response can be fetched. | `string` | false |
| count | Specifies the number of indexed objects to be fetched for the specified pagination cookie. | `number` | false |
| object_type | Specifies the object type to be searched for. | `string` | true |
| use_cached_data | Specifies whether we can serve the GET request from the read replica cache. There is a lag of 15 seconds between the read replica and primary data source. | `bool` | false |
| files | Specifies the request parameters to search for files and file folders. | `` | false |
| public_folders | Specifies the request parameters to search for Public Folder items. | `` | false |

### Resource: ibm_source_registration

```hcl
resource "ibm_source_registration" "source_registration_instance" {
  environment = var.source_registration_environment
  name = var.source_registration_name
  connection_id = var.source_registration_connection_id
  connections = var.source_registration_connections
  connector_group_id = var.source_registration_connector_group_id
  advanced_configs = var.source_registration_advanced_configs
  physical_params = var.source_registration_physical_params
  oracle_params = var.source_registration_oracle_params
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| ibmcloud\_api\_key | IBM Cloud API key | `string` | true |
| environment | Specifies the environment type of the Protection Source. | `string` | true |
| name | The user specified name for this source. | `string` | false |
| connection_id | Specifies the id of the connection from where this source is reachable. This should only be set for a source being registered by a tenant user. This field will be depricated in future. Use connections field. | `number` | false |
| connections | Specfies the list of connections for the source. | `list()` | false |
| connector_group_id | Specifies the connector group id of connector groups. | `number` | false |
| advanced_configs | Specifies the advanced configuration for a protection source. | `list()` | false |
| physical_params | Physical Params params. | `` | false |
| oracle_params | Physical Params params. | `` | false |

#### Outputs

| Name | Description |
|------|-------------|
| source_id | ID of top level source object discovered after the registration. |
| source_info | Specifies information about an object. |
| authentication_status | Specifies the status of the authentication during the registration of a Protection Source. 'Pending' indicates the authentication is in progress. 'Scheduled' indicates the authentication is scheduled. 'Finished' indicates the authentication is completed. 'RefreshInProgress' indicates the refresh is in progress. |
| registration_time_msecs | Specifies the time when the source was registered in milliseconds. |
| last_refreshed_time_msecs | Specifies the time when the source was last refreshed in milliseconds. |

### Resource: ibm_update_protection_group_run_request

```hcl
resource "ibm_update_protection_group_run_request" "update_protection_group_run_request_instance" {
  update_protection_group_run_params = var.update_protection_group_run_request_update_protection_group_run_params
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| ibmcloud\_api\_key | IBM Cloud API key | `string` | true |
| update_protection_group_run_params |  | `list()` | true |

#### Outputs

| Name | Description |
|------|-------------|
| run_id | The unique ID. |

### Resource: ibm_protection_group_state

```hcl
resource "ibm_protection_group_state" "protection_group_state_instance" {
  action = var.protection_group_state_action
  ids = var.protection_group_state_ids
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| ibmcloud\_api\_key | IBM Cloud API key | `string` | true |
| action | Specifies the action to be performed on all the specfied Protection Groups. 'kActivate' specifies that Protection Group should be activated. 'kDeactivate' sepcifies that Protection Group should be deactivated. 'kPause' specifies that Protection Group should be paused. 'kResume' specifies that Protection Group should be resumed. | `string` | true |
| ids | Specifies a list of Protection Group ids for which the state should change. | `list(string)` | true |

## IBM Backup recovery API data sources

### Data source: ibm_run_debug_logs

```hcl
data "ibm_run_debug_logs" "run_debug_logs_instance" {
  run_debug_logs_id = var.run_debug_logs_run_debug_logs_id
  run_id = var.run_debug_logs_run_id
  object_id = var.run_debug_logs_object_id
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| run_debug_logs_id | Specifies a unique id of the Protection Group. | `string` | true |
| run_id | Specifies a unique run id of the Protection Group run. | `string` | true |
| object_id | Specifies the id of the object for which debug logs are to be returned. | `string` | false |

### Data source: ibm_object_run_debug_logs

```hcl
data "ibm_object_run_debug_logs" "object_run_debug_logs_instance" {
  object_run_debug_logs_id = var.object_run_debug_logs_object_run_debug_logs_id
  run_id = var.object_run_debug_logs_run_id
  object_id = var.object_run_debug_logs_object_id
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| object_run_debug_logs_id | Specifies a unique id of the Protection Group. | `string` | true |
| run_id | Specifies a unique run id of the Protection Group run. | `string` | true |
| object_id | Specifies the id of the object for which debug logs are to be returned. | `string` | true |

### Data source: ibm_run_error_report

```hcl
data "ibm_run_error_report" "run_error_report_instance" {
  run_error_report_id = var.run_error_report_run_error_report_id
  run_id = var.run_error_report_run_id
  object_id = var.run_error_report_object_id
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| run_error_report_id | Specifies a unique id of the Protection Group. | `string` | true |
| run_id | Specifies a unique run id of the Protection Group run. | `string` | true |
| object_id | Specifies the id of the object for which errors/warnings are to be returned. | `string` | true |

### Data source: ibm_runs_report

```hcl
data "ibm_runs_report" "runs_report_instance" {
  runs_report_id = var.runs_report_runs_report_id
  run_id = var.runs_report_run_id
  object_id = var.runs_report_object_id
  file_type = var.runs_report_file_type
  name = var.runs_report_name
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| runs_report_id | Specifies a unique id of the Protection Group. | `string` | true |
| run_id | Specifies a unique run id of the Protection Group run. | `string` | true |
| object_id | Specifies the id of the object for which errors/warnings are to be returned. | `string` | true |
| file_type | Specifies the downloaded type, i.e: success_files_list, default: success_files_list. | `string` | false |
| name | Specifies the name of the source being backed up. | `string` | false |

### Data source: ibm_recovery_debug_logs

```hcl
data "ibm_recovery_debug_logs" "recovery_debug_logs_instance" {
  recovery_debug_logs_id = var.recovery_debug_logs_recovery_debug_logs_id
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| recovery_debug_logs_id | Specifies the id of a Recovery job. | `string` | true |

### Data source: ibm_recovery_download_messages

```hcl
data "ibm_recovery_download_messages" "recovery_download_messages_instance" {
  recovery_download_messages_id = var.recovery_download_messages_recovery_download_messages_id
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| recovery_download_messages_id | Specifies a unique ID of a Recovery. | `string` | true |

### Data source: ibm_recovery_download_files

```hcl
data "ibm_recovery_download_files" "recovery_download_files_instance" {
  recovery_download_files_id = var.recovery_download_files_recovery_download_files_id
  start_offset = var.recovery_download_files_start_offset
  length = var.recovery_download_files_length
  file_type = var.recovery_download_files_file_type
  source_name = var.recovery_download_files_source_name
  start_time = var.recovery_download_files_start_time
  include_tenants = var.recovery_download_files_include_tenants
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| recovery_download_files_id | Specifies the id of a Recovery. | `string` | true |
| start_offset | Specifies the start offset of file chunk to be downloaded. | `number` | false |
| length | Specifies the length of bytes to download. This can not be greater than 8MB (8388608 byets). | `number` | false |
| file_type | Specifies the downloaded type, i.e: error, success_files_list. | `string` | false |
| source_name | Specifies the name of the source on which restore is done. | `string` | false |
| start_time | Specifies the start time of restore task. | `string` | false |
| include_tenants | Specifies if objects of all the organizations under the hierarchy of the logged in user's organization should be returned. | `bool` | false |

### Data source: ibm_recovery_fetch_uptier_data

```hcl
data "ibm_recovery_fetch_uptier_data" "recovery_fetch_uptier_data_instance" {
  archive_u_id = var.recovery_fetch_uptier_data_archive_u_id
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| archive_u_id | Archive UID of the current restore. | `string` | true |

#### Outputs

| Name | Description |
|------|-------------|
| data_size | Specifies the amount of data in bytes estimated to be uptiered as part of the current restore job. |

### Data source: ibm_protection_run_progress

```hcl
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
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| run_id | Specifies a unique run id of the Protection Run. | `string` | true |
| objects | Specifies the objects whose progress will be returned. This only applies to protection group runs and will be ignored for object runs. If the objects are specified, the run progress will not be returned and only the progress of the specified objects will be returned. | `list(number)` | false |
| tenant_ids | TenantIds contains ids of the tenants for which the run is to be returned. | `list(string)` | false |
| include_tenants | If true, the response will include Protection Group Runs which were created by all tenants which the current user has permission to see. If false, then only Protection Groups created by the current user will be returned. If it's not specified, it is true by default. | `bool` | false |
| include_finished_tasks | Specifies whether to return finished tasks. By default only active tasks are returned. | `bool` | false |
| start_time_usecs | Specifies the time after which the progress task starts in Unix epoch Timestamp(in microseconds). | `number` | false |
| end_time_usecs | Specifies the time before which the progress task ends in Unix epoch Timestamp(in microseconds). | `number` | false |
| max_tasks_num | Specifies the maximum number of tasks to return. | `number` | false |
| exclude_object_details | Specifies whether to return objects. By default all the task tree are returned. | `bool` | false |
| include_event_logs | Specifies whether to include event logs. | `bool` | false |
| max_log_level | Specifies the number of levels till which to fetch the event logs. This is applicable only when includeEventLogs is true. | `number` | false |
| run_task_path | Specifies the task path of the run or object run. This is applicable only if progress of a protection group with one or more object is required.If provided this will be used to fetch progress details directly without looking actual task path of the object. Objects field is stil expected else it changes the response format. | `string` | false |
| object_task_paths | Specifies the object level task path. This relates to the objectID. If provided this will take precedence over the objects, and will be used to fetch progress details directly without looking actuall task path of the object. | `list(string)` | false |

#### Outputs

| Name | Description |
|------|-------------|
| local_run | Specifies the progress of a local backup run. |
| archival_run | Progress for the archival run. |
| replication_run | Progress for the replication run. |

### Data source: ibm_protection_run_stat

```hcl
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
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| run_id | Specifies a unique run id of the Protection Run. | `string` | true |
| objects | Specifies the objects whose stats will be returned. This only applies to protection group runs and will be ignored for object runs. If the objects are specified, the run stats will not be returned and only the stats of the specified objects will be returned. | `list(number)` | false |
| tenant_ids | TenantIds contains ids of the tenants for which the run is to be returned. | `list(string)` | false |
| include_tenants | If true, the response will include Protection Group Runs which were created by all tenants which the current user has permission to see. If false, then only Protection Groups created by the current user will be returned. If it's not specified, it is true by default. | `bool` | false |
| include_finished_tasks | Specifies whether to return finished tasks. By default only active tasks are returned. | `bool` | false |
| start_time_usecs | Specifies the time after which the stats task starts in Unix epoch Timestamp(in microseconds). | `number` | false |
| end_time_usecs | Specifies the time before which the stats task ends in Unix epoch Timestamp(in microseconds). | `number` | false |
| max_tasks_num | Specifies the maximum number of tasks to return. | `number` | false |
| exclude_object_details | Specifies whether to return objects. By default all the task tree are returned. | `bool` | false |
| run_task_path | Specifies the task path of the run or object run. This is applicable only if stats of a protection group with one or more object is required. If provided this will be used to fetch stats details directly without looking actual task path of the object. Objects field is stil expected else it changes the response format. | `string` | false |
| object_task_paths | Specifies the object level task path. This relates to the objectID. If provided this will take precedence over the objects, and will be used to fetch stats details directly without looking actuall task path of the object. | `list(string)` | false |

#### Outputs

| Name | Description |
|------|-------------|
| local_run | Specifies the stats of a local backup run. |
| archival_run | Stats for the archival run. |

### Data source: ibm_search_objects

```hcl
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
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| request_initiator_type | Specifies the type of request from UI, which is used for services like magneto to determine the priority of requests. | `string` | false |
| search_string | Specifies the search string to filter the objects. This search string will be applicable for objectnames. User can specify a wildcard character '*' as a suffix to a string where all object names are matched with the prefix string. For example, if vm1 and vm2 are the names of objects, user can specify vm* to list the objects. If not specified, then all the objects will be returned which will match other filtering criteria. | `string` | false |
| environments | Specifies the environment type to filter objects. | `list(string)` | false |
| protection_types | Specifies the protection type to filter objects. | `list(string)` | false |
| tenant_ids | TenantIds contains ids of the tenants for which objects are to be returned. | `list(string)` | false |
| include_tenants | If true, the response will include Objects which belongs to all tenants which the current user has permission to see. | `bool` | false |
| protection_group_ids | Specifies a list of Protection Group ids to filter the objects. If specified, the objects protected by specified Protection Group ids will be returned. | `list(string)` | false |
| object_ids | Specifies a list of Object ids to filter. | `list(number)` | false |
| os_types | Specifies the operating system types to filter objects on. | `list(string)` | false |
| source_ids | Specifies a list of Protection Source object ids to filter the objects. If specified, the object which are present in those Sources will be returned. | `list(number)` | false |
| source_uuids | Specifies a list of Protection Source object uuids to filter the objects. If specified, the object which are present in those Sources will be returned. | `list(string)` | false |
| is_protected | Specifies the protection status of objects. If set to true, only protected objects will be returned. If set to false, only unprotected objects will be returned. If not specified, all objects will be returned. | `bool` | false |
| is_deleted | If set to true, then objects which are deleted on atleast one cluster will be returned. If not set or set to false then objects which are registered on atleast one cluster are returned. | `bool` | false |
| last_run_status_list | Specifies a list of status of the object's last protection run. Only objects with last run status of these will be returned. | `list(string)` | false |
| region_ids | Specifies a list of region ids. Only records from clusters having these region ids will be returned. | `list(string)` | false |
| cluster_identifiers | Specifies the list of cluster identifiers. Format is clusterId:clusterIncarnationId. Only records from clusters having these identifiers will be returned. | `list(string)` | false |
| storage_domain_ids | Specifies the list of storage domain ids. Format is clusterId:clusterIncarnationId:storageDomainId. Only objects having protection in these storage domains will be returned. | `list(string)` | false |
| include_deleted_objects | Specifies whether to include deleted objects in response. These objects can't be protected but can be recovered. This field is deprecated. | `bool` | false |
| pagination_cookie | Specifies the pagination cookie with which subsequent parts of the response can be fetched. | `string` | false |
| count | Specifies the number of objects to be fetched for the specified pagination cookie. | `number` | false |
| must_have_tag_ids | Specifies tags which must be all present in the document. | `list(string)` | false |
| might_have_tag_ids | Specifies list of tags, one or more of which might be present in the document. These are OR'ed together and the resulting criteria AND'ed with the rest of the query. | `list(string)` | false |
| must_have_snapshot_tag_ids | Specifies snapshot tags which must be all present in the document. | `list(string)` | false |
| might_have_snapshot_tag_ids | Specifies list of snapshot tags, one or more of which might be present in the document. These are OR'ed together and the resulting criteria AND'ed with the rest of the query. | `list(string)` | false |

#### Outputs

| Name | Description |
|------|-------------|
| objects | Specifies the list of Objects. |

### Data source: ibm_search_protected_objects

```hcl
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
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| request_initiator_type | Specifies the type of request from UI, which is used for services like magneto to determine the priority of requests. | `string` | false |
| search_string | Specifies the search string to filter the objects. This search string will be applicable for objectnames and Protection Group names. User can specify a wildcard character '*' as a suffix to a string where all object and their Protection Group names are matched with the prefix string. For example, if vm1 and vm2 are the names of objects, user can specify vm* to list the objects. If not specified, then all the objects with Protection Groups will be returned which will match other filtering criteria. | `string` | false |
| environments | Specifies the environment type to filter objects. | `list(string)` | false |
| snapshot_actions | Specifies a list of recovery actions. Only snapshots that applies to these actions will be returned. | `list(string)` | false |
| object_action_key | Filter by ObjectActionKey, which uniquely represents protection of an object. An object can be protected in multiple ways but atmost once for a given combination of ObjectActionKey. When specified, latest snapshot info matching the objectActionKey is for corresponding object. | `string` | false |
| tenant_ids | TenantIds contains ids of the tenants for which objects are to be returned. | `list(string)` | false |
| include_tenants | If true, the response will include Objects which belongs to all tenants which the current user has permission to see. | `bool` | false |
| protection_group_ids | Specifies a list of Protection Group ids to filter the objects. If specified, the objects protected by specified Protection Group ids will be returned. | `list(string)` | false |
| object_ids | Specifies a list of Object ids to filter. | `list(number)` | false |
| storage_domain_ids | Specifies the Storage Domain ids to filter objects for which Protection Groups are writing data to Cohesity Views on the specified Storage Domains. | `list(number)` | false |
| sub_result_size | Specifies the size of objects to be fetched for a single subresult. | `number` | false |
| filter_snapshot_from_usecs | Specifies the timestamp in Unix time epoch in microseconds to filter the objects if the Object has a successful snapshot after this value. | `number` | false |
| filter_snapshot_to_usecs | Specifies the timestamp in Unix time epoch in microseconds to filter the objects if the Object has a successful snapshot before this value. | `number` | false |
| os_types | Specifies the operating system types to filter objects on. | `list(string)` | false |
| source_ids | Specifies a list of Protection Source object ids to filter the objects. If specified, the object which are present in those Sources will be returned. | `list(number)` | false |
| run_instance_ids | Specifies a list of run instance ids. If specified only objects belonging to the provided run id will be retunrned. | `list(number)` | false |
| cdp_protected_only | Specifies whether to only return the CDP protected objects. | `bool` | false |
| region_ids | Specifies a list of region ids. Only records from clusters having these region ids will be returned. | `list(string)` | false |
| use_cached_data | Specifies whether we can serve the GET request to the read replica cache cache. There is a lag of 15 seconds between the read replica and primary data source. | `bool` | false |

#### Outputs

| Name | Description |
|------|-------------|
| objects | Specifies the list of Protected Objects. |
| metadata | Specifies the metadata information about the Protection Groups, Protection Policy etc., for search result. |
| num_results | Specifies the total number of search results which matches the search criteria. |

### Data source: ibm_protection_group

```hcl
data "ibm_protection_group" "protection_group_instance" {
  protection_group_id = var.data_protection_group_protection_group_id
  request_initiator_type = var.data_protection_group_request_initiator_type
  include_last_run_info = var.data_protection_group_include_last_run_info
  prune_excluded_source_ids = var.data_protection_group_prune_excluded_source_ids
  prune_source_ids = var.data_protection_group_prune_source_ids
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| protection_group_id | Specifies a unique id of the Protection Group. | `string` | true |
| request_initiator_type | Specifies the type of request from UI, which is used for services like magneto to determine the priority of requests. | `string` | false |
| include_last_run_info | If true, the response will include last run info. If it is false or not specified, the last run info won't be returned. | `bool` | false |
| prune_excluded_source_ids | If true, the response will not include the list of excluded source IDs in groups that contain this field. This can be set to true in order to improve performance if excluded source IDs are not needed by the user. | `bool` | false |
| prune_source_ids | If true, the response will exclude the list of source IDs within the group specified. | `bool` | false |

#### Outputs

| Name | Description |
|------|-------------|
| name | Specifies the name of the Protection Group. |
| cluster_id | Specifies the cluster ID. |
| region_id | Specifies the region ID. |
| policy_id | Specifies the unique id of the Protection Policy associated with the Protection Group. The Policy provides retry settings Protection Schedules, Priority, SLA, etc. |
| priority | Specifies the priority of the Protection Group. |
| storage_domain_id | Specifies the Storage Domain (View Box) ID where this Protection Group writes data. |
| description | Specifies a description of the Protection Group. |
| start_time | Specifies the time of day. Used for scheduling purposes. |
| end_time_usecs | Specifies the end time in micro seconds for this Protection Group. If this is not specified, the Protection Group won't be ended. |
| last_modified_timestamp_usecs | Specifies the last time this protection group was updated. If this is passed into a PUT request, then the backend will validate that the timestamp passed in matches the time that the protection group was actually last modified. If the two timestamps do not match, then the request will be rejected with a stale error. |
| alert_policy | Specifies a policy for alerting users of the status of a Protection Group. |
| sla | Specifies the SLA parameters for this Protection Group. |
| qos_policy | Specifies whether the Protection Group will be written to HDD or SSD. |
| abort_in_blackouts | Specifies whether currently executing jobs should abort if a blackout period specified by a policy starts. Available only if the selected policy has at least one blackout period. Default value is false. |
| pause_in_blackouts | Specifies whether currently executing jobs should be paused if a blackout period specified by a policy starts. Available only if the selected policy has at least one blackout period. Default value is false. This field should not be set to true if 'abortInBlackouts' is sent as true. |
| is_active | Specifies if the Protection Group is active or not. |
| is_deleted | Specifies if the Protection Group has been deleted. |
| is_paused | Specifies if the the Protection Group is paused. New runs are not scheduled for the paused Protection Groups. Active run if any is not impacted. |
| environment | Specifies the environment of the Protection Group. |
| last_run | Specifies the parameters which are common between Protection Group runs of all Protection Groups. |
| permissions | Specifies the list of tenants that have permissions for this protection group. |
| is_protect_once | Specifies if the the Protection Group is using a protect once type of policy. This field is helpful to identify run happen for this group. |
| missing_entities | Specifies the Information about missing entities. |
| invalid_entities | Specifies the Information about invalid entities. An entity will be considered invalid if it is part of an active protection group but has lost compatibility for the given backup type. |
| num_protected_objects | Specifies the number of protected objects of the Protection Group. |
| advanced_configs | Specifies the advanced configuration for a protection job. |
| physical_params |  |
| oracle_params | Specifies the parameters to create Oracle Protection Group. |

### Data source: ibm_protection_groups

```hcl
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
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| request_initiator_type | Specifies the type of request from UI, which is used for services like magneto to determine the priority of requests. | `string` | false |
| ids | Filter by a list of Protection Group ids. | `list(string)` | false |
| names | Filter by a list of Protection Group names. | `list(string)` | false |
| policy_ids | Filter by Policy ids that are associated with Protection Groups. Only Protection Groups associated with the specified Policy ids, are returned. | `list(string)` | false |
| storage_domain_id | Filter by Storage Domain id. Only Protection Groups writing data to this Storage Domain will be returned. | `number` | false |
| include_groups_with_datalock_only | Whether to only return Protection Groups with a datalock. | `bool` | false |
| environments | Filter by environment types such as 'kVMware', 'kView', etc. Only Protection Groups protecting the specified environment types are returned. | `list(string)` | false |
| is_active | Filter by Inactive or Active Protection Groups. If not set, all Inactive and Active Protection Groups are returned. If true, only Active Protection Groups are returned. If false, only Inactive Protection Groups are returned. When you create a Protection Group on a Primary Cluster with a replication schedule, the Cluster creates an Inactive copy of the Protection Group on the Remote Cluster. In addition, when an Active and running Protection Group is deactivated, the Protection Group becomes Inactive. | `bool` | false |
| is_deleted | If true, return only Protection Groups that have been deleted but still have Snapshots associated with them. If false, return all Protection Groups except those Protection Groups that have been deleted and still have Snapshots associated with them. A Protection Group that is deleted with all its Snapshots is not returned for either of these cases. | `bool` | false |
| is_paused | Filter by paused or non paused Protection Groups, If not set, all paused and non paused Protection Groups are returned. If true, only paused Protection Groups are returned. If false, only non paused Protection Groups are returned. | `bool` | false |
| last_run_local_backup_status | Filter by last local backup run status.<br> 'Running' indicates that the run is still running.<br> 'Canceled' indicates that the run has been canceled.<br> 'Canceling' indicates that the run is in the process of being canceled.<br> 'Failed' indicates that the run has failed.<br> 'Missed' indicates that the run was unable to take place at the scheduled time because the previous run was still happening.<br> 'Succeeded' indicates that the run has finished successfully.<br> 'SucceededWithWarning' indicates that the run finished successfully, but there were some warning messages.<br> 'Paused' indicates that the ongoing run has been paused.<br> 'Skipped' indicates that the run was skipped. | `list(string)` | false |
| last_run_replication_status | Filter by last remote replication run status.<br> 'Running' indicates that the run is still running.<br> 'Canceled' indicates that the run has been canceled.<br> 'Canceling' indicates that the run is in the process of being canceled.<br> 'Failed' indicates that the run has failed.<br> 'Missed' indicates that the run was unable to take place at the scheduled time because the previous run was still happening.<br> 'Succeeded' indicates that the run has finished successfully.<br> 'SucceededWithWarning' indicates that the run finished successfully, but there were some warning messages.<br> 'Paused' indicates that the ongoing run has been paused.<br> 'Skipped' indicates that the run was skipped. | `list(string)` | false |
| last_run_archival_status | Filter by last cloud archival run status.<br> 'Running' indicates that the run is still running.<br> 'Canceled' indicates that the run has been canceled.<br> 'Canceling' indicates that the run is in the process of being canceled.<br> 'Failed' indicates that the run has failed.<br> 'Missed' indicates that the run was unable to take place at the scheduled time because the previous run was still happening.<br> 'Succeeded' indicates that the run has finished successfully.<br> 'SucceededWithWarning' indicates that the run finished successfully, but there were some warning messages.<br> 'Paused' indicates that the ongoing run has been paused.<br> 'Skipped' indicates that the run was skipped. | `list(string)` | false |
| last_run_cloud_spin_status | Filter by last cloud spin run status.<br> 'Running' indicates that the run is still running.<br> 'Canceled' indicates that the run has been canceled.<br> 'Canceling' indicates that the run is in the process of being canceled.<br> 'Failed' indicates that the run has failed.<br> 'Missed' indicates that the run was unable to take place at the scheduled time because the previous run was still happening.<br> 'Succeeded' indicates that the run has finished successfully.<br> 'SucceededWithWarning' indicates that the run finished successfully, but there were some warning messages.<br> 'Paused' indicates that the ongoing run has been paused.<br> 'Skipped' indicates that the run was skipped. | `list(string)` | false |
| last_run_any_status | Filter by last any run status.<br> 'Running' indicates that the run is still running.<br> 'Canceled' indicates that the run has been canceled.<br> 'Canceling' indicates that the run is in the process of being canceled.<br> 'Failed' indicates that the run has failed.<br> 'Missed' indicates that the run was unable to take place at the scheduled time because the previous run was still happening.<br> 'Succeeded' indicates that the run has finished successfully.<br> 'SucceededWithWarning' indicates that the run finished successfully, but there were some warning messages.<br> 'Paused' indicates that the ongoing run has been paused.<br> 'Skipped' indicates that the run was skipped. | `list(string)` | false |
| is_last_run_sla_violated | If true, return Protection Groups for which last run SLA was violated. | `bool` | false |
| tenant_ids | TenantIds contains ids of the tenants for which objects are to be returned. | `list(string)` | false |
| include_tenants | If true, the response will include Protection Groups which were created by all tenants which the current user has permission to see. If false, then only Protection Groups created by the current user will be returned. | `bool` | false |
| include_last_run_info | If true, the response will include last run info. If it is false or not specified, the last run info won't be returned. | `bool` | false |
| prune_excluded_source_ids | If true, the response will not include the list of excluded source IDs in groups that contain this field. This can be set to true in order to improve performance if excluded source IDs are not needed by the user. | `bool` | false |
| prune_source_ids | If true, the response will exclude the list of source IDs within the group specified. | `bool` | false |
| use_cached_data | Specifies whether we can serve the GET request from the read replica cache. There is a lag of 15 seconds between the read replica and primary data source. | `bool` | false |

#### Outputs

| Name | Description |
|------|-------------|
| protection_groups | Specifies the list of Protection Groups which were returned by the request. |

### Data source: ibm_protection_group_run

```hcl
data "ibm_protection_group_run" "protection_group_run_instance" {
  protection_group_run_id = var.protection_group_run_protection_group_run_id
  run_id = var.protection_group_run_run_id
  request_initiator_type = var.protection_group_run_request_initiator_type
  tenant_ids = var.protection_group_run_tenant_ids
  include_tenants = var.protection_group_run_include_tenants
  include_object_details = var.protection_group_run_include_object_details
  use_cached_data = var.protection_group_run_use_cached_data
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| protection_group_run_id | Specifies a unique id of the Protection Group. | `string` | true |
| run_id | Specifies a unique run id of the Protection Group run. | `string` | true |
| request_initiator_type | Specifies the type of request from UI, which is used for services like magneto to determine the priority of requests. | `string` | false |
| tenant_ids | TenantIds contains ids of the tenants for which the run is to be returned. | `list(string)` | false |
| include_tenants | If true, the response will include Protection Group Runs which were created by all tenants which the current user has permission to see. If false, then only Protection Groups created by the current user will be returned. If it's not specified, it is true by default. | `bool` | false |
| include_object_details | Specifies if the result includes the object details for a protection run. If set to true, details of the protected object will be returned. If set to false or not specified, details will not be returned. | `bool` | false |
| use_cached_data | Specifies whether we can serve the GET request from the read replica cache. There is a lag of 15 seconds between the read replica and primary data source. | `bool` | false |

#### Outputs

| Name | Description |
|------|-------------|
| protection_group_instance_id | Protection Group instance Id. This field will be removed later. |
| protection_group_id | ProtectionGroupId to which this run belongs. |
| is_replication_run | Specifies if this protection run is a replication run. |
| origin_cluster_identifier | Specifies the information about a cluster. |
| origin_protection_group_id | ProtectionGroupId to which this run belongs on the primary cluster if this run is a replication run. |
| protection_group_name | Name of the Protection Group to which this run belongs. |
| is_local_snapshots_deleted | Specifies if snapshots for this run has been deleted. |
| objects | Snapahot, replication, archival results for each object. |
| local_backup_info | Specifies summary information about local snapshot run across all objects. |
| original_backup_info | Specifies summary information about local snapshot run across all objects. |
| replication_info | Specifies summary information about replication run. |
| archival_info | Specifies summary information about archival run. |
| cloud_spin_info | Specifies summary information about cloud spin run. |
| on_legal_hold | Specifies if the Protection Run is on legal hold. |
| permissions | Specifies the list of tenants that have permissions for this protection group run. |
| is_cloud_archival_direct | Specifies whether the run is a CAD run if cloud archive direct feature is enabled. If this field is true, the primary backup copy will only be available at the given archived location. |
| has_local_snapshot | Specifies whether the run has a local snapshot. For cloud retrieved runs there may not be local snapshots. |
| environment | Specifies the environment of the Protection Group. |
| externally_triggered_backup_tag | The tag of externally triggered backup job. |

### Data source: ibm_protection_group_runs

```hcl
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
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| protection_group_runs_id | Specifies a unique id of the Protection Group. | `string` | true |
| request_initiator_type | Specifies the type of request from UI, which is used for services like magneto to determine the priority of requests. | `string` | false |
| run_id | Specifies the protection run id. | `string` | false |
| start_time_usecs | Start time for time range filter. Specify the start time as a Unix epoch Timestamp (in microseconds), only runs executing after this time will be returned. By default it is endTimeUsecs minus an hour. | `number` | false |
| end_time_usecs | End time for time range filter. Specify the end time as a Unix epoch Timestamp (in microseconds), only runs executing before this time will be returned. By default it is current time. | `number` | false |
| tenant_ids | TenantIds contains ids of the tenants for which objects are to be returned. | `list(string)` | false |
| include_tenants | If true, the response will include Protection Group Runs which were created by all tenants which the current user has permission to see. If false, then only Protection Group Runs created by the current user will be returned. | `bool` | false |
| run_types | Filter by run type. Only protection run matching the specified types will be returned. | `list(string)` | false |
| include_object_details | Specifies if the result includes the object details for each protection run. If set to true, details of the protected object will be returned. If set to false or not specified, details will not be returned. | `bool` | false |
| local_backup_run_status | Specifies a list of local backup status, runs matching the status will be returned.<br> 'Running' indicates that the run is still running.<br> 'Canceled' indicates that the run has been canceled.<br> 'Canceling' indicates that the run is in the process of being canceled.<br> 'Failed' indicates that the run has failed.<br> 'Missed' indicates that the run was unable to take place at the scheduled time because the previous run was still happening.<br> 'Succeeded' indicates that the run has finished successfully.<br> 'SucceededWithWarning' indicates that the run finished successfully, but there were some warning messages.<br> 'Paused' indicates that the ongoing run has been paused.<br> 'Skipped' indicates that the run was skipped. | `list(string)` | false |
| replication_run_status | Specifies a list of replication status, runs matching the status will be returned.<br> 'Running' indicates that the run is still running.<br> 'Canceled' indicates that the run has been canceled.<br> 'Canceling' indicates that the run is in the process of being canceled.<br> 'Failed' indicates that the run has failed.<br> 'Missed' indicates that the run was unable to take place at the scheduled time because the previous run was still happening.<br> 'Succeeded' indicates that the run has finished successfully.<br> 'SucceededWithWarning' indicates that the run finished successfully, but there were some warning messages.<br> 'Paused' indicates that the ongoing run has been paused.<br> 'Skipped' indicates that the run was skipped. | `list(string)` | false |
| archival_run_status | Specifies a list of archival status, runs matching the status will be returned.<br> 'Running' indicates that the run is still running.<br> 'Canceled' indicates that the run has been canceled.<br> 'Canceling' indicates that the run is in the process of being canceled.<br> 'Failed' indicates that the run has failed.<br> 'Missed' indicates that the run was unable to take place at the scheduled time because the previous run was still happening.<br> 'Succeeded' indicates that the run has finished successfully.<br> 'SucceededWithWarning' indicates that the run finished successfully, but there were some warning messages.<br> 'Paused' indicates that the ongoing run has been paused.<br> 'Skipped' indicates that the run was skipped. | `list(string)` | false |
| cloud_spin_run_status | Specifies a list of cloud spin status, runs matching the status will be returned.<br> 'Running' indicates that the run is still running.<br> 'Canceled' indicates that the run has been canceled.<br> 'Canceling' indicates that the run is in the process of being canceled.<br> 'Failed' indicates that the run has failed.<br> 'Missed' indicates that the run was unable to take place at the scheduled time because the previous run was still happening.<br> 'Succeeded' indicates that the run has finished successfully.<br> 'SucceededWithWarning' indicates that the run finished successfully, but there were some warning messages.<br> 'Paused' indicates that the ongoing run has been paused.<br> 'Skipped' indicates that the run was skipped. | `list(string)` | false |
| num_runs | Specifies the max number of runs. If not specified, at most 100 runs will be returned. | `number` | false |
| exclude_non_restorable_runs | Specifies whether to exclude non restorable runs. Run is treated restorable only if there is atleast one object snapshot (which may be either a local or an archival snapshot) which is not deleted or expired. Default value is false. | `bool` | false |
| run_tags | Specifies a list of tags for protection runs. If this is specified, only the runs which match these tags will be returned. | `list(string)` | false |
| use_cached_data | Specifies whether we can serve the GET request from the read replica cache. There is a lag of 15 seconds between the read replica and primary data source. | `bool` | false |
| filter_by_end_time | If true, the runs with backup end time within the specified time range will be returned. Otherwise, the runs with start time in the time range are returned. | `bool` | false |
| snapshot_target_types | Specifies the snapshot's target type which should be filtered. | `list(string)` | false |
| only_return_successful_copy_run | only successful copyruns are returned. | `bool` | false |
| filter_by_copy_task_end_time | If true, then the details of the runs for which any copyTask completed in the given timerange will be returned. Only one of filterByEndTime and filterByCopyTaskEndTime can be set. | `bool` | false |

#### Outputs

| Name | Description |
|------|-------------|
| runs | Specifies the list of Protection Group runs. |
| total_runs | Specifies the count of total runs exist for the given set of filters. The number of runs in single API call are limited and this count can be used to estimate query filter values to get next set of remaining runs. Please note that this field will only be populated if startTimeUsecs or endTimeUsecs or both are specified in query parameters. |

### Data source: ibm_protection_policies

```hcl
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
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| request_initiator_type | Specifies the type of request from UI, which is used for services like magneto to determine the priority of requests. | `string` | false |
| ids | Filter policies by a list of policy ids. | `list(string)` | false |
| policy_names | Filter policies by a list of policy names. | `list(string)` | false |
| tenant_ids | TenantIds contains ids of the organizations for which objects are to be returned. | `list(string)` | false |
| include_tenants | IncludeTenantPolicies specifies if objects of all the organizations under the hierarchy of the logged in user's organization should be returned. | `bool` | false |
| types | Types specifies the policy type of policies to be returned. | `list(string)` | false |
| exclude_linked_policies | If excludeLinkedPolicies is set to true then only local policies created on cluster will be returned. The result will exclude all linked policies created from policy templates. | `bool` | false |
| include_replicated_policies | If includeReplicatedPolicies is set to true, then response will also contain replicated policies. By default, replication policies are not included in the response. | `bool` | false |
| include_stats | If includeStats is set to true, then response will return number of protection groups and objects. By default, the protection stats are not included in the response. | `bool` | false |

#### Outputs

| Name | Description |
|------|-------------|
| policies | Specifies a list of protection policies. |

### Data source: ibm_protection_policy

```hcl
data "ibm_protection_policy" "protection_policy_instance" {
  protection_policy_id = var.data_protection_policy_protection_policy_id
  request_initiator_type = var.data_protection_policy_request_initiator_type
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| protection_policy_id | Specifies a unique id of the Protection Policy to return. | `string` | true |
| request_initiator_type | Specifies the type of request from UI, which is used for services like magneto to determine the priority of requests. | `string` | false |

#### Outputs

| Name | Description |
|------|-------------|
| name | Specifies the name of the Protection Policy. |
| backup_policy | Specifies the backup schedule and retentions of a Protection Policy. |
| description | Specifies the description of the Protection Policy. |
| blackout_window | List of Blackout Windows. If specified, this field defines blackout periods when new Group Runs are not started. If a Group Run has been scheduled but not yet executed and the blackout period starts, the behavior depends on the policy field AbortInBlackoutPeriod. |
| extended_retention | Specifies additional retention policies that should be applied to the backup snapshots. A backup snapshot will be retained up to a time that is the maximum of all retention policies that are applicable to it. |
| remote_target_policy | Specifies the replication, archival and cloud spin targets of Protection Policy. |
| cascaded_targets_config | Specifies the configuration for cascaded replications. Using cascaded replication, replication cluster(Rx) can further replicate and archive the snapshot copies to further targets. Its recommended to create cascaded configuration where protection group will be created. |
| retry_options | Retry Options of a Protection Policy when a Protection Group run fails. |
| data_lock | This field is now deprecated. Please use the DataLockConfig in the backup retention. |
| version | Specifies the current policy verison. Policy version is incremented for optionally supporting new features and differentialting across releases. |
| is_cbs_enabled | Specifies true if Calender Based Schedule is supported by client. Default value is assumed as false for this feature. |
| last_modification_time_usecs | Specifies the last time this Policy was updated. If this is passed into a PUT request, then the backend will validate that the timestamp passed in matches the time that the policy was actually last modified. If the two timestamps do not match, then the request will be rejected with a stale error. |
| template_id | Specifies the parent policy template id to which the policy is linked to. This field is set only when policy is created from template. |
| is_usable | This field is set to true if the linked policy which is internally created from a policy templates qualifies as usable to create more policies on the cluster. If the linked policy is partially filled and can not create a working policy then this field will be set to false. In case of normal policy created on the cluster, this field wont be populated. |
| is_replicated | This field is set to true when policy is the replicated policy. |
| num_protection_groups | Specifies the number of protection groups using the protection policy. |
| num_protected_objects | Specifies the number of protected objects using the protection policy. |

### Data source: ibm_protection_run_summary

```hcl
data "ibm_protection_run_summary" "protection_run_summary_instance" {
  start_time_usecs = var.protection_run_summary_start_time_usecs
  end_time_usecs = var.protection_run_summary_end_time_usecs
  run_status = var.protection_run_summary_run_status
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| start_time_usecs | Start time for time range filter. Specify the start time as a Unix epoch Timestamp (in microseconds), only runs executing after this time will be returned. By default it is endTimeUsecs minus an hour. | `number` | false |
| end_time_usecs | End time for time range filter. Specify the end time as a Unix epoch Timestamp (in microseconds), only runs executing before this time will be returned. By default it is current time. | `number` | false |
| run_status | Specifies a list of status, runs matching the status will be returned.<br> 'Running' indicates that the run is still running.<br> 'Canceled' indicates that the run has been canceled.<br> 'Canceling' indicates that the run is in the process of being canceled.<br> 'Failed' indicates that the run has failed.<br> 'Missed' indicates that the run was unable to take place at the scheduled time because the previous run was still happening.<br> 'Succeeded' indicates that the run has finished successfully.<br> 'SucceededWithWarning' indicates that the run finished successfully, but there were some warning messages.<br> 'Skipped' indicates that the run was skipped. | `list(string)` | false |

#### Outputs

| Name | Description |
|------|-------------|
| protection_runs_summary | Specifies a list of summaries of protection runs. |

### Data source: ibm_protection_sources

```hcl
data "ibm_protection_sources" "protection_sources_instance" {
  request_initiator_type = var.protection_sources_request_initiator_type
  tenant_ids = var.protection_sources_tenant_ids
  include_tenants = var.protection_sources_include_tenants
  include_source_credentials = var.protection_sources_include_source_credentials
  encryption_key = var.protection_sources_encryption_key
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| request_initiator_type | Specifies the type of request from UI, which is used for services like magneto to determine the priority of requests. | `string` | false |
| tenant_ids | TenantIds contains ids of the tenants for which Sources are to be returned. | `list(string)` | false |
| include_tenants | If true, the response will include Sources which belong belong to all tenants which the current user has permission to see. If false, then only Sources for the current user will be returned. | `bool` | false |
| include_source_credentials | If true, the encrypted crednetial for the registered sources will be included. Credential is first encrypted with internal key and then reencrypted with user supplied encryption key. | `bool` | false |
| encryption_key | Specifies the key to be used to encrypt the source credential. If includeSourceCredentials is set to true this key must be specified. | `string` | false |

#### Outputs

| Name | Description |
|------|-------------|
| sources | Specifies the list of Protection Sources. |

### Data source: ibm_recovery

```hcl
data "ibm_recovery" "recovery_instance" {
  recovery_id = var.data_recovery_recovery_id
  include_tenants = var.data_recovery_include_tenants
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| recovery_id | Specifies the id of a Recovery. | `string` | true |
| include_tenants | Specifies if objects of all the organizations under the hierarchy of the logged in user's organization should be returned. | `bool` | false |

#### Outputs

| Name | Description |
|------|-------------|
| name | Specifies the name of the Recovery. |
| start_time_usecs | Specifies the start time of the Recovery in Unix timestamp epoch in microseconds. |
| end_time_usecs | Specifies the end time of the Recovery in Unix timestamp epoch in microseconds. This field will be populated only after Recovery is finished. |
| status | Status of the Recovery. 'Running' indicates that the Recovery is still running. 'Canceled' indicates that the Recovery has been cancelled. 'Canceling' indicates that the Recovery is in the process of being cancelled. 'Failed' indicates that the Recovery has failed. 'Succeeded' indicates that the Recovery has finished successfully. 'SucceededWithWarning' indicates that the Recovery finished successfully, but there were some warning messages. 'Skipped' indicates that the Recovery task was skipped. |
| progress_task_id | Progress monitor task id for Recovery. |
| snapshot_environment | Specifies the type of snapshot environment for which the Recovery was performed. |
| recovery_action | Specifies the type of recover action. |
| permissions | Specifies the list of tenants that have permissions for this recovery. |
| creation_info | Specifies the information about the creation of the protection group or recovery. |
| can_tear_down | Specifies whether it's possible to tear down the objects created by the recovery. |
| tear_down_status | Specifies the status of the tear down operation. This is only set when the canTearDown is set to true. 'DestroyScheduled' indicates that the tear down is ready to schedule. 'Destroying' indicates that the tear down is still running. 'Destroyed' indicates that the tear down succeeded. 'DestroyError' indicates that the tear down failed. |
| tear_down_message | Specifies the error message about the tear down operation if it fails. |
| messages | Specifies messages about the recovery. |
| is_parent_recovery | Specifies whether the current recovery operation has created child recoveries. This is currently used in SQL recovery where multiple child recoveries can be tracked under a common/parent recovery. |
| parent_recovery_id | If current recovery is child recovery triggered by another parent recovery operation, then this field willt specify the id of the parent recovery. |
| retrieve_archive_tasks | Specifies the list of persistent state of a retrieve of an archive task. |
| is_multi_stage_restore | Specifies whether the current recovery operation is a multi-stage restore operation. This is currently used by VMware recoveres for the migration/hot-standby use case. |
| physical_params | Specifies the recovery options specific to Physical environment. |
| oracle_params | Specifies the recovery options specific to oracle environment. |

### Data source: ibm_recoveries

```hcl
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
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| ids | Filter Recoveries for given ids. | `list(string)` | false |
| return_only_child_recoveries | Returns only child recoveries if passed as true. This filter should always be used along with 'ids' filter. | `bool` | false |
| tenant_ids | TenantIds contains ids of the organizations for which recoveries are to be returned. | `list(string)` | false |
| include_tenants | Specifies if objects of all the organizations under the hierarchy of the logged in user's organization should be returned. | `bool` | false |
| start_time_usecs | Returns the recoveries which are started after the specific time. This value should be in Unix timestamp epoch in microseconds. | `number` | false |
| end_time_usecs | Returns the recoveries which are started before the specific time. This value should be in Unix timestamp epoch in microseconds. | `number` | false |
| storage_domain_id | Filter by Storage Domain id. Only recoveries writing data to this Storage Domain will be returned. | `number` | false |
| snapshot_target_type | Specifies the snapshot's target type from which recovery has been performed. | `list(string)` | false |
| archival_target_type | Specifies the snapshot's archival target type from which recovery has been performed. This parameter applies only if 'snapshotTargetType' is 'Archival'. | `list(string)` | false |
| snapshot_environments | Specifies the list of snapshot environment types to filter Recoveries. If empty, Recoveries related to all environments will be returned. | `list(string)` | false |
| status | Specifies the list of run status to filter Recoveries. If empty, Recoveries with all run status will be returned. | `list(string)` | false |
| recovery_actions | Specifies the list of recovery actions to filter Recoveries. If empty, Recoveries related to all actions will be returned. | `list(string)` | false |

#### Outputs

| Name | Description |
|------|-------------|
| recoveries | Specifies list of Recoveries. |

### Data source: ibm_source_registrations

```hcl
data "ibm_source_registrations" "source_registrations_instance" {
  ids = var.source_registrations_ids
  tenant_ids = var.source_registrations_tenant_ids
  include_tenants = var.source_registrations_include_tenants
  include_source_credentials = var.source_registrations_include_source_credentials
  encryption_key = var.source_registrations_encryption_key
  use_cached_data = var.source_registrations_use_cached_data
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| ids | Ids specifies the list of source registration ids to return. If left empty, every source registration will be returned by default. | `list(number)` | false |
| tenant_ids | TenantIds contains ids of the tenants for which objects are to be returned. | `list(string)` | false |
| include_tenants | If true, the response will include Registrations which were created by all tenants which the current user has permission to see. If false, then only Registrations created by the current user will be returned. | `bool` | false |
| include_source_credentials | If true, the encrypted crednetial for the registered sources will be included. Credential is first encrypted with internal key and then reencrypted with user supplied encryption key. | `bool` | false |
| encryption_key | Specifies the key to be used to encrypt the source credential. If includeSourceCredentials is set to true this key must be specified. | `string` | false |
| use_cached_data | Specifies whether we can serve the GET request from the read replica cache. There is a lag of 15 seconds between the read replica and primary data source. | `bool` | false |

#### Outputs

| Name | Description |
|------|-------------|
| registrations | Specifies the list of Protection Source Registrations. |

### Data source: ibm_source_registration

```hcl
data "ibm_source_registration" "source_registration_instance" {
  source_registration_id = var.data_source_registration_source_registration_id
  request_initiator_type = var.data_source_registration_request_initiator_type
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| source_registration_id | Specifies the id of the Protection Source registration. | `number` | true |
| request_initiator_type | Specifies the type of request from UI, which is used for services like magneto to determine the priority of requests. | `string` | false |

#### Outputs

| Name | Description |
|------|-------------|
| source_id | ID of top level source object discovered after the registration. |
| source_info | Specifies information about an object. |
| environment | Specifies the environment type of the Protection Source. |
| name | The user specified name for this source. |
| connection_id | Specifies the id of the connection from where this source is reachable. This should only be set for a source being registered by a tenant user. This field will be depricated in future. Use connections field. |
| connections | Specfies the list of connections for the source. |
| connector_group_id | Specifies the connector group id of connector groups. |
| advanced_configs | Specifies the advanced configuration for a protection source. |
| authentication_status | Specifies the status of the authentication during the registration of a Protection Source. 'Pending' indicates the authentication is in progress. 'Scheduled' indicates the authentication is scheduled. 'Finished' indicates the authentication is completed. 'RefreshInProgress' indicates the refresh is in progress. |
| registration_time_msecs | Specifies the time when the source was registered in milliseconds. |
| last_refreshed_time_msecs | Specifies the time when the source was last refreshed in milliseconds. |
| physical_params | Physical Params params. |
| oracle_params | Physical Params params. |

## Assumptions

1. TODO

## Notes

### Resources with Incomplete CRUD Operations
This service includes certain resources that do not have fully implemented CRUD (Create, Read, Update, Delete) operations due to limitations in the underlying APIs. Specifically:

#### Protection Group Run:

***Create:*** A `ibm_protection_group_run_request` resource is available for creating new protection group run.

***Update:*** protection group run updates are managed through a separate `ibm_update_protection_group_run_request` resource. Note that the `ibm_protection_group_run_request` and `ibm_update_protection_group_run_request` resources must be used in tandem to manage Protection Group Runs.

***Delete:*** There is no delete operation available for the protection group run resource. If  ibm_update_protection_group_run_request or ibm_protection_group_run_request resource is removed from the `main.tf` file, Terraform will remove it from the state file but not from the backend. The resource will continue to exist in the backend system.


#### Other resources that do not support update and delete:

Some resources in this service do not support update or delete operations due to the absence of corresponding API endpoints. As a result, Terraform cannot manage these operations for those resources. Users should be aware that removing these resources from the configuration (main.tf) will only remove them from the Terraform state and will not affect the actual resources in the backend.
- ibm_perform_action_on_protection_group_run_request
- ibm_recovery_download_files_folders
- ibm_recovery_cancel
- ibm_recovery_teardown
- ibm_search_indexed_object
- ibm_protection_group_state

**Important:** When managing resources that lack complete CRUD operations, users should exercise caution and consider the limitations described above. Manual intervention may be required to manage these resources in the backend if updates or deletions are necessary.**


## Requirements

| Name | Version |
|------|---------|
| terraform | ~> 0.12 |

## Providers

| Name | Version |
|------|---------|
| ibm | 1.13.1 |
