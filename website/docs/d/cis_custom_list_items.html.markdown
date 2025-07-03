---
subcategory: "Internet services"
layout: "ibm"
page_title: "IBM: ibm_cis_custom_list_items"
description: |-
  Get information on an IBM Cloud Internet Services custom list items.
---

# ibm_cis_custom_list_items

Retrieve information about IBM Cloud Internet Services custom list items data sources. For more information, see [IBM Cloud Internet Services].

## Example usage

```terraform
 data ibm_cis_custom_list_items custom_list_items {
    cis_id    = ibm_cis.instance.id
    list_id   = ibm_cis.lists.list_id 
    item_id   = ibm_cis.lists.item.item_id
}
```

## Argument reference

Review the argument references that you can specify for your data source.

- `cis_id` - (Required, String) The ID of the CIS service instance.
- `list_id` - (Required, String) The ID of the custom list.
- `item_id` - (Optional, String) The ID of the item. If `item_id` is provided, you will get the information for the partiular item otherwise information for all the items will be provided.

## Attributes reference

In addition to the argument reference list, you can access the following attribute references after your data source is created.

- `items` - (List)
  - `item_id` - (string) The unique ID of the item.
  - `asn` - (int) Defines a non-negative 32 bit integer.
  - `ip` - (string) An IPv4 address, an IPv4 CIDR, or an IPv6 CIDR. IPv6 CIDRs are limited to a maximum of /64.
  - `hostname` - (string) Defines the hostname.
  - `comment` - (string) Defines an informative summary of the list item.
  - `created_on` - (string) The timestamp of when the item was created.
  - `modified_on` - (string) The timestamp of when the item was last modified.
