---

subcategory: "Transit Gateway"
layout: "ibm"
page_title: "IBM : tg_redundancy_groups"
description: |-
  Lists IBM Cloud Transit Gateway redundancy groups.
---

# ibm_tg_redundancy_groups
Retrieve a list of all redundancy groups in the IBM Cloud account. For more information, about transit gateways, see [managing transit gateways](https://cloud.ibm.com/docs/transit-gateway?topic=transit-gateway-edit-gateway).

## Example usage

```terraform
# List all redundancy groups
data "ibm_tg_redundancy_groups" "all" {
}

# Filter by name
data "ibm_tg_redundancy_groups" "named" {
  name = "my-redundancy-group"
}
```

## Argument reference
Review the argument references that you can specify for your data source.

- `name` - (Optional, String) Filter the results to only the redundancy group with this name.

## Attribute reference
You can access the following attribute references after your data source is created.

- `redundancy_groups` - (List) List of redundancy groups.

  Nested scheme for `redundancy_groups`:
  - `id` - (String) The unique identifier of the redundancy group.
  - `name` - (String) The name of the redundancy group.
  - `created_at` - (String) The date and time the redundancy group was created.
  - `updated_at` - (String) The date and time the redundancy group was last updated.
