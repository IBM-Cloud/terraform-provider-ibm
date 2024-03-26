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
  - `crn` - (String) The VPN gateway's CRN.
  - `created_at`- (Timestamp) The date and time the VPN gateway was created.
  - `id` - (String) The ID of the VPN gateway.
  - `name`-  (String) The VPN gateway instance name.
  - `members` - (List) Collection of VPN gateway members.</n>
  
      Nested scheme for `members`:
	    - `address` - (String) The public IP address assigned to the VPN gateway member.</br>
	    - `role`-  (String) The high availability role assigned to the VPN gateway member.</br>
      - `private_ip` - (List) The primary IP address to bind to the network interface. This can be specified using an existing reserved IP, or a prototype object for a new reserved IP.
        
          Nested scheme for `private_ip`:
          - `address` - (String) The IP address. If the address has not yet been selected, the value will be 0.0.0.0. This property may add support for IPv6 addresses in the future. When processing a value in this property, verify that the address is in an expected format. If it is not, log an error. Optionally halt processing and surface the error, or bypass the resource on which the unexpected IP address format was encountered.
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