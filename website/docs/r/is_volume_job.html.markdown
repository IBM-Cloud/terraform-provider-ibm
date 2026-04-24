---
layout: "ibm"
page_title: "IBM : ibm_is_volume_job"
description: |-
  Manages is_volume_job.
subcategory: "Virtual Private Cloud API"
---

# ibm_is_volume_job

Create, update, and delete is_volume_job with this resource.

~> When using volume_job to migrate volume from one profile to another, the `profile` on the actual volume must be the same as the `profile` on the `ibm_is_volume_job`, otherwise terraform would show change on the next refresh/plan/apply. Either update the volume profile on the `ibm_is_volume` or use meta_arguments, ignore_changes on `ibm_is_volume`

## Example Usage

```terraform
resource "ibm_is_volume" "example" {
  name     = "my-volume"
  capacity = 10
  profile  = "general-purpose"   // either update the profile (eg to `sdp` for this example )or add a lifecycle block to ignore changes on the profile
  zone     = "us-south-2"
  lifecycle {
    ignore_changes = [ profile ]
  }
}


resource "ibm_is_volume_job" "example" {
  volume_id = ibm_is_volume.example.id
  job_type 	= "migrate"
  name 		= "my-volume-job"
  parameters {
		bandwidth 	= 1000
		iops 		= 10000
		profile {
			name = "sdp"
		}
  }
}
```

## Argument Reference

You can specify the following arguments for this resource.

- `volume_id` - (Required, Forces new resource, String) The volume identifier.
- `job_type` - (Required, String) The type of volume job.The enumerated values for this property may [expand](https://cloud.ibm.com/apidocs/vpc#property-value-expansion) in the future. Constraints: Allowable values are: `migrate`. 
- `name` - (Optional, String) The name for this volume job. The name must not be used by another volume job for this volume.
- `parameters` - (Optional, List) The parameters to use after the volume is migrated.
	
	Nested schema for **parameters**:
	- `bandwidth` - (Optional, Integer) The maximum bandwidth (in megabits per second) for the volume.If specified, the volume profile must not have a `bandwidth.type` of `dependent`. Constraints: The maximum value is `8192`. The minimum value is `1000`.
	- `iops` - (Optional, Integer) The maximum I/O operations per second (IOPS) for this volume.If specified, the volume profile must not have a `iops.type` of `dependent`.
	- `profile` - (Required, List) Identifies a volume profile by a unique property.
		
		Nested schema for **profile**:
		- `href` - (Optional, String) The URL for this volume profile.
		- `name` - (Optional, String) The globally unique name for this volume profile.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

- `id` - The unique identifier of the `ibm_is_volume_job`. &lt;volume_id&gt;/&lt;volume_job_id&gt;
- `auto_delete` - (Boolean) Indicates whether this volume job will be automatically deleted after it completes. At present, this is always `false`, but may be modifiable in the future.
- `completed_at` - (String) The date and time that the volume job was completed.If absent, the volume job has not yet completed.
- `created_at` - (String) The date and time that the volume job was created.
- `estimated_completion_at` - (String) The date and time that the volume job is estimated to complete.If absent, the volume job is still queued and has not yet started.
- `href` - (String) The URL for this volume job.
- `volume_job_id` - (String) The unique identifier for this volume job.
- `resource_type` - (String) The resource type. Constraints: Allowable values are: `volume_job`.
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
	The enumerated values for this property may [expand](https://cloud.ibm.com/apidocs/vpc#property-value-expansion) in the future. Constraints: Allowable values are: `canceled`, `canceling`, `deleting`, `failed`, `queued`, `running`, `succeeded`. 
- `status_reasons` - (List) The reasons for the current status (if any).

	Nested schema for **status_reasons**:
	- `code` - (String) A snake case string succinctly identifying the status reason.The enumerated values for this property may [expand](https://cloud.ibm.com/apidocs/vpc#property-value-expansion) in the future. Constraints: Allowable values are: `internal_error`, `virtual_instance_powered_off`, `volume_detached_from_virtual_instance`. T
	- `message` - (String) An explanation of the status reason.
	- `more_info` - (String) A link to documentation about this status reason.


## Import

You can import the `ibm_is_volume_job` resource by using `id`.
The `id` property can be formed from `volume_id`, and `volume_job_id` in the following format:

<pre>
&lt;volume_id&gt;/&lt;volume_job_id&gt;
</pre>
- `volume_id`: A string. The volume identifier.
- `volume_job_id`: A string in the format `r006-095e9baf-01d4-4e29-986e-20d26606b82a`. The unique identifier for this volume job.

# Syntax
<pre>
$ terraform import ibm_is_volume_job.is_volume_job &lt;volume_id&gt;/&lt;volume_job_id&gt;
</pre>
