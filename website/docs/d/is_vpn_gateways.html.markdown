---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : is_vpn_gateways"
description: |-
  Manages IBM Cloud VPN gateways.
---

# ibm_is_vpn_gateways
Retrieve information of an existing VPN gateways. For more information, about IBM Cloud VPN gateways, see [configuring ACLs and security groups for use with VPN](https://cloud.ibm.com/docs/vpc?topic=vpc-acls-security-groups-vpn).

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

data "ibm_is_vpn_gateways" "example" {
}

```
## Argument reference

Review the argument references that you can specify for your data source. 

- `resource_group` - (Optional, String) The ID of the Resource group this vpn gateway belongs to
- `mode` - (Optional, String) The mode of this VPN Gateway. Available options are `policy` and `route`.

## Attribute reference
In addition to all argument reference list, you can access the following attribute references after your data source is created. 

- `vpn_gateways` - (List) Collection of VPN Gateways.

  Nested scheme for `vpn_gateways`:
  - `access_tags`  - (List) Access management tags associated for the vpn gateway.
	- `availability_mode` - (String) The availability mode of the VPN gateway:- `zonal`: The availability of this VPN gateway is limited only to a single zone of a  given region as provided by the `zone` of the VPN gateway.
	  * Constraints: Allowable values are: `zonal`. 
  - `crn` - (String) The VPN gateway's CRN.
  - `created_at`- (Timestamp) The date and time the VPN gateway was created.
  - `id` - (String) The ID of the VPN gateway.
  - `advertised_cidrs` - (Optional, List) The additional CIDRs advertised through any enabled routing protocol (for example, BGP). The routing protocol will advertise routes with these CIDRs and VPC prefixes as route destinations.
  - `name`-  (String) The VPN gateway instance name.
	- `members` - (List) The members for the VPN gateway.
	  Nested schema for **members**:
    - `address` - (String) The public IP address assigned to the VPN gateway member.</br>
    - `role`-  (String) The high availability role assigned to the VPN gateway member.</br>
		- `health_reasons` - (List) The reasons for the current `health_state` (if any).
		  Nested schema for **health_reasons**:
			- `code` - (String) A reason code for this health state:- `cannot_reserve_ip_address`: IP address exhaustion (release addresses on the VPN's  subnet)- `internal_error`: Internal error (contact IBM support)The enumerated values for this property may[expand](https://cloud.ibm.com/apidocs/vpc#property-value-expansion) in the future.
			  * Constraints: Allowable values are: `cannot_reserve_ip_address`, `internal_error`.
			- `message` - (String) An explanation of the reason for this health state.
			- `more_info` - (String) A link to documentation about the reason for this health state.
		- `health_state` - (String) The health of this resource:- `ok`: No abnormal behavior detected- `degraded`: Experiencing compromised performance, capacity, or connectivity- `faulted`: Completely unreachable, inoperative, or otherwise entirely incapacitated- `inapplicable`: The health state does not apply because of the current lifecycle   state. A resource with a lifecycle state of `failed` or `deleting` will have a   health state of `inapplicable`. A `pending` resource may also have this state.
		  * Constraints: Allowable values are: `degraded`, `faulted`, `inapplicable`, `ok`
		- `id` - (String) The unique identifier for this VPN gateway member.
		- `lifecycle_reasons` - (List) The reasons for the current `lifecycle_state` (if any).
		  Nested schema for **lifecycle_reasons**:
			- `code` - (String) A reason code for this lifecycle state:- `internal_error`: internal error (contact IBM support)- `resource_suspended_by_provider`: The resource has been suspended (contact IBM  support)The enumerated values for this property may[expand](https://cloud.ibm.com/apidocs/vpc#property-value-expansion) in the future.
			  * Constraints: Allowable values are: `internal_error`, `resource_suspended_by_provider`.
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
	  - `private_address` - (String) The private IP address assigned to the VPN gateway member. Same as `private_ip.0.address`.</br>
  

  - `resource_type` - (String) The resource type, supported value is `vpn_gateway`.
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
  - `lifecycle_reasons` - (List) The reasons for the current lifecycle_reasons (if any).

      Nested scheme for `lifecycle_reasons`:
      - `code` - (String) A snake case string succinctly identifying the reason for this lifecycle reason.
      - `message` - (String) An explanation of the reason for this lifecycle reason.
      - `more_info` - (String) Link to documentation about the reason for this lifecycle reason.
  - `lifecycle_state` - (String) The lifecycle state of the VPN gateway.
  - `local_asn` - (Integer) The local autonomous system number (ASN) for this VPN gateway and its connections.
  - `subnet` - (String) The VPN gateway subnet information.
  - `tags`- (Optional, Array of Strings) A list of tags associated with the instance.
  - `vpc` - (String) 	The VPC this VPN server resides in.
  
      Nested scheme for `vpc`:
      - `crn` - (String) The CRN for this VPC.
      - `deleted` - (List) 	If present, this property indicates the referenced resource has been deleted and provides some supplementary information.
        Nested scheme for **deleted**:
        - `more_info` - (String) Link to documentation about deleted resources.
      - `href` - (String) - The URL for this VPC
      - `id` - (String) - The unique identifier for this VPC.
      - `name` - (String) - The unique user-defined name for this VPC.
  - `resource_group` - (String) The resource group ID.
  - `mode` - (String) The VPN gateway mode, supported values are `policy` and `route`.