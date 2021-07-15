---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : snapshots"
description: |-
  Reads IBM Cloud snapshots.
---
# ibm\_is_snapshots

Import the details of an existing IBM Cloud Infrastructure snapshot collection as a read-only data source. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax, see [viewing snapshots](https://cloud.ibm.com/docs/vpc?topic=vpc-snapshots-vpc-view).


## Example Usage

```terraform

data "ibm_is_snapshots" "ds_snapshots" {
}

```

## Attribute reference
In addition to all argument reference list, you can access the following attribute references after your data source is created.

- `snapshots` - (List) List of snapshots in the IBM Cloud Infrastructure.
  
  Nested scheme for `snapshots`:
  - `id` - (String) The unique identifier for this snapshot.
  - `bootable` - (Bool) Indicates if a boot volume attachment can be created with a volume created from this snapshot.
  - `crn` - (String) The CRN for this snapshot.
  - `encryption` - (String) The type of encryption used on the source volume(One of [ **provider_managed**, **user_managed** ]).
  - `href` - (String) The URL for this snapshot.
  - `lifecycle_state` - (String) The lifecycle state of this snapshot. Supported values are **deleted**, **deleting**, **failed**, **pending**, **stable**, **updating**, **waiting**, **suspended**.
  - `minimum_capacity` - (Integer) The minimum capacity of a volume created from this snapshot. When a snapshot is created, this will be set to the capacity of the source_volume.
  - `operating_system` - (String) The globally unique name for the operating system included in this image.  
  - `resource_type` - (String) The resource type.
  - `size` - (Integer) The size of this snapshot rounded up to the next gigabyte.
  - `source_image` - (String) If present, the unique identifier for the image from which the data on this volume was most directly provisioned.

