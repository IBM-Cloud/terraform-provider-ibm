---
layout: "ibm"
page_title: "IBM : ibm_iam_access_group_template"
description: |-
  Get information about iam_access_group_template
subcategory: "Identity & Access Management (IAM)"
---

# ibm_iam_access_group_template

Provides a read-only data source to retrieve information about an iam_access_group_template. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_iam_access_group_template" "iam_access_group_template" {
	account_id = "accountID-123"
	verbose = true
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `account_id` - (Required, String) Enterprise account ID.
  * Constraints: The maximum length is `50` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9_-]+$/`.
* `transaction_id` - (Optional, String) An optional transaction id for the request.
  * Constraints: The maximum length is `50` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9_-]+$/`.
* `verbose` - (Optional, Boolean) If `verbose=true`, IAM resource details are returned. If performance is a concern, leave the `verbose` parameter off so that details are not retrieved.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the iam_access_group_template.
* `first` - (List) A link object.
Nested schema for **first**:
	* `href` - (String) A string containing the link’s URL.

* `group_templates` - (List) A list of access group templates.
  * Constraints: The maximum length is `50` items. The minimum length is `0` items.
Nested schema for **group_templates**:
	* `committed` - (Boolean) A boolean indicating whether the access group template is committed. You must commit a template before you can assign it to child accounts.
	* `created_at` - (String) The date and time when the access group template was created.
	* `created_by_id` - (String) The ID of the user who created the access group template.
	  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9!@#$%^&*()_+{}:;"'<>,.?\/|\\-\\s]+$/`.
	* `description` - (String) The description of the access group template.
	  * Constraints: The maximum length is `250` characters. The minimum length is `0` characters. The value must match regular expression `/^[a-zA-Z0-9!@#$%^&*()_+{}:;"'<>,.?\/|\\-\\s]+$/`.
	* `group` - (List) Access Group Component.
	Nested schema for **group**:
		* `action_controls` - (List) Access group action controls component.
		Nested schema for **action_controls**:
			* `access` - (List) Control whether or not access group administrators in child accounts can add access policies to the enterprise-managed access group in their account.
			Nested schema for **access**:
				* `add` - (Boolean) Action control for adding access policies to an enterprise-managed access group in a child account. If an access group administrator in a child account adds a policy, they can always update or remove it.
		* `assertions` - (List) Assertions Input Component.
		Nested schema for **assertions**:
			* `action_controls` - (List) Control whether or not access group administrators in child accounts can add, remove, and update dynamic rules for the enterprise-managed access group in their account. The inner level RuleActionControls override these `remove` and `update` action controls.
			Nested schema for **action_controls**:
				* `add` - (Boolean) Action control for adding dynamic rules to an enterprise-managed access group. If an access group administrator in a child account adds a dynamic rule, they can always update or remove it.
				* `remove` - (Boolean) Action control for removing enterprise-managed dynamic rules in an enterprise-managed access group.
				* `update` - (Boolean) Action control for updating enterprise-managed dynamic rules in an enterprise-managed access group.
			* `rules` - (List) Dynamic rules to automatically add federated users to access groups based on specific identity attributes.
			  * Constraints: The maximum length is `50` items. The minimum length is `0` items.
			Nested schema for **rules**:
				* `action_controls` - (List) Control whether or not access group administrators in child accounts can update and remove this dynamic rule in the enterprise-managed access group in their account.This overrides outer level AssertionsActionControls.
				Nested schema for **action_controls**:
					* `remove` - (Boolean) Action control for removing this enterprise-managed dynamic rule.
					* `update` - (Boolean) Action control for updating this enterprise-managed dynamic rule.
				* `conditions` - (List) Conditions of membership. You can think of this as a key:value pair.
				  * Constraints: The maximum length is `50` items. The minimum length is `0` items.
				Nested schema for **conditions**:
					* `claim` - (String) The key in the key:value pair.
					  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9_-]+$/`.
					* `operator` - (String) Compares the claim and the value.
					  * Constraints: The maximum length is `10` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z-]+$/`.
					* `value` - (String) The value in the key:value pair.
					  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9_-]+$/`.
				* `expiration` - (Integer) Session duration in hours. Access group membership is revoked after this time period expires. Users must log back in to refresh their access group membership.
				* `name` - (String) Dynamic rule name.
				  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9!@#$%^&*()_+{}:;"'<>,.?\/|\\-\\s]+$/`.
				* `realm_name` - (String) The identity provider (IdP) URL.
				  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9!@#$%^&*()_+{}:;"'<>,.?\/|\\-\\s]+$/`.
		* `description` - (String) Access group description. This is shown in child accounts.
		  * Constraints: The maximum length is `250` characters. The minimum length is `0` characters. The value must match regular expression `/^[a-zA-Z0-9!@#$%^&*()_+{}:;"'<>,.?\/|\\-\\s]+$/`.
		* `members` - (List) Array of enterprise users to add to the template. All enterprise users that you add to the template must be invited to the child accounts where the template is assigned.
		Nested schema for **members**:
			* `action_controls` - (List) Control whether or not access group administrators in child accounts can add and remove members from the enterprise-managed access group in their account.
			Nested schema for **action_controls**:
				* `add` - (Boolean) Action control for adding child account members to an enterprise-managed access group. If an access group administrator in a child account adds a member, they can always remove them.
				* `remove` - (Boolean) Action control for removing enterprise-managed members from an enterprise-managed access group.
			* `services` - (List) Array of service IDs to add to the template.
			  * Constraints: The list items must match regular expression `/^[a-zA-Z0-9_-]+$/`. The maximum length is `50` items. The minimum length is `0` items.
			* `users` - (List) Array of enterprise users to add to the template. All enterprise users that you add to the template must be invited to the child accounts where the template is assigned.
			  * Constraints: The list items must match regular expression `/^[a-zA-Z0-9!@#$%^&*()_+{}:;"'<>,.?\/|\\-\\s]+$/`. The maximum length is `50` items. The minimum length is `0` items.
		* `name` - (String) Give the access group a unique name that doesn't conflict with other templates access group name in the given account. This is shown in child accounts.
		  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9!@#$%^&*()_+{}:;"'<>,.?\/|\\-\\s]+$/`.
	* `href` - (String) The URL of the access group template resource.
	  * Constraints: The maximum length is `1000` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9!@#$%^&*()_+{}:;"'<>,.?\/|\\-\\s]+$/`.
	* `id` - (String) The ID of the access group template.
	  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9_-]+$/`.
	* `last_modified_at` - (String) The date and time when the access group template was last modified.
	* `last_modified_by_id` - (String) The ID of the user who last modified the access group template.
	  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9!@#$%^&*()_+{}:;"'<>,.?\/|\\-\\s]+$/`.
	* `name` - (String) The name of the access group template.
	  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9!@#$%^&*()_+{}:;"'<>,.?\/|\\-\\s]+$/`.
	* `policy_template_references` - (List) References to policy templates assigned to the access group template.
	  * Constraints: The maximum length is `100` items. The minimum length is `0` items.
	Nested schema for **policy_template_references**:
		* `id` - (String) Policy template ID.
		  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9_-]+$/`.
		* `version` - (String) Policy template version.
		  * Constraints: The maximum length is `2` characters. The minimum length is `1` character. The value must match regular expression `/^[0-9]+$/`.
	* `version` - (String) The version of the access group template.
	  * Constraints: The maximum length is `2` characters. The minimum length is `1` character. The value must match regular expression `/^[0-9]+$/`.

* `last` - (List) A link object.
Nested schema for **last**:
	* `href` - (String) A string containing the link’s URL.

* `previous` - (List) A link object.
Nested schema for **previous**:
	* `href` - (String) A string containing the link’s URL.

