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

```

## Argument Reference

You can specify the following arguments for this data source.

- `profile` - (Optional, String) The name of a bare metal server profile. Filters the collection to resources with a profile.name property matching the specified name.
- `zone` - (Optional, String) The name of a zone. Filters the collection to resources with a zone.name property matching the specified name.

## Attribute Reference

Review the attribute references that you can access after you retrieve your data source.

- `id` - (String) The unique identifier of the bare metal server capacities data source.
- `capacities` - (List) A page of available bare metal server capacities. Each element represents a zone that has available bare metal servers with a profile.

  Nested schema for **capacities**:
  - `profile` - (List) The profile available in the zone.
    
    Nested schema for **profile**:
    - `href` - (String) The URL for this bare metal server profile.
    - `name` - (String) The name for this bare metal server profile.
    - `resource_type` - (String) The resource type. Value: `bare_metal_server_profile`.
  
  - `zone` - (List) The zone where one or more bare metal servers of the profile are available.
    
    Nested schema for **zone**:
    - `href` - (String) The URL for this zone.
    - `name` - (String) The globally unique name for this zone.