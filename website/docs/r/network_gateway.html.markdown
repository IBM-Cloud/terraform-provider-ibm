---
layout: "ibm"
page_title: "IBM : Network_Gateway"
sidebar_current: "docs-ibm-resource-Network-Gateway"
description: |-
  Manages IBM Network Gateway
---

# ibm\_network_gateway

Provides a network gateway resource. This allows a network gateway to be created, updated and deleted. This resource supports both HA (High Availability) and non HA models. For more detail on Networking solutions, refer to the [IBM Cloud Network page](https://www.ibm.com/cloud/network).

## Example Usage


```hcl

provider "ibm" {
}

resource "ibm_network_gateway" "gateway01" {
    hostname =        "Gateway01"
    domain =          "exampleDomain.com"
    datacenter =      "ams01"
    network_speed =   100
    memory =          4
    private_vlan_id = 123456
    public_vlan_id =  123456
    ipv6_enabled =    true
}

```


## Argument Reference

The following arguments are supported:

* `hostname` - (Required, string) The Network Gateway name.
* `domain` - (Required, string) The Network Gateway domain name.
* `datacenter` - (Required) The Datacenter in which you want to provision the Network Gateway.
* `network_speed` - (Required) The interface speed of the Network Gateway expressed in MPBS.
* `memory` - (Required) The amount of memory RAM that would be provisioned to the Network Gateway.
* `private_vlan_id` - (Optional) The Private VLAN where the Network Gateway would be provisioned.
* `public_vlan_id` - (Optional) The Public VLAN where the Network Gateway would be provisioned.
* `notes` - (Optional) additional notes added to the description of the Network Gateway.
* `process_key_name` - (Optional) Model of the processor to include in the Network Gateway order, consult Softlayer documentation to identify the possible options for your account.
* `os_key_name` - (Optional) Operating system for the Network Gateway, Soflayer supported are limited to Brocade options.
* `ipv6_enabled` - (Required) Indicates if the Network Gateway has ipv6 support, Softlayer default requires it to be true.
* `public_bandwidth` - (Optional) Bandwidth is measured from the port onto which your Network Gateway is connected to the Public Network.


## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier of the Network Gateway.
* `public_ipv4_address` - The public IPv4 address of the Network Gateway.
* `private_ipv4_address` - The private IPv4 address of the Network Gateway.
* `private_ipv4_address` - The public ipv6_address of the Network Gateway.
