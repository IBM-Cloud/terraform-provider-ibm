---
subcategory: "Power Systems"
layout: "ibm"
page_title: "IBM: pi_volume_clone"
description: |-
   Manages IBM Volume Clone in the Power Virtual Server cloud.
---

# ibm_pi_volume_clone

Create a volume clone. For more information, about managing volume clone, see [getting started with IBM Power Systems Virtual Servers](https://cloud.ibm.com/docs/power-iaas?topic=power-iaas-getting-started).

## Example Usage

The following example creates a volume clone.

```terraform
resource "ibm_pi_volume_clone" "testacc_volume_clone" {
  pi_cloud_instance_id    = "<value of the cloud_instance_id>"
  pi_volume_clone_name    = "test-volume-clone"
  pi_volume_ids           = ["<Volume ID>"]
  pi_target_storage_tier  = "<storage tier>"
  pi_replication_enabled  = true
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

ibm_pi_volume_clone provides the following [timeouts](https://www.terraform.io/docs/language/resources/syntax.html) configuration options:

- **create** - (Default 15 minutes) Used for creating volume clone.
- **delete** - (Default 15 minutes) Used for deleting volume clone.

## Argument Reference

Review the argument references that you can specify for your resource.

- `pi_cloud_instance_id` - (Required, String) The GUID of the service instance associated with an account.
- `pi_replication_enabled` - (Optional, Boolean) Indicates whether the cloned volume should have replication enabled. If no value is provided, it will default to the replication status of the source volume(s).
- `pi_target_storage_tier` - (Optional, String) The storage tier for the cloned volume(s). To get a list of available storage tiers, please use the [ibm_pi_storage_types_capacity](https://registry.terraform.io/providers/IBM-Cloud/ibm/latest/docs/data-sources/pi_storage_types_capacity) data source.
- `pi_user_tags` - (Optional, List) The user tags attached to this resource.
- `pi_volume_clone_name` - (Required, String) The base name of the newly cloned volume(s).
- `pi_volume_ids` - (Required, Set of String) List of volumes to be cloned.

## Attribute Reference

In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `cloned_volumes` - (List of objects) The List of cloned volumes.
  
  Nested scheme for `cloned_volumes`:
  - `clone_volume_id` - (String) The ID of the newly cloned volume.
  - `source_volume_id` - (String) The ID of the source volume.
- `failure_reason` - (String) The reason for the failure of the volume clone task.
- `id` - (String) The unique identifier of the volume clone. The ID is composed of `<pi_cloud_instance_id>/<task_id>`.
- `percent_complete` - (Integer) The completion percentage of the volume clone task.
- `status` - (String) The status of the volume clone task.
- `task_id` - (String) The ID of the volume clone task.

## Import

The `ibm_pi_volume_clone` resource can be imported by using `pi_cloud_instance_id` and `task_id`.

### Example

```bash
terraform import ibm_pi_volume_clone.example d7bec597-4726-451f-8a63-e62e6f19c32c/cea6651a-bc0a-4438-9f8a-a0770bbf3ebb
```
