---
subcategory: "Catalog Management"
layout: "ibm"
page_title: "IBM : cm_catalog"
description: |-
  Manages cm_catalog.
---

# ibm\_cm_catalog

Provides a resource for cm_catalog. This allows cm_catalog to be created, updated and deleted.

## Example Usage

```hcl
resource "ibm_cm_catalog" "cm_catalog" {
  label = "placeholder"
  short_description = "placeholder"
}
```

## Argument Reference

The following arguments are supported:

* `label` - (Required, Forces new resource, string) Display Name in the requested language.
* `short_description` - (Optional, Forces new resource, string) Description in the requested language.

## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier of the cm_catalog.
* `url` - The url for this specific catalog.
* `crn` - CRN associated with the catalog.
* `offerings_url` - URL path to offerings.