---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : network_acl_rules"
description: |-
  Manages IBM Network Acl Rules.
---

# ibm_is_network_acl_rules

Import the details of an existing IBM Cloud Infrastructure Network ACL Rules as a read-only data source. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax. For more information, about managing IBM Cloud Network ACL , see [about network acl](https://cloud.ibm.com/docs/vpc?topic=vpc-using-acls).

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
data "ibm_is_network_acl_rules" "example"{
  network_acl = ibm_is_network_acl.example.id
}
```

## Argument reference

Review the argument references that you can specify for your resource.

- `network_acl` - (Required, String) The network ACL identifier.
- `direction` - (Optional, String) The direction of the rules to filter. Available options are `inbound` and `outbound`

## Attribute reference

In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `rules` - (List of Objects) List of all rules in the network acl.

  - `action` - (String) Whether to allow or deny matching traffic.
  - `before` - (String) The unique identifier of the rule that this rule is immediately before.If absent, this is the last rule.
  - `destination` - (String) The destination IP address or CIDR block.
  - `direction` - (String) Whether the traffic to be matched is inbound or outbound.
  - `href` - (String) The URL for this network ACL rule.  
  - `icmp` - (List) The protocol ICMP 

    Nested scheme for `icmp`:
    - `code` - (Integer) The ICMP traffic code to allow. Valid values from 0 to 255. If unspecified, all codes are allowed. This can only be specified if type is also specified.
    - `type` - (Integer) The ICMP traffic type to allow. Valid values from 0 to 254. If unspecified, all types are allowed by this rule.
  - `ip_version` - (String) The IP version for this rule.
  - `name` - (String) The user-defined name for this rule.  
  - `protocol` - (String) The protocol to enforce.  
  - `rule_id` - (String) The network ACL rule identifier.
  - `source` - (String) The source IP address or CIDR block.
  - `tcp` - (List) TCP protocol.

    Nested scheme for `tcp`:
    - `port_max` - (Integer) The highest port in the range of ports to be matched; if unspecified, 65535 is used.
    - `port_min` - (Integer) The lowest port in the range of ports to be matched; if unspecified, 1 is used.
    - `source_port_max` - (Integer) The highest port in the range of ports to be matched; if unspecified, 65535 is used.
    - `source_port_min` - (Integer) The lowest port in the range of ports to be matched; if unspecified, 1 is used.
  - `udp` - (List) UDP protocol

    Nested scheme for `udp`:
    - `port_max` - (Integer) The highest port in the range of ports to be matched; if unspecified, 65535 is used.
    - `port_min` - (Integer) The lowest port in the range of ports to be matched; if unspecified, 1 is used.
    - `source_port_max` - (Integer) The highest port in the range of ports to be matched; if unspecified, 65535 is used.
    - `source_port_min` - (Integer) The lowest port in the range of ports to be matched; if unspecified, 1 is used.