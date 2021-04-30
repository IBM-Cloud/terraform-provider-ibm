---

subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : volume"
description: |-
  Manages IBM Volume.
---

# ibm\_is_volume

Provides a volume resource. This allows volume to be created, updated, and cancelled.


## Example Usage

In the following example, you can create a `10iops-tier` volume:

```hcl
resource "ibm_is_volume" "testacc_volume" {
  name     = "test_volume"
  profile  = "10iops-tier"
  zone     = "us-south-1"
}

```
In the following example, you can create a `custom` volume:

```hcl
resource "ibm_is_volume" "testacc_volume" {
  name     = "test_volume"
  profile  = "custom"
  zone     = "us-south-1"
  iops     = 1000
  capacity = 200
}

```

## Timeouts

ibm_is_volume provides the following [Timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) configuration options:

* `create` - (Default 10 minutes) Used for Creating Instance.
* `delete` - (Default 10 minutes) Used for Deleting Instance.


## Argument Reference

The following arguments are supported:

* `name` - (Required, string) The user-defined name for this volume.
* `profile` - (Required, Forces new resource, string) The profile to use for this volume.
* `zone` - (Required, Forces new resource, string) The location of the volume.
* `iops` - (Optional, Forces new resource, int) The bandwidth for the volume. This is required only for the `custom` profile volume.
* `capacity` - (Optional, Forces new resource, int) The capacity of the volume in gigabytes. This defaults to `100`.
* `encryption_key` - (Optional, Forces new resource, string) The key to use for encrypting this volume.
* `resource_group` - (Optional, Forces new resource, string) The resource group ID for this volume.
* `tags` - (Optional, array of strings) Tags associated with the volume.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The unique identifier of the volume.
* `status` - The status of volume.
* `crn` - The CRN for the volume.
* `status` - The status of the volume. One of [ available, failed, pending, unusable, pending_deletion ].
* `status_reasons` - Array of reasons for the current status
  * `code` - A snake case string succinctly identifying the status reason
  * `message` - An explanation of the status reason

## Import

ibm_is_volume can be imported using volume ID, eg

```
$ terraform import ibm_is_volume.example d7bec597-4726-451f-8a63-e62e6f19c32c
```