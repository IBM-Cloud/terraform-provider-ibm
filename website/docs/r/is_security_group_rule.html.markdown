---
layout: "ibm"
page_title: "IBM : security_group_rule"
sidebar_current: "docs-ibm-resource-is-security-group-rule"
description: |-
  Manages IBM Security Group Rule.
---

# ibm\_is_security_group_rule

Provides a security group rule resource. This allows security group rule to be created, updated, and cancelled.


## Example Usage

In the following example, you can create a different types of protocol rules `ALL`, `ICMP`, `UDP` and `TCP`.

```hcl
resource "ibm_is_vpc" "testacc_vpc" {
	name = "test"
}

resource "ibm_is_security_group" "testacc_security_group" {
	name = "test"
	vpc = "${ibm_is_vpc.testacc_vpc.id}"
}

resource "ibm_is_security_group_rule" "testacc_security_group_rule_all" {
	group = "${ibm_is_security_group.testacc_security_group.id}"
	direction = "ingress"
	remote = "127.0.0.1"
 }
 
 resource "ibm_is_security_group_rule" "testacc_security_group_rule_icmp" {
	group = "${ibm_is_security_group.testacc_security_group.id}"
	direction = "ingress"
	remote = "127.0.0.1"
	icmp = {
		code = 20
		type = 30
	}

 }

 resource "ibm_is_security_group_rule" "testacc_security_group_rule_udp" {
	group = "${ibm_is_security_group.testacc_security_group.id}"
	direction = "ingress"
	remote = "127.0.0.1"
	udp = {
		port_min = 805
		port_max = 807
	}
 }

 resource "ibm_is_security_group_rule" "testacc_security_group_rule_tcp" {
	group = "${ibm_is_security_group.testacc_security_group.id}"
	direction = "egress"
	remote = "127.0.0.1"
	tcp = {
		port_min = 8080
		port_max = 8080
	}
 }
```

## Argument Reference

The following arguments are supported:

* `group` - (Required, string) The security group id.
* `direction` - (Required, string)  The direction of the traffic either `ingress` or `egress`.
* `remote` - (Required, string) Security group id - an IP address, a CIDR block, or a single security group identifier.
* `ip_version` - (Optional, string) IP version either `IPv4` or `IPv6`. Default `IPv4`.
* `icmp` - (Optional, list) A nested block describing the `icmp` protocol of this security group rule.
  * `type` - (Required, int) The ICMP traffic type to allow. Valid values from 0 to 254.
  * `code` - (Optional, int) The ICMP traffic code to allow. Valid values from 0 to 255.
* `tcp` - (Optional, list) A nested block describing the `tcp` protocol of this security group rule.
  * `port_min` - (Required, int) The inclusive lower bound of TCP port range. Valid values are from 1 to 65535.
  * `port-max` - (Required, int) The inclusive upper bound of TCP port range. Valid values are from 1 to 65535.
* `udp` - (Optional, list) A nested block describing the `udp` protocol of this security group rule.
  * `port_min` - (Required, int) The inclusive lower bound of UDP port range. Valid values are from 1 to 65535.
  * `port-max` - (Required, int) The inclusive upper bound of UDP port range. Valid values are from 1 to 65535.

**NOTE**: If any of the `icmp` , `tcp` or `udp` is not specified it creates a rule with protocol `ALL`. 


## Attribute Reference

The following attributes are exported:

* `id` - The id of the security group rule. The id is composed of \<security_group_id\>/\<security_group_rule_id\>.
* `rule_id` - The unique identifier of the rule.

## Import

ibm_is_security_group_rule can be imported using security group ID and security group rule ID, eg

```
$ terraform import ibm_is_security_group_rule.example d7bec597-4726-451f-8a63-e62e6f19c32c/cea6651a-bc0a-4438-9f8a-a0770bbf3ebb
```
