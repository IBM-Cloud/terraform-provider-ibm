---
subcategory: "Enterprise Management"
layout: "ibm"
page_title: "IBM : enterprises"
description: |-
  Get information about enterprises
---

# ibm\_enterprises

Provides a read-only data source for enterprises. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "ibm_enterprises" "enterprises" {
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional, string) The name of the enterprise.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The unique identifier of the enterprises.

* `enterprises` - A list of enterprise objects. Nested `enterprises` blocks have the following structure:
	* `url` - The URL of the enterprise.
	* `id` - The enterprise ID.
	* `enterprise_account_id` - The enterprise account ID.
	* `crn` - The Cloud Resource Name (CRN) of the enterprise.
	* `name` - The name of the enterprise.
	* `domain` - The domain of the enterprise.
	* `state` - The state of the enterprise.
	* `primary_contact_iam_id` - The IAM ID of the primary contact of the enterprise, such as `IBMid-0123ABC`.
	* `primary_contact_email` - The email of the primary contact of the enterprise.
	* `created_at` - The time stamp at which the enterprise was created.
	* `created_by` - The IAM ID of the user or service that created the enterprise.
	* `updated_at` - The time stamp at which the enterprise was last updated.
	* `updated_by` - The IAM ID of the user or service that updated the enterprise.

