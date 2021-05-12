---
subcategory: "Catalog Management"
layout: "ibm"
page_title: "IBM : cm_catalog"
description: |-
  Get information about cm_catalog.
---

# `ibm_cm_catalog`

Create, modify, or delete an `cm_catalog` resources. You can manage the settings for all catalogs across your account. For more information, about managing catalog, refer to [catalog management settings](https://cloud.ibm.com/docs/account?topic=account-account-getting-started).


## Example usage

```
data "cm_catalog" "cm_catalog" {
	catalog_identifier = "catalog_identifier"
}
```

## Argument reference
Review the input parameters that you can specify for your data source. 
 
- `catalog_identifier` - (Required, String) The catalog identifier.


## Attribute reference
Review the output parameters that you can access after your data source is created. 

- `catalog_icon_url` - (String) The URL for an icon associated with the catalog.
- `crn` - (String) The CRN associated with the catalog.
- `id` - (String) The unique identifier of the `cm_catalog`.
- `label` - (String) Display the name in the requested language.
- `offerings_url` - (String) URL path to the offerings.
- `short_description` - (String) The description in the requested language.
- `tags` - (String) The list of tags associated with this catalog.
- `url` - (String) The URL for the specific catalog.


