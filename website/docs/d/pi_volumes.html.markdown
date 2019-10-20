---
layout: "ibm"
page_title: "IBM: Volumes"
sidebar_current: "docs-ibm-datasources-pi-volumes"
description: |-
  Manages volumes in the Power Virtual Server Cloud.
---

# ibm\_pi_volumes

Import the details of existing IBM Power Virtual Server Cloud volumes as a read-only data source. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "ibm_pi_volumes" "ds_volumes" {
    instance_name   = "volume_1"
    powerinstanceid = "49fba6c9-23f8-40bc-9899-aca322ee7d5b"
}
```

## Argument Reference

The following arguments are supported:

* `instance_name` - (Required, string) The name of the instance whose volumes to retrieve.
* `powerinstanceid` - (Required, string) The service instance associated with the account.

## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier for this volume.
* `type` - The disk type for this volume.
* `state` - The state of the volume.
* `bootable` - If this volume is bootable or not.
* `size` - The size of this volume.
* `shareable` - If this volume is shareable or not.
* `href` - The href of this volume.
* `name` - The name of this volume.
