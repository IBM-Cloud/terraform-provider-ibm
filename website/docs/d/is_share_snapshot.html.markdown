---
layout: "ibm"
page_title: "IBM : ibm_is_share_snapshot"
description: |-
  Get information about ShareSnapshot
subcategory: "VPC infrastructure"
---

# ibm_is_share_snapshot

Provides a read-only data source to retrieve information about a ShareSnapshot. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
resource "ibm_is_share" "example" {
	zone    = "us-south-1"
	size    = 220
	name    = "%s"
	profile = "dp2"
}

resource "ibm_is_share_snapshot" "example" {
  name = "my-example-share-snapshot"
  share = ibm_is_share.example.id
  tags = ["my-example-share-snapshot-tag"]
}
data "ibm_is_share_snapshot" "example" {
	share_snapshot = ibm_is_share_snapshot.is_share_snapshot_instance.is_share_snapshot_id
	share = ibm_is_share.example.id
}
```

## Argument Reference

You can specify the following arguments for this data source.

- `share_snapshot` - (Required, Forces new resource, String) The share snapshot identifier.
- `share` - (Required, Forces new resource, String) The file share identifier.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

- `access_tags`  - (Array of Strings) Access management tags associated with the snapshot.
- `id` - The unique identifier of the ShareSnapshot.
- `backup_policy_plan` - (List) If present, the backup policy plan which created this share snapshot.
	Nested schema for **backup_policy_plan**:
	- `deleted` - (List) If present, this property indicates the referenced resource has been deleted, and providessome supplementary information.
		Nested schema for **deleted**:
		- `more_info` - (String) Link to documentation about deleted resources.
	- `href` - (String) The URL for this backup policy plan.
	- `id` - (String) The unique identifier for this backup policy plan.
	- `name` - (String) The name for this backup policy plan. The name is unique across all plans in the backup policy.
	- `remote` - (List) If present, this property indicates that the resource associated with this referenceis remote and therefore may not be directly retrievable.
		Nested schema for **remote**:
		- `region` - (List) If present, this property indicates that the referenced resource is remote to thisregion, and identifies the native region.
			Nested schema for **region**:
			- `href` - (String) The URL for this region.
			- `name` - (String) The globally unique name for this region.
	- `resource_type` - (String) The resource type.
- `captured_at` - (String) The date and time the data capture for this share snapshot was completed.If absent, this snapshot's data has not yet been captured.
- `created_at` - (String) The date and time that the share snapshot was created.
- `crn` - (String) The CRN for this share snapshot.
- `fingerprint` - (String) The fingerprint for this snapshot.
- `href` - (String) The URL for this share snapshot.
- `lifecycle_state` - (String) The lifecycle state of this share snapshot.
- `minimum_size` - (Integer) The minimum size of a share created from this snapshot. When a snapshot is created, this will be set to the size of the `source_share`.
- `name` - (String) The name for this share snapshot. The name is unique across all snapshots for the file share.
- `resource_group` - (List) The resource group for this file share.
	Nested schema for **resource_group**:
	- `href` - (String) The URL for this resource group.
	- `id` - (String) The unique identifier for this resource group.
	- `name` - (String) The name for this resource group.
- `resource_type` - (String) The resource type.
- `status` - (String) The status of the share snapshot:- `available`: The share snapshot is available for use.- `failed`: The share snapshot is irrecoverably unusable.- `pending`: The share snapshot is being provisioned and is not yet usable.- `unusable`: The share snapshot is not currently usable (see `status_reasons`)The enumerated values for this property may[expand](https://cloud.ibm.com/apidocs/vpc#property-value-expansion) in the future.
- `status_reasons` - (List) The reasons for the current status (if any).
	Nested schema for **status_reasons**:
	- `code` - (String) A reason code for the status:- `encryption_key_deleted`: File share snapshot is unusable  because its `encryption_key` was deleted- `internal_error`: Internal error (contact IBM support)The enumerated values for this property may[expand](https://cloud.ibm.com/apidocs/vpc#property-value-expansion) in the future.
	- `message` - (String) An explanation of the status reason.
	- `more_info` - (String) Link to documentation about this status reason.
- `tags` - (List) The [user tags](https://cloud.ibm.com/apidocs/tagging#types-of-tags) associated with this share snapshot.
- `zone` - (List) The zone this share snapshot resides in.
	Nested schema for **zone**:
	- `href` - (String) The URL for this zone.
	- `name` - (String) The globally unique name for this zone.

