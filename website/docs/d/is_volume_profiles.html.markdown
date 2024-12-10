---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : Volume profiles"
description: |-
  Manages IBM Cloud virtual server volume profiles.
---

# ibm_is_volume_profiles
Retrieve information of an existing IBM Cloud VSI. For more information, about the volumes and profiles, see [block storage profiles](https://cloud.ibm.com/docs/vpc?topic=vpc-block-storage-profiles).

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

data "ibm_is_volume_profiles" "example" {
}

```

## Attribute reference
You can access the following attribute references after your data source is created. 

- `profiles` - (List)  Lists all server volume profiles in the region.

  Nested scheme for `profiles`:
  - `adjustable_capacity_states` - (List) 
    Nested schema for **adjustable_capacity_states**:
    - `type` - (String) The type for this profile field.
    - `values` - (List) The attachment states that support adjustable capacity for a volume with this profile. Allowable list items are: `attached`, `unattached`, `unusable`. 
  - `adjustable_iops_states` - (List) 
    Nested schema for **adjustable_iops_states**:
    - `type` - (String) The type for this profile field.
    - `values` - (List) The attachment states that support adjustable IOPS for a volume with this profile. Allowable list items are: `attached`, `unattached`, `unusable`.
	- `name` - (String) The name of the virtual server volume profile.
	- `family` - (String) The family of the virtual server volume profile.

