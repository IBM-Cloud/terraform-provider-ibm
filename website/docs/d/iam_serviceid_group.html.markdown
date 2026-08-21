---
layout: "ibm"
page_title: "IBM : ibm_iam_serviceid_group"
description: |-
  Get information about iam_serviceid_group
subcategory: "IAM Identity Services"
---

# ibm_iam_serviceid_group

Provides a read-only data source to retrieve information about an iam_serviceid_group. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_iam_serviceid_group" "iam_serviceid_group" {
	iam_serviceid_group_id = "iam_serviceid_group_id"
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `iam_serviceid_group_id` - (Required, Forces new resource, String) Unique ID of the service ID group.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the iam_serviceid_group.
* `account_id` - (String) ID of the account the service ID group belongs to.
* `created_at` - (String) Timestamp of when the service ID group was created.
* `created_by` - (String) IAM ID of the user or service which created the Service Id group.
* `crn` - (String) Cloud Resource Name of the item.
* `description` - (String) Description of the service ID group.
* `entity_tag` - (String) Version of the service ID group details object. You need to specify this value when updating the service ID group to avoid stale updates.
* `modified_at` - (String) Timestamp of when the service ID group was modified.
* `name` - (String) Name of the service ID group. Unique in the account.

