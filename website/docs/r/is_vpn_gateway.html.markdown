---
layout: "ibm"
page_title: "IBM : VPN-gateway"
sidebar_current: "docs-ibm-resource-is-vpn-gateway"
description: |-
  Manages IBM VPN Gateway
---

# ibm\_is_vpn_gateway

Provides a VPN gateway resource. This allows VPN gateway to be created, updated, and cancelled.


## Example Usage

In the following example, you can create a VPN gateway:

```hcl
resource "ibm_is_vpn_gateway" "testacc_vpn_gateway" {
  name   = "test"
  subnet = "a4ce411d-e118-4802-95ad-525e6ea0cfc9"
}

```

## Timeouts

ibm_is_vpn_gateway provides the following [Timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) configuration options:

* `create` - (Default 10 minutes) Used for creating vpn gateway Instance.
* `delete` - (Default 10 minutes) Used for deleting vpn gateway Instance.

## Argument Reference

The following arguments are supported:

* `name` - (Required, string) Name of the VPN gateway.
* `subnet` - (Required, Forces new resource, string) The unique identifier for this subnet.
* `resource_group` - (Optional, Forces new resource, string) The resource group where the VPN gateway to be created.
* `tags` - (Optional, array of strings) Tags associated with the VPN Gateway.

## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier of the VPN gateway.
* `status` - The status of VPN gateway.
* `public_ip_address` -  The IP address assigned to this VPN gateway.


## Import

ibm_is_vpn_gateway can be imported using ID, eg

```
$ terraform import ibm_is_vpn_gateway.example d7bec597-4726-451f-8a63-e62e6f19c32c
```