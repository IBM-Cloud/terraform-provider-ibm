---
subcategory: "Schematics"
layout: "ibm"
page_title: "IBM : ibm_schematics_action"
sidebar_current: "docs-ibm-datasource-schematics-action"
description: |-
  Get information about ibm_schematics_action
---

# ibm\_schematics_action

Provides a read-only data source for ibm_schematics_action. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "ibm_schematics_action" "schematics_action" {
	action_id = "action_id"
}
```

## Argument Reference

The following arguments are supported:

* `action_id` - (Required, string) Use GET or actions API to look up the action IDs in your IBM Cloud account.

## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier of the schematics_action.
* `name` - Action name (unique for an account).

* `description` - Action description.

* `location` - List of action locations supported by IBM Cloud Schematics service.  **Note** this does not limit the location of the resources provisioned using Schematics.

* `resource_group` - Resource-group name for an action.  By default, action is created in default resource group.

* `tags` - Action tags.

* `user_state` - User defined status of the Schematics object. Nested `user_state` blocks have the following structure:
	* `state` - User defined states  * `draft` Object can be modified, and can be used by jobs run by an author, during execution  * `live` Object can be modified, and can be used by jobs during execution  * `locked` Object cannot be modified, and can be used by jobs during execution  * `disable` Object can be modified, and cannot be used by Jobs during execution.
	* `set_by` - Name of the user who set the state of an Object.
	* `set_at` - When the user who set the state of an Object.

* `source_readme_url` - URL of the `README` file, for the source.

* `source` - Source of templates, playbooks, or controls. Nested `source` blocks have the following structure:
	* `source_type` - Type of source for the Template.
	* `git` - Connection details to Git source. Nested `git` blocks have the following structure:
		* `git_repo_url` - URL to the GIT Repo that can be used to clone the template.
		* `git_token` - Personal Access Token to connect to Git URLs.
		* `git_repo_folder` - Name of the folder in the Git Repo, that contains the template.
		* `git_release` - Name of the release tag, used to fetch the Git Repo.
		* `git_branch` - Name of the branch, used to fetch the Git Repo.

* `source_type` - Type of source for the Template.

* `command_parameter` - Schematics job command parameter (playbook-name, capsule-name or flow-name).

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

* `targets_ini` - Inventory of host and host group for the playbook in `INI` file format. For example, `"targets_ini": "[webserverhost]  172.22.192.6  [dbhost]  172.22.192.5"`. For more information, about an inventory host group syntax, see [Inventory host groups](/docs/schematics?topic=schematics-schematics-cli-reference#schematics-inventory-host-grps).

* `credentials` - credentials of the Action. Nested `credentials` blocks have the following structure:
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
  
* `inputs` - Input variables for an action. Nested `inputs` blocks have the following structure:
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

* `outputs` - Output variables for an action. Nested `outputs` blocks have the following structure:
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

* `settings` - Environment variables for an action. Nested `settings` blocks have the following structure:
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

* `trigger_record_id` - ID to the trigger.

* `id` - Action ID.

* `crn` - Action Cloud Resource Name.

* `account` - Action account ID.

* `source_created_at` - Action Playbook Source creation time.

* `source_created_by` - E-mail address of user who created the Action Playbook Source.

* `source_updated_at` - The action playbook updation time.

* `source_updated_by` - E-mail address of user who updated the action playbook source.

* `created_at` - Action creation time.

* `created_by` - E-mail address of the user who created an action.

* `updated_at` - Action updation time.

* `updated_by` - E-mail address of the user who updated an action.

* `namespace` - Name of the namespace.

* `state` - Computed state of an action. Nested `state` blocks have the following structure:
	* `status_code` - Status of automation (workspace or action).
	* `status_message` - Automation status message - to be displayed along with the status_code.

* `playbook_names` - Playbook names retrieved from the respository.

* `sys_lock` - System lock status. Nested `sys_lock` blocks have the following structure:
	* `sys_locked` - Is the Workspace locked by the Schematic action ?.
	* `sys_locked_by` - Name of the user who performed the action, that lead to lock the Workspace.
	* `sys_locked_at` - When the user performed the action that lead to lock the Workspace ?.

