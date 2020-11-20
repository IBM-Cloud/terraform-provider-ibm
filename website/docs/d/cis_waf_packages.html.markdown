---
layout: "ibm"
page_title: "IBM : ibm_cis_waf_packages"
sidebar_current: "docs-ibm-datasource-cis-waf-packages"
description: |-
  Get information on an IBM Cloud Internet Services WAF Packages.
---

# ibm_cis_waf_packages

Imports a read only copy of an existing Internet Services WAF Package resource.

## Example Usage

```hcl
data "ibm_cis_rate_limit" "ratelimit" {
    cis_id = data.ibm_cis.cis.id
    domain_id = data.ibm_cis_domain.cis_domain.domain_id
}
```

## Argument Reference

The following arguments are supported:

- `cis_id` - (Required,string) The ID of the CIS service instance
- `domain_id` - (Required,string) The ID of the domain.

## Attribute Reference

The following attributes are exported:

- `id` - The ID of resource. Id id the combination of <package_id>:<domain_id>:<cis_id>.
- `description` - The WAF Package description.
- `package_id` - The WAF package ID.
- `detection_mode` - Thw WAF Package detection mode.
