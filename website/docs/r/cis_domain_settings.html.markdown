---
layout: "ibm"
page_title: "IBM: ibm_cis_domain_settings"
sidebar_current: "docs-ibm-cis-domain-settings"
description: |-
  Provides a resource which customizes IBM Cloud Internet Services domain settings.
---

# ibm_zone_settings

Provides a resource which customizes IBM Cloud Internet Services domain settings. 

## Example Usage

```hcl
resource "ibm_zone_settings_override" "test" {
    cis_id = "${ibm_cis.instance.id}"  
    domain_id = "${ibm_cis_domain.example.id}"
    name = "${var.ibm_zone}"
    waf = true
    min_tls_version = "1.3"
    }
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The DNS zone to which apply settings.
* `waf`. Allowed values: "off", "on"
* `min_tls_version`. Allowed values: 1.1", "1.2", "1.3".
* `ssl`. Allowed values: "off", "flexible", "full", "strict", "origin_pull".
* `automatic_https_rewrites`. Allowed values: "off", "on"
* `opportunistic_encryption`. Allowed values: "off", "on"
* `cname_flattening`. Allowed values: "flatten_at_root", "flatten_all", "flatten_none".


Additional settings not implemented in this version of the provider. 


## Attributes Reference

The following attributes are exported:
* `certificate_status`. Value of: "none", "initializing", "authorizing", "active"
