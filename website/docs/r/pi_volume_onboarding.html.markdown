---
subcategory: "Power Systems"
layout: "ibm"
page_title: "IBM: pi_volume_onboarding"
description: |-
  Manages IBM volume onboarding in the Power Virtual Server cloud.
---

# ibm_pi_volume_onboarding

Creates volume onboarding. For more information, about managing volume groups, see [getting started with IBM Power Systems Virtual Servers](https://cloud.ibm.com/docs/power-iaas?topic=power-iaas-getting-started).

## Example Usage

The following example attaches volume to a power systems virtual server instance.

```terraform
resource "ibm_pi_volume_onboarding" "testacc_volume_onboarding" {
  pi_cloud_instance_id = "<value of the cloud_instance_id>"
  pi_description       = "volume-onboarding-1"
  pi_onboarding_volumes {
    pi_source_crn = "< source crn >"
    pi_auxiliary_volumes {
      pi_auxiliary_volume_name = "< auxiliary volume name >"
      pi_display_name = "< display name >"
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

ibm_pi_volume_onboarding provides the following [timeouts](https://www.terraform.io/docs/language/resources/syntax.html) configuration options:

- **create** - (Default 15 minutes) Used for attaching volume.
- **delete** - (Default 15 minutes) Used for detaching volume.

## Argument Reference

Review the argument references that you can specify for your resource.

- `pi_cloud_instance_id` - (Required, Forces new resource, String) The GUID of the service instance associated with an account.
- `pi_description` - (Optional, String) The description of the volume onboarding operation.
- `pi_onboarding_volumes` - (Required, Forces new resource, List of objects) List of onboarding volumes.
  - Constraints: The minimum length is `1` items.

  Nested scheme for **pi_onboarding_volumes**:
  - `pi_auxiliary_volumes` - (Required, List of objects) List auxiliary volumes.
  - `pi_source_crn` - (Required, String) The crn of source service broker instance from where auxiliary volumes need to be onboarded.
    - Constraints: The minimum length is `1` items.

    Nested scheme for **pi_auxiliary_volumes**:
    - `pi_auxiliary_volume_name` - (Required, String) The auxiliary volume name.
    - `pi_display_name` - (Optional, String) The display name of auxiliary volume which is to be onboarded.

## Attribute Reference

In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `create_time` - (String) The create time of volume onboarding operation.
- `id` - (String) The unique identifier of the volume attach. The ID is composed of `<pi_cloud_instance_id>/<onboarding_id>`.
- `input_volumes` - (List of strings) List of volumes requested to be onboarded.
- `onboarding_id` - (String) The volume onboarding ID.
- `progress` - (Float) The progress of volume onboarding operation.
- `results_onboarded_volumes` - (List of strings) List of volumes which are onboarded successfully.
- `results_volume_onboarding_failures` - (List of objects) - The volume onboarding failure details.

  Nested scheme for `results_volume_onboarding_failures`:
  - `failure_message` - (String) The failure reason for the volumes which have failed to be onboarded.
  - `volumes` - (List of strings) List of volumes which have failed to be onboarded.
- `status` - (String) The status of volume onboarding operation.

## Import

The `ibm_pi_volume_onboarding` resource can be imported by using `pi_cloud_instance_id` and `onboarding_id`.

### Example

```bash
terraform import ibm_pi_volume_onboarding.example d7bec597-4726-451f-8a63-e62e6f19c32c/49fba6c9-23f8-40bc-9899-aca322ee7d5b
```
