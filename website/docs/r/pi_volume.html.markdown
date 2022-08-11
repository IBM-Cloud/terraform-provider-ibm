---

subcategory: "Power Systems"
layout: "ibm"
page_title: "IBM: pi_volume"
description: |-
  Manages IBM Volume in the Power Virtual Server cloud.
---

# ibm_pi_volume
Create, update, or delete a volume to attach it to a Power Systems Virtual Server instance. For more information, about managing volume, see [cloning a volume](https://cloud.ibm.com/docs/power-iaas?topic=power-iaas-volume-snapshot-clone#cloning-volume).

## Example usage
The following example creates a 20 GB volume.

```terraform
resource "ibm_pi_volume" "testacc_volume"{
  pi_volume_size       = 20
  pi_volume_name       = "test-volume"
  pi_volume_type       = "ssd"
  pi_volume_shareable  = true
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

ibm_pi_volume provides the following [timeouts](https://www.terraform.io/docs/language/resources/syntax.html) configuration options:

- **create** - (Default 30 minutes) Used for creating volume.
- **update** - (Default 30 minutes) Used for updating volume.
- **delete** - (Default 10 minutes) Used for deleting volume.

## Argument reference 
Review the argument references that you can specify for your resource. 

- `pi_affinity_instance` - (Optional, String) PVM Instance (ID or Name) to base volume affinity policy against; required if requesting `affinity` and `pi_affinity_volume` is not provided.
- `pi_affinity_policy` - (Optional, String) Affinity policy for data volume being created; ignored if `pi_volume_pool` provided; for policy 'affinity' requires one of `pi_affinity_instance` or `pi_affinity_volume` to be specified; for policy 'anti-affinity' requires one of `pi_anti_affinity_instances` or `pi_anti_affinity_volumes` to be specified; Allowable values: `affinity`, `anti-affinity`
- `pi_affinity_volume`- (Optional, String) Volume (ID or Name) to base volume affinity policy against; required if requesting `affinity` and `pi_affinity_instance` is not provided.
- `pi_anti_affinity_instances` - (Optional, String) List of pvmInstances to base volume anti-affinity policy against; required if requesting `anti-affinity` and `pi_anti_affinity_volumes` is not provided.
- `pi_anti_affinity_volumes`- (Optional, String) List of volumes to base volume anti-affinity policy against; required if requesting `anti-affinity` and `pi_anti_affinity_instances` is not provided.
- `pi_cloud_instance_id` - (Required, String) The GUID of the service instance associated with an account.
- `pi_volume_name` - (Required, String) The name of the volume.
- `pi_volume_pool` - (Optional, String) Volume pool where the volume will be created; if provided then `pi_volume_type` and `pi_affinity_policy` values will be ignored.
- `pi_volume_shareable` - (Required, Bool) If set to **true**, the volume can be shared across Power Systems Virtual Server instances. If set to **false**, you can attach it only to one instance. 
- `pi_volume_size`  - (Required, Integer) The size of the volume in gigabytes. 
- `pi_volume_type` - (Optional, String) Type of Disk, required if `pi_affinity_policy` and `pi_volume_pool` not provided, otherwise ignored. Supported values are `ssd`, `standard`, `tier1`, and `tier3`.

## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `delete_on_termination` - (Bool) Indicates if the volume should be deleted when the server terminates.
- `id` - (String) The unique identifier of the volume. The ID is composed of `<power_instance_id>/<volume_id>`.
- `volume_id` - (String) The unique identifier of the volume.
- `volume_status` - (String) The status of the volume.
- `wwn` - (String) The world wide name of the volume.

## Import

The `ibm_pi_volume` resource can be imported by using `power_instance_id` and `volume_id`.

**Example**

```
$ terraform import ibm_pi_volume.example d7bec597-4726-451f-8a63-e62e6f19c32c/cea6651a-bc0a-4438-9f8a-a0770bbf3ebb
```
