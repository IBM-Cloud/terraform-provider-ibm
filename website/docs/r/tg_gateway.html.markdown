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
resource "ibm_tg_gateway" "new_tg_gw" {
  name           = "transit-gateway-1"
  location       = "us-south"
  global         = true
  resource_group = "30951d2dff914dafb26455a88c0c0092"
}

# Gateway in a redundancy group
resource "ibm_tg_gateway" "redundant_tg_gw" {
  name             = "transit-gateway-redundant"
  location         = "us-south"
  global           = true
  redundancy_group = "my-redundancy-group"
  resource_group   = "30951d2dff914dafb26455a88c0c0092"
}
```

## Argument reference
Review the argument references that you can specify for your resource.

- `location` - (Required, Forces new resource, String) The location of the transit gateway. For example, `us-south`.
- `name` - (Required, String) The unique user-defined name for the gateway. For example, `myGateway`.
- `global` - (Optional, Bool) The gateways with global routing (true) to connect to the networks outside their associated region. Cannot be changed if `redundancy_group` is set.
- `gre_enhanced_route_propagation` - (Optional, Bool) Allows route propagation across all GREs connected to the same transit gateway. This affects connections on the gateway of type `redundant_gre`, `unbound_gre_tunnel`, and `gre_tunnel`.
- `redundancy_group` - (Optional, String) The name of the redundancy group to add this global transit gateway to. If the redundancy group does not exist in the account it will be created. Cannot be specified together with `redundancy_group_id`.
- `resource_group` - (Optional, Forces new resource, String) The resource group ID where the transit gateway is to be created.

## Attribute reference
In addition to all argument reference list, you can access the following attribute references after your resource is created.

- `connection_count` - (Integer) The number of connections associated with this gateway.
- `connection_needs_attention` - (Bool) Indicates if this gateway has a connection that needs attention, such as a cross-account approval.
- `crn` - (String) The CRN of the gateway.
- `created_at` - (Timestamp) The date and time the gateway is created.
- `id` - (String) The unique identifier of the gateway.
- `redundancy_group_id` - (String) The unique identifier of the redundancy group this gateway belongs to.
- `status` - (String) The configuration status of the gateway, such as **available**, **pending**.
- `updated_at` - (Timestamp) The date and time the gateway is last updated.

## Import
The `ibm_tg_gateway` resource can be imported by using the transit gateway ID.

**Example**

```
$ terraform import ibm_tg_gateway.example 5ffda12064634723b079acdb018ef308
```
