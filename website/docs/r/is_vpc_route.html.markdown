---

subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : vpc-route"
description: |-
  Manages IBM IS VPC Route.
---

# ibm\_is_vpc_route

Provides a vpc route resource. This allows vpc route to be created, updated, and cancelled.


## Example Usage

```hcl
resource "ibm_is_vpc" "testacc_vpc" {
  name = "testvpc"
}

resource "ibm_is_vpc_route" "testacc_vpc_route" {
  name        = "routetest"
  vpc         = ibm_is_vpc.testacc_vpc.id
  zone        = "us-south-1"
  destination = "192.168.4.0/24"
  next_hop    = "10.0.0.4"
}

```

## Timeouts

ibm_is_vpc_route provides the following [Timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) configuration options:

* `create` - (Default 10 minutes) Used for Creating VPC Route.
* `delete` - (Default 10 minutes) Used for Deleting VPC Route.


## Argument Reference

The following arguments are supported:

* `name` - (Required, string) The route name.
* `vpc` - (Required, Forces new resource, string) The vpc id. 
* `zone` - (Required, Forces new resource, string) Name of the zone. 
* `destination` - (Required, Forces new resource, string) The destination of the route. 
* `next_hop` - (Required, string) The next hop of the route. 

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The id of the route. The id is composed of \<vpc_id\>/\<vpc_route_id\>
* `status` - The status of the VPC Route.

## Import

ibm_is_vpc_route can be imported using VPC ID and VPC Route ID, eg

```
$ terraform import ibm_is_vpc_route.example 56738c92-4631-4eb5-8938-8af9211a6ea4/fc2667e0-9e6f-4993-a0fd-cabab477c4d1
```