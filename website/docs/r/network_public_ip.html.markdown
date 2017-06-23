---
layout: "ibm"
page_title: "IBM: network_private_ip"
sidebar_current: "docs-ibm-resource-network-public-ip"
description: |-
  Manages IBM Network Public IP.
---

# ibm\_network_public_ip

Provides a public IP resource to route between servers. This allows public IPs to be created, updated, and deleted. public IPs are not restricted to routing within the same data center.

For additional details, see the [Bluemix Infrastructure (SoftLayer) API docs](http://sldn.softlayer.com/reference/services/SoftLayer_Network_Subnet_IpAddress_Global) and [public IP address overview](https://knowledgelayer.softlayer.com/learning/global-ip-addresses).

## Example Usage

```hcl
resource "ibm_network_public_ip" "test_public_ip " {
    routes_to = "119.81.82.163"
}
```

## Argument Reference

The following arguments are supported:

* `routes_to` - (Required, string) Destination IP address that the public IP routes traffic through. The destination IP address can be a public IP address of IBM resources in the same account, such as a public IP address of vm and public virtual IP address of NetScaler VPXs. 

## Attributes Reference

The following attributes are exported:

* `id` - ID of the public IP.
* `ip_address` - Address of the public IP.
