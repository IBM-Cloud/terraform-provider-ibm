---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : is_instance_network_interface"
description: |-
  Get information about NetworkInterface
---

# ibm_is_instance_network_interface

Retrieve information of an exisitng network interface. For more information, about instance network interface, see [managing an network interfaces](https://cloud.ibm.com/docs/vpc?topic=vpc-using-instance-vnics).


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

resource "ibm_is_subnet" "example" {
  name            = "example-vpc"
  vpc             = ibm_is_vpc.example.id
  zone            = "us-south-1"
  ipv4_cidr_block = "10.240.0.0/24"
}

resource "ibm_is_ssh_key" "example" {
  name       = "example-ssh"
  public_key = file("~/.ssh/id_rsa.pub")
}

resource "ibm_is_instance" "example" {
  name    = "example-vpc"
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

data "ibm_is_instance_network_interface" "example" {
	instance_name = ibm_is_instance.example.name
	network_interface_name = is_instance_network_interface.example.name
}
```

## Argument reference

The following arguments are supported:

- `instance_name` - (Required, string) The name of the instance.
- `network_interface_name` - (Required, string) The name of the network interface.

## Attribute reference

In addition to all arguments above, the following attributes are exported:

- `id` - (String) The unique identifier of the NetworkInterface.
- `allow_ip_spoofing` - (Boolean) Indicates whether source IP spoofing is allowed on this interface. If false, source IP spoofing is prevented on this interface. If true, source IP spoofing is allowed on this interface.

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

- `name` - (String) The user-defined name for this network interface.

- `port_speed` - (Integer) The network interface port speed in Mbps.

- `primary_ip` - (List) The primary IP address to bind to the network interface. This can be specified using an existing reserved IP, or a prototype object for a new reserved IP.

    Nested scheme for `primary_ip`:
    - `address` - (String) The IP address. If the address has not yet been selected, the value will be 0.0.0.0. This property may add support for IPv6 addresses in the future. When processing a value in this property, verify that the address is in an expected format. If it is not, log an error. Optionally halt processing and surface the error, or bypass the resource on which the unexpected IP address format was encountered.
    - `href`- (String) The URL for this reserved IP
    - `name`- (String) The user-defined or system-provided name for this reserved IP
    - `reserved_ip`- (String) The unique identifier for this reserved IP
    - `resource_type`- (String) The resource type.
    
- `primary_ipv4_address` - (String) The primary IPv4 address. Same as `primary_ip.0.address`

- `resource_type` - (String) The resource type.

- `security_groups` - (List) Collection of security groups. Nested `security_groups` blocks have the following structure:
	- `crn` - (String) The security group's CRN.
	- `deleted` - (List) If present, this property indicates the referenced resource has been deleted and providessome supplementary information. Nested `deleted` blocks have the following structure:
		- `more_info` - (String) Link to documentation about deleted resources.
	- `href` - (String) The security group's canonical URL.
	- `id` - (String) The unique identifier for this security group.
	- `name` - (String) The user-defined name for this security group. Names must be unique within the VPC the security group resides in.

- `status` - (String) The status of the network interface.

- `subnet` - (List) The associated subnet. Nested `subnet` blocks have the following structure:
	- `crn` - (String) The CRN for this subnet.
	- `deleted` - (List) If present, this property indicates the referenced resource has been deleted and providessome supplementary information. Nested `deleted` blocks have the following structure:
		- `more_info` - (String) Link to documentation about deleted resources.
	- `href` - (String) The URL for this subnet.
	- `id` - (String) The unique identifier for this subnet.
	- `name` - (String) The user-defined name for this subnet.

- `type` - (String) The type of this network interface as it relates to an instance.

