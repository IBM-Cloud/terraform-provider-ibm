---
layout: "ibm"
page_title: "IBM : ibm_iam_policy_template_version"
description: |-
  Manages policy_template_version
subcategory: "Identity & Access Management (IAM)"
---

# ibm_iam_policy_template_version

Create, update, and delete a policy_template versions with this resource.

## Example Usage

```hcl
resource "ibm_iam_policy_template_version" "policy_template_instance" {
  description = "Template description"
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

* `name` - (Optional) field when creating a new template. Otherwise this field is optional. If the field is included it will change the name value for all existing versions of the template.
* `committed` - (Optional, Boolean) Committed status of the template version.
* `description` - (Optional, String) Description of the policy template. This is shown to users in the enterprise account. Use this to describe the purpose or context of the policy for enterprise users managing IAM templates.
  * Constraints: The maximum length is `300` characters. The minimum length is `1` character. The value must match regular expression `/^.*$/`.
* `policy` - (Required, List) The core set of properties associated with the template's policy objet.
Nested schema for **policy**:
	* `control` - (Required, List) Specifies the type of access granted by the policy.
	Nested schema for **control**:
		* `grant` - (Required, List) Permission granted by the policy.
		Nested schema for **grant**:
			* `roles` - (Required, List) A set of displayNames.
			  * Constraints: The minimum length is `1` item.
	* `description` - (Optional, String) Description of the policy. This is shown in child accounts when an access group or trusted profile template uses the policy template to assign access.
	  * Constraints: The maximum length is `300` characters. The minimum length is `1` character. The value must match regular expression `/^.*$/`.
	* `pattern` - (Optional, String) Indicates pattern of rule, either 'time-based-conditions:once', 'time-based-conditions:weekly:all-day', or 'time-based-conditions:weekly:custom-hours'.
	  * Constraints: The maximum length is `42` characters. The minimum length is `1` character. The value must match regular expression `/^[a-z:-]*$/`.
	* `resource` - (Required, List) The resource attributes to which the policy grants access.
	Nested schema for **resource**:
		* `attributes` - (Required, List) List of resource attributes to which the policy grants access.
		  * Constraints: The minimum length is `1` item.
		Nested schema for **attributes**:
			* `key` - (Required, String) The name of a resource attribute.
			  * Constraints: The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9_]*$/`.
			* `operator` - (Required, String) The operator of an attribute.
			  * Constraints: Allowable values are: `stringEquals`, `stringExists`, `stringMatch`. The minimum length is `1` character.
			* `value` - (Required, String) The value of a rule or resource attribute; can be boolean or string for resource attribute. Can be string or an array of strings (e.g., array of days to permit access) for rule attribute.
		* `tags` - (Optional, List) Optional list of resource tags to which the policy grants access.
		  * Constraints: The minimum length is `1` item.
		Nested schema for **tags**:
			* `key` - (Required, String) The name of an access management tag.
			  * Constraints: The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9 _.-]*$/`.
			* `operator` - (Required, String) The operator of an access management tag.
			  * Constraints: Allowable values are: `stringEquals`, `stringMatch`. The minimum length is `1` character.
			* `value` - (Required, String) The value of an access management tag.
			  * Constraints: The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9 _*?.-]*$/`.
	* `rule` - (Optional, List) Additional access conditions associated with the policy.
	Nested schema for **rule**:
		* `conditions` - (Optional, List) List of conditions associated with a policy, e.g., time-based conditions that grant access over a certain time period.
		  * Constraints: The maximum length is `10` items. The minimum length is `2` items.
		Nested schema for **conditions**:
			* `key` - (Required, String) The name of an attribute.
			  * Constraints: The minimum length is `1` character.
			* `operator` - (Required, String) The operator of an attribute.
			  * Constraints: Allowable values are: `timeLessThan`, `timeLessThanOrEquals`, `timeGreaterThan`, `timeGreaterThanOrEquals`, `dateTimeLessThan`, `dateTimeLessThanOrEquals`, `dateTimeGreaterThan`, `dateTimeGreaterThanOrEquals`, `dayOfWeekEquals`, `dayOfWeekAnyOf`. The minimum length is `1` character.
			* `value` - (Required, String) The value of a rule or resource attribute; can be boolean or string for resource attribute. Can be string or an array of strings (e.g., array of days to permit access) for rule attribute.
		* `key` - (Optional, String) The name of an attribute.
		  * Constraints: The minimum length is `1` character.
		* `operator` - (Optional, String) The operator of an attribute.
		  * Constraints: Allowable values are: `timeLessThan`, `timeLessThanOrEquals`, `timeGreaterThan`, `timeGreaterThanOrEquals`, `dateTimeLessThan`, `dateTimeLessThanOrEquals`, `dateTimeGreaterThan`, `dateTimeGreaterThanOrEquals`, `dayOfWeekEquals`, `dayOfWeekAnyOf`. The minimum length is `1` character.
		* `value` - (Optional, String) The value of a rule or resource attribute; can be boolean or string for resource attribute. Can be string or an array of strings (e.g., array of days to permit access) for rule attribute.
	* `type` - (Required, String) The policy type; either 'access' or 'authorization'.
	  * Constraints: Allowable values are: `access`, `authorization`.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the policy_template. The ID is composed of `<template_id>/<template_version>`.
* `template_id` - (String) The policy template ID.
* `name` - (String) Required field when creating a new template. Otherwise this field is optional. If the field is included it will change the name value for all existing versions of the template.
  * Constraints: The maximum length is `30` characters. The minimum length is `1` character.
* `value` - (String) The policy template version.


* `etag` - ETag identifier for policy_template.

## Import

You can import the `ibm_iam_policy_template_version` resource by using `version`.
The `version` property can be formed from `policy_template_id`, and `version` in the following format:

```
<policy_template_id>/<version>
```
* `policy_template_id`: A string. The policy template ID.
* `version`: A string. The policy template version.

# Syntax
```
$ terraform import ibm_iam_policy_template_version.policy_template <policy_template_id>/<version>
```
