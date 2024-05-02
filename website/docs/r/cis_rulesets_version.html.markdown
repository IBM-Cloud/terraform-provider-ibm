---

subcategory: "Internet services"
layout: "ibm"
page_title: "IBM: ibm_cis_rulesets_version"
description: |-
  Provides an IBM CIS ruleset resource.
---

# ibm_cis_rulesets_version
Provides an IBM Cloud Internet Services ruleset resource to delete a ruleset of an instance or domain. For more information about IBM Cloud Internet Services ruleset, see [ruleset instance]().

## Example usage

```terraform
# delete ruleset of a domain or instance

resource "ibm_cis_rulesets_version" "tests" {
    cis_id    = ibm_cis.instance.id
    domain_id = data.ibm_cis_domain.cis_domain.domain_id
    ruleset_id = "<id of the ruleset>"
    ruleset_version = "<ruleset version>"
    }
```

## Argument reference
Review the argument references that you can specify for your resource. 

- `cis_id` - (Required, String) The ID of the CIS service instance.
- `domain_id` - (Optional, String) The Domain/Zone ID of the CIS service instance. If domain_id is provided the request will be made at the zone/domain level otherwise the request will be made at the instance level.
- `ruleset_id` - (Required, String) ID of the ruleset.
- `ruleset_version` - (Required, String) Version of the ruleset to be deleted.

## Import
The `ibm_cis_rulesets_version` resource is imported by using the ID. The ID is formed from the ruleset version, the ruleset ID, the domain ID of the domain and the Cloud Resource Name (CRN) concatenated using a `:` character.

The domain ID and CRN are located on the **Overview** page of the Internet Services instance of the domain heading of the console, or by using the `ibm cis` CLI commands.

- **Ruleset Version** is a string of the form: `10`.

- **Ruleset ID** is a 32-digit character string of the form: `489d96f0da6ed76251b475971b097205c`.

- **Domain ID** is a 32-digit character string of the form: `9caf68812ae9b3f0377fdf986751a78f`.

- **CRN** is a 120-digit character string of the form: `crn:v1:bluemix:public:internet-svcs:global:a/4ea1882a2d3401ed1e459979941966ea:31fa970d-51d0-4b05-893e-251cba75a7b3::`.

**Syntax**

```
$ terraform import ibm_cis_rulesets_version.config <ruleset version>:<ruleset_id>:<domain-id>:<crn>
```

**Example**

```
$ terraform import ibm_cis_rulesets_version.config 10:48996f0da6ed76251b475971b097205c:9caf68812ae9b3f0377fdf986751a78f:crn:v1:bluemix:public:internet-svcs:global:a/4ea1882a2d3401ed1e459979941966ea:31fa970d-51d0-4b05-893e-251cba75a7b3::
```

