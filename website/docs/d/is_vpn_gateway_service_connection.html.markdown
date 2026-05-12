---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : ibm_is_vpn_gateway_service_connection"
description: |-
  Get information about IBM Cloud VPN Connection
---

# ibm_is_vpn_gateway_service_connection

Provides a read-only data source for VPN gateway service Connection. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "ibm_is_vpn_gateway_service_connection" "example" {
  vpn_gateway = ibm_is_vpn_gateway.example.id
  vpn_gateway_service_connection = "3066f374-97f7-4138-b59d-20a8414f49a8"
}
data "ibm_is_vpn_gateway_service_connection" "example-1" {
  vpn_gateway_name = ibm_is_vpn_gateway.example.name
  vpn_gateway_service_connection = "3066f374-97f7-4138-b59d-20a8414f49a8"
}
```

## Argument Reference

Review the argument reference that you can specify for your data source.

- `vpn_gateway` - (Optional, String) The VPN gateway identifier.
- `vpn_gateway_name` - (Optional, String) The VPN gateway name.
- `vpn_gateway_service_connection` - (Required, String) The VPN gateway service connection identifier.

  ~> **Note** Provide either one of `vpn_gateway`, `vpn_gateway_name` to identifiy vpn gateway.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

- `created_at` - (String) The date and time that this VPN gateway connection was created.
- `creator` - (List)
  Nested scheme for **creator**:
	- `crn` - (String) The CRN for this transit gateway.
	- `id` - (String) The unique identifier for this transit gateway.
	- `resource_type` - (String) The resource type.
- `id` - The unique identifier for this VPN gateway service connection.
- `lifecycle_reasons` - (List) The reasons for the current lifecycle_state (if any).
  Nested scheme for **lifecycle_reasons**:
	- `code` - (String) A reason code for this lifecycle state:- `internal_error`: internal error (contact IBM support)- `resource_suspended_by_provider`: The resource has been suspended (contact IBM  support)The enumerated values for this property may[expand](https://cloud.ibm.com/apidocs/vpc#property-value-expansion) in the future.
	- `message` - (String) An explanation of the reason for this lifecycle state.
	- `message` - (String) Link to documentation about the reason for this lifecycle state.
- `lifecycle_state` - (List) The lifecycle state of the VPN service connection.
- `status` - (String) The status of this service connection.
- `status_reasons` - (List) The reasons for the current VPN gateway service connection status (if any).
  Nested `status_reasons`:
    - `code` - (String) The status reason code.
    - `message` - (String) An explanation of the reason for this VPN service connection's status.
    - `more_info` - (String) Link to documentation about this status reason


