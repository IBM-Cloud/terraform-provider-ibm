---

subcategory: "Internet services"
layout: "ibm"
page_title: "IBM: ibm_cis_domain_settings"
description: |-
  Provides a resource which customizes IBM Cloud Internet Services domain settings.
---

# ibm_cis_domain_settings

Customize the IBM Cloud Internet Services domain settings. For more information, about Internet Services domain settings, see [adding domains to your CIS instance](https://cloud.ibm.com/docs/cis?topic=cis-multi-domain-support).

## Example usage

```terraform
resource "ibm_cis_domain_settings" "test_domain_settings" {
  cis_id    = data.ibm_cis.cis.id
  domain_id = data.ibm_cis_domain.cis_domain.domain_id
  dnssec                      = "disabled"
  waf                         = "off"
  ssl                         = "flexible"
  min_tls_version             = "1.2"
  cname_flattening            = "flatten_all"
  opportunistic_encryption    = "off"
  automatic_https_rewrites    = "on"
  always_use_https            = "off"
  ipv6                        = "off"
  browser_check               = "off"
  hotlink_protection          = "off"
  http2                       = "on"
  image_load_optimization     = "off"
  image_size_optimization     = "lossless"
  ip_geolocation              = "off"
  origin_error_page_pass_thru = "off"
  brotli                      = "off"
  pseudo_ipv4                 = "off"
  prefetch_preload            = "off"
  response_buffering          = "off"
  script_load_optimization    = "off"
  server_side_exclude         = "off"
  tls_client_auth             = "off"
  true_client_ip_header       = "off"
  websockets                  = "off"
  challenge_ttl               = 31536000
  max_upload                  = 300
  cipher                      = ["AES128-SHA256"]
  minify {
    css  = "off"
    js   = "off"
    html = "off"
  }
  security_header {
    enabled            = false
    include_subdomains = false
    max_age            = 0
    nosniff            = false
  }
  mobile_redirect {
    status           = "on"
    mobile_subdomain = "m.domain.com"
    strip_uri        = true
  }
}

resource "ibm_cis_domain_settings" "test" {
  cis_id          = ibm_cis.instance.id
  domain_id       = ibm_cis_domain.example.id
  waf             = "on"
  ssl             = "full"
  min_tls_version = "1.2"
}
```

## Argument reference
Review the argument references that you can specify for your resource. 

- `always_use_https` - (Optional, String) Supported values are `off` and `on`.
- `automatic_https_rewrites` - (Optional, String) Enable HTTPS rewrites. Allowed values are `off` and `on`.
- `browser_check` - (Optional, String) Enable a client browser check to look for common HTTP headers that are used by malicious users. If HTTP headers are found,  access to your website is blocked. Supported values are `off` and `on`.
- `brotli` - (Optional, String) Supported values are `off` and `on`.
- `challenge_ttl` - (Optional, String) Challenge TTL values are `300`, `900`, `1800`, `2700`, `3600`, `7200`, `10800`, `14400`, `28800`, `57600`, `86400`, `604800`, `2592000`, and `31536000`.
- `cipher` - (Optional, String) Cipher setting values are  `ECDHE-ECDSA-AES128-GCM-SHA256`, `ECDHE-ECDSA-CHACHA20-POLY1305`,`ECDHE-RSA-AES128-GCM-SHA256`, `ECDHE-RSA-CHACHA20-POLY1305`, `ECDHE-ECDSA-AES128-SHA256`, `ECDHE-ECDSA-AES128-SHA`, `ECDHE-RSA-AES128-SHA256`, `ECDHE-RSA-AES128-SHA`, `AES128-GCM-SHA256`, `AES128-SHA256`, `AES128-SHA`, `ECDHE-ECDSA-AES256-GCM-SHA384`, `ECDHE-ECDSA-AES256-SHA384`, `ECDHE-RSA-AES256-GCM-SHA384`, `ECDHE-RSA-AES256-SHA384`, `ECDHE-RSA-AES256-SHA`, `AES256-GCM-SHA384`, `AES256-SHA256`, `AES256-SHA`, `DES-CBC3-SHA`, `AEAD-AES128-GCM-SHA256`, `AEAD-AES256-GCM-SHA384`, `AEAD-CHACHA20-POLY1305-SHA256`.
- `cis_id` - (Required, String) The ID of the IBM Cloud Internet Services instance.
- `cname_flattening` - (Optional, String) Supported values are `flatten_at_root`, `flatten_all`, and `flatten_none`.
- `domain_id` - (Required, String) The ID of the domain that you want to customize.
- `dnssec` - (Optional, String) Can set to `active` only once. Allowed values are `active`, `disabled`.
- `hotlink_protection` - (Optional, String) Supported values are `off` and `on`.
- `http2` - (Optional, String) Supported values are `off` and `on`.
- `image_load_optimization` - (Optional, String) Supported values are `off` and `on`.
- `image_size_optimization` - (Optional, String) Supported values are `lossless`,  `off`, and `lossy`.
- `ipv6` - (Optional, String) Supported values are `off` and `on`.
- `ip_geolocation` - (Optional, String) Supported values are `off` and `on`.
- `max_upload` - (Optional, String) Maximum upload values are `100`, `125`, `150`, `175`, `200`, `225`, `250`, `275`, `300`, `325`, `350`, `375`, `400`, `425`, `450`, `475`, and `500`.
- `min_tls_version` - (Optional, String) The minimum TLS version that you want to allow. Allowed values are `1.1`, `1.2`, or `1.3`.
- `minify`  (Optional, List) Minify the setting as stated.

  Nested scheme for `minify`:
  - `css` - (Required, String) CSS supported values are `on` and `off`.
  - `html` - (Required, String) HTML supported values are `on` and `off`.
  - `js` - (Required, String) JS supported values are `on` and `off`.
- `mobile_redirect`  (Optional, List) Mobile redirect setting.

  Nested scheme for `mobile_redirect`:
  - `mobile_subdomain` - (Optional, String) Mobile redirect subdomain. For example `m.domain.com`.
  - `status`- (Bool) Required-Mobile redirect setting status values are **true** and **false**.
  - `strip_uri` - (Optional, Bool) Strip URI for mobile redirect.
- `origin_error_page_pass_thru` - (Optional, String) Supported values are `off` and `on`.
- `opportunistic_encryption` - (Optional, String) Supported values are `off` and `on`.
- `pseudo_ipv4` - (Optional, String) Supported values are `overwrite_header`, `off`, and `add_header`.
- `prefetch_preload` - (Optional, String) Supported values are `off` and `on`.
- `response_buffering` - (Optional, String) Supported values are `off` and `on`.
- `script_load_optimization` - (Optional, String) Supported values are `off` and `on`.
- `server_side_exclude` - (Optional, String) Supported values are `off` and `on`.
- `security_header`  (Optional, List) Security headers as stated.
- `security_header.enabled`- (Bool) Required-Supported values are **true** and **false**.
- `security_header.include_subdomains`- (Bool) Required-Supported values are **true** and **false**.
- `security_header.max_age`- (Required, Integer) Maximum age of the security header.
- `security_header.nosniff`- (Bool) Required-No sniff.
- `ssl` - (Optional, String) Allowed values: `off`, `flexible`, `full`, `strict`, `origin_pull`.
- `tls_client_auth` - (Optional, String) Supported values are `off` and `on`.
- `true_client_ip_header` - (Optional, String) Supported values are `off` and `on`.
- `waf` - (Optional, String) Enable a web application firewall (WAF). Supported values are `off` and `on`.
- `websockets` - (Optional, String) Supported values are `off` and `on`.

**Note**

Extra settings are not implemented in this version of the provider.
 
## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `certificate_status` - (String)  The value is displayed as `none`, `initializing`, `authorizing`, or `active`.
