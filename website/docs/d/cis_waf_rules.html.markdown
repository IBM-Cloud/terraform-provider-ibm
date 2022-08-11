---
subcategory: "Internet services"
layout: "ibm"
page_title: "IBM: ibm_cis_waf_rules"
description: |-
  Get information of IBM Cloud Internet Services WAF rules resource.
---

# ibm_cis_waf_rule
Retrieve information about an existing IBM Cloud Internet Services WAF rules resource. For more information, see [CIS rule sets](https://cloud.ibm.com/docs/cis?topic=cis-waf-settings#cis-ruleset-for-waf).

## Example usage

```terraform
data "ibm_cis_waf_rules" "rules" {
		cis_id    = data.ibm_cis.cis.id
		domain_id = data.ibm_cis_domain.cis_domain.id
		package_id = "1e334934fd7ae32ad705667f8c1057aa"
}
```

## Argument reference
Review the argument references that you can specify for your data source. 

- `cis_id` - (Required, String) The ID of the IBM Cloud Internet Services service instance.
- `domain_id` - (Required, String) The ID of the domain to add the rate limit rule.
- `package_id` - (Required, String) The ID of WAF rule package.

## Attribute reference
In addition to all argument reference list, you can access the following attribute references after your data source is created. 

- `waf_rules` (List) The list of WAF rules.

  Nested scheme for `waf_rules`:
  - `allowed_modes` - (String) The allowed modes for setting the WAF rule mode.
  - `description` - (String) The WAF rule description.
  - `group` - (String) The WAF rule group.

    Nested scheme for `group`:
	- `id` - (String) The WAF rule group ID.
	- `name` - (String) The name of the WAF rule group.
  - `id` - (String)  It is a combination of `<rule_id>,<package_id>,<domain_id>,<cis_id>` attributes concatenated with `:` character.
  - `mode` - (String) The mode setting that can be set only once. Valid values are `on`, `off`, `default`, `disable`, `simulate`, `block`, `challenge`.
  - `package_id` - (String) The ID of WAF rule package.
  - `priority` - (String) The WAF rule priority.
  - `rule_id` - (String) The ID of WAF rule package.