---
subcategory: "Enterprise Management"
layout: "ibm"
page_title: "IBM : enterprises"
description: |-
  Get information about the enterprises
---

# ibm_enterprises

Retrieve an information from an `ibm_enterprise` data source. For more information, about enterprise management, see [setting up an enterprise](https://cloud.ibm.com/docs/account?topic=account-create-enterprise).


## Example usage

```terraform
data "ibm_enterprises" "enterprises" {
}
```

## Argument reference
Review the argument reference that you can specify for your data source. 

- `name`  - (Optional, String) The name of an enterprise.

## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your data source is created.

- `enterprises` - (String) A list of enterprise objects. Nested `resources` blocks has the following structure.

  Nested scheme for `enterprises`:
  - `crn` - (String) The Cloud Resource Name (CRN) of an enterprise.
  - `created_at` - (Timestamp) The time stamp at which an enterprise is created.
  - `created_by` - (String) The IAM ID of an user or service that created an enterprise.
  - `domain` - (String) The domain of an enterprise.
  - `enterprise_account_id` - (String) The enterprise account ID.
  - `id` - (String) The enterprise ID.
  - `name` - (String) The name of an enterprise.
  - `primary_contact_iam_id` - (String) The IAM ID of the primary contact of an enterprise, such as `IBMid-0123ABC`.
  - `primary_contact_email` - (String) The Email address of the primary contact of an enterprise.
  - `state` - (String) The state of an enterprise.
  - `updated_at` - (Timestamp) The time stamp at which an enterprise was last updated.
  - `updated_by` - (String) The IAM ID of the user or service that updated an enterprise.
  - `url` - (String) The URL of an enterprise.
- `id` - (String) The unique identifier of an enterprises.

