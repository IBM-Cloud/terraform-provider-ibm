---

subcategory: "Transit Gateway"
layout: "ibm"
page_title: "IBM : tg_gateways"
description: |-
  Manages IBM Cloud Infrastructure Transit Gateways.
---

# ibm_tg_gateways
Imports the information of an existing IBM Cloud infrastructure transit gateway as a read only data source. For more information, about transit gateways, see [managing transit gateways](https://cloud.ibm.com/docs/transit-gateway?topic=transit-gateway-edit-gateway).

## Example usage

```terraform
# List all gateways
data "ibm_tg_gateways" "all" {
}

# List only gateways in a specific redundancy group
data "ibm_tg_gateways" "in_rg" {
  redundancy_group = "my-redundancy-group"
}
```

## Argument reference
Review the argument references that you can specify for your data source.

- `redundancy_group` - (Optional, String) Filter the results to only gateways belonging to this redundancy group name.

## Attribute reference
You can access the following attribute references after your data source is created.

- `transit_gateways` - (List) List of all transit gateways.

  Nested scheme for `transit_gateways`:
   - `connection_count` - (Integer) The number of connections associated with this gateway.
   - `connection_needs_attention` - (Bool) Indicates if this gateway has a connection that needs attention, such as a cross-account approval.
   - `created_at` - (String) The date and time resource is created.
   - `crn` - (String) The CRN of the gateway.
   - `global` - (Bool) The gateways with global routing true to connect to the networks outside the associated region.
   - `gre_enhanced_route_propagation` - (Bool) The gateways with GRE enhanced route propagation true to share routes across all GRE connections on the same gateway.
   - `id` - (String) The unique identifier of this gateway.
   - `location` - (String) The gateway location.
   - `name` - (String) The user-defined name for the transit gateway.
   - `redundancy_group` - (String) The name of the redundancy group this gateway belongs to.
   - `redundancy_group_id` - (String) The unique identifier of the redundancy group this gateway belongs to.
   - `resource_group` - (String) The resource group identifier.
   - `status` - (String) The gateway status.
   - `updated_at` - (String) The date and time resource is last updated.
