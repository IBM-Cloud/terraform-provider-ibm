---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : Routing Table Routes"
description: |-
  Get information about IBM VPC Routing Table Routes.
---

# ibm\_is_vpc_routing_table_routes

Import the details of an existing IBM Cloud Infrastructure Virtual Private Cloud routing table routes as a read-only data source. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.


## Example Usage

```hcl

resource "ibm_is_vpc" "test_vpc" {
  name = "test-vpc"
}

resource "ibm_is_vpc_routing_table" "test_routing_table" {
  name   = "test-routing-table"
  vpc    = ibm_is_vpc.test_vpc.id
}


data "ibm_is_vpc_routing_table_routes" "ds_routing_table_routes" {
	vpc = ibm_is_vpc.test_vpc.id
	routing_table = ibm_is_vpc_routing_tables.test_routing_table.routing_table
}

```

## Argument Reference

The following arguments are supported:

* `vpc` - (Required, string) The id of the VPC.
* `routing_table` - (Required, string) The id of the Routing Table.

## Attribute Reference

The following attributes are exported:

* `routing table routes` - List of all routes in a Routing Table in a VPC.
  * `name` - The name for the route.
  * `route_id` - The unique identifier for the route.
  * `lifecycle_state` - The lifecycle state of the route.
  * `href` - The URL for the route.
  * `created_at` - The date and time that the route was created.
  * `action` - The action to perform with a packet matching the route.
  * `destination` - The destination of the route.
  * `nexthop` - The next_hop address of the route.
  * `zone` - The zone name of the route.
