---
layout: "ibm"
page_title: "IBM : tg_gateway"
sidebar_current: "docs-ibm-resource-tg-gateway"
description: |-
  Manages IBM Transit Gateway.
---

# ibm\_dl_gateway

Provides a transit gateway resource. This allows transit gateway to be created, and updated and deleted.

## Example Usage

```hcl
resource "ibm_tg_gateway" "new_tg_gw"{
name="transit-gateway-1"
location="us-south"
global=true
resource_group=""
}  
```

## Argument Reference

The following arguments are supported:

* `name` - (Required, boolean) The unique user-defined name for this gateway. Example: myGateway
* `location` - (Required, Forces new resource, integer) Transit Gateway location. Example: us-south
* `global` - (Required, boolean) Gateways with global routing (true) can connect to networks outside their associated region.

## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier of this gateway. 
* `name` - The unique user-defined name for this gateway. 
* `crn` - The CRN (Cloud Resource Name) of this gateway.
* `location` -  The Transit Gateway location. Example: us-south
* `created_at` - The date and time resource was created.
* `updated_at` - The date and time resource was created.
* `status` - The status of the transit gateway. Example Available/Pending
* `resource_group` - Resource group reference


## Import

ibm_tg_gateway can be imported using transit gateway id, eg

```
$ terraform import ibm_tg_gateway.example 5ffda12064634723b079acdb018ef308
```