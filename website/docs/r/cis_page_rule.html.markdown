---
layout: "ibm"
page_title: "IBM: ibm_cis_page_rule"
sidebar_current: "docs-ibm-resource-cis-page-rule"
description: |-
  Provides a IBM CIS Page Rule resource.
---

# ibm_cis_page_rule

Provides a IBM CIS Page Rule resource. This resource is associated with an IBM Cloud Internet Services instance and a CIS Domain resource. It allows to create, update, delete page rules of a domain of a CIS instance

## Example Usage

```hcl
# Add a page rule to the domain

resource "ibm_cis_page_rule" "page_rule" {
  cis_id    = var.cis_crn
  domain_id = var.zone_id
  targets {
    target = "url"
    constraint {
      operator = "matches"
      value    = "example.com"
    }
  }
  actions {
    id    = "email_obfuscation"
    value = "on"
  }
  actions {
    id          = "forwarding_url"
    url         = "https://ibm.travis-kuganes1.sdk.cistest-load.com/*"
    status_code = 302
  }
}
```

## Argument Reference

The following arguments are supported:

- `cis_id` - (Required,string) The ID of the CIS service instance.
- `domain_id` - (Required,string) The ID of the domain.
- `priority` - (Optional,integer) The priority of page rule. Default value is `1`. Set/Update not supported yet.
- `status` - (Optional,string) The status of page rule. Valid values are `active` and `disabled`. Default value is `disabled`.
- `targets` - (Required,set) The Targets, which rule is added.
  - `target` - (Required,string) The Target type. Valid value: `url`.
  - `constraint` - (Required,list(MaxItems: 1)) The Constrint of page rule.
    - `operator` - (Required,string) The Operation on page rule. Valid value is `matches`.
    - `value` - (Required,string) The URL value on which page rule is applied.
- `actions` - (Required, list(MinItems: 1)) The List ofactions to be performed on url.
  - `id` - (Required, string) The action identifier. Valid values are : (Page rule action field map from UI to API CF-UI map API)
    - 'Disable Security': `disable_security` - this action conflicts with `email_obfuscation`, `server_side_exclude`, `waf` ,
    - 'Always Online': `always_online` - this action conflicts with with all other settings,
    - 'Forwarding URL': `forwarding_url` - this action conflicts with all other settings,
    - 'Always Use HTTPS': `always_use_https` - this actions conflicts with all other settings,
    - 'TLS': `ssl`,
    - 'Browser Cache TTL': `browser_cache_ttl`,
    - 'Security Level': `security_level`,
    - 'Cache Level': `cache_level`,
    - 'Edge Cache TTL': `edge_cache_ttl`,
    - 'Bypass Cache on Cookie': `bypass_cache_on_cookie`,
    - 'Browser Integrity Check': `browser_check`,
    - 'Server Side Excludes': `server_side_exclude`,
    - 'Server stale content': `serve_stale_content`,
    - 'Email Obfuscation': `email_obfuscation`,
    - 'Automatic HTTPS Rewrites': `automatic_https_rewrites`,
    - 'Opportunistic Encryption': `opportunistic_encryption`,
    - 'IP Geolocation Header': `ip_geolocation`,
    - 'Origin Cache Control': `explicit_cache_control`,
    - 'Cache Deception Armor': `cache_deception_armor`,
    - 'Web Application Firewall': `waf`,
    - 'Host header override': `host_header_override`,
    - 'Resolve override': `resolve_override`,
    - 'Cache on cookie': `cache_on_cookie`,
    - 'Disable apps': `disable_apps`,
    - 'Disable Performance': `disable_performance`,
    - 'Image load optimization': `image_load_optimization`,
    - 'Origin error page pass-through': `origin_error_page_pass_thru`,
    - 'Response buffering': `response_buffering`,
    - 'Image size optimization': `image_size_optimization`,
    - 'Script load optimization': `script_load_optimization`,
    - 'True client IP header': `true_client_ip_header`,
    - 'Sort query string': `sort_query_string_for_cache`,
  - `value` - (Required, string) The Values for corresponding actions are below,
    - `always_online` - valid values: `on`, `off`.
    - `ssl` - valid values: `off`, `flexible`, `full`, `strict`, `origin_pull`.
    - `browser_cache_ttl` - valid values: `0`, `1800`, `3600`, `7200`, `10800`, `14400`, `18000`, `28800`, `43200`, `57600`, `72000`, `86400`, `172800`, `259200`, `345600`, `432000`, `691200`, `1382400`, `2073600`, `2678400`, `5356800`, `16070400`, `31536000`.
    - `security_level` - valid values: `disable_security`, `always_use_https`.
    - `cache_level` - valid values: `bypass`, `aggressive`, `basic`, `simplified`, `cache_everything`.
    - `edge_cache_ttl` - valid values: `0`, `30`, `60`, `300`, `600`, `1200`, `1800`, `3600`, `7200`, `10800`, `14400`, `18000`, `28800`, `43200`, `57600`, `72000`, `86400`, `172800`, `259200`, `345600`, `432000`, `518400`, `604800`, `1209600`, `2419200`.
    - `bypass_cache_on_cookie` - valid values: (string) cookie tags.
    - `browser_check` - valid values: `on`, `off`.
    - `server_side_exclude` - valid values: `on`, `off`.
    - `serve_stale_content` - valid values: `on`, `off`.
    - `email_obfuscation` - valid values: `on`, `off`.
    - `automatic_https_rewrites` - valid values: `on`, `off`.
    - `opportunistic_encryption` - valid values: `on`, `off`.
    - `ip_geolocation` - valid values: `on`, `off`.
    - `explicit_cache_control` - valid values: `on`, `off`.
    - `cache_deception_armor` - valid values: `on`, `off`.
    - `waf` - valid values: `on`, `off`.
    - `host_header_override` - (string) Header value.
    - `resolve_override` - (string) The value for resolving URL override.
    - `cache_on_cookie` - (string) The cookie value.
    - `disable_apps` - no value required.
    - `disable_performance` - no value required.
    - `image_load_optimization` - valid values: `on`, `off`.
    - `origin_error_page_pass_thru` - valid values: `on`, `off`.
    - `response_buffering` - valid values: `on`, `off`.
    - `image_size_optimization` - valid values: `off`, `lossless`, `lossy`.
    - `script_load_optimization` - valid values: `on`, `off`.
    - `true_client_ip_header` - valid values: `on`, `off`.
    - `sort_query_string_for_cache` - valid values: `on`, `off`.
    - `minify` - This is not supported yet.
  - `url` - (Optional,string) The URL for forward rule. This attribute is required for `forwarding_url` action.
  - `status_code` - (Optional,string) The Status code to check for url forwarding. This attribute is required for `forwarding_url` action. valid values are: `301` and `302`. This returns `0` for all other actions.

## Attributes Reference

The following attributes are exported:

- `id` - The record ID. It is a combination of <`rule_id`>,<`domain_id`>,<`cis_id`> attributes concatenated with ":".
- `rule_id` - The Page rule ID.

## Import

The `ibm_cis_page_rule` resource can be imported using the `id`. The ID is formed from the `Rule ID`, the `Domain ID` of the domain and the `CRN` (Cloud Resource Name) concatentated using a `:` character.

The Domain ID and CRN will be located on the **Overview** page of the Internet Services instance under the **Domain** heading of the UI, or via using the `ibmcloud cis` CLI commands.

- **Domain ID** is a 32 digit character string of the form: `9caf68812ae9b3f0377fdf986751a78f`

- **CRN** is a 120 digit character string of the form: `crn:v1:bluemix:public:internet-svcs:global:a/4ea1882a2d3401ed1e459979941966ea:31fa970d-51d0-4b05-893e-251cba75a7b3::`

- **Rule ID** is a 32 digit character string of the form: `489d96f0da6ed76251b475971b097205c`.

```
$ terraform import ibm_cis_page_rule.myorg <rule_id>:<domain-id>:<crn>

$ terraform import ibm_cis_page_rule.myorg page_rule 48996f0da6ed76251b475971b097205c:9caf68812ae9b3f0377fdf986751a78f:crn:v1:bluemix:public:internet-svcs:global:a/4ea1882a2d3401ed1e459979941966ea:31fa970d-51d0-4b05-893e-251cba75a7b3::
```
