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

# Gateway joined to a redundancy group by name (creates group if absent)
# redundancy_group requires global = true
resource "ibm_tg_gateway" "redundant_tg_gw" {
  name             = "transit-gateway-redundant"
  location         = "us-south"
  global           = true
  redundancy_group = "my-redundancy-group"
  resource_group   = "30951d2dff914dafb26455a88c0c0092"
}

# Gateway joined to an existing redundancy group by ID
resource "ibm_tg_gateway" "redundant_tg_gw_by_id" {
  name                = "transit-gateway-redundant-2"
  location            = "eu-de"
  global              = true
  redundancy_group_id = "a7fb9f77-0a3c-4d1e-abc1-1234567890ab"
  resource_group      = "30951d2dff914dafb26455a88c0c0092"
}
```

## Argument reference
Review the argument references that you can specify for your resource.

- `location` - (Required, Forces new resource, String) The location of the transit gateway. For example, `us-south`.
- `name` - (Required, String) The unique user-defined name for the gateway. For example, `myGateway`.
- `global` - (Optional, Bool) Allow global routing for a Transit Gateway so it can connect to networks outside its associated region. Default value is `false`. **Note:** Attempting to change this attribute on a gateway that belongs to a redundancy group will result in an error; the `global` flag cannot be updated while `redundancy_group` is set.
- `gre_enhanced_route_propagation` - (Optional, Computed, Bool) Allows route propagation across all GREs connected to the same transit gateway. This affects connections on the gateway of type `redundant_gre`, `unbound_gre_tunnel`, and `gre_tunnel`.
- `redundancy_group` - (Optional, Computed, String) The name of the redundancy group to add this global transit gateway to. If the redundancy group does not exist in the account it will be created. Requires `global = true`. Once set, `global` cannot be changed for the lifetime of the resource. Conflicts with `redundancy_group_id`.
- `redundancy_group_id` - (Optional, Computed, String) The unique identifier of an existing redundancy group to add this global transit gateway to. Requires `global = true`. Conflicts with `redundancy_group`.
- `resource_group` - (Optional, Computed, Forces new resource, String) The resource group ID where the transit gateway is to be created.
- `tags` - (Optional, Array of Strings) Tags associated with the transit gateway instance.

## Attribute reference
In addition to all argument reference list, you can access the following attribute references after your resource is created.

- `connection_count` - (Integer) The number of connections associated with this gateway.
- `connection_needs_attention` - (Bool) Indicates if this gateway has a connection that needs attention, such as a cross-account approval.
- `crn` - (String) The CRN of the gateway.
- `created_at` - (Timestamp) The date and time the gateway is created.
- `id` - (String) The unique identifier of the gateway.
- `redundancy_group_id` - (String) The unique identifier of the redundancy group this gateway belongs to.
- `status` - (String) The configuration status of the gateway, such as **available**, **pending**, **failed**.
- `updated_at` - (Timestamp) The date and time the gateway is last updated.

## Timeouts
The following timeouts are available for `ibm_tg_gateway`:

- **Create**: Default 10 minutes
- **Delete**: Default 32 minutes
- **Update**: Default 10 minutes

## Import
The `ibm_tg_gateway` resource can be imported by using the transit gateway ID.

**Example**

```
$ terraform import ibm_tg_gateway.example 5ffda12064634723b079acdb018ef308
```
