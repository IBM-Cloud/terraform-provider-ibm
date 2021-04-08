---

subcategory: "Power Systems"
layout: "ibm"
page_title: "IBM: pi_network"
description: |-
  Manages networks in the IBM Power Virtual Server Cloud.
---

# ibm\_pi_network

Provides a network resource. This allows network to be created, updated and deleted.

## Example Usage

In the following example, you can create a network:

```hcl
resource "ibm_pi_network" "power_networks" {
  count                = 1
  pi_network_name      = "power-network"
  pi_cloud_instance_id = "<value of the cloud_instance_id>"
  pi_network_type      = "vlan"
  pi_cidr              = "<Network in CIDR notation (192.168.0.0/24)>"
  pi_dns               = [<"DNS Servers">]
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

## Timeouts

ibm_pi_network provides the following [timeout](https://www.terraform.io/docs/configuration/resources.html#timeouts) configuration options:

* `create` - (Default 60 minutes) Used for creating a network.
* `delete` - (Default 60 minutes) Used for deleting a network.

## Argument Reference

The following arguments are supported:

* `pi_network_name` - (Required, string) The name of the network.
* `pi_network_type` - (Required, string) The type of network (e.g., pub-vlan, vlan).
* `pi_cloud_instance_id` - (Required, string) The GUID of the service instance associated with the account
* `pi_dns` - (Optional, list(strings)) List of DNS entries for the network. Required for `vlan` network type.
* `pi_cidr` - (Optional, string) The network CIDR. Required for `vlan` network type.


## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The unique identifier of the network.The id is composed of \<power_instance_id\>/\<network_id\>.
* `network_id` - The unique identifier (string) of the network.
* `vlan_id` - The unique identifier (int) of the network VLAN.

## Import

ibm_pi_network can be imported using `power_instance_id` and `network_id`, eg

```
$ terraform import ibm_pi_network.example d7bec597-4726-451f-8a63-e62e6f19c32c/cea6651a-bc0a-4438-9f8a-a0770bbf3ebb
```
