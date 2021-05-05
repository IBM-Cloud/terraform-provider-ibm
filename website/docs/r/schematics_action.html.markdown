---
subcategory: "Schematics"
layout: "ibm"
page_title: "IBM : ibm_schematics_action"
sidebar_current: "docs-ibm-resource-schematics-action"
description: |-
  Manages schematics_action.
---

# ibm\_schematics_action

Provides a resource for ibm_schematics_action. This allows ibm_schematics_action to be created, updated and deleted.

## Example Usage

```hcl
resource "ibm_schematics_action" "schematics_action" {
  name = "<action_name>"
  description = "<action_description>"
  location = "us-east"
  resource_group = "default"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional, string) Action name (unique for an account).
* `description` - (Optional, string) Action description.
* `location` - (Optional, string) List of action locations supported by IBM Cloud Schematics service.  **Note** this does not limit the location of the resources provisioned using Schematics.
* `resource_group` - (Optional, string) Resource-group name for an action.  By default, action is created in default resource group.
* `tags` - (Optional, List) Action tags.
* `user_state` - (Optional, List) User defined status of the Schematics object.
  * `state` - (Optional, string) User defined states  * `draft` Object can be modified, and can be used by jobs run by an author, during execution  * `live` Object can be modified, and can be used by jobs during execution  * `locked` Object cannot be modified, and can be used by jobs during execution  * `disable` Object can be modified, and cannot be used by Jobs during execution.
  * `set_by` - (Optional, string) Name of the user who set the state of an Object.
  * `set_at` - (Optional, TypeString) When the user who set the state of an Object.
* `source_readme_url` - (Optional, string) URL of the `README` file, for the source.
* `source` - (Optional, List) Source of templates, playbooks, or controls.
  * `source_type` - (Required, string) Type of source for the Template.
  * `git` - (Optional, ExternalSourceGit) Connection details to Git source.
* `source_type` - (Optional, string) Type of source for the Template.
* `command_parameter` - (Optional, string) Schematics job command parameter (playbook-name, capsule-name or flow-name).
* `bastion` - (Optional, List) Complete target details with the user inputs and the system generated data.
  * `name` - (Optional, string) Target name.
  * `type` - (Optional, string) Target type (`cluster`, `vsi`, `icd`, `vpc`).
  * `description` - (Optional, string) Target description.
  * `resource_query` - (Optional, string) Resource selection query string.
  * `credential` - (Optional, string) Override credential for each resource.  Reference to credentials values, used by all the resources.
  * `id` - (Optional, string) Target ID.
  * `created_at` - (Optional, TypeString) Targets creation time.
  * `created_by` - (Optional, string) E-mail address of the user who created the targets.
  * `updated_at` - (Optional, TypeString) Targets updation time.
  * `updated_by` - (Optional, string) E-mail address of user who updated the targets.
  * `sys_lock` - (Optional, SystemLock) System lock status.
  * `resource_ids` - (Optional, []interface{}) Array of the resource IDs.
* `targets_ini` - (Optional, string) Inventory of host and host group for the playbook in `INI` file format. For example, `"targets_ini": "[webserverhost]  172.22.192.6  [dbhost]  172.22.192.5"`. For more information, about an inventory host group syntax, see [Inventory host groups](/docs/schematics?topic=schematics-schematics-cli-reference#schematics-inventory-host-grps).
* `credentials` - (Optional, List) credentials of the Action.
  * `name` - (Optional, string) Name of the variable.
  * `value` - (Optional, string) Value for the variable or reference to the value.
  * `metadata` - (Optional, VariableMetadata) User editable metadata for the variables.
  * `link` - (Optional, string) Reference link to the variable value By default the expression will point to self.value.
* `inputs` - (Optional, List) Input variables for an action.
  * `name` - (Optional, string) Name of the variable.
  * `value` - (Optional, string) Value for the variable or reference to the value.
  * `metadata` - (Optional, VariableMetadata) User editable metadata for the variables.
  * `link` - (Optional, string) Reference link to the variable value By default the expression will point to self.value.
* `outputs` - (Optional, List) Output variables for an action.
  * `name` - (Optional, string) Name of the variable.
  * `value` - (Optional, string) Value for the variable or reference to the value.
  * `metadata` - (Optional, VariableMetadata) User editable metadata for the variables.
  * `link` - (Optional, string) Reference link to the variable value By default the expression will point to self.value.
* `settings` - (Optional, List) Environment variables for an action.
  * `name` - (Optional, string) Name of the variable.
  * `value` - (Optional, string) Value for the variable or reference to the value.
  * `metadata` - (Optional, VariableMetadata) User editable metadata for the variables.
  * `link` - (Optional, string) Reference link to the variable value By default the expression will point to self.value.
* `trigger_record_id` - (Optional, string) ID to the trigger.
* `state` - (Optional, List) Computed state of an action.
  * `status_code` - (Optional, string) Status of automation (workspace or action).
  * `status_message` - (Optional, string) Automation status message - to be displayed along with the status_code.
* `sys_lock` - (Optional, List) System lock status.
  * `sys_locked` - (Optional, bool) Is the Workspace locked by the Schematic action ?.
  * `sys_locked_by` - (Optional, string) Name of the user who performed the action, that lead to lock the Workspace.
  * `sys_locked_at` - (Optional, TypeString) When the user performed the action that lead to lock the Workspace ?.
* `x_github_token` - (Optional, string) The personal access token to authenticate with your private GitHub or GitLab repository and access your Terraform template.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The unique identifier of the schematics_action.
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
* `playbook_names` - Playbook names retrieved from the respository.
