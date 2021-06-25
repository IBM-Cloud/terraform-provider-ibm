---

subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : bare metal server network interface"
description: |-
  Manages IBM bare metal sever network interface.
---

# ibm\_is_bare_metal_server

Provides a Bare Metal Server  network interface resource. This allows Bare Metal Server  network interface to be created, updated, and cancelled.


## Example Usage

In the following example, you can create a Bare Metal Server:

```terraform
resource "ibm_is_bare_metal_server" "bms" {
    profile = "mx2d-metal-32x192"
    name = "my-bms"
    image = "r134-31c8ca90-2623-48d7-8cf7-737be6fc4c3e"
    zone = "us-south-3"
    keys = [ibm_is_ssh_key.sshkey.id]
    primary_network_interface {
      subnet     = ibm_is_subnet.subnet1.id
    }
    vpc = ibm_is_vpc.vpc1.id
}
resource ibm_is_bare_metal_server_network_interface this {
  bare_metal_server = ibm_is_bare_metal_server.bms.id

    subnet = ibm_is_subnet.this.id
    name   = "eth2"
    allow_ip_spoofing = true
    allowed_vlans = [101, 102]
}
```

## Timeouts

ibm_is_bare-metal_server_network_interface provides the following [Timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) configuration options:

* `create` - (Default 10 minutes) Used for creating bare metal server network interface .
* `update` - (Default 10 minutes) Used for updating bare metal server network interface.
* `delete` - (Default 10 minutes) Used for deleting bare metal server network interface.


## Argument reference
Review the argument references that you can specify for your resource. 

- `allow_interface_to_float` - (Optional, Boolean) Indicates if the interface can float to any other server within the same resource_group. The interface will float automatically if the network detects a GARP or RARP on another bare metal server in the resource group. Applies only to vlan type interfaces.
- `allowed_vlans` - (Optional, Integer) Indicates what VLAN IDs (for VLAN type only) can use this physical (PCI type) interface. A given VLAN can only be in the allowed_vlans array for one PCI type adapter per bare metal server.
- `allow_ip_spoofing` - (Optional, Boolean) Indicates whether source IP spoofing is allowed on this interface. If false, source IP spoofing is prevented on this interface. If true, source IP spoofing is allowed on this interface.
- `bare_metal_server` - (Required, String) The id for this bare metal server.
- `enable_infrastructure_nat` - (Optional, Boolean) If true, the VPC infrastructure performs any needed NAT operations. If false, the packet is passed unmodified to/from the network interface, allowing the workload to perform any needed NAT operations.
- `ips` - (Optional, List) The reserved IPs bound to this network interface.
- `name` - (Optional, String) The user-defined name for this network interface
- `primary_ip` - (Optional, List)
	- `address` - (Optional, String) title: IPv4 The IP address. This property may add support for IPv6 addresses in the future. When processing a value in this property, verify that the address is in an expected format. If it is not, log an error. Optionally halt processing and surface the error, or bypass the resource on which the unexpected IP address format was encountered.
	- `name` - (String) The user-defined or system-provided name for this reserved IP
- `security_groups` - (Optional, List) Collection of security groups
- `subnet` - (Required, List) The associated subnet
- `vlan` - (Optional, Integer) Indicates the 802.1Q VLAN ID tag that must be used for all traffic on this interface

## Attribute reference
In addition to the argument reference list, you can access the following attribute references after your data source is created. 

- `floating_ips` - (List) The floating IPs associated with this network interface.
- `href` - (String) The URL for this network interface
- `id` - (String) The unique identifier for this network interface
- `interface_type` - (String) The network interface type, supported values are [ **pci**, **vlan** ]
- `mac_address` - (String) The MAC address of the interface. If absent, the value is not known.
- `port_speed` - (Integer) The network interface port speed in Mbps
- `primary_ip` - (, List)

	- `href` - (String) The URL for this reserved IP
	- `id` - (String) The unique identifier for this reserved IP
	- `name` - (String) The user-defined or system-provided name for this reserved IP
- `resource_type` - (String)The resource type [ **subnet_reserved_ip** ]
- `status` - (String) The status of the network interface. Supported values are [ **available**, **deleting**, **failed**, **pending** ]
- `type` - (String) The type of this bare metal server network interface. Supported values are [ **primary**, **secondary** ]



## Import

ibm_is_bare_metal_server can be imported using bare metal server ID , eg

```
$ terraform import ibm_is_bare_metal_server.example d7bec597-4726-451f-8a63-e62e6f19c32c
```
