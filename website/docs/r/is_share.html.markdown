---
layout: "ibm"
page_title: "IBM : is_share"
description: |-
  Manages Share.
subcategory: "VPC infrastructure"
---

# ibm\_is_share

Provides a resource for Share. This allows Share to be created, updated and deleted.


~> **NOTE**
IBM Cloud® File Storage for VPC is available for customers with special approval. Contact your IBM Sales representative if you are interested in getting access.

~> **NOTE**
This is a Beta feature and it is subject to change in the GA release 


## Example Usage

```terraform
resource "ibm_is_share" "example" {
  name    = "my-share"
  size    = 200
  profile = "tier-3iops"
  zone    = "us-south-2"
}
```
## Example Usage (Create a replica share)

```terraform
resource "ibm_is_share" "example-1" {
  zone                  = "us-south-3"
  source_share          = ibm_is_share.example.id
  name                  = "my-replica1"
  profile               = "tier-3iops"
  replication_cron_spec = "0 */5 * * *"
}
```
## Example Usage (Create a file share with inline replica share)

```terraform
resource "ibm_is_share" "example-2" {
  zone    = "us-south-1"
  size    = 220
  name    = "my-share"
  profile = "tier-3iops"
  replica_share {
    name                  = "my-replica"
    replication_cron_spec = "0 */5 * * *"
    profile               = "tier-3iops"
    zone                  = "us-south-3"
  }
}
```
## Argument Reference

The following arguments are supported:

- `access_tags`  - (Optional, List of Strings) The list of access management tags to attach to the share. **Note** For more information, about creating access tags, see [working with tags](https://cloud.ibm.com/docs/account?topic=account-tag).
- `encryption_key` - (Optional, String) The CRN of the [Key Protect Root Key](https://cloud.ibm.com/docs/key-protect?topic=key-protect-getting-started-tutorial) or [Hyper Protect Crypto Service Root Key](https://cloud.ibm.com/docs/hs-crypto?topic=hs-crypto-get-started) for this resource.
- `initial_owner` - (Optional, List) The initial owner for the file share.

  Nested scheme for `initial_owner`:
  - `gid` - (Optional, Integer) The initial group identifier for the file share.
  - `uid` - (Optional, Integer) The initial user identifier for the file share.
- `iops` - (Optional, Integer) The maximum input/output operation performance bandwidth per second for the file share.
- `name` - (Required, String) The unique user-defined name for this file share. If unspecified, the name will be a hyphenated list of randomly-selected words.
- `profile` - (Required, String) The globally unique name for this share profile.

  ~> **NOTE** 
  While updating `profile` from 'custom' to a tiered profile make sure to remove `iops` from the configuration.
  
- `replica_share` - (Optional, List) Configuration for a replica file share to create and associate with this file share.
  - `iops` - (Optional, Integer) The maximum input/output operations per second (IOPS) for the file share. The share must be in the custom profile family, and the value must be in the range supported by the share's specified size.
  - `name` - (Required, String) The name for this share. The name must not be used by another share in the region.
  - `profile` - (Required, String) The profile to use for this file share.
  - `replication_cron_spec` - (Required, Forces new resource, String)
  - `mount_targets` - (Optional, List) Mount targets for the replica file share.
    - `name` - (Optional, String) The user-defined name for this share target. Names must be unique within the share the share target resides in.
    - `vpc` - (Required, String) The VPC in which instances can mount the file share using this share target.
  - `zone` - (Required, Forces new resource, String) The zone this replica file share will reside in. Must be a different zone in the same region as the source share.
  - `access_tags`  - (Optional, List of Strings) The list of access management tags to attach to the share. **Note** For more information, about creating access tags, see [working with tags](https://cloud.ibm.com/docs/account?topic=account-tag).
  - `tags`  - (Optional, List of Strings) The list of user tags to attach to the share.
- `resource_group` - (Optional, String) The unique identifier for this resource group. If unspecified, the account's [default resourcegroup](https://cloud.ibm.com/apidocs/resource-manager#introduction) is used.
- `size` - (Required, Integer) The size of the file share rounded up to the next gigabyte. 
- `mount_targets` - (Optional, List) Share targets for the file share.
  - `name` - (Optional, String) The user-defined name for this share target. Names must be unique within the share the share target resides in.
  - `vpc` - (Required, String) The VPC in which instances can mount the file share using this share target.
- `source_share` - (Optional, String) The ID of the source file share for this replica file share. The specified file share must not already have a replica, and must not be a replica.
- `replication_cron_spec` - (Optional, String) The cron specification for the file share replication schedule.
- `tags`  - (Optional, List of Strings) The list of user tags to attach to the share.
- `zone` - (Required, String) The globally unique name for this zone.

## Attribute Reference

The following attributes are exported:

- `access_tags`  - (String) Access management tags associated to the share.
- `created_at` - (String) The date and time that the file share is created.
- `crn` - (String) The CRN for this share.
- `encryption` - (String) The type of encryption used for this file share.
- `href` - (String) The URL for this share.
- `id` - (String) The unique identifier of the Share.
- `iops` - (Integer) The maximum input/output operation performance bandwidth per second for the file share.
- `latest_job` - (List) The latest job associated with this file share.This property will be absent if no jobs have been created for this file share. Nested `latest_job` blocks have the following structure:
  - `status` - (String) The status of the file share job
  - `status_reasons` - (List) The reasons for the file share job status (if any). Nested `status_reasons` blocks have the following structure:
    - `code` - (String) A snake case string succinctly identifying the status reason.
    - `message` - (String) An explanation of the status reason.
    - `more_info` - (String) Link to documentation about this status reason.
  - `type` - (String) The type of the file share job
- `lifecycle_state` - (String) The lifecycle state of the file share.
- `replica_share` - (List) Configuration for a replica file share to create and associate with this file share.
  - `crn` - (String) The CRN for this replica share.
  - `href` - (String) The href for this replica share.
  - `id` - (String) The id for this replica share.
  - `iops` - (Integer) The maximum input/output operations per second (IOPS) for the file share. The share must be in the custom profile family, and the value must be in the range supported by the share's specified size.
  - `name` - (String) The name for this share. The name must not be used by another share in the region.
  - `profile` - (String) The profile to use for this file share.
  - `replication_cron_spec` - (String) The cron specification for the file share replication schedule.
  - `replication_role` - (String) The replication role of the file share.
  - `replication_status` - (String) The replication status of the file share.
  - `replication_status_reasons` - (List) The reasons for the current replication status.
    - `code` - (String) A snake case string succinctly identifying the status reason
    - `message` - (String) An explanation of the status reason
    - `more_info` - (String) Link to documentation about this status reason
  - `mount_targets` - (List) Mount targets for the replica file share.
    - `href` - (String) The href for this mount target.
    - `id` - (String) The id for this mount target.
    - `name` - (String) The user-defined name for this share target. Names must be unique within the share the share target resides in.
    - `vpc` - (String) The VPC in which instances can mount the file share using this share target.
    - `resource_type` - (String) Resource type of mount target
  - `zone` - (String) The zone this replica file share will reside in.
  - `access_tags`  - (List of Strings) The list of access management tags to attach to the share. **Note** For more information, about creating access tags, see [working with tags](https://cloud.ibm.com/docs/account?topic=account-tag).
  - `tags`  - (List of Strings) The list of user tags to attach to the share.
- `resource_type` - The type of resource referenced.
- `encryption_key` - The CRN of the [Key Protect Root Key](https://cloud.ibm.com/docs/key-protect?topic=key-protect-getting-started-tutorial) or [Hyper Protect Crypto Service Root Key](https://cloud.ibm.com/docs/hs-crypto?topic=hs-crypto-get-started) for this resource.

- `resource_group` - The unique identifier for this resource group. If unspecified, the account's [default resourcegroup](https://cloud.ibm.com/apidocs/resource-manager#introduction) is used.
- `replication_cron_spec` - (String) The cron specification for the file share replication schedule.
- `replication_role`  - (String) The replication role of the file share.
  
  -> **replication_role could be one of the below:**
   &#x2022; `none`: This share is not participating in replication. </br>
   &#x2022; `replica`: This share is a replication target. </br>
   &#x2022; `source`: This share is a replication source. </br>
  
- `replication_status` - (String) "The replication status of the file share.

  -> **replication_status could be one of the below:**
   &#x2022; `initializing`: This share is initializing replication. </br>
   &#x2022; `active`: This share is actively participating in replication. </br>
   &#x2022; `failover_pending`: This share is performing a replication failover. </br>
   &#x2022; `split_pending`: This share is performing a replication split. </br>
   &#x2022; `none`: This share is not participating in replication. </br>
   &#x2022; `degraded`: This share's replication sync is degraded. </br>
   &#x2022; `sync_pending`: This share is performing a replication sync. </br>
  
- `replication_status_reasons` - (List) The reasons for the current replication status (if any). Nested `replication_status_reasons` blocks have the following structure:
  - `code` - (String) A snake case string succinctly identifying the status reason.
  - `message` - (String) An explanation of the status reason.
  - `more_info` - (String) Link to documentation about this status reason.
- `mount_targets` - (List) Mount targets for the file share. Nested `mount_targets` blocks have the following structure:
	- `href` - (String) The href for this mount target.
  - `id` - (String) The id for this mount target.
  - `name` - (String) The user-defined name for this mount target.
  - `vpc` - (String) The VPC in which instances can mount the file share using this mount target.
  - `resource_type` - (String) Resource type of mount target
- `tags`  - (String) User tags associated for to the share.


## Import

The `ibm_is_share` can be imported using ID.

**Syntax**

```
$ terraform import ibm_is_share.example <id>
```

**Example**

```
$ terraform import ibm_is_share.example d7bec597-4726-451f-8a63-e62e6f19c32c
```
