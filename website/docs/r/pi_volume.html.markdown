---
layout: "ibm"
page_title: "IBM: pi_volume"
sidebar_current: "docs-ibm-resource-pi-volume"
description: |-
  Manages IBM Volume in the Power Virtual Server Cloud.
---

# ibm\_pi_volume

Provides a volume resource. This allows volume to be created, updated, and cancelled in the Power Virtual Server Cloud.

## Example Usage

In the following example, you can create a volume:

```hcl
resource "ibm_pi_volume" "testacc_volume"{
  pi_volume_size       = 20
  pi_volume_name       = test-volume
  pi_volume_type       = ssd
  pi_volume_shareable  = true
  pi_cloud_instance_id = "<value of the cloud_instance_id>"
}
```

## Timeouts

ibm_pi_volume provides the following [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) configuration options:

* `create` - (Default 60 minutes) Used for Creating Volume.
* `delete` - (Default 60 minutes) Used for Deleting Volume.

## Argument Reference

The following arguments are supported:

* `pi_volume_size` - (Required, int) The size for this volume.
* `pi_volume_name` - (Required, string) The name of this volume.
* `pi_volume_type` - (Required, string) The volume type - supported types are (ssd/standard/tier1/tier3).
* `pi_volume_shareable` - (Optional, boolean) If the volume can be shared or not (true/false).
* `pi_cloud_instance_id` - (Required, string) The cloud_instance_id for this account.

## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier of the network.The id is composed of \<power_instance_id\>/\<volume_id\>.
* `volume_id` - The unique identifier of the volume.
* `status` - The status of the volume.

## Import

ibm_pi_volume can be imported using `power_instance_id` and `volume_id`, eg

```
$ terraform import ibm_pi_volume.example d7bec597-4726-451f-8a63-e62e6f19c32c/cea6651a-bc0a-4438-9f8a-a0770bbf3ebb
```