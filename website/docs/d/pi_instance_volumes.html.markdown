---

subcategory: "Power Systems"
layout: "ibm"
page_title: "IBM: pi_instance_volumes"
description: |-
  Manages Instance volumes in the Power Virtual Server cloud.
---

# ibm_pi_instance_volumes
Retrieves information about a persistent storage volume that is mounted to a Power Systems Virtual Server instance. For more information, about power instance volume, see [snapshotting, cloning, and restoring](https://cloud.ibm.com/docs/power-iaas?topic=power-iaas-volume-snapshot-clone).

## Example usage
The following example retrieves information about the `volume_1` volume that is mounted to the Power Systems Virtual Server instance with the ID.

```terraform
data "ibm_pi_instance_volumes" "ds_volumes" {
  pi_instance_name     = "volume_1"
  pi_cloud_instance_id = "49fba6c9-23f8-40bc-9899-aca322ee7d5b"
}
```

**Notes**
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

## Argument reference
Review the argument references that you can specify for your data source. 

- `pi_cloud_instance_id` - (Required, String) The GUID of the service instance associated with an account.
- `pi_volume_name` - (Required, String) The name of the volume for which you want to retrieve detailed information.

## Attribute reference
In addition to all argument reference list, you can access the following attribute references after your data source is created. 

- `boot_volume_id` - (String) The unique identifier of the boot volume.
- `instance_volumes` - List of volumes - List of volumes attached to instance.

  Nested scheme for `instance_volumes`:
    - `bootable`- (Bool) Indicates if the volume is bootable (**true**) or not (**false**).
	- `href` - (String) The hyper link of the volume.
	- `id` - (String) The unique identifier of the volume.
	- `name` - (String) The name of the volume.
	- `shareable` -  (Bool) If set to **true**, the volume can be shared across multiple Power Systems Virtual Server instances. If set to **false**, the volume can be mounted to one instance only.
	- `size` - (Integer) The size of this volume in gigabytes.
	- `state` - (String) The state of the volume.
	- `type` - (String) The disk type that is used for this volume.
