---
layout: "ibm"
page_title: "IBM : ibm_iam_serviceid_group"
description: |-
  Manages iam_serviceid_group.
subcategory: "IAM Identity Services"
---

# ibm_iam_serviceid_group

Create, update, and delete iam_serviceid_groups with this resource.

## Example Usage

```hcl
resource "ibm_iam_serviceid_group" "iam_serviceid_group_instance" {
  account_id = "account_id"
  name = "name"
}
```

## Argument Reference

You can specify the following arguments for this resource.

* `account_id` - (Required, String) ID of the account the service ID group belongs to.
* `description` - (Optional, String) Description of the service ID group.
* `name` - (Required, String) Name of the service ID group. Unique in the account.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the iam_serviceid_group.
* `created_at` - (String) Timestamp of when the service ID group was created.
* `created_by` - (String) IAM ID of the user or service which created the Service Id group.
* `crn` - (String) Cloud Resource Name of the item.
* `entity_tag` - (String) Version of the service ID group details object. You need to specify this value when updating the service ID group to avoid stale updates.
* `modified_at` - (String) Timestamp of when the service ID group was modified.


## Import

You can import the `ibm_iam_serviceid_group` resource by using `id`. ID of the the service ID group.

# Syntax
<pre>
$ terraform import ibm_iam_serviceid_group.iam_serviceid_group &lt;id&gt;
</pre>
