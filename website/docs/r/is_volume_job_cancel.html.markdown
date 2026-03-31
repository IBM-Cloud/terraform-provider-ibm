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
resource "ibm_is_volume_job" "is_volume_job_instance" {
  volume_id = ibm_is_volume.example.id
  job_type = "migrate"
  name = "my-volume-job"
  parameters {
		bandwidth = 1000
		iops = 10000
		profile {
			name = "general-purpose"
		}
  }
}

resource "ibm_is_volume_job_cancel" "cancel_migration" {
  volume_id     = ibm_is_volume_job.is_volume_job_instance.volume_id
  volume_job_id = ibm_is_volume_job.is_volume_job_instance.volume_job_id
}
```

## Argument Reference

You can specify the following arguments for this resource.

* `volume_job_id` - (Required, Forces new resource, String) The volume job identifier.
  * Constraints: The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-z_]+$/`.
* `volume_id` - (Required, Forces new resource, String) The volume identifier.
  * Constraints: The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-z_]+$/`.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the is_volume_job_cancel.
* `job_type` - (String) The type of volume job.
  * Constraints: Allowable values are: `migrate`.
* `name` - (String) The name for this volume job.
* `parameters` - (List) The parameters to use after the volume is migrated.
Nested schema for **parameters**:
	* `bandwidth` - (Integer) The maximum bandwidth (in megabits per second) for the volume.
	  * Constraints: The maximum value is `8192`. The minimum value is `1000`.
	* `iops` - (Integer) The maximum I/O operations per second (IOPS) for this volume.
	* `profile` - (List) Identifies a volume profile by a unique property.
	Nested schema for **profile**:
		* `href` - (String) The URL for this volume profile.
		* `name` - (String) The globally unique name for this volume profile.
* `auto_delete` - (Boolean) Indicates whether this volume job will be automatically deleted after it completes. At present, this is always `false`, but may be modifiable in the future.
* `completed_at` - (String) The date and time that the volume job was completed.If absent, the volume job has not yet completed.
* `created_at` - (String) The date and time that the volume job was created.
* `estimated_completion_at` - (String) The date and time that the volume job is estimated to complete.If absent, the volume job is still queued and has not yet started.
* `href` - (String) The URL for this volume job.
  * Constraints: The maximum length is `8000` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
* `volume_job_id` - (String) The unique identifier for this volume job.
  * Constraints: The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-z_]+$/`.
* `resource_type` - (String) The resource type.
  * Constraints: Allowable values are: `volume_job`. The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `/^[a-z][a-z0-9]*(_[a-z0-9]+)*$/`.
* `started_at` - (String) The date and time that the volume job was started.If absent, the volume job has not yet started.
* `status` - (String) The status of this volume job:- `deleting`:   job is being deleted- `failed`:     job could not be completed successfully- `queued`:     job is queued- `running`:    job is in progress- `succeeded`:  job was completed successfully- `canceling`: job is being canceled- `canceled`:  job is canceledThe enumerated values for this property may[expand](https://cloud.ibm.com/apidocs/vpc#property-value-expansion) in the future.
  * Constraints: Allowable values are: `canceled`, `canceling`, `deleting`, `failed`, `queued`, `running`, `succeeded`. The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `/^[a-z][a-z0-9]*(_[a-z0-9]+)*$/`.
* `status_reasons` - (List) The reasons for the current status (if any).
  * Constraints: The minimum length is `0` items.
Nested schema for **status_reasons**:
	* `code` - (String) A snake case string succinctly identifying the status reason.The enumerated values for this property may[expand](https://cloud.ibm.com/apidocs/vpc#property-value-expansion) in the future.
	  * Constraints: Allowable values are: `internal_error`, `virtual_instance_powered_off`, `volume_detached_from_virtual_instance`. The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `/^[a-z][a-z0-9]*(_[a-z0-9]+)*$/`.
	* `message` - (String) An explanation of the status reason.
	  * Constraints: The maximum length is `512` characters. The minimum length is `1` character. The value must match regular expression `/^[ -~\\n\\r\\t]*$/`.
	* `more_info` - (String) A link to documentation about this status reason.
	  * Constraints: The maximum length is `8000` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.


## Import

You can import the `ibm_is_volume_job_cancel` resource by using `id`.
The `id` property can be formed from `volume_id`, and `volume_job_id` in the following format:

<pre>
<volume_id>/<volume_job_id>
</pre>
* `volume_id`: A string. The volume identifier.
* `volume_job_id`: A string in the format `r006-095e9baf-01d4-4e29-986e-20d26606b82a`. The unique identifier for this volume job.

# Syntax
<pre>
$ terraform import ibm_is_volume_job_cancel.cancel_migration <volume_id>/<volume_job_id>
</pre>
