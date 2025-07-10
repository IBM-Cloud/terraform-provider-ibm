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
data "ibm_tg_gateways" "ds_tggateways" {
}
```


## Argument reference
There is no argument reference for `ibm_tg_gateways`.

## Attribute reference
You can access the following attribute references after your data source is created. 

- `transit_gateways` - (String) List of all transit gateways.

  Nested scheme for `transit_gateways`:
   - `created_at` - (String) The date and time resource is created.
   - `crn` - (String) The CRN of the gateway.
   - `global` - (String) The gateways with global routing true to connect to the networks outside the associated region.
   - `id` - (String) The unique identifier of this gateway.
   - `location` - (String) The gateway location.
   - `name` - (String) The user defined name for the transit gateway connection.
   - `resource_group` - (String) The resource group identifier.
   - `gre_enhanced_route_propagation` - (Bool) The gateways with GRE enhanced route propagation true to share routes across all GRE connections on the same gateway.
   - `status` - (String) The gateway status.
   - `updated_at` - (String) The date and time resource is last updated.
