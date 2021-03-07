---
layout: "ibm"
page_title: "IBM : schematics_job"
sidebar_current: "docs-ibm-datasource-schematics-job"
description: |-
  Get information about schematics_job
---

# ibm\_schematics_job

Provides a read-only data source for schematics_job. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "schematics_job" "schematics_job" {
	job_id = "job_id"
}
```

## Argument Reference

The following arguments are supported:

* `job_id` - (Required, string) Use GET jobs API to look up the Job IDs in your IBM Cloud account.

## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier of the schematics_job.
* `command_object` - Name of the Schematics automation resource.

* `command_object_id` - Job command object ID (`workspace-id, action-id or control-id`).

* `command_name` - Schematics job command name.

* `command_parameter` - Schematics job command parameter (`playbook-name, capsule-name or flow-name`).

* `command_options` - Command line options for the command.

* `inputs` - Job inputs used by an action. Nested `inputs` blocks have the following structure:
	* `name` - Name of the variable.
	* `value` - Value for the variable or reference to the value.
	* `metadata` - User editable metadata for the variables. Nested `metadata` blocks have the following structure:
		* `type` - Type of the variable.
		* `aliases` - List of aliases for the variable name.
		* `description` - Description of the meta data.
		* `default_value` - Default value for the variable, if the override value is not specified.
		* `secure` - Is the variable secure or sensitive ?.
		* `immutable` - Is the variable readonly ?.
		* `hidden` - If true, the variable will not be displayed on UI or CLI.
		* `options` - List of possible values for this variable.  If type is integer or date, then the array of string will be  converted to array of integers or date during runtime.
		* `min_value` - Minimum value of the variable. Applicable for integer type.
		* `max_value` - Maximum value of the variable. Applicable for integer type.
		* `min_length` - Minimum length of the variable value. Applicable for string type.
		* `max_length` - Maximum length of the variable value. Applicable for string type.
		* `matches` - Regex for the variable value.
		* `position` - Relative position of this variable in a list.
		* `group_by` - Display name of the group this variable belongs to.
		* `source` - Source of this meta-data.
	* `link` - Reference link to the variable value By default the expression will point to self.value.

* `settings` - Environment variables used by the job while performing an action. Nested `settings` blocks have the following structure:
	* `name` - Name of the variable.
	* `value` - Value for the variable or reference to the value.
	* `metadata` - User editable metadata for the variables. Nested `metadata` blocks have the following structure:
		* `type` - Type of the variable.
		* `aliases` - List of aliases for the variable name.
		* `description` - Description of the meta data.
		* `default_value` - Default value for the variable, if the override value is not specified.
		* `secure` - Is the variable secure or sensitive ?.
		* `immutable` - Is the variable readonly ?.
		* `hidden` - If true, the variable will not be displayed on UI or CLI.
		* `options` - List of possible values for this variable.  If type is integer or date, then the array of string will be  converted to array of integers or date during runtime.
		* `min_value` - Minimum value of the variable. Applicable for integer type.
		* `max_value` - Maximum value of the variable. Applicable for integer type.
		* `min_length` - Minimum length of the variable value. Applicable for string type.
		* `max_length` - Maximum length of the variable value. Applicable for string type.
		* `matches` - Regex for the variable value.
		* `position` - Relative position of this variable in a list.
		* `group_by` - Display name of the group this variable belongs to.
		* `source` - Source of this meta-data.
	* `link` - Reference link to the variable value By default the expression will point to self.value.

* `tags` - User defined tags, while running the job.

* `id` - Job ID.

* `name` - Job name, uniquely derived from the related action.

* `description` - Job description derived from the related action.

* `location` - List of action locations supported by IBM Cloud Schematics service.  **Note** this does not limit the location of the resources provisioned using Schematics.

* `resource_group` - Resource group name derived from the related action.

* `submitted_at` - Job submission time.

* `submitted_by` - E-mail address of the user who submitted the job.

* `start_at` - Job start time.

* `end_at` - Job end time.

* `duration` - Duration of job execution, for example, `40 sec`.

* `status` - Job Status. Nested `status` blocks have the following structure:
	* `action_job_status` - Action Job Status. Nested `action_job_status` blocks have the following structure:
		* `action_name` - Action name.
		* `status_code` - Status of the jobs.
		* `status_message` - Action job status message to be displayed along with the `action_status_code`.
		* `bastion_status_code` - Status of the resources.
		* `bastion_status_message` - Bastion status message to be displayed along with the `bastion_status_code`.
		* `targets_status_code` - Status of the resources.
		* `targets_status_message` - Aggregated status message for all target resources, to be displayed along with the `targets_status_code`.
		* `updated_at` - Job status updation timestamp.

* `data` - Job data. Nested `data` blocks have the following structure:
	* `job_type` - Type of the job.
	* `action_job_data` - Action Job data. Nested `action_job_data` blocks have the following structure:
		* `action_name` - Flow name.
		* `inputs` - Input variables data used by an action job. Nested `inputs` blocks have the following structure:
			* `name` - Name of the variable.
			* `value` - Value for the variable or reference to the value.
			* `metadata` - User editable metadata for the variables. Nested `metadata` blocks have the following structure:
				* `type` - Type of the variable.
				* `aliases` - List of aliases for the variable name.
				* `description` - Description of the meta data.
				* `default_value` - Default value for the variable, if the override value is not specified.
				* `secure` - Is the variable secure or sensitive ?.
				* `immutable` - Is the variable readonly ?.
				* `hidden` - If true, the variable will not be displayed on UI or CLI.
				* `options` - List of possible values for this variable.  If type is integer or date, then the array of string will be  converted to array of integers or date during runtime.
				* `min_value` - Minimum value of the variable. Applicable for integer type.
				* `max_value` - Maximum value of the variable. Applicable for integer type.
				* `min_length` - Minimum length of the variable value. Applicable for string type.
				* `max_length` - Maximum length of the variable value. Applicable for string type.
				* `matches` - Regex for the variable value.
				* `position` - Relative position of this variable in a list.
				* `group_by` - Display name of the group this variable belongs to.
				* `source` - Source of this meta-data.
			* `link` - Reference link to the variable value By default the expression will point to self.value.
		* `outputs` - Output variables data from an action job. Nested `outputs` blocks have the following structure:
			* `name` - Name of the variable.
			* `value` - Value for the variable or reference to the value.
			* `metadata` - User editable metadata for the variables. Nested `metadata` blocks have the following structure:
				* `type` - Type of the variable.
				* `aliases` - List of aliases for the variable name.
				* `description` - Description of the meta data.
				* `default_value` - Default value for the variable, if the override value is not specified.
				* `secure` - Is the variable secure or sensitive ?.
				* `immutable` - Is the variable readonly ?.
				* `hidden` - If true, the variable will not be displayed on UI or CLI.
				* `options` - List of possible values for this variable.  If type is integer or date, then the array of string will be  converted to array of integers or date during runtime.
				* `min_value` - Minimum value of the variable. Applicable for integer type.
				* `max_value` - Maximum value of the variable. Applicable for integer type.
				* `min_length` - Minimum length of the variable value. Applicable for string type.
				* `max_length` - Maximum length of the variable value. Applicable for string type.
				* `matches` - Regex for the variable value.
				* `position` - Relative position of this variable in a list.
				* `group_by` - Display name of the group this variable belongs to.
				* `source` - Source of this meta-data.
			* `link` - Reference link to the variable value By default the expression will point to self.value.
		* `settings` - Environment variables used by all the templates in an action. Nested `settings` blocks have the following structure:
			* `name` - Name of the variable.
			* `value` - Value for the variable or reference to the value.
			* `metadata` - User editable metadata for the variables. Nested `metadata` blocks have the following structure:
				* `type` - Type of the variable.
				* `aliases` - List of aliases for the variable name.
				* `description` - Description of the meta data.
				* `default_value` - Default value for the variable, if the override value is not specified.
				* `secure` - Is the variable secure or sensitive ?.
				* `immutable` - Is the variable readonly ?.
				* `hidden` - If true, the variable will not be displayed on UI or CLI.
				* `options` - List of possible values for this variable.  If type is integer or date, then the array of string will be  converted to array of integers or date during runtime.
				* `min_value` - Minimum value of the variable. Applicable for integer type.
				* `max_value` - Maximum value of the variable. Applicable for integer type.
				* `min_length` - Minimum length of the variable value. Applicable for string type.
				* `max_length` - Maximum length of the variable value. Applicable for string type.
				* `matches` - Regex for the variable value.
				* `position` - Relative position of this variable in a list.
				* `group_by` - Display name of the group this variable belongs to.
				* `source` - Source of this meta-data.
			* `link` - Reference link to the variable value By default the expression will point to self.value.
		* `updated_at` - Job status updation timestamp.

* `targets_ini` - Inventory of host and host group for the playbook in `INI` file format. For example, `"targets_ini": "[webserverhost]  172.22.192.6  [dbhost]  172.22.192.5"`. For more information, about an inventory host group syntax, see [Inventory host groups](/docs/schematics?topic=schematics-schematics-cli-reference#schematics-inventory-host-grps).

* `bastion` - Complete target details with the user inputs and the system generated data. Nested `bastion` blocks have the following structure:
	* `name` - Target name.
	* `type` - Target type (`cluster`, `vsi`, `icd`, `vpc`).
	* `description` - Target description.
	* `resource_query` - Resource selection query string.
	* `credential` - Override credential for each resource.  Reference to credentials values, used by all the resources.
	* `id` - Target ID.
	* `created_at` - Targets creation time.
	* `created_by` - E-mail address of the user who created the targets.
	* `updated_at` - Targets updation time.
	* `updated_by` - E-mail address of user who updated the targets.
	* `sys_lock` - System lock status. Nested `sys_lock` blocks have the following structure:
		* `sys_locked` - Is the Workspace locked by the Schematic action ?.
		* `sys_locked_by` - Name of the user who performed the action, that lead to lock the Workspace.
		* `sys_locked_at` - When the user performed the action that lead to lock the Workspace ?.
	* `resource_ids` - Array of the resource IDs.

* `log_summary` - Job log summary record. Nested `log_summary` blocks have the following structure:
	* `job_id` - Workspace ID.
	* `job_type` - Type of Job.
	* `log_start_at` - Job log start timestamp.
	* `log_analyzed_till` - Job log update timestamp.
	* `elapsed_time` - Job log elapsed time (`log_analyzed_till - log_start_at`).
	* `log_errors` - Job log errors. Nested `log_errors` blocks have the following structure:
		* `error_code` - Error code in the Log.
		* `error_msg` - Summary error message in the log.
		* `error_count` - Number of occurrence.
	* `repo_download_job` - Repo download Job log summary. Nested `repo_download_job` blocks have the following structure:
		* `scanned_file_count` - Number of files scanned.
		* `quarantined_file_count` - Number of files quarantined.
		* `detected_filetype` - Detected template or data file type.
		* `inputs_count` - Number of inputs detected.
		* `outputs_count` - Number of outputs detected.
	* `action_job` - Flow Job log summary. Nested `action_job` blocks have the following structure:
		* `target_count` - number of targets or hosts.
		* `task_count` - number of tasks in playbook.
		* `play_count` - number of plays in playbook.
		* `recap` - Recap records. Nested `recap` blocks have the following structure:
			* `target` - List of target or host name.
			* `ok` - Number of OK.
			* `changed` - Number of changed.
			* `failed` - Number of failed.
			* `skipped` - Number of skipped.
			* `unreachable` - Number of unreachable.

* `log_store_url` - Job log store URL.

* `state_store_url` - Job state store URL.

* `results_url` - Job results store URL.

* `updated_at` - Job status updation timestamp.

