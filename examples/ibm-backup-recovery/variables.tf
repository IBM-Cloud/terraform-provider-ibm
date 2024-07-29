variable "ibmcloud_api_key" {
  description = "IBM Cloud API key"
  type        = string
}

// Resource arguments for protection_group_run_request
variable "protection_group_run_request_run_type" {
  description = "Type of protection run. 'kRegular' indicates an incremental (CBT) backup. Incremental backups utilizing CBT (if supported) are captured of the target protection objects. The first run of a kRegular schedule captures all the blocks. 'kFull' indicates a full (no CBT) backup. A complete backup (all blocks) of the target protection objects are always captured and Change Block Tracking (CBT) is not utilized. 'kLog' indicates a Database Log backup. Capture the database transaction logs to allow rolling back to a specific point in time. 'kSystem' indicates system volume backup. It produces an image for bare metal recovery."
  type        = string
  default     = "kRegular"
}

// Resource arguments for recovery_download_files_folders
variable "recovery_download_files_folders_name" {
  description = "Specifies the name of the recovery task. This field must be set and must be a unique name."
  type        = string
  default     = "name"
}
variable "recovery_download_files_folders_parent_recovery_id" {
  description = "If current recovery is child task triggered through another parent recovery operation, then this field will specify the id of the parent recovery."
  type        = string
  default     = "parent_recovery_id"
}
variable "recovery_download_files_folders_glacier_retrieval_type" {
  description = "Specifies the glacier retrieval type when restoring or downloding files or folders from a Glacier-based cloud snapshot."
  type        = string
  default     = "kStandard"
}

// Resource arguments for perform_action_on_protection_group_run_request
variable "perform_action_on_protection_group_run_request_action" {
  description = "Specifies the type of the action which will be performed on protection runs."
  type        = string
  default     = "Pause"
}

// Resource arguments for protection_group
variable "protection_group_name" {
  description = "Specifies the name of the Protection Group."
  type        = string
  default     = "name"
}
variable "protection_group_policy_id" {
  description = "Specifies the unique id of the Protection Policy associated with the Protection Group. The Policy provides retry settings Protection Schedules, Priority, SLA, etc."
  type        = string
  default     = "policy_id"
}
variable "protection_group_priority" {
  description = "Specifies the priority of the Protection Group."
  type        = string
  default     = "kLow"
}
variable "protection_group_storage_domain_id" {
  description = "Specifies the Storage Domain (View Box) ID where this Protection Group writes data."
  type        = number
  default     = 1
}
variable "protection_group_description" {
  description = "Specifies a description of the Protection Group."
  type        = string
  default     = "description"
}
variable "protection_group_end_time_usecs" {
  description = "Specifies the end time in micro seconds for this Protection Group. If this is not specified, the Protection Group won't be ended."
  type        = number
  default     = 1
}
variable "protection_group_last_modified_timestamp_usecs" {
  description = "Specifies the last time this protection group was updated. If this is passed into a PUT request, then the backend will validate that the timestamp passed in matches the time that the protection group was actually last modified. If the two timestamps do not match, then the request will be rejected with a stale error."
  type        = number
  default     = 1
}
variable "protection_group_qos_policy" {
  description = "Specifies whether the Protection Group will be written to HDD or SSD."
  type        = string
  default     = "kBackupHDD"
}
variable "protection_group_abort_in_blackouts" {
  description = "Specifies whether currently executing jobs should abort if a blackout period specified by a policy starts. Available only if the selected policy has at least one blackout period. Default value is false."
  type        = bool
  default     = true
}
variable "protection_group_pause_in_blackouts" {
  description = "Specifies whether currently executing jobs should be paused if a blackout period specified by a policy starts. Available only if the selected policy has at least one blackout period. Default value is false. This field should not be set to true if 'abortInBlackouts' is sent as true."
  type        = bool
  default     = true
}
variable "protection_group_is_paused" {
  description = "Specifies if the the Protection Group is paused. New runs are not scheduled for the paused Protection Groups. Active run if any is not impacted."
  type        = bool
  default     = true
}
variable "protection_group_environment" {
  description = "Specifies the environment of the Protection Group."
  type        = string
  default     = "kPhysical"
}

// Resource arguments for protection_policy
variable "protection_policy_name" {
  description = "Specifies the name of the Protection Policy."
  type        = string
  default     = "name"
}
variable "protection_policy_description" {
  description = "Specifies the description of the Protection Policy."
  type        = string
  default     = "description"
}
variable "protection_policy_data_lock" {
  description = "This field is now deprecated. Please use the DataLockConfig in the backup retention."
  type        = string
  default     = "Compliance"
}
variable "protection_policy_version" {
  description = "Specifies the current policy verison. Policy version is incremented for optionally supporting new features and differentialting across releases."
  type        = number
  default     = 1
}
variable "protection_policy_is_cbs_enabled" {
  description = "Specifies true if Calender Based Schedule is supported by client. Default value is assumed as false for this feature."
  type        = bool
  default     = true
}
variable "protection_policy_last_modification_time_usecs" {
  description = "Specifies the last time this Policy was updated. If this is passed into a PUT request, then the backend will validate that the timestamp passed in matches the time that the policy was actually last modified. If the two timestamps do not match, then the request will be rejected with a stale error."
  type        = number
  default     = 1
}
variable "protection_policy_template_id" {
  description = "Specifies the parent policy template id to which the policy is linked to. This field is set only when policy is created from template."
  type        = string
  default     = "template_id"
}

// Resource arguments for recovery
variable "recovery_request_initiator_type" {
  description = "Specifies the type of request from UI, which is used for services like magneto to determine the priority of requests."
  type        = string
  default     = "UIUser"
}
variable "recovery_name" {
  description = "Specifies the name of the Recovery."
  type        = string
  default     = "name"
}
variable "recovery_snapshot_environment" {
  description = "Specifies the type of snapshot environment for which the Recovery was performed."
  type        = string
  default     = "kPhysical"
}

// Resource arguments for search_indexed_object
variable "search_indexed_object_protection_group_ids" {
  description = "Specifies a list of Protection Group ids to filter the indexed objects. If specified, the objects indexed by specified Protection Group ids will be returned."
  type        = list(string)
  default     = [ "protectionGroupIds" ]
}
variable "search_indexed_object_storage_domain_ids" {
  description = "Specifies the Storage Domain ids to filter indexed objects for which Protection Groups are writing data to Cohesity Views on the specified Storage Domains."
  type        = list(number)
  default     = [ 1 ]
}
variable "search_indexed_object_tenant_id" {
  description = "TenantId contains id of the tenant for which objects are to be returned."
  type        = string
  default     = "tenant_id"
}
variable "search_indexed_object_include_tenants" {
  description = "If true, the response will include objects which belongs to all tenants which the current user has permission to see. Default value is false."
  type        = bool
  default     = true
}
variable "search_indexed_object_tags" {
  description = ""This field is deprecated. Please use mightHaveTagIds."."
  type        = list(string)
  default     = [ "tags" ]
}
variable "search_indexed_object_snapshot_tags" {
  description = ""This field is deprecated. Please use mightHaveSnapshotTagIds."."
  type        = list(string)
  default     = [ "snapshotTags" ]
}
variable "search_indexed_object_must_have_tag_ids" {
  description = "Specifies tags which must be all present in the document."
  type        = list(string)
  default     = [ "mustHaveTagIds" ]
}
variable "search_indexed_object_might_have_tag_ids" {
  description = "Specifies list of tags, one or more of which might be present in the document. These are OR'ed together and the resulting criteria AND'ed with the rest of the query."
  type        = list(string)
  default     = [ "mightHaveTagIds" ]
}
variable "search_indexed_object_must_have_snapshot_tag_ids" {
  description = "Specifies snapshot tags which must be all present in the document."
  type        = list(string)
  default     = [ "mustHaveSnapshotTagIds" ]
}
variable "search_indexed_object_might_have_snapshot_tag_ids" {
  description = "Specifies list of snapshot tags, one or more of which might be present in the document. These are OR'ed together and the resulting criteria AND'ed with the rest of the query."
  type        = list(string)
  default     = [ "mightHaveSnapshotTagIds" ]
}
variable "search_indexed_object_pagination_cookie" {
  description = "Specifies the pagination cookie with which subsequent parts of the response can be fetched."
  type        = string
  default     = "pagination_cookie"
}
variable "search_indexed_object_count" {
  description = "Specifies the number of indexed objects to be fetched for the specified pagination cookie."
  type        = number
  default     = 1
}
variable "search_indexed_object_object_type" {
  description = "Specifies the object type to be searched for."
  type        = string
  default     = "Files"
}
variable "search_indexed_object_use_cached_data" {
  description = "Specifies whether we can serve the GET request from the read replica cache. There is a lag of 15 seconds between the read replica and primary data source."
  type        = bool
  default     = true
}

// Resource arguments for source_registration
variable "source_registration_environment" {
  description = "Specifies the environment type of the Protection Source."
  type        = string
  default     = "kPhysical"
}
variable "source_registration_name" {
  description = "The user specified name for this source."
  type        = string
  default     = "name"
}
variable "source_registration_connection_id" {
  description = "Specifies the id of the connection from where this source is reachable. This should only be set for a source being registered by a tenant user. This field will be depricated in future. Use connections field."
  type        = number
  default     = 1
}
variable "source_registration_connector_group_id" {
  description = "Specifies the connector group id of connector groups."
  type        = number
  default     = 1
}

// Resource arguments for update_protection_group_run_request

// Resource arguments for protection_group_state
variable "protection_group_state_action" {
  description = "Specifies the action to be performed on all the specfied Protection Groups. 'kActivate' specifies that Protection Group should be activated. 'kDeactivate' sepcifies that Protection Group should be deactivated. 'kPause' specifies that Protection Group should be paused. 'kResume' specifies that Protection Group should be resumed."
  type        = string
  default     = "kPause"
}
variable "protection_group_state_ids" {
  description = "Specifies a list of Protection Group ids for which the state should change."
  type        = list(string)
  default     = [ "ids" ]
}

// Data source arguments for run_debug_logs
variable "run_debug_logs_run_debug_logs_id" {
  description = "Specifies a unique id of the Protection Group."
  type        = string
  default     = "run_debug_logs_id"
}
variable "run_debug_logs_run_id" {
  description = "Specifies a unique run id of the Protection Group run."
  type        = string
  default     = "run_id"
}
variable "run_debug_logs_object_id" {
  description = "Specifies the id of the object for which debug logs are to be returned."
  type        = string
  default     = "placeholder"
}

// Data source arguments for object_run_debug_logs
variable "object_run_debug_logs_object_run_debug_logs_id" {
  description = "Specifies a unique id of the Protection Group."
  type        = string
  default     = "object_run_debug_logs_id"
}
variable "object_run_debug_logs_run_id" {
  description = "Specifies a unique run id of the Protection Group run."
  type        = string
  default     = "run_id"
}
variable "object_run_debug_logs_object_id" {
  description = "Specifies the id of the object for which debug logs are to be returned."
  type        = string
  default     = "object_id"
}

// Data source arguments for run_error_report
variable "run_error_report_run_error_report_id" {
  description = "Specifies a unique id of the Protection Group."
  type        = string
  default     = "run_error_report_id"
}
variable "run_error_report_run_id" {
  description = "Specifies a unique run id of the Protection Group run."
  type        = string
  default     = "run_id"
}
variable "run_error_report_object_id" {
  description = "Specifies the id of the object for which errors/warnings are to be returned."
  type        = string
  default     = "object_id"
}

// Data source arguments for runs_report
variable "runs_report_runs_report_id" {
  description = "Specifies a unique id of the Protection Group."
  type        = string
  default     = "runs_report_id"
}
variable "runs_report_run_id" {
  description = "Specifies a unique run id of the Protection Group run."
  type        = string
  default     = "run_id"
}
variable "runs_report_object_id" {
  description = "Specifies the id of the object for which errors/warnings are to be returned."
  type        = string
  default     = "object_id"
}
variable "runs_report_file_type" {
  description = "Specifies the downloaded type, i.e: success_files_list, default: success_files_list."
  type        = string
  default     = "placeholder"
}
variable "runs_report_name" {
  description = "Specifies the name of the source being backed up."
  type        = string
  default     = "placeholder"
}

// Data source arguments for recovery_debug_logs
variable "recovery_debug_logs_recovery_debug_logs_id" {
  description = "Specifies the id of a Recovery job."
  type        = string
  default     = "recovery_debug_logs_id"
}

// Data source arguments for recovery_download_messages
variable "recovery_download_messages_recovery_download_messages_id" {
  description = "Specifies a unique ID of a Recovery."
  type        = string
  default     = "recovery_download_messages_id"
}

// Data source arguments for recovery_download_files
variable "recovery_download_files_recovery_download_files_id" {
  description = "Specifies the id of a Recovery."
  type        = string
  default     = "recovery_download_files_id"
}
variable "recovery_download_files_start_offset" {
  description = "Specifies the start offset of file chunk to be downloaded."
  type        = number
  default     = 0
}
variable "recovery_download_files_length" {
  description = "Specifies the length of bytes to download. This can not be greater than 8MB (8388608 byets)."
  type        = number
  default     = 0
}
variable "recovery_download_files_file_type" {
  description = "Specifies the downloaded type, i.e: error, success_files_list."
  type        = string
  default     = "placeholder"
}
variable "recovery_download_files_source_name" {
  description = "Specifies the name of the source on which restore is done."
  type        = string
  default     = "placeholder"
}
variable "recovery_download_files_start_time" {
  description = "Specifies the start time of restore task."
  type        = string
  default     = "placeholder"
}
variable "recovery_download_files_include_tenants" {
  description = "Specifies if objects of all the organizations under the hierarchy of the logged in user's organization should be returned."
  type        = bool
  default     = false
}

// Data source arguments for recovery_fetch_uptier_data
variable "recovery_fetch_uptier_data_archive_u_id" {
  description = "Archive UID of the current restore."
  type        = string
  default     = "archive_u_id"
}

// Data source arguments for protection_run_progress
variable "protection_run_progress_run_id" {
  description = "Specifies a unique run id of the Protection Run."
  type        = string
  default     = "run_id"
}
variable "protection_run_progress_objects" {
  description = "Specifies the objects whose progress will be returned. This only applies to protection group runs and will be ignored for object runs. If the objects are specified, the run progress will not be returned and only the progress of the specified objects will be returned."
  type        = list(number)
  default     = [ 0 ]
}
variable "protection_run_progress_tenant_ids" {
  description = "TenantIds contains ids of the tenants for which the run is to be returned."
  type        = list(string)
  default     = [ "placeholder" ]
}
variable "protection_run_progress_include_tenants" {
  description = "If true, the response will include Protection Group Runs which were created by all tenants which the current user has permission to see. If false, then only Protection Groups created by the current user will be returned. If it's not specified, it is true by default."
  type        = bool
  default     = false
}
variable "protection_run_progress_include_finished_tasks" {
  description = "Specifies whether to return finished tasks. By default only active tasks are returned."
  type        = bool
  default     = false
}
variable "protection_run_progress_start_time_usecs" {
  description = "Specifies the time after which the progress task starts in Unix epoch Timestamp(in microseconds)."
  type        = number
  default     = 0
}
variable "protection_run_progress_end_time_usecs" {
  description = "Specifies the time before which the progress task ends in Unix epoch Timestamp(in microseconds)."
  type        = number
  default     = 0
}
variable "protection_run_progress_max_tasks_num" {
  description = "Specifies the maximum number of tasks to return."
  type        = number
  default     = 0
}
variable "protection_run_progress_exclude_object_details" {
  description = "Specifies whether to return objects. By default all the task tree are returned."
  type        = bool
  default     = false
}
variable "protection_run_progress_include_event_logs" {
  description = "Specifies whether to include event logs."
  type        = bool
  default     = false
}
variable "protection_run_progress_max_log_level" {
  description = "Specifies the number of levels till which to fetch the event logs. This is applicable only when includeEventLogs is true."
  type        = number
  default     = 0
}
variable "protection_run_progress_run_task_path" {
  description = "Specifies the task path of the run or object run. This is applicable only if progress of a protection group with one or more object is required.If provided this will be used to fetch progress details directly without looking actual task path of the object. Objects field is stil expected else it changes the response format."
  type        = string
  default     = "placeholder"
}
variable "protection_run_progress_object_task_paths" {
  description = "Specifies the object level task path. This relates to the objectID. If provided this will take precedence over the objects, and will be used to fetch progress details directly without looking actuall task path of the object."
  type        = list(string)
  default     = [ "placeholder" ]
}

// Data source arguments for protection_run_stat
variable "protection_run_stat_run_id" {
  description = "Specifies a unique run id of the Protection Run."
  type        = string
  default     = "run_id"
}
variable "protection_run_stat_objects" {
  description = "Specifies the objects whose stats will be returned. This only applies to protection group runs and will be ignored for object runs. If the objects are specified, the run stats will not be returned and only the stats of the specified objects will be returned."
  type        = list(number)
  default     = [ 0 ]
}
variable "protection_run_stat_tenant_ids" {
  description = "TenantIds contains ids of the tenants for which the run is to be returned."
  type        = list(string)
  default     = [ "placeholder" ]
}
variable "protection_run_stat_include_tenants" {
  description = "If true, the response will include Protection Group Runs which were created by all tenants which the current user has permission to see. If false, then only Protection Groups created by the current user will be returned. If it's not specified, it is true by default."
  type        = bool
  default     = false
}
variable "protection_run_stat_include_finished_tasks" {
  description = "Specifies whether to return finished tasks. By default only active tasks are returned."
  type        = bool
  default     = false
}
variable "protection_run_stat_start_time_usecs" {
  description = "Specifies the time after which the stats task starts in Unix epoch Timestamp(in microseconds)."
  type        = number
  default     = 0
}
variable "protection_run_stat_end_time_usecs" {
  description = "Specifies the time before which the stats task ends in Unix epoch Timestamp(in microseconds)."
  type        = number
  default     = 0
}
variable "protection_run_stat_max_tasks_num" {
  description = "Specifies the maximum number of tasks to return."
  type        = number
  default     = 0
}
variable "protection_run_stat_exclude_object_details" {
  description = "Specifies whether to return objects. By default all the task tree are returned."
  type        = bool
  default     = false
}
variable "protection_run_stat_run_task_path" {
  description = "Specifies the task path of the run or object run. This is applicable only if stats of a protection group with one or more object is required. If provided this will be used to fetch stats details directly without looking actual task path of the object. Objects field is stil expected else it changes the response format."
  type        = string
  default     = "placeholder"
}
variable "protection_run_stat_object_task_paths" {
  description = "Specifies the object level task path. This relates to the objectID. If provided this will take precedence over the objects, and will be used to fetch stats details directly without looking actuall task path of the object."
  type        = list(string)
  default     = [ "placeholder" ]
}

// Data source arguments for search_objects
variable "search_objects_request_initiator_type" {
  description = "Specifies the type of request from UI, which is used for services like magneto to determine the priority of requests."
  type        = string
  default     = "placeholder"
}
variable "search_objects_search_string" {
  description = "Specifies the search string to filter the objects. This search string will be applicable for objectnames. User can specify a wildcard character '*' as a suffix to a string where all object names are matched with the prefix string. For example, if vm1 and vm2 are the names of objects, user can specify vm* to list the objects. If not specified, then all the objects will be returned which will match other filtering criteria."
  type        = string
  default     = "placeholder"
}
variable "search_objects_environments" {
  description = "Specifies the environment type to filter objects."
  type        = list(string)
  default     = [ "placeholder" ]
}
variable "search_objects_protection_types" {
  description = "Specifies the protection type to filter objects."
  type        = list(string)
  default     = [ "placeholder" ]
}
variable "search_objects_tenant_ids" {
  description = "TenantIds contains ids of the tenants for which objects are to be returned."
  type        = list(string)
  default     = [ "placeholder" ]
}
variable "search_objects_include_tenants" {
  description = "If true, the response will include Objects which belongs to all tenants which the current user has permission to see."
  type        = bool
  default     = false
}
variable "search_objects_protection_group_ids" {
  description = "Specifies a list of Protection Group ids to filter the objects. If specified, the objects protected by specified Protection Group ids will be returned."
  type        = list(string)
  default     = [ "placeholder" ]
}
variable "search_objects_object_ids" {
  description = "Specifies a list of Object ids to filter."
  type        = list(number)
  default     = [ 0 ]
}
variable "search_objects_os_types" {
  description = "Specifies the operating system types to filter objects on."
  type        = list(string)
  default     = [ "placeholder" ]
}
variable "search_objects_source_ids" {
  description = "Specifies a list of Protection Source object ids to filter the objects. If specified, the object which are present in those Sources will be returned."
  type        = list(number)
  default     = [ 0 ]
}
variable "search_objects_source_uuids" {
  description = "Specifies a list of Protection Source object uuids to filter the objects. If specified, the object which are present in those Sources will be returned."
  type        = list(string)
  default     = [ "placeholder" ]
}
variable "search_objects_is_protected" {
  description = "Specifies the protection status of objects. If set to true, only protected objects will be returned. If set to false, only unprotected objects will be returned. If not specified, all objects will be returned."
  type        = bool
  default     = false
}
variable "search_objects_is_deleted" {
  description = "If set to true, then objects which are deleted on atleast one cluster will be returned. If not set or set to false then objects which are registered on atleast one cluster are returned."
  type        = bool
  default     = false
}
variable "search_objects_last_run_status_list" {
  description = "Specifies a list of status of the object's last protection run. Only objects with last run status of these will be returned."
  type        = list(string)
  default     = [ "placeholder" ]
}
variable "search_objects_region_ids" {
  description = "Specifies a list of region ids. Only records from clusters having these region ids will be returned."
  type        = list(string)
  default     = [ "placeholder" ]
}
variable "search_objects_cluster_identifiers" {
  description = "Specifies the list of cluster identifiers. Format is clusterId:clusterIncarnationId. Only records from clusters having these identifiers will be returned."
  type        = list(string)
  default     = [ "placeholder" ]
}
variable "search_objects_storage_domain_ids" {
  description = "Specifies the list of storage domain ids. Format is clusterId:clusterIncarnationId:storageDomainId. Only objects having protection in these storage domains will be returned."
  type        = list(string)
  default     = [ "placeholder" ]
}
variable "search_objects_include_deleted_objects" {
  description = "Specifies whether to include deleted objects in response. These objects can't be protected but can be recovered. This field is deprecated."
  type        = bool
  default     = false
}
variable "search_objects_pagination_cookie" {
  description = "Specifies the pagination cookie with which subsequent parts of the response can be fetched."
  type        = string
  default     = "placeholder"
}
variable "search_objects_count" {
  description = "Specifies the number of objects to be fetched for the specified pagination cookie."
  type        = number
  default     = 0
}
variable "search_objects_must_have_tag_ids" {
  description = "Specifies tags which must be all present in the document."
  type        = list(string)
  default     = [ "placeholder" ]
}
variable "search_objects_might_have_tag_ids" {
  description = "Specifies list of tags, one or more of which might be present in the document. These are OR'ed together and the resulting criteria AND'ed with the rest of the query."
  type        = list(string)
  default     = [ "placeholder" ]
}
variable "search_objects_must_have_snapshot_tag_ids" {
  description = "Specifies snapshot tags which must be all present in the document."
  type        = list(string)
  default     = [ "placeholder" ]
}
variable "search_objects_might_have_snapshot_tag_ids" {
  description = "Specifies list of snapshot tags, one or more of which might be present in the document. These are OR'ed together and the resulting criteria AND'ed with the rest of the query."
  type        = list(string)
  default     = [ "placeholder" ]
}

// Data source arguments for search_protected_objects
variable "search_protected_objects_request_initiator_type" {
  description = "Specifies the type of request from UI, which is used for services like magneto to determine the priority of requests."
  type        = string
  default     = "placeholder"
}
variable "search_protected_objects_search_string" {
  description = "Specifies the search string to filter the objects. This search string will be applicable for objectnames and Protection Group names. User can specify a wildcard character '*' as a suffix to a string where all object and their Protection Group names are matched with the prefix string. For example, if vm1 and vm2 are the names of objects, user can specify vm* to list the objects. If not specified, then all the objects with Protection Groups will be returned which will match other filtering criteria."
  type        = string
  default     = "placeholder"
}
variable "search_protected_objects_environments" {
  description = "Specifies the environment type to filter objects."
  type        = list(string)
  default     = [ "placeholder" ]
}
variable "search_protected_objects_snapshot_actions" {
  description = "Specifies a list of recovery actions. Only snapshots that applies to these actions will be returned."
  type        = list(string)
  default     = [ "placeholder" ]
}
variable "search_protected_objects_object_action_key" {
  description = "Filter by ObjectActionKey, which uniquely represents protection of an object. An object can be protected in multiple ways but atmost once for a given combination of ObjectActionKey. When specified, latest snapshot info matching the objectActionKey is for corresponding object."
  type        = string
  default     = "placeholder"
}
variable "search_protected_objects_tenant_ids" {
  description = "TenantIds contains ids of the tenants for which objects are to be returned."
  type        = list(string)
  default     = [ "placeholder" ]
}
variable "search_protected_objects_include_tenants" {
  description = "If true, the response will include Objects which belongs to all tenants which the current user has permission to see."
  type        = bool
  default     = false
}
variable "search_protected_objects_protection_group_ids" {
  description = "Specifies a list of Protection Group ids to filter the objects. If specified, the objects protected by specified Protection Group ids will be returned."
  type        = list(string)
  default     = [ "placeholder" ]
}
variable "search_protected_objects_object_ids" {
  description = "Specifies a list of Object ids to filter."
  type        = list(number)
  default     = [ 0 ]
}
variable "search_protected_objects_storage_domain_ids" {
  description = "Specifies the Storage Domain ids to filter objects for which Protection Groups are writing data to Cohesity Views on the specified Storage Domains."
  type        = list(number)
  default     = [ 0 ]
}
variable "search_protected_objects_sub_result_size" {
  description = "Specifies the size of objects to be fetched for a single subresult."
  type        = number
  default     = 0
}
variable "search_protected_objects_filter_snapshot_from_usecs" {
  description = "Specifies the timestamp in Unix time epoch in microseconds to filter the objects if the Object has a successful snapshot after this value."
  type        = number
  default     = 0
}
variable "search_protected_objects_filter_snapshot_to_usecs" {
  description = "Specifies the timestamp in Unix time epoch in microseconds to filter the objects if the Object has a successful snapshot before this value."
  type        = number
  default     = 0
}
variable "search_protected_objects_os_types" {
  description = "Specifies the operating system types to filter objects on."
  type        = list(string)
  default     = [ "placeholder" ]
}
variable "search_protected_objects_source_ids" {
  description = "Specifies a list of Protection Source object ids to filter the objects. If specified, the object which are present in those Sources will be returned."
  type        = list(number)
  default     = [ 0 ]
}
variable "search_protected_objects_run_instance_ids" {
  description = "Specifies a list of run instance ids. If specified only objects belonging to the provided run id will be retunrned."
  type        = list(number)
  default     = [ 0 ]
}
variable "search_protected_objects_cdp_protected_only" {
  description = "Specifies whether to only return the CDP protected objects."
  type        = bool
  default     = false
}
variable "search_protected_objects_region_ids" {
  description = "Specifies a list of region ids. Only records from clusters having these region ids will be returned."
  type        = list(string)
  default     = [ "placeholder" ]
}
variable "search_protected_objects_use_cached_data" {
  description = "Specifies whether we can serve the GET request to the read replica cache cache. There is a lag of 15 seconds between the read replica and primary data source."
  type        = bool
  default     = false
}

// Data source arguments for protection_group
variable "data_protection_group_protection_group_id" {
  description = "Specifies a unique id of the Protection Group."
  type        = string
  default     = "protection_group_id"
}
variable "data_protection_group_request_initiator_type" {
  description = "Specifies the type of request from UI, which is used for services like magneto to determine the priority of requests."
  type        = string
  default     = "placeholder"
}
variable "data_protection_group_include_last_run_info" {
  description = "If true, the response will include last run info. If it is false or not specified, the last run info won't be returned."
  type        = bool
  default     = false
}
variable "data_protection_group_prune_excluded_source_ids" {
  description = "If true, the response will not include the list of excluded source IDs in groups that contain this field. This can be set to true in order to improve performance if excluded source IDs are not needed by the user."
  type        = bool
  default     = false
}
variable "data_protection_group_prune_source_ids" {
  description = "If true, the response will exclude the list of source IDs within the group specified."
  type        = bool
  default     = false
}

// Data source arguments for protection_groups
variable "protection_groups_request_initiator_type" {
  description = "Specifies the type of request from UI, which is used for services like magneto to determine the priority of requests."
  type        = string
  default     = "placeholder"
}
variable "protection_groups_ids" {
  description = "Filter by a list of Protection Group ids."
  type        = list(string)
  default     = [ "placeholder" ]
}
variable "protection_groups_names" {
  description = "Filter by a list of Protection Group names."
  type        = list(string)
  default     = [ "placeholder" ]
}
variable "protection_groups_policy_ids" {
  description = "Filter by Policy ids that are associated with Protection Groups. Only Protection Groups associated with the specified Policy ids, are returned."
  type        = list(string)
  default     = [ "placeholder" ]
}
variable "protection_groups_storage_domain_id" {
  description = "Filter by Storage Domain id. Only Protection Groups writing data to this Storage Domain will be returned."
  type        = number
  default     = 0
}
variable "protection_groups_include_groups_with_datalock_only" {
  description = "Whether to only return Protection Groups with a datalock."
  type        = bool
  default     = false
}
variable "protection_groups_environments" {
  description = "Filter by environment types such as 'kVMware', 'kView', etc. Only Protection Groups protecting the specified environment types are returned."
  type        = list(string)
  default     = [ "placeholder" ]
}
variable "protection_groups_is_active" {
  description = "Filter by Inactive or Active Protection Groups. If not set, all Inactive and Active Protection Groups are returned. If true, only Active Protection Groups are returned. If false, only Inactive Protection Groups are returned. When you create a Protection Group on a Primary Cluster with a replication schedule, the Cluster creates an Inactive copy of the Protection Group on the Remote Cluster. In addition, when an Active and running Protection Group is deactivated, the Protection Group becomes Inactive."
  type        = bool
  default     = false
}
variable "protection_groups_is_deleted" {
  description = "If true, return only Protection Groups that have been deleted but still have Snapshots associated with them. If false, return all Protection Groups except those Protection Groups that have been deleted and still have Snapshots associated with them. A Protection Group that is deleted with all its Snapshots is not returned for either of these cases."
  type        = bool
  default     = false
}
variable "protection_groups_is_paused" {
  description = "Filter by paused or non paused Protection Groups, If not set, all paused and non paused Protection Groups are returned. If true, only paused Protection Groups are returned. If false, only non paused Protection Groups are returned."
  type        = bool
  default     = false
}
variable "protection_groups_last_run_local_backup_status" {
  description = "Filter by last local backup run status.<br> 'Running' indicates that the run is still running.<br> 'Canceled' indicates that the run has been canceled.<br> 'Canceling' indicates that the run is in the process of being canceled.<br> 'Failed' indicates that the run has failed.<br> 'Missed' indicates that the run was unable to take place at the scheduled time because the previous run was still happening.<br> 'Succeeded' indicates that the run has finished successfully.<br> 'SucceededWithWarning' indicates that the run finished successfully, but there were some warning messages.<br> 'Paused' indicates that the ongoing run has been paused.<br> 'Skipped' indicates that the run was skipped."
  type        = list(string)
  default     = [ "placeholder" ]
}
variable "protection_groups_last_run_replication_status" {
  description = "Filter by last remote replication run status.<br> 'Running' indicates that the run is still running.<br> 'Canceled' indicates that the run has been canceled.<br> 'Canceling' indicates that the run is in the process of being canceled.<br> 'Failed' indicates that the run has failed.<br> 'Missed' indicates that the run was unable to take place at the scheduled time because the previous run was still happening.<br> 'Succeeded' indicates that the run has finished successfully.<br> 'SucceededWithWarning' indicates that the run finished successfully, but there were some warning messages.<br> 'Paused' indicates that the ongoing run has been paused.<br> 'Skipped' indicates that the run was skipped."
  type        = list(string)
  default     = [ "placeholder" ]
}
variable "protection_groups_last_run_archival_status" {
  description = "Filter by last cloud archival run status.<br> 'Running' indicates that the run is still running.<br> 'Canceled' indicates that the run has been canceled.<br> 'Canceling' indicates that the run is in the process of being canceled.<br> 'Failed' indicates that the run has failed.<br> 'Missed' indicates that the run was unable to take place at the scheduled time because the previous run was still happening.<br> 'Succeeded' indicates that the run has finished successfully.<br> 'SucceededWithWarning' indicates that the run finished successfully, but there were some warning messages.<br> 'Paused' indicates that the ongoing run has been paused.<br> 'Skipped' indicates that the run was skipped."
  type        = list(string)
  default     = [ "placeholder" ]
}
variable "protection_groups_last_run_cloud_spin_status" {
  description = "Filter by last cloud spin run status.<br> 'Running' indicates that the run is still running.<br> 'Canceled' indicates that the run has been canceled.<br> 'Canceling' indicates that the run is in the process of being canceled.<br> 'Failed' indicates that the run has failed.<br> 'Missed' indicates that the run was unable to take place at the scheduled time because the previous run was still happening.<br> 'Succeeded' indicates that the run has finished successfully.<br> 'SucceededWithWarning' indicates that the run finished successfully, but there were some warning messages.<br> 'Paused' indicates that the ongoing run has been paused.<br> 'Skipped' indicates that the run was skipped."
  type        = list(string)
  default     = [ "placeholder" ]
}
variable "protection_groups_last_run_any_status" {
  description = "Filter by last any run status.<br> 'Running' indicates that the run is still running.<br> 'Canceled' indicates that the run has been canceled.<br> 'Canceling' indicates that the run is in the process of being canceled.<br> 'Failed' indicates that the run has failed.<br> 'Missed' indicates that the run was unable to take place at the scheduled time because the previous run was still happening.<br> 'Succeeded' indicates that the run has finished successfully.<br> 'SucceededWithWarning' indicates that the run finished successfully, but there were some warning messages.<br> 'Paused' indicates that the ongoing run has been paused.<br> 'Skipped' indicates that the run was skipped."
  type        = list(string)
  default     = [ "placeholder" ]
}
variable "protection_groups_is_last_run_sla_violated" {
  description = "If true, return Protection Groups for which last run SLA was violated."
  type        = bool
  default     = false
}
variable "protection_groups_tenant_ids" {
  description = "TenantIds contains ids of the tenants for which objects are to be returned."
  type        = list(string)
  default     = [ "placeholder" ]
}
variable "protection_groups_include_tenants" {
  description = "If true, the response will include Protection Groups which were created by all tenants which the current user has permission to see. If false, then only Protection Groups created by the current user will be returned."
  type        = bool
  default     = false
}
variable "protection_groups_include_last_run_info" {
  description = "If true, the response will include last run info. If it is false or not specified, the last run info won't be returned."
  type        = bool
  default     = false
}
variable "protection_groups_prune_excluded_source_ids" {
  description = "If true, the response will not include the list of excluded source IDs in groups that contain this field. This can be set to true in order to improve performance if excluded source IDs are not needed by the user."
  type        = bool
  default     = false
}
variable "protection_groups_prune_source_ids" {
  description = "If true, the response will exclude the list of source IDs within the group specified."
  type        = bool
  default     = false
}
variable "protection_groups_use_cached_data" {
  description = "Specifies whether we can serve the GET request from the read replica cache. There is a lag of 15 seconds between the read replica and primary data source."
  type        = bool
  default     = false
}

// Data source arguments for protection_group_run
variable "protection_group_run_protection_group_run_id" {
  description = "Specifies a unique id of the Protection Group."
  type        = string
  default     = "protection_group_run_id"
}
variable "protection_group_run_run_id" {
  description = "Specifies a unique run id of the Protection Group run."
  type        = string
  default     = "run_id"
}
variable "protection_group_run_request_initiator_type" {
  description = "Specifies the type of request from UI, which is used for services like magneto to determine the priority of requests."
  type        = string
  default     = "placeholder"
}
variable "protection_group_run_tenant_ids" {
  description = "TenantIds contains ids of the tenants for which the run is to be returned."
  type        = list(string)
  default     = [ "placeholder" ]
}
variable "protection_group_run_include_tenants" {
  description = "If true, the response will include Protection Group Runs which were created by all tenants which the current user has permission to see. If false, then only Protection Groups created by the current user will be returned. If it's not specified, it is true by default."
  type        = bool
  default     = false
}
variable "protection_group_run_include_object_details" {
  description = "Specifies if the result includes the object details for a protection run. If set to true, details of the protected object will be returned. If set to false or not specified, details will not be returned."
  type        = bool
  default     = false
}
variable "protection_group_run_use_cached_data" {
  description = "Specifies whether we can serve the GET request from the read replica cache. There is a lag of 15 seconds between the read replica and primary data source."
  type        = bool
  default     = false
}

// Data source arguments for protection_group_runs
variable "protection_group_runs_protection_group_runs_id" {
  description = "Specifies a unique id of the Protection Group."
  type        = string
  default     = "protection_group_runs_id"
}
variable "protection_group_runs_request_initiator_type" {
  description = "Specifies the type of request from UI, which is used for services like magneto to determine the priority of requests."
  type        = string
  default     = "placeholder"
}
variable "protection_group_runs_run_id" {
  description = "Specifies the protection run id."
  type        = string
  default     = "placeholder"
}
variable "protection_group_runs_start_time_usecs" {
  description = "Start time for time range filter. Specify the start time as a Unix epoch Timestamp (in microseconds), only runs executing after this time will be returned. By default it is endTimeUsecs minus an hour."
  type        = number
  default     = 0
}
variable "protection_group_runs_end_time_usecs" {
  description = "End time for time range filter. Specify the end time as a Unix epoch Timestamp (in microseconds), only runs executing before this time will be returned. By default it is current time."
  type        = number
  default     = 0
}
variable "protection_group_runs_tenant_ids" {
  description = "TenantIds contains ids of the tenants for which objects are to be returned."
  type        = list(string)
  default     = [ "placeholder" ]
}
variable "protection_group_runs_include_tenants" {
  description = "If true, the response will include Protection Group Runs which were created by all tenants which the current user has permission to see. If false, then only Protection Group Runs created by the current user will be returned."
  type        = bool
  default     = false
}
variable "protection_group_runs_run_types" {
  description = "Filter by run type. Only protection run matching the specified types will be returned."
  type        = list(string)
  default     = [ "placeholder" ]
}
variable "protection_group_runs_include_object_details" {
  description = "Specifies if the result includes the object details for each protection run. If set to true, details of the protected object will be returned. If set to false or not specified, details will not be returned."
  type        = bool
  default     = false
}
variable "protection_group_runs_local_backup_run_status" {
  description = "Specifies a list of local backup status, runs matching the status will be returned.<br> 'Running' indicates that the run is still running.<br> 'Canceled' indicates that the run has been canceled.<br> 'Canceling' indicates that the run is in the process of being canceled.<br> 'Failed' indicates that the run has failed.<br> 'Missed' indicates that the run was unable to take place at the scheduled time because the previous run was still happening.<br> 'Succeeded' indicates that the run has finished successfully.<br> 'SucceededWithWarning' indicates that the run finished successfully, but there were some warning messages.<br> 'Paused' indicates that the ongoing run has been paused.<br> 'Skipped' indicates that the run was skipped."
  type        = list(string)
  default     = [ "placeholder" ]
}
variable "protection_group_runs_replication_run_status" {
  description = "Specifies a list of replication status, runs matching the status will be returned.<br> 'Running' indicates that the run is still running.<br> 'Canceled' indicates that the run has been canceled.<br> 'Canceling' indicates that the run is in the process of being canceled.<br> 'Failed' indicates that the run has failed.<br> 'Missed' indicates that the run was unable to take place at the scheduled time because the previous run was still happening.<br> 'Succeeded' indicates that the run has finished successfully.<br> 'SucceededWithWarning' indicates that the run finished successfully, but there were some warning messages.<br> 'Paused' indicates that the ongoing run has been paused.<br> 'Skipped' indicates that the run was skipped."
  type        = list(string)
  default     = [ "placeholder" ]
}
variable "protection_group_runs_archival_run_status" {
  description = "Specifies a list of archival status, runs matching the status will be returned.<br> 'Running' indicates that the run is still running.<br> 'Canceled' indicates that the run has been canceled.<br> 'Canceling' indicates that the run is in the process of being canceled.<br> 'Failed' indicates that the run has failed.<br> 'Missed' indicates that the run was unable to take place at the scheduled time because the previous run was still happening.<br> 'Succeeded' indicates that the run has finished successfully.<br> 'SucceededWithWarning' indicates that the run finished successfully, but there were some warning messages.<br> 'Paused' indicates that the ongoing run has been paused.<br> 'Skipped' indicates that the run was skipped."
  type        = list(string)
  default     = [ "placeholder" ]
}
variable "protection_group_runs_cloud_spin_run_status" {
  description = "Specifies a list of cloud spin status, runs matching the status will be returned.<br> 'Running' indicates that the run is still running.<br> 'Canceled' indicates that the run has been canceled.<br> 'Canceling' indicates that the run is in the process of being canceled.<br> 'Failed' indicates that the run has failed.<br> 'Missed' indicates that the run was unable to take place at the scheduled time because the previous run was still happening.<br> 'Succeeded' indicates that the run has finished successfully.<br> 'SucceededWithWarning' indicates that the run finished successfully, but there were some warning messages.<br> 'Paused' indicates that the ongoing run has been paused.<br> 'Skipped' indicates that the run was skipped."
  type        = list(string)
  default     = [ "placeholder" ]
}
variable "protection_group_runs_num_runs" {
  description = "Specifies the max number of runs. If not specified, at most 100 runs will be returned."
  type        = number
  default     = 0
}
variable "protection_group_runs_exclude_non_restorable_runs" {
  description = "Specifies whether to exclude non restorable runs. Run is treated restorable only if there is atleast one object snapshot (which may be either a local or an archival snapshot) which is not deleted or expired. Default value is false."
  type        = bool
  default     = false
}
variable "protection_group_runs_run_tags" {
  description = "Specifies a list of tags for protection runs. If this is specified, only the runs which match these tags will be returned."
  type        = list(string)
  default     = [ "placeholder" ]
}
variable "protection_group_runs_use_cached_data" {
  description = "Specifies whether we can serve the GET request from the read replica cache. There is a lag of 15 seconds between the read replica and primary data source."
  type        = bool
  default     = false
}
variable "protection_group_runs_filter_by_end_time" {
  description = "If true, the runs with backup end time within the specified time range will be returned. Otherwise, the runs with start time in the time range are returned."
  type        = bool
  default     = false
}
variable "protection_group_runs_snapshot_target_types" {
  description = "Specifies the snapshot's target type which should be filtered."
  type        = list(string)
  default     = [ "placeholder" ]
}
variable "protection_group_runs_only_return_successful_copy_run" {
  description = "only successful copyruns are returned."
  type        = bool
  default     = false
}
variable "protection_group_runs_filter_by_copy_task_end_time" {
  description = "If true, then the details of the runs for which any copyTask completed in the given timerange will be returned. Only one of filterByEndTime and filterByCopyTaskEndTime can be set."
  type        = bool
  default     = false
}

// Data source arguments for protection_policies
variable "protection_policies_request_initiator_type" {
  description = "Specifies the type of request from UI, which is used for services like magneto to determine the priority of requests."
  type        = string
  default     = "placeholder"
}
variable "protection_policies_ids" {
  description = "Filter policies by a list of policy ids."
  type        = list(string)
  default     = [ "placeholder" ]
}
variable "protection_policies_policy_names" {
  description = "Filter policies by a list of policy names."
  type        = list(string)
  default     = [ "placeholder" ]
}
variable "protection_policies_tenant_ids" {
  description = "TenantIds contains ids of the organizations for which objects are to be returned."
  type        = list(string)
  default     = [ "placeholder" ]
}
variable "protection_policies_include_tenants" {
  description = "IncludeTenantPolicies specifies if objects of all the organizations under the hierarchy of the logged in user's organization should be returned."
  type        = bool
  default     = false
}
variable "protection_policies_types" {
  description = "Types specifies the policy type of policies to be returned."
  type        = list(string)
  default     = [ "placeholder" ]
}
variable "protection_policies_exclude_linked_policies" {
  description = "If excludeLinkedPolicies is set to true then only local policies created on cluster will be returned. The result will exclude all linked policies created from policy templates."
  type        = bool
  default     = false
}
variable "protection_policies_include_replicated_policies" {
  description = "If includeReplicatedPolicies is set to true, then response will also contain replicated policies. By default, replication policies are not included in the response."
  type        = bool
  default     = false
}
variable "protection_policies_include_stats" {
  description = "If includeStats is set to true, then response will return number of protection groups and objects. By default, the protection stats are not included in the response."
  type        = bool
  default     = false
}

// Data source arguments for protection_policy
variable "data_protection_policy_protection_policy_id" {
  description = "Specifies a unique id of the Protection Policy to return."
  type        = string
  default     = "protection_policy_id"
}
variable "data_protection_policy_request_initiator_type" {
  description = "Specifies the type of request from UI, which is used for services like magneto to determine the priority of requests."
  type        = string
  default     = "placeholder"
}

// Data source arguments for protection_run_summary
variable "protection_run_summary_start_time_usecs" {
  description = "Start time for time range filter. Specify the start time as a Unix epoch Timestamp (in microseconds), only runs executing after this time will be returned. By default it is endTimeUsecs minus an hour."
  type        = number
  default     = 0
}
variable "protection_run_summary_end_time_usecs" {
  description = "End time for time range filter. Specify the end time as a Unix epoch Timestamp (in microseconds), only runs executing before this time will be returned. By default it is current time."
  type        = number
  default     = 0
}
variable "protection_run_summary_run_status" {
  description = "Specifies a list of status, runs matching the status will be returned.<br> 'Running' indicates that the run is still running.<br> 'Canceled' indicates that the run has been canceled.<br> 'Canceling' indicates that the run is in the process of being canceled.<br> 'Failed' indicates that the run has failed.<br> 'Missed' indicates that the run was unable to take place at the scheduled time because the previous run was still happening.<br> 'Succeeded' indicates that the run has finished successfully.<br> 'SucceededWithWarning' indicates that the run finished successfully, but there were some warning messages.<br> 'Skipped' indicates that the run was skipped."
  type        = list(string)
  default     = [ "placeholder" ]
}

// Data source arguments for protection_sources
variable "protection_sources_request_initiator_type" {
  description = "Specifies the type of request from UI, which is used for services like magneto to determine the priority of requests."
  type        = string
  default     = "placeholder"
}
variable "protection_sources_tenant_ids" {
  description = "TenantIds contains ids of the tenants for which Sources are to be returned."
  type        = list(string)
  default     = [ "placeholder" ]
}
variable "protection_sources_include_tenants" {
  description = "If true, the response will include Sources which belong belong to all tenants which the current user has permission to see. If false, then only Sources for the current user will be returned."
  type        = bool
  default     = false
}
variable "protection_sources_include_source_credentials" {
  description = "If true, the encrypted crednetial for the registered sources will be included. Credential is first encrypted with internal key and then reencrypted with user supplied encryption key."
  type        = bool
  default     = false
}
variable "protection_sources_encryption_key" {
  description = "Specifies the key to be used to encrypt the source credential. If includeSourceCredentials is set to true this key must be specified."
  type        = string
  default     = "placeholder"
}

// Data source arguments for recovery
variable "data_recovery_recovery_id" {
  description = "Specifies the id of a Recovery."
  type        = string
  default     = "recovery_id"
}
variable "data_recovery_include_tenants" {
  description = "Specifies if objects of all the organizations under the hierarchy of the logged in user's organization should be returned."
  type        = bool
  default     = false
}

// Data source arguments for recoveries
variable "recoveries_ids" {
  description = "Filter Recoveries for given ids."
  type        = list(string)
  default     = [ "placeholder" ]
}
variable "recoveries_return_only_child_recoveries" {
  description = "Returns only child recoveries if passed as true. This filter should always be used along with 'ids' filter."
  type        = bool
  default     = false
}
variable "recoveries_tenant_ids" {
  description = "TenantIds contains ids of the organizations for which recoveries are to be returned."
  type        = list(string)
  default     = [ "placeholder" ]
}
variable "recoveries_include_tenants" {
  description = "Specifies if objects of all the organizations under the hierarchy of the logged in user's organization should be returned."
  type        = bool
  default     = false
}
variable "recoveries_start_time_usecs" {
  description = "Returns the recoveries which are started after the specific time. This value should be in Unix timestamp epoch in microseconds."
  type        = number
  default     = 0
}
variable "recoveries_end_time_usecs" {
  description = "Returns the recoveries which are started before the specific time. This value should be in Unix timestamp epoch in microseconds."
  type        = number
  default     = 0
}
variable "recoveries_storage_domain_id" {
  description = "Filter by Storage Domain id. Only recoveries writing data to this Storage Domain will be returned."
  type        = number
  default     = 0
}
variable "recoveries_snapshot_target_type" {
  description = "Specifies the snapshot's target type from which recovery has been performed."
  type        = list(string)
  default     = [ "placeholder" ]
}
variable "recoveries_archival_target_type" {
  description = "Specifies the snapshot's archival target type from which recovery has been performed. This parameter applies only if 'snapshotTargetType' is 'Archival'."
  type        = list(string)
  default     = [ "placeholder" ]
}
variable "recoveries_snapshot_environments" {
  description = "Specifies the list of snapshot environment types to filter Recoveries. If empty, Recoveries related to all environments will be returned."
  type        = list(string)
  default     = [ "placeholder" ]
}
variable "recoveries_status" {
  description = "Specifies the list of run status to filter Recoveries. If empty, Recoveries with all run status will be returned."
  type        = list(string)
  default     = [ "placeholder" ]
}
variable "recoveries_recovery_actions" {
  description = "Specifies the list of recovery actions to filter Recoveries. If empty, Recoveries related to all actions will be returned."
  type        = list(string)
  default     = [ "placeholder" ]
}

// Data source arguments for source_registrations
variable "source_registrations_ids" {
  description = "Ids specifies the list of source registration ids to return. If left empty, every source registration will be returned by default."
  type        = list(number)
  default     = [ 0 ]
}
variable "source_registrations_tenant_ids" {
  description = "TenantIds contains ids of the tenants for which objects are to be returned."
  type        = list(string)
  default     = [ "placeholder" ]
}
variable "source_registrations_include_tenants" {
  description = "If true, the response will include Registrations which were created by all tenants which the current user has permission to see. If false, then only Registrations created by the current user will be returned."
  type        = bool
  default     = false
}
variable "source_registrations_include_source_credentials" {
  description = "If true, the encrypted crednetial for the registered sources will be included. Credential is first encrypted with internal key and then reencrypted with user supplied encryption key."
  type        = bool
  default     = false
}
variable "source_registrations_encryption_key" {
  description = "Specifies the key to be used to encrypt the source credential. If includeSourceCredentials is set to true this key must be specified."
  type        = string
  default     = "placeholder"
}
variable "source_registrations_use_cached_data" {
  description = "Specifies whether we can serve the GET request from the read replica cache. There is a lag of 15 seconds between the read replica and primary data source."
  type        = bool
  default     = false
}

// Data source arguments for source_registration
variable "data_source_registration_source_registration_id" {
  description = "Specifies the id of the Protection Source registration."
  type        = number
  default     = 2
}
variable "data_source_registration_request_initiator_type" {
  description = "Specifies the type of request from UI, which is used for services like magneto to determine the priority of requests."
  type        = string
  default     = "placeholder"
}
