---
subcategory: "Catalog Management"
layout: "ibm"
page_title: "IBM : cm_offering"
description: |-
  Get information about cm_offering.
---

# `ibm_cm_offering`

Create, modify, or delete an `cm_offering` data source. You can manage the settings for all catalogs across your account. For more information, about managing catalog, refer to [catalog management settings](https://cloud.ibm.com/docs/account?topic=account-account-getting-started).


## Example usage

```
data "cm_offering" "cm_offering" {
	catalog_identifier = "catalog_identifier"
	offering_id = "offering_id"
}
```


## Argument reference
Review the input parameters that you can specify for your data source. 

- `catalog_identifier` - (Required, String) The catalog identifier.
- `offering_id` - (Required, String) The offering identification.


## Attribute reference
Review the output parameters that you can access after your data source is created. 

- `catalog_id` - (String) The ID of the catalog containing this offering.
- `catalog_name` - (String) The name of the catalog.
- `crn` - (String) The CRN for the specific offering.
- `disclaimer` - (String) A disclaimer for the offering.
- `hidden` - (String) Determine if the offering should be displayed in the consumption console.
- `ibm_publish_approved` - (String) Indicates if the offering has been approved for use by all IBMers.
- `id` - (String) The unique identifier of the `cm_offering`.
- `label` - (String) Display the name in the requested language.
- `long_description` - (String) The long description in the requested language.
- `name` - (String) The programmatic name of the offering.
- `offering_icon_url` - (String) The URL for an icon associated with the offering.
- `offering_docs_url` - (String) The URL for an extra documentation with the offering.
- `offering_support_url` - (String) The URL to be displayed in the consumption console for getting support on the offering.
- `permit_request_ibm_public_publish` - (String) Is it permitted to request publishing to IBM or public.
- `public_publish_approved` - (String) Indicates if the offering has been approved to all IBM Cloud users.
- `public_original_crn` - (String) The original offering CRN that is published.
- `publish_public_crn` - (String) The CRN of the public catalog entry of the offering.
- `portal_approval_record` - (String) The portal's approval record ID.
- `portal_ui_url` - (String) The portal console URL.
- `provider` - (String) Provider of this offering.
- `repo_info` - (List) Repository information for offerings. Nested `repo_info` blocks have the following structure.
	- `token` - (String) The token for the private repository.
	- `type` - (String) The public or enterprise GitHub.
- `short_description` - (String) The short description in the requested language.
- `url` - (String) The URL for the specific offering.

