---
subcategory: "Internet services"
layout: "ibm"
page_title: "IBM: ibm_cis_ai_security_settings"
description: |-
  Provides an IBM Cloud Internet Services AI Security for Apps settings resource.
---

# ibm_cis_ai_security_settings

Create, update, or delete the AI Security for Apps enabled setting on a zone of an IBM Cloud Internet Services (CIS) instance. For more information, see [IBM Cloud Internet Services](https://cloud.ibm.com/docs/cis).

## Example usage

```terraform
resource "ibm_cis_ai_security_settings" "example" {
  cis_id    = data.ibm_cis.cis.id
  domain_id = data.ibm_cis_domain.cis_domain.domain_id
  enabled   = true
}
```

## Argument reference

Review the argument references that you can specify for your resource.

- `cis_id` - (Required, String) The ID of the IBM Cloud Internet Services instance.
- `domain_id` - (Required, String) The ID of the domain to change AI Security settings. The ID is a combination of the zone ID and CRN, separated by a `:`.
- `enabled` - (Optional, Bool) Whether AI Security for Apps is enabled on the zone. Defaults to `false`.

## Attribute reference

In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `id` - (String) The record ID. It is a combination of `<domain_id>:<cis_id>` attributes concatenated with `:`.

## Import

The `ibm_cis_ai_security_settings` resource can be imported using the `id`. The ID is formed from the `Domain ID` and the `CRN` (Cloud Resource Name) concatenated using a `:` character.

- **Domain ID** is a 32 digit character string of the form: `9caf68812ae9b3f0377fdf986751a78f`

- **CRN** is a 120 digit character string of the form: `crn:v1:bluemix:public:internet-svcs:global:a/4ea1882a2d3401ed1e459979941966ea:31fa970d-51d0-4b05-893e-251cba75a7b3::`

**Syntax**

```
$ terraform import ibm_cis_ai_security_settings.example <domain-id>:<crn>
```

**Example**

```
$ terraform import ibm_cis_ai_security_settings.example 9caf68812ae9b3f0377fdf986751a78f:crn:v1:bluemix:public:internet-svcs:global:a/4ea1882a2d3401ed1e459979941966ea:31fa970d-51d0-4b05-893e-251cba75a7b3::
```
