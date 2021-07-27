---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : Routing Table Routes"
description: |-
  Get information about IBM VPC routing table routes.
---

# ibm_is_vpc_routing_table_routes
Retrieve information of an existing IBM Cloud Infrastructure Virtual Private Cloud routing table routes as a read-only data source. For more information, about VPC default routing table, see [about routing tables and routes](https://cloud.ibm.com/docs/vpc?topic=vpc-about-custom-routes).


## Example usage

```terraform

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
## Argument reference
Review the argument references that you can specify for your data source. 

- `vpc` - (Required, String) The ID of the VPC.
- `routing_table` - (Required, String) The ID of the routing table.

## Attribute reference
In addition to all argument reference list, you can access the following attribute references after your data source is created. 

- `routing_table_routes` (List) List of all the routing table in a VPC.

  Nested scheme for `routing_table_routes`:
	- `name` - (String) The name for the default routing table.
	- `route_id` - (String) The unique ID for the route.
	- `lifecycle_state` - (String) The lifecycle state of the route.
	- `href` - (String) The routing table URL.
- `created_at` - (Timestamp)  The date and time that the route was created.
	- `action` - (String) The action to perform with a packet matching the route.
	- `destination` - (String) The destination of the route.
	- `next_hop` - (String) The next hop address of the route.
	- `zone` - (String) The zone name of the route.
