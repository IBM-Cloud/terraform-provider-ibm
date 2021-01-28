---
layout: "ibm"
page_title: "IBM: pi_instance"
sidebar_current: "docs-ibm-datasources-pi-instance"
description: |-
  Manages an instance in the Power Virtual Server Cloud.
---

# ibm\_pi_instance

Import the details of an existing IBM Power Virtual Server instance as a read-only data source. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "ibm_pi_instance" "ds_instance" {
  pi_instance_name     = "terraform-test-instance"
  pi_cloud_instance_id = "49fba6c9-23f8-40bc-9899-aca322ee7d5b"
}
```
## Notes:
* Please find [supported Regions](https://cloud.ibm.com/apidocs/power-cloud#endpoint) for endpoints.
* If a Power cloud instance is provisioned at `lon04`, The provider level attributes should be as follows:
  * `region` - `lon`
  * `zone` - `lon04`
  Example Usage:
  ```hcl
    provider "ibm" {
      region    =   "lon"
      zone      =   "lon04"
    }
  ```
## Argument Reference

The following arguments are supported:

* `pi_instance_name` - (Required, string) The name of the instance.
* `pi_cloud_instance_id` - (Required, string) The GUID of the service instance associated with the account

## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier for this instance.
* `memory` - The memory of the instance.
* `processors` - The processors of the instance.
* `status` - The status of the instance.
* `proctype` - The proctype of the instance.
* `volumes` - The list of the volumes attached to the instance.
* `health_status` - The health status of the instance.
* `state` - The state of the instance.
* `min_processors` - Minimum number of processors that were  allocated (for resize)
* `min_memory` - Minimum memory  that was allocated (for resize)
* `max_processors` - Maximumx number of processors that can be allocated (for resize) without a shutdown/reboot of the lpar
* `max_memory` - Maximum amount of memory that can be allocated (for resize) without a shutdown/reboot of the lpar
* `virtual_cores_assigned` - The virtual cores that are assigned to the instance
* `max_virtual_cores` - The max value that we are increase to without a reboot
* `min_virutal_cores` - The min cores assigned to the instance
 * `addresses` - The addresses associated with this instance.  Nested `addresses` blocks have the following structure:
	* `ip` - IP of the instance.
  * `macaddress` - The macaddress of the instance.
  * `network_id` - The networkID of the instance.
  * `network_name` - The network name of the instance.
  * `type` - The type of the network
  * `external_ip` - The externalIP address of the instance.

