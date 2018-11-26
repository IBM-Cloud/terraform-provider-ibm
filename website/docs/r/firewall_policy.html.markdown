---
layout: "ibm"
page_title: "IBM : firewal_policy"
sidebar_current: "docs-ibm-resource-firewall-policy"
description: |-
  Manages IBM Firewall Policy.
---

# ibm\_firewall_policy

Provides rules for firewall resources in IBM. One rule resource is allowed per firewall. However, a rule resource can contain multiple firewall rules within it.

For more details about how to configure a firewall, see the [docs](https://knowledgelayer.softlayer.com/procedure/configure-hardware-firewall-dedicated).  

**NOTE**: The target VLAN should have at least one subnet for rule configuration. To express any IP addresses externally, configure `src_ip_address` as `0.0.0.0` and `src_ip_cidr` as `0`. To express API IP addresses internally, configure `dst_ip_address` as `any` and `src_ip_cidr` as `32`.

When a rules resource is created, it cannot be deleted. IBM does not allow entire rule deletion.

Firewalls should have at least one rule. If Terraform destroys the rules resources, _permit from any to any with TCP, UDP, ICMP, GRE, PPTP, ESP, and HA_ rules will be configured.

## Example Usage

```hcl
resource "ibm_firewall" "demofw" {
  ha_enabled = false
  public_vlan_id = 1234567
}

resource "ibm_firewall_policy" "rules" {
 firewall_id = "${ibm_firewall.demofw.id}"
 rules = {
      "action" = "permit"
      "src_ip_address"= "10.1.1.0"
      "src_ip_cidr"= 24
      "dst_ip_address"= "any"
      "dst_ip_cidr"= 32
      "dst_port_range_start"= 80
      "dst_port_range_end"= 80
      "notes"= "Permit from 10.1.1.0"
      "protocol"= "udp"
 }
  rules = {
       "action" = "deny"
       "src_ip_address"= "10.1.1.0"
       "src_ip_cidr"= 24
       "dst_ip_address"= "any"
       "dst_ip_cidr"= 32
       "dst_port_range_start"= 81
       "dst_port_range_end"= 81
       "notes"= "Permit from 10.1.1.0"
       "protocol"= "udp"
  }
}
```

## Argument Reference

The following arguments are supported:

* `firewall_id` - (Required, integer) The device ID for the target hardware firewall.
* `rules` - (Required, array) The firewall rules. At least one rule is required.
* `rules.action` - (Required, string) Specifies whether traffic is allowed when rules are matched. Accepted values are `permit` or `deny`.
* `rules.src_ip_address` - (Required, string) Specifies either a specific IP address or the network address for a specific subnet.
* `rules.src_ip_cidr` - (Required, string) Specifies the standard CIDR notation for the selected source. `32` implements the rule for a single IP while, for example, `24` implements the rule for 256 IPs.
* `rules.dst_ip_address` - (Required, string) Accepted values are `any`, a specific IP address, or the network address for a specific subnet.
* `rules.dst_ip_cidr` - (Required, string) Specifies the standard CIDR notation for the selected destination.
* `rules.dst_port_range_start` - (Optional, string) The start of the range of ports for TCP and UDP. Accepted values are `1` - `65535`.
* `rules.dst_port_range_end` - (Optional, string) The end of the range of ports for TCP and UDP. Accepted values are `1` - `65535`.
* `rules.notes` - (Optional, string) Descriptive text about the rule.
* `rules.protocol` - (Required, string) The protocol for the rule. Accepted values are `tcp`,`udp`,`icmp`,`gre`,`pptp`,`ah`, or `esp`.
* `tags` - (Optional, array of strings) Tags associated with the firewall policy instance.  
  **NOTE**: `Tags` are managed locally and not stored on the IBM Cloud service endpoint at this moment.
