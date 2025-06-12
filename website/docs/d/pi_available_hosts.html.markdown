---
subcategory: "Power Systems"
layout: "ibm"
page_title: "IBM : pi_available_hosts"
description: |-
  List all available hosts
---

# ibm_pi_available_hosts

Retrieve the details information about available hosts. For more information, about available host, see [dedicated hosts](https://cloud.ibm.com/docs/virtual-servers?topic=virtual-servers-provisioning-dedicated-hosts-instances).

## Example Usage

The following example shows how to retrieve information using `ibm_pi_available_hosts`.

```terraform
data "ibm_pi_available_hosts" "pi_available_hosts" {
   pi_cloud_instance_id = "<value of the cloud_instance_id>"
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

## Attribute Reference

In addition to the argument reference list, you can access the following attribute references after your data source is created.

- `available_hosts` (List) Lists of all availabe hosts.

    Nested scheme for `available_hosts`:
       - `available_cores`- (Float) Core capacity of the host.
       - `available_memory`- (Float) Memory capacity of the host (in GB).
       - `count`- (int) The number of hosts with similar types/capacities that are available.
       - `sys_type`- (String) System type.
  