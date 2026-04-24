---
layout: "ibm"
page_title: "IBM : ibm_is_volume_job_cancel"
description: |-
  Helps cancelling is_volume_job.
subcategory: "Virtual Private Cloud API"
---

# ibm_is_volume_job_cancel

Cancel ibm_is_volume_job with this resource.

## Example Usage

```hcl
resource "ibm_is_volume_job" "example" {
  volume_id = ibm_is_volume.example.id
  job_type = "migrate"
  name = "my-volume-job"
  parameters {
		bandwidth 	= 1000
		iops 		= 10000
		profile {
			name = "sdp"
		}
  }
}

resource "ibm_is_volume_job_cancel" "cancel_migration" {
  volume_id     = ibm_is_volume_job.example.volume_id
  volume_job_id = ibm_is_volume_job.example.volume_job_id
}
```

## Argument Reference

You can specify the following arguments for this resource.

- `volume_job_id` - (Required, Forces new resource, String) The volume job identifier.
- `volume_id` - (Required, Forces new resource, String) The volume identifier.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

- `id` - The unique identifier of the is_volume_job_cancel.
- `job_type` - (String) The type of volume job. Constraints: Allowable values are: `migrate`.
- `name` - (String) The name for this volume job.
- `parameters` - (List) The parameters to use after the volume is migrated.
	
	Nested schema for **parameters**:
	- `bandwidth` - (Integer) The maximum bandwidth (in megabits per second) for the volume.
	- `iops` - (Integer) The maximum I/O operations per second (IOPS) for this volume.
	- `profile` - (List) Identifies a volume profile by a unique property.
		
		Nested schema for **profile**:
		- `href` - (String) The URL for this volume profile.
		- `name` - (String) The globally unique name for this volume profile.
- `auto_delete` - (Boolean) Indicates whether this volume job will be automatically deleted after it completes. At present, this is always `false`, but may be modifiable in the future.
- `completed_at` - (String) The date and time that the volume job was completed.If absent, the volume job has not yet completed.
- `created_at` - (String) The date and time that the volume job was created.
- `estimated_completion_at` - (String) The date and time that the volume job is estimated to complete.If absent, the volume job is still queued and has not yet started.
- `href` - (String) The URL for this volume job.
- `volume_job_id` - (String) The unique identifier for this volume job.
- `resource_type` - (String) The resource type. Constraints: Allowable values are: `volume_job`. T
- `started_at` - (String) The date and time that the volume job was started.If absent, the volume job has not yet started.
	- `status` - (String) The status of this volume job:
	- `deleting`:   job is being deleted
	- `failed`:     job could not be completed successfully
	- `queued`:     job is queued
	- `running`:    job is in progress
	- `succeeded`:  job was completed successfully
	- `canceling`: job is being canceled
	- `canceled`:  job is canceled
	<br/>
	The enumerated values for this property may [expand](https://cloud.ibm.com/apidocs/vpc#property-value-expansion) in the future. Constraints: Allowable values are: `canceled`, `canceling`, `deleting`, `failed`, `queued`, `running`, `succeeded`. The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `/^[a-z][a-z0-9]*(_[a-z0-9]+)*$/`.
- `status_reasons` - (List) The reasons for the current status (if any).
  * Constraints: The minimum length is `0` items.
	Nested schema for **status_reasons**:
	- `code` - (String) A snake case string succinctly identifying the status reason.The enumerated values for this property may [expand](https://cloud.ibm.com/apidocs/vpc#property-value-expansion) in the future. Constraints: Allowable values are: `internal_error`, `virtual_instance_powered_off`, `volume_detached_from_virtual_instance`. 
	- `message` - (String) An explanation of the status reason.
	- `more_info` - (String) A link to documentation about this status reason.

