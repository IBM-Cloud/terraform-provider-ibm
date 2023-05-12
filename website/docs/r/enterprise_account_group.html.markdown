---
subcategory: "Enterprise Management"
layout: "ibm"
page_title: "IBM : enterprise_account_group"
sidebar_current: "docs-ibm-resource-enterprise-account-group"
description: |-
  Manages an enterprise account group.
---

# ibm_enterprise_account_group

Create and update and delete `enterprise_account_group`resource. For more information, about enterprise account group, refer to [setting up access groups](https://cloud.ibm.com/docs/account?topic=account-groups).

## Example usage

```terraform
resource "ibm_enterprise_account_group" "enterprise_account_group" {
  parent = "parent"
  name = "name"
  primary_contact_iam_id = "primary_contact_iam_id"
}
```

## Argument reference
Review the argument reference that you can specify for your resource. 

- `name` - (Required, String) The name of an enterprise. The minimum and maximum character should be from `3 to 60` characters.
- `parent` - (Required, String) The CRN of the parent in which the account group is created. The parent can be an existing account group or an enterprise itself.
- `primary_contact_iam_id` - (Required, String) The IAM ID of an enterprise primary contact, such as `IBMid-0123ABC.` The IAM ID must already exist.

## Attribute reference
In addition to all argument reference list, you can access the following attribute references after your resource is created. 

- `created_at` - (Timestamp) The time stamp at which an account group is created.
- `created_by` - (String) The IAM ID of an user or service that created an account group.
- `crn` - (String) The Cloud Resource Name (CRN) of an account group.
- `enterprise_account_id` - (String) The enterprise account ID.
- `enterprise_id` - (String) The enterprise ID that the account group is a part of.
- `enterprise_path` - (String) The path from the enterprise to the particular account group.
- `id` - (String) The unique identifier of an enterprise account group.
- `state` - (String) The state of an account group.
- `primary_contact_email` - (String) The Email address of the primary contact of an account group.
- `updated_at` - (Timestamp) The time stamp at which an account group was last updated.
- `updated_by` - (String) The IAM ID of the user or service that updated an account group.
- `url` - (String) The URL of an account group.

## Import

The `ibm_enterprise_account_group` resource can be imported by using account_group_id.

**Example**

```
$ terraform import ibm_enterprise_account_group.example ae337d0b6cf6485a918a47e289ab4628
```
