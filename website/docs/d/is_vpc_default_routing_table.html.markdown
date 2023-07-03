---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : Default Routing Table"
description: |-
  Get Information about IBM VPC default routing table.
---

# ibm_is_vpc_default_routing_table
Retrieve information of an existing IBM Cloud Infrastructure Virtual Pricate Cloud default routing table as a read-only data source. For more information, about VPC default routing table, see [about routing tables and routes](https://cloud.ibm.com/docs/vpc?topic=vpc-about-custom-routes).

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
resource "ibm_is_vpc" "example" {
  name = "example-vpc"
}

data "ibm_is_vpc_default_routing_table" "example" {
  vpc = ibm_is_vpc.example.id
}
```

## Argument reference
Review the argument references that you can specify for your data source. 

- `vpc` - (Required, String) The ID of the VPC.

## Attribute reference
In addition to the argument reference list, you can access the following attribute references after your data source is created. 

- `created_at` - (Timestamp)  The date and time that the default routing table was created.
- `default_routing_table` - (String) The unique identifier for this routing table.
- `href` - (String) The routing table URL.
- `id` - (String) The unique identifier for this routing table. Same as `default_routing_table`.
- `is_default` - (String)  Indicates the default routing table for this VPC.
- `lifecycle_state` - (String) The lifecycle state of the routing table.
- `name` - (String) The name for the default routing table.
- `resource_type` - (String) The resource type.
- `route_direct_link_ingress`- (Bool)  Indicates the routing table is used to route traffic that originates from Direct Link to the VPC.
- `route_internet_ingress` - (Bool) Indicates whether this routing table is used to route traffic that originates from the internet.
- `route_transit_gateway_ingress`- (Bool) Indicates the routing table is used to route traffic that originates from Transit Gateway to the VPC.
- `route_vpc_zone_ingress`- (Bool) Indicates the routing table is used to route traffic that originates from subnets in other zones in the VPC.
- `routes` - (List) The routes for the default routing table.

  Nested scheme for `routes`:
	- `id` - (String) The unique ID of the route.
	- `name` -  (String) The name of the route.
- `subnets` - (List) The subnets to which routing table is attached.

  Nested scheme for `subnets`:
	- `id` - (String) The unique ID of the subnet.
	- `name` - (String) The name of the subnet.
