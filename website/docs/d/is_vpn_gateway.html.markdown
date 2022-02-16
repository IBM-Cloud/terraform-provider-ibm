---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : ibm_is_vpn_gateway"
description: |-
  Get information about IBM Cloud VPN Gateway
---

# ibm_is_vpn_gateway

Provides a read-only data source for VPN Gateway. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "ibm_is_vpn_gateway" "example" {
  vpn_gateway = ibm_is_vpn_gateway.example.id
}

data "ibm_is_vpn_gateway" "example-1" {
  vpn_gateway_name = ibm_is_vpn_gateway.example.name
}
```

## Argument Reference

Review the argument reference that you can specify for your data source.

- `vpn_gateway` - (Optional, String) The VPN gateway identifier.
- `vpn_gateway_name` - (Optional, String) The VPN gateway name.
  ~> **Note** Provide either `vpn_gateway` or `vpn_gateway_name`

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

- `id` - The unique identifier of the is_vpn_gateway.
- `connections` - (List) Connections for this VPN gateway.
  Nested scheme for **connections**:
	- `deleted` - (List) If present, this property indicates the referenced resource has been deleted and provides some supplementary information.
	  Nested scheme for **deleted**:
		- `more_info` - (String) Link to documentation about deleted resources.
	- `href` - (String) The VPN connection's canonical URL.
	- `id` - (String) The unique identifier for this VPN gateway connection.
	- `name` - (String) The user-defined name for this VPN connection.
	- `resource_type` - (String) The resource type.

- `created_at` - (String) The date and time that this VPN gateway was created.

- `crn` - (String) The VPN gateway's CRN.

- `href` - (String) The VPN gateway's canonical URL.

- `members` - (List) Collection of VPN gateway members.
  Nested scheme for **members**:
	- `private_ip_address` - (String) The private IP address assigned to the VPN gateway member. This property will be present only when the VPN gateway status is `available`. This property may add support for IPv6 addresses in the future. When processing a value in this property, verify that the address is in an expected format. If it is not, log an error. Optionally halt processing and surface the error, or bypass the resource on which the unexpected IP address format was encountered.
	- `public_ip_address` - (String) The public IP address assigned to the VPN gateway member. This property may add support for IPv6 addresses in the future. When processing a value in this property, verify that the address is in an expected format. If it is not, log an error. Optionally halt processing and surface the error, or bypass the resource on which the unexpected IP address format was encountered.
	- `role` - (String) The high availability role assigned to the VPN gateway member.
	- `status` - (String) The status of the VPN gateway member.

- `mode` - (String) Route mode VPN gateway.

- `name` - (String) The user-defined name for this VPN gateway.

- `resource_group` - (List) The resource group for this VPN gateway.
  Nested scheme for **resource_group**:
	- `href` - (String) The URL for this resource group.
	- `id` - (String) The unique identifier for this resource group.
	- `name` - (String) The user-defined name for this resource group.

- `resource_type` - (String) The resource type.

- `status` - (String) The status of the VPN gateway.

- `subnet` - (List) 
  Nested scheme for **subnet**:
	- `crn` - (String) The CRN for this subnet.
	- `deleted` - (List) If present, this property indicates the referenced resource has been deleted and provides some supplementary information.
	  Nested scheme for **deleted**:
		- `more_info` - (String) Link to documentation about deleted resources.
	- `href` - (String) The URL for this subnet.
	- `id` - (String) The unique identifier for this subnet.
	- `name` - (String) The user-defined name for this subnet.

