---
subcategory: "Catalog Management"
layout: "ibm"
page_title: "IBM : cm_catalog"
description: |-
  Manages cm_catalog.
---

# `ibm_cm_catalog`

Create, modify, or delete an `cm_catalog` resources. You can manage the settings for all catalogs across your account. For more information, about managing catalog, refer to [catalog management settings](https://cloud.ibm.com/docs/account?topic=account-account-getting-started).


## Example usage

```
resource "ibm_cm_catalog" "cm_catalog" {
  label = "placeholder"
  short_description = "placeholder"
}
```


## Argument reference
Review the input parameters that you can specify for your resource. 

- `label` - (Required, Forces new resource, String) The display name in the requested language.
- `short_description` - (Optional, Forces new resource, String) The short description in the requested language.


## Attribute reference
Review the output parameters that you can access after your resource is created. 

- `crn` - (String) The CRN associated with the catalog.
- `id` - (String) The unique identifier of the `cm_catalog`.
- `offerings_url` - (String) The URL path to the offerings.
- `url` - (String) The URL for this specific catalog.


