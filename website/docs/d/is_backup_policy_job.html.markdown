---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : ibm_is_backup_policy_job"
description: |-
  Get information about BackupPolicyJob
---

# ibm_is_backup_policy_job

Provides a read-only data source for BackupPolicyJob. For more information, about backup policy in your IBM Cloud VPC, see [Backup policy jobs](https://cloud.ibm.com/docs/vpc?topic=vpc-backup-view-policy-jobs).

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
data "ibm_is_backup_policy_job" "example" {
	backup_policy_id = ibm_is_backup_policy.example.id
	identifier = "0fe9e5d8-0a4d-4818-96ec-e99708644a58"
}
```

## Argument Reference

Review the argument reference that you can specify for your data source.

- `backup_policy_id` - (Required, String) The backup policy identifier.
- `identifier` - (Required, String) The backup policy job identifier.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

- `id` - The unique identifier of the BackupPolicyJob.
- `auto_delete` - (Boolean) Indicates whether this backup policy job will be automatically deleted after it completes. At present, this is always `true`, but may be modifiable in the future.
- `auto_delete_after` - (Integer) If `auto_delete` is `true`, the days after completion that this backup policy job will be deleted. This value may be modifiable in the future.
- `backup_policy_plan` - (List) The backup policy plan operated this backup policy job (may be [deleted](https://cloud.ibm.com/apidocs/vpc#deleted-resources)).
	
	Nested scheme for `backup_policy_plan`:
	- `deleted` - (List) If present, this property indicates the referenced resource has been deleted and providessome supplementary information.
		
		Nested scheme for `deleted`:
			- `more_info` - (String) Link to documentation about deleted resources.
	- `href` - (String) The URL for this backup policy plan.
	- `id` - (String) The unique identifier for this backup policy plan.
	- `name` - (String) The unique user-defined name for this backup policy plan.
	- `resource_type` - (String) The resource type.
- `completed_at` - (String) The date and time that the backup policy job was completed.
- `created_at` - (String) The date and time that the backup policy job was created.
- `href` - (String) The URL for this backup policy job.
- `job_type` - (String) The type of backup policy job.The enumerated values for this property will expand in the future. When processing this property, check for and log unknown values. Optionally halt processing and surface the error, or bypass the backup policy job on which the unexpected property value was encountered.
  - Constraints: Allowable values are: `creation`, `deletion`.
- `resource_type` - (String) The resource type.
- `source_volume` - (List) The source volume this backup was created from (may be [deleted](https://cloud.ibm.com/apidocs/vpc#deleted-resources)).
	
	Nested scheme for `source_volume`:
	- `crn` - (String) The CRN for this volume.
	- `deleted` - (List) If present, this property indicates the referenced resource has been deleted and providessome supplementary information.
		
		Nested scheme for `deleted`:
		- `more_info` - (String) Link to documentation about deleted resources.
	- `href` - (String) The URL for this volume.
	- `id` - (String) The unique identifier for this volume.
	- `name` - (String) The unique user-defined name for this volume.
	- `remote` - (List) If present, this property indicates that the referenced resource is remote to this region, and identifies the native region.

		Nested scheme for `remote`:
		- `href` - (String) The URL for this region.
		- `name` - (String) The globally unique name for this region.
	- `resource_type` - (String) The resource type.
- `source_instance` - (List) The source instance this backup was created from (may be [deleted](https://cloud.ibm.com/apidocs/vpc#deleted-resources)).
	
	Nested scheme for `source_instance`:
	- `crn` - (String) The CRN for this virtual server instance.
	- `deleted` - (List) If present, this property indicates the referenced resource has been deleted and providessome supplementary information.
		
		Nested scheme for `deleted`:
		- `more_info` - (String) Link to documentation about deleted resources.
	- `href` - (String) The URL for this virtual server instance.
	- `id` - (String) The unique identifier for this virtual server instance.
	- `name` - (String) The unique user-defined name for this virtual server instance.
	- `resource_type` - (String) The resource type.
- `source_share` - (List) The source share this backup was created from (may be [deleted](https://cloud.ibm.com/apidocs/vpc#deleted-resources))
		
	Nested scheme for `source_volume`:
	- `crn` - (String) The CRN for this share.
	- `deleted` - (List) If present, this property indicates the referenced resource has been deleted and provides some supplementary information.
		
		Nested scheme for `deleted`:
		- `more_info` - (String) Link to documentation about deleted resources.
	- `href` - (String) The URL for this share.
	- `id` - (String) The unique identifier for this share.
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
	- `name` - (String) The unique user-defined name for this share.
	- `resource_type` - (String) The resource type.
- `status` - (String) The status of the backup policy job.The enumerated values for this property will expand in the future. When processing this property, check for and log unknown values. Optionally halt processing and surface the error, or bypass the backup policy job on which the unexpected property value was encountered.

- `status_reasons` - (List) The reasons for the current status (if any).The enumerated reason code values for this property will expand in the future. When processing this property, check for and log unknown values. Optionally halt processing and surface the error, or bypass the resource on which the unexpected reason code was encountered.
	Nested scheme for `status_reasons`:
	- `code` - (String) A snake case string succinctly identifying the status reason
		- `internal_error`: Internal error (contact IBM support)
		- `snapshot_pending`: Cannot delete backup (snapshot) in the `pending` lifecycle state
		- `snapshot_volume_limit`: The snapshot limit for the source volume has been reached
		- `source_volume_busy`: The source volume has `busy` set (after multiple retries).
	- `message` - (String) An explanation of the status reason.
	- `more_info` - (String) Link to documentation about this status reason.

- `target_snapshot` - (List) The snapshot operated on by this backup policy job (may be [deleted](https://cloud.ibm.com/apidocs/vpc#deleted-resources)).
	
	Nested scheme for `target_snapshot`:
	- `crn` - (String) The CRN for this snapshot.
	- `deleted` - (List) If present, this property indicates the referenced resource has been deleted and provides some supplementary information.
		
		Nested scheme for `deleted`:
		- `more_info` - (String) Link to documentation about deleted resources.
	- `href` - (String) The URL for this snapshot.
	- `id` - (String) The unique identifier for this snapshot.
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
	- `name` - (String) The user-defined name for this snapshot.
	- `resource_type` - (String) The resource type.

