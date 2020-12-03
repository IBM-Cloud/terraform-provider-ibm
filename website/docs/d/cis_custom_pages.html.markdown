---
layout: "ibm"
page_title: "IBM: ibm_cis_custom_pages"
sidebar_current: "docs-ibm-datasource-cis-custom-pages"
description: |-
  Get information on an IBM Cloud Internet Services Custom Pages resource.
---

# ibm_cis_custom_pages

Imports a read only copy of an existing Internet Services custom pages resource.

## Example Usage

```hcl
# Get custom pages of the domain

data "ibm_cis_custom_pages" "custom_pages" {
    cis_id    = data.ibm_cis.cis.id
    domain_id = data.ibm_cis_domain.cis_domain.domain_id
}
```

## Argument Reference

The following arguments are supported:

- `id` - The custom page ID. It is a combination of <`page_id`>,<`domain_id`>,<`cis_id`> attributes concatenated with ":".
- `cis_id` - The ID of the CIS service instance.
- `domain_id` - The ID of the domain to change Custom Page.
- `page_id` - The Custom page identifier. Valid values are `basic_challenge, waf_challenge, waf_block, ratelimit_block, country_challenge, ip_block, under_attack, 500_errors, 1000_errors, always_online`
- `url` - The URL for custom page settings. By default `url` is set with empty string `""`. If this field is being set with empty string, when it is already set with empty string, then it throws error.
- `description` - The description of custom page.
- `required_tokens` - The custom page required token which is expected from `url` page.
- `preview_target` - The custom page target
- `state` - The custom page state. This is set `default` when there is empty `url` and set to `customized` when `url` is set with some url.
- `created_on` - The custom page created date and time.
- `modified_on` - The custom page modified date and time.
