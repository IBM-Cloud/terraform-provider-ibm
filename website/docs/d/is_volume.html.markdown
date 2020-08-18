---
layout: "ibm"
page_title: "IBM : volume"
sidebar_current: "docs-ibm-datasource-is-volume"
description: |-
  Manages IBM Volume.
---

# ibm\_is_volume

Provides a vpc volume datasource. This allows to fetch an existing volume.


## Example Usage

```hcl
resource "ibm_is_volume" "testacc_volume"{
    name = "testvol"
    profile = "10iops-tier"
    zone = "us-south-1"
}
data "ibm_is_volume" "testacc_dsvol" {
    name = ibm_is_volume.testacc_volume.name
}

```

## Argument Reference

The following arguments are supported:

* `name` - (Required, string) The name of the volume.
* `zone` - (Optional, string) The zone of the volume.

## Attribute Reference

The following attributes are exported:

* `profile` - The profile to use for this volume.
* `iops` - The bandwidth for the volume.
* `capacity` - The capacity of the volume in gigabytes.
* `encryption_key` - The key to use for encrypting this volume.
* `resource_group` - The resource group ID for this volume.
* `tags` - Tags associated with the volume.
