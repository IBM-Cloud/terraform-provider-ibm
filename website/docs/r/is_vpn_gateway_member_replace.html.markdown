
---

subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM: is_vpn_gateway_member_replace"
description: |-
  Manages IBM VPC VPN gateway member replacement.
---

# ibm_is_vpn_gateway_member_replace

Moves a member of a route-based VPN gateway to another subnet, and so to another zone if needed. The API recreates the member in the given subnet with a new reserved IP. Use this, for example, to rebalance the members of a regional VPN gateway across zones. For more information about VPC VPN gateways, see [IBM Cloud Docs: Virtual Private Cloud - VPN Gateway](https://cloud.ibm.com/docs/vpc?topic=vpc-vpn-onprem-example).

**Note:** 
VPC infrastructure services are a regional specific based endpoint, by default targets to `us-south`. Please make sure to target right region in the provider block as shown in the `provider.tf` file, if VPC service is created in region other than `us-south`.

**provider.tf**

```terraform
provider "ibm" {
  region = "eu-gb"
}
```

## Timeouts

The `ibm_is_vpn_gateway_member_replace` provides the following [Timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) configuration options:

- **create** - (Default 10 minutes) Used for replacing the VPN gateway member and waiting until its `lifecycle_state` is `stable`.
- **delete** - (Default 10 minutes) Not used. Delete only removes the resource from the Terraform state.

## Example usage

```terraform
resource "ibm_is_vpc" "example" {
  name = "example-vpc"
}

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

resource "ibm_is_subnet" "new_subnet" {
  name                     = "example-subnet-new"
  vpc                      = ibm_is_vpc.example.id
  zone                     = "us-south-3"
  total_ipv4_address_count = 16
}

resource "ibm_is_vpn_gateway" "example" {
  name              = "example-vpn-gateway"
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

# Replace the VPN gateway member's subnet
resource "ibm_is_vpn_gateway_member_replace" "example" {
  vpn_gateway_id        = ibm_is_vpn_gateway.example.id
  vpn_gateway_member_id = ibm_is_vpn_gateway.example.members[0].id
  subnet {
    id = ibm_is_subnet.new_subnet.id
  }
}
```

## Example usage with subnet CRN

```terraform
resource "ibm_is_vpn_gateway_member_replace" "example" {
  vpn_gateway_id        = ibm_is_vpn_gateway.example.id
  vpn_gateway_member_id = ibm_is_vpn_gateway.example.members[0].id
  subnet {
    crn = ibm_is_subnet.new_subnet.crn
  }
}
```

## Example usage with subnet href

```terraform
resource "ibm_is_vpn_gateway_member_replace" "example" {
  vpn_gateway_id        = ibm_is_vpn_gateway.example.id
  vpn_gateway_member_id = ibm_is_vpn_gateway.example.members[0].id
  subnet {
    href = ibm_is_subnet.new_subnet.href
  }
}
```

## Argument reference

Review the argument references that you can specify for your resource. 

- `vpn_gateway_id` - (Required, Forces new resource, String) The unique identifier of the VPN gateway.
- `vpn_gateway_member_id` - (Required, Forces new resource, String) The unique identifier of the VPN gateway member to be replaced.
- `subnet` - (Required, Forces new resource, List) The subnet to move the member to. A reserved IP is allocated from this subnet. Specify exactly one of `id`, `crn` or `href`; the other two are filled in after create.
  - `id` - (Optional, String) The unique identifier of the subnet.
  - `crn` - (Optional, String) The CRN of the subnet.
  - `href` - (Optional, String) The URL of the subnet.

## Attribute reference

In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `id` - (String) The ID of this resource, in the format `vpn_gateway_id/vpn_gateway_member_id`. The member ID part is the ID returned by the replace call.
- `health_state` - (String) The health state of the VPN gateway member.
- `lifecycle_state` - (String) The lifecycle state of the VPN gateway member. It is `stable` once the replace completes.
- `private_ip_address` - (String) The private IP address of the member in the new subnet.
- `public_ip_address` - (String) The public IP address of the member.
- `role` - (String) The high availability role of the member, for example `active` or `standby`.

## Notes

- This resource calls the replace VPN gateway member API once, at create time.
- `delete` only removes the resource from the Terraform state. It does not move the member back.
- All arguments force a new resource. Changing `subnet` moves the member again.
- The move can disrupt traffic on that member. Move one member at a time (use `depends_on` between two of these resources) so the gateway keeps one working member.
- When the gateway was created with `members`, moving a member does not cause a difference on `ibm_is_vpn_gateway`.

## Import

You can import this resource by using the ID in the format `<vpn_gateway_id>/<vpn_gateway_member_id>`.

```
$ terraform import ibm_is_vpn_gateway_member_replace.example r006-1234abcd/r006-5678efgh
```
