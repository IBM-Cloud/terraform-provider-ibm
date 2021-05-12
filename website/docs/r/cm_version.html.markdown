---
subcategory: "Catalog Management"
layout: "ibm"
page_title: "IBM : cm_version"
description: |-
  Manages cm_version.
---

# `ibm_cm_version`

Create, modify, or delete an `cm_version` resources. For more information, about managing catalog version, refer to [updating your software](https://cloud.ibm.com/docs/account?topic=account-update-private).


## Example usage

```
resource "cm_version" "cm_version" {
  catalog_identifier = "catalog_identifier"
  offering_id = "offering_id"
  zipurl = "placeholder"
}
```


## Argument reference
Review the input parameters that you can specify for your resource. 
 
- `catalog_identifier` - (Required, Forces new resource, String) Catalog identifier.
- `content` - (Optional, Forces new resource, String) The byte array representing the content to import. Currently supports only `OVA` images.
- `offering_id` - (Required, Forces new resource, String) Offering identification.
- `tags` - (Optional, Forces new resource, List) The tags array.
- `target_kinds` - (Optional, Forces new resource, List) The target kinds. Supported values are `iks`, `roks`, `vcenter`, and `terraform`.
- `target_version` - (Optional, Forces new resource, String) The semver value for the new version, if not found in the `zip` URL package content.
- `zipurl` - (Optional, Forces new resource, String) The URL path to `.zip` location. If not specified, must provide content in the body of the call.


## Attribute reference
Review the output parameters that you can access after your resource is created. 

- `catalog_id` - (String) The catalog ID.
- `crn` - (String) The CRN version.
- `id` - (String) The unique identifier and version locator of the version.
- `kind_id` - (String) The kind ID.
- `repo_url` - (String) The URL of the content repository.
- `sha` - (String) The hash of the content.
- `source_url` - (String) The source URL of the content repository, for example, Git repository.
- `tgz_url` - (String) File used to onboard the version.
- `url` - (String) The URL for the specific offering.
- `version` - (String) Version of the content type.
