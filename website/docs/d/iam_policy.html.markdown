---
layout: "ibm"
page_title: "IBM : iam_policy"
sidebar_current: "docs-ibm-datasource-iam-policy"
description: |-
  Get information about iam_policy
---

# ibm\_iam_policy

Provides a read-only data source for iam_policy. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "iam_policy" "iam_policy" {
	policy_id = "policy_id"
}
```

## Argument Reference

The following arguments are supported:

* `policy_id` - (Required, string) The policy ID.

## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier of the iam_policy.
* `id` - The policy ID.

* `type` - The policy type; either 'access' or 'authorization'.

* `description` - Customer-defined description.

* `subjects` - The subjects associated with a policy. Nested `subjects` blocks have the following structure:
	* `attributes` - List of subject attributes. Nested `attributes` blocks have the following structure:
		* `name` - The name of an attribute.
		* `value` - The value of an attribute.

* `roles` - A set of role cloud resource names (CRNs) granted by the policy. Nested `roles` blocks have the following structure:
	* `role_id` - The role cloud resource name granted by the policy.
	* `display_name` - The display name of the role.
	* `description` - The description of the role.

* `resources` - The resources associated with a policy. Nested `resources` blocks have the following structure:
	* `attributes` - List of resource attributes. Nested `attributes` blocks have the following structure:
		* `name` - The name of an attribute.
		* `value` - The value of an attribute.
		* `operator` - The operator of an attribute.

* `href` - The href link back to the policy.

* `created_at` - The UTC timestamp when the policy was created.

* `created_by_id` - The iam ID of the entity that created the policy.

* `last_modified_at` - The UTC timestamp when the policy was last modified.

* `last_modified_by_id` - The iam ID of the entity that last modified the policy.

