---
subcategory: "Internet services"
layout: "ibm"
page_title: "IBM: ibm_cis_api_gateway_operations"
description: |-
  Provides an IBM Cloud Internet Services API Gateway Operations data source.
---

# ibm_cis_api_gateway_operations

Retrieve a list of API Gateway operations saved to Endpoint Management for a zone in IBM Cloud Internet Services (CIS). For more information, see [IBM Cloud Internet Services](https://cloud.ibm.com/docs/cis).

## Example usage

```terraform
data "ibm_cis_api_gateway_operations" "example" {
  cis_id    = data.ibm_cis.cis.id
  domain_id = data.ibm_cis_domain.cis_domain.domain_id
  order     = "method"
  direction = "asc"
}
```

## Argument reference

Review the argument references that you can specify for your data source.

- `cis_id` - (Required, String) The ID of the IBM Cloud Internet Services instance.
- `domain_id` - (Required, String) The ID of the domain. The ID is a combination of the zone ID and CRN, separated by a `:`.
- `host` - (Optional, List of String) Filter results to only include the specified hosts.
- `method` - (Optional, List of String) Filter results to only include the specified HTTP methods.
- `endpoint` - (Optional, String) Filter results to only include endpoints containing this pattern.
- `user_labels` - (Optional, List of String) Filter results to only include operations with these labels applied.
- `direction` - (Optional, String) Direction to order results (`asc` or `desc`).
- `order` - (Optional, String) Field to order results by (`endpoint`, `host`, `last_updated`, `method`).
- `page` - (Optional, Int) Page number of paginated results.
- `per_page` - (Optional, Int) Maximum number of results per page.

## Attribute reference

In addition to all argument reference list, you can access the following attribute reference after your data source is read.

- `id` - (String) The record ID. It is a combination of `<domain_id>:<cis_id>` attributes concatenated with `:`.
- `total_count` - (Int) Total count of operations matching the query.
- `operations` - (List) A list of API Gateway operations.
  - `operation_id` - (String) UUID of the operation.
  - `method` - (String) HTTP method.
  - `host` - (String) API host.
  - `endpoint` - (String) Endpoint path.
  - `last_updated` - (String) Timestamp of the last update.
