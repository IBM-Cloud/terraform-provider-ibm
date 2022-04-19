---
layout: "ibm"
page_title: "IBM : is_share"
description: |-
  Manages Share.
subcategory: "Virtual Private Cloud API"
---

# ibm\_is_share

Provides a resource for Share. This allows Share to be created, updated and deleted.

## Example Usage

```terraform
resource "ibm_is_share" "example" {
  name = "my-share"
  size = 200
  profile = "tier-3iops"
  zone = "us-south-2"
}
```
## Example Usage (Create a replica share)

```terraform
resource "ibm_is_share" "example-1" {
    zone = "us-south-3"
    source_share = ibm_is_share.example.id
    name = "my-replica1"
    profile = "tier-3iops"
    replication_cron_spec = "0 */5 * * *"
}
```
## Example Usage (Create a source share with replica share)

```terraform
resource "ibm_is_share" "example-2" {
  zone = "us-south-1"
  size = 220
  name = "my-share"
  profile = "tier-3iops"
  replica_share {
    name = "my-replica" 
    replication_cron_spec = "0 */5 * * *"
    profile = "tier-3iops"
    zone = "us-south-3"
  }
}
```
## Argument Reference

The following arguments are supported:

- `encryption_key` - (Optional, String) The CRN of the [Key Protect Root Key](https://cloud.ibm.com/docs/key-protect?topic=key-protect-getting-started-tutorial) or [Hyper Protect Crypto Service Root Key](https://cloud.ibm.com/docs/hs-crypto?topic=hs-crypto-get-started) for this resource.
- `initial_owner_gid` - (Optional, int) The initial group identifier for the file share.
- `initial_owner_uid` - (Optional, int) The initial user identifier for the file share.
- `iops` - (Optional, int) The maximum input/output operation performance bandwidth per second for the file share.
- `name` - (Required, string) The unique user-defined name for this file share. If unspecified, the name will be a hyphenated list of randomly-selected words.
- `profile` - (Required, string) The globally unique name for this share profile.
- `replica_share` - (Optional, List) Configuration for a replica file share to create and associate with this file share.
  - `iops` - (Optional, Int)
  - `name` - (Optional, String)
  - `profile` - (Optional, String)
  - `replication_cron_spec` - (Optional, String)
  - `targets`
    - `name` - (Optional, String)
    - `subnet` - (Optional, String)
    - `vpc` - (Required, String) 
  - `zone` - (Required, String)
- `resource_group` - (Optional, string) The unique identifier for this resource group. If unspecified, the account's [default resourcegroup](https://cloud.ibm.com/apidocs/resource-manager#introduction) is used.
- `size` - (Required, int) The size of the file share rounded up to the next gigabyte.
- `share_target_prototype` - (Optional, List) Share targets for the file share.
  - `name` - (Optional, string) The user-defined name for this share target. Names must be unique within the share the share target resides in. If unspecified, the name will be a hyphenated list of randomly-selected words.
  - `vpc` - (Required, string) The VPC in which instances can mount the file share using this share target.This property will be removed in a future release.The `subnet` property should be used instead.
- `source_share` - (Optional, String) The ID of the source file share for this replica file share. The specified file share must not already have a replica, and must not be a replica.
- `replication_cron_spec` - (Optional, String) The cron specification for the file share replication schedule.
- `zone` - (Required, string) The globally unique name for this zone.
- `access_tags`  - (Optional, List of Strings) The list of access management tags to attach to the share. **Note** For more information, about creating access tags, see [working with tags](https://cloud.ibm.com/docs/account?topic=account-tag).
- `tags`  - (Optional, List of Strings) The list of user tags to attach to the share.

## Attribute Reference

The following attributes are exported:

- `id` - The unique identifier of the Share.
- `created_at` - The date and time that the file share is created.
- `crn` - The CRN for this share.
- `encryption` - The type of encryption used for this file share.
- `href` - The URL for this share.
- `last_sync_at` - The date and time that the file share was last synchronized to its replica.This property will be present when the `replication_role` is `source`.
- `latest_job` - The latest job associated with this file share.This property will be absent if no jobs have been created for this file share. Nested `latest_job` blocks have the following structure:
  - `status` - The status of the file share job
  - `status_reasons` - The reasons for the file share job status (if any). Nested `status_reasons` blocks have the following structure:
    - `code` - A snake case string succinctly identifying the status reason.
    - `message` - An explanation of the status reason.
    - `more_info` - Link to documentation about this status reason.
  - `type` - The type of the file share job
- `lifecycle_state` - The lifecycle state of the file share.
- `resource_type` - The type of resource referenced.
- `encryption_key` - The CRN of the [Key Protect Root Key](https://cloud.ibm.com/docs/key-protect?topic=key-protect-getting-started-tutorial) or [Hyper Protect Crypto Service Root Key](https://cloud.ibm.com/docs/hs-crypto?topic=hs-crypto-get-started) for this resource.
- `iops` - The maximum input/output operation performance bandwidth per second for the file share.
- `resource_group` - The unique identifier for this resource group. If unspecified, the account's [default resourcegroup](https://cloud.ibm.com/apidocs/resource-manager#introduction) is used.
- `replication_role`  - The replication role of the file share.* `none`: This share is not participating in replication.* `replica`: This share is a replication target.* `source`: This share is a replication source.
- `replication_status` - "The replication status of the file share.* `initializing`: This share is initializing replication.* `active`: This share is actively participating in replication.* `failover_pending`: This share is performing a replication failover.* `split_pending`: This share is performing a replication split.* `none`: This share is not participating in replication.* `degraded`: This share's replication sync is degraded.* `sync_pending`: This share is performing a replication sync.
- `replication_status_reasons` - The reasons for the current replication status (if any). Nested `replication_status_reasons` blocks have the following structure:
  - `code` - A snake case string succinctly identifying the status reason.
  - `message` - An explanation of the status reason.
  - `more_info` - Link to documentation about this status reason.
- `share_targets` - Mount targets for the file share. Nested `share_targets` blocks have the following structure:
	- `deleted` - If present, this property indicates the referenced resource has been deleted and providessome supplementary information. Nested `deleted` blocks have the following structure:
		- `more_info` - Link to documentation about deleted resources.
	- `href` - The URL for this share target.
	- `id` - The unique identifier for this share target.
	- `name` - The user-defined name for this share target.
	- `resource_type` - The type of resource referenced.
- `access_tags`  - (String) Access management tags associated to the share.
- `tags`  - (String) User tags associated for to the share.