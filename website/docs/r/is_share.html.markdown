---
layout: "ibm"
page_title: "IBM : is_share"
description: |-
  Manages Share.
subcategory: "VPC infrastructure"
---

# ibm\_is_share

Provides a resource for Share. This allows Share to be created, updated and deleted. For more information, about share replication, see [Share replication](https://cloud.ibm.com/docs/vpc?topic=vpc-file-storage-replication).

~> **NOTE**
  Regional file share `rfs` profile is available for accounts that have been granted special approval to preview the feature.

## Example Usage

```terraform
resource "ibm_is_share" "example" {
  access_control_mode = "security_group"
  name    = "my-share"
  size    = 200
  profile = "dp2"
  zone    = "us-south-2"
}
```
## Example Usage (Create a replica share)

```terraform
resource "ibm_is_share" "example-1" {
  zone                  = "us-south-3"
  source_share          = ibm_is_share.example.id
  name                  = "my-replica1"
  profile               = "dp2"
  replication_cron_spec = "0 */5 * * *"
}
```
## Example Usage (Create a file share with inline replica share)

```terraform
resource "ibm_is_share" "example-2" {
  zone    = "us-south-1"
  size    = 220
  name    = "my-share"
  profile = "dp2"
  replica_share {
    name                  = "my-replica"
    replication_cron_spec = "0 */5 * * *"
    profile               = "dp2"
    zone                  = "us-south-3"
  }
}
```
## Example Usage (Create a file share with inline mount target with a VNI)

```terraform
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
resource "ibm_is_share" "example-3" {
  zone    = "us-south-1"
  size    = 220
  name    = "my-share-1"
  profile = "dp2"
  mount_targets {
    name = "my-mount-target"
    virtual_network_interface {
      id = ibm_is_virtual_network_interface.example.id
    }
  }
}
```

## Example Usage (Create a cross regional replication)
```terraform
resource "ibm_is_share" "example-3" {
  provider = ibm.syd
  access_control_mode = "security_group"
  name    = "my-share"
  size    = 200
  profile = "dp2"
  zone    = "au-syd-2"
}
resource "ibm_is_share" "example-4" {
  provider = ibm.ussouth
  zone                  = "us-south-3"
  source_share_crn      = ibm_is_share.example-3.crn
  name                  = "my-replica1"
  profile               = "dp2"
  replication_cron_spec = "0 */5 * * *"
}
```

## Example share (Create accessor share for an origin share)
```terraform
resource "ibm_is_share" "example-4" {
  allowed_transit_encryption_modes = ["ipsec", "none"]
  access_control_mode = "security_group"
  name    = "my-share"
  size    = 200
  profile = "dp2"
  zone    = "au-syd-2"
}
resource "ibm_is_share" "example-5" {
  origin_share {
    crn = ibm_is_share.example-4.crn
  }     
  name                  = "my-replica1"
}
resource "ibm_is_share" "example-6" {
  origin_share {
    id = ibm_is_share.example-4.id
  }     
  name                  = "my-replica1"
}
```
## Example share (Create a regional file share)
```terraform
resource "ibm_is_share" "example-4" {
  allowed_transit_encryption_modes = ["stunnel", "none"]
  access_control_mode = "security_group"
  bandwidth = 210
  name    = "my-share"
  size    = 200
  profile = "rfs"
}

## Argument Reference

The following arguments are supported:

- `allowed_access_protocols` - (Optional, List) List of allowed access protocols for the share. Supported values are **nfs4** 
- `access_control_mode` - (Optional, Boolean) The access control mode for the share. Supported values are **security_group** and **vpc**. Default value is **security_group**
- `allowed_transit_encryption_modes` - (Optional, List of string) The transit encryption modes allowed for this share.
- `access_tags`  - (Optional, List of Strings) The list of access management tags to attach to the share. **Note** For more information, about creating access tags, see [working with tags](https://cloud.ibm.com/docs/account?topic=account-tag).
- `bandwidth` - (Optional, Integer) The bandwidth for the file share
- `encryption_key` - (Optional, String) The CRN of the [Key Protect Root Key](https://cloud.ibm.com/docs/key-protect?topic=key-protect-getting-started-tutorial) or [Hyper Protect Crypto Service Root Key](https://cloud.ibm.com/docs/hs-crypto?topic=hs-crypto-get-started) for this resource.
- `initial_owner` - (Optional, List) The initial owner for the file share.

  Nested scheme for `initial_owner`:
  - `gid` - (Optional, Integer) The initial group identifier for the file share.
  - `uid` - (Optional, Integer) The initial user identifier for the file share.
- `iops` - (Optional, Integer) The maximum input/output operation performance bandwidth per second for the file share. For more information about the iops range for the given size, refer [File Storage for VPC profiles](https://cloud.ibm.com/docs/vpc?topic=vpc-file-storage-profiles&interface=ui)
- `mount_targets` - (Optional, List) Share targets for the file share.
  - `name` - (Required, string) The user-defined name for this share target. Names must be unique within the share the share target resides in.
  - `virtual_network_interface` (Optional, List) The virtual network interface for this share mount target. Required if the share's `access_control_mode` is `security_group`.

    Nested scheme for `virtual_network_interface`:
    - `name` - (Required, String) Name for this virtual network interface. The name must not be used by another virtual network interface in the VPC.
    - `id` - (Optional) The ID for virtual network interface. Mutually exclusive with other `virtual_network_interface` arguments.
    
    ~> **Note**
    `id` is mutually exclusive with other `virtual_network_interface` prototype arguments
    - `allow_ip_spoofing` - (Optional, Bool) Indicates whether source IP spoofing is allowed on this interface. If false, source IP spoofing is prevented on this interface. If true, source IP spoofing is allowed on this interface.
    - `auto_delete` - (Optional, Bool) Indicates whether this virtual network interface will be automatically deleted when target is deleted
    - `enable_infrastructure_nat` - (Optional, Bool) If `true`:- The VPC infrastructure performs any needed NAT operations.- `floating_ips` must not have more than one floating IP.If `false`:- Packets are passed unchanged to/from the network interface,  allowing the workload to perform any needed NAT operations.- `allow_ip_spoofing` must be `false`.- If the virtual network interface is attached:  - The target `resource_type` must be `bare_metal_server_network_attachment`.  - The target `interface_type` must not be `hipersocket`.
    - `primary_ip` - (Optional, List) The primary IP address to bind to the virtual network interface. May be either a reserved IP identity, or a reserved IP prototype object which will be used to create a new reserved IP.

        Nested scheme for `primary_ip`:
        - `auto_delete` - (Optional, Bool) Indicates whether this reserved IP member will be automatically deleted when either target is deleted, or the reserved IP is unbound. Defaults to `true`
        - `address` - (Optional, Forces new resource, String) The IP address to reserve. If unspecified, an available address on the subnet will automatically be selected.
        - `name`- (Optional, String) The name for this reserved IP. The name must not be used by another reserved IP in the subnet. Names starting with ibm- are reserved for provider-owned resources, and are not allowed.
        - `reserved_ip`- (Optional, String) The unique identifier for this reserved IP
    - `protocol_state_filtering_mode` - (Optional, String) The protocol state filtering mode to use for this virtual network interface. 

        ~> **If auto, protocol state packet filtering is enabled or disabled based on the virtual network interface's target resource type:** 
            **&#x2022;** bare_metal_server_network_attachment: disabled </br>
            **&#x2022;** instance_network_attachment: enabled </br>
            **&#x2022;** share_mount_target: enabled </br>
    - `resource_group` - (Optional, String) The ID of the resource group to use.
    - `security_groups`- (Optional, List of string) The security groups to use for this virtual network interface.
    - `subnet` - (Optional, string) The associated subnet.
    
      ~> **Note**
      Within `primary_ip`, `reserved_ip` is mutually exclusive to  `auto_delete`, `address` and `name`

  - `vpc` - (Optional, string) The VPC in which instances can mount the file share using this share target. Required if the share's `access_control_mode` is vpc.
  - `transit_encryption` - (Optional, String) The transit encryption mode for this share target. Supported values are **none**, **user_managed**. Default is **none**

~> **Note**
  `transit_encryption` can only be provided to create mount target for a share with `access_control_mode` `security_group`. It is not supported with shares that has `access_control_mode` `vpc`
  ~> **Note**
    `virtual_network_interface` and `vpc` are mutually exclusive and one of them must be provided.
- `name` - (Required, string) The unique user-defined name for this file share. If unspecified, the name will be a hyphenated list of randomly-selected words.
- `origin_share` - (Optional, List) The origin share this accessor share is referring to.
  Nested schema for **origin_share**:
	- `crn` - (Optional, String) The CRN for this file share.
	- `id` - (Optional, String) The unique identifier for this file share.
- `profile` - (Required, string) The globally unique name for this share profile.

  ~> **NOTE** 
  While updating `profile` from 'custom' to a tiered profile make sure to remove `iops` from the configuration.
  
- `replica_share` - (Optional, List) Configuration for a replica file share to create and associate with this file share.
  - `access_control_mode` - (Optional, Boolean) The access control mode for the share. Supported values are **security_group** and **vpc**. Default value is **vpc**
  - `access_tags`  - (Optional, List of Strings) The list of access management tags to attach to the share. **Note** For more information, about creating access tags, see [working with tags](https://cloud.ibm.com/docs/account?topic=account-tag).
  - `iops` - (Optional, Int)
  - `mount_targets` - (List) List of mount targets
    - `name` - (Optional, String)
    - `virtual_network_interface` (Optional, List) The virtual network interface for this share mount target. Required if the share's `access_control_mode` is `security_group`.
      Nested scheme for `virtual_network_interface`:
      - `name` - (Required, String) Name for this virtual network interface. The name must not be used by another virtual network interface in the VPC.
      - `id` - (Optional) The ID for virtual network interface. Mutually exclusive with other `virtual_network_interface` arguments.
      
      ~> **Note**
        `id` is mutually exclusive with other `virtual_network_interface` prototype arguments
      - `allow_ip_spoofing` - (Optional, Bool) Indicates whether source IP spoofing is allowed on this interface. If false, source IP spoofing is prevented on this interface. If true, source IP spoofing is allowed on this interface.
      - `auto_delete` - (Optional, Bool) Indicates whether this virtual network interface will be automatically deleted when target is deleted
      - `enable_infrastructure_nat` - (Optional, Bool) If `true`:- The VPC infrastructure performs any needed NAT operations.- `floating_ips` must not have more than one floating IP.If `false`:- Packets are passed unchanged to/from the network interface,  allowing the workload to perform any needed NAT operations.- `allow_ip_spoofing` must be `false`.- If the virtual network interface is attached:  - The target `resource_type` must be `bare_metal_server_network_attachment`.  - The target `interface_type` must not be `hipersocket`.
      - `primary_ip` - (Optional, List) The primary IP address to bind to the virtual network interface. May be either a reserved IP identity, or a reserved IP prototype object which will be used to create a new reserved IP.
        Nested scheme for `primary_ip`:
        - `auto_delete` - (Optional, Bool) Indicates whether this reserved IP member will be automatically deleted when either target is deleted, or the reserved IP is unbound. Defaults to `true`
        - `address` - (Optional, Forces new resource, String) The IP address to reserve. If unspecified, an available address on the subnet will automatically be selected.
        - `name`- (Optional, String) The name for this reserved IP. The name must not be used by another reserved IP in the subnet. Names starting with ibm- are reserved for provider-owned resources, and are not allowed.
        - `reserved_ip`- (Optional, String) The unique identifier for this reserved IP
      - `protocol_state_filtering_mode` - (Optional, String) The protocol state filtering mode to use for this virtual network interface. 

        ~> **If auto, protocol state packet filtering is enabled or disabled based on the virtual network interface's target resource type:** 
            **&#x2022;** bare_metal_server_network_attachment: disabled </br>
            **&#x2022;** instance_network_attachment: enabled </br>
            **&#x2022;** share_mount_target: enabled </br>
      - `resource_group` - (Optional, String) The ID of the resource group to use.
      - `security_groups`- (Optional, List of string) The security groups to use for this virtual network interface.
      - `subnet` - (Optional, string) The associated subnet.
        ~> **Note**
        Within `primary_ip`, `reserved_ip` is mutually exclusive to  `auto_delete`, `address` and `name`
    - `vpc` - (Optional, string) The VPC in which instances can mount the file share using this share target. Required if the share's `access_control_mode` is vpc.
    ~> **Note**
    `virtual_network_interface` and `vpc` are mutually exclusive and one of them must be provided.

  - `name` - (Optional, String)
  - `profile` - (Optional, String)
  - `replication_cron_spec` - (Optional, String)
  - `tags`  - (Optional, List of Strings) The list of user tags to attach to the share.
  - `zone` - (Required, String)
- `resource_group` - (Optional, String) The unique identifier for this resource group.
- `replication_cron_spec` - (Optional, String) The cron specification for the file share replication schedule.
- `size` - (Required, Integer) The size of the file share rounded up to the next gigabyte.
- `source_share` - (Optional, String) The ID of the source file share for this replica file share. The specified file share must not already have a replica, and must not be a replica.
- `source_share_crn` - (Optional, String) The CRN of the source file share. 
- `source_snapshot` - (List) The snapshot from which this share was cloned.This property will be present when the share was created from a snapshot.The resources supported by this property may[expand](https://cloud.ibm.com/apidocs/vpc#property-value-expansion) in thefuture.
  Nested schema for **source_snapshot**:
	- `crn` - (Optional, String) The CRN for this share snapshot.
	- `id` - (Optional, String) The unique identifier for this share snapshot.

  ~> **NOTE** 
  `crn` and `id` are mutually exclusive for source_snapshot

  ~> **NOTE** 
  `source_snapshot` is mutually exclusive with `zone`, `source_share_crn`, `source_share`, `origin_share`

- `tags`  - (Optional, List of Strings) The list of user tags to attach to the share.
- `zone` - (Required, string) The globally unique name for this zone.

## Attribute Reference

The following attributes are exported:

- `access_control_mode` - (Boolean) The access control mode for the share.
- `access_tags`  - (String) Access management tags associated to the share.
- `accessor_binding_role` - (String) The accessor binding role of this file share:- `none`: This file share is not participating in access with another file share- `origin`: This file share is the origin for one or more file shares  (which may be in other accounts)- `accessor`: This file share is providing access to another file share  (which may be in another account).
- `accessor_bindings` - (List) The accessor bindings for this file share. Each accessor binding identifies a resource (possibly in another account) with access to this file share's data.
  Nested schema for **accessor_bindings**:
	- `href` - (String) The URL for this share accessor binding.
	- `id` - (String) The unique identifier for this share accessor binding.
	- `resource_type` - (String) The resource type.
- `allowed_transit_encryption_modes` - (List of string) The transit encryption modes allowed for this share.
- `created_at` - (String) The date and time that the file share is created.
- `crn` - (String) The CRN for this share.
- `encryption` - (String) The type of encryption used for this file share.
- `encryption_key` - (String) The CRN of the [Key Protect Root Key](https://cloud.ibm.com/docs/key-protect?topic=key-protect-getting-started-tutorial) or [Hyper Protect Crypto Service Root Key](https://cloud.ibm.com/docs/hs-crypto?topic=hs-crypto-get-started) for this resource.
- `href` - (String) The URL for this share.
- `id` - (String) The unique identifier of the Share.
- `iops` - (Integer) The maximum input/output operation performance bandwidth per second for the file share.
- `latest_sync` - (List) Information about the latest synchronization for this file share.
Nested `latest_sync` blocks have the following structure:
  - `completed_at` - (String) The completed date and time of last synchronization between the replica share and its source.
  - `data_transferred` - (Integer) The data transferred (in bytes) in the last synchronization between the replica and its source.
  - `started_at` - (String) The start date and time of last synchronization between the replica share and its source.
- `latest_job` - (List) The latest job associated with this file share.This property will be absent if no jobs have been created for this file share. Nested `latest_job` blocks have the following structure:
  - `status` - (String) The status of the file share job
  - `status_reasons` - (List) The reasons for the file share job status (if any). Nested `status_reasons` blocks have the following structure:
    - `code` - (String) A snake case string succinctly identifying the status reason.
    - `message` - (String) An explanation of the status reason.
    - `more_info` - (String) Link to documentation about this status reason.
  - `type` - (String) The type of the file share job
- `lifecycle_state` - (String) The lifecycle state of the file share.
- `mount_targets` - (List) Mount targets for the file share. Nested `mount_targets` blocks have the following structure:
  - `name` - (String) The name for this share. The name is unique across all shares in the region.
  - `href` - (String) Href of this mount target.
  - `id` - (String) Unique ID of this mount target.
	- `virtual_network_interface` (List) The virtual network interface for this share mount target. Required if the share's `access_control_mode` is security_group.
    Nested scheme for `virtual_network_interface`:
    - `crn` - (String) CRN of virtual network interface
    - `name` - (String) The name for this virtual network interface.
    - `href` - (String) Href of this virtual network interface.
    - `id` - (String) Unique ID of this virtual network interface.
    - `primary_ip` - (List) The primary IP address to bind to the virtual network interface. May be either a reserved IP identity, or a reserved IP prototype object which will be used to create a new reserved IP.

        Nested scheme for `primary_ip`:
        - `auto_delete` - (Boolean) Indicates whether this reserved IP member will be automatically deleted when either target is deleted, or the reserved IP is unbound. Defaults to `true`
        - `address` - (String) The IP address to reserve. If unspecified, an available address on the subnet will automatically be selected.
        - `href` - (String) Href of this primary ip.
        - `name`- (String) The name for this reserved IP. The name must not be used by another reserved IP in the subnet. Names starting with ibm- are reserved for provider-owned resources, and are not allowed.
        - `reserved_ip`- (String) The unique identifier for this reserved IP
        - `resource_type` - (String) Resource type of primary ip
    - `resource_group` - (String) The ID of the resource group to use.
    - `resource_type` - (String) Resource type of this virtual network interface.
    - `security_groups`- (List of string) The security groups to use for this virtual network interface.
    - `subnet` - (string) The associated subnet.
- `origin_share` - (Optional, List) The origin share this accessor share is referring to.This property will be present when the `accessor_binding_role` is `accessor`.
  Nested schema for **origin_share**:
	- `crn` - (Computed, String) The CRN for this file share.
	- `deleted` - (Optional, List) If present, this property indicates the referenced resource has been deleted, and providessome supplementary information.
	  Nested schema for **deleted**:
		- `more_info` - (Computed, String) Link to documentation about deleted resources.
	- `href` - (Computed, String) The URL for this file share.
	- `id` - (Computed, String) The unique identifier for this file share.
	- `name` - (Computed, String) The name for this share. The name is unique across all shares in the region.
	- `remote` - (Optional, List) If present, this property indicates that the resource associated with this referenceis remote and therefore may not be directly retrievable.
	  Nested schema for **remote**:
		- `account` - (Optional, List) If present, this property indicates that the referenced resource is remote to thisaccount, and identifies the owning account.
		  Nested schema for **account**:
			- `id` - (Computed, String) The unique identifier for this account.
			- `resource_type` - (Computed, String) The resource type.
		- `region` - (Optional, List) If present, this property indicates that the referenced resource is remote to thisregion, and identifies the native region.
		  Nested schema for **region**:
			- `href` - (Computed, String) The URL for this region.
			- `name` - (Computed, String) The globally unique name for this region.
	- `resource_type` - (Computed, String) The resource type.
- `resource_type` - (String) The type of resource referenced.
- `replica_share` - (List) Configuration for a replica file share to create and associate with this file share.
  - `crn` - (String) CRN of replica share
  - `href` - (String) Href of replica share
  - `id` - (String) ID of replica share
  - `iops` - (Integer)
  - `name` - (String)
  - `profile` - (String)
  - `replication_cron_spec` - (String)
  - `mount_targets` - (List) List of mount targets
    - `name` - (String) Name of the mount target
    - `href` - (String) Href of this mount target.
    - `id` - (String) Unique ID of this mount target.
    - `virtual_network_interface` (List) The virtual network interface for this share mount target. Required if the share's `access_control_mode` is `security_group`.
      Nested scheme for `virtual_network_interface`:
      - `crn` - (String) CRN of virtual network interface
      - `name` - (String) Name for this virtual network interface.
      - `href` - (String) Href of this mount target.
      - `id` - (String) Unique ID of this mount target.
      - `primary_ip` - (List) The primary IP address to bind to the virtual network interface. May be either a reserved IP identity, or a reserved IP prototype object which will be used to create a new reserved IP.

          Nested scheme for `primary_ip`:
          - `auto_delete` - (Boolean) Indicates whether this reserved IP member will be automatically deleted when either target is deleted, or the reserved IP is unbound. Defaults to `true`
          - `address` - (String) The IP address to reserve. If unspecified, an available address on the subnet will automatically be selected.
          - `name`- (String) The name for this reserved IP. The name must not be used by another reserved IP in the subnet. Names starting with ibm- are reserved for provider-owned resources, and are not allowed.
          - `reserved_ip`- (String) The unique identifier for this reserved IP
      - `resource_group` - (String) The ID of the resource group to use.
      - `security_groups`- (List of string) The security groups to use for this virtual network interface.
      - `subnet` - (String) The associated subnet.
    - `vpc` - (String) The VPC in which instances can mount the file share using this share target. Required if the share's `access_control_mode` is vpc.
  - `zone` - (String) The zone this replica file share will reside in.
- `resource_group` - The unique identifier of the resource group of this share.
- `replication_role`  - The replication role of the file share.
- `replication_status` - "The replication status of the file share.
- `replication_status_reasons` - The reasons for the current replication status (if any). 
  Nested `replication_status_reasons` blocks have the following structure:
  - `code` - A snake case string succinctly identifying the status reason.
  - `message` - An explanation of the status reason.
  - `more_info` - Link to documentation about this status reason.
- `source_snapshot` - (List) The snapshot from which this share was cloned.This property will be present when the share was created from a snapshot.The resources supported by this property may[expand](https://cloud.ibm.com/apidocs/vpc#property-value-expansion) in thefuture.
  Nested schema for **source_snapshot**:
	- `crn` - (String) The CRN for this share snapshot.
	- `deleted` - ( List) If present, this property indicates the referenced resource has been deleted, and providessome supplementary information.
	  Nested schema for **deleted**:
		- `more_info` - (Computed, String) Link to documentation about deleted resources.
	- `href` - (String) The URL for this share snapshot.
	- `id` - (String) The unique identifier for this share snapshot.
	- `name` - (String) The name for this share snapshot. The name is unique across all snapshots for the file share.
	- `resource_type` - (String) The resource type.
- `snapshot_count` - (Integer) The total number of snapshots for this share.
- `snapshot_size` - (Integer) The total size (in gigabytes) of snapshots used for this file share.
- `tags`  - (String) User tags associated for to the share.


## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import the `ibm_is_share` resource by using `id`.
The `id` property can be formed using the share id. For example:

```terraform
import {
  to = ibm_is_share.example
  id = "<id>"
}
```

Using `terraform import`. For example:

```console
% terraform import ibm_is_share.example <id>
```