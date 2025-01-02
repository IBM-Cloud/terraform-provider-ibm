---
subcategory: "Power Systems"
layout: "ibm"
page_title: "IBM: ibm_pi_volume_group_action"
description: |-
  Manages IBM Volume Group Action in the Power Virtual Server cloud.
---

# ibm_pi_volume_group_action

Perfoms action on a volume group. For more information, about managing volume, see [getting started with IBM Power Systems Virtual Servers](https://cloud.ibm.com/docs/power-iaas?topic=power-iaas-getting-started).

## Example usage

The following example attaches volume to a power systems virtual server instance.

```terraform
  resource "ibm_pi_volume_group_action" "testacc_volume_group_action" {
    pi_cloud_instance_id = "<value of the cloud_instance_id>"
    pi_volume_group_id = "<id of the volume group>"
    pi_volume_group_action {
      stop {
      access = true
      }
    }
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

ibm_pi_volume_group_action provides the following [timeouts](https://www.terraform.io/docs/language/resources/syntax.html) configuration options:

* **create** - (Default 15 minutes) Used for performing action on volume group.
* **delete** - (Default 15 minutes) Used for deleting volume group action resource.

## Argument reference

Review the argument references that you can specify for your resource.

* `pi_cloud_instance_id` - (Required, Forces new resource, String) The GUID of the service instance associated with an account.
* `pi_volume_group_action` - (Required, Forces new resource, List) Performs an action (`start` / `stop` / `reset`) on a volume group(one at a time).
  * Constraints: The maximum length is `1` items. The minimum length is `1` items.
  Nested scheme for `pi_volume_group_action`:
    * `reset` - (Optional, Forces new resource, List) Performs reset action on the volume group to update its status value.
      * Constraints: The maximum length is `1` items.
      Nested scheme for `reset`:
        * `status` - (Required, String) New status to be set for a volume group.
    * `start` - (Optional, Forces new resource, List) Performs start action on a volume group.
      * Constraints: The maximum length is `1` items.
      Nested scheme for `start`:
        * `source` - (Required, String) Indicates the source of the action `master` or `aux`.
    * `stop` - (Optional, Forces new resource, List) Performs stop action on a volume group.
      * Constraints: The maximum length is `1` items.
      Nested scheme for `reset`:
        * `access` - (Required, Boolean) Indicates the access mode of aux volumes.
* `pi_volume_group_id` - (Required, Forces new resource, String) The ID of volume group on which action is to performed.

## Attribute reference

In addition to all argument reference list, you can access the following attribute reference after your resource is created.

* `id` - (String) The unique identifier of the volume group action. The ID is composed of `<pi_cloud_instance_id>/<volume_group_id>`.
* `replication_status` - (String) The replication status of volume group.
* `volume_group_name` - (String) The name of the volume group.
* `volume_group_status` - (String) The status of the volume group.

## Import

The `ibm_pi_volume_group_action` resource can be imported by using `pi_cloud_instance_id` and `volume_group_id`.

### Example

```bash
terraform import ibm_pi_volume_group_action.example d7bec597-4726-451f-8a63-e62e6f19c32c/49fba6c9-23f8-40bc-9899-aca322ee7d5b
```
