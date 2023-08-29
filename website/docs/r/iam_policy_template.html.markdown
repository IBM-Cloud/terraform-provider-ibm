---
layout: "ibm"
page_title: "IBM : ibm_iam_policy_template"
description: |-
  Manages policy_template.
subcategory: "Identity & Access Management (IAM)"
---

# ibm_iam_policy_template

Create, update, and delete a policy_template with this resource.

## Example Usage

```hcl
resource "ibm_iam_policy_template" "policy_template_instance" {
  name = "TestTemplates"
  policy {
		type = "access"
		description = "description"
		resource {
			attributes {
				key = "key"
				operator = "stringEquals"
				value = "anything as a string"
			}
			tags {
				key = "key"
				value = "value"
				operator = "stringEquals"
			}
		}
		pattern = "pattern"
		rule {
			key = "key"
			operator = "timeLessThan"
			value = "anything as a string"
		}
		roles = ["Viewer"]
  }
  committed = "true"
}
```

## Argument Reference

You can specify the following arguments for this resource.

* `name` - (String) Required field when creating a new template. Otherwise this field is optional. If the field is included it will change the name value for all existing versions of the template.

	**Note** "Name" will be out of sync when anyone of the version resource updates this parameter. Please update this parameter with the latest version name
* `committed` - (Optional, Boolean) Committed status of the template version.
* `description` - (Optional, String) Description of the policy template. This is shown to users in the enterprise account. Use this to describe the purpose or context of the policy for enterprise users managing IAM templates.
* `policy` - (Required, List) The core set of properties associated with the template's policy objet.
Nested schema for **policy**:
	* `roles` - (Required, List) A set of displayNames.
	* `description` - (Optional, String) Description of the policy. This is shown in child accounts when an access group or trusted profile template uses the policy template to assign access.
	* `pattern` - (Optional, String) Indicates pattern of rule, either 'time-based-conditions:once', 'time-based-conditions:weekly:all-day', or 'time-based-conditions:weekly:custom-hours'.
	* `resource` - (Required, List) The resource attributes to which the policy grants access.
	Nested schema for **resource**:
		* `attributes` - (Required, List) List of resource attributes to which the policy grants access.
		Nested schema for **attributes**:
			* `key` - (Required, String) The name of a resource attribute.
			* `operator` - (Required, String) The operator of an attribute.
			* `value` - (Required, String) The value of a rule or resource attribute; can be boolean or string for resource attribute. Can be string or an array of strings (e.g., array of days to permit access) for rule attribute.
		* `tags` - (Optional, List) Optional list of resource tags to which the policy grants access.
		Nested schema for **tags**:
			* `key` - (Required, String) The name of an access management tag.
			* `operator` - (Required, String) The operator of an access management tag.
			* `value` - (Required, String) The value of an access management tag.
	* `rule` - (Optional, List) Additional access conditions associated with the policy.
	Nested schema for **rule**:
		Nested schema for **conditions**:
			* `key` - (Required, String) The name of an attribute.
			* `operator` - (Required, String) The operator of an attribute.
			* `value` - (Required, String) The value of a rule or resource attribute; can be boolean or string for resource attribute. Can be string or an array of strings (e.g., array of days to permit access) for rule attribute.
		* `key` - (Optional, String) The name of an attribute.
		* `operator` - (Optional, String) The operator of an attribute.
		* `value` - (Optional, String) The value of a rule or resource attribute; can be boolean or string for resource attribute. Can be string or an array of strings (e.g., array of days to permit access) for rule attribute.
	* `type` - (Required, String) The policy type: 'access'.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the policy_template. The ID is composed of `<template_id>/<template_version>`.
* `template_id` - (String) The policy template ID.
* `etag` - ETag identifier for policy_template.
* `account_id` - (String) Enterprise account ID where this template will be created.

## Import

You can import the `ibm_iam_policy_template` resource by using `version`.
The `version` property can be formed from `policy_template_id`, and `version` in the following format:

```
<policy_template_id>/<version>
```
* `policy_template_id`: A string. The policy template ID.
* `version`: A string. The policy template version.

# Syntax
```
$ terraform import ibm_iam_policy_template.policy_template <policy_template_id>/<version>
```
