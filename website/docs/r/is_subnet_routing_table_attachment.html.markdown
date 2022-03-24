---

subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : subnet_routing_table_attachment"
description: |-
  Manages IBM subnet routing table attachment.
---

# ibm_is_subnet_routing_table_attachment
Create, update, or delete a subnet routing table attachment resource. For more information, about subnet routing table attachment, see [setting up routing tables](https://cloud.ibm.com/docs/vpc?topic=vpc-attach-subnets-routing-table).

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
resource "ibm_is_vpc_routing_table" "example" {
  vpc   = ibm_is_vpc.example.id
  name  = "example-rt"
}

resource "ibm_is_subnet" "example" {
  name                        = "example-subnet"
  vpc                         = ibm_is_vpc.example.id
  zone                        = "eu-gb-1"
  total_ipv4_address_count    = 16
}

resource "ibm_is_subnet_routing_table_attachment" "example" {
  subnet        = ibm_is_subnet.example.id
  routing_table = ibm_is_vpc_routing_table.example.routing_table
}

```
## Argument reference
Review the argument references that you can specify for your resource. 

- `routing_table` - (Required, String) The routing table identity.
- `subnet` - (Required, Forces new resource, String) The subnet identifier.


## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `created_at` - (Timestamp) The creation date and time the routing table.
- `href` - (String) The URL of this routing table.
- `id` - (String) The unique identifier of the subnet.
- `is_default` - (Boolean) Indicates whether this is the default routing table for this VPC.
- `lifecycle_state` - (String) The lifecycle state of the routing table.
- `name` - (String) The user-defined name of this routing table.
- `resource_type` - (String) The resource type.
- `route_direct_link_ingress` - (Boolean) Indicates whether this routing table is used to route traffic that originates from [Direct Link](https://cloud.ibm.com/docs/dl/) to this VPC. Incoming traffic will be routed according to the routing table with one exception: routes with an action of deliver are treated as drop unless the next_hop is an IP address within the VPC's address prefix ranges. Therefore, if an incoming packet matches a route with a next_hop of an internet-bound IP address or a VPN gateway connection, the packet will be dropped..
- `route_transit_gateway_ingress` - (Boolean) Indicates whether this routing table is used to route traffic that originates from from [Transit Gateway](https://cloud.ibm.com/cloud/transit-gateway/) to this VPC.
Incoming traffic will be routed according to the routing table with one exception: routes with an action of deliver are treated as drop unless the next_hop is an IP address within the VPC's address prefix ranges. Therefore, if an incoming packet matches a route with a next_hop of an internet-bound IP address or a VPN gateway connection, the packet will be dropped.
- `route_vpc_zone_ingress` - (Boolean) Indicates whether this routing table is used to route traffic that originates from subnets in other zones in this VPC. Incoming traffic will be routed according to the routing table with one exception: routes with an action of deliver are treated as drop unless the next_hop is an IP address within the VPC's address prefix ranges. Therefore, if an incoming packet matches a route with a next_hop of an internet-bound IP address or a VPN gateway connection, the packet will be dropped..
- `routes` - (List) The routes for this routing table.
	
  Nested scheme for `routes`:
  - `href` - (List) The URL for this route.
  - `id` - (List) The unique identifier for this route.
  - `name` - (List) The user-defined name for this route.

- `subnets` - (List) The subnets to which this routing table is attached.

  Nested scheme for `subnets`:
	- `id` - (String) The unique identifier for this subnet.
	- `name` - (String) The user-defined name for this subnet.


## Import
The `ibm_is_subnet_routing_table_attachment` resource can be imported by using the subnet ID. 

**Syntax**

```
$ terraform import ibm_is_subnet_routing_table_attachment.example <subnet_ID>
```

**Example**

```
$ terraform import ibm_is_subnet_routing_table_attachment.example d7bec597-4726-451f-8a63-1111e6f19c32c
```
