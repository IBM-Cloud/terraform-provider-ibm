---
layout: "ibm"
page_title: "IBM : snapshot"
sidebar_current: "docs-ibm-resource-is-snapshot"
description: |-
  Manages IBM snapshot.
---

# ibm\_is_snapshot

Provides a snapshot resource. This allows snapshot to be created, updated, and cancelled.


## Example Usage

```hcl
resource "ibm_is_snapshot" "testacc_snapshot" {
  name            = "test_snapshot"
  source_volume   = r006-1772e102-0671-48c7-a97a-504247e61e4
}
```


## Argument Reference

The following arguments are supported:


* `name` - (Optional, Forces new resource, string)   The name of the snapshot.
* `source_volume` - (Required, string) The unique identifier for the volume for which snaphot is to be created. 
* `resource_group` - (Optional, Forces new resource, string) The resource group ID where the Snapshot to be created

## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier for this snapshot.
* `bootable` - Indicates if a boot volume attachment can be created with a volume created from this snapshot.
* `crn` - The CRN for this snapshot.
* `delatable` - Indicates whether this snapshot can be deleted. This value will not be true if any other snapshots depend on it.
* `encryption` - The type of encryption used on the source volume(One of [ provider_managed, user_managed ]).
* `href` - The URL for this snapshot.
* `lifecycle_state` - The lifecycle state of this snapshot (One of [ deleted, deleting, failed, pending, stable, updating, waiting, suspended ]).
* `minimum_capacity` - The minimum capacity of a volume created from this snapshot. When a snapshot is created, this will be set to the capacity of the source_volume.
* `resource_type` - The resource type.
* `size` - The size of this snapshot rounded up to the next gigabyte.

## Import

ibm_is_snapshot can be imported using ID, eg

```
$ terraform import ibm_is_snapshot.example d7bec597-4726-451f-8a63-e62e6f19c32c
```
