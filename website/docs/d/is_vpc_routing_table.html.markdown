---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : is_vpc_routing_table"
description: |-
  Get information about RoutingTable
---

# ibm_is_vpc_routing_table

Provides a read-only data source for RoutingTable. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax. For more information, about VPC routing tables, see [about routing tables and routes](https://cloud.ibm.com/docs/vpc?topic=vpc-about-custom-routes)

**Note:** 
VPC infrastructure services are a regional specific based endpoint, by default targets to `us-south`. Please make sure to target right region in the provider block as shown in the `provider.tf` file, if VPC service is created in region other than `us-south`.

**provider.tf**

```terraform
provider "ibm" {
  region = "eu-gb"
}
```

## Example Usage (using routing table id)
```terraform
data "ibm_is_vpc_routing_table" "example_routing_table" {
  vpc 				= ibm_is_vpc.example_vpc.id
  routing_table 	= ibm_is_vpc_routing_table.example_rt.routing_table
}
```

## Example Usage (using routing table name)
```terraform	
data "ibm_is_vpc_routing_table" "example_routing_table_name" {
  vpc 			= ibm_is_vpc.example_vpc.id
  name 			= ibm_is_vpc_routing_table.example_rt.name
}
```

## Argument Reference

Review the argument reference that you can specify for your data source.

- `name` - (Optional, String) The VPC routing table name. Mutually exclusive with `routing_table`, one of them is required
- `routing_table` - (Optional, String) The VPC routing table identifier. Mutually exclusive with `name`, one of them is required
- `vpc` - (Required, String) The VPC identifier.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

- `created_at` - (String) The date and time that this routing table was created.
- `href` - (String) The URL for this routing table.
- `id` - (String) The unique identifier of the RoutingTable.
- `is_default` - (Boolean) Indicates whether this is the default routing table for this VPC.
- `lifecycle_state` - (String) The lifecycle state of the routing table.
  - Constraints: Allowable values are: `deleting`, `failed`, `pending`, `stable`, `suspended`, `updating`, `waiting`.
- `name` - (String) The user-defined name for this routing table.
- `resource_type` - (String) The resource type.
- `route_direct_link_ingress` - (Boolean) Indicates whether this routing table is used to route traffic that originates from [Direct Link](https://cloud.ibm.com/docs/dl/) to this VPC.Incoming traffic will be routed according to the routing table with one exception: routes with an `action` of `deliver` are treated as `drop` unless the `next_hop` is an IP address within the VPC's address prefix ranges. Therefore, if an incoming packet matches a route with a `next_hop` of an internet-bound IP address or a VPN gateway connection, the packet will be dropped.
- `route_transit_gateway_ingress` - (Boolean) Indicates whether this routing table is used to route traffic that originates from from [Transit Gateway](https://cloud.ibm.com/cloud/transit-gateway/) to this VPC.Incoming traffic will be routed according to the routing table with one exception: routes with an `action` of `deliver` are treated as `drop` unless the `next_hop` is an IP address within the VPC's address prefix ranges. Therefore, if an incoming packet matches a route with a `next_hop` of an internet-bound IP address or a VPN gateway connection, the packet will be dropped.
- `route_vpc_zone_ingress` - (Boolean) Indicates whether this routing table is used to route traffic that originates from subnets in other zones in this VPC.Incoming traffic will be routed according to the routing table with one exception: routes with an `action` of `deliver` are treated as `drop` unless the `next_hop` is an IP address within the VPC's address prefix ranges. Therefore, if an incoming packet matches a route with a `next_hop` of an internet-bound IP address or a VPN gateway connection, the packet will be dropped.
- `routes` - (List) The routes for this routing table.
	Nested scheme for **routes**:
	- `deleted` - (List) If present, this property indicates the referenced resource has been deleted and providessome supplementary information.
		Nested scheme for **deleted**:
		- `more_info` - (String) Link to documentation about deleted resources.
	- `href` - (String) The URL for this route.
	- `id` - (String) The unique identifier for this route.
	- `name` - (String) The user-defined name for this route.
- `subnets` - (List) The subnets to which this routing table is attached.
	Nested scheme for **subnets**:
	- `crn` - (String) The CRN for this subnet.
	- `deleted` - (List) If present, this property indicates the referenced resource has been deleted and provides some supplementary information.
		Nested scheme for **deleted**:
		- `more_info` - (String) Link to documentation about deleted resources.
	- `href` - (String) The URL for this subnet.
	- `id` - (String) The unique identifier for this subnet.
	- `name` - (String) The user-defined name for this subnet.