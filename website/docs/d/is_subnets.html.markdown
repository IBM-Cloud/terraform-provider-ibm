---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : Subnets"
description: |-
  Manages IBM Cloud infrastructure subnets.
---

# ibm_is_subnets
Retrieve information about of an existing VPC subnets in an IBM Cloud account as a read only data source. For more information, about infrastructure subnets, see [attaching subnets to a routing table](https://cloud.ibm.com/docs/vpc?topic=vpc-attach-subnets-routing-table).

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
data "ibm_resource_group" "example" {
  name = "Default"
}

resource "ibm_is_vpc" "example" {
  name = "example-vpc"
}

resource "ibm_is_vpc_routing_table" "example" {
  name = "example-vpc-routing-table"
  vpc  = ibm_is_vpc.example.id
}

resource "ibm_is_subnet" "example" {
  name            = "example-subnet"
  vpc             = ibm_is_vpc.example.id
  zone            = "us-south-1"
  ipv4_cidr_block = "10.240.0.0/24"
  routing_table   = ibm_is_vpc_routing_table.example.routing_table
  resource_group  = data.ibm_resource_group.example.id
}

data "ibm_is_subnets" "example1" {
  resource_group = data.ibm_resource_group.example.id
}

data "ibm_is_subnets" "example2" {
  routing_table_name = ibm_is_vpc_routing_table.example.name
}

data "ibm_is_subnets" "example3" {
  routing_table = ibm_is_vpc_routing_table.example.id
}

data "ibm_is_subnets" "example4" {
}
```

## Argument reference

Review the argument references that you can specify for your data source. 

- `resource_group` - (Optional, string) The id of the resource group.
- `routing_table` - (Optional, string) The id of the routing table.
- `routing_table_name` - (Optional, string) The name of the routing table.
- `vpc` - (Optional, string) The id of the vpc.
- `vpc_crn` - (Optional, string) The crn of the vpc.
- `vpc_name` - (Optional, string) The name of vpc.
- `zone` - (Optional, string) The name of the zone.

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
	- `resource_group` - (String) The resource group id, that the subnet belongs to.
    - `total_ipv4_address_count`- (Integer) The total number of IPv4 addresses in the subnet.
    - `status` - (String) The status of the subnet.
  - `routing_table` -  (List) The routing table for this subnet. 
    Nested scheme for `routing_table`:
      - `crn` -  (String) The crn for this routing table.
      - `deleted` -  (List) If present, this property indicates the referenced resource has been deleted and provides some supplementary information.
      Nested scheme for `deleted`:
        - `more_info` -  (String) Link to documentation about deleted resources.
      - `href` -  (String) The URL for this routing table.
      - `id` -  (String) The unique identifier for this routing table.
      - `name` -  (String) The user-defined name for this routing table.
      - `resource_type` -  (String) The type of resource referenced.
	- `vpc` - (String) The ID of the VPC that this subnet belongs to.
	- `zone` - (String) The zone where the subnet was created.
