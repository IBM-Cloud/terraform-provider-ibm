---
subcategory: "Enterprise Management"
layout: "ibm"
page_title: "IBM : enterprise"
sidebar_current: "docs-ibm-resource-enterprise"
description: |-
  Manages enterprise.
---

# ibm\_enterprise

Provides a resource for enterprise. This allows enterprise to be created and updated. Delete operation is not supported.

## Example Usage

```hcl
resource "ibm_enterprise" "enterprise" {
  source_account_id = "source_account_id"
  name = "name"
  primary_contact_iam_id = "primary_contact_iam_id"
}
```

## Argument Reference

The following arguments are supported:

* `source_account_id` - (Required, string) The ID of the account that is used to create the enterprise.
* `name` - (Required, string) The name of the enterprise. This field must have 3 - 60 characters.
* `primary_contact_iam_id` - (Required, string) The IAM ID of the enterprise primary contact, such as `IBMid-0123ABC`. The IAM ID must already exist.
* `domain` - (Optional, string) A domain or subdomain for the enterprise, such as `example.com` or `my.example.com`.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The unique identifier of the enterprise.
* `url` - The URL of the enterprise.
* `enterprise_account_id` - The enterprise account ID.
* `crn` - The Cloud Resource Name (CRN) of the enterprise.
* `state` - The state of the enterprise.
* `primary_contact_email` - The email of the primary contact of the enterprise.
* `created_at` - The time stamp at which the enterprise was created.
* `created_by` - The IAM ID of the user or service that created the enterprise.
* `updated_at` - The time stamp at which the enterprise was last updated.
* `updated_by` - The IAM ID of the user or service that updated the enterprise.
