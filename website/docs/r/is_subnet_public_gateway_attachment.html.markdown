---

subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : subnet public gateway attachment"
description: |-
  Manages IBM Subnet public gateway attachment.
---

# ibm_is_subnet_public_gateway_attachment
Create, update, or delete a public gateway attachment for a VPC subnet. Public gateways enable a VPC subnet and all the instances that are connected to the subnet to connect to the internet. For more information, see [use a Public Gateway for external connectivity of a subnet](https://cloud.ibm.com/docs/vpc?topic=vpc-about-networking-for-vpc#public-gateway-for-external-connectivity).

## Example usage

```terraform
	resource "ibm_is_vpc" "example" {
		name = "example-vpc"
	}

	resource "ibm_is_subnet" "example" {
		name 				      = "example-subnet"
		vpc 				      = ibm_is_vpc.example.id
		zone 				      = "us-south-1"
		ipv4_cidr_block 	= "10.240.0.0/24"
	}

	resource "ibm_is_public_gateway" "example" {
		name = "example-public-gateway"
		vpc  = ibm_is_vpc.example.id
		zone = "us-south-1"
	}

	resource "ibm_is_subnet_public_gateway_attachment" "example" {
		depends_on 		  = [ibm_is_public_gateway.example, ibm_is_subnet.example]
		subnet      	  = ibm_is_subnet.example.id
		public_gateway 	= ibm_is_public_gateway.example.id
	}

```
## Argument reference
Review the argument references that you can specify for your resource. 

- `public_gateway` - (Required, String) The public gateway identifier.
- `subnet` - (Required, Forces new resource, String) The subnet identifier.


## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `floating_ip` - (List) A list of floating IP addresses that you want to assign to the public gateway.
	- `id` - (String) The unique identifier of the floating IP address. If you specify this parameter, do not specify `address` at the same time. 
	- `address` - (String) The floating IP address. If you specify this parameter, do not specify `id` at the same time.
- `name` -  (String) Enter a name for your public gateway.
- `resource_group` - (String) Enter the ID of the resource group where you want to create the public gateway. To list available resource groups, run `ibmcloud resource groups`. If you do not specify a resource group, the public gateway is created in the `default` resource group.
- `vpc` - (String) Enter the ID of the VPC, for which you want to create a public gateway. To list available VPCs, run `ibmcloud is vpcs`.
- `zone` - (String) Enter the zone where you want to create the public gateway. 

## Import
The `ibm_is_subnet_public_gateway_attachment` resource can be imported by using the ID. 

**Syntax**

```
$ terraform import ibm_is_subnet_public_gateway_attachment.example <subnet__ID>
```

**Example**

```
$ terraform import ibm_is_subnet_public_gateway_attachment.example d7bec597-4726-451f-8a63-1111e6f19c32c
```
