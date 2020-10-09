---
layout: "ibm"
page_title: "IBM: ibm_cis_domain_settings"
sidebar_current: "docs-ibm-cis-domain-settings"
description: |-
  Provides a resource which customizes IBM Cloud Internet Services domain settings.
---

# ibm_cis_domain_settings

Provides a resource which customizes IBM Cloud Internet Services domain settings.

## Example Usage

```hcl
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
```

## Argument Reference

The following arguments are supported:

- `cis_id` - (Required) The ID of the CIS service instance.
- `domain_id` - (Required) The ID of the domain.
- `waf`. (Optional, string) Allowed values: "off", "on"
- `min_tls_version`. (Optional, string) Allowed values: 1.1", "1.2", "1.3".
- `ssl`. (Optional, string) Allowed values: "off", "flexible", "full", "strict", "origin_pull".
- `automatic_https_rewrites`. (Optional, string) Allowed values: "off", "on"
- `opportunistic_encryption`. (Optional, string) Allowed values: "off", "on"
- `cname_flattening`. (Optional, string) Allowed values: "flatten_at_root", "flatten_all", "flatten_none".
- `always_use_https` . (Optional, string) Allowed values: "off", "on"
- `ipv6` . (Optional, string) Allowed values: "off", "on"
- `browser_check` . (Optional, string) Allowed values: "off", "on"
- `hotlink_protection` . (Optional, string) Allowed values: "off", "on"
- `http2` . (Optional, string) Allowed values: "off", "on"
- `image_load_optimization` . (Optional, string) Allowed values: "off", "on"
- `image_size_optimization` . (Optional, string) Allowed values: "lossless", "off", "lossy"
- `ip_geolocation` . (Optional, string) Allowed values: "off", "on"
- `origin_error_page_pass_thru` . (Optional, string) Allowed values: "off", "on"
- `brotli` . (Optional, string) Allowed values: "off", "on"
- `pseudo_ipv4` . (Optional, string) Allowed values: "overwrite_header", "off", "add_header"
- `prefetch_preload` . (Optional, string) Allowed values: "off", "on"
- `response_buffering` . (Optional, string) Allowed values: "off", "on"
- `script_load_optimization` . (Optional, string) Allowed values: "off", "on"
- `server_side_exclude` . (Optional, string) Allowed values: "off", "on"
- `tls_client_auth` . (Optional, string) Allowed values: "off", "on"
- `true_client_ip_header` . (Optional, string) Allowed values: "off", "on"
- `websockets` . (Optional, string) Allowed values: "off", "on"
- `challenge_ttl` . (Optional, string) Challenge TTL values: 300, 900, 1800, 2700, 3600, 7200, 10800, 14400, 28800, 57600, 86400, 604800, 2592000, 31536000
- `max_upload` . (Optional, string) Maximum upload values: 100, 125, 150, 175, 200, 225, 250, 275, 300, 325, 350, 375, 400, 425, 450, 475, 500
- `cipher` . (Optional, string) Cipher setting values: "ECDHE-ECDSA-AES128-GCM-SHA256", "ECDHE-ECDSA-CHACHA20-POLY1305","ECDHE-RSA-AES128-GCM-SHA256", "ECDHE-RSA-CHACHA20-POLY1305", "ECDHE-ECDSA-AES128-SHA256", "ECDHE-ECDSA-AES128-SHA", "ECDHE-RSA-AES128-SHA256", "ECDHE-RSA-AES128-SHA", "AES128-GCM-SHA256", "AES128-SHA256", "AES128-SHA", "ECDHE-ECDSA-AES256-GCM-SHA384", "ECDHE-ECDSA-AES256-SHA384", "ECDHE-RSA-AES256-GCM-SHA384", "ECDHE-RSA-AES256-SHA384", "ECDHE-RSA-AES256-SHA", "AES256-GCM-SHA384", "AES256-SHA256", "AES256-SHA", "DES-CBC3-SHA"
- `minify` . (Optional, list) Minify setting
  - `css` - (Required, string) CSS values: "on", "off"
  - `html` - (Required, string) HTML values: "on", "off"
  - `js` - (Required, string) JS values: "on", "off"
- `security_header` . (Optional, list) Security Headers
  - `enabled` - (Required, boolean) Enabled values : true, false
  - `include_subdomains` - (Required, boolean) Subdomain Included : true, false
  - `max_age` - (Required, integer) Maximum age
  - `nosniff` - (Required, boolean) No Sniff
- `mobile_redirect` . (Optional, list) Mobile Redirect Setting
  - `status` - (Required, boolean) Mobile redirect setting status values: true, false
  - `mobile_subdomain` . (Optional, string) Mobile redirect subdomain. Ex. m.domain.com
  - `strip_uri` . (Optional, boolean) Strip URI for mobile redirect.

Additional settings not implemented in this version of the provider.

## Attributes Reference

The following attributes are exported:

- `certificate_status`. (deprecated) Value of: "none", "initializing", "authorizing", "active"
