---
subcategory: "Schematics"
layout: "ibm"
page_title: "IBM : ibm_schematics_action"
sidebar_current: "docs-ibm-resource-schematics-action"
description: |-
  Manages the Schematics action.
---

# ibm_schematics_action
Create, update, and delete `ibm_schematics_action`. For more information, about Schematics action, refer to [setting up actions](https://cloud.ibm.com/docs/schematics?topic=schematics-action-setup).

## Example usage

```terraform
resource "ibm_schematics_action" "schematics_action" {
  name = "<action_name>"
  description = "<action_description>"
  location = "us-east"
  resource_group = "default"
}
```

## Argument reference

Review the argument reference that you can specify for your resource.

- `action_inputs` - (Optional, List) Input variables for the Action.
Nested scheme for **action_inputs**:
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
* `action_outputs` - (Optional, List) Output variables for the Action.
Nested scheme for **action_outputs**:
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
* `bastion` - (Optional, List) Describes a bastion resource. MaxItems: 1.
Nested scheme for **bastion**:
	* `name` - (Optional, String) Bastion Name(Unique).
	* `host` - (Optional, String) Reference to the Inventory resource definition.
* `bastion_credential` - (Optional, List) User editable variable data & system generated reference to value. MaxItems: 1
Nested scheme for **bastion_credential**:
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
* `command_parameter` - (Optional, String) Schematics job command parameter (playbook-name).
* `credentials` - (Optional, List) credentials of the Action. MaxItems: 1.
Nested scheme for **credentials**:
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
* `description` - (Optional, String) Action description.
* `inventory` - (Optional, String) Target inventory record ID, used by the action or ansible playbook.
* `location` - (Optional, String) Location supported by IBM Cloud Schematics service.  While creating your workspace or action, choose the right region, since it cannot be changed.  Note, this does not limit the location of the IBM Cloud resources, provisioned using Schematics.
  * Constraints: Allowable values are: us-south, us-east, eu-gb, eu-de
* `name` - (Required, String) The unique name of your action. The name can be up to 128 characters long and can include alphanumeric characters, spaces, dashes, and underscores. **Example** you can use the name to stop action.
* `resource_group` - (Optional, String) Resource-group name for an action.  By default, action is created in default resource group.
* `settings` - (Optional, List) Environment variables for the Action.
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
* `source` - (Optional, List) Source of templates, playbooks, or controls.
Nested scheme for **source**:
	* `source_type` - (Required, String) Type of source for the Template.
	  * Constraints: Allowable values are: local, git_hub, git_hub_enterprise, git_lab, ibm_git_lab, ibm_cloud_catalog, external_scm, cos_bucket
	* `git` - (Optional, List) Connection details to Git source.
	Nested scheme for **git**:
		* `computed_git_repo_url` - (Optional, String) The Complete URL which is computed by git_repo_url, git_repo_folder and branch.
		* `git_repo_url` - (Optional, String) URL to the GIT Repo that can be used to clone the template.
		* `git_token` - (Optional, String) Personal Access Token to connect to Git URLs.
		* `git_repo_folder` - (Optional, String) Name of the folder in the Git Repo, that contains the template.
		* `git_release` - (Optional, String) Name of the release tag, used to fetch the Git Repo.
		* `git_branch` - (Optional, String) Name of the branch, used to fetch the Git Repo.
	* `catalog` - (Optional, List) Connection details to IBM Cloud Catalog source. MaxItems:1.
	Nested scheme for **catalog**:
		* `catalog_name` - (Optional, String) name of the private catalog.
		* `offering_name` - (Optional, String) Name of the offering in the IBM Catalog.
		* `offering_version` - (Optional, String) Version string of the offering in the IBM Catalog.
		* `offering_kind` - (Optional, String) Type of the offering, in the IBM Catalog.
		* `offering_id` - (Optional, String) Id of the offering the IBM Catalog.
		* `offering_version_id` - (Optional, String) Id of the offering version the IBM Catalog.
		* `offering_repo_url` - (Optional, String) Repo Url of the offering, in the IBM Catalog.
* `source_readme_url` - (Optional, String) URL of the `README` file, for the source URL.
* `source_type` - (Optional, String) Type of source for the Template.
  * Constraints: Allowable values are: local, git_hub, git_hub_enterprise, git_lab, ibm_git_lab, ibm_cloud_catalog, external_scm, cos_bucket
* `tags` - (Optional, List) Action tags.
* `targets_ini` - (Optional, String) Inventory of host and host group for the playbook in `INI` file format. For example, `"targets_ini": "[webserverhost]  172.22.192.6  [dbhost]  172.22.192.5"`. For more information, about an inventory host group syntax, see [Inventory host groups](https://cloud.ibm.com/docs/schematics?topic=schematics-schematics-cli-reference#schematics-inventory-host-grps).
* `user_state` - (Optional, List) User defined status of the Schematics object.
Nested scheme for **user_state**:
	* `state` - (Optional, String) User-defined states  * `draft` Object can be modified; can be used by Jobs run by the author, during execution  * `live` Object can be modified; can be used by Jobs during execution  * `locked` Object cannot be modified; can be used by Jobs during execution  * `disable` Object can be modified. cannot be used by Jobs during execution.
	  * Constraints: Allowable values are: draft, live, locked, disable
	* `set_by` - (Optional, String) Name of the User who set the state of the Object.
	* `set_at` - (Optional, String) When the User who set the state of the Object.
* `x_github_token` - (Optional, String) The personal access token to authenticate with your private GitHub or GitLab repository and access your Terraform template.

## Attribute reference

In addition to all argument references listed, you can access the following attribute references after your resource is created.

* `id` - The unique identifier of the schematics_action.
* `account` - (Optional, String) Action account ID.
* `created_at` - (String) Action creation time.
* `created_by` - (String) E-mail address of the user who created an action.
* `crn` - (Optional, String) Action Cloud Resource Name.
* `playbook_names` - (Optional, List) Playbook names retrieved from the respository.
* `source_created_at` - (String) Action Playbook Source creation time.
* `source_created_by` - (String) E-mail address of user who created the Action Playbook Source.
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
* `updated_at` - (String) Action updation time.
* `updated_by` - (String) E-mail address of the user who updated an action.

## Import

You can import the `ibm_schematics_action` resource by using `id`. Action ID.

# Syntax
```
$ terraform import ibm_schematics_action.schematics_action <id>
```