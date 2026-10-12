---
subcategory: "Internet services"
layout: "ibm"
page_title: "IBM: ibm_cis_api_gateway_schemas"
description: |-
  Retrieve IBM Cloud Internet Services API Gateway schemas.
---

# ibm_cis_api_gateway_schemas

Retrieve API Gateway schemas for a zone of an IBM Cloud Internet Services (CIS) instance. For more information, see [IBM Cloud Internet Services](https://cloud.ibm.com/docs/cis).

## Example usage

```terraform
data "ibm_cis_api_gateway_schemas" "example" {
  cis_id    = data.ibm_cis.cis.id
  domain_id = data.ibm_cis_domain.cis_domain.domain_id
}

output "schemas" {
  value = data.ibm_cis_api_gateway_schemas.example.schemas
}
```

## Argument reference

Review the argument references that you can specify for your data source.

- `cis_id` - (Required, String) The ID of the IBM Cloud Internet Services instance.
- `domain_id` - (Required, String) The ID of the domain. The ID is a combination of the zone ID and CRN, separated by a `:`.

## Attributes reference

In addition to all argument reference list, you can access the following attribute references after your data source is created.

- `id` - (String) The record ID. It is a combination of `<domain_id>:<cis_id>` attributes concatenated with `:`.
- `schemas` - (String) API Gateway schemas rendered as a JSON-encoded OpenAPI schema document.
