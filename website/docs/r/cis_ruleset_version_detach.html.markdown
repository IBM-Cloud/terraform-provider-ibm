---

subcategory: "Internet services"
layout: "ibm"
page_title: "IBM: ibm_cis_ruleset_version_detach"
description: |-
  Provides an IBM CIS ruleset version resource.
---

# ibm_cis_ruleset_version_detach
Provides an IBM Cloud Internet Services ruleset version resource of an instance or domain to be detached. This allow ruleset version to delete. For more information about IBM Cloud Internet Services ruleset version detach, see [ruleset instance](https://cloud.ibm.com/docs/cis?topic=cis-managed-rules-overview).

## Example usage

```terraform
# delete ruleset of a domain or instance

resource "ibm_cis_ruleset_version_detach" "tests" {
    cis_id    = ibm_cis.instance.id
    domain_id = data.ibm_cis_domain.cis_domain.domain_id
    ruleset_id = "<id of the ruleset>"
    version = "<ruleset version>"
    }
```

## Argument reference
Review the argument references that you can specify for your resource. 

- `cis_id` - (Required, String) The ID of the CIS service instance.
- `domain_id` - (Optional, String) The Domain/Zone ID of the CIS service instance. If domain_id is provided the request will be made at the zone/domain level otherwise the request will be made at the instance level.
- `ruleset_id` - (Required, String) ID of the ruleset.
- `version` - (Required, String) Version of the ruleset to be deleted. You can not delete the latest version of the ruleset.

## Attribute reference

This resource does not provide attribute reference.

## Import

Import is not possible, as there is no way to read the resource module.