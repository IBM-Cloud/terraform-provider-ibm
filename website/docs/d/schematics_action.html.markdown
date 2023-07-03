---
subcategory: "Schematics"
layout: "ibm"
page_title: "IBM : ibm_schematics_action"
sidebar_current: "docs-ibm-datasource-schematics-action"
description: |-
  Get information about Schematics action.
---

# ibm_schematics_action
Retrieve information about a Schematics action. For more details about the Schematics and Schematics actions, see [Setting up an action](https://cloud.ibm.com/docs/schematics?topic=schematics-action-setup).

## Example usage

```terraform
data "ibm_schematics_action" "schematics_action" {
	action_id = "action_id"
}
```

## Argument reference
Review the argument references that you can specify for your data source. 

- `action_id` - (Required, String) Action Id.  Use GET /actions API to look up the Action Ids in your IBM Cloud account.

* `location` - (Optional,String) Location supported by IBM Cloud Schematics service.  While creating your workspace or action, choose the right region, since it cannot be changed.  Note, this does not limit the location of the IBM Cloud resources, provisioned using Schematics.
  * Constraints: Allowable values are: us-south, us-east, eu-gb, eu-de


## Attribute reference
In addition to the argument reference list, you can access the following attribute references after your data source is created. 

  * `action_inputs` - (List) Input variables for the Action.
Nested scheme for **action_inputs**:
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

* `action_outputs` - (List) Output variables for the Action.
Nested scheme for **action_outputs**:
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

* `bastion` - (List) Describes a bastion resource.
Nested scheme for **bastion**:
	* `name` - (String) Bastion Name(Unique).
	* `host` - (String) Reference to the Inventory resource definition.

* `bastion_credential` - (List) User editable variable data & system generated reference to value.
Nested scheme for **bastion_credential**:
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

* `command_parameter` - (String) Schematics job command parameter (playbook-name).

* `created_at` - (String) Action creation time.

* `created_by` - (String) E-mail address of the user who created an action.

* `credentials` - (List) credentials of the Action.
Nested scheme for **credentials**:
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

* `crn` - (String) Action Cloud Resource Name.

* `description` - (String) Action description.

* `id` - (String) Action ID.

* `inventory` - (String) Target inventory record ID, used by the action or ansible playbook.

* `name` - (String) The unique name of your action. The name can be up to 128 characters long and can include alphanumeric characters, spaces, dashes, and underscores. **Example** you can use the name to stop action.

* `playbook_names` - (List) Playbook names retrieved from the respository.

* `resource_group` - (String) Resource-group name for an action.  By default, action is created in default resource group.

* `settings` - (List) Environment variables for the Action.
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

* `source` - (List) Source of templates, playbooks, or controls.
Nested scheme for **source**:
	* `source_type` - (Required, String) Type of source for the Template.
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

* `source_created_at` - (String) Action Playbook Source creation time.

* `source_created_by` - (String) E-mail address of user who created the Action Playbook Source.

* `source_readme_url` - (String) URL of the `README` file, for the source URL.

* `source_type` - (String) Type of source for the Template.
  * Constraints: Allowable values are: local, git_hub, git_hub_enterprise, git_lab, ibm_git_lab, ibm_cloud_catalog, external_scm, cos_bucket

* `source_updated_at` - (String) The action playbook updation time.

* `source_updated_by` - (String) E-mail address of user who updated the action playbook source.

* `state` - (List) Computed state of the Action.
Nested scheme for **state**:
	* `status_code` - (String) Status of automation (workspace or action).
	  * Constraints: Allowable values are: normal, pending, disabled, critical
	* `status_job_id` - (String) Job id reference for this status.
	* `status_message` - (String) Automation status message - to be displayed along with the status_code.

* `sys_lock` - (List) System lock status.
Nested scheme for **sys_lock**:
	* `sys_locked` - (Boolean) Is the automation locked by a Schematic job ?.
	* `sys_locked_by` - (String) Name of the User who performed the job, that lead to the locking of the automation.
	* `sys_locked_at` - (String) When the User performed the job that lead to locking of the automation ?.

* `tags` - (List) Action tags.

* `targets_ini` - (String) Inventory of host and host group for the playbook in `INI` file format. For example, `"targets_ini": "[webserverhost]  172.22.192.6  [dbhost]  172.22.192.5"`. For more information, about an inventory host group syntax, see [Inventory host groups](https://cloud.ibm.com/docs/schematics?topic=schematics-schematics-cli-reference#schematics-inventory-host-grps).

* `updated_at` - (String) Action updation time.

* `updated_by` - (String) E-mail address of the user who updated an action.

* `user_state` - (List) User defined status of the Schematics object.
Nested scheme for **user_state**:
	* `state` - (String) User-defined states  * `draft` Object can be modified; can be used by Jobs run by the author, during execution  * `live` Object can be modified; can be used by Jobs during execution  * `locked` Object cannot be modified; can be used by Jobs during execution  * `disable` Object can be modified. cannot be used by Jobs during execution.
	  * Constraints: Allowable values are: draft, live, locked, disable
	* `set_by` - (String) Name of the User who set the state of the Object.
	* `set_at` - (String) When the User who set the state of the Object.

