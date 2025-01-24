# Examples for IBM Backup Recovery API

These examples illustrate how to use the resources and data sources associated with IBM Backup Recovery API.

The following resources are supported:
* ibm_backup_recovery_connection_registration_token
* ibm_backup_recovery_agent_upgrade_task
* ibm_backup_recovery_protection_group_run_request
* ibm_backup_recovery_data_source_connection
* ibm_backup_recovery_download_files_folders
* ibm_backup_recovery_restore_points
* ibm_backup_recovery_perform_action_on_protection_group_run_request
* ibm_backup_recovery_protection_group
* ibm_backup_recovery_protection_policy
* ibm_backup_recovery
* ibm_backup_recovery_source_registration
* ibm_backup_recovery_update_protection_group_run_request

The following data sources are supported:
* ibm_backup_recovery_agent_upgrade_tasks
* ibm_backup_recovery_connectors_metadata
* ibm_backup_recovery_data_source_connections
* ibm_backup_recovery_data_source_connectors
* ibm_backup_recovery_download_agent
* ibm_backup_recovery_download_indexed_files
* ibm_backup_recovery_object_snapshots
* ibm_backup_recovery_protection_group_runs
* ibm_backup_recovery_protection_group
* ibm_backup_recovery_protection_groups
* ibm_backup_recovery_protection_policies
* ibm_backup_recovery_protection_policy
* ibm_backup_recoveries
* ibm_backup_recovery_download_files
* ibm_backup_recovery
* ibm_backup_recovery_search_indexed_object
* ibm_backup_recovery_search_objects
* ibm_backup_recovery_search_protected_objects
* ibm_backup_recovery_source_registrations
* ibm_backup_recovery_source_registration
* ibm_backup_recovery_protection_sources


## Usage

To run this example, execute the following commands:

```bash
$ terraform init
$ terraform plan
$ terraform apply
```

Run `terraform destroy` when you don't need these resources.

## IBM Backup recovery API resources

### Resource: ibm_backup_recovery_connection_registration_token

```hcl
resource "ibm_backup_recovery_connection_registration_token" "backup_recovery_connection_registration_token_instance" {
  connection_id = var.backup_recovery_connection_registration_token_connection_id
  x_ibm_tenant_id = var.backup_recovery_connection_registration_token_x_ibm_tenant_id
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| ibmcloud\_api\_key | IBM Cloud API key | `string` | true |
| connection_id | Specifies the ID of the connection, connectors belonging to which are to be fetched. | `string` | true |
| x_ibm_tenant_id | Specifies the key to be used to encrypt the source credential. If includeSourceCredentials is set to true this key must be specified. | `string` | true |

#### Outputs

| Name | Description |
|------|-------------|
| registration_token | Specifies a token that can be used to register a connector against this connection |

### Resource: ibm_backup_recovery_agent_upgrade_task

```hcl
resource "ibm_backup_recovery_agent_upgrade_task" "backup_recovery_agent_upgrade_task_instance" {
  x_ibm_tenant_id = var.backup_recovery_agent_upgrade_tasks_x_ibm_tenant_id
  agent_ids = var.backup_recovery_agent_upgrade_tasks_ids
  description = ""
  name = ""
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| x_ibm_tenant_id | Specifies the key to be used to encrypt the source credential. If includeSourceCredentials is set to true this key must be specified. | `string` | true |
| agent_ids | Specifies IDs of tasks to be fetched. | `list(number)` | false |
| retry_task_id | Specifies the retry task id | `number` | false |
| description | task description| `string` | false |
| name | name of the task | `string` | false |
| schedule_end_time_usecs | Specifies the end time specified as a Unix epoch Timestamp in microseconds. | `number` | false |
| schedule_time_usecs | Specifies the schedule time specified as a Unix epoch Timestamp in microseconds. | `number` | false |

### Resource: ibm_backup_recovery_protection_group_run_request

```hcl
resource "ibm_backup_recovery_protection_group_run_request" "backup_recovery_protection_group_run_request_instance" {
  x_ibm_tenant_id = var.backup_recovery_protection_group_run_request_x_ibm_tenant_id
  run_type = var.backup_recovery_protection_group_run_request_run_type
  objects = var.backup_recovery_protection_group_run_request_objects
  targets_config = var.backup_recovery_protection_group_run_request_targets_config
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| ibmcloud\_api\_key | IBM Cloud API key | `string` | true |
| x_ibm_tenant_id | Specifies the key to be used to encrypt the source credential. If includeSourceCredentials is set to true this key must be specified. | `string` | true |
| run_type | Type of protection run. 'kRegular' indicates an incremental (CBT) backup. Incremental backups utilizing CBT (if supported) are captured of the target protection objects. The first run of a kRegular schedule captures all the blocks. 'kFull' indicates a full (no CBT) backup. A complete backup (all blocks) of the target protection objects are always captured and Change Block Tracking (CBT) is not utilized. 'kLog' indicates a Database Log backup. Capture the database transaction logs to allow rolling back to a specific point in time. 'kSystem' indicates system volume backup. It produces an image for bare metal recovery. | `string` | true |
| objects | Specifies the list of objects to be protected by this Protection Group run. These can be leaf objects or non-leaf objects in the protection hierarchy. This must be specified only if a subset of objects from the Protection Groups needs to be protected. | `list()` | false |
| targets_config | Specifies the replication and archival targets. | `object` | false |

### Resource: ibm_backup_recovery_data_source_connection

```hcl
resource "ibm_backup_recovery_data_source_connection" "backup_recovery_data_source_connection_instance" {
  x_ibm_tenant_id = var.backup_recovery_data_source_connection_x_ibm_tenant_id
  connection_name = var.backup_recovery_data_source_connection_connection_name
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| ibmcloud\_api\_key | IBM Cloud API key | `string` | true |
| x_ibm_tenant_id | Id of the tenant accessing the cluster. | `string` | false |
| connection_name | Specifies the name of the connection. For a given tenant, different connections can't have the same name. However, two (or more) different tenants can each have a connection with the same name. | `string` | true |

#### Outputs

| Name | Description |
|------|-------------|
| connector_ids | Specifies the IDs of the connectors in this connection. |
| network_settings | Specifies the common network settings for the connectors associated with this connection. |
| registration_token | Specifies a token that can be used to register a connector against this connection. |
| tenant_id | Specifies the tenant ID of the connection. |

### Resource: ibm_backup_recovery_download_files_folders

```hcl
resource "ibm_backup_recovery_download_files_folders" "backup_recovery_download_files_folders_instance" {
  x_ibm_tenant_id = var.backup_recovery_download_files_folders_x_ibm_tenant_id
  documents = var.backup_recovery_download_files_folders_documents
  name = var.backup_recovery_download_files_folders_name
  object = var.backup_recovery_download_files_folders_object
  parent_recovery_id = var.backup_recovery_download_files_folders_parent_recovery_id
  files_and_folders = var.backup_recovery_download_files_folders_files_and_folders
  glacier_retrieval_type = var.backup_recovery_download_files_folders_glacier_retrieval_type
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| ibmcloud\_api\_key | IBM Cloud API key | `string` | true |
| x_ibm_tenant_id | Specifies the key to be used to encrypt the source credential. If includeSourceCredentials is set to true this key must be specified. | `string` | true |
| documents | Specifies the list of documents to download using item ids. Only one of filesAndFolders or documents should be used. Currently only files are supported by documents. | `list()` | false |
| name | Specifies the name of the recovery task. This field must be set and must be a unique name. | `string` | true |
| object | Specifies the common snapshot parameters for a protected object. | `object` | true |
| parent_recovery_id | If current recovery is child task triggered through another parent recovery operation, then this field will specify the id of the parent recovery. | `string` | false |
| files_and_folders | Specifies the list of files and folders to download. | `list()` | true |
| glacier_retrieval_type | Specifies the glacier retrieval type when restoring or downloding files or folders from a Glacier-based cloud snapshot. | `string` | false |

### Resource: ibm_backup_recovery_restore_points

```hcl
resource "ibm_backup_recovery_restore_points" "backup_recovery_restore_points_instance" {
  x_ibm_tenant_id = var.backup_recovery_restore_points_x_ibm_tenant_id
  end_time_usecs = var.backup_recovery_restore_points_end_time_usecs
  environment = var.backup_recovery_restore_points_environment
  protection_group_ids = var.backup_recovery_restore_points_protection_group_ids
  source_id = var.backup_recovery_restore_points_source_id
  start_time_usecs = var.backup_recovery_restore_points_start_time_usecs
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| ibmcloud\_api\_key | IBM Cloud API key | `string` | true |
| x_ibm_tenant_id | Specifies the key to be used to encrypt the source credential. If includeSourceCredentials is set to true this key must be specified. | `string` | true |
| end_time_usecs | Specifies the end time specified as a Unix epoch Timestamp in microseconds. | `number` | true |
| environment | Specifies the protection source environment type. | `string` | true |
| protection_group_ids | Specifies the jobs for which to get the full snapshot information. | `list(string)` | true |
| source_id | Specifies the id of the Protection Source which is to be restored. | `number` | false |
| start_time_usecs | Specifies the start time specified as a Unix epoch Timestamp in microseconds. | `number` | true |

### Resource: ibm_backup_recovery_perform_action_on_protection_group_run_request

```hcl
resource "ibm_backup_recovery_perform_action_on_protection_group_run_request" "backup_recovery_perform_action_on_protection_group_run_request_instance" {
  x_ibm_tenant_id = var.backup_recovery_perform_action_on_protection_group_run_request_x_ibm_tenant_id
  action = var.backup_recovery_perform_action_on_protection_group_run_request_action
  pause_params = var.backup_recovery_perform_action_on_protection_group_run_request_pause_params
  resume_params = var.backup_recovery_perform_action_on_protection_group_run_request_resume_params
  cancel_params = var.backup_recovery_perform_action_on_protection_group_run_request_cancel_params
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| ibmcloud\_api\_key | IBM Cloud API key | `string` | true |
| x_ibm_tenant_id | Specifies the key to be used to encrypt the source credential. If includeSourceCredentials is set to true this key must be specified. | `string` | true |
| action | Specifies the type of the action which will be performed on protection runs. | `string` | true |
| pause_params | Specifies the pause action params for a protection run. | `list()` | false |
| resume_params | Specifies the resume action params for a protection run. | `list()` | false |
| cancel_params | Specifies the cancel action params for a protection run. | `list()` | false |

### Resource: ibm_backup_recovery_protection_group

```hcl
resource "ibm_backup_recovery_protection_group" "backup_recovery_protection_group_instance" {
  x_ibm_tenant_id = var.backup_recovery_protection_group_x_ibm_tenant_id
  name = var.backup_recovery_protection_group_name
  policy_id = var.backup_recovery_protection_group_policy_id
  priority = var.backup_recovery_protection_group_priority
  description = var.backup_recovery_protection_group_description
  start_time = var.backup_recovery_protection_group_start_time
  end_time_usecs = var.backup_recovery_protection_group_end_time_usecs
  last_modified_timestamp_usecs = var.backup_recovery_protection_group_last_modified_timestamp_usecs
  alert_policy = var.backup_recovery_protection_group_alert_policy
  sla = var.backup_recovery_protection_group_sla
  qos_policy = var.backup_recovery_protection_group_qos_policy
  abort_in_blackouts = var.backup_recovery_protection_group_abort_in_blackouts
  pause_in_blackouts = var.backup_recovery_protection_group_pause_in_blackouts
  is_paused = var.backup_recovery_protection_group_is_paused
  environment = var.backup_recovery_protection_group_environment
  advanced_configs = var.backup_recovery_protection_group_advanced_configs
  physical_params = var.backup_recovery_protection_group_physical_params
  mssql_params = var.backup_recovery_protection_group_mssql_params
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| ibmcloud\_api\_key | IBM Cloud API key | `string` | true |
| x_ibm_tenant_id | Specifies the key to be used to encrypt the source credential. If includeSourceCredentials is set to true this key must be specified. | `string` | true |
| name | Specifies the name of the Protection Group. | `string` | true |
| policy_id | Specifies the unique id of the Protection Policy associated with the Protection Group. The Policy provides retry settings Protection Schedules, Priority, SLA, etc. | `string` | true |
| priority | Specifies the priority of the Protection Group. | `string` | false |
| description | Specifies a description of the Protection Group. | `string` | false |
| start_time | Specifies the time of day. Used for scheduling purposes. | `object` | false |
| end_time_usecs | Specifies the end time in micro seconds for this Protection Group. If this is not specified, the Protection Group won't be ended. | `number` | false |
| last_modified_timestamp_usecs | Specifies the last time this protection group was updated. If this is passed into a PUT request, then the backend will validate that the timestamp passed in matches the time that the protection group was actually last modified. If the two timestamps do not match, then the request will be rejected with a stale error. | `number` | false |
| alert_policy | Specifies a policy for alerting users of the status of a Protection Group. | `object` | false |
| sla | Specifies the SLA parameters for this Protection Group. | `list()` | false |
| qos_policy | Specifies whether the Protection Group will be written to HDD or SSD. | `string` | false |
| abort_in_blackouts | Specifies whether currently executing jobs should abort if a blackout period specified by a policy starts. Available only if the selected policy has at least one blackout period. Default value is false. | `bool` | false |
| pause_in_blackouts | Specifies whether currently executing jobs should be paused if a blackout period specified by a policy starts. Available only if the selected policy has at least one blackout period. Default value is false. This field should not be set to true if 'abortInBlackouts' is sent as true. | `bool` | false |
| is_paused | Specifies if the the Protection Group is paused. New runs are not scheduled for the paused Protection Groups. Active run if any is not impacted. | `bool` | false |
| environment | Specifies the environment of the Protection Group. | `string` | true |
| advanced_configs | Specifies the advanced configuration for a protection job. | `list()` | false |
| physical_params | Specifies the parameters for Physical object. | `object` | false |
| mssql_params | Specifies the parameters specific to MSSQL Protection Group. | `object` | false |

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

### Resource: ibm_backup_recovery_protection_policy

```hcl
resource "ibm_backup_recovery_protection_policy" "backup_recovery_protection_policy_instance" {
  x_ibm_tenant_id = var.backup_recovery_protection_policy_x_ibm_tenant_id
  name = var.backup_recovery_protection_policy_name
  backup_policy = var.backup_recovery_protection_policy_backup_policy
  description = var.backup_recovery_protection_policy_description
  blackout_window = var.backup_recovery_protection_policy_blackout_window
  extended_retention = var.backup_recovery_protection_policy_extended_retention
  remote_target_policy = var.backup_recovery_protection_policy_remote_target_policy
  cascaded_targets_config = var.backup_recovery_protection_policy_cascaded_targets_config
  retry_options = var.backup_recovery_protection_policy_retry_options
  data_lock = var.backup_recovery_protection_policy_data_lock
  version = var.backup_recovery_protection_policy_version
  is_cbs_enabled = var.backup_recovery_protection_policy_is_cbs_enabled
  last_modification_time_usecs = var.backup_recovery_protection_policy_last_modification_time_usecs
  template_id = var.backup_recovery_protection_policy_template_id
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| ibmcloud\_api\_key | IBM Cloud API key | `string` | true |
| x_ibm_tenant_id | Specifies the key to be used to encrypt the source credential. If includeSourceCredentials is set to true this key must be specified. | `string` | true |
| name | Specifies the name of the Protection Policy. | `string` | true |
| backup_policy | Specifies the backup schedule and retentions of a Protection Policy. | `object` | true |
| description | Specifies the description of the Protection Policy. | `string` | false |
| blackout_window | List of Blackout Windows. If specified, this field defines blackout periods when new Group Runs are not started. If a Group Run has been scheduled but not yet executed and the blackout period starts, the behavior depends on the policy field AbortInBlackoutPeriod. | `list()` | false |
| extended_retention | Specifies additional retention policies that should be applied to the backup snapshots. A backup snapshot will be retained up to a time that is the maximum of all retention policies that are applicable to it. | `list()` | false |
| remote_target_policy | Specifies the replication, archival and cloud spin targets of Protection Policy. | `object` | false |
| cascaded_targets_config | Specifies the configuration for cascaded replications. Using cascaded replication, replication cluster(Rx) can further replicate and archive the snapshot copies to further targets. Its recommended to create cascaded configuration where protection group will be created. | `list()` | false |
| retry_options | Retry Options of a Protection Policy when a Protection Group run fails. | `object` | false |
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

### Resource: ibm_backup_recovery

```hcl
resource "ibm_backup_recovery" "backup_recovery_instance" {
  x_ibm_tenant_id = var.backup_recovery_x_ibm_tenant_id
  request_initiator_type = var.backup_recovery_request_initiator_type
  name = var.backup_recovery_name
  snapshot_environment = var.backup_recovery_snapshot_environment
  physical_params = var.backup_recovery_physical_params
  mssql_params = var.backup_recovery_mssql_params
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| ibmcloud\_api\_key | IBM Cloud API key | `string` | true |
| x_ibm_tenant_id | Specifies the key to be used to encrypt the source credential. If includeSourceCredentials is set to true this key must be specified. | `string` | true |
| request_initiator_type | Specifies the type of request from UI, which is used for services like magneto to determine the priority of requests. | `string` | false |
| name | Specifies the name of the Recovery. | `string` | true |
| snapshot_environment | Specifies the type of snapshot environment for which the Recovery was performed. | `string` | true |
| physical_params | Specifies the recovery options specific to Physical environment. | `object` | false |
| mssql_params | Specifies the recovery options specific to Sql environment. | `object` | false |

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

### Resource: ibm_backup_recovery_search_indexed_object

```hcl
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
  cassandra_params = var.backup_recovery_search_indexed_object_cassandra_params
  couchbase_params = var.backup_recovery_search_indexed_object_couchbase_params
  email_params = var.backup_recovery_search_indexed_object_email_params
  exchange_params = var.backup_recovery_search_indexed_object_exchange_params
  file_params = var.backup_recovery_search_indexed_object_file_params
  hbase_params = var.backup_recovery_search_indexed_object_hbase_params
  hdfs_params = var.backup_recovery_search_indexed_object_hdfs_params
  hive_params = var.backup_recovery_search_indexed_object_hive_params
  mongodb_params = var.backup_recovery_search_indexed_object_mongodb_params
  ms_groups_params = var.backup_recovery_search_indexed_object_ms_groups_params
  ms_teams_params = var.backup_recovery_search_indexed_object_ms_teams_params
  one_drive_params = var.backup_recovery_search_indexed_object_one_drive_params
  public_folder_params = var.backup_recovery_search_indexed_object_public_folder_params
  sfdc_params = var.backup_recovery_search_indexed_object_sfdc_params
  sharepoint_params = var.backup_recovery_search_indexed_object_sharepoint_params
  uda_params = var.backup_recovery_search_indexed_object_uda_params
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| ibmcloud\_api\_key | IBM Cloud API key | `string` | true |
| x_ibm_tenant_id | Specifies the key to be used to encrypt the source credential. If includeSourceCredentials is set to true this key must be specified. | `string` | true |
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
| cassandra_params | Parameters required to search Cassandra on a cluster. | `object` | false |
| couchbase_params | Parameters required to search CouchBase on a cluster. | `object` | false |
| email_params | Specifies the request parameters to search for emails and email folders. | `object` | false |
| exchange_params | Specifies the parameters which are specific for searching Exchange mailboxes. | `object` | false |
| file_params | Specifies the request parameters to search for files and file folders. | `object` | false |
| hbase_params | Parameters required to search Hbase on a cluster. | `object` | false |
| hdfs_params | Parameters required to search HDFS on a cluster. | `object` | false |
| hive_params | Parameters required to search Hive on a cluster. | `object` | false |
| mongodb_params | Parameters required to search Mongo DB on a cluster. | `object` | false |
| ms_groups_params | Specifies the request params to search for Groups items. | `object` | false |
| ms_teams_params | Specifies the request params to search for Teams items. | `object` | false |
| one_drive_params | Specifies the request parameters to search for files/folders in document libraries. | `object` | false |
| public_folder_params | Specifies the request parameters to search for Public Folder items. | `object` | false |
| sfdc_params | Specifies the parameters which are specific for searching Salesforce records. | `object` | false |
| sharepoint_params | Specifies the request parameters to search for files/folders in document libraries. | `object` | false |
| uda_params | Parameters required to search Universal Data Adapter objects. | `object` | false |

### Resource: ibm_backup_recovery_source_registration

```hcl
resource "ibm_backup_recovery_source_registration" "backup_recovery_source_registration_instance" {
  x_ibm_tenant_id = var.backup_recovery_source_registration_x_ibm_tenant_id
  environment = var.backup_recovery_source_registration_environment
  name = var.backup_recovery_source_registration_name
  connection_id = var.backup_recovery_source_registration_connection_id
  connections = var.backup_recovery_source_registration_connections
  connector_group_id = var.backup_recovery_source_registration_connector_group_id
  data_source_connection_id = var.backup_recovery_source_registration_data_source_connection_id
  advanced_configs = var.backup_recovery_source_registration_advanced_configs
  physical_params = var.backup_recovery_source_registration_physical_params
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| ibmcloud\_api\_key | IBM Cloud API key | `string` | true |
| x_ibm_tenant_id | Specifies the key to be used to encrypt the source credential. If includeSourceCredentials is set to true this key must be specified. | `string` | true |
| environment | Specifies the environment type of the Protection Source. | `string` | true |
| name | The user specified name for this source. | `string` | false |
| connection_id | Specifies the id of the connection from where this source is reachable. This should only be set for a source being registered by a tenant user. This field will be depricated in future. Use connections field. | `string` | false |
| connections | Specfies the list of connections for the source. | `list()` | false |
| connector_group_id | Specifies the connector group id of connector groups. | `number` | false |
| data_source_connection_id | Specifies the id of the connection from where this source is reachable. This should only be set for a source being registered by a tenant user. Also, this is the 'string' of connectionId. This property was added to accommodate for ID values that exceed 2^53 - 1, which is the max value for which JS maintains precision. | `string` | false |
| advanced_configs | Specifies the advanced configuration for a protection source. | `list()` | false |
| physical_params | Specifies parameters to register physical server. | `object` | false |

#### Outputs

| Name | Description |
|------|-------------|
| source_id | ID of top level source object discovered after the registration. |
| source_info | Specifies information about an object. |
| authentication_status | Specifies the status of the authentication during the registration of a Protection Source. 'Pending' indicates the authentication is in progress. 'Scheduled' indicates the authentication is scheduled. 'Finished' indicates the authentication is completed. 'RefreshInProgress' indicates the refresh is in progress. |
| registration_time_msecs | Specifies the time when the source was registered in milliseconds. |
| last_refreshed_time_msecs | Specifies the time when the source was last refreshed in milliseconds. |
| external_metadata | Specifies the External metadata of an entity. |

### Resource: ibm_backup_recovery_update_protection_group_run_request

```hcl
resource "ibm_backup_recovery_update_protection_group_run_request" "backup_recovery_update_protection_group_run_request_instance" {
  x_ibm_tenant_id = var.backup_recovery_update_protection_group_run_request_x_ibm_tenant_id
  update_protection_group_run_params = var.backup_recovery_update_protection_group_run_request_update_protection_group_run_params
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| ibmcloud\_api\_key | IBM Cloud API key | `string` | true |
| x_ibm_tenant_id | Specifies the key to be used to encrypt the source credential. If includeSourceCredentials is set to true this key must be specified. | `string` | true |
| update_protection_group_run_params |  | `list()` | true |

## IBM Backup recovery API data sources

### Data source: ibm_backup_recovery_agent_upgrade_tasks

```hcl
data "ibm_backup_recovery_agent_upgrade_tasks" "backup_recovery_agent_upgrade_tasks_instance" {
  x_ibm_tenant_id = var.backup_recovery_agent_upgrade_tasks_x_ibm_tenant_id
  ids = var.backup_recovery_agent_upgrade_tasks_ids
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| x_ibm_tenant_id | Specifies the key to be used to encrypt the source credential. If includeSourceCredentials is set to true this key must be specified. | `string` | true |
| ids | Specifies IDs of tasks to be fetched. | `list(number)` | false |

#### Outputs

| Name | Description |
|------|-------------|
| tasks | Specifies the list of agent upgrade tasks. |

### Data source: ibm_backup_recovery_connectors_metadata

```hcl
data "ibm_backup_recovery_connectors_metadata" "backup_recovery_connectors_metadata_instance" {
  x_ibm_tenant_id = var.backup_recovery_agent_upgrade_tasks_x_ibm_tenant_id
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| x_ibm_tenant_id | Specifies the key to be used to encrypt the source credential. If includeSourceCredentials is set to true this key must be specified. | `string` | true |

#### Outputs

| Name | Description |
|------|-------------|
| connector_image_metadata | Specifies the list of connector image medata that contains image type and connector image url. |

### Data source: ibm_backup_recovery_data_source_connections

```hcl
data "ibm_backup_recovery_data_source_connections" "backup_recovery_data_source_connections_instance" {
  x_ibm_tenant_id = var.backup_recovery_data_source_connections_x_ibm_tenant_id
  connection_ids = var.backup_recovery_data_source_connections_connection_ids
  connection_names = var.backup_recovery_data_source_connections_connection_names
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| x_ibm_tenant_id | Specifies the key to be used to encrypt the source credential. If includeSourceCredentials is set to true this key must be specified. | `string` | true |
| connection_ids | Specifies the unique IDs of the connections which are to be fetched. | `list(string)` | false |
| connection_names | Specifies the names of the connections which are to be fetched. | `list(string)` | false |

#### Outputs

| Name | Description |
|------|-------------|
| connections |  |

### Data source: ibm_backup_recovery_data_source_connectors

```hcl
data "ibm_backup_recovery_data_source_connectors" "backup_recovery_data_source_connectors_instance" {
  x_ibm_tenant_id = var.backup_recovery_data_source_connectors_x_ibm_tenant_id
  connector_ids = var.backup_recovery_data_source_connectors_connector_ids
  connector_names = var.backup_recovery_data_source_connectors_connector_names
  connection_id = var.backup_recovery_data_source_connectors_connection_id
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| x_ibm_tenant_id | Specifies the key to be used to encrypt the source credential. If includeSourceCredentials is set to true this key must be specified. | `string` | true |
| connector_ids | Specifies the unique IDs of the connectors which are to be fetched. | `list(string)` | false |
| connector_names | Specifies the names of the connectors which are to be fetched. | `list(string)` | false |
| connection_id | Specifies the ID of the connection, connectors belonging to which are to be fetched. | `string` | false |

#### Outputs

| Name | Description |
|------|-------------|
| connectors |  |

### Data source: ibm_backup_recovery_download_agent

```hcl
data "ibm_backup_recovery_download_agent" "backup_recovery_download_agent_instance" {
  x_ibm_tenant_id = var.backup_recovery_download_agent_x_ibm_tenant_id
  platform = var.backup_recovery_download_agent_platform
  linux_params = var.backup_recovery_download_agent_linux_params
  file_path = ""
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| x_ibm_tenant_id | Specifies the key to be used to encrypt the source credential. If includeSourceCredentials is set to true this key must be specified. | `string` | true |
| file_path | Specifies file path. | `string` | true |
| platform | Specifies the platform for which agent needs to be downloaded. | `string` | true |
| linux_params | Linux agent parameters. | `object` | false |

### Data source: ibm_backup_recovery_object_snapshots

```hcl
data "ibm_backup_recovery_object_snapshots" "backup_recovery_object_snapshots_instance" {
  object_id = var.backup_recovery_object_snapshots_backup_recovery_object_snapshots_id
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
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| object_id | Specifies the id of the Object. | `number` | true |
| x_ibm_tenant_id | Specifies the key to be used to encrypt the source credential. If includeSourceCredentials is set to true this key must be specified. | `string` | true |
| from_time_usecs | Specifies the timestamp in Unix time epoch in microseconds to filter Object's snapshots which were taken after this value. | `number` | false |
| to_time_usecs | Specifies the timestamp in Unix time epoch in microseconds to filter Object's snapshots which were taken before this value. | `number` | false |
| run_start_from_time_usecs | Specifies the timestamp in Unix time epoch in microseconds to filter Object's snapshots which were run after this value. | `number` | false |
| run_start_to_time_usecs | Specifies the timestamp in Unix time epoch in microseconds to filter Object's snapshots which were run before this value. | `number` | false |
| snapshot_actions | Specifies a list of recovery actions. Only snapshots that apply to these actions will be returned. | `list(string)` | false |
| run_types | Filter by run type. Only protection runs matching the specified types will be returned. By default, CDP hydration snapshots are not included unless explicitly queried using this field. | `list(string)` | false |
| protection_group_ids | If specified, this returns only the snapshots of the specified object ID, which belong to the provided protection group IDs. | `list(string)` | false |
| run_instance_ids | Filter by a list of run instance IDs. If specified, only snapshots created by these protection runs will be returned. | `list(number)` | false |
| region_ids | Filter by a list of region IDs. | `list(string)` | false |
| object_action_keys | Filter by ObjectActionKey, which uniquely represents the protection of an object. An object can be protected in multiple ways but at most once for a given combination of ObjectActionKey. When specified, only snapshots matching the given action keys are returned for the corresponding object. | `list(string)` | false |

#### Outputs

| Name | Description |
|------|-------------|
| snapshots | Specifies the list of snapshots. |

### Data source: ibm_backup_recovery_search_objects

```hcl
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
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| x_ibm_tenant_id | Specifies the key to be used to encrypt the source credential. If includeSourceCredentials is set to true this key must be specified. | `string` | true |
| request_initiator_type | Specifies the type of request from UI, which is used for services like magneto to determine the priority of requests. | `string` | false |
| search_string | Specifies the search string to filter the objects. This search string will be applicable for objectnames. User can specify a wildcard character '*' as a suffix to a string where all object names are matched with the prefix string. For example, if vm1 and vm2 are the names of objects, user can specify vm* to list the objects. If not specified, then all the objects will be returned which will match other filtering criteria. | `string` | false |
| environments | Specifies the environment type to filter objects. | `list(string)` | false |
| protection_types | Specifies the protection type to filter objects. | `list(string)` | false |
| protection_group_ids | Specifies a list of Protection Group ids to filter the objects. If specified, the objects protected by specified Protection Group ids will be returned. | `list(string)` | false |
| object_ids | Specifies a list of Object ids to filter. | `list(number)` | false |
| os_types | Specifies the operating system types to filter objects on. | `list(string)` | false |
| source_ids | Specifies a list of Protection Source object ids to filter the objects. If specified, the object which are present in those Sources will be returned. | `list(number)` | false |
| source_uuids | Specifies a list of Protection Source object uuids to filter the objects. If specified, the object which are present in those Sources will be returned. | `list(string)` | false |
| is_protected | Specifies the protection status of objects. If set to true, only protected objects will be returned. If set to false, only unprotected objects will be returned. If not specified, all objects will be returned. | `bool` | false |
| is_deleted | If set to true, then objects which are deleted on atleast one cluster will be returned. If not set or set to false then objects which are registered on atleast one cluster are returned. | `bool` | false |
| last_run_status_list | Specifies a list of status of the object's last protection run. Only objects with last run status of these will be returned. | `list(string)` | false |
| cluster_identifiers | Specifies the list of cluster identifiers. Format is clusterId:clusterIncarnationId. Only records from clusters having these identifiers will be returned. | `list(string)` | false |
| include_deleted_objects | Specifies whether to include deleted objects in response. These objects can't be protected but can be recovered. This field is deprecated. | `bool` | false |
| pagination_cookie | Specifies the pagination cookie with which subsequent parts of the response can be fetched. | `string` | false |
| count | Specifies the number of objects to be fetched for the specified pagination cookie. | `number` | false |
| must_have_tag_ids | Specifies tags which must be all present in the document. | `list(string)` | false |
| might_have_tag_ids | Specifies list of tags, one or more of which might be present in the document. These are OR'ed together and the resulting criteria AND'ed with the rest of the query. | `list(string)` | false |
| must_have_snapshot_tag_ids | Specifies snapshot tags which must be all present in the document. | `list(string)` | false |
| might_have_snapshot_tag_ids | Specifies list of snapshot tags, one or more of which might be present in the document. These are OR'ed together and the resulting criteria AND'ed with the rest of the query. | `list(string)` | false |
| tag_search_name | Specifies the tag name to filter the tagged objects and snapshots. User can specify a wildcard character '*' as a suffix to a string where all object's tag names are matched with the prefix string. | `string` | false |
| tag_names | Specifies the tag names to filter the tagged objects and snapshots. | `list(string)` | false |
| tag_types | Specifies the tag names to filter the tagged objects and snapshots. | `list(string)` | false |
| tag_categories | Specifies the tag category to filter the objects and snapshots. | `list(string)` | false |
| tag_sub_categories | Specifies the tag subcategory to filter the objects and snapshots. | `list(string)` | false |
| include_helios_tag_info_for_objects | pecifies whether to include helios tags information for objects in response. Default value is false. | `bool` | false |
| external_filters | Specifies the key-value pairs to filtering the results for the search. Each filter is of the form 'key:value'. The filter 'externalFilters:k1:v1&externalFilters:k2:v2&externalFilters:k2:v3' returns the documents where each document will match the query (k1=v1) AND (k2=v2 OR k2 = v3). Allowed keys: - vmBiosUuid - graphUuid - arn - instanceId - bucketName - azureId. | `list(string)` | false |

#### Outputs

| Name | Description |
|------|-------------|
| objects | Specifies the list of Objects. |

### Data source: ibm_backup_recovery_search_protected_objects

```hcl
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
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| x_ibm_tenant_id | Specifies the key to be used to encrypt the source credential. If includeSourceCredentials is set to true this key must be specified. | `string` | true |
| request_initiator_type | Specifies the type of request from UI, which is used for services like magneto to determine the priority of requests. | `string` | false |
| search_string | Specifies the search string to filter the objects. This search string will be applicable for objectnames and Protection Group names. User can specify a wildcard character '*' as a suffix to a string where all object and their Protection Group names are matched with the prefix string. For example, if vm1 and vm2 are the names of objects, user can specify vm* to list the objects. If not specified, then all the objects with Protection Groups will be returned which will match other filtering criteria. | `string` | false |
| environments | Specifies the environment type to filter objects. | `list(string)` | false |
| snapshot_actions | Specifies a list of recovery actions. Only snapshots that applies to these actions will be returned. | `list(string)` | false |
| object_action_key | Filter by ObjectActionKey, which uniquely represents protection of an object. An object can be protected in multiple ways but atmost once for a given combination of ObjectActionKey. When specified, latest snapshot info matching the objectActionKey is for corresponding object. | `string` | false |
| protection_group_ids | Specifies a list of Protection Group ids to filter the objects. If specified, the objects protected by specified Protection Group ids will be returned. | `list(string)` | false |
| object_ids | Specifies a list of Object ids to filter. | `list(number)` | false |
| sub_result_size | Specifies the size of objects to be fetched for a single subresult. | `number` | false |
| filter_snapshot_from_usecs | Specifies the timestamp in Unix time epoch in microseconds to filter the objects if the Object has a successful snapshot after this value. | `number` | false |
| filter_snapshot_to_usecs | Specifies the timestamp in Unix time epoch in microseconds to filter the objects if the Object has a successful snapshot before this value. | `number` | false |
| os_types | Specifies the operating system types to filter objects on. | `list(string)` | false |
| source_ids | Specifies a list of Protection Source object ids to filter the objects. If specified, the object which are present in those Sources will be returned. | `list(number)` | false |
| run_instance_ids | Specifies a list of run instance ids. If specified only objects belonging to the provided run id will be retunrned. | `list(number)` | false |
| cdp_protected_only | Specifies whether to only return the CDP protected objects. | `bool` | false |
| use_cached_data | Specifies whether we can serve the GET request to the read replica cache cache. There is a lag of 15 seconds between the read replica and primary data source. | `bool` | false |

#### Outputs

| Name | Description |
|------|-------------|
| objects | Specifies the list of Protected Objects. |
| metadata | Specifies the metadata information about the Protection Groups, Protection Policy etc., for search result. |
| num_results | Specifies the total number of search results which matches the search criteria. |

### Data source: ibm_backup_recovery_protection_group

```hcl
data "ibm_backup_recovery_protection_group" "backup_recovery_protection_group_instance" {
  protection_group_id = var.data_backup_recovery_protection_group_protection_group_id
  x_ibm_tenant_id = var.data_backup_recovery_protection_group_x_ibm_tenant_id
  request_initiator_type = var.data_backup_recovery_protection_group_request_initiator_type
  include_last_run_info = var.data_backup_recovery_protection_group_include_last_run_info
  prune_excluded_source_ids = var.data_backup_recovery_protection_group_prune_excluded_source_ids
  prune_source_ids = var.data_backup_recovery_protection_group_prune_source_ids
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| protection_group_id | Specifies a unique id of the Protection Group. | `string` | true |
| x_ibm_tenant_id | Specifies the key to be used to encrypt the source credential. If includeSourceCredentials is set to true this key must be specified. | `string` | true |
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
| physical_params | Specifies the parameters for Physical object. |
| mssql_params | Specifies the parameters specific to MSSQL Protection Group. |

### Data source: ibm_backup_recovery_protection_groups

```hcl
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
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| x_ibm_tenant_id | Specifies the key to be used to encrypt the source credential. If includeSourceCredentials is set to true this key must be specified. | `string` | true |
| request_initiator_type | Specifies the type of request from UI, which is used for services like magneto to determine the priority of requests. | `string` | false |
| ids | Filter by a list of Protection Group ids. | `list(string)` | false |
| names | Filter by a list of Protection Group names. | `list(string)` | false |
| policy_ids | Filter by Policy ids that are associated with Protection Groups. Only Protection Groups associated with the specified Policy ids, are returned. | `list(string)` | false |
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
| include_last_run_info | If true, the response will include last run info. If it is false or not specified, the last run info won't be returned. | `bool` | false |
| prune_excluded_source_ids | If true, the response will not include the list of excluded source IDs in groups that contain this field. This can be set to true in order to improve performance if excluded source IDs are not needed by the user. | `bool` | false |
| prune_source_ids | If true, the response will exclude the list of source IDs within the group specified. | `bool` | false |
| use_cached_data | Specifies whether we can serve the GET request from the read replica cache. There is a lag of 15 seconds between the read replica and primary data source. | `bool` | false |
| source_ids | Filter by Source ids that are associated with Protection Groups. Only Protection Groups associated with the specified Source ids, are returned. | `list(number)` | false |

#### Outputs

| Name | Description |
|------|-------------|
| protection_groups | Specifies the list of Protection Groups which were returned by the request. |

### Data source: ibm_backup_recovery_protection_group_run

```hcl
data "ibm_backup_recovery_protection_group_run" "backup_recovery_protection_group_run_instance" {
  protection_group_run_id = var.backup_recovery_protection_group_run_protection_group_run_id
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
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| protection_group_run_id | Specifies a unique id of the Protection Group. | `string` | true |
| x_ibm_tenant_id | Specifies the key to be used to encrypt the source credential. If includeSourceCredentials is set to true this key must be specified. | `string` | true |
| request_initiator_type | Specifies the type of request from UI, which is used for services like magneto to determine the priority of requests. | `string` | false |
| run_id | Specifies the protection run id. | `string` | false |
| start_time_usecs | Start time for time range filter. Specify the start time as a Unix epoch Timestamp (in microseconds), only runs executing after this time will be returned. By default it is endTimeUsecs minus an hour. | `number` | false |
| end_time_usecs | End time for time range filter. Specify the end time as a Unix epoch Timestamp (in microseconds), only runs executing before this time will be returned. By default it is current time. | `number` | false |
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

### Data source: ibm_backup_recovery_protection_group_runs

```hcl
data "ibm_backup_recovery_protection_group_runs" "backup_recovery_protection_group_runs_instance" {
  protection_group_runs_id = var.backup_recovery_protection_group_runs_protection_group_runs_id
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
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| protection_group_runs_id | Specifies a unique id of the Protection Group. | `string` | true |
| x_ibm_tenant_id | Specifies the key to be used to encrypt the source credential. If includeSourceCredentials is set to true this key must be specified. | `string` | true |
| request_initiator_type | Specifies the type of request from UI, which is used for services like magneto to determine the priority of requests. | `string` | false |
| run_id | Specifies the protection run id. | `string` | false |
| start_time_usecs | Start time for time range filter. Specify the start time as a Unix epoch Timestamp (in microseconds), only runs executing after this time will be returned. By default it is endTimeUsecs minus an hour. | `number` | false |
| end_time_usecs | End time for time range filter. Specify the end time as a Unix epoch Timestamp (in microseconds), only runs executing before this time will be returned. By default it is current time. | `number` | false |
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

### Data source: ibm_backup_recovery_protection_policies

```hcl
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
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| x_ibm_tenant_id | Specifies the key to be used to encrypt the source credential. If includeSourceCredentials is set to true this key must be specified. | `string` | true |
| request_initiator_type | Specifies the type of request from UI, which is used for services like magneto to determine the priority of requests. | `string` | false |
| ids | Filter policies by a list of policy ids. | `list(string)` | false |
| policy_names | Filter policies by a list of policy names. | `list(string)` | false |
| types | Types specifies the policy type of policies to be returned. | `list(string)` | false |
| exclude_linked_policies | If excludeLinkedPolicies is set to true then only local policies created on cluster will be returned. The result will exclude all linked policies created from policy templates. | `bool` | false |
| include_replicated_policies | If includeReplicatedPolicies is set to true, then response will also contain replicated policies. By default, replication policies are not included in the response. | `bool` | false |
| include_stats | If includeStats is set to true, then response will return number of protection groups and objects. By default, the protection stats are not included in the response. | `bool` | false |

#### Outputs

| Name | Description |
|------|-------------|
| policies | Specifies a list of protection policies. |

### Data source: ibm_backup_recovery_protection_policy

```hcl
data "ibm_backup_recovery_protection_policy" "backup_recovery_protection_policy_instance" {
  protection_policy_id = var.data_backup_recovery_protection_policy_protection_policy_id
  x_ibm_tenant_id = var.data_backup_recovery_protection_policy_x_ibm_tenant_id
  request_initiator_type = var.data_backup_recovery_protection_policy_request_initiator_type
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| protection_policy_id | Specifies a unique id of the Protection Policy to return. | `string` | true |
| x_ibm_tenant_id | Specifies the key to be used to encrypt the source credential. If includeSourceCredentials is set to true this key must be specified. | `string` | true |
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

### Data source: ibm_backup_recovery

```hcl
data "ibm_backup_recovery" "backup_recovery_instance" {
  recovery_id = var.data_backup_recovery_backup_recovery_id
  x_ibm_tenant_id = var.data_backup_recovery_x_ibm_tenant_id
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| recovery_id | Specifies the id of a Recovery. | `string` | true |
| x_ibm_tenant_id | Specifies the key to be used to encrypt the source credential. If includeSourceCredentials is set to true this key must be specified. | `string` | true |

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
| mssql_params | Specifies the recovery options specific to Sql environment. |



### Data source: ibm_backup_recovery_download_files

```hcl
data "ibm_backup_recovery_download_files" "backup_recovery_download_files_instance" {
  recovery_download_files_id = "recovery_download_files_id"
  x_ibm_tenant_id = var.data_backup_recovery_x_ibm_tenant_id
  start_offset = 0
  length = 0
  file_type = "file_type"
  source_name = "source_name"
  start_time = "start_time"
  include_tenants = false
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| recovery_download_files_id | Specifies the id of a Recovery. | `string` | true |
| x_ibm_tenant_id | Specifies the key to be used to encrypt the source credential. If includeSourceCredentials is set to true this key must be specified. | `string` | true |
| start_offset | Specifies the start offset of file chunk to be downloaded. | `int` | false |
| length | Specifies the length of bytes to download. This can not be greater than 8MB (8388608 byets). | `int` | false |
| file_type | Specifies the downloaded type, i.e: error, success_files_list. | `string` | false |
| source_name | Specifies the name of the source on which restore is done. | `string` | false |
| start_time | Specifies the start time of restore task. | `string` | false |
| include_tenants | Specifies if objects of all the organizations under the hierarchy of the logged in user's organization should be returned.| `bool` | false |

### Data source: ibm_backup_recoveries

```hcl
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
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| x_ibm_tenant_id | Specifies the key to be used to encrypt the source credential. If includeSourceCredentials is set to true this key must be specified. | `string` | true |
| ids | Filter Recoveries for given ids. | `list(string)` | false |
| return_only_child_recoveries | Returns only child recoveries if passed as true. This filter should always be used along with 'ids' filter. | `bool` | false |
| start_time_usecs | Returns the recoveries which are started after the specific time. This value should be in Unix timestamp epoch in microseconds. | `number` | false |
| end_time_usecs | Returns the recoveries which are started before the specific time. This value should be in Unix timestamp epoch in microseconds. | `number` | false |
| snapshot_target_type | Specifies the snapshot's target type from which recovery has been performed. | `list(string)` | false |
| archival_target_type | Specifies the snapshot's archival target type from which recovery has been performed. This parameter applies only if 'snapshotTargetType' is 'Archival'. | `list(string)` | false |
| snapshot_environments | Specifies the list of snapshot environment types to filter Recoveries. If empty, Recoveries related to all environments will be returned. | `list(string)` | false |
| status | Specifies the list of run status to filter Recoveries. If empty, Recoveries with all run status will be returned. | `list(string)` | false |
| recovery_actions | Specifies the list of recovery actions to filter Recoveries. If empty, Recoveries related to all actions will be returned. | `list(string)` | false |

#### Outputs

| Name | Description |
|------|-------------|
| recoveries | Specifies list of Recoveries. |

### Data source: ibm_backup_recovery_source_registrations

```hcl
data "ibm_backup_recovery_source_registrations" "backup_recovery_source_registrations_instance" {
  x_ibm_tenant_id = var.backup_recovery_source_registrations_x_ibm_tenant_id
  ids = var.backup_recovery_source_registrations_ids
  include_source_credentials = var.backup_recovery_source_registrations_include_source_credentials
  encryption_key = var.backup_recovery_source_registrations_encryption_key
  use_cached_data = var.backup_recovery_source_registrations_use_cached_data
  include_external_metadata = var.backup_recovery_source_registrations_include_external_metadata
  ignore_tenant_migration_in_progress_check = var.backup_recovery_source_registrations_ignore_tenant_migration_in_progress_check
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| x_ibm_tenant_id | Specifies the key to be used to encrypt the source credential. If includeSourceCredentials is set to true this key must be specified. | `string` | true |
| ids | Ids specifies the list of source registration ids to return. If left empty, every source registration will be returned by default. | `list(number)` | false |
| include_source_credentials | If true, the encrypted crednetial for the registered sources will be included. Credential is first encrypted with internal key and then reencrypted with user supplied encryption key. | `bool` | false |
| encryption_key | Specifies the key to be used to encrypt the source credential. If includeSourceCredentials is set to true this key must be specified. | `string` | false |
| use_cached_data | Specifies whether we can serve the GET request from the read replica cache. There is a lag of 15 seconds between the read replica and primary data source. | `bool` | false |
| include_external_metadata | If true, the external entity metadata like maintenance mode config for the registered sources will be included. | `bool` | false |
| ignore_tenant_migration_in_progress_check | If true, tenant migration check will be ignored. | `bool` | false |

#### Outputs

| Name | Description |
|------|-------------|
| registrations | Specifies the list of Protection Source Registrations. |

### Data source: ibm_backup_recovery_source_registration

```hcl
data "ibm_backup_recovery_source_registration" "backup_recovery_source_registration_instance" {
  source_registration_id = var.data_backup_recovery_source_registration_source_registration_id
  x_ibm_tenant_id = var.data_backup_recovery_source_registration_x_ibm_tenant_id
  request_initiator_type = var.data_backup_recovery_source_registration_request_initiator_type
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| source_registration_id | Specifies the id of the Protection Source registration. | `number` | true |
| x_ibm_tenant_id | Specifies the key to be used to encrypt the source credential. If includeSourceCredentials is set to true this key must be specified. | `string` | true |
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
| data_source_connection_id | Specifies the id of the connection from where this source is reachable. This should only be set for a source being registered by a tenant user. Also, this is the 'string' of connectionId. This property was added to accommodate for ID values that exceed 2^53 - 1, which is the max value for which JS maintains precision. |
| advanced_configs | Specifies the advanced configuration for a protection source. |
| authentication_status | Specifies the status of the authentication during the registration of a Protection Source. 'Pending' indicates the authentication is in progress. 'Scheduled' indicates the authentication is scheduled. 'Finished' indicates the authentication is completed. 'RefreshInProgress' indicates the refresh is in progress. |
| registration_time_msecs | Specifies the time when the source was registered in milliseconds. |
| last_refreshed_time_msecs | Specifies the time when the source was last refreshed in milliseconds. |
| external_metadata | Specifies the External metadata of an entity. |
| physical_params | Specifies parameters to register physical server. |

### Data source: ibm_backup_recovery_download_indexed_files

```hcl
data "ibm_backup_recovery_download_indexed_files" "backup_recovery_download_indexed_files_instance" {
  snapshots_id = var.backup_recovery_download_indexed_files_snapshots_id
  x_ibm_tenant_id = var.backup_recovery_download_indexed_files_x_ibm_tenant_id
  file_path = var.backup_recovery_download_indexed_files_file_path
  nvram_file = var.backup_recovery_download_indexed_files_nvram_file
  retry_attempt = var.backup_recovery_download_indexed_files_retry_attempt
  start_offset = var.backup_recovery_download_indexed_files_start_offset
  length = var.backup_recovery_download_indexed_files_length
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| snapshots_id | Specifies the snapshot id to download from. | `string` | true |
| x_ibm_tenant_id | Specifies the key to be used to encrypt the source credential. If includeSourceCredentials is set to true this key must be specified. | `string` | true |
| file_path | Specifies the path to the file to download. If no path is specified and snapshot environment is kVMWare, VMX file for VMware will be downloaded. For other snapshot environments, this field must be specified. | `string` | false |
| nvram_file | Specifies if NVRAM file for VMware should be downloaded. | `bool` | false |
| retry_attempt | Specifies the number of attempts the protection run took to create this file. | `number` | false |
| start_offset | Specifies the start offset of file chunk to be downloaded. | `number` | false |
| length | Specifies the length of bytes to download. This can not be greater than 8MB (8388608 byets). | `number` | false |

### Data source: ibm_backup_recovery_protection_sources

```hcl
data "ibm_backup_recovery_protection_sources" "backup_recovery_protection_sources_instance" {
  x_ibm_tenant_id = var.backup_recovery_protection_sources_x_ibm_tenant_id
  exclude_office365_types = var.backup_recovery_protection_sources_exclude_office365_types
  get_teams_channels = var.backup_recovery_protection_sources_get_teams_channels
  after_cursor_entity_id = var.backup_recovery_protection_sources_after_cursor_entity_id
  before_cursor_entity_id = var.backup_recovery_protection_sources_before_cursor_entity_id
  node_id = var.backup_recovery_protection_sources_node_id
  page_size = var.backup_recovery_protection_sources_page_size
  has_valid_mailbox = var.backup_recovery_protection_sources_has_valid_mailbox
  has_valid_onedrive = var.backup_recovery_protection_sources_has_valid_onedrive
  is_security_group = var.backup_recovery_protection_sources_is_security_group
  backup_recovery_protection_sources_id = var.backup_recovery_protection_sources_backup_recovery_protection_sources_id
  num_levels = var.backup_recovery_protection_sources_num_levels
  exclude_types = var.backup_recovery_protection_sources_exclude_types
  exclude_aws_types = var.backup_recovery_protection_sources_exclude_aws_types
  exclude_kubernetes_types = var.backup_recovery_protection_sources_exclude_kubernetes_types
  include_datastores = var.backup_recovery_protection_sources_include_datastores
  include_networks = var.backup_recovery_protection_sources_include_networks
  include_vm_folders = var.backup_recovery_protection_sources_include_vm_folders
  include_sfdc_fields = var.backup_recovery_protection_sources_include_sfdc_fields
  include_system_v_apps = var.backup_recovery_protection_sources_include_system_v_apps
  environments = var.backup_recovery_protection_sources_environments
  environment = var.backup_recovery_protection_sources_environment
  include_entity_permission_info = var.backup_recovery_protection_sources_include_entity_permission_info
  sids = var.backup_recovery_protection_sources_sids
  include_source_credentials = var.backup_recovery_protection_sources_include_source_credentials
  encryption_key = var.backup_recovery_protection_sources_encryption_key
  include_object_protection_info = var.backup_recovery_protection_sources_include_object_protection_info
  prune_non_critical_info = var.backup_recovery_protection_sources_prune_non_critical_info
  prune_aggregation_info = var.backup_recovery_protection_sources_prune_aggregation_info
  request_initiator_type = var.backup_recovery_protection_sources_request_initiator_type
  use_cached_data = var.backup_recovery_protection_sources_use_cached_data
  all_under_hierarchy = var.backup_recovery_protection_sources_all_under_hierarchy
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| x_ibm_tenant_id | Specifies the key to be used to encrypt the source credential. If includeSourceCredentials is set to true this key must be specified. | `string` | true |
| exclude_office365_types | Specifies the Object types to be filtered out for Office 365 that match the passed in types such as 'kDomain', 'kOutlook', 'kMailbox', etc. For example, set this parameter to 'kMailbox' to exclude Mailbox Objects from being returned. | `list(string)` | false |
| get_teams_channels | Filter policies by a list of policy ids. | `bool` | false |
| after_cursor_entity_id | Specifies the entity id starting from which the items are to be returned. | `number` | false |
| before_cursor_entity_id | Specifies the entity id upto which the items are to be returned. | `number` | false |
| node_id | Specifies the entity id for the Node at any level within the Source entity hierarchy whose children are to be paginated. | `number` | false |
| page_size | Specifies the maximum number of entities to be returned within the page. | `number` | false |
| has_valid_mailbox | If set to true, users with valid mailbox will be returned. | `bool` | false |
| has_valid_onedrive | If set to true, users with valid onedrive will be returned. | `bool` | false |
| is_security_group | If set to true, Groups which are security enabled will be returned. | `bool` | false |
| backup_recovery_protection_sources_id | Return the Object subtree for the passed in Protection Source id. | `number` | false |
| num_levels | Specifies the expected number of levels from the root node to be returned in the entity hierarchy response. | `number` | false |
| exclude_types | Filter out the Object types (and their subtrees) that match the passed in types such as 'kVCenter', 'kFolder', 'kDatacenter', 'kComputeResource', 'kResourcePool', 'kDatastore', 'kHostSystem', 'kVirtualMachine', etc. For example, set this parameter to 'kResourcePool' to exclude Resource Pool Objects from being returned. | `list(string)` | false |
| exclude_aws_types | Specifies the Object types to be filtered out for AWS that match the passed in types such as 'kEC2Instance', 'kRDSInstance', 'kAuroraCluster', 'kTag', 'kAuroraTag', 'kRDSTag', kS3Bucket, kS3Tag. For example, set this parameter to 'kEC2Instance' to exclude ec2 instance from being returned. | `list(string)` | false |
| exclude_kubernetes_types | Specifies the Object types to be filtered out for Kubernetes that match the passed in types such as 'kService'. For example, set this parameter to 'kService' to exclude services from being returned. | `list(string)` | false |
| include_datastores | Set this parameter to true to also return kDatastore object types found in the Source in addition to their Object subtrees. By default, datastores are not returned. | `bool` | false |
| include_networks | Set this parameter to true to also return kNetwork object types found in the Source in addition to their Object subtrees. By default, network objects are not returned. | `bool` | false |
| include_vm_folders | Set this parameter to true to also return kVMFolder object types found in the Source in addition to their Object subtrees. By default, VM folder objects are not returned. | `bool` | false |
| include_sfdc_fields | Set this parameter to true to also return fields of the object found in the Source in addition to their Object subtrees. By default, Sfdc object fields are not returned. | `bool` | false |
| include_system_v_apps | Set this parameter to true to also return system VApp object types found in the Source in addition to their Object subtrees. By default, VM folder objects are not returned. | `bool` | false |
| environments | Return only Protection Sources that match the passed in environment type such as 'kVMware', 'kSQL', 'kView' 'kPhysical', 'kPuppeteer', 'kPure', 'kNetapp', 'kGenericNas', 'kHyperV', 'kAcropolis', or 'kAzure'. For example, set this parameter to 'kVMware' to only return the Sources (and their Object subtrees) found in the 'kVMware' (VMware vCenter Server) environment. | `list(string)` | false |
| environment | This field is deprecated. Use environments instead. | `string` | false |
| include_entity_permission_info | If specified, then a list of entites with permissions assigned to them are returned. | `bool` | false |
| sids | Filter the object subtree for the sids given in the list. | `list(string)` | false |
| include_source_credentials | If specified, then crednetial for the registered sources will be included. Credential is first encrypted with internal key and then reencrypted with user supplied 'encryption_key'. | `bool` | false |
| encryption_key | Key to be used to encrypt the source credential. If include_source_credentials is set to true this key must be specified. | `string` | false |
| include_object_protection_info | If specified, the object protection of entities(if any) will be returned. | `bool` | false |
| prune_non_critical_info | Specifies whether to prune non critical info within entities. Incase of VMs, virtual disk information will be pruned. Incase of Office365, metadata about user entities will be pruned. This can be used to limit the size of the response by caller. | `bool` | false |
| prune_aggregation_info | Specifies whether to prune the aggregation information about the number of entities protected/unprotected. | `bool` | false |
| request_initiator_type | Specifies the type of the request. Possible values are UIUser and UIAuto, which means the request is triggered by user or is an auto refresh request. Services like magneto will use this to determine the priority of the requests, so that it can more intelligently handle overload situations by prioritizing higher priority requests. | `string` | false |
| use_cached_data | Specifies whether we can serve the GET request to the read replica cache. setting this to true ensures that the API request is served to the read replica. setting this to false will serve the request to the master. | `bool` | false |
| all_under_hierarchy | AllUnderHierarchy specifies if objects of all the tenants under the hierarchy of the logged in user's organization should be returned. | `bool` | false |

#### Outputs

| Name | Description |
|------|-------------|
| protection_sources | Specifies list of protection sources. |

## Assumptions

1. End user has connector endpoint URL
2. End user has tenantID
3. End user has backup recovery cluster URL

## Notes

### Backup Recovery Cluster URL can be set using environment variable or in endpoints.json respectively
Example:  
```
  export IBMCLOUD_BACKUP_RECOVERY_ENDPOINT=https://brs-stage-us-south-02.backup-recovery.test.cloud.ibm.com/v2
  export BACKUP_RECOVERY_CONNECTOR_ENDPOINT=https://1.2.3.4
```
OR
```
  {
    "IBMCLOUD_BACKUP_RECOVERY_ENDPOINT" : {
        "public" : {
            "us-south": "https://brs-stage-us-south-01.backup-recovery.test.cloud.ibm.com/v2"
        }
    }
}
```

### Resources with incomplete CRUD operations
This service includes certain resources that do not have fully implemented CRUD (Create, Read, Update, Delete) operations due to limitations in the underlying APIs. Specifically:

#### Protection Group Run:

***Create:*** A `ibm_backup_recovery_protection_group_run_request` resource is available for creating new protection group run.

***Update:*** protection group run updates are managed through a separate resource `ibm_backup_recovery_update_protection_group_run_request`. 
Note that the `ibm_backup_recovery_protection_group_run_request` and `ibm_backup_recovery_update_protection_group_run_request` resources must be used in tandem to manage Protection Group Runs.

***Delete:*** There is no delete operation available for the protection group run resource. If  ibm_backup_recovery_update_protection_group_run_request or ibm_backup_recovery_protection_group_run_request resource is removed from the `main.tf` file, Terraform will remove it from the state file but not from the backend. The resource will continue to exist in the backend system.


#### Other resources that do not support update and delete:

Some resources in this service do not support update or delete operations due to the absence of corresponding API endpoints. As a result, Terraform cannot manage these operations for those resources. Users should be aware that removing these resources from the configuration (main.tf) will only remove them from the Terraform state and will not affect the actual resources in the backend. Similarly updating these resources will throw an error in the plan phase stating that the resource cannot be updated.
- ibm_backup_recovery_perform_action_on_protection_group_run_request
- ibm_backup_recovery_download_files_folders
- ibm_backup_recovery_agent_upgrade_task
- ibm_backup_recovery_protection_group_state
- ibm_backup_recovery_connection_registration_token
- ibm_backup_recovery_restore_points
- ibm_backup_recovery
- ibm_backup_recovery_data_source_connector_patch
- ibm_backup_recovery_data_source_connector_registration
- ibm_backup_recovery_update_protection_group_run_request


**Important:** When managing resources that lack complete CRUD operations, users should exercise caution and consider the limitations described above. Manual intervention may be required to manage these resources in the backend if updates or deletions are necessary.**


### Import
Backup and recovery resources can be imported by using the id. The ID is formed using tenantID and resourceId.
`id = <tenantId>::<resourceId>`. 

#### Syntax
```
import {
	to = <ibm_backup_recovery_resource>
	id = "<tenantId>::<resourceId>"
}
```

#### Example
```
resource "ibm_backup_recovery_data_source_connection" "baas_data_source_connection_instance" {
	x_ibm_tenant_id = "jhxqx715r9/"
	connection_name = "terraform-conn-1"
}

import {
	to = ibm_backup_recovery_data_source_connection.baas_data_source_connection_instance
	id = "jhxqx715r9/::3309023926479362560"
}
```
#### List of resources that support import:
- ibm_backup_recovery_data_source_connection
- ibm_backup_recovery_protection_group
- ibm_backup_recovery_protection_policy
- ibm_backup_recovery
- ibm_backup_recovery_source_registration

## Requirements

| Name | Version |
|------|---------|
| terraform | ~> 0.12 |

## Providers

| Name | Version |
|------|---------|
| ibm | 1.13.1 |
