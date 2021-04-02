---
subcategory: "Catalog Management"
layout: "ibm"
page_title: "IBM : cm_version"
description: |-
  Manages cm_version.
---

# ibm\_cm_version

Provides a resource for cm_version. This allows cm_version to be created, updated and deleted.

## Example Usage

```hcl
resource "cm_version" "cm_version" {
  catalog_identifier = "catalog_identifier"
  offering_id = "offering_id"
  zipurl = "placeholder"
}
```

## Argument Reference

The following arguments are supported:

* `catalog_identifier` - (Required, Forces new resource, string) Catalog identifier.
* `offering_id` - (Required, Forces new resource, string) Offering identification.
* `tags` - (Optional, Forces new resource, List) Tags array.
* `target_kinds` - (Optional, Forces new resource, List) Target kinds.  Current valid values are 'iks', 'roks', 'vcenter', and 'terraform'.
* `content` - (Optional, Forces new resource, TypeString) byte array representing the content to be imported.  Only supported for OVA images at this time.
* `zipurl` - (Optional, Forces new resource, string) URL path to zip location.  If not specified, must provide content in the body of this call.
* `target_version` - (Optional, Forces new resource, string) The semver value for this new version, if not found in the zip url package content.

## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier and version locator of the version.
* `crn` - Version's CRN.
* `version` - Version of content type.
* `sha` - hash of the content.
* `catalog_id` - Catalog ID.
* `kind_id` - Kind ID.
* `repo_url` - Content's repo URL.
* `source_url` - Content's source URL (e.g git repo).
* `tgz_url` - File used to on-board this version.
