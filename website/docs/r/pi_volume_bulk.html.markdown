---
subcategory: "Power Systems"
layout: "ibm"
page_title: "IBM: pi_volume_bulk"
description: |-
  Manages IBM Volume in the Power Virtual Server cloud.
---

# ibm_pi_volume_bulk

Create, update, or delete a volume to attach it to a Power Systems Virtual Server instance. For more information, about managing volume, see [cloning a volume](https://cloud.ibm.com/docs/power-iaas?topic=power-iaas-snapshots-cloning).

## Example usage

The following example creates 3 20 GB volumes.

```terraform
resource "ibm_pi_volume_bulk" "testacc_volume"{
  pi_cloud_instance_id = "<value of the cloud_instance_id>"
  pi_count             = 3
  pi_volume_size       = 20
  pi_volume_name       = "test-volume"
  pi_volume_type       = "tier3"
  pi_volume_shareable  = true
}
```

### Notes

- Please find [supported Regions](https://cloud.ibm.com/apidocs/power-cloud#endpoint) for endpoints.
- If a Power cloud instance is provisioned at `lon04`, The provider level attributes should be as follows:
  - `region` - `lon`
  - `zone` - `lon04`

Example usage:

  ```terraform
    provider "ibm" {
      region    =   "lon"
      zone      =   "lon04"
    }
  ```
  
## Timeouts

ibm_pi_volume_bulk provides the following [timeouts](https://www.terraform.io/docs/language/resources/syntax.html) configuration options:

- **create** - (Default 30 minutes) Used for creating volume.
- **delete** - (Default 30 minutes) Used for deleting volume.

## Argument reference

Review the argument references that you can specify for your resource.

- `pi_affinity_instance` - (Optional, String) PVM Instance (ID or Name) to base volume affinity policy against; required if requesting `affinity` and `pi_affinity_volume` is not provided.
- `pi_affinity_policy` - (Optional, String) Affinity policy for data volume being created; ignored if `pi_volume_pool` provided; for policy 'affinity' requires one of `pi_affinity_instance` or `pi_affinity_volume` to be specified; for policy 'anti-affinity' requires one of `pi_anti_affinity_instances` or `pi_anti_affinity_volumes` to be specified; Allowable values: `affinity`, `anti-affinity`.
- `pi_affinity_volume`- (Optional, String) Volume (ID or Name) to base volume affinity policy against; required if requesting `affinity` and `pi_affinity_instance` is not provided.
- `pi_anti_affinity_instances` - (Optional, String) List of pvmInstances to base volume anti-affinity policy against; required if requesting `anti-affinity` and `pi_anti_affinity_volumes` is not provided.
- `pi_anti_affinity_volumes`- (Optional, String) List of volumes to base volume anti-affinity policy against; required if requesting `anti-affinity` and `pi_anti_affinity_instances` is not provided.
- `pi_cloud_instance_id` - (Required, String) The GUID of the service instance associated with an account.
- `pi_count` - (Optional, Integer) Number of volumes to create. Default 1. Maximum is 500 for public workspaces, and 250 for private workspaces.
- `pi_replication_enabled` - (Optional, Boolean) Indicates if the volume should be replication enabled or not.

  **Note:** `replication_sites` will be populated automatically with default sites if set to true and sites are not specified.

- `pi_replication_sites` - (Optional, List) List of replication sites for volume replication. Must set `pi_replication_enabled` to true to use.
- `pi_user_tags` - (Optional, List) The user tags attached to this resource.
- `pi_volume_name` - (Required, String) The name of the volume.
- `pi_volume_pool` - (Optional, String) Volume pool where the volume will be created; if provided then `pi_affinity_policy` values will be ignored.
- `pi_volume_shareable` - (Required, Boolean) If set to **true**, the volume can be shared across Power Systems Virtual Server instances. If set to **false**, you can attach it only to one instance.
- `pi_volume_size`  - (Required, Integer) The size of the volume in GB.
- `pi_volume_type` - (Optional, String) Type of volume, if this field is not provided, it will default to `tier3`. To get a list of available volume types, please use the [ibm_pi_storage_types_capacity](https://registry.terraform.io/providers/IBM-Cloud/ibm/latest/docs/data-sources/pi_storage_types_capacity) data source.

## Attribute reference

In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `id` - (String) The unique identifier of the bulk volume resource. The ID is composed of `<cloud_instance_id>/<volume_id_1>/.../<volume_id_n>`.
- `volumes` - (List) List of volumes to create.

  - `auxiliary` - (Boolean) Indicates if the volume is auxiliary or not.
  - `auxiliary_volume_name` - (String) The auxiliary volume name.
  - `consistency_group_name` - (String) The consistency group name if volume is a part of volume group.
  - `crn` - (String) The CRN of this resource.
  - `delete_on_termination` - (Boolean) Indicates if the volume should be deleted when the server terminates.
  - `group_id` - (String) The volume group id to which volume belongs.
  - `io_throttle_rate` - (String) Amount of iops assigned to the volume.
  - `master_volume_name` - (String) The master volume name.
  - `mirroring_state` - (String) Mirroring state for replication enabled volume.
  - `primary_role` - (String) Indicates whether `master`/`auxiliary` volume is playing the primary role.
  - `replication_status` - (String) The replication status of the volume.
  - `replication_sites` - (List) List of replication sites for volume replication.
  - `replication_type` - (String) The replication type of the volume `metro` or `global`.
  - `volume_id` - (String) The unique identifier of the volume.
  - `volume_status` - (String) The status of the volume.
  - `wwn` - (String) The world wide name of the volume.

## Import

The `ibm_pi_volume_bulk` resource can be imported by using `pi_cloud_instance_id` and `volume_id`.

### Example

```bash
terraform import ibm_pi_volume_bulk.example d7bec597-4726-451f-8a63-e62e6f19c32c/cea6651a-bc0a-4438-9f8a-a0770bbf3ebb
```
