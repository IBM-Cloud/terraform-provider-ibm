---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : "
description: |-
  Manages IBM vpn gateway connections.
---

# ibm\_is_vpn_gateway_connections

Import the details of an existing IBM VPN Gateway connections as a read-only data source. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.


## Example Usage

```hcl

data "ibm_is_vpn_gateway_connections" "ds_vpn_gateway_connections" {
  vpn_gateway = ibm_is_vpn_gateway.testacc_vpnGateway.id
}

```

## Argument Reference

The following arguments are supported:

* `status` - (Optional, string) Filters the collection to VPN gateway connections with the specified status .
* `vpn_gateway` - (Required, string) The VPN gateway identifier(ID).

## Attribute Reference

The following attributes are exported:

* `id` - ID of the VPN Gateway connection.
* `admin_state_up` - VPN gateway connection admin state,default: true.
* `authentication_mode` - The authentication mode.
* `created_at` - The date and time that this VPN gateway connection was created.
* `ike_policy` - VPN gateway connection IKE Policy.
* `interval` - Interval for dead peer detection interval.
* `ipsec_policy` - IP security policy for vpn gateway connection.
* `local_cidrs` - VPN gateway connection local CIDRs.
* `mode` - The mode of the VPN gateway.
* `name` - VPN Gateway connection name.
* `peer_address` - VPN gateway connection peer address.
* `peer_cidrs` - VPN gateway connection peer CIDRs.
* `resource_type` - The resource type.
* `timeout` - Timeout for dead peer detection
* `action` - Action detection for dead peer detection action
* `tunnels` - The VPN tunnel configuration for this VPN gateway connection (in static route mode)
  * `address` - The IP address of the VPN gateway member in which the tunnel resides
  * `status` - The status of the VPN Tunnel