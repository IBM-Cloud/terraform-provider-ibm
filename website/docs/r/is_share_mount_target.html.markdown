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
  share = ibm_is_share.example.id
  vpc = ibm_is_vpc.example.id
  name = "my-share-target"
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
  share = ibm_is_share.example1.id
  virtual_network_interface {
    primary_ip {
      name = "my-example-pip"
    }
    subnet = ibm_is_subnet.example1.id
    name = "my-example-vni"
  }
  name  = "my-example-mount-target"
}

//Create a mount target with subnet id
resource "ibm_is_share_mount_target" "example2" {
  share = ibm_is_share.example.id
  virtual_network_interface {
    subnet = ibm_is_subnet.example.id
    name = "my-example-vni"
  }
  name  = "my-example-mount-target"
}

//Create mount target with reserved ip id
resource "ibm_is_subnet_reserved_ip" "example" {
  subnet = ibm_is_subnet.example.id
  name = "my-example-resip"
}
resource "ibm_is_share_mount_target" "example" {
  share = ibm_is_share.example.id
  virtual_network_interface {
    primary_ip {
      reserved_ip = ibm_is_subnet_reserved_ip.example.reserved_ip
    }
    name = "my-example-vni"
  }
  name  = "my-example-mount-target"
}
```
## Argument Reference

The following arguments are supported:

- `share` - (Required, String) The file share identifier.
- `virtual_network_interface` (Optional, List) The virtual network interface for this share mount target. Required if the share's `access_control_mode` is `security_group`.
  - `name` - (Required, String) Name for this virtual network interface.
  Nested scheme for `virtual_network_interface`:
  - `primary_ip` - (Optional, List) The primary IP address to bind to the virtual network interface. May be either a reserved IP identity, or a reserved IP prototype object which will be used to create a new reserved IP.

      Nested scheme for `primary_ip`:
      - `auto_delete` - (Optional, Bool) Indicates whether this reserved IP member will be automatically deleted when either target is deleted, or the reserved IP is unbound. Defaults to `true`
      - `address` - (Optional, Forces new resource, String) The IP address to reserve. If unspecified, an available address on the subnet will automatically be selected.
      - `name`- (Optional, String) The name for this reserved IP. The name must not be used by another reserved IP in the subnet. Names starting with ibm- are reserved for provider-owned resources, and are not allowed.
      - `reserved_ip`- (Optional, String) The unique identifier for this reserved IP
  - `resource_group` - (Optional, String) The ID of the resource group to use.
  - `security_groups`- (Optional, List of string) The security groups to use for this virtual network interface.
  - `subnet` - (Optional, string) The associated subnet.
    
    ~> **Note**
    Within `primary_ip`, `reserved_ip` is mutually exclusive to  `auto_delete`, `address` and `name`

- `vpc` - (Optional, string) The VPC in which instances can mount the file share using this share target. Required if the share's `access_control_mode` is vpc.
  ~> **Note**
  `virtual_network_interface` and `vpc` are mutually exclusive and one of them must be provided.
- `name` - (Required, String) The user-defined name for this share target. Names must be unique within the share the share target resides in. If unspecified, the name will be a hyphenated list of randomly-selected words.
- `transit_encryption` - (Optional, String) The transit encryption mode for this share target. Supported values are **none**, **user_managed**. Default is **none**

~> **Note**
  `transit_encryption` can only be provided to create mount target for a share with `access_control_mode` `security_group`. It is not supported with shares that has `access_control_mode` `vpc`

## Attribute Reference

The following attributes are exported:

- `access_control_mode` - (String) The access control mode for the share.
- `mount_target` - (String) The unique identifier of the share target
- `created_at` - (String) The date and time that the share target was created.
- `href` - (String) The URL for this share target.
- `id` - (String) The unique identifier of the ShareTarget. The id is composed of \<ibm_is_share_id\>/\<ibm_is_share_mount_target_id\>
- `lifecycle_state` - (String) The lifecycle state of the mount target.
- `mount_path` - (String) The mount path for the share. The server component of the mount path may be either an IP address or a fully qualified domain name.

    This property will be absent if the lifecycle_state of the mount target is 'pending', failed, or deleting.

    -> **If the share's access_control_mode is:**
    &#x2022; security_group: The IP address used in the mount path is the primary_ip address of the virtual network interface for this share mount target. </br>
    &#x2022; vpc: The fully-qualified domain name used in the mount path is an address that resolves to the share mount target. </br>
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