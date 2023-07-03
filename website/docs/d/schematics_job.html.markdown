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

Review the argument reference that you can specify for your data source.

* `job_id` - (String) Job Id. Use `GET /v2/jobs` API to look up the Job Ids in your IBM Cloud account.

* `location` - (Optional,String) Location supported by IBM Cloud Schematics service.  While creating your workspace or action, choose the right region, since it cannot be changed.  Note, this does not limit the location of the IBM Cloud resources, provisioned using Schematics.
  * Constraints: Allowable values are: us-south, us-east, eu-gb, eu-de

## Attribute reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

* `id` - The unique identifier of the schematics_job.
* `bastion` - (List) Describes a bastion resource.
Nested scheme for **bastion**:
	* `name` - (String) Bastion Name(Unique).
	* `host` - (String) Reference to the Inventory resource definition.

* `command_name` - (String) Schematics job command name.
  * Constraints: Allowable values are: workspace_plan, workspace_apply, workspace_destroy, workspace_refresh, ansible_playbook_run, ansible_playbook_check, create_action, put_action, patch_action, delete_action, system_key_enable, system_key_delete, system_key_disable, system_key_rotate, system_key_restore, create_workspace, put_workspace, patch_workspace, delete_workspace, create_cart, create_environment, put_environment, delete_environment, environment_init, environment_install, environment_uninstall, repository_process

* `command_object` - (String) Name of the Schematics automation resource.
  * Constraints: Allowable values are: workspace, action, system, environment

* `command_object_id` - (String) Job command object id (workspace-id, action-id).

* `command_options` - (List) Command line options for the command.

* `command_parameter` - (String) Schematics job command parameter (playbook-name).

* `data` - (List) Job data.
Nested scheme for **data**:
	* `job_type` - (String) Type of Job.
	  * Constraints: Allowable values are: repo_download_job, workspace_job, action_job, system_job, flow-job
	* `workspace_job_data` - (List) Workspace Job data.
	Nested scheme for **workspace_job_data**:
		* `workspace_name` - (String) Workspace name.
		* `flow_id` - (String) Flow Id.
		* `flow_name` - (String) Flow name.
		* `inputs` - (List) Input variables data used by the Workspace Job.
		Nested scheme for **inputs**:
			* `name` - (String) Name of the variable.
			* `value` - (String) Value for the variable or reference to the value.
			* `metadata` - (List) User editable metadata for the variables.
			Nested scheme for **metadata**:
				* `type` - (String) Type of the variable.
				  * Constraints: Allowable values are: boolean, string, integer, date, array, list, map, complex
				* `aliases` - (List) List of aliases for the variable name.
				* `description` - (String) Description of the meta data.
				* `default_value` - (String) Default value for the variable, if the override value is not specified.
				* `secure` - (Boolean) Is the variable secure or sensitive ?.
				* `immutable` - (Boolean) Is the variable readonly ?.
				* `hidden` - (Boolean) If true, the variable will not be displayed on UI or CLI.
				* `options` - (List) List of possible values for this variable.  If type is integer or date, then the array of string will be  converted to array of integers or date during runtime.
				* `min_value` - (Integer) Minimum value of the variable. Applicable for integer type.
				* `max_value` - (Integer) Maximum value of the variable. Applicable for integer type.
				* `min_length` - (Integer) Minimum length of the variable value. Applicable for string type.
				* `max_length` - (Integer) Maximum length of the variable value. Applicable for string type.
				* `matches` - (String) Regex for the variable value.
				* `position` - (Integer) Relative position of this variable in a list.
				* `group_by` - (String) Display name of the group this variable belongs to.
				* `source` - (String) Source of this meta-data.
			* `link` - (String) Reference link to the variable value By default the expression will point to self.value.
		* `outputs` - (List) Output variables data from the Workspace Job.
		Nested scheme for **outputs**:
			* `name` - (String) Name of the variable.
			* `value` - (String) Value for the variable or reference to the value.
			* `metadata` - (List) User editable metadata for the variables.
			Nested scheme for **metadata**:
				* `type` - (String) Type of the variable.
				  * Constraints: Allowable values are: boolean, string, integer, date, array, list, map, complex
				* `aliases` - (List) List of aliases for the variable name.
				* `description` - (String) Description of the meta data.
				* `default_value` - (String) Default value for the variable, if the override value is not specified.
				* `secure` - (Boolean) Is the variable secure or sensitive ?.
				* `immutable` - (Boolean) Is the variable readonly ?.
				* `hidden` - (Boolean) If true, the variable will not be displayed on UI or CLI.
				* `options` - (List) List of possible values for this variable.  If type is integer or date, then the array of string will be  converted to array of integers or date during runtime.
				* `min_value` - (Integer) Minimum value of the variable. Applicable for integer type.
				* `max_value` - (Integer) Maximum value of the variable. Applicable for integer type.
				* `min_length` - (Integer) Minimum length of the variable value. Applicable for string type.
				* `max_length` - (Integer) Maximum length of the variable value. Applicable for string type.
				* `matches` - (String) Regex for the variable value.
				* `position` - (Integer) Relative position of this variable in a list.
				* `group_by` - (String) Display name of the group this variable belongs to.
				* `source` - (String) Source of this meta-data.
			* `link` - (String) Reference link to the variable value By default the expression will point to self.value.
		* `settings` - (List) Environment variables used by all the templates in the Workspace.
		Nested scheme for **settings**:
			* `name` - (String) Name of the variable.
			* `value` - (String) Value for the variable or reference to the value.
			* `metadata` - (List) User editable metadata for the variables.
			Nested scheme for **metadata**:
				* `type` - (String) Type of the variable.
				  * Constraints: Allowable values are: boolean, string, integer, date, array, list, map, complex
				* `aliases` - (List) List of aliases for the variable name.
				* `description` - (String) Description of the meta data.
				* `default_value` - (String) Default value for the variable, if the override value is not specified.
				* `secure` - (Boolean) Is the variable secure or sensitive ?.
				* `immutable` - (Boolean) Is the variable readonly ?.
				* `hidden` - (Boolean) If true, the variable will not be displayed on UI or CLI.
				* `options` - (List) List of possible values for this variable.  If type is integer or date, then the array of string will be  converted to array of integers or date during runtime.
				* `min_value` - (Integer) Minimum value of the variable. Applicable for integer type.
				* `max_value` - (Integer) Maximum value of the variable. Applicable for integer type.
				* `min_length` - (Integer) Minimum length of the variable value. Applicable for string type.
				* `max_length` - (Integer) Maximum length of the variable value. Applicable for string type.
				* `matches` - (String) Regex for the variable value.
				* `position` - (Integer) Relative position of this variable in a list.
				* `group_by` - (String) Display name of the group this variable belongs to.
				* `source` - (String) Source of this meta-data.
			* `link` - (String) Reference link to the variable value By default the expression will point to self.value.
		* `template_data` - (List) Input / output data of the Template in the Workspace Job.
		Nested scheme for **template_data**:
			* `template_id` - (String) Template Id.
			* `template_name` - (String) Template name.
			* `flow_index` - (Integer) Index of the template in the Flow.
			* `inputs` - (List) Job inputs used by the Templates.
			Nested scheme for **inputs**:
				* `name` - (String) Name of the variable.
				* `value` - (String) Value for the variable or reference to the value.
				* `metadata` - (List) User editable metadata for the variables.
				Nested scheme for **metadata**:
					* `type` - (String) Type of the variable.
					  * Constraints: Allowable values are: boolean, string, integer, date, array, list, map, complex
					* `aliases` - (List) List of aliases for the variable name.
					* `description` - (String) Description of the meta data.
					* `default_value` - (String) Default value for the variable, if the override value is not specified.
					* `secure` - (Boolean) Is the variable secure or sensitive ?.
					* `immutable` - (Boolean) Is the variable readonly ?.
					* `hidden` - (Boolean) If true, the variable will not be displayed on UI or CLI.
					* `options` - (List) List of possible values for this variable.  If type is integer or date, then the array of string will be  converted to array of integers or date during runtime.
					* `min_value` - (Integer) Minimum value of the variable. Applicable for integer type.
					* `max_value` - (Integer) Maximum value of the variable. Applicable for integer type.
					* `min_length` - (Integer) Minimum length of the variable value. Applicable for string type.
					* `max_length` - (Integer) Maximum length of the variable value. Applicable for string type.
					* `matches` - (String) Regex for the variable value.
					* `position` - (Integer) Relative position of this variable in a list.
					* `group_by` - (String) Display name of the group this variable belongs to.
					* `source` - (String) Source of this meta-data.
				* `link` - (String) Reference link to the variable value By default the expression will point to self.value.
			* `outputs` - (List) Job output from the Templates.
			Nested scheme for **outputs**:
				* `name` - (String) Name of the variable.
				* `value` - (String) Value for the variable or reference to the value.
				* `metadata` - (List) User editable metadata for the variables.
				Nested scheme for **metadata**:
					* `type` - (String) Type of the variable.
					  * Constraints: Allowable values are: boolean, string, integer, date, array, list, map, complex
					* `aliases` - (List) List of aliases for the variable name.
					* `description` - (String) Description of the meta data.
					* `default_value` - (String) Default value for the variable, if the override value is not specified.
					* `secure` - (Boolean) Is the variable secure or sensitive ?.
					* `immutable` - (Boolean) Is the variable readonly ?.
					* `hidden` - (Boolean) If true, the variable will not be displayed on UI or CLI.
					* `options` - (List) List of possible values for this variable.  If type is integer or date, then the array of string will be  converted to array of integers or date during runtime.
					* `min_value` - (Integer) Minimum value of the variable. Applicable for integer type.
					* `max_value` - (Integer) Maximum value of the variable. Applicable for integer type.
					* `min_length` - (Integer) Minimum length of the variable value. Applicable for string type.
					* `max_length` - (Integer) Maximum length of the variable value. Applicable for string type.
					* `matches` - (String) Regex for the variable value.
					* `position` - (Integer) Relative position of this variable in a list.
					* `group_by` - (String) Display name of the group this variable belongs to.
					* `source` - (String) Source of this meta-data.
				* `link` - (String) Reference link to the variable value By default the expression will point to self.value.
			* `settings` - (List) Environment variables used by the template.
			Nested scheme for **settings**:
				* `name` - (String) Name of the variable.
				* `value` - (String) Value for the variable or reference to the value.
				* `metadata` - (List) User editable metadata for the variables.
				Nested scheme for **metadata**:
					* `type` - (String) Type of the variable.
					  * Constraints: Allowable values are: boolean, string, integer, date, array, list, map, complex
					* `aliases` - (List) List of aliases for the variable name.
					* `description` - (String) Description of the meta data.
					* `default_value` - (String) Default value for the variable, if the override value is not specified.
					* `secure` - (Boolean) Is the variable secure or sensitive ?.
					* `immutable` - (Boolean) Is the variable readonly ?.
					* `hidden` - (Boolean) If true, the variable will not be displayed on UI or CLI.
					* `options` - (List) List of possible values for this variable.  If type is integer or date, then the array of string will be  converted to array of integers or date during runtime.
					* `min_value` - (Integer) Minimum value of the variable. Applicable for integer type.
					* `max_value` - (Integer) Maximum value of the variable. Applicable for integer type.
					* `min_length` - (Integer) Minimum length of the variable value. Applicable for string type.
					* `max_length` - (Integer) Maximum length of the variable value. Applicable for string type.
					* `matches` - (String) Regex for the variable value.
					* `position` - (Integer) Relative position of this variable in a list.
					* `group_by` - (String) Display name of the group this variable belongs to.
					* `source` - (String) Source of this meta-data.
				* `link` - (String) Reference link to the variable value By default the expression will point to self.value.
			* `updated_at` - (String) Job status updation timestamp.
		* `updated_at` - (String) Job status updation timestamp.
	* `action_job_data` - (List) Action Job data.
	Nested scheme for **action_job_data**:
		* `action_name` - (String) Flow name.
		* `inputs` - (List) Input variables data used by the Action Job.
		Nested scheme for **inputs**:
			* `name` - (String) Name of the variable.
			* `value` - (String) Value for the variable or reference to the value.
			* `metadata` - (List) User editable metadata for the variables.
			Nested scheme for **metadata**:
				* `type` - (String) Type of the variable.
				  * Constraints: Allowable values are: boolean, string, integer, date, array, list, map, complex
				* `aliases` - (List) List of aliases for the variable name.
				* `description` - (String) Description of the meta data.
				* `default_value` - (String) Default value for the variable, if the override value is not specified.
				* `secure` - (Boolean) Is the variable secure or sensitive ?.
				* `immutable` - (Boolean) Is the variable readonly ?.
				* `hidden` - (Boolean) If true, the variable will not be displayed on UI or CLI.
				* `options` - (List) List of possible values for this variable.  If type is integer or date, then the array of string will be  converted to array of integers or date during runtime.
				* `min_value` - (Integer) Minimum value of the variable. Applicable for integer type.
				* `max_value` - (Integer) Maximum value of the variable. Applicable for integer type.
				* `min_length` - (Integer) Minimum length of the variable value. Applicable for string type.
				* `max_length` - (Integer) Maximum length of the variable value. Applicable for string type.
				* `matches` - (String) Regex for the variable value.
				* `position` - (Integer) Relative position of this variable in a list.
				* `group_by` - (String) Display name of the group this variable belongs to.
				* `source` - (String) Source of this meta-data.
			* `link` - (String) Reference link to the variable value By default the expression will point to self.value.
		* `outputs` - (List) Output variables data from the Action Job.
		Nested scheme for **outputs**:
			* `name` - (String) Name of the variable.
			* `value` - (String) Value for the variable or reference to the value.
			* `metadata` - (List) User editable metadata for the variables.
			Nested scheme for **metadata**:
				* `type` - (String) Type of the variable.
				  * Constraints: Allowable values are: boolean, string, integer, date, array, list, map, complex
				* `aliases` - (List) List of aliases for the variable name.
				* `description` - (String) Description of the meta data.
				* `default_value` - (String) Default value for the variable, if the override value is not specified.
				* `secure` - (Boolean) Is the variable secure or sensitive ?.
				* `immutable` - (Boolean) Is the variable readonly ?.
				* `hidden` - (Boolean) If true, the variable will not be displayed on UI or CLI.
				* `options` - (List) List of possible values for this variable.  If type is integer or date, then the array of string will be  converted to array of integers or date during runtime.
				* `min_value` - (Integer) Minimum value of the variable. Applicable for integer type.
				* `max_value` - (Integer) Maximum value of the variable. Applicable for integer type.
				* `min_length` - (Integer) Minimum length of the variable value. Applicable for string type.
				* `max_length` - (Integer) Maximum length of the variable value. Applicable for string type.
				* `matches` - (String) Regex for the variable value.
				* `position` - (Integer) Relative position of this variable in a list.
				* `group_by` - (String) Display name of the group this variable belongs to.
				* `source` - (String) Source of this meta-data.
			* `link` - (String) Reference link to the variable value By default the expression will point to self.value.
		* `settings` - (List) Environment variables used by all the templates in the Action.
		Nested scheme for **settings**:
			* `name` - (String) Name of the variable.
			* `value` - (String) Value for the variable or reference to the value.
			* `metadata` - (List) User editable metadata for the variables.
			Nested scheme for **metadata**:
				* `type` - (String) Type of the variable.
				  * Constraints: Allowable values are: boolean, string, integer, date, array, list, map, complex
				* `aliases` - (List) List of aliases for the variable name.
				* `description` - (String) Description of the meta data.
				* `default_value` - (String) Default value for the variable, if the override value is not specified.
				* `secure` - (Boolean) Is the variable secure or sensitive ?.
				* `immutable` - (Boolean) Is the variable readonly ?.
				* `hidden` - (Boolean) If true, the variable will not be displayed on UI or CLI.
				* `options` - (List) List of possible values for this variable.  If type is integer or date, then the array of string will be  converted to array of integers or date during runtime.
				* `min_value` - (Integer) Minimum value of the variable. Applicable for integer type.
				* `max_value` - (Integer) Maximum value of the variable. Applicable for integer type.
				* `min_length` - (Integer) Minimum length of the variable value. Applicable for string type.
				* `max_length` - (Integer) Maximum length of the variable value. Applicable for string type.
				* `matches` - (String) Regex for the variable value.
				* `position` - (Integer) Relative position of this variable in a list.
				* `group_by` - (String) Display name of the group this variable belongs to.
				* `source` - (String) Source of this meta-data.
			* `link` - (String) Reference link to the variable value By default the expression will point to self.value.
		* `updated_at` - (String) Job status updation timestamp.
		* `inventory_record` - (List) Complete inventory resource details with user inputs and system generated data.
		Nested scheme for **inventory_record**:
			* `name` - (String) The unique name of your Inventory.  The name can be up to 128 characters long and can include alphanumeric  characters, spaces, dashes, and underscores.
			* `id` - (String) Inventory id.
			* `description` - (String) The description of your Inventory.  The description can be up to 2048 characters long in size.
			* `location` - (String) List of locations supported by IBM Cloud Schematics service.  While creating your workspace or action, choose the right region, since it cannot be changed.  Note, this does not limit the location of the IBM Cloud resources, provisioned using Schematics.
			  * Constraints: Allowable values are: us-south, us-east, eu-gb, eu-de
			* `resource_group` - (String) Resource-group name for the Inventory definition.  By default, Inventory will be created in Default Resource Group.
			* `created_at` - (String) Inventory creation time.
			* `created_by` - (String) Email address of user who created the Inventory.
			* `updated_at` - (String) Inventory updation time.
			* `updated_by` - (String) Email address of user who updated the Inventory.
			* `inventories_ini` - (String) Input inventory of host and host group for the playbook,  in the .ini file format.
			* `resource_queries` - (List) Input resource queries that is used to dynamically generate  the inventory of host and host group for the playbook.
		* `materialized_inventory` - (String) Materialized inventory details used by the Action Job, in .ini format.
	* `system_job_data` - (List) Controls Job data.
	Nested scheme for **system_job_data**:
		* `key_id` - (String) Key ID for which key event is generated.
		* `schematics_resource_id` - (List) List of the schematics resource id.
		* `updated_at` - (String) Job status updation timestamp.
	* `flow_job_data` - (List) Flow Job data.
	Nested scheme for **flow_job_data**:
		* `flow_id` - (String) Flow ID.
		* `flow_name` - (String) Flow Name.
		* `workitems` - (List) Job data used by each workitem Job.
		Nested scheme for **workitems**:
			* `command_object_id` - (String) command object id.
			* `command_object_name` - (String) command object name.
			* `layers` - (String) layer name.
			* `source_type` - (String) Type of source for the Template.
			  * Constraints: Allowable values are: local, git_hub, git_hub_enterprise, git_lab, ibm_git_lab, ibm_cloud_catalog, external_scm, cos_bucket
			* `source` - (List) Source of templates, playbooks, or controls.
			Nested scheme for **source**:
				* `source_type` - (String) Type of source for the Template.
				  * Constraints: Allowable values are: local, git_hub, git_hub_enterprise, git_lab, ibm_git_lab, ibm_cloud_catalog, external_scm, cos_bucket
				* `git` - (List) Connection details to Git source.
				Nested scheme for **git**:
					* `computed_git_repo_url` - (String) The Complete URL which is computed by git_repo_url, git_repo_folder and branch.
					* `git_repo_url` - (String) URL to the GIT Repo that can be used to clone the template.
					* `git_token` - (String) Personal Access Token to connect to Git URLs.
					* `git_repo_folder` - (String) Name of the folder in the Git Repo, that contains the template.
					* `git_release` - (String) Name of the release tag, used to fetch the Git Repo.
					* `git_branch` - (String) Name of the branch, used to fetch the Git Repo.
				* `catalog` - (List) Connection details to IBM Cloud Catalog source.
				Nested scheme for **catalog**:
					* `catalog_name` - (String) name of the private catalog.
					* `offering_name` - (String) Name of the offering in the IBM Catalog.
					* `offering_version` - (String) Version string of the offering in the IBM Catalog.
					* `offering_kind` - (String) Type of the offering, in the IBM Catalog.
					* `offering_id` - (String) Id of the offering the IBM Catalog.
					* `offering_version_id` - (String) Id of the offering version the IBM Catalog.
					* `offering_repo_url` - (String) Repo Url of the offering, in the IBM Catalog.
			* `inputs` - (List) Input variables data for the workItem used in FlowJob.
			Nested scheme for **inputs**:
				* `name` - (String) Name of the variable.
				* `value` - (String) Value for the variable or reference to the value.
				* `metadata` - (List) User editable metadata for the variables.
				Nested scheme for **metadata**:
					* `type` - (String) Type of the variable.
					  * Constraints: Allowable values are: boolean, string, integer, date, array, list, map, complex
					* `aliases` - (List) List of aliases for the variable name.
					* `description` - (String) Description of the meta data.
					* `default_value` - (String) Default value for the variable, if the override value is not specified.
					* `secure` - (Boolean) Is the variable secure or sensitive ?.
					* `immutable` - (Boolean) Is the variable readonly ?.
					* `hidden` - (Boolean) If true, the variable will not be displayed on UI or CLI.
					* `options` - (List) List of possible values for this variable.  If type is integer or date, then the array of string will be  converted to array of integers or date during runtime.
					* `min_value` - (Integer) Minimum value of the variable. Applicable for integer type.
					* `max_value` - (Integer) Maximum value of the variable. Applicable for integer type.
					* `min_length` - (Integer) Minimum length of the variable value. Applicable for string type.
					* `max_length` - (Integer) Maximum length of the variable value. Applicable for string type.
					* `matches` - (String) Regex for the variable value.
					* `position` - (Integer) Relative position of this variable in a list.
					* `group_by` - (String) Display name of the group this variable belongs to.
					* `source` - (String) Source of this meta-data.
				* `link` - (String) Reference link to the variable value By default the expression will point to self.value.
			* `outputs` - (List) Output variables for the workItem.
			Nested scheme for **outputs**:
				* `name` - (String) Name of the variable.
				* `value` - (String) Value for the variable or reference to the value.
				* `metadata` - (List) User editable metadata for the variables.
				Nested scheme for **metadata**:
					* `type` - (String) Type of the variable.
					  * Constraints: Allowable values are: boolean, string, integer, date, array, list, map, complex
					* `aliases` - (List) List of aliases for the variable name.
					* `description` - (String) Description of the meta data.
					* `default_value` - (String) Default value for the variable, if the override value is not specified.
					* `secure` - (Boolean) Is the variable secure or sensitive ?.
					* `immutable` - (Boolean) Is the variable readonly ?.
					* `hidden` - (Boolean) If true, the variable will not be displayed on UI or CLI.
					* `options` - (List) List of possible values for this variable.  If type is integer or date, then the array of string will be  converted to array of integers or date during runtime.
					* `min_value` - (Integer) Minimum value of the variable. Applicable for integer type.
					* `max_value` - (Integer) Maximum value of the variable. Applicable for integer type.
					* `min_length` - (Integer) Minimum length of the variable value. Applicable for string type.
					* `max_length` - (Integer) Maximum length of the variable value. Applicable for string type.
					* `matches` - (String) Regex for the variable value.
					* `position` - (Integer) Relative position of this variable in a list.
					* `group_by` - (String) Display name of the group this variable belongs to.
					* `source` - (String) Source of this meta-data.
				* `link` - (String) Reference link to the variable value By default the expression will point to self.value.
			* `settings` - (List) Environment variables for the workItem.
			Nested scheme for **settings**:
				* `name` - (String) Name of the variable.
				* `value` - (String) Value for the variable or reference to the value.
				* `metadata` - (List) User editable metadata for the variables.
				Nested scheme for **metadata**:
					* `type` - (String) Type of the variable.
					  * Constraints: Allowable values are: boolean, string, integer, date, array, list, map, complex
					* `aliases` - (List) List of aliases for the variable name.
					* `description` - (String) Description of the meta data.
					* `default_value` - (String) Default value for the variable, if the override value is not specified.
					* `secure` - (Boolean) Is the variable secure or sensitive ?.
					* `immutable` - (Boolean) Is the variable readonly ?.
					* `hidden` - (Boolean) If true, the variable will not be displayed on UI or CLI.
					* `options` - (List) List of possible values for this variable.  If type is integer or date, then the array of string will be  converted to array of integers or date during runtime.
					* `min_value` - (Integer) Minimum value of the variable. Applicable for integer type.
					* `max_value` - (Integer) Maximum value of the variable. Applicable for integer type.
					* `min_length` - (Integer) Minimum length of the variable value. Applicable for string type.
					* `max_length` - (Integer) Maximum length of the variable value. Applicable for string type.
					* `matches` - (String) Regex for the variable value.
					* `position` - (Integer) Relative position of this variable in a list.
					* `group_by` - (String) Display name of the group this variable belongs to.
					* `source` - (String) Source of this meta-data.
				* `link` - (String) Reference link to the variable value By default the expression will point to self.value.
			* `last_job` - (List) Status of the last job executed by the workitem.
			Nested scheme for **last_job**:
				* `command_object` - (String) Name of the Schematics automation resource.
				  * Constraints: Allowable values are: workspace, action, system, environment
				* `command_object_name` - (String) command object name (workspace_name/action_name).
				* `command_object_id` - (String) Workitem command object id, maps to workspace_id or action_id.
				* `command_name` - (String) Schematics job command name.
				  * Constraints: Allowable values are: workspace_plan, workspace_apply, workspace_destroy, workspace_refresh, ansible_playbook_run, ansible_playbook_check, create_action, put_action, patch_action, delete_action, system_key_enable, system_key_delete, system_key_disable, system_key_rotate, system_key_restore, create_workspace, put_workspace, patch_workspace, delete_workspace, create_cart, create_environment, put_environment, delete_environment, environment_init, environment_install, environment_uninstall, repository_process
				* `job_id` - (String) Workspace job id.
				* `job_status` - (String) Status of Jobs.
				  * Constraints: Allowable values are: job_pending, job_in_progress, job_finished, job_failed, job_cancelled
			* `updated_at` - (String) Job status updation timestamp.
		* `updated_at` - (String) Job status updation timestamp.

* `description` - (String) The description of your job is derived from the related action or workspace.  The description can be up to 2048 characters long in size.

* `duration` - (String) Duration of job execution; example 40 sec.

* `end_at` - (String) Job end time.

* `id` - (String) Job ID.

* `job_env_settings` - (List) Environment variables used by the Job while performing Action or Workspace.
Nested scheme for **job_env_settings**:
	* `name` - (String) Name of the variable.
	* `value` - (String) Value for the variable or reference to the value.
	* `metadata` - (List) User editable metadata for the variables.
	Nested scheme for **metadata**:
		* `type` - (String) Type of the variable.
		  * Constraints: Allowable values are: boolean, string, integer, date, array, list, map, complex
		* `aliases` - (List) List of aliases for the variable name.
		* `description` - (String) Description of the meta data.
		* `default_value` - (String) Default value for the variable, if the override value is not specified.
		* `secure` - (Boolean) Is the variable secure or sensitive ?.
		* `immutable` - (Boolean) Is the variable readonly ?.
		* `hidden` - (Boolean) If true, the variable will not be displayed on UI or CLI.
		* `options` - (List) List of possible values for this variable.  If type is integer or date, then the array of string will be  converted to array of integers or date during runtime.
		* `min_value` - (Integer) Minimum value of the variable. Applicable for integer type.
		* `max_value` - (Integer) Maximum value of the variable. Applicable for integer type.
		* `min_length` - (Integer) Minimum length of the variable value. Applicable for string type.
		* `max_length` - (Integer) Maximum length of the variable value. Applicable for string type.
		* `matches` - (String) Regex for the variable value.
		* `position` - (Integer) Relative position of this variable in a list.
		* `group_by` - (String) Display name of the group this variable belongs to.
		* `source` - (String) Source of this meta-data.
	* `link` - (String) Reference link to the variable value By default the expression will point to self.value.

* `job_inputs` - (List) Job inputs used by Action or Workspace.
Nested scheme for **job_inputs**:
	* `name` - (String) Name of the variable.
	* `value` - (String) Value for the variable or reference to the value.
	* `metadata` - (List) User editable metadata for the variables.
	Nested scheme for **metadata**:
		* `type` - (String) Type of the variable.
		  * Constraints: Allowable values are: boolean, string, integer, date, array, list, map, complex
		* `aliases` - (List) List of aliases for the variable name.
		* `description` - (String) Description of the meta data.
		* `default_value` - (String) Default value for the variable, if the override value is not specified.
		* `secure` - (Boolean) Is the variable secure or sensitive ?.
		* `immutable` - (Boolean) Is the variable readonly ?.
		* `hidden` - (Boolean) If true, the variable will not be displayed on UI or CLI.
		* `options` - (List) List of possible values for this variable.  If type is integer or date, then the array of string will be  converted to array of integers or date during runtime.
		* `min_value` - (Integer) Minimum value of the variable. Applicable for integer type.
		* `max_value` - (Integer) Maximum value of the variable. Applicable for integer type.
		* `min_length` - (Integer) Minimum length of the variable value. Applicable for string type.
		* `max_length` - (Integer) Maximum length of the variable value. Applicable for string type.
		* `matches` - (String) Regex for the variable value.
		* `position` - (Integer) Relative position of this variable in a list.
		* `group_by` - (String) Display name of the group this variable belongs to.
		* `source` - (String) Source of this meta-data.
	* `link` - (String) Reference link to the variable value By default the expression will point to self.value.


* `log_store_url` - (String) Job log store URL.

* `log_summary` - (List) Job log summary record.
Nested scheme for **log_summary**:
	* `job_id` - (String) Workspace Id.
	* `job_type` - (String) Type of Job.
	  * Constraints: Allowable values are: repo_download_job, workspace_job, action_job, system_job, flow_job
	* `log_start_at` - (String) Job log start timestamp.
	* `log_analyzed_till` - (String) Job log update timestamp.
	* `elapsed_time` - (Float) Job log elapsed time (log_analyzed_till - log_start_at).
	* `log_errors` - (List) Job log errors.
	Nested scheme for **log_errors**:
		* `error_code` - (String) Error code in the Log.
		* `error_msg` - (String) Summary error message in the log.
		* `error_count` - (Float) Number of occurrence.
	* `repo_download_job` - (List) Repo download Job log summary.
	Nested scheme for **repo_download_job**:
		* `scanned_file_count` - (Float) Number of files scanned.
		* `quarantined_file_count` - (Float) Number of files quarantined.
		* `detected_filetype` - (String) Detected template or data file type.
		* `inputs_count` - (String) Number of inputs detected.
		* `outputs_count` - (String) Number of outputs detected.
	* `workspace_job` - (List) Workspace Job log summary.
	Nested scheme for **workspace_job**:
		* `resources_add` - (Float) Number of resources add.
		* `resources_modify` - (Float) Number of resources modify.
		* `resources_destroy` - (Float) Number of resources destroy.
	* `flow_job` - (List) Flow Job log summary.
	Nested scheme for **flow_job**:
		* `workitems_completed` - (Float) Number of workitems completed successfully.
		* `workitems_pending` - (Float) Number of workitems pending in the flow.
		* `workitems_failed` - (Float) Number of workitems failed.
		* `workitems` - (List)
		Nested scheme for **workitems**:
			* `workspace_id` - (String) workspace ID.
			* `job_id` - (String) workspace JOB ID.
			* `resources_add` - (Float) Number of resources add.
			* `resources_modify` - (Float) Number of resources modify.
			* `resources_destroy` - (Float) Number of resources destroy.
			* `log_url` - (String) Log url for job.
	* `action_job` - (List) Flow Job log summary.
	Nested scheme for **action_job**:
		* `target_count` - (Float) number of targets or hosts.
		* `task_count` - (Float) number of tasks in playbook.
		* `play_count` - (Float) number of plays in playbook.
		* `recap` - (List) Recap records.
		Nested scheme for **recap**:
			* `target` - (List) List of target or host name.
			* `ok` - (Float) Number of OK.
			* `changed` - (Float) Number of changed.
			* `failed` - (Float) Number of failed.
			* `skipped` - (Float) Number of skipped.
			* `unreachable` - (Float) Number of unreachable.
	* `system_job` - (List) System Job log summary.
	Nested scheme for **system_job**:
		* `target_count` - (Float) number of targets or hosts.
		* `success` - (Float) Number of passed.
		* `failed` - (Float) Number of failed.

* `name` - (String) Job name, uniquely derived from the related Workspace or Action.

* `resource_group` - (String) Resource-group name derived from the related Workspace or Action.

* `results_url` - (String) Job results store URL.

* `start_at` - (String) Job start time.

* `state_store_url` - (String) Job state store URL.

* `status` - (List) Job Status.
Nested scheme for **status**:
	* `workspace_job_status` - (List) Workspace Job Status.
	Nested scheme for **workspace_job_status**:
		* `workspace_name` - (String) Workspace name.
		* `status_code` - (String) Status of Jobs.
		  * Constraints: Allowable values are: job_pending, job_in_progress, job_finished, job_failed, job_cancelled
		* `status_message` - (String) Workspace job status message (eg. App1_Setup_Pending, for a 'Setup' flow in the 'App1' Workspace).
		* `flow_status` - (List) Environment Flow JOB Status.
		Nested scheme for **flow_status**:
			* `flow_id` - (String) flow id.
			* `flow_name` - (String) flow name.
			* `status_code` - (String) Status of Jobs.
			  * Constraints: Allowable values are: job_pending, job_in_progress, job_finished, job_failed, job_cancelled
			* `status_message` - (String) Flow Job status message - to be displayed along with the status_code;.
			* `workitems` - (List) Environment's individual workItem status details;.
			Nested scheme for **workitems**:
				* `workspace_id` - (String) Workspace id.
				* `workspace_name` - (String) workspace name.
				* `job_id` - (String) workspace job id.
				* `status_code` - (String) Status of Jobs.
				  * Constraints: Allowable values are: job_pending, job_in_progress, job_finished, job_failed, job_cancelled
				* `status_message` - (String) workitem job status message;.
				* `updated_at` - (String) workitem job status updation timestamp.
			* `updated_at` - (String) Job status updation timestamp.
		* `template_status` - (List) Workspace Flow Template job status.
		Nested scheme for **template_status**:
			* `template_id` - (String) Template Id.
			* `template_name` - (String) Template name.
			* `flow_index` - (Integer) Index of the template in the Flow.
			* `status_code` - (String) Status of Jobs.
			  * Constraints: Allowable values are: job_pending, job_in_progress, job_finished, job_failed, job_cancelled
			* `status_message` - (String) Template job status message (eg. VPCt1_Apply_Pending, for a 'VPCt1' Template).
			* `updated_at` - (String) Job status updation timestamp.
		* `updated_at` - (String) Job status updation timestamp.
	* `action_job_status` - (List) Action Job Status.
	Nested scheme for **action_job_status**:
		* `action_name` - (String) Action name.
		* `status_code` - (String) Status of Jobs.
		  * Constraints: Allowable values are: job_pending, job_in_progress, job_finished, job_failed, job_cancelled
		* `status_message` - (String) Action Job status message - to be displayed along with the action_status_code.
		* `bastion_status_code` - (String) Status of Resources.
		  * Constraints: Allowable values are: none, ready, processing, error
		* `bastion_status_message` - (String) Bastion status message - to be displayed along with the bastion_status_code;.
		* `targets_status_code` - (String) Status of Resources.
		  * Constraints: Allowable values are: none, ready, processing, error
		* `targets_status_message` - (String) Aggregated status message for all target resources,  to be displayed along with the targets_status_code;.
		* `updated_at` - (String) Job status updation timestamp.
	* `system_job_status` - (List) System Job Status.
	Nested scheme for **system_job_status**:
		* `system_status_message` - (String) System job message.
		* `system_status_code` - (String) Status of Jobs.
		  * Constraints: Allowable values are: job_pending, job_in_progress, job_finished, job_failed, job_cancelled
		* `schematics_resource_status` - (List) job staus for each schematics resource.
		Nested scheme for **schematics_resource_status**:
			* `status_code` - (String) Status of Jobs.
			  * Constraints: Allowable values are: job_pending, job_in_progress, job_finished, job_failed, job_cancelled
			* `status_message` - (String) system job status message.
			* `schematics_resource_id` - (String) id for each resource which is targeted as a part of system job.
			* `updated_at` - (String) Job status updation timestamp.
		* `updated_at` - (String) Job status updation timestamp.
	* `flow_job_status` - (List) Environment Flow JOB Status.
	Nested scheme for **flow_job_status**:
		* `flow_id` - (String) flow id.
		* `flow_name` - (String) flow name.
		* `status_code` - (String) Status of Jobs.
		  * Constraints: Allowable values are: job_pending, job_in_progress, job_finished, job_failed, job_cancelled
		* `status_message` - (String) Flow Job status message - to be displayed along with the status_code;.
		* `workitems` - (List) Environment's individual workItem status details;.
		Nested scheme for **workitems**:
			* `workspace_id` - (String) Workspace id.
			* `workspace_name` - (String) workspace name.
			* `job_id` - (String) workspace job id.
			* `status_code` - (String) Status of Jobs.
			  * Constraints: Allowable values are: job_pending, job_in_progress, job_finished, job_failed, job_cancelled
			* `status_message` - (String) workitem job status message;.
			* `updated_at` - (String) workitem job status updation timestamp.
		* `updated_at` - (String) Job status updation timestamp.

* `submitted_at` - (String) Job submission time.

* `submitted_by` - (String) Email address of user who submitted the job.

* `tags` - (List) User defined tags, while running the job.

* `updated_at` - (String) Job status updation timestamp.

