---
subcategory: "Enterprise Management"
layout: "ibm"
page_title: "IBM : enterprise_accounts"
description: |-
  Get information about accounts
---

# ibm\_enterprise_accounts

Provides a read-only data source for accounts. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "ibm_enterprise_accounts" "accounts" {
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional, string) The name of the account.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The unique identifier of the accounts.

* `accounts` - A list of accounts. Nested `resources` blocks have the following structure:
	* `url` - The URL of the account.
	* `id` - The account ID.
	* `crn` - The Cloud Resource Name (CRN) of the account.
	* `parent` - The CRN of the parent of the account.
	* `enterprise_account_id` - The enterprise account ID.
	* `enterprise_id` - The enterprise ID that the account is a part of.
	* `enterprise_path` - The path from the enterprise to this particular account.
	* `name` - The name of the account.
	* `state` - The state of the account.
	* `owner_iam_id` - The IAM ID of the owner of the account.
	* `paid` - The type of account - whether it is free or paid.
	* `owner_email` - The email address of the owner of the account.
	* `is_enterprise_account` - The flag to indicate whether the account is an enterprise account or not.
	* `created_at` - The time stamp at which the account was created.
	* `created_by` - The IAM ID of the user or service that created the account.
	* `updated_at` - The time stamp at which the account was last updated.
	* `updated_by` - The IAM ID of the user or service that updated the account.

