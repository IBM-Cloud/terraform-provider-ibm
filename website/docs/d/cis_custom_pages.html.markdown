---
subcategory: "Internet services"
layout: "ibm"
page_title: "IBM: ibm_cis_custom_pages"
description: |-
  Get information on an IBM Cloud Internet Services custom pages resource.
---

# ibm_cis_custom_pages
Retrieve information of an existing IBM Cloud Internet Services custom pages resource. For more information, about custom page, refer to [CIS custom page](https://cloud.ibm.com/docs/cis?topic=cis-custom-page).

## Example usage

```terraform
# Get custom pages of the domain

data "ibm_cis_custom_pages" "custom_pages" {
    cis_id    = data.ibm_cis.cis.id
    domain_id = data.ibm_cis_domain.cis_domain.domain_id
}
```

## Argument reference
Review the argument references that you can specify for your data source. 

- `cis_id` - (String) The ID of the CIS service instance.
- `created_on` - (String) Created date and time of the custom page.
- `description` - (String) The description of the custom page.
- `domain_id` - (String) The domain ID to change custom page.
- `id` - (String) The custom page ID. It is a combination of `<page_id>, <domain_id>, <cis_id>` attributes concatenated with `:`.
- `modified_on` - (String) Modified date and time of the custom page.
- `page_id ` - (String) The custom page identifier. Valid values are `basic_challenge`, `waf_challenge`, `waf_block`, `ratelimit_block`, `country_challenge`, `ip_block`, `under_attack`, `500_errors`, `1000_errors`, `always_online`.
- `preview_target` - (String) The target custom page.
- `required_tokens` - (String) The custom page required token which is expected from the URL page.
- `state` - (String) The custom page state. This is set default when there is an empty URL and can customize when URL is set with some URL.
- `url` - (String) The URL for custom page settings. By default URL is set with empty string `""`. Setting a duplicate empty string throws an error.
