---

subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : bare_metal_server_network_interface"
description: |-
  Manages IBM bare metal sever network interface.
---

# ibm\_is_bare_metal_server_network_interface

Create, update, or delete a Network Interface on an existing Bare Metal Server for VPC. For more information, about managing VPC Bare Metal Server, see [Network of Bare Metal Servers for VPC](https://cloud.ibm.com/docs/vpc?topic=vpc-bare-metal-servers-network). User `is_bare_metal_server_network_interface_allow_float` resource to create a vlan type network interface with allow float.

**Note:** 
VPC infrastructure services are a regional specific based endpoint, by default targets to `us-south`. Please make sure to target right region in the provider block as shown in the `provider.tf` file, if VPC service is created in region other than `us-south`.

**provider.tf**

```terraform
provider "ibm" {
  region = "eu-gb"
}
```

## Example Usage

In the following example, you can create a Bare Metal Server and add a network interface to it:

```terraform
resource "ibm_is_vpc" "vpc" {
  name = "testvpc"
}

resource "ibm_is_subnet" "subnet" {
  name            = "testsubnet"
  vpc             = ibm_is_vpc.vpc.id
  zone            = "us-south-3"
  ipv4_cidr_block = "10.240.129.0/24"
}

resource "ibm_is_ssh_key" "ssh" {
  name       = "testssh"
  public_key = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR"
}

resource "ibm_is_bare_metal_server" "bms" {
  profile = "mx2d-metal-32x192"
  name    = "my-bms"
  image   = "r134-31c8ca90-2623-48d7-8cf7-737be6fc4c3e"
  zone    = "us-south-3"
  keys    = [ibm_is_ssh_key.sshkey.id]
  primary_network_interface {
    subnet     = ibm_is_subnet.subnet.id
  }
  vpc = ibm_is_vpc.vpc.id
}

resource "ibm_is_bare_metal_server_network_interface" "bms_nic" {
  bare_metal_server = ibm_is_bare_metal_server.bms.id

  subnet = ibm_is_subnet.subnet.id
  name   = "eth2"
  allow_ip_spoofing = true
  allowed_vlans = [101, 102]
}

```
  ~> **NOTE**
    Creating/Deleting a PCI type network interface would stop and start the bare metal server. Use `hard_stop` to configure the stopping type, by default `hard` is enabled, make it `false` to `soft` stop the server 

## Timeouts

ibm_is_bare-metal_server_network_interface provides the following [Timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) configuration options:

* `create` - (Default 10 minutes) Used for creating bare metal server network interface .
* `update` - (Default 10 minutes) Used for updating bare metal server network interface.
* `delete` - (Default 10 minutes) Used for deleting bare metal server network interface.


## Argument reference
Review the argument references that you can specify for your resource. 

- `allowed_vlans` - (Optional, Integer) Indicates what VLAN IDs (for VLAN type only) can use this physical (PCI type) interface. A given VLAN can only be in the allowed_vlans array for one PCI type adapter per bare metal server. This property which controls the VLANs that will be permitted to use the pci interface.

  ~> **NOTE**
    Creates a PCI type interface, a physical PCI device can only be created or deleted when the bare metal server is stopped. Use `hard_stop` as `false` to `soft` stop the server, by default its `hard`

- `allow_ip_spoofing` - (Optional, Boolean) Indicates whether source IP spoofing is allowed on this interface. If false, source IP spoofing is prevented on this interface. If true, source IP spoofing is allowed on this interface.
- `bare_metal_server` - (Required, String) The id for this bare metal server.
- `enable_infrastructure_nat` - (Optional, Boolean) If true, the VPC infrastructure performs any needed NAT operations. If false, the packet is passed unmodified to/from the network interface, allowing the workload to perform any needed NAT operations.
- `hard_stop` - (Optional, Boolean) Default is `true`. Applicable for `pci` type only, controls if the server should be hard stopped.
- `name` - (Optional, String) The user-defined name for this network interface
- `primary_ip` - (Optional, List)
	- `address` - (Optional, String) title: IPv4 The IP address. This property may add support for IPv6 addresses in the future. When processing a value in this property, verify that the address is in an expected format. If it is not, log an error. Optionally halt processing and surface the error, or bypass the resource on which the unexpected IP address format was encountered.
  - `auto_delete` - (Optional, Bool) Indicates whether this reserved IP member will be automatically deleted when either target is deleted, or the reserved IP is unbound.
  - `reserved_ip`- (Optional, String) The unique identifier for this reserved IP. `id` is mutually exclusive with rest of the `primary_ip` attributes.
  - `name`- (Optional, String) The user-defined or system-provided name for this reserved IP

- `security_groups` - (Optional, List) Collection of security groups
- `subnet` - (Required, String) The associated subnet
- `vlan` - (Optional, Integer) Indicates the 802.1Q VLAN ID tag that must be used for all traffic on this interface

  ~> **NOTE**
    Creates a vlan type network interface, a virtual device, used through a pci device that has the vlan in its array of allowed_vlans. 

## Attribute reference
In addition to the argument reference list, you can access the following attribute references after your data source is created. 

- `allow_interface_to_float` - (Boolean) Indicates if the interface can float to any other server within the same resource_group. The interface will float automatically if the network detects a GARP or RARP on another bare metal server in the resource group. Applies only to vlan type interfaces.
- `floating_ips` - (List) The floating IPs associated with this network interface.

  Nested scheme for `floating_ips`:
  - `address` - (String) The floating IP address.
  - `id` - (String) The unique identifier of the floating IP.
- `href` - (String) The URL for this network interface
- `id` - (String) The unique identifier for this network interface. Its of the format <bare_metal_server_id>/<network_interface_id>
- `interface_type` - (String) The network interface type, supported values are [ **pci**, **vlan** ]
- `mac_address` - (String) The MAC address of the interface. If absent, the value is not known.
- `network_interface` - (String) The network interface id.
- `port_speed` - (Integer) The network interface port speed in Mbps
- `resource_type` - (String)The resource type [ **subnet_reserved_ip** ]
- `status` - (String) The status of the network interface. Supported values are [ **available**, **deleting**, **failed**, **pending** ]
- `type` - (String) The type of this bare metal server network interface. Supported values are [ **primary**, **secondary** ]



## Import

ibm_is_bare_metal_server can be imported using bare metal server ID and network interface id

## Syntax

```
$ terraform import ibm_is_bare_metal_server_network_interface.example <bare_metal_server_id>/<bare_metal_server_network_interface_id>
```

## Example 

```
$ terraform import ibm_is_bare_metal_server_network_interface.example d7bec597-4726-451f-8a63-e62e6f19c32c/e7bec597-4726-451f-8a63-e62e6f19c32d
```
