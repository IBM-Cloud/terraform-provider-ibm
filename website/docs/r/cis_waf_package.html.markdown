---

subcategory: "Internet services"
layout: "ibm"
page_title: "IBM: ibm_cis_waf_package"
description: |-
  Provides a IBM CIS WAF package resource.
---

# ibm_cis_waf_package
Provides an IBM Cloud Internet Services WAF package resource. This resource is associated with an IBM Cloud Internet Services instance and a CIS domain resource. It allows to change WAF package settings of a domain of an IBM Cloud Internet Services instance. It is also named as `OWASP` rule set. For more information, about WAF, see [Web Application Firewall concepts](https://cloud.ibm.com/docs/cis?topic=cis-waf-q-and-a).

## Example usage
The following example shows how you can add a WAF package resource to an IBM Cloud Internet Services domain. 

```terraform
# Change WAF Package settings of the domain

resource "ibm_cis_waf_package" "waf_package" {
	cis_id      = data.ibm_cis.cis.id
	domain_id   = data.ibm_cis_domain.cis_domain.domain_id
	package_id  = "c504870194831cd12c3fc0284f294abb"
	sensitivity = "low"
	action_mode = "block"
}
```

## Argument reference
Review the argument references that you can specify for your resource.

- `action_mode` - (Required, String) The WAF package action mode. Valid values are `simulate`, `block`, `challenge`.
- `cis_id` - (Required, String) The ID of the IBM Cloud Internet Services instance.
- `domain_id` - (Required, String) The ID of the domain where you want to change TLS settings.
- `package_id` - (Required, String) The WAF package ID. This cannot be modified.
- `sensitivity` - (Required, String) The WAF package sensitivity. Valid values are `high`, `medium`, `low`, `off`.

## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `description` - The WAF package description.
- `detection_mode` - Thw WAF package detection mode.
- `id` - (String) The WAF package ID. It is a combination of `<package_id>:<domain_id>:<cis_id>` attributes concatenated with `:`.

## Import

The `ibm_cis_waf_package` resource can be imported by using the ID. The ID is formed from the package ID, domain ID of the domain and the Cloud Resource Name (CRN) concatenated by using a : character.

The domain ID and CRN will be located on the **Overview** page of the Internet Services instance of the Domain heading of the console, or by using the `ibmcloud cis` command line commands.

- **Domain ID** is a 32 digit character string of the form: `9caf68812ae9b3f0377fdf986751a78f`

- **CRN** is a 120 digit character string of the form: `crn:v1:bluemix:public:internet-svcs:global:a/4ea1882a2d3401ed1e459979941966ea:31fa970d-51d0-4b05-893e-251cba75a7b3::`

- **WAF Package ID** is a 32 digit character string of the form: `489d96f0da6ed76251b475971b097205c`.

**Syntax**

```
$ terraform import ibm_cis_waf_package.waf_package <package-id>:<domain-id>:<crn>
```

**Example**

```
$ terraform import ibm_cis_waf_package.waf_package 489d96f0da6ed76251b475971b097205c:9caf68812ae9b3f0377fdf986751a78f:crn:v1:bluemix:public:internet-svcs:global:a/4ea1882a2d3401ed1e459979941966ea:31fa970d-51d0-4b05-893e-251cba75a7b3::
```
