---
layout: "ibm"
page_title: "IBM : ibm_is_vpn_gateway_connection_local_cidrs"
description: |-
  Get information about VPNGatewayConnectionCIDRs
subcategory: "VPC infrastructure"
---

# ibm_is_vpn_gateway_connection_local_cidrs

Provides a read-only data source to retrieve information about VPNGatewayConnectionCIDRs. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```terraform
data "ibm_is_vpn_gateway_connection_local_cidrs" "is_vpn_gateway_connection_cidrs" {
	vpn_gateway_connection = "vpn_gateway_connection"
	vpn_gateway = "vpn_gateway"
}
```

## Argument Reference

You can specify the following arguments for this data source.

- `vpn_gateway_connection` - (Required, Forces new resource, String) The VPN gateway connection identifier.
- `vpn_gateway` - (Required, Forces new resource, String) The VPN gateway identifier.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

- `id` - The unique identifier of the VPNGatewayConnectionCIDRs.
- `cidrs` - (List) The CIDRs for this resource.

