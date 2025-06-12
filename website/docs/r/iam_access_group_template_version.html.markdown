---
layout: "ibm"
page_title: "IBM : ibm_iam_access_group_template_version"
description: |-
  Manages iam_access_group_template_version.
subcategory: "Identity & Access Management (IAM)"
---

# ibm_iam_access_group_template_version

Create, update, and delete iam_access_group_template_versions with this resource.

## Example Usage

```hcl
resource "ibm_iam_access_group_template_version" "iam_access_group_template_version_instance" {
  name = "IAM Admin Group template 2"
  template_id = ibm_iam_access_group_template.iam_access_group_template_v1.template_id
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
				}
			}
			action_controls {
				add = true
				remove = true
			}
		}
		action_controls {
			access {
				add = true
			}
		}
  }
  policy_template_references {
		id = "id"
		version = "version"
  }
}
```

## Argument Reference

You can specify the following arguments for this resource.

* `description` - (Optional, String) The description of the access group template.
  * Constraints: The maximum length is `250` characters. The minimum length is `0` characters. The value must match regular expression `/^[a-zA-Z0-9!@#$%^&*()_+{}:;"'<>,.?\/|\\-\\s]+$/`.
* `group` - (Optional, List) Access Group Component.
Nested schema for **group**:
	* `action_controls` - (Optional, List) Access group action controls component.
	Nested schema for **action_controls**:
		* `access` - (Optional, List) Control whether or not access group administrators in child accounts can add access policies to the enterprise-managed access group in their account.
		Nested schema for **access**:
			* `add` - (Optional, Boolean) Action control for adding access policies to an enterprise-managed access group in a child account. If an access group administrator in a child account adds a policy, they can always update or remove it.
	* `assertions` - (Optional, List) Assertions Input Component.
	Nested schema for **assertions**:
		* `action_controls` - (Optional, List) Control whether or not access group administrators in child accounts can add, remove, and update dynamic rules for the enterprise-managed access group in their account. The inner level RuleActionControls override these action controls.
		Nested schema for **action_controls**:
			* `add` - (Optional, Boolean) Action control for adding dynamic rules to an enterprise-managed access group. If an access group administrator in a child account adds a dynamic rule, they can always update or remove it.
			* `remove` - (Optional, Boolean) Action control for removing enterprise-managed dynamic rules in an enterprise-managed access group.
		* `rules` - (Optional, List) Dynamic rules to automatically add federated users to access groups based on specific identity attributes.
		  * Constraints: The maximum length is `50` items. The minimum length is `0` items.
		Nested schema for **rules**:
			* `action_controls` - (Optional, List) Control whether or not access group administrators in child accounts can update and remove this dynamic rule in the enterprise-managed access group in their account.This overrides outer level AssertionsActionControls.
			Nested schema for **action_controls**:
				* `remove` - (Optional, Boolean) Action control for removing this enterprise-managed dynamic rule.
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
		* `services` - (Optional, List) Array of service IDs' IAM ID to add to the template.
		  * Constraints: The list items must match regular expression `/^[a-zA-Z0-9_-]+$/`. The maximum length is `50` items. The minimum length is `0` items.
		* `users` - (Optional, List) Array of enterprise users' IAM ID to add to the template. All enterprise users that you add to the template must be invited to the child accounts where the template is assigned.
		  * Constraints: The list items must match regular expression `/^[a-zA-Z0-9!@#$%^&*()_+{}:;"'<>,.?\/|\\-\\s]+$/`. The maximum length is `50` items. The minimum length is `0` items.
	* `name` - (Required, String) Give the access group a unique name that doesn't conflict with other templates access group name in the given account. This is shown in child accounts.
	  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9!@#$%^&*()_+{}:;"'<>,.?\/|\\-\\s]+$/`.
* `name` - (Optional, String) The name of the access group template.
  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9!@#$%^&*()_+{}:;"'<>,.?\/|\\-\\s]+$/`.
* `policy_template_references` - (Optional, List) References to policy templates assigned to the access group template.
  * Constraints: The maximum length is `100` items. The minimum length is `0` items.
Nested schema for **policy_template_references**:
	* `id` - (Optional, String) Policy template ID.
	  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9_-]+$/`.
	* `version` - (Optional, String) Policy template version.
	  * Constraints: The maximum length is `2` characters. The minimum length is `1` character. The value must match regular expression `/^[0-9]+$/`.
* `template_id` - (Required, Forces new resource, String) ID of the template that you want to create a new version of.
  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9_-]+$/`.
* `transaction_id` - (Optional, String) An optional transaction id for the request.
  * Constraints: The maximum length is `50` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9_-]+$/`.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the iam_access_group_template.
* `name` - (String) The name of the access group template.
* `description` - (String) The description of the access group template.
* `account_id` - (String) Enterprise account id.
* `version` - (String) The version of the access group template.
* `committed` - (Boolean) A boolean indicating whether the access group template is committed. You must commit a template before you can assign it to child accounts.
* `group` - (List) Access Group Component.
Nested schema for **group**:
	* `name` - (String) Give the access group a unique name that doesn't conflict with other templates access group name in the given account. This is shown in child accounts.
	* `description` - (String) Access group description. This is shown in child accounts.
	* `members` - (List) Array of enterprise users to add to the template. All enterprise users that you add to the template must be invited to the child accounts where the template is assigned.
    Nested schema for **members**:
        * `users` - (List) Array of enterprise users to add to the template. All enterprise users that you add to the template must be invited to the child accounts where the template is assigned.
        * `services` - (List) Array of service IDs to add to the template.
        * `action_controls` - (List) Control whether or not access group administrators in child accounts can add and remove members from the enterprise-managed access group in their account.
    	Nested schema for **action_controls**:
    		* `add` - (Boolean) Action control for adding child account members to an enterprise-managed access group. If an access group administrator in a child account adds a member, they can always remove them.
    		* `remove` - (Boolean) Action control for removing enterprise-managed members from an enterprise-managed access group.
    * `assertions` - (List) Assertions Input Component.
    Nested schema for **assertions**:
    	* `action_controls` - (List) Control whether or not access group administrators in child accounts can add, remove, and update dynamic rules for the enterprise-managed access group in their account. The inner level RuleActionControls override these action controls.
    	Nested schema for **action_controls**:
    		* `add` - (Boolean) Action control for adding dynamic rules to an enterprise-managed access group. If an access group administrator in a child account adds a dynamic rule, they can always update or remove it.
    		* `remove` - (Boolean) Action control for removing enterprise-managed dynamic rules in an enterprise-managed access group.
    	* `rules` - (List) Dynamic rules to automatically add federated users to access groups based on specific identity attributes.
    	Nested schema for **rules**:
    		* `action_controls` - (List) Control whether or not access group administrators in child accounts can update and remove this dynamic rule in the enterprise-managed access group in their account.This overrides outer level AssertionsActionControls.
    		Nested schema for **action_controls**:
    			* `remove` - (Boolean) Action control for removing this enterprise-managed dynamic rule.
    		* `conditions` - (List) Conditions of membership. You can think of this as a key:value pair.
    		Nested schema for **conditions**:
    			* `claim` - (String) The key in the key:value pair.
    			* `operator` - (String) Compares the claim and the value.
    			* `value` - (String) The value in the key:value pair.
	* `action_controls` - (List) Access group action controls component.
	Nested schema for **action_controls**:
		* `access` - (List) Control whether or not access group administrators in child accounts can add access policies to the enterprise-managed access group in their account.
		Nested schema for **access**:
			* `add` - (Boolean) Action control for adding access policies to an enterprise-managed access group in a child account. If an access group administrator in a child account adds a policy, they can always update or remove it.
* `policy_template_references` - (List) References to policy templates assigned to the access group template.
Nested schema for **policy_template_references**:
	* `id` - (String) Policy template ID.
	* `version` - (String) Policy template version.
* `href` - (String) The URL of the access group template resource.
* `created_at` - (String) The date and time when the access group template was created.
* `created_by_id` - (String) The ID of the user who created the access group template.
* `last_modified_at` - (String) The date and time when the access group template was last modified.
* `last_modified_by_id` - (String) The ID of the user who last modified the access group template.

## Import

You can import the `ibm_iam_access_group_template_version` resource by using `id`.
The `id` property can be formed from `template_id`, and `version` in the following format: `<template_id>/<version>`

* `template_id`: A string. ID of the template that you want to create a new version of.
* `version`: A string. version number in path.

### Syntax

```bash
$ terraform import ibm_iam_access_group_template_version.iam_access_group_template_version $template_id/$version
```
