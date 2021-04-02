---
subcategory: "Catalog Management"
layout: "ibm"
page_title: "IBM : cm_version"
description: |-
  Get information about cm_version
---

# ibm\_cm_version

Provides a read-only data source for cm_version. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "cm_version" "cm_version" {
	version_loc_id = "version_loc_id"
}
```

## Argument Reference

The following arguments are supported:

* `version_loc_id` - (Required, string) A dotted value of `catalogID`.`versionID`.

## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier of the cm_version.

* `crn` - Version's CRN.

* `version` - Version of content type.

* `sha` - hash of the content.

* `offering_id` - Offering ID.

* `catalog_id` - Catalog ID.

* `repo_url` - Content's repo URL.

* `source_url` - Content's source URL (e.g git repo).

* `tgz_url` - File used to on-board this version.


