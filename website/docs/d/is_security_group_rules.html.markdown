---
layout: "ibm"
page_title: "IBM : ibm_is_security_group_rules"
description: |-
  Get information about SecurityGroupRuleCollection
subcategory: "VPC infrastructure"
---

# ibm_is_security_group_rules

Provides a read-only data source for SecurityGroupRuleCollection. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.
For more information, about security group rule, see [security in your VPC](https://cloud.ibm.com/docs/vpc?topic=vpc-security-in-your-vpc).

## Example Usage

```hcl
data "ibm_is_security_group_rules" "example" {
  security_group = ibm_is_security_group.example.id
}
```

## Argument Reference

Review the argument reference that you can specify for your data source.

- `security_group` - (Required, String) The security group identifier.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

- `id` - The unique identifier of the SecurityGroupRuleCollection.
- `rules` - (List) Array of rules.
Nested scheme for `rules`:
	- `code` - (Integer) The ICMP traffic code to allow.
	- `direction` - (String) The direction of traffic to enforce, either `inbound` or `outbound`.
	- `href` - (String) The URL for this security group rule.
	- `id` - (String) The unique identifier for this security group rule.
	- `ip_version` - (String) The IP version to enforce. The format of `remote.address` or `remote.cidr_block` must match this property, if they are used. Alternatively, if `remote` references a security group, then this rule only applies to IP addresses (network interfaces) in that group matching this IP version.
	- `port_max` - (Integer) The inclusive upper bound of TCP/UDP port range.
	- `port_min` - (Integer) The inclusive lower bound of TCP/UDP port range.
	- `protocol` - (String) The protocol to enforce.
	- `remote` - (List) The IP addresses or security groups from which this rule allows traffic (or to which,for outbound rules). Can be specified as an IP address, a CIDR block, or a securitygroup. A CIDR block of `0.0.0.0/0` allows traffic from any source (or to any source,for outbound rules).
	Nested scheme for `remote`:
		- `address` - (String) The IP address.This property may add support for IPv6 addresses in the future. When processing a value in this property, verify that the address is in an expected format. If it is not, log an error. Optionally halt processing and surface the error, or bypass the resource on which the unexpected IP address format was encountered.
		- `cidr_block` - (String) The CIDR block. This property may add support for IPv6 CIDR blocks in the future. When processing a value in this property, verify that the CIDR block is in an expected format. If it is not, log an error. Optionally halt processing and surface the error, or bypass the resource on which the unexpected CIDR block format was encountered.
		- `crn` - (String) The security group's CRN.
		- `deleted` - (List) If present, this property indicates the referenced resource has been deleted and providessome supplementary information.
		Nested scheme for `deleted`:
			- `more_info` - (String) Link to documentation about deleted resources.
		- `href` - (String) The security group's canonical URL.
		- `id` - (String) The unique identifier for this security group.
		- `name` - (String) The user-defined name for this security group. Names must be unique within the VPC the security group resides in.
	- `type` - (Integer) The ICMP traffic type to allow.

