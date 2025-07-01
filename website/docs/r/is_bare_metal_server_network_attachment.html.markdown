---
layout: "ibm"
page_title: "IBM : ibm_is_bare_metal_server_network_attachment"
description: |-
  Manages is_bare_metal_server_network_attachment.
subcategory: "VPC infrastructure"
---

# ibm_is_bare_metal_server_network_attachment

Create, update, and delete is_bare_metal_server_network_attachments with this resource.

## Example Usage

```terraform
resource "ibm_is_bare_metal_server" "testacc_bms" {
    profile 			= "cx2-metal-96x192"
    name 				  = "${var.name}-bms"
    image 				= "r134-f47cc24c-e020-4db5-ad96-1e5be8b5853b"
    zone 				  = "${var.region}-2"
    keys 				  = ["r134-349aac10-ed14-4dc6-a95d-2ce66c3b447c"]
    primary_network_attachment {
        name = "vni-2"
        virtual_network_interface { 
            id = "0726-b1755e04-1430-48d7-971d-661ba2836b54"
        }
        allowed_vlans = [100, 102]
    }
    vpc 				  = "r134-6d509c8a-470e-4cdd-a82c-103f2353f5fc"
}
resource "ibm_is_bare_metal_server_network_attachment" "na" {
	bare_metal_server   = "0726-e17dbe53-25d2-42da-8532-bcb3b5a19f37"
		allowed_vlans     = [200, 202, 203]
        virtual_network_interface { 
            id = "0726-b1755e04-1430-48d7-971d-661ba2836b54"
        }
}

```

## Argument Reference

You can specify the following arguments for this resource.

- `allow_to_float` - (Optional, Boolean) Indicates if the bare metal server network attachment can automatically float to any other server within the same `resource_group`. The bare metal server network attachment will float automatically if the network detects a GARP or RARP on another bare metal server in the resource group. Applies only to bare metal server network attachments with `vlan` interface type.
- `allowed_vlans` - (Optional, List) Indicates what VLAN IDs (for VLAN type only) can use this physical (PCI type) attachment.
- `bare_metal_server` - (Required, Forces new resource, String) The bare metal server identifier.
- `interface_type` - (Required, String) The network attachment's interface type:- `hipersocket`: a virtual network device that provides high-speed TCP/IP connectivity  within a `s390x` based system- `pci`: a physical PCI device which can only be created or deleted when the bare metal  server is stopped  - Has an `allowed_vlans` property which controls the VLANs that will be permitted    to use the PCI attachment  - Cannot directly use an IEEE 802.1q VLAN tag.- `vlan`: a virtual device, used through a `pci` device that has the `vlan` in its  array of `allowed_vlans`.  - Must use an IEEE 802.1q tag.The enumerated values for this property are expected to expand in the future. When processing this property, check for and log unknown values. Optionally halt processing and surface the error, or bypass the resource on which the unexpected property value was encountered.
- `name` - (Optional, String) The name for this bare metal server network attachment. The name is unique across all network attachments for the bare metal server.
- `virtual_network_interface` - (Optional, List) The virtual network interface for this bare metal server network attachment.
	Nested schema for **virtual_network_interface**:
    - `allow_ip_spoofing` - (Optional, Boolean) Indicates whether source IP spoofing is allowed on this interface. If `false`, source IP spoofing is prevented on this interface. If `true`, source IP spoofing is allowed on this interface.
    - `auto_delete` - (Optional, Boolean) Indicates whether this virtual network interface will be automatically deleted when`target` is deleted.
    - `crn` - (String) The CRN for this virtual network interface.
    - `enable_infrastructure_nat` - (Optional, Boolean) If `true`:- The VPC infrastructure performs any needed NAT operations.- `floating_ips` must not have more than one floating IP. If `false`:- Packets are passed unchanged to/from the network interface,  allowing the workload to perform any needed NAT operations.- `allow_ip_spoofing` must be `false`.- If the virtual network interface is attached:  - The target `resource_type` must be `bare_metal_server_network_attachment`.  - The target `interface_type` must not be `hipersocket`.
    - `ips` - (Optional, List) The reserved IPs bound to this virtual network interface.May be empty when `lifecycle_state` is `pending`.
      Nested schema for **ips**:
      - `address` - (Required, String) The IP address.If the address has not yet been selected, the value will be `0.0.0.0`.This property may add support for IPv6 addresses in the future. When processing a value in this property, verify that the address is in an expected format. If it is not, log an error. Optionally halt processing and surface the error, or bypass the resource on which the unexpected IP address format was encountered.
      - `deleted` - (Optional, List) If present, this property indicates the referenced resource has been deleted, and providessome supplementary information.
      Nested schema for **deleted**:
        - `more_info` - (String) Link to documentation about deleted resources.
      - `href` - (String) The URL for this reserved IP.
      - `reserved_ip` - (Required, String) The unique identifier for this reserved IP.
      - `name` - (Required, String) The name for this reserved IP. The name is unique across all reserved IPs in a subnet.
      - `resource_type` - (Computed, String) The resource type.
    - `name` - (Optional, String) The name for this virtual network interface. The name is unique across all virtual network interfaces in the VPC.
    - `primary_ip` - (Optional, List) The reserved IP for this virtual network interface.
      Nested schema for **primary_ip**:
      - `address` - (Required, String) The IP address.If the address has not yet been selected, the value will be `0.0.0.0`.This property may add support for IPv6 addresses in the future. When processing a value in this property, verify that the address is in an expected format. If it is not, log an error. Optionally halt processing and surface the error, or bypass the resource on which the unexpected IP address format was encountered.
      - `deleted` - (Optional, List) If present, this property indicates the referenced resource has been deleted, and providessome supplementary information.
      Nested schema for **deleted**:
        - `more_info` - (Required, String) Link to documentation about deleted resources.
      - `href` - (Required, String) The URL for this reserved IP.
      - `reserved_ip` - (Required, String) The unique identifier for this reserved IP.
      - `name` - (Required, String) The name for this reserved IP. The name is unique across all reserved IPs in a subnet.
      - `resource_type` - (Computed, String) The resource type.
    - `protocol_state_filtering_mode` - (Optional, String) The protocol state filtering mode to use for this virtual network interface. 

        ~> **If auto, protocol state packet filtering is enabled or disabled based on the virtual network interface's target resource type:** 
            **&#x2022;** bare_metal_server_network_attachment: disabled </br>
            **&#x2022;** instance_network_attachment: enabled </br>
            **&#x2022;** share_mount_target: enabled </br>
    - `resource_group` - (Optional, List) The resource group id for this virtual network interface.
    - `security_groups` - (Optional, Array of string) The security group ids list for this virtual network interface.
    - `subnet` - (Optional, List) The associated subnet id.
- `vlan` - (Optional, Integer) Indicates the 802.1Q VLAN ID tag that must be used for all traffic on this attachment.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

- `id` - The unique identifier of the is_bare_metal_server_network_attachment.
- `bare_metal_server_network_attachment_id` - (String) The unique identifier for this bare metal server network attachment.
- `created_at` - (String) The date and time that the bare metal server network attachment was created.
- `href` - (String) The URL for this bare metal server network attachment.
- `lifecycle_state` - (String) The lifecycle state of the bare metal server network attachment.
- `port_speed` - (Integer) The port speed for this bare metal server network attachment in Mbps.
- `primary_ip` - (List) The primary IP address of the virtual network interface for the bare metal servernetwork attachment.
Nested schema for **primary_ip**:
	- `address` - (String) The IP address.If the address has not yet been selected, the value will be `0.0.0.0`.This property may add support for IPv6 addresses in the future. When processing a value in this property, verify that the address is in an expected format. If it is not, log an error. Optionally halt processing and surface the error, or bypass the resource on which the unexpected IP address format was encountered.
	- `deleted` - (List) If present, this property indicates the referenced resource has been deleted, and providessome supplementary information.
	Nested schema for **deleted**:
		- `more_info` - (String) Link to documentation about deleted resources.
	- `href` - (String) The URL for this reserved IP.
	- `id` - (String) The unique identifier for this reserved IP.
	- `name` - (String) The name for this reserved IP. The name is unique across all reserved IPs in a subnet.
	- `resource_type` - (String) The resource type.
- `resource_type` - (String) The resource type.
- `subnet` - (List) The subnet of the virtual network interface for the bare metal server networkattachment.
Nested schema for **subnet**:
	- `crn` - (String) The CRN for this subnet.
	- `deleted` - (List) If present, this property indicates the referenced resource has been deleted, and providessome supplementary information.
	Nested schema for **deleted**:
		- `more_info` - (String) Link to documentation about deleted resources.
	- `href` - (String) The URL for this subnet.
	- `id` - (String) The unique identifier for this subnet.
	- `name` - (String) The name for this subnet. The name is unique across all subnets in the VPC.
	- `resource_type` - (String) The resource type.
- `type` - (String) The bare metal server network attachment type.


## Import

You can import the `ibm_is_bare_metal_server_network_attachment` resource by using `id`.
The `id` property can be formed from `bare_metal_server`, and `id` in the following format:

```
<bare_metal_server>/<id>
```
- `bare_metal_server`: A string. The bare metal server identifier.
- `id`: A string. The bare metal server network attachment identifier.

# Syntax
```
$ terraform import ibm_is_bare_metal_server_network_attachment.is_bare_metal_server_network_attachment <bare_metal_server>/<id>
```
