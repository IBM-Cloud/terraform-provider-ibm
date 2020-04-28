---
layout: "ibm"
page_title: "IBM: pi_network_port"
sidebar_current: "docs-ibm-resource-pi-network_port"
description: |-
  Manages IBM Network Ports in the Power Virtual Server Cloud.
---

# ibm\_network_port

Provides a Network Port resource. This allows Network Port to be created, updated, and cancelled in the Power Virtual Server Cloud.

## Example Usage

In the following example, you can create a network port :

```hcl
resource "ibm_pi_network_port" "network_port" {
  pi_network_name          = "<id of the network name>"
  pi_cloud_instance_id = "<value of the cloud_instance_id>"
}
```

## Timeouts

ibm_network_port provides the following [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) configuration options:

* `create` - (Default 60 minutes) Used for creating a Network Port.
* `delete` - (Default 60 minutes) Used for deleting a Network Port.
* `update` - (Default 60 minutes) Used for Updating a Network Port.

## Argument Reference

The following arguments are supported:

* `pi_network_name` - (Required, int) The key name.
* `pi_cloud_instance_id` - (Required, string) The cloud_instance_id for this account.


## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier of the port.The id is composed of \<power_instance_id\>/\<portid\>.
* `portid` -  The unique identifier of the key.
* `ipaddress` - The ip address defined for this port
* `macaddress` - The mac address defined for the port
* `status`  - Status of this port. WHen created it will always be DOWN



## Import

ibm_pi_network_port can be imported using `power_instance_id` and `port_id`, eg

```
$ terraform import ibm_pi_network_port.example d7bec597-4726-451f-8a63-e62e6f19c32c/a16685ed-4d7c-4faa-ac23-a9ae24215241
```