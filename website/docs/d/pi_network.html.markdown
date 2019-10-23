---
layout: "ibm"
page_title: "IBM: pi_network"
sidebar_current: "docs-ibm-datasources-pi-network"
description: |-
  Manages a network in the IBM Power Virtual Server Cloud.
---

# ibm\_pi_network

Import the details of an existing IBM Power Virtual Server Cloud network as a read-only data source. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "ibm_pi_network" "ds_network" {
  pi_network_name = "APP"
  powerinstanceid = "49fba6c9-23f8-40bc-9899-aca322ee7d5b"
}
```

## Argument Reference

The following arguments are supported:

* `pi_network_name` - (Required, string) The name of the network.
* `pi_cloud_instance_id` - (Required, string) The service instance associated with the account

## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier for this network.
* `gateway_ip` - The gateway for this network.
* `available_ip_count` - The available IPs for this network.
* `used_ip_count` - The IPs that are in use for this network.
* `used_ip_percent` - The used ip in percent.
* `vlan_type` - The type of vlan for this network.
