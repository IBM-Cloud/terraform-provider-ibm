---
subcategory: "Internet services"
layout: "ibm"
page_title: "IBM: ibm_cis_managed_lists"
description: |-
  Get information on an IBM Cloud Internet Services managed lists.
---

# ibm_cis_managed_lists

Retrieve information about IBM Cloud Internet Services managed lists data sources. For more information, see [IBM Cloud Internet Services].

## Example usage

```terraform
  data ibm_cis_managed_lists managed_lists {
    cis_id    = ibm_cis.instance.id
  }
```

## Argument reference

Review the argument references that you can specify for your data source.

- `cis_id` - (Required, String) The ID of the CIS service instance.

## Attributes reference

In addition to the argument reference list, you can access the following attribute references after your data source is created.

- `lists` - (List)
  - `description` - (string) Description of the managed list.
  - `kind` - (string) The kind of the managed list.
  - `name` - (string) Name of the managed list.
