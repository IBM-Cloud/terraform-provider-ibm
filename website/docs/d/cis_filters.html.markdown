---
subcategory: "Internet services"
layout: "ibm"
page_title: "IBM: ibm_cis_filters"
description: |-
  Get information on an IBM Cloud Internet Services filters.
---

# ibm_cis_filters

Retrieve information about an IBM Cloud Internet Services filters data sources. For more information, see [IBM Cloud Internet Services](https://cloud.ibm.com/docs/cis?topic=cis-about-ibm-cloud-internet-services-cis).

## Example usage

```terraform
data "ibm_cis_filters" "test" {
  cis_id    = ibm_cis.instance.id
  domain_id = ibm_cis_domain.example.id
}
```

## Argument reference
Review the argument references that you can specify for your data source.

- `cis_id` - (Required, String) The ID of the CIS service instance.
- `domain_id` - (Required, String) The ID of the domain.

## Attributes reference
In addition to all argument reference list, you can access the following attribute references after your data source is created.

- `id` - (String) The Filter ID. It is a combination of <`filter_id`>,<`domain_id`>,<`cis_id`> attributes concatenated with ":"
- `cis_filters_list` - (List)
   - `expression` - (String) The expression of filter.
   - `paused` - (Boolean). Whether this filter is currently disabled.
   - `description` - (String) The information about this filter to help identify the purpose of it.
   - `filter_id` - (String) The filter ID.

