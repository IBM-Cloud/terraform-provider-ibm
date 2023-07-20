---
layout: "ibm"
page_title: "IBM : ibm_is_vpn_server_route"
description: |-
  Get information about VPNServerRoute
subcategory: "VPC infrastructure"
---

# ibm_is_vpn_server_route

Provides a read-only data source for VPNServerRoute. For more information, about VPN Server Routes, see [Managing VPN Server routes](https://cloud.ibm.com/docs/vpc?topic=vpc-vpn-client-to-site-routes&interface=ui).

## Example Usage

```terraform
data "ibm_is_vpn_server_route" "example" {
	id = ibm_is_vpn_server_route.example.vpn_route
	vpn_server = ibm_is_vpn_server.example.id
}
```

## Argument Reference

Review the argument reference that you can specify for your data source.

- `identifier` - (Optional, String) The VPN route identifier.
- `name` - (Optional, String) The VPN route identifier.
- `vpn_server` - (Required, String) The VPN server identifier.

  ~> **NOTE:** `identifier` and `name` are mutually exclusive.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

- `id` - The unique identifier of the VPNServerRoute and it has format VPNServerID/VPNServerRouteID.
- `action` - (String) The action to perform with a packet matching the VPN route:- `translate`: translate the source IP address to one of the private IP addresses of the VPN server.- `deliver`: deliver the packet into the VPC.- `drop`: drop the packet The enumerated values for this property are expected to expand in the future. When processing this property, check for and log unknown values. Optionally halt processing and surface the error, or bypass the VPN route on which the unexpected property value was encountered.
- `created_at` - (String) The date and time that the VPN route was created.
- `destination` - (String) The destination for this VPN route in the VPN server. If an incoming packet does not match any destination, it will be dropped.
- `href` - (String) The URL for this VPN route.
- `lifecycle_state` - (String) The lifecycle state of the VPN route.
- `name` - (String) The user-defined name for this VPN route.
- `resource_type` - (String) The resource type.

