---
subcategory: "Power Systems"
layout: "ibm"
page_title: "IBM: pi_volume_flash_copy_mappings"
description: |-
  Manages a volume flash copy mappings in the Power Virtual Server cloud.
---

# ibm_pi_volume_flash_copy_mappings
Retrieves information about flash copy mappings of a volume. For more information, about managing a volume group, see [moving data to the cloud](https://cloud.ibm.com/docs/power-iaas?topic=power-iaas-moving-data-to-the-cloud).

## Example usage
The following example retrieves information about flash copy mappings of a volume in Power Systems Virtual Server.

```terraform
data "pi_volume_flash_copy_mappings" "ds_volume_flash_copy_mappings" {
  pi_volume_id         = "810b5fde-e054-4577-ab5e-3f866a1f6f66"
  pi_cloud_instance_id = "49fba6c9-23f8-40bc-9899-aca322ee7d5b"
}
```

**Notes**
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
Review the argument references that you can specify for your data source. 

- `pi_cloud_instance_id` - (Required, String) The GUID of the service instance associated with an account.
- `pi_volume_id` - (Required, String) The ID of the volume for which you want to retrieve detailed information.

## Attribute reference
In addition to all argument reference list, you can access the following attribute references after your data source is created. 

- `flash_copy_mappings` - (List) List of flash copy mappings details of a volume.

  Nested scheme for `flash_copy_mappings`:
      - `copy_rate` - (Integer) The rate of flash copy operation of a volume.
      - `flash_copy_name` - (String) The flash copy name of the volume.
      - `progress` - (Integer) The progress of flash copy operation.
      - `source_volume_name` (String) The name of the source volume.
      - `start_time` - (String) The start time of flash copy operation.
      - `status` - (String) The copy status of a volume.
      - `target_volume_name` (String) The name of the target volume.
