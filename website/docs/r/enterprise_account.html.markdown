---
subcategory: "Enterprise Management"
layout: "ibm"
page_title: "IBM : enterprise_account"
sidebar_current: "docs-ibm-resource-enterprise-account"
description: |-
  Manages enterprise_account.
---

# ibm\_enterprise_account

Provides a resource for enterprise_account. This allows enterprise_account to be created, updated and imported. Delete operation is not supported.

## Example Usage

```hcl
resource "ibm_enterprise_account" "enterprise_account" {
  parent = "parent"
  name = "name"
  owner_iam_id = "owner_iam_id"
}

resource "ibm_enterprise_account" "enterprise_import_account"{
  parent = "parent"
  enterprise_id = "enterprise_id"
  account_id = "account_id"
}
```


## Argument Reference

The following arguments are supported to create a new account in enterprise:

* `parent` - (Required, string) The CRN of the parent under which the account will be created. The parent can be an existing account group or the enterprise itself.
* `name` - (Required, string) The name of the account. This field must have 3 - 60 characters.
* `owner_iam_id` - (Required, string) The IAM ID of the account owner, such as `IBMid-0123ABC`. The IAM ID must already exist.

The following arguments are supported to import a new account in enterprise:

* `parent` - (Required, string) The CRN of the parent under which the account will be created. The parent can be an existing account group or the enterprise itself.
* `enterprise_id` - (Required, string) The enterprise ID where the account should be imported to.
* `account_id` - (Required, string) The standalone account ID which needs to be imported, such as `521ac39afd1b40aaad96fde2c6ad97xx`


## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The unique identifier of the enterprise_account.
* `account_id` - The source account ID.
* `url` - The URL of the account.
* `crn` - The Cloud Resource Name (CRN) of the account.
* `enterprise_account_id` - The enterprise account ID.
* `enterprise_id` - The enterprise ID that the account is a part of.
* `enterprise_path` - The path from the enterprise to this particular account.
* `state` - The state of the account.
* `paid` - The type of account - whether it is free or paid.
* `owner_email` - The email address of the owner of the account.
* `is_enterprise_account` - The flag to indicate whether the account is an enterprise account or not.
* `created_at` - The time stamp at which the account was created.
* `created_by` - The IAM ID of the user or service that created the account.
* `updated_at` - The time stamp at which the account was last updated.
* `updated_by` - The IAM ID of the user or service that updated the account.
