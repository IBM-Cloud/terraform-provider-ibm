---
layout: "ibm"
page_title: "IBM : ibm_is_vpn_server_client_configuration"
description: |-
  Get information about VPN Server Client Configuration
subcategory: "VPC infrastructure"
---

# ibm_is_vpn_server_client_configuration

Provides a read-only data source for VPN Server Client Configuration. For more information, about VPN Server Client Configuration, see [Setting up a client VPN environment and connecting to a VPN server](https://cloud.ibm.com/docs/vpc?topic=vpc-vpn-client-environment-setup&interface=ui).

## Example Usage

```terraform
data "ibm_is_vpn_server_client_configuration" "example" {
	vpn_server = ibm_is_vpn_server.example.id
	file_path = "vpn_server_client_configuration.txt"
}
```

## Argument Reference

Review the argument reference that you can specify for your data source.

- `vpn_server` - (Required, String) The VPN server identifier.
- `file_path` - (Optional, String) The File path to store configuration.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

- `id` - The unique identifier of the VPNServerClientConfiguration.
- `vpn_server_client_configuration` - (String) The client configuration of vpn server.
