---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : is_volume"
description: |-
  Manages IBM Cloud VPC volume.
---

# ibm_is_volume
Retrieve information of an existing IBM Cloud VSI volume. For more information, about the volume concepts, see [expandable volume concepts for VPC](https://cloud.ibm.com/docs/vpc?topic=vpc-expanding-block-storage-volumes#expandable-volume-concepts).

**Note:** 
VPC infrastructure services are a regional specific based endpoint, by default targets to `us-south`. Please make sure to target right region in the provider block as shown in the `provider.tf` file, if VPC service is created in region other than `us-south`.

**provider.tf**

```terraform
provider "ibm" {
  region = "eu-gb"
}
```

## Example usage

```terraform
resource "ibm_is_volume" "example" {
  name    = "example-volume"
  profile = "10iops-tier"
  zone    = "us-south-1"
}
data "ibm_is_volume" "example" {
  name = ibm_is_volume.example.name
}

```

## Argument reference
Review the argument references that you can specify for your data source. 

- `name` - (Required, String) The name of the volume.
- `zone` - (Optional, String) The zone of the volume.

## Attribute reference
In addition to all argument reference list, you can access the following attribute references after your data source is created.

- `bandwidth` - The maximum bandwidth (in megabits per second) for the volume
- `capacity` - (String) The capacity of the volume in gigabytes.
- `crn` - (String) The crn of this volume.
- `encryption_key` - (String) The key to use for encrypting this volume.
- `encryption_type` - (String) The type of ecryption used in the volume [**provider_managed**, **user_managed**].
- `iops` - (String) The bandwidth for the volume.
- `profile` - (String) The profile to use for this volume.
- `resource_group` - (String) The resource group ID for this volume.
- `source_snapshot` - ID of the snapshot, if volume was created from it.
- `status` - (String) The status of the volume. Supported values are **available**, **failed**, **pending**, **unusable**, **pending_deletion**.
- `status_reasons` - (List) Array of reasons for the current status.
  
  Nested scheme for `status_reasons`:
  - `code` - (String)  A snake case string identifying the status reason.
  - `message` - (String)  An explanation of the status reason
  - `more_info` - (String) Link to documentation about this status reason
- `tags` - (String) Tags associated with the volume.
