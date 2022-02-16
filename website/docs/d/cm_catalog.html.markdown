---
subcategory: "Catalog Management"
layout: "ibm"
page_title: "IBM : ibm_cm_catalog"
description: |-
  Get information about ibm_cm_catalog.
---

# ibm_cm_catalog

Provides a read-only data source for `ibm_cm_catalog`. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.


## Example usage

```terraform
data "ibm_cm_catalog" "cm_catalog" {
	catalog_identifier = "catalog_identifier"
}
```

## Argument reference
Review the argument reference that you can specify for your data source. 
 
- `catalog_identifier` - (Required, String) The catalog identifier.


## Attribute reference
In addition to the argument reference list, you can access the following attribute references after your data source is created.

- `catalog_icon_url` - (String) The URL for an icon associated with the catalog.
- `crn` - (String) The CRN associated with the catalog.
- `id` - (String) The unique identifier of the `ibm_cm_catalog`.
- `kind` - (String) Kind of catalog, offering or vpe.
- `label` - (String) Display the name in the requested language.
- `offerings_url` - (String) URL path to the offerings.
- `resource_group_id` - (String) The ID of the resource group this catalog is in
- `short_description` - (String) The description in the requested language.
- `tags` - (String) The list of tags associated with this catalog.
- `url` - (String) The URL for the specific catalog.
