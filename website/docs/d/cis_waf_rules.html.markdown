---
subcategory: "Internet services"
layout: "ibm"
page_title: "IBM: ibm_cis_waf_rules"
description: |-
  Get information of IBM Cloud Internet Services WAF Rules resource.
---

# ibm_cis_waf_rule

Imports a read only copy of an existing Internet Services WAF Rules resource.

## Example Usage

```hcl
data "ibm_cis_waf_rules" "rules" {
		cis_id    = data.ibm_cis.cis.id
		domain_id = data.ibm_cis_domain.cis_domain.id
		package_id = "1e334934fd7ae32ad705667f8c1057aa"
}
```

## Argument Reference

The following arguments are supported:

- `cis_id` - (Required,string) The ID of the CIS service instance
- `domain_id` - (Required,string) The ID of the domain to add the Rate Limit rule.
- `package_id` - (Required,string) The ID of WAF rule package.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

- `waf_rules` - The list of waf rules.
  - `id` - It is a combination of <`rule_id`>,<`package_id`>,<`domain_id`>,<`cis_id`> attributes concatenated with ":".
  - `rule_id` - The ID of waf rule.
  - `package_id` - The ID of waf rule package.
  - `mode` - The mode setting. This field only once can be set. Valid values: `on`, `off`, `default`, `disable`, `simulate`, `block`, `challenge`.
  - `description` - The WAF rule description.
  - `priority` - The WAF Rule priority.
  - `group` - The waf rule group.
    - `id` - The waf rule group id.
    - `name` - The name of waf rule group.
  - `allowed_modes` - The allowed modes for setting the waf rule mode.
