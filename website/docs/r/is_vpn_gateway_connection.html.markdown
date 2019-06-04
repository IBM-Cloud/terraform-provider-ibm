---
layout: "ibm"
page_title: "IBM : VPN-gateway-connection"
sidebar_current: "docs-ibm-resource-is-vpn-gateway-connection"
description: |-
  Manages IBM VPN Gateway Connection
---

# ibm\_is_vpn_gateway_connection

Provides a VPN gateway connection resource. This allows VPN gateway connection to be created, updated, and cancelled.


## Example Usage

In the following example, you can create a VPN gateway:

```hcl
resource "ibm_is_vpn_gateway_connection" "VPNGatewayConnection" {
  name          = "test2"
  vpn_gateway   = "${ibm_is_vpn_gateway.testacc_VPNGateway2.id}"
  peer_address  = "${ibm_is_vpn_gateway.testacc_VPNGateway2.public_ip_address}"
  preshared_key = "VPNDemoPassword"
  local_cidrs   = ["${ibm_is_subnet.testacc_subnet2.ipv4_cidr_block}"]
  peer_cidrs    = ["${ibm_is_subnet.testacc_subnet1.ipv4_cidr_block}"]
}

```

## Timeouts

ibm_is_vpn_gateway_connection provides the following [Timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) configuration options:

* `delete` - (Default 60 minutes) Used for deleting Instance.

## Argument Reference

The following arguments are supported:

* `name` - (Required, string) Name of the VPN gateway connection.
* `vpn_gateway` - (Required, string) The unique identifier of VPN gateway.
* `peer_address` - (Required, string) The IP address of the peer VPN gateway.
* `preshared_key`: The preshared key.
* `local_cidrs` - (Optional, list) List of CIDRs for this resource.
* `peer-cidrs` - (Optional, list) List of CIDRs for this resource.
* `admin-state-up` - (Optional, bool) VPN gateway connection status. Default false. If set to false, the VPN gateway connection is shut down
* `dead-peer-detection-action` - (Optional, string) Dead Peer Detection actions. Supported values are restart, clear, hold, none. Default `none`
* `dead-peer-detection-interval` - (Optional, int) Dead Peer Detection interval in seconds. Default 30.
* `dead-peer-detection-timeout` - (Optional, int) Dead Peer Detection timeout in seconds. Default 120.
* `ike-policy` - (Optional, string) ID of the IKE policy.
* `ipsec-policy` - (Optional, string) ID of the IPSec policy.

## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier of the VPN gateway connection. The id is composed of \<vpn_gateway_id\>/\<vpn_gateway_connection_id\>.
* `status` - The status of VPN gateway connection.

## Import

ibm_is_vpn_gateway_connection can be imported using vpn gateway ID and vpn gateway connection ID, eg

```
$ terraform import ibm_is_vpn_gateway_connection.example d7bec597-4726-451f-8a63-e62e6f19c32c/cea6651a-bc0a-4438-9f8a-a0770bbf3ebb
```
