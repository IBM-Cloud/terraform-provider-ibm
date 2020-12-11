---
layout: "ibm"
page_title: "IBM : atracker_target"
sidebar_current: "docs-ibm-resource-atracker-target"
description: |-
  Manages ATracker Target.
---

# ibm\_atracker_target

Provides a resource for ATracker Target. This allows ATracker Target to be created, updated and deleted.

## Example Usage

```hcl
resource "atracker_target" "atracker_target" {
  name = "my-cos-target"
  target_type = "example"
  cos_endpoint = { example: "object" }
}
```

## Timeouts

atracker_target provides the following [Timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) configuration options:

* `create` - (Default 60 minutes) Used for creating a ATracker Target.
* `update` - (Default 20 minutes) Used for updating a ATracker Target.
* `delete` - (Default 10 minutes) Used for deleting a ATracker Target.

## Argument Reference

The following arguments are supported:

* `name` - (Required, string) The name of the target. Must be 256 characters or less.
* `target_type` - (Required, Forces new resource, string) The type of the target.
* `cos_endpoint` - (Required, List) Property values for a Cloud Object Storage Endpoint.
  * `endpoint` - (Required, string) The host name of this COS endpoint.
  * `target_crn` - (Required, string) The CRN of this COS instance.
  * `bucket` - (Required, string) The bucket name under this COS instance.
  * `api_key` - (Required, string) The IAM Api key that have writer access to this cos instance. This credential will be masked in the response.

## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier of the ATracker Target.
* `instance_id` - The uuid of ATracker services in this region.
* `crn` - The crn of this target type resource.
* `encrypt_key` - The encryption key used to encrypt events before ATracker services buffer them on storage. This credential will be masked in the response.
