---
subcategory: "Enterprise Management"
layout: "ibm"
page_title: "IBM : enterprise"
sidebar_current: "docs-ibm-resource-enterprise"
description: |-
  Manages enterprise.
---

# ibm_enterprise

Create and update an enterprise. Delete operation is not supported. For more information, about enterprise management, refer to [setting up an enterprise](https://cloud.ibm.com/docs/account?topic=account-create-enterprise).

## Example usage

```terraform
resource "ibm_enterprise" "enterprise" {
  source_account_id = "source_account_id"
  name = "name"
  primary_contact_iam_id = "primary_contact_iam_id"
}
```

## Argument reference

Review the argument reference that you can specify for your resource. 

- `domain` - (Optional, String) A domain or subdomain for an enterprise, such as `example.com`, or `my.example.com`.
- `name` - (Required, String) The name of an enterprise. The minimum and maximum character should be from `3 to 60` characters.
- `primary_contact_iam_id` - (Required, String) The IAM ID of an enterprise primary contact, such as `IBMid-0123ABC.` The IAM ID must already exist.
- `source_account_id` - (Required, String) The ID of an account that is used to create the enterprise.


## Attribute reference

In addition to all argument reference list, you can access the following attribute references after your resource is created. 

- `crn` - (String) The Cloud Resource Name (CRN) of an enterprise.
- `created_at`  - (Timestamp) The time stamp at which an enterprise is created.
- `created_by` - (String) The IAM ID of an user or service that created an enterprise.
- `enterprise_account_id` - (String) The enterprise account ID.
- `id` - (String) The unique identifier of an enterprise.
- `state` - (String) The state of an enterprise.
- `primary_contact_email` - (String) The Email of the primary contact of an enterprise.
- `updated_at` - (Timestamp) The time stamp at which an enterprise was last updated.
- `updated_by` - (String) The IAM ID of the user or service that updated an enterprise.
- `url` - (String) The URL of an enterprise.

## Import

The `ibm_enterprise` resource can be imported by using enterprise_id.

**Example**

```
$  terraform import ibm_enterprise.enterprise_example c117bf3cb7a448fca830645865e3f1f2

```
