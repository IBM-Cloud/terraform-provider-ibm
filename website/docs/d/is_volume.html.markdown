---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : is_volume"
description: |-
  Manages IBM Cloud VPC volume.
---

# ibm_is_volume
Retrieve information of an existing IBM Cloud VSI volume. For more information, about the volume concepts, see [expandable volume concepts for VPC](https://cloud.ibm.com/docs/vpc?topic=vpc-expanding-block-storage-volumes#expandable-volume-concepts).


## Example usage

```terraform
resource "ibm_is_volume" "testacc_volume"{
    name = "testvol"
    profile = "10iops-tier"
    zone = "us-south-1"
}
data "ibm_is_volume" "testacc_dsvol" {
    name = ibm_is_volume.testacc_volume.name
}

```

## Argument reference
Review the argument references that you can specify for your data source. 

- `name` - (Required, String) The name of the volume.
- `zone` - (Optional, String) The zone of the volume.

## Attribute reference
In addition to all argument reference list, you can access the following attribute references after your data source is created.

- `access_tags`  - (String) Access management tags associated for the instance.
- `capacity` - (String) The capacity of the volume in gigabytes.
- `encryption_key` - (String) The key to use for encrypting this volume.
- `iops` - (String) The bandwidth for the volume.
- `profile` - (String) The profile to use for this volume.
- `resource_group` - (String) The resource group ID for this volume.
- `source_snapshot` - ID of the snapshot, if volume was created from it.
- `status` - (String) The status of the volume. Supported values are **available**, **failed**, **pending**, **unusable**, **pending_deletion**.
- `status_reasons` - (List) Array of reasons for the current status.
  
  Nested scheme for `status_reasons`:
  - `code` - (String)  A snake case string identifying the status reason.
  - `message` - (String)  An explanation of the status reason
- `tags` - (String) Tags associated with the volume.
