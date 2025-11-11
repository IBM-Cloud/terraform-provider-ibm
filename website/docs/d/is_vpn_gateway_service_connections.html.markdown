---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : "
description: |-
  Manages IBM VPN gateway connections.
---

# ibm_is_vpn_gateway_service_connections
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

data "ibm_is_vpn_gateway_service_connections" "example" {
  vpn_gateway = ibm_is_vpn_gateway.example.id
}

```

## Argument reference
Review the argument references that you can specify for your data source. 

- `vpn_gateway` - (Required, String) The VPN gateway ID.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.
- `service_connections` - (List) List VPN gateway service connections.

	- `created_at` - (String) The date and time that this VPN gateway connection was created.
	- `creator` - (List) 
	Nested scheme for **creator**:
		- `crn` - (String) The CRN for this transit gateway.
		- `id` - (String) The unique identifier for this transit gateway.
		- `resource_type` - (String) The resource type.
	- `id` - The unique identifier for this VPN gateway service connection.
	- `lifecycle_reasons` - (List) The reasons for the current lifecycle_state (if any).
	Nested scheme for **lifecycle_reasons**:
		- `code` - (String) A reason code for this lifecycle state:- `internal_error`: internal error (contact IBM support)- `resource_suspended_by_provider`: The resource has been suspended (contact IBM  support)The enumerated values for this property may  [expand](https://cloud.ibm.com/apidocs/vpc#property-value-expansion) in the future.
		- `message` - (String) An explanation of the reason for this lifecycle state.
		- `message` - (String) Link to documentation about the reason for this lifecycle state.
	- `lifecycle_state` - (List) The lifecycle state of the VPN service connection.
	- `status` - (String) The status of this service connection.
	- `status_reasons` - (List) The reasons for the current VPN gateway service connection status (if any).
	Nested `status_reasons`:
		- `code` - (String) The status reason code.
		- `message` - (String) An explanation of the reason for this VPN service connection's status.
		- `more_info` - (String) Link to documentation about this status reason


