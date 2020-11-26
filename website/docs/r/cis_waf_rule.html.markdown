---
layout: "ibm"
page_title: "IBM: ibm_cis_waf_rule"
sidebar_current: "docs-ibm-resource-cis-waf-rule"
description: |-
  Provides a IBM CIS WAF Rule Settings resource.
---

# ibm_cis_waf_rule

Provides a IBM CIS WAF Rule Settings resource. This resource is associated with an IBM Cloud Internet Services instance and a CIS Domain resource. It allows to change WAF Rule settings of a domain of a CIS instance.

## Example Usage

```hcl
resource "ibm_cis_waf_rule" "test" {
	cis_id     = data.ibm_cis.cis.id
	domain_id  = data.ibm_cis_domain.cis_domain.id
	package_id = "c504870194831cd12c3fc0284f294abb"
	rule_id    = "100000356"
	mode       = "on"
}
```

## Argument Reference

The following arguments are supported:

- `cis_id` - (Required,string) The ID of the CIS service instance.
- `domain_id` - (Required,string) The ID of the domain to change TLS settings.
- `package_id` - (Required,string) The ID of waf rule package. This field can not be modified.
- `rule_id` - (Required,string) The ID of waf rule. This field can not be modified.
- `mode` - (Required,string) The mode to use when the rule is triggered. Value is restricted based on the allowed_modes of the rule. Valid values: `on`, `off`, `default`, `disable`, `simulate`, `block`, `challenge`.

## Attributes Reference

The following attributes are exported:

- `id` - It is a combination of <`rule_id`>,<`package_id`>,<`domain_id`>,<`cis_id`> attributes concatenated with ":".
- `description` - The WAF rule description.
- `priority` - The WAF Rule priority.
- `group` - The waf rule group.
  - `id` - The waf rule group id.
  - `name` - The name of waf rule group.
- `allowed_modes` - The allowed modes for setting the waf rule mode.

## Import

The `ibm_cis_waf_rule` resource can be imported using the `id`. The ID is formed from the `rule_id`, `package_id`, `Domain ID` of the domain and the `CRN` (Cloud Resource Name) concatentated using a `:` character.

The Domain ID and CRN will be located on the **Overview** page of the Internet Services instance under the **Domain** heading of the UI, or via using the `ibmcloud cis` CLI commands.

- **Rule ID** is a digit character string of the form: `100000356`

- **Package ID** is a 32 digit character string of the form: `c504870194831cd12c3fc0284f294abb`

- **Domain ID** is a 32 digit character string of the form: `9caf68812ae9b3f0377fdf986751a78f`

- **CRN** is a 120 digit character string of the form: `crn:v1:bluemix:public:internet-svcs:global:a/4ea1882a2d3401ed1e459979941966ea:31fa970d-51d0-4b05-893e-251cba75a7b3::`

```
$ terraform import ibm_cis_waf_rule.waf_rule <rule_id>:<package_id>:<domain-id>:<crn>

$ terraform import ibm_cis_waf_rule.waf_rule 100000356:c504870194831cd12c3fc0284f294abb:9caf68812ae9b3f0377fdf986751a78f:crn:v1:bluemix:public:internet-svcs:global:a/4ea1882a2d3401ed1e459979941966ea:31fa970d-51d0-4b05-893e-251cba75a7b3::
```
