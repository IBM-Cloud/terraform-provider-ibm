---
subcategory: "Enterprise Management"
layout: "ibm"
page_title: "IBM : enterprise_account_groups"
description: |-
  Get information about account_groups
---

# ibm\_enterprise_account_groups

Provides a read-only data source for account_groups. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "ibm_enterprise_account_groups" "account_groups" {
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional, string) The name of the account group.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The unique identifier of the account_groups.

* `account_groups` - A list of account groups. Nested `account_groups` blocks have the following structure:
	* `url` - The URL of the account group.
	* `id` - The account group ID.
	* `crn` - The Cloud Resource Name (CRN) of the account group.
	* `parent` - The CRN of the parent of the account group.
	* `enterprise_account_id` - The enterprise account ID.
	* `enterprise_id` - The enterprise ID that the account group is a part of.
	* `enterprise_path` - The path from the enterprise to this particular account group.
	* `name` - The name of the account group.
	* `state` - The state of the account group.
	* `primary_contact_iam_id` - The IAM ID of the primary contact of the account group.
	* `primary_contact_email` - The email address of the primary contact of the account group.
	* `created_at` - The time stamp at which the account group was created.
	* `created_by` - The IAM ID of the user or service that created the account group.
	* `updated_at` - The time stamp at which the account group was last updated.
	* `updated_by` - The IAM ID of the user or service that updated the account group.

