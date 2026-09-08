---
layout: "ibm"
page_title: "IBM : ibm_brs_migration_volume"
description: |-
  Get information about brs_migration_volume
subcategory: "IBM Cloud Backup and Recovery Migration API"
---

# ibm_brs_migration_volume

Provides a read-only data source to retrieve information about a brs_migration_volume. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_brs_migration_volume" "brs_migration_volume" {
	migration_id = ibm_brs_migration_volume.brs_migration_volume_instance.migration_id
	volume_id = ibm_brs_migration_volume.brs_migration_volume_instance.volume_id
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `migration_id` - (Required, Forces new resource, String) The migration project ID (mgr-{uuid4} format).
  * Constraints: Length must be `40` characters. The value must match regular expression `/^mgr-[0-9a-f-]{36}$/`.
* `volume_id` - (Required, Forces new resource, String) The migration service volume ID (vol-{uuid4} format).
  * Constraints: Length must be `40` characters. The value must match regular expression `/^vol-[0-9a-f-]{36}$/`.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the brs_migration_volume.
* `attachment_state` - (String) Migration-service-computed attachment status. Set to `attached` when `host_attachments` is non-empty, `unattached` when empty.
  * Constraints: Allowable values are: `attached`, `unattached`.
* `env` - (String) Infrastructure environment this volume belongs to.
  * Constraints: Allowable values are: `vpc`, `classic`.
* `host_attachments` - (List) Per-host attachment records for this volume.
  * Constraints: The maximum length is `1000` items. The minimum length is `0` items.
Nested schema for **host_attachments**:
	* `block_device` - (String) Block device path on this host. Present for block/local volumes; empty for file/NFS.
	  * Constraints: The maximum length is `1024` characters. The minimum length is `1` character. The value must match regular expression `/^.+$/`.
	* `device_id` - (String) The storage device identifier for this volume on the host. For VPC block volumes, this is the volume attachment device ID. For Classic SAN and local volumes, this is the numeric device slot index. Present for VPC block, Classic SAN, and Classic local volumes; absent for Classic iSCSI block and all file volumes.
	  * Constraints: The maximum length is `1024` characters. The minimum length is `1` character. The value must match regular expression `/^.+$/`.
	* `host_id` - (String) Migration service host ID (host-* prefix).
	  * Constraints: The maximum length is `1024` characters. The minimum length is `1` character. The value must match regular expression `/^.+$/`.
	* `mount_path` - (String) OS-level mount path of the volume on this host (e.g. /mnt/data).
	  * Constraints: The maximum length is `1024` characters. The minimum length is `1` character. The value must match regular expression `/^.+$/`.
	* `type` - (String) The filesystem type of this volume on the host (e.g. ext4, xfs, nfs4).
	  * Constraints: The maximum length is `1024` characters. The minimum length is `1` character. The value must match regular expression `/^.+$/`.
* `migrated` - (Boolean) Set to true when the workload covering this volume is completed.
  * Constraints: The default value is `false`.
* `registered_at` - (String) Timestamp when this volume was registered in the Migration API.
* `storage` - (List) Enriched storage details.
Nested schema for **storage**:
	* `access_control_mode` - (String) Access control mode for the VPC file share. Determines how mount target access is governed (`security_group` or `vpc`). Set to `none` for VPC block volumes and Classic volumes where this concept does not apply.
	  * Constraints: Allowable values are: `none`, `security_group`, `vpc`.
	* `availability_mode` - (String) Availability zone scope of the VPC file share. `regional` means the share is accessible across all zones in its region. `zonal` means the share is pinned to a single zone (profile `dp2`). Set to `none` for VPC block volumes and Classic volumes where this concept does not apply.
	  * Constraints: Allowable values are: `none`, `regional`, `zonal`.
	* `capacity_gib` - (Integer) Provisioned capacity in gibibytes (GiB).
	  * Constraints: The maximum value is `65536`. The minimum value is `0`.
	* `crn` - (String) IBM Cloud CRN for this VPC volume.
	  * Constraints: The maximum length is `1024` characters. The minimum length is `1` character. The value must match regular expression `/^.+$/`.
	* `datacenter` - (String) Classic datacenter slug where this volume is provisioned (e.g. dal10).
	  * Constraints: The maximum length is `1024` characters. The minimum length is `1` character. The value must match regular expression `/^.+$/`.
	* `encryption` - (String) Encryption type for the volume. Classic: e.g. `aes256`. VPC block: `provider_managed` or `user_managed`. VPC file shares: `provider_managed` or `user_managed`.
	  * Constraints: The maximum length is `256` characters. The minimum length is `1` character. The value must match regular expression `/^.+$/`.
	* `global_identifier` - (String) Raw infrastructure volume ID. Classic: numeric string (e.g. "98765432"). VPC: UUID (e.g. r134-abcdef01-…).
	  * Constraints: The maximum length is `1024` characters. The minimum length is `1` character. The value must match regular expression `/^.+$/`.
	* `iops` - (Integer) Maximum I/O operations per second. VPC block: from VPC API (required). Classic: from Endurance/Performance tier. Absent for file/san/local volumes.
	  * Constraints: The maximum value is `96000`. The minimum value is `0`.
	* `iscsi_target_ips` - (List) iSCSI portal IPs for this volume. Classic block (iSCSI) volumes only; absent for Classic file, SAN, and local volumes. The orchestrator passes all IPs to discover.sh; the script tries each in sequence until an active session is found.
	  * Constraints: The list items must match regular expression `/^.+$/`. The maximum length is `100` items. The minimum length is `0` items.
	* `lifecycle_state` - (String) Normalised lifecycle state of the volume. All infrastructure-specific states are mapped to this canonical set before storage. Classic `ready` → `stable`; Classic `provisioning` → `pending`. VPC block `available` → `stable`; `pending_deletion` → `deleting`; `unusable` → `failed`. VPC file share states (`stable`, `pending`, `updating`, `deleting`, `suspended`, `waiting`, `failed`) map directly.
	  * Constraints: Allowable values are: `deleting`, `failed`, `pending`, `stable`, `suspended`, `unusable`, `updating`, `waiting`.
	* `name` - (String) Human-readable name of the volume as set in IBM Cloud.
	  * Constraints: The maximum length is `1024` characters. The minimum length is `1` character. The value must match regular expression `/^.+$/`.
	* `profile` - (String) Volume profile or tier. Classic: e.g. Endurance, Performance. VPC: e.g. general-purpose, 5iops-tier.
	  * Constraints: The maximum length is `1024` characters. The minimum length is `1` character. The value must match regular expression `/^.+$/`.
	* `region` - (String) VPC region (e.g. us-south).
	  * Constraints: The maximum length is `1024` characters. The minimum length is `1` character. The value must match regular expression `/^.+$/`.
	* `replication_role` - (String) Replication role of the VPC file share. `none` means replication is not configured. `source` means this share is the replication source. `replica` means this share is the replication target (replica). File shares only.
	  * Constraints: Allowable values are: `none`, `replica`, `source`.
	* `resource_group_id` - (String) ID of the IBM Cloud resource group this volume belongs to.
	  * Constraints: The maximum length is `1024` characters. The minimum length is `1` character. The value must match regular expression `/^.+$/`.
	* `source_paths` - (List) NFS mount paths for this volume. Present for file volumes only (classic/file and vpc/file); absent for block, SAN, and local volumes. Classic file: one or more export paths, no vpc_id or mount_target_id. VPC file: one entry per VPC mount target, each with vpc_id and mount_target_id. The orchestrator passes all entries to discover.sh; the script tries each in sequence until an active mount is found.
	  * Constraints: The maximum length is `1000` items. The minimum length is `0` items.
	Nested schema for **source_paths**:
		* `mount_target_id` - (String) VPC file share mount target ID (share_mount_target_id from the VPC API). Present for VPC file volumes only; absent for Classic file volumes. Useful for lifecycle operations such as mount target teardown during workload completion.
		  * Constraints: The maximum length is `1024` characters. The minimum length is `1` character. The value must match regular expression `/^.+$/`.
		* `path` - (String) NFS mount path (host:/export format).
		  * Constraints: The maximum length is `2048` characters. The minimum length is `1` character. The value must match regular expression `/^.+$/`.
		* `vpc_id` - (String) VPC ID this mount path belongs to. Present for VPC file volumes only; absent for Classic file volumes.
		  * Constraints: The maximum length is `1024` characters. The minimum length is `1` character. The value must match regular expression `/^.+$/`.
	* `throughput_mbps` - (Integer) Maximum throughput in Mbps. VPC block: from the VPC API. Classic and VPC file shares: absent if not applicable.
	  * Constraints: The maximum value is `102400`. The minimum value is `0`.
	* `zone` - (String) VPC zone (e.g. us-south-1).
	  * Constraints: The maximum length is `1024` characters. The minimum length is `1` character. The value must match regular expression `/^.+$/`.
* `storage_type` - (String) Storage type of the volume.
  * Constraints: Allowable values are: `block`, `file`, `san`, `local`.
* `workload_id` - (String) ID of the workload this volume is associated with.
  * Constraints: The maximum length is `1024` characters. The minimum length is `1` character. The value must match regular expression `/^.+$/`.

