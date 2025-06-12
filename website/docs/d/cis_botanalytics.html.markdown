---
subcategory: "Internet services"
layout: "ibm"
page_title: "IBM: ibm_cis_bot_analytics"
description: |-
  Get information on an IBM Cloud Internet Services Bot Analytics APIs.
---

# ibm_cis_bot_analytics

Retrieve information about an IBM Cloud Internet Services Bot Analytics data sources for a zone, on the basis of type. 3 types of bot analytics are Score Source, TimeSeries and Top Attributes. For more information, see [IBM Cloud Internet Services Bot Management](https://cloud.ibm.com/docs/cis?topic=cis-about-bot-mgmt).
a
## Example usage

```terraform
data "ibm_cis_bot_analytics" "tests" {
    cis_id = data.ibm_cis.cis.id
    domain_id = data.ibm_cis_domain.cis_domain.domain_id
    since  = "2023-06-12T00:00:00Z"
    until  = "2023-06-13T00:00:00Z"
    type   = "score_source"

}
```

## Argument reference
Review the argument references that you can specify for your data source.

- `cis_id` - (Required, String) The ID of the CIS service instance.
- `domain_id` - (Required, String) The Domain of the CIS service instance.
- `type`   - (Required, String) The type of bot analytics score - `score_source`, `timeseries` and `top_ns`.
- `since`  - (Required, String) The time from which the analytics data is requested. 
- `until`  - (Required, String) The time till which the analytics data is requested.

## Attributes reference
In addition to all argument reference list, you can access the following attribute references after your data source is created.

- `id` - (String) The ID of the CIS service instance.
- `domain_id` - (String) The Domain of the CIS service instance.
- `type`   - (String) The type of bot analytics score - `score_source`, `timeseries` and `top_ns`.
- `result` - (String) The bot analytics data as per the requested type.
