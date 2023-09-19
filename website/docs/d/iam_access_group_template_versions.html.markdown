---
layout: "ibm"
page_title: "IBM : ibm_iam_access_group_template_version"
description: |-
  Get information about ibm_iam_access_group_template_version
subcategory: "Identity & Access Management (IAM)"
---

# ibm_iam_access_group_template_version

Provides a read-only data source to retrieve information about an ibm_iam_access_group_template_version. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_iam_access_group_template_versions" "iam_access_group_template_version_instance" {
	template_id = ibm_iam_access_group_template_versions.iam_access_group_template_version_instance.template_id
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `template_id` - (Required, Forces new resource, String) ID of the template that you want to list all versions of.
  * Constraints: The maximum length is `50` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9_-]+$/`.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

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
