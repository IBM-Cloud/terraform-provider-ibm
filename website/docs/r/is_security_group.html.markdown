---
layout: "ibm"
page_title: "IBM : security_group"
sidebar_current: "docs-ibm-resource-is-security-group"
description: |-
  Manages IBM Security Group.
---

# ibm\_is_security_group

Provides a security group resource. This allows security group to be created, updated, and cancelled.


## Example Usage

```hcl
resource "ibm_is_vpc" "testacc_vpc" {
	name = "test"
}

resource "ibm_is_security_group" "testacc_security_group" {
	name = "test"
	vpc = "${ibm_is_vpc.testacc_vpc.id}"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional, string) The security group name.
* `vpc` - (Required, string) The vpc id. 

## Attribute Reference

The following attributes are exported:

* `id` - The id of the security group.
* `rules` - A nested block describing the rules of this security group.
Nested `rules` blocks have the following structure:
  * `direction` -  The direction of the traffic either `inbound` or `outbound`.
  * `ip_version` - IP version either `ipv4` or `ipv6`.
  * `remote` - Security group id - an IP address, a CIDR block, or a single security group identifier.
  * `protocol` - The type of the protocol `all`, `icmp`, `tcp`, `udp`. 
  * `type` - The ICMP traffic type to allow.
  * `code` - The ICMP traffic code to allow.
  * `port-max` - The inclusive upper bound of TCP/UDP port range.
  * `port_min` - The inclusive lower bound of TCP/UDP port range. 
   
## Import

ibm_is_security_group can be imported using lbID, eg

```
$ terraform import ibm_is_security_group.example d7bec597-4726-451f-8a63-e62e6f19c32c
```
