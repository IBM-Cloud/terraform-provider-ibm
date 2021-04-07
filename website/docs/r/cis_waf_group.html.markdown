---

subcategory: "Internet services"
layout: "ibm"
page_title: "IBM: cis_waf_group"
description: |-
  Provides an IBM Cloud Internet Services WAF Rule Group resource.
---

# ibm_cis_waf_group

Provides a IBM CIS WAF Rule Group resource. This resource is associated with an IBM Cloud Internet Services instance and a CIS Domain resource. It allows to change WAF Groups mode of a domain of a CIS instance. This is also named as CIS rule set. Please find OWASP rule set set tab under WAF of your instance in UI.

## Example Usage

```hcl
resource "ibm_cis_waf_group" "test" {
  cis_id     = data.ibm_cis.cis.id
  domain_id  = data.ibm_cis_domain.cis_domain.domain_id
  package_id = "c504870194831cd12c3fc0284f294abb"
  group_id   = "3d8fb0c18b5a6ba7682c80e94c7937b2"
  mode       = "on"
}
```

## Argument Reference

The following arguments are supported:

- `cis_id` - (Required,string) The ID of the CIS service instance
- `domain_id` - (Required,string) The ID of the domain to change WAF Rule Group mode.
- `package_id` - (Required,string) The WAF Rule Group package ID.
- `group_id` - (Required,string) The WAF Rule Group ID.
- `mode` - (Required,string) The WAF Group mode. Valid values: `on` and `off`.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

- `id` - Unique identifier for the WAF Rule Group. Ex. <group-id>:<package-id>:<domain-id>:<crn>
- `name` - The WAF Rule Group name.
- `description` - The WAF Rule Group description.
- `rules_count` - No. of rules in WAF Group.
- `modified_rules_count` - No. of rules modified in WAF Group.

## Import

The `ibm_cis_waf_group` resource can be imported using the `id`. The ID is formed from the `WAF Rule Group ID`, the `WAF Rule Package ID`, the `Domain ID` of the domain and the `CRN` (Cloud Resource Name) concatentated usinga `:` character.

The Domain ID and CRN will be located on the **Overview** page of the Internet Services instance under the **Domain** heading of the UI, or via using the `ibmcloud cis` CLI commands.

- **Domain ID** is a 32 digit character string of the form: `9caf68812ae9b3f0377fdf986751a78f`

- **CRN** is a 120 digit character string of the form: `crn:v1:bluemix:public:internet-svcs:global:a/4ea1882a2d3401ed1e459979941966ea:31fa970d-51d0-4b05-893e-251cba75a7b3::`

- **Group ID** is a 32 digit character string of the form: `57d96f0da6ed76251b475971b097205c`. The id of an existing WAF Rule Group is not avaiable via the UI. It can be retrieved programatically via the CIS API or via the CLI using the CIS command to list the defined WAF Groups: `ibmcloud cis waf-groups <domain_id> <waf_package_id>`

```
$ terraform import ibm_cis_waf_group.myorg <group_id>:<package_id>:<domain-id>:<crn>

$ terraform import ibm_cis_domain.myorg  3d8fb0c18b5a6ba7682c80e94c7937b2:57d96f0da6ed76251b475971b097205c:9caf68812ae9b3f0377fdf986751a78f:crn:v1:bluemix:public:internet-svcs:global:a/4ea1882a2d3401ed1e459979941966ea:31fa970d-51d0-4b05-893e-251cba75a7b3::
```
