---
subcategory: "Internet services"
layout: "ibm"
page_title: "IBM: ibm_cis_api_gateway_discovery_operations"
description: |-
  Retrieve IBM Cloud Internet Services API Gateway discovered operations.
---

# ibm_cis_api_gateway_discovery_operations

List discovered API Gateway operations for a zone of an IBM Cloud Internet Services (CIS) instance, with optional filtering and pagination. For more information, see [IBM Cloud Internet Services](https://cloud.ibm.com/docs/cis).

## Example usage

```terraform
data "ibm_cis_api_gateway_discovery_operations" "example" {
  cis_id    = data.ibm_cis.cis.id
  domain_id = data.ibm_cis_domain.cis_domain.domain_id
  state     = "review"
  per_page  = 25
}
```

## Argument reference

Review the argument references that you can specify for your data source.

- `cis_id` - (Required, String) The ID of the IBM Cloud Internet Services instance.
- `domain_id` - (Required, String) The ID of the domain. The ID is a combination of the zone ID and CRN, separated by a `:`.
- `diff` - (Optional, Bool) When `true`, returns only operations not yet saved into API Shield Endpoint Management.
- `direction` - (Optional, String) Direction to order results. Valid values are `asc` or `desc`.
- `endpoint` - (Optional, String) Filter results to only include endpoints containing this pattern.
- `host` - (Optional, List of String) Filter results to only include the specified hosts.
- `method` - (Optional, List of String) Filter results to only include the specified HTTP methods.
- `order` - (Optional, String) Field to order results by. Valid values are `endpoint`, `host`, `method`, `traffic_stats.last_updated`, `traffic_stats.requests`.
- `origin` - (Optional, String) Filter by discovery engine source. Valid values are `ML`, `LabelDiscovery`, `SessionIdentifier`.
- `state` - (Optional, String) Filter results by discovery state. Valid values are `review`, `saved`, `ignored`.
- `page` - (Optional, Integer) Page number of paginated results.
- `per_page` - (Optional, Integer) Maximum number of results per page.

## Attributes reference

In addition to all argument reference list, you can access the following attribute references after your data source is created.

- `id` - (String) The record ID. It is a combination of `<domain_id>:<cis_id>` attributes concatenated with `:`.
- `total_count` - (Integer) Total number of matching operations.
- `operations` - (List) List of discovered API Gateway operations.
  - `id` - (String) UUID of the discovered operation.
  - `endpoint` - (String) Endpoint path.
  - `host` - (String) Host.
  - `method` - (String) HTTP method.
  - `state` - (String) Discovery state (`review`, `saved`, or `ignored`).
  - `origin` - (List of String) Discovery engine origins.
  - `last_updated` - (String) Timestamp of last update in RFC3339 format.
