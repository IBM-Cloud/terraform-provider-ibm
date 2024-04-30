---

subcategory: "Internet services"
layout: "ibm"
page_title: "IBM: ibm_cis_rulesets_version"
description: |-
  Provides a IBM CIS ruleset resource.
---

# ibm_cis_rulesets_version
Provides an IBM Cloud Internet Services ruleset resource to delete a ruleset of an instance or domain. For more information, about IBM Cloud Internet Services ruleset, see [ruleset instance]().

## Example usage

```terraform
# delete ruleset of a domain or instance

resource "ibm_cis_rulesets_version" "tests" {
    cis_id    = ibm_cis.instance.id
    domain_id = data.ibm_cis_domain.cis_domain.domain_id
    ruleset_id = "<id of the ruleset>"
    }
```

## Argument reference
Review the argument references that you can specify for your resource. 

- `cis_id` - (Required, String) The ID of the CIS service instance.
- `domain_id` - (Optional, String) The Domain/Zone ID of the CIS service instance. If domain_id is provided request will be made at zone/domain level else request will be made at instance level.
- `ruleset_id` - (Required, String) Id of the ruleset.

        



