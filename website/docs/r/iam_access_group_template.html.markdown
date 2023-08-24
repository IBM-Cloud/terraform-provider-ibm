---
layout: "ibm"
page_title: "IBM : ibm_iam_access_group_template"
description: |-
  Manages iam_access_group_template.
subcategory: "Identity & Access Management (IAM)"
---

# ibm_iam_access_group_template

Create, update, and delete iam_access_group_templates with this resource.

## Example Usage

```hcl
resource "ibm_iam_access_group_template" "iam_access_group_template_instance" {
  account_id = "accountID-123"
  description = "This access group template allows admin access to all IAM platform services in the account."
  group {
		name = "name"
		description = "description"
		members {
			users = [ "users" ]
			services = [ "services" ]
			action_controls {
				add = true
				remove = true
			}
		}
		assertions {
			rules {
				name = "name"
				expiration = 1
				realm_name = "realm_name"
				conditions {
					claim = "claim"
					operator = "operator"
					value = "value"
				}
				action_controls {
					remove = true
					update = true
				}
			}
			action_controls {
				add = true
				remove = true
				update = true
			}
		}
		action_controls {
			access {
				add = true
			}
		}
  }
  name = "IAM Admin Group template"
  policy_template_references {
		id = "id"
		version = "version"
  }
}
```

## Argument Reference

You can specify the following arguments for this resource.

* `account_id` - (Required, Forces new resource, String) The ID of the account to which the access group template is assigned.
  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9_-]+$/`.
* `description` - (Optional, Forces new resource, String) The description of the access group template.
  * Constraints: The maximum length is `250` characters. The minimum length is `0` characters. The value must match regular expression `/^[a-zA-Z0-9!@#$%^&*()_+{}:;"'<>,.?\/|\\-\\s]+$/`.
* `group` - (Optional, Forces new resource, List) Access Group Component.
Nested schema for **group**:
	* `action_controls` - (Optional, List) Access group action controls component.
	Nested schema for **action_controls**:
		* `access` - (Optional, List) Control whether or not access group administrators in child accounts can add access policies to the enterprise-managed access group in their account.
		Nested schema for **access**:
			* `add` - (Optional, Boolean) Action control for adding access policies to an enterprise-managed access group in a child account. If an access group administrator in a child account adds a policy, they can always update or remove it.
	* `assertions` - (Optional, List) Assertions Input Component.
	Nested schema for **assertions**:
		* `action_controls` - (Optional, List) Control whether or not access group administrators in child accounts can add, remove, and update dynamic rules for the enterprise-managed access group in their account. The inner level RuleActionControls override these `remove` and `update` action controls.
		Nested schema for **action_controls**:
			* `add` - (Optional, Boolean) Action control for adding dynamic rules to an enterprise-managed access group. If an access group administrator in a child account adds a dynamic rule, they can always update or remove it.
			* `remove` - (Optional, Boolean) Action control for removing enterprise-managed dynamic rules in an enterprise-managed access group.
			* `update` - (Optional, Boolean) Action control for updating enterprise-managed dynamic rules in an enterprise-managed access group.
		* `rules` - (Optional, List) Dynamic rules to automatically add federated users to access groups based on specific identity attributes.
		  * Constraints: The maximum length is `50` items. The minimum length is `0` items.
		Nested schema for **rules**:
			* `action_controls` - (Optional, List) Control whether or not access group administrators in child accounts can update and remove this dynamic rule in the enterprise-managed access group in their account.This overrides outer level AssertionsActionControls.
			Nested schema for **action_controls**:
				* `remove` - (Optional, Boolean) Action control for removing this enterprise-managed dynamic rule.
				* `update` - (Optional, Boolean) Action control for updating this enterprise-managed dynamic rule.
			* `conditions` - (Optional, List) Conditions of membership. You can think of this as a key:value pair.
			  * Constraints: The maximum length is `50` items. The minimum length is `0` items.
			Nested schema for **conditions**:
				* `claim` - (Optional, String) The key in the key:value pair.
				  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9_-]+$/`.
				* `operator` - (Optional, String) Compares the claim and the value.
				  * Constraints: The maximum length is `10` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z-]+$/`.
				* `value` - (Optional, String) The value in the key:value pair.
				  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9_-]+$/`.
			* `expiration` - (Optional, Integer) Session duration in hours. Access group membership is revoked after this time period expires. Users must log back in to refresh their access group membership.
			* `name` - (Optional, String) Dynamic rule name.
			  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9!@#$%^&*()_+{}:;"'<>,.?\/|\\-\\s]+$/`.
			* `realm_name` - (Optional, String) The identity provider (IdP) URL.
			  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9!@#$%^&*()_+{}:;"'<>,.?\/|\\-\\s]+$/`.
	* `description` - (Optional, String) Access group description. This is shown in child accounts.
	  * Constraints: The maximum length is `250` characters. The minimum length is `0` characters. The value must match regular expression `/^[a-zA-Z0-9!@#$%^&*()_+{}:;"'<>,.?\/|\\-\\s]+$/`.
	* `members` - (Optional, List) Array of enterprise users to add to the template. All enterprise users that you add to the template must be invited to the child accounts where the template is assigned.
	Nested schema for **members**:
		* `action_controls` - (Optional, List) Control whether or not access group administrators in child accounts can add and remove members from the enterprise-managed access group in their account.
		Nested schema for **action_controls**:
			* `add` - (Optional, Boolean) Action control for adding child account members to an enterprise-managed access group. If an access group administrator in a child account adds a member, they can always remove them.
			* `remove` - (Optional, Boolean) Action control for removing enterprise-managed members from an enterprise-managed access group.
		* `services` - (Optional, List) Array of service IDs to add to the template.
		  * Constraints: The list items must match regular expression `/^[a-zA-Z0-9_-]+$/`. The maximum length is `50` items. The minimum length is `0` items.
		* `users` - (Optional, List) Array of enterprise users to add to the template. All enterprise users that you add to the template must be invited to the child accounts where the template is assigned.
		  * Constraints: The list items must match regular expression `/^[a-zA-Z0-9!@#$%^&*()_+{}:;"'<>,.?\/|\\-\\s]+$/`. The maximum length is `50` items. The minimum length is `0` items.
	* `name` - (Required, String) Give the access group a unique name that doesn't conflict with other templates access group name in the given account. This is shown in child accounts.
	  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9!@#$%^&*()_+{}:;"'<>,.?\/|\\-\\s]+$/`.
* `name` - (Required, Forces new resource, String) The name of the access group template.
  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9!@#$%^&*()_+{}:;"'<>,.?\/|\\-\\s]+$/`.
* `policy_template_references` - (Optional, Forces new resource, List) References to policy templates assigned to the access group template.
  * Constraints: The maximum length is `100` items. The minimum length is `0` items.
Nested schema for **policy_template_references**:
	* `id` - (Optional, String) Policy template ID.
	  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9_-]+$/`.
	* `version` - (Optional, String) Policy template version.
	  * Constraints: The maximum length is `2` characters. The minimum length is `1` character. The value must match regular expression `/^[0-9]+$/`.
* `transaction_id` - (Optional, Forces new resource, String) An optional transaction id for the request.
  * Constraints: The maximum length is `50` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9_-]+$/`.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the iam_access_group_template.
* `committed` - (Boolean) A boolean indicating whether the access group template is committed. You must commit a template before you can assign it to child accounts.
* `created_at` - (String) The date and time when the access group template was created.
* `created_by_id` - (String) The ID of the user who created the access group template.
  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9!@#$%^&*()_+{}:;"'<>,.?\/|\\-\\s]+$/`.
* `href` - (String) The URL of the access group template resource.
  * Constraints: The maximum length is `1000` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9!@#$%^&*()_+{}:;"'<>,.?\/|\\-\\s]+$/`.
* `last_modified_at` - (String) The date and time when the access group template was last modified.
* `last_modified_by_id` - (String) The ID of the user who last modified the access group template.
  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9!@#$%^&*()_+{}:;"'<>,.?\/|\\-\\s]+$/`.
* `version` - (String) The version of the access group template.
  * Constraints: The maximum length is `2` characters. The minimum length is `1` character. The value must match regular expression `/^[0-9]+$/`.


## Import

You can import the `ibm_iam_access_group_template` resource by using `id`. The ID of the access group template.

# Syntax
```
$ terraform import ibm_iam_access_group_template.iam_access_group_template_instance <template_id>/<version>
```
