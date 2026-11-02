---
layout: "ibm"
page_title: "IBM : ibm_brs_migration_discover"
description: |-
  Get information about brs_migration_discover
subcategory: "IBM Backup Recovery API"
---

# ibm_brs_migration_discover

Provides a read-only data source to retrieve information about a brs_migration_discover. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_brs_migration_discover" "brs_migration_discover" {
	brs_instance = ibm_brs_migration_discover.brs_migration_discover_instance.brs_instance
	job_id = "job-12345678-abcd-ef01-2345-678901234567"
	migration_id = ibm_brs_migration_discover.brs_migration_discover_instance.migration_id
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `brs_instance` - (Required, Forces new resource, String) The BRS service instance identifier — either a bare UUID (instance ID) or a full BRS service instance CRN (percent-encoded when supplied in a URL, since CRNs contain colons and forward-slashes).
  * Constraints: The maximum length is `512` characters. The minimum length is `1` character. The value must match regular expression `/^([0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}|crn:v1:bluemix:public:backup-recovery:[a-z0-9-]+:a\/[a-zA-Z0-9]+:[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}::)$/`.
* `job_id` - (Required, Forces new resource, String) The unique ID of the discovery job (job-{uuid4} format).
  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^job-[0-9a-f-]{36}$/`.
* `migration_id` - (Required, Forces new resource, String) The migration project ID (uuid4 format).
  * Constraints: Length must be `36` characters. The value must match regular expression `/^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$/`.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the brs_migration_discover.
* `end_time` - (String) End of the time window used for this discovery run.
* `env` - (String) Infrastructure environment being discovered.
  * Constraints: Allowable values are: `classic`, `vpc`.
* `location` - (String) Location scope used for this discovery job (datacenter or zone).
  * Constraints: The maximum length is `32` characters. The minimum length is `3` characters. The value must match regular expression `/^[a-z][a-z0-9-]+[0-9]$/`.
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
* `vpc_id` - (String) VPC ID used to scope discovery. Present for env=vpc only.
  * Constraints: The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `/^r[0-9]+-[0-9a-zA-Z-]+$/`.

