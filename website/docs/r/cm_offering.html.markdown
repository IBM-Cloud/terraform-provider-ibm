---
subcategory: "Catalog Management"
layout: "ibm"
page_title: "IBM : cm_offering"
description: |-
  Manages cm_offering.
---

# `ibm_cm_offering`

Create, modify, or delete an `cm_offering` resources. You can manage the settings for all catalogs across your account. For more information, about managing catalog, refer to [catalog management settings](https://cloud.ibm.com/docs/account?topic=account-account-getting-started).


## Example usage

```
resource "ibm_cm_offering" "cm_offering" {
  catalog_id = "catalog_id"
  label = "placeholder"
  tags = [ "placeholder" ]
}
```


## Argument reference
Review the input parameters that you can specify for your resource. 

- `catalog_identifier` - (Required, Forces new resrouce, String) Catalog identifier.
- `label` - (Optional, Forces new resrouce, String) Display the name in the requested language.
- `tags` - (Optional, Forces new resrouce, List) The list of tags associated with the catalog.

## Attribute reference
Review the output parameters that you can access after your resource is created. 

- `catalog_id` - (String) The ID of the catalog containing this offering.
- `catalog_name` - (String) The name of the catalog.
- `crn` - (String) The CRN for the specific offering.
- `disclaimer` - (String) A disclaimer for the offering.
- `ibm_publish_approved` - (String) Indicates if the offering has been approved for use by all IBMers.
- `id` - (String) The unique identifier of the `cm_offering`.
- `long_description` - (String) The long description in the requested language.
- `name` - (String) The programmatic name of the offering.
- `permit_request_ibm_public_publish` - (String) Is it permitted to request publishing to IBM or public.
- `public_publish_approved` - (String) Indicates if the offering has been approved for use by all IBM Cloud users.
- `public_original_crn` - (String) The original offering CRN has published.
- `publish_public_crn` - (String) The CRN of the public catalog entry of an offering.
- `portal_approval_record` - (String) The portal's approval record ID.
- `portal_ui_url` - (String) The portal console URL.
- `repo_info` - (String) Repository information for an offerings.
  - `token` - (String) Token for the private repository.
  - `type` - (String) The public or enterprise GitHub.
- `short_description` - (String) The short description in the requested language.
- `url` - (String) The URL for the specific offering.
