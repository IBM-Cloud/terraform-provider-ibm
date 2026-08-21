
---

subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM: is_vpn_gateway_member_replace"
description: |-
  Manages IBM VPC VPN gateway member replacement.
---

# ibm_is_vpn_gateway_member_replace

Provide support to replace a VPN gateway member's subnet. This resource replaces the subnet associated with a VPN gateway member by updating its reserved IP address. After successful replacement, the VPN gateway member will use a reserved IP from the new subnet. For more information about VPC VPN gateways, see [IBM Cloud Docs: Virtual Private Cloud - VPN Gateway](https://cloud.ibm.com/docs/vpc?topic=vpc-vpn-onprem-example).

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

- **create** - (Default 10 minutes) Used for replacing the VPN gateway member.
- **delete** - (Default 10 minutes) Used for deleting the resource reference.

## Example usage

```terraform
resource "ibm_is_vpc" "example" {
  name = "example-vpc"
}

resource "ibm_is_subnet" "old_subnet" {
  name                     = "old-subnet"
  vpc                      = ibm_is_vpc.example.id
  zone                     = "us-south-1"
  ipv4_cidr_block          = "10.240.0.0/24"
}

resource "ibm_is_subnet" "new_subnet" {
  name                     = "new-subnet"
  vpc                      = ibm_is_vpc.example.id
  zone                     = "us-south-1"
  ipv4_cidr_block          = "10.240.1.0/24"
}

resource "ibm_is_vpn_gateway" "example" {
  name   = "example-vpn-gateway"
  subnet = ibm_is_subnet.old_subnet.id
  mode   = "route"
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
- `subnet` - (Required, Forces new resource, List) The subnet that the VPN gateway member will use for its reserved IP. Specify exactly one of `id`, `crn`, or `href` within this block.
  - `id` - (Optional, String) The unique identifier of the subnet.
  - `crn` - (Optional, String) The CRN of the subnet.
  - `href` - (Optional, String) The href of the subnet.

## Attribute reference

In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `id` - (String) The unique identifier of the VPN gateway member replacement resource, formatted as `vpn_gateway_id/vpn_gateway_member_id`.
- `status` - (String) The status of the VPN gateway member replacement operation. Returns `204` on successful replacement.


## Notes

- This resource performs a **replace** operation on the VPN gateway member, updating its associated subnet.
- The operation succeeds when a `204` status code is returned from the API.
- The resource only supports `create` and `delete` operations. The `delete` operation simply removes the resource from the Terraform state and does not modify the actual VPN gateway member.
- All arguments are `ForceNew`, meaning any change will cause the resource to be recreated.
- The VPN gateway member's reserved IP address will be automatically updated to one from the new subnet.
- This operation may affect VPN connectivity. Ensure proper planning before performing the replacement.
