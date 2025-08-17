---
layout: "ibm"
page_title: "IBM : ibm_iam_action_control_template"
description: |-
  Manages action_control_template.
subcategory: "Identity & Access Management (IAM)"
---

# ibm_iam_action_control_template

Create, update, and delete a action_control_template with this resource.

## Example Usage

```hcl
resource "ibm_iam_action_control_template" "action_control_template_instance" {
  name = "TestTemplates"
  description = "Base template Testing"
  action_control {
	actions = ["am-test-service.test.create" ]
	service_name="am-test-service"
	}
  committed = "true"
}
```

## Argument Reference

You can specify the following arguments for this resource.

* `name` - (String) Required field when creating a new template. Otherwise this field is optional. If the field is included it will change the name value for all existing versions of the template.

	**Note** "Name" will be out of sync when anyone of the version resource updates this parameter. Please update this parameter with the latest version name
* `committed` - (Optional, Boolean) Committed status of the template. If committed is set to true, then the template version can no longer be updated.
* `description` - (Optional, String) Description of the action control template. This is shown to users in the enterprise account. Use this to describe the purpose or context of the action control for enterprise users managing IAM templates.
* `action_control` - (Optional, List) The core set of properties associated with the template's action control objet.
Nested schema for **action_control**:
	* `actions` - (Required, List) List of actions to control access.
	* `description` - (Optional, String) Description of the action control.
	* `service_name` - (Required, String) The service name that the action control refers.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the action_control_template. The ID is composed of `<action_control_template_id>/<template_version>`.
* `action_control_template_id` - (String) The action control template ID.
* `etag` - ETag identifier for action_control_template.
* `account_id` - (String) Enterprise account ID where this template will be created.

## Import

You can import the `ibm_iam_action_control_template` resource by using `version`.
The `version` property can be formed from `template_id`, and `version` in the following format: `<template_id>/<version>`

* `action_control_template_id`: A string. The action control template ID.
* `version`: A string. The action control template version.

### Syntax

```bash
$ terraform import ibm_iam_action_control_template.action_control_template $action_control_template_id/$version
```
