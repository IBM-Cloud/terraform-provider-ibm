---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : ibm_is_vpn_gateway_connection"
description: |-
  Get information about IBM Cloud VPN Connection
---

# ibm_is_vpn_gateway_connection

Provides a read-only data source for VPN Connection. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "ibm_is_vpn_gateway_connection" "example" {
  vpn_gateway = ibm_is_vpn_gateway.example.id
  vpn_gateway_connection = ibm_is_vpn_gateway_connection.example.gateway_connection
}
data "ibm_is_vpn_gateway_connection" "example-1" {
  vpn_gateway = ibm_is_vpn_gateway.example.id
  vpn_gateway_connection_name = ibm_is_vpn_gateway_connection.example.name
}
data "ibm_is_vpn_gateway_connection" "example-2" {
  vpn_gateway_name = ibm_is_vpn_gateway.example.name
  vpn_gateway_connection = ibm_is_vpn_gateway_connection.example.gateway_connection
}
data "ibm_is_vpn_gateway_connection" "example-3" {
  vpn_gateway_name = ibm_is_vpn_gateway.example.name
  vpn_gateway_connection_name = ibm_is_vpn_gateway_connection.example.name
}
```

## Argument Reference

Review the argument reference that you can specify for your data source.

- `vpn_gateway` - (Optional, String) The VPN gateway identifier.
- `vpn_gateway_name` - (Optional, String) The VPN gateway name.
- `vpn_gateway_connection` - (Optional, String) The VPN gateway connection identifier.
- `vpn_gateway_connection_name` - (Optional, String) The VPN gateway connection name.

  ~> **Note** Provide either one of `vpn_gateway`, `vpn_gateway_name` to identifiy vpn gateway and either one of `vpn_gateway_connection`, `vpn_gateway_connection_name` to identify vpn gateway connection.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

- `id` - The unique identifier of the VPN gateway connection. The ID is composed of `<vpn_gateway_id>/<vpn_gateway_connection_id>`.
- `admin_state_up` - (Boolean) If set to false, the VPN gateway connection is shut down.

- `authentication_mode` - (String) The authentication mode. Only `psk` is currently supported.

- `created_at` - (String) The date and time that this VPN gateway connection was created.

- `dead_peer_detection` - (List) The Dead Peer Detection settings.
  Nested scheme for **dead_peer_detection**:
	- `action` - (String) Dead Peer Detection actions.
	- `interval` - (Integer) Dead Peer Detection interval in seconds.
	- `timeout` - (Integer) Dead Peer Detection timeout in seconds. Must be at least the interval.

- `href` - (String) The VPN connection's canonical URL.

- `ike_policy` - (List) The IKE policy. If absent, [auto-negotiation isused](https://cloud.ibm.com/docs/vpc?topic=vpc-using-vpn&interface=ui#ike-auto-negotiation-phase-1).
  Nested scheme for **ike_policy**:
	- `deleted` - (List) If present, this property indicates the referenced resource has been deleted and provides some supplementary information.
	  Nested scheme for **deleted**:
		- `more_info` - (String) Link to documentation about deleted resources.
	- `href` - (String) The IKE policy's canonical URL.
	- `id` - (String) The unique identifier for this IKE policy.
	- `name` - (String) The user-defined name for this IKE policy.
	- `resource_type` - (String) The resource type.

- `ipsec_policy` - (List) The IPsec policy. If absent, [auto-negotiation isused](https://cloud.ibm.com/docs/vpc?topic=vpc-using-vpn&interface=ui#ipsec-auto-negotiation-phase-2).
  Nested scheme for **ipsec_policy**:
	- `deleted` - (List) If present, this property indicates the referenced resource has been deleted and provides some supplementary information.
	  Nested scheme for **deleted**:
		- `more_info` - (String) Link to documentation about deleted resources.
	- `href` - (String) The IPsec policy's canonical URL.
	- `id` - (String) The unique identifier for this IPsec policy.
	- `name` - (String) The user-defined name for this IPsec policy.
	- `resource_type` - (String) The resource type.

- `local_cidrs` - (List) The local CIDRs for this resource.

- `mode` - (String) The mode of the VPN gateway.

- `name` - (String) The user-defined name for this VPN gateway connection.

- `peer_address` - (String) The IP address of the peer VPN gateway.

- `peer_cidrs` - (List) The peer CIDRs for this resource.

- `psk` - (String) The preshared key.

- `resource_type` - (String) The resource type.

- `routing_protocol` - (String) Routing protocols are disabled for this VPN gateway connection.

- `status` - (String) The status of a VPN gateway connection.

- `tunnels` - (List) The VPN tunnel configuration for this VPN gateway connection (in static route mode).
  Nested scheme for **tunnels**:
	- `public_ip_address` - (String) The IP address of the VPN gateway member in which the tunnel resides. This property may add support for IPv6 addresses in the future. When processing a value in this property, verify that the address is in an expected format. If it is not, log an error. Optionally halt processing and surface the error, or bypass the resource on which the unexpected IP address format was encountered.
	- `status` - (String) The status of the VPN Tunnel.

