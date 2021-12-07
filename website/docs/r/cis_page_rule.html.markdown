---

subcategory: "Internet services"
layout: "ibm"
page_title: "IBM: ibm_cis_page_rule"
description: |-
  Provides a IBM CIS Page rule resource.
---

# ibm_cis_page_rule
Provides an IBM Cloud Internet Services page rule resource, to create, update, delete page rules of a domain. This resource is associated with an IBM Cloud Internet Services instance and an IBM Cloud Internet Services domain resource. For more information, about IBM Cloud Internet Services page rules, see [using page rules](https://cloud.ibm.com/docs/cis?topic=cis-use-page-rules).

## Example usage

```terraform
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
  actions {
    id    = "minify"
    html  = "off"
    css   = "on"
    js    = "off"
  }
}
```

## Argument reference
Review the argument references that you can specify for your resource. 

- `actions` - (Required, List) The list of actions performed on URL. Minimum items is `1`.

  Nested scheme for `actions`:
  - `id` - (Required, String) The action ID. Valid values are `page rule action field map from console` to `API CF-UI map API`).
  - `css` - (String) The required attribute for `minify` action. CSS supported values are `on` and `off`.
  - `html` - (String) The required attribute for `minify` action. HTML supported values are `on` and `off`.
  - `js` - (String) The required attribute for `minify` action. JS supported values are `on` and `off`.
  - `status_code` - (Optional, String) The status code to check for URL forwarding. The required attribute for `forwarding_url` action. Valid values are `301` and `302`. It returns `0` for all other actions.
  - `url` - (Optional, String) The forward rule URL, a required attribute for `forwarding_url` action.
  - `value` - (Required, String) The values for corresponding actions.
    
    
    ~> **NOTE**
        Below table shows corresponding valid values of `id` and `value` of `actions` argument

      |         id                |                 description                   |valid values|
      |---------------------------|-----------------------------------------------|------------|
      |`always_online`              |The action conflicts with all other settings.	|`on`, `off`|
      |`always_use_https`           |The action conflicts with all other settings.	|The value is not required.|
      |`automatic_https_rewrites`   |The automatic HTTPS rewrites.				 	        |`on`, `off`|
      |`browser_cache_ttl`          |The browser cache TTL.  					  	          |`0`, `1800`, `3600`, `7200`, `10800`, `14400`, `18000`, `28800`, `43200`, `57600`, `72000`, `86400`, `172800`, `259200`, `345600`, `432000`, `691200`, `1382400`, `2073600`, `2678400`, `5356800`, `16070400`, `31536000`|
      |`bypass_cache_on_cookie`     |The bypass cache on cookie.				  	        |cookie tags|
      |`browser_check`              |The browser integrity check.				  	        |`on`, `off`|
      |`cache_level`                |The cache level.							  	              |`bypass`, `aggressive`, `basic`, `simplified`, `cache_everything`|
      |`cache_on_cookie`            |The cache on cookie.						  	            |The cookie value|
      |`cache_deception_armor`      |The cache deception armor.					  	        |`on`, `off`|
      |`disable_security`           |The action conflicts with `email_obfuscation`, `server_side_exclude`, `waf`.|The value is not required.|
      |`disable_apps`               |The disable apps.							  	            |The value is not required|
      |`disable_performance`        |The disable performance.					  	          |The value is not required.|
      |`edge_cache_ttl`             |The edge cache TTL.						  	            |`0`, `30`, `60`, `300`, `600`, `1200`, `1800`, `3600`, `7200`, `10800`, `14400`, `18000`, `28800`, `43200`, `57600`, `72000`, `86400`, `172800`, `259200`, `345600`, `432000`, `518400`, `604800`, `1209600`, `2419200`|
      |`email_obfuscation`          |The Email obfuscation.						  	          |`on`, `off`|
      |`explicit_cache_control`     |The origin cache control.					  	        |`on`, `off`|
      |`forwarding_url`             |The action conflicts with all other settings.	|The value is not required.|
      |`host_header_override`       |The host header override.					  	        |The header value.|
      |`image_load_optimization`  	|The image load optimization.				  	        |`on`, `off`|
      |`image_size_optimization`  	|The image size optimization.				 	          |`on`, `off`|
      |`ip_geolocation`  			      |The IP geography location header.			  	    |`on`, `off`|
      |`opportunistic_encryption`   |The opportunistic encryption.				  	      |`on`, `off`|
      |`origin_error_page_pass_thru`|The origin error page pass-through.		  	    |`on`, `off`|
      |`resolve_override`  			    |The resolve override.						  	          |The value for resolving URL override.|
      |`response_buffering` 		    |The response buffering.					  	          |`on`, `off`|
      |`script_load_optimization`   |The script load optimization.				  	      |`off`, `lossless`, `lossy`|
      |`ssl` 						            |The TLS settings.							  	            |`off`, `flexible`, `full`, `strict`,`origin_pull`|
      |`security_level`  			      |The security level.						  	            |`disable_security`, `always_use_https`|
      |`server_side_exclude`  	  	|The server side excludes.					  	        |`on`, `off`|
      |`server_stale_content`  		  |The server stale content.					  	        |`on`, `off`|
      |`sort_query_string_for_cache`|The sort query string.						  	          |`on`, `off`|
      |`true_client_ip_header`  	  |The true client IP header.					  	        |`on`, `off`|
      |`waf`  						          |The Web Application Firewall.				  	      |`on`, `off`|
      |`minify`  					          |The Minify web content						  	          |The value is not required|
- `cis_id` - (Required, String) The ID of the IBM Cloud Internet Services instance.
- `domain_id` - (Required, String) The ID of the IBM Cloud Internet Services domain.
- `priority` - (Optional, Integer) The priority of the page rule. Default value is `1`. `Set` and `Update` are not supported yet.
- `status` - (Optional, String) The status of the page rule. Valid values are `active` and `disabled`. Default value is `disabled`.
- `targets`- (Required, Set) The targets, where rule is added.

  Nested scheme for `targets`:
  - `target` - (Required, String) The target type. Valid value is `url`.
  - `constraint` -(Required, List) The constraint of the page rule. Maximum item is `1`.

    Nested scheme for `constraint`:
    - `operator` - (Required, String) The operation on the page rule. The valid value is `matches`.
    - `value` - (Required, String) The URL value on which page rule is applied.

## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `id` - (String) The record ID. It is a combination of `<rule_id>:<domain_id>:<cis_id>` attributes of the origin pool.
- `rule_id` - (String) The page rule ID.

## Import
The `ibm_cis_page_rule` resource can be imported by using the ID. The ID is formed from the rule ID, the domain ID of the domain and the CRN concatenated by using a `:` character.

The domain ID and CRN is located on the **Overview** page of the Internet Services instance under the **Domain** heading of the console, or via the `ibmcloud cis` CLI.

- **Domain ID** is a 32 digit character string of the form: `9caf68812ae9b3f0377fdf986751a78f`.

- **CRN** is a 120 digit character string of the format `crn:v1:bluemix:public:internet-svcs:global:a/1aa1111a1a1111aa1a111111111111aa:11aa111a-11a1-1a11-111a-111aaa11a1a1::` 

- **Rule ID** is a 32 digit character string in the format `489d96f0da6ed76251b475971b097205c`.

**syntax**

```
$ terraform import ibm_cis_page_rule.myorg <rule_id>:<domain-id>:<crn>
```
**Example**

```
$ terraform import ibm_cis_page_rule.myorg page_rule 48996f0da6ed76251b475971b097205c:9caf68812ae9b3f0377fdf986751a78f:crn:v1:bluemix:public:internet-svcs:global:a/4ea1882a2d3401ed1e459979941966ea:31fa970d-51d0-4b05-893e-251cba75a7b3::
```
