---
layout: "ibm"
page_title: "IBM : ibm_is_vpn_server_client"
description: |-
  Get information about VPNServerClient
subcategory: "VPC infrastructure"
---

# ibm_is_vpn_server_client

Provides a read-only data source for VPNServerClient. For more information, about VPN Server Client, see [Setting up a client VPN environment and connecting to a VPN server](https://cloud.ibm.com/docs/vpc?topic=vpc-vpn-client-environment-setup&interface=ui).

## Example Usage

```terraform
data "ibm_is_vpn_server_client" "example" {
	identifier = "id"
	vpn_server = ibm_is_vpn_server.example.id
}
```

## Argument Reference

Review the argument reference that you can specify for your data source.

- `identifier` - (Required, String) The VPN client identifier.
- `vpn_server` - (Required, String) The VPN server identifier.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

- `id` - The unique identifier of the VPNServerClient.
- `client_ip` - (List) The IP address assigned to this VPN client from `client_ip_pool`.

  Nested scheme for **client_ip**:
    - `address` - (String) The IP address. This property may add support for IPv6 addresses in the future. When processing a value in this property, verify that the address is in an expected format. If it is not, log an error. Optionally halt processing and surface the error, or bypass the resource on which the unexpected IP address format was encountered.
      - Constraints: The value must match regular expression `/^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$/`

- `common_name` - (String) The common name of client certificate that the VPN client provided when connecting to the server.This property will be present only when the `certificate` client authentication method is enabled on the VPN server.
  - Constraints: The maximum length is `64` characters. The minimum length is `1` character.

- `created_at` - (String) The date and time that the VPN client was created.

- `disconnected_at` - (String) The date and time that the VPN client was disconnected.

- `href` - (String) The URL for this VPN client.
  - Constraints: The value must match regular expression `/^http(s)?:\/\/([^\/?#]-)([^?#]-)(\\?([^#]-))?(#(.-))?$/`

- `remote_ip` - (List) The remote IP address of this VPN client.

    Nested scheme for **remote_ip**:
    - `address` - (String) The IP address. This property may add support for IPv6 addresses in the future. When processing a value in this property, verify that the address is in an expected format. If it is not, log an error. Optionally halt processing and surface the error, or bypass the resource on which the unexpected IP address format was encountered.
      - Constraints: The value must match regular expression `/^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$/`

- `remote_port` - (Integer) The remote port of this VPN client.
  - Constraints: The maximum value is `65535`. The minimum value is `1`.

- `resource_type` - (String) The resource type.
  - Constraints: Allowable values are: vpn_server_client

- `status` - (String) The status of the VPN client:- `connected`: the VPN client is `connected` to this VPN server.- `disconnected`: the VPN client is `disconnected` from this VPN server.The enumerated values for this property are expected to expand in the future. When processing this property, check for and log unknown values. Optionally halt processing and surface the error, or bypass the VPN client on which the unexpected property value was encountered.
  - Constraints: Allowable values are: connected, disconnected

- `username` - (String) The username that this VPN client provided when connecting to the VPN server.This property will be present only when  the`username` client authentication method is enabled on the VPN server.

