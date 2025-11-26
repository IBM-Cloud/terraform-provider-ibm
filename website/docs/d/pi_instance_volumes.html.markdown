---
subcategory: "Power Systems"
layout: "ibm"
page_title: "IBM: pi_instance_volumes"
description: |-
  Manages Instance volumes in the Power Virtual Server cloud.
---

# ibm_pi_instance_volumes

Retrieves information about the persistent storage volumes that are mounted to a Power Systems Virtual Server instance. For more information, about power instance volume, see [snapshotting, cloning, and restoring](https://cloud.ibm.com/docs/power-iaas?topic=power-iaas-volume-snapshot-clone).

## Example Usage

The following example retrieves information about the volumes attached to the `terraform-test-instance` instance.

```terraform
data "ibm_pi_instance_volumes" "ds_volumes" {
  pi_cloud_instance_id = "49fba6c9-23f8-40bc-9899-aca322ee7d5b"
  pi_instance_id       = "e6b579b7-d94b-42e5-a19d-5d1e0b2547c4"
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

## Argument Reference

Review the argument references that you can specify for your data source.

- `pi_cloud_instance_id` - (Required, String) The GUID of the service instance associated with an account.
- `pi_instance_id` - (Optional, String) The PVM instance ID.
- `pi_instance_name` - (Deprecated, Optional, String) The unique identifier or name of the instance. Passing the name of the instance could fail or fetch stale data. Please pass an id and use `pi_instance_id` instead.

## Attribute Reference

In addition to all argument reference list, you can access the following attribute references after your data source is created.

- `boot_volume_id` - (String) The unique identifier of the boot volume.
- `instance_volumes` - (List) List of volumes attached to instance.

  Nested scheme for `instance_volumes`:
  - `auxiliary_volume_name` - (String) The auxiliary volume name.
  - `auxiliary` - (Boolean) Indicates if the volume is auxiliary or not.
  - `bootable`- (Boolean) Indicates if the volume is boot capable.
  - `consistency_group_name` - (String) The consistency group name if volume is a part of volume group.
  - `creation_date` - (String) Date of volume creation.
  - `crn` - (String) The CRN of this resource.
  - `delete_on_termination` - (Boolean) Indicates if the volume should be deleted when the server terminates.
  - `freeze_time` - (String) Time of remote copy relationship.
  - `group_id` - (String) The volume group id to which volume belongs.
  - `href` - (String) The hyper link of the volume.
  - `id` - (String) The unique identifier of the volume.
  - `io_throttle_rate` - (String) Amount of iops assigned to the volume.
  - `last_update_date` - (String) The date when the volume last updated.
  - `master_volume_name` - (String) The master volume name.
  - `mirroring_state` - (String) Mirroring state for replication enabled volume.
  - `name` - (String) The name of the volume.
  - `out_of_band_deleted` - (Bool) Indicates if the volume does not exist on storage controller.
  - `pool` - (String) Volume pool, name of storage pool where the volume is located.
  - `primary_role` - (String) Indicates whether `master`/`auxiliary` volume is playing the primary role.
  - `replication_enabled` - (Boolean) Indicates whether replication is enabled on the volume.
  - `replication_sites` - (List) List of replication sites for volume replication.
  - `replication_status` - (String) The replication status of the volume.
  - `replication_type` - (String) The replication type of the volume, `metro` or `global`.
  - `shareable` - (Boolean) Indicates if the volume is shareable between VMs.
  - `size` - (Integer) The size of this volume in GB.
  - `state` - (String) The state of the volume.
  - `type` - (String) The disk type that is used for this volume.
  - `user_tags` - (List) List of user tags attached to the resource.
  - `volume_pool` - (String) Name of the storage pool where the volume is located.
  - `volume_type` - (String) Name of storage template used to create the volume.
  - `wwn` - (String) The world wide name of the volume.
