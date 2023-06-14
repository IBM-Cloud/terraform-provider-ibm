---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : snapshot clones"
description: |-
  Reads IBM Cloud snapshot clones.
---
# ibm_is_snapshot_clones

Import the details of an existing IBM Cloud infrastructure snapshot's clone collection as a read-only data source. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax, see [viewing snapshots](https://cloud.ibm.com/docs/vpc?topic=vpc-snapshots-vpc-view).


## Example usage

```terraform

data "ibm_is_snapshot_clones" "ds_snapshotclones" {
  snapshot = "6284-8230x-1234-33ae"
}

```


## Argument reference
Review the argument references that you can specify for your data source. 

- `snapshot` - (Required, String) The unique identifier of the snapshot.


## Attribute reference
In addition to all argument reference list, you can access the following attribute references after your data source is created.

- `clones` - (List) List of snapshots in the IBM Cloud Infrastructure.
  
  Nested scheme for `clones`:
  - `available` - (Bool) Indicates whether this snapshot clone is available for use.
  - `created_at` - (String) The date and time that this snapshot clone was created.
  - `zone` - (String) The zone this snapshot clone resides in.

