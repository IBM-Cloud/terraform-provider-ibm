---
subcategory: "Power Systems"
layout: "ibm"
page_title: "IBM: pi_volume_groups_details"
description: |-
  Manages a volume groups in the Power Virtual Server cloud.
---

# ibm_pi_volume_groups_details

Retrieves information about all volume groups with details. about managing a volume group, see [moving data to the cloud](https://cloud.ibm.com/docs/power-iaas?topic=power-iaas-moving-data-to-the-cloud).

## Example Usage

The following example retrieves information about the volume groups with details that is present in Power Systems Virtual Server.

```terraform
data "ibm_pi_volume_groups_details" "ds_volume_groups_details" {
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

- `volume_groups` - (List) List of all volume group.
  
  Nested scheme for `volume_groups`:
  - `auxiliary` - (Boolean) Indicates if the volume group is auxiliary.
  - `consistency_group_name` - (String) The name of consistency group at storage controller level.
  - `id` - (String) The unique identifier of the volume group.
  - `replication_status` - (String) The replication status of volume group.
  - `replication_sites` - (List) Indicates the replication sites of the volume group.
  - `status` - (String) The status of the volume group.
  - `status_description_errors` - (List) The status details of the volume group.

      Nested scheme for `status_description_errors`:
        - `key` - (String) The volume group error key.
        - `message` - (String) The failure message providing more details about the error key.
        - `volume_ids` - (List) List of volume IDs, which failed to be added/removed to/from the volume group, with the given error.
  - `storage_pool` - (String) Storage pool of the volume group.
  - `volume_group_name` - (String) The name of the volume group.
  - `volume_ids` - (List) List of volume IDs, member of volume group.
