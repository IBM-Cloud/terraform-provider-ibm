---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : Volume profiles"
description: |-
  Manages IBM Cloud virtual server volume profiles.
---

# ibm_is_volume_profiles
Retrieve information of an existing IBM Cloud VSI. For more information, about the volumes and profiles, see [block storage profiles](https://cloud.ibm.com/docs/vpc?topic=vpc-block-storage-profiles).

## Example usage

```terraform

data "ibm_is_volume_profiles" "volprofiles"{
}

```

## Attribute reference
You can access the following attribute references after your data source is created. 

- `profiles` - (List)  Lists all server volume profiles in the region.

  Nested scheme for `profiles`:
	- `name` - (String) The name of the virtual server volume profile.
	- `family` - (String) The family of the virtual server volume profile.

