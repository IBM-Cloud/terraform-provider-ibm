---
layout: "ibm"
page_title: "IBM : ibm_is_vpn_gateway_member"
description: |-
  Get information about VPNGatewayMemberIndividual
subcategory: "Virtual Private Cloud API"
---

# ibm_is_vpn_gateway_member

Provides a read-only data source to retrieve information about a VPNGatewayMemberIndividual. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_is_vpn_gateway_member" "is_vpn_gateway_member" {
	vpn_gateway_member_id = "vpn_gateway_member_id"
	vpn_gateway_id = "vpn_gateway_id"
}
```

## Argument Reference

You can specify the following arguments for this data source.

- `vpn_gateway_member_id` - (Required, Forces new resource, String) The VPN gateway member identifier.
  * Constraints: The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-z_]+$/`.
- `vpn_gateway_id` - (Required, Forces new resource, String) The VPN gateway identifier.
  * Constraints: The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-z_]+$/`.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

- `id` - The unique identifier of the VPNGatewayMemberIndividual.
- `health_reasons` - (List) The reasons for the current `health_state` (if any).
  * Constraints: The minimum length is `0` items.
Nested schema for **health_reasons**:
	- `code` - (String) A reason code for this health state:- `cannot_reserve_ip_address`: IP address exhaustion (release addresses on the VPN's  subnet)- `internal_error`: Internal error (contact IBM support)The enumerated values for this property may[expand](https://cloud.ibm.com/apidocs/vpc#property-value-expansion) in the future.
	  * Constraints: Allowable values are: `cannot_reserve_ip_address`, `internal_error`. The value must match regular expression `/^[a-z][a-z0-9]*(_[a-z0-9]+)*$/`.
	- `message` - (String) An explanation of the reason for this health state.
	- `more_info` - (String) A link to documentation about the reason for this health state.
	  * Constraints: The maximum length is `8000` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
- `health_state` - (String) The health of this resource:- `ok`: No abnormal behavior detected- `degraded`: Experiencing compromised performance, capacity, or connectivity- `faulted`: Completely unreachable, inoperative, or otherwise entirely incapacitated- `inapplicable`: The health state does not apply because of the current lifecycle   state. A resource with a lifecycle state of `failed` or `deleting` will have a   health state of `inapplicable`. A `pending` resource may also have this state.
  * Constraints: Allowable values are: `degraded`, `faulted`, `inapplicable`, `ok`. The value must match regular expression `/^[a-z][a-z0-9]*(_[a-z0-9]+)*$/`.
- `lifecycle_reasons` - (List) The reasons for the current `lifecycle_state` (if any).
  * Constraints: The minimum length is `0` items.
Nested schema for **lifecycle_reasons**:
	- `code` - (String) A reason code for this lifecycle state:- `internal_error`: internal error (contact IBM support)- `resource_suspended_by_provider`: The resource has been suspended (contact IBM  support)The enumerated values for this property may[expand](https://cloud.ibm.com/apidocs/vpc#property-value-expansion) in the future.
	  * Constraints: Allowable values are: `internal_error`, `resource_suspended_by_provider`. The value must match regular expression `/^[a-z][a-z0-9]*(_[a-z0-9]+)*$/`.
	- `message` - (String) An explanation of the reason for this lifecycle state.
	- `more_info` - (String) A link to documentation about the reason for this lifecycle state.
	  * Constraints: The maximum length is `8000` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
- `lifecycle_state` - (String) The lifecycle state of the VPN gateway member.
  * Constraints: Allowable values are: `deleting`, `failed`, `pending`, `stable`, `suspended`, `updating`, `waiting`. The value must match regular expression `/^[a-z][a-z0-9]*(_[a-z0-9]+)*$/`.
- `private_ip` - (List) The reserved IP address assigned to the VPN gateway member.This property will be present only when the VPN gateway status is `available`.
Nested schema for **private_ip**:
	- `address` - (String) The IP address.If the address has not yet been selected, the value will be `0.0.0.0`.This property may [expand](https://cloud.ibm.com/apidocs/vpc#property-value-expansion) to support IPv6 addresses in the future.
	  * Constraints: The maximum length is `15` characters. The minimum length is `7` characters. The value must match regular expression `/^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$/`.
	- `deleted` - (List) If present, this property indicates the referenced resource has been deleted, and providessome supplementary information.
	Nested schema for **deleted**:
		- `more_info` - (String) A link to documentation about deleted resources.
		  * Constraints: The maximum length is `8000` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
	- `href` - (String) The URL for this reserved IP.
	  * Constraints: The maximum length is `8000` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
	- `id` - (String) The unique identifier for this reserved IP.
	  * Constraints: The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-z_]+$/`.
	- `name` - (String) The name for this reserved IP. The name is unique across all reserved IPs in a subnet.
	  * Constraints: The maximum length is `63` characters. The minimum length is `1` character. The value must match regular expression `/^-?([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$/`.
	- `resource_type` - (String) The resource type.
	  * Constraints: Allowable values are: `subnet_reserved_ip`. The value must match regular expression `/^[a-z][a-z0-9]*(_[a-z0-9]+)*$/`.
	- `subnet` - (List)
	Nested schema for **subnet**:
		- `crn` - (String) The CRN for this subnet.
		  * Constraints: The maximum length is `512` characters. The minimum length is `17` characters. The value must match regular expression `/^crn:v[0-9]+:[a-z0-9-]+:[a-z0-9-]+:[a-z0-9-]+:[a-z0-9-]*:([a-z]\/[a-z0-9-]+)?:[a-z0-9-]*:[a-z0-9-]*:[a-zA-Z0-9-_\\.\/]*$/`.
		- `deleted` - (List) If present, this property indicates the referenced resource has been deleted, and providessome supplementary information.
		Nested schema for **deleted**:
			- `more_info` - (String) A link to documentation about deleted resources.
			  * Constraints: The maximum length is `8000` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
		- `href` - (String) The URL for this subnet.
		  * Constraints: The maximum length is `8000` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
		- `id` - (String) The unique identifier for this subnet.
		  * Constraints: The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^[-0-9a-z_]+$/`.
		- `name` - (String) The name for this subnet. The name is unique across all subnets in the VPC.
		  * Constraints: The maximum length is `63` characters. The minimum length is `1` character. The value must match regular expression `/^-?([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$/`.
		- `resource_type` - (String) The resource type.
		  * Constraints: Allowable values are: `subnet`. The value must match regular expression `/^[a-z][a-z0-9]*(_[a-z0-9]+)*$/`.
- `public_ip` - (List) The public IP address assigned to the VPN gateway member.
Nested schema for **public_ip**:
	- `address` - (String) The IP address.This property may [expand](https://cloud.ibm.com/apidocs/vpc#property-value-expansion) to support IPv6 addresses in the future.
	  * Constraints: The maximum length is `15` characters. The minimum length is `7` characters. The value must match regular expression `/^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$/`.
- `role` - (String) The high availability role assigned to the VPN gateway member.The enumerated values for this property may[expand](https://cloud.ibm.com/apidocs/vpc#property-value-expansion) in the future.
  * Constraints: Allowable values are: `active`, `standby`. The value must match regular expression `/^[a-z][a-z0-9]*(_[a-z0-9]+)*$/`.

