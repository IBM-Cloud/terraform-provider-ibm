---
layout: "ibm"
page_title: "IBM : subnet"
sidebar_current: "docs-ibm-datasources-is-subnet"
description: |-
  Manages IBM subnet.
---

# ibm\_is_subnet

Import the details of an existing IBM cloud subnet as a read-only data source. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.


## Example Usage

```hcl
resource "ibm_is_vpc" "testacc_vpc" {
	name = "test"
}

resource "ibm_is_subnet" "testacc_subnet" {
	name = "test_subnet"
	vpc = "${ibm_is_vpc.testacc_vpc.id}"
	zone = "us-south-1"
	ipv4_cidr_block = "192.168.0.0/1"
}
data "ibm_is_subnet" "ds_subnet" {
	identifier = "${ibm_is_subnet.testacc_subnet.id}"
}

```

## Argument Reference

The following arguments are supported:

* `identifier` - (Required, string) The id of the subnet.

## Attribute Reference

The following attributes are exported:

* `ipv4_cidr_block` -  The IPv4 range of the subnet.
* `ipv6_cidr_block` - The IPv6 range of the subnet.
* `total_ipv4_address_count` - The total number of IPv4 addresses.
* `ip_version` - The Ip Version.
* `name` - The name of the subnet.
* `network_acl` - The ID of the network ACL for the subnet.
* `public_gateway` - The ID of the public-gateway for the subnet.
* `status` - The status of the subnet.
* `vpc` - The vpc id.
* `zone` - The subnet zone name.
* `available_ipv4_address_count` - The total number of available IPv4 addresses.