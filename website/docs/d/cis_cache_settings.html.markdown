---
subcategory: "Internet services"
layout: "ibm"
page_title: "IBM: ibm_cis_cache_settings"
description: |-
  Get information on IBM Cloud Internet Services cache settings.
---

# ibm_cis_cache_settings
Retrieve an information of an existing internet services cache settings. For more information, about understanding CIS cache settings, see [caching concepts](https://cloud.ibm.com/docs/cis?topic=cis-caching-concepts).

## Example usage

```terraform
data "ibm_cis_cache_settings" "test" {
  cis_id    = data.ibm_cis_cache_settings.test.cis_id
  domain_id = data.ibm_cis_cache_settings.test.domain_id
}
```

## Argument reference
Review the argument references that you can specify for your data source. 

- `cis_id` - (Required, String) The resource CIS ID of the CIS on which zones were created.
- `domain_id` - (Required, String) The resource domain ID of the DNS on which zones were created.

## Attribute reference
In addition to all argument reference list, you can access the following attribute references after your data source is created. 

- `browser_expiration` - (String) The browser cache TTL (in seconds) specifies how long `CDN` edge servers cached resources will remain on your visitors' computers.

  Nested scheme for `browser_expiration`:
  - `editable` - (String) The browser expiration editable value.
  - `id` - (String) The browser expiration TTL type ID.
  - `modified_on` - (String) The browser expiration modified date.
  - `value` - (String) The browser expiration TTL value.
- `caching_level` - (String) The cache level setting of a specific zone.

  Nested scheme for `caching_level`:
  - `editable` - (String) The cache level editable value.
  - `id` - (String) The cache level ID.
  - `modified_on` - (String) The cache level modified date.
  - `value` - (String) The cache level value `basic`, `simplified`, or `aggressive`.
- `development_mode` - (String) The development mode settings of a specific zone.

  Nested scheme for `development_mode`:
  - `editable` - (String) The development mode editable value.
  - `id` - (String) The development mode object ID.
  - `modified_on` - (String) The development mode modified date.
  - `value` - (String) The development mode value. Supported values are `on`, and `off`.
- `query_string_sort` - (String) Enables query string sort settings.

  Nested scheme for `query_string_sort`:
	- `editable` - (String) The query string sort editable property.
	- `id` - (String) The query string sort cache ID.
	- `modified_on` - (String) The query string sort modified date.
	- `value` - (String) The query string sort value. Supported values are `on`, and `off`.
- `serve_stale_content` - (String) The serve stale content will serve pages from `CDN` edge servers cache if your server is offline.

  Nested scheme for `server_stale_content`:
	- `editable` - (String) The serve stale content editable value.
	- `id` - (String) The serve stale content cache ID.
	- `modified_on` - (String) The serve stale content modified date.
	- `value` - (String) The serve stale content value. Supported values are `on`, and `off`.
