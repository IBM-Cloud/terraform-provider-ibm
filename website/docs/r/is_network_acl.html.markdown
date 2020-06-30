---
layout: "ibm"
page_title: "IBM : network acl"
sidebar_current: "docs-ibm-resource-is-network-acl"
description: |-
  Manages IBM network acl.
---

# ibm\_is_network_acl

Provides a network ACL resourcewith icmp protocol. This allows network ACL to be created, updated, and cancelled.


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
```
Provides a network ACL resource with tcp/udp protocol. This allows network ACL to be created, updated, and cancelled.


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
    tcp {
      port_max        = 65535
      port_min        = 1
      source_port_max = 60000
      source_port_min = 22
    }
  }
  rules {
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
}
```

Provides a NextGen VPC Network ACL resource with icmp protocol. This allows network ACL to be created, updated, and cancelled.


## Example Usage

```hcl
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

## Argument Reference

The following arguments are supported:



* `name` - (Required, string) The name of the network ACL.
* `vpc` - (Optional, Forces new resource, string) The VPC Id. This is a Required field and to be set only when the generation parameter is `2`
* `resource_group` - (Optional, Forces new resource, string) The resource group ID where the Network ACL is to be created. Should be set only when the generation parameter is `2`
* `rules` - (Optional, array)   The rules for a network ACL. The order of rules priority depends on the order of rules specified in the template.
Nested `rules` blocks have the following structure:
	* `name` - (Required, string) The user-defined name for this rule.
	* `action` - (Required, string) Whether to allow or deny matching traffic.
	* `source` - (Required, string) The source IP address or CIDR block.
	* `destination` - (Required, string) The destination IP address or CIDR block.
	* `direction` - (Required, string) Whether the traffic to be matched is inbound or outbound.
	* `icmp` - (Optional, array) The protocol ICMP
		* `code` - (Optional, int) The ICMP traffic code to allow. Valid values from 0 to 255. If unspecified, all codes are allowed. This can only be specified if type is also specified.
		* `type` - (Optional, int) The ICMP traffic type to allow. Valid values from 0 to 254. If unspecified, all types are allowed by this rule.
	* `tcp` - (Optional, array) TCP protocol.
		* `port_max` - (Optional, int) The highest port in the range of ports to be matched; if unspecified, 65535 is used.
		* `port_min` - (Optional, int) The lowest port in the range of ports to be matched; if unspecified, 1 is used.
		* `source_port_max` - (Optional, int) The highest port in the range of ports to be matched; if unspecified, 65535 is used.
		* `source_port_min` - (Optional, int) The lowest port in the range of ports to be matched; if unspecified, 1 is used.
	* `udp` - (Optional, array) UDP protocol
		* `port_max` - (Optional, int) The highest port in the range of ports to be matched; if unspecified, 65535 is used.
		* `port_min` - (Optional, int) The lowest port in the range of ports to be matched; if unspecified, 1 is used.
		* `source_port_max` - (Optional, int) The highest port in the range of ports to be matched; if unspecified, 65535 is used.
		* `source_port_min` - (Optional, int) The lowest port in the range of ports to be matched; if unspecified, 1 is used.
		

## Attribute Reference

The following attributes are exported:

* `id` - The id of the network ACL.
* `rules` - The rules for a network ACL.
Nested `rules` blocks have the following structure:
	* `id` - The rule id.
	* `ip_version` - The IP version of the rule.
	* `subnets` - The subnets for the ACL rule.

## Import

ibm_is_network_acl can be imported using ID, eg

```
$ terraform import ibm_is_network_acl.example d7bec597-4726-451f-8a63-e62e6f19c32c
```
