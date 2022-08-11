---

subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : public_gateway"
description: |-
  Manages IBM public gateway.
---

# ibm_is_public_gateway
Create, update, or delete a public gateway for a VPC subnet. Public gateways enable a VPC subnet and all the instances that are connected to the subnet to connect to the internet. For more information, see [use a Public Gateway for external connectivity of a subnet](https://cloud.ibm.com/docs/vpc?topic=vpc-about-networking-for-vpc#public-gateway-for-external-connectivity).

**Note:** 
VPC infrastructure services are a regional specific based endpoint, by default targets to `us-south`. Please make sure to target right region in the provider block as shown in the `provider.tf` file, if VPC service is created in region other than `us-south`.

**provider.tf**

```terraform
provider "ibm" {
  region = "eu-gb"
}
```

## Example usage
The following example shows how you can create a public gateway for all the subnets that are located in a specific zone.

```terraform
resource "ibm_is_vpc" "example" {
  name = "example-vpc"
}

resource "ibm_is_public_gateway" "example" {
  name = "example-gateway"
  vpc  = ibm_is_vpc.example.id
  zone = "us-south-1"

  //User can configure timeouts
  timeouts {
    create = "90m"
  }
}

```

## Timeouts
The `ibm_is_public_gateway` resource provides the following [Timeouts](https://www.terraform.io/docs/language/resources/syntax.html) configuration options:

- **create** The creation of the public gateway is considered `failed` when no response is received for 10 minutes. 
- **delete** The deletion of the public gateway is considered `failed` when no response is received for 10 minutes.

## Argument reference
Review the argument references that you can specify for your resource. 

- `floating_ip` - (Optional, List) A list of floating IP addresses that you want to assign to the public gateway.
	- `id` - (Optional, String) The unique identifier of the floating IP address. If you specify this parameter, do not specify `address` at the same time. 
	- `address` - (Optional, String) The floating IP address. If you specify this parameter, do not specify `id` at the same time.
- `name` -  (Required, String) Enter a name for your public gateway.
- `resource_group` - (Optional, Forces new resource, String) Enter the ID of the resource group where you want to create the public gateway. To list available resource groups, run `ibmcloud resource groups`. If you do not specify a resource group, the public gateway is created in the `default` resource group.
- `tags` (Optional, Array of Strings) Enter any tags that you want to associate with your VPC. Tags might help you find your VPC more easily after it is created. Separate multiple tags with a comma (`,`).
- `vpc` - (Required, Forces new resource, String) Enter the ID of the VPC, for which you want to create a public gateway. To list available VPCs, run `ibmcloud is vpcs`.
- `zone` - (Required, Forces new resource, String) Enter the zone where you want to create the public gateway. To list available zones, run `ibmcloud is zones`.

## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `crn` - (String) The crn for the public gateway.
- `id` - (String) The unique identifier that was assigned to your public gateway.
- `status` - (String) The provisioning status of your public gateway.

## Import
The `ibm_is_public_gateway` resource can be imported by using ID.

**Example**

```
$ terraform import ibm_is_public_gateway.example d7bec597-4726-451f-8a63-e62e6f19c32c
```
