---
layout: "ibm"
page_title: "IBM : subnet network acl attachment"
sidebar_current: "docs-ibm-resource-is-subnet-network-acl-attachment"
description: |-
  Manages IBM Subnet Network ACL Attachment.
---

# ibm\_is_subnet_network_acl_attachment

Provides a subnet network ACL attachment resource. This allows subnet network ACL attachment to be created, updated, and cancelled.


## Example Usage

```hcl
resource "ibm_is_network_acl" "isExampleACL" {
  name = "is-example-acl"
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

resource "ibm_is_subnet" "testacc_subnet" {
  name            = "test_subnet"
  vpc             = ibm_is_vpc.testacc_vpc.id
  zone            = "us-south-1"
  ipv4_cidr_block = "192.168.0.0/1"

}

resource "ibm_is_subnet_network_acl_attachment" attach {
  subnet      = ibm_is_subnet.testacc_subnet.id
  network_acl = ibm_is_network_acl.isExampleACL.id
}

```

## Argument Reference

The following arguments are supported:

* `subnet` - (Required, Forces new resource string) The subnet identifier.
* `network_acl` - (Required, string) The network ACL identity.

## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier for this network ACL.
* `name` - The user-defined name for this network ACL.
* `created_at` - The date and time that the network ACL was created.
* `crn` - The CRN for this network ACL.
* `href` - The URL for this network ACL.
* `resource_group` - The resource group for this network ACL.
* `rules` - The ordered rules for this network ACL. If no rules exist, all traffic will be denied.
Nested `rules` blocks have the following structure:

        * `action` - Whether to allow or deny matching traffic.
        * `created_at` - The date and time that the rule was created
        * `source` - The source CIDR block. The CIDR block 0.0.0.0/0 applies to all addresses.
        * `destination` - The destination CIDR block. The CIDR block 0.0.0.0/0 applies to all addresses.
        * `direction` - Whether the traffic to be matched is inbound or outbound.
        * `href` - The URL for this Network ACL rule.
        * `id` - The unique identifier for this Network ACL rule.
        * `ip_version` - The IP version for this rule
        * `name` - The user-defined name for this rule.
	* `protocol` - The protocol to enforce.
        * `icmp` - The protocol ICMP
                * `code` - The ICMP traffic code to allow. If unspecified, all codes are allowed. This can only be specified if type is also specified.
                * `type` - The ICMP traffic type to allow. If unspecified, all types are allowed by this rule.
        * `tcp` - TCP protocol.
                * `destination_port_max` - The inclusive upper bound of TCP destination port range.  
                * `destination_port_min` - The inclusive lower bound of TCP destination port range.
                * `source_port_max` - The inclusive upper bound of TCP source port range.
                * `source_port_min` - The inclusive lower bound of TCP source port range.
        * `udp` - UDP protocol
                * `destination_port_max` - The inclusive upper bound of UDP destination port range. 
                * `destination_port_min` - The inclusive lower bound of UDP destination port range.
                * `source_port_max` - The inclusive upper bound of UDP source port range.
                * `source_port_min` - The inclusive lower bound of UDP source port range.
        * `subnets` - The subnets to which this network ACL is attached.
        * `vpc` - The VPC this network ACL is a part of. 


## Import

ibm_is_subnet_network_acl_attachment can be imported using ID, eg

```
$ terraform import ibm_is_subnet_network_acl_attachment.example d7bec597-4726-451f-8a63-e62e6f19c32c
```
