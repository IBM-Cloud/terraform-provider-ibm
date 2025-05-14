---
layout: "ibm"
page_title: "IBM : is_share_mount_target"
description: |-
  Manages ShareTarget.
subcategory: "VPC infrastructure"
---


# ibm\_is_share_mount_target

Provides a resource for ShareMountTarget. This allows ShareTarget to be created, updated and deleted.


## Example Usage

```hcl
resource "ibm_is_vpc" "example" {
  name = "my-vpc"
}

resource "ibm_is_share" "example" {
  name = "my-share"
  size = 200
  profile = "dp2"
  zone = "us-south-2"
}

resource "ibm_is_share_mount_target" "example" {
  access_protocol = "nfs4"
  share = ibm_is_share.example.id
  vpc = ibm_is_vpc.example.id
  name = "my-share-target"
  transit_encryption = "none"
}`
```
```
//Create mount target with virtual network interface that has primary ip name and subnet id
resource "ibm_is_vpc" "example1" {
  name = "my-vpc"
}

resource "ibm_is_share" "example1" {
  access_control_mode = "security_group"
  zone    = "br-sao-2"
  size    = 9600
  name    = "my-example-share1"
  profile = "dp2"
}
resource "ibm_is_subnet" "example1" {
  # provider = ibm.sao
  name                     = "my-subnet"
  vpc                      = ibm_is_vpc.example1.id
  zone                     = "br-sao-2"
  total_ipv4_address_count = 16
}

resource "ibm_is_share_mount_target" "example1" {
  access_protocol = "nfs4"
  share = ibm_is_share.example1.id
  virtual_network_interface {
    primary_ip {
      name = "my-example-pip"
    }
    subnet = ibm_is_subnet.example1.id
    name = "my-example-vni"
  }
  name  = "my-example-mount-target"
  transit_encryption = "ipsec"
}

//Create a mount target with subnet id
resource "ibm_is_share_mount_target" "example2" {
  access_protocol = "nfs4"
  share = ibm_is_share.example.id
  virtual_network_interface {
    subnet = ibm_is_subnet.example.id
    name = "my-example-vni"
  }
  name  = "my-example-mount-target"
  transit_encryption = "ipsec"
}

//Create mount target with reserved ip id
resource "ibm_is_subnet_reserved_ip" "example" {
  subnet = ibm_is_subnet.example.id
  name = "my-example-resip"
}
resource "ibm_is_share_mount_target" "example" {
  access_protocol = "nfs4"
  share = ibm_is_share.example.id
  virtual_network_interface {
    primary_ip {
      reserved_ip = ibm_is_subnet_reserved_ip.example.reserved_ip
    }
    name = "my-example-vni"
  }
  name  = "my-example-mount-target"
  transit_encryption = "ipsec"
}

//Create mount target with VNI ID
resource "ibm_is_subnet" "example" {
  name                     = "my-subnet"
  vpc                      = ibm_is_vpc.vpc2.id
  zone                     = "br-sao-2"
  total_ipv4_address_count = 16
}
resource "ibm_is_virtual_network_interface" "example" {
  name   = "my-example-vni"
  subnet = ibm_is_subnet.example.id
}
resource "ibm_is_share_mount_target" "mtarget1" {
  access_protocol = "nfs4"
  share = ibm_is_share.share.id
  virtual_network_interface {
    id = ibm_is_virtual_network_interface.example.id
  }
  name = "my-example-mount-target"
  transit_encryption = "ipsec"
}
```
## Argument Reference

The following arguments are supported:

- `share` - (Required, String) The file share identifier.
- `access_protocol` - (Required, String) The protocol to use to access the share for this share mount target. The specified value must be listed in the share's allowed_access_protocols. Available values are `nfs4`
- `virtual_network_interface` (Optional, List) The virtual network interface for this share mount target. Required if the share's `access_control_mode` is `security_group`.
  - `name` - (Required, String) Name for this virtual network interface. The name must not be used by another virtual network interface in the VPC.
  Nested scheme for `virtual_network_interface`:
  - `id` - (Optional) The ID for virtual network interface. Mutually exclusive with other `virtual_network_interface` arguments.
  
  ~> **Note**
    `id` is mutually exclusive with other `virtual_network_interface` prototype arguments
  - `primary_ip` - (Optional, List) The primary IP address to bind to the virtual network interface. May be either a reserved IP identity, or a reserved IP prototype object which will be used to create a new reserved IP.

      Nested scheme for `primary_ip`:
      
      - `auto_delete` - (Optional, Bool) Indicates whether this reserved IP member will be automatically deleted when either target is deleted, or the reserved IP is unbound. Defaults to `true`
      - `address` - (Optional, Forces new resource, String) The IP address to reserve. If unspecified, an available address on the subnet will automatically be selected.
      
      - `name`- (Optional, String) The name for this reserved IP. The name must not be used by another reserved IP in the subnet. Names starting with ibm- are reserved for provider-owned resources, and are not allowed.
      - `reserved_ip`- (Optional, String) The unique identifier for this reserved IP.

  ~> **Note**
    Within `primary_ip`, `reserved_ip` is mutually exclusive to  `auto_delete`, `address` and `name`
  - `protocol_state_filtering_mode` - (Optional, String) The protocol state filtering mode to use for this virtual network interface. 

        ~> **If auto, protocol state packet filtering is enabled or disabled based on the virtual network interface's target resource type:** 
            **&#x2022;** bare_metal_server_network_attachment: disabled </br>
            **&#x2022;** instance_network_attachment: enabled </br>
            **&#x2022;** share_mount_target: enabled </br>
  - `resource_group` - (Optional, String) The ID of the resource group to use.
  - `security_groups`- (Optional, List of string) The security groups to use for this virtual network interface.
  - `subnet` - (Optional, string) The associated subnet.
    
    

- `vpc` - (Optional, string) The VPC in which instances can mount the file share using this share target. Required if the share's `access_control_mode` is vpc.
  ~> **Note**
  `virtual_network_interface` and `vpc` are mutually exclusive and one of them must be provided.
- `name` - (Required, String) The user-defined name for this share target. Names must be unique within the share the share target resides in. If unspecified, the name will be a hyphenated list of randomly-selected words.
- `transit_encryption` - (Required, String) The transit encryption mode for this share target. Supported values are **none**, **ipsec** and **stunnel**

## Attribute Reference

The following attributes are exported:


- `access_control_mode` - (String) The access control mode for the share.
- `allow_ip_spoofing` - (Bool) Indicates whether source IP spoofing is allowed on this interface. If false, source IP spoofing is prevented on this interface. If true, source IP spoofing is allowed on this interface.
- `auto_delete` - (Bool) Indicates whether this virtual network interface will be automatically deleted when target is deleted
- `enable_infrastructure_nat` - (Bool) If `true`:- The VPC infrastructure performs any needed NAT operations.- `floating_ips` must not have more than one floating IP.If `false`:- Packets are passed unchanged to/from the network interface,  allowing the workload to perform any needed NAT operations.- `allow_ip_spoofing` must be `false`.- If the virtual network interface is attached:  - The target `resource_type` must be `bare_metal_server_network_attachment`.  - The target `interface_type` must not be `hipersocket`.
- `mount_path` - (String) The mount path for the share. The server component of the mount path may be either an IP address or a fully qualified domain name.

    This property will be absent if the lifecycle_state of the mount target is 'pending', failed, or deleting.

    -> **If the share's access_control_mode is:**
    &#x2022; security_group: The IP address used in the mount path is the primary_ip address of the virtual network interface for this share mount target. </br>
    &#x2022; vpc: The fully-qualified domain name used in the mount path is an address that resolves to the share mount target. </br>
- `created_at` - (String)The date and time that the share target was created.
- `href` - (String)The URL for this share target.
- `id` - (String)The unique identifier of the ShareTarget. The id is composed of \<ibm_is_share_id\>/\<ibm_is_share_mount_target_id\>
- `lifecycle_state` - (String)The lifecycle state of the mount target.
- `resource_type` - (String) The type of resource referenced.
- `transit_encryption` - (String) The transit encryption mode for this share target.

## Import

The `ibm_is_share_mount_target` can be imported using ID.

**Syntax**

```
$ terraform import ibm_is_share_mount_target.example `\<ibm_is_share_id\>/\<ibm_is_share_mount_target_id\>`
```

**Example**

```
$ terraform import ibm_is_share_mount_target.example d7bec597-4726-451f-8a63-e62e6f19c32c/d7bec597-4726-451f-8a63-e62e6f19c32c
```