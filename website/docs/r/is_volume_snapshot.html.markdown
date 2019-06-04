---
layout: "ibm"
page_title: "IBM : volume_snapshot"
sidebar_current: "docs-ibm-resource-is-volume-snapshot"
description: |-
  Manages IBM Volume Snapshot.
---

# ibm\_is_volume_snapshot

Provides a volume snapshot resource. This allows volume snapshot to be created, updated, and cancelled.


## Example Usage

In the following example, you can create a volume snapshots:

```hcl
resource "ibm_is_volume" "testacc_volume" {
    name 		= "test_volume"
    type 		= "boot"
    zone 		= "test"
    iops 		= 10000
    capacity    = 100
    auto_delete = true
}
resource "ibm_is_volume_snapshot" "testacc_volume_snapshot" {
    name 		= "test_volume"
    volume_id   = "${ibm_is_volume.testacc_volume.id}"
}

```

## Argument Reference

The following arguments are supported:

* `name` - (Required, string) The user-defined name for this volume.
* `volume_id` - (Required, string) The volume id.
* `resource_group` - (Optional, string) The resource group for this volume snapshot.

## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier of the volume snapshot.The id is composed of \<volume_id\>/\<volume_snapshot_id\>.
* `status` - The status of volume snapshot.