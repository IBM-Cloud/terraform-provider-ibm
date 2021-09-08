---

subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : vpc-routing-tables"
description: |-
  Manages IBM IS VPC routing tables.
---

# ibm_is_vpc_routing_table
Create, update, or delete an VPC routing tables. For more information, about VPC routes, see [routing tables for VPC](https://cloud.ibm.com/docs/vpc?topic=vpc (List)routing-tables-for-vpc).


## Example usage

```terraform
resource "ibm_is_vpc" "testacc_vpc" {
  name = "testvpc"
}
resource "ibm_is_vpc_routing_table" "test_ibm_is_vpc_routing_table" {
  vpc = ibm_is_vpc.testacc_vpc.id
  name = "routTabletest"
  route_direct_link_ingress = true
  route_transit_gateway_ingress = false
  route_vpc_zone_ingress = false
}

```

## Argument reference
Review the argument references that you can specify for your resource. 

- `name` - (Optional, String) The routing table name.
- `route_direct_link_ingress` - (Optional, Bool)  If set to **true**, the routing table is used to route traffic that originates from Direct Link to the VPC. To succeed, the VPC must not already have a routing table with the property set to **true**.
- `route_transit_gateway_ingress` - (Optional, Bool) If set to **true**, the routing table is used to route traffic that originates from Transit Gateway to the VPC. To succeed, the VPC must not already have a routing table with the property set to **true**.
- `route_vpc_zone_ingress` - (Optional, Bool) If set to true, the routing table is used to route traffic that originates from subnets in other zones in the VPC. To succeed, the VPC must not already have a routing table with the property set to **true**.
- `vpc` - (Required, Forces new resource, String) The VPC ID. 

## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `href` - (String) The routing table URL.
- `id` - (String) The routing table ID. The ID is composed of `<vpc_id>/<vpc_route_table_id>` of the VPC route.
- `is_default` - (String)  Indicates the default routing table for this VPC.
- `lifecycle_state` - (String) The lifecycle state of the routing table.
- `resource_type` - (String) The resource type.
- `routing_table` - (String) The generated routing table ID.
- `routes` - (List) The routes for the routing table.

  Nested scheme for `routes`:
  - `id` - (String) The unique ID of the route.
  - `name`-  (String) The user-defined name of the route.
- `subnets` - (List) The subnets to which routing table is attached.

  Nested scheme for `subnets`:
  - `id` - (String) The unique ID of the subnet.
  - `name` - (String) The user defined name of the subnet.

## Import
The `ibm_is_vpc_routing_table` resource can be imported by using VPC ID and VPC Route table ID.

**Example**

```
$ terraform import ibm_is_vpc_routing_table.example 56738c92-4631-4eb5-8938-8af9211a6ea4/fc2667e0-9e6f-4993-a0fd-cabab477c4d1
```
