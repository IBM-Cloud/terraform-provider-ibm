---

subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : ibm_is_network_acl"
description: |-
  Manages IBM network ACL.
---

# ibm_is_network_acl
Create, update, or delete a network access control list (ACL). For more information, about network ACL, see [setting up network ACLs](https://cloud.ibm.com/docs/vpc?topic=vpc-using-acls).

## Example usage

```terraform
resource "ibm_is_vpc" "testacc_vpc" {
  name = "vpctest"
}

resource "ibm_is_network_acl" "isExampleACL" {
  name = "is-example-acl"
  vpc  = ibm_is_vpc.testacc_vpc.id
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
```

## Argument reference
Review the argument references that you can specify for your resource. 
 
- `name` - (Required, String) The name of the network ACL.
- `resource_group` - (Optional, Forces new resource, String) The ID of the resource group where you want to create the network ACL.
- `rules`- (Optional, Array of Strings) A list of rules for a network ACL. The order in which the rules are added to the list determines the priority of the rules. For example, the first rule that you want to enforce must be specified as the first rule in this list.

  Nested scheme for `rules`:
  - `name` - (Required, String) The user-defined name for this rule.
  - `action` - (Required, String)  `Allow` or `deny` matching network traffic.
  - `source` - (Required, String) The source IP address or CIDR block.
  - `destination` - (Required, String) The destination IP address or CIDR block.
  - `direction` - (Required, String) Indicates whether the traffic to be matched is `inbound` or `outbound`.
  - `icmp`- (Optional, List) The protocol ICMP.

    Nested scheme for `icmp`:
    - `code` - (Optional, Integer) The ICMP traffic code to allow. Valid values from 0 to 255. If unspecified, all codes are allowed. This can only be specified if type is also specified.
    - `type` - (Optional, Integer) The ICMP traffic type to allow. Valid values from 0 to 254. If unspecified, all types are allowed by this rule.
  - `tcp`- (Optional, List) The TCP protocol.

    Nested scheme for `tcp`:
    - `port_max` - (Optional, Integer) The highest port in the range of ports to be matched; if unspecified, 65535 is used.
    - `port_min` - (Optional, Integer) The lowest port in the range of ports to be matched, if unspecified, 1 is used as default.
    - `source_port_max` - (Optional, Integer) The highest port in the range of ports to be matched; if unspecified, 65535 is used as default.
    - `source_port_min` - (Optional, Integer) The lowest port in the range of ports to be matched; if unspecified, 1 is used as default.
  - `udp`- (Optional, List) The UDP protocol.

    Nested scheme for `udp`:
    - `port_max` - (Optional, Integer) The highest port in the range of ports to be matched; if unspecified, 65535 is used.
    - `port_min` - (Optional, Integer) The lowest port in the range of ports to be matched; if unspecified, 1 is used.
    - `source_port_max` - (Optional, Integer) The highest port in the range of ports to be matched; if unspecified, 65535 is used.
    - `source_port_min` - (Optional, Integer) The lowest port in the range of ports to be matched; if unspecified, 1 is used.
- `tags`- (Optional, List of Strings) Tags associated with the network ACL.
- `vpc` - (Optional, Forces new resource, String) The VPC ID. This parameter is required if you want to create a network ACL for a Generation 2 VPC.

## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `crn` - (String) The CRN of the network ACL.
- `id` - (String) The ID of the network ACL.
- `rules`- (List) The rules for a network ACL.

  Nested scheme for `rules`:
  - `id` - (String) The rule ID.
  - `ip_version` - (String) The IP version of the rule.
  - `subnets` - (String) The subnets for the ACL rule.

## Import
The `ibm_is_network_acl` resource can be imported by using the network ACL ID. 

**Syntax**

```
$ terraform import ibm_is_network_acl.example <network_acl_id>
```

**Example**

```
$ terraform import ibm_is_network_acl.example d7bec597-4726-451f-8a63-1111132c
```

