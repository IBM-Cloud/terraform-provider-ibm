---
subcategory: "Internet services"
layout: "ibm"
page_title: "IBM : ibm_cis_waf_packages"
description: |-
  Get information on an IBM Cloud Internet Services WAF packages.
---

# ibm_cis_waf_packages
Retrieve information about an existing IBM Cloud Internet Services WAF package resource. For more information, about WAF refer to [CIS rule sets](https://cloud.ibm.com/docs/cis?topic=cis-waf-settings).

## Example usage

```terraform
data "ibm_cis_rate_limit" "ratelimit" {
    cis_id = data.ibm_cis.cis.id
    domain_id = data.ibm_cis_domain.cis_domain.domain_id
}
```

## Argument reference
Review the argument references that you can specify for your data source. 

- `cis_id` - (Required, String) The ID of the IBM Cloud Internet Services service instance.
- `domain_id` - (Required, String) The ID of the domain.

## Attribute reference
In addition to all argument reference list, you can access the following attribute references after your data source is created. 

- `description` - (String)   The WAF package description.
- `detection_mode` - (String) The WAF package detection mode.
- `id` - (String) The ID of resource. It is the combination of `<package_id>:<domain_id>:<cis_id>`.
- `package_id` - (String) The WAF package ID.
