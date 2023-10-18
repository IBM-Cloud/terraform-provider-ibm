---
layout: "ibm"
page_title: "IBM : ibm_is_vpn_server_routes"
description: |-
  Get information about VPNServerRouteCollection
subcategory: "VPC infrastructure"
---

# ibm_is_vpn_server_routes

Provides a read-only data source for VPNServerRouteCollection. For more information, about VPN Server Routes, see [Managing VPN Server routes](https://cloud.ibm.com/docs/vpc?topic=vpc-vpn-client-to-site-routes&interface=ui).

## Example Usage

```terraform
data "ibm_is_vpn_server_routes" "example" {
	vpn_server = ibm_is_vpn_server.example.id
}
```

## Argument Reference

Review the argument reference that you can specify for your data source.

- `vpn_server` - (Required, String) The VPN server identifier.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

- `id` - The unique identifier of the VPNServerRouteCollection.
- `routes` - (List) Collection of VPN routes.
	Nested scheme for `routes`:
	- `action` - (String) The action to perform with a packet matching the VPN route:- `translate`: translate the source IP address to one of the private IP addresses of the VPN server.- `deliver`: deliver the packet into the VPC.- `drop`: drop the packet The enumerated values for this property are expected to expand in the future. When processing this property, check for and log unknown values. Optionally halt processing and surface the error, or bypass the VPN route on which the unexpected property value was encountered.
	- `created_at` - (String) The date and time that the VPN route was created.
	- `destination` - (String) The destination for this VPN route in the VPN server. If an incoming packet does not match any destination, it will be dropped.
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
      		</br>**Note:** A resource with a lifecycle state of `failed` or `deleting` will have a health state of `inapplicable`. A `pending` resource may also have this state.
	- `href` - (String) The URL for this VPN route.
	- `id` - (String) The unique identifier for this VPN route.
	- `lifecycle_reasons` - (List) The reasons for the current lifecycle_reasons (if any).

  		Nested scheme for `lifecycle_reasons`:
  		- `code` - (String) A snake case string succinctly identifying the reason for this lifecycle reason.
  		- `message` - (String) An explanation of the reason for this lifecycle reason.
  		- `more_info` - (String) Link to documentation about the reason for this lifecycle reason.
	- `lifecycle_state` - (String) The lifecycle state of the VPN route.
	- `name` - (String) The user-defined name for this VPN route.
	- `resource_type` - (String) The resource type.
