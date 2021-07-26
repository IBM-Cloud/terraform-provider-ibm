---

subcategory: "Classic infrastructure"
layout: "ibm"
page_title: "IBM : firewal_policy"
description: |-
  Manages IBM Cloud firewall Policy.
---

# ibm_firewall_policy
Provides rules for firewall resources in IBM. One rule resource is allowed per firewall. However, a rule resource can contain multiple firewall rules within it. For an overview of supported firewalls rules, see [hardware firewall (Dedicated) rules](https://cloud.ibm.com/docs/hardware-firewall-dedicated?topic=hardware-firewall-dedicated-bypassing-hardware-firewall-dedicated-rules).

**Note** 

The target VLAN should have at least one subnet for rule configuration. To express any IP addresses externally, configure `src_ip_address` as `0.0.0.0` and `src_ip_cidr` as `0`. To express API IP addresses internally, configure `dst_ip_address` as `any` and `src_ip_cidr` as `32`.

When a rules resource is created, it cannot be deleted. IBM does not allow entire rule deletion.

Firewalls should have at least one rule. If  Terraform destroys the rules resources, _permit from any to any with `TCP`, `UDP`, `ICMP`, `GRE`, `PPTP`, `ESP`, and `HA_` rule to be configured.

## Example usage

```terraform
resource "ibm_firewall" "demofw" {
  ha_enabled     = false
  public_vlan_id = 1234567
}


resource "ibm_firewall_policy" "rules" {
  firewall_id = ibm_firewall.demofw.id
  rules {
    action               = "permit"
    src_ip_address       = "10.1.1.0"
    src_ip_cidr          = 24
    dst_ip_address       = "any"
    dst_ip_cidr          = 32
    dst_port_range_start = 80
    dst_port_range_end   = 80
    notes                = "Permit from 10.1.1.0"
    protocol             = "udp"
  }
  rules {
    action               = "deny"
    src_ip_address       = "10.1.1.0"
    src_ip_cidr          = 24
    dst_ip_address       = "any"
    dst_ip_cidr          = 32
    dst_port_range_start = 81
    dst_port_range_end   = 81
    notes                = "Permit from 10.1.1.0"
    protocol             = "udp"
  }
}

```

## Argument reference
Review the argument references that you can specify for your resource. 

- `firewall_id` - (Required, Force new resource, Integer) The device ID for the target hardware firewall.
- `rules`- (Required, List) The firewall rules. At least one rule is required.

  Nested scheme for `rules`:
  - `action` - (Required, String) Specifies whether traffic is allowed when rules are matched. Accepted values are `permit` or `deny`.
  - `dst_ip_address` - (Required, String) Accepted values are `any`, a specific IP address, or the network address for a specific subnet.
  - `dst_ip_cidr` - (Required, String) Specifies the standard CIDR notation for the selected destination.
  - `dst_port_range_start`- (Optional, String) The start of the range of ports for TCP and UDP. Accepted values are `1`- `65535`.
  - `dst_port_range_end`-  (Optional, String) The end of the range of ports for TCP and UDP. Accepted values are `1`- `65535`.
  - `notes`-  (Optional, String)  Descriptive text about the rule.
  - `protocol` - (Required, String) The protocol for the rule. Accepted values are `tcp`,`udp`,`icmp`,`gre`,`pptp`,`ah`, or `esp`.
  - `src_ip_address` - (Required, String) Specifies either a specific IP address or the network address for a specific subnet.
  - `src_ip_cidr`- (Required, String) Specifies the standard CIDR notation for the selected source. `32` implements the rule for a single IP while, for example, `24` implements the rule for 256 IP's.
- `tags`- (Optional, Array of Strings) Tags associated with the firewall policy instance. **Note** `Tags` are managed locally and not stored on the IBM Cloud Service Endpoint at this moment.
