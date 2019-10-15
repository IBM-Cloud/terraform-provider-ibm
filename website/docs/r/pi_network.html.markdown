---
layout: "ibm"
page_title: "IBM: pi_network"
sidebar_current: "docs-ibm-resource-pi-network"
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
}
```

## Timeouts

ibm_pi_network provides the following [timeout](https://www.terraform.io/docs/configuration/resources.html#timeouts) configuration options:

* `create` - (Default 60 minutes) Used for creating a network.
* `delete` - (Default 60 minutes) Used for deleting a network.

## Argument Reference

The following arguments are supported:

* `pi_network_name` - (Required, string) The name of the network.
* `pi_network_dns` - (Required, list(strings)) List of DNS entries for the network.
* `pi_network_cidr` - (Required, string) The network CIDR.
* `pi_network_type` - (Required, string) The type of network (e.g., pub-vlan, vlan).
* `pi_network_gateway` - (Optional, string) The network gateway address.
* `pi_network_available_ip_count` - (Optional, float) The number of available IP addresses.
* `pi_network_used_ip_count` - (Optional, float) The number of used IP addresses.
* `pi_network_used_ip_percent` - (Optional, float) The percentage of IP addresses used.
* `pi_cloud_instance_id` - (Required, string) The cloud_instance_id for this account.

## Attribute Reference

The following attributes are exported:

* `networkid` - The unique identifier (string) of the network.
* `vlanid` - The unique identifier (int) of the network VLAN.
