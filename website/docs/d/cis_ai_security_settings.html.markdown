---
subcategory: "Internet services"
layout: "ibm"
page_title: "IBM: ibm_cis_ai_security_settings"
description: |-
  Retrieve IBM Cloud Internet Services AI Security for Apps settings.
---

# ibm_cis_ai_security_settings

Retrieve the AI Security for Apps enabled setting for a zone of an IBM Cloud Internet Services (CIS) instance. For more information, see [IBM Cloud Internet Services](https://cloud.ibm.com/docs/cis).

## Example usage

```terraform
data "ibm_cis_ai_security_settings" "example" {
  cis_id    = data.ibm_cis.cis.id
  domain_id = data.ibm_cis_domain.cis_domain.domain_id
}
```

## Argument reference

Review the argument references that you can specify for your data source.

- `cis_id` - (Required, String) The ID of the IBM Cloud Internet Services instance.
- `domain_id` - (Required, String) The ID of the domain. The ID is a combination of the zone ID and CRN, separated by a `:`.

## Attributes reference

In addition to all argument reference list, you can access the following attribute references after your data source is created.

- `id` - (String) The record ID. It is a combination of `<domain_id>:<cis_id>` attributes concatenated with `:`.
- `enabled` - (Bool) Whether AI Security for Apps is enabled on the zone.
