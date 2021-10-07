---

subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : public_gateway"
description: |-
  Manages IBM public gateway.
---

# ibm_is_public_gateway
Create, update, or delete a public gateway for a VPC subnet. Public gateways enable a VPC subnet and all the instances that are connected to the subnet to connect to the internet. For more information, see [use a Public Gateway for external connectivity of a subnet](hhttps://cloud.ibm.com/docs/vpc?topic=vpc-about-networking-for-vpc#public-gateway-for-external-connectivity).

## Example usage
The following example shows how you can create a public gateway for all the subnets that are located in a specific zone.

```terraform
resource "ibm_is_vpc" "testacc_vpc" {
  name = "test"
}

resource "ibm_is_public_gateway" "testacc_gateway" {
  name = "test-gateway"
  vpc  = ibm_is_vpc.testacc_vpc.id
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

- `access_tags`  - (Optional, List of Strings) A list of access management tags to attach to the public gateway.
  **Note** 
  - Create access tag using `ibm_resource_tag` resource. You can attach only the access tags that already exists.
  - For more information, about creating access tags, see [working with tags](https://cloud.ibm.com/docs/account?topic=account-tag).
  - You must have the access listed in the [Granting users access to tag resources](https://cloud.ibm.com/docs/account?topic=account-access) for creating `access_tags`
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
