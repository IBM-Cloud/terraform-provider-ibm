---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : "
description: |-
  Manages IBM VPN gateway connections.
---

# ibm_is_vpn_gateway_connections
Retrieve information of an existing VPN gateway connections. For more information, see [adding connections to a VPN gateway](https://cloud.ibm.com/docs/vpc?topic=vpc-vpn-adding-connections).

**Note:** 
VPC infrastructure services are a regional specific based endpoint, by default targets to `us-south`. Please make sure to target right region in the provider block as shown in the `provider.tf` file, if VPC service is created in region other than `us-south`.

**provider.tf**

```terraform
provider "ibm" {
  region = "eu-gb"
}
```

## Example usage

```terraform

data "ibm_is_vpn_gateway_connections" "example" {
  vpn_gateway = ibm_is_vpn_gateway.example.id
}

```

## Argument reference
Review the argument references that you can specify for your data source. 

- `status` - (Optional, String) Filters the collection to VPN gateway connections with the specified status.
- `vpn_gateway` - (Required, String) The VPN gateway ID.

## Attribute reference
In addition to all argument reference list, you can access the following attribute references after your data source is created. 

- `action` - (String) Action detection for dead peer detection action.
- `admin_state_up` - (String) The VPN gateway connection admin state. Default value is **true**.
- `authentication_mode` - (String) The authentication mode.
- `created_at`- (Timestamp) The date and time the VPN gateway connection was created.
- `id` - (String) The ID of the VPN gateway connection.
- `ike_policy` - (String) The VPN gateway connection IKE Policy.
- `interval`-  (String) Interval for dead peer detection.
- `ipsec_policy` - (String) The IP security policy VPN gateway connection.
- `local_cidrs` - (String) The VPN gateway connection local CIDRs.
- `mode` - (String) The mode of the VPN gateway.
- `name`-  (String) The VPN gateway connection name.
- `peer_address` - (String) The VPN gateway connection peer address.
- `peer_cidrs` - (String) The VPN gateway connection peer CIDRs.
- `resource_type` - (String) The resource type.
- `timeout` - (String) Timeout for dead peer detection.
- `tunnels` - (List) The VPN tunnel configuration for the VPN gateway connection (in static route mode).

  Nested scheme for `tunnels`:
	- `address` - (String) The IP address of the VPN gateway member in which the tunnel resides.
	- `status` - (String) The status of the VPN tunnel.
- `status_reasons` - (List) Array of reasons for the current status (if any).

  Nested `status_reasons`:
  - `code` - (String) The status reason code.
  - `message` - (String) An explanation of the status reason.
  - `more_info` - (String) Link to documentation about this status reason.
