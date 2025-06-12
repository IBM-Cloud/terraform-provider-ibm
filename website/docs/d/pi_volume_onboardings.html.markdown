---
subcategory: "Power Systems"
layout: "ibm"
page_title: "IBM: pi_volume_onboardings"
description: |-
  Manages a volume onboardings in the Power Virtual Server cloud.
---

# ibm_pi_volume_onboardings

Retrieves information about volume onboardings. For more information, about managing a volume group, see [moving data to the cloud](https://cloud.ibm.com/docs/power-iaas?topic=power-iaas-moving-data-to-the-cloud).

## Example Usage

The following example retrieves information about about volume onboardings in Power Systems Virtual Server.

```terraform
data "ibm_pi_volume_onboardings" "ds_volume_onboardings" {
  pi_cloud_instance_id = "49fba6c9-23f8-40bc-9899-aca322ee7d5b"
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
  
## Argument Reference

Review the argument references that you can specify for your data source.

- `pi_cloud_instance_id` - (Required, String) The GUID of the service instance associated with an account.

## Attribute Reference

In addition to all argument reference list, you can access the following attribute references after your data source is created.

- `onboardings` - (List) List of volume onboardings.

  Nested scheme for `onboardings`:
      - `description` - (String) The description of the volume onboarding operation.
      - `id` - (String) The type of cycling mode used.
      - `input_volumes` - (List) List of volumes requested to be onboarded.
      - `status` - (String) The status of volume onboarding operation.
