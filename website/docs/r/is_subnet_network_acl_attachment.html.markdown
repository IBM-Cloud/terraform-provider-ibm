---

subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : subnet network acl attachment"
description: |-
  Manages IBM Subnet network ACL attachment.
---

# ibm_is_subnet_network_acl_attachment
Create, update, or delete a subnet network ACL attachment resource. For more information, about subnet network ACL attachment, see [setting up network ACLs](https://cloud.ibm.com/docs/vpc?topic=vpc-using-acls).

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
resource "ibm_is_network_acl" "example" {
  name = "example-acl"
  rules {
    name        = "outbound"
    action      = "allow"
    source      = "0.0.0.0/0"
    destination = "0.0.0.0/0"
    direction   = "outbound"
    icmp {
      code = 1
      type = 1
    }
  }
  rules {
    name        = "inbound"
    action      = "allow"
    source      = "0.0.0.0/0"
    destination = "0.0.0.0/0"
    direction   = "inbound"
    icmp {
      code = 1
      type = 1
    }
  }
}

resource "ibm_is_subnet" "example" {
  name            = "example-subnet"
  vpc             = ibm_is_vpc.example.id
  zone            = "us-south-1"
  ipv4_cidr_block = "192.168.0.0/1"

}

resource "ibm_is_subnet_network_acl_attachment" "example" {
  subnet      = ibm_is_subnet.example.id
  network_acl = ibm_is_network_acl.example.id
}

```
## Argument reference
Review the argument references that you can specify for your resource. 

- `network_acl` - (Optional, String) The network ACL identity.
- `subnet` - (Optional, Forces new resource, String) The subnet identifier.


## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `created_at` - (Timestamp) The creation date and time the network ACL.
- `crn` - (String) The CRN of this network ACL.
- `href` - (String) The URL of this network ACL.
- `id` - (String) The unique identifier of this network ACL.
- `name` - (String) The user-defined name of this network ACL.
- `protocol` - (List) The protocol list to enforce.
	
  Nested scheme for `protocol`:
  - `icmp` - (List) The protocol ICMP.

    Nested scheme for `icmp`:
	  - `code` - (String) The ICMP traffic code to allow. If unspecified, all codes are allowed. This can only be specified if type is also specified.
	  - `type` - (String) The ICMP traffic type to allow. If unspecified, all types are allowed by this rule.
	- `tcp` - (List) The TCP protocol.

    Nested scheme for `tcp`:
	  - `destination_port_max` - (String) The inclusive maximum bound of TCP destination port range.
	  - `destination_port_min` - (String) The inclusive minimum bound of TCP destination port range.
	  - `source_port_max` - (String) The inclusive maximum bound of TCP source port range.
	  - `source_port_min` - (String) The inclusive minimum bound of TCP source port range.
	- `udp` - (List) The UDP protocol.

    Nested scheme for `udp`:
	  - `destination_port_max` - (String) The inclusive maximum bound of UDP destination port range.
	  - `destination_port_min` - (String) The inclusive minimum bound of UDP destination port range.
	  - `source_port_max` - (String) The inclusive maximum bound of UDP source port range.
	  - `source_port_min` - (String) The inclusive minimum bound of UDP source port range.
	- `subnets` - (String) The subnets to which this network ACL is attached.
	- `vpc` - (String) The VPC to which this network ACL is a part of.
- `resource_group` - (String) The resource group (Id), of this network ACL.
- `rules` - (List) The ordered rules of this network ACL. If rules does not exist, all traffic will be denied. Nested rules blocks has the following structure.

  Nested scheme for `rules`:
	- `action` - (String) Specify to allow or deny matching traffic.
	- `created_at` - (String) The rule creation date and time.
	- `source` - (String) The source CIDR block. The CIDR block 0.0.0.0/0 applies to all addresses.
	- `destination` - (String) The destination CIDR block. The CIDR block 0.0.0.0/0 applies to all addresses.
	- `href` - (String) The URL of the Network ACL rule.
	- `id` - (String) The unique identifier of the Network ACL rule.
	- `ip_version` - (String) The IP version of the rule.
	- `name` - (String) The user-defined name of the rule.

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import the `ibm_is_subnet_network_acl_attachment` resource by using `id`.
The `id` property can be formed from subnet ID. For example:

```terraform
import {
  to = ibm_is_subnet_network_acl_attachment.example
  id = "<subnet_ID>"
}
```

Using `terraform import`. For example:

```console
% terraform import ibm_is_subnet_network_acl_attachment.example <subnet_network_acl_ID>
```