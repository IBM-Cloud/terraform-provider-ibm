---

subcategory: "Internet services"
layout: "ibm"
page_title: "IBM: ibm_cis_waf_package"
description: |-
  Provides a IBM CIS WAF Package resource.
---

# ibm_cis_waf_package

Provides a IBM CIS WAF Package resource. This resource is associated with an IBM Cloud Internet Services instance and a CIS Domain resource. It allows to change WAF Package settings of a domain of a CIS instance. This is also named as OWASP rule set. Please find OWASP rule set set tab under WAF of your instance in UI.

## Example Usage

```hcl
# Change WAF Package settings of the domain

resource "ibm_cis_waf_package" "waf_package" {
	cis_id      = data.ibm_cis.cis.id
	domain_id   = data.ibm_cis_domain.cis_domain.domain_id
	package_id  = "c504870194831cd12c3fc0284f294abb"
	sensitivity = "low"
	action_mode = "block"
}
```

## Argument Reference

The following arguments are supported:

- `cis_id` - (Required,string) The ID of the CIS service instance.
- `domain_id` - (Required,string) The ID of the domain to change TLS settings.
- `package_id` - (Required, string) The WAF package ID. This can not be modified.
- `sensitivity` - (Required,string) The WAF package sensitivity. Valid values are `high`, `medium`, `low`, `off`.
- `action_mode` - (Required, string) The WAF package action mode. Valid values are `simulate`, `block`, `challenge`.

## Attributes Reference

The following attributes are exported:

- `id` - The WAF package ID. It is a combination of <`package_id`>:<`domain_id`>,<`cis_id`> attributes concatenated with ":".
- `description` - The WAF Package description.
- `detection_mode` - Thw WAF Package detection mode.

## Import

The `ibm_cis_waf_package` resource can be imported using the `id`. The ID is formed from the `Package ID`, `Domain ID` of the domain and the `CRN` (Cloud Resource Name) concatentated using a `:` character.

The Domain ID and CRN will be located on the **Overview** page of the Internet Services instance under the **Domain** heading of the UI, or via using the `ibmcloud cis` CLI commands.

- **Domain ID** is a 32 digit character string of the form: `9caf68812ae9b3f0377fdf986751a78f`

- **CRN** is a 120 digit character string of the form: `crn:v1:bluemix:public:internet-svcs:global:a/4ea1882a2d3401ed1e459979941966ea:31fa970d-51d0-4b05-893e-251cba75a7b3::`

- **WAF Package ID** is a 32 digit character string of the form: `489d96f0da6ed76251b475971b097205c`.

```
$ terraform import ibm_cis_waf_package.waf_package <package-id>:<domain-id>:<crn>

$ terraform import ibm_cis_waf_package.waf_package 489d96f0da6ed76251b475971b097205c:9caf68812ae9b3f0377fdf986751a78f:crn:v1:bluemix:public:internet-svcs:global:a/4ea1882a2d3401ed1e459979941966ea:31fa970d-51d0-4b05-893e-251cba75a7b3::
```
