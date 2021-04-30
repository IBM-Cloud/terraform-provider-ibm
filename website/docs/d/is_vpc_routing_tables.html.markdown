---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : Routing Tables"
description: |-
  Get information about IBM VPC Routing Tables.
---

# ibm\_is_vpc_routing_tables

Import the details of an existing IBM Cloud Infrastructure Virtual Private Cloud routing tables as a read-only data source. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.


## Example Usage

```hcl

resource "ibm_is_vpc" "test_vpc" {
  name = "test-vpc"
}

data "ibm_is_vpc_routing_tables" "ds_routing_tables" {
	vpc = ibm_is_vpc.test_vpc.id
}

```

## Argument Reference

The following arguments are supported:

* `vpc` - (Required, string) The id of the VPC.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `routing tables` - List of all routing tables in a VPC in the IBM Cloud Infrastructure.
  * `name` - The name for the routing table.
  * `routing_table` - The id of the routing table.
  * `lifecycle_state` - The lifecycle state of the routing table.
  * `href` - The URL for the routing table.
  * `resource_type` - The type of resource referenced.
  * `created_at` - The date and time that the routing table was created.
  * `is_default` - Indicates whether this is the default routing table for this VPC.
  * `route_direct_link_ingress` - Indicates if this routing table will be used to route traffic that originates from Direct Link to this VPC.
  * `route_transit_gateway_ingress` - Indicates if this routing table will be used to route traffic that originates from Transit Gateway to this VPC.
  * `route_vpc_zone_ingress` - Indicates if this routing table will be used to route traffic that originates from subnets in other zones in this VPC.
  * `routes` - The routes for the routing table.
    * `name` - The name for the route.
    * `id` - The unique identifier for the route.
  * `subnets` - The subnets to which the routing table is attached.
    * `name` - The name for the subnet.
    * `id` - The unique identifier for the subnet.
