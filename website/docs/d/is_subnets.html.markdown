---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : Subnets"
description: |-
  Manages IBM Cloud Infrastructure Subnets.
---

# ibm\_is_subnets

Import the details of an existing IBM Cloud Infrastructure subnets as a read-only data source. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.


## Example Usage

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

The following arguments are supported:

* `resource_group` - (Optional, string) The id of the resource group.
* `routing_table` - (Optional, string) The id of the routing table.
* `routing_table_name` - (Optional, string) The name of the routing table.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `subnets` - List of all subnets in the IBM Cloud Infrastructure.
  * `name` - The name for this subnet.
  * `id` - The unique identifier for this subnet.
  * `ipv4_cidr_block` - The IPv4 CIDR block for this subnet.
  * `ipv6_cidr_block` - The IPv6 CIDR block for this subnet when used.
  * `status` - The status of this subnet.
  * `crn` - The CRN for this image.
  * `available_ipv4_address_count` - Amount of addresses available within this subnet.
  * `total_ipv4_address_count` - Amount of addresses used within this subnet.
  * `network_acl` - Security group attached to this subnet.
  * `public_gateway` - Public gateway attached to this subnet.
  * `resource_group` - Resource group where this subnet is created.
  * `vpc` - VPC where this subnet is created.
  * `zone` - Zone where this subnet is created.