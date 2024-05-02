---
subcategory: "Internet services"
layout: "ibm"
page_title: "IBM: ibm_cis_rulesets"
description: |-
  Get information on an IBM Cloud Internet Services ruleset.
---

# ibm_cis_rulesets

Retrieve information about IBM Cloud Internet Services Instance/Zone rulesets data sources. For more information, see [IBM Cloud Internet Services].

## Example usage

```terraform
data "ibm_cis_rulesets" "tests" {
    cis_id    = ibm_cis.instance.id
    domain_id = data.ibm_cis_domain.cis_domain.domain_id
    }
```

## Argument reference
Review the argument references that you can specify for your data source.

- `cis_id` - (Required, String) The ID of the CIS service instance.
- `domain_id` - (Optional, String) The Domain/Zone ID of the CIS service instance. If domain_id is provided the request will be made at the zone/domain level, otherwise the request will be made at the instance level.  

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

