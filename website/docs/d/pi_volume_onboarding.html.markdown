---
subcategory: "Power Systems"
layout: "ibm"
page_title: "IBM: pi_volume_onboarding"
description: |-
  Manages a volume onboarding in the Power Virtual Server cloud.
---

# ibm_pi_volume_onboarding

Retrieves information about volume onboarding. For more information, about managing a volume group, see [moving data to the cloud](https://cloud.ibm.com/docs/power-iaas?topic=power-iaas-moving-data-to-the-cloud).

## Example Usage

The following example retrieves information about about volume onboarding in Power Systems Virtual Server.

```terraform
data "ibm_pi_volume_onboarding" "ds_volume_onboarding" {
  pi_cloud_instance_id    = "49fba6c9-23f8-40bc-9899-aca322ee7d5b"
  pi_volume_onboarding_id = "1212a6c9-23f8-40bc-9899-aca322ee7343"
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
- `pi_volume_onboarding_id` - (Required, String) The ID of volume onboarding for which you want to retrieve detailed information.

## Attribute Reference

In addition to all argument reference list, you can access the following attribute references after your data source is created.

- `create_time` - (String) The create-time of volume onboarding operation.
- `description` - (String) The description of the volume onboarding operation.
- `id` - (String) The volume onboarding operation id.
- `input_volumes` - (List) List of volumes requested to be onboarded.
- `progress` - (String) The progress of volume onboarding operation.
- `results_onboarded_volumes` - (List) List of volumes which are onboarded successfully.
- `results_volume_onboarding_failures` - (List) The volume onboarding failure details.

  Nested scheme for `results_volume_onboarding_failures`:
  - `failure_message` - (String) The failure reason for the volumes which have failed to be onboarded.
  - `volumes` - (List) List of volumes which have failed to be onboarded.
- `status` - (String) The status of volume onboarding operation.
