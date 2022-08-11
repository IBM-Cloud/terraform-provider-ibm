---
subcategory: "Internet services"
layout: "ibm"
page_title: "IBM: ibm_cis_cache_settings"
description: |-
  Provides a IBM CIS cache settings resource.
---

# ibm_cis_cache_settings
 Provides an IBM Cloud Internet Services cache settings resource. The resource allows to create, update, or delete cache settings of a domain of an IBM Cloud Internet Services CIS instance. For more information about cache setting, see [CIS cache concepts](https://cloud.ibm.com/docs/cis?topic=cis-caching-concepts).

## Example usage

```terraform
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

## Argument reference

Review the argument references that you can specify for your resource. 

- `browser_expiration` - (Optional, Integer) The Browser expiration settings. Valid values are `0, 30, 60, 300, 1200, 1800, 3600, 7200, 10800, 14400, 18000, 28800, 43200, 57600, 72000, 86400, 172800, 259200, 345600, 432000, 691200, 1382400, 2073600, 2678400, 5356800, 16070400, 31536000`.
- `cis_id` - (Required, String) The ID of the IBM Cloud Internet Services instance.
- `caching_level` - (Optional, String) The cache level settings. Valid values are `basic`, `simplified`, `aggressive`.
- `domain_id` - (Required, String) The ID of the domain to change cache settings.
- `development_mode` - (Optional, String) The development mode enable or disable settings. Valid values are `on`, and `off`.
- `purge_all` - (Optional, Bool)  Purge all cached files.
- `purge_by_urls` - (Optional, List of Strings) Purge cached URLs.
- `purge_by_hosts` - (Optional, List of Strings) Purge cached hosts.
- `purge_by_tags` - (Optional, List of Strings) Purge cached item that matches the tags.
- `query_string_sort` - (Optional, String) The query string sort settings. Valid values are `on`, and `off`.
- `serve_stale_content` - (Optional, String) Enable (`on`) or disable (`off`) the serve stale content setting.

**Note**

Among all the purge actions `purge_all`, `purge_by-urls`, `purge_by_hosts`, and `purge_by_tags`, only one is allowed to give inside a resource.

## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `id` - (String) The record ID. It is a combination of `<domain_id>,<cis_id>` attributes concatenated with `:`.

## Import
The `ibm_cis_cache_settings` resource can be imported using the ID. The ID is formed from the domain ID of the domain and the CRN concatenated  using a `:` character.

The domain ID and CRN will be located on the overview page of the IBM Cloud Internet Services instance of the console domain heading, or by using the `ibmcloud cis` command line commands.

- **Domain ID** is a 32 digit character string of the form: `9caf68812ae9b3f0377fdf986751a78f`

- **CRN** is a 120 digit character string of the form: `crn:v1:bluemix:public:internet-svcs:global:a/4ea1882a2d3401ed1e459979941966ea:31fa970d-51d0-4b05-893e-251cba75a7b3::`

**Syntax**

```
$ terraform import ibm_cis_cache_settings.cache_settings <domain-id>:<crn>
```

**Example**

```
$ terraform import ibm_cis_cache_settings.cache_settings 9caf68812ae9b3f0377fdf986751a78f:crn:v1:bluemix:public:internet-svcs:global:a/4ea1882a2d3401ed1e459979941966ea:31fa970d-51d0-4b05-893e-251cba75a7b3::
```

