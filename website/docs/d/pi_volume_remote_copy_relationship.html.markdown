---
subcategory: "Power Systems"
layout: "ibm"
page_title: "IBM: pi_volume_remote_copy_relationship"
description: |-
  Manages a remote copy relationship of a volume in the Power Virtual Server cloud.
---

# ibm_pi_volume_remote_copy_relationship
Retrieves information about remote copy relationship of a volume. For more information, about managing a volume group, see [moving data to the cloud](https://cloud.ibm.com/docs/power-iaas?topic=power-iaas-moving-data-to-the-cloud).

## Example usage
The following example retrieves information about about remote copy relationship of a volume in Power Systems Virtual Server.

```terraform
data "ibm_pi_volume_remote_copy_relationship" "ds_volume_remote_copy_relationships" {
  pi_volume_id         = "810b5fde-e054-4577-ab5e-3f866a1f6f60"
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
- `pi_volume_id` - (Required, String) The ID of the volume for which you want to retrieve detailed information.

## Attribute reference
In addition to all argument reference list, you can access the following attribute references after your data source is created. 

- `auxiliary_changed_volume_name` - (String) The name of the volume that is acting as the auxiliary change volume for the relationship.
- `auxiliary_volume_name` - (String) The auxiliary volume name at storage host level.
- `consistency_group_name` - (String) The consistency group name if volume is a part of volume group.
- `copy_type` (String) The copy type.
- `cycling_mode` - (String) The type of cycling mode used.
- `cycle_period_seconds` - (Integer) The minimum period in seconds between multiple cycles.
- `freeze_time` - (String) The freeze time of remote copy relationship.
- `id` - (String) The unique identifier of the volume.
- `master_changed_volume_name` (String) The name of the volume that is acting as the master change volume for the relationship.
- `master_volume_name` - (String) The master volume name at storage host level.
- `name` - (String) The remote copy relationship name.
- `primary_role` (String) Indicates whether master/aux volume is playing the primary role.
- `progress` - (Integer) The relationship progress.
- `remote_copy_id` - (String) The remote copy relationship ID.
- `state` - (String) The relationship state.
- `synchronized` - (String) Indicates whether the relationship is synchronized.
