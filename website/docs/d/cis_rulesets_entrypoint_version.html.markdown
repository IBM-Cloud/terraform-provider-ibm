---
subcategory: "Internet services"
layout: "ibm"
page_title: "IBM: ibm_cis_rulesets_entrypoint_versions"
description: |-
  Get information on an IBM Cloud Internet Services ruleset version.
---

# ibm_cis_rulesets_entrypoint_versions

Retrieve information about an IBM Cloud Internet Services Instance/Zone Entry Point ruleset's versions data sources. For more information, see [IBM Cloud Internet Services].

## Example usage

```terraform
data "ibm_cis_rulesets_entrypoint_versions" "tests" {
    cis_id    = ibm_cis.instance.id
    domain_id = data.ibm_cis_domain.cis_domain.domain_id
    ruleset_phase = data.ibm_cis_ruleset.cis_ruleset.ruleset_phase
    }
```

## Argument reference
Review the argument references that you can specify for your data source.

- `cis_id` - (Required, String) The ID of the CIS service instance.
- `domain_id` - (Optional, String) The Domain/Zone ID of the CIS service instance. If domain_id is provided the request will be made at the zone/domain level otherwise the request will be made at the instance level.  
- `ruleset_phase` - (Required, String) The phase of the ruleset.

## Attributes reference
In addition to the argument reference list, you can access the following attribute references after your data source is created.

- `result` - (list)
    - `id` - (string) Ruleset ID.
    - `description` - (string) Description of the ruleset.
    - `kind` - (string) The kind of the ruleset.
    - `Phase` - (string) Phase of the ruleset.
    - `name` - (string) Name of the ruleset.
    - `last updated` - (string) Last update date of the ruleset.
    - `version` - (string) Version of the ruleset.

