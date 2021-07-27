---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : subnet"
description: |-
  Manages IBM Cloud subnet.
---

# ibm_is_subnet
Retrieve information of an existing VPC Generation 2 compute subnet as a read only data source. For more information, about the IBM Cloud subnet, see [attaching subnets to a routing table](https://cloud.ibm.com/docs/vpc?topic=vpc-attach-subnets-routing-table).


## Example usage
Example to retrieve the subnet information by using subnet name.

```terraform
resource "ibm_is_vpc" "testacc_vpc" {
  name = "test"
}

resource "ibm_is_subnet" "testacc_subnet" {
  name            = "test_subnet"
  vpc             = ibm_is_vpc.testacc_vpc.id
  zone            = "us-south-1"
  ipv4_cidr_block = "192.168.0.0/1"
}

data "ibm_is_subnet" "ds_subnet" {
  name = ibm_is_subnet.testacc_subnet.name
}

```
// Example to retrieve the subnet information by using subnet ID.

```terraform
resource "ibm_is_vpc" "testacc_vpc" {
  name = "test"
}

resource "ibm_is_subnet" "testacc_subnet" {
  name            = "test-subnet"
  vpc             = ibm_is_vpc.testacc_vpc.id
  zone            = "us-south-1"
  ipv4_cidr_block = "192.168.0.0/1"
}

data "ibm_is_subnet" "ds_subnet" {
  identifier = ibm_is_subnet.testacc_subnet.id
}

```

## Argument reference
Review the argument references that you can specify for your data source. 

- `identifier` - (Optional, String) The ID of the subnet.
- `name` - (Optional, String) The name of the subnet.

## Attribute reference
In addition to all argument reference list, you can access the following attribute references after your data source is created. 

- `available_ipv4_address_count` - (Integer) The total number of available IPv4 addresses.
- `crn` - (String) The CRN of subnet.
- `ipv4_cidr_block` -  (String) The IPv4 range of the subnet.
- `ip_version` - (String) The IP version.
- `name` - (String) The name of the subnet.
- `network_acl` - (String) The ID of the network ACL for the subnet.
- `public_gateway` - (String) The ID of the public gateway for the subnet.
- `resource_group` - (String) The subnet resource group.
- `status` - (String) The status of the subnet.
- `tags`  - (String) Tags associated for the instance.
- `total_ipv4_address_count` - (Integer) The total number of IPv4 addresses.
- `vpc` - (String) The ID of the VPC that the subnet belongs to.
- `vpc_name` - (String) The name of the VPC that the subnet belongs to.
- `zone` - (String) The subnet zone name.
