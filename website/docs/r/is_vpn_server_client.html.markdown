---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : VPN-Server-Client"
description: |-
  Manages IBM VPN Server Client Disconnect or Delete.
---

# ibm_is_vpn_server_client

Provides VPNServer client delete or disconnect functionality for VPNServer. For more information, about VPN Server Client, see [Setting up a client VPN environment and connecting to a VPN server](https://cloud.ibm.com/docs/vpc?topic=vpc-vpn-client-environment-setup&interface=ui).

## Example Usage

```terraform
resource "ibm_is_vpn_server_client" "example" {
  vpn_server    = ibm_is_vpn_server.example.vpn_server
  vpn_client    = "id"
  delete        = true
}
```

## Argument Reference
Review the argument references that you can specify for your resource. 

- `vpn_server` - (Required, Forces new resource, String) The VPN server identifier.
- `vpn_client` - (Required, Forces new resource, String) The VPN client identifier.
- `delete` - (Optional, String) The delete to use for this VPN client to be deleted or not, when false, client is disconnected and when set to true client is deleted.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your resource is created.

- `id` - The unique identifier of the VPNServerClient, it is combination of vpn_server, vpn_client, status_code.
- `status_code` - status code of the result.
- `description` - description of the result.

## Import

You can import the `ibm_is_vpn_server_client` resource by using `id`. The unique identifier for this VPN server client.

# Syntax
```
$ terraform import ibm_is_vpn_server_client.example <id>
```

# Example
```
$ terraform import ibm_is_vpn_server_client.example r006-d7cc5196-9864-48c4-82d8-3f30da41fcc5/r006-d7cc5196-9864-48c4-82d8-3h30db41acd5/202
```
