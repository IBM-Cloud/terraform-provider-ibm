---
subcategory: "Power Systems"
layout: "ibm"
page_title: "IBM: pi_virtual_serial_number"
description: |-
  Manages a virtual serial number in IBM Power
---

# ibm_pi_virtual_serial_number

Create, get, update or delete an existing virtual serial number.

## Example Usage

The following example enables you to create a virtual serial number:

```terraform
resource "ibm_pi_virtual_serial_number" "testacc_virtual_serial_number" {
  pi_serial            = "<existing virtual serial number>"
  pi_cloud_instance_id = "<value of the cloud_instance_id>"
  pi_description       = "<desired description for virtual serial number>"
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

~> **Note** This resource is used to create a virtual serial number by assigning to a power instance using 'auto-assign'. Otherwise, it can only be used to manage an existing virtual serial number. If deleting an instance with a `ibm_pi_virtual_serial_number` resource assigned, it is recommended to unassign the virtual serial number resource from the instance before destruction.

## Timeouts

ibm_pi_virtual_serial_number provides the following [timeouts](https://www.terraform.io/docs/language/resources/syntax.html) configuration options:

- **create** - (Default 45 minutes) Used for creating a virtual serial number.
- **update** - (Default 45 minutes) Used for updating a virtual serial number.
- **delete** - (Default 45 minutes) Used for deleting a reserved virtual serial number.

## Argument Reference

Review the argument references that you can specify for your resource.

- `pi_cloud_instance_id` - (Required, String) The GUID of the service instance associated with an account.
- `pi_description` - (Optional, String) Desired description for virtual serial number.
- `pi_instance_id` - (Optional, String) Power instance ID to assign created or existing virtual serial number to. Must unassign from previous power instance if different than current assignment. Cannot use the instance name, only ID. Please see note on `pi_virtual_serial_number` in the `ibm_pi_instance` resource documentation.
- `pi_retain_virtual_serial_number` - (Optional, Boolean) Indicates whether to reserve or delete virtual serial number when detached from power instance during deletion. Required with `pi_instance_id`. Default behavior does not retain virtual serial number after deletion.
- `pi_serial` - (Required, String) Virtual serial number of existing serial. Cannot use 'auto-assign' unless `pi_instance_id` is specified.

    ~> **Note** When set to "auto-assign" in the configuration, changes to `pi_serial` outside of terraform will not be detected.

## Attribute Reference

 In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `id` - (String) The unique identifier of the virtual serial number. Composed of `<cloud_instance_id>/<virtual_serial_number>`

## Import

The `ibm_virtual_serial_number` resource can be imported by using `pi_cloud_instance_id` and `serial`.

### Example

```bash
terraform import ibm_pi_virtual_serial_number.example d7bec597-4726-451f-8a63-e62e6f19c32c/VS0762Y
```
