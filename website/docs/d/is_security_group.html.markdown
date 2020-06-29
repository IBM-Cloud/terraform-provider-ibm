---
layout: "ibm"
page_title: "IBM : security_group"
sidebar_current: "docs-ibm-resource-is-security-group"
description: |-
  Reads IBM Cloud Security Group.
---

# ibm\_is_security_group

Import the details of a security group as a read-only data source. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.


## Example Usage

In the following example, you can create a different types of protocol rules `ALL`, `ICMP`, `UDP`, `TCP` and read the security group.

```hcl
resource "ibm_is_vpc" "testacc_vpc" {
  name = "test"
}

resource "ibm_is_security_group" "testacc_security_group" {
  name = "test"
  vpc  = ibm_is_vpc.testacc_vpc.id
}

resource "ibm_is_security_group_rule" "testacc_security_group_rule_all" {
  group     = ibm_is_security_group.testacc_security_group.id
  direction = "inbound"
  remote    = "127.0.0.1"
}

resource "ibm_is_security_group_rule" "testacc_security_group_rule_icmp" {
  group     = ibm_is_security_group.testacc_security_group.id
  direction = "inbound"
  remote    = "127.0.0.1"
  icmp {
    code = 20
    type = 30
  }
}

resource "ibm_is_security_group_rule" "testacc_security_group_rule_udp" {
  group     = ibm_is_security_group.testacc_security_group.id
  direction = "inbound"
  remote    = "127.0.0.1"
  udp {
    port_min = 805
    port_max = 807
  }
}

resource "ibm_is_security_group_rule" "testacc_security_group_rule_tcp" {
  group     = ibm_is_security_group.testacc_security_group.id
  direction = "egress"
  remote    = "127.0.0.1"
  tcp {
    port_min = 8080
    port_max = 8080
  }
}

data "ibm_is_security_group" "sg1_rule" {
  name = ibm_is_security_group.testacc_security_group.name
}
```

## Argument Reference

The following arguments are supported:

`name` - (Required, string) The name of the security group.

## Attribute Reference

The following attributes are exported:

* `id` - The id of the security group. 
* `rules` - Rules associated with security group. Each rule has follwoing attributes
  * `rule_id` - ID of the rule.
  * `direction` - Direction of traffic to enforce, either inbound or outbound.
  * `ip_version` - IP version: ipv4 or ipv6.
  * `remote` - Security group id, an IP address, a CIDR block, or a single security group identifier.
  * `type` - The traffic type to allow.
  * `code` - The traffic code to allow.
  * `port_max` - The inclusive upper bound of TCP/UDP port range.
  * `port_min` - The inclusive lower bound of TCP/UDP port range.
  * `protocol` - The type of the protocol all, icmp, tcp, udp.
  


