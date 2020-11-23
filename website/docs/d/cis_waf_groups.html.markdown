---
layout: "ibm"
page_title: "IBM: ibm_cis_waf_groups"
sidebar_current: "docs-ibm-datasource-cis-waf-groups"
description: |-
  List the WAF Rule Groups in Cloud Internet Services.
---

# ibm_cis_waf_groups

Import the details of an existing IBM Cloud Internet Service WAF Rule Groups as a read-only data source. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "ibm_cis_waf_groups" "waf_groups" {
  cis_id     = data.ibm_cis.cis.id
  domain_id  = data.ibm_cis_domain.cis_domain.id
  package_id = "c504870194831cd12c3fc0284f294abb"
}
```

## Argument Reference

- `cis_id` - (Required, string) The resource crn id of the CIS on which zones were created.
- `domain_id` - (Required, string) The ID of the domain to retrive the load balancers from.
- `package_id` - (Required, string) The WAF Rule Package ID.

## Attribute Reference

The following attributes are exported:

- `name` - The name of WAF Rule Group.
- `group_id` - The WAF group ID.
- `mode` - The `on`/`off` mode setting of WAF rule group.
- `description` - The WAF rule group description.
- `rules_count` - No. of rules in WAF Group.
- `modified_rules_count` - No. of rules modified in WAF Group.
