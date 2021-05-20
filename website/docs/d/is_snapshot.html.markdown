---
layout: "ibm"
page_title: "IBM : Snapshots"
sidebar_current: "docs-ibm-datasources-is-snapshot"
description: |-
  Manages IBM Cloud Infrastructure Snapshots.
---

# ibm\_is_snapshots

Import the details of an existing IBM Cloud Infrastructure snapshots as a read-only data source. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.


## Example Usage

```hcl

data "ibm_is_snapshot" "ds_snapshots" {
    identifier = ibm_is_snapshot.testacc_snapshot.id
}

```

## Argument Reference

The following arguments are supported:

* `name` - (Optional, string) The name of the Snapshot.
* `identifier` - (Optional, string) The unique identifier for this snapshot.

## Attribute Reference

The following attributes are exported:

* `bootable` - Indicates if a boot volume attachment can be created with a volume created from this snapshot.
* `crn` - The CRN for this snapshot.
* `delatable` - Indicates whether this snapshot can be deleted. This value will not be true if any other snapshots depend on it.
* `encryption` - The type of encryption used on the source volume(One of [ provider_managed, user_managed ]).
* `href` - The URL for this snapshot.
* `lifecycle_state` - The lifecycle state of this snapshot (One of [ deleted, deleting, failed, pending, stable, updating, waiting, suspended ]).
* `minimum_capacity` - The minimum capacity of a volume created from this snapshot. When a snapshot is created, this will be set to the capacity of the source_volume.
* `resource_type` - The resource type.
* `size` - The size of this snapshot rounded up to the next gigabyte.
