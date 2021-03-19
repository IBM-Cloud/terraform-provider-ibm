---

subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : vpc-routing-tables-route"
description: |-
  Manages IBM IS VPC Routing tables.
---

# ibm\_is_vpc_routing_table_route

Provides a vpc routing tables resource. This allows vpc routing tables to be created, updated, and cancelled.


## Example Usage

```hcl
resource "ibm_is_vpc_routing_table_route" "test_ibm_is_vpc_routing_table_route" {
  vpc = ""
  routing_table = ""
  zone = "us-south-1"
  name = "custom-route-2"
  destination = "192.168.4.0/24"
  action = "deliver"
  next_hop_address    = "10.0.0.4"
}

```

## Argument Reference

The following arguments are supported:

* `name` - (Optional, string) The user-defined name for this route. If unspecified, the name will be a hyphenated list of randomly-selected words. Names must be unique within the VPC routing table the route resides in.
* `vpc` - (Required, Forces new resource, string) The vpc id.
* `routing_table` - (Required, Forces new resource, string) The routing table identifier
* `action` - (Optional,string) The action to perform with a packet matching the route.
* `zone` - (Required, Forces new resource, string) Name of the zone.
* `destination` - (Required, Forces new resource, string) The destination of the route.
* `next_hop_address` - (Optional, Forces new resource, string) The next hop of the route.
 **NOTE**: Conflicts with `next_hop_vpn_connection`
* `next_hop_vpn_connection` - (Optional, Forces new resource, string) The next hop to vpn connection gateway.

## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier for this routing table. The id is composed of \<vpc_route_table_id\>/\<vpc_route_table_route_id\>
* `is_default` - Indicates whether this is the default routing table for this VPC
* `lifecycle_state` - The lifecycle state of the route
* `resource_type` - The resource type
* `href` - The URL for this route


## Import

ibm_is_vpc_routing_table_route can be imported using VPC ID, VPC Route table ID and VPC Route table Route ID , eg

```
$ terraform import ibm_is_vpc_routing_table_route.example 56738c92-4631-4eb5-8938-8af9211a6ea4/4993-a0fd-cabab477c4d1-8af9211a6ea4/fc2667e0-9e6f-4993-a0fd-cabab477c4d1
```
