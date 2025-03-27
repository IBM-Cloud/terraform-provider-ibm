---
subcategory: "Power Systems"
layout: "ibm"
page_title: "IBM: pi_volume_group_storage_details"
description: |-
  Manages a volume group in the Power Virtual Server cloud.
---

# ibm_pi_volume_group_storage_details

Retrieves information about the storage details of a volume group. For more information, about managing a volume group, see [moving data to the cloud](https://cloud.ibm.com/docs/power-iaas?topic=power-iaas-moving-data-to-the-cloud).

## Example Usage

The following example retrieves information about the storage details of a volume group that is present in Power Systems Virtual Server.

```terraform
data "ibm_pi_volume_group_storage_details" "ds_volume_group_storage_details" {
  pi_volume_group_id   = "cf2ea8d3-cfc8-40e0-80c9-b096581be676"
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
- `pi_volume_group_id` - (Required, String) The ID of the volume group.

## Attribute Reference

In addition to all argument reference list, you can access the following attribute references after your data source is created.

- `consistency_group_name` - (String) The name of consistency group at storage controller level.
- `cycle_period_seconds` - (Integer) The minimum period in seconds between multiple cycles.
- `cycling_mode` - (String) The type of cycling mode used.
- `id` - (String) The unique identifier of the volume group.
- `number_of_volumes` - (Integer) The number of volumes in volume group.
- `primary_role` - (String) Indicates whether master/aux volume is playing the primary role.
- `remote_copy_relationship_names` - (List) List of remote-copy relationship names in a volume group.
- `replication_type` - (String) The type of replication (metro, global).
- `state` - (String) The relationship state.
- `synchronized` - (String) Indicates whether the relationship is synchronized.
