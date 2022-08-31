---

subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : ibm_is_vpn_server_route"
description: |-
  Manages IBM VPN Server Route.
---

# ibm_is_vpn_server_route

Provides a resource for VPNServerRoute. This allows VPNServerRoute to be created, updated and deleted. For more information, about VPN Server Routes, see [Managing VPN Server routes](https://cloud.ibm.com/docs/vpc?topic=vpc-vpn-client-to-site-routes&interface=ui).
## Example Usage

```terraform
resource "ibm_is_vpn_server_route" "example" {
  vpn_server_id = ibm_is_vpn_server.example.vpn_server
  destination   = "172.16.0.0/16"
  action        = "translate"
  name          = "example-vpn-server-route"
}
```

## Argument Reference

Review the argument reference that you can specify for your resource.

- `action` - (Optional, String) The action to perform with a packet matching the VPN route:
  - `translate`: translate the source IP address to one of the private IP addresses of the VPN server, then deliver the packet to target.
  - `deliver`: deliver the packet to the target.
  - `drop`: drop the packet. The enumerated values for this property are expected to expand in the future. 
  When processing this property, check for and log unknown values. Optionally halt processing and surface the error, or bypass the VPN route on which the unexpected property value was encountered.
  - Constraints: The default value is `deliver`. Allowable values are: translate, deliver, drop
- `destination` - (Required, String) The destination to use for this VPN route in the VPN server. Must be unique within the VPN server. If an incoming packet does not match any destination, it will be dropped.
  - Constraints: The value must match regular expression `/^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])(\/(3[0-2]|[1-2][0-9]|[0-9]))$/`
- `name` - (Optional, String) The user-defined name for this VPN route. If unspecified, the name will be a hyphenated list of randomly-selected words. Names must be unique within the VPN server the VPN route resides in.
  - Constraints: The maximum length is `63` characters. The minimum length is `1` character. The value must match regular expression `/^([a-z]|[a-z][-a-z0-9]*[a-z0-9])$/`
- `vpn_server_id` - (Required, Forces new resource, String) The VPN server identifier.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your resource is created.

- `vpn_route` - The identifier of the VPNServerRoute.
- `created_at` - (String) The date and time that the VPN route was created.
- `href` - (String) The URL for this VPN route.
  - Constraints: The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`
- `lifecycle_state` - (String) The lifecycle state of the VPN route.
  - Constraints: Allowable values are: deleting, failed, pending, stable, updating, waiting, suspended
- `resource_type` - (String) The resource type.
  - Constraints: Allowable values are: vpn_server_route

## Import

You can import the `ibm_is_vpn_server_route` resource by using `id`.
The `id` property can be formed from `vpn_server_id`, and `vpn_route` in the following format:
- `vpn_server_id`: A string. The VPN server identifier.
- `vpn_route`: A string. The VPN route identifier.

```
 r134-0b5b3bed-8c95-4ded-81bf-913bf8ec5fa9/r134-299ed4f0-b279-4db3-8dfe-eff69a9fb66a
```

# Syntax
```
$ terraform import ibm_is_vpn_server_route.is_vpn_server_route r134-0b5b3bed-8c95-4ded-81bf-913bf8ec5fa9/r134-299ed4f0-b279-4db3-8dfe-eff69a9fb66a
```
