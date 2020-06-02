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
resource "ibm_cis_domain_settings" "test" {
  cis_id          = ibm_cis.instance.id
  domain_id       = ibm_cis_domain.example.id
  waf             = "on"
  ssl             = "full"
  min_tls_version = "1.2"
}
```

## Argument Reference

The following arguments are supported:

* `cis_id` - (Required) The ID of the CIS service instance.
* `domain_id` - (Required) The ID of the domain.
* `waf`. (Optional, string) Allowed values: "off", "on"
* `min_tls_version`. (Optional, string) Allowed values: 1.1", "1.2", "1.3".
* `ssl`. (Optional, string) Allowed values: "off", "flexible", "full", "strict", "origin_pull".
* `automatic_https_rewrites`. (Optional, string) Allowed values: "off", "on"
* `opportunistic_encryption`. (Optional, string) Allowed values: "off", "on"
* `cname_flattening`. (Optional, string) Allowed values: "flatten_at_root", "flatten_all", "flatten_none".
* `always_use_https` . (Optional, string) Allowed values: "off", "on"
* `ipv6` . (Optional, string) Allowed values: "off", "on"
* `browser_check` . (Optional, string) Allowed values: "off", "on"
* `hotlink_protection` . (Optional, string) Allowed values: "off", "on"
* `http2` . (Optional, string) Allowed values: "off", "on"
* `image_load_optimization` . (Optional, string) Allowed values: "off", "on"
* `image_size_optimization` . (Optional, string) Allowed values: "lossless", "off", "lossy"
* `ip_geolocation` . (Optional, string) Allowed values: "off", "on"
* `origin_error_page_pass_thru` . (Optional, string) Allowed values: "off", "on"
* `brotli` . (Optional, string) Allowed values: "off", "on"
* `pseudo_ipv4` . (Optional, string) Allowed values: "overwrite_header", "off", "add_header"
* `prefetch_preload` . (Optional, string) Allowed values: "off", "on"
* `response_buffering` . (Optional, string) Allowed values: "off", "on"
* `script_load_optimization` . (Optional, string) Allowed values: "off", "on"
* `server_side_exclude` . (Optional, string) Allowed values: "off", "on"
* `tls_client_auth` . (Optional, string) Allowed values: "off", "on"
* `true_client_ip_header` . (Optional, string) Allowed values: "off", "on"
* `websockets` . (Optional, string) Allowed values: "off", "on"

Additional settings not implemented in this version of the provider. 


## Attributes Reference

The following attributes are exported: 

* `certificate_status`. Value of: "none", "initializing", "authorizing", "active"
