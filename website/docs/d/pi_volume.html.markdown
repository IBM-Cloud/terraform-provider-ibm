---
subcategory: "Power Systems"
layout: "ibm"
page_title: "IBM: pi_volume"
description: |-
  Manages a volume in the Power Virtual Server cloud.
---

# ibm_pi_volume
Retrieves information about a persistent storage volume that is mounted to a Power Systems Virtual Server instance. For more information, about managin a volume, see [moving data to the cloud](https://cloud.ibm.com/docs/power-iaas?topic=power-iaas-moving-data-to-the-cloud).

## Example usage
The following example retrieves information about the `volume_1` volume that is mounted to the Power Systems Virtual Server instance with the ID.

```terraform
data "ibm_pi_volume" "ds_volume" {
  pi_volume_name       = "volume_1"
  pi_cloud_instance_id = "49fba6c9-23f8-40bc-9899-aca322ee7d5b"
}
```

**Notes**
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
  
## Argument reference
Review the argument references that you can specify for your data source. 

- `pi_cloud_instance_id` - (Required, String) The GUID of the service instance associated with an account.
- `pi_volume_name` - (Required, String) The name of the volume for which you want to retrieve detailed information.

## Attribute reference
In addition to all argument reference list, you can access the following attribute references after your data source is created. 

- `auxiliary` - (Boolean) Indicates if the volume is auxiliary or not.
- `auxiliary_volume_name` - (String) The auxiliary volume name.
- `bootable` -  (Boolean) Indicates if the volume is boot capable.
- `consistency_group_name` - (String) Consistency group name if volume is a part of volume group.
- `disk_type` - (String) The disk type that is used for the volume.
- `group_id` - (String) The volume group id in which the volume belongs.
- `id` - (String) The unique identifier of the volume.
- `io_throttle_rate` - (String) Amount of iops assigned to the volume.
- `master_volume_name` - (String) The master volume name.
- `mirroring_state` - (String) Mirroring state for replication enabled volume.
- `primary_role` - (String) Indicates whether `master`/`auxiliary` volume is playing the primary role.
- `replication_enabled` - (Boolean) Indicates if the volume should be replication enabled or not.
- `replication_status` - (String) The replication status of the volume.
- `replication_type` - (String) The replication type of the volume, `metro` or `global`.
- `shareable` - (String) Indicates if the volume is shareable between VMs. 
- `size` - (Integer) The size of the volume in gigabytes.
- `state` - (String) The state of the volume.
- `volume_pool` - (String) Volume pool, name of storage pool where the volume is located.
- `wwn` - (String) The world wide name of the volume.
