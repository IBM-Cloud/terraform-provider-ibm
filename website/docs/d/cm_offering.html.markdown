---
subcategory: "Catalog Management"
layout: "ibm"
page_title: "IBM : cm_offering"
description: |-
  Get information about cm_offering
---

# ibm\_cm_offering

Provides a read-only data source for cm_offering. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "cm_offering" "cm_offering" {
	catalog_identifier = "catalog_identifier"
	offering_id = "offering_id"
}
```

## Argument Reference

The following arguments are supported:

* `catalog_identifier` - (Required, string) Catalog identifier.
* `offering_id` - (Required, string) Offering identification.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The unique identifier of the cm_offering.

* `url` - The url for this specific offering.

* `crn` - The crn for this specific offering.

* `label` - Display Name in the requested language.

* `name` - The programmatic name of this offering.

* `offering_icon_url` - URL for an icon associated with this offering.

* `offering_docs_url` - URL for an additional docs with this offering.

* `offering_support_url` - URL to be displayed in the Consumption UI for getting support on this offering.

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

* `hidden` - Determine if this offering should be displayed in the Consumption UI.

* `provider` - Provider of this offering.

* `repo_info` - Repository info for offerings. Nested `repo_info` blocks have the following structure:
	* `token` - Token for private repos.
	* `type` - Public or enterprise GitHub.

