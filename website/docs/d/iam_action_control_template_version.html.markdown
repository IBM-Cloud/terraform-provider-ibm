---
layout: "ibm"
page_title: "IBM : ibm_iam_action_control_template_version"
description: |-
  Get information about action_control_template_version
subcategory: "Identity & Access Management (IAM)"
---

# ibm_iam_action_control_template_version

Provides a read-only data source to retrieve information about a action_control_template. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_iam_action_control_template_version" "action_control_template" {
	action_control_template_id = ibm_iam_action_control_template_version.action_control_template.action_control_template_id
	version = "version"
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `action_control_template_id` - (Required, String) The policy template ID.
* `version` - (Required, Forces new resource, String) The policy template version.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the action_control_template.
* `version` - The action_control_template version.
* `account_id` - (String) Enterprise account ID where this template will be created.

* `committed` - (Boolean) Committed status of the template version.

* `description` - (String) Description of the policy template. This is shown to users in the enterprise account. Use this to describe the purpose or context of the policy for enterprise users managing IAM templates.
  * Constraints: The maximum length is `300` characters. The minimum length is `1` character. The value must match regular expression `/^.*$/`.

* `name` - (String) Required field when creating a new template. Otherwise this field is optional. If the field is included it will change the name value for all existing versions of the template.
  * Constraints: The maximum length is `30` characters. The minimum length is `1` character.

* `action_control` - (Optional, List) The core set of properties associated with the template's action control objet.
Nested schema for **action_control**:
	* `actions` - (Required, List) List of actions to control access.
	* `description` - (Optional, String) Description of the action control.
	* `service_name` - (Required, String) The service name that the action control refers.