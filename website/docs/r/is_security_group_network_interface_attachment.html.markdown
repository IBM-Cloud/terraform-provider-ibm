---

subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : security_group_network_interface_attachment"
description: |-
  Manages IBM security group network interface attachment.
---

# ibm_is_security_group_network_interface_attachment
Create, update, or delete a security group network interface attachment. For more information, about security group network interface attachment, see [attaching and detaching security groups](https://cloud.ibm.com/docs/vpc?topic=vpc-alb-integration-with-security-groups#attaching-detaching-sg-to-alb).

## Example usage

```terraform
resource "ibm_is_security_group_network_interface_attachment" "sgnic" {
  security_group    = "2d364f0a-a870-42c3-a554-000001352417"
  network_interface = "6d6128aa-badc-45c4-bb0e-7c2c1c47be55"
}
```
**Note** This resource is deprecated. Use `ibm_is_security_group_target` to attach a network interface to a security group

## Argument reference
Review the argument references that you can specify for your resource. 

- `security_group` - (Required, Forces new resource, String) The security group ID. 
- `network_interface` - (Required, Forces new resource, String) The network interface ID.

## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `floating_ips` - (List of Objects) A nested block describes the floating IP's of this network interface.

  Nested scheme for `floating_ips`:
	- `address` - (String) The globally unique IP address.
	- `crn` - (String) The CRN of this floating IP.
	- `id` - (String) The ID of this floating IP.
	- `name` - (String) The name of this floating IP.
- `id` - (String) The ID of the security group network interface. The ID is composed of `<security_group_id>/<network_interface_id>`.
- `instance_network_interface` - (String) The instance network interface ID.
- `name` - (String) The user-defined name for this network interface.
- `port_speed`- (Integer) The network interface port speed in Mbps.
- `primary_ipv4_address` - (String) The primary IPv4 address.
- `secondary_address` - (Array) Collection secondary IP addresses.
- `status` - (String) The status of the volume.
- `subnet` - (String) The Subnet ID.
- `security_groups` - (List of Objects) A nested block describes the security groups of this network interface.

  Nested scheme for `security_groups`:
	- `crn` - (String) The CRN of this security group.
	- `id` - (String) The ID of this security group.
	- `name` - (String) The name of this security group.
- `type` - (String) The type of this network interface as it relates to a instance.



## Import
The `ibm_is_security_group_network_interface_attachment` resource can be imported by using security group ID and network interface ID.

**Syntax**

```
$ terraform import ibm_is_security_group_network_interface_attachment.example <security_group_ID>/<network_interface_ID>
```

**Example**

```
$ terraform import ibm_is_security_group_network_interface_attachment.example d7bec597-4726-451f-8a63-e62e6f19c32c/cea6651a-bc0a-4438-9f8a-a0770bbf3ebb
```
