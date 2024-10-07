---

subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : subnet"
description: |-
  Manages IBM subnet.
---

# ibm_is_subnet
Create, update, or delete a subnet. For more information, about subnet, see [configuring ACLs and security groups for use with VPN](https://cloud.ibm.com/docs/vpc?topic=vpc-acls-security-groups-vpn).

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
  name = "example-routing-table"
  vpc  =  ibm_is_vpc.example.id
}


resource "ibm_is_subnet" "example" {
  name            = "example-subnet"
  vpc             = ibm_is_vpc.example.id
  zone            = "us-south-1"
  ipv4_cidr_block = "10.240.0.0/24"
  routing_table   = ibm_is_vpc_routing_table.example.routing_table

  //User can configure timeouts
  timeouts {
    create = "90m"
    delete = "30m"
  }
}
```

## Example usage with address prefix
```terraform
resource "ibm_is_vpc" "example" {
  name = "example-vpc"
}

resource "ibm_is_vpc_address_prefix" "example" {
  cidr = "10.0.1.0/24"
  name = "example-add-prefix"
  vpc  = ibm_is_vpc.example.id
  zone = "us-south-1"
}

resource "ibm_is_subnet" "example" {
  depends_on = [
    ibm_is_vpc_address_prefix.example
  ]
  ipv4_cidr_block = "10.0.1.0/24"
  name            = "example-subnet"
  vpc             = ibm_is_vpc.example.id
  zone            = "us-south-1"
}
```


## Timeouts
The `ibm_is_subnet` resource provides the following [Timeouts](https://www.terraform.io/docs/language/resources/syntax.html) configuration options:

- **create** - (Default 10 minutes) Used for creating Instance.
- **update** - (Default 10 minutes) Used for creating Instance.
- **delete** - (Default 10 minutes) Used for deleting Instance.


## Argument reference
Review the argument references that you can specify for your resource. 

- `access_tags`  - (Optional, List of Strings) A list of access management tags to attach to the bare metal server.

  ~> **Note:** 
  **&#x2022;** You can attach only those access tags that already exists.</br>
  **&#x2022;** For more information, about creating access tags, see [working with tags](https://cloud.ibm.com/docs/account?topic=account-tag&interface=ui#create-access-console).</br>
  **&#x2022;** You must have the access listed in the [Granting users access to tag resources](https://cloud.ibm.com/docs/account?topic=account-access) for `access_tags`</br>
  **&#x2022;** `access_tags` must be in the format `key:value`.
- `ipv4_cidr_block` - (Optional, Forces new resource, String) The IPv4 range of the subnet.

  ~> **NOTE:**
    If using a IPv4 range from a `ibm_is_vpc_address_prefix` resource, add a `depends_on` to handle hidden `ibm_is_vpc_address_prefix` dependency if not using interpolation.

- `ip_version` - (Optional, Forces new resource, String) The IP Version. The default is `ipv4`.
- `name` - (Required, String) The name of the subnet.
- `network_acl` - (Optional, String) The ID of the network ACL for the subnet.
- `public_gateway` - (Optional, String) The ID of the public gateway for the subnet that you want to attach to the subnet. You create the public gateway with the [`ibm_is_public_gateway` resource](#provider-public-gateway).
- `resource_group` - (Optional, Forces new resource, String) The ID of the resource group where you want to create the subnet.
- `routing_table` - (Optional, String) The routing table ID associated with the subnet.
- `routing_table_crn` - (Optional, String) The routing table crn associated with the subnet.
  ~> **Note** 
  `routing_table` and `routing_table_crn` are mutually exclusive.
- `tags`  - (Optional, List of Strings) The tags associated with the subnet.
- `total_ipv4_address_count` - (Optional, Forces new resource, String) The total number of IPv4 addresses. Either `ipv4_cidr_block` or `total_pv4_address_count` input parameters must be provided in the resource.
  
  ~> **Note** 
  The VPC must have a default address prefix in the specified zone, and that prefix must have a free CIDR range with at least this number of addresses.

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
