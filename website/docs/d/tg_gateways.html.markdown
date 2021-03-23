---

subcategory: "Transit Gateway"
layout: "ibm"
page_title: "IBM : tg_gateways"
description: |-
  Manages IBM Cloud Infrastructure Transit Gateway.
---

# ibm\_tg_gateways

Import the details of an existing IBM Cloud Infrastructure transit gateways as a read-only data source. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.


## Example Usage

```hcl
data "ibm_tg_gateways" "ds_tggateways" {
}
```

## Attribute Reference

The following attributes are exported:

* `transit_gateways` - List of all Transit gateways in the IBM Cloud Infrastructure.
  * `created_at` - The date and time resource was created.
  * `updated_at` - The date and time resource was last updated.
  * `crn` - The CRN (Cloud Resource Name) of this gateway.
  * `global` - Gateways with global routing (true) can connect to networks outside their associated region.
  * `id` - The unique identifier of this gateway.
  * `location` - Gateway location.
  * `name` - The unique user-defined name for this gateway.
  * `status` - Gateway status.
  * `resource_group` - Resource group identifier.