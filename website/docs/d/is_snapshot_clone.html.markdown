---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : snapshot clone"
description: |-
  Reads IBM Cloud snapshot clone.
---
# ibm_is_snapshot_clone

Import the details of an existing IBM Cloud infrastructure snapshot's clone in a zone as a read-only data source. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax, see [viewing snapshots](https://cloud.ibm.com/docs/vpc?topic=vpc-snapshots-vpc-view).


## Example usage

```terraform

data "ibm_is_snapshot_clone" "ds_snapshotclone" {
  snapshot = "xxxx-xxxx-xxxx-xxxx-xxxx"
  zone     = "us-south-1"
}

```


## Argument reference
Review the argument references that you can specify for your data source. 

- `snapshot` - (Required, String) The unique identifier of the snapshot.
- `zone` - (Required, String) The zone in which clone resides in.


## Attribute reference
In addition to all argument reference list, you can access the following attribute references after your data source is created.


- `available` - (Bool) Indicates whether this snapshot clone is available for use.
- `created_at` - (String) The date and time that this snapshot clone was created.
- `id` - (String) The zone this snapshot clone resides in.
- `zone` - (String) The zone this snapshot clone resides in.

