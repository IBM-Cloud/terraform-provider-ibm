---

subcategory: "Power Systems"
layout: "ibm"
page_title: "IBM: pi_snapshot"
description: |-
  Manages Snapshots in the Power Virtual Server Cloud.
---

# ibm\_pi_snapshot

Provides a snapshot resource. This allows snapshots to be created, updated, and deleted in the Power Virtual Server Cloud.

## Example Usage

In the following example, you can create a snapshot:

```hcl
resource "ibm_pi_snapshot" "testacc_snapshot"{
  pi_instance_name       = test-instance
  pi_snap_shot_name       = test-snapshot
  pi_volume_ids       = ["volumeid1","volumeid2"]
  description  = "Testing snapshot for instance"
  pi_cloud_instance_id = "<value of the cloud_instance_id>"
}
```
## Notes:
* Please find [supported Regions](https://cloud.ibm.com/apidocs/power-cloud#endpoint) for endpoints.
* If a Power cloud instance is provisioned at `lon04`, The provider level attributes should be as follows:
  * `region` - `lon`
  * `zone` - `lon04`
  Example Usage:
  ```hcl
    provider "ibm" {
      region    =   "lon"
      zone      =   "lon04"
    }
  ```
## Timeouts

ibm_pi_snapshot provides the following [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) configuration options:

* `create` - (Default 60 minutes) Used for Creating snapshot.
* `delete` - (Default 60 minutes) Used for Deleting snapshot.
* `update` - (Default 60 minutes) Used for Updating snapshot.

## Argument Reference

The following arguments are supported:

* `pi_instance_name` - (Required, string) The instance name that we want to take a snapshot of.
* `pi_snapshot_name` - (Required, string) The name of this snasphot ( make sure it's unique) .
* `description` - (Optional, string) Description of the snapshot.
* `pi_volume_ids` - (Optional, List) String of volumeids. If none provided then all volumes of the instance
will be part of the snapshot.
* `pi_cloud_instance_id` - (Required, string) The GUID of the service instance associated with the account

## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier of the snapshot.The id is composed of \<power_instance_id\>/\<pi_snap_shot_id\>.
* `volume_snapshots` - A map of the source and target volumes.


## Import

ibm_pi_snapshot can be imported using `power_instance_id` and `pi_snap_shot_id`, eg

```
$ terraform import ibm_pi_snapshot.example d7bec597-4726-451f-8a63-e62e6f19c32c/cea6651a-bc0a-4438-9f8a-a0770bbf3ebb
```