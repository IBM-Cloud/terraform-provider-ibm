---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : Subnets"
description: |-
  Manages IBM Cloud infrastructure subnets.
---

# ibm_is_subnets
Retrieve information about of an existing VPC subnets in an IBM Cloud account. For more information, about infrastructure subnets, see [attaching subnets to a routing table](https://cloud.ibm.com/docs/vpc?topic=vpc-attach-subnets-routing-table).

## Example usage

```hcl
data "ibm_resource_group" "resourceGroup" {
  name = "Default"
}

resource "ibm_is_vpc" "testacc_vpc" {
  name = "test"
}

resource "ibm_is_vpc_routing_table" "test_cr_route_table1" {
  name = "test-cr-route-table1"
  vpc  = ibm_is_vpc.testacc_vpc.id
}

resource "ibm_is_subnet" "testacc_subnet" {
  name            = "test_subnet"
  vpc             = ibm_is_vpc.testacc_vpc.id
  zone            = "us-south-1"
  ipv4_cidr_block = "192.168.0.0/1"
  routing_table   = ibm_is_vpc_routing_table.test_cr_route_table1.routing_table
  resource_group  = data.ibm_resource_group.resourceGroup.id
}

data "ibm_is_subnets" "ds_subnets_resource_group" {
  resource_group = data.ibm_resource_group.resourceGroup.id
}

data "ibm_is_subnets" "ds_subnets_routing_table_name" {
  routing_table_name = ibm_is_vpc_routing_table.test_cr_route_table1.name
}

data "ibm_is_subnets" "ds_subnets_routing_table" {
  routing_table = ibm_is_vpc_routing_table.test_cr_route_table1.id
}

data "ibm_is_subnets" "ds_subnets" {
}
```

## Argument Reference

Review the argument references that you can specify for your data source. 

* `resource_group` - (Optional, string) The id of the resource group.
* `routing_table` - (Optional, string) The id of the routing table.
* `routing_table_name` - (Optional, string) The name of the routing table.

## Attribute reference
You can access the following attribute references after your data source is created. 

- `subnets` - (List) A list of subnets in the IBM Cloud infrastructure.

  Nested scheme for `subnets`:
    - `available_ipv4_address_count`- (Integer) The number of IPv4 addresses that are available in the subnet.
	- `crn` - (String) The CRN of the subnet.
	- `id` - (String) The ID of the subnet.
	- `ipv4_cidr_block` - (String) The IPv4 CIDR block of this subnet.
	- `ipv6_cidr_block` - (String) The IPv6 CIDR block of this subnet.
	- `name` - (String) The name of the subnet.
	- `network_acl` - (String) The access control list (ACL) that is attached to the subnet.
    - `public_gateway`- (Bool) If set to **true**, a public gateway is attached to the subnet. If set to **false**, no public gateway for this subnet exists.
	- `resource_group` - (String) The resource group that the subnet belongs to.
    - `total_ipv4_address_count`- (Integer) The total number of IPv4 addresses in the subnet.
    - `status` - (String) The status of the subnet.
	- `vpc` - (String) The ID of the VPC that this subnet belongs to.
	- `zone` - (String) The zone where the subnet was created.
