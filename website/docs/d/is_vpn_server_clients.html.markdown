---
layout: "ibm"
page_title: "IBM : ibm_is_vpn_server_clients"
description: |-
  Get information about VPNServerClientCollection
subcategory: "VPC infrastructure"
---

# ibm_is_vpn_server_clients

Provides a read-only data source for VPNServerClientCollection. For more information, about VPN Server Clients, see [Setting up a client VPN environment and connecting to a VPN server](https://cloud.ibm.com/docs/vpc?topic=vpc-vpn-client-environment-setup&interface=ui).
## Example Usage

```terraform
data "ibm_is_vpn_server_clients" "example" {
	vpn_server = ibm_is_vpn_server.example.id
}
```

## Argument Reference

Review the argument reference that you can specify for your data source.

- `vpn_server` - (Required, String) The VPN server identifier.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

- `id` - The unique identifier of the VPNServerClientCollection.
- `clients` - (List) Collection of VPN clients.
	Nested scheme for `clients`:
	- `client_ip` - (List) The IP address assigned to this VPN client from `client_ip_pool`.
		Nested scheme for `client_ip`:
		- `address` - (String) The IP address. This property may add support for IPv6 addresses in the future. When processing a value in this property, verify that the address is in an expected format. If it is not, log an error. Optionally halt processing and surface the error, or bypass the resource on which the unexpected IP address format was encountered.
	- `common_name` - (String) The common name of client certificate that the VPN client provided when connecting to the server.This property will be present only when the `certificate` client authentication method is enabled on the VPN server.
	- `created_at` - (String) The date and time that the VPN client was created.
	- `disconnected_at` - (String) The date and time that the VPN client was disconnected.
	- `href` - (String) The URL for this VPN client.
	- `id` - (String) The unique identifier for this VPN client.
	- `remote_ip` - (List) The remote IP address of this VPN client.
	</br>Nested scheme for `remote_ip`:
		- `address` - (String) The IP address. This property may add support for IPv6 addresses in the future. When processing a value in this property, verify that the address is in an expected format. If it is not, log an error. Optionally halt processing and surface the error, or bypass the resource on which the unexpected IP address format was encountered.
	- `remote_port` - (Integer) The remote port of this VPN client.
	- `resource_type` - (String) The resource type.
	- `status` - (String) The status of the VPN client.
	- `username` - (String) The username that this VPN client provided when connecting to the VPN server.