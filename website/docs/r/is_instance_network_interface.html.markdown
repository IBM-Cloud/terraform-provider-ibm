---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : is_instance_network_interface"
description: |-
  Manages NetworkInterface.
---

# ibm_is_instance_network_interface

Create, update, or delete an instance network interface on VPC. For more information, about instance network interface, see [managing an network interface](https://cloud.ibm.com/docs/vpc?topic=vpc-using-instance-vnics).

**Note:**
- IBM Cloud terraform provider currently provides both a standalone `ibm_is_instance_network_interface` resource and a `network_interfaces` block defined in-line in the `ibm_is_instance` resource. At this time you cannot use the `network_interfaces` block inline with `ibm_is_instance` in conjunction with the standalone resource `ibm_is_instance_network_interface`. Doing so will create a conflict of network interfaces and will overwrite it.
- IBM Cloud terraform provider currently provides both a standalone `ibm_is_security_group_target` resource and a `security_groups` block defined in-line in the `ibm_is_instance_network_interface` resource to attach security group to a network interface target. At this time you cannot use the `security_groups` block inline with `ibm_is_instance_network_interface` in conjunction with the standalone resource `ibm_is_security_group_target`. Doing so will create a conflict of security groups attaching to the network interface and will overwrite it.
- VPC infrastructure services are a regional specific based endpoint, by default targets to `us-south`. Please make sure to target right region in the provider block as shown in the `provider.tf` file, if VPC service is created in region other than `us-south`.

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

resource "ibm_is_subnet" "example" {
  name            = "example-subnet"
  vpc             = ibm_is_vpc.example.id
  zone            = "us-south-1"
  ipv4_cidr_block = "10.240.0.0/24"
}

resource "ibm_is_ssh_key" "example" {
  name       = "example-ssh"
  public_key = file("~/.ssh/id_rsa.pub")
}

resource "ibm_is_instance" "example" {
  name    = "example-instance"
  image   = "a7a0626c-f97e-4180-afbe-0331ec62f32a"
  profile = "bc1-2x8"

  primary_network_interface {
    subnet = ibm_is_subnet.example.id
  }

  network_interfaces {
    name   = "eth1"
    subnet = ibm_is_subnet.example.id
  }

  vpc  = ibm_is_vpc.example.id
  zone = "us-south-1"
  keys = [ibm_is_ssh_key.example.id]
}

resource "ibm_is_instance_network_interface" "example" {
  instance = ibm_is_instance.example.id
  subnet = ibm_is_subnet.example.id
  allow_ip_spoofing = true
  name = "example-network-interface"
  primary_ipv4_address = "10.0.0.5"
}
```

## Argument reference

The following arguments are supported:

- `allow_ip_spoofing` - (Optional, Boolean) Indicates whether source IP spoofing is allowed on this interface. If false, source IP spoofing is prevented on this interface. If true, source IP spoofing is allowed on this interface. The default value is `false`.
- `floating_ip` - (Optional, String) The ID of the floating IP to attach to this network interface.
- `instance` - (Required, Forces new resource, String) The instance identifier.
- `name` - (Required, String) The user-defined name for this network interface.
- `primary_ip` - (Optional, List) The primary IP address to bind to the network interface. This can be specified using an existing reserved IP, or a prototype object for a new reserved IP.
    Nested scheme for `primary_ip`:
    - `auto_delete` - (Optional, Bool) Indicates whether this reserved IP member will be automatically deleted when either target is deleted, or the reserved IP is unbound.
    - `address` - (Optional, String) The IP address. If the address has not yet been selected, the value will be 0.0.0.0. This property may add support for IPv6 addresses in the future. When processing a value in this property, verify that the address is in an expected format. If it is not, log an error. Optionally halt processing and surface the error, or bypass the resource on which the unexpected IP address format was encountered.
    - `name`- (Optional, String) The user-defined or system-provided name for this reserved IP
    - `reserved_ip`- (Optional, String) The unique identifier for this reserved IP
- `primary_ipv4_address` - (Optional, Forces new resource, String) The primary IPv4 address. If specified, it must be an available address on the network interface's subnet. If unspecified, an available address on the subnet will be automatically selected.
- `security_groups` - (Optional, List of strings) A comma separated list of security groups to add to the primary network interface.
- `subnet` - (Required, Forces new resource, String) The unique identifier of the associated subnet.
  

~> **Note**
  - Only 1 floating IP can be attached to a VSI at any given time. Floating IP can be de-attached from one network interface and attached to a different network interface, but be sure to remove the floating_ip field from the previous network interface resource before adding it to a new resource. 
~> **Note**
  - `floating_ip` cannot be used in conjunction with the `target` argument of `ibm_is_floating_ip` resource and might cause cyclic dependency/unexpected issues if used used both ways.
~> **Note**
  - Using `ibm_is_security_group_target` to attach security groups to the network interface along with `security_groups` field in this resource could cause undesired behavior. Use either one of them to associate network interface to a security group.

## Attribute reference

In addition to all arguments above, the following attributes are exported:

- `created_at` - (String) The date and time that the network interface was created.
- `floating_ips` - (List) The floating IPs associated with this network interface. Nested `floating_ips` blocks have the following structure:
	- `address` - (String) The globally unique IP address.
	- `crn` - (String) The CRN for this floating IP.
	- `deleted` - (List) If present, this property indicates the referenced resource has been deleted and providessome supplementary information. Nested `deleted` blocks have the following structure:
		- `more_info` - (String) Link to documentation about deleted resources.
	- `href` - (String) The URL for this floating IP.
	- `id` - (String) The unique identifier for this floating IP.
	- `name` - (String) The unique user-defined name for this floating IP.
- `href` - (String) The URL for this network interface.
- `id` - (String) The unique identifier of the resource. Follows the format <instance_id>/<network_interface_id>.
- `network_interface` - (String) The unique identifier of the NetworkInterface.
- `port_speed` - (Integer) The network interface port speed in Mbps.
- `resource_type` - (String) The resource type.
- `status` - (String) The status of the network interface.
- `type` - (String) The type of this network interface as it relates to an instance.

## Import

You can import the `ibm_is_instance_network_interface` resource by using `id`.
The `id` property can be formed from `instance_ID`, and `network_interface_ID` in the following format:

```
<instance>/<network_interface>
```
- `instance`: A string. The instance identifier.
- `network_interface`: A string. The network interface identifier.

```
$ terraform import ibm_is_instance_network_interface.is_instance_network_interface <instance>/<network_interface>
```
