---
layout: "ibm"
page_title: "IBM : security_group_network_interface_attachment"
sidebar_current: "docs-ibm-resource-is-security-group-network-interface-attachment"
description: |-
  Manages IBM Security Group Network Interface Attachment.
---

# ibm\_is_security_group_network_interface_attachment

Provides a security group network interface attachment resource. This allows security group network interface attachment to be created, updated, and cancelled.


## Example Usage

```hcl
resource "ibm_is_security_group_network_interface_attachment" "sgnic" {
  security_group    = "2d364f0a-a870-42c3-a554-000001352417"
  network_interface = "6d6128aa-badc-45c4-bb0e-7c2c1c47be55"
}
```

## Argument Reference

The following arguments are supported:

* `security_group` - (Required, string) The security group id.
* `network_interface` - (Required, string) The network interface id. 

## Attribute Reference

The following attributes are exported:

* `id` - The id of the security group network interface. The id is composed of \<security_group_id\>/\<network_interface_id\>.
* `instance_network_interface` - The instance network interface id.
* `name` - The user-defined name for this network interface.
* `port_speed` - The network interface port speed in Mbp.
* `primary_ipv4_address` - The network interface port speed in Mbp.
* `primary_ipv6_address` - The primary IPv6 address in compressed notation as specified by RFC 5952.
* `secondary_address` - Collection seconary IP addresses.
* `status` - The status of the volume.
* `subnet` - The Subnet id.
* `type` - The type of this network interface as it relates to a instance.
* `security_groups` -  A nested block describing the security groups of this network interface.
Nested `security_groups` blocks have the following structure:
	* `id` - The id of this security group.
	* `crn` - The CRN of this security group.
	* `name` - The name of this security group.

* `floating_ips` - A nested block describing the floating ip's of this network interface.
Nested `floating_ips` blocks have the following structure:
  * `id` - The id of this floating Ip.
  * `crn` - The CRN of this floating Ip.
  * `name` - The name of this floating Ip.
  * `address` - The globally unique IP address

## Import

ibm_is_security_group_network_interface_attachment can be imported using security group ID and network interface ID, eg

```
$ terraform import ibm_is_security_group_network_interface_attachment.example d7bec597-4726-451f-8a63-e62e6f19c32c/cea6651a-bc0a-4438-9f8a-a0770bbf3ebb
```
