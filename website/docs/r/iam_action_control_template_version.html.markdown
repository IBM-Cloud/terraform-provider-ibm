---
layout: "ibm"
page_title: "IBM : ibm_iam_action_control_template_version"
description: |-
  Manages action_control_template_version
subcategory: "Identity & Access Management (IAM)"
---

# ibm_iam_action_control_template_version

Create, update, and delete a action_control_template versions with this resource.

## Example Usage

```hcl
resource "ibm_iam_action_control_template_version" "action_control_template_v2" {
  action_control_template_id = ibm_iam_action_control_template.action_control_template_v1.action_control_template_id
  description = "Template description"
  action_control {
	actions = ["am-test-service.test.create" ]
	service_name="am-test-service"
	}
  committed = "true"
}
```

## Argument Reference

You can specify the following arguments for this resource.

* `action_control_template_id` - (Required, String) Template id for the action control template to create a new version.
* `name` - (Optional) field when creating a new template. Otherwise this field is optional. If the field is included it will change the name value for all existing versions of the template.
* `committed` - (Optional, Boolean) Committed status of the template version. If committed is set to true, then the template version can no longer be updated.
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
* `name` - (String) Required field when creating a new template. Otherwise this field is optional. If the field is included it will change the name value for all existing versions of the template.
* `version` - (String) The action control template version.
* `account_id` - (String) Enterprise account ID where template will be created.
* `etag` - ETag identifier for action_control_template.

## Import

You can import the `ibm_iam_action_control_template_version` resource by using `version`.
The `version` property can be formed from `template_id`, and `version` in the following format: `<action_control_template_id>/<version>`

* `action_control_template_id`: A string. The action control template ID.
* `version`: A string. The action control template version.

### Syntax

```bash
$ terraform import ibm_iam_action_control_template_version.action_control_template $action_control_template_id/$version
```
