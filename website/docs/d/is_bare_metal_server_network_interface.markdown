---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : bare_metal_server_network_interface"
description: |-
  Manages IBM Cloud Bare Metal Server Network Interface.
---

# ibm\_is_bare_metal_server_network_interface

Import the details of an existing IBM Cloud Bare Metal Server Network Interface as a read-only data source. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax. For more information, about bare metal servers, see [Network of Bare Metal Servers for VPC](https://cloud.ibm.com/docs/vpc?topic=vpc-bare-metal-servers-network).

**provider.tf**

```terraform
provider "ibm" {
  region = "eu-gb"
}
```


## Example Usage

```terraform

data "ibm_is_bare_metal_server_network_interface" "ds_bms_nic" {
  bare_metal_server         = "xxxx-xxxxx-xxxxx-xxxx"
  network_interface         = "xxxx-xxxxx-xxxxx-xxxx"
}

```

## Argument reference
Review the argument references that you can specify for your data source. 

- `bare_metal_server` - (Required, String) The id for this bare metal server.
- `network_interface` - (Required, String) The id for this bare metal server network interface.

## Attribute reference
In addition to the argument reference list, you can access the following attribute references after your data source is created. 

- `allow_interface_to_float` - (Boolean) Indicates if the interface can float to any other server within the same resource_group. The interface will float automatically if the network detects a GARP or RARP on another bare metal server in the resource group. Applies only to vlan type interfaces.
- `allow_ip_spoofing` - (Boolean) Indicates whether source IP spoofing is allowed on this interface. If false, source IP spoofing is prevented on this interface. If true, source IP spoofing is allowed on this interface.
- `allowed_vlans` - (Array) Indicates what VLAN IDs (for VLAN type only) can use this physical (PCI type) interface. A given VLAN can only be in the allowed_vlans array for one PCI type adapter per bare metal server.
- `enable_infrastructure_nat` - (Boolean) If true, the VPC infrastructure performs any needed NAT operations. If false, the packet is passed unmodified to/from the network interface, allowing the workload to perform any needed NAT operations.
- `floating_ips` - (List) The floating IPs associated with this network interface.
- `href` - (String) The URL for this network interface
- `id` - (String) The unique identifier for this network interface
- `interface_type` - (String) The network interface type, supported values are [ **pci**, **vlan** ]
- `mac_address` - (String) The MAC address of the interface. If absent, the value is not known.
- `name` - (String) The user-defined name for this network interface
- `port_speed` - (Integer) The network interface port speed in Mbps
- `primary_ip` - (List)
	- `address` - (String) title: IPv4 The IP address. This property may add support for IPv6 addresses in the future. When processing a value in this property, verify that the address is in an expected format. If it is not, log an error. Optionally halt processing and surface the error, or bypass the resource on which the unexpected IP address format was encountered.
	- `href` - (String) The URL for this reserved IP
	- `reserved_ip` - (String) The unique identifier for this reserved IP
	- `name` - (String) The user-defined or system-provided name for this reserved IP
  - `resource_type` - (String)The resource type [ **subnet_reserved_ip** ]
- `security_groups` - (Array) Collection of security groups
- `status` - (String) The status of the network interface. Supported values are [ **available**, **deleting**, **failed**, **pending** ]
- `subnet` - (List) The associated subnet
- `type` - (String) The type of this bare metal server network interface. Supported values are [ **primary**, **secondary** ]
- `vlan` - (Integer) Indicates the 802.1Q VLAN ID tag that must be used for all traffic on this interface