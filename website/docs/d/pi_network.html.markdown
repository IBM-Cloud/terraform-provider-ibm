---
layout: "ibm"
page_title: "IBM : Network"
sidebar_current: "docs-ibm-datasources-pi-network"
description: |-
  Manages IBM Cloud Infrastructure Networks for IBM Power
---

# ibm\_pi_network

Import the details of an existing IBM Cloud Infrastructure Network for IBM Power as a read-only data source. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.


## Example Usage

```hcl

data "ibm_pi_network" "ds_network" {
    name = "APP"
    powerinstanceid="49fba6c9-23f8-40bc-9899-aca322ee7d5b"

}

```

## Argument Reference

The following arguments are supported:

* `networkname` - (Required, string) The name of the network.
* `powerinstanceid` - (Required, string) The service instance associated with the account


## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier for this network.
* `gateway_ip` - The gateway for this network.
* `available_ip_count` - The available ips for this network
* `used_ip_count` - The ips that are in use for this network
* `used_ip_percent` - The used ip in percent.
* `vlan_type` - The type of vlan for this network.
