---
layout: "ibm"
page_title: "IBM: security_group_rule"
sidebar_current: "docs-ibm-resource-security-group-rule"
description: |-
  Manages IBM Security Group Rules
---

# ibm\_security_group_rule

Provide a rule for a security group. You can set the IP range to manage incoming (ingress) and outgoing (egress) traffic to a virtual server instance. This resources allows rules for security groups to be created, updated, and deleted. To create the security group, use the `security_group` resource.

For additional details, see the [IBM Cloud Infrastructure (SoftLayer) API docs](http://sldn.softlayer.com/reference/datatypes/SoftLayer_Network_SecurityGroup_Rule).

## Example Usage

```
resource "ibm_security_group_rule" "allow_port_8080" {
    direction = "ingress"
    ether_type = "IPv4"
    port_range_min = 8080
    port_range_max = 8080
    protocol = "tcp"
    security_group_id = 123456
}
```

## Argument Reference

The following arguments are supported:

* `direction` - (Required, string) The direction of traffic. Accepted values: `ingress` or `egress`.
* `ether_type` - (Optional, string) The IP version. Accepted values  (case sensitive): `IPv4` or `IPv6`. Default value: 'IPv4'.
* `port_range_min` - (Optional, int) The start of the port range for allowed traffic.
* `port_range_max` - (Optional, int) The end of the port range for allowed traffic.
* `protocol` - (Optional, string) The IP protocol type. Accepted values (case sensitive): `icmp`,`tcp`, or `udp`.
* `remote_group_id` - (Optional, int) The ID of the remote security group allowed as part of the rule.  

    **NOTE**: Conflicts with `remote_ip`.
* `remote_ip` - (Optional, string) The CIDR or IP address for allowed connections.  

    **NOTE**: Conflicts with `remote_group_id`.
* `security_group_id` - (Required, int) The ID of the security group this rule belongs to.

## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier of the security group rule.
