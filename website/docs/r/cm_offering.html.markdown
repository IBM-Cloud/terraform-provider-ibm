---
subcategory: "Catalog Management"
layout: "ibm"
page_title: "IBM : cm_offering"
description: |-
  Manages cm_offering.
---

# ibm\_cm_offering

Provides a resource for cm_offering. This allows cm_offering to be created, updated and deleted.

## Example Usage

```hcl
resource "ibm_cm_offering" "cm_offering" {
  catalog_id = "catalog_id"
  label = "placeholder"
  tags = [ "placeholder" ]
}
```

## Argument Reference

The following arguments are supported:

* `catalog_identifier` - (Required, Forces new resource, string) Catalog identifier.
* `label` - (Optional, Forces new resource, string) Display Name in the requested language.
* `tags` - (Optional, Forces new resrouce, List) List of tags associated with this catalog.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The unique identifier of the cm_offering.
* `url` - The url for this specific offering.
* `crn` - The crn for this specific offering.
* `name` - The programmatic name of this offering.
* `short_description` - Short description in the requested language.
* `long_description` - Long description in the requested language.
* `permit_request_ibm_public_publish` - Is it permitted to request publishing to IBM or Public.
* `ibm_publish_approved` - Indicates if this offering has been approved for use by all IBMers.
* `public_publish_approved` - Indicates if this offering has been approved for use by all IBM Cloud users.
* `public_original_crn` - The original offering CRN that this publish entry came from.
* `publish_public_crn` - The crn of the public catalog entry of this offering.
* `portal_approval_record` - The portal's approval record ID.
* `portal_ui_url` - The portal UI URL.
* `catalog_id` - The id of the catalog containing this offering.
* `catalog_name` - The name of the catalog.
* `disclaimer` - A disclaimer for this offering.
* `repo_info` - Repository info for offerings.
  * `token` - Token for private repos.
  * `type` - Public or enterprise GitHub.
