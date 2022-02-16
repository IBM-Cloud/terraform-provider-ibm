---
subcategory: "Catalog Management"
layout: "ibm"
page_title: "IBM : cm_catalog"
description: |-
  Manages cm_catalog.
---

# ibm_cm_catalog

Create, modify, or delete an `cm_catalog` resources. You can manage the settings for all catalogs across your account. For more information, about managing catalog, refer to [catalog management settings](https://cloud.ibm.com/docs/account?topic=account-account-getting-started).


## Example usage

```terraform
resource "ibm_cm_catalog" "cm_catalog" {
  label = "placeholder"
  short_description = "placeholder"
}
```


## Argument reference
Review the argument reference that you can specify for your resource. 

- `label` - (Required, Forces new resource, String) The display name in the requested language.
- `kind` - (Optional, Forces new resource, Defaults to offering, String) The kind of the catalog, offering or vpe.
- `resource_group_id` - (Optional, Forces new resource, String) The ID of the resource group this catalog will be provisioned in
- `short_description` - (Optional, Forces new resource, String) The short description in the requested language.


## Attribute reference
In addition to all argument references list, you can access the following attribute references after your resource is created. 

- `crn` - (String) The CRN associated with the catalog.
- `id` - (String) The unique identifier of the `cm_catalog`.
- `offerings_url` - (String) The URL path to the offerings.
- `url` - (String) The URL for this specific catalog.
