---
subcategory: "Power Systems"
layout: "ibm"
page_title: "IBM: ibm_pi_virtual_serial_numbers"
description: |-
  Provides data for a virtual_serial_number in an IBM Power Virtual Server cloud.
---

# ibm_virtual_serial_numbers

Retrieve information about existing virtual serial numbers as a read-only data source. For more information, about IBM power virtual server cloud, see [getting started with IBM Power Systems Virtual Servers](https://cloud.ibm.com/docs/power-iaas?topic=power-iaas-getting-started).

## Example Usage

```terraform
data "ibm_pi_virtual_serial_numbers" "ds_virtual_serial_number" {
  pi_cloud_instance_id     = "<cloud instance id>"
  pi_virtual_serial_number = "<virtual serial number>"
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

Review the argument reference that you can specify for your data source.

- `pi_cloud_instance_id` - (Required, String) The GUID of the service instance associated with an account.
- `pi_instance_id` - (Optional, String) Power virtual server instance ID.

## Attribute Reference

In addition to the argument reference list, you can access the following attribute references after your data source is created.

- `virtual_serial_numbers` - (List) List of virtual serial numbers

  Nested scheme for `virtual_serial_numbers`:
  - `description` - (String) Description for virtual serial number.
  - `pvm_instance_id` - (String) ID of PVM virtual serial number is attached to.
  - `serial` - (String) Virtual serial number.
