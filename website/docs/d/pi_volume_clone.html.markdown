---
subcategory: "Power Systems"
layout: "ibm"
page_title: "IBM: pi_volume_clone"
description: |-
   Manages IBM Volume Clone in the Power Virtual Server cloud.
---

# ibm_pi_volume_clone

Retrieves information about a volume clone. For more information, about managing volume clone, see [getting started with IBM Power Systems Virtual Servers](https://cloud.ibm.com/docs/power-iaas?topic=power-iaas-getting-started).

## Example usage

The following example retrieves information about the volume clone task that is present in Power Systems Virtual Server.

```terraform
data "ibm_pi_volume_clone" "ds_volume_clone" {
  pi_cloud_instance_id        = "<value of the cloud_instance_id>"
  pi_volume_clone_task_id     = "<clone task id>"
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

## Argument reference

Review the argument references that you can specify for your resource.

- `pi_cloud_instance_id` - (Required, String) The GUID of the service instance associated with an account.
- `pi_volume_clone_task_id` - (Required, String) The ID of the volume clone task.

## Attribute reference

In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `cloned_volumes` - (List of objects) The List of cloned volumes.
  
  Nested scheme for `cloned_volumes`:
  - `clone_volume_id` - (String) The ID of the newly cloned volume.
  - `source_volume_id` - (String) The ID of the source volume.
- `failure_reason` - (String) The reason for the failure of the clone volume task.
- `id` - (String) The unique identifier of the volume clone task.
- `percent_complete` - (Integer) The completion percentage of the volume clone task.
- `status` - (String) The status of the volume clone task.
