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

- `action_id` - (Required, String) Use GET or actions API to see the action IDs in your IBM Cloud account.


## Attribute reference
In addition to the argument reference list, you can access the following attribute references after your data source is created. 

- `account` - (String) The action account ID.
- `bastion` - (String) The complete target details with the user inputs and the system generated data. Nested bastion blocks have the following structure.

  Nested scheme for `bastion`:
  - `name` - (String) The target name.
  - `type` - (String) The target type such as, `cluster`, `vsi`, `icd`, `vpc`.
  - `description` - (String) The target description.
  - `resource_query` - (String) The resource selection query string.
  - `credential` - (String) Override credential for each resource. Reference to credentials values, used by all the resources.
  - `id` - (String) The target ID.
  - `created_at` - (String) The targets creation time.
  - `created_by` - (String) The Email address of the user who created the targets.
  - `updated_at` - (String) The targets update time.
  - `updated_by` - (String) The Email address of user who updated the targets.
  - `sys_lock` - (String) The system lock status.Nested sys_lock blocks have the following structure.

    Nested scheme for `sys_lock`:
    - `sys_locked` - (String) Is the workspace locked by the Schematics action?
    - `sys_locked_by` - (String) The name of the user who performed the action, that lead to lock the workspace.
    - `sys_locked_at` - (String) When the user performed the action that lead to lock the workspace?-
 - `resource_ids` - (String) Array of the resource IDs.
- `command_parameter` - (String) Schematics job command parameter such as  `playbook-name`, `capsule-name`, or `flow-name`.
- `created_at` - (Timestamp) The action creation time.
- `created_by` - (String) The Email address of the user who created an action.
- `crn` - (String) The action Cloud Resource Name.
- `credentials` - (String) Credentials of an action. Nested `credentials` blocks have the following structure.

  Nested scheme for `credentials`:
  - `metadata` - (String) User editable metadata for the variables. Nested `metadata` blocks have the following structure.

    Nested scheme for `metadata`:
    - `aliases` - (String) The list of an aliases for the variable name.
    - `description` - (String) The description of the metadata.
    - `default_value` - (String) The default value for the variable, if the override value is not specified.
    - `group_by` - (String) Display name of the group this variable belongs to.
    - `immutable` - (String) Is the variable read only?-
    - `hidden` - (String) If set **true**, the variable will not be displayed on console or command-line.
    - `min_value` - (String) Minimum value of the variable. Applicable for integer type.
    - `max_value` - (String) Maximum value of the variable. Applicable for integer type.
    - `min_length` - (String) Minimum length of the variable value. Applicable for string type.
    - `max_length` - (String) Maximum length of the variable value. Applicable for string type.
    - `matches` - (String) Regular expression for the variable value.
    - `options` - (String) The list of possible values for this variable. If type is integer or date, then the array of string will be converted to array of integers or date during runtime.
    - `position` - (String) Relative position of this variable in a list.
    - `secure` - (String) Is the variable secure or sensitive?
    - `source` - (String) Source of the meta-data.
    - `type` - (String) The type of the variable.
  - `link` - (String) Reference link to the variable value By default the expression will point to `self.value`.
  - `name` - (String) The name of the variable.
  - `value` - (String) The value for the variable or reference to the value.
- `description` - (String) The action description.
- `id` - (String) The unique ID of the Schematics action.
- `inputs` - (String) Input variables for an action. Nested `inputs` blocks have the following structure.

  Nested scheme for `inputs`:
  - `metadata` - (String) User editable metadata for the variables. Nested `metadata` blocks have the following structure.

    Nested scheme for `metadata`:
    - `aliases` - (String) The list of an aliases for the variable name.
    - `description` - (String) The description of the metadata.
    - `default_value` - (String) The default value for the variable, if the override value is not specified.
    - `group_by` - (String) The display name of the group this variable belongs to.
    - `immutable` - (String) Is the variable read only?
    - `hidden` - (String) If set to **true**, the variable will not be displayed on console or command-line.
    - `min_value` - (String) Minimum value of the variable. Applicable for integer type.
    - `max_value` - (String) Maximum value of the variable. Applicable for integer type.
    - `min_length` - (String) Minimum length of the variable value. Applicable for string type.
    - `max_length` - (String) Maximum length of the variable value. Applicable for string type.
    - `matches` - (String) Regular expression for the variable value.
    - `options` - (String) The list of possible values for this variable. If type is integer or date, then the array of string will be converted to array of integers or date during runtime.
    - `position` - (String) Relative position of this variable in a list.
    - `secure` - (Bool) Is the variable secure or sensitive?
    - `source` - (String) The source of this metadata.
    - `type` - (String) The type of the variable.
 - `link` - (String) Reference link to the variable value By default the expression will point to `self.value`.
 - `name` - (String) The name of the variable.
 - `value` - (String) The value for the variable or reference to the value.
- `location` - (String) List of action locations supported by Schematics service. **Note** this does not limit the location of the resources provisioned using Schematics.
- `name` - (String) The unique action name.
- `namespace` - (String) The name of the namespace.
- `outputs` - (String) The output variables for an action. Nested `outputs` blocks have the following structure.

  Nested scheme for `outputs`:
  - `metadata` - (String) User editable metadata for the variables. Nested metadata blocks have the following structure.

    Nested scheme for `metadata`:
    - `aliases` - (String) List of aliases for the variable name.
    - `description` - (String) Description of the meta data.
    - `default_value` - (String) Default value for the variable, if the override value is not specified.
    - `group_by` - (String) Display name of the group this variable belongs to.
    - `immutable` - (String) Is the variable read only ?
    - `hidden` - (String) If **true**, the variable will not be displayed on console or command-line.
    - `min_value` - (String) Minimum value of the variable. Applicable for integer type.
    - `max_value` - (String) Maximum value of the variable. Applicable for integer type.
    - `min_length` - (String) Minimum length of the variable value. Applicable for string type.
    - `max_length` - (String) Maximum length of the variable value. Applicable for string type.
    - `matches` - (String) Regex for the variable value.
    - `options` - (String) List of possible values for this variable. If type is integer or date, then the array of string will be converted to array of integers or date during runtime.
    - `position` - (String) Relative position of this variable in a list.
    - `secure` - (String) Is the variable secure or sensitive?
    - `source` - (String) Source of this meta-data.
    - `type` - (String) Type of the variable.
  - `link` - (String) Reference link to the variable value By default  the expression will point to self.value.
  - `name` - (String) Name of the variable.
  - `value` - (String) Value for the variable or reference to the value.
- `playbook_names` - (String) Playbook names retrieved from the repository.
- `resource_group` - (String) The resource group name for an action. By default, action is created in default resource group.
- `source_readme_url` - (String) URL of the `README` file, for the source.
- `source` - (String) Source of templates, playbooks, or controls. Nested `source` blocks have the following structure.

  Nested scheme for `source`:
  - `source_type` - (String) Type of source for the template.
  - `git` - (String) The connection details to Git source. Nested `Git` blocks have the following structure.

    Nested scheme for `git`:
    - `git_branch` - (String) The name of the branch, used to fetch the Git Repository.
	- `git_repo_folder` - (String) The name of the folder in the Git Repository, that contains the template.
	- `git_release` - (String) The name of the release tag, used to fetch the Git Repository.
    - `git_token` - (String) The personal access token to connect to Git URLs.
	- `git_repo_url` - (String) The URL to the Git repository that can be used to clone the template.
- `source_type` - (String) Type of source for the template.
- `settings` - (String) Environment variables for an action. Nested settings blocks have the following structure.

   Nested scheme for `settings`:
   - `metadata` - (String) User editable metadata for the variables. Nested metadata blocks have the following structure.

     Nested scheme for `metadata`:
     - `aliases` - (String) List of aliases for the variable name.
     - `description` - (String) Description of the meta data.
     - `default_value` - (String) Default value for the variable, if the override value is not specified.
     - `group_by` - (String) Display name of the group this variable belongs to.
     - `immutable` - (String) Is the variable read only ?.
     - `hidden` - (String) If true, the variable will not be displayed on console or command-line.
     - `min_value` - (String) Minimum value of the variable. Applicable for integer type.
     - `max_value` - (String) Maximum value of the variable. Applicable for integer type.
     - `min_length` - (String) Minimum length of the variable value. Applicable for string type.
     - `max_length` - (String) Maximum length of the variable value. Applicable for string type.
     - `matches` - (String) Regex for the variable value.
     - `options` - (String) List of possible values for this variable. If type is integer or date, then the array of string will be converted to array of integers or date during runtime.
     - `position` - (String) Relative position of this variable in a list.
     - `secure` - (String) Is the variable secure or sensitive?
     - `source` - (String) Source of this meta-data.
     - `type` - (String) Type of the variable.
  - `link` - (String) Reference link to the variable value By default the expression will point to `self.value`.
  - `name` - (String) Name of the variable.
  - `value` - (String) Value for the variable or reference to the value.
- `state` - (String) Computed state of an action. Nested `state` blocks have the following structure.

  Nested scheme for `state`:
  - `status_code` - (String) The status of automation such as `workdspace`, or `action`.
  - `status_message` - (String) Automation status message to be displayed along with the status code.
- `sys_lock` - (String) System lock status. Nested sys_lock blocks have the following structure.

  Nested scheme for `sys_lock`:
  - `sys_locked` - (String) Is the Workspace locked by the Schematic action?-
  - `sys_locked_by` - (String) Name of the user who performed the action, that lead to lock the Workspace.
  - `sys_locked_at` - (String) When the user performed the action that lead to lock the Workspace?-
- `source_created_at` - (String) The Ansible playbook source creation time.
- `source_created_by` - (String) The Email address of user who created the Ansible playbook Source.
- `source_updated_at` - (String) The Ansible playbook update time.
- `source_updated_by` - (String) The Email address of user who updated the Ansible playbook source.
- `tags` - (String) The action tags.
- `targets_ini` - (String) Inventory of host and host group for the playbook in `INI` file format. For example, `"targets_ini": "[webserverhost] 172.22.192.6 [dbhost] 172.22.192.5"`. For more information, about an inventory host group syntax, see [Inventory host groups](https://cloud.ibm.com/docs/schematics?topic=schematics-schematics-cli-reference#inventory-host-grps).
- `trigger_record_id` - (String) The trigger ID.
- `updated_at` - (Timestamp) The action update time.
- `updated_by` - (String) The Email address of the user who updated an action.
- `user_state` - (String) User defined status of the Schematics object. Nested user_state blocks have the following structure.

  Nested scheme for `user_state`:
  - `state` - (String) User defined states * **draft Object** can be modified, and can be used by jobs run by an author, during execution * **live Object** can be modified, and can be used by jobs during execution * **locked Object** cannot be modified, and can be used by jobs during execution * **disable Object** can be modified, and cannot be used by Jobs during execution.
  - `set_by` - (String) The name of an user who set the state of an Object.
  - `set_at` - (String) The user who sets the state of an object.

