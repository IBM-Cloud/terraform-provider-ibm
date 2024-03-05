---

subcategory: "Power Systems"
layout: "ibm"
page_title: "IBM: pi_snapshot"
description: |-
  Manages snapshots in the Power Virtual Server cloud.
---

# ibm_pi_snapshot
Creates, updates, deletes, and manages snapshots in the Power Virtual Server Cloud. For more information, about snapshots in the Power Virutal Server, see [snapshotting, cloning, and restoring](https://cloud.ibm.com/docs/power-iaas?topic=power-iaas-volume-snapshot-clone).

## Example usage
The following example enables you to create a snapshot:

```terraform
resource "ibm_pi_snapshot" "testacc_snapshot"{
  pi_instance_name       = test-instance
  pi_snap_shot_name       = test-snapshot
  pi_volume_ids       = ["volumeid1","volumeid2"]
  description  = "Testing snapshot for instance"
  pi_cloud_instance_id = "<value of the cloud_instance_id>"
}
```

**Note**
* Please find [supported Regions](https://cloud.ibm.com/apidocs/power-cloud#endpoint) for endpoints.
* If a Power cloud instance is provisioned at `lon04`, The provider level attributes should be as follows:
  * `region` - `lon`
  * `zone` - `lon04`

  Example usage:
  
  ```terraform
    provider "ibm" {
      region    =   "lon"
      zone      =   "lon04"
    }
  ```
  
## Timeouts

The `ibm_pi_snapshot` provides the following [Timeouts](https://www.terraform.io/docs/language/resources/syntax.html) configuration options:

- **create** - (Default 60 minutes) Used for Creating snapshot.
- **update** - (Default 60 minutes) Used for Updating snapshot.
- **delete** - (Default 10 minutes) Used for Deleting snapshot.

## Argument reference
Review the argument references that you can specify for your resource.
 
- `pi_instance_name` - (Required, String) The name of the instance you want to take a snapshot of.
- `pi_snapshot_name` - (Required, String) The unique name of the snapshot.
- `pi_description` - (Optional, String) Description of the PVM instance snapshot.
- `pi_volume_ids` - (Optional, String) The volume ID, if none provided then all volumes of the instance will be part of the snapshot.
- `pi_cloud_instance_id` - (Required, String) The GUID of the service instance associated with an account.

## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `id` - (String) The unique identifier of the snapshot. The ID is composed of <power_instance_id>/<snapshot_id>.
- `snapshot_id` - (String) ID of the PVM instance snapshot.
- `status` - (String) Status of the PVM instance snapshot.
- `creation_date` - (String) Creation Date.
- `last_update_date` - (String) Last Update Date.
- `volume_snapshots` - (String) A map of volume snapshots included in the PVM instance snapshot.

## Import

The `ibm_pi_snapshot` resource can be imported by using `power_instance_id` and `snapshot_id`.

**Example**

```
$ terraform import ibm_pi_snapshot.example d7bec597-4726-451f-8a63-e62e6f19c32c/cea6651a-bc0a-4438-9f8a-a0770bbf3ebb
```
