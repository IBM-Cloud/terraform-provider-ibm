---
layout: "ibm"
page_title: "IBM : ibm_is_volume_jobs"
description: |-
  Get information about VolumeJobPaginatedCollection
subcategory: "Virtual Private Cloud API"
---

# ibm_is_volume_jobs

Provides a read-only data source to retrieve information about a VolumeJobPaginatedCollection. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_is_volume_jobs" "example" {
	volume_id = ibm_is_volume_job.example.volume_id
}
```

## Argument Reference

You can specify the following arguments for this data source.

- `volume_id` - (Required, Forces new resource, String) The volume identifier.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

- `id` - The unique identifier of the VolumeJobPaginatedCollection.
- `jobs` - (List) The jobs for this volume.
	
	Nested schema for **jobs**:
	- `auto_delete` - (Boolean) Indicates whether this volume job will be automatically deleted after it completes. At present, this is always `false`, but may be modifiable in the future.
	- `completed_at` - (String) The date and time that the volume job was completed.If absent, the volume job has not yet completed.
	- `created_at` - (String) The date and time that the volume job was created.
	- `estimated_completion_at` - (String) The date and time that the volume job is estimated to complete.If absent, the volume job is still queued and has not yet started.
	- `href` - (String) The URL for this volume job.
	- `id` - (String) The unique identifier for this volume job.
	- `job_type` - (String) The type of volume job.The enumerated values for this property may [expand](https://cloud.ibm.com/apidocs/vpc#property-value-expansion) in the future.
	- `name` - (String) The name for this volume job. The name must not be used by another volume job for this volume.
	- `parameters` - (List) The parameters to use after the volume is migrated.
		
		Nested schema for **parameters**:
		- `bandwidth` - (Integer) The maximum bandwidth (in megabits per second) for the volume.If specified, the volume profile must not have a `bandwidth.type` of `dependent`.
		- `iops` - (Integer) The maximum I/O operations per second (IOPS) for this volume.If specified, the volume profile must not have a `iops.type` of `dependent`.
		- `profile` - (List) Identifies a volume profile by a unique property.
			
			Nested schema for **profile**:
			- `href` - (String) The URL for this volume profile.
			- `name` - (String) The globally unique name for this volume profile.
	- `resource_type` - (String) The resource type.
	- `started_at` - (String) The date and time that the volume job was started.If absent, the volume job has not yet started.
	- `status` - (String) The status of this volume job:

		- `deleting`:   job is being deleted
		- `failed`:     job could not be completed successfully
		- `queued`:     job is queued
		- `running`:    job is in progress
		- `succeeded`:  job was completed successfully
		- `canceling`: job is being canceled
		- `canceled`:  job is canceled
	- `status_reasons` - (List) The reasons for the current status (if any).
		
		Nested schema for **status_reasons**:
		- `code` - (String) A snake case string succinctly identifying the status reason.The enumerated values for this property may [expand](https://cloud.ibm.com/apidocs/vpc#property-value-expansion) in the future. Constraints: Allowable values are: `internal_error`, `virtual_instance_powered_off`, `volume_detached_from_virtual_instance`. The maximum length is `128` characters.
		- `message` - (String) An explanation of the status reason.
		- `more_info` - (String) A link to documentation about this status reason.

