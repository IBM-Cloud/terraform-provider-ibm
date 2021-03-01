---
layout: "ibm"
page_title: "IBM : iam_custom_role"
sidebar_current: "docs-ibm-datasource-iam-custom-role"
description: |-
  Get information about iam_custom_role
---

# ibm\_iam_custom_role

Provides a read-only data source for iam_custom_role. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "iam_custom_role" "iam_custom_role" {
	role_id = "role_id"
}
```

## Argument Reference

The following arguments are supported:

* `role_id` - (Required, string) The role ID.

## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier of the iam_custom_role.
* `id` - The role ID.

* `display_name` - The display name of the role that is shown in the console.

* `description` - The description of the role.

* `actions` - The actions of the role.

* `crn` - The role CRN.

* `name` - The name of the role that is used in the CRN. Can only be alphanumeric and has to be capitalized.

* `account_id` - The account GUID.

* `service_name` - The service name.

* `created_at` - The UTC timestamp when the role was created.

* `created_by_id` - The iam ID of the entity that created the role.

* `last_modified_at` - The UTC timestamp when the role was last modified.

* `last_modified_by_id` - The iam ID of the entity that last modified the policy.

* `href` - The href link back to the role.

