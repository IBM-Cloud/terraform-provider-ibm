---
layout: "ibm"
page_title: "IBM : volume"
sidebar_current: "docs-ibm-resource-is-volume"
description: |-
  Manages IBM Volume.
---

# ibm\_is_volume

Provides a volume resource. This allows volume to be created, updated, and cancelled.


## Example Usage

In the following example, you can create a volume:

```hcl
resource "ibm_is_volume" "testacc_volume" {
  name     = "test_volume"
  profile  = "10iops-tier"
  zone     = "us-south-1"
  iops     = 10000
  capacity = 100
}

```

## Timeouts

ibm_is_volume provides the following [Timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) configuration options:

* `create` - (Default 60 minutes) Used for Creating Instance.
* `delete` - (Default 60 minutes) Used for Deleting Instance.


## Argument Reference

The following arguments are supported:

* `name` - (Required, string) The user-defined name for this volume.
* `profile` - (Required, Forces new resource, string) The profile to use for this volume.
* `zone` - (Required, Forces new resource, string) The location of the volume.
* `iops` - (Optional, Forces new resource, int) The bandwidth for the volume.
* `capacity` - (Optional, Forces new resource, int) The capacity of the volume in gigabytes. This defaults to `100`.
* `encryption_key` - (Optional, Forces new resource, string) The key to use for encrypting this volume.
* `resource_group` - (Optional, Forces new resource, string) The resource group ID for this volume.
* `resource_controller_url` - The URL of the IBM Cloud dashboard that can be used to explore and view details about this instance.


## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier of the volume.
* `status` - The status of volume.
* `crn` - The CRN for the volume.