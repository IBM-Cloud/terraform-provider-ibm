---
subcategory: "Catalog Management"
layout: "ibm"
page_title: "IBM : cm_catalog"
description: |-
  Get information about cm_catalog
---

# ibm\_cm_catalog

Provides a read-only data source for cm_catalog. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "cm_catalog" "cm_catalog" {
	catalog_identifier = "catalog_identifier"
}
```

## Argument Reference

The following arguments are supported:

* `catalog_identifier` - (Required, string) Catalog identifier.

## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier of the cm_catalog.

* `label` - Display Name in the requested language.

* `short_description` - Description in the requested language.

* `catalog_icon_url` - URL for an icon associated with this catalog.

* `tags` - List of tags associated with this catalog.

* `url` - The url for this specific catalog.

* `crn` - CRN associated with the catalog.

* `offerings_url` - URL path to offerings.