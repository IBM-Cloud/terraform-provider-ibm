---

subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : subnet_public_gateway_attachment"
description: |-
  Manages IBM Subnet public gateway attachment.
---

# ibm_is_subnet_public_gateway_attachment
Create, update, or delete a public gateway attachment for a VPC subnet. Public gateways enable a VPC subnet and all the instances that are connected to the subnet to connect to the internet. For more information, see [use a Public Gateway for external connectivity of a subnet](https://cloud.ibm.com/docs/vpc?topic=vpc-about-networking-for-vpc#public-gateway-for-external-connectivity).

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
resource "ibm_is_vpc" "example" {
  name = "example-vpc"
}

resource "ibm_is_subnet" "example" {
  name 				              = "example-subnet"
  vpc 				              = ibm_is_vpc.example.id
  zone 				              = "eu-gb-1"
  total_ipv4_address_count  = 16
}

resource "ibm_is_public_gateway" "example" {
  name = "example-public-gateway"
  vpc  = ibm_is_vpc.example.id
  zone = "eu-gb-1"
}

resource "ibm_is_subnet_public_gateway_attachment" "example" {
  subnet      	 	  = ibm_is_subnet.example.id
  public_gateway 		= ibm_is_public_gateway.example.id
}

```

## Argument reference
Review the argument references that you can specify for your resource. 

- `public_gateway` - (Required, String) The public gateway identifier.
- `subnet` - (Required, Forces new resource, String) The subnet identifier.


## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `crn` - (String) The CRN for this public gateway.
- `floating_ip` - (List) The floating IP bound to this public gateway.
  Nested scheme for `floating_ip`:
	- `address` - (String) The globally unique IP address for this floating ip.
	- `id` - (String) The unique identifier of the floating IP address.
- `id` - (String) The unique identifier of the subnet.
- `name` -  (String) The user-defined name for this public gateway.
- `resource_group` - (String) The resource group identifier for this public gateway.
- `resource_group_name` - (String) The name for the resource group for this public gateway.
- `resource_type` - (String) The resource type for this public gateway.
- `status` - (String) The status of this public gateway.
- `vpc` - (String) The identifier of the VPC this public gateway serves.
- `zone` - (String) The zone this public gateway resides in.

## Import
The `ibm_is_subnet_public_gateway_attachment` resource can be imported by using the subnet ID. 

**Syntax**

```
$ terraform import ibm_is_subnet_public_gateway_attachment.example <subnet_ID>
```

**Example**

```
$ terraform import ibm_is_subnet_public_gateway_attachment.example d7bec597-4726-451f-8a63-1111e6f19c32c
```
