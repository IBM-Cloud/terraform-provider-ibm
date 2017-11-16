---
layout: "ibm"
page_title: "IBM: security_group_rule"
sidebar_current: "docs-ibm-resource-security-group-rule"
description: |-
  Manages IBM Security Group Rules
---

# ibm\_security_group_rule

Provide a Rule inside a Security Group resource. This allows rules to be created, updated and deleted.

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

* `direction` - (Required, string) Direction of traffic: ingress or egress
* `ether_type` - (Optional, string) IP version: IPv4 or IPv6 (case sensitive). Defaults to 'IPv4'
* `port_range_min` - (Optional, int) Lower bound of port range to allow
* `port_range_max` - (Optional, int) Upper bound of port range to allow
* `protocol` - (Optional, string) Traffic protocol: icmp or tcp or udp (case sensitive)
* `remote_group_id` - (Optional, int) The ID of the remote security group allowed as part of the rule.

    **NOTE**: Conflicts with `remote_ip`.
* `remote_ip` - (Optional, string) CIDR or IP address for allowed connections.

    **NOTE**: Conflicts with `remote_group_id`.
* `protocol` - (Optional, string) Traffic protocol: icmp or tcp or udp (case sensitive)
* `security_group_id` - (Required, int) The ID of the security group this rule belongs to.

## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier of the new security group
