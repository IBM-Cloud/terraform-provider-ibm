---
subcategory: "Power Systems"
layout: "ibm"
page_title: "IBM: ibm_pi_volumes"
description: |-
  Manages volumes in the Power Virtual Server cloud.
---

# ibm_pi_volumes

Retrieves information about all persistent storage volumes that in a Power Systems Virtual Server workspace. For more information, about managing volumes, see [moving data to the cloud](https://cloud.ibm.com/docs/power-iaas?topic=power-iaas-moving-data-to-the-cloud).

## Example Usage

The following example retrieves information about all volumes present in a Power Systems Virtual Server workspace.

```terraform
data "ibm_pi_volumes" "ds_volume" {
  pi_cloud_instance_id = "49fba6c9-23f8-40bc-9899-aca322ee7d5b"
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

## Attribute Reference

In addition to all argument reference list, you can access the following attribute references after your data source is created.

- `volumes` - (List) The list of volumes.

  Nested schema for `volumes`:
  - `auxiliary_volume_name` - (String) The auxiliary volume name.
  - `auxiliary` - (Boolean) Indicates if the volume is auxiliary.
  - `bootable` -  (Boolean) Indicates if the volume is boot capable.
  - `consistency_group_name` - (String) Consistency group name if volume is a part of volume group.
  - `creation_date` - (String) Date of volume creation.
  - `crn` - (String) The CRN of this resource.
  - `disk_type` - (String) The disk type that is used for the volume.
  - `freeze_time` - (String) Time of remote copy relationship.
  - `group_id` - (String) The volume group id in which the volume belongs.
  - `id` - (String) The unique identifier of the volume.
  - `io_throttle_rate` - (String) Amount of iops assigned to the volume.
  - `last_update_date` - (String) The date when the volume last updated.
  - `master_volume_name` - (String) The master volume name.
  - `mirroring_state` - (String) Mirroring state for replication enabled volume.
  - `name` - (String) The name of the volume.
  - `out_of_band_deleted` - (Bool) Indicates if the volume does not exist on storage controller.
  - `primary_role` - (String) Indicates whether `master`/`auxiliary` volume is playing the primary role.
  - `replication_enabled` - (Boolean) Indicates if the volume should be replication enabled or not.
  - `replication_sites` - (List) List of replication sites for volume replication.
  - `replication_status` - (String) The replication status of the volume.
  - `replication_type` - (String) The replication type of the volume, `metro` or `global`.
  - `shareable` - (String) Indicates if the volume is shareable between VMs.
  - `size` - (Integer) The size of the volume in GB.
  - `state` - (String) The state of the volume.
  - `user_tags` - (List) List of user tags attached to the resource.
  - `volume_pool` - (String) The name of storage pool where the volume is located.
  - `volume_type` - (String) The name of storage template used to create the volume.
  - `wwn` - (String) The world wide name of the volume.
