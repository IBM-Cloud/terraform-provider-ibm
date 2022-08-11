---
subcategory: "Enterprise Management"
layout: "ibm"
page_title: "IBM : enterprise_account_groups"
description: |-
  Get information about account groups
---

# ibm_enterprise_account_groups

Retrieve an information from an `account_groups` data source.  For more information, about enterprise account groups, refer to [setting up access groups](https://cloud.ibm.com/docs/account?topic=account-groups).


## Example usage

```terraform
data "ibm_enterprise_account_groups" "account_groups" {
}
```

## Argument reference
Review the argument reference that you can specify for your data source. 

- `name` - (Optional, String) The name of an account group.

## Attribute reference
In addition to the argument reference list, you can access the following attribute reference after your data source is created. 

- `id`  - (String) The unique identifier of an account groups.
- `account_groups`  - (List)  A list of account groups. Nested `resources` blocks has the following structure.
  
  Nested scheme for `account_groups`:
  - `crn`  - (String) The Cloud Resource Name (CRN) of an account group.
  - `created_at`  - (Timestamp) The time stamp at which an account is created.
  - `created_by`  - (String) The IAM ID of an user or service that created an account group.
  - `enterprise_account_id`  - (String) The enterprise account ID.
  - `enterprise_id` - (String) The enterprise ID that an account group is a part of.
  - `enterprise_path` - (String) The path from an enterprise to the particular account group.
  - `id`  - (String) The account group ID.
  - `name` - (String) The name of an account group.
  - `parent` - (String) The CRN of the parent of an account group.
  - `primary_contact_iam_id` - (String) The IAM ID of the owner of an account group.
  - `primary_contact_email`  - (String) The Email address of the owner of an account group.
  - `state`  - (String) The state of an account group.
  - `updated_at`  - (Timestamp) The time stamp at which an account was last updated.
  - `updated_by`  - (String) The IAM ID of the user or service that updated an account group.
  - `url`  - (String) The URL of an account group.


