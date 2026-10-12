---

subcategory: "Transit Gateway"
layout: "ibm"
page_title: "IBM : tg_redundancy_group"
description: |-
  Retrieves an IBM Cloud Transit Gateway redundancy group.
---

# ibm_tg_redundancy_group
Retrieve information about a specific IBM Cloud Transit Gateway redundancy group by name. For more information, about transit gateways, see [managing transit gateways](https://cloud.ibm.com/docs/transit-gateway?topic=transit-gateway-edit-gateway).

## Example usage

```terraform
data "ibm_tg_redundancy_group" "rg" {
  name = "my-redundancy-group"
}

output "redundancy_group_id" {
  value = data.ibm_tg_redundancy_group.rg.id
}
```

## Argument reference
Review the argument references that you can specify for your data source.

- `name` - (Required, String) The name of the redundancy group.

## Attribute reference
In addition to the argument reference list, you can access the following attribute references after your data source is created.

- `id` - (String) The unique identifier of the redundancy group.
- `created_at` - (String) The date and time the redundancy group was created.
- `updated_at` - (String) The date and time the redundancy group was last updated.
