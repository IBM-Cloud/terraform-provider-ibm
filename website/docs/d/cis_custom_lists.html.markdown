---
subcategory: "Internet services"
layout: "ibm"
page_title: "IBM: ibm_cis_custom_lists"
description: |-
  Get information on an IBM Cloud Internet Services custom lists.
---

# ibm_cis_custom_lists

Retrieve information about IBM Cloud Internet Services custom lists data sources. For more information, see [IBM Cloud Internet Services].

## Example usage

```terraform
  data ibm_cis_custom_lists custom_lists {
    cis_id    = ibm_cis.instance.id
    list_id   = ibm_cis.lists.list_id 
  }
```

## Argument reference

Review the argument references that you can specify for your data source.

- `cis_id` - (Required, String) The ID of the CIS service instance.
- `list_id` - (Optional, String) The ID of the custom list. If `list_id` is provided, details will be given for this particular list otherwise you will get the details of all the lists.

## Attributes reference

In addition to the argument reference list, you can access the following attribute references after your data source is created.

- `lists` - (List)
  - `list_id` - (string) The unique ID of the list.
  - `description` - (string) Description of the custom list.
  - `kind` - (string) The kind of the custom list.
  - `name` - (string) Name of the custom list.
  - `num_items` - (int) The number of items in the list.
  - `num_referencing_filters` - (int) The number of filters referencing the list.
