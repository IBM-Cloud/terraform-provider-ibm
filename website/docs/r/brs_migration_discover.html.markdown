---
layout: "ibm"
page_title: "IBM : ibm_brs_migration_discover"
description: |-
  Manages brs_migration_discover.
subcategory: "IBM Cloud Backup and Recovery Migration API"
---

# ibm_brs_migration_discover

Create, update, and delete brs_migration_discovers with this resource.

## Example Usage

```hcl
resource "ibm_brs_migration_discover" "brs_migration_discover_instance" {
  env = "classic"
  migration_id = "mgr-a1b2c3d4-e5f6-7890-abcd-ef12345678ab"
}
```

## Argument Reference

You can specify the following arguments for this resource.

* `env` - (Required, Forces new resource, String) Infrastructure environment being discovered.
  * Constraints: Allowable values are: `classic`, `vpc`.
* `migration_id` - (Required, Forces new resource, String) The migration project ID (mgr-{uuid4} format).
  * Constraints: Length must be `40` characters. The value must match regular expression `/^mgr-[0-9a-f-]{36}$/`.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the brs_migration_discover.
* `end_time` - (String) End of the time window used for this discovery run.
* `job_id` - (String) Unique discovery job ID (job-{uuid4} format).
  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^[0-9a-zA-Z-]+$/`.
* `message` - (String) Human-readable status or error message.
  * Constraints: The maximum length is `1024` characters. The minimum length is `1` character. The value must match regular expression `/^.+$/`.
* `start_time` - (String) Start of the time window used for this discovery run.
* `state` - (String) Current lifecycle state of the discovery job.
  * Constraints: Allowable values are: `pending`, `running`, `completed`, `failed`, `canceled`.
* `summary` - (List) Counts of discovered resources by compute and storage type.
Nested schema for **summary**:
	* `compute` - (List) Compute resource counts by type.
	Nested schema for **compute**:
		* `bare_metal` - (Integer) Number of bare metal servers discovered.
		  * Constraints: The maximum value is `100000`. The minimum value is `0`.
		* `virtual_server` - (Integer) Number of Virtual Server Instances discovered.
		  * Constraints: The maximum value is `100000`. The minimum value is `0`.
	* `storage` - (List) Storage volume counts by type.
	Nested schema for **storage**:
		* `block` - (Integer) Number of block volumes discovered.
		  * Constraints: The maximum value is `100000`. The minimum value is `0`.
		* `file` - (Integer) Number of file shares discovered.
		  * Constraints: The maximum value is `100000`. The minimum value is `0`.
		* `local` - (Integer) Number of local disks discovered.
		  * Constraints: The maximum value is `100000`. The minimum value is `0`.
		* `san` - (Integer) Number of SAN volumes discovered (Classic only).
		  * Constraints: The maximum value is `100000`. The minimum value is `0`.
	* `total` - (Integer) Total number of compute resources discovered.
	  * Constraints: The maximum value is `100000`. The minimum value is `0`.


## Import

You can import the `ibm_brs_migration_discover` resource by using `id`.
The `id` property can be formed from `migration_id`, and `job_id` in the following format:

<pre>
&lt;migration_id&gt;/&lt;job_id&gt;
</pre>
* `migration_id`: A string in the format `mgr-a1b2c3d4-e5f6-7890-abcd-ef12345678ab`. The migration project ID (mgr-{uuid4} format).
* `job_id`: A string in the format `job-12345678-abcd-ef01-2345-678901234567`. Unique discovery job ID (job-{uuid4} format).

# Syntax
<pre>
$ terraform import ibm_brs_migration_discover.brs_migration_discover &lt;migration_id&gt;/&lt;job_id&gt;
</pre>
