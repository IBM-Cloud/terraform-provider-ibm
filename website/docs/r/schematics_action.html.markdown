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
Review the argument references that you can specify for your resource.

- `bastion` - (Optional, List) Complete target details with the user inputs and the system generated data.

  Nested scheme for `bastion`:
  - `created_at` - (Optional, TypeString) The targets creation time.
  - `created_by` - (Optional, String) E-mail address of the user who created the targets.
  - `credential` - (Optional, String) Override credential for each resource.  Reference to credentials values, that are used by the resources.
  - `description` - (Optional, String) The target description.
  - `id` - (Optional, String) The target ID.
  - `name` - (Optional, String) The target name.
  - `resource_ids` - (Optional, []interface{}) Array of the resource IDs.
  - `resource_query` - (Optional, String) The resource selection query string.
  - `sys_lock` - (Optional, SystemLock) The System lock status.
  - `type` - (Optional, String) The target type such as `cluster`, `vsi`, `icd`, `vpc`.
  - `updated_at` - (Optional, TypeString) The targets updation time.
  - `updated_by` - (Optional, String) E-mail address of user who updated the targets.
- `command_parameter` - (Optional, String) The Schematics job command parameter such as `playbook-name`, `capsule-name` or `flow-name`.
- `credentials` - (Optional, List) The credentials of an action.

  Nested scheme for `credentials`:
  - `link` - (Optional, String) The reference link to the variable value. By default the expression will point to `self.value`.
  - `metadata` - (Optional, VariableMetadata) An user editable metadata for the variables.
  - `name` - (Optional, String) The name of the variable.
  - `value` - (Optional, String) The value for the variable or reference to the value.
- `inputs` - (Optional, List) The input variables for an action.

  Nested scheme for `inputs`:
  - `link` - (Optional, String) The reference link to the variable value. By default the expression will point to `self.value`.
  - `metadata` - (Optional, VariableMetadata) An user editable metadata for the variables.
  - `name` - (Optional, String) The name of the variable.
  - `value` - (Optional, String) The value for the variable or reference to the value.
- `outputs` - (Optional, List) The output variables for an action.

    Nested scheme for `outputs`:
	- `link` - (Optional, String) The reference link to the variable value. By default the expression will point to `self.value`.
    - `metadata` - (Optional, VariableMetadata) An user editable metadata for the variables.
	- `name` - (Optional, String) The name of the variable.
    - `value` - (Optional, String) The value for the variable or reference to the value.
- `description` - (Optional, String) The description of an action.
- `location` - (Optional, String) List of an action locations supported by Schematics service. **Note** this does not limit the location of the resources provisioned using Schematics.
- `name` - (Optional, String) The unique name of an action.
- `resource_group` - (Optional, String) The resource group name for an action. By default, action is created in default resource group.
- `source_readme_url` - (Optional, String) The URL of the `README` file, for the source.
- `source` - (Optional, List) The source of templates, playbooks, or controls.
 
  Nested scheme for `source`:
  - `git` - (Optional, ExternalSourceGit) The connection details to the Git source.
  - `source_type` - (Required, String) Type of source for the Template.
- `source_type` - (Optional, String) The source type for the template.
- `settings` - (Optional, List) The environment variables for an action.

  Nested scheme for `settings`:
  - `link` - (Optional, String) The reference link to the variable value. By default the expression will point to `self.value`.
  - `metadata` - (Optional, VariableMetadata) An user editable metadata for the variables.
  - `name` - (Optional, String) The name of the variable.
  - `value` - (Optional, String) The value for the variable or reference to the value.
- `state` - (Optional, List) The computed state of an action.

  Nested scheme for `state`:
  - `status_code` - (Optional, String) The status of automation, such as `workspace` or `action`.
  - `status_message` - (Optional, String) An automation status message that are displayed along with the status_code.
- `sys_lock` - (Optional, List) The system lock status.

  Nested scheme for `sys_lock`:
  - `sys_locked` - (Optional, Bool) Is the Workspace locked by the Schematic action?
  - `sys_locked_by` - (Optional, String) The name of the user who performed the action, that lead to lock the Workspace.
  - `sys_locked_at` - (Optional, TypeString) When the user performed an action has lead to lock the Workspace?
- `targets_ini` - (Optional, String) Inventory of host and host group for the playbook in `INI` file format. For example, `"targets_ini": "[webserverhost]  172.22.192.6  [dbhost]  172.22.192.5"`. For more information, about an inventory host group syntax, see [Inventory host groups](/docs/schematics?topic=schematics-schematics-cli-reference#schematics-inventory-host-grps).
- `tags`- (Optional, List) An actions tags.
- `trigger_record_id` - (Optional, String) The trigger record ID.
- `user_state` - (Optional, List) User defined status of the Schematics object.

  Nested scheme for `user_state`:
  - `state` - (Optional, String) User defined states * **draft Object** can be modified, and can be used by jobs run by an author, during execution * **live Object** can be modified, and can be used by jobs during execution * **locked Object** cannot be modified, and can be used by jobs during execution * disable Object can be modified, and cannot be used by Jobs during execution.
  - `set_by` - (Optional, String) Name of the user who set the state of an object.
  - `set_at` - (Optional, TypeString) The user who set the state of an object.
- `x_github_token` - (Optional, String) The personal access token to authenticate with your private GitHub or GitLab repository and access your Terraform template.

## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `account` - (String) An action account ID.
- `crn` - (String)  An action Cloud Resource Name (`CRN`).
- `created_at` - (Timestamp)  An action creation time.
- `created_by` - (String) An Email address of the user who created an action.
- `id` - (String) The unique identifier of the Schematics workspace.
- `namespace` - (String) The name of the namespace.
- `playbook_names` - (String) The playbook names retrieved from the repository.
- `source_created_at` - (Timestamp) An Ansible playbook source creation time.
- `source_created_by` - (String)  An Email address of the user who created an Ansible playbook source.
- `source_updated_at` - (Timestamp) The action playbook update time.
- `source_updated_by` - (String) An Email address of the user who updated the action playbook source.
- `updated_at` - (Timestamp) An action update time.
- `updated_by` - (String) An Email address of the user who updated an action.
