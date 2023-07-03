---
subcategory: "Schematics"
layout: "ibm"
page_title: "IBM : ibm_schematics_job"
sidebar_current: "docs-ibm-resource-schematics-job"
description: |-
  Manages Schematics job.
---

# ibm_schematics_job
Create, update, and delete `ibm_schematics_job`. For more information, about IBM Cloud Schematics job, refer to [setting up jobs](https://cloud.ibm.com/docs/schematics?topic=schematics-action-setup#action-jobs).

## Example usage

```terraform
resource "ibm_schematics_job" "schematics_job" {
  command_object = "action"
  command_object_id = "<action_id>"
  command_name = "ansible_playbook_run | ansible_playbook_check"
  command_parameter = "<yml_file_name>"
  location = "us-east"
}
```

## Argument reference

Review the argument reference that you can specify for your resource.

* `bastion` - (Optional, List) Describes a bastion resource. MaxItems: 1.
Nested scheme for **bastion**:
	* `name` - (Optional, String) Bastion Name(Unique).
	* `host` - (Optional, String) Reference to the Inventory resource definition.
* `command_name` - (Required, String) Schematics job command name.
  * Constraints: Allowable values are: workspace_plan, workspace_apply, workspace_destroy, workspace_refresh, ansible_playbook_run, ansible_playbook_check, create_action, put_action, patch_action, delete_action, system_key_enable, system_key_delete, system_key_disable, system_key_rotate, system_key_restore, create_workspace, put_workspace, patch_workspace, delete_workspace, create_cart, create_environment, put_environment, delete_environment, environment_init, environment_install, environment_uninstall, repository_process
* `command_object` - (Required, String) Name of the Schematics automation resource.
  * Constraints: Allowable values are: workspace, action, system, environment
* `command_object_id` - (Required, String) Job command object id (workspace-id, action-id).
* `command_options` - (Optional, List) Command line options for the command.
* `command_parameter` - (Optional, String) Schematics job command parameter (playbook-name).
* `data` - (Optional, List) Job data.
Nested scheme for **data**:
	* `job_type` - (Required, String) Type of Job.
	  * Constraints: Allowable values are: repo_download_job, workspace_job, action_job, system_job, flow-job
	* `workspace_job_data` - (Optional, List) Workspace Job data. MaxItems: 1.
	Nested scheme for **workspace_job_data**:
		* `workspace_name` - (Optional, String) Workspace name.
		* `flow_id` - (Optional, String) Flow Id.
		* `flow_name` - (Optional, String) Flow name.
		* `inputs` - (Optional, List) Input variables data used by the Workspace Job.
		Nested scheme for **inputs**:
			* `name` - (Optional, String) Name of the variable.
			* `value` - (Optional, String) Value for the variable or reference to the value.
			* `metadata` - (Optional, List) User editable metadata for the variables. MaxItems: 1.
			Nested scheme for **metadata**:
				* `type` - (Optional, String) Type of the variable.
				  * Constraints: Allowable values are: boolean, string, integer, date, array, list, map, complex
				* `aliases` - (Optional, List) List of aliases for the variable name.
				* `description` - (Optional, String) Description of the meta data.
				* `default_value` - (Optional, String) Default value for the variable, if the override value is not specified.
				* `secure` - (Optional, Boolean) Is the variable secure or sensitive ?.
				* `immutable` - (Optional, Boolean) Is the variable readonly ?.
				* `hidden` - (Optional, Boolean) If true, the variable will not be displayed on UI or CLI.
				* `options` - (Optional, List) List of possible values for this variable.  If type is integer or date, then the array of string will be  converted to array of integers or date during runtime.
				* `min_value` - (Optional, Integer) Minimum value of the variable. Applicable for integer type.
				* `max_value` - (Optional, Integer) Maximum value of the variable. Applicable for integer type.
				* `min_length` - (Optional, Integer) Minimum length of the variable value. Applicable for string type.
				* `max_length` - (Optional, Integer) Maximum length of the variable value. Applicable for string type.
				* `matches` - (Optional, String) Regex for the variable value.
				* `position` - (Optional, Integer) Relative position of this variable in a list.
				* `group_by` - (Optional, String) Display name of the group this variable belongs to.
				* `source` - (Optional, String) Source of this meta-data.
			* `link` - (Optional, String) Reference link to the variable value By default the expression will point to self.value.
		* `outputs` - (Optional, List) Output variables data from the Workspace Job.
		Nested scheme for **outputs**:
			* `name` - (Optional, String) Name of the variable.
			* `value` - (Optional, String) Value for the variable or reference to the value.
			* `metadata` - (Optional, List) User editable metadata for the variables. MaxItems: 1.
			Nested scheme for **metadata**:
				* `type` - (Optional, String) Type of the variable.
				  * Constraints: Allowable values are: boolean, string, integer, date, array, list, map, complex
				* `aliases` - (Optional, List) List of aliases for the variable name.
				* `description` - (Optional, String) Description of the meta data.
				* `default_value` - (Optional, String) Default value for the variable, if the override value is not specified.
				* `secure` - (Optional, Boolean) Is the variable secure or sensitive ?.
				* `immutable` - (Optional, Boolean) Is the variable readonly ?.
				* `hidden` - (Optional, Boolean) If true, the variable will not be displayed on UI or CLI.
				* `options` - (Optional, List) List of possible values for this variable.  If type is integer or date, then the array of string will be  converted to array of integers or date during runtime.
				* `min_value` - (Optional, Integer) Minimum value of the variable. Applicable for integer type.
				* `max_value` - (Optional, Integer) Maximum value of the variable. Applicable for integer type.
				* `min_length` - (Optional, Integer) Minimum length of the variable value. Applicable for string type.
				* `max_length` - (Optional, Integer) Maximum length of the variable value. Applicable for string type.
				* `matches` - (Optional, String) Regex for the variable value.
				* `position` - (Optional, Integer) Relative position of this variable in a list.
				* `group_by` - (Optional, String) Display name of the group this variable belongs to.
				* `source` - (Optional, String) Source of this meta-data.
			* `link` - (Optional, String) Reference link to the variable value By default the expression will point to self.value.
		* `settings` - (Optional, List) Environment variables used by all the templates in the Workspace. MaxItems: 1.
		Nested scheme for **settings**:
			* `name` - (Optional, String) Name of the variable.
			* `value` - (Optional, String) Value for the variable or reference to the value.
			* `metadata` - (Optional, List) User editable metadata for the variables.
			Nested scheme for **metadata**:
				* `type` - (Optional, String) Type of the variable.
				  * Constraints: Allowable values are: boolean, string, integer, date, array, list, map, complex
				* `aliases` - (Optional, List) List of aliases for the variable name.
				* `description` - (Optional, String) Description of the meta data.
				* `default_value` - (Optional, String) Default value for the variable, if the override value is not specified.
				* `secure` - (Optional, Boolean) Is the variable secure or sensitive ?.
				* `immutable` - (Optional, Boolean) Is the variable readonly ?.
				* `hidden` - (Optional, Boolean) If true, the variable will not be displayed on UI or CLI.
				* `options` - (Optional, List) List of possible values for this variable.  If type is integer or date, then the array of string will be  converted to array of integers or date during runtime.
				* `min_value` - (Optional, Integer) Minimum value of the variable. Applicable for integer type.
				* `max_value` - (Optional, Integer) Maximum value of the variable. Applicable for integer type.
				* `min_length` - (Optional, Integer) Minimum length of the variable value. Applicable for string type.
				* `max_length` - (Optional, Integer) Maximum length of the variable value. Applicable for string type.
				* `matches` - (Optional, String) Regex for the variable value.
				* `position` - (Optional, Integer) Relative position of this variable in a list.
				* `group_by` - (Optional, String) Display name of the group this variable belongs to.
				* `source` - (Optional, String) Source of this meta-data.
			* `link` - (Optional, String) Reference link to the variable value By default the expression will point to self.value.
		* `template_data` - (Optional, List) Input / output data of the Template in the Workspace Job.
		Nested scheme for **template_data**:
			* `template_id` - (Optional, String) Template Id.
			* `template_name` - (Optional, String) Template name.
			* `flow_index` - (Optional, Integer) Index of the template in the Flow.
			* `inputs` - (Optional, List) Job inputs used by the Templates.
			Nested scheme for **inputs**:
				* `name` - (Optional, String) Name of the variable.
				* `value` - (Optional, String) Value for the variable or reference to the value.
				* `metadata` - (Optional, List) User editable metadata for the variables. MaxItems: 1.
				Nested scheme for **metadata**:
					* `type` - (Optional, String) Type of the variable.
					  * Constraints: Allowable values are: boolean, string, integer, date, array, list, map, complex
					* `aliases` - (Optional, List) List of aliases for the variable name.
					* `description` - (Optional, String) Description of the meta data.
					* `default_value` - (Optional, String) Default value for the variable, if the override value is not specified.
					* `secure` - (Optional, Boolean) Is the variable secure or sensitive ?.
					* `immutable` - (Optional, Boolean) Is the variable readonly ?.
					* `hidden` - (Optional, Boolean) If true, the variable will not be displayed on UI or CLI.
					* `options` - (Optional, List) List of possible values for this variable.  If type is integer or date, then the array of string will be  converted to array of integers or date during runtime.
					* `min_value` - (Optional, Integer) Minimum value of the variable. Applicable for integer type.
					* `max_value` - (Optional, Integer) Maximum value of the variable. Applicable for integer type.
					* `min_length` - (Optional, Integer) Minimum length of the variable value. Applicable for string type.
					* `max_length` - (Optional, Integer) Maximum length of the variable value. Applicable for string type.
					* `matches` - (Optional, String) Regex for the variable value.
					* `position` - (Optional, Integer) Relative position of this variable in a list.
					* `group_by` - (Optional, String) Display name of the group this variable belongs to.
					* `source` - (Optional, String) Source of this meta-data.
				* `link` - (Optional, String) Reference link to the variable value By default the expression will point to self.value.
			* `outputs` - (Optional, List) Job output from the Templates.
			Nested scheme for **outputs**:
				* `name` - (Optional, String) Name of the variable.
				* `value` - (Optional, String) Value for the variable or reference to the value.
				* `metadata` - (Optional, List) User editable metadata for the variables. MaxItems: 1.
				Nested scheme for **metadata**:
					* `type` - (Optional, String) Type of the variable.
					  * Constraints: Allowable values are: boolean, string, integer, date, array, list, map, complex
					* `aliases` - (Optional, List) List of aliases for the variable name.
					* `description` - (Optional, String) Description of the meta data.
					* `default_value` - (Optional, String) Default value for the variable, if the override value is not specified.
					* `secure` - (Optional, Boolean) Is the variable secure or sensitive ?.
					* `immutable` - (Optional, Boolean) Is the variable readonly ?.
					* `hidden` - (Optional, Boolean) If true, the variable will not be displayed on UI or CLI.
					* `options` - (Optional, List) List of possible values for this variable.  If type is integer or date, then the array of string will be  converted to array of integers or date during runtime.
					* `min_value` - (Optional, Integer) Minimum value of the variable. Applicable for integer type.
					* `max_value` - (Optional, Integer) Maximum value of the variable. Applicable for integer type.
					* `min_length` - (Optional, Integer) Minimum length of the variable value. Applicable for string type.
					* `max_length` - (Optional, Integer) Maximum length of the variable value. Applicable for string type.
					* `matches` - (Optional, String) Regex for the variable value.
					* `position` - (Optional, Integer) Relative position of this variable in a list.
					* `group_by` - (Optional, String) Display name of the group this variable belongs to.
					* `source` - (Optional, String) Source of this meta-data.
				* `link` - (Optional, String) Reference link to the variable value By default the expression will point to self.value.
			* `settings` - (Optional, List) Environment variables used by the template.
			Nested scheme for **settings**:
				* `name` - (Optional, String) Name of the variable.
				* `value` - (Optional, String) Value for the variable or reference to the value.
				* `metadata` - (Optional, List) User editable metadata for the variables. MaxItems: 1.
				Nested scheme for **metadata**:
					* `type` - (Optional, String) Type of the variable.
					  * Constraints: Allowable values are: boolean, string, integer, date, array, list, map, complex
					* `aliases` - (Optional, List) List of aliases for the variable name.
					* `description` - (Optional, String) Description of the meta data.
					* `default_value` - (Optional, String) Default value for the variable, if the override value is not specified.
					* `secure` - (Optional, Boolean) Is the variable secure or sensitive ?.
					* `immutable` - (Optional, Boolean) Is the variable readonly ?.
					* `hidden` - (Optional, Boolean) If true, the variable will not be displayed on UI or CLI.
					* `options` - (Optional, List) List of possible values for this variable.  If type is integer or date, then the array of string will be  converted to array of integers or date during runtime.
					* `min_value` - (Optional, Integer) Minimum value of the variable. Applicable for integer type.
					* `max_value` - (Optional, Integer) Maximum value of the variable. Applicable for integer type.
					* `min_length` - (Optional, Integer) Minimum length of the variable value. Applicable for string type.
					* `max_length` - (Optional, Integer) Maximum length of the variable value. Applicable for string type.
					* `matches` - (Optional, String) Regex for the variable value.
					* `position` - (Optional, Integer) Relative position of this variable in a list.
					* `group_by` - (Optional, String) Display name of the group this variable belongs to.
					* `source` - (Optional, String) Source of this meta-data.
				* `link` - (Optional, String) Reference link to the variable value By default the expression will point to self.value.
			* `updated_at` - (Optional, String) Job status updation timestamp.
		* `updated_at` - (Optional, String) Job status updation timestamp.
	* `action_job_data` - (Optional, List) Action Job data.
	Nested scheme for **action_job_data**:
		* `action_name` - (Optional, String) Flow name.
		* `inputs` - (Optional, List) Input variables data used by the Action Job.
		Nested scheme for **inputs**:
			* `name` - (Optional, String) Name of the variable.
			* `value` - (Optional, String) Value for the variable or reference to the value.
			* `metadata` - (Optional, List) User editable metadata for the variables.
			Nested scheme for **metadata**:
				* `type` - (Optional, String) Type of the variable.
				  * Constraints: Allowable values are: boolean, string, integer, date, array, list, map, complex
				* `aliases` - (Optional, List) List of aliases for the variable name.
				* `description` - (Optional, String) Description of the meta data.
				* `default_value` - (Optional, String) Default value for the variable, if the override value is not specified.
				* `secure` - (Optional, Boolean) Is the variable secure or sensitive ?.
				* `immutable` - (Optional, Boolean) Is the variable readonly ?.
				* `hidden` - (Optional, Boolean) If true, the variable will not be displayed on UI or CLI.
				* `options` - (Optional, List) List of possible values for this variable.  If type is integer or date, then the array of string will be  converted to array of integers or date during runtime.
				* `min_value` - (Optional, Integer) Minimum value of the variable. Applicable for integer type.
				* `max_value` - (Optional, Integer) Maximum value of the variable. Applicable for integer type.
				* `min_length` - (Optional, Integer) Minimum length of the variable value. Applicable for string type.
				* `max_length` - (Optional, Integer) Maximum length of the variable value. Applicable for string type.
				* `matches` - (Optional, String) Regex for the variable value.
				* `position` - (Optional, Integer) Relative position of this variable in a list.
				* `group_by` - (Optional, String) Display name of the group this variable belongs to.
				* `source` - (Optional, String) Source of this meta-data.
			* `link` - (Optional, String) Reference link to the variable value By default the expression will point to self.value.
		* `outputs` - (Optional, List) Output variables data from the Action Job.
		Nested scheme for **outputs**:
			* `name` - (Optional, String) Name of the variable.
			* `value` - (Optional, String) Value for the variable or reference to the value.
			* `metadata` - (Optional, List) User editable metadata for the variables.
			Nested scheme for **metadata**:
				* `type` - (Optional, String) Type of the variable.
				  * Constraints: Allowable values are: boolean, string, integer, date, array, list, map, complex
				* `aliases` - (Optional, List) List of aliases for the variable name.
				* `description` - (Optional, String) Description of the meta data.
				* `default_value` - (Optional, String) Default value for the variable, if the override value is not specified.
				* `secure` - (Optional, Boolean) Is the variable secure or sensitive ?.
				* `immutable` - (Optional, Boolean) Is the variable readonly ?.
				* `hidden` - (Optional, Boolean) If true, the variable will not be displayed on UI or CLI.
				* `options` - (Optional, List) List of possible values for this variable.  If type is integer or date, then the array of string will be  converted to array of integers or date during runtime.
				* `min_value` - (Optional, Integer) Minimum value of the variable. Applicable for integer type.
				* `max_value` - (Optional, Integer) Maximum value of the variable. Applicable for integer type.
				* `min_length` - (Optional, Integer) Minimum length of the variable value. Applicable for string type.
				* `max_length` - (Optional, Integer) Maximum length of the variable value. Applicable for string type.
				* `matches` - (Optional, String) Regex for the variable value.
				* `position` - (Optional, Integer) Relative position of this variable in a list.
				* `group_by` - (Optional, String) Display name of the group this variable belongs to.
				* `source` - (Optional, String) Source of this meta-data.
			* `link` - (Optional, String) Reference link to the variable value By default the expression will point to self.value.
		* `settings` - (Optional, List) Environment variables used by all the templates in the Action.
		Nested scheme for **settings**:
			* `name` - (Optional, String) Name of the variable.
			* `value` - (Optional, String) Value for the variable or reference to the value.
			* `metadata` - (Optional, List) User editable metadata for the variables.
			Nested scheme for **metadata**:
				* `type` - (Optional, String) Type of the variable.
				  * Constraints: Allowable values are: boolean, string, integer, date, array, list, map, complex
				* `aliases` - (Optional, List) List of aliases for the variable name.
				* `description` - (Optional, String) Description of the meta data.
				* `default_value` - (Optional, String) Default value for the variable, if the override value is not specified.
				* `secure` - (Optional, Boolean) Is the variable secure or sensitive ?.
				* `immutable` - (Optional, Boolean) Is the variable readonly ?.
				* `hidden` - (Optional, Boolean) If true, the variable will not be displayed on UI or CLI.
				* `options` - (Optional, List) List of possible values for this variable.  If type is integer or date, then the array of string will be  converted to array of integers or date during runtime.
				* `min_value` - (Optional, Integer) Minimum value of the variable. Applicable for integer type.
				* `max_value` - (Optional, Integer) Maximum value of the variable. Applicable for integer type.
				* `min_length` - (Optional, Integer) Minimum length of the variable value. Applicable for string type.
				* `max_length` - (Optional, Integer) Maximum length of the variable value. Applicable for string type.
				* `matches` - (Optional, String) Regex for the variable value.
				* `position` - (Optional, Integer) Relative position of this variable in a list.
				* `group_by` - (Optional, String) Display name of the group this variable belongs to.
				* `source` - (Optional, String) Source of this meta-data.
			* `link` - (Optional, String) Reference link to the variable value By default the expression will point to self.value.
		* `updated_at` - (Optional, String) Job status updation timestamp.
		* `inventory_record` - (Optional, List) Complete inventory resource details with user inputs and system generated data. MaxItems: 1.
		Nested scheme for **inventory_record**: 
			* `name` - (Optional, String) The unique name of your Inventory.  The name can be up to 128 characters long and can include alphanumeric  characters, spaces, dashes, and underscores.
			* `id` - (Optional, String) Inventory id.
			* `description` - (Optional, String) The description of your Inventory.  The description can be up to 2048 characters long in size.
			* `location` - (Optional, String) List of locations supported by IBM Cloud Schematics service.  While creating your workspace or action, choose the right region, since it cannot be changed.  Note, this does not limit the location of the IBM Cloud resources, provisioned using Schematics.
			  * Constraints: Allowable values are: us-south, us-east, eu-gb, eu-de
			* `resource_group` - (Optional, String) Resource-group name for the Inventory definition.  By default, Inventory will be created in Default Resource Group.
			* `created_at` - (Optional, String) Inventory creation time.
			* `created_by` - (Optional, String) Email address of user who created the Inventory.
			* `updated_at` - (Optional, String) Inventory updation time.
			* `updated_by` - (Optional, String) Email address of user who updated the Inventory.
			* `inventories_ini` - (Optional, String) Input inventory of host and host group for the playbook,  in the .ini file format.
			* `resource_queries` - (Optional, List) Input resource queries that is used to dynamically generate  the inventory of host and host group for the playbook.
		* `materialized_inventory` - (Optional, String) Materialized inventory details used by the Action Job, in .ini format.
	* `system_job_data` - (Optional, List) Controls Job data. MaxItems: 1.
	Nested scheme for **system_job_data**:
		* `key_id` - (Optional, String) Key ID for which key event is generated.
		* `schematics_resource_id` - (Optional, List) List of the schematics resource id.
		* `updated_at` - (Optional, String) Job status updation timestamp.
	* `flow_job_data` - (Optional, List) Flow Job data.
	Nested scheme for **flow_job_data**:
		* `flow_id` - (Optional, String) Flow ID.
		* `flow_name` - (Optional, String) Flow Name.
		* `workitems` - (Optional, List) Job data used by each workitem Job.
		Nested scheme for **workitems**:
			* `command_object_id` - (Optional, String) command object id.
			* `command_object_name` - (Optional, String) command object name.
			* `layers` - (Optional, String) layer name.
			* `source_type` - (Optional, String) Type of source for the Template.
			  * Constraints: Allowable values are: local, git_hub, git_hub_enterprise, git_lab, ibm_git_lab, ibm_cloud_catalog, external_scm, cos_bucket
			* `source` - (Optional, List) Source of templates, playbooks, or controls. MaxItems: 1.
			Nested scheme for **source**:
				* `source_type` - (Required, String) Type of source for the Template.
				  * Constraints: Allowable values are: local, git_hub, git_hub_enterprise, git_lab, ibm_git_lab, ibm_cloud_catalog, external_scm, cos_bucket
				* `git` - (Optional, List) Connection details to Git source. MaxItems: 1.
				Nested scheme for **git**:
					* `computed_git_repo_url` - (Optional, String) The Complete URL which is computed by git_repo_url, git_repo_folder and branch.
					* `git_repo_url` - (Optional, String) URL to the GIT Repo that can be used to clone the template.
					* `git_token` - (Optional, String) Personal Access Token to connect to Git URLs.
					* `git_repo_folder` - (Optional, String) Name of the folder in the Git Repo, that contains the template.
					* `git_release` - (Optional, String) Name of the release tag, used to fetch the Git Repo.
					* `git_branch` - (Optional, String) Name of the branch, used to fetch the Git Repo.
				* `catalog` - (Optional, List) Connection details to IBM Cloud Catalog source. MaxItems: 1.
				Nested scheme for **catalog**:
					* `catalog_name` - (Optional, String) name of the private catalog.
					* `offering_name` - (Optional, String) Name of the offering in the IBM Catalog.
					* `offering_version` - (Optional, String) Version string of the offering in the IBM Catalog.
					* `offering_kind` - (Optional, String) Type of the offering, in the IBM Catalog.
					* `offering_id` - (Optional, String) Id of the offering the IBM Catalog.
					* `offering_version_id` - (Optional, String) Id of the offering version the IBM Catalog.
					* `offering_repo_url` - (Optional, String) Repo Url of the offering, in the IBM Catalog.
			* `inputs` - (Optional, List) Input variables data for the workItem used in FlowJob.
			Nested scheme for **inputs**:
				* `name` - (Optional, String) Name of the variable.
				* `value` - (Optional, String) Value for the variable or reference to the value.
				* `metadata` - (Optional, List) User editable metadata for the variables. MaxItems: 1.
				Nested scheme for **metadata**:
					* `type` - (Optional, String) Type of the variable.
					  * Constraints: Allowable values are: boolean, string, integer, date, array, list, map, complex
					* `aliases` - (Optional, List) List of aliases for the variable name.
					* `description` - (Optional, String) Description of the meta data.
					* `default_value` - (Optional, String) Default value for the variable, if the override value is not specified.
					* `secure` - (Optional, Boolean) Is the variable secure or sensitive ?.
					* `immutable` - (Optional, Boolean) Is the variable readonly ?.
					* `hidden` - (Optional, Boolean) If true, the variable will not be displayed on UI or CLI.
					* `options` - (Optional, List) List of possible values for this variable.  If type is integer or date, then the array of string will be  converted to array of integers or date during runtime.
					* `min_value` - (Optional, Integer) Minimum value of the variable. Applicable for integer type.
					* `max_value` - (Optional, Integer) Maximum value of the variable. Applicable for integer type.
					* `min_length` - (Optional, Integer) Minimum length of the variable value. Applicable for string type.
					* `max_length` - (Optional, Integer) Maximum length of the variable value. Applicable for string type.
					* `matches` - (Optional, String) Regex for the variable value.
					* `position` - (Optional, Integer) Relative position of this variable in a list.
					* `group_by` - (Optional, String) Display name of the group this variable belongs to.
					* `source` - (Optional, String) Source of this meta-data.
				* `link` - (Optional, String) Reference link to the variable value By default the expression will point to self.value.
			* `outputs` - (Optional, List) Output variables for the workItem.
			Nested scheme for **outputs**:
				* `name` - (Optional, String) Name of the variable.
				* `value` - (Optional, String) Value for the variable or reference to the value.
				* `metadata` - (Optional, List) User editable metadata for the variables. MaxItems: 1.
				Nested scheme for **metadata**:
					* `type` - (Optional, String) Type of the variable.
					  * Constraints: Allowable values are: boolean, string, integer, date, array, list, map, complex
					* `aliases` - (Optional, List) List of aliases for the variable name.
					* `description` - (Optional, String) Description of the meta data.
					* `default_value` - (Optional, String) Default value for the variable, if the override value is not specified.
					* `secure` - (Optional, Boolean) Is the variable secure or sensitive ?.
					* `immutable` - (Optional, Boolean) Is the variable readonly ?.
					* `hidden` - (Optional, Boolean) If true, the variable will not be displayed on UI or CLI.
					* `options` - (Optional, List) List of possible values for this variable.  If type is integer or date, then the array of string will be  converted to array of integers or date during runtime.
					* `min_value` - (Optional, Integer) Minimum value of the variable. Applicable for integer type.
					* `max_value` - (Optional, Integer) Maximum value of the variable. Applicable for integer type.
					* `min_length` - (Optional, Integer) Minimum length of the variable value. Applicable for string type.
					* `max_length` - (Optional, Integer) Maximum length of the variable value. Applicable for string type.
					* `matches` - (Optional, String) Regex for the variable value.
					* `position` - (Optional, Integer) Relative position of this variable in a list.
					* `group_by` - (Optional, String) Display name of the group this variable belongs to.
					* `source` - (Optional, String) Source of this meta-data.
				* `link` - (Optional, String) Reference link to the variable value By default the expression will point to self.value.
			* `settings` - (Optional, List) Environment variables for the workItem.
			Nested scheme for **settings**:
				* `name` - (Optional, String) Name of the variable.
				* `value` - (Optional, String) Value for the variable or reference to the value.
				* `metadata` - (Optional, List) User editable metadata for the variables. MaxItems: 1.
				Nested scheme for **metadata**:
					* `type` - (Optional, String) Type of the variable.
					  * Constraints: Allowable values are: boolean, string, integer, date, array, list, map, complex
					* `aliases` - (Optional, List) List of aliases for the variable name.
					* `description` - (Optional, String) Description of the meta data.
					* `default_value` - (Optional, String) Default value for the variable, if the override value is not specified.
					* `secure` - (Optional, Boolean) Is the variable secure or sensitive ?.
					* `immutable` - (Optional, Boolean) Is the variable readonly ?.
					* `hidden` - (Optional, Boolean) If true, the variable will not be displayed on UI or CLI.
					* `options` - (Optional, List) List of possible values for this variable.  If type is integer or date, then the array of string will be  converted to array of integers or date during runtime.
					* `min_value` - (Optional, Integer) Minimum value of the variable. Applicable for integer type.
					* `max_value` - (Optional, Integer) Maximum value of the variable. Applicable for integer type.
					* `min_length` - (Optional, Integer) Minimum length of the variable value. Applicable for string type.
					* `max_length` - (Optional, Integer) Maximum length of the variable value. Applicable for string type.
					* `matches` - (Optional, String) Regex for the variable value.
					* `position` - (Optional, Integer) Relative position of this variable in a list.
					* `group_by` - (Optional, String) Display name of the group this variable belongs to.
					* `source` - (Optional, String) Source of this meta-data.
				* `link` - (Optional, String) Reference link to the variable value By default the expression will point to self.value.
			* `last_job` - (Optional, List) Status of the last job executed by the workitem. MaxItems: 1.
			Nested scheme for **last_job**:
				* `command_object` - (Optional, String) Name of the Schematics automation resource.
				  * Constraints: Allowable values are: workspace, action, system, environment
				* `command_object_name` - (Optional, String) command object name (workspace_name/action_name).
				* `command_object_id` - (Optional, String) Workitem command object id, maps to workspace_id or action_id.
				* `command_name` - (Optional, String) Schematics job command name.
				  * Constraints: Allowable values are: workspace_plan, workspace_apply, workspace_destroy, workspace_refresh, ansible_playbook_run, ansible_playbook_check, create_action, put_action, patch_action, delete_action, system_key_enable, system_key_delete, system_key_disable, system_key_rotate, system_key_restore, create_workspace, put_workspace, patch_workspace, delete_workspace, create_cart, create_environment, put_environment, delete_environment, environment_init, environment_install, environment_uninstall, repository_process
				* `job_id` - (Optional, String) Workspace job id.
				* `job_status` - (Optional, String) Status of Jobs.
				  * Constraints: Allowable values are: job_pending, job_in_progress, job_finished, job_failed, job_cancelled
			* `updated_at` - (Optional, String) Job status updation timestamp.
		* `updated_at` - (Optional, String) Job status updation timestamp.
* `job_env_settings` - (Optional, List) Environment variables used by the Job while performing Action or Workspace.
Nested scheme for **job_env_settings**:
	* `name` - (Required, String) Name of the variable.
	* `value` - (Required, String) Value for the variable or reference to the value.
	* `metadata` - (Optional, List) User editable metadata for the variables.
	Nested scheme for **metadata**:
		* `type` - (Required, String) Type of the variable.
		  * Constraints: Allowable values are: boolean, string, integer, date, array, list, map, complex
		* `aliases` - (Optional, List) List of aliases for the variable name.
		* `description` - (Optional, String) Description of the meta data.
		* `default_value` - (Optional, String) Default value for the variable, if the override value is not specified.
		* `secure` - (Optional, Boolean) Is the variable secure or sensitive ?.
		* `immutable` - (Optional, Boolean) Is the variable readonly ?.
		* `hidden` - (Optional, Boolean) If true, the variable will not be displayed on UI or CLI.
		* `options` - (Optional, List) List of possible values for this variable.  If type is integer or date, then the array of string will be  converted to array of integers or date during runtime.
		* `min_value` - (Optional, Integer) Minimum value of the variable. Applicable for integer type.
		* `max_value` - (Optional, Integer) Maximum value of the variable. Applicable for integer type.
		* `min_length` - (Optional, Integer) Minimum length of the variable value. Applicable for string type.
		* `max_length` - (Optional, Integer) Maximum length of the variable value. Applicable for string type.
		* `matches` - (Optional, String) Regex for the variable value.
		* `position` - (Optional, Integer) Relative position of this variable in a list.
		* `group_by` - (Optional, String) Display name of the group this variable belongs to.
		* `source` - (Optional, String) Source of this meta-data.
	* `link` - (Optional, String) Reference link to the variable value By default the expression will point to self.value.
* `job_inputs` - (Optional, List) Job inputs used by Action or Workspace.
Nested scheme for **job_inputs**:
	* `name` - (Required, String) Name of the variable.
	* `value` - (Required, String) Value for the variable or reference to the value.
	* `metadata` - (Optional, List) User editable metadata for the variables.
	Nested scheme for **metadata**:
		* `type` - (Required, String) Type of the variable.
		  * Constraints: Allowable values are: boolean, string, integer, date, array, list, map, complex
		* `aliases` - (Optional, List) List of aliases for the variable name.
		* `description` - (Optional, String) Description of the meta data.
		* `default_value` - (Optional, String) Default value for the variable, if the override value is not specified.
		* `secure` - (Optional, Boolean) Is the variable secure or sensitive ?.
		* `immutable` - (Optional, Boolean) Is the variable readonly ?.
		* `hidden` - (Optional, Boolean) If true, the variable will not be displayed on UI or CLI.
		* `options` - (Optional, List) List of possible values for this variable.  If type is integer or date, then the array of string will be  converted to array of integers or date during runtime.
		* `min_value` - (Optional, Integer) Minimum value of the variable. Applicable for integer type.
		* `max_value` - (Optional, Integer) Maximum value of the variable. Applicable for integer type.
		* `min_length` - (Optional, Integer) Minimum length of the variable value. Applicable for string type.
		* `max_length` - (Optional, Integer) Maximum length of the variable value. Applicable for string type.
		* `matches` - (Optional, String) Regex for the variable value.
		* `position` - (Optional, Integer) Relative position of this variable in a list.
		* `group_by` - (Optional, String) Display name of the group this variable belongs to.
		* `source` - (Optional, String) Source of this meta-data.
	* `link` - (Optional, String) Reference link to the variable value By default the expression will point to self.value.
* `location` - (Optional, String) Location supported by IBM Cloud Schematics service.  While creating your workspace or action, choose the right region, since it cannot be changed.  Note, this does not limit the location of the IBM Cloud resources, provisioned using Schematics.
  * Constraints: Allowable values are: us-south, us-east, eu-gb, eu-de
* `log_summary` - (Optional, List) Job log summary record.
Nested scheme for **log_summary**:
	* `job_id` - (Optional, String) Workspace Id.
	* `job_type` - (Optional, String) Type of Job.
	  * Constraints: Allowable values are: repo_download_job, workspace_job, action_job, system_job, flow_job
	* `log_start_at` - (Optional, String) Job log start timestamp.
	* `log_analyzed_till` - (Optional, String) Job log update timestamp.
	* `elapsed_time` - (Optional, Float) Job log elapsed time (log_analyzed_till - log_start_at).
	* `log_errors` - (Optional, List) Job log errors.
	Nested scheme for **log_errors**:
		* `error_code` - (Optional, String) Error code in the Log.
		* `error_msg` - (Optional, String) Summary error message in the log.
		* `error_count` - (Optional, Float) Number of occurrence.
	* `repo_download_job` - (Optional, List) Repo download Job log summary.
	Nested scheme for **repo_download_job**:
		* `scanned_file_count` - (Optional, Float) Number of files scanned.
		* `quarantined_file_count` - (Optional, Float) Number of files quarantined.
		* `detected_filetype` - (Optional, String) Detected template or data file type.
		* `inputs_count` - (Optional, String) Number of inputs detected.
		* `outputs_count` - (Optional, String) Number of outputs detected.
	* `workspace_job` - (Optional, List) Workspace Job log summary. MaxItems: 1.
	Nested scheme for **workspace_job**:
		* `resources_add` - (Optional, Float) Number of resources add.
		* `resources_modify` - (Optional, Float) Number of resources modify.
		* `resources_destroy` - (Optional, Float) Number of resources destroy.
	* `flow_job` - (Optional, List) Flow Job log summary. MaxItems: 1.
	Nested scheme for **flow_job**:
		* `workitems_completed` - (Optional, Float) Number of workitems completed successfully.
		* `workitems_pending` - (Optional, Float) Number of workitems pending in the flow.
		* `workitems_failed` - (Optional, Float) Number of workitems failed.
		* `workitems` - (Optional, List)
		Nested scheme for **workitems**:
			* `workspace_id` - (Optional, String) workspace ID.
			* `job_id` - (Optional, String) workspace JOB ID.
			* `resources_add` - (Optional, Float) Number of resources add.
			* `resources_modify` - (Optional, Float) Number of resources modify.
			* `resources_destroy` - (Optional, Float) Number of resources destroy.
			* `log_url` - (Optional, String) Log url for job.
	* `action_job` - (Optional, List) Flow Job log summary.
	Nested scheme for **action_job**:
		* `target_count` - (Optional, Float) number of targets or hosts.
		* `task_count` - (Optional, Float) number of tasks in playbook.
		* `play_count` - (Optional, Float) number of plays in playbook.
		* `recap` - (Optional, List) Recap records.
		Nested scheme for **recap**:
			* `target` - (Optional, List) List of target or host name.
			* `ok` - (Optional, Float) Number of OK.
			* `changed` - (Optional, Float) Number of changed.
			* `failed` - (Optional, Float) Number of failed.
			* `skipped` - (Optional, Float) Number of skipped.
			* `unreachable` - (Optional, Float) Number of unreachable.
	* `system_job` - (Optional, List) System Job log summary. MaxItems: 1.
	Nested scheme for **system_job**:
		* `target_count` - (Optional, Float) number of targets or hosts.
		* `success` - (Optional, Float) Number of passed.
		* `failed` - (Optional, Float) Number of failed.
* `status` - (Optional, List) Job Status. MaxItems: 1.
Nested scheme for **status**:
	* `workspace_job_status` - (Optional, List) Workspace Job Status.
	Nested scheme for **workspace_job_status**:
		* `workspace_name` - (Optional, String) Workspace name.
		* `status_code` - (Optional, String) Status of Jobs.
		  * Constraints: Allowable values are: job_pending, job_in_progress, job_finished, job_failed, job_cancelled
		* `status_message` - (Optional, String) Workspace job status message (eg. App1_Setup_Pending, for a 'Setup' flow in the 'App1' Workspace).
		* `flow_status` - (Optional, List) Environment Flow JOB Status.MaxItems: 1.
		Nested scheme for **flow_status**:
			* `flow_id` - (Optional, String) flow id.
			* `flow_name` - (Optional, String) flow name.
			* `status_code` - (Optional, String) Status of Jobs.
			  * Constraints: Allowable values are: job_pending, job_in_progress, job_finished, job_failed, job_cancelled
			* `status_message` - (Optional, String) Flow Job status message - to be displayed along with the status_code;.
			* `workitems` - (Optional, List) Environment's individual workItem status details;.
			Nested scheme for **workitems**:
				* `workspace_id` - (Optional, String) Workspace id.
				* `workspace_name` - (Optional, String) workspace name.
				* `job_id` - (Optional, String) workspace job id.
				* `status_code` - (Optional, String) Status of Jobs.
				  * Constraints: Allowable values are: job_pending, job_in_progress, job_finished, job_failed, job_cancelled
				* `status_message` - (Optional, String) workitem job status message;.
				* `updated_at` - (Optional, String) workitem job status updation timestamp.
			* `updated_at` - (Optional, String) Job status updation timestamp.
		* `template_status` - (Optional, List) Workspace Flow Template job status.
		Nested scheme for **template_status**:
			* `template_id` - (Optional, String) Template Id.
			* `template_name` - (Optional, String) Template name.
			* `flow_index` - (Optional, Integer) Index of the template in the Flow.
			* `status_code` - (Optional, String) Status of Jobs.
			  * Constraints: Allowable values are: job_pending, job_in_progress, job_finished, job_failed, job_cancelled
			* `status_message` - (Optional, String) Template job status message (eg. VPCt1_Apply_Pending, for a 'VPCt1' Template).
			* `updated_at` - (Optional, String) Job status updation timestamp.
		* `updated_at` - (Optional, String) Job status updation timestamp.
	* `action_job_status` - (Optional, List) Action Job Status.
	Nested scheme for **action_job_status**:
		* `action_name` - (Optional, String) Action name.
		* `status_code` - (Optional, String) Status of Jobs.
		  * Constraints: Allowable values are: job_pending, job_in_progress, job_finished, job_failed, job_cancelled
		* `status_message` - (Optional, String) Action Job status message - to be displayed along with the action_status_code.
		* `bastion_status_code` - (Optional, String) Status of Resources.
		  * Constraints: Allowable values are: none, ready, processing, error
		* `bastion_status_message` - (Optional, String) Bastion status message - to be displayed along with the bastion_status_code;.
		* `targets_status_code` - (Optional, String) Status of Resources.
		  * Constraints: Allowable values are: none, ready, processing, error
		* `targets_status_message` - (Optional, String) Aggregated status message for all target resources,  to be displayed along with the targets_status_code;.
		* `updated_at` - (Optional, String) Job status updation timestamp.
	* `system_job_status` - (Optional, List) System Job Status. MaxItems: 1.
	Nested scheme for **system_job_status**:
		* `system_status_message` - (Optional, String) System job message.
		* `system_status_code` - (Optional, String) Status of Jobs.
		  * Constraints: Allowable values are: job_pending, job_in_progress, job_finished, job_failed, job_cancelled
		* `schematics_resource_status` - (Optional, List) job staus for each schematics resource.
		Nested scheme for **schematics_resource_status**:
			* `status_code` - (Optional, String) Status of Jobs.
			  * Constraints: Allowable values are: job_pending, job_in_progress, job_finished, job_failed, job_cancelled
			* `status_message` - (Optional, String) system job status message.
			* `schematics_resource_id` - (Optional, String) id for each resource which is targeted as a part of system job.
			* `updated_at` - (Optional, String) Job status updation timestamp.
		* `updated_at` - (Optional, String) Job status updation timestamp.
	* `flow_job_status` - (Optional, List) Environment Flow JOB Status. MaxItems: 1.
	Nested scheme for **flow_job_status**:
		* `flow_id` - (Optional, String) flow id.
		* `flow_name` - (Optional, String) flow name.
		* `status_code` - (Optional, String) Status of Jobs.
		  * Constraints: Allowable values are: job_pending, job_in_progress, job_finished, job_failed, job_cancelled
		* `status_message` - (Optional, String) Flow Job status message - to be displayed along with the status_code;.
		* `workitems` - (Optional, List) Environment's individual workItem status details;.
		Nested scheme for **workitems**:
			* `workspace_id` - (Optional, String) Workspace id.
			* `workspace_name` - (Optional, String) workspace name.
			* `job_id` - (Optional, String) workspace job id.
			* `status_code` - (Optional, String) Status of Jobs.
			  * Constraints: Allowable values are: job_pending, job_in_progress, job_finished, job_failed, job_cancelled
			* `status_message` - (Optional, String) workitem job status message;.
			* `updated_at` - (Optional, String) workitem job status updation timestamp.
		* `updated_at` - (Optional, String) Job status updation timestamp.
* `tags` - (Optional, List) User defined tags, while running the job.

## Attribute reference

In addition to all argument references listed, you can access the following attribute references after your resource is created.

* `id` - The unique identifier of the schematics_job.
* `description` - (Optional, String) The description of your job is derived from the related action or workspace.  The description can be up to 2048 characters long in size.
* `duration` - (Optional, String) Duration of job execution; example 40 sec.
* `end_at` - (String) Job end time.
* `log_store_url` - (Optional, String) Job log store URL.
* `name` - (Optional, String) Job name, uniquely derived from the related Workspace or Action.
* `resource_group` - (Optional, String) Resource-group name derived from the related Workspace or Action.
* `results_url` - (Optional, String) Job results store URL.
* `start_at` - (Optional, String) Job start time.
* `state_store_url` - (Optional, String) Job state store URL.
* `submitted_at` - (String) Job submission time.
* `submitted_by` - (String) Email address of user who submitted the job.
* `updated_at` - (String) Job status updation timestamp.

## Import

You can import the `ibm_schematics_job` resource by using `id`. Job ID.

# Syntax
```
$ terraform import ibm_schematics_job.schematics_job <id>
```
