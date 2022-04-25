---
subcategory: "Schematics"
layout: "ibm"
page_title: "IBM : ibm_schematics_job"
sidebar_current: "docs-ibm-datasource-schematics-job"
description: |-
  Get information about schematics_job
subcategory: "Schematics Service API"
---

# ibm_schematics_job

Provides a read-only data source for schematics_job. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "ibm_schematics_job" "schematics_job" {
	job_id = "job_id"
}
```

## Argument Reference

Review the argument reference that you can specify for your data source.

* `job_id` - (Required, Forces new resource, String) Job Id. Use `GET /v2/jobs` API to look up the Job Ids in your IBM Cloud account.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

* `id` - The unique identifier of the schematics_job.
* `bastion` - (Optional, List) Describes a bastion resource.
Nested scheme for **bastion**:
	* `host` - (Optional, String) Reference to the Inventory resource definition.
	* `name` - (Optional, String) Bastion Name(Unique).

* `command_name` - (Optional, String) Schematics job command name.
  * Constraints: Allowable values are: `workspace_plan`, `workspace_apply`, `workspace_destroy`, `workspace_refresh`, `ansible_playbook_run`, `ansible_playbook_check`, `create_action`, `put_action`, `patch_action`, `delete_action`, `system_key_enable`, `system_key_delete`, `system_key_disable`, `system_key_rotate`, `system_key_restore`, `create_workspace`, `put_workspace`, `patch_workspace`, `delete_workspace`, `create_cart`, `create_environment`, `put_environment`, `delete_environment`, `environment_init`, `environment_install`, `environment_uninstall`, `repository_process`, `terraform_commands`.

* `command_object` - (Optional, String) Name of the Schematics automation resource.
  * Constraints: Allowable values are: `workspace`, `action`, `system`, `environment`.

* `command_object_id` - (Optional, String) Job command object id (workspace-id, action-id).

* `command_options` - (Optional, List) Command line options for the command.

* `command_parameter` - (Optional, String) Schematics job command parameter (playbook-name).

* `data` - (Optional, List) Job data.
Nested scheme for **data**:
	* `action_job_data` - (Optional, List) Action Job data.
	Nested scheme for **action_job_data**:
		* `action_name` - (Optional, String) Flow name.
		* `inputs` - (Optional, List) Input variables data used by the Action Job.
		Nested scheme for **inputs**:
			* `link` - (Optional, String) The reference link to the variable value By default the expression points to `$self.value`.
			* `metadata` - (Optional, List) An user editable metadata for the variables.
			Nested scheme for **metadata**:
				* `aliases` - (Optional, List) The list of aliases for the variable name.
				* `cloud_data_type` - (Optional, String) Cloud data type of the variable. eg. resource_group_id, region, vpc_id.
				* `default_value` - (Optional, String) Default value for the variable only if the override value is not specified.
				* `description` - (Optional, String) The description of the meta data.
				* `group_by` - (Optional, String) The display name of the group this variable belongs to.
				* `hidden` - (Optional, Boolean) If **true**, the variable is not displayed on UI or Command line.
				* `immutable` - (Optional, Boolean) Is the variable readonly ?.
				* `link_status` - (Optional, String) The status of the link.
				  * Constraints: Allowable values are: `normal`, `broken`.
				* `matches` - (Optional, String) The regex for the variable value.
				* `max_length` - (Optional, Integer) The maximum length of the variable value. Applicable for the string type.
				* `max_value` - (Optional, Integer) The maximum value of the variable. Applicable for the integer type.
				* `min_length` - (Optional, Integer) The minimum length of the variable value. Applicable for the string type.
				* `min_value` - (Optional, Integer) The minimum value of the variable. Applicable for the integer type.
				* `options` - (Optional, List) The list of possible values for this variable.  If type is **integer** or **date**, then the array of string is  converted to array of integers or date during the runtime.
				* `position` - (Optional, Integer) The relative position of this variable in a list.
				* `required` - (Optional, Boolean) If the variable required?.
				* `secure` - (Optional, Boolean) Is the variable secure or sensitive ?.
				* `source` - (Optional, String) The source of this meta-data.
				* `type` - (Optional, String) Type of the variable.
				  * Constraints: Allowable values are: `boolean`, `string`, `integer`, `date`, `array`, `list`, `map`, `complex`, `link`.
			* `name` - (Optional, String) The name of the variable. For example, `name = "inventory username"`.
			* `use_default` - (Optional, Boolean) True, will ignore the data in the value attribute, instead the data in metadata.default_value will be used.
			* `value` - (Optional, String) The value for the variable or reference to the value. For example, `value = "<provide your ssh_key_value with \n>"`. **Note** The SSH key should contain `\n` at the end of the key details in case of command line or API calls.
		* `inventory_record` - (Optional, List) Complete inventory resource details with user inputs and system generated data.
		Nested scheme for **inventory_record**:
			* `created_at` - (Optional, String) Inventory creation time.
			* `created_by` - (Optional, String) Email address of user who created the Inventory.
			* `description` - (Optional, String) The description of your Inventory.  The description can be up to 2048 characters long in size.
			* `id` - (Optional, String) Inventory id.
			* `inventories_ini` - (Optional, String) Input inventory of host and host group for the playbook,  in the .ini file format.
			* `location` - (Optional, String) List of locations supported by IBM Cloud Schematics service.  While creating your workspace or action, choose the right region, since it cannot be changed.  Note, this does not limit the location of the IBM Cloud resources, provisioned using Schematics.
			  * Constraints: Allowable values are: `us-south`, `us-east`, `eu-gb`, `eu-de`.
			* `name` - (Optional, String) The unique name of your Inventory.  The name can be up to 128 characters long and can include alphanumeric  characters, spaces, dashes, and underscores.
			* `resource_group` - (Optional, String) Resource-group name for the Inventory definition.  By default, Inventory will be created in Default Resource Group.
			* `resource_queries` - (Optional, List) Input resource queries that is used to dynamically generate  the inventory of host and host group for the playbook.
			* `updated_at` - (Optional, String) Inventory updation time.
			* `updated_by` - (Optional, String) Email address of user who updated the Inventory.
		* `materialized_inventory` - (Optional, String) Materialized inventory details used by the Action Job, in .ini format.
		* `outputs` - (Optional, List) Output variables data from the Action Job.
		Nested scheme for **outputs**:
			* `link` - (Optional, String) The reference link to the variable value By default the expression points to `$self.value`.
			* `metadata` - (Optional, List) An user editable metadata for the variables.
			Nested scheme for **metadata**:
				* `aliases` - (Optional, List) The list of aliases for the variable name.
				* `cloud_data_type` - (Optional, String) Cloud data type of the variable. eg. resource_group_id, region, vpc_id.
				* `default_value` - (Optional, String) Default value for the variable only if the override value is not specified.
				* `description` - (Optional, String) The description of the meta data.
				* `group_by` - (Optional, String) The display name of the group this variable belongs to.
				* `hidden` - (Optional, Boolean) If **true**, the variable is not displayed on UI or Command line.
				* `immutable` - (Optional, Boolean) Is the variable readonly ?.
				* `link_status` - (Optional, String) The status of the link.
				  * Constraints: Allowable values are: `normal`, `broken`.
				* `matches` - (Optional, String) The regex for the variable value.
				* `max_length` - (Optional, Integer) The maximum length of the variable value. Applicable for the string type.
				* `max_value` - (Optional, Integer) The maximum value of the variable. Applicable for the integer type.
				* `min_length` - (Optional, Integer) The minimum length of the variable value. Applicable for the string type.
				* `min_value` - (Optional, Integer) The minimum value of the variable. Applicable for the integer type.
				* `options` - (Optional, List) The list of possible values for this variable.  If type is **integer** or **date**, then the array of string is  converted to array of integers or date during the runtime.
				* `position` - (Optional, Integer) The relative position of this variable in a list.
				* `required` - (Optional, Boolean) If the variable required?.
				* `secure` - (Optional, Boolean) Is the variable secure or sensitive ?.
				* `source` - (Optional, String) The source of this meta-data.
				* `type` - (Optional, String) Type of the variable.
				  * Constraints: Allowable values are: `boolean`, `string`, `integer`, `date`, `array`, `list`, `map`, `complex`, `link`.
			* `name` - (Optional, String) The name of the variable. For example, `name = "inventory username"`.
			* `use_default` - (Optional, Boolean) True, will ignore the data in the value attribute, instead the data in metadata.default_value will be used.
			* `value` - (Optional, String) The value for the variable or reference to the value. For example, `value = "<provide your ssh_key_value with \n>"`. **Note** The SSH key should contain `\n` at the end of the key details in case of command line or API calls.
		* `settings` - (Optional, List) Environment variables used by all the templates in the Action.
		Nested scheme for **settings**:
			* `link` - (Optional, String) The reference link to the variable value By default the expression points to `$self.value`.
			* `metadata` - (Optional, List) An user editable metadata for the variables.
			Nested scheme for **metadata**:
				* `aliases` - (Optional, List) The list of aliases for the variable name.
				* `cloud_data_type` - (Optional, String) Cloud data type of the variable. eg. resource_group_id, region, vpc_id.
				* `default_value` - (Optional, String) Default value for the variable only if the override value is not specified.
				* `description` - (Optional, String) The description of the meta data.
				* `group_by` - (Optional, String) The display name of the group this variable belongs to.
				* `hidden` - (Optional, Boolean) If **true**, the variable is not displayed on UI or Command line.
				* `immutable` - (Optional, Boolean) Is the variable readonly ?.
				* `link_status` - (Optional, String) The status of the link.
				  * Constraints: Allowable values are: `normal`, `broken`.
				* `matches` - (Optional, String) The regex for the variable value.
				* `max_length` - (Optional, Integer) The maximum length of the variable value. Applicable for the string type.
				* `max_value` - (Optional, Integer) The maximum value of the variable. Applicable for the integer type.
				* `min_length` - (Optional, Integer) The minimum length of the variable value. Applicable for the string type.
				* `min_value` - (Optional, Integer) The minimum value of the variable. Applicable for the integer type.
				* `options` - (Optional, List) The list of possible values for this variable.  If type is **integer** or **date**, then the array of string is  converted to array of integers or date during the runtime.
				* `position` - (Optional, Integer) The relative position of this variable in a list.
				* `required` - (Optional, Boolean) If the variable required?.
				* `secure` - (Optional, Boolean) Is the variable secure or sensitive ?.
				* `source` - (Optional, String) The source of this meta-data.
				* `type` - (Optional, String) Type of the variable.
				  * Constraints: Allowable values are: `boolean`, `string`, `integer`, `date`, `array`, `list`, `map`, `complex`, `link`.
			* `name` - (Optional, String) The name of the variable. For example, `name = "inventory username"`.
			* `use_default` - (Optional, Boolean) True, will ignore the data in the value attribute, instead the data in metadata.default_value will be used.
			* `value` - (Optional, String) The value for the variable or reference to the value. For example, `value = "<provide your ssh_key_value with \n>"`. **Note** The SSH key should contain `\n` at the end of the key details in case of command line or API calls.
		* `updated_at` - (Optional, String) Job status updation timestamp.
	* `flow_job_data` - (Optional, List) Flow Job data.
	Nested scheme for **flow_job_data**:
		* `flow_id` - (Optional, String) Flow ID.
		* `flow_name` - (Optional, String) Flow Name.
		* `updated_at` - (Optional, String) Job status updation timestamp.
		* `workitems` - (Optional, List) Job data used by each workitem Job.
		Nested scheme for **workitems**:
			* `command_object_id` - (Optional, String) command object id.
			* `command_object_name` - (Optional, String) command object name.
			* `inputs` - (Optional, List) Input variables data for the workItem used in FlowJob.
			Nested scheme for **inputs**:
				* `link` - (Optional, String) The reference link to the variable value By default the expression points to `$self.value`.
				* `metadata` - (Optional, List) An user editable metadata for the variables.
				Nested scheme for **metadata**:
					* `aliases` - (Optional, List) The list of aliases for the variable name.
					* `cloud_data_type` - (Optional, String) Cloud data type of the variable. eg. resource_group_id, region, vpc_id.
					* `default_value` - (Optional, String) Default value for the variable only if the override value is not specified.
					* `description` - (Optional, String) The description of the meta data.
					* `group_by` - (Optional, String) The display name of the group this variable belongs to.
					* `hidden` - (Optional, Boolean) If **true**, the variable is not displayed on UI or Command line.
					* `immutable` - (Optional, Boolean) Is the variable readonly ?.
					* `link_status` - (Optional, String) The status of the link.
					  * Constraints: Allowable values are: `normal`, `broken`.
					* `matches` - (Optional, String) The regex for the variable value.
					* `max_length` - (Optional, Integer) The maximum length of the variable value. Applicable for the string type.
					* `max_value` - (Optional, Integer) The maximum value of the variable. Applicable for the integer type.
					* `min_length` - (Optional, Integer) The minimum length of the variable value. Applicable for the string type.
					* `min_value` - (Optional, Integer) The minimum value of the variable. Applicable for the integer type.
					* `options` - (Optional, List) The list of possible values for this variable.  If type is **integer** or **date**, then the array of string is  converted to array of integers or date during the runtime.
					* `position` - (Optional, Integer) The relative position of this variable in a list.
					* `required` - (Optional, Boolean) If the variable required?.
					* `secure` - (Optional, Boolean) Is the variable secure or sensitive ?.
					* `source` - (Optional, String) The source of this meta-data.
					* `type` - (Optional, String) Type of the variable.
					  * Constraints: Allowable values are: `boolean`, `string`, `integer`, `date`, `array`, `list`, `map`, `complex`, `link`.
				* `name` - (Optional, String) The name of the variable. For example, `name = "inventory username"`.
				* `use_default` - (Optional, Boolean) True, will ignore the data in the value attribute, instead the data in metadata.default_value will be used.
				* `value` - (Optional, String) The value for the variable or reference to the value. For example, `value = "<provide your ssh_key_value with \n>"`. **Note** The SSH key should contain `\n` at the end of the key details in case of command line or API calls.
			* `last_job` - (Optional, List) Status of the last job executed by the workitem.
			Nested scheme for **last_job**:
				* `command_name` - (Optional, String) Schematics job command name.
				  * Constraints: Allowable values are: `workspace_plan`, `workspace_apply`, `workspace_destroy`, `workspace_refresh`, `ansible_playbook_run`, `ansible_playbook_check`, `create_action`, `put_action`, `patch_action`, `delete_action`, `system_key_enable`, `system_key_delete`, `system_key_disable`, `system_key_rotate`, `system_key_restore`, `create_workspace`, `put_workspace`, `patch_workspace`, `delete_workspace`, `create_cart`, `create_environment`, `put_environment`, `delete_environment`, `environment_init`, `environment_install`, `environment_uninstall`, `repository_process`, `terraform_commands`.
				* `command_object` - (Optional, String) Name of the Schematics automation resource.
				  * Constraints: Allowable values are: `workspace`, `action`, `system`, `environment`.
				* `command_object_id` - (Optional, String) Workitem command object id, maps to workspace_id or action_id.
				* `command_object_name` - (Optional, String) command object name (workspace_name/action_name).
				* `job_id` - (Optional, String) Workspace job id.
				* `job_status` - (Optional, String) Status of Jobs.
				  * Constraints: Allowable values are: `job_pending`, `job_in_progress`, `job_finished`, `job_failed`, `job_cancelled`.
			* `layers` - (Optional, String) layer name.
			* `outputs` - (Optional, List) Output variables for the workItem.
			Nested scheme for **outputs**:
				* `link` - (Optional, String) The reference link to the variable value By default the expression points to `$self.value`.
				* `metadata` - (Optional, List) An user editable metadata for the variables.
				Nested scheme for **metadata**:
					* `aliases` - (Optional, List) The list of aliases for the variable name.
					* `cloud_data_type` - (Optional, String) Cloud data type of the variable. eg. resource_group_id, region, vpc_id.
					* `default_value` - (Optional, String) Default value for the variable only if the override value is not specified.
					* `description` - (Optional, String) The description of the meta data.
					* `group_by` - (Optional, String) The display name of the group this variable belongs to.
					* `hidden` - (Optional, Boolean) If **true**, the variable is not displayed on UI or Command line.
					* `immutable` - (Optional, Boolean) Is the variable readonly ?.
					* `link_status` - (Optional, String) The status of the link.
					  * Constraints: Allowable values are: `normal`, `broken`.
					* `matches` - (Optional, String) The regex for the variable value.
					* `max_length` - (Optional, Integer) The maximum length of the variable value. Applicable for the string type.
					* `max_value` - (Optional, Integer) The maximum value of the variable. Applicable for the integer type.
					* `min_length` - (Optional, Integer) The minimum length of the variable value. Applicable for the string type.
					* `min_value` - (Optional, Integer) The minimum value of the variable. Applicable for the integer type.
					* `options` - (Optional, List) The list of possible values for this variable.  If type is **integer** or **date**, then the array of string is  converted to array of integers or date during the runtime.
					* `position` - (Optional, Integer) The relative position of this variable in a list.
					* `required` - (Optional, Boolean) If the variable required?.
					* `secure` - (Optional, Boolean) Is the variable secure or sensitive ?.
					* `source` - (Optional, String) The source of this meta-data.
					* `type` - (Optional, String) Type of the variable.
					  * Constraints: Allowable values are: `boolean`, `string`, `integer`, `date`, `array`, `list`, `map`, `complex`, `link`.
				* `name` - (Optional, String) The name of the variable. For example, `name = "inventory username"`.
				* `use_default` - (Optional, Boolean) True, will ignore the data in the value attribute, instead the data in metadata.default_value will be used.
				* `value` - (Optional, String) The value for the variable or reference to the value. For example, `value = "<provide your ssh_key_value with \n>"`. **Note** The SSH key should contain `\n` at the end of the key details in case of command line or API calls.
			* `settings` - (Optional, List) Environment variables for the workItem.
			Nested scheme for **settings**:
				* `link` - (Optional, String) The reference link to the variable value By default the expression points to `$self.value`.
				* `metadata` - (Optional, List) An user editable metadata for the variables.
				Nested scheme for **metadata**:
					* `aliases` - (Optional, List) The list of aliases for the variable name.
					* `cloud_data_type` - (Optional, String) Cloud data type of the variable. eg. resource_group_id, region, vpc_id.
					* `default_value` - (Optional, String) Default value for the variable only if the override value is not specified.
					* `description` - (Optional, String) The description of the meta data.
					* `group_by` - (Optional, String) The display name of the group this variable belongs to.
					* `hidden` - (Optional, Boolean) If **true**, the variable is not displayed on UI or Command line.
					* `immutable` - (Optional, Boolean) Is the variable readonly ?.
					* `link_status` - (Optional, String) The status of the link.
					  * Constraints: Allowable values are: `normal`, `broken`.
					* `matches` - (Optional, String) The regex for the variable value.
					* `max_length` - (Optional, Integer) The maximum length of the variable value. Applicable for the string type.
					* `max_value` - (Optional, Integer) The maximum value of the variable. Applicable for the integer type.
					* `min_length` - (Optional, Integer) The minimum length of the variable value. Applicable for the string type.
					* `min_value` - (Optional, Integer) The minimum value of the variable. Applicable for the integer type.
					* `options` - (Optional, List) The list of possible values for this variable.  If type is **integer** or **date**, then the array of string is  converted to array of integers or date during the runtime.
					* `position` - (Optional, Integer) The relative position of this variable in a list.
					* `required` - (Optional, Boolean) If the variable required?.
					* `secure` - (Optional, Boolean) Is the variable secure or sensitive ?.
					* `source` - (Optional, String) The source of this meta-data.
					* `type` - (Optional, String) Type of the variable.
					  * Constraints: Allowable values are: `boolean`, `string`, `integer`, `date`, `array`, `list`, `map`, `complex`, `link`.
				* `name` - (Optional, String) The name of the variable. For example, `name = "inventory username"`.
				* `use_default` - (Optional, Boolean) True, will ignore the data in the value attribute, instead the data in metadata.default_value will be used.
				* `value` - (Optional, String) The value for the variable or reference to the value. For example, `value = "<provide your ssh_key_value with \n>"`. **Note** The SSH key should contain `\n` at the end of the key details in case of command line or API calls.
			* `source` - (Optional, List) Source of templates, playbooks, or controls.
			Nested scheme for **source**:
				* `catalog` - (Optional, List) The connection details to the IBM Cloud Catalog source.
				Nested scheme for **catalog**:
					* `catalog_name` - (Optional, String) The name of the private catalog.
					* `offering_id` - (Optional, String) The ID of an offering in the IBM Cloud Catalog.
					* `offering_kind` - (Optional, String) The type of an offering, in the IBM Cloud Catalog.
					* `offering_name` - (Optional, String) The name of an offering in the IBM Cloud Catalog.
					* `offering_repo_url` - (Optional, String) The repository URL of an offering, in the IBM Cloud Catalog.
					* `offering_version` - (Optional, String) The version string of an offering in the IBM Cloud Catalog.
					* `offering_version_id` - (Optional, String) The ID of an offering version the IBM Cloud Catalog.
				* `git` - (Optional, List) The connection details to the Git source repository.
				Nested scheme for **git**:
					* `computed_git_repo_url` - (Optional, String) The complete URL which is computed by the **git_repo_url**, **git_repo_folder**, and **branch**.
					* `git_branch` - (Optional, String) The name of the branch that are used to fetch the Git repository.
					* `git_release` - (Optional, String) The name of the release tag that are used to fetch the Git repository.
					* `git_repo_folder` - (Optional, String) The name of the folder in the Git repository, that contains the template.
					* `git_repo_url` - (Optional, String) The URL to the Git repository that can be used to clone the template.
					* `git_token` - (Optional, String) The Personal Access Token (PAT) to connect to the Git URLs.
				* `source_type` - (Required, String) Type of source for the Template.
				  * Constraints: Allowable values are: `local`, `git_hub`, `git_hub_enterprise`, `git_lab`, `ibm_git_lab`, `ibm_cloud_catalog`, `external_scm`.
			* `source_type` - (Optional, String) Type of source for the Template.
			  * Constraints: Allowable values are: `local`, `git_hub`, `git_hub_enterprise`, `git_lab`, `ibm_git_lab`, `ibm_cloud_catalog`, `external_scm`.
			* `updated_at` - (Optional, String) Job status updation timestamp.
	* `job_type` - (Required, String) Type of Job.
	  * Constraints: Allowable values are: `repo_download_job`, `workspace_job`, `action_job`, `system_job`, `flow-job`.
	* `system_job_data` - (Optional, List) Controls Job data.
	Nested scheme for **system_job_data**:
		* `key_id` - (Optional, String) Key ID for which key event is generated.
		* `schematics_resource_id` - (Optional, List) List of the schematics resource id.
		* `updated_at` - (Optional, String) Job status updation timestamp.
	* `workspace_job_data` - (Optional, List) Workspace Job data.
	Nested scheme for **workspace_job_data**:
		* `flow_id` - (Optional, String) Flow Id.
		* `flow_name` - (Optional, String) Flow name.
		* `inputs` - (Optional, List) Input variables data used by the Workspace Job.
		Nested scheme for **inputs**:
			* `link` - (Optional, String) The reference link to the variable value By default the expression points to `$self.value`.
			* `metadata` - (Optional, List) An user editable metadata for the variables.
			Nested scheme for **metadata**:
				* `aliases` - (Optional, List) The list of aliases for the variable name.
				* `cloud_data_type` - (Optional, String) Cloud data type of the variable. eg. resource_group_id, region, vpc_id.
				* `default_value` - (Optional, String) Default value for the variable only if the override value is not specified.
				* `description` - (Optional, String) The description of the meta data.
				* `group_by` - (Optional, String) The display name of the group this variable belongs to.
				* `hidden` - (Optional, Boolean) If **true**, the variable is not displayed on UI or Command line.
				* `immutable` - (Optional, Boolean) Is the variable readonly ?.
				* `link_status` - (Optional, String) The status of the link.
				  * Constraints: Allowable values are: `normal`, `broken`.
				* `matches` - (Optional, String) The regex for the variable value.
				* `max_length` - (Optional, Integer) The maximum length of the variable value. Applicable for the string type.
				* `max_value` - (Optional, Integer) The maximum value of the variable. Applicable for the integer type.
				* `min_length` - (Optional, Integer) The minimum length of the variable value. Applicable for the string type.
				* `min_value` - (Optional, Integer) The minimum value of the variable. Applicable for the integer type.
				* `options` - (Optional, List) The list of possible values for this variable.  If type is **integer** or **date**, then the array of string is  converted to array of integers or date during the runtime.
				* `position` - (Optional, Integer) The relative position of this variable in a list.
				* `required` - (Optional, Boolean) If the variable required?.
				* `secure` - (Optional, Boolean) Is the variable secure or sensitive ?.
				* `source` - (Optional, String) The source of this meta-data.
				* `type` - (Optional, String) Type of the variable.
				  * Constraints: Allowable values are: `boolean`, `string`, `integer`, `date`, `array`, `list`, `map`, `complex`, `link`.
			* `name` - (Optional, String) The name of the variable. For example, `name = "inventory username"`.
			* `use_default` - (Optional, Boolean) True, will ignore the data in the value attribute, instead the data in metadata.default_value will be used.
			* `value` - (Optional, String) The value for the variable or reference to the value. For example, `value = "<provide your ssh_key_value with \n>"`. **Note** The SSH key should contain `\n` at the end of the key details in case of command line or API calls.
		* `outputs` - (Optional, List) Output variables data from the Workspace Job.
		Nested scheme for **outputs**:
			* `link` - (Optional, String) The reference link to the variable value By default the expression points to `$self.value`.
			* `metadata` - (Optional, List) An user editable metadata for the variables.
			Nested scheme for **metadata**:
				* `aliases` - (Optional, List) The list of aliases for the variable name.
				* `cloud_data_type` - (Optional, String) Cloud data type of the variable. eg. resource_group_id, region, vpc_id.
				* `default_value` - (Optional, String) Default value for the variable only if the override value is not specified.
				* `description` - (Optional, String) The description of the meta data.
				* `group_by` - (Optional, String) The display name of the group this variable belongs to.
				* `hidden` - (Optional, Boolean) If **true**, the variable is not displayed on UI or Command line.
				* `immutable` - (Optional, Boolean) Is the variable readonly ?.
				* `link_status` - (Optional, String) The status of the link.
				  * Constraints: Allowable values are: `normal`, `broken`.
				* `matches` - (Optional, String) The regex for the variable value.
				* `max_length` - (Optional, Integer) The maximum length of the variable value. Applicable for the string type.
				* `max_value` - (Optional, Integer) The maximum value of the variable. Applicable for the integer type.
				* `min_length` - (Optional, Integer) The minimum length of the variable value. Applicable for the string type.
				* `min_value` - (Optional, Integer) The minimum value of the variable. Applicable for the integer type.
				* `options` - (Optional, List) The list of possible values for this variable.  If type is **integer** or **date**, then the array of string is  converted to array of integers or date during the runtime.
				* `position` - (Optional, Integer) The relative position of this variable in a list.
				* `required` - (Optional, Boolean) If the variable required?.
				* `secure` - (Optional, Boolean) Is the variable secure or sensitive ?.
				* `source` - (Optional, String) The source of this meta-data.
				* `type` - (Optional, String) Type of the variable.
				  * Constraints: Allowable values are: `boolean`, `string`, `integer`, `date`, `array`, `list`, `map`, `complex`, `link`.
			* `name` - (Optional, String) The name of the variable. For example, `name = "inventory username"`.
			* `use_default` - (Optional, Boolean) True, will ignore the data in the value attribute, instead the data in metadata.default_value will be used.
			* `value` - (Optional, String) The value for the variable or reference to the value. For example, `value = "<provide your ssh_key_value with \n>"`. **Note** The SSH key should contain `\n` at the end of the key details in case of command line or API calls.
		* `settings` - (Optional, List) Environment variables used by all the templates in the Workspace.
		Nested scheme for **settings**:
			* `link` - (Optional, String) The reference link to the variable value By default the expression points to `$self.value`.
			* `metadata` - (Optional, List) An user editable metadata for the variables.
			Nested scheme for **metadata**:
				* `aliases` - (Optional, List) The list of aliases for the variable name.
				* `cloud_data_type` - (Optional, String) Cloud data type of the variable. eg. resource_group_id, region, vpc_id.
				* `default_value` - (Optional, String) Default value for the variable only if the override value is not specified.
				* `description` - (Optional, String) The description of the meta data.
				* `group_by` - (Optional, String) The display name of the group this variable belongs to.
				* `hidden` - (Optional, Boolean) If **true**, the variable is not displayed on UI or Command line.
				* `immutable` - (Optional, Boolean) Is the variable readonly ?.
				* `link_status` - (Optional, String) The status of the link.
				  * Constraints: Allowable values are: `normal`, `broken`.
				* `matches` - (Optional, String) The regex for the variable value.
				* `max_length` - (Optional, Integer) The maximum length of the variable value. Applicable for the string type.
				* `max_value` - (Optional, Integer) The maximum value of the variable. Applicable for the integer type.
				* `min_length` - (Optional, Integer) The minimum length of the variable value. Applicable for the string type.
				* `min_value` - (Optional, Integer) The minimum value of the variable. Applicable for the integer type.
				* `options` - (Optional, List) The list of possible values for this variable.  If type is **integer** or **date**, then the array of string is  converted to array of integers or date during the runtime.
				* `position` - (Optional, Integer) The relative position of this variable in a list.
				* `required` - (Optional, Boolean) If the variable required?.
				* `secure` - (Optional, Boolean) Is the variable secure or sensitive ?.
				* `source` - (Optional, String) The source of this meta-data.
				* `type` - (Optional, String) Type of the variable.
				  * Constraints: Allowable values are: `boolean`, `string`, `integer`, `date`, `array`, `list`, `map`, `complex`, `link`.
			* `name` - (Optional, String) The name of the variable. For example, `name = "inventory username"`.
			* `use_default` - (Optional, Boolean) True, will ignore the data in the value attribute, instead the data in metadata.default_value will be used.
			* `value` - (Optional, String) The value for the variable or reference to the value. For example, `value = "<provide your ssh_key_value with \n>"`. **Note** The SSH key should contain `\n` at the end of the key details in case of command line or API calls.
		* `template_data` - (Optional, List) Input / output data of the Template in the Workspace Job.
		Nested scheme for **template_data**:
			* `flow_index` - (Optional, Integer) Index of the template in the Flow.
			* `inputs` - (Optional, List) Job inputs used by the Templates.
			Nested scheme for **inputs**:
				* `link` - (Optional, String) The reference link to the variable value By default the expression points to `$self.value`.
				* `metadata` - (Optional, List) An user editable metadata for the variables.
				Nested scheme for **metadata**:
					* `aliases` - (Optional, List) The list of aliases for the variable name.
					* `cloud_data_type` - (Optional, String) Cloud data type of the variable. eg. resource_group_id, region, vpc_id.
					* `default_value` - (Optional, String) Default value for the variable only if the override value is not specified.
					* `description` - (Optional, String) The description of the meta data.
					* `group_by` - (Optional, String) The display name of the group this variable belongs to.
					* `hidden` - (Optional, Boolean) If **true**, the variable is not displayed on UI or Command line.
					* `immutable` - (Optional, Boolean) Is the variable readonly ?.
					* `link_status` - (Optional, String) The status of the link.
					  * Constraints: Allowable values are: `normal`, `broken`.
					* `matches` - (Optional, String) The regex for the variable value.
					* `max_length` - (Optional, Integer) The maximum length of the variable value. Applicable for the string type.
					* `max_value` - (Optional, Integer) The maximum value of the variable. Applicable for the integer type.
					* `min_length` - (Optional, Integer) The minimum length of the variable value. Applicable for the string type.
					* `min_value` - (Optional, Integer) The minimum value of the variable. Applicable for the integer type.
					* `options` - (Optional, List) The list of possible values for this variable.  If type is **integer** or **date**, then the array of string is  converted to array of integers or date during the runtime.
					* `position` - (Optional, Integer) The relative position of this variable in a list.
					* `required` - (Optional, Boolean) If the variable required?.
					* `secure` - (Optional, Boolean) Is the variable secure or sensitive ?.
					* `source` - (Optional, String) The source of this meta-data.
					* `type` - (Optional, String) Type of the variable.
					  * Constraints: Allowable values are: `boolean`, `string`, `integer`, `date`, `array`, `list`, `map`, `complex`, `link`.
				* `name` - (Optional, String) The name of the variable. For example, `name = "inventory username"`.
				* `use_default` - (Optional, Boolean) True, will ignore the data in the value attribute, instead the data in metadata.default_value will be used.
				* `value` - (Optional, String) The value for the variable or reference to the value. For example, `value = "<provide your ssh_key_value with \n>"`. **Note** The SSH key should contain `\n` at the end of the key details in case of command line or API calls.
			* `outputs` - (Optional, List) Job output from the Templates.
			Nested scheme for **outputs**:
				* `link` - (Optional, String) The reference link to the variable value By default the expression points to `$self.value`.
				* `metadata` - (Optional, List) An user editable metadata for the variables.
				Nested scheme for **metadata**:
					* `aliases` - (Optional, List) The list of aliases for the variable name.
					* `cloud_data_type` - (Optional, String) Cloud data type of the variable. eg. resource_group_id, region, vpc_id.
					* `default_value` - (Optional, String) Default value for the variable only if the override value is not specified.
					* `description` - (Optional, String) The description of the meta data.
					* `group_by` - (Optional, String) The display name of the group this variable belongs to.
					* `hidden` - (Optional, Boolean) If **true**, the variable is not displayed on UI or Command line.
					* `immutable` - (Optional, Boolean) Is the variable readonly ?.
					* `link_status` - (Optional, String) The status of the link.
					  * Constraints: Allowable values are: `normal`, `broken`.
					* `matches` - (Optional, String) The regex for the variable value.
					* `max_length` - (Optional, Integer) The maximum length of the variable value. Applicable for the string type.
					* `max_value` - (Optional, Integer) The maximum value of the variable. Applicable for the integer type.
					* `min_length` - (Optional, Integer) The minimum length of the variable value. Applicable for the string type.
					* `min_value` - (Optional, Integer) The minimum value of the variable. Applicable for the integer type.
					* `options` - (Optional, List) The list of possible values for this variable.  If type is **integer** or **date**, then the array of string is  converted to array of integers or date during the runtime.
					* `position` - (Optional, Integer) The relative position of this variable in a list.
					* `required` - (Optional, Boolean) If the variable required?.
					* `secure` - (Optional, Boolean) Is the variable secure or sensitive ?.
					* `source` - (Optional, String) The source of this meta-data.
					* `type` - (Optional, String) Type of the variable.
					  * Constraints: Allowable values are: `boolean`, `string`, `integer`, `date`, `array`, `list`, `map`, `complex`, `link`.
				* `name` - (Optional, String) The name of the variable. For example, `name = "inventory username"`.
				* `use_default` - (Optional, Boolean) True, will ignore the data in the value attribute, instead the data in metadata.default_value will be used.
				* `value` - (Optional, String) The value for the variable or reference to the value. For example, `value = "<provide your ssh_key_value with \n>"`. **Note** The SSH key should contain `\n` at the end of the key details in case of command line or API calls.
			* `settings` - (Optional, List) Environment variables used by the template.
			Nested scheme for **settings**:
				* `link` - (Optional, String) The reference link to the variable value By default the expression points to `$self.value`.
				* `metadata` - (Optional, List) An user editable metadata for the variables.
				Nested scheme for **metadata**:
					* `aliases` - (Optional, List) The list of aliases for the variable name.
					* `cloud_data_type` - (Optional, String) Cloud data type of the variable. eg. resource_group_id, region, vpc_id.
					* `default_value` - (Optional, String) Default value for the variable only if the override value is not specified.
					* `description` - (Optional, String) The description of the meta data.
					* `group_by` - (Optional, String) The display name of the group this variable belongs to.
					* `hidden` - (Optional, Boolean) If **true**, the variable is not displayed on UI or Command line.
					* `immutable` - (Optional, Boolean) Is the variable readonly ?.
					* `link_status` - (Optional, String) The status of the link.
					  * Constraints: Allowable values are: `normal`, `broken`.
					* `matches` - (Optional, String) The regex for the variable value.
					* `max_length` - (Optional, Integer) The maximum length of the variable value. Applicable for the string type.
					* `max_value` - (Optional, Integer) The maximum value of the variable. Applicable for the integer type.
					* `min_length` - (Optional, Integer) The minimum length of the variable value. Applicable for the string type.
					* `min_value` - (Optional, Integer) The minimum value of the variable. Applicable for the integer type.
					* `options` - (Optional, List) The list of possible values for this variable.  If type is **integer** or **date**, then the array of string is  converted to array of integers or date during the runtime.
					* `position` - (Optional, Integer) The relative position of this variable in a list.
					* `required` - (Optional, Boolean) If the variable required?.
					* `secure` - (Optional, Boolean) Is the variable secure or sensitive ?.
					* `source` - (Optional, String) The source of this meta-data.
					* `type` - (Optional, String) Type of the variable.
					  * Constraints: Allowable values are: `boolean`, `string`, `integer`, `date`, `array`, `list`, `map`, `complex`, `link`.
				* `name` - (Optional, String) The name of the variable. For example, `name = "inventory username"`.
				* `use_default` - (Optional, Boolean) True, will ignore the data in the value attribute, instead the data in metadata.default_value will be used.
				* `value` - (Optional, String) The value for the variable or reference to the value. For example, `value = "<provide your ssh_key_value with \n>"`. **Note** The SSH key should contain `\n` at the end of the key details in case of command line or API calls.
			* `template_id` - (Optional, String) Template Id.
			* `template_name` - (Optional, String) Template name.
			* `updated_at` - (Optional, String) Job status updation timestamp.
		* `updated_at` - (Optional, String) Job status updation timestamp.
		* `workspace_name` - (Optional, String) Workspace name.

* `description` - (Optional, String) The description of your job is derived from the related action or workspace.  The description can be up to 2048 characters long in size.

* `duration` - (Optional, String) Duration of job execution; example 40 sec.

* `end_at` - (Optional, String) Job end time.

* `id` - (Optional, String) Job ID.

* `job_env_settings` - (Optional, List) Environment variables used by the Job while performing Action or Workspace.
Nested scheme for **job_env_settings**:
	* `link` - (Optional, String) The reference link to the variable value By default the expression points to `$self.value`.
	* `metadata` - (Optional, List) An user editable metadata for the variables.
	Nested scheme for **metadata**:
		* `aliases` - (Optional, List) The list of aliases for the variable name.
		* `cloud_data_type` - (Optional, String) Cloud data type of the variable. eg. resource_group_id, region, vpc_id.
		* `default_value` - (Optional, String) Default value for the variable only if the override value is not specified.
		* `description` - (Optional, String) The description of the meta data.
		* `group_by` - (Optional, String) The display name of the group this variable belongs to.
		* `hidden` - (Optional, Boolean) If **true**, the variable is not displayed on UI or Command line.
		* `immutable` - (Optional, Boolean) Is the variable readonly ?.
		* `link_status` - (Optional, String) The status of the link.
		  * Constraints: Allowable values are: `normal`, `broken`.
		* `matches` - (Optional, String) The regex for the variable value.
		* `max_length` - (Optional, Integer) The maximum length of the variable value. Applicable for the string type.
		* `max_value` - (Optional, Integer) The maximum value of the variable. Applicable for the integer type.
		* `min_length` - (Optional, Integer) The minimum length of the variable value. Applicable for the string type.
		* `min_value` - (Optional, Integer) The minimum value of the variable. Applicable for the integer type.
		* `options` - (Optional, List) The list of possible values for this variable.  If type is **integer** or **date**, then the array of string is  converted to array of integers or date during the runtime.
		* `position` - (Optional, Integer) The relative position of this variable in a list.
		* `required` - (Optional, Boolean) If the variable required?.
		* `secure` - (Optional, Boolean) Is the variable secure or sensitive ?.
		* `source` - (Optional, String) The source of this meta-data.
		* `type` - (Optional, String) Type of the variable.
		  * Constraints: Allowable values are: `boolean`, `string`, `integer`, `date`, `array`, `list`, `map`, `complex`, `link`.
	* `name` - (Optional, String) The name of the variable. For example, `name = "inventory username"`.
	* `use_default` - (Optional, Boolean) True, will ignore the data in the value attribute, instead the data in metadata.default_value will be used.
	* `value` - (Optional, String) The value for the variable or reference to the value. For example, `value = "<provide your ssh_key_value with \n>"`. **Note** The SSH key should contain `\n` at the end of the key details in case of command line or API calls.

* `job_inputs` - (Optional, List) Job inputs used by Action or Workspace.
Nested scheme for **job_inputs**:
	* `link` - (Optional, String) The reference link to the variable value By default the expression points to `$self.value`.
	* `metadata` - (Optional, List) An user editable metadata for the variables.
	Nested scheme for **metadata**:
		* `aliases` - (Optional, List) The list of aliases for the variable name.
		* `cloud_data_type` - (Optional, String) Cloud data type of the variable. eg. resource_group_id, region, vpc_id.
		* `default_value` - (Optional, String) Default value for the variable only if the override value is not specified.
		* `description` - (Optional, String) The description of the meta data.
		* `group_by` - (Optional, String) The display name of the group this variable belongs to.
		* `hidden` - (Optional, Boolean) If **true**, the variable is not displayed on UI or Command line.
		* `immutable` - (Optional, Boolean) Is the variable readonly ?.
		* `link_status` - (Optional, String) The status of the link.
		  * Constraints: Allowable values are: `normal`, `broken`.
		* `matches` - (Optional, String) The regex for the variable value.
		* `max_length` - (Optional, Integer) The maximum length of the variable value. Applicable for the string type.
		* `max_value` - (Optional, Integer) The maximum value of the variable. Applicable for the integer type.
		* `min_length` - (Optional, Integer) The minimum length of the variable value. Applicable for the string type.
		* `min_value` - (Optional, Integer) The minimum value of the variable. Applicable for the integer type.
		* `options` - (Optional, List) The list of possible values for this variable.  If type is **integer** or **date**, then the array of string is  converted to array of integers or date during the runtime.
		* `position` - (Optional, Integer) The relative position of this variable in a list.
		* `required` - (Optional, Boolean) If the variable required?.
		* `secure` - (Optional, Boolean) Is the variable secure or sensitive ?.
		* `source` - (Optional, String) The source of this meta-data.
		* `type` - (Optional, String) Type of the variable.
		  * Constraints: Allowable values are: `boolean`, `string`, `integer`, `date`, `array`, `list`, `map`, `complex`, `link`.
	* `name` - (Optional, String) The name of the variable. For example, `name = "inventory username"`.
	* `use_default` - (Optional, Boolean) True, will ignore the data in the value attribute, instead the data in metadata.default_value will be used.
	* `value` - (Optional, String) The value for the variable or reference to the value. For example, `value = "<provide your ssh_key_value with \n>"`. **Note** The SSH key should contain `\n` at the end of the key details in case of command line or API calls.

* `job_runner_id` - (Optional, String) ID of the Job Runner.

* `location` - (Optional, String) List of locations supported by IBM Cloud Schematics service.  While creating your workspace or action, choose the right region, since it cannot be changed.  Note, this does not limit the location of the IBM Cloud resources, provisioned using Schematics.
  * Constraints: Allowable values are: `us-south`, `us-east`, `eu-gb`, `eu-de`.

* `log_store_url` - (Optional, String) Job log store URL.

* `log_summary` - (Optional, List) Job log summary record.
Nested scheme for **log_summary**:
	* `action_job` - (Optional, List) Flow Job log summary.
	Nested scheme for **action_job**:
		* `play_count` - (Optional, Float) number of plays in playbook.
		* `recap` - (Optional, List) Recap records.
		Nested scheme for **recap**:
			* `changed` - (Optional, Float) Number of changed.
			* `failed` - (Optional, Float) Number of failed.
			* `ok` - (Optional, Float) Number of OK.
			* `skipped` - (Optional, Float) Number of skipped.
			* `target` - (Optional, List) List of target or host name.
			* `unreachable` - (Optional, Float) Number of unreachable.
		* `target_count` - (Optional, Float) number of targets or hosts.
		* `task_count` - (Optional, Float) number of tasks in playbook.
	* `elapsed_time` - (Optional, Float) Job log elapsed time (log_analyzed_till - log_start_at).
	* `flow_job` - (Optional, List) Flow Job log summary.
	Nested scheme for **flow_job**:
		* `workitems` - (Optional, List)
		Nested scheme for **workitems**:
			* `job_id` - (Optional, String) workspace JOB ID.
			* `log_url` - (Optional, String) Log url for job.
			* `resources_add` - (Optional, Float) Number of resources add.
			* `resources_destroy` - (Optional, Float) Number of resources destroy.
			* `resources_modify` - (Optional, Float) Number of resources modify.
			* `workspace_id` - (Optional, String) workspace ID.
		* `workitems_completed` - (Optional, Float) Number of workitems completed successfully.
		* `workitems_failed` - (Optional, Float) Number of workitems failed.
		* `workitems_pending` - (Optional, Float) Number of workitems pending in the flow.
	* `job_id` - (Optional, String) Workspace Id.
	* `job_type` - (Optional, String) Type of Job.
	  * Constraints: Allowable values are: `repo_download_job`, `workspace_job`, `action_job`, `system_job`, `flow_job`.
	* `log_analyzed_till` - (Optional, String) Job log update timestamp.
	* `log_errors` - (Optional, List) Job log errors.
	Nested scheme for **log_errors**:
		* `error_code` - (Optional, String) Error code in the Log.
		* `error_count` - (Optional, Float) Number of occurrence.
		* `error_msg` - (Optional, String) Summary error message in the log.
	* `log_start_at` - (Optional, String) Job log start timestamp.
	* `repo_download_job` - (Optional, List) Repo download Job log summary.
	Nested scheme for **repo_download_job**:
		* `detected_filetype` - (Optional, String) Detected template or data file type.
		* `inputs_count` - (Optional, String) Number of inputs detected.
		* `outputs_count` - (Optional, String) Number of outputs detected.
		* `quarantined_file_count` - (Optional, Float) Number of files quarantined.
		* `scanned_file_count` - (Optional, Float) Number of files scanned.
	* `system_job` - (Optional, List) System Job log summary.
	Nested scheme for **system_job**:
		* `failed` - (Optional, Float) Number of failed.
		* `success` - (Optional, Float) Number of passed.
		* `target_count` - (Optional, Float) number of targets or hosts.
	* `workspace_job` - (Optional, List) Workspace Job log summary.
	Nested scheme for **workspace_job**:
		* `resources_add` - (Optional, Float) Number of resources add.
		* `resources_destroy` - (Optional, Float) Number of resources destroy.
		* `resources_modify` - (Optional, Float) Number of resources modify.

* `name` - (Optional, String) Job name, uniquely derived from the related Workspace or Action.

* `resource_group` - (Optional, String) Resource-group name derived from the related Workspace or Action.

* `results_url` - (Optional, String) Job results store URL.

* `start_at` - (Optional, String) Job start time.

* `state_store_url` - (Optional, String) Job state store URL.

* `status` - (Optional, List) Job Status.
Nested scheme for **status**:
	* `action_job_status` - (Optional, List) Action Job Status.
	Nested scheme for **action_job_status**:
		* `action_name` - (Optional, String) Action name.
		* `bastion_status_code` - (Optional, String) Status of Resources.
		  * Constraints: Allowable values are: `none`, `ready`, `processing`, `error`.
		* `bastion_status_message` - (Optional, String) Bastion status message - to be displayed along with the bastion_status_code;.
		* `status_code` - (Optional, String) Status of Jobs.
		  * Constraints: Allowable values are: `job_pending`, `job_in_progress`, `job_finished`, `job_failed`, `job_cancelled`.
		* `status_message` - (Optional, String) Action Job status message - to be displayed along with the action_status_code.
		* `targets_status_code` - (Optional, String) Status of Resources.
		  * Constraints: Allowable values are: `none`, `ready`, `processing`, `error`.
		* `targets_status_message` - (Optional, String) Aggregated status message for all target resources,  to be displayed along with the targets_status_code;.
		* `updated_at` - (Optional, String) Job status updation timestamp.
	* `flow_job_status` - (Optional, List) Environment Flow JOB Status.
	Nested scheme for **flow_job_status**:
		* `flow_id` - (Optional, String) flow id.
		* `flow_name` - (Optional, String) flow name.
		* `status_code` - (Optional, String) Status of Jobs.
		  * Constraints: Allowable values are: `job_pending`, `job_in_progress`, `job_finished`, `job_failed`, `job_cancelled`.
		* `status_message` - (Optional, String) Flow Job status message - to be displayed along with the status_code;.
		* `updated_at` - (Optional, String) Job status updation timestamp.
		* `workitems` - (Optional, List) Environment's individual workItem status details;.
		Nested scheme for **workitems**:
			* `job_id` - (Optional, String) workspace job id.
			* `status_code` - (Optional, String) Status of Jobs.
			  * Constraints: Allowable values are: `job_pending`, `job_in_progress`, `job_finished`, `job_failed`, `job_cancelled`.
			* `status_message` - (Optional, String) workitem job status message;.
			* `updated_at` - (Optional, String) workitem job status updation timestamp.
			* `workspace_id` - (Optional, String) Workspace id.
			* `workspace_name` - (Optional, String) workspace name.
	* `position_in_queue` - (Optional, Float) Position of job in pending queue.
	* `system_job_status` - (Optional, List) System Job Status.
	Nested scheme for **system_job_status**:
		* `schematics_resource_status` - (Optional, List) job staus for each schematics resource.
		Nested scheme for **schematics_resource_status**:
			* `schematics_resource_id` - (Optional, String) id for each resource which is targeted as a part of system job.
			* `status_code` - (Optional, String) Status of Jobs.
			  * Constraints: Allowable values are: `job_pending`, `job_in_progress`, `job_finished`, `job_failed`, `job_cancelled`.
			* `status_message` - (Optional, String) system job status message.
			* `updated_at` - (Optional, String) Job status updation timestamp.
		* `system_status_code` - (Optional, String) Status of Jobs.
		  * Constraints: Allowable values are: `job_pending`, `job_in_progress`, `job_finished`, `job_failed`, `job_cancelled`.
		* `system_status_message` - (Optional, String) System job message.
		* `updated_at` - (Optional, String) Job status updation timestamp.
	* `total_in_queue` - (Optional, Float) Total no. of jobs in pending queue.
	* `workspace_job_status` - (Optional, List) Workspace Job Status.
	Nested scheme for **workspace_job_status**:
		* `commands` - (Optional, List) List of terraform commands executed and their status.
		Nested scheme for **commands**:
			* `name` - (Optional, String) Name of the command.
			* `outcome` - (Optional, String) outcome of the command.
		* `flow_status` - (Optional, List) Environment Flow JOB Status.
		Nested scheme for **flow_status**:
			* `flow_id` - (Optional, String) flow id.
			* `flow_name` - (Optional, String) flow name.
			* `status_code` - (Optional, String) Status of Jobs.
			  * Constraints: Allowable values are: `job_pending`, `job_in_progress`, `job_finished`, `job_failed`, `job_cancelled`.
			* `status_message` - (Optional, String) Flow Job status message - to be displayed along with the status_code;.
			* `updated_at` - (Optional, String) Job status updation timestamp.
			* `workitems` - (Optional, List) Environment's individual workItem status details;.
			Nested scheme for **workitems**:
				* `job_id` - (Optional, String) workspace job id.
				* `status_code` - (Optional, String) Status of Jobs.
				  * Constraints: Allowable values are: `job_pending`, `job_in_progress`, `job_finished`, `job_failed`, `job_cancelled`.
				* `status_message` - (Optional, String) workitem job status message;.
				* `updated_at` - (Optional, String) workitem job status updation timestamp.
				* `workspace_id` - (Optional, String) Workspace id.
				* `workspace_name` - (Optional, String) workspace name.
		* `status_code` - (Optional, String) Status of Jobs.
		  * Constraints: Allowable values are: `job_pending`, `job_in_progress`, `job_finished`, `job_failed`, `job_cancelled`.
		* `status_message` - (Optional, String) Workspace job status message (eg. App1_Setup_Pending, for a 'Setup' flow in the 'App1' Workspace).
		* `template_status` - (Optional, List) Workspace Flow Template job status.
		Nested scheme for **template_status**:
			* `flow_index` - (Optional, Integer) Index of the template in the Flow.
			* `status_code` - (Optional, String) Status of Jobs.
			  * Constraints: Allowable values are: `job_pending`, `job_in_progress`, `job_finished`, `job_failed`, `job_cancelled`.
			* `status_message` - (Optional, String) Template job status message (eg. VPCt1_Apply_Pending, for a 'VPCt1' Template).
			* `template_id` - (Optional, String) Template Id.
			* `template_name` - (Optional, String) Template name.
			* `updated_at` - (Optional, String) Job status updation timestamp.
		* `updated_at` - (Optional, String) Job status updation timestamp.
		* `workspace_name` - (Optional, String) Workspace name.

* `submitted_at` - (Optional, String) Job submission time.

* `submitted_by` - (Optional, String) Email address of user who submitted the job.

* `tags` - (Optional, List) User defined tags, while running the job.

* `updated_at` - (Optional, String) Job status updation timestamp.

