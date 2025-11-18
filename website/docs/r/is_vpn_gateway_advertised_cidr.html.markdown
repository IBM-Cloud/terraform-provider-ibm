---

subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : VPN-gateway-advertised-cidr"
description: |-
  Manages IBM VPN gateway connection.
---

# ibm_is_vpn_gateway_advertised_cidr
Update, or delete a VPN gateway advertised cidr. For more information, about VPN gateway, see [adding connections to a VPN gateway](https://cloud.ibm.com/docs/vpc?topic=vpc-vpn-adding-connections).

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

resource "ibm_is_subnet" "example2" {
  name            = "example-subnet2"
  vpc             = ibm_is_vpc.example.id
  zone            = "us-south-2"
  ipv4_cidr_block = "10.240.68.0/24"
}

resource "ibm_is_vpn_gateway" "example" {
  name   = "example-vpn-gateway"
  subnet = ibm_is_subnet.example.id
  mode   = "route"
}

resource "ibm_is_vpn_gateway_advertised_cidr" "example" {
  vpn_gateway = ibm_is_vpn_gateway.example.id
  cidr        = "10.45.0.0/24"
}

```

## Timeouts
The `ibm_is_vpn_gateway_advertised_cidr` resource provides the following [Timeouts](https://www.terraform.io/docs/language/resources/syntax.html) configuration options:

- **delete** - (Default 10 minutes) Used for deleting instance.


## Argument reference
Review the argument references that you can specify for your resource. 

- `cidr` - (Required, Force new resource, String) The IP address range in CIDR block notation.
- `vpn_gateway` - (Required, Force new resource, String) The unique identifier of the VPN gateway.

## Import
The `ibm_is_vpn_gateway_advertised_cidr` resource can be imported by using the VPN gateway ID and the Advertised Cidr. 

**Syntax**

```
$ terraform import ibm_is_vpn_gateway_advertised_cidr.example <vpn_gateway_ID>/<cidr>
```

**Example**

```
$ terraform import ibm_is_vpn_gateway_advertised_cidr.example d7bec597-4726-451f-8a63-e62e6f19c32c/10.45.0.0/24
```
