---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : network_acl"
description: |-
  Get information about IBM Network ACL.

---

# ibm_is_network_acl
Retrieve information of an network ACL data source. For more information, about managing network ACL, see [create a network acl](hhttps://cloud.ibm.com/docs/vpc?topic=vpc-acl-create-ui&interface=ui).

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

data "ibm_is_network_acl" "example" {
  network_acl = ibm_is_network_acl.example.id
}

data "ibm_is_network_acl" "is_network_acl1" {
  name     = ibm_is_network_acl.example.name
  vpc_name = ibm_is_vpc.example.name
}
```

## Argument reference

Review the argument reference that you can specify for your data source.

- `name` - (Optional, String) The name of the network ACL.
- `network_acl` - (Optional, String) The network ACL identifier.
- `vpc_name` - (Optional, String) The name of the VPC.
  **Note** Provide `network_acl` or the combination of `vpc_name` and `name`.

## Attribute reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

- `created_at` - (String) The date and time that the network ACL was created.
- `crn` - (String) The CRN for this network ACL.
- `href` - (String) The URL for this network ACL.
- `id` - The unique identifier of the NetworkACL.
- `name` - (String) The user-defined name for this network ACL.
- `resource_group` - (List) The resource group object, for this network ACL.

  Nested scheme for `resource_group`:
	- `href` - (String) The URL for this resource group.
	- `id` - (String) The unique identifier for this resource group.
	- `name` - (String) The user-defined name for this resource group.
- `rules`- (Array of Strings) A list of rules for a network ACL. The order in which the rules are added to the list determines the priority of the rules. For example, the first rule that you want to enforce must be specified as the first rule in this list.

  Nested scheme for `rules`:
  - `name` - (String) The user-defined name for this rule.
  - `action` - (String)  `Allow` or `deny` matching network traffic.
  - `source` - (String) The source IP address or CIDR block.
  - `destination` - (String) The destination IP address or CIDR block.
  - `direction` - (String) Indicates whether the traffic to be matched is `inbound` or `outbound`.
  - `icmp`- (List) The protocol ICMP.

    Nested scheme for `icmp`:
    - `code` - (Integer) The ICMP traffic code to allow. Valid values from `0 to 255`. If unspecified, all codes are allowed. This can only be specified if type is also specified.
    - `type` - (Integer) The ICMP traffic type to allow. Valid values from `0 to 254`. If unspecified, all types are allowed by this rule.
  - `tcp`- (List) The TCP protocol.

    Nested scheme for `tcp`:
    - `port_max` - (Integer) The highest port in the range of ports to be matched; if unspecified, `65535` is used.
    - `port_min` - (Integer) The lowest port in the range of ports to be matched, if unspecified, `1` is used as default.
    - `source_port_max` - (Integer) The highest port in the range of ports to be matched; if unspecified, `65535` is used as default.
    - `source_port_min` - (Integer) The lowest port in the range of ports to be matched; if unspecified, `1` is used as default.
  - `udp`- (List) The UDP protocol.

    Nested scheme for `udp`:
    - `port_max` - (Integer) The highest port in the range of ports to be matched; if unspecified, `65535` is used.
    - `port_min` - (Integer) The lowest port in the range of ports to be matched; if unspecified, `1` is used.
    - `source_port_max` - (Integer) The highest port in the range of ports to be matched; if unspecified, `65535` is used.
    - `source_port_min` - (Integer) The lowest port in the range of ports to be matched; if unspecified, `1` is used.
- `subnets` - (List) The subnets to which this network ACL is attached.

  Nested scheme for **subnets**:
	- `crn` - (String) The CRN for this subnet.
	- `deleted` - (List) If present, this property indicates the referenced resource has been deleted and provides some supplementary information.

	  Nested scheme for **deleted**:
		- `more_info` - (String) Link to documentation about deleted resources.

	- `href` - (String) The URL for this subnet.
	- `id` - (String) The unique identifier for this subnet.
	- `name` - (String) The user-defined name for this subnet.
- `vpc` - (List) The VPC this network ACL is a part of.

  Nested scheme for **vpc**:
	- `crn` - (String) The CRN for this VPC.
	- `deleted` - (List) If present, this property indicates the referenced resource has been deleted and provides some supplementary information.

	  Nested scheme for **deleted**:
		- `more_info` - (String) Link to documentation about deleted resources.

	- `href` - (String) The URL for this VPC.
	- `id` - (String) The unique identifier for this VPC.
	- `name` - (String) The unique user-defined name for this VPC.

