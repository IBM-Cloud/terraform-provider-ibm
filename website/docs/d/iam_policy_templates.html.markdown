---
layout: "ibm"
page_title: "IBM : ibm_iam_policy_template"
description: |-
  Get information about policy_templates for an enterprise account_id
subcategory: "Identity & Access Management (IAM)"
---

# ibm_iam_policy_template

Provides a read-only data source to retrieve information about policy_templates under enterprise account_id. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_iam_policy_template" "policy_template" {

}
```

## Argument Reference

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
	* `description` - (String) Description of the policy. This is shown in child accounts when an access group or trusted profile template uses the policy template to assign access.
	  * Constraints: The maximum length is `300` characters. The minimum length is `1` character. The value must match regular expression `/^.*$/`.
	* `pattern` - (String) Indicates pattern of rule, either 'time-based-conditions:once', 'time-based-conditions:weekly:all-day', or 'time-based-conditions:weekly:custom-hours'.
	  * Constraints: The maximum length is `42` characters. The minimum length is `1` character. The value must match regular expression `/^[a-z:-]*$/`.
	* `resource` - (List) The resource attributes to which the policy grants access.
	Nested schema for **resource**:
		* `attributes` - (List) List of resource attributes to which the policy grants access.
		  * Constraints: The minimum length is `1` item.
		Nested schema for **attributes**:
			* `key` - (String) The name of a resource attribute.
			  * Constraints: The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9_]*$/`.
			* `operator` - (String) The operator of an attribute.
			  * Constraints: Allowable values are: `stringEquals`, `stringExists`, `stringMatch`. The minimum length is `1` character.
			* `value` - (String) The value of a rule or resource attribute; can be boolean or string for resource attribute. Can be string or an array of strings (e.g., array of days to permit access) for rule attribute.
		* `tags` - (List) Optional list of resource tags to which the policy grants access.
		  * Constraints: The minimum length is `1` item.
		Nested schema for **tags**:
			* `key` - (String) The name of an access management tag.
			  * Constraints: The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9 _.-]*$/`.
			* `operator` - (String) The operator of an access management tag.
			  * Constraints: Allowable values are: `stringEquals`, `stringMatch`. The minimum length is `1` character.
			* `value` - (String) The value of an access management tag.
			  * Constraints: The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9 _*?.-]*$/`.
	* `rule` - (List) Additional access conditions associated with the policy.
	Nested schema for **rule**:
		* `conditions` - (List) List of conditions associated with a policy, e.g., time-based conditions that grant access over a certain time period.
		  * Constraints: The maximum length is `10` items. The minimum length is `2` items.
		Nested schema for **conditions**:
			* `key` - (String) The name of an attribute.
			  * Constraints: The minimum length is `1` character.
			* `operator` - (String) The operator of an attribute.
			  * Constraints: Allowable values are: `timeLessThan`, `timeLessThanOrEquals`, `timeGreaterThan`, `timeGreaterThanOrEquals`, `dateTimeLessThan`, `dateTimeLessThanOrEquals`, `dateTimeGreaterThan`, `dateTimeGreaterThanOrEquals`, `dayOfWeekEquals`, `dayOfWeekAnyOf`. The minimum length is `1` character.
			* `value` - (String) The value of a rule or resource attribute; can be boolean or string for resource attribute. Can be string or an array of strings (e.g., array of days to permit access) for rule attribute.
		* `key` - (String) The name of an attribute.
		  * Constraints: The minimum length is `1` character.
		* `operator` - (String) The operator of an attribute.
		  * Constraints: Allowable values are: `timeLessThan`, `timeLessThanOrEquals`, `timeGreaterThan`, `timeGreaterThanOrEquals`, `dateTimeLessThan`, `dateTimeLessThanOrEquals`, `dateTimeGreaterThan`, `dateTimeGreaterThanOrEquals`, `dayOfWeekEquals`, `dayOfWeekAnyOf`. The minimum length is `1` character.
		* `value` - (String) The value of a rule or resource attribute; can be boolean or string for resource attribute. Can be string or an array of strings (e.g., array of days to permit access) for rule attribute.
	* `type` - (String) The policy type; either 'access' or 'authorization'.
	  * Constraints: Allowable values are: `access`, `authorization`.

