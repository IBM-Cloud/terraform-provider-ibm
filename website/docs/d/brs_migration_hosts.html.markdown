---
layout: "ibm"
page_title: "IBM : ibm_brs_migration_hosts"
description: |-
  Get information about brs_migration_hosts
subcategory: "IBM Cloud Backup and Recovery Migration API"
---

# ibm_brs_migration_hosts

Provides a read-only data source to retrieve information about brs_migration_hosts. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_brs_migration_hosts" "brs_migration_hosts" {
	migration_id = ibm_brs_migration_host.brs_migration_host_instance.migration_id
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `env` - (Optional, String) Filter by infrastructure environment.
  * Constraints: Allowable values are: `classic`, `vpc`.
* `migrated` - (Optional, Boolean) Filter by migration status.
* `migration_id` - (Required, Forces new resource, String) The migration project ID (mgr-{uuid4} format).
  * Constraints: Length must be `40` characters. The value must match regular expression `/^mgr-[0-9a-f-]{36}$/`.
* `type` - (Optional, String) Filter by compute type.
  * Constraints: Allowable values are: `virtual_server`, `bare_metal`.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the brs_migration_hosts.
* `hosts` - (List) The list of registered hosts on this page.
  * Constraints: The maximum length is `100` items. The minimum length is `0` items.
Nested schema for **hosts**:
	* `compute` - (List) Enriched compute details. Schema variant matches the sibling `env` field: `classic` → `ClassicComputeDetails`, `vpc` → `VPCComputeDetails`.
	Nested schema for **compute**:
		* `boot_volume_id` - (String) Migration service `vol-*` ID of the boot volume attachment.
		  * Constraints: The maximum length is `1024` characters. The minimum length is `1` character. The value must match regular expression `/^.+$/`.
		* `cpu_architecture` - (String) CPU architecture of the instance.
		  * Constraints: Allowable values are: `amd64`, `s390x`.
		* `crn` - (String) IBM Cloud Resource Name for this instance.
		  * Constraints: The maximum length is `1024` characters. The minimum length is `1` character. The value must match regular expression `/^.+$/`.
		* `datacenter` - (String) Classic datacenter (e.g. dal10).
		  * Constraints: The maximum length is `1024` characters. The minimum length is `1` character. The value must match regular expression `/^.+$/`.
		* `global_identifier` - (String) GUID that uniquely identifies this instance in the infrastructure.
		  * Constraints: The maximum length is `1024` characters. The minimum length is `1` character. The value must match regular expression `/^.+$/`.
		* `health_state` - (String) Health state as reported by the VPC API. Same enum on both Virtual Server Instance and bare metal.
		  * Constraints: Allowable values are: `ok`, `degraded`, `faulted`, `inapplicable`.
		* `image_id` - (String) Boot image ID. Optional — VPC Virtual Server Instance only, not present on bare metal.
		  * Constraints: The maximum length is `1024` characters. The minimum length is `1` character. The value must match regular expression `/^.+$/`.
		* `ip_address` - (String) Primary IP address of the instance.
		  * Constraints: The maximum length is `1024` characters. The minimum length is `1` character. The value must match regular expression `/^.+$/`.
		* `lifecycle_state` - (String) Lifecycle state as returned by the VPC API. Present on both VPC Virtual Server Instance and bare metal.
		  * Constraints: Allowable values are: `deleting`, `failed`, `pending`, `stable`, `suspended`, `updating`, `waiting`.
		* `memory_gib` - (Integer) Memory in gibibytes (GiB).
		  * Constraints: The maximum value is `4096`. The minimum value is `0`.
		* `name` - (String) Display name or hostname of the instance.
		  * Constraints: The maximum length is `1024` characters. The minimum length is `1` character. The value must match regular expression `/^.+$/`.
		* `os_family` - (String) OS family of the instance.
		  * Constraints: Allowable values are: `linux`, `windows`.
		* `os_type` - (String) OS image identifier as returned by the infrastructure API.
		  * Constraints: The maximum length is `1024` characters. The minimum length is `1` character. The value must match regular expression `/^.+$/`.
		* `profile` - (String) Instance profile name.
		  * Constraints: The maximum length is `1024` characters. The minimum length is `1` character. The value must match regular expression `/^.+$/`.
		* `public_ips` - (List) Public IP addresses on this host. For VPC this maps from `floating_ips` on the primary network interface. For Classic this maps from `primaryPublicIpAddress`. Empty array when none.
		  * Constraints: The list items must match regular expression `/^.+$/`. The maximum length is `1000` items. The minimum length is `0` items.
		* `region` - (String) VPC region.
		  * Constraints: The maximum length is `1024` characters. The minimum length is `1` character. The value must match regular expression `/^.+$/`.
		* `resource_group_id` - (String) ID of the IBM Cloud resource group this instance belongs to.
		  * Constraints: The maximum length is `1024` characters. The minimum length is `1` character. The value must match regular expression `/^.+$/`.
		* `security_groups` - (List) Security group IDs on the primary network interface. Empty array when none.
		  * Constraints: The list items must match regular expression `/^.+$/`. The maximum length is `1000` items. The minimum length is `0` items.
		* `status` - (String) Current power/lifecycle status (union of VPC Virtual Server Instance and bare metal status enums).
		  * Constraints: Allowable values are: `pending`, `starting`, `running`, `restarting`, `stopping`, `stopped`, `deleting`, `failed`, `maintenance`, `unknown`.
		* `subnet_id` - (String) ID of the subnet the primary network interface is attached to.
		  * Constraints: The maximum length is `1024` characters. The minimum length is `1` character. The value must match regular expression `/^.+$/`.
		* `throughput_mbps` - (Integer) Total network throughput in Mbps. Required on both VPC Virtual Server Instance and bare metal.
		  * Constraints: The maximum value is `100000`. The minimum value is `0`.
		* `vcpu_count` - (Integer) Number of virtual CPUs (from vcpu.count on VPC Virtual Server Instance).
		  * Constraints: The maximum value is `512`. The minimum value is `0`.
		* `vpc_id` - (String) ID of the VPC this instance belongs to.
		  * Constraints: The maximum length is `1024` characters. The minimum length is `1` character. The value must match regular expression `/^.+$/`.
		* `vpc_name` - (String) Name of the VPC.
		  * Constraints: The maximum length is `1024` characters. The minimum length is `1` character. The value must match regular expression `/^.+$/`.
		* `zone` - (String) VPC zone.
		  * Constraints: The maximum length is `1024` characters. The minimum length is `1` character. The value must match regular expression `/^.+$/`.
	* `env` - (String) Infrastructure environment this host belongs to.
	  * Constraints: Allowable values are: `vpc`, `classic`.
	* `id` - (String) Migration service host ID (host-{uuid4} format).
	  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^[0-9a-zA-Z-]+$/`.
	* `migrated` - (Boolean) Set to true when `POST /migrations/{migration_id}/workloads/{workload_id}/complete` is called for a workload that includes this host.
	  * Constraints: The default value is `false`.
	* `registered_at` - (String) Timestamp when this host was registered in the Migration API.
	* `type` - (String) Whether the host is a Virtual Server Instance or bare metal server.
	  * Constraints: Allowable values are: `virtual_server`, `bare_metal`.
	* `volume_attachments` - (List) Per-volume attachment records for this host.
	  * Constraints: The maximum length is `1000` items. The minimum length is `0` items.
	Nested schema for **volume_attachments**:
		* `volume_id` - (String) Migration service volume ID (vol-* prefix).
		  * Constraints: The maximum length is `1024` characters. The minimum length is `1` character. The value must match regular expression `/^.+$/`.
	* `workload_id` - (String) ID of the workload this host is associated with. Null when the host has not been added to any workload yet.
	  * Constraints: The maximum length is `1024` characters. The minimum length is `1` character. The value must match regular expression `/^.+$/`.

