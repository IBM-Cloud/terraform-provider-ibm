---
layout: "ibm"
page_title: "IBM : is_share_target"
description: |-
  Manages ShareTarget.
subcategory: "VPC infrastructure"
---


# ibm\_is_share_target

Provides a resource for ShareTarget. This allows ShareTarget to be created, updated and deleted.


~> **NOTE**
IBM Cloud® File Storage for VPC is available for customers with special approval. Contact your IBM Sales representative if you are interested in getting access.

~> **NOTE**
This is a Beta feature and it is subject to change in the GA release 

~> **NOTE**
This resource is being deprecated. Please use `ibm_is_share_mount_target` instead

## Example Usage

```hcl
resource "ibm_is_vpc" "vpc" {
  name = "my-vpc"
}

resource "ibm_is_share" "is_share" {
  name = "my-share"
  size = 200
  profile = "dp2"
  zone = "us-south-2"
}

resource "ibm_is_share_mount_target" "example" {
  share = ibm_is_share.example.id
  vpc = ibm_is_vpc.vpc.id
  name = "my-share-target"
}

//Example to create share with access_control_mode security_group and mount target with virtual network interface

resource "ibm_is_share" "example1" {
  name = "my-share"
  access_control_mode = "security_group"
  size = 200
  profile = "dp2"
  zone = "us-south-2"
}

// Example with virtual network interface with reserved ip prototype 
resource "ibm_is_share_mount_target" "example1" {
  share = ibm_is_share.example1.id
  virtual_network_interface {
    name = "my-virtual_network_interface"
    primary_ip {
      address = "10.240.64.5"
      auto_delete = true
      name = "my-reserved-ip"
    }
    name = "my-share-target"
  }
}

// Example with virtual network interface with existing reserved ip 
resource "ibm_is_subnet_reserved_ip" "example" {
  subnet      = ibm_is_subnet.example.id
  name        = "example-subnet-reserved-ip"
  auto_delete = true
}
resource "ibm_is_share_mount_target" "example2" {
  share = ibm_is_share.example1.id
  virtual_network_interface {
    name = "my-virtual_network_interface"
    primary_ip {
      reserved_ip = ibm_is_subnet_reserved_ip.example.id
    }
    name = "my-share-target"
  }
}

// Example with virtual network interface with subnet
resource "ibm_is_share_mount_target" "example3" {
  share = ibm_is_share.example1.id
  virtual_network_interface {
    name = "my-virtual_network_interface"
    subnet = ibm_is_subnet.example.id
    name = "my-share-target"
  }
}
```

## Argument Reference

The following arguments are supported:

- `share` - (Required, Forces new resource, String) The file share identifier.
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
  
- `name` - (Required, String) The user-defined name for this share target. Names must be unique within the share the share target resides in.

## Attribute Reference

The following attributes are exported:

- `access_control_mode` (String) The access control mode for the share.
- `created_at` - The date and time that the share target was created.
- `href` - The URL for this share target.
- `lifecycle_state` - The lifecycle state of the mount target.
- `id` - The unique identifier of the ShareTarget. The id is composed of \<ibm_is_share_id\>/\<ibm_is_share_target_id\>
- `mount_path` - The mount path for the share.The IP addresses used in the mount path are currently within the IBM services IP range, but are expected to change to be within one of the VPC's subnets in the future.
- `virtual_network_interface` (List) The virtual network interface for this share mount target. Required if the share's `access_control_mode` is security_group.
  - `href` - (String) Href of this virtual network interface.
  - `id` - (String) Unique ID of this virtual network interface.
  Nested scheme for `virtual_network_interface`:
  - `primary_ip` - (List) The primary IP address to bind to the virtual network interface. May be either a reserved IP identity, or a reserved IP prototype object which will be used to create a new reserved IP.

      Nested scheme for `primary_ip`:
      - `auto_delete` - (Bool) Indicates whether this reserved IP member will be automatically deleted when either target is deleted, or the reserved IP is unbound. Defaults to `true`
      - `address` - (String) The IP address to reserve. If unspecified, an available address on the subnet will automatically be selected.
      - `href` - (String) Href of this primary ip.
      - `name`- (String) The name for this reserved IP. The name must not be used by another reserved IP in the subnet. Names starting with ibm- are reserved for provider-owned resources, and are not allowed.
      - `reserved_ip`- (String) The unique identifier for this reserved IP
      - `resource_type` - (String) Resource type of primary ip
  - `resource_group` - (String) The ID of the resource group to use.
  - `resource_type` - (String) Resource type of this virtual network interface.
  - `security_groups`- (List of string) The security groups to use for this virtual network interface.
  - `subnet` - (string) The associated subnet.
- `share_target` - (String) The unique identifier of the share target
- `resource_type` - (String) The type of resource referenced.
- `vpc` - (String) Unique ID of the VPC
  ~> **Note**
  If access_control_mode is:
  `security_group`: ID of the VPC for the virtual network interface of this share mount target
  `vpc`: ID of the VPC in which instances can mount the file share using this share mount target



## Import

The `ibm_is_share_target` can be imported using ID.

**Syntax**

```
$ terraform import ibm_is_share_mount_target.example `\<ibm_is_share_id\>/\<ibm_is_share_target_id\>`
```

**Example**

```
$ terraform import ibm_is_share_mount_target.example d7bec597-4726-451f-8a63-e62e6f19c32c/d7bec597-4726-451f-8a63-e62e6f19c32c
```