---
layout: "ibm"
page_title: "IBM : is_shares"
description: |-
  Get information about ShareCollection
subcategory: "VPC infrastructure"
---

# ibm\_is_shares

Provides a read-only data source for ShareCollection. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "ibm_is_shares" "example" {
}
```

## Argument Reference

The following arguments are supported:

- `name` - (Optional, string) The unique user-defined name for this file share to filter the collection.
- `resource_group` - (Optional, string) The unique identifier for this resource group to filter the collection.

## Attribute Reference

The following attributes are exported:

- `shares` - Collection of file shares. Nested `shares` blocks have the following structure:
	- `created_at` - The date and time that the file share is created.
	- `crn` - The CRN for this share.
	- `encryption` - The type of encryption used for this file share.
	- `encryption_key` - The CRN of the key used to encrypt this file share.
	- `href` - The URL for this share.
	- `id` - The unique identifier for this file share.
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
	- `name` - The unique user-defined name for this file share. If unspecified, the name will be a hyphenated list of randomly-selected words.
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
	- `mount_targets` - Mount targets for the file share. Nested `targets` blocks have the following structure:
    	- `deleted` - If present, this property indicates the referenced resource has been deleted and providessome supplementary information. Nested `deleted` blocks have the following structure:
    		- `more_info` - Link to documentation about deleted resources.
    	- `href` - The URL for this share target.
    	- `id` - The unique identifier for this share target.
    	- `name` - The user-defined name for this share target.
    	- `resource_type` - The type of resource referenced.
	- `zone` - The name of the zone this file share will reside in.
	- `access_tags`  - (String) Access management tags associated to the share.
	- `tags`  - (String) User tags associated for to the share.
	- `source_share` - (List) The source file share for this replica file share.This property will be present when the `replication_role` is `replica`. Nested `source_share` blocks have the following structure:
      - `crn` - (String) The CRN for this file share.
      - `deleted` - (List)  If present, this property indicates the referenced resource has been deleted and providessome supplementary information. Nested `deleted` blocks have the following structure:
        - `more_info` - (String) Link to documentation about deleted resources.
      - `href` - (String) The URL for this file share.
      - `id` - (String) The unique identifier for this file share.
      - `name` - (String) The unique user-defined name for this file share.
      - `resource_type` - (String) The resource type.
- `total_count` - The total number of resources across all pages.

