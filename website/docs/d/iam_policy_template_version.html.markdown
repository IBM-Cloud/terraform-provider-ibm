---
layout: "ibm"
page_title: "IBM : ibm_iam_policy_template_version"
description: |-
  Get information about policy_template_version
subcategory: "Identity & Access Management (IAM)"
---

# ibm_iam_policy_template_version

Provides a read-only data source to retrieve information about a policy_template. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_iam_policy_template_version" "policy_template" {
	policy_template_id = ibm_iam_policy_template_version.policy_template.policy_template_id
	version = "version"
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `policy_template_id` - (Required, String) The policy template ID.
* `version` - (Required, Forces new resource, String) The policy template version.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the policy_template.
* `version` - The policy_template version.
* `account_id` - (String) Enterprise account ID where this template will be created.

* `committed` - (Boolean) Committed status of the template version.

* `description` - (String) Description of the policy template. This is shown to users in the enterprise account. Use this to describe the purpose or context of the policy for enterprise users managing IAM templates.
  * Constraints: The maximum length is `300` characters. The minimum length is `1` character. The value must match regular expression `/^.*$/`.

* `name` - (String) Required field when creating a new template. Otherwise this field is optional. If the field is included it will change the name value for all existing versions of the template.
  * Constraints: The maximum length is `30` characters. The minimum length is `1` character.

* `policy` - (List) The core set of properties associated with the template's policy objet.
Nested schema for **policy**:
	* `roles` - (List) A set of displayNames.
	* `role_template_references` - (Optional, List) A set of role templates.
	Nested schema for **role_template_references**:
		* `id` - (Required, String) The role template id
		* `version` - (Required, String) The role template version
	* `description` - (String) Description of the policy. This is shown in child accounts when an access group or trusted profile template uses the policy template to assign access.
	* `pattern` - (String) Indicates pattern of rule, either 'time-based-conditions:once', 'time-based-conditions:weekly:all-day', or 'time-based-conditions:weekly:custom-hours'.
	* `resource` - (List) The resource attributes to which the policy grants access.
	Nested schema for **resource**:
		* `attributes` - (List) List of resource attributes to which the policy grants access.
		Nested schema for **attributes**:
			* `key` - (String) The name of a resource attribute.
			* `operator` - (String) The operator of an attribute.
			* `value` - (String) The value of a rule or resource attribute; can be boolean or string for resource attribute. Can be string or an array of strings (e.g., array of days to permit access) for rule attribute.
		* `tags` - (List) Optional list of resource tags to which the policy grants access.
		Nested schema for **tags**:
			* `key` - (String) The name of an access management tag.
			* `operator` - (String) The operator of an access management tag.
			* `value` - (String) The value of an access management tag.
	* `rule` - (List) Additional access conditions associated with the policy.
	Nested schema for **rule**:
		* `conditions` - (List) List of conditions associated with a policy, e.g., time-based conditions that grant access over a certain time period.
		Nested schema for **conditions**:
			* `key` - (String) The name of an attribute.
			  * Constraints: The minimum length is `1` character.
			* `operator` - (String) The operator of an attribute.
			  * Constraints: Allowable values are: `timeLessThan`, `timeLessThanOrEquals`, `timeGreaterThan`, `timeGreaterThanOrEquals`, `dateTimeLessThan`, `dateTimeLessThanOrEquals`, `dateTimeGreaterThan`, `dateTimeGreaterThanOrEquals`, `dayOfWeekEquals`, `dayOfWeekAnyOf`. The minimum length is `1` character.
			* `value` - (String) The value of a rule or resource attribute; can be boolean or string for resource attribute. Can be string or an array of strings (e.g., array of days to permit access) for rule attribute.
		* `key` - (String) The name of an attribute.
		* `operator` - (String) The operator of an attribute.
		* `value` - (String) The value of a rule or resource attribute; can be boolean or string for resource attribute. Can be string or an array of strings (e.g., array of days to permit access) for rule attribute.
	* `type` - (String) The policy type: 'access'.
