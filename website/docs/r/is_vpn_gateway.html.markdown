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

- `local_asn` - (Optional, Integer) The local autonomous system number (ASN) for this VPN gateway and its connections.
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
The `ibm_is_vpn_gateway` resource can be imported by using the VPN gateway ID. 

**Syntax**

```
$ terraform import ibm_is_vpn_gateway.example <vpn_gateway_ID>
```

**Example**

```
$ terraform import ibm_is_vpn_gateway.example d7bec597-4726-451f-8a63-e621111119c32c
```
