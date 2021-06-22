---
subcategory: "Schematics"
layout: "ibm"
page_title: "IBM : ibm_schematics_job"
sidebar_current: "docs-ibm-datasource-schematics-job"
description: |-
  Get information about Schematics job.
---

# ibm_schematics_job
Retrieve information about a Schematics job. For more details about the Schematics and Schematics job, see [setting up jobs](https://cloud.ibm.com/docs/schematics?topic=schematics-action-setup#action-jobs).

## Example usage

```terraform
data "ibm_schematics_job" "schematics_job" {
	job_id = "job_id"
}
```
## Argument reference
Review the argument references that you can specify for your data source. 

- `job_id` - (Required, String) Use GET or jobs API to see the job IDs in your IBM Cloud account.

## Attribute reference
In addition to the argument reference list, you can access the following attribute references after your data source is created. 

- `bastion` - (String) The complete target details with the user inputs and the system generated data. Nested bastion blocks have the following structure.

  Nested scheme for `bastion`:
  - `credential` - (String) Override credential for each resource. Reference to credentials values, used by all the resources.
  - `created_at` - (String) The targets creation time.
  - `created_by` - (String) The Email address of the user who created the targets.
  - `description` - (String) The target description.
  - `id` - (String) The target ID.
  - `name` - (String) The target name.
  - `resource_query` - (String) The resource selection query string.
 `resource_ids` - (String) Array of the resource IDs.
  - `sys_lock` - (String) The system lock status.Nested sys_lock blocks have the following structure.
  - `sys_lock.sys_locked` - (String) Is the workspace locked by the Schematics action?-
  - `sys_lock.sys_locked_by` - (String) The name of the user who performed the action, that lead to lock the workspace.
  - `sys_lock.sys_locked_at` - (String) The action that lead to lock the workspace?
   - `type` - (String) The target type such as, `cluster`, `vsi`, `icd`, `vpc`.
  - `updated_at` - (String) The targets update time.
  - `updated_by` - (String) The Email address of user who updated the targets.
- `command_object` - (String) The name of the Schematics automation resource.
- `command_object_id` - (String) The job command object ID such as `workspace-id`, `action-id`, or `control-id`.
- `command_name` - (String) The Schematics job command name.
- `command_parameter` - (String) The Schematics job command parameter  such as `playbook-name`, `capsule-name`, or `flow-name`.
- `command_options` - (String) The command line options for the command.
- `data` (String) The Job data. Nested data blocks have the following structure.

  Nested scheme for `data`:
  - `job_type` - (String) Type of the job.
  - `action_job_data`-  (String) The action job data. Nested `action_job_data` blocks have the following structure.

    Nested scheme for `action_job_data`:
    - `action_name` - (String) The flow name.
	- `inputs` - (String) The input variables for an action job. Nested `inputs` blocks have the following structure.

	  Nested scheme for `inputs`:
	  - `metadata` - (String) User editable metadata for the variables. Nested `metadata` blocks have the following structure.

	    Nested scheme for `metadata`:
	    - `type` - (String) The type of the variable.
	    - `aliases` - (String) The list of an aliases for the variable name.
	    - `description` - (String) The description of the metadata.
	    - `default_value` - (String) The default value for the variable, if the override value is not specified.
	    - `secure` - (String) Is the variable secure or sensitive?-
	    - `immutable` - (String) Is the variable read only?-
	    - `hidden` - (String) If set to **true**, the variable will not be displayed on console or command-line.
	    - `options` - (String) The list of possible values for this variable. If type is integer or date, then the array of string will be converted to array of integers or date during runtime.
	    - `min_value` - (String) Minimum value of the variable. Applicable for integer type.
	    - `max_value` - (String) Maximum value of the variable. Applicable for integer type.
	    - `min_length` - (String) Minimum length of the variable value. Applicable for string type.
	    - `max_length` - (String) Maximum length of the variable value. Applicable for string type.
	    - `matches` - (String) Regular expression for the variable value.
	    - `position` - (String) Relative position of this variable in a list.
	    - `group_by` - (String) The display name of the group this variable belongs to.
	    - `source` - (String) The source of this metadata.
  - `link` - (String) Reference link to the variable value By default the expression will point to `self.value`.
  - `name` - (String) The name of the variable.
  - `value` - (String) The value for the variable or reference to the value.
- `description` - (String) The job description derived from the related action.
- `duration` - (String) Duration of job execution, for example, `40 seconds`.
- `end_at` -  (String) The job end time.
- `id` - (String) The unique ID of the Schematics job.
- `id` - (String) The job ID.
- `location` - (String) List of an action locations supported by Schematics service. Note this does not limit the location of the resources provisioned using Schematics.
- `log_store_url` - (String) The job log store URL.
- `name` -  (String) The unique job name, derived from the related action.
- `outputs` - (String) The output variables for an action. Nested `outputs` blocks have the following structure.

  Nested scheme for `output`:
  - `metadata` - (String) User editable metadata for the variables. Nested metadata blocks have the following structure.

	Nested scheme for `metadata`:
	- `type` - (String) Type of the variable.
	- `aliases` - (String) List of aliases for the variable name.
	- `description` - (String) Description of the meta data.
	- `default_value` - (String) Default value for the variable, if the override value is not specified.
	- `secure` - (String) Is the variable secure or sensitive?-
	- `immutable` - (String) Is the variable read only ?-
	- `hidden` - (String) If set **true**, the variable will not be displayed on console or command-line.
	- `options` - (String) List of possible values for this variable. If type is integer or date, then the array of string will be converted to array of integers or date during runtime.
	- `min_value` - (String) Minimum value of the variable. Applicable for integer type.
	- `max_value` - (String) Maximum value of the variable. Applicable for integer type.
	- `min_length` - (String) Minimum length of the variable value. Applicable for string type.
	- `max_length` - (String) Maximum length of the variable value. Applicable for string type.
	- `matches` - (String) Regex for the variable value.
	- `position` - (String) Relative position of this variable in a list.
	- `group_by` - (String) Display name of the group this variable belongs to.
	- `source` - (String) Source of this meta-data.
  - `link` - (String) Reference link to the variable value By default the expression will point to `self.value`.
  - `name` - (String) Name of the variable.
  - `value` - (String) Value for the variable or reference to the value.
- `resource_group` - (String) The resource group name derived from the related action.
- `settings` - (String) Environment variables for an action. Nested settings blocks have the following structure.

  Nested scheme for `id`:
  - `metadata` - (String) User editable metadata for the variables. Nested metadata blocks have the following structure.

    Nested scheme for `metadata`:
    - `type` - (String) Type of the variable.
    - `aliases` - (String) List of aliases for the variable name.
    - `description` - (String) Description of the meta data.
    - `default_value` - (String) Default value for the variable, if the override value is not specified.
	- `secure` - (String) Is the variable secure or sensitive?-
	- `immutable` - (String) Is the variable read only?
	- `hidden` - (String) If set **true**, the variable will not be displayed on console or command-line.
	- `options` - (String) List of possible values for this variable. If type is integer or date, then the array of string will be converted to array of integers or date during runtime.
	- `min_value` - (String) Minimum value of the variable. Applicable for integer type.
	- `max_value` - (String) Maximum value of the variable. Applicable for integer type.
	- `min_length` - (String) Minimum length of the variable value. Applicable for string type.
	- `max_length` - (String) Maximum length of the variable value. Applicable for string type.
	- `matches` - (String) Regex for the variable value.
	- `position` - (String) Relative position of this variable in a list.
	- `group_by` - (String) Display name of the group this variable belongs to.
	- `source` - (String) Source of this meta-data.
  - `link` - (String) Reference link to the variable value By default the expression will point to `self.value`.
  - `name` - (String) Name of the variable.
  - `value` - (String) Value for the variable or reference to the value.
- `submitted_at` -  (String) The job submission time.
- `submitted_by` - (String) The Email address of the user who submitted the job.
- `start_at` - (String) The job start time.
- `status` - (String) The job status. Nested `status` blocks have the following structure.
 
  Nested scheme for `status`:
  - `action_job_status` -  (String) The action job status. Nested `action_job_status` blocks have the following structure.

    Nested scheme for `action_job_status`:
	- `action_name` - (String) The action name.
	- `bastion_status_code` -  (String) The status of the resources.
	- `bastion_status_message` -  (String) The bastion status message to be displayed along with the `bastion_status_code`.
	- `status_code` -  (String) The status of the jobs.
	- `status_message` -  (String) An action job status message to be displayed along with the `action_status_code`.
	- `targets_status_code` -  (String) -Status of the resources.
	- `targets_status_message` -  (String) -Aggregated status message for all target resources, to be displayed along with the targets_status_code.
	- `updated_at` -  (String) The job status update timestamp.
- `settings` - (String) The environment variables used by all the templates in an action. Nested `settings` blocks have the following structure.

  Nested scheme for `settings`:
  - `name` - (String) Name of the variable.
  - `value` - (String) Value for the variable or reference to the value.
  - `metadata` - (String) User editable metadata for the variables. Nested metadata blocks have the following structure.

    Nested scheme for `metadata`:
	- `type` - (String) Type of the variable.
	- `aliases` - (String) List of aliases for the variable name.
	- `description` - (String) Description of the meta data.
	- `default_value` - (String) Default value for the variable, if the override value is not specified.
	- `secure` - (String) Is the variable secure or sensitive?-
	- `immutable` - (String) Is the variable read only ?-
	- `hidden` - (String) If set **true**, the variable will not be displayed on console or command-line.
	- `options` - (String) List of possible values for this variable. If type is integer or date, then the array of string will be converted to array of integers or date during runtime.
	- `min_value` - (String) Minimum value of the variable. Applicable for integer type.
	- `max_value` - (String) Maximum value of the variable. Applicable for integer type.
	- `min_length` - (String) Minimum length of the variable value. Applicable for string type.
	- `max_length` - (String) Maximum length of the variable value. Applicable for string type.
	- `matches` - (String) Regex for the variable value.
	- `position` - (String) Relative position of this variable in a list.
	- `group_by` - (String) Display name of the group this variable belongs to.
	- `source` - (String) Source of this meta-data.
  - `link` - (String) Reference link to the variable value By default the expression will point to `self.value`.
- `log_summary` -  (String) -Job log summary record. Nested `log_summary` blocks have the following structure.

  Nested scheme for `log_summary`:
  - `elapsed_time` -  (String) The job log elapsed time `log_analyzed_till`, `log_start_at`.
  - `job_id` - (String) The workspace ID.
  - `job_type` -  (String) The type of job.
  - `log_start_at` -  (String) The job log start timestamp.
  - `log_analyzed_till` - (String) The job log update timestamp.
  - `log_errors` -  (String) -Job log errors. Nested log_errors blocks have the following structure.

    Nested scheme for `log_errors`:
    - `error_code` -  (String) The error code in the Log.
    - `error_msg` -  (String) The summary error message in the log.
    - `error_count` - (String) The number of occurrence.
  - `repo_download_job` -  (String) -Repo download Job log summary. Nested `repo_download_job `blocks have the following structure.
    
	 Nested scheme for `repo_download_job`:
     - `detected_filetype` -  (String) Detected template or data file type.
	 - `quarantined_file_count` - (String) The number of files quarantined.
     - `inputs_count` - (String) The number of inputs detected.
     - `outputs_count` - (String) The number of outputs detected.
	 - `scanned_file_count` - (String) The number of files scanned.
  - `action_job` -  (String) The flow job log summary. Nested `action_job` blocks have the following structure.

    Nested scheme for `action_job`:
    - `play_count` -  (String) -number of plays in playbook.
	- `target_count` - (String) The number of targets or hosts. 
    - `task_count` -  (String) The number of tasks in playbook.
    - `recap` -  (String) The recap records. Nested recap blocks have the following structure.

	  Nested scheme for `recap`:
	  - `changed` - (String) The number of changed.
	  - `failed` - (String) The number of failed.
	  - `ok` -  (String) -the number of OK.
	  - `skipped` -  (String) The number of skipped.
	  - `target` -  (String) The list of target or host name.
	  - `unreachable` -  (String) The number of unreachable.
- `results_url` -  (String) The job results store URL.
- `state_store_url` -  (String) The job state store URL.
- `targets_ini` - (String) Inventory of host and host group for the playbook in `INI` file format. For example, `"targets_ini": "[webserverhost] 172.22.192.6 [dbhost] 172.22.192.5"`. For more information, about an inventory host group syntax, see [Inventory host groups](https://cloud.ibm.com/docs/schematics?topic=schematics-schematics-cli-reference#inventory-host-grps).
- `tags` -  (String) The user defined tags, while running the job.
- `updated_at` - (Timestamp) The job status update timestamp.

