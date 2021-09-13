---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : security_group"
description: |-
  Reads IBM Cloud security group.
---

# ibm_is_security_group
Retrieve information about a security group as a read-only data source. For more information, about managing IBM Cloud security group , see [about security group](https://cloud.ibm.com/docs/vpc?topic=vpc-using-security-groups).


## Example usage
The following example allows to create a different types of protocol rules `ALL`, `ICMP`, `UDP`, `TCP` and read the security group.

```terraform
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

## Argument reference
Review the argument references that you can specify for your resource. 

- `name` - (Required, String) The name of the security group.

## Attribute reference
In addition to all argument reference list, you can access the following attribute references after your data source is created. 

- `crn` - The CRN of the security group.
- `id` - (String) The ID of the security group.
- `rules` - (List of Objects) The rules associated with security group. Each rule has following attributes.

  Nested scheme for `rules`:
  - `rule_id`-  (String) ID of the rule.
  - `direction` - (String) Direction of traffic to enforce, either inbound or outbound.
  - `ip_version` - (String) IP version: IPv4
  - `protocol` - (String) The type of the protocol `all`, `icmp`, `tcp`, `udp`.
  - `type` - (String) The traffic type to allow.
  - `code` - (String) The traffic code to allow.
  - `port_max`- (Integer) The TCP/UDP port range that includes the maximum bound.
  - `port_min`- (Integer) The TCP/UDP port range that includes the minimum bound.
  - `remote`- (Integer)  Security group ID, an IP address, a CIDR block, or a single security group identifier.
- `tags` - Tags associated with the security group.
  


