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
  vpn_server    = ibm_is_vpn_server.example.vpn_server
  destination   = "172.16.0.0/16"
  action        = "translate"
  name          = "example-vpn-server-route"
}
```

## Argument Reference

Review the argument reference that you can specify for your resource.

- `action` - (Optional, String) The action to perform with a packet matching the VPN route.

 ~> **Note:** </br> Allowed values are : </br>
  **&#x2022;** `translate`: translate the source IP address to one of the private IP addresses of the VPN server, then deliver the packet to target.</br>
  **&#x2022;** `deliver`: deliver the packet to the target.</br>
  **&#x2022;** `drop`: drop the packet. The enumerated values for this property are expected to expand in the future.</br>
- `destination` - (Required, String) The destination to use for this VPN route in the VPN server. Must be unique within the VPN server. If an incoming packet does not match any destination, it will be dropped.
- `name` - (Optional, String) The user-defined name for this VPN route. If unspecified, the name will be a hyphenated list of randomly-selected words.Names must be unique within the VPN server the VPN route resides in.
- `vpn_server` - (Required, Forces new resource,String) The VPN server identifier.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your resource is created.

- `id` - The unique identifier of the VPNServerRoute and it has format VPNServerID/VPNServerRouteID.
- `vpn_route` - The identifier of the VPNServerRoute.
- `created_at` - (String) The date and time that the VPN route was created.
- `href` - (String) The URL for this VPN route.
- `health_reasons` - (List) The reasons for the current health_state (if any).

  Nested scheme for `health_reasons`:
  - `code` - (String) A snake case string succinctly identifying the reason for this health state.
  - `message` - (String) An explanation of the reason for this health state.
  - `more_info` - (String) Link to documentation about the reason for this health state.
- `health_state` - (String) The health of this resource.

  -> **Supported health_state values:** 
    </br>&#x2022; `ok`: Healthy
    </br>&#x2022; `degraded`: Suffering from compromised performance, capacity, or connectivity
    </br>&#x2022; `faulted`: Completely unreachable, inoperative, or otherwise entirely incapacitated
    </br>&#x2022; `inapplicable`: The health state does not apply because of the current lifecycle state. 
      **Note:** A resource with a lifecycle state of `failed` or `deleting` will have a health state of `inapplicable`. A `pending` resource may also have this state.
- `lifecycle_state` - (String) The lifecycle state of the VPN route.
- `lifecycle_reasons` - (List) The reasons for the current lifecycle_reasons (if any).

  Nested scheme for `lifecycle_reasons`:
  - `code` - (String) A snake case string succinctly identifying the reason for this lifecycle reason.
  - `message` - (String) An explanation of the reason for this lifecycle reason.
  - `more_info` - (String) Link to documentation about the reason for this lifecycle reason.
- `resource_type` - (String) The resource type.

## Import

You can import the `ibm_is_vpn_server_route` resource by using `id`.
The `id` property can be formed from `vpn_server`, and `vpn_route` in the following format:
- `vpn_server`: A string. The VPN server identifier.
- `vpn_route`: A string. The VPN route identifier.

```
 r134-0b5b3bed-8c95-4ded-81bf-913bf8ec5fa9/r134-299ed4f0-b279-4db3-8dfe-eff69a9fb66a
```

# Syntax
```
$ terraform import ibm_is_vpn_server_route.is_vpn_server_route r134-0b5b3bed-8c95-4ded-81bf-913bf8ec5fa9/r134-299ed4f0-b279-4db3-8dfe-eff69a9fb66a
```
