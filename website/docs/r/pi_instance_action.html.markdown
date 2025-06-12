---
subcategory: "Power Systems"
layout: "ibm"
page_title: "IBM: pi_instance_action"
description: |-
  Performs an action start, stop, reboot, immediate-shutdown, reset on a PVM instance.
---

# ibm_pi_instance_action

Performs an action on a [Power Systems Virtual Server instance](https://cloud.ibm.com/docs/power-iaas?topic=power-iaas-creating-power-virtual-server).

## Example Usage

The following example perform an action "hard-reboot" on a Power Systems Virtual Server instance.

```terraform
resource "ibm_pi_instance_action" "example" {
  pi_cloud_instance_id  = "d7bec597-4726-451f-8a63-e62e6f19c32c"
  pi_instance_id        = "cea6651a-bc0a-4438-9f8a-a0770b112ebb"
  pi_action             = "hard-reboot"
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

The `ibm_pi_instance_action` provides the following [timeouts](https://www.terraform.io/docs/language/resources/syntax.html) configuration options:

- **create** - (Default 15 minutes) Used for taking action on the instance.
- **update** - (Default 15 minutes) Used for updating action on the instance.

## Argument References

Review the argument references that you can specify for your resource.

- `pi_action` - (Required, String) Name of the action to take. Allowed values are `start`, `stop`, `hard-reboot`, `soft-reboot`, `immediate-shutdown`, `reset-state`.
- `pi_cloud_instance_id` - (Required, String) The GUID of the service instance associated with an account.
- `pi_health_status` - (Optional, String) Specifies if Terraform should poll for the health status to be `OK` or `WARNING`. The default value is `OK`. Ignored for `pi_action = "reset-state"`.
- `pi_instance_id` - (Required, String) The ID of the pvm instance to perform an action on.

## Attribute Reference

In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `health_status` - (String) The health status of the VM.
- `id` - (String) The unique identifier of the instance. The ID is composed of `<pi_cloud_instance_id>/<pi_instance_id>`.
- `progress` - (Float) - The progress of the instance.
- `status` - (String) - The status of the instance.

## Import

The `ibm_pi_instance_action` can be imported using `pi_cloud_instance_id` and `pi_instance_id`.

### Example

```bash
terraform import ibm_pi_instance_action.example d7bec597-4726-451f-8a63-e62e6f19c32c/cea6651a-bc0a-4438-9f8a-a0770b112ebb
```
