---

subcategory: "Transit Gateway"
layout: "ibm"
page_title: "IBM : tg_gateway"
description: |-
  Manages an IBM Transit Gateway.
---

# ibm_tg_gateway
Create, update and delete for the transit gateway resource. For more information, about transit location, see [managing transit gateways](https://cloud.ibm.com/docs/transit-gateway?topic=transit-gateway-edit-gateway).

## Example usage

```terraform
resource "ibm_tg_gateway" "new_tg_gw"{
name="transit-gateway-1"
location="us-south"
global=true
resource_group="30951d2dff914dafb26455a88c0c0092"
}  
```

## Argument reference
Review the argument references that you can specify for your resource. 

- `location` - (Optional, Forces new resource, Integer) The location of the transit gateway. For example, `us-south`.
- `name` - (Required, String) The unique user-defined name for the gateway. For example, `myGateway`.
- `global` - (Required, Bool) The gateways with global routing (true) to connect to the networks outside their associated region.
- `gre_enhanced_route_propagation` - (Optional, Bool) Allows route propagation across all GREs connected to the same transit gateway. This affects connections on the gateway of type redundant_gre, unbound_gre_tunnel and gre_tunnel.
- `resource_group` -  (Optional, Forces new resource, String) The resource group ID where the transit gateway to be created.

## Attribute reference
In addition to all argument reference list, you can access the following attribute references after your resource is created.

- `crn` - (String) The CRN of the gateway.
- `created_at` - (Timestamp) The date and time the connection is created. 
- `id` - (String) The unique identifier of the gateway ID or connection ID resource.
- `status` - (String) The configuration status of the connection, such as **Available**, **pending**.
- `updated_at` - (Timestamp) The date and time the connection is last updated.

## Import
The `ibm_tg_gateway` resource can be imported by using transit gateway ID and connection ID.

**Example**

```
$ terraform import ibm_tg_gateway.example 5ffda12064634723b079acdb018ef308
```
