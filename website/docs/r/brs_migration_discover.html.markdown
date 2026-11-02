---
layout: "ibm"
page_title: "IBM : ibm_brs_migration_discover"
description: |-
  Manages brs_migration_discover.
subcategory: "IBM Backup Recovery API"
---

# ibm_brs_migration_discover

Create, update, and delete brs_migration_discovers with this resource.

## Example Usage

```hcl
resource "ibm_brs_migration_discover" "brs_migration_discover_instance" {
  brs_instance = "00000000-0000-4000-8000-000000000001"
  env = "classic"
  location = "us-south-1"
  migration_id = "a1b2c3d4-e5f6-7890-abcd-ef12345678ab"
  vpc_id = "r134-abcdef01-2345-6789-abcd-ef0123456789"
}
```

## Argument Reference

You can specify the following arguments for this resource.

* `brs_instance` - (Required, Forces new resource, String) The BRS service instance identifier — either a bare UUID (instance ID) or a full BRS service instance CRN (percent-encoded when supplied in a URL, since CRNs contain colons and forward-slashes).
  * Constraints: The maximum length is `512` characters. The minimum length is `1` character. The value must match regular expression `/^([0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}|crn:v1:bluemix:public:backup-recovery:[a-z0-9-]+:a\/[a-zA-Z0-9]+:[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}::)$/`.
* `env` - (Required, Forces new resource, String) Infrastructure environment being discovered.
  * Constraints: Allowable values are: `classic`, `vpc`.
* `location` - (Required, Forces new resource, String) Location scope used for this discovery job (datacenter or zone).
  * Constraints: The maximum length is `32` characters. The minimum length is `3` characters. The value must match regular expression `/^[a-z][a-z0-9-]+[0-9]$/`.
* `migration_id` - (Required, Forces new resource, String) The migration project ID (uuid4 format).
  * Constraints: Length must be `36` characters. The value must match regular expression `/^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$/`.
* `vpc_id` - (Optional, Forces new resource, String) VPC ID used to scope discovery. Present for env=vpc only.
  * Constraints: The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `/^r[0-9]+-[0-9a-zA-Z-]+$/`.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the brs_migration_discover.
* `brs_migration_discover_id` - (String) Unique discovery job ID (job-{uuid4} format).
  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^[0-9a-zA-Z-]+$/`.
* `end_time` - (String) End of the time window used for this discovery run.
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
The `id` property can be formed from `brs_instance`, `migration_id`, and `brs_migration_discover_id` in the following format:

<pre>
&lt;brs_instance&gt;/&lt;migration_id&gt;/&lt;brs_migration_discover_id&gt;
</pre>
* `brs_instance`: A string in the format `00000000-0000-4000-8000-000000000001`. The BRS service instance identifier — either a bare UUID (instance ID) or a full BRS service instance CRN (percent-encoded when supplied in a URL, since CRNs contain colons and forward-slashes).
* `migration_id`: A string in the format `a1b2c3d4-e5f6-7890-abcd-ef12345678ab`. The migration project ID (uuid4 format).
* `brs_migration_discover_id`: A string in the format `job-12345678-abcd-ef01-2345-678901234567`. Unique discovery job ID (job-{uuid4} format).

# Syntax
<pre>
$ terraform import ibm_brs_migration_discover.brs_migration_discover &lt;brs_instance&gt;/&lt;migration_id&gt;/&lt;brs_migration_discover_id&gt;
</pre>
