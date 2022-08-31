---

subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : vpc-route"
description: |-
  Manages IBM IS VPC route.
---

~>**Note**  This resource is deprecated, use `ibm_is_vpc_routing_table_route` instead.
# ibm_is_vpc_route
Create, update, or delete a VPC route. For more information, about VPC routes, see [setting up advanced routing in VPC](https://cloud.ibm.com/docs/vpc?topic=vpc-about-custom-routes).

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

resource "ibm_is_vpc_route" "example" {
  name        = "example-route"
  vpc         = ibm_is_vpc.example.id
  zone        = "us-south-1"
  destination = "192.168.4.0/24"
  next_hop    = "10.0.0.4"
}

```

## Timeouts
The `ibm_is_vpc_route` resource provides the following [Timeouts](https://www.terraform.io/docs/language/resources/syntax.html) configuration options:

- **create**: The creation of the route is considered `failed` when no response is received for 10 minutes. 
- **delete**: The deletion of the route is considered `failed` when no response is received for 10 minutes. 


## Argument reference
Review the argument references that you can specify for your resource. 

- `destination` - (Required, Forces new resource, String) The destination IP address or CIDR that network traffic from your VPC must match to be routed to the `next_hop`.
- `name` - (Required, String) The name of the route that you want to create.
- `next_hop` - (Required, String) The IP address where network traffic is sent next.
- `vpc` - (Required, Forces new resource, String) The ID of the VPC where you want to create the route. 
- `zone` - (Required, Forces new resource, String) The name of the VPC zone where you want to create the route.

## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `id` - (String) The ID of the VPC route. The ID is composed of `<vpc_id>/<vpc_route_id>`.
- `status` - (String) The status of the VPC route.

## Import
The `ibm_is_vpc_route` resource can be imported by using the VPC and route IDs. 

**Syntax**

```
$ terraform import ibm_is_vpc_route.example <vpc_ID>/<vpc_route_ID>
```
