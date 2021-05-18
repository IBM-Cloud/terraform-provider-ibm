---
subcategory: "Catalog Management"
layout: "ibm"
page_title: "IBM : cm_version"
description: |-
  Get information about Catalog Management version.
---

# `ibm_cm_version`

Create, modify, or delete an `cm_version` data source. For more information, about managing catalog version, refer to [updating your software](https://cloud.ibm.com/docs/account?topic=account-update-private).


## Example usage

```
data "cm_version" "cm_version" {
	version_loc_id = "version_loc_id"
}
```


## Argument reference
Review the input parameters that you can specify for your data source. 

- `version_loc_id` - (Required, String) A dotted value of `catalogID.versionID`. 

## Attribute reference
Review the output parameters that you can access after your data source is created. 

- `catalog_id` - (String) The catalog ID.
- `crn` - (String) The CRN version.
- `id` - (String) The unique identifier of the `cm_version`.
- `offering_id` - (String) The offering ID.
- `repo_url` - (String) The URL of the content repository.
- `sha` - (String) The hash of the content.
- `source_url` - (String) The source URL of the content repository, for example, Git repository.
- `tgz_url` - (String) File used to onboard the version.
- `version` - (String) Version of the content type.
