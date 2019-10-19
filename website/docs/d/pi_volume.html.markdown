---
layout: "ibm"
page_title: "IBM : Volume"
sidebar_current: "docs-ibm-datasources-pi-volume"
description: |-
  Manages IBM Cloud Infrastructure Volumes for IBM Power
---

# ibm\_pi_image

Import the details of an existing IBM Cloud Infrastructure Volume for IBM Power as a read-only data source. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.


## Example Usage

```hcl

data "ibm_pi_volume" "ds_volume" {
    name = "volume_1"
    powerinstanceid="49fba6c9-23f8-40bc-9899-aca322ee7d5b"

}

```

## Argument Reference

The following arguments are supported:

* `volumename` - (Required, string) The name of the volume.
* `powerinstanceid` - (Required, string) The service instance associated with the account


## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier for this volume.
* `type` - The disktype for this volume.
* `state` - The state of the volume.
* `bootable` - If this volume is bootable or not.
* `size` - The size of this volume



