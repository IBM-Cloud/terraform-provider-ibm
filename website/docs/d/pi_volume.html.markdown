---

subcategory: "Power Systems"
layout: "ibm"
page_title: "IBM: pi_volume"
description: |-
  Manages a volume in the Power Virtual Server cloud.
---

# ibm_pi_volume
Retrieves information about a persistent storage volume that is mounted to a Power Systems Virtual Server instance. For more information, about managin a volume, see [moving data to the cloud](https://cloud.ibm.com/docs/power-iaas?topic=power-iaas-moving-data-to-the-cloud).

## Example usage
The following example retrieves information about the `volume_1` volume that is mounted to the Power Systems Virtual Server instance with the ID.

```terraform
data "ibm_pi_volume" "ds_volume" {
  pi_volume_name       = "volume_1"
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

- `bootable` -  (Bool) If set to **true**, the Power Systems Virtual Server instance can boot from this volume. If set to **false**, this volume is not used during the boot process of the instance.
- `id` - (String) The unique identifier of the volume.
- `size` - (Integer) The size of the volume in gigabytes.
- `state` - (String) The state of the volume.
- `type` - (String) The disk type that is used for the volume.
- `wwn` - (String) The world wide name of the volume.
