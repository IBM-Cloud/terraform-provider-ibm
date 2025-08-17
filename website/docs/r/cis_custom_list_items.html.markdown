---
subcategory: "Internet services"
layout: "ibm"
page_title: "IBM: ibm_cis_custom_list_items"
description: |-
  Provides an IBM CIS custom list items resource.
---

# ibm_cis_custom_list

Provides an IBM Cloud Internet Services custom list items resource to create, update, and delete the custom list items of an instance. For more information, see [Using custom lists](https:/cloud.ibm.com/docs/cis?group=custom-lists).

## Example usage

```terraform
  # adding items in ip list
  resource ibm_cis_custom_list_items items {
    cis_id    = ibm_cis.instance.id
    list_id   = ibm_cis.custom_list.list_id 
    items {
        ip = var.ip1
    }
    items {
        ip = var.ip2
    }
  }

  # adding items in asn list
  resource ibm_cis_custom_list_items items {
    cis_id    = ibm_cis.instance.id
    list_id   = ibm_cis.custom_list.list_id 
    items {
        asn = 23
    }
    items {
        asn = 213
    }
  }

```

## Argument reference

Review the argument references that you can specify for your resource.

- `cis_id` - (Required, String) The ID of the CIS service instance.
- `list_id` - (Required, String) ID of the custom list.
- `items` - (Required, List) Items list
  - `asn` - (Optional, int) Defines a non-negative 32-bit integer. It is used with the list where kind is `asn`.
  - `ip` - (Optional, string) An IPv4 address, an IPv4 CIDR, or an IPv6 CIDR. IPv6 CIDRs are limited to a maximum of /64. It is used with the list where kind is `ip`.
  - `hostname` - (optional, string) Defines the hostname. It is used with the list where kind is `hostname`.
  - `comment` - (optional, string) Defines an informative summary of the list item.

## Attribute reference

In addition to the argument reference list, you can access the following attribute reference after your resource is created.

- `created_on` - (string) The timestamp of when the item was created.
- `modified_on` - (string) The timestamp of when the item was last modified.

## Import

Import is not possible.
