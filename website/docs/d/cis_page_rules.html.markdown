---
subcategory: "Internet services"
layout: "ibm"
page_title: "IBM: ibm_cis_page_rules"
description: |-
  Get information on an IBM Cloud Internet Services page rules.
---

# ibm_cis_page_rules
Retrieve an information of an IBM Cloud Internet Services page rules resource. For more information, about IBM Cloud Internet Services page rules, see [using page rules](https://cloud.ibm.com/docs/cis?topic=cis-use-page-rules).

## Example usage

```terraform
data "ibm_cis_page_rules" "rules" {
  cis_id    = ibm_cis.instance.id
  domain_id = ibm_cis_domain.example.id
}
```

## Argument reference
Review the argument references that you can specify for your data source. 

- `cis_id` - (Required, String) The ID of the IBM Cloud Internet Services instance .
- `domain_id` - (Required, String) The ID of the domain.


## Attribute reference
In addition to all argument reference list, you can access the following attribute references after your data source is created. 

- `cis_page_rules` - (String) The page rules detail.

  Nested scheme for `cis_page_rules`:
	- `priority` - (String) The priority of the page rule.
	- `status` - (String)   The status of the page rule. Default value is `active`.
	- `rule_id` - (String) The page rule ID.
	- `targets` - (String)   The targets, of the added rule.

	  Nested scheme for `targets`:
		- `constraint` - (String) The constraint of the page rule.

		  Nested scheme for `constraint`:
			- `operator` - (String) The operation on page rule. Valid value is `matches`.
			- `value` - (String) The URL value on the applied page rule.
		- `target` - (String) The target type. Valid value is `url`.
	- `actions` - (String)   The actions to be performed on the URL.

	  Nested scheme for `actions`:
		- `id` - (String) The action ID. Valid values are `page rule action field map from console` to `API CF-UI map API`).

		  Nested scheme for `id`:
		  - `automatic_https_rewrites` - (String) The automatic HTTPS rewrites.
		  - `always_use_https` - (String) The action conflicts with all other settings.
		  - `always_online` - (String) The action conflicts with all other settings.
		  - `browser_cache_ttl` - (String) The browser cache TTL.
		  - `bypass_cache_on_cookie` - (String) The bypass cache on cookie.
		  - `browser_check` - (String) The browser integrity check.
		  - `cache_deception_armor` - (String) The cache deception armor.
		  - `cache_level` - (String) The cache level.
		  - `cache_on_cookie` - (String) The cache on cookie.
		  - `disable_apps` - (String) The disable apps.
		  - `disable_performance` - (String) The disable
		  - `disable_security` - (String) The action conflicts with `email_obfuscation`, `server_side_exclude`, `waf`.
		  - `edge_cache_ttl` - (String) The edge cache TTL.
		  - `email_obfuscation` - (String) The Email obfuscation.
		  - `explicit_cache_control` - (String) The origin cache control.
		  - `forwarding_url` - (String) The action conflicts with all other settings.
		  - `host_header_override` - (String) The host header override.
		  - `image_load_optimization` - (String) The image load optimization.
		  - `image_size_optimization` - (String) The image size optimization.
		  - `ip_geolocation` - (String) The IP geography location header.
		  - `origin_error_page_pass_thru` - (String) The origin error page pass-through.
		  - `opportunistic_encryption` - (String) The opportunistic encryption.
		  - `resolve_override` - (String) The resolve override.
			 performance.
		  - `response_buffering` - (String) The response buffering.
		  - `script_load_optimization` - (String) The script load optimization.
	   	  - `ssl` - (String) The TLS settings.
		  - `security_level` - (String) The security level.
		  - `server_side_exclude` - (String) The server side excludes.
		  - `server_stale_content` - (String) The server stale content.
		  - `sort_query_string_for_cache` - (String) The sort query string.
		  - `true_client_ip_header` - (String) The true client IP header.
		  - `waf` - (String) The Web Application Firewall.
		  - `minify` - (String) The Minify web content.
		- `value` - (String) The values for corresponding actions.

		  Nested scheme for `value`:
		  - `always_online` - (String) The valid values are `on`, `off`.
		  - `automatic_https_rewrites` - (String) The valid values are `on`, `off`.
		  - `browser_cache_ttl`- (Integer) The valid values are `0, 1800, 3600, 7200, 10800, 14400, 18000, 28800, 43200, 57600, 72000, 86400, 172800, 259200, 345600, 432000, 691200, 1382400, 2073600, 2678400, 5356800, 16070400, 31536000`.
		  - `browser_check` - (String) The valid values are `on`, `off`.
		  - `bypass_cache_on_cookie` - (String) The valid values are `cookie tags`.
		  - `cache_deception_armor` - (String) The valid values are `on`, `off`.
		  - `cache_on_cookie` - (String) The cookie value.
		  - `cache_level` - (String) The valid values are `bypass`, `aggressive`, `basic`, `simplified`, `cache_everything`.
		  - `edge_cache_ttl` - (String) The valid values are `0, 30, 60, 300, 600, 1200, 1800, 3600, 7200, 10800, 14400, 18000, 28800, 43200, 57600, 72000, 86400, 172800, 259200, 345600, 432000, 518400, 604800, 1209600, 2419200`.
		  - `disable_apps` - (String) The value is not required.
		  - `disable_performance` - (String) The value is not required.
		  - `email_obfuscation` - (String) The valid values are `on`, `off`.
		  - `explicit_cache_control` - (String) The valid values are `on`, `off`.
		  - `host_header_override` - (String) The header value.
		  - `ip_geolocation` - (String) The valid values are `on`, `off`.
		  - `image_load_optimization` - (String) The valid values are `on`, `off`.
		  - `image_size_optimization` - (String) The valid values are `on`, `off`.
		  - `opportunistic_encryption` - (String) The valid values are `on`, `off`.
		  - `origin_error_page_pass_thru` - (String) The valid values are `on`, `off`.
		  - `resolve_override` - (String) The value for resolving URL override.
		  - `response_buffering` - (String) The valid values are `on`, `off`.
		  - `script_load_optimization` - (String) The valid values are `off`, `lossless`, `lossy`.
		  - `ssl` - (String) The valid values are `off`, `flexible`, `full`, `strict`, `origin_pull`.
		  - `security_level` - (String) The valid values are `disable_security`, `always_use_https`.
		  - `server_side_exclude` - (String) The valid values are `on`, `off`.
		  - `server_stale_content` - (String) The valid values are `on`, `off`.
		  - `sort_query_string_for_cache` - (String) The valid values are `on`, `off`.
		  - `true_client_ip_header` - (String) The valid values are `on`, `off`.
		  - `waf` - (String) The valid values are `on`, `off`.
	- `status_code` - (String) The status code to check for URL forwarding. The required attribute for `forwarding_url` action. Valid values are `301` and `302`. It returns `0` for all other actions.
	- `url` - (String) The forward rule URL, a required attribute for `forwarding_url` action.
	- `css` - (String) The required attribute for `minify` action. CSS supported values are `on` and `off`.
    - `html` - (String) The required attribute for `minify` action. HTML supported values are `on` and `off`.
    - `js` - (String) The required attribute for `minify` action. JS supported values are `on` and `off`.