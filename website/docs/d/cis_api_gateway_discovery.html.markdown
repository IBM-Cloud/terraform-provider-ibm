---
subcategory: "Internet services"
layout: "ibm"
page_title: "IBM: ibm_cis_api_gateway_discovery"
description: |-
  Retrieve IBM Cloud Internet Services API Gateway discovery document.
---

# ibm_cis_api_gateway_discovery

Retrieve discovered API Gateway operations rendered as an OpenAPI schema document for a zone of an IBM Cloud Internet Services (CIS) instance. For more information, see [IBM Cloud Internet Services](https://cloud.ibm.com/docs/cis).

## Example usage

```terraform
data "ibm_cis_api_gateway_discovery" "example" {
  cis_id    = data.ibm_cis.cis.id
  domain_id = data.ibm_cis_domain.cis_domain.domain_id
}

output "discovery_doc" {
  value = data.ibm_cis_api_gateway_discovery.example.result
}
```

## Argument reference

Review the argument references that you can specify for your data source.

- `cis_id` - (Required, String) The ID of the IBM Cloud Internet Services instance.
- `domain_id` - (Required, String) The ID of the domain. The ID is a combination of the zone ID and CRN, separated by a `:`.

## Attributes reference

In addition to all argument reference list, you can access the following attribute references after your data source is created.

- `id` - (String) The record ID. It is a combination of `<domain_id>:<cis_id>` attributes concatenated with `:`.
- `result` - (String) Discovered operations rendered as a JSON-encoded OpenAPI schema document.
