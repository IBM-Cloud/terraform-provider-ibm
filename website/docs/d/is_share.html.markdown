---
layout: "ibm"
page_title: "IBM : is_share"
description: |-
  Get information about Share
subcategory: "VPC infrastructure"
---

# ibm\_is_share

Provides a read-only data source for Share. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
resource "ibm_is_vpc" "example" {
  name = "example-vpc"
}
resource "ibm_is_share" "example" {
  name    = "example-share"
  size    = 200
  profile = "dp2"
  zone    = "us-south-2"
}

data "ibm_is_share" "example" {
  share = ibm_is_share.example.id
}

data "ibm_is_share" "example1" {
  name = ibm_is_share.example.name
}
```

## Argument Reference

The following arguments are supported:

- `share` - (Optional, String) The file share identifier.
- `name` - (Optional, String) The file share name
**Note** One of the aurgument is mandatory

## Attribute Reference

The following attributes are exported:

- `access_control_mode` - (Boolean) The access control mode for the share.
- `accessor_binding_role` - (String) The accessor binding role of this file share:- `none`: This file share is not participating in access with another file share- `origin`: This file share is the origin for one or more file shares  (which may be in other accounts)- `accessor`: This file share is providing access to another file share  (which may be in another account).
- `accessor_bindings` - (List) The accessor bindings for this file share. Each accessor binding identifies a resource (possibly in another account) with access to this file share's data.
  Nested schema for **accessor_bindings**:
	- `href` - (String) The URL for this share accessor binding.
	- `id` - (String) The unique identifier for this share accessor binding.
	- `resource_type` - (String) The resource type.
- `allowed_transit_encryption_modes` - (List of string) The transit encryption modes allowed for this share.
- `access_tags`  - (String) Access management tags associated to the share.
- `created_at` - The date and time that the file share is created.
- `crn` - The CRN for this share.
- `encryption` - The type of encryption used for this file share.
- `encryption_key` - The CRN of the key used to encrypt this file share.
- `href` - The URL for this share.
- `iops` - The maximum input/output operation performance bandwidth per second for the file share.
- `latest_sync` - (List) Information about the latest synchronization for this file share.
Nested `latest_sync` blocks have the following structure:
  - `completed_at` - (String) The completed date and time of last synchronization between the replica share and its source.
  - `data_transferred` - (Integer) The data transferred (in bytes) in the last synchronization between the replica and its source.
  - `started_at` - (String) The start date and time of last synchronization between the replica share and its source.
- `latest_job` - The latest job associated with this file share.This property will be absent if no jobs have been created for this file share. Nested `latest_job` blocks have the following structure:
  - `status` - The status of the file share job
  - `status_reasons` - The reasons for the file share job status (if any). Nested `status_reasons` blocks have the following structure:
    - `code` - A snake case string succinctly identifying the status reason.
    - `message` - An explanation of the status reason.
    - `more_info` - Link to documentation about this status reason.
  - `type` - The type of the file share job
- `lifecycle_state` - The lifecycle state of the file share.
- `name` - The unique user-defined name for this file share.
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
- `profile` - The name of the profile this file share uses.
- `replication_role`  - The replication role of the file share.* `none`: This share is not participating in replication.* `replica`: This share is a replication target.* `source`: This share is a replication source.
- `replication_status` - "The replication status of the file share.* `initializing`: This share is initializing replication.* `active`: This share is actively participating in replication.* `failover_pending`: This share is performing a replication failover.* `split_pending`: This share is performing a replication split.* `none`: This share is not participating in replication.* `degraded`: This share's replication sync is degraded.* `sync_pending`: This share is performing a replication sync.
- `replication_status_reasons` - The reasons for the current replication status (if any). Nested `replication_status_reasons` blocks have the following structure:
  - `code` - A snake case string succinctly identifying the status reason.
  - `message` - An explanation of the status reason.
  - `more_info` - Link to documentation about this status reason. 
- `replica_share` - The replica file share for this source file share.This property will be present when the `replication_role` is `source`. Nested `replica_share` blocks have the following structure:
  - `crn` - The CRN for this file share.
  - `deleted` - If present, this property indicates the referenced resource has been deleted and providessome supplementary information. Nested `deleted` blocks have the following structure:
    - `more_info` - Link to documentation about deleted resources.
  - `href` - The URL for this file share.
  - `id` - The unique identifier for this file share.
  - `name` - The unique user-defined name for this file share.
  - `resource_type` - The resource type.
- `resource_group` - The ID of the resource group for this file share.
- `resource_type` - The type of resource referenced.
- `size` - The size of the file share rounded up to the next gigabyte.
- `share_targets` - Mount targets for the file share. Nested `targets` blocks have the following structure:
	- `deleted` - If present, this property indicates the referenced resource has been deleted and providessome supplementary information. Nested `deleted` blocks have the following structure:
		- `more_info` - Link to documentation about deleted resources.
	- `href` - The URL for this share target.
	- `id` - The unique identifier for this share target.
	- `name` - The user-defined name for this share target.
	- `resource_type` - The type of resource referenced.
- `source_share` - The source file share for this replica file share. This property will be present when the `replication_role` is `replica`. Nested `source_share` blocks have the following structure:
  - `crn` - The CRN for this file share.
  - `deleted` - If present, this property indicates the referenced resource has been deleted and providessome supplementary information. Nested `deleted` blocks have the following structure:
    - `more_info` - Link to documentation about deleted resources.
  - `href` - The URL for this file share.
  - `id` - The unique identifier for this file share.
  - `name` - The unique user-defined name for this file share.
  - `resource_type` - The resource type.
- `tags`  - (String) User tags associated for to the share.
- `zone` - The name of the zone this file share will reside in.

