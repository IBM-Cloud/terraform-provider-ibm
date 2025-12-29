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
resource "ibm_iam_role_template" "role_template" {
	name = "TerraformPolicyRoleTest"
	description = "Create role template and reference in policy template through Terraform resources"
	role {
		name = "TerrPolicyRole"
		display_name = "TestingTerraformPolicyRole"
		actions = ["cloud-object-storage.bucket.get", "cloud-object-storage.bucket.delete_bucket" ]
		service_name="cloud-object-storage"
	}
	committed = true
}

resource "ibm_iam_policy_template_version" "policy_template_v2" {
  template_id = ibm_iam_policy_template.policy_template_v1.template_id
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
		role_template_references {
			id = ibm_iam_role_template.role_template.role_template_id
			version = ibm_iam_role_template.role_template.version
		}
  }
  committed = "true"
}
```

## Argument Reference

You can specify the following arguments for this resource.

* `template_id` - (Required, String) Template id for the policy template to create a new version.
* `name` - (Optional) field when creating a new template. Otherwise this field is optional. If the field is included it will change the name value for all existing versions of the template.
* `committed` - (Optional, Boolean) Committed status of the template version.
* `description` - (Optional, String) Description of the policy template. This is shown to users in the enterprise account. Use this to describe the purpose or context of the policy for enterprise users managing IAM templates.
* `policy` - (Required, List) The core set of properties associated with the template's policy objet.
Nested schema for **policy**:
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
		* `conditions` - (Optional, List) List of conditions associated with a policy, e.g., time-based conditions that grant access over a certain time period.
		Nested schema for **conditions**:
			* `key` - (Required, String) The name of an attribute.
			* `operator` - (Required, String) The operator of an attribute.
			* `value` - (Required, String) The value of a rule or resource attribute; can be boolean or string for resource attribute. Can be string or an array of strings (e.g., array of days to permit access) for rule attribute.
		* `key` - (Optional, String) The name of an attribute.
		* `operator` - (Optional, String) The operator of an attribute.
		* `value` - (Optional, String) The value of a rule or resource attribute; can be boolean or string for resource attribute. Can be string or an array of strings (e.g., array of days to permit access) for rule attribute.
	* `type` - (Required, String) The policy type: 'access'.
	* `roles` - (Optional, List) A set of displayNames.
	* `role_template_references` - (Optional, List) A set of role templates.
	Nested schema for **role_template_references**:
		* `id` - (Required, String) The role template id
		* `version` - (Required, String) The role template version
	
	**NOTE**: Policies may include roles, role_template_references, or a combination of both
## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the policy_template. The ID is composed of `<template_id>/<template_version>`.
* `template_id` - (String) The policy template ID.
* `name` - (String) Required field when creating a new template. Otherwise this field is optional. If the field is included it will change the name value for all existing versions of the template.
* `version` - (String) The policy template version.
* `account_id` - (String) Enterprise account ID where template will be created.
* `etag` - ETag identifier for policy_template.

## Import

You can import the `ibm_iam_policy_template_version` resource by using `version`.
The `version` property can be formed from `template_id`, and `version` in the following format: `<template_id>/<version>`

* `template_id`: A string. The policy template ID.
* `version`: A string. The policy template version.

### Syntax

```bash
$ terraform import ibm_iam_policy_template_version.policy_template $template_id/$version
```
