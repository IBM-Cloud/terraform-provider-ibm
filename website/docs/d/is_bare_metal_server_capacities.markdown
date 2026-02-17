---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : Bare Metal Server Capacities"
description: |-
  Manages IBM Cloud Bare Metal Server Capacities.
---

# ibm\_is_bare_metal_server_capacities

Import the details of existing IBM Cloud Bare Metal Server capacity information as a read-only data source. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax. For more information, about bare metal server capacities, see [Bare Metal Servers for VPC](https://cloud.ibm.com/docs/vpc?topic=vpc-bare-metal-servers-profile).

This data source provides information about which bare metal server profiles have available capacity in which zones within the region.

**Note:** 
VPC infrastructure services are a regional specific based endpoint, by default targets to `us-south`. Please make sure to target right region in the provider block as shown in the `provider.tf` file, if VPC service is created in region other than `us-south`.

**provider.tf**

```terraform
provider "ibm" {
  region = "eu-gb"
}
```

## Example Usage

```terraform
# Get all bare metal server capacities in the region
data "ibm_is_bare_metal_server_capacities" "ds_bmscapacities" {
}

# Filter by profile name
data "ibm_is_bare_metal_server_capacities" "ds_bmscapacities_by_profile" {
  profile = "bx2-metal-192x768"
}

# Filter by zone name
data "ibm_is_bare_metal_server_capacities" "ds_bmscapacities_by_zone" {
  zone = "us-south-1"
}

# Filter by both profile and zone
data "ibm_is_bare_metal_server_capacities" "ds_bmscapacities_specific" {
  profile = "bx2-metal-192x768"
  zone    = "us-south-1"
}

# Use the output to check if a profile is available in a specific zone
output "profile_available_zones" {
  value = data.ibm_is_bare_metal_server_capacities.ds_bmscapacities_by_profile.capacities[0].zones
}

# Check if capacity exists for a specific profile
locals {
  has_capacity = length(data.ibm_is_bare_metal_server_capacities.ds_bmscapacities_specific.capacities) > 0
}
```

## Argument Reference

Review the argument references that you can specify for your data source.

- `profile` - (Optional, String) The name of a bare metal server profile. Filters the collection to resources with a profile.name property matching the specified name.
- `zone` - (Optional, String) The name of a zone. Filters the collection to resources with a zone.name property matching the specified name.

## Attribute Reference

Review the attribute references that you can access after you retrieve your data source.

- `id` - (String) The unique identifier of the bare metal server capacities data source.
- `capacities` - (List) List of capacities for each profile. The results will include all profile capacities unless a zone or profile filter are specified. The API returns individual profile+zone pairs which are aggregated client-side by profile name.

  Nested scheme for an element of `capacities`:
  - `name` - (String) The name of the bare metal server profile.
  - `zones` - (List) List of zones in the region that have capacity for the profile. This list represents availability zones where the profile can be provisioned.