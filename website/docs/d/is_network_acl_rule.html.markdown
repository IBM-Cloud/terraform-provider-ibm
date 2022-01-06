---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : network_acl_rule"
description: |-
  Manages IBM Cloud Network ACL rule.
---

# ibm_is_network_acl_rule

Retrieve information of an network ACL rule data source. For more information, about managing IBM Cloud Network ACL , see [about network acl](https://cloud.ibm.com/docs/vpc?topic=vpc-using-acls).

**Note:** 
VPC infrastructure services are a regional specific based endpoint, by default targets to `us-south`. Please make sure to target right region in the provider block as shown in the `provider.tf` file, if VPC service is created in region other than `us-south`.

**provider.tf**

```terraform
provider "ibm" {
  region = "eu-gb"
}
```

## Example usage

```terraform
resource "ibm_is_vpc" "example" {
  name = "example-vpc"
}

resource "ibm_is_network_acl" "example" {
  name = "example-network-acl"
  vpc  = ibm_is_vpc.example.id
}

resource "ibm_is_network_acl_rule" "example" {
  network_acl = ibm_is_network_acl.example.id
  name        = "example-network-acl-rule"
  action      = "allow"
  source      = "0.0.0.0/0"
  destination = "0.0.0.0/0"
  direction   = "outbound"
}

data "ibm_is_network_acl_rule" "example" {
  network_acl = ibm_is_network_acl.example.id
  name        = "example-network-acl-rule"
}
```

## Argument reference

Review the argument references that you can specify for your resource.

- `name` - (Required, String) The network ACL rule name.
- `network_acl` - (Required, String) The network ACL identifier.

## Attribute reference

In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `action` - (String) Whether to allow or deny matching traffic.
- `before` - (String) The unique identifier of the rule that this rule is immediately before.If absent, this is the last rule.
- `destination` - (String) The destination IP address or CIDR block.
- `direction` - (String) Whether the traffic to be matched is inbound or outbound.
- `href` - (String) The URL for this network ACL rule.
- `icmp` - (List) The protocol ICMP
    - `code` - (Integer) The ICMP traffic code to allow. Valid values from **0 to 255**. If unspecified, all codes are allowed. This can only be specified if type is also specified.
    - `type` - (Integer) The ICMP traffic type to allow. Valid values from **0 to 254**. If unspecified, all types are allowed by this rule.
- `id` - (String) The network ACL rule identfier. This id is composed of `<network_acl\>/\<rule_id\>`.
- `ip_version` - (String) The IP version for this rule.
- `protocol` - (String) The protocol to enforce.
- `rule_id` - (String) The network ACL rule ID.
- `source` - (String) The source IP address or CIDR block.
- `tcp` - (List) TCP protocol.
    - `port_max` - (Integer) The highest port in the range of ports to be matched; if unspecified, **65535** is used.
    - `port_min` - (Integer) The lowest port in the range of ports to be matched; if unspecified, **1** is used.
    - `source_port_max` - (Integer) The highest port in the range of ports to be matched; if unspecified, **65535** is used.
    - `source_port_min` - (Integer) The lowest port in the range of ports to be matched; if unspecified, **1** is used.
- `udp` - (List) UDP protocol
    - `port_max` - (Integer) The highest port in the range of ports to be matched; if unspecified, **65535** is used.
    - `port_min` - (Integer) The lowest port in the range of ports to be matched; if unspecified, **1** is used.
    - `source_port_max` - (Integer) The highest port in the range of ports to be matched; if unspecified, **65535** is used.
    - `source_port_min` - (Integer) The lowest port in the range of ports to be matched; if unspecified, **1** is used.
    