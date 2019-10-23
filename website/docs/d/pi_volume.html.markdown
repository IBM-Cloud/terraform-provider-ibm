---
layout: "ibm"
page_title: "IBM: pi_volume"
sidebar_current: "docs-ibm-datasources-pi-volume"
description: |-
  Manages a volume in the Power Virtual Server Cloud.
---

# ibm\_pi_volume

Import the details of an existing IBM Power Virtual Server Cloud volume as a read-only data source. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "ibm_pi_volume" "ds_volume" {
  pi_volume_name       = "volume_1"
  pi_cloud_instance_id = "49fba6c9-23f8-40bc-9899-aca322ee7d5b"
}
```

## Argument Reference

The following arguments are supported:

* `pi_volume_name` - (Required, string) The name of the volume.
* `pi_cloud_instance_id` - (Required, string) The service instance associated with the account

## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier for this volume.
* `type` - The disk type for this volume.
* `state` - The state of the volume.
* `bootable` - If this volume is bootable or not.
* `size` - The size of this volume.
