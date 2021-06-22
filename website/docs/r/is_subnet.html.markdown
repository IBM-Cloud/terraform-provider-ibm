---

subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : subnet"
description: |-
  Manages IBM subnet.
---

# ibm_is_subnet
Create, update, or delete a subnet. For more information, about subnet, see [configuring ACLs and security groups for use with VPN](https://cloud.ibm.com/docs/vpc?topic=vpc-acls-security-groups-vpn).


## Example usage

```terraform
resource "ibm_is_vpc" "testacc_vpc" {
  name = "test"
}

resource "ibm_is_vpc_routing_table" "test_cr_route_table1" {
  name   = "test-cr-route-table1"
  vpc    = data.ibm_is_vpc.testacc_vpc.id
}


resource "ibm_is_subnet" "testacc_subnet" {
  name            = "test_subnet"
  vpc             = ibm_is_vpc.testacc_vpc.id
  zone            = "us-south-1"
  ipv4_cidr_block = "192.168.0.0/1"
  routing_table   = ibm_is_vpc_routing_table.test_cr_route_table1.routing_table  

  //User can configure timeouts
  timeouts {
    create = "90m"
    delete = "30m"
  }
}
```

## Timeouts
The `ibm_is_subnet` resource provides the following [Timeouts](https://www.terraform.io/docs/language/resources/syntax.html) configuration options:

- **create** - (Default 10 minutes) Used for creating Instance.
- **update** - (Default 10 minutes) Used for creating Instance.
- **delete** - (Default 10 minutes) Used for deleting Instance.


## Argument reference
Review the argument references that you can specify for your resource. 

- `ipv4_cidr_block` - (Optional, Forces new resource, String) The IPv4 range of the subnet.
- `ip_version` - (Optional, Forces new resource, String) The IP Version. The default is `ipv4`.
- `name` - (Required, String) The name of the subnet.
- `network_acl` - (Optional, String) The ID of the network ACL for the subnet.
- `public_gateway` - (Optional, String) The ID of the public gateway for the subnet that you want to attach to the subnet. You create the public gateway with the [`ibm_is_public_gateway` resource](#provider-public-gateway).
- `resource_group` - (Optional, Forces new resource, String) The ID of the resource group where you want to create the subnet.
- `routing_table` - (Optional, String) The routing table ID associated with the subnet.
- `tags`  - (Optional, List of Strings) A list of tags with the service policy instance. **Note** Tags are managed locally and not stored in the IBM Cloud service endpoint at this moment.
- `total_ipv4_address_count` - (Optional, Forces new resource, String) The total number of IPv4 addresses. Either `ipv4_cidr_block` or `total_pv4_address_count` input parameters must be provided in the resource.
- `vpc` - (Required, Forces new resource, String) The VPC ID.
- `zone` - (Required, Forces new resource, String) The subnet zone name.


## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `available_ipv4_address_count` - (String) The total number of available IPv4 addresses.
- `crn` - (String) The CRN of subnet.
- `id` - (String) The ID of the subnet.
- `ipv6_cidr_block` - (String) The IPv6 range of the subnet.
- `status` - (String) The status of the subnet.

## Import
The `ibm_is_subnet` resource can be imported by using the ID. 

**Syntax**

```
$ terraform import ibm_is_subnet.example <subnet_ID>
```

**Example**

```
$ terraform import ibm_is_subnet.example d7bec597-4726-451f-8a63-e62e6f12122c
```
