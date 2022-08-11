---
subcategory: "Enterprise Management"
layout: "ibm"
page_title: "IBM : enterprise_account"
sidebar_current: "docs-ibm-resource-enterprise-account"
description: |-
  Manages an enterprise account.
---

# ibm_enterprise_account

Create and update an `enterprise_account` resource. Delete operation is not supported. For more information, about enterprise account, refer to [setting up accounts to an enterprise](https://cloud.ibm.com/docs/account?topic=account-enterprise-add).

## Example usage

```terraform
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

## Argument reference

Review the argument reference that you can specify to create a new account in an enterprise resource.

- `name` - (Required, String) The name of an enterprise. The minimum and maximum character should be from `3 to 60` characters.
- `owneriam_id` - (Required, String) The IAM ID of an account owner, such as `IBMid-0123ABC.` The IAM ID must already exist.
- `parent` - (Required, String) The CRN of the parent in which the account is created. The parent can be an existing account group or an enterprise itself.

Review the argument reference that you can specify to import a new account in an enterprise resource. 

- `account_id` - (Required, String) The stand-alone account ID that needs to be imported, such as `521ac39afd1b40aaad96fde2c6ad97xx`.
- `enterprise_id` - (Required, String) The enterprise ID where the account is imported.
- `parent` - (Required, String) The CRN of the parent in which the account is created. The parent can be an existing account group or an enterprise itself.

## Attribute reference

In addition to all argument reference list, you can access the following attribute references after your resource is created. 

- `account_id` - (String) The source account ID.
- `crn` - (String) The Cloud Resource Name (CRN) of an account.
- `created_at` - (Timestamp) The time stamp at which an account is created.
- `created_by` - (String) The IAM ID of an user or service that created an account.
- `enterprise_account_id` - (String) The enterprise account ID.
- `enterprise_id` - (String) The enterprise ID that the account is a part of.
- `enterprise_path` - (String) The path from the enterprise to the particular account.
- `id` - (String) The unique identifier of an enterprise account.
- `is_enterprise_account` - (String) The flag to indicate whether the account is an enterprise account or not.
- `owner_email` - (String) The Email address of the owner of an account.
- `paid` - (String) The type of account, whether it is `free`, or `paid`.
- `state` - (String) The state of an account.
- `updated_at` - (Timestamp) The time stamp at which an account was last updated.
- `updated_by` - (String) The IAM ID of the user or service that updated an account.
- `url` - (String) The URL of an account.

## Import

The `ibm_enterprise_account` resource can be imported by using account_group_id.

**Example**

```
$ terraform import ibm_enterprise_account.example 907ec1a69a354afc94d3a7b499d6784f
```
