---
layout: "ibm"
page_title: "IBM : ibm_atracker_target"
description: |-
  Manages Activity Tracking Target.
subcategory: "Activity Tracking API"
---

# ibm_atracker_target

Provides a resource for Activity Tracking Target. This allows Activity Tracking Target to be created, updated and deleted.

## Example Usage

```hcl
resource "ibm_atracker_target" "atracker_target" {
}
```

## Timeouts

atracker_target provides the following [Timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) configuration options:

* `create` - (Default 60 minutes) Used for creating a Activity Tracking Target.
* `update` - (Default 20 minutes) Used for updating a Activity Tracking Target.
* `delete` - (Default 10 minutes) Used for deleting a Activity Tracking Target.

## Argument Reference

The following arguments are supported:


## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The unique identifier of the Activity Tracking Target.
* `cos_endpoint` - Property values for a Cloud Object Storage Endpoint.
* `cos_write_status` - The status of the write attempt with the provided cos_endpoint parameters.
* `created` - The timestamp of the target creation time.
* `crn` - The crn of the target resource.
* `encrypt_key` - The encryption key that is used to encrypt events before Activity Tracking services buffer them on storage. This credential is masked in the response.
* `name` - The name of the target resource.
* `target_type` - The type of the target.
  * Constraints: Allowable values are: cloud_object_storage
* `updated` - The timestamp of the target last updated time.

## Import

You can import the `ibm_atracker_target` resource by using `id`.
The `id` property can be formed from `id`, and `id` in the following format:

```
<id>/<id>
```
* `id`: A string. The v4 UUID that uniquely identifies the target.
* `id`: A string. The v4 UUID that uniquely identifies the target.

```
$ terraform import ibm_atracker_target.atracker_target <id>/<id>
```
