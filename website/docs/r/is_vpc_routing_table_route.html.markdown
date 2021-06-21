---

subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : vpc-routing-tables-route"
description: |-
  Manages IBM IS VPC routing tables.
---

# ibm_is_vpc_routing_table_route
Create, update, or delete of an VPC routing tables. For more information, about VPC routes, see [about routing tables and routes](https://cloud.ibm.com/docs/vpc?topic=vpc-about-custom-routes).


## Example usage

```terraform
resource "ibm_is_vpc_routing_table_route" "test_ibm_is_vpc_routing_table_route" {
  vpc = ""
  routing_table = ""
  zone = "us-south-1"
  name = "custom-route-2"
  destination = "192.168.4.0/24"
  action = "deliver"
  next_hop = "10.0.0.4"
}
```

```terraform
resource "ibm_is_vpc_routing_table_route" "test_ibm_is_vpc_routing_table_route" {
  vpc = ""
  routing_table = ""
  zone = "us-south-1"
  name = "custom-route-2"
  destination = "192.168.4.0/24"
  action = "deliver"
  next_hop = ibm_is_vpn_gateway_connection.VPNGatewayConnection.gateway_connection
}

```

## Argument reference
Review the argument references that you can specify for your resource. 

- `action` - (Optional, String) The action to perform with a packet matching the route `delegate`, `delegate_vpc`, `deliver`, `drop`.
- `destination` - (Required, Forces new resource, String) The destination of the route. 
- `name` - (Optional, String) The user-defined name of the route. If unspecified, the name will be a hyphenated list of randomly selected words. You need to provide unique name within the VPC routing table the route resides in.
- `next_hop` - (Required, Forces new resource, String) The next hop of the route. It accepts IP address or a VPN connection ID. For `action` other than `deliver`, you must specify `0.0.0.0`. 
- `routing_table` - (Required, String) The routing table ID.
- `vpc` - (Required, Forces new resource, String) The VPC ID.
- `zone` - (Required, Forces new resource, String)  Name of the zone. 


## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `href` - (String) The routing table URL.
- `id` - (String) The routing table ID. The ID is composed of `<vpc_route_table_id>/<vpc_route_table_route_id>`.
- `is_default` - (String) Indicates the default routing table for this VPC.
- `lifecycle_state` - (String) The lifecycle state of the route.
- `resource_type` - (String) The resource type.

## Import
The `ibm_is_vpc_routing_table_route` resource  can be imported by using VPC ID, VPC Route table ID, and VPC Route table Route ID.

**Example**

```
$ terraform import ibm_is_vpc_routing_table_route.example 56738c92-4631-4eb5-8938-8af90000006ea4/4993-a0fd-cabab477c4d1-8af911111a4/fc2667e0-9e6f-4993-a0fd-cabab55557c4d1
```

