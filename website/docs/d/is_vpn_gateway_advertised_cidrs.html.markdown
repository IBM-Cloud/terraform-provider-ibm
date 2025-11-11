---
layout: "ibm"
page_title: "IBM : ibm_is_vpn_gateway_advertised_cidrs"
description: |-
  Get information about VPNGatewayAdvertisedCIDRs
subcategory: "VPC infrastructure"
---

# ibm_is_vpn_gateway_advertised_cidrs

Provides a read-only data source to retrieve information about VPNGatewayAdvertisedCIDRs. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_is_vpn_gateway_advertised_cidrs" "example" {
  vpn_gateway = ibm_is_vpn_gateway.example.id
}

data "ibm_is_vpn_gateway_advertised_cidrs" "example-2" {
  vpn_gateway_name = ibm_is_vpn_gateway.example.name
}
```

## Argument Reference

You can specify the following arguments for this data source.

- `vpn_gateway` - (Optional, String) The VPN gateway identifier.
- `vpn_gateway_name` - (Optional, String) The VPN gateway name.

  ~> **Note** Provide either one of `vpn_gateway`, `vpn_gateway_name` to identifiy vpn gateway 

## Attribute Reference

After your data source is created, you can read values from the following attributes.

- `advertised_cidrs` - (List) The additional CIDRs advertised through any enabled routing protocol (for example, BGP). The routing protocol will advertise routes with these CIDRs and VPC prefixes as route destinations.

