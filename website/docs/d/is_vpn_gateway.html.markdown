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
- `access_tags`  - (List) Access management tags associated for the vpn gateway.
- `advertised_cidrs` - (Optional, List) The additional CIDRs advertised through any enabled routing protocol (for example, BGP). The routing protocol will advertise routes with these CIDRs and VPC prefixes as route destinations.
- `availability_mode` - (String) The availability mode of the VPN gateway:- `zonal`: The availability of this VPN gateway is limited only to a single zone of a  given region as provided by the `zone` of the VPN gateway.
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
- `local_asn` - (Integer) The local autonomous system number (ASN) for this VPN gateway and its connections.
- `members` - (List) The members for the VPN gateway.
	Nested schema for **members**:
	- `private_ip_address` - (String) The private IP address assigned to the VPN gateway member. This property will be present only when the VPN gateway status is `available`. This property may add support for IPv6 addresses in the future. When processing a value in this property, verify that the address is in an expected format. If it is not, log an error. Optionally halt processing and surface the error, or bypass the resource on which the unexpected IP address format was encountered. Same as `primary_ip.0.address`
	- `public_ip_address` - (String) The public IP address assigned to the VPN gateway member. This property may add support for IPv6 addresses in the future. When processing a value in this property, verify that the address is in an expected format. If it is not, log an error. Optionally halt processing and surface the error, or bypass the resource on which the unexpected IP address format was encountered.
	- `health_reasons` - (List) The reasons for the current `health_state` (if any).
		Nested schema for **health_reasons**:
		- `code` - (String) A reason code for this health state:- `cannot_reserve_ip_address`: IP address exhaustion (release addresses on the VPN's  subnet)- `internal_error`: Internal error (contact IBM support)The enumerated values for this property may[expand](https://cloud.ibm.com/apidocs/vpc#property-value-expansion) in the future.
		  * Constraints: Allowable values are: `cannot_reserve_ip_address`, `internal_error`.
		- `message` - (String) An explanation of the reason for this health state.
		- `more_info` - (String) A link to documentation about the reason for this health state.
	- `health_state` - (String) The health of this resource:- `ok`: No abnormal behavior detected- `degraded`: Experiencing compromised performance, capacity, or connectivity- `faulted`: Completely unreachable, inoperative, or otherwise entirely incapacitated- `inapplicable`: The health state does not apply because of the current lifecycle   state. A resource with a lifecycle state of `failed` or `deleting` will have a   health state of `inapplicable`. A `pending` resource may also have this state.
	  * Constraints: Allowable values are: `degraded`, `faulted`, `inapplicable`, `ok`. 
	- `id` - (String) The unique identifier for this VPN gateway member.
	- `lifecycle_reasons` - (List) The reasons for the current `lifecycle_state` (if any).
		Nested schema for **lifecycle_reasons**:
		- `code` - (String) A reason code for this lifecycle state:- `internal_error`: internal error (contact IBM support)- `resource_suspended_by_provider`: The resource has been suspended (contact IBM  support)The enumerated values for this property may[expand](https://cloud.ibm.com/apidocs/vpc#property-value-expansion) in the future.
		  * Constraints: Allowable values are: `internal_error`, `resource_suspended_by_provider`. The value must match regular expression `/^[a-z][a-z0-9]*(_[a-z0-9]+)*$/`.
		- `message` - (String) An explanation of the reason for this lifecycle state.
		- `more_info` - (String) A link to documentation about the reason for this lifecycle state.
	- `lifecycle_state` - (String) The lifecycle state of the VPN gateway member.
	  * Constraints: Allowable values are: `deleting`, `failed`, `pending`, `stable`, `suspended`, `updating`, `waiting`.
	- `private_ip` - (List) The reserved IP address assigned to the VPN gateway member.This property will be present only when the VPN gateway status is `available`.
		Nested schema for **private_ip**:
		- `address` - (String) The IP address.If the address has not yet been selected, the value will be `0.0.0.0`.This property may [expand](https://cloud.ibm.com/apidocs/vpc#property-value-expansion) to support IPv6 addresses in the future.
		- `deleted` - (List) If present, this property indicates the referenced resource has been deleted, and providessome supplementary information.
			Nested schema for **deleted**:
			- `more_info` - (String) A link to documentation about deleted resources.
		- `href` - (String) The URL for this reserved IP.
		- `id` - (String) The unique identifier for this reserved IP.
		- `reserved_ip`- (String) The unique identifier for this reserved IP
		- `name` - (String) The name for this reserved IP. The name is unique across all reserved IPs in a subnet.
		- `resource_type` - (String) The resource type.
		- `subnet` - (List)
			Nested schema for **subnet**:
			- `crn` - (String) The CRN for this subnet.
			- `deleted` - (List) If present, this property indicates the referenced resource has been deleted, and providessome supplementary information.
				Nested schema for **deleted**:
				- `more_info` - (String) A link to documentation about deleted resources.
			- `href` - (String) The URL for this subnet.
			- `id` - (String) The unique identifier for this subnet.
			- `name` - (String) The name for this subnet. The name is unique across all subnets in the VPC.
			- `resource_type` - (String) The resource type.
	- `public_ip` - (List) The public IP address assigned to the VPN gateway member.
		Nested schema for **public_ip**:
		- `address` - (String) The IP address.This property may [expand](https://cloud.ibm.com/apidocs/vpc#property-value-expansion) to support IPv6 addresses in the future.
	- `role` - (String) The high availability role assigned to the VPN gateway member.The enumerated values for this property may[expand](https://cloud.ibm.com/apidocs/vpc#property-value-expansion) in the future.
- `mode` - (String) Route mode VPN gateway.
- `name` - (String) The user-defined name for this VPN gateway.

- `resource_group` - (List) The resource group object, for this VPN gateway.
  Nested scheme for **resource_group**:
	- `href` - (String) The URL for this resource group.
	- `id` - (String) The unique identifier for this resource group.
	- `name` - (String) The user-defined name for this resource group.

- `resource_type` - (String) The resource type.

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
- `lifecycle_reasons` - (List) The reasons for the current lifecycle_reasons (if any).

  Nested scheme for `lifecycle_reasons`:
  - `code` - (String) A snake case string succinctly identifying the reason for this lifecycle reason.
  - `message` - (String) An explanation of the reason for this lifecycle reason.
  - `more_info` - (String) Link to documentation about the reason for this lifecycle reason.
- `lifecycle_state` - (String) The lifecycle state of the VPN gateway.
- `subnet` - (List) 
  Nested scheme for **subnet**:
	- `crn` - (String) The CRN for this subnet.
	- `deleted` - (List) If present, this property indicates the referenced resource has been deleted and provides some supplementary information.
	  Nested scheme for **deleted**:
		- `more_info` - (String) Link to documentation about deleted resources.
	- `href` - (String) The URL for this subnet.
	- `id` - (String) The unique identifier for this subnet.
	- `name` - (String) The user-defined name for this subnet.
- `tags`- (Optional, Array of Strings) A list of tags associated with the instance.
- `vpc` - (String) The VPC this VPN server resides in.
  Nested scheme for `vpc`:
  - `crn` - (String) The CRN for this VPC.
  - `deleted` - (List) 	If present, this property indicates the referenced resource has been deleted and provides some supplementary information.
	  Nested scheme for **deleted**:
		- `more_info` - (String) Link to documentation about deleted resources.
  - `href` - (String) - The URL for this VPC
  - `id` - (String) - The unique identifier for this VPC.
  - `name` - (String) - The unique user-defined name for this VPC.
- `resource_type` - (String) - The resource type.

