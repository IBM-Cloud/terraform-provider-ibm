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
  cos_endpoint = { "endpoint" : "endpoint", "target_crn" : "target_crn", "bucket" : "bucket", "api_key" : "api_key" }
  name = "my-cos-target"
  target_type = "target_type"
}
```

## Timeouts

atracker_target provides the following [Timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) configuration options:

* `create` - (Default 60 minutes) Used for creating a Activity Tracking Target.
* `update` - (Default 20 minutes) Used for updating a Activity Tracking Target.
* `delete` - (Default 10 minutes) Used for deleting a Activity Tracking Target.

## Argument Reference

The following arguments are supported:

* `cos_endpoint` - (Required, List) Property values for a Cloud Object Storage Endpoint.
  * `endpoint` - (Required, string) The host name of the Cloud Object Storage endpoint.
  * `target_crn` - (Required, string) The CRN of the Cloud Object Storage instance.
  * `bucket` - (Required, string) The bucket name under the Cloud Object Storage instance.
  * `api_key` - (Required, string) The IAM API key that has writer access to the Cloud Object Storage instance. This credential is masked in the response.
* `name` - (Required, string) The name of the target. The name must be 1000 characters or less, and cannot include any special characters other than `(space) - . _ :`.
  * Constraints: The maximum length is `1000` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9 -._:]+$/`
* `target_type` - (Required, Forces new resource, string) The type of the target.
  * Constraints: Allowable values are: cloud_object_storage

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The unique identifier of the Activity Tracking Target.
* `cos_write_status` - The status of the write attempt with the provided cos_endpoint parameters.
* `created` - The timestamp of the target creation time.
* `crn` - The crn of the target resource.
* `encrypt_key` - The encryption key that is used to encrypt events before Activity Tracking services buffer them on storage. This credential is masked in the response.
* `updated` - The timestamp of the target last updated time.

## Import

You can import the `ibm_atracker_target` resource by using `id`. The uuid of the target resource.

```
$ terraform import ibm_atracker_target.atracker_target f7dcfae6-e7c5-08ca-451b-fdfa696c9bb6
```
