---

subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : network_acl_rule"
description: |-
  Manages IBM Network ACL rule.
---

# ibm_is_network_acl_rule

Provides a network ACL rule resource with `icmp`, `tcp`, `udp` or `all` protocol. This allows Network ACL rule to create, update, and delete an existing network ACL. For more information, about managing IBM Cloud Network ACL , see [about network acl](https://cloud.ibm.com/docs/vpc?topic=vpc-using-acls).

**Note:** 
VPC infrastructure services are a regional specific based endpoint, by default targets to `us-south`. Please make sure to target right region in the provider block as shown in the `provider.tf` file, if VPC service is created in region other than `us-south`.

**provider.tf**

```terraform
provider "ibm" {
  region = "eu-gb"
}
```

## Example usage (all)

```terraform
resource "ibm_is_vpc" "example" {
  name = "example-vpc"
}

resource "ibm_is_network_acl" "example" {
  name = "example-acl"
  vpc  = ibm_is_vpc.example.id
}
resource "ibm_is_network_acl_rule" "example" {
  network_acl = ibm_is_network_acl.example.id
  name        = "outbound"
  action      = "allow"
  source      = "0.0.0.0/0"
  destination = "0.0.0.0/0"
  direction   = "outbound"
}
resource "ibm_is_network_acl_rule" "example1" {
  network_acl = ibm_is_network_acl.example.id
  name        = "inbound"
  action      = "allow"
  source      = "0.0.0.0/0"
  destination = "0.0.0.0/0"
  direction   = "inbound"
}
```

## Example usage (icmp)

```terraform
resource "ibm_is_network_acl_rule" "example" {
  network_acl = ibm_is_network_acl.example.id
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
resource "ibm_is_network_acl_rule" "example1" {
  network_acl = ibm_is_network_acl.example.id
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

```

## Example usage (tcp/udp)

```terraform
resource "ibm_is_network_acl_rule" "example" {
  network_acl = ibm_is_network_acl.example.id
  name        = "outbound"
  action      = "allow"
  source      = "0.0.0.0/0"
  destination = "0.0.0.0/0"
  direction   = "outbound"
  tcp {
    port_max        = 65535
    port_min        = 1
    source_port_max = 60000
    source_port_min = 22
  }
}
resource "ibm_is_network_acl_rule" "example1" {
  network_acl = ibm_is_network_acl.example.id
  name        = "inbound"
  action      = "allow"
  source      = "0.0.0.0/0"
  destination = "0.0.0.0/0"
  direction   = "inbound"
  tcp {
    port_max        = 65535
    port_min        = 1
    source_port_max = 60000
    source_port_min = 22
  }
}
```

## Argument reference
Review the argument references that you can specify for your resource.

- `action` - (Required, String) Whether to **allow** or **deny** matching traffic.
- `before` - (Optional, String) The unique identifier of the rule that this rule is immediately before. If unspecified, this rule will be inserted after all existing rules. While modifying the resource, specify **"null"** (within double quotes) to move this rule after all existing rules.
- `destination` - (Required, String) The destination IP address or CIDR block.
- `direction` - (Required, String) Whether the traffic to be matched is **inbound** or **outbound**.
- `icmp` - (Optional, List) The protocol ICMP.

   Nested scheme for `icmp`:
   - `code` - (Optional, Integer) The ICMP traffic code to allow. Valid values from 0 to 255. If unspecified, all codes are allowed. This can only be specified if type is also specified.
   - `type` - (Optional, Integer) The ICMP traffic type to allow. Valid values from 0 to 254. If unspecified, all types are allowed by this rule.
- `network_acl` - (Required, String) The ID of the network ACL.
- `name` - (Required, String) The user-defined name for this rule.
- `source` - (Required, String) The source IP address or CIDR block.
- `tcp` - (Optional, List) TCP protocol.

   Nested scheme for `tcp`:
   - `port_max` - (Optional, Integer) The highest port in the range of ports to be matched; if unspecified, **65535** is used.
   - `port_min` - (Optional, Integer) The lowest port in the range of ports to be matched; if unspecified, **1** is used.
   - `source_port_max` - (Optional, Integer) The highest port in the range of ports to be matched; if unspecified, **65535** is used.
   - `source_port_min` - (Optional, Integer) The lowest port in the range of ports to be matched; if unspecified, **1** is used.
- `udp` - (Optional, List) UDP protocol

   Nested scheme for `udp`:
   - `port_max` - (Optional, Integer) The highest port in the range of ports to be matched; if unspecified, **65535** is used.
   - `port_min` - (Optional, Integer) The lowest port in the range of ports to be matched; if unspecified, **1** is used.
   - `source_port_max` - (Optional, Integer) The highest port in the range of ports to be matched; if unspecified, **65535** is used.
   - `source_port_min` - (Optional, Integer) The lowest port in the range of ports to be matched; if unspecified, **1** is used.

~> **NOTE:**: Only one type of protocol out of **icmp**, **tcp**, or **udp** can be used to create a new rule. If none is provided, **all** is selected.

## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `id` - (String) The ID of the network ACL rule. The ID is composed of `\<network_acl\>/\<rule_id\>.`
- `href` - (String) The URL for this network ACL rule.
- `protocol` - (String) The protocol to enforce.
- `rule_id` - (String) The unique identifier of the rule.


## Import
The `ibm_is_network_acl_rule` can be imported using ID `\<network_acl\>/\<rule_id\>`

**Example**

```
$ terraform import ibm_is_network_acl_rule.example d7bec597-4726-451f-8a63-e62e6f19c32c/cea6651a-bc0a-4438-9f8a-a0770bbf3ebb
```
 