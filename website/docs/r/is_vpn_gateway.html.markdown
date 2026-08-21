---

subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : VPN-gateway"
description: |-
  Manages IBM VPN gateway.
---

# ibm_is_vpn_gateway
Create, update, or delete a VPN gateway. For more information, about VPN gateway, see [adding connections to a VPN gateway](https://cloud.ibm.com/docs/vpc?topic=vpc-vpn-adding-connections).

**Note:** 
VPC infrastructure services are a regional specific based endpoint, by default targets to `us-south`. Please make sure to target right region in the provider block as shown in the `provider.tf` file, if VPC service is created in region other than `us-south`.

**provider.tf**

```terraform
provider "ibm" {
  region = "eu-gb"
}
```

## Example usage
The following example creates a VPN gateway:

```terraform

resource "ibm_is_vpc" "example" {
  name = "example-vpc"
}

resource "ibm_is_subnet" "example" {
  name            = "example-subnet"
  vpc             = ibm_is_vpc.example.id
  zone            = "us-south-1"
  ipv4_cidr_block = "10.240.0.0/24"
}

resource "ibm_is_vpn_gateway" "example" {
  name      = "example-vpn-gateway"
  subnet    = ibm_is_subnet.example.id
  mode      = "route"
  local_asn = 64520
}

```

## Timeouts
The `ibm_is_vpn_gateway` resource provides the following [Timeouts](https://www.terraform.io/docs/language/resources/syntax.html) configuration options:

- **create**: The creation of the VPN gateway is considered `failed` when no response is received for 10 minutes. 
- **delete**: The deletion of the VPN gateway is considered `failed` when no response is received for 10 minutes. 


## Argument reference
Review the argument references that you can specify for your resource. 

- `availability_mode` - (Optional, String) The availability mode of the VPN gateway:- `zonal`: The availability of this VPN gateway is limited only to a single zone of a  given region as provided by the `zone` of the VPN gateway.
  * Constraints: Allowable values are: `zonal`. 
- `local_asn` - (Optional, Integer) The local autonomous system number (ASN) for this VPN gateway and its connections.
- `members` - (Optional, List) The members for the VPN gateway.
  Nested schema for **members**:

	- `private_ip` - (Required, List) The reserved IP address assigned to the VPN gateway member.This property will be present only when the VPN gateway status is `available`.
	  Nested schema for **private_ip**:
		- `subnet` - (Required, List)
		  Nested schema for **subnet**: (one of the three, all three are mutually exclusive)
			- `crn` - (Required, String) The CRN for this subnet.
			- `href` - (Required, String) The URL for this subnet.
			- `id` - (Required, String) The unique identifier for this subnet.

- `mode`- (Optional, String) Mode in VPN gateway. Supported values are `route` or `policy`. The default value is `route`.
- `name` - (Required, String) The name of the VPN gateway.
- `resource_group` - (Optional, Forces new resource, String) The resource group (id), where the VPN gateway to be created.
- `subnet` - (Required, Forces new resource, String) The unique identifier for this subnet.
- `tags`- (Optional, Array of Strings) A list of tags that you want to add to your VPN gateway. Tags can help you find your VPN gateway more easily later.


## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `created_at` -  (String) The Second IP address assigned to this VPN gateway.
- `crn` - (String) The CRN for this VPN gateway.
- `id` - (String) The unique identifier of the VPN gateway.
- `members` - (List) Collection of VPN gateway members.

  Nested scheme for `members`:
  - `address` -  (String) The public IP address assigned to the VPN gateway member.
  - `private_address` -  (String) The private IP address assigned to the VPN gateway member.
  - `role` -  (String) The high availability role assigned to the VPN gateway member.
- `public_ip_address` - (String) The IP address assigned to this VPN gateway.
- `public_ip_address2` -  (String) The Second Public IP address assigned to this VPN gateway member.

  ~>**Note:** If one of the public IP addresses is "0.0.0.0", you can use a conditional expression to get the valid IP address: `ibm_is_vpn_gateway.example.public_ip_address == "0.0.0.0" ? ibm_is_vpn_gateway.example.public_ip_address2 : ibm_is_vpn_gateway.example.public_ip_address`

- `members` - (Optional, List) The members for the VPN gateway.
  Nested schema for **members**:
	- `health_reasons` - (Required, List) The reasons for the current `health_state` (if any).
	  Nested schema for **health_reasons**:
		- `code` - (Required, String) A reason code for this health state:- `cannot_reserve_ip_address`: IP address exhaustion (release addresses on the VPN's  subnet)- `internal_error`: Internal error (contact IBM support)The enumerated values for this property may[expand](https://cloud.ibm.com/apidocs/vpc#property-value-expansion) in the future.
		  * Constraints: Allowable values are: `cannot_reserve_ip_address`, `internal_error`. 
		- `message` - (Required, String) An explanation of the reason for this health state.
		- `more_info` - (Optional, String) A link to documentation about the reason for this health state.
	- `health_state` - (Required, String) The health of this resource:- `ok`: No abnormal behavior detected- `degraded`: Experiencing compromised performance, capacity, or connectivity- `faulted`: Completely unreachable, inoperative, or otherwise entirely incapacitated- `inapplicable`: The health state does not apply because of the current lifecycle   state. A resource with a lifecycle state of `failed` or `deleting` will have a   health state of `inapplicable`. A `pending` resource may also have this state.
	  * Constraints: Allowable values are: `degraded`, `faulted`, `inapplicable`, `ok`. 
	- `id` - (Optional, String) The unique identifier for this VPN gateway member.
	- `lifecycle_reasons` - (Required, List) The reasons for the current `lifecycle_state` (if any).
	  Nested schema for **lifecycle_reasons**:
		- `code` - (Required, String) A reason code for this lifecycle state:- `internal_error`: internal error (contact IBM support)- `resource_suspended_by_provider`: The resource has been suspended (contact IBM  support)The enumerated values for this property may[expand](https://cloud.ibm.com/apidocs/vpc#property-value-expansion) in the future.
		  * Constraints: Allowable values are: `internal_error`, `resource_suspended_by_provider`.
		- `message` - (Required, String) An explanation of the reason for this lifecycle state.
		- `more_info` - (Optional, String) A link to documentation about the reason for this lifecycle state.
	- `lifecycle_state` - (Required, String) The lifecycle state of the VPN gateway member.
	  * Constraints: Allowable values are: `deleting`, `failed`, `pending`, `stable`, `suspended`, `updating`, `waiting`.
	- `private_ip` - (Required, List) The reserved IP address assigned to the VPN gateway member.This property will be present only when the VPN gateway status is `available`.
	  Nested schema for **private_ip**:
		- `address` - (Required, String) The IP address.If the address has not yet been selected, the value will be `0.0.0.0`.This property may [expand](https://cloud.ibm.com/apidocs/vpc#property-value-expansion) to support IPv6 addresses in the future.
		- `deleted` - (Optional, List) If present, this property indicates the referenced resource has been deleted, and providessome supplementary information.
		  Nested schema for **deleted**:
			- `more_info` - (Required, String) A link to documentation about deleted resources.
		- `href` - (Required, String) The URL for this reserved IP.
		- `id` - (Required, String) The unique identifier for this reserved IP.
		- `name` - (Required, String) The name for this reserved IP. The name is unique across all reserved IPs in a subnet.
		- `resource_type` - (Required, String) The resource type.
		- `subnet` - (Required, List)
		  Nested schema for **subnet**:
			- `crn` - (Required, String) The CRN for this subnet.
			- `deleted` - (Optional, List) If present, this property indicates the referenced resource has been deleted, and  providessome supplementary information.
			  Nested schema for **deleted**:
				- `more_info` - (Computed, String) A link to documentation about deleted resources.
			- `href` - (Required, String) The URL for this subnet.
			- `id` - (Required, String) The unique identifier for this subnet.
			- `name` - (Computed, String) The name for this subnet. The name is unique across all subnets in the VPC.
			- `resource_type` - (Computed, String) The resource type.
	- `public_ip` - (Required, List) The public IP address assigned to the VPN gateway member.
	  Nested schema for **public_ip**:
		- `address` - (Required, String) The IP address.This property may [expand](https://cloud.ibm.com/apidocs/vpc#property-value-expansion) to support IPv6 addresses in the future.
	- `role` - (Required, String) The high availability role assigned to the VPN gateway member.The enumerated values for this property may[expand](https://cloud.ibm.com/apidocs/vpc#property-value-expansion) in the future.
	  * Constraints: Allowable values are: `active`, `standby`.
- `mode` - (Optional, String) The mode for this VPN gateway.
  * Constraints: Allowable values are: `policy`. The value must match regular expression `/^[a-z][a-z0-9]*(_[a-z0-9]+)*$/`.
- `name` - (Optional, String) The name for this VPN gateway. The name is unique across all VPN gateways in the VPC.
  * Constraints: The maximum length is `63` characters. The minimum length is `1` character. The value must match regular expression `/^-?([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$/`.
- `resource_group` - (Optional, List) The resource group for this VPN gateway.
  Nested schema for **resource_group**:
	- `href` - (Computed, String) The URL for this resource group.
	- `id` - (Required, String) The unique identifier for this resource group.
	- `name` - (Computed, String) The name for this resource group.
- `subnet` - (Optional, List) Identifies a subnet by a unique property.
  Nested schema for **subnet**:
	- `crn` - (Required, String) The CRN for this subnet.
	- `deleted` - (Optional, List) If present, this property indicates the referenced resource has been deleted, and providessome supplementary information.
	  Nested schema for **deleted**:
		- `more_info` - (Computed, String) A link to documentation about deleted resources.
	- `href` - (Required, String) The URL for this subnet.
	- `id` - (Required, String) The unique identifier for this subnet.
	- `name` - (Computed, String) The name for this subnet. The name is unique across all subnets in the VPC.
	- `resource_type` - (Computed, String) The resource type.

- `private_ip_address` -  (String) The Private IP address assigned to this VPN gateway member.
- `private_ip_address2` -  (String) The Second Private IP address assigned to this VPN gateway.
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
- `local_asn` - (Integer) The local autonomous system number (ASN) for this VPN gateway and its connections.
- `vpc` - (String) 	The VPC this VPN server resides in.

  Nested scheme for `vpc`:
  - `crn` - (String) The CRN for this VPC.
  - `deleted` - (List) 	If present, this property indicates the referenced resource has been deleted and provides some supplementary information.

    Nested scheme for **deleted**:
    - `more_info` - (String) Link to documentation about deleted resources.
  - `href` - (String) - The URL for this VPC
  - `id` - (String) - The unique identifier for this VPC.
  - `name` - (String) - The unique user-defined name for this VPC.
- `resource_type` - (String) - The resource type.



## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import the `ibm_is_vpn_gateway` resource by using `id`.
The `id` property can be formed from `VPN gateway ID`. For example:

```terraform
import {
  to = ibm_is_vpn_gateway.example
  id = "<vpn_gateway_ID>"
}
```

Using `terraform import`. For example:

```console
% terraform import ibm_is_vpn_gateway.example <vpn_gateway_ID>
```