---

subcategory: "Internet services"
layout: "ibm"
page_title: "IBM: ibm_cis_waf_rule"
description: |-
  Provides a IBM CIS WAF Rule Settings resource.
---

# ibm_cis_waf_rule
Create, update, or delete an IBM Cloud Internet Services WAF rule settings resource. This resource is associated with an IBM Cloud Internet Services instance and a CIS Domain resource. It allows to change WAF rule settings of a domain of a CIS instance. For more information, refer to [IBM Cloud Internet Services rule sets](https://cloud.ibm.com/docs/cis?topic=cis-waf-settings#cis-ruleset-for-waf).

## Example usage
The following example shows how you can add a WAF rule resource to an IBM Cloud Internet Services domain.

```terraform
resource "ibm_cis_waf_rule" "test" {
	cis_id     = data.ibm_cis.cis.id
	domain_id  = data.ibm_cis_domain.cis_domain.id
	package_id = "c504870194831cd12c3fc0284f294abb"
	rule_id    = "100000356"
	mode       = "on"
}
```

## Argument reference
Review the argument references that you can specify for your resource. 

- `cis_id` - (Required, String) The ID of the IBM Cloud Internet Services instance.
- `domain_id` - (Required, String) The ID of the domain where you want to change TLS settings.
- `package_id` - (Required, String) The WAF rule package ID. This cannot be modified.
- `rule_id` - (Required, String) The WAF rule ID. The filed cannot be modified.
- `mode` - (Required, String) The mode to use when the rule is triggered. Value is restricted based on the allowed_modes of the rule. Valid values are `on`, `off`, `default`, `disable`, `simulate`, `block`, `challenge`.

## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `allowed_modes` - (String) The allowed modes for setting the WAF rule mode.
- `description` - (String) The WAF rule description.
- `group` - (String) The WAF rule group.
 
  Nested scheme for `group`: 
	- `id` - (String) The WAF rule group ID.
	- `name` - (String) The name of the WAF rule group.
- `id` - (String) The WAF package ID. It is a combination of `<rule_id>,<package_id>,<domain_id>,<cis_id>` attributes concatenated with `:`.
- `priority` - (String) The WAF rule priority.

## Import
The `ibm_cis_waf_rule` resource can be imported by using the ID. The ID is formed from the rule_id, `<package_id>, <domain ID>, <package ID>` of the domain and the CRN (Cloud Resource Name)  Concatenated  by using a `:` character.

The domain ID and CRN will be located on the **Overview** page of the Internet Services instance of the domain heading of the console, or by using the `ibmcloud cis` command line commands.

- **Rule ID** is a digit character string of the form: `100000356`

- **Package ID** is a 32 digit character string of the form: `c504870194831cd12c3fc0284f294abb`

- **Domain ID** is a 32 digit character string of the form: `9caf68812ae9b3f0377fdf986751a78f`

- **CRN** is a 120 digit character string of the form: `crn:v1:bluemix:public:internet-svcs:global:a/4ea1882a2d3401ed1e459979941966ea:31fa970d-51d0-4b05-893e-251cba75a7b3::`

**Syntax**

```
$ terraform import ibm_cis_waf_rule.waf_rule <rule_id>:<package_id>:<domain-id>:<crn>

```

**Example**

```
$ terraform import ibm_cis_waf_rule.waf_rule 100000356:c504870194831cd12c3fc0284f294abb:9caf68812ae9b3f0377fdf986751a78f:crn:v1:bluemix:public:internet-svcs:global:a/4ea1882a2d3401ed1e459979941966ea:31fa970d-51d0-4b05-893e-251cba75a7b3::
```
