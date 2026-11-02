---
layout: "ibm"
page_title: "IBM : ibm_brs_migration_host"
description: |-
  Manages brs_migration_host.
subcategory: "IBM Cloud Backup and Recovery Migration API"
---

# ibm_brs_migration_host

Create, update, and delete brs_migration_hosts with this resource.

## Example Usage

```hcl
resource "ibm_brs_migration_host" "brs_migration_host_instance" {
  env = "classic"
  migration_id = "mgr-a1b2c3d4-e5f6-7890-abcd-ef12345678ab"
  type = "virtual_server"
}
```

## Argument Reference

You can specify the following arguments for this resource.

* `env` - (Required, Forces new resource, String) Infrastructure environment this host belongs to.
  * Constraints: Allowable values are: `vpc`, `classic`.
* `migration_id` - (Required, Forces new resource, String) The migration project ID (mgr-{uuid4} format).
  * Constraints: Length must be `40` characters. The value must match regular expression `/^mgr-[0-9a-f-]{36}$/`.
* `type` - (Required, Forces new resource, String) Whether the host is a Virtual Server Instance or bare metal server.
  * Constraints: Allowable values are: `virtual_server`, `bare_metal`.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the brs_migration_host.
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
* `host_id` - (String) Migration service host ID (host-{uuid4} format).
  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^[0-9a-zA-Z-]+$/`.
* `migrated` - (Boolean) Set to true when `POST /migrations/{migration_id}/workloads/{workload_id}/complete` is called for a workload that includes this host.
  * Constraints: The default value is `false`.
* `registered_at` - (String) Timestamp when this host was registered in the Migration API.
* `volume_attachments` - (List) Per-volume attachment records for this host.
  * Constraints: The maximum length is `1000` items. The minimum length is `0` items.
Nested schema for **volume_attachments**:
	* `volume_id` - (String) Migration service volume ID (vol-* prefix).
	  * Constraints: The maximum length is `1024` characters. The minimum length is `1` character. The value must match regular expression `/^.+$/`.
* `workload_id` - (String) ID of the workload this host is associated with. Null when the host has not been added to any workload yet.
  * Constraints: The maximum length is `1024` characters. The minimum length is `1` character. The value must match regular expression `/^.+$/`.


## Import

You can import the `ibm_brs_migration_host` resource by using `id`.
The `id` property can be formed from `migration_id`, and `host_id` in the following format:

<pre>
&lt;migration_id&gt;/&lt;host_id&gt;
</pre>
* `migration_id`: A string in the format `mgr-a1b2c3d4-e5f6-7890-abcd-ef12345678ab`. The migration project ID (mgr-{uuid4} format).
* `host_id`: A string in the format `host-a1b2c3d4-e5f6-7890-abcd-ef1234567890`. Migration service host ID (host-{uuid4} format).

# Syntax
<pre>
$ terraform import ibm_brs_migration_host.brs_migration_host &lt;migration_id&gt;/&lt;host_id&gt;
</pre>
