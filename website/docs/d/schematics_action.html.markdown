---
subcategory: "Schematics"
layout: "ibm"
page_title: "IBM : ibm_schematics_action"
sidebar_current: "docs-ibm-datasource-schematics-action"
description: |-
  Get information about schematics_action
subcategory: "Schematics Service API"
---

# ibm_schematics_action

Provides a read-only data source for schematics_action. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "ibm_schematics_action" "schematics_action" {
	action_id = "action_id"
}
```

## Argument Reference

Review the argument reference that you can specify for your data source.

* `action_id` - (Required, Forces new resource, String) Action Id.  Use GET /actions API to look up the Action Ids in your IBM Cloud account.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

* `id` - The unique identifier of the schematics_action.
* `account` - (Optional, String) Action account ID.

* `action_inputs` - (Optional, List) Input variables for the Action.
Nested scheme for **action_inputs**:
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

* `action_outputs` - (Optional, List) Output variables for the Action.
Nested scheme for **action_outputs**:
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

* `bastion` - (Optional, List) Describes a bastion resource.
Nested scheme for **bastion**:
	* `host` - (Optional, String) Reference to the Inventory resource definition.
	* `name` - (Optional, String) Bastion Name(Unique).

* `bastion_connection_type` - (Optional, String) Type of connection to be used when connecting to bastion host. If the `inventory_connection_type=winrm`, then `bastion_connection_type` is not supported.
  * Constraints: Allowable values are: `ssh`.

* `bastion_credential` - (Optional, List) User editable credential variable data and system generated reference to the value.
Nested scheme for **bastion_credential**:
	* `link` - (Optional, String) The reference link to the variable value By default the expression points to `$self.value`.
	* `metadata` - (Optional, List) An user editable metadata for the credential variables.
	Nested scheme for **metadata**:
		* `aliases` - (Optional, List) The list of aliases for the variable name.
		* `cloud_data_type` - (Optional, String) Cloud data type of the credential variable. eg. api_key, iam_token, profile_id.
		* `default_value` - (Optional, String) Default value for the variable only if the override value is not specified.
		* `description` - (Optional, String) The description of the meta data.
		* `group_by` - (Optional, String) The display name of the group this variable belongs to.
		* `hidden` - (Optional, Boolean) If **true**, the variable is not displayed on UI or Command line.
		* `immutable` - (Optional, Boolean) Is the variable readonly ?.
		* `link_status` - (Optional, String) The status of the link.
		  * Constraints: Allowable values are: `normal`, `broken`.
		* `position` - (Optional, Integer) The relative position of this variable in a list.
		* `required` - (Optional, Boolean) If the variable required?.
		* `source` - (Optional, String) The source of this meta-data.
		* `type` - (Optional, String) Type of the variable.
		  * Constraints: Allowable values are: `string`, `link`.
	* `name` - (Optional, String) The name of the credential variable.
	* `use_default` - (Optional, Boolean) True, will ignore the data in the value attribute, instead the data in metadata.default_value will be used.
	* `value` - (Optional, String) The credential value for the variable or reference to the value. For example, `value = "<provide your ssh_key_value with \n>"`. **Note** The SSH key should contain `\n` at the end of the key details in case of command line or API calls.

* `command_parameter` - (Optional, String) Schematics job command parameter (playbook-name).

* `created_at` - (Optional, String) Action creation time.

* `created_by` - (Optional, String) E-mail address of the user who created an action.

* `credentials` - (Optional, List) credentials of the Action.
Nested scheme for **credentials**:
	* `link` - (Optional, String) The reference link to the variable value By default the expression points to `$self.value`.
	* `metadata` - (Optional, List) An user editable metadata for the credential variables.
	Nested scheme for **metadata**:
		* `aliases` - (Optional, List) The list of aliases for the variable name.
		* `cloud_data_type` - (Optional, String) Cloud data type of the credential variable. eg. api_key, iam_token, profile_id.
		* `default_value` - (Optional, String) Default value for the variable only if the override value is not specified.
		* `description` - (Optional, String) The description of the meta data.
		* `group_by` - (Optional, String) The display name of the group this variable belongs to.
		* `hidden` - (Optional, Boolean) If **true**, the variable is not displayed on UI or Command line.
		* `immutable` - (Optional, Boolean) Is the variable readonly ?.
		* `link_status` - (Optional, String) The status of the link.
		  * Constraints: Allowable values are: `normal`, `broken`.
		* `position` - (Optional, Integer) The relative position of this variable in a list.
		* `required` - (Optional, Boolean) If the variable required?.
		* `source` - (Optional, String) The source of this meta-data.
		* `type` - (Optional, String) Type of the variable.
		  * Constraints: Allowable values are: `string`, `link`.
	* `name` - (Optional, String) The name of the credential variable.
	* `use_default` - (Optional, Boolean) True, will ignore the data in the value attribute, instead the data in metadata.default_value will be used.
	* `value` - (Optional, String) The credential value for the variable or reference to the value. For example, `value = "<provide your ssh_key_value with \n>"`. **Note** The SSH key should contain `\n` at the end of the key details in case of command line or API calls.

* `crn` - (Optional, String) Action Cloud Resource Name.

* `description` - (Optional, String) Action description.

* `id` - (Optional, String) Action ID.

* `inventory` - (Optional, String) Target inventory record ID, used by the action or ansible playbook.

* `inventory_connection_type` - (Optional, String) Type of connection to be used when connecting to remote host. **Note** Currently, WinRM supports only Windows system with the public IPs and do not support Bastion host.
  * Constraints: Allowable values are: `ssh`, `winrm`.

* `location` - (Optional, String) List of locations supported by IBM Cloud Schematics service.  While creating your workspace or action, choose the right region, since it cannot be changed.  Note, this does not limit the location of the IBM Cloud resources, provisioned using Schematics.
  * Constraints: Allowable values are: `us-south`, `us-east`, `eu-gb`, `eu-de`.

* `name` - (Optional, String) The unique name of your action. The name can be up to 128 characters long and can include alphanumeric characters, spaces, dashes, and underscores. **Example** you can use the name to stop action.

* `playbook_names` - (Optional, List) Playbook names retrieved from the repository.

* `resource_group` - (Optional, String) Resource-group name for an action. By default, an action is created in `Default` resource group.

* `settings` - (Optional, List) Environment variables for the Action.
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

* `source_created_at` - (Optional, String) Action Playbook Source creation time.

* `source_created_by` - (Optional, String) E-mail address of user who created the Action Playbook Source.

* `source_readme_url` - (Optional, String) URL of the `README` file, for the source URL.

* `source_type` - (Optional, String) Type of source for the Template.
  * Constraints: Allowable values are: `local`, `git_hub`, `git_hub_enterprise`, `git_lab`, `ibm_git_lab`, `ibm_cloud_catalog`, `external_scm`.

* `source_updated_at` - (Optional, String) The action playbook updation time.

* `source_updated_by` - (Optional, String) E-mail address of user who updated the action playbook source.

* `state` - (Optional, List) Computed state of the Action.
Nested scheme for **state**:
	* `status_code` - (Optional, String) Status of automation (workspace or action).
	  * Constraints: Allowable values are: `normal`, `pending`, `disabled`, `critical`.
	* `status_job_id` - (Optional, String) Job id reference for this status.
	* `status_message` - (Optional, String) Automation status message - to be displayed along with the status_code.

* `sys_lock` - (Optional, List) System lock status.
Nested scheme for **sys_lock**:
	* `sys_locked` - (Optional, Boolean) Is the automation locked by a Schematic job ?.
	* `sys_locked_at` - (Optional, String) When the User performed the job that lead to locking of the automation ?.
	* `sys_locked_by` - (Optional, String) Name of the User who performed the job, that lead to the locking of the automation.

* `tags` - (Optional, List) Action tags.

* `targets_ini` - (Optional, String) Inventory of host and host group for the playbook in `INI` file format. For example, `"targets_ini": "[webserverhost]  172.22.192.6  [dbhost] 172.22.192.5"`. For more information, about an inventory host group syntax, see [Inventory host groups](https://cloud.ibm.com/docs/schematics?topic=schematics-schematics-cli-reference#schematics-inventory-host-grps).

* `updated_at` - (Optional, String) Action updation time.

* `updated_by` - (Optional, String) E-mail address of the user who updated an action.

* `user_state` - (Optional, List) User defined status of the Schematics object.
Nested scheme for **user_state**:
	* `set_at` - (Optional, String) When the User who set the state of the Object.
	* `set_by` - (Optional, String) Name of the User who set the state of the Object.
	* `state` - (Optional, String) User-defined states  * `draft` Object can be modified; can be used by Jobs run by the author, during execution  * `live` Object can be modified; can be used by Jobs during execution  * `locked` Object cannot be modified; can be used by Jobs during execution  * `disable` Object can be modified. cannot be used by Jobs during execution.
	  * Constraints: Allowable values are: `draft`, `live`, `locked`, `disable`.

