---
subcategory: "Internet services"
layout: "ibm"
page_title: "IBM: ibm_cis_api_gateway_operation"
description: |-
  Retrieve an IBM Cloud Internet Services API Gateway Operation.
---

# ibm_cis_api_gateway_operation

Retrieve details of a specific API Gateway operation saved in IBM Cloud Internet Services (CIS) API Shield Endpoint Management. For more information, see [IBM Cloud Internet Services](https://cloud.ibm.com/docs/cis).

## Example usage

```terraform
data "ibm_cis_api_gateway_operation" "example" {
  cis_id       = data.ibm_cis.cis.id
  domain_id    = data.ibm_cis_domain.cis_domain.domain_id
  operation_id = "f174e90a-fafe-4643-bbbc-4a0ed4fc8415"
}
```

## Argument reference

Review the argument references that you can specify for your data source.

- `cis_id` - (Required, String) The ID of the IBM Cloud Internet Services instance.
- `domain_id` - (Required, String) The ID of the domain. The ID is a combination of the zone ID and CRN, separated by a `:`.
- `operation_id` - (Required, String) The UUID of the API Gateway operation to retrieve.

## Attributes reference

In addition to all argument reference list, you can access the following attribute references after your data source is created.

- `id` - (String) The record ID. It is a combination of `<operation_id>:<domain_id>:<cis_id>` attributes concatenated with `:`.
- `endpoint` - (String) The API endpoint path, for example `/api/v1/chat`.
- `host` - (String) The API hostname, for example `api.example.com`.
- `method` - (String) The HTTP method, for example `POST`.
