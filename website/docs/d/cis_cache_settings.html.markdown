---
subcategory: "Internet services"
layout: "ibm"
page_title: "IBM: ibm_cis_cache_settings"
description: |-
  Get information on IBM Cloud Internet Services Cache Settings.
---

# ibm_cis_cache_settings

Imports a read only copy of an existing Internet Services Cache Settings.

## Example Usage

```hcl
data "ibm_cis_cache_settings" "test" {
  cis_id    = data.ibm_cis_cache_settings.test.cis_id
  domain_id = data.ibm_cis_cache_settings.test.domain_id
}
```

## Argument Reference

The following arguments are supported:

- `cis_id` - (Required, string) The resource cis id of the CIS on which zones were created.
- `domain_id` - (Required, string) The resource domain id of the DNS on which zones were created.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

- `caching_level` - cache level setting of a specific zone.
    - `id` - The cache level id 
    - `value` - The cache level value `basic`, `simplified`, `aggressive`.
    - `editable` - The cache level editable value.
    - `modified_on` - The cache level modified date.
- `browser_expiration` - Browser Cache TTL (in seconds) specifies how long CDN edge servers cached resources will remain on your visitors' computers.
    - `id` - The browser expiration ttl type id.
    - `value` - The browser expiration ttl value.
    - `editable` - The browser expiration editable value.
    - `modified_on` - The browser expiration modified date. 
- `development_mode` -  The development mode setting of a specific zone. 
    - `id` - The development mode object id.
    - `value` - The development mode value. `on` and `off`.
    - `editable` - The development mode  editable value. 
    - `modified_on` - The development mode modified date.
- `query_string_sort` -  Enables query string sort setting.
    - `id` - The query string sort cache id.
    - `value` - The query string sort value.`on` and `off`.
    - `editable` - The query string sort editable propery. 
    - `modified_on` - The query string sort modified date.
- `serve_stale_content` -  Serve Stale Content will serve pages from CDN edge servers' cache if your server is offline.
    - `id` - The serve stale content cache id.
    - `value` - The serve stale content value.`on` and `off`.
    - `editable` - The serve stale content editable value. 
    - `modified_on` - The serve stale content modified date.

