---

subcategory: "Power Systems"
layout: "ibm"
page_title: "IBM: pi_instance_ip"
description: |-
  Obtains the information about the ip address for a specific subnet on an instance.
---

# ibm\_pi_instance_ip

Import the details of an existing IBM Power Virtual Server instance as a read-only data source. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "ibm_pi_instance_ip" "ds_instance_ip" {
  pi_instance_name     = "terraform-test-instance"
  pi_network_name = "APP"
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
* `pi_network_name` - (Required,string) - The subnet that the vm belongs to. This should have been created.
* `pi_cloud_instance_id` - (Required, string) The GUID of the service instance associated with the account

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The unique identifier of the network.
* `ip` - The IP Address that is attached to this instance from that specific subnet. 
* `macaddress` - The macaddress of the network that is attached to this instance.
* `external_ip` - The external ip of the network that is attached to this instance. 
* `type` - The type of the network that is attached to this instance.
* `ipoctet` - The ipoctet of the network that is attached to this instance.