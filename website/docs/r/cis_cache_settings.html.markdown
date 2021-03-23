---

subcategory: "Internet services"
layout: "ibm"
page_title: "IBM: ibm_cis_cache_settings"
sidebar_current: "docs-ibm-resource-cis-cache-settings"
description: |-
  Provides a IBM CIS Cache Settings resource.
---

# ibm_cis_cache_settings

Provides a IBM CIS Cache Settings resource. This resource is associated with an IBM Cloud Internet Services instance and a CIS Domain resource. It allows to change Cache Settings of a domain of a CIS instance

## Example Usage

```hcl
# Change Cache Settings of the domain

resource "ibm_cis_cache_settings" "cache_settings" {
  cis_id             = data.ibm_cis.cis.id
  domain_id          = data.ibm_cis_domain.cis_domain.domain_id
  caching_level      = "aggressive"
  browser_expiration = 14400
  development_mode   = "off"
  query_string_sort  = "off"
  purge_all          = true
  serve_stale_content = "off"
}
```

## Argument Reference

The following arguments are supported:

- `cis_id` - (Required,string) The ID of the CIS service instance.
- `domain_id` - (Required,string) The ID of the domain to change Cache Settings.
- `caching_level` - (Optional, string) The cache level setting. Valid values are `basic`, `simplified`, `aggressive`.
- `browser_expiration` - (Optional, integer) The Browser expiration setting. Valid values are `0, 30, 60, 300, 1200, 1800, 3600, 7200, 10800, 14400, 18000, 28800, 43200, 57600, 72000, 86400, 172800, 259200, 345600, 432000, 691200, 1382400, 2073600, 2678400, 5356800, 16070400, 31536000`.
- `development_mode` - (Optional, string) The Development mode enable/disable setting. Valid values are `on` and `off`.
- `query_string_sort` - (Optional, string) The Query string sort settings. Valid values are `on` and `off`.
- `purge_all` - (Optional, boolean) Purge all cached files.
- `purge_by_urls` - (Optional, list(string)) Purge cached urls.
- `purge_by_hosts` - (Optional, list(string)) Purge cached hosts.
- `purge_by_tags` - (Optional, list(string)) Purge cached item which matches the tags.
- `serve_stale_content` - (Optional, string) The Serve stale content  enable/disable setting. Valid values are `on` and `off`.

**Note**:

- Among all the purge actions `purge_all`, `purge_by-urls`, `purge_by_hosts` and `purge_by_tags`, only one is allowed to give inside a resource.
- `serve_stale_content` is supported now.

## Attributes Reference

The following attributes are exported:

- `id` - The record ID. It is a combination of <`domain_id`>,<`cis_id`> attributes concatenated with ":".

## Import

The `ibm_cis_cache_settings` resource can be imported using the `id`. The ID is formed from the `Domain ID` of the domain and the `CRN` (Cloud Resource Name) concatentated using a `:` character.

The Domain ID and CRN will be located on the **Overview** page of the Internet Services instance under the **Domain** heading of the UI, or via using the `ibmcloud cis` CLI commands.

- **Domain ID** is a 32 digit character string of the form: `9caf68812ae9b3f0377fdf986751a78f`

- **CRN** is a 120 digit character string of the form: `crn:v1:bluemix:public:internet-svcs:global:a/4ea1882a2d3401ed1e459979941966ea:31fa970d-51d0-4b05-893e-251cba75a7b3::`

```
$ terraform import ibm_cis_cache_settings.cache_settings <domain-id>:<crn>

$ terraform import ibm_cis_cache_settings.cache_settings 9caf68812ae9b3f0377fdf986751a78f:crn:v1:bluemix:public:internet-svcs:global:a/4ea1882a2d3401ed1e459979941966ea:31fa970d-51d0-4b05-893e-251cba75a7b3::
```
