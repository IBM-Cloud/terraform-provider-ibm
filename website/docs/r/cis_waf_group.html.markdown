---

subcategory: "Internet services"
layout: "ibm"
page_title: "IBM: cis_waf_group"
description: |-
  Provides an IBM Cloud Internet Services WAF rule group resource.
---

# ibm_cis_waf_group
Create, update, or delete an IBM Cloud Internet Services instance and a CIS Domain resource. It allows to change WAF Groups mode of a domain of a CIS instance. It is also named as CIS rule set. Find `OWASP` rule set set tab in WAF of your instance console. For more information, refer to [IBM Cloud Internet Services rule sets](https://cloud.ibm.com/docs/cis?topic=cis-waf-settings#cis-ruleset-for-waf).

## Example usage
The following example shows how you can add a WAF group resource to an IBM Cloud Internet Services domain.

```terraform
resource "ibm_cis_waf_group" "test" {
  cis_id     = data.ibm_cis.cis.id
  domain_id  = data.ibm_cis_domain.cis_domain.domain_id
  package_id = "c504870194831cd12c3fc0284f294abb"
  group_id   = "3d8fb0c18b5a6ba7682c80e94c7937b2"
  mode       = "on"
}
```

## Argument reference
Review the argument references that you can specify for your resource.

- `cis_id` - (Required, String) The ID of the IBM Cloud Internet Services instance.
- `domain_id` - (Required, String) The ID of the domain to change WAF rule group mode.
- `package_id` - (Required, String) The WAF rule group package ID.
- `group_id` - (Required, String) The WAF rule group ID.
- `mode` - (Required, String) The WAF group mode. Valid values are `on` and `off`.

## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `description` - (String) The WAF rule group description.
- `id` - (String) The WAF rule group ID. It is a combination of `<group_id>:<package_id>:<domain-id>:<crn>` attributes concatenated with `:`.
- `modified_rules_count`-  (Integer) Number of rules modified in WAF Group.
- `name` - (String) The WAF rule group name.
- `rules_count` - (Integer)  Number of rules in WAF Group.
- `check_mode` - (Boolean) If `true`, then updating the mode with same ON>ON or OFF>OFF value, will be skipped. By default, check_mode attribute is false, and it won't check the backend value before update.


## Import
The `ibm_cis_waf_group` resource can be imported by using the ID. The ID is formed from the WAF Rule Group ID, the WAF rule package ID, the domain ID of the domain and the Cloud Resource Name (CRN) Concatenated  by using `:` character.

The domain ID and CRN will be located on the **Overview** page of the Internet Services instance of the domain heading of the console, or by using the `ibmcloud cis` command line commands.

- **Domain ID** is a 32 digit character string of the form: `9caf68812ae9b3f0377fdf986751a78f`

- **CRN** is a 120 digit character string of the form: `crn:v1:bluemix:public:internet-svcs:global:a/4ea1882a2d3401ed1e459979941966ea:31fa970d-51d0-4b05-893e-251cba75a7b3::`

- **Group ID** is a 32 digit character string of the form: `57d96f0da6ed76251b475971b097205c`. The id of an existing WAF Rule Group is not avaiable via the UI. It can be retrieved programatically via the CIS API or via the CLI using the CIS command to list the defined WAF Groups: `ibmcloud cis waf-groups <domain_id> <waf_package_id>`

**Syntax**

```
$ terraform import ibm_cis_waf_group.myorg <group_id>:<package_id>:<domain-id>:<crn>

```

**Example**

```
$ terraform import ibm_cis_domain.myorg  3d8fb0c18b5a6ba7682c80e94c7937b2:57d96f0da6ed76251b475971b097205c:9caf68812ae9b3f0377fdf986751a78f:crn:v1:bluemix:public:internet-svcs:global:a/4ea1882a2d3401ed1e459979941966ea:31fa970d-51d0-4b05-893e-251cba75a7b3::
```
