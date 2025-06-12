---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : ibm_is_snapshot_consistency_groups"
description: |-
  Get information about SnapshotConsistencyGroupCollection
---

# ibm_is_snapshot_consistency_groups

Provides a read-only data source to retrieve information about a SnapshotConsistencyGroupCollection. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

**Note:** 
VPC infrastructure services are a regional specific based endpoint, by default targets to `us-south`. Please make sure to target right region in the provider block as shown in the `provider.tf` file, if VPC service is created in region other than `us-south`.

**provider.tf**

```terraform
provider "ibm" {
  region = "eu-gb"
}
```

## Example Usage

```terraform
data "ibm_is_snapshot_consistency_groups" "is_snapshot_consistency_groups" {
	name = "example-snapshot-consistency-group"
}
```

## Argument Reference

You can specify the following arguments for this data source.

- `backup_policy_plan` - (Optional, String) Filters the collection to backup policy jobs with a `backup_policy_plan.id` property matching the specified identifier.
- `name` - (Optional, String) Filters the collection to resources with a `name` property matching the exact specified name.
- `resource_group` - (Optional, String) Filters the collection to resources with a `resource_group.id` property matching the specified identifier.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

- `id` - The unique identifier of the SnapshotConsistencyGroupCollection.
- `snapshot_consistency_groups` - (List) Collection of snapshot consistency groups.
	
	Nested schema for `snapshot_consistency_groups`:
	- `backup_policy_plan` - (List) If present, the backup policy plan which created this snapshot consistency group.
		Nested schema for `backup_policy_plan`:
		- `deleted` - (List) If present, this property indicates the referenced resource has been deleted, and providessome supplementary information.
			
			Nested schema for `deleted`:
			- `more_info` - (String) Link to documentation about deleted resources.
		- `href` - (String) The URL for this backup policy plan.
		- `id` - (String) The unique identifier for this backup policy plan.
		- `name` - (String) The name for this backup policy plan. The name is unique across all plans in the backup policy.
		- `remote` - (List) If present, this property indicates that the resource associated with this referenceis remote and therefore may not be directly retrievable.
			
			Nested schema for `remote`:
			- `href` - (String) The URL for this region.
			- `name` - (String) The globally unique name for this region.
		- `resource_type` - (String) The resource type.
	- `created_at` - (String) The date and time that this snapshot consistency group was created.
	- `crn` - (String) The CRN of this snapshot consistency group.
	- `delete_snapshots_on_delete` - (Boolean) Indicates whether deleting the snapshot consistency group will also delete the snapshots in the group.
	- `href` - (String) The URL for this snapshot consistency group.
	- `id` - (String) The unique identifier for this snapshot consistency group.
	- `lifecycle_state` - (String) The lifecycle state of this snapshot consistency group.
	- `name` - (String) The name for this snapshot consistency group. The name is unique across all snapshot consistency groups in the region.
	- `resource_group` - (List) The resource group identifier for this snapshot consistency group.
	- `resource_type` - (String) The resource type.
	- `service_tags` - (List) The [service tags](https://cloud.ibm.com/apidocs/tagging#types-of-tags)[`is.instance:` prefix](https://cloud.ibm.com/docs/vpc?topic=vpc-snapshots-vpc-faqs) associated with this snapshot consistency group.
	- `snapshots` - (List) The member snapshots that are data-consistent with respect to captured time. (may be[deleted](https://cloud.ibm.com/apidocs/vpc#deleted-resources)).
		
		Nested schema for `snapshots`:
		- `crn` - (String) The CRN of this snapshot.
		- `deleted` - (List) If present, this property indicates the referenced resource has been deleted, and providessome supplementary information.
			
			Nested schema for `deleted`:
			- `more_info` - (String) Link to documentation about deleted resources.
		- `href` - (String) The URL for this snapshot.
		- `id` - (String) The unique identifier for this snapshot.
		- `name` - (String) The name for this snapshot. The name is unique across all snapshots in the region.
		- `remote` - (List) If present, this property indicates that the resource associated with this referenceis remote and therefore may not be directly retrievable.
			
			Nested schema for `remote`:
			- `href` - (String) The URL for this region.
			- `name` - (String) The globally unique name for this region.
		- `resource_type` - (String) The resource type.
