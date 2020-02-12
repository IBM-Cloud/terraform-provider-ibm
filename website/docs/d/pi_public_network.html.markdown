---
layout: "ibm"
page_title: "IBM: pi_public_network"
sidebar_current: "docs-ibm-datasources-pi-public-network"
description: |-
  Manages a public network in the IBM Power Virtual Server Cloud.
---

# ibm\_pi_public_network

Import the details of an existing IBM Power Virtual Server Cloud public network as a read-only data source. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "ibm_pi_public_network" "ds_public_network" {
  pi_network_name      = "PUBLIC"
  pi_cloud_instance_id = "49fba6c9-23f8-40bc-9899-aca322ee7d5b"
}
```

## Argument Reference

The following arguments are supported:

* `pi_network_name` - (Deprecated, string) The name of the network.
* `pi_cloud_instance_id` - (Required, string) The service instance associated with the account

## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier for this network.
* `type` - The network type for this network.
* `name` - The name of the network.
* `vlan_id` - The VLAN id for the network.