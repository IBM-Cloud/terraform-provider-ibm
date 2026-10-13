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

The following example creates a regional, route-based VPN gateway with one member in each of two zones:

```terraform
resource "ibm_is_subnet" "zone1" {
  name                     = "example-subnet-zone1"
  vpc                      = ibm_is_vpc.example.id
  zone                     = "us-south-1"
  total_ipv4_address_count = 16
}

resource "ibm_is_subnet" "zone2" {
  name                     = "example-subnet-zone2"
  vpc                      = ibm_is_vpc.example.id
  zone                     = "us-south-2"
  total_ipv4_address_count = 16
}

resource "ibm_is_vpn_gateway" "regional" {
  name              = "example-regional-vpn-gateway"
  availability_mode = "regional"
  mode              = "route"
  members {
    private_ip {
      subnet {
        id = ibm_is_subnet.zone1.id
      }
    }
  }
  members {
    private_ip {
      subnet {
        id = ibm_is_subnet.zone2.id
      }
    }
  }
}
```

## Timeouts
The `ibm_is_vpn_gateway` resource provides the following [Timeouts](https://www.terraform.io/docs/language/resources/syntax.html) configuration options:

- **create**: The creation of the VPN gateway is considered `failed` when no response is received for 10 minutes. 
- **delete**: The deletion of the VPN gateway is considered `failed` when no response is received for 10 minutes. 


## Argument reference
Review the argument references that you can specify for your resource. 

- `availability_mode` - (Optional, Forces new resource, String) The availability mode of the VPN gateway. Allowable values are `zonal` and `regional`. If not set, the gateway is `zonal`.
  - `zonal`: the gateway lives in the single zone of `subnet`.
  - `regional`: the gateway has two `members` that can be in the same zone or in different zones of the region. Supported only when `mode` is `route`.

  ~>**Note:** `availability_mode` cannot be updated in place. Changing it destroys the gateway and creates a new one.
- `local_asn` - (Optional, Integer) The local autonomous system number (ASN) for this VPN gateway and its connections.
- `members` - (Optional, Forces new resource, List) The members of a regional VPN gateway. Required when `availability_mode` is `regional`, with exactly 2 items. Must not be set for a zonal gateway.

  Nested scheme for `members`:
  - `private_ip` - (Required, List) The reserved IP for the member. Exactly one item.

    Nested scheme for `private_ip`:
    - `subnet` - (Required, List) The subnet to create the member in. Exactly one item. Set exactly one of the following:
      - `crn` - (Optional, String) The CRN of the subnet.
      - `href` - (Optional, String) The URL of the subnet.
      - `id` - (Optional, String) The unique identifier of the subnet.

  ~>**Note:** The member subnet is used only when the gateway is created. To move a member to another subnet or zone later, use [ibm_is_vpn_gateway_member_replace](is_vpn_gateway_member_replace.html). After such a move, Terraform does not report a difference for `members`, and does not recreate the gateway.
- `mode`- (Optional, String) Mode in VPN gateway. Supported values are `route` or `policy`. The default value is `route`.
- `name` - (Required, String) The name of the VPN gateway.
- `resource_group` - (Optional, Forces new resource, String) The resource group (id), where the VPN gateway to be created.
- `subnet` - (Optional, Forces new resource, String) The unique identifier of the subnet for a zonal VPN gateway. Required when `availability_mode` is `zonal` or not set. Must not be set when `availability_mode` is `regional`.
- `tags`- (Optional, Array of Strings) A list of tags that you want to add to your VPN gateway. Tags can help you find your VPN gateway more easily later.


## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `created_at` -  (String) The Second IP address assigned to this VPN gateway.
- `crn` - (String) The CRN for this VPN gateway.
- `id` - (String) The unique identifier of the VPN gateway.
- `availability_mode` - (String) The availability mode of the VPN gateway: `zonal` or `regional`.
- `members` - (List) The members of the VPN gateway.

  Nested scheme for `members`:
  - `address` - (String) The public IP address assigned to the VPN gateway member. Same as `public_ip.0.address`.
  - `health_reasons` - (List) The reasons for the current `health_state` (if any).

    Nested scheme for `health_reasons`:
    - `code` - (String) A reason code for this health state, for example `cannot_reserve_ip_address` or `internal_error`.
    - `message` - (String) An explanation of the reason for this health state.
    - `more_info` - (String) A link to documentation about the reason for this health state.
  - `health_state` - (String) The health of the member: `ok`, `degraded`, `faulted` or `inapplicable`.
  - `id` - (String) The unique identifier for this VPN gateway member. Use it with `ibm_is_vpn_gateway_member_replace` or the `ibm_is_vpn_gateway_member` data source.
  - `lifecycle_reasons` - (List) The reasons for the current `lifecycle_state` (if any).

    Nested scheme for `lifecycle_reasons`:
    - `code` - (String) A reason code for this lifecycle state, for example `internal_error` or `resource_suspended_by_provider`.
    - `message` - (String) An explanation of the reason for this lifecycle state.
    - `more_info` - (String) A link to documentation about the reason for this lifecycle state.
  - `lifecycle_state` - (String) The lifecycle state of the member: `deleting`, `failed`, `pending`, `stable`, `suspended`, `updating` or `waiting`.
  - `private_address` - (String) The private IP address assigned to the VPN gateway member. Same as `private_ip.0.address`.
  - `private_ip` - (List) The reserved IP assigned to the VPN gateway member. Present only when the VPN gateway status is `available`.

    Nested scheme for `private_ip`:
    - `address` - (String) The IP address. The value is `0.0.0.0` if the address has not been selected yet.
    - `deleted` - (List) If present, the reserved IP has been deleted. Contains `more_info`.
    - `href` - (String) The URL for this reserved IP.
    - `id` - (String) The unique identifier for this reserved IP.
    - `name` - (String) The name for this reserved IP.
    - `resource_type` - (String) The resource type.
    - `subnet` - (List) The subnet of the member, with `crn`, `deleted`, `href`, `id`, `name` and `resource_type`.
  - `public_ip` - (List) The public IP assigned to the VPN gateway member, with `address`.
  - `role` - (String) The high availability role assigned to the VPN gateway member, for example `active` or `standby`.
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