---
subcategory: "Power Systems"
layout: "ibm"
page_title: "IBM: pi_volume_group"
description: |-
  Manages IBM Volume Group in the Power Virtual Server cloud.
---

# ibm_pi_volume_group

Create, update, or delete a volume group. For more information, about managing volume groups, see [getting started with IBM Power Systems Virtual Servers](https://cloud.ibm.com/docs/power-iaas?topic=power-iaas-getting-started).

## Example Usage

The following example creates a volume group.

```terraform
resource "ibm_pi_volume_group" "testacc_volume_group"{
  pi_cloud_instance_id = "<value of the cloud_instance_id>"
  pi_volume_group_name = "test-volume-group"
  pi_volume_ids        = ["<Volume ID>"]
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

ibm_pi_volume_group provides the following [timeouts](https://www.terraform.io/docs/language/resources/syntax.html) configuration options:

- **create** - (Default 30 minutes) Used for creating volume group.
- **update** - (Default 30 minutes) Used for updating volume group.
- **delete** - (Default 10 minutes) Used for deleting volume group.

## Argument Reference

Review the argument references that you can specify for your resource.

- `pi_cloud_instance_id` - (Required, String) The GUID of the service instance associated with an account.
- `pi_consistency_group_name` - (Optional, String) The name of consistency group at storage controller level, required if `pi_volume_group_name` is not provided.
- `pi_volume_group_name` - (Optional, String) The name of the volume group, required if `pi_consistency_group_name` is not provided.
- `pi_volume_ids` - (Required, Set of String) List of volume IDs to add in volume group.

## Attribute Reference

In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `auxiliary` - (Boolean) Indicates if the volume group is auxiliary.
- `consistency_group_name` - (String) The consistency Group Name if volume is a part of volume group.
- `id` - (String) The unique identifier of the volume group. The ID is composed of `<pi_cloud_instance_id>/<volume_group_id>`.
- `replication_sites` - (List) Indicates the replication sites of the volume group.
- `replication_status` - (String) The replication status of volume group.
- `status_description_errors` - (Set) The status details of the volume group.
  
  Nested scheme for `status_description_errors`:
  - `key` - (String) The volume group error key.
  - `message` - (String) The failure message providing more details about the error key.
  - `volume_ids` - (List of String) List of volume IDs, which failed to be added to or removed from the volume group, with the given error.
- `storage_pool` - (String) Storage pool of the volume group.
- `volume_group_id` - (String) The unique identifier of the volume group.
- `volume_group_status` - (String) The status of the volume group.

## Import

The `ibm_pi_volume_group` resource can be imported by using `pi_cloud_instance_id` and `volume_group_id`.

### Example

```bash
terraform import ibm_pi_volume_group.example d7bec597-4726-451f-8a63-e62e6f19c32c/cea6651a-bc0a-4438-9f8a-a0770bbf3ebb
```
