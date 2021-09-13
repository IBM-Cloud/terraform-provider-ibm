---
subcategory: "Catalog Management"
layout: "ibm"
page_title: "IBM : ibm_cm_version"
description: |-
  Get information about Catalog Management version.
---

# ibm_cm_version

Provides a read-only data source for `ibm_cm_version`. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example usage

```terraform
data "ibm_cm_version" "cm_version" {
	version_loc_id = "version_loc_id"
}
```


## Argument reference
Review the argument reference that you can specify for your data source. 

- `version_loc_id` - (Required, String) A dotted value of `catalogID.versionID`. 

## Attribute reference
In addition to the argument reference list, you can access the following attribute references after your data source is created.

- `catalog_id` - (String) The catalog ID.
- `crn` - (String) The CRN version.
- `id` - (String) The unique identifier of the `ibm_cm_version`.
- `offering_id` - (String) The offering ID.
- `repo_url` - (String) The URL of the content repository.
- `sha` - (String) The hash of the content.
- `source_url` - (String) The source URL of the content repository, for example, Git repository.
- `tgz_url` - (String) File used to onboard the version.
- `version` - (String) Version of the content type.
