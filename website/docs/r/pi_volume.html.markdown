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

- **create** - (Default 60 minutes) Used for creating volume.
- **delete** - (Default 60 minutes) Used for deleting volume.

## Argument reference 
Review the argument references that you can specify for your resource. 
- `pi_affinity_instance` - (Optional, String) PVM Instance (ID or Name) to base volume affinity policy against. This argument is Required if `pi_affinity_policy` is provided and Conflicts with `pi_affinity_volume`.
- `pi_affinity_policy` - (Optional, String) Affinity policy for data volume being created; requires `pi_affinity_instance` or `pi_affinity_volume` to be specified; Allowable values: `affinity`, `anti-affinity`
- `pi_affinity_volume`- (Optional, String) Volume (ID or Name) to base volume affinity policy against; required if pi_affinity_policy is provided and Conflicts with `pi_affinity_instance`.
- `pi_cloud_instance_id` - (Required, String) The GUID of the service instance associated with an account.
- `pi_volume_name` - (Required, String) The name of the volume.
- `pi_volume_shareable` - (Required, Bool) If set to **true**, the volume can be shared across Power Systems Virtual Server instances. If set to **false**, you can attach it only to one instance. 
- `pi_volume_size`  - (Required, Integer) The size of the volume in gigabytes. 
- `pi_volume_type` - (Optional, String) The type of volume that you want to create. Supported values are `ssd`, `standard`, `tier1`, and `tier3`.

## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `id` - (String) The unique identifier of the volume. The ID is composed of `<power_instance_id>/<volume_id>`. 
- `status` - (String) The status of the volume. 
- `volume_id` - (String) The unique identifier of the volume.
- `wwn` - (String) The world wide name of the volume.

## Import

The `ibm_pi_volume` resource can be imported by using `power_instance_id` and `volume_id`.

**Example**

```
$ terraform import ibm_pi_volume.example d7bec597-4726-451f-8a63-e62e6f19c32c/cea6651a-bc0a-4438-9f8a-a0770bbf3ebb
```
