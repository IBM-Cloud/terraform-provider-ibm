---
layout: "ibm"
page_title: "IBM : enterprise_account_group"
sidebar_current: "docs-ibm-resource-enterprise-account-group"
description: |-
  Manages enterprise_account_group.
---

# ibm\_enterprise_account_group

Provides a resource for enterprise_account_group. This allows enterprise_account_group to be created and updated. Delete operation is not supported.

## Example Usage

```hcl
resource "ibm_enterprise_account_group" "enterprise_account_group" {
  parent = "parent"
  name = "name"
  primary_contact_iam_id = "primary_contact_iam_id"
}
```

## Argument Reference

The following arguments are supported:

* `parent` - (Required, string) The CRN of the parent under which the account group will be created. The parent can be an existing account group or the enterprise itself.
* `name` - (Required, string) The name of the account group. This field must have 3 - 60 characters.
* `primary_contact_iam_id` - (Required, string) The IAM ID of the primary contact for this account group, such as `IBMid-0123ABC`. The IAM ID must already exist.

## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier of the enterprise_account_group.
* `url` - The URL of the account group.
* `crn` - The Cloud Resource Name (CRN) of the account group.
* `enterprise_account_id` - The enterprise account ID.
* `enterprise_id` - The enterprise ID that the account group is a part of.
* `enterprise_path` - The path from the enterprise to this particular account group.
* `state` - The state of the account group.
* `primary_contact_email` - The email address of the primary contact of the account group.
* `created_at` - The time stamp at which the account group was created.
* `created_by` - The IAM ID of the user or service that created the account group.
* `updated_at` - The time stamp at which the account group was last updated.
* `updated_by` - The IAM ID of the user or service that updated the account group.
